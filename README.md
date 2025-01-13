# go-quiz

## Getting started
Unfortunately I ran into issues generating executables (see cmd/api/Dockerfile). As a result one must open 2 terminals to run the project:
1. In the first execute `go run .`, this will start the API
2. Navigate to `main.go`, comment line 8, uncomment line 9 and execute `go run . start` to begin the interactive CLI

I could have circumvented this by throwing both applications in the main package but that didn't feel like an architecturally sound decision. Plus this wouldn't have been an issue if the WSL dev experience was not a PITA :)

## Highlights
* Robust API designed for expandability
* Interactive CLI experience with `promptui`
* OpenAPI specification and client generation for decoupled applications
* Use of stores for data allowing to easily implement a database

## What can be improved?
* Use a .env file (or similar) to store configuration values that are currently hard coded
* Error handling in the CLI with custom error types
* Add rate limiting to API
* Explore concurrency to optimise request execution time