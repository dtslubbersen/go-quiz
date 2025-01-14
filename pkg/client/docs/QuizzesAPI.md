# \QuizzesAPI

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QuizzesGet**](QuizzesAPI.md#QuizzesGet) | **Get** /quizzes | Retrieves all quizzes
[**QuizzesQuizIdGet**](QuizzesAPI.md#QuizzesQuizIdGet) | **Get** /quizzes/{quizId} | Retrieves a quiz by ID
[**QuizzesQuizIdResultsGet**](QuizzesAPI.md#QuizzesQuizIdResultsGet) | **Get** /quizzes/{quizId}/results | Retrieves quiz results for a user
[**QuizzesQuizIdSubmitPost**](QuizzesAPI.md#QuizzesQuizIdSubmitPost) | **Post** /quizzes/{quizId}/submit | Submits answers for a quiz



## QuizzesGet

> QuizzesGet200Response QuizzesGet(ctx).Execute()

Retrieves all quizzes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dtslubbersen/go-quiz/pkg/client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuizzesAPI.QuizzesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuizzesAPI.QuizzesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuizzesGet`: QuizzesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `QuizzesAPI.QuizzesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQuizzesGetRequest struct via the builder pattern


### Return type

[**QuizzesGet200Response**](QuizzesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuizzesQuizIdGet

> QuizzesQuizIdGet200Response QuizzesQuizIdGet(ctx, quizId).Execute()

Retrieves a quiz by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dtslubbersen/go-quiz/pkg/client"
)

func main() {
	quizId := int32(56) // int32 | Quiz ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuizzesAPI.QuizzesQuizIdGet(context.Background(), quizId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuizzesAPI.QuizzesQuizIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuizzesQuizIdGet`: QuizzesQuizIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `QuizzesAPI.QuizzesQuizIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quizId** | **int32** | Quiz ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuizzesQuizIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuizzesQuizIdGet200Response**](QuizzesQuizIdGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuizzesQuizIdResultsGet

> QuizzesQuizIdResultsGet200Response QuizzesQuizIdResultsGet(ctx, quizId).Execute()

Retrieves quiz results for a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dtslubbersen/go-quiz/pkg/client"
)

func main() {
	quizId := int32(56) // int32 | Quiz ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuizzesAPI.QuizzesQuizIdResultsGet(context.Background(), quizId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuizzesAPI.QuizzesQuizIdResultsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuizzesQuizIdResultsGet`: QuizzesQuizIdResultsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `QuizzesAPI.QuizzesQuizIdResultsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quizId** | **int32** | Quiz ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuizzesQuizIdResultsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuizzesQuizIdResultsGet200Response**](QuizzesQuizIdResultsGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuizzesQuizIdSubmitPost

> QuizzesQuizIdResultsGet200Response QuizzesQuizIdSubmitPost(ctx, quizId).Payload(payload).Execute()

Submits answers for a quiz



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dtslubbersen/go-quiz/pkg/client"
)

func main() {
	quizId := int32(56) // int32 | Quiz ID
	payload := *openapiclient.NewApiSubmitQuizAnswersPayload([]openapiclient.ApiSubmitQuizAnswersPayloadAnswersInner{*openapiclient.NewApiSubmitQuizAnswersPayloadAnswersInner(int32(123), int32(123))}) // ApiSubmitQuizAnswersPayload | User's answers

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuizzesAPI.QuizzesQuizIdSubmitPost(context.Background(), quizId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuizzesAPI.QuizzesQuizIdSubmitPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuizzesQuizIdSubmitPost`: QuizzesQuizIdResultsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `QuizzesAPI.QuizzesQuizIdSubmitPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quizId** | **int32** | Quiz ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuizzesQuizIdSubmitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**ApiSubmitQuizAnswersPayload**](ApiSubmitQuizAnswersPayload.md) | User&#39;s answers | 

### Return type

[**QuizzesQuizIdResultsGet200Response**](QuizzesQuizIdResultsGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

