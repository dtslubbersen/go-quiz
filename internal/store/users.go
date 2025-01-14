package store

import (
	"golang.org/x/crypto/bcrypt"
	"sync"
)

type UserId int64

type User struct {
	Id        UserId   `json:"id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	Password  password `json:"password"`
}

type password struct {
	value *string
	hash  []byte
}

type UserStore interface {
	GetByEmail(string) (*User, error)
	GetById(UserId) (*User, error)
}

type InMemoryUserStore struct {
	mu    sync.Mutex
	users map[UserId]*User
}

func (p *password) Set(value string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.value = &value
	p.hash = hash

	return nil
}

func (p *password) Compare(text string) error {
	return bcrypt.CompareHashAndPassword(p.hash, []byte(text))
}

func (s *InMemoryUserStore) GetByEmail(email string) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, user := range s.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, NotFoundError
}

func (s *InMemoryUserStore) GetById(id UserId) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]

	if !exists {
		return nil, NotFoundError
	}

	return user, nil
}
