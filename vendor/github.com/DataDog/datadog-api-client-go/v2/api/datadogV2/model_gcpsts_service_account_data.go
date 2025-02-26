// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPSTSServiceAccountData Additional metadata on your generated service account.
type GCPSTSServiceAccountData struct {
	// Attributes associated with your service account.
	Attributes *GCPSTSServiceAccountAttributes `json:"attributes,omitempty"`
	// The type of account.
	Type *GCPServiceAccountType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewGCPSTSServiceAccountData instantiates a new GCPSTSServiceAccountData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGCPSTSServiceAccountData() *GCPSTSServiceAccountData {
	this := GCPSTSServiceAccountData{}
	var typeVar GCPServiceAccountType = GCPSERVICEACCOUNTTYPE_GCP_SERVICE_ACCOUNT
	this.Type = &typeVar
	return &this
}

// NewGCPSTSServiceAccountDataWithDefaults instantiates a new GCPSTSServiceAccountData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGCPSTSServiceAccountDataWithDefaults() *GCPSTSServiceAccountData {
	this := GCPSTSServiceAccountData{}
	var typeVar GCPServiceAccountType = GCPSERVICEACCOUNTTYPE_GCP_SERVICE_ACCOUNT
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GCPSTSServiceAccountData) GetAttributes() GCPSTSServiceAccountAttributes {
	if o == nil || o.Attributes == nil {
		var ret GCPSTSServiceAccountAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPSTSServiceAccountData) GetAttributesOk() (*GCPSTSServiceAccountAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GCPSTSServiceAccountData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given GCPSTSServiceAccountAttributes and assigns it to the Attributes field.
func (o *GCPSTSServiceAccountData) SetAttributes(v GCPSTSServiceAccountAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GCPSTSServiceAccountData) GetType() GCPServiceAccountType {
	if o == nil || o.Type == nil {
		var ret GCPServiceAccountType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPSTSServiceAccountData) GetTypeOk() (*GCPServiceAccountType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GCPSTSServiceAccountData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given GCPServiceAccountType and assigns it to the Type field.
func (o *GCPSTSServiceAccountData) SetType(v GCPServiceAccountType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GCPSTSServiceAccountData) MarshalJSON() ([]byte, error) {
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
func (o *GCPSTSServiceAccountData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *GCPSTSServiceAccountAttributes `json:"attributes,omitempty"`
		Type       *GCPServiceAccountType          `json:"type,omitempty"`
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
