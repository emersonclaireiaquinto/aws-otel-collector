// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookSplitBy Object describing how to split the graph to display multiple visualizations per request.
type NotebookSplitBy struct {
	// Keys to split on.
	Keys []string `json:"keys"`
	// Tags to split on.
	Tags []string `json:"tags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewNotebookSplitBy instantiates a new NotebookSplitBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotebookSplitBy(keys []string, tags []string) *NotebookSplitBy {
	this := NotebookSplitBy{}
	this.Keys = keys
	this.Tags = tags
	return &this
}

// NewNotebookSplitByWithDefaults instantiates a new NotebookSplitBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotebookSplitByWithDefaults() *NotebookSplitBy {
	this := NotebookSplitBy{}
	return &this
}

// GetKeys returns the Keys field value.
func (o *NotebookSplitBy) GetKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value
// and a boolean to check if the value has been set.
func (o *NotebookSplitBy) GetKeysOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Keys, true
}

// SetKeys sets field value.
func (o *NotebookSplitBy) SetKeys(v []string) {
	o.Keys = v
}

// GetTags returns the Tags field value.
func (o *NotebookSplitBy) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *NotebookSplitBy) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *NotebookSplitBy) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotebookSplitBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["keys"] = o.Keys
	toSerialize["tags"] = o.Tags

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotebookSplitBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Keys *[]string `json:"keys"`
		Tags *[]string `json:"tags"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Keys == nil {
		return fmt.Errorf("required field keys missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"keys", "tags"})
	} else {
		return err
	}
	o.Keys = *all.Keys
	o.Tags = *all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
