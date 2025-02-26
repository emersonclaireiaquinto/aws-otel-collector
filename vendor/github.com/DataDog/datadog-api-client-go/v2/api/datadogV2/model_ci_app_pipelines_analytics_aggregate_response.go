// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppPipelinesAnalyticsAggregateResponse The response object for the pipeline events aggregate API endpoint.
type CIAppPipelinesAnalyticsAggregateResponse struct {
	// The query results.
	Data *CIAppPipelinesAggregationBucketsResponse `json:"data,omitempty"`
	// Links attributes.
	Links *CIAppResponseLinks `json:"links,omitempty"`
	// The metadata associated with a request.
	Meta *CIAppResponseMetadata `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewCIAppPipelinesAnalyticsAggregateResponse instantiates a new CIAppPipelinesAnalyticsAggregateResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCIAppPipelinesAnalyticsAggregateResponse() *CIAppPipelinesAnalyticsAggregateResponse {
	this := CIAppPipelinesAnalyticsAggregateResponse{}
	return &this
}

// NewCIAppPipelinesAnalyticsAggregateResponseWithDefaults instantiates a new CIAppPipelinesAnalyticsAggregateResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCIAppPipelinesAnalyticsAggregateResponseWithDefaults() *CIAppPipelinesAnalyticsAggregateResponse {
	this := CIAppPipelinesAnalyticsAggregateResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CIAppPipelinesAnalyticsAggregateResponse) GetData() CIAppPipelinesAggregationBucketsResponse {
	if o == nil || o.Data == nil {
		var ret CIAppPipelinesAggregationBucketsResponse
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppPipelinesAnalyticsAggregateResponse) GetDataOk() (*CIAppPipelinesAggregationBucketsResponse, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CIAppPipelinesAnalyticsAggregateResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given CIAppPipelinesAggregationBucketsResponse and assigns it to the Data field.
func (o *CIAppPipelinesAnalyticsAggregateResponse) SetData(v CIAppPipelinesAggregationBucketsResponse) {
	o.Data = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CIAppPipelinesAnalyticsAggregateResponse) GetLinks() CIAppResponseLinks {
	if o == nil || o.Links == nil {
		var ret CIAppResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppPipelinesAnalyticsAggregateResponse) GetLinksOk() (*CIAppResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CIAppPipelinesAnalyticsAggregateResponse) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given CIAppResponseLinks and assigns it to the Links field.
func (o *CIAppPipelinesAnalyticsAggregateResponse) SetLinks(v CIAppResponseLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CIAppPipelinesAnalyticsAggregateResponse) GetMeta() CIAppResponseMetadata {
	if o == nil || o.Meta == nil {
		var ret CIAppResponseMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppPipelinesAnalyticsAggregateResponse) GetMetaOk() (*CIAppResponseMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CIAppPipelinesAnalyticsAggregateResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given CIAppResponseMetadata and assigns it to the Meta field.
func (o *CIAppPipelinesAnalyticsAggregateResponse) SetMeta(v CIAppResponseMetadata) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CIAppPipelinesAnalyticsAggregateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
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
func (o *CIAppPipelinesAnalyticsAggregateResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data  *CIAppPipelinesAggregationBucketsResponse `json:"data,omitempty"`
		Links *CIAppResponseLinks                       `json:"links,omitempty"`
		Meta  *CIAppResponseMetadata                    `json:"meta,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "links", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = all.Data
	if all.Links != nil && all.Links.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Links = all.Links
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
