// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppTestsAggregationBucketsResponse The query results.
type CIAppTestsAggregationBucketsResponse struct {
	// The list of matching buckets, one item per bucket.
	Buckets []CIAppTestsBucketResponse `json:"buckets,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewCIAppTestsAggregationBucketsResponse instantiates a new CIAppTestsAggregationBucketsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCIAppTestsAggregationBucketsResponse() *CIAppTestsAggregationBucketsResponse {
	this := CIAppTestsAggregationBucketsResponse{}
	return &this
}

// NewCIAppTestsAggregationBucketsResponseWithDefaults instantiates a new CIAppTestsAggregationBucketsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCIAppTestsAggregationBucketsResponseWithDefaults() *CIAppTestsAggregationBucketsResponse {
	this := CIAppTestsAggregationBucketsResponse{}
	return &this
}

// GetBuckets returns the Buckets field value if set, zero value otherwise.
func (o *CIAppTestsAggregationBucketsResponse) GetBuckets() []CIAppTestsBucketResponse {
	if o == nil || o.Buckets == nil {
		var ret []CIAppTestsBucketResponse
		return ret
	}
	return o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppTestsAggregationBucketsResponse) GetBucketsOk() (*[]CIAppTestsBucketResponse, bool) {
	if o == nil || o.Buckets == nil {
		return nil, false
	}
	return &o.Buckets, true
}

// HasBuckets returns a boolean if a field has been set.
func (o *CIAppTestsAggregationBucketsResponse) HasBuckets() bool {
	return o != nil && o.Buckets != nil
}

// SetBuckets gets a reference to the given []CIAppTestsBucketResponse and assigns it to the Buckets field.
func (o *CIAppTestsAggregationBucketsResponse) SetBuckets(v []CIAppTestsBucketResponse) {
	o.Buckets = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CIAppTestsAggregationBucketsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Buckets != nil {
		toSerialize["buckets"] = o.Buckets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CIAppTestsAggregationBucketsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Buckets []CIAppTestsBucketResponse `json:"buckets,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"buckets"})
	} else {
		return err
	}
	o.Buckets = all.Buckets

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
