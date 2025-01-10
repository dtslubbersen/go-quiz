package store

const defaultPassword string = "password"

func getSeedUsers() map[UserId]*User {
	demoUser := User{
		Id:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "demo@quiz.com",
	}
	_ = demoUser.Password.Set(defaultPassword)

	bobUser := User{
		Id:        2,
		FirstName: "Bob",
		LastName:  "Ross",
		Email:     "bob.ross@gmail.com",
	}
	_ = bobUser.Password.Set(defaultPassword)

	janeUser := User{
		Id:        3,
		FirstName: "Jane",
		LastName:  "Goodall",
		Email:     "jane.goodall@gmail.com",
	}
	_ = bobUser.Password.Set(defaultPassword)

	return map[UserId]*User{
		demoUser.Id: &demoUser,
		bobUser.Id:  &bobUser,
		janeUser.Id: &janeUser,
	}
}
