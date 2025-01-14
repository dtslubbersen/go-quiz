package store

import (
	"errors"
)

var (
	DuplicateEntryError = errors.New("the provided item already exists")
	NotFoundError       = errors.New("the requested item could not be found")
)

type Storage interface {
	Questions() QuestionStore
	Quizzes() QuizStore
	Results() ResultStore
	UserAnswers() UserAnswerStore
	Users() UserStore
}

type InMemoryStorage struct {
	questions   QuestionStore
	quizzes     QuizStore
	results     ResultStore
	userAnswers UserAnswerStore
	users       UserStore
}

func (s *InMemoryStorage) Questions() QuestionStore {
	return s.questions
}

func (s *InMemoryStorage) Quizzes() QuizStore {
	return s.quizzes
}

func (s *InMemoryStorage) Results() ResultStore {
	return s.results
}

func (s *InMemoryStorage) UserAnswers() UserAnswerStore {
	return s.userAnswers
}

func (s *InMemoryStorage) Users() UserStore {
	return s.users
}

func NewStorage(seed *Seed) Storage {
	return &InMemoryStorage{
		questions: &InMemoryQuestionStore{
			questions: seed.questions,
		},
		quizzes: &InMemoryQuizStore{
			quizzes: seed.quizzes,
		},
		results: &InMemoryResultStore{
			results: make(map[ResultId]*Result),
			nextId:  1,
		},
		userAnswers: &InMemoryUserAnswerStore{
			userAnswers: seed.userAnswers,
			nextId:      55,
		},
		users: &InMemoryUserStore{
			users: seed.users,
		},
	}
}
