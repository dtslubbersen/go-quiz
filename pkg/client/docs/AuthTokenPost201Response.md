# AuthTokenPost201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ApiTokenCreatedResponse**](ApiTokenCreatedResponse.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 

## Methods

### NewAuthTokenPost201Response

`func NewAuthTokenPost201Response() *AuthTokenPost201Response`

NewAuthTokenPost201Response instantiates a new AuthTokenPost201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthTokenPost201ResponseWithDefaults

`func NewAuthTokenPost201ResponseWithDefaults() *AuthTokenPost201Response`

NewAuthTokenPost201ResponseWithDefaults instantiates a new AuthTokenPost201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AuthTokenPost201Response) GetData() ApiTokenCreatedResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AuthTokenPost201Response) GetDataOk() (*ApiTokenCreatedResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AuthTokenPost201Response) SetData(v ApiTokenCreatedResponse)`

SetData sets Data field to given value.

### HasData

`func (o *AuthTokenPost201Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *AuthTokenPost201Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AuthTokenPost201Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AuthTokenPost201Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AuthTokenPost201Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStatusCode

`func (o *AuthTokenPost201Response) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *AuthTokenPost201Response) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *AuthTokenPost201Response) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *AuthTokenPost201Response) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


