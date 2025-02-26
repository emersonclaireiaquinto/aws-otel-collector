// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceDefinitionV2Link Service's external links.
type ServiceDefinitionV2Link struct {
	// Link name.
	Name string `json:"name"`
	// Link type.
	Type ServiceDefinitionV2LinkType `json:"type"`
	// Link URL.
	Url string `json:"url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewServiceDefinitionV2Link instantiates a new ServiceDefinitionV2Link object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceDefinitionV2Link(name string, typeVar ServiceDefinitionV2LinkType, url string) *ServiceDefinitionV2Link {
	this := ServiceDefinitionV2Link{}
	this.Name = name
	this.Type = typeVar
	this.Url = url
	return &this
}

// NewServiceDefinitionV2LinkWithDefaults instantiates a new ServiceDefinitionV2Link object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceDefinitionV2LinkWithDefaults() *ServiceDefinitionV2Link {
	this := ServiceDefinitionV2Link{}
	return &this
}

// GetName returns the Name field value.
func (o *ServiceDefinitionV2Link) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Link) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ServiceDefinitionV2Link) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *ServiceDefinitionV2Link) GetType() ServiceDefinitionV2LinkType {
	if o == nil {
		var ret ServiceDefinitionV2LinkType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Link) GetTypeOk() (*ServiceDefinitionV2LinkType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ServiceDefinitionV2Link) SetType(v ServiceDefinitionV2LinkType) {
	o.Type = v
}

// GetUrl returns the Url field value.
func (o *ServiceDefinitionV2Link) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Link) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value.
func (o *ServiceDefinitionV2Link) SetUrl(v string) {
	o.Url = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceDefinitionV2Link) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceDefinitionV2Link) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *string                      `json:"name"`
		Type *ServiceDefinitionV2LinkType `json:"type"`
		Url  *string                      `json:"url"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Url == nil {
		return fmt.Errorf("required field url missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "type", "url"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Url = *all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
