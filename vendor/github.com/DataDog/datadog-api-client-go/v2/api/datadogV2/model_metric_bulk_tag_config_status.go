// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricBulkTagConfigStatus The status of a request to bulk configure metric tags.
// It contains the fields from the original request for reference.
type MetricBulkTagConfigStatus struct {
	// Optional attributes for the status of a bulk tag configuration request.
	Attributes *MetricBulkTagConfigStatusAttributes `json:"attributes,omitempty"`
	// A text prefix to match against metric names.
	Id string `json:"id"`
	// The metric bulk configure tags resource.
	Type MetricBulkConfigureTagsType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewMetricBulkTagConfigStatus instantiates a new MetricBulkTagConfigStatus object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricBulkTagConfigStatus(id string, typeVar MetricBulkConfigureTagsType) *MetricBulkTagConfigStatus {
	this := MetricBulkTagConfigStatus{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewMetricBulkTagConfigStatusWithDefaults instantiates a new MetricBulkTagConfigStatus object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricBulkTagConfigStatusWithDefaults() *MetricBulkTagConfigStatus {
	this := MetricBulkTagConfigStatus{}
	var typeVar MetricBulkConfigureTagsType = METRICBULKCONFIGURETAGSTYPE_BULK_MANAGE_TAGS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MetricBulkTagConfigStatus) GetAttributes() MetricBulkTagConfigStatusAttributes {
	if o == nil || o.Attributes == nil {
		var ret MetricBulkTagConfigStatusAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricBulkTagConfigStatus) GetAttributesOk() (*MetricBulkTagConfigStatusAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MetricBulkTagConfigStatus) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given MetricBulkTagConfigStatusAttributes and assigns it to the Attributes field.
func (o *MetricBulkTagConfigStatus) SetAttributes(v MetricBulkTagConfigStatusAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *MetricBulkTagConfigStatus) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MetricBulkTagConfigStatus) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *MetricBulkTagConfigStatus) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *MetricBulkTagConfigStatus) GetType() MetricBulkConfigureTagsType {
	if o == nil {
		var ret MetricBulkConfigureTagsType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MetricBulkTagConfigStatus) GetTypeOk() (*MetricBulkConfigureTagsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *MetricBulkTagConfigStatus) SetType(v MetricBulkConfigureTagsType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricBulkTagConfigStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricBulkTagConfigStatus) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *MetricBulkTagConfigStatusAttributes `json:"attributes,omitempty"`
		Id         *string                              `json:"id"`
		Type       *MetricBulkConfigureTagsType         `json:"type"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
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
	o.Id = *all.Id
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
