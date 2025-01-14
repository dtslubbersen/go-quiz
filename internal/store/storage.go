package store

import (
	"errors"
)

var (
	DuplicateEntryError = errors.New("the provided item already exists")
	NotFoundError       = errors.New("the requested item could not be found")
)

type Storage interface {
	GetQuizById(quizId QuizId) (*Quiz, error)
	ListQuizzes() ([]*Quiz, error)
	UpdateQuiz(*Quiz) error

	AddResult(*Result) (*Result, error)
	GetResultByQuizAndUserId(QuizId, UserId) (*Result, error)

	AddUserAnswer(*UserAnswer) error
	ListUserAnswersByQuizId(QuizId) ([]*UserAnswer, error)
	ListUserAnswersByUserAndQuizId(UserId, QuizId) ([]*UserAnswer, error)

	GetUserByEmail(string) (*User, error)
	GetUserById(UserId) (*User, error)
}

type InMemoryStorage struct {
	Quizzes     *InMemoryQuizStore
	Results     *InMemoryResultStore
	UserAnswers *InMemoryUserAnswerStore
	Users       *InMemoryUserStore
}

func NewStorage(seed *Seed) Storage {

	return &InMemoryStorage{
		Quizzes: &InMemoryQuizStore{
			entities: seed.quizzes,
		},
		Results: &InMemoryResultStore{
			items:  make(map[ResultId]*Result),
			nextId: 1,
		},
		UserAnswers: &InMemoryUserAnswerStore{
			items:  seed.userAnswers,
			nextId: 55,
		},
		Users: &InMemoryUserStore{
			items: seed.users,
		},
	}
}
