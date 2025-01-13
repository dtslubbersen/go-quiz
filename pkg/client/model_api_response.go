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

// checks if the ApiResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiResponse{}

// ApiResponse struct for ApiResponse
type ApiResponse struct {
	Data  map[string]interface{} `json:"data,omitempty"`
	Error *string                `json:"error,omitempty"`
}

// NewApiResponse instantiates a new ApiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiResponse() *ApiResponse {
	this := ApiResponse{}
	return &this
}

// NewApiResponseWithDefaults instantiates a new ApiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiResponseWithDefaults() *ApiResponse {
	this := ApiResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ApiResponse) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ApiResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *ApiResponse) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ApiResponse) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ApiResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ApiResponse) SetError(v string) {
	o.Error = &v
}

func (o ApiResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableApiResponse struct {
	value *ApiResponse
	isSet bool
}

func (v NullableApiResponse) Get() *ApiResponse {
	return v.value
}

func (v *NullableApiResponse) Set(val *ApiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiResponse(val *ApiResponse) *NullableApiResponse {
	return &NullableApiResponse{value: val, isSet: true}
}

func (v NullableApiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
