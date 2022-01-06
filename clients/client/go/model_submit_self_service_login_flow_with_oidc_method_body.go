/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.39
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SubmitSelfServiceLoginFlowWithOidcMethodBody SubmitSelfServiceLoginFlowWithOidcMethodBody is used to decode the login form payload when using the oidc method.
type SubmitSelfServiceLoginFlowWithOidcMethodBody struct {
	// The CSRF Token
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Method to use  This field must be set to `oidc` when using the oidc method.
	Method string `json:"method"`
	// The provider to register with
	Provider string `json:"provider"`
	// The identity traits. This is a placeholder for the registration flow.
	Traits map[string]interface{} `json:"traits,omitempty"`
}

// NewSubmitSelfServiceLoginFlowWithOidcMethodBody instantiates a new SubmitSelfServiceLoginFlowWithOidcMethodBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitSelfServiceLoginFlowWithOidcMethodBody(method string, provider string) *SubmitSelfServiceLoginFlowWithOidcMethodBody {
	this := SubmitSelfServiceLoginFlowWithOidcMethodBody{}
	this.Method = method
	this.Provider = provider
	return &this
}

// NewSubmitSelfServiceLoginFlowWithOidcMethodBodyWithDefaults instantiates a new SubmitSelfServiceLoginFlowWithOidcMethodBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitSelfServiceLoginFlowWithOidcMethodBodyWithDefaults() *SubmitSelfServiceLoginFlowWithOidcMethodBody {
	this := SubmitSelfServiceLoginFlowWithOidcMethodBody{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *SubmitSelfServiceLoginFlowWithOidcMethodBody) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceLoginFlowWithOidcMethodBody) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *SubmitSelfServiceLoginFlowWithOidcMethodBody) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *SubmitSelfServiceLoginFlowWithOidcMethodBody) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetMethod returns the Method field value
func (o *SubmitSelfServiceLoginFlowWithOidcMethodBody) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceLoginFlowWithOidcMethodBody) GetMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *SubmitSelfServiceLoginFlowWithOidcMethodBody) SetMethod(v string) {
	o.Method = v
}

// GetProvider returns the Provider field value
func (o *SubmitSelfServiceLoginFlowWithOidcMethodBody) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceLoginFlowWithOidcMethodBody) GetProviderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *SubmitSelfServiceLoginFlowWithOidcMethodBody) SetProvider(v string) {
	o.Provider = v
}

// GetTraits returns the Traits field value if set, zero value otherwise.
func (o *SubmitSelfServiceLoginFlowWithOidcMethodBody) GetTraits() map[string]interface{} {
	if o == nil || o.Traits == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceLoginFlowWithOidcMethodBody) GetTraitsOk() (map[string]interface{}, bool) {
	if o == nil || o.Traits == nil {
		return nil, false
	}
	return o.Traits, true
}

// HasTraits returns a boolean if a field has been set.
func (o *SubmitSelfServiceLoginFlowWithOidcMethodBody) HasTraits() bool {
	if o != nil && o.Traits != nil {
		return true
	}

	return false
}

// SetTraits gets a reference to the given map[string]interface{} and assigns it to the Traits field.
func (o *SubmitSelfServiceLoginFlowWithOidcMethodBody) SetTraits(v map[string]interface{}) {
	o.Traits = v
}

func (o SubmitSelfServiceLoginFlowWithOidcMethodBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if o.Traits != nil {
		toSerialize["traits"] = o.Traits
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitSelfServiceLoginFlowWithOidcMethodBody struct {
	value *SubmitSelfServiceLoginFlowWithOidcMethodBody
	isSet bool
}

func (v NullableSubmitSelfServiceLoginFlowWithOidcMethodBody) Get() *SubmitSelfServiceLoginFlowWithOidcMethodBody {
	return v.value
}

func (v *NullableSubmitSelfServiceLoginFlowWithOidcMethodBody) Set(val *SubmitSelfServiceLoginFlowWithOidcMethodBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitSelfServiceLoginFlowWithOidcMethodBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitSelfServiceLoginFlowWithOidcMethodBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitSelfServiceLoginFlowWithOidcMethodBody(val *SubmitSelfServiceLoginFlowWithOidcMethodBody) *NullableSubmitSelfServiceLoginFlowWithOidcMethodBody {
	return &NullableSubmitSelfServiceLoginFlowWithOidcMethodBody{value: val, isSet: true}
}

func (v NullableSubmitSelfServiceLoginFlowWithOidcMethodBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitSelfServiceLoginFlowWithOidcMethodBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


