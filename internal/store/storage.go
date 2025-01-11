package store

import (
	"errors"
)

var (
	DuplicateEntryError = errors.New("the provided item already exists")
	NotFoundError       = errors.New("the requested item could not be found")
)

type Storage struct {
	Questions interface {
		GetByQuizId(QuizId) ([]*Question, error)
	}
	Quizzes interface {
		GetById(QuizId) (*Quiz, error)
		GetAll() ([]*Quiz, error)
	}
	Results interface {
		Add(*Result) error
		GetByQuizAndUserId(QuizId, UserId) (*Result, error)
	}
	UserAnswers interface {
		Add(*UserAnswer) error
		GetByQuizId(QuizId) ([]*UserAnswer, error)
		GetByUserAndQuizId(UserId, QuizId) ([]*UserAnswer, error)
	}
	Users interface {
		GetByEmail(string) (*User, error)
		GetById(UserId) (*User, error)
	}
}

func NewStorage(seed *Seed) Storage {
	return Storage{
		Questions: &QuestionStore{
			questions: seed.questions,
		},
		Quizzes: &QuizStore{
			quizzes: seed.quizzes,
		},
		Results: &ResultStore{},
		UserAnswers: &UserAnswerStore{
			userAnswers: seed.userAnswers,
		},
		Users: &UserStore{
			users: seed.users,
		},
	}
}