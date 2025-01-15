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

docker-build:
	@docker build -t go-quiz -f cmd/go-quiz/Dockerfile .