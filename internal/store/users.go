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

type InMemoryUserStore struct {
	mu    sync.Mutex
	items map[UserId]*User
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

func (s *InMemoryStorage) GetUserByEmail(email string) (*User, error) {
	s.Users.mu.Lock()
	defer s.Users.mu.Unlock()

	for _, user := range s.Users.items {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, NotFoundError
}

func (s *InMemoryStorage) GetUserById(id UserId) (*User, error) {
	s.Users.mu.Lock()
	defer s.Users.mu.Unlock()

	user, exists := s.Users.items[id]

	if !exists {
		return nil, NotFoundError
	}

	return user, nil
}
