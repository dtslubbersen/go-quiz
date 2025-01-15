/*
go-quiz

This is the API documentation for go-quiz, a simple Quiz API allowing users to obtain quizzes, answer the questions and see their results compared to other users.

API version: 1.0
Contact: dtslubbersen@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ApiCreateTokenPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiCreateTokenPayload{}

// ApiCreateTokenPayload struct for ApiCreateTokenPayload
type ApiCreateTokenPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type _ApiCreateTokenPayload ApiCreateTokenPayload

// NewApiCreateTokenPayload instantiates a new ApiCreateTokenPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiCreateTokenPayload(email string, password string) *ApiCreateTokenPayload {
	this := ApiCreateTokenPayload{}
	this.Email = email
	this.Password = password
	return &this
}

// NewApiCreateTokenPayloadWithDefaults instantiates a new ApiCreateTokenPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiCreateTokenPayloadWithDefaults() *ApiCreateTokenPayload {
	this := ApiCreateTokenPayload{}
	return &this
}

// GetEmail returns the Email field value
func (o *ApiCreateTokenPayload) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ApiCreateTokenPayload) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ApiCreateTokenPayload) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *ApiCreateTokenPayload) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ApiCreateTokenPayload) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ApiCreateTokenPayload) SetPassword(v string) {
	o.Password = v
}

func (o ApiCreateTokenPayload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiCreateTokenPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

func (o *ApiCreateTokenPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"password",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApiCreateTokenPayload := _ApiCreateTokenPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiCreateTokenPayload)

	if err != nil {
		return err
	}

	*o = ApiCreateTokenPayload(varApiCreateTokenPayload)

	return err
}

type NullableApiCreateTokenPayload struct {
	value *ApiCreateTokenPayload
	isSet bool
}

func (v NullableApiCreateTokenPayload) Get() *ApiCreateTokenPayload {
	return v.value
}

func (v *NullableApiCreateTokenPayload) Set(val *ApiCreateTokenPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableApiCreateTokenPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableApiCreateTokenPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiCreateTokenPayload(val *ApiCreateTokenPayload) *NullableApiCreateTokenPayload {
	return &NullableApiCreateTokenPayload{value: val, isSet: true}
}

func (v NullableApiCreateTokenPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiCreateTokenPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
