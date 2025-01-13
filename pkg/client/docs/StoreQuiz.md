# StoreQuiz

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Performance** | Pointer to [**StorePerformance**](StorePerformance.md) |  | [optional] 
**Questions** | Pointer to [**[]StoreQuestion**](StoreQuestion.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewStoreQuiz

`func NewStoreQuiz() *StoreQuiz`

NewStoreQuiz instantiates a new StoreQuiz object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreQuizWithDefaults

`func NewStoreQuizWithDefaults() *StoreQuiz`

NewStoreQuizWithDefaults instantiates a new StoreQuiz object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *StoreQuiz) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StoreQuiz) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StoreQuiz) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StoreQuiz) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *StoreQuiz) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StoreQuiz) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StoreQuiz) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StoreQuiz) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPerformance

`func (o *StoreQuiz) GetPerformance() StorePerformance`

GetPerformance returns the Performance field if non-nil, zero value otherwise.

### GetPerformanceOk

`func (o *StoreQuiz) GetPerformanceOk() (*StorePerformance, bool)`

GetPerformanceOk returns a tuple with the Performance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformance

`func (o *StoreQuiz) SetPerformance(v StorePerformance)`

SetPerformance sets Performance field to given value.

### HasPerformance

`func (o *StoreQuiz) HasPerformance() bool`

HasPerformance returns a boolean if a field has been set.

### GetQuestions

`func (o *StoreQuiz) GetQuestions() []StoreQuestion`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *StoreQuiz) GetQuestionsOk() (*[]StoreQuestion, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *StoreQuiz) SetQuestions(v []StoreQuestion)`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *StoreQuiz) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.

### GetTitle

`func (o *StoreQuiz) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StoreQuiz) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StoreQuiz) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *StoreQuiz) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


