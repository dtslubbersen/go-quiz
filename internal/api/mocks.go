package api

import (
	"github.com/dtslubbersen/go-quiz/internal/auth"
	"github.com/dtslubbersen/go-quiz/internal/store"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewMockApplication(t *testing.T, cfg apiCfg) *Application {
	t.Helper()

	authenticator := &auth.MockAuthenticator{}
	logger := zap.NewNop().Sugar()
	storage := store.NewMockStorage()

	return &Application{
		authenticator: authenticator,
		cfg:           cfg,
		logger:        logger,
		storage:       storage,
	}
}

func executeRequest(req *http.Request, router http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}

func validateStatusCode(t *testing.T, expectedStatus, status int) {
	if expectedStatus != status {
		t.Errorf("Expected status code %d. Received %d instead", expectedStatus, status)
	}
}
