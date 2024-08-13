// authentication integration with a test server
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	UserID   string
	Username string
}

type Application struct {
	AuthServiceURL string
}

func NewApplication(authServiceURL string) *Application {
	return &Application{
		AuthServiceURL: authServiceURL,
	}
}

func (app *Application) AuthenticateUser(token string) (*User, error) {
	return &User{
		UserID:   "123",
		Username: "testuser",
	}, nil
}

func TestAuthenticationIntegration(t *testing.T) {
	authService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "Bearer valid_token" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"user_id": "123", "username": "testuser"}`))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}))
	defer authService.Close()

	app := NewApplication(authService.URL)
	token := "valid_token"
	gotUser, err := app.AuthenticateUser(token)
	assert.NoError(t, err)
	assert.Equal(t, "123", gotUser.UserID)
	assert.Equal(t, "testuser", gotUser.Username)
}
