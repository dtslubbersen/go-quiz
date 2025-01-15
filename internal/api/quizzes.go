package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/dtslubbersen/go-quiz/internal/store"
	"github.com/go-chi/chi/v5"
	"math"
	"net/http"
	"strconv"
)

var (
	SubmitAnswersInvalidAnswerCountError = errors.New("amount of answers and question count do not match")
	QuizAlreadyAnsweredError             = errors.New("quiz already answered")
	ResultsNotFoundError                 = errors.New("results not found")
	InvalidQuizIdError                   = func(quizId string) error { return fmt.Errorf("invalid quiz id: %s", quizId) }
	QuizNotFoundError                    = errors.New("quiz not found")
)

type quizKey string

const quizCtxKey quizKey = "post"

// GetQuizzes godoc
//
//	@Summary		Retrieves all quizzes
//	@Description	Fetches a list of all quizzes from the in memory store
//	@Tags			quizzes
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Response{status_code=int,data=[]store.Quiz}
//	@Failure		400	{object}	Response{status_code=int,error=string}
//	@Failure		500	{object}	Response{status_code=int,error=string}
//	@Security		BearerAuth
//	@Router			/quizzes [get]
func (a *Application) getQuizzesHandler(w http.ResponseWriter, r *http.Request) {
	quizzes, err := a.storage.ListQuizzes()

	if err != nil {
		a.badRequest(w, r, err)
		return
	}

	a.dataResponse(w, r, http.StatusOK, quizzes)
}

// GetQuizById godoc
//
//	@Summary		Retrieves a quiz by ID
//	@Description	Fetches a specific quiz using its ID from the in-memory store
//	@Tags			quizzes
//	@Accept			json
//	@Produce		json
//	@Param			quizId	path		int	true	"Quiz ID"
//	@Success		200		{object}	Response{status_code=int,data=store.Quiz}
//	@Failure		400		{object}	Response{status_code=int,error=string}
//	@Failure		404		{object}	Response{status_code=int,error=string}
//	@Failure		500		{object}	Response{status_code=int,error=string}
//	@Security		BearerAuth
//	@Router			/quizzes/{quizId} [get]
func (a *Application) getQuizByIdHandler(w http.ResponseWriter, r *http.Request) {
	quiz := getQuizFromCtx(r)
	questions, err := a.storage.ListQuestionsByQuizId(quiz.Id)

	if err != nil {
		a.internalServerError(w, r, err)
	}

	quiz.Questions = questions
	a.dataResponse(w, r, http.StatusOK, quiz)
}

type SubmitQuizAnswersPayload struct {
	Answers []QuestionAnswerPayload `json:"answers" validate:"required"`
}

type QuestionAnswerPayload struct {
	QuestionId  int64 `json:"question_id" validate:"required"`
	AnswerIndex int   `json:"answer_index" validate:"required"`
}

