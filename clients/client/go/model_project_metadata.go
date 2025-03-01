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
	"time"
	"fmt"
)

// checks if the ProjectMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectMetadata{}

// ProjectMetadata struct for ProjectMetadata
type ProjectMetadata struct {
	// The Project's Creation Date
	CreatedAt time.Time `json:"created_at"`
	// The environment of the project. prod Production dev Development
	Environment string `json:"environment"`
	Hosts []string `json:"hosts"`
	// The project's ID.
	Id string `json:"id"`
	// The project's name if set
	Name string `json:"name"`
	// The project's slug
	Slug *string `json:"slug,omitempty"`
	// The state of the project. running Running halted Halted deleted Deleted
	State string `json:"state"`
	SubscriptionId NullableString `json:"subscription_id,omitempty"`
	SubscriptionPlan NullableString `json:"subscription_plan,omitempty"`
	// Last Time Project was Updated
	UpdatedAt time.Time `json:"updated_at"`
	WorkspaceId NullableString `json:"workspace_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProjectMetadata ProjectMetadata

// NewProjectMetadata instantiates a new ProjectMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectMetadata(createdAt time.Time, environment string, hosts []string, id string, name string, state string, updatedAt time.Time) *ProjectMetadata {
	this := ProjectMetadata{}
	this.CreatedAt = createdAt
	this.Environment = environment
	this.Hosts = hosts
	this.Id = id
	this.Name = name
	this.State = state
	this.UpdatedAt = updatedAt
	return &this
}

// NewProjectMetadataWithDefaults instantiates a new ProjectMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectMetadataWithDefaults() *ProjectMetadata {
	this := ProjectMetadata{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *ProjectMetadata) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadata) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ProjectMetadata) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEnvironment returns the Environment field value
func (o *ProjectMetadata) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadata) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *ProjectMetadata) SetEnvironment(v string) {
	o.Environment = v
}

// GetHosts returns the Hosts field value
func (o *ProjectMetadata) GetHosts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadata) GetHostsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hosts, true
}

// SetHosts sets field value
func (o *ProjectMetadata) SetHosts(v []string) {
	o.Hosts = v
}

// GetId returns the Id field value
func (o *ProjectMetadata) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadata) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectMetadata) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ProjectMetadata) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadata) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectMetadata) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *ProjectMetadata) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMetadata) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *ProjectMetadata) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *ProjectMetadata) SetSlug(v string) {
	o.Slug = &v
}

// GetState returns the State field value
func (o *ProjectMetadata) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadata) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ProjectMetadata) SetState(v string) {
	o.State = v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectMetadata) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId.Get()) {
		var ret string
		return ret
	}
	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectMetadata) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *ProjectMetadata) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given NullableString and assigns it to the SubscriptionId field.
func (o *ProjectMetadata) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}
// SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil
func (o *ProjectMetadata) SetSubscriptionIdNil() {
	o.SubscriptionId.Set(nil)
}

// UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
func (o *ProjectMetadata) UnsetSubscriptionId() {
	o.SubscriptionId.Unset()
}

// GetSubscriptionPlan returns the SubscriptionPlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectMetadata) GetSubscriptionPlan() string {
	if o == nil || IsNil(o.SubscriptionPlan.Get()) {
		var ret string
		return ret
	}
	return *o.SubscriptionPlan.Get()
}

// GetSubscriptionPlanOk returns a tuple with the SubscriptionPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectMetadata) GetSubscriptionPlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionPlan.Get(), o.SubscriptionPlan.IsSet()
}

// HasSubscriptionPlan returns a boolean if a field has been set.
func (o *ProjectMetadata) HasSubscriptionPlan() bool {
	if o != nil && o.SubscriptionPlan.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionPlan gets a reference to the given NullableString and assigns it to the SubscriptionPlan field.
func (o *ProjectMetadata) SetSubscriptionPlan(v string) {
	o.SubscriptionPlan.Set(&v)
}
// SetSubscriptionPlanNil sets the value for SubscriptionPlan to be an explicit nil
func (o *ProjectMetadata) SetSubscriptionPlanNil() {
	o.SubscriptionPlan.Set(nil)
}

// UnsetSubscriptionPlan ensures that no value is present for SubscriptionPlan, not even an explicit nil
func (o *ProjectMetadata) UnsetSubscriptionPlan() {
	o.SubscriptionPlan.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ProjectMetadata) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadata) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ProjectMetadata) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectMetadata) GetWorkspaceId() string {
	if o == nil || IsNil(o.WorkspaceId.Get()) {
		var ret string
		return ret
	}
	return *o.WorkspaceId.Get()
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectMetadata) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkspaceId.Get(), o.WorkspaceId.IsSet()
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *ProjectMetadata) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId.IsSet() {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given NullableString and assigns it to the WorkspaceId field.
func (o *ProjectMetadata) SetWorkspaceId(v string) {
	o.WorkspaceId.Set(&v)
}
// SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil
func (o *ProjectMetadata) SetWorkspaceIdNil() {
	o.WorkspaceId.Set(nil)
}

// UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil
func (o *ProjectMetadata) UnsetWorkspaceId() {
	o.WorkspaceId.Unset()
}

func (o ProjectMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["environment"] = o.Environment
	toSerialize["hosts"] = o.Hosts
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	toSerialize["state"] = o.State
	if o.SubscriptionId.IsSet() {
		toSerialize["subscription_id"] = o.SubscriptionId.Get()
	}
	if o.SubscriptionPlan.IsSet() {
		toSerialize["subscription_plan"] = o.SubscriptionPlan.Get()
	}
	toSerialize["updated_at"] = o.UpdatedAt
	if o.WorkspaceId.IsSet() {
		toSerialize["workspace_id"] = o.WorkspaceId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectMetadata) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_at",
		"environment",
		"hosts",
		"id",
		"name",
		"state",
		"updated_at",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProjectMetadata := _ProjectMetadata{}

	err = json.Unmarshal(bytes, &varProjectMetadata)

	if err != nil {
		return err
	}

	*o = ProjectMetadata(varProjectMetadata)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "environment")
		delete(additionalProperties, "hosts")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "state")
		delete(additionalProperties, "subscription_id")
		delete(additionalProperties, "subscription_plan")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "workspace_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectMetadata struct {
	value *ProjectMetadata
	isSet bool
}

func (v NullableProjectMetadata) Get() *ProjectMetadata {
	return v.value
}

func (v *NullableProjectMetadata) Set(val *ProjectMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectMetadata(val *ProjectMetadata) *NullableProjectMetadata {
	return &NullableProjectMetadata{value: val, isSet: true}
}

func (v NullableProjectMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


