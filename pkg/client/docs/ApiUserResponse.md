# ApiUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiUserResponse

`func NewApiUserResponse() *ApiUserResponse`

NewApiUserResponse instantiates a new ApiUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUserResponseWithDefaults

`func NewApiUserResponseWithDefaults() *ApiUserResponse`

NewApiUserResponseWithDefaults instantiates a new ApiUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ApiUserResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ApiUserResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ApiUserResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ApiUserResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *ApiUserResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiUserResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiUserResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ApiUserResponse) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


