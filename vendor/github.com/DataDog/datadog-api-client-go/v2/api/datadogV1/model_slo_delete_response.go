// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLODeleteResponse A response list of all service level objective deleted.
type SLODeleteResponse struct {
	// An array containing the ID of the deleted service level objective object.
	Data []string `json:"data,omitempty"`
	// An dictionary containing the ID of the SLO as key and a deletion error as value.
	Errors map[string]string `json:"errors,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSLODeleteResponse instantiates a new SLODeleteResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSLODeleteResponse() *SLODeleteResponse {
	this := SLODeleteResponse{}
	return &this
}

// NewSLODeleteResponseWithDefaults instantiates a new SLODeleteResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSLODeleteResponseWithDefaults() *SLODeleteResponse {
	this := SLODeleteResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SLODeleteResponse) GetData() []string {
	if o == nil || o.Data == nil {
		var ret []string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLODeleteResponse) GetDataOk() (*[]string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SLODeleteResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []string and assigns it to the Data field.
func (o *SLODeleteResponse) SetData(v []string) {
	o.Data = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SLODeleteResponse) GetErrors() map[string]string {
	if o == nil || o.Errors == nil {
		var ret map[string]string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLODeleteResponse) GetErrorsOk() (*map[string]string, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return &o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SLODeleteResponse) HasErrors() bool {
	return o != nil && o.Errors != nil
}

// SetErrors gets a reference to the given map[string]string and assigns it to the Errors field.
func (o *SLODeleteResponse) SetErrors(v map[string]string) {
	o.Errors = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SLODeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SLODeleteResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data   []string          `json:"data,omitempty"`
		Errors map[string]string `json:"errors,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "errors"})
	} else {
		return err
	}
	o.Data = all.Data
	o.Errors = all.Errors

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
