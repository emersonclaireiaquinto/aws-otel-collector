// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FullApplicationKey Datadog application key.
type FullApplicationKey struct {
	// Attributes of a full application key.
	Attributes *FullApplicationKeyAttributes `json:"attributes,omitempty"`
	// ID of the application key.
	Id *string `json:"id,omitempty"`
	// Resources related to the application key.
	Relationships *ApplicationKeyRelationships `json:"relationships,omitempty"`
	// Application Keys resource type.
	Type *ApplicationKeysType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewFullApplicationKey instantiates a new FullApplicationKey object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFullApplicationKey() *FullApplicationKey {
	this := FullApplicationKey{}
	var typeVar ApplicationKeysType = APPLICATIONKEYSTYPE_APPLICATION_KEYS
	this.Type = &typeVar
	return &this
}

// NewFullApplicationKeyWithDefaults instantiates a new FullApplicationKey object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFullApplicationKeyWithDefaults() *FullApplicationKey {
	this := FullApplicationKey{}
	var typeVar ApplicationKeysType = APPLICATIONKEYSTYPE_APPLICATION_KEYS
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *FullApplicationKey) GetAttributes() FullApplicationKeyAttributes {
	if o == nil || o.Attributes == nil {
		var ret FullApplicationKeyAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullApplicationKey) GetAttributesOk() (*FullApplicationKeyAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *FullApplicationKey) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given FullApplicationKeyAttributes and assigns it to the Attributes field.
func (o *FullApplicationKey) SetAttributes(v FullApplicationKeyAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FullApplicationKey) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullApplicationKey) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FullApplicationKey) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FullApplicationKey) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *FullApplicationKey) GetRelationships() ApplicationKeyRelationships {
	if o == nil || o.Relationships == nil {
		var ret ApplicationKeyRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullApplicationKey) GetRelationshipsOk() (*ApplicationKeyRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *FullApplicationKey) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given ApplicationKeyRelationships and assigns it to the Relationships field.
func (o *FullApplicationKey) SetRelationships(v ApplicationKeyRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FullApplicationKey) GetType() ApplicationKeysType {
	if o == nil || o.Type == nil {
		var ret ApplicationKeysType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullApplicationKey) GetTypeOk() (*ApplicationKeysType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FullApplicationKey) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ApplicationKeysType and assigns it to the Type field.
func (o *FullApplicationKey) SetType(v ApplicationKeysType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FullApplicationKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
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
func (o *FullApplicationKey) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *FullApplicationKeyAttributes `json:"attributes,omitempty"`
		Id            *string                       `json:"id,omitempty"`
		Relationships *ApplicationKeyRelationships  `json:"relationships,omitempty"`
		Type          *ApplicationKeysType          `json:"type,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
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
