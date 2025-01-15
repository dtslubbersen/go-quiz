# go-quiz
This project is the first I've built in `go`, I spent approximately 28 hours on it.

## Getting started
1. Start the API => `go run ./cmd/go-quiz api`
2. Start the CLI => `go run ./cmd/go-quiz cli`
3. Answer a quiz!

## What can be improved?
* Add more test coverage
* Replace "primary keys" of structs like `QuizId` and `QuestionId` with UUIDs
* Containerise applications
* Request handlers are bloated, move business logic from to service layer. Because of this logic to calculate rank percentile isnt easily testable
* Integrate database, there are some bugs caused by not having transactions
* Use a `.env` file (or similar) to store configuration values that are currently hard coded
* Error handling in the CLI with custom error types
* Add rate limiting to API
* Explore concurrency to optimise request execution time
* Add caching
