// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceDefinitionV2Repo Service code repositories.
type ServiceDefinitionV2Repo struct {
	// Repository name.
	Name string `json:"name"`
	// Repository provider.
	Provider *string `json:"provider,omitempty"`
	// Repository URL.
	Url string `json:"url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewServiceDefinitionV2Repo instantiates a new ServiceDefinitionV2Repo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceDefinitionV2Repo(name string, url string) *ServiceDefinitionV2Repo {
	this := ServiceDefinitionV2Repo{}
	this.Name = name
	this.Url = url
	return &this
}

// NewServiceDefinitionV2RepoWithDefaults instantiates a new ServiceDefinitionV2Repo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceDefinitionV2RepoWithDefaults() *ServiceDefinitionV2Repo {
	this := ServiceDefinitionV2Repo{}
	return &this
}

// GetName returns the Name field value.
func (o *ServiceDefinitionV2Repo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Repo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ServiceDefinitionV2Repo) SetName(v string) {
	o.Name = v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Repo) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Repo) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Repo) HasProvider() bool {
	return o != nil && o.Provider != nil
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *ServiceDefinitionV2Repo) SetProvider(v string) {
	o.Provider = &v
}

// GetUrl returns the Url field value.
func (o *ServiceDefinitionV2Repo) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Repo) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value.
func (o *ServiceDefinitionV2Repo) SetUrl(v string) {
	o.Url = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceDefinitionV2Repo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceDefinitionV2Repo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name     *string `json:"name"`
		Provider *string `json:"provider,omitempty"`
		Url      *string `json:"url"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Url == nil {
		return fmt.Errorf("required field url missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "provider", "url"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Provider = all.Provider
	o.Url = *all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
