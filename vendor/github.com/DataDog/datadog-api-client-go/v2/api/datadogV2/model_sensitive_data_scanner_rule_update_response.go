// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerRuleUpdateResponse Update rule response.
type SensitiveDataScannerRuleUpdateResponse struct {
	// Meta payload containing information about the API.
	Meta *SensitiveDataScannerMetaVersionOnly `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSensitiveDataScannerRuleUpdateResponse instantiates a new SensitiveDataScannerRuleUpdateResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSensitiveDataScannerRuleUpdateResponse() *SensitiveDataScannerRuleUpdateResponse {
	this := SensitiveDataScannerRuleUpdateResponse{}
	return &this
}

// NewSensitiveDataScannerRuleUpdateResponseWithDefaults instantiates a new SensitiveDataScannerRuleUpdateResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSensitiveDataScannerRuleUpdateResponseWithDefaults() *SensitiveDataScannerRuleUpdateResponse {
	this := SensitiveDataScannerRuleUpdateResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SensitiveDataScannerRuleUpdateResponse) GetMeta() SensitiveDataScannerMetaVersionOnly {
	if o == nil || o.Meta == nil {
		var ret SensitiveDataScannerMetaVersionOnly
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerRuleUpdateResponse) GetMetaOk() (*SensitiveDataScannerMetaVersionOnly, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SensitiveDataScannerRuleUpdateResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given SensitiveDataScannerMetaVersionOnly and assigns it to the Meta field.
func (o *SensitiveDataScannerRuleUpdateResponse) SetMeta(v SensitiveDataScannerMetaVersionOnly) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SensitiveDataScannerRuleUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
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
func (o *SensitiveDataScannerRuleUpdateResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Meta *SensitiveDataScannerMetaVersionOnly `json:"meta,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"meta"})
	} else {
		return err
	}

	hasInvalidField := false
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
