package store

import "golang.org/x/crypto/bcrypt"

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

type UserStore struct {
	users map[UserId]*User
}

func (s *UserStore) GetByEmail(email string) (*User, error) {
	for _, user := range s.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, NotFoundError
}

func (s *UserStore) GetById(id UserId) (*User, error) {
	user, exists := s.users[id]

	if !exists {
		return nil, NotFoundError
	}

	return user, nil
}
