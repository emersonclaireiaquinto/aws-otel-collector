// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOBulkDeleteError Object describing the error.
type SLOBulkDeleteError struct {
	// The ID of the service level objective object associated with
	// this error.
	Id string `json:"id"`
	// The error message.
	Message string `json:"message"`
	// The timeframe of the threshold associated with this error
	// or "all" if all thresholds are affected.
	Timeframe SLOErrorTimeframe `json:"timeframe"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSLOBulkDeleteError instantiates a new SLOBulkDeleteError object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSLOBulkDeleteError(id string, message string, timeframe SLOErrorTimeframe) *SLOBulkDeleteError {
	this := SLOBulkDeleteError{}
	this.Id = id
	this.Message = message
	this.Timeframe = timeframe
	return &this
}

// NewSLOBulkDeleteErrorWithDefaults instantiates a new SLOBulkDeleteError object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSLOBulkDeleteErrorWithDefaults() *SLOBulkDeleteError {
	this := SLOBulkDeleteError{}
	return &this
}

// GetId returns the Id field value.
func (o *SLOBulkDeleteError) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SLOBulkDeleteError) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *SLOBulkDeleteError) SetId(v string) {
	o.Id = v
}

// GetMessage returns the Message field value.
func (o *SLOBulkDeleteError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SLOBulkDeleteError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *SLOBulkDeleteError) SetMessage(v string) {
	o.Message = v
}

// GetTimeframe returns the Timeframe field value.
func (o *SLOBulkDeleteError) GetTimeframe() SLOErrorTimeframe {
	if o == nil {
		var ret SLOErrorTimeframe
		return ret
	}
	return o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value
// and a boolean to check if the value has been set.
func (o *SLOBulkDeleteError) GetTimeframeOk() (*SLOErrorTimeframe, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeframe, true
}

// SetTimeframe sets field value.
func (o *SLOBulkDeleteError) SetTimeframe(v SLOErrorTimeframe) {
	o.Timeframe = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SLOBulkDeleteError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["message"] = o.Message
	toSerialize["timeframe"] = o.Timeframe

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SLOBulkDeleteError) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id        *string            `json:"id"`
		Message   *string            `json:"message"`
		Timeframe *SLOErrorTimeframe `json:"timeframe"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.Timeframe == nil {
		return fmt.Errorf("required field timeframe missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "message", "timeframe"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	o.Message = *all.Message
	if !all.Timeframe.IsValid() {
		hasInvalidField = true
	} else {
		o.Timeframe = *all.Timeframe
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
