package api

import (
	"fmt"
	"net/http"
	"testing"
)

const authorizationHeaderKey = "Authorization"

func setAuthorization(t *testing.T, request *http.Request, token string) {
	authorizationHeaderValue := fmt.Sprintf("Bearer %s", token)
	request.Header.Set(authorizationHeaderKey, authorizationHeaderValue)
}

func TestJwtTokenMiddleware(t *testing.T) {

}
