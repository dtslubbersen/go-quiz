# AuthTokenPost400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthTokenPost400Response

`func NewAuthTokenPost400Response() *AuthTokenPost400Response`

NewAuthTokenPost400Response instantiates a new AuthTokenPost400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthTokenPost400ResponseWithDefaults

`func NewAuthTokenPost400ResponseWithDefaults() *AuthTokenPost400Response`

NewAuthTokenPost400ResponseWithDefaults instantiates a new AuthTokenPost400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AuthTokenPost400Response) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AuthTokenPost400Response) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AuthTokenPost400Response) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *AuthTokenPost400Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *AuthTokenPost400Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AuthTokenPost400Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AuthTokenPost400Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AuthTokenPost400Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


