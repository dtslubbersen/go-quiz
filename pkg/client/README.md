# Go API client for openapi

This is the API documentation for go-quiz, a simple Quiz API allowing users to obtain quizzes, answer the questions and see their results compared to other users.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Generator version: 7.10.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://github.com/dtslubbersen/go-quiz](https://github.com/dtslubbersen/go-quiz)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/dtslubbersen/go-quiz/pkg/client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8080/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthAPI* | [**AuthTokenPost**](docs/AuthAPI.md#authtokenpost) | **Post** /auth/token | Generates an authentication token
*QuizzesAPI* | [**QuizzesGet**](docs/QuizzesAPI.md#quizzesget) | **Get** /quizzes | Retrieves all quizzes
*QuizzesAPI* | [**QuizzesQuizIdGet**](docs/QuizzesAPI.md#quizzesquizidget) | **Get** /quizzes/{quizId} | Retrieves a quiz by ID
*QuizzesAPI* | [**QuizzesQuizIdResultsGet**](docs/QuizzesAPI.md#quizzesquizidresultsget) | **Get** /quizzes/{quizId}/results | Retrieves quiz results for a user
*QuizzesAPI* | [**QuizzesQuizIdSubmitPost**](docs/QuizzesAPI.md#quizzesquizidsubmitpost) | **Post** /quizzes/{quizId}/submit | Submits answers for a quiz


## Documentation For Models

 - [ApiCreateTokenPayload](docs/ApiCreateTokenPayload.md)
 - [ApiQuestionAnswerPayload](docs/ApiQuestionAnswerPayload.md)
 - [ApiResponse](docs/ApiResponse.md)
 - [ApiSubmitQuizAnswersPayload](docs/ApiSubmitQuizAnswersPayload.md)
 - [ApiTokenCreatedResponse](docs/ApiTokenCreatedResponse.md)
 - [ApiUserResponse](docs/ApiUserResponse.md)
 - [AuthTokenPost201Response](docs/AuthTokenPost201Response.md)
 - [AuthTokenPost400Response](docs/AuthTokenPost400Response.md)
 - [QuizzesGet200Response](docs/QuizzesGet200Response.md)
 - [QuizzesGet400Response](docs/QuizzesGet400Response.md)
 - [QuizzesQuizIdGet200Response](docs/QuizzesQuizIdGet200Response.md)
 - [QuizzesQuizIdResultsGet200Response](docs/QuizzesQuizIdResultsGet200Response.md)
 - [QuizzesQuizIdSubmitPost200Response](docs/QuizzesQuizIdSubmitPost200Response.md)
 - [QuizzesQuizIdSubmitPostRequest](docs/QuizzesQuizIdSubmitPostRequest.md)
 - [StorePerformance](docs/StorePerformance.md)
 - [StoreQuestion](docs/StoreQuestion.md)
 - [StoreQuiz](docs/StoreQuiz.md)
 - [StoreResult](docs/StoreResult.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### BearerAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: BearerAuth and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"BearerAuth": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

dtslubbersen@gmail.com

