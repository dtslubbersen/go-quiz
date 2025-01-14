build:
	@go build ./cmd/go-quiz

gen-docs:
	@swag fmt
	@swag init -g ./cmd/go-quiz/main.go

gen-client:
	@openapi-generator-cli generate -i ./docs/swagger.yaml -g go -o ./pkg/client  --git-repo-id go-quiz/pkg/client --git-user-id dtslubbersen
	@rm -f ./pkg/client/go.mod  # Remove the generated go.mod file

gen-mock:
	@mockgen -package store -destination internal/store/mock/storage.go github.com/dtslubbersen/go-quiz/internal/store Storage
#	@mockgen -package store -destination internal/store/mock/quizzes.go github.com/dtslubbersen/go-quiz/internal/store QuizStore
#	@mockgen -package store -destination internal/store/mock/results.go github.com/dtslubbersen/go-quiz/internal/store ResultStore
#	@mockgen -package store -destination internal/store/mock/user_answers.go github.com/dtslubbersen/go-quiz/internal/store UserAnswerStore
#	@mockgen -package store -destination internal/store/mock/users.go github.com/dtslubbersen/go-quiz/internal/store UserStore

docker-build:
	@docker build -t go-quiz -f cmd/go-quiz/Dockerfile .