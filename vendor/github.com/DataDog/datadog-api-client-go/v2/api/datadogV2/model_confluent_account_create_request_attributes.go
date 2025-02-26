// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConfluentAccountCreateRequestAttributes Attributes associated with the account creation request.
type ConfluentAccountCreateRequestAttributes struct {
	// The API key associated with your Confluent account.
	ApiKey string `json:"api_key"`
	// The API secret associated with your Confluent account.
	ApiSecret string `json:"api_secret"`
	// A list of Confluent resources associated with the Confluent account.
	Resources []ConfluentAccountResourceAttributes `json:"resources,omitempty"`
	// A list of strings representing tags. Can be a single key, or key-value pairs separated by a colon.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewConfluentAccountCreateRequestAttributes instantiates a new ConfluentAccountCreateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConfluentAccountCreateRequestAttributes(apiKey string, apiSecret string) *ConfluentAccountCreateRequestAttributes {
	this := ConfluentAccountCreateRequestAttributes{}
	this.ApiKey = apiKey
	this.ApiSecret = apiSecret
	return &this
}

// NewConfluentAccountCreateRequestAttributesWithDefaults instantiates a new ConfluentAccountCreateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConfluentAccountCreateRequestAttributesWithDefaults() *ConfluentAccountCreateRequestAttributes {
	this := ConfluentAccountCreateRequestAttributes{}
	return &this
}

// GetApiKey returns the ApiKey field value.
func (o *ConfluentAccountCreateRequestAttributes) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *ConfluentAccountCreateRequestAttributes) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value.
func (o *ConfluentAccountCreateRequestAttributes) SetApiKey(v string) {
	o.ApiKey = v
}

// GetApiSecret returns the ApiSecret field value.
func (o *ConfluentAccountCreateRequestAttributes) GetApiSecret() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApiSecret
}

// GetApiSecretOk returns a tuple with the ApiSecret field value
// and a boolean to check if the value has been set.
func (o *ConfluentAccountCreateRequestAttributes) GetApiSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiSecret, true
}

// SetApiSecret sets field value.
func (o *ConfluentAccountCreateRequestAttributes) SetApiSecret(v string) {
	o.ApiSecret = v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *ConfluentAccountCreateRequestAttributes) GetResources() []ConfluentAccountResourceAttributes {
	if o == nil || o.Resources == nil {
		var ret []ConfluentAccountResourceAttributes
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfluentAccountCreateRequestAttributes) GetResourcesOk() (*[]ConfluentAccountResourceAttributes, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return &o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *ConfluentAccountCreateRequestAttributes) HasResources() bool {
	return o != nil && o.Resources != nil
}

// SetResources gets a reference to the given []ConfluentAccountResourceAttributes and assigns it to the Resources field.
func (o *ConfluentAccountCreateRequestAttributes) SetResources(v []ConfluentAccountResourceAttributes) {
	o.Resources = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ConfluentAccountCreateRequestAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfluentAccountCreateRequestAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ConfluentAccountCreateRequestAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ConfluentAccountCreateRequestAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ConfluentAccountCreateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["api_key"] = o.ApiKey
	toSerialize["api_secret"] = o.ApiSecret
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ConfluentAccountCreateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiKey    *string                              `json:"api_key"`
		ApiSecret *string                              `json:"api_secret"`
		Resources []ConfluentAccountResourceAttributes `json:"resources,omitempty"`
		Tags      []string                             `json:"tags,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApiKey == nil {
		return fmt.Errorf("required field api_key missing")
	}
	if all.ApiSecret == nil {
		return fmt.Errorf("required field api_secret missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"api_key", "api_secret", "resources", "tags"})
	} else {
		return err
	}
	o.ApiKey = *all.ApiKey
	o.ApiSecret = *all.ApiSecret
	o.Resources = all.Resources
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
