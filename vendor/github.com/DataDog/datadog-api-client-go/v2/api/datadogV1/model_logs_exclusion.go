// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsExclusion Represents the index exclusion filter object from configuration API.
type LogsExclusion struct {
	// Exclusion filter is defined by a query, a sampling rule, and a active/inactive toggle.
	Filter *LogsExclusionFilter `json:"filter,omitempty"`
	// Whether or not the exclusion filter is active.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Name of the index exclusion filter.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewLogsExclusion instantiates a new LogsExclusion object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsExclusion(name string) *LogsExclusion {
	this := LogsExclusion{}
	this.Name = name
	return &this
}

// NewLogsExclusionWithDefaults instantiates a new LogsExclusion object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsExclusionWithDefaults() *LogsExclusion {
	this := LogsExclusion{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *LogsExclusion) GetFilter() LogsExclusionFilter {
	if o == nil || o.Filter == nil {
		var ret LogsExclusionFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsExclusion) GetFilterOk() (*LogsExclusionFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *LogsExclusion) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given LogsExclusionFilter and assigns it to the Filter field.
func (o *LogsExclusion) SetFilter(v LogsExclusionFilter) {
	o.Filter = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *LogsExclusion) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsExclusion) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *LogsExclusion) HasIsEnabled() bool {
	return o != nil && o.IsEnabled != nil
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *LogsExclusion) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetName returns the Name field value.
func (o *LogsExclusion) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LogsExclusion) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *LogsExclusion) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsExclusion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsExclusion) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter    *LogsExclusionFilter `json:"filter,omitempty"`
		IsEnabled *bool                `json:"is_enabled,omitempty"`
		Name      *string              `json:"name"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filter", "is_enabled", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Filter != nil && all.Filter.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Filter = all.Filter
	o.IsEnabled = all.IsEnabled
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
