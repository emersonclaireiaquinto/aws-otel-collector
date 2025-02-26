// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsMetricUpdateData The new log-based metric properties.
type LogsMetricUpdateData struct {
	// The log-based metric properties that will be updated.
	Attributes LogsMetricUpdateAttributes `json:"attributes"`
	// The type of the resource. The value should always be logs_metrics.
	Type LogsMetricType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewLogsMetricUpdateData instantiates a new LogsMetricUpdateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsMetricUpdateData(attributes LogsMetricUpdateAttributes, typeVar LogsMetricType) *LogsMetricUpdateData {
	this := LogsMetricUpdateData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewLogsMetricUpdateDataWithDefaults instantiates a new LogsMetricUpdateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsMetricUpdateDataWithDefaults() *LogsMetricUpdateData {
	this := LogsMetricUpdateData{}
	var typeVar LogsMetricType = LOGSMETRICTYPE_LOGS_METRICS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *LogsMetricUpdateData) GetAttributes() LogsMetricUpdateAttributes {
	if o == nil {
		var ret LogsMetricUpdateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *LogsMetricUpdateData) GetAttributesOk() (*LogsMetricUpdateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *LogsMetricUpdateData) SetAttributes(v LogsMetricUpdateAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *LogsMetricUpdateData) GetType() LogsMetricType {
	if o == nil {
		var ret LogsMetricType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsMetricUpdateData) GetTypeOk() (*LogsMetricType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LogsMetricUpdateData) SetType(v LogsMetricType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsMetricUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsMetricUpdateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *LogsMetricUpdateAttributes `json:"attributes"`
		Type       *LogsMetricType             `json:"type"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
