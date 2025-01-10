package api

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"go-quiz/internal/store"
	"net/http"
	"strconv"
)

type quizKey string

const quizCtxKey quizKey = "post"

func (a *application) getQuizzesHandler(w http.ResponseWriter, r *http.Request) {
	quizzes, err := a.store.Quizzes.GetAll()

	if err != nil {
		a.badRequest(w, r, err)
		return
	}

	if err := a.writeDataResponse(w, http.StatusOK, quizzes); err != nil {
		a.internalServerError(w, r, err)
	}
}

func (a *application) getQuizByIdHandler(w http.ResponseWriter, r *http.Request) {
	quiz := getQuizFromCtx(r)

	if err := a.writeDataResponse(w, http.StatusOK, quiz); err != nil {
		a.internalServerError(w, r, err)
	}
}

type SubmitQuizAnswersPayload struct {
	Answers []struct {
		QuestionId  int64 `json:"question_id" validate:"required"`
		AnswerIndex int   `json:"answer_index" validate:"required"`
	} `json:"answers" validate:"required"`
}

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
	answersMap := make(map[int64]int)

	for _, answer := range payload.Answers {
		answersMap[answer.QuestionId] = answer.AnswerIndex
	}

	var userAnswers []store.UserAnswer

	for _, question := range quiz.Questions {
		answerIndex := answersMap[int64(question.Id)]
		userAnswers = append(userAnswers, store.UserAnswer{
			QuestionId:  question.Id,
			UserId:      user.Id,
			AnswerIndex: answerIndex,
			IsCorrect:   question.CorrectAnswerIndex == answerIndex,
		})
	}

	w.WriteHeader(http.StatusOK)
}

func (a *application) getResultsHandler(w http.ResponseWriter, r *http.Request) {
	quiz, user := getQuizFromCtx(r), getUserFromCtx(r)
	result, err := a.store.Results.GetByQuizAndUserId(quiz.Id, user.Id)

	if errors.Is(err, store.NotFoundError) {
		a.notFound(w, r, err)
		return
	}

	if err := a.writeDataResponse(w, http.StatusOK, result); err != nil {
		a.internalServerError(w, r, err)
	}
}

func (a *application) quizzesContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParameter := chi.URLParam(r, "quizId")
		id, err := strconv.ParseInt(idParameter, 10, 64)

		if err != nil {
			a.internalServerError(w, r, err)
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
