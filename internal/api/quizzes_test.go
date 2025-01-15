package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dtslubbersen/go-quiz/internal/auth"
	"github.com/dtslubbersen/go-quiz/internal/store"
	mockStore "github.com/dtslubbersen/go-quiz/internal/store/mock"
	"github.com/dtslubbersen/go-quiz/pkg/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetQuizzesApi(t *testing.T) {
	quizCount := 10
	quizzes := make([]*store.Quiz, quizCount)

	for i := 0; i < quizCount; i++ {
		quizzes[i] = newTestQuiz(t, false)
	}

	user, _ := newTestUser(t)
	claims := newTestClaims(t, user)

	tests := []struct {
		name             string
		setupMocks       func(storage *mockStore.MockStorage)
		setAuthorization func(t *testing.T, request *http.Request, authenticator auth.Authenticator)
		validateResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "ok",
			setupMocks: func(storage *mockStore.MockStorage) {
				storage.EXPECT().GetUserById(gomock.Eq(user.Id)).Times(1).Return(user, nil)
				storage.EXPECT().ListQuizzes().Times(1).Return(quizzes, nil)
			},
			setAuthorization: func(t *testing.T, request *http.Request, authenticator auth.Authenticator) {
				token, err := authenticator.GenerateToken(claims)
				require.NoError(t, err)
				setAuthorization(t, request, token)
			},
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireQuizzesResponseBodyMatch(t, recorder.Body, quizzes)
			},
		},
		{
			name:             "unauthorized",
			setupMocks:       func(storage *mockStore.MockStorage) {},
			setAuthorization: func(t *testing.T, request *http.Request, authenticator auth.Authenticator) {},
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
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
			test.setupMocks(storage)

			url := "/api/v1/quizzes"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			test.setAuthorization(t, request, api.authenticator)
			api.router.ServeHTTP(recorder, request)
			test.validateResponse(t, recorder)
		})
	}
}

func requireQuizzesResponseBodyMatch(t *testing.T, body *bytes.Buffer, expectedQuizzes []*store.Quiz) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	type Response struct {
		Data []*store.Quiz `json:"data"`
	}
	var response Response
	err = json.Unmarshal(data, &response)

	require.NoError(t, err)
	require.Equal(t, expectedQuizzes, response.Data)
}

func TestGetQuizByIdApi(t *testing.T) {
	quiz := newTestQuiz(t, true)
	user, _ := newTestUser(t)
	claims := newTestClaims(t, user)

	tests := []struct {
		name             string
		quizId           store.QuizId
		setupMocks       func(storage *mockStore.MockStorage)
		setAuthorization func(t *testing.T, request *http.Request, authenticator auth.Authenticator)
		validateResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "ok",
			quizId: quiz.Id,
			setupMocks: func(storage *mockStore.MockStorage) {
				storage.EXPECT().GetUserById(gomock.Eq(user.Id)).Times(1).Return(user, nil)
				storage.EXPECT().GetQuizById(gomock.Eq(quiz.Id)).Times(1).Return(quiz, nil)
				storage.EXPECT().ListQuestionsByQuizId(gomock.Eq(quiz.Id)).Times(1).Return(quiz.Questions, nil)
			},
			setAuthorization: func(t *testing.T, request *http.Request, authenticator auth.Authenticator) {
				token, err := authenticator.GenerateToken(claims)
				require.NoError(t, err)
				setAuthorization(t, request, token)
			},
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireQuizResponseBodyMatch(t, recorder.Body, quiz)
			},
		},
		{
			name:             "unauthorized",
			quizId:           quiz.Id,
			setupMocks:       func(storage *mockStore.MockStorage) {},
			setAuthorization: func(t *testing.T, request *http.Request, authenticator auth.Authenticator) {},
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:   "quiz not found",
			quizId: quiz.Id,
			setupMocks: func(storage *mockStore.MockStorage) {
				storage.EXPECT().GetUserById(gomock.Eq(user.Id)).Times(1).Return(user, nil)
				storage.EXPECT().GetQuizById(gomock.Eq(quiz.Id)).Times(1).Return(nil, store.NotFoundError)
			},
			setAuthorization: func(t *testing.T, request *http.Request, authenticator auth.Authenticator) {
				token, err := authenticator.GenerateToken(claims)
				require.NoError(t, err)
				setAuthorization(t, request, token)
			},
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
				requireErrorResponseBodyMatch(t, recorder.Body, QuizNotFoundError)
			},
		},
		{
			name:   "invalid quiz id",
			quizId: 0,
			setupMocks: func(storage *mockStore.MockStorage) {
				storage.EXPECT().GetUserById(gomock.Eq(user.Id)).Times(1).Return(user, nil)
			},
			setAuthorization: func(t *testing.T, request *http.Request, authenticator auth.Authenticator) {
				token, err := authenticator.GenerateToken(claims)
				require.NoError(t, err)
				setAuthorization(t, request, token)
			},
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
				requireErrorResponseBodyMatch(t, recorder.Body, InvalidQuizIdError("0"))
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
			test.setupMocks(storage)

			url := fmt.Sprintf("/api/v1/quizzes/%d", test.quizId)
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

	type Response struct {
		Data *store.Quiz `json:"data"`
	}
	var response Response
	err = json.Unmarshal(data, &response)

	require.NoError(t, err)
	require.Equal(t, expectedQuiz, response.Data)
}

