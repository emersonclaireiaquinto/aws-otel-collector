// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"time"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageIndexedSpansHour The hours of indexed spans usage.
type UsageIndexedSpansHour struct {
	// The hour for the usage.
	Hour *time.Time `json:"hour,omitempty"`
	// Contains the number of spans indexed.
	IndexedEventsCount datadog.NullableInt64 `json:"indexed_events_count,omitempty"`
	// The organization name.
	OrgName *string `json:"org_name,omitempty"`
	// The organization public ID.
	PublicId *string `json:"public_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewUsageIndexedSpansHour instantiates a new UsageIndexedSpansHour object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageIndexedSpansHour() *UsageIndexedSpansHour {
	this := UsageIndexedSpansHour{}
	return &this
}

// NewUsageIndexedSpansHourWithDefaults instantiates a new UsageIndexedSpansHour object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageIndexedSpansHourWithDefaults() *UsageIndexedSpansHour {
	this := UsageIndexedSpansHour{}
	return &this
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageIndexedSpansHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageIndexedSpansHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageIndexedSpansHour) HasHour() bool {
	return o != nil && o.Hour != nil
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageIndexedSpansHour) SetHour(v time.Time) {
	o.Hour = &v
}

// GetIndexedEventsCount returns the IndexedEventsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageIndexedSpansHour) GetIndexedEventsCount() int64 {
	if o == nil || o.IndexedEventsCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IndexedEventsCount.Get()
}

// GetIndexedEventsCountOk returns a tuple with the IndexedEventsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageIndexedSpansHour) GetIndexedEventsCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndexedEventsCount.Get(), o.IndexedEventsCount.IsSet()
}

// HasIndexedEventsCount returns a boolean if a field has been set.
func (o *UsageIndexedSpansHour) HasIndexedEventsCount() bool {
	return o != nil && o.IndexedEventsCount.IsSet()
}

// SetIndexedEventsCount gets a reference to the given datadog.NullableInt64 and assigns it to the IndexedEventsCount field.
func (o *UsageIndexedSpansHour) SetIndexedEventsCount(v int64) {
	o.IndexedEventsCount.Set(&v)
}

// SetIndexedEventsCountNil sets the value for IndexedEventsCount to be an explicit nil.
func (o *UsageIndexedSpansHour) SetIndexedEventsCountNil() {
	o.IndexedEventsCount.Set(nil)
}

// UnsetIndexedEventsCount ensures that no value is present for IndexedEventsCount, not even an explicit nil.
func (o *UsageIndexedSpansHour) UnsetIndexedEventsCount() {
	o.IndexedEventsCount.Unset()
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *UsageIndexedSpansHour) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageIndexedSpansHour) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *UsageIndexedSpansHour) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *UsageIndexedSpansHour) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *UsageIndexedSpansHour) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageIndexedSpansHour) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *UsageIndexedSpansHour) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *UsageIndexedSpansHour) SetPublicId(v string) {
	o.PublicId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageIndexedSpansHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Hour != nil {
		if o.Hour.Nanosecond() == 0 {
			toSerialize["hour"] = o.Hour.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["hour"] = o.Hour.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.IndexedEventsCount.IsSet() {
		toSerialize["indexed_events_count"] = o.IndexedEventsCount.Get()
	}
	if o.OrgName != nil {
		toSerialize["org_name"] = o.OrgName
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageIndexedSpansHour) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Hour               *time.Time            `json:"hour,omitempty"`
		IndexedEventsCount datadog.NullableInt64 `json:"indexed_events_count,omitempty"`
		OrgName            *string               `json:"org_name,omitempty"`
		PublicId           *string               `json:"public_id,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"hour", "indexed_events_count", "org_name", "public_id"})
	} else {
		return err
	}
	o.Hour = all.Hour
	o.IndexedEventsCount = all.IndexedEventsCount
	o.OrgName = all.OrgName
	o.PublicId = all.PublicId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
