// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansAggregateRequest The object sent with the request to retrieve a list of aggregated spans from your organization.
type SpansAggregateRequest struct {
	// The object containing the query content.
	Data *SpansAggregateData `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSpansAggregateRequest instantiates a new SpansAggregateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSpansAggregateRequest() *SpansAggregateRequest {
	this := SpansAggregateRequest{}
	return &this
}

// NewSpansAggregateRequestWithDefaults instantiates a new SpansAggregateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSpansAggregateRequestWithDefaults() *SpansAggregateRequest {
	this := SpansAggregateRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SpansAggregateRequest) GetData() SpansAggregateData {
	if o == nil || o.Data == nil {
		var ret SpansAggregateData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAggregateRequest) GetDataOk() (*SpansAggregateData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SpansAggregateRequest) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given SpansAggregateData and assigns it to the Data field.
func (o *SpansAggregateRequest) SetData(v SpansAggregateData) {
	o.Data = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SpansAggregateRequest) MarshalJSON() ([]byte, error) {
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
func (o *SpansAggregateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *SpansAggregateData `json:"data,omitempty"`
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