func TestPostSubmitAnswersApi(t *testing.T) {
	quiz := newTestQuiz(t, true)
	user, _ := newTestUser(t)
	claims := newTestClaims(t, user)

	tests := []struct {
		name             string
		quizId           store.QuizId
		preparePayload   func(t *testing.T) map[string]interface{}
		setupMocks       func(storage *mockStore.MockStorage)
		setAuthorization func(t *testing.T, request *http.Request, authenticator auth.Authenticator)
		validateResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "ok",
			quizId: quiz.Id,
			preparePayload: func(t *testing.T) map[string]interface{} {
				return newTestSubmitAnswersPayload(quiz, len(quiz.Questions))
			},
			setupMocks: func(storage *mockStore.MockStorage) {
				storage.EXPECT().GetUserById(gomock.Eq(user.Id)).Times(1).Return(user, nil)
				storage.EXPECT().GetQuizById(gomock.Eq(quiz.Id)).Times(1).Return(quiz, nil)
				storage.EXPECT().ListQuestionsByQuizId(gomock.Eq(quiz.Id)).Times(1).Return(quiz.Questions, nil)
				storage.EXPECT().GetResultByQuizAndUserId(gomock.Eq(quiz.Id), gomock.Eq(user.Id)).Times(1).Return(nil, store.NotFoundError)
				storage.EXPECT().AddUserAnswer(gomock.Any()).Times(len(quiz.Questions)).Return(nil)
				storage.EXPECT().UpdateQuiz(gomock.Any()).Times(1).Return(nil)
				storage.EXPECT().AddResult(gomock.Any()).Times(1).Return(&store.Result{}, nil)
			},
			setAuthorization: func(t *testing.T, request *http.Request, authenticator auth.Authenticator) {
				token, err := authenticator.GenerateToken(claims)
				require.NoError(t, err)
				setAuthorization(t, request, token)
			},
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:             "unauthorized",
			quizId:           quiz.Id,
			preparePayload:   func(t *testing.T) map[string]interface{} { return map[string]interface{}(nil) },
			setupMocks:       func(storage *mockStore.MockStorage) {},
			setAuthorization: func(t *testing.T, request *http.Request, authenticator auth.Authenticator) {},
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:           "bad request - no payload",
			quizId:         quiz.Id,
			preparePayload: func(t *testing.T) map[string]interface{} { return map[string]interface{}(nil) },
			setupMocks: func(storage *mockStore.MockStorage) {
				storage.EXPECT().GetUserById(gomock.Eq(user.Id)).Times(1).Return(user, nil)
				storage.EXPECT().GetQuizById(gomock.Eq(quiz.Id)).Times(1).Return(quiz, nil)
			},
			setAuthorization: func(t *testing.T, request *http.Request, authenticator auth.Authenticator) {
				token, err := authenticator.GenerateToken(claims)
				require.NoError(t, err)
				setAuthorization(t, request, token)
			},
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:   "bad request - malformed payload",
			quizId: quiz.Id,
			preparePayload: func(t *testing.T) map[string]interface{} {
				return map[string]interface{}{
					"answers": map[string]interface{}{
						"i_dont_exist_id": 1,
					},
				}
			},
			setupMocks: func(storage *mockStore.MockStorage) {
				storage.EXPECT().GetUserById(gomock.Eq(user.Id)).Times(1).Return(user, nil)
				storage.EXPECT().GetQuizById(gomock.Eq(quiz.Id)).Times(1).Return(quiz, nil)
			},
			setAuthorization: func(t *testing.T, request *http.Request, authenticator auth.Authenticator) {
				token, err := authenticator.GenerateToken(claims)
				require.NoError(t, err)
				setAuthorization(t, request, token)
			},
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:   "bad request - misaligned question and answer count",
			quizId: quiz.Id,
			preparePayload: func(t *testing.T) map[string]interface{} {
				return newTestSubmitAnswersPayload(quiz, len(quiz.Questions)-1)
			},
			setupMocks: func(storage *mockStore.MockStorage) {
				storage.EXPECT().GetUserById(gomock.Eq(user.Id)).Times(1).Return(user, nil)
				storage.EXPECT().GetQuizById(gomock.Eq(quiz.Id)).Times(1).Return(quiz, nil)
				storage.EXPECT().ListQuestionsByQuizId(gomock.Eq(quiz.Id)).Times(1).Return(quiz.Questions, nil)
			},
			setAuthorization: func(t *testing.T, request *http.Request, authenticator auth.Authenticator) {
				token, err := authenticator.GenerateToken(claims)
				require.NoError(t, err)
				setAuthorization(t, request, token)
			},
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
				requireErrorResponseBodyMatch(t, recorder.Body, SubmitAnswersInvalidAnswerCountError)
			},
		},
		{
			name:   "bad request - quiz already answered",
			quizId: quiz.Id,
			preparePayload: func(t *testing.T) map[string]interface{} {
				return newTestSubmitAnswersPayload(quiz, len(quiz.Questions))
			},
			setupMocks: func(storage *mockStore.MockStorage) {
				storage.EXPECT().GetUserById(gomock.Eq(user.Id)).Times(1).Return(user, nil)
				storage.EXPECT().GetQuizById(gomock.Eq(quiz.Id)).Times(1).Return(quiz, nil)
				storage.EXPECT().ListQuestionsByQuizId(gomock.Eq(quiz.Id)).Times(1).Return(quiz.Questions, nil)
				storage.EXPECT().GetResultByQuizAndUserId(gomock.Eq(quiz.Id), gomock.Eq(user.Id)).Times(1).Return(&store.Result{}, nil)
			},
			setAuthorization: func(t *testing.T, request *http.Request, authenticator auth.Authenticator) {
				token, err := authenticator.GenerateToken(claims)
				require.NoError(t, err)
				setAuthorization(t, request, token)
			},
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
				requireErrorResponseBodyMatch(t, recorder.Body, QuizAlreadyAnsweredError)
			},
		},
		{
			name:   "invalid quiz id",
			quizId: 0,
			preparePayload: func(t *testing.T) map[string]interface{} {
				return newTestSubmitAnswersPayload(quiz, len(quiz.Questions))
			},
			setupMocks: func(storage *mockStore.MockStorage) {
				storage.EXPECT().GetUserById(gomock.Eq(user.Id)).Times(1).Return(user, nil)
			},
			setAuthorization: func(t *testing.T, request *http.Request, authenticator auth.Authenticator) {
				token, err := authenticator.GenerateToken(claims)
				require.NoError(t, err)
				setAuthorization(t, request, token)
			},
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
				requireErrorResponseBodyMatch(t, recorder.Body, InvalidQuizIdError("0"))
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
			test.setupMocks(storage)

			data, err := json.Marshal(test.preparePayload(t))
			require.NoError(t, err)

			url := fmt.Sprintf("/api/v1/quizzes/%d/submit", test.quizId)
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			test.setAuthorization(t, request, api.authenticator)
			api.router.ServeHTTP(recorder, request)
			test.validateResponse(t, recorder)
		})
	}
}

