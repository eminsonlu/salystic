package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"github.com/eminsonlu/salystic/internal/config"
	"github.com/eminsonlu/salystic/internal/service"
	"github.com/eminsonlu/salystic/pkg/responses"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService service.AuthService
	config      *config.Config
}

func NewAuthHandler(authService service.AuthService, cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		config:      cfg,
	}
}

func (h *AuthHandler) LinkedInLogin(c echo.Context) error {
	state, err := generateRandomState()
	if err != nil {
		return responses.InternalServerError(c, "Failed to generate state")
	}

	c.SetCookie(&http.Cookie{
		Name:     "oauth_state",
		Value:    state,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	authURL := h.authService.GetLinkedInAuthURL(state)

	return c.Redirect(http.StatusFound, authURL)
}

func (h *AuthHandler) LinkedInCallback(c echo.Context) error {
	code := c.QueryParam("code")
	state := c.QueryParam("state")
	errorParam := c.QueryParam("error")

	if errorParam != "" {
		errorDescription := c.QueryParam("error_description")
		redirectURL := fmt.Sprintf("%s?error=%s&error_description=%s",
			h.config.FrontendCallbackURL,
			url.QueryEscape(errorParam),
			url.QueryEscape(errorDescription))
		return c.Redirect(http.StatusFound, redirectURL)
	}

	if code == "" {
		redirectURL := fmt.Sprintf("%s?error=missing_code&error_description=%s",
			h.config.FrontendCallbackURL,
			url.QueryEscape("Authorization code is required"))
		return c.Redirect(http.StatusFound, redirectURL)
	}

	if state == "" {
		redirectURL := fmt.Sprintf("%s?error=missing_state&error_description=%s",
			h.config.FrontendCallbackURL,
			url.QueryEscape("State parameter is required"))
		return c.Redirect(http.StatusFound, redirectURL)
	}

	storedState := ""
	if cookie, err := c.Cookie("oauth_state"); err == nil {
		storedState = cookie.Value
	}

	if state != storedState {
		redirectURL := fmt.Sprintf("%s?error=invalid_state&error_description=%s",
			h.config.FrontendCallbackURL,
			url.QueryEscape("Invalid state parameter"))
		return c.Redirect(http.StatusFound, redirectURL)
	}

	c.SetCookie(&http.Cookie{
		Name:   "oauth_state",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	authResponse, err := h.authService.AuthenticateWithLinkedIn(c.Request().Context(), code)
	if err != nil {
		redirectURL := fmt.Sprintf("%s?error=auth_failed&error_description=%s",
			h.config.FrontendCallbackURL, url.QueryEscape(err.Error()))
		return c.Redirect(http.StatusFound, redirectURL)
	}

	redirectURL := fmt.Sprintf("%s?success=true&token=%s&expires_in=%d",
		h.config.FrontendCallbackURL,
		url.QueryEscape(authResponse.AccessToken),
		authResponse.ExpiresIn)

	return c.Redirect(http.StatusFound, redirectURL)
}

func (h *AuthHandler) Me(c echo.Context) error {
	userID := c.Get("user_id").(string)

	user, err := h.authService.GetUserByID(c.Request().Context(), userID)
	if err != nil {
		return responses.InternalServerError(c, "Failed to get user information")
	}

	return responses.Success(c, user)
}

func (h *AuthHandler) Logout(c echo.Context) error {
	userID := c.Get("user_id").(string)

	if err := h.authService.Logout(c.Request().Context(), userID); err != nil {
		return responses.InternalServerError(c, "Failed to logout")
	}

	return responses.SuccessWithMessage(c, "Successfully logged out", nil)
}

func generateRandomState() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
