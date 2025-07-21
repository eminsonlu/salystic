package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/eminsonlu/salystic/internal/model"
	"strconv"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/linkedin"
)

type LinkedInOAuth struct {
	config     *oauth2.Config
	hmacSecret string
}

func NewLinkedInOAuth(clientID, clientSecret, redirectURL, hmacSecret string) *LinkedInOAuth {
	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint:     linkedin.Endpoint,
	}

	return &LinkedInOAuth{
		config:     config,
		hmacSecret: hmacSecret,
	}
}

func (l *LinkedInOAuth) GetAuthURL(state string) string {
	return l.config.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (l *LinkedInOAuth) ExchangeCodeForToken(ctx context.Context, code string) (*oauth2.Token, error) {
	return l.config.Exchange(ctx, code)
}

func (l *LinkedInOAuth) GetLinkedInProfile(ctx context.Context, token *oauth2.Token) (*model.LinkedInProfile, error) {
	client := l.config.Client(ctx, token)

	resp, err := client.Get("https://api.linkedin.com/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to get LinkedIn profile: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("LinkedIn API returned status %d", resp.StatusCode)
	}

	var profileResponse struct {
		Sub           string      `json:"sub"`
		Name          string      `json:"name"`
		GivenName     string      `json:"given_name"`
		FamilyName    string      `json:"family_name"`
		Picture       string      `json:"picture"`
		Locale        interface{} `json:"locale"` // string or object
		Email         string      `json:"email"`
		EmailVerified bool        `json:"email_verified"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&profileResponse); err != nil {
		return nil, fmt.Errorf("failed to decode LinkedIn profile: %w", err)
	}

	var localeStr string
	switch v := profileResponse.Locale.(type) {
	case string:
		localeStr = v
	case map[string]interface{}:
		if lang, ok := v["language"].(string); ok {
			localeStr = lang
		} else {
			localeStr = "en-US"
		}
	default:
		localeStr = "en-US"
	}

	return &model.LinkedInProfile{
		Sub:           profileResponse.Sub,
		Name:          profileResponse.Name,
		GivenName:     profileResponse.GivenName,
		FamilyName:    profileResponse.FamilyName,
		Picture:       profileResponse.Picture,
		Locale:        localeStr,
		Email:         profileResponse.Email,
		EmailVerified: profileResponse.EmailVerified,
	}, nil
}

func (l *LinkedInOAuth) PseudonymizeLinkedInID(linkedInID string) string {
	h := hmac.New(sha256.New, []byte(l.hmacSecret))
	h.Write([]byte(linkedInID))
	h.Write([]byte(strconv.FormatInt(time.Now().Unix()/86400, 10)))
	return hex.EncodeToString(h.Sum(nil))
}
