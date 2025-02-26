// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsAggregateRequestPage Paging settings
type LogsAggregateRequestPage struct {
	// The returned paging point to use to get the next results. Note: at most 1000 results can be paged.
	Cursor *string `json:"cursor,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewLogsAggregateRequestPage instantiates a new LogsAggregateRequestPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsAggregateRequestPage() *LogsAggregateRequestPage {
	this := LogsAggregateRequestPage{}
	return &this
}

// NewLogsAggregateRequestPageWithDefaults instantiates a new LogsAggregateRequestPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsAggregateRequestPageWithDefaults() *LogsAggregateRequestPage {
	this := LogsAggregateRequestPage{}
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *LogsAggregateRequestPage) GetCursor() string {
	if o == nil || o.Cursor == nil {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsAggregateRequestPage) GetCursorOk() (*string, bool) {
	if o == nil || o.Cursor == nil {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *LogsAggregateRequestPage) HasCursor() bool {
	return o != nil && o.Cursor != nil
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *LogsAggregateRequestPage) SetCursor(v string) {
	o.Cursor = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsAggregateRequestPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Cursor != nil {
		toSerialize["cursor"] = o.Cursor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsAggregateRequestPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cursor *string `json:"cursor,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cursor"})
	} else {
		return err
	}
	o.Cursor = all.Cursor

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
