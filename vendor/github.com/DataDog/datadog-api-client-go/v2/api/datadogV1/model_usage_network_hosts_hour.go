// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"time"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageNetworkHostsHour Number of active NPM hosts for each hour for a given organization.
type UsageNetworkHostsHour struct {
	// Contains the number of active NPM hosts.
	HostCount datadog.NullableInt64 `json:"host_count,omitempty"`
	// The hour for the usage.
	Hour *time.Time `json:"hour,omitempty"`
	// The organization name.
	OrgName *string `json:"org_name,omitempty"`
	// The organization public ID.
	PublicId *string `json:"public_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewUsageNetworkHostsHour instantiates a new UsageNetworkHostsHour object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageNetworkHostsHour() *UsageNetworkHostsHour {
	this := UsageNetworkHostsHour{}
	return &this
}

// NewUsageNetworkHostsHourWithDefaults instantiates a new UsageNetworkHostsHour object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageNetworkHostsHourWithDefaults() *UsageNetworkHostsHour {
	this := UsageNetworkHostsHour{}
	return &this
}

// GetHostCount returns the HostCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageNetworkHostsHour) GetHostCount() int64 {
	if o == nil || o.HostCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.HostCount.Get()
}

// GetHostCountOk returns a tuple with the HostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageNetworkHostsHour) GetHostCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.HostCount.Get(), o.HostCount.IsSet()
}

// HasHostCount returns a boolean if a field has been set.
func (o *UsageNetworkHostsHour) HasHostCount() bool {
	return o != nil && o.HostCount.IsSet()
}

// SetHostCount gets a reference to the given datadog.NullableInt64 and assigns it to the HostCount field.
func (o *UsageNetworkHostsHour) SetHostCount(v int64) {
	o.HostCount.Set(&v)
}

// SetHostCountNil sets the value for HostCount to be an explicit nil.
func (o *UsageNetworkHostsHour) SetHostCountNil() {
	o.HostCount.Set(nil)
}

// UnsetHostCount ensures that no value is present for HostCount, not even an explicit nil.
func (o *UsageNetworkHostsHour) UnsetHostCount() {
	o.HostCount.Unset()
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageNetworkHostsHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageNetworkHostsHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageNetworkHostsHour) HasHour() bool {
	return o != nil && o.Hour != nil
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageNetworkHostsHour) SetHour(v time.Time) {
	o.Hour = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *UsageNetworkHostsHour) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageNetworkHostsHour) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *UsageNetworkHostsHour) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *UsageNetworkHostsHour) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *UsageNetworkHostsHour) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageNetworkHostsHour) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *UsageNetworkHostsHour) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *UsageNetworkHostsHour) SetPublicId(v string) {
	o.PublicId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageNetworkHostsHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.HostCount.IsSet() {
		toSerialize["host_count"] = o.HostCount.Get()
	}
	if o.Hour != nil {
		if o.Hour.Nanosecond() == 0 {
			toSerialize["hour"] = o.Hour.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["hour"] = o.Hour.Format("2006-01-02T15:04:05.000Z07:00")
		}
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
func (o *UsageNetworkHostsHour) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HostCount datadog.NullableInt64 `json:"host_count,omitempty"`
		Hour      *time.Time            `json:"hour,omitempty"`
		OrgName   *string               `json:"org_name,omitempty"`
		PublicId  *string               `json:"public_id,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"host_count", "hour", "org_name", "public_id"})
	} else {
		return err
	}
	o.HostCount = all.HostCount
	o.Hour = all.Hour
	o.OrgName = all.OrgName
	o.PublicId = all.PublicId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
