// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageCustomReportsResponse Response containing available custom reports.
type UsageCustomReportsResponse struct {
	// An array of available custom reports.
	Data []UsageCustomReportsData `json:"data,omitempty"`
	// The object containing document metadata.
	Meta *UsageCustomReportsMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewUsageCustomReportsResponse instantiates a new UsageCustomReportsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageCustomReportsResponse() *UsageCustomReportsResponse {
	this := UsageCustomReportsResponse{}
	return &this
}

// NewUsageCustomReportsResponseWithDefaults instantiates a new UsageCustomReportsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageCustomReportsResponseWithDefaults() *UsageCustomReportsResponse {
	this := UsageCustomReportsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UsageCustomReportsResponse) GetData() []UsageCustomReportsData {
	if o == nil || o.Data == nil {
		var ret []UsageCustomReportsData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCustomReportsResponse) GetDataOk() (*[]UsageCustomReportsData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UsageCustomReportsResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []UsageCustomReportsData and assigns it to the Data field.
func (o *UsageCustomReportsResponse) SetData(v []UsageCustomReportsData) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *UsageCustomReportsResponse) GetMeta() UsageCustomReportsMeta {
	if o == nil || o.Meta == nil {
		var ret UsageCustomReportsMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCustomReportsResponse) GetMetaOk() (*UsageCustomReportsMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *UsageCustomReportsResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given UsageCustomReportsMeta and assigns it to the Meta field.
func (o *UsageCustomReportsResponse) SetMeta(v UsageCustomReportsMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageCustomReportsResponse) MarshalJSON() ([]byte, error) {
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
func (o *UsageCustomReportsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data []UsageCustomReportsData `json:"data,omitempty"`
		Meta *UsageCustomReportsMeta  `json:"meta,omitempty"`
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
