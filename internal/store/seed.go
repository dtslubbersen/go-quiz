package store

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
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
	questions, _ := readMapFromJsonFile[Question, QuestionId](getDataFilePath("questions.json"), "Id")
	quizzes, _ := readMapFromJsonFile[Quiz, QuizId](getDataFilePath("quizzes.json"), "Id")
	userAnswers, _ := readMapFromJsonFile[UserAnswer, UserAnswerId](getDataFilePath("user_answers.json"), "Id")

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

func readMapFromJsonFile[T any, K comparable](fileName string, keyFieldName string) (map[K]*T, error) {
	file, err := os.ReadFile(fileName)

	if err != nil {
		return nil, fmt.Errorf("could not read file %s: %v", fileName, err)
	}

	var items []T

	if err := json.Unmarshal(file, &items); err != nil {
		return nil, fmt.Errorf("could not unmarshall json %s: %v", fileName, err)
	}

	itemsMap := make(map[K]*T)
	for _, item := range items {
		v := reflect.ValueOf(item)
		field := v.FieldByName(keyFieldName)
		if !field.IsValid() {
			log.Fatalf("Missing '%s' field in struct %T", keyFieldName, item)
		}

		mapKeyValue := field.Interface()
		itemsMap[mapKeyValue.(K)] = &item
	}

	return itemsMap, nil
}