func newTestSubmitAnswersPayload(quiz *store.Quiz, questionCount int) map[string]interface{} {
	answers := make([]QuestionAnswerPayload, 0)

	for i := 0; i < questionCount; i++ {
		question := quiz.Questions[i]
		answers = append(answers, QuestionAnswerPayload{
			QuestionId:  int64(question.Id),
			AnswerIndex: 1,
		})
	}

	return map[string]interface{}{
		"answers": answers,
	}
}

func TestGetQuizResultsApi(t *testing.T) {
	quiz := newTestQuiz(t, true)
	user, _ := newTestUser(t)
	claims := newTestClaims(t, user)
	result := newTestResult(t, quiz)

	tests := []struct {
		name             string
		quizId           store.QuizId
		setupMocks       func(storage *mockStore.MockStorage)
		setAuthorization func(t *testing.T, request *http.Request, authenticator auth.Authenticator)
		validateResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "ok",
			quizId: quiz.Id,
			setupMocks: func(storage *mockStore.MockStorage) {
				storage.EXPECT().GetUserById(gomock.Eq(user.Id)).Times(1).Return(user, nil)
				storage.EXPECT().GetQuizById(gomock.Eq(quiz.Id)).Times(1).Return(quiz, nil)
				storage.EXPECT().GetResultByQuizAndUserId(gomock.Eq(quiz.Id), gomock.Eq(user.Id)).Times(1).Return(result, nil)
			},
			setAuthorization: func(t *testing.T, request *http.Request, authenticator auth.Authenticator) {
				token, err := authenticator.GenerateToken(claims)
				require.NoError(t, err)
				setAuthorization(t, request, token)
			},
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireResultResponseBodyMatch(t, recorder.Body, result)
			},
		},
		{
			name:             "unauthorized",
			setupMocks:       func(storage *mockStore.MockStorage) {},
			setAuthorization: func(t *testing.T, request *http.Request, authenticator auth.Authenticator) {},
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:   "result not found",
			quizId: quiz.Id,
			setupMocks: func(storage *mockStore.MockStorage) {
				storage.EXPECT().GetUserById(gomock.Eq(user.Id)).Times(1).Return(user, nil)
				storage.EXPECT().GetQuizById(gomock.Eq(quiz.Id)).Times(1).Return(quiz, nil)
				storage.EXPECT().GetResultByQuizAndUserId(gomock.Eq(quiz.Id), gomock.Eq(user.Id)).Times(1).Return(nil, store.NotFoundError)
			},
			setAuthorization: func(t *testing.T, request *http.Request, authenticator auth.Authenticator) {
				token, err := authenticator.GenerateToken(claims)
				require.NoError(t, err)
				setAuthorization(t, request, token)
			},
			validateResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
				requireErrorResponseBodyMatch(t, recorder.Body, ResultsNotFoundError)
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
			test.setupMocks(storage)

			url := fmt.Sprintf("/api/v1/quizzes/%d/results", test.quizId)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			test.setAuthorization(t, request, api.authenticator)
			api.router.ServeHTTP(recorder, request)
			test.validateResponse(t, recorder)
		})
	}
}

