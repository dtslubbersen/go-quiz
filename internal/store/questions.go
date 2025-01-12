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

type QuestionStore struct {
	mu        sync.Mutex
	questions map[QuestionId]*Question
}

func (s *QuestionStore) GetByQuizId(quizId QuizId) ([]*Question, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var questions []*Question

	for _, question := range s.questions {
		if question.QuizId == quizId {
			questions = append(questions, question)
		}
	}

	return questions, nil
}
