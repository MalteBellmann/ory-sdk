/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.6.1
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the StripeCustomer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StripeCustomer{}

// StripeCustomer struct for StripeCustomer
type StripeCustomer struct {
	Id *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StripeCustomer StripeCustomer

// NewStripeCustomer instantiates a new StripeCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripeCustomer() *StripeCustomer {
	this := StripeCustomer{}
	return &this
}

// NewStripeCustomerWithDefaults instantiates a new StripeCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripeCustomerWithDefaults() *StripeCustomer {
	this := StripeCustomer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StripeCustomer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeCustomer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StripeCustomer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StripeCustomer) SetId(v string) {
	o.Id = &v
}

func (o StripeCustomer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StripeCustomer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StripeCustomer) UnmarshalJSON(bytes []byte) (err error) {
	varStripeCustomer := _StripeCustomer{}

	err = json.Unmarshal(bytes, &varStripeCustomer)

	if err != nil {
		return err
	}

	*o = StripeCustomer(varStripeCustomer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStripeCustomer struct {
	value *StripeCustomer
	isSet bool
}

func (v NullableStripeCustomer) Get() *StripeCustomer {
	return v.value
}

func (v *NullableStripeCustomer) Set(val *StripeCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullableStripeCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableStripeCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripeCustomer(val *StripeCustomer) *NullableStripeCustomer {
	return &NullableStripeCustomer{value: val, isSet: true}
}

func (v NullableStripeCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripeCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


