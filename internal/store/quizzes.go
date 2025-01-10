package store

type QuizId int64

type Quiz struct {
	Id          QuizId      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Questions   []*Question `json:"questions"`
}

type QuizStore struct {
	quizzes map[QuizId]*Quiz
}

func (s *QuizStore) GetById(id QuizId) (*Quiz, error) {
	quiz, exists := s.quizzes[id]

	if !exists {
		return nil, NotFoundError
	}

	return quiz, nil
}

func (s *QuizStore) GetAll() ([]*Quiz, error) {
	quizzes := make([]*Quiz, 0, len(s.quizzes))

	for _, quiz := range s.quizzes {
		quizzes = append(quizzes, quiz)
	}

	return quizzes, nil
}