// SubmitQuizAnswers godoc
//
//	@Summary		Submits answers for a quiz
//	@Description	Allows a user to submit answers for a given quiz
//	@Tags			quizzes
//	@Accept			json
//	@Produce		json
//	@Param			quizId	path		int															true	"Quiz ID"
//	@Param			payload body		SubmitQuizAnswersPayload{answers=[]QuestionAnswerPayload}	true	"User's answers"
//	@Success		200		{object}	Response{status_code=int}
//	@Failure		400		{object}	Response{status_code=int,error=string}
//	@Failure		500		{object}	Response{status_code=int,error=string}
//	@Security		BearerAuth
//	@Router			/quizzes/{quizId}/submit [post]
func (a *Application) postSubmitAnswersHandler(w http.ResponseWriter, r *http.Request) {
	var payload SubmitQuizAnswersPayload

	if err := readJson(w, r, &payload); err != nil {
		a.badRequest(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		a.badRequest(w, r, err)
		return
	}

	quiz, user := getQuizFromCtx(r), getUserFromCtx(r)
	questions, err := a.storage.ListQuestionsByQuizId(quiz.Id)

	if err != nil {
		a.internalServerError(w, r, err)
	}

	quiz.Questions = questions

	if len(quiz.Questions) != len(payload.Answers) {
		a.badRequest(w, r, SubmitAnswersInvalidAnswerCountError)
		return
	}

	if result, _ := a.storage.GetResultByQuizAndUserId(quiz.Id, user.Id); result != nil {
		a.badRequest(w, r, QuizAlreadyAnsweredError)
		return
	}

	answersMap := make(map[int64]int)

	for _, answer := range payload.Answers {
		answersMap[answer.QuestionId] = answer.AnswerIndex
	}

	correctAnswersCount := 0

	for _, question := range quiz.Questions {
		answerIndex := answersMap[int64(question.Id)]
		isCorrect := question.CorrectAnswerIndex == answerIndex

		if isCorrect {
			correctAnswersCount++
		}

		userAnswer := &store.UserAnswer{
			QuestionId:  question.Id,
			UserId:      user.Id,
			AnswerIndex: answerIndex,
			IsCorrect:   isCorrect,
		}

		if err := a.storage.AddUserAnswer(userAnswer); err != nil {
			a.badRequest(w, r, err)
			return
		}
	}

	quiz.Performance.UsersTakenCount++

	if err := a.storage.UpdateQuiz(quiz); err != nil {
		a.internalServerError(w, r, err)
		return
	}

	result := &store.Result{
		QuizId:              quiz.Id,
		QuestionCount:       len(quiz.Questions),
		UserId:              user.Id,
		CorrectAnswersCount: correctAnswersCount,
	}

	_, err = a.storage.AddResult(result)

	if err != nil {
		a.badRequest(w, r, err)
		return
	}

	a.dataResponse(w, r, http.StatusOK, nil)
}

func calculatePercentileRank(quiz *store.Quiz, correctAnswersCount int) float64 {
	totalUsers := quiz.Performance.UsersTakenCount
	if totalUsers == 0 {
		return 100
	}

	usersWithLessCorrectAnswers := 0
	for count, numUsers := range quiz.Performance.CorrectAnswersCount {
		if count < correctAnswersCount {
			usersWithLessCorrectAnswers += numUsers
		}
	}

	percentileRank := float64(usersWithLessCorrectAnswers) / float64(totalUsers) * 100
	return math.Round(percentileRank*100) / 100
}

// GetQuizResults godoc
//
//	@Summary		Retrieves quiz results for a user
//	@Description	Fetches the result of a quiz attempt by the current user.
//	@Tags			quizzes
//	@Accept			json
//	@Produce		json
//	@Param			quizId	path		int	true	"Quiz ID"
//	@Success		200		{object}	Response{status_code=int,data=store.Result}
//	@Failure		400		{object}	Response{status_code=int,error=string}
//	@Failure		404		{object}	Response{status_code=int,error=string}
//	@Failure		500		{object}	Response{status_code=int,error=string}
//	@Security		BearerAuth
//	@Router			/quizzes/{quizId}/results [get]
func (a *Application) getQuizResultsHandler(w http.ResponseWriter, r *http.Request) {
	quiz, user := getQuizFromCtx(r), getUserFromCtx(r)
	result, err := a.storage.GetResultByQuizAndUserId(quiz.Id, user.Id)

	if errors.Is(err, store.NotFoundError) {
		a.notFound(w, r, ResultsNotFoundError)
		return
	}

	if err != nil {
		a.internalServerError(w, r, err)
		return
	}

	result.PercentileRank = calculatePercentileRank(quiz, result.CorrectAnswersCount)

	a.dataResponse(w, r, http.StatusOK, result)
}

func (a *Application) quizzesContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParameter := chi.URLParam(r, "quizId")
		id, err := strconv.ParseInt(idParameter, 10, 64)

		if err != nil || id < 1 {
			a.badRequest(w, r, InvalidQuizIdError(idParameter))
			return
		}

		ctx := r.Context()

		quiz, err := a.storage.GetQuizById(store.QuizId(id))

		if err != nil {
			switch {
			case errors.Is(err, store.NotFoundError):
				a.notFound(w, r, QuizNotFoundError)
			default:
				a.internalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, quizCtxKey, quiz)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getQuizFromCtx(r *http.Request) *store.Quiz {
	quiz, _ := r.Context().Value(quizCtxKey).(*store.Quiz)
	return quiz
}
