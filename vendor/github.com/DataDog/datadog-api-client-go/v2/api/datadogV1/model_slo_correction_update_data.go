// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOCorrectionUpdateData The data object associated with the SLO correction to be updated.
type SLOCorrectionUpdateData struct {
	// The attribute object associated with the SLO correction to be updated.
	Attributes *SLOCorrectionUpdateRequestAttributes `json:"attributes,omitempty"`
	// SLO correction resource type.
	Type *SLOCorrectionType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSLOCorrectionUpdateData instantiates a new SLOCorrectionUpdateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSLOCorrectionUpdateData() *SLOCorrectionUpdateData {
	this := SLOCorrectionUpdateData{}
	var typeVar SLOCorrectionType = SLOCORRECTIONTYPE_CORRECTION
	this.Type = &typeVar
	return &this
}

// NewSLOCorrectionUpdateDataWithDefaults instantiates a new SLOCorrectionUpdateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSLOCorrectionUpdateDataWithDefaults() *SLOCorrectionUpdateData {
	this := SLOCorrectionUpdateData{}
	var typeVar SLOCorrectionType = SLOCORRECTIONTYPE_CORRECTION
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SLOCorrectionUpdateData) GetAttributes() SLOCorrectionUpdateRequestAttributes {
	if o == nil || o.Attributes == nil {
		var ret SLOCorrectionUpdateRequestAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOCorrectionUpdateData) GetAttributesOk() (*SLOCorrectionUpdateRequestAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SLOCorrectionUpdateData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given SLOCorrectionUpdateRequestAttributes and assigns it to the Attributes field.
func (o *SLOCorrectionUpdateData) SetAttributes(v SLOCorrectionUpdateRequestAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SLOCorrectionUpdateData) GetType() SLOCorrectionType {
	if o == nil || o.Type == nil {
		var ret SLOCorrectionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOCorrectionUpdateData) GetTypeOk() (*SLOCorrectionType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SLOCorrectionUpdateData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given SLOCorrectionType and assigns it to the Type field.
func (o *SLOCorrectionUpdateData) SetType(v SLOCorrectionType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SLOCorrectionUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
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
func (o *SLOCorrectionUpdateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *SLOCorrectionUpdateRequestAttributes `json:"attributes,omitempty"`
		Type       *SLOCorrectionType                    `json:"type,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
