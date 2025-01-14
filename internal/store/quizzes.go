package store

import "sync"

type QuizId int64

type Quiz struct {
	Id          QuizId      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Questions   []*Question `json:"questions,omitempty"`
	Performance Performance `json:"performance"`
}

type Performance struct {
	UsersTakenCount     int64       `json:"users_taken_count" default:"0"`
	CorrectAnswersCount map[int]int `json:"correct_answers_count"`
}

type QuizStore interface {
	GetById(QuizId) (*Quiz, error)
	GetAll() ([]*Quiz, error)
	Update(*Quiz) error
}

type InMemoryQuizStore struct {
	mu      sync.Mutex
	quizzes map[QuizId]*Quiz
}

func (s *InMemoryQuizStore) GetById(id QuizId) (*Quiz, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	quiz, exists := s.quizzes[id]

	if !exists {
		return nil, NotFoundError
	}

	return quiz, nil
}

func (s *InMemoryQuizStore) GetAll() ([]*Quiz, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	quizzes := make([]*Quiz, 0, len(s.quizzes))

	for _, quiz := range s.quizzes {
		quizzes = append(quizzes, quiz)
	}

	return quizzes, nil
}

func (s *InMemoryQuizStore) Update(quiz *Quiz) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.quizzes[quiz.Id]; !exists {
		return NotFoundError
	}

	s.quizzes[quiz.Id] = quiz
	return nil
}
