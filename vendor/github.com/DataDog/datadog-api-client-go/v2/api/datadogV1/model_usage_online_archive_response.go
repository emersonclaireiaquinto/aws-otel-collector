// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageOnlineArchiveResponse Online Archive usage response.
type UsageOnlineArchiveResponse struct {
	// Response containing Online Archive usage.
	Usage []UsageOnlineArchiveHour `json:"usage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewUsageOnlineArchiveResponse instantiates a new UsageOnlineArchiveResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageOnlineArchiveResponse() *UsageOnlineArchiveResponse {
	this := UsageOnlineArchiveResponse{}
	return &this
}

// NewUsageOnlineArchiveResponseWithDefaults instantiates a new UsageOnlineArchiveResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageOnlineArchiveResponseWithDefaults() *UsageOnlineArchiveResponse {
	this := UsageOnlineArchiveResponse{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageOnlineArchiveResponse) GetUsage() []UsageOnlineArchiveHour {
	if o == nil || o.Usage == nil {
		var ret []UsageOnlineArchiveHour
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageOnlineArchiveResponse) GetUsageOk() (*[]UsageOnlineArchiveHour, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return &o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageOnlineArchiveResponse) HasUsage() bool {
	return o != nil && o.Usage != nil
}

// SetUsage gets a reference to the given []UsageOnlineArchiveHour and assigns it to the Usage field.
func (o *UsageOnlineArchiveResponse) SetUsage(v []UsageOnlineArchiveHour) {
	o.Usage = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageOnlineArchiveResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageOnlineArchiveResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Usage []UsageOnlineArchiveHour `json:"usage,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"usage"})
	} else {
		return err
	}
	o.Usage = all.Usage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
