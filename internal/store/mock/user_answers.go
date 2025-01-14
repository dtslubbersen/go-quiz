// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dtslubbersen/go-quiz/internal/store (interfaces: UserAnswerStore)
//
// Generated by this command:
//
//	mockgen -package store -destination internal/store/mock/user_answers.go github.com/dtslubbersen/go-quiz/internal/store UserAnswerStore
//

// Package store is a generated GoMock package.
package store

import (
	reflect "reflect"

	store "github.com/dtslubbersen/go-quiz/internal/store"
	gomock "go.uber.org/mock/gomock"
)

// MockUserAnswerStore is a mock of UserAnswerStore interface.
type MockUserAnswerStore struct {
	ctrl     *gomock.Controller
	recorder *MockUserAnswerStoreMockRecorder
	isgomock struct{}
}

// MockUserAnswerStoreMockRecorder is the mock recorder for MockUserAnswerStore.
type MockUserAnswerStoreMockRecorder struct {
	mock *MockUserAnswerStore
}

// NewMockUserAnswerStore creates a new mock instance.
func NewMockUserAnswerStore(ctrl *gomock.Controller) *MockUserAnswerStore {
	mock := &MockUserAnswerStore{ctrl: ctrl}
	mock.recorder = &MockUserAnswerStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserAnswerStore) EXPECT() *MockUserAnswerStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockUserAnswerStore) Add(arg0 *store.UserAnswer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockUserAnswerStoreMockRecorder) Add(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockUserAnswerStore)(nil).Add), arg0)
}

// GetByQuizId mocks base method.
func (m *MockUserAnswerStore) GetByQuizId(arg0 store.QuizId) ([]*store.UserAnswer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListQuestionsByQuizId", arg0)
	ret0, _ := ret[0].([]*store.UserAnswer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByQuizId indicates an expected call of GetByQuizId.
func (mr *MockUserAnswerStoreMockRecorder) GetByQuizId(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListQuestionsByQuizId", reflect.TypeOf((*MockUserAnswerStore)(nil).GetByQuizId), arg0)
}

// GetByUserAndQuizId mocks base method.
func (m *MockUserAnswerStore) GetByUserAndQuizId(arg0 store.UserId, arg1 store.QuizId) ([]*store.UserAnswer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserAndQuizId", arg0, arg1)
	ret0, _ := ret[0].([]*store.UserAnswer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserAndQuizId indicates an expected call of GetByUserAndQuizId.
func (mr *MockUserAnswerStoreMockRecorder) GetByUserAndQuizId(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserAndQuizId", reflect.TypeOf((*MockUserAnswerStore)(nil).GetByUserAndQuizId), arg0, arg1)
}
