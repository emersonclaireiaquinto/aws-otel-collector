// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricEstimateResponse Response object that includes metric cardinality estimates.
type MetricEstimateResponse struct {
	// Object for a metric cardinality estimate.
	Data *MetricEstimate `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewMetricEstimateResponse instantiates a new MetricEstimateResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricEstimateResponse() *MetricEstimateResponse {
	this := MetricEstimateResponse{}
	return &this
}

// NewMetricEstimateResponseWithDefaults instantiates a new MetricEstimateResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricEstimateResponseWithDefaults() *MetricEstimateResponse {
	this := MetricEstimateResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MetricEstimateResponse) GetData() MetricEstimate {
	if o == nil || o.Data == nil {
		var ret MetricEstimate
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricEstimateResponse) GetDataOk() (*MetricEstimate, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MetricEstimateResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given MetricEstimate and assigns it to the Data field.
func (o *MetricEstimateResponse) SetData(v MetricEstimate) {
	o.Data = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricEstimateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricEstimateResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *MetricEstimate `json:"data,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
