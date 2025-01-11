package store

import (
	"go-quiz/pkg"
	"path/filepath"
	"runtime"
)

const defaultPassword string = "password"

type Seed struct {
	questions   map[QuestionId]*Question
	quizzes     map[QuizId]*Quiz
	userAnswers map[UserAnswerId]*UserAnswer
	users       map[UserId]*User
}

func NewSeed() *Seed {
	questions, _ := pkg.ReadMapFromJsonFile[Question, QuestionId](getDataFilePath("questions.json"))
	quizzes, _ := pkg.ReadMapFromJsonFile[Quiz, QuizId](getDataFilePath("quizzes.json"))
	userAnswers, _ := pkg.ReadMapFromJsonFile[UserAnswer, UserAnswerId](getDataFilePath("user_answers.json"))

	return &Seed{
		questions:   questions,
		quizzes:     quizzes,
		userAnswers: userAnswers,
		users:       getSeedUsers(),
	}
}

func getDataFilePath(fileName string) string {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(filepath.Dir(filepath.Dir(b)))
	return filepath.Join(basePath, "data", fileName)
}

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
