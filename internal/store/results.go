package store

import (
	"sync"
)

type ResultId int64

type Result struct {
	Id                  ResultId `json:"id"`
	QuizId              QuizId   `json:"quiz_id"`
	QuestionCount       int      `json:"question_count"`
	UserId              UserId   `json:"user_id"`
	CorrectAnswersCount int      `json:"correct_answers_count"`
	PercentileRank      float64  `json:"top_percentile"`
}

type ResultStore interface {
	Add(*Result) (*Result, error)
	GetByQuizAndUserId(QuizId, UserId) (*Result, error)
}

type InMemoryResultStore struct {
	mu      sync.Mutex
	results map[ResultId]*Result
	nextId  ResultId
}

func (s *InMemoryResultStore) Add(result *Result) (*Result, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, err := s.getByCompositeKey(result.QuizId, result.UserId); err == nil {
		return nil, DuplicateEntryError
	}

	result.Id = s.nextId
	s.results[result.Id] = result
	s.nextId++
	return result, nil

}

func (s *InMemoryResultStore) GetByQuizAndUserId(quizId QuizId, userId UserId) (*Result, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.getByCompositeKey(quizId, userId)
}

func (s *InMemoryResultStore) getByCompositeKey(quizId QuizId, userId UserId) (*Result, error) {
	for _, result := range s.results {
		if result.QuizId == quizId && result.UserId == userId {
			return result, nil
		}
	}

	return nil, NotFoundError
}
