.PHONY: gen-client
gen-client:
	@openapi-generator-cli generate -i ./docs/swagger.yaml -g go -o ./pkg/client
	@rm -f ./pkg/client/go.mod  # Remove the generated go.mod file

.PHONY: gen-docs
gen-docs:
	@swag init -g ./api/main.go -d cmd,internal && swag fmt

.PHONE: docker-build
docker-build:
	@docker build -t go-quiz-api -f cmd/api/Dockerfile .