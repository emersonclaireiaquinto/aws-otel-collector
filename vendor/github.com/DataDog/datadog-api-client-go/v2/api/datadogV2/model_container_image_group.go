// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainerImageGroup Container Image Group object.
type ContainerImageGroup struct {
	// Attributes for a Container Image Group.
	Attributes *ContainerImageGroupAttributes `json:"attributes,omitempty"`
	// Container Image Group ID.
	Id *string `json:"id,omitempty"`
	// Relationships inside a Container Image Group.
	Relationships *ContainerImageGroupRelationships `json:"relationships,omitempty"`
	// Type of Container Image Group.
	Type *ContainerImageGroupType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewContainerImageGroup instantiates a new ContainerImageGroup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewContainerImageGroup() *ContainerImageGroup {
	this := ContainerImageGroup{}
	var typeVar ContainerImageGroupType = CONTAINERIMAGEGROUPTYPE_CONTAINER_IMAGE_GROUP
	this.Type = &typeVar
	return &this
}

// NewContainerImageGroupWithDefaults instantiates a new ContainerImageGroup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewContainerImageGroupWithDefaults() *ContainerImageGroup {
	this := ContainerImageGroup{}
	var typeVar ContainerImageGroupType = CONTAINERIMAGEGROUPTYPE_CONTAINER_IMAGE_GROUP
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ContainerImageGroup) GetAttributes() ContainerImageGroupAttributes {
	if o == nil || o.Attributes == nil {
		var ret ContainerImageGroupAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageGroup) GetAttributesOk() (*ContainerImageGroupAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ContainerImageGroup) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given ContainerImageGroupAttributes and assigns it to the Attributes field.
func (o *ContainerImageGroup) SetAttributes(v ContainerImageGroupAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContainerImageGroup) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageGroup) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContainerImageGroup) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ContainerImageGroup) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ContainerImageGroup) GetRelationships() ContainerImageGroupRelationships {
	if o == nil || o.Relationships == nil {
		var ret ContainerImageGroupRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageGroup) GetRelationshipsOk() (*ContainerImageGroupRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ContainerImageGroup) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given ContainerImageGroupRelationships and assigns it to the Relationships field.
func (o *ContainerImageGroup) SetRelationships(v ContainerImageGroupRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ContainerImageGroup) GetType() ContainerImageGroupType {
	if o == nil || o.Type == nil {
		var ret ContainerImageGroupType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageGroup) GetTypeOk() (*ContainerImageGroupType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ContainerImageGroup) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ContainerImageGroupType and assigns it to the Type field.
func (o *ContainerImageGroup) SetType(v ContainerImageGroupType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ContainerImageGroup) MarshalJSON() ([]byte, error) {
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
func (o *ContainerImageGroup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *ContainerImageGroupAttributes    `json:"attributes,omitempty"`
		Id            *string                           `json:"id,omitempty"`
		Relationships *ContainerImageGroupRelationships `json:"relationships,omitempty"`
		Type          *ContainerImageGroupType          `json:"type,omitempty"`
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
