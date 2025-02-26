// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesResponseSeries
type TimeseriesResponseSeries struct {
	// List of tags that apply to a single response value.
	GroupTags []string `json:"group_tags,omitempty"`
	// The index of the query in the "formulas" array (or "queries" array if no "formulas" was specified).
	QueryIndex *int32 `json:"query_index,omitempty"`
	// Detailed information about the unit.
	// The first element describes the "primary unit" (for example, `bytes` in `bytes per second`).
	// The second element describes the "per unit" (for example, `second` in `bytes per second`).
	// If the second element is not present, the API returns null.
	Unit []Unit `json:"unit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewTimeseriesResponseSeries instantiates a new TimeseriesResponseSeries object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesResponseSeries() *TimeseriesResponseSeries {
	this := TimeseriesResponseSeries{}
	return &this
}

// NewTimeseriesResponseSeriesWithDefaults instantiates a new TimeseriesResponseSeries object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesResponseSeriesWithDefaults() *TimeseriesResponseSeries {
	this := TimeseriesResponseSeries{}
	return &this
}

// GetGroupTags returns the GroupTags field value if set, zero value otherwise.
func (o *TimeseriesResponseSeries) GetGroupTags() []string {
	if o == nil || o.GroupTags == nil {
		var ret []string
		return ret
	}
	return o.GroupTags
}

// GetGroupTagsOk returns a tuple with the GroupTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesResponseSeries) GetGroupTagsOk() (*[]string, bool) {
	if o == nil || o.GroupTags == nil {
		return nil, false
	}
	return &o.GroupTags, true
}

// HasGroupTags returns a boolean if a field has been set.
func (o *TimeseriesResponseSeries) HasGroupTags() bool {
	return o != nil && o.GroupTags != nil
}

// SetGroupTags gets a reference to the given []string and assigns it to the GroupTags field.
func (o *TimeseriesResponseSeries) SetGroupTags(v []string) {
	o.GroupTags = v
}

// GetQueryIndex returns the QueryIndex field value if set, zero value otherwise.
func (o *TimeseriesResponseSeries) GetQueryIndex() int32 {
	if o == nil || o.QueryIndex == nil {
		var ret int32
		return ret
	}
	return *o.QueryIndex
}

// GetQueryIndexOk returns a tuple with the QueryIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesResponseSeries) GetQueryIndexOk() (*int32, bool) {
	if o == nil || o.QueryIndex == nil {
		return nil, false
	}
	return o.QueryIndex, true
}

// HasQueryIndex returns a boolean if a field has been set.
func (o *TimeseriesResponseSeries) HasQueryIndex() bool {
	return o != nil && o.QueryIndex != nil
}

// SetQueryIndex gets a reference to the given int32 and assigns it to the QueryIndex field.
func (o *TimeseriesResponseSeries) SetQueryIndex(v int32) {
	o.QueryIndex = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *TimeseriesResponseSeries) GetUnit() []Unit {
	if o == nil || o.Unit == nil {
		var ret []Unit
		return ret
	}
	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesResponseSeries) GetUnitOk() (*[]Unit, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return &o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *TimeseriesResponseSeries) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given []Unit and assigns it to the Unit field.
func (o *TimeseriesResponseSeries) SetUnit(v []Unit) {
	o.Unit = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesResponseSeries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.GroupTags != nil {
		toSerialize["group_tags"] = o.GroupTags
	}
	if o.QueryIndex != nil {
		toSerialize["query_index"] = o.QueryIndex
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesResponseSeries) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GroupTags  []string `json:"group_tags,omitempty"`
		QueryIndex *int32   `json:"query_index,omitempty"`
		Unit       []Unit   `json:"unit,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"group_tags", "query_index", "unit"})
	} else {
		return err
	}
	o.GroupTags = all.GroupTags
	o.QueryIndex = all.QueryIndex
	o.Unit = all.Unit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
