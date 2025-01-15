package store

import "sync"

type QuestionId int64

type Question struct {
	Id                 QuestionId `json:"id"`
	QuizId             QuizId     `json:"quiz_id"`
	Value              string     `json:"value"`
	Answers            []Answer   `json:"answers"`
	CorrectAnswerIndex int        `json:"correct_answer_index"`
}

type Answer string

type InMemoryQuestionStore struct {
	mu    sync.Mutex
	items map[QuestionId]*Question
}

func (s *InMemoryStorage) ListQuestionsByQuizId(quizId QuizId) ([]*Question, error) {
	s.Questions.mu.Lock()
	defer s.Questions.mu.Unlock()

	var questions []*Question

	for _, question := range s.Questions.items {
		if question.QuizId == quizId {
			questions = append(questions, question)
		}
	}

	return questions, nil
}
