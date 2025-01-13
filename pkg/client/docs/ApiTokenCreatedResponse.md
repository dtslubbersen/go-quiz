# ApiTokenCreatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresIn** | Pointer to **int32** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**ApiUserResponse**](ApiUserResponse.md) |  | [optional] 

## Methods

### NewApiTokenCreatedResponse

`func NewApiTokenCreatedResponse() *ApiTokenCreatedResponse`

NewApiTokenCreatedResponse instantiates a new ApiTokenCreatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTokenCreatedResponseWithDefaults

`func NewApiTokenCreatedResponseWithDefaults() *ApiTokenCreatedResponse`

NewApiTokenCreatedResponseWithDefaults instantiates a new ApiTokenCreatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresIn

`func (o *ApiTokenCreatedResponse) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *ApiTokenCreatedResponse) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *ApiTokenCreatedResponse) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *ApiTokenCreatedResponse) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetToken

`func (o *ApiTokenCreatedResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ApiTokenCreatedResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ApiTokenCreatedResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ApiTokenCreatedResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUser

`func (o *ApiTokenCreatedResponse) GetUser() ApiUserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ApiTokenCreatedResponse) GetUserOk() (*ApiUserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ApiTokenCreatedResponse) SetUser(v ApiUserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *ApiTokenCreatedResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


