/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.13
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InternalProvisionMockSubscription Internal Provision Mock Subscription Request Body
type InternalProvisionMockSubscription struct {
	// Identity ID
	IdentityId string `json:"identity_id"`
	// Plan or Price
	PlanOrPrice string `json:"plan_or_price"`
}

// NewInternalProvisionMockSubscription instantiates a new InternalProvisionMockSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalProvisionMockSubscription(identityId string, planOrPrice string) *InternalProvisionMockSubscription {
	this := InternalProvisionMockSubscription{}
	this.IdentityId = identityId
	this.PlanOrPrice = planOrPrice
	return &this
}

// NewInternalProvisionMockSubscriptionWithDefaults instantiates a new InternalProvisionMockSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalProvisionMockSubscriptionWithDefaults() *InternalProvisionMockSubscription {
	this := InternalProvisionMockSubscription{}
	return &this
}

// GetIdentityId returns the IdentityId field value
func (o *InternalProvisionMockSubscription) GetIdentityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value
// and a boolean to check if the value has been set.
func (o *InternalProvisionMockSubscription) GetIdentityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityId, true
}

// SetIdentityId sets field value
func (o *InternalProvisionMockSubscription) SetIdentityId(v string) {
	o.IdentityId = v
}

// GetPlanOrPrice returns the PlanOrPrice field value
func (o *InternalProvisionMockSubscription) GetPlanOrPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanOrPrice
}

// GetPlanOrPriceOk returns a tuple with the PlanOrPrice field value
// and a boolean to check if the value has been set.
func (o *InternalProvisionMockSubscription) GetPlanOrPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanOrPrice, true
}

// SetPlanOrPrice sets field value
func (o *InternalProvisionMockSubscription) SetPlanOrPrice(v string) {
	o.PlanOrPrice = v
}

func (o InternalProvisionMockSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["identity_id"] = o.IdentityId
	}
	if true {
		toSerialize["plan_or_price"] = o.PlanOrPrice
	}
	return json.Marshal(toSerialize)
}

type NullableInternalProvisionMockSubscription struct {
	value *InternalProvisionMockSubscription
	isSet bool
}

func (v NullableInternalProvisionMockSubscription) Get() *InternalProvisionMockSubscription {
	return v.value
}

func (v *NullableInternalProvisionMockSubscription) Set(val *InternalProvisionMockSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalProvisionMockSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalProvisionMockSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalProvisionMockSubscription(val *InternalProvisionMockSubscription) *NullableInternalProvisionMockSubscription {
	return &NullableInternalProvisionMockSubscription{value: val, isSet: true}
}

func (v NullableInternalProvisionMockSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalProvisionMockSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


