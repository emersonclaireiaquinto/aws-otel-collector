// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// APIKeyCreateData Object used to create an API key.
type APIKeyCreateData struct {
	// Attributes used to create an API Key.
	Attributes APIKeyCreateAttributes `json:"attributes"`
	// API Keys resource type.
	Type APIKeysType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewAPIKeyCreateData instantiates a new APIKeyCreateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAPIKeyCreateData(attributes APIKeyCreateAttributes, typeVar APIKeysType) *APIKeyCreateData {
	this := APIKeyCreateData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewAPIKeyCreateDataWithDefaults instantiates a new APIKeyCreateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAPIKeyCreateDataWithDefaults() *APIKeyCreateData {
	this := APIKeyCreateData{}
	var typeVar APIKeysType = APIKEYSTYPE_API_KEYS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *APIKeyCreateData) GetAttributes() APIKeyCreateAttributes {
	if o == nil {
		var ret APIKeyCreateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *APIKeyCreateData) GetAttributesOk() (*APIKeyCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *APIKeyCreateData) SetAttributes(v APIKeyCreateAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *APIKeyCreateData) GetType() APIKeysType {
	if o == nil {
		var ret APIKeysType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *APIKeyCreateData) GetTypeOk() (*APIKeysType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *APIKeyCreateData) SetType(v APIKeysType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o APIKeyCreateData) MarshalJSON() ([]byte, error) {
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
func (o *APIKeyCreateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *APIKeyCreateAttributes `json:"attributes"`
		Type       *APIKeysType            `json:"type"`
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
