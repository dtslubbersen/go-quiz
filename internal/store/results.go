package store

import (
	"sync"
)

type ResultId int64

type Result struct {
	Id            ResultId `json:"id"`
	QuizId        QuizId   `json:"quiz_id"`
	QuestionCount int      `json:"question_count"`
	UserId        UserId   `json:"user_id"`
	Score         int      `json:"user_score"`
	TopPercentile int      `json:"user_percentile"`
}

type ResultStore struct {
	mu      sync.Mutex
	results map[ResultId]*Result
	nextId  ResultId
}

func (s *ResultStore) Add(result *Result) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, err := s.getByCompositeKey(result.QuizId, result.UserId); err == nil {
		return DuplicateEntryError
	}

	result.Id = s.nextId
	s.results[result.Id] = result
	s.nextId++
	return nil

}

func (s *ResultStore) GetByQuizAndUserId(quizId QuizId, userId UserId) (*Result, error) {
	return s.getByCompositeKey(quizId, userId)
}

func (s *ResultStore) getByCompositeKey(quizId QuizId, userId UserId) (*Result, error) {
	for _, result := range s.results {
		if result.QuizId == quizId && result.UserId == userId {
			return result, nil
		}
	}

	return nil, NotFoundError
}
