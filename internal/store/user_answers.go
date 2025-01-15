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

type InMemoryUserAnswerStore struct {
	mu     sync.RWMutex
	items  map[UserAnswerId]*UserAnswer
	nextId UserAnswerId
}

func (s *InMemoryStorage) AddUserAnswer(userAnswer *UserAnswer) error {
	s.UserAnswers.mu.Lock()
	defer s.UserAnswers.mu.Unlock()

	if s.UserAnswers.compositeKeyExists(userAnswer.UserId, userAnswer.QuizId, userAnswer.QuestionId) {
		return DuplicateEntryError
	}

	userAnswer.Id = s.UserAnswers.nextId
	s.UserAnswers.items[userAnswer.Id] = userAnswer
	s.UserAnswers.nextId++

	return nil
}

func (s *InMemoryStorage) ListUserAnswersByQuizId(quizId QuizId) ([]*UserAnswer, error) {
	s.UserAnswers.mu.Lock()
	defer s.UserAnswers.mu.Unlock()

	var userAnswers []*UserAnswer

	for _, userAnswer := range s.UserAnswers.items {
		if userAnswer.QuizId == quizId {
			userAnswers = append(userAnswers, userAnswer)
		}
	}

	return userAnswers, nil
}

func (s *InMemoryStorage) ListUserAnswersByUserAndQuizId(userId UserId, quizId QuizId) ([]*UserAnswer, error) {
	s.UserAnswers.mu.Lock()
	defer s.UserAnswers.mu.Unlock()

	var userAnswers []*UserAnswer

	for _, userAnswer := range s.UserAnswers.items {
		if userAnswer.UserId == userId && userAnswer.QuizId == quizId {
			userAnswers = append(userAnswers, userAnswer)
		}
	}

	return userAnswers, nil
}

func (s *InMemoryUserAnswerStore) compositeKeyExists(userId UserId, quizId QuizId, questionId QuestionId) bool {
	for _, ua := range s.items {
		if ua.UserId == userId && ua.QuizId == quizId && ua.QuestionId == questionId {
			return true
		}
	}

	return false
}
