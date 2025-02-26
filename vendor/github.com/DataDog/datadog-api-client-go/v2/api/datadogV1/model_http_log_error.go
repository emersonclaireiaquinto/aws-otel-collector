// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HTTPLogError Invalid query performed.
type HTTPLogError struct {
	// Error code.
	Code int32 `json:"code"`
	// Error message.
	Message string `json:"message"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewHTTPLogError instantiates a new HTTPLogError object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHTTPLogError(code int32, message string) *HTTPLogError {
	this := HTTPLogError{}
	this.Code = code
	this.Message = message
	return &this
}

// NewHTTPLogErrorWithDefaults instantiates a new HTTPLogError object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHTTPLogErrorWithDefaults() *HTTPLogError {
	this := HTTPLogError{}
	return &this
}

// GetCode returns the Code field value.
func (o *HTTPLogError) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *HTTPLogError) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value.
func (o *HTTPLogError) SetCode(v int32) {
	o.Code = v
}

// GetMessage returns the Message field value.
func (o *HTTPLogError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *HTTPLogError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *HTTPLogError) SetMessage(v string) {
	o.Message = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HTTPLogError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HTTPLogError) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Code    *int32  `json:"code"`
		Message *string `json:"message"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Code == nil {
		return fmt.Errorf("required field code missing")
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"code", "message"})
	} else {
		return err
	}
	o.Code = *all.Code
	o.Message = *all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
