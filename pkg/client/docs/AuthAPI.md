# \AuthAPI

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthTokenPost**](AuthAPI.md#AuthTokenPost) | **Post** /auth/token | Generates an authentication token



## AuthTokenPost

> AuthTokenPost201Response AuthTokenPost(ctx).Payload(payload).Execute()

Generates an authentication token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	payload := *openapiclient.NewApiCreateTokenPayload("demo@quiz.com", "password") // ApiCreateTokenPayload | User credentials

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthTokenPost(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthTokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthTokenPost`: AuthTokenPost201Response
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthTokenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**ApiCreateTokenPayload**](ApiCreateTokenPayload.md) | User credentials | 

### Return type

[**AuthTokenPost201Response**](AuthTokenPost201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

