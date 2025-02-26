// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RestrictionPolicyAttributes Restriction policy attributes.
type RestrictionPolicyAttributes struct {
	// An array of bindings.
	Bindings []RestrictionPolicyBinding `json:"bindings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewRestrictionPolicyAttributes instantiates a new RestrictionPolicyAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRestrictionPolicyAttributes(bindings []RestrictionPolicyBinding) *RestrictionPolicyAttributes {
	this := RestrictionPolicyAttributes{}
	this.Bindings = bindings
	return &this
}

// NewRestrictionPolicyAttributesWithDefaults instantiates a new RestrictionPolicyAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRestrictionPolicyAttributesWithDefaults() *RestrictionPolicyAttributes {
	this := RestrictionPolicyAttributes{}
	return &this
}

// GetBindings returns the Bindings field value.
func (o *RestrictionPolicyAttributes) GetBindings() []RestrictionPolicyBinding {
	if o == nil {
		var ret []RestrictionPolicyBinding
		return ret
	}
	return o.Bindings
}

// GetBindingsOk returns a tuple with the Bindings field value
// and a boolean to check if the value has been set.
func (o *RestrictionPolicyAttributes) GetBindingsOk() (*[]RestrictionPolicyBinding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bindings, true
}

// SetBindings sets field value.
func (o *RestrictionPolicyAttributes) SetBindings(v []RestrictionPolicyBinding) {
	o.Bindings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RestrictionPolicyAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["bindings"] = o.Bindings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RestrictionPolicyAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Bindings *[]RestrictionPolicyBinding `json:"bindings"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Bindings == nil {
		return fmt.Errorf("required field bindings missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bindings"})
	} else {
		return err
	}
	o.Bindings = *all.Bindings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
