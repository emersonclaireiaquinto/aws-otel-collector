// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PowerpackRelationships Powerpack relationship object.
type PowerpackRelationships struct {
	// Creator of the object.
	Author *Creator `json:"author,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewPowerpackRelationships instantiates a new PowerpackRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPowerpackRelationships() *PowerpackRelationships {
	this := PowerpackRelationships{}
	return &this
}

// NewPowerpackRelationshipsWithDefaults instantiates a new PowerpackRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPowerpackRelationshipsWithDefaults() *PowerpackRelationships {
	this := PowerpackRelationships{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *PowerpackRelationships) GetAuthor() Creator {
	if o == nil || o.Author == nil {
		var ret Creator
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackRelationships) GetAuthorOk() (*Creator, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *PowerpackRelationships) HasAuthor() bool {
	return o != nil && o.Author != nil
}

// SetAuthor gets a reference to the given Creator and assigns it to the Author field.
func (o *PowerpackRelationships) SetAuthor(v Creator) {
	o.Author = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PowerpackRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PowerpackRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Author *Creator `json:"author,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Author != nil && all.Author.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Author = all.Author

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
