// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SearchSLOResponseDataAttributesFacetsObjectString Facet
type SearchSLOResponseDataAttributesFacetsObjectString struct {
	// Count
	Count *int64 `json:"count,omitempty"`
	// Facet
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSearchSLOResponseDataAttributesFacetsObjectString instantiates a new SearchSLOResponseDataAttributesFacetsObjectString object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSearchSLOResponseDataAttributesFacetsObjectString() *SearchSLOResponseDataAttributesFacetsObjectString {
	this := SearchSLOResponseDataAttributesFacetsObjectString{}
	return &this
}

// NewSearchSLOResponseDataAttributesFacetsObjectStringWithDefaults instantiates a new SearchSLOResponseDataAttributesFacetsObjectString object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSearchSLOResponseDataAttributesFacetsObjectStringWithDefaults() *SearchSLOResponseDataAttributesFacetsObjectString {
	this := SearchSLOResponseDataAttributesFacetsObjectString{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SearchSLOResponseDataAttributesFacetsObjectString) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSLOResponseDataAttributesFacetsObjectString) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SearchSLOResponseDataAttributesFacetsObjectString) HasCount() bool {
	return o != nil && o.Count != nil
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *SearchSLOResponseDataAttributesFacetsObjectString) SetCount(v int64) {
	o.Count = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SearchSLOResponseDataAttributesFacetsObjectString) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSLOResponseDataAttributesFacetsObjectString) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SearchSLOResponseDataAttributesFacetsObjectString) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SearchSLOResponseDataAttributesFacetsObjectString) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SearchSLOResponseDataAttributesFacetsObjectString) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SearchSLOResponseDataAttributesFacetsObjectString) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count *int64  `json:"count,omitempty"`
		Name  *string `json:"name,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"count", "name"})
	} else {
		return err
	}
	o.Count = all.Count
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
