// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansAggregateResponse The response object for the spans aggregate API endpoint.
type SpansAggregateResponse struct {
	// The list of matching buckets, one item per bucket.
	Data []SpansAggregateBucket `json:"data,omitempty"`
	// The metadata associated with a request.
	Meta *SpansAggregateResponseMetadata `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSpansAggregateResponse instantiates a new SpansAggregateResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSpansAggregateResponse() *SpansAggregateResponse {
	this := SpansAggregateResponse{}
	return &this
}

// NewSpansAggregateResponseWithDefaults instantiates a new SpansAggregateResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSpansAggregateResponseWithDefaults() *SpansAggregateResponse {
	this := SpansAggregateResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SpansAggregateResponse) GetData() []SpansAggregateBucket {
	if o == nil || o.Data == nil {
		var ret []SpansAggregateBucket
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAggregateResponse) GetDataOk() (*[]SpansAggregateBucket, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SpansAggregateResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []SpansAggregateBucket and assigns it to the Data field.
func (o *SpansAggregateResponse) SetData(v []SpansAggregateBucket) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SpansAggregateResponse) GetMeta() SpansAggregateResponseMetadata {
	if o == nil || o.Meta == nil {
		var ret SpansAggregateResponseMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAggregateResponse) GetMetaOk() (*SpansAggregateResponseMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SpansAggregateResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given SpansAggregateResponseMetadata and assigns it to the Meta field.
func (o *SpansAggregateResponse) SetMeta(v SpansAggregateResponseMetadata) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SpansAggregateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SpansAggregateResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data []SpansAggregateBucket          `json:"data,omitempty"`
		Meta *SpansAggregateResponseMetadata `json:"meta,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = all.Data
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
