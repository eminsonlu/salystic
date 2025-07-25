package middleware

import (
	"github.com/eminsonlu/salystic/internal/service"
	"github.com/eminsonlu/salystic/pkg/responses"
	"strings"

	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	authService service.AuthService
}

func NewAuthMiddleware(authService service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

func (m *AuthMiddleware) RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return responses.Unauthorized(c, "Authorization header required")
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return responses.Unauthorized(c, "Invalid authorization header format")
		}

		token := tokenParts[1]
		claims, err := m.authService.ValidateToken(token)
		if err != nil {
			return responses.Unauthorized(c, "Invalid or expired token")
		}

		c.Set("user_id", claims.UserID)
		c.Set("pseudonymized_id", claims.PseudonymizedID)
		c.Set("claims", claims)

		return next(c)
	}
}
