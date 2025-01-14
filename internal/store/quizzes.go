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

type InMemoryQuizStore struct {
	mu       sync.Mutex
	entities map[QuizId]*Quiz
}

func (s *InMemoryStorage) GetQuizById(id QuizId) (*Quiz, error) {
	s.Quizzes.mu.Lock()
	defer s.Quizzes.mu.Unlock()

	quiz, exists := s.Quizzes.entities[id]

	if !exists {
		return nil, NotFoundError
	}

	return quiz, nil
}

func (s *InMemoryStorage) ListQuizzes() ([]*Quiz, error) {
	s.Quizzes.mu.Lock()
	defer s.Quizzes.mu.Unlock()

	quizzes := make([]*Quiz, 0, len(s.Quizzes.entities))

	for _, quiz := range s.Quizzes.entities {
		quizzes = append(quizzes, quiz)
	}

	return quizzes, nil
}

func (s *InMemoryStorage) UpdateQuiz(quiz *Quiz) error {
	s.Quizzes.mu.Lock()
	defer s.Quizzes.mu.Unlock()

	if _, exists := s.Quizzes.entities[quiz.Id]; !exists {
		return NotFoundError
	}

	s.Quizzes.entities[quiz.Id] = quiz
	return nil
}