func requireResultResponseBodyMatch(t *testing.T, body *bytes.Buffer, expectedResult *store.Result) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	type Response struct {
		Data *store.Result `json:"data"`
	}
	var response Response
	err = json.Unmarshal(data, &response)

	require.NoError(t, err)
	require.Equal(t, expectedResult, response.Data)
}

func newTestResult(t *testing.T, quiz *store.Quiz) *store.Result {
	return &store.Result{}
}

func newTestQuiz(t *testing.T, includeQuestions bool) *store.Quiz {
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

	if includeQuestions {
		var questions []*store.Question

		for i := 0; i < 10; i++ {
			question := &store.Question{
				Id:                 store.QuestionId(i + 1),
				QuizId:             quiz.Id,
				Value:              util.RandomString(8),
				Answers:            nil,
				CorrectAnswerIndex: 1,
			}

			var answers []store.Answer

			for i := 0; i < 3; i++ {
				answers = append(answers, store.Answer(util.RandomString(8)))
			}

			question.Answers = answers
			questions = append(questions, question)
		}

		quiz.Questions = questions
	}

	return quiz
}

func newTestUser(t *testing.T) (*store.User, string) {
	password := util.RandomString(8)
	user := &store.User{
		Id:        store.UserId(util.RandomInt(1, 128)),
		FirstName: util.RandomString(8),
		LastName:  util.RandomString(8),
		Email:     util.RandomEmail(),
	}

	err := user.Password.Set(password)
	require.NoError(t, err)

	return user, password
}

func newTestClaims(t *testing.T, user *store.User) jwt.Claims {
	claims := jwt.MapClaims{
		"aud": "test-aud",
		"iss": "test-aud",
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour).Unix(),
	}

	return claims
}

func requireErrorResponseBodyMatch(t *testing.T, body *bytes.Buffer, expectedError error) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	type Response struct {
		Error string `json:"error"`
	}
	var response Response
	err = json.Unmarshal(data, &response)
	require.NoError(t, err)
	require.Equal(t, expectedError.Error(), response.Error)
}
