// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorStateGroup Monitor state for a single group.
type MonitorStateGroup struct {
	// Latest timestamp the monitor was in NO_DATA state.
	LastNodataTs *int64 `json:"last_nodata_ts,omitempty"`
	// Latest timestamp of the notification sent for this monitor group.
	LastNotifiedTs *int64 `json:"last_notified_ts,omitempty"`
	// Latest timestamp the monitor group was resolved.
	LastResolvedTs *int64 `json:"last_resolved_ts,omitempty"`
	// Latest timestamp the monitor group triggered.
	LastTriggeredTs *int64 `json:"last_triggered_ts,omitempty"`
	// The name of the monitor.
	Name *string `json:"name,omitempty"`
	// The different states your monitor can be in.
	Status *MonitorOverallStates `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewMonitorStateGroup instantiates a new MonitorStateGroup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorStateGroup() *MonitorStateGroup {
	this := MonitorStateGroup{}
	return &this
}

// NewMonitorStateGroupWithDefaults instantiates a new MonitorStateGroup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorStateGroupWithDefaults() *MonitorStateGroup {
	this := MonitorStateGroup{}
	return &this
}

// GetLastNodataTs returns the LastNodataTs field value if set, zero value otherwise.
func (o *MonitorStateGroup) GetLastNodataTs() int64 {
	if o == nil || o.LastNodataTs == nil {
		var ret int64
		return ret
	}
	return *o.LastNodataTs
}

// GetLastNodataTsOk returns a tuple with the LastNodataTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStateGroup) GetLastNodataTsOk() (*int64, bool) {
	if o == nil || o.LastNodataTs == nil {
		return nil, false
	}
	return o.LastNodataTs, true
}

// HasLastNodataTs returns a boolean if a field has been set.
func (o *MonitorStateGroup) HasLastNodataTs() bool {
	return o != nil && o.LastNodataTs != nil
}

// SetLastNodataTs gets a reference to the given int64 and assigns it to the LastNodataTs field.
func (o *MonitorStateGroup) SetLastNodataTs(v int64) {
	o.LastNodataTs = &v
}

// GetLastNotifiedTs returns the LastNotifiedTs field value if set, zero value otherwise.
func (o *MonitorStateGroup) GetLastNotifiedTs() int64 {
	if o == nil || o.LastNotifiedTs == nil {
		var ret int64
		return ret
	}
	return *o.LastNotifiedTs
}

// GetLastNotifiedTsOk returns a tuple with the LastNotifiedTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStateGroup) GetLastNotifiedTsOk() (*int64, bool) {
	if o == nil || o.LastNotifiedTs == nil {
		return nil, false
	}
	return o.LastNotifiedTs, true
}

// HasLastNotifiedTs returns a boolean if a field has been set.
func (o *MonitorStateGroup) HasLastNotifiedTs() bool {
	return o != nil && o.LastNotifiedTs != nil
}

// SetLastNotifiedTs gets a reference to the given int64 and assigns it to the LastNotifiedTs field.
func (o *MonitorStateGroup) SetLastNotifiedTs(v int64) {
	o.LastNotifiedTs = &v
}

// GetLastResolvedTs returns the LastResolvedTs field value if set, zero value otherwise.
func (o *MonitorStateGroup) GetLastResolvedTs() int64 {
	if o == nil || o.LastResolvedTs == nil {
		var ret int64
		return ret
	}
	return *o.LastResolvedTs
}

// GetLastResolvedTsOk returns a tuple with the LastResolvedTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStateGroup) GetLastResolvedTsOk() (*int64, bool) {
	if o == nil || o.LastResolvedTs == nil {
		return nil, false
	}
	return o.LastResolvedTs, true
}

// HasLastResolvedTs returns a boolean if a field has been set.
func (o *MonitorStateGroup) HasLastResolvedTs() bool {
	return o != nil && o.LastResolvedTs != nil
}

// SetLastResolvedTs gets a reference to the given int64 and assigns it to the LastResolvedTs field.
func (o *MonitorStateGroup) SetLastResolvedTs(v int64) {
	o.LastResolvedTs = &v
}

// GetLastTriggeredTs returns the LastTriggeredTs field value if set, zero value otherwise.
func (o *MonitorStateGroup) GetLastTriggeredTs() int64 {
	if o == nil || o.LastTriggeredTs == nil {
		var ret int64
		return ret
	}
	return *o.LastTriggeredTs
}

// GetLastTriggeredTsOk returns a tuple with the LastTriggeredTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStateGroup) GetLastTriggeredTsOk() (*int64, bool) {
	if o == nil || o.LastTriggeredTs == nil {
		return nil, false
	}
	return o.LastTriggeredTs, true
}

// HasLastTriggeredTs returns a boolean if a field has been set.
func (o *MonitorStateGroup) HasLastTriggeredTs() bool {
	return o != nil && o.LastTriggeredTs != nil
}

// SetLastTriggeredTs gets a reference to the given int64 and assigns it to the LastTriggeredTs field.
func (o *MonitorStateGroup) SetLastTriggeredTs(v int64) {
	o.LastTriggeredTs = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MonitorStateGroup) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStateGroup) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MonitorStateGroup) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MonitorStateGroup) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MonitorStateGroup) GetStatus() MonitorOverallStates {
	if o == nil || o.Status == nil {
		var ret MonitorOverallStates
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStateGroup) GetStatusOk() (*MonitorOverallStates, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MonitorStateGroup) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given MonitorOverallStates and assigns it to the Status field.
func (o *MonitorStateGroup) SetStatus(v MonitorOverallStates) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorStateGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.LastNodataTs != nil {
		toSerialize["last_nodata_ts"] = o.LastNodataTs
	}
	if o.LastNotifiedTs != nil {
		toSerialize["last_notified_ts"] = o.LastNotifiedTs
	}
	if o.LastResolvedTs != nil {
		toSerialize["last_resolved_ts"] = o.LastResolvedTs
	}
	if o.LastTriggeredTs != nil {
		toSerialize["last_triggered_ts"] = o.LastTriggeredTs
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorStateGroup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LastNodataTs    *int64                `json:"last_nodata_ts,omitempty"`
		LastNotifiedTs  *int64                `json:"last_notified_ts,omitempty"`
		LastResolvedTs  *int64                `json:"last_resolved_ts,omitempty"`
		LastTriggeredTs *int64                `json:"last_triggered_ts,omitempty"`
		Name            *string               `json:"name,omitempty"`
		Status          *MonitorOverallStates `json:"status,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"last_nodata_ts", "last_notified_ts", "last_resolved_ts", "last_triggered_ts", "name", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.LastNodataTs = all.LastNodataTs
	o.LastNotifiedTs = all.LastNotifiedTs
	o.LastResolvedTs = all.LastResolvedTs
	o.LastTriggeredTs = all.LastTriggeredTs
	o.Name = all.Name
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
