package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go-quiz/internal/store"
	"math"
	"net/http"
	"strconv"
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
//	@Success		200	{object}	Response{data=[]store.Quiz}
//	@Failure		400	{object}	Response{error=string}
//	@Failure		500	{object}	Response{error=string}
//	@Security		BearerAuth
//	@Router			/quizzes [get]
func (a *application) getQuizzesHandler(w http.ResponseWriter, r *http.Request) {
	quizzes, err := a.store.Quizzes.GetAll()

	if err != nil {
		a.badRequest(w, r, err)
		return
	}

	if err := a.dataResponse(w, http.StatusOK, quizzes); err != nil {
		a.internalServerError(w, r, err)
	}
}

// GetQuizById godoc
//
//	@Summary		Retrieves a quiz by ID
//	@Description	Fetches a specific quiz using its ID from the in-memory store
//	@Tags			quizzes
//	@Accept			json
//	@Produce		json
//	@Param			quizId	path		int	true	"Quiz ID"
//	@Success		200		{object}	Response{data=store.Quiz}
//	@Failure		400		{object}	Response{error=string}
//	@Failure		404		{object}	Response{error=string}
//	@Failure		500		{object}	Response{error=string}
//	@Security		BearerAuth
//	@Router			/quizzes/{quizId} [get]
func (a *application) getQuizByIdHandler(w http.ResponseWriter, r *http.Request) {
	quiz := getQuizFromCtx(r)

	if err := a.dataResponse(w, http.StatusOK, quiz); err != nil {
		a.internalServerError(w, r, err)
	}
}

type SubmitQuizAnswersPayload struct {
	Answers []struct {
		QuestionId  int64 `json:"question_id" validate:"required"`
		AnswerIndex int   `json:"answer_index" validate:"required"`
	} `json:"answers" validate:"required"`
}

// SubmitQuizAnswers godoc
//
//	@Summary		Submits answers for a quiz
//	@Description	Allows a user to submit answers for a given quiz
//	@Tags			quizzes
//	@Accept			json
//	@Produce		json
//	@Param			quizId	path		int							true	"Quiz ID"
//	@Param			payload	body		SubmitQuizAnswersPayload	true	"User's answers"
//	@Success		200		{object}	Response{data=store.Result}
//	@Failure		400		{object}	Response{error=string}
//	@Failure		500		{object}	Response{error=string}
//	@Security		BearerAuth
//	@Router			/quizzes/{quizId}/submit [post]
func (a *application) submitAnswersHandler(w http.ResponseWriter, r *http.Request) {
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

	if result, _ := a.store.Results.GetByQuizAndUserId(quiz.Id, user.Id); result != nil {
		a.badRequest(w, r, errors.New("quiz already answered"))
		return
	}

	if len(quiz.Questions) != len(payload.Answers) {
		a.badRequest(w, r, errors.New("amount of answers and question count do not match"))
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

		if err := a.store.UserAnswers.Add(userAnswer); err != nil {
			a.badRequest(w, r, err)
			return
		}
	}

	percentileRank := calculatePercentileRank(quiz, correctAnswersCount)
	quiz.Performance.UsersTakenCount++

	if err := a.store.Quizzes.Update(quiz); err != nil {
		a.internalServerError(w, r, err)
		return
	}

	result := &store.Result{
		QuizId:              quiz.Id,
		QuestionCount:       len(quiz.Questions),
		UserId:              user.Id,
		CorrectAnswersCount: correctAnswersCount,
		PercentileRank:      percentileRank,
	}

	newResult, err := a.store.Results.Add(result)

	if err != nil {
		a.badRequest(w, r, err)
		return
	}

	if err := a.dataResponse(w, http.StatusOK, newResult); err != nil {
		a.internalServerError(w, r, err)
	}
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
//	@Success		200		{object}	Response{data=store.Result}
//	@Failure		400		{object}	Response{error=string}
//	@Failure		404		{object}	Response{error=string}
//	@Failure		500		{object}	Response{error=string}
//	@Security		BearerAuth
//	@Router			/quizzes/{quizId}/results [get]
func (a *application) getQuizResultsHandler(w http.ResponseWriter, r *http.Request) {
	quiz, user := getQuizFromCtx(r), getUserFromCtx(r)
	result, err := a.store.Results.GetByQuizAndUserId(quiz.Id, user.Id)

	if errors.Is(err, store.NotFoundError) {
		a.notFound(w, r, err)
		return
	}

	_, err = a.store.UserAnswers.GetByQuizId(quiz.Id)

	if err != nil {
		a.internalServerError(w, r, err)
		return
	}

	result.PercentileRank = calculatePercentileRank(quiz, result.CorrectAnswersCount)

	if err := a.dataResponse(w, http.StatusOK, result); err != nil {
		a.internalServerError(w, r, err)
	}
}

func (a *application) quizzesContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParameter := chi.URLParam(r, "quizId")
		id, err := strconv.ParseInt(idParameter, 10, 64)

		if err != nil {
			a.badRequest(w, r, fmt.Errorf("invalid quizId %s", idParameter))
			return
		}

		ctx := r.Context()

		quiz, err := a.store.Quizzes.GetById(store.QuizId(id))

		if err != nil {
			switch {
			case errors.Is(err, store.NotFoundError):
				a.notFound(w, r, err)
			default:
				a.internalServerError(w, r, err)
			}
			return
		}

		questions, err := a.store.Questions.GetByQuizId(store.QuizId(id))

		if err != nil {
			a.internalServerError(w, r, err)
			return
		}

		quiz.Questions = questions

		ctx = context.WithValue(ctx, quizCtxKey, quiz)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getQuizFromCtx(r *http.Request) *store.Quiz {
	quiz, _ := r.Context().Value(quizCtxKey).(*store.Quiz)
	return quiz
}
