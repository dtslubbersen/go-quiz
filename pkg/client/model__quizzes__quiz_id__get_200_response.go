/*
go-quiz

This is the API documentation for go-quiz, a simple Quiz API allowing users to obtain quizzes, answer the questions and see their results compared to other users.

API version: 1.0
Contact: dtslubbersen@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the QuizzesQuizIdGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuizzesQuizIdGet200Response{}

// QuizzesQuizIdGet200Response struct for QuizzesQuizIdGet200Response
type QuizzesQuizIdGet200Response struct {
	Data       *StoreQuiz `json:"data,omitempty"`
	Error      *string    `json:"error,omitempty"`
	StatusCode *int32     `json:"status_code,omitempty"`
}

// NewQuizzesQuizIdGet200Response instantiates a new QuizzesQuizIdGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuizzesQuizIdGet200Response() *QuizzesQuizIdGet200Response {
	this := QuizzesQuizIdGet200Response{}
	return &this
}

// NewQuizzesQuizIdGet200ResponseWithDefaults instantiates a new QuizzesQuizIdGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuizzesQuizIdGet200ResponseWithDefaults() *QuizzesQuizIdGet200Response {
	this := QuizzesQuizIdGet200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *QuizzesQuizIdGet200Response) GetData() StoreQuiz {
	if o == nil || IsNil(o.Data) {
		var ret StoreQuiz
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuizzesQuizIdGet200Response) GetDataOk() (*StoreQuiz, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *QuizzesQuizIdGet200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given StoreQuiz and assigns it to the Data field.
func (o *QuizzesQuizIdGet200Response) SetData(v StoreQuiz) {
	o.Data = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *QuizzesQuizIdGet200Response) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuizzesQuizIdGet200Response) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *QuizzesQuizIdGet200Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *QuizzesQuizIdGet200Response) SetError(v string) {
	o.Error = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *QuizzesQuizIdGet200Response) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuizzesQuizIdGet200Response) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *QuizzesQuizIdGet200Response) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *QuizzesQuizIdGet200Response) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o QuizzesQuizIdGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuizzesQuizIdGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.StatusCode) {
		toSerialize["status_code"] = o.StatusCode
	}
	return toSerialize, nil
}

type NullableQuizzesQuizIdGet200Response struct {
	value *QuizzesQuizIdGet200Response
	isSet bool
}

func (v NullableQuizzesQuizIdGet200Response) Get() *QuizzesQuizIdGet200Response {
	return v.value
}

func (v *NullableQuizzesQuizIdGet200Response) Set(val *QuizzesQuizIdGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableQuizzesQuizIdGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableQuizzesQuizIdGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuizzesQuizIdGet200Response(val *QuizzesQuizIdGet200Response) *NullableQuizzesQuizIdGet200Response {
	return &NullableQuizzesQuizIdGet200Response{value: val, isSet: true}
}

func (v NullableQuizzesQuizIdGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuizzesQuizIdGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
