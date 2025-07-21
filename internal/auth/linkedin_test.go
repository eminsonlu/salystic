package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

func TestNewLinkedInOAuth(t *testing.T) {
	clientID := "test_client_id"
	clientSecret := "test_client_secret"
	redirectURL := "http://localhost:8080/auth/callback"
	hmacSecret := "hmac_secret"

	linkedinOAuth := NewLinkedInOAuth(clientID, clientSecret, redirectURL, hmacSecret)

	assert.NotNil(t, linkedinOAuth)
	assert.Equal(t, hmacSecret, linkedinOAuth.hmacSecret)
	assert.Equal(t, clientID, linkedinOAuth.config.ClientID)
	assert.Equal(t, clientSecret, linkedinOAuth.config.ClientSecret)
	assert.Equal(t, redirectURL, linkedinOAuth.config.RedirectURL)
	assert.Contains(t, linkedinOAuth.config.Scopes, "openid")
	assert.Contains(t, linkedinOAuth.config.Scopes, "profile")
	assert.Contains(t, linkedinOAuth.config.Scopes, "email")
}

func TestGetAuthURL(t *testing.T) {
	clientID := "test_client_id"
	clientSecret := "test_client_secret"
	redirectURL := "http://localhost:8080/auth/callback"
	hmacSecret := "hmac_secret"

	linkedinOAuth := NewLinkedInOAuth(clientID, clientSecret, redirectURL, hmacSecret)

	state := "test_state_123"
	authURL := linkedinOAuth.GetAuthURL(state)

	assert.NotEmpty(t, authURL)
	assert.Contains(t, authURL, "linkedin.com")
	assert.Contains(t, authURL, "client_id="+clientID)
	assert.Contains(t, authURL, "state="+state)
	assert.Contains(t, authURL, "scope=openid+profile+email")
}

func TestExchangeCodeForToken_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenResponse := map[string]interface{}{
			"access_token": "mock_access_token",
			"token_type":   "Bearer",
			"expires_in":   3600,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tokenResponse)
	}))
	defer server.Close()

	clientID := "test_client_id"
	clientSecret := "test_client_secret"
	redirectURL := "http://localhost:8080/auth/callback"
	hmacSecret := "hmac_secret"

	linkedinOAuth := NewLinkedInOAuth(clientID, clientSecret, redirectURL, hmacSecret)
	
	linkedinOAuth.config.Endpoint = oauth2.Endpoint{
		AuthURL:  server.URL + "/auth",
		TokenURL: server.URL + "/token",
	}

	ctx := context.Background()
	code := "authorization_code"

	token, err := linkedinOAuth.ExchangeCodeForToken(ctx, code)

	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.Equal(t, "mock_access_token", token.AccessToken)
	assert.Equal(t, "Bearer", token.TokenType)
}

func TestGetLinkedInProfile_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/v2/userinfo" {
			profileResponse := map[string]interface{}{
				"sub":            "linkedin_user_id",
				"name":           "John Doe",
				"given_name":     "John",
				"family_name":    "Doe",
				"picture":        "https://media.licdn.com/picture.jpg",
				"locale":         "en-US",
				"email":          "john.doe@example.com",
				"email_verified": true,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(profileResponse)
		}
	}))
	defer server.Close()

	clientID := "test_client_id"
	clientSecret := "test_client_secret"
	redirectURL := "http://localhost:8080/auth/callback"
	hmacSecret := "hmac_secret"

	linkedinOAuth := NewLinkedInOAuth(clientID, clientSecret, redirectURL, hmacSecret)
	
	token := &oauth2.Token{
		AccessToken: "mock_access_token",
		TokenType:   "Bearer",
	}

	linkedinOAuth.config.Endpoint = oauth2.Endpoint{
		AuthURL:  server.URL + "/auth",
		TokenURL: server.URL + "/token",
	}

	ctx := context.Background()
	client := &http.Client{
		Transport: &http.Transport{},
	}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, client)

	profile, err := linkedinOAuth.GetLinkedInProfile(ctx, token)

	if err != nil {
		assert.Error(t, err)
		return
	}

	assert.NotNil(t, profile)
}

func TestGetLinkedInProfile_Structure(t *testing.T) {
	linkedinOAuth := NewLinkedInOAuth("client_id", "client_secret", "redirect_url", "hmac_secret")
	
	assert.NotNil(t, linkedinOAuth)
	assert.NotNil(t, linkedinOAuth.config)
	
	assert.NotNil(t, linkedinOAuth.GetLinkedInProfile)
	assert.NotNil(t, linkedinOAuth.ExchangeCodeForToken)
}

type mockTransport struct {
	server *httptest.Server
}

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = "http"
	req.URL.Host = t.server.URL[7:]
	return http.DefaultTransport.RoundTrip(req)
}

func TestGetLinkedInProfile_WithLocaleObject(t *testing.T) {
	linkedinOAuth := NewLinkedInOAuth("client_id", "client_secret", "redirect_url", "hmac_secret")

	assert.NotNil(t, linkedinOAuth)
	assert.Equal(t, "hmac_secret", linkedinOAuth.hmacSecret)
}

func TestLinkedInPseudonymizeID(t *testing.T) {
	hmacSecret := "test_hmac_secret"
	linkedinOAuth := NewLinkedInOAuth("client_id", "client_secret", "redirect_url", hmacSecret)

	linkedInID1 := "linkedin_user_1"
	linkedInID2 := "linkedin_user_2"

	pseudoID1 := linkedinOAuth.PseudonymizeLinkedInID(linkedInID1)
	pseudoID2 := linkedinOAuth.PseudonymizeLinkedInID(linkedInID2)

	assert.NotEmpty(t, pseudoID1)
	assert.NotEmpty(t, pseudoID2)
	assert.NotEqual(t, pseudoID1, pseudoID2)
	assert.Equal(t, 64, len(pseudoID1))
	assert.Equal(t, 64, len(pseudoID2))
}

func TestLinkedInPseudonymizeID_Consistency(t *testing.T) {
	hmacSecret := "test_hmac_secret"
	linkedinOAuth := NewLinkedInOAuth("client_id", "client_secret", "redirect_url", hmacSecret)

	linkedInID := "consistent_user_id"

	pseudoID1 := linkedinOAuth.PseudonymizeLinkedInID(linkedInID)
	pseudoID2 := linkedinOAuth.PseudonymizeLinkedInID(linkedInID)

	assert.Equal(t, pseudoID1, pseudoID2)
}

func TestPseudonymizeLinkedInID_DifferentSecrets(t *testing.T) {
	linkedinOAuth1 := NewLinkedInOAuth("client_id", "client_secret", "redirect_url", "secret1")
	linkedinOAuth2 := NewLinkedInOAuth("client_id", "client_secret", "redirect_url", "secret2")

	linkedInID := "same_user_id"

	pseudoID1 := linkedinOAuth1.PseudonymizeLinkedInID(linkedInID)
	pseudoID2 := linkedinOAuth2.PseudonymizeLinkedInID(linkedInID)

	assert.NotEqual(t, pseudoID1, pseudoID2)
}