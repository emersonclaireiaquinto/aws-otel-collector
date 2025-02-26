// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorGroupSearchResponseCounts The counts of monitor groups per different criteria.
type MonitorGroupSearchResponseCounts struct {
	// Search facets.
	Status []MonitorSearchCountItem `json:"status,omitempty"`
	// Search facets.
	Type []MonitorSearchCountItem `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewMonitorGroupSearchResponseCounts instantiates a new MonitorGroupSearchResponseCounts object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorGroupSearchResponseCounts() *MonitorGroupSearchResponseCounts {
	this := MonitorGroupSearchResponseCounts{}
	return &this
}

// NewMonitorGroupSearchResponseCountsWithDefaults instantiates a new MonitorGroupSearchResponseCounts object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorGroupSearchResponseCountsWithDefaults() *MonitorGroupSearchResponseCounts {
	this := MonitorGroupSearchResponseCounts{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MonitorGroupSearchResponseCounts) GetStatus() []MonitorSearchCountItem {
	if o == nil || o.Status == nil {
		var ret []MonitorSearchCountItem
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorGroupSearchResponseCounts) GetStatusOk() (*[]MonitorSearchCountItem, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MonitorGroupSearchResponseCounts) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given []MonitorSearchCountItem and assigns it to the Status field.
func (o *MonitorGroupSearchResponseCounts) SetStatus(v []MonitorSearchCountItem) {
	o.Status = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MonitorGroupSearchResponseCounts) GetType() []MonitorSearchCountItem {
	if o == nil || o.Type == nil {
		var ret []MonitorSearchCountItem
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorGroupSearchResponseCounts) GetTypeOk() (*[]MonitorSearchCountItem, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MonitorGroupSearchResponseCounts) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given []MonitorSearchCountItem and assigns it to the Type field.
func (o *MonitorGroupSearchResponseCounts) SetType(v []MonitorSearchCountItem) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorGroupSearchResponseCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
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
func (o *MonitorGroupSearchResponseCounts) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status []MonitorSearchCountItem `json:"status,omitempty"`
		Type   []MonitorSearchCountItem `json:"type,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"status", "type"})
	} else {
		return err
	}
	o.Status = all.Status
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
