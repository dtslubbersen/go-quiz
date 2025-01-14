build:
	@go build ./cmd/go-quiz

.PHONY: gen-docs
gen-docs:
	@swag fmt
	@swag init -g ./cmd/go-quiz/main.go

.PHONY: gen-client
gen-client:
	@openapi-generator-cli generate -i ./docs/swagger.yaml -g go -o ./pkg/client
	@rm -f ./pkg/client/go.mod  # Remove the generated go.mod file

.PHONY: docker-build
docker-build:
	@docker build -t go-quiz -f cmd/go-quiz/Dockerfile .