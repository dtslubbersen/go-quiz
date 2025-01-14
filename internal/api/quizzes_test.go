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
		assert           func(container *MockContainer)
		setAuthorization func(t *testing.T, request *http.Request, authenticator auth.Authenticator)
		validateResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			quizId: quiz.Id,
			assert: func(container *MockContainer) {
				container.MockQuizStore.EXPECT().GetById(gomock.Eq(quiz.Id)).Times(1).Return(&quiz, nil)
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

			container := newMockContainer(ctrl)
			api := newTestApplication(t, container.MockStorage)
			recorder := httptest.NewRecorder()
			test.assert(container)

			url := fmt.Sprintf("/quizes/%d", test.quizId)
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
	quizzes := make([]store.Quiz, quizCount)

	for i := 0; i < quizCount; i++ {
		quizzes[i] = newTestQuiz(t)
	}

	tests := []struct {
		name             string
		assert           func(container *MockContainer)
		setAuthorization func(t *testing.T, request *http.Request, authenticator auth.Authenticator)
		validateResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			assert: func(container *MockContainer) {
				container.MockStorage.EXPECT().Quizzes().Times(1).Return(quizzes, nil)

				//container.MockQuizStore.EXPECT().GetAll().Times(1).Return(&quizzes, nil)
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

			container := newMockContainer(ctrl)
			api := newTestApplication(t, container.MockStorage)
			recorder := httptest.NewRecorder()
			test.assert(container)

			url := "/quizzes"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			test.setAuthorization(t, request, api.authenticator)
			api.router.ServeHTTP(recorder, request)
			test.validateResponse(t, recorder)
		})
	}
}

func requireQuizResponseBodyMatch(t *testing.T, body *bytes.Buffer, expectedQuiz store.Quiz) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var response Response
	err = json.Unmarshal(data, &response)
	require.NoError(t, err)
	require.Equal(t, expectedQuiz, response.Data)
}

func requireQuizzesResponseBodyMatch(t *testing.T, body *bytes.Buffer, expectedQuizzes []store.Quiz) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var response Response
	err = json.Unmarshal(data, &response)
	require.NoError(t, err)
	require.Equal(t, expectedQuizzes, response.Data)
}

func newTestQuiz(t *testing.T) store.Quiz {
	quiz := store.Quiz{
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

type MockContainer struct {
	MockStorage         *mockStore.MockStorage
	MockQuizStore       *mockStore.MockQuizStore
	MockQuestionStore   *mockStore.MockQuestionStore
	MockResultStore     *mockStore.MockResultStore
	MockUserStore       *mockStore.MockUserStore
	MockUserAnswerStore *mockStore.MockUserAnswerStore
}

func newMockContainer(ctrl *gomock.Controller) *MockContainer {
	mockStorage := mockStore.NewMockStorage(ctrl)
	mockQuizStore := mockStore.NewMockQuizStore(ctrl)
	mockQuestionStore := mockStore.NewMockQuestionStore(ctrl)
	mockResultStore := mockStore.NewMockResultStore(ctrl)
	mockUserStore := mockStore.NewMockUserStore(ctrl)
	mockUserAnswerStore := mockStore.NewMockUserAnswerStore(ctrl)

	mockStorage.EXPECT().Quizzes().Return(mockQuizStore).AnyTimes()
	mockStorage.EXPECT().Questions().Return(mockQuestionStore).AnyTimes()
	mockStorage.EXPECT().Results().Return(mockResultStore).AnyTimes()
	mockStorage.EXPECT().Users().Return(mockUserStore).AnyTimes()
	mockStorage.EXPECT().UserAnswers().Return(mockUserAnswerStore).AnyTimes()

	return &MockContainer{
		MockStorage:         mockStorage,
		MockQuizStore:       mockQuizStore,
		MockQuestionStore:   mockQuestionStore,
		MockResultStore:     mockResultStore,
		MockUserStore:       mockUserStore,
		MockUserAnswerStore: mockUserAnswerStore,
	}
}
