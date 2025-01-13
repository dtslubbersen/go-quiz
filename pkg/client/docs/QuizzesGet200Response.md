# QuizzesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]StoreQuiz**](StoreQuiz.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewQuizzesGet200Response

`func NewQuizzesGet200Response() *QuizzesGet200Response`

NewQuizzesGet200Response instantiates a new QuizzesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuizzesGet200ResponseWithDefaults

`func NewQuizzesGet200ResponseWithDefaults() *QuizzesGet200Response`

NewQuizzesGet200ResponseWithDefaults instantiates a new QuizzesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *QuizzesGet200Response) GetData() []StoreQuiz`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *QuizzesGet200Response) GetDataOk() (*[]StoreQuiz, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *QuizzesGet200Response) SetData(v []StoreQuiz)`

SetData sets Data field to given value.

### HasData

`func (o *QuizzesGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *QuizzesGet200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *QuizzesGet200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *QuizzesGet200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *QuizzesGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


