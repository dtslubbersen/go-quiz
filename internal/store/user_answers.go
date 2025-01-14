package store

import (
	"sync"
)

type UserAnswerId int64

type UserAnswer struct {
	Id          UserAnswerId `json:"id"`
	UserId      UserId       `json:"user_id"`
	QuizId      QuizId       `json:"quiz_id"`
	QuestionId  QuestionId   `json:"question_id"`
	AnswerIndex int          `json:"answer_index"`
	IsCorrect   bool         `json:"is_correct"`
}

type UserAnswerStore interface {
	Add(*UserAnswer) error
	GetByQuizId(QuizId) ([]*UserAnswer, error)
	GetByUserAndQuizId(UserId, QuizId) ([]*UserAnswer, error)
}

type InMemoryUserAnswerStore struct {
	mu          sync.RWMutex
	userAnswers map[UserAnswerId]*UserAnswer
	nextId      UserAnswerId
}

func (s *InMemoryUserAnswerStore) Add(userAnswer *UserAnswer) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.compositeKeyExists(userAnswer.UserId, userAnswer.QuizId, userAnswer.QuestionId) {
		return DuplicateEntryError
	}

	userAnswer.Id = s.nextId
	s.userAnswers[userAnswer.Id] = userAnswer
	s.nextId++

	return nil
}

func (s *InMemoryUserAnswerStore) GetByQuizId(quizId QuizId) ([]*UserAnswer, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var userAnswers []*UserAnswer

	for _, userAnswer := range s.userAnswers {
		if userAnswer.QuizId == quizId {
			userAnswers = append(userAnswers, userAnswer)
		}
	}

	return userAnswers, nil
}

func (s *InMemoryUserAnswerStore) GetByUserAndQuizId(userId UserId, quizId QuizId) ([]*UserAnswer, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var userAnswers []*UserAnswer

	for _, userAnswer := range s.userAnswers {
		if userAnswer.UserId == userId && userAnswer.QuizId == quizId {
			userAnswers = append(userAnswers, userAnswer)
		}
	}

	return userAnswers, nil
}

func (s *InMemoryUserAnswerStore) compositeKeyExists(userId UserId, quizId QuizId, questionId QuestionId) bool {
	for _, ua := range s.userAnswers {
		if ua.UserId == userId && ua.QuizId == quizId && ua.QuestionId == questionId {
			return true
		}
	}

	return false
}
