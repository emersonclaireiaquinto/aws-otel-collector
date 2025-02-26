// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeMonitorIncludedItem Information about the monitor identified by the downtime.
type DowntimeMonitorIncludedItem struct {
	// Attributes of the monitor identified by the downtime.
	Attributes *DowntimeMonitorIncludedAttributes `json:"attributes,omitempty"`
	// ID of the monitor identified by the downtime.
	Id *int64 `json:"id,omitempty"`
	// Monitor resource type.
	Type *DowntimeIncludedMonitorType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewDowntimeMonitorIncludedItem instantiates a new DowntimeMonitorIncludedItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDowntimeMonitorIncludedItem() *DowntimeMonitorIncludedItem {
	this := DowntimeMonitorIncludedItem{}
	var typeVar DowntimeIncludedMonitorType = DOWNTIMEINCLUDEDMONITORTYPE_MONITORS
	this.Type = &typeVar
	return &this
}

// NewDowntimeMonitorIncludedItemWithDefaults instantiates a new DowntimeMonitorIncludedItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDowntimeMonitorIncludedItemWithDefaults() *DowntimeMonitorIncludedItem {
	this := DowntimeMonitorIncludedItem{}
	var typeVar DowntimeIncludedMonitorType = DOWNTIMEINCLUDEDMONITORTYPE_MONITORS
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *DowntimeMonitorIncludedItem) GetAttributes() DowntimeMonitorIncludedAttributes {
	if o == nil || o.Attributes == nil {
		var ret DowntimeMonitorIncludedAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeMonitorIncludedItem) GetAttributesOk() (*DowntimeMonitorIncludedAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DowntimeMonitorIncludedItem) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given DowntimeMonitorIncludedAttributes and assigns it to the Attributes field.
func (o *DowntimeMonitorIncludedItem) SetAttributes(v DowntimeMonitorIncludedAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DowntimeMonitorIncludedItem) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeMonitorIncludedItem) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DowntimeMonitorIncludedItem) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DowntimeMonitorIncludedItem) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DowntimeMonitorIncludedItem) GetType() DowntimeIncludedMonitorType {
	if o == nil || o.Type == nil {
		var ret DowntimeIncludedMonitorType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeMonitorIncludedItem) GetTypeOk() (*DowntimeIncludedMonitorType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DowntimeMonitorIncludedItem) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given DowntimeIncludedMonitorType and assigns it to the Type field.
func (o *DowntimeMonitorIncludedItem) SetType(v DowntimeIncludedMonitorType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DowntimeMonitorIncludedItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DowntimeMonitorIncludedItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *DowntimeMonitorIncludedAttributes `json:"attributes,omitempty"`
		Id         *int64                             `json:"id,omitempty"`
		Type       *DowntimeIncludedMonitorType       `json:"type,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
