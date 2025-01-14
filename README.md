# go-quiz

## To-do
* Add unit tests for storage, auth, and request handlers

## Getting started
1. Start the API => `go run ./cmd/go-quiz api`
2. Start the CLI => `go run ./cmd/go-quiz cli`
3. Answer a quiz!

## Highlights
* Robust API designed for expandability
* Interactive CLI experience with `promptui`
* OpenAPI specification and client generation for decoupled applications
* Use of stores for data allowing to easily implement a database

## What can be improved?
* Replace "primary keys" of structs like `QuizId` and `QuestionId` with UUIDs
* Containerise applications
* Move business logic from request handlers to service layer
* Integrate database, there are some bugs caused by not having transactions
* Use a `.env` file (or similar) to store configuration values that are currently hard coded
* Error handling in the CLI with custom error types
* Add rate limiting to API
* Explore concurrency to optimise request execution time
* Add caching