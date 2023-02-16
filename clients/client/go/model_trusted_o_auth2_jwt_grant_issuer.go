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
	"time"
)

// TrustedOAuth2JwtGrantIssuer OAuth2 JWT Bearer Grant Type Issuer Trust Relationship
type TrustedOAuth2JwtGrantIssuer struct {
	// The \"allow_any_subject\" indicates that the issuer is allowed to have any principal as the subject of the JWT.
	AllowAnySubject *bool `json:"allow_any_subject,omitempty"`
	// The \"created_at\" indicates, when grant was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The \"expires_at\" indicates, when grant will expire, so we will reject assertion from \"issuer\" targeting \"subject\".
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	Id *string `json:"id,omitempty"`
	// The \"issuer\" identifies the principal that issued the JWT assertion (same as \"iss\" claim in JWT).
	Issuer *string `json:"issuer,omitempty"`
	PublicKey *TrustedOAuth2JwtGrantJsonWebKey `json:"public_key,omitempty"`
	// The \"scope\" contains list of scope values (as described in Section 3.3 of OAuth 2.0 [RFC6749])
	Scope []string `json:"scope,omitempty"`
	// The \"subject\" identifies the principal that is the subject of the JWT.
	Subject *string `json:"subject,omitempty"`
}

// NewTrustedOAuth2JwtGrantIssuer instantiates a new TrustedOAuth2JwtGrantIssuer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrustedOAuth2JwtGrantIssuer() *TrustedOAuth2JwtGrantIssuer {
	this := TrustedOAuth2JwtGrantIssuer{}
	return &this
}

// NewTrustedOAuth2JwtGrantIssuerWithDefaults instantiates a new TrustedOAuth2JwtGrantIssuer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrustedOAuth2JwtGrantIssuerWithDefaults() *TrustedOAuth2JwtGrantIssuer {
	this := TrustedOAuth2JwtGrantIssuer{}
	return &this
}

// GetAllowAnySubject returns the AllowAnySubject field value if set, zero value otherwise.
func (o *TrustedOAuth2JwtGrantIssuer) GetAllowAnySubject() bool {
	if o == nil || o.AllowAnySubject == nil {
		var ret bool
		return ret
	}
	return *o.AllowAnySubject
}

// GetAllowAnySubjectOk returns a tuple with the AllowAnySubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOAuth2JwtGrantIssuer) GetAllowAnySubjectOk() (*bool, bool) {
	if o == nil || o.AllowAnySubject == nil {
		return nil, false
	}
	return o.AllowAnySubject, true
}

// HasAllowAnySubject returns a boolean if a field has been set.
func (o *TrustedOAuth2JwtGrantIssuer) HasAllowAnySubject() bool {
	if o != nil && o.AllowAnySubject != nil {
		return true
	}

	return false
}

// SetAllowAnySubject gets a reference to the given bool and assigns it to the AllowAnySubject field.
func (o *TrustedOAuth2JwtGrantIssuer) SetAllowAnySubject(v bool) {
	o.AllowAnySubject = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TrustedOAuth2JwtGrantIssuer) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOAuth2JwtGrantIssuer) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TrustedOAuth2JwtGrantIssuer) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TrustedOAuth2JwtGrantIssuer) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *TrustedOAuth2JwtGrantIssuer) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOAuth2JwtGrantIssuer) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *TrustedOAuth2JwtGrantIssuer) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *TrustedOAuth2JwtGrantIssuer) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TrustedOAuth2JwtGrantIssuer) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOAuth2JwtGrantIssuer) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TrustedOAuth2JwtGrantIssuer) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TrustedOAuth2JwtGrantIssuer) SetId(v string) {
	o.Id = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *TrustedOAuth2JwtGrantIssuer) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOAuth2JwtGrantIssuer) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *TrustedOAuth2JwtGrantIssuer) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *TrustedOAuth2JwtGrantIssuer) SetIssuer(v string) {
	o.Issuer = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *TrustedOAuth2JwtGrantIssuer) GetPublicKey() TrustedOAuth2JwtGrantJsonWebKey {
	if o == nil || o.PublicKey == nil {
		var ret TrustedOAuth2JwtGrantJsonWebKey
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOAuth2JwtGrantIssuer) GetPublicKeyOk() (*TrustedOAuth2JwtGrantJsonWebKey, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *TrustedOAuth2JwtGrantIssuer) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given TrustedOAuth2JwtGrantJsonWebKey and assigns it to the PublicKey field.
func (o *TrustedOAuth2JwtGrantIssuer) SetPublicKey(v TrustedOAuth2JwtGrantJsonWebKey) {
	o.PublicKey = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *TrustedOAuth2JwtGrantIssuer) GetScope() []string {
	if o == nil || o.Scope == nil {
		var ret []string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOAuth2JwtGrantIssuer) GetScopeOk() ([]string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *TrustedOAuth2JwtGrantIssuer) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given []string and assigns it to the Scope field.
func (o *TrustedOAuth2JwtGrantIssuer) SetScope(v []string) {
	o.Scope = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *TrustedOAuth2JwtGrantIssuer) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOAuth2JwtGrantIssuer) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *TrustedOAuth2JwtGrantIssuer) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *TrustedOAuth2JwtGrantIssuer) SetSubject(v string) {
	o.Subject = &v
}

func (o TrustedOAuth2JwtGrantIssuer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowAnySubject != nil {
		toSerialize["allow_any_subject"] = o.AllowAnySubject
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.PublicKey != nil {
		toSerialize["public_key"] = o.PublicKey
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	return json.Marshal(toSerialize)
}

type NullableTrustedOAuth2JwtGrantIssuer struct {
	value *TrustedOAuth2JwtGrantIssuer
	isSet bool
}

func (v NullableTrustedOAuth2JwtGrantIssuer) Get() *TrustedOAuth2JwtGrantIssuer {
	return v.value
}

func (v *NullableTrustedOAuth2JwtGrantIssuer) Set(val *TrustedOAuth2JwtGrantIssuer) {
	v.value = val
	v.isSet = true
}

func (v NullableTrustedOAuth2JwtGrantIssuer) IsSet() bool {
	return v.isSet
}

func (v *NullableTrustedOAuth2JwtGrantIssuer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrustedOAuth2JwtGrantIssuer(val *TrustedOAuth2JwtGrantIssuer) *NullableTrustedOAuth2JwtGrantIssuer {
	return &NullableTrustedOAuth2JwtGrantIssuer{value: val, isSet: true}
}

func (v NullableTrustedOAuth2JwtGrantIssuer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrustedOAuth2JwtGrantIssuer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


