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

type InMemoryResultStore struct {
	mu     sync.Mutex
	items  map[ResultId]*Result
	nextId ResultId
}

func (s *InMemoryStorage) AddResult(result *Result) (*Result, error) {
	s.Results.mu.Lock()
	defer s.Results.mu.Unlock()

	if _, err := s.Results.getByCompositeKey(result.QuizId, result.UserId); err == nil {
		return nil, DuplicateEntryError
	}

	result.Id = s.Results.nextId
	s.Results.items[result.Id] = result
	s.Results.nextId++
	return result, nil

}

func (s *InMemoryStorage) GetResultByQuizAndUserId(quizId QuizId, userId UserId) (*Result, error) {
	s.Results.mu.Lock()
	defer s.Results.mu.Unlock()

	return s.Results.getByCompositeKey(quizId, userId)
}

func (s *InMemoryResultStore) getByCompositeKey(quizId QuizId, userId UserId) (*Result, error) {
	for _, result := range s.items {
		if result.QuizId == quizId && result.UserId == userId {
			return result, nil
		}
	}

	return nil, NotFoundError
}
