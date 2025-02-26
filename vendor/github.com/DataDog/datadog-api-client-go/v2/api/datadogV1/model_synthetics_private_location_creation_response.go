// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsPrivateLocationCreationResponse Object that contains the new private location, the public key for result encryption, and the configuration skeleton.
type SyntheticsPrivateLocationCreationResponse struct {
	// Configuration skeleton for the private location. See installation instructions of the private location on how to use this configuration.
	Config interface{} `json:"config,omitempty"`
	// Object containing information about the private location to create.
	PrivateLocation *SyntheticsPrivateLocation `json:"private_location,omitempty"`
	// Public key for the result encryption.
	ResultEncryption *SyntheticsPrivateLocationCreationResponseResultEncryption `json:"result_encryption,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSyntheticsPrivateLocationCreationResponse instantiates a new SyntheticsPrivateLocationCreationResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsPrivateLocationCreationResponse() *SyntheticsPrivateLocationCreationResponse {
	this := SyntheticsPrivateLocationCreationResponse{}
	return &this
}

// NewSyntheticsPrivateLocationCreationResponseWithDefaults instantiates a new SyntheticsPrivateLocationCreationResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsPrivateLocationCreationResponseWithDefaults() *SyntheticsPrivateLocationCreationResponse {
	this := SyntheticsPrivateLocationCreationResponse{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *SyntheticsPrivateLocationCreationResponse) GetConfig() interface{} {
	if o == nil || o.Config == nil {
		var ret interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsPrivateLocationCreationResponse) GetConfigOk() (*interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return &o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *SyntheticsPrivateLocationCreationResponse) HasConfig() bool {
	return o != nil && o.Config != nil
}

// SetConfig gets a reference to the given interface{} and assigns it to the Config field.
func (o *SyntheticsPrivateLocationCreationResponse) SetConfig(v interface{}) {
	o.Config = v
}

// GetPrivateLocation returns the PrivateLocation field value if set, zero value otherwise.
func (o *SyntheticsPrivateLocationCreationResponse) GetPrivateLocation() SyntheticsPrivateLocation {
	if o == nil || o.PrivateLocation == nil {
		var ret SyntheticsPrivateLocation
		return ret
	}
	return *o.PrivateLocation
}

// GetPrivateLocationOk returns a tuple with the PrivateLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsPrivateLocationCreationResponse) GetPrivateLocationOk() (*SyntheticsPrivateLocation, bool) {
	if o == nil || o.PrivateLocation == nil {
		return nil, false
	}
	return o.PrivateLocation, true
}

// HasPrivateLocation returns a boolean if a field has been set.
func (o *SyntheticsPrivateLocationCreationResponse) HasPrivateLocation() bool {
	return o != nil && o.PrivateLocation != nil
}

// SetPrivateLocation gets a reference to the given SyntheticsPrivateLocation and assigns it to the PrivateLocation field.
func (o *SyntheticsPrivateLocationCreationResponse) SetPrivateLocation(v SyntheticsPrivateLocation) {
	o.PrivateLocation = &v
}

// GetResultEncryption returns the ResultEncryption field value if set, zero value otherwise.
func (o *SyntheticsPrivateLocationCreationResponse) GetResultEncryption() SyntheticsPrivateLocationCreationResponseResultEncryption {
	if o == nil || o.ResultEncryption == nil {
		var ret SyntheticsPrivateLocationCreationResponseResultEncryption
		return ret
	}
	return *o.ResultEncryption
}

// GetResultEncryptionOk returns a tuple with the ResultEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsPrivateLocationCreationResponse) GetResultEncryptionOk() (*SyntheticsPrivateLocationCreationResponseResultEncryption, bool) {
	if o == nil || o.ResultEncryption == nil {
		return nil, false
	}
	return o.ResultEncryption, true
}

// HasResultEncryption returns a boolean if a field has been set.
func (o *SyntheticsPrivateLocationCreationResponse) HasResultEncryption() bool {
	return o != nil && o.ResultEncryption != nil
}

// SetResultEncryption gets a reference to the given SyntheticsPrivateLocationCreationResponseResultEncryption and assigns it to the ResultEncryption field.
func (o *SyntheticsPrivateLocationCreationResponse) SetResultEncryption(v SyntheticsPrivateLocationCreationResponseResultEncryption) {
	o.ResultEncryption = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsPrivateLocationCreationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.PrivateLocation != nil {
		toSerialize["private_location"] = o.PrivateLocation
	}
	if o.ResultEncryption != nil {
		toSerialize["result_encryption"] = o.ResultEncryption
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsPrivateLocationCreationResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Config           interface{}                                                `json:"config,omitempty"`
		PrivateLocation  *SyntheticsPrivateLocation                                 `json:"private_location,omitempty"`
		ResultEncryption *SyntheticsPrivateLocationCreationResponseResultEncryption `json:"result_encryption,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"config", "private_location", "result_encryption"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Config = all.Config
	if all.PrivateLocation != nil && all.PrivateLocation.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PrivateLocation = all.PrivateLocation
	if all.ResultEncryption != nil && all.ResultEncryption.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ResultEncryption = all.ResultEncryption

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
