package store

import "github.com/stretchr/testify/mock"

func NewMockStorage() Storage {
	return Storage{
		Questions:   &MockQuestionStore{},
		Quizzes:     &MockQuizStore{},
		Results:     &MockResultStore{},
		UserAnswers: &MockUserAnswerStore{},
		Users:       &MockUserStore{},
	}
}

type MockQuestionStore struct {
	mock.Mock
}

func (m *MockQuestionStore) GetByQuizId(quizId QuizId) ([]*Question, error) {
	args := m.Called(quizId)
	return args.Get(0).([]*Question), args.Error(1)
}

type MockQuizStore struct {
	mock.Mock
}

func (m *MockQuizStore) GetById(quizId QuizId) (*Quiz, error) {
	args := m.Called(quizId)
	return args.Get(0).(*Quiz), args.Error(1)
}

func (m *MockQuizStore) GetAll() ([]*Quiz, error) {
	args := m.Called()
	return args.Get(0).([]*Quiz), args.Error(1)
}

func (m *MockQuizStore) Update(quiz *Quiz) error {
	args := m.Called(quiz)
	return args.Error(0)
}

type MockResultStore struct {
	mock.Mock
}

func (m *MockResultStore) Add(result *Result) (*Result, error) {
	args := m.Called(result)
	return args.Get(0).(*Result), args.Error(1)
}

func (m *MockResultStore) GetByQuizAndUserId(quizId QuizId, userId UserId) (*Result, error) {
	args := m.Called(quizId, userId)
	return args.Get(0).(*Result), args.Error(1)
}

type MockUserAnswerStore struct {
	mock.Mock
}

func (m *MockUserAnswerStore) Add(userAnswer *UserAnswer) error {
	args := m.Called(userAnswer)
	return args.Error(0)
}

func (m *MockUserAnswerStore) GetByQuizId(quizId QuizId) ([]*UserAnswer, error) {
	args := m.Called(quizId)
	return args.Get(0).([]*UserAnswer), args.Error(1)
}

func (m *MockUserAnswerStore) GetByUserAndQuizId(userId UserId, quizId QuizId) ([]*UserAnswer, error) {
	args := m.Called(userId, quizId)
	return args.Get(0).([]*UserAnswer), args.Error(1)
}

type MockUserStore struct {
	mock.Mock
}

func (m *MockUserStore) GetByEmail(email string) (*User, error) {
	args := m.Called(email)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserStore) GetById(userId UserId) (*User, error) {
	args := m.Called(userId)
	return args.Get(0).(*User), args.Error(1)
}
