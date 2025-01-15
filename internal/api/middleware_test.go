package api

import (
	"fmt"
	"github.com/dtslubbersen/go-quiz/internal/store"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"testing"
	"time"
)

const authorizationHeaderKey = "Authorization"

func setAuthorization(t *testing.T, request *http.Request, token string) {
	t.Helper()
	authorizationHeaderValue := fmt.Sprintf("Bearer %s", token)
	request.Header.Set(authorizationHeaderKey, authorizationHeaderValue)
}

func newTestClaims(t *testing.T, user *store.User) jwt.Claims {
	t.Helper()
	claims := jwt.MapClaims{
		"aud": "test-aud",
		"iss": "test-aud",
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour).Unix(),
	}

	return claims
}
