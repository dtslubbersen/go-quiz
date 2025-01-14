package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dtslubbersen/go-quiz/internal/auth"
	"github.com/dtslubbersen/go-quiz/internal/store"
	mockStore "github.com/dtslubbersen/go-quiz/internal/store/mock"
	"github.com/dtslubbersen/go-quiz/pkg/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetQuizByIdApi(t *testing.T) {
	quiz := newTestQuiz(t)

	tests := []struct {
		name             string
		quizId           store.QuizId
		assert           func(storage *mockStore.MockStorage)
		setAuthorization func(t *testing.T, request *http.Request, authenticator auth.Authenticator)
		validateResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			quizId: quiz.Id,
			assert: func(storage *mockStore.MockStorage) {
				storage.EXPECT().GetQuizById(gomock.Eq(quiz.Id)).Times(1).Return(&quiz, nil)
			},
			setAuthorization: func(t *testing.T, request *http.Request, authenticator auth.Authenticator) {
				token, err := authenticator.GenerateToken(nil)
				require.NoError(t, err)
				setAuthorization(t, request, token)
			},
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireQuizResponseBodyMatch(t, recorder.Body, quiz)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			storage := mockStore.NewMockStorage(ctrl)
			api := newTestApplication(t, storage)
			recorder := httptest.NewRecorder()
			test.assert(storage)

			url := fmt.Sprintf("/api/v1/quizzes/%d", test.quizId)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			test.setAuthorization(t, request, api.authenticator)
			api.router.ServeHTTP(recorder, request)
			test.validateResponse(t, recorder)
		})
	}
}

func TestGetQuizzesApi(t *testing.T) {
	quizCount := 10
	quizzes := make([]*store.Quiz, quizCount)

	for i := 0; i < quizCount; i++ {
		quizzes[i] = newTestQuiz(t)
	}

	tests := []struct {
		name             string
		assert           func(storage *mockStore.MockStorage)
		setAuthorization func(t *testing.T, request *http.Request, authenticator auth.Authenticator)
		validateResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			assert: func(storage *mockStore.MockStorage) {
				storage.EXPECT().ListQuizzes().Times(1).Return(quizzes, nil)
			},
			setAuthorization: func(t *testing.T, request *http.Request, authenticator auth.Authenticator) {
				token, err := authenticator.GenerateToken(nil)
				require.NoError(t, err)
				setAuthorization(t, request, token)
			},
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireQuizzesResponseBodyMatch(t, recorder.Body, quizzes)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			storage := mockStore.NewMockStorage(ctrl)
			api := newTestApplication(t, storage)
			recorder := httptest.NewRecorder()
			test.assert(storage)

			url := "/api/v1/quizzes"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			test.setAuthorization(t, request, api.authenticator)
			api.router.ServeHTTP(recorder, request)
			test.validateResponse(t, recorder)
		})
	}
}

func requireQuizResponseBodyMatch(t *testing.T, body *bytes.Buffer, expectedQuiz *store.Quiz) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var response Response
	err = json.Unmarshal(data, &response)
	require.NoError(t, err)
	require.Equal(t, expectedQuiz, response.Data)
}

func requireQuizzesResponseBodyMatch(t *testing.T, body *bytes.Buffer, expectedQuizzes []*store.Quiz) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var response Response
	err = json.Unmarshal(data, &response)
	require.NoError(t, err)
	require.Equal(t, expectedQuizzes, response.Data)
}

func newTestQuiz(t *testing.T) *store.Quiz {
	quiz := &store.Quiz{
		Id:          store.QuizId(util.RandomInt(1, 128)),
		Title:       util.RandomString(8),
		Description: util.RandomString(32),
		Questions:   nil,
		Performance: store.Performance{
			UsersTakenCount:     0,
			CorrectAnswersCount: nil,
		},
	}

	return quiz
}

func newTestUser(t *testing.T) (store.User, string) {
	password := util.RandomString(8)
	user := store.User{
		Id:        store.UserId(util.RandomInt(1, 128)),
		FirstName: util.RandomString(8),
		LastName:  util.RandomString(8),
		Email:     util.RandomEmail(),
	}

	err := user.Password.Set(password)
	require.NoError(t, err)

	return user, password
}
