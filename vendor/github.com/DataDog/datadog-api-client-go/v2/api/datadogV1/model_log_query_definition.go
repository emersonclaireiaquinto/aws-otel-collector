// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogQueryDefinition The log query.
type LogQueryDefinition struct {
	// Define computation for a log query.
	Compute *LogsQueryCompute `json:"compute,omitempty"`
	// List of tag prefixes to group by in the case of a cluster check.
	GroupBy []LogQueryDefinitionGroupBy `json:"group_by,omitempty"`
	// A coma separated-list of index names. Use "*" query all indexes at once. [Multiple Indexes](https://docs.datadoghq.com/logs/indexes/#multiple-indexes)
	Index *string `json:"index,omitempty"`
	// This field is mutually exclusive with `compute`.
	MultiCompute []LogsQueryCompute `json:"multi_compute,omitempty"`
	// The query being made on the logs.
	Search *LogQueryDefinitionSearch `json:"search,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewLogQueryDefinition instantiates a new LogQueryDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogQueryDefinition() *LogQueryDefinition {
	this := LogQueryDefinition{}
	return &this
}

// NewLogQueryDefinitionWithDefaults instantiates a new LogQueryDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogQueryDefinitionWithDefaults() *LogQueryDefinition {
	this := LogQueryDefinition{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *LogQueryDefinition) GetCompute() LogsQueryCompute {
	if o == nil || o.Compute == nil {
		var ret LogsQueryCompute
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogQueryDefinition) GetComputeOk() (*LogsQueryCompute, bool) {
	if o == nil || o.Compute == nil {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *LogQueryDefinition) HasCompute() bool {
	return o != nil && o.Compute != nil
}

// SetCompute gets a reference to the given LogsQueryCompute and assigns it to the Compute field.
func (o *LogQueryDefinition) SetCompute(v LogsQueryCompute) {
	o.Compute = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *LogQueryDefinition) GetGroupBy() []LogQueryDefinitionGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []LogQueryDefinitionGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogQueryDefinition) GetGroupByOk() (*[]LogQueryDefinitionGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *LogQueryDefinition) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []LogQueryDefinitionGroupBy and assigns it to the GroupBy field.
func (o *LogQueryDefinition) SetGroupBy(v []LogQueryDefinitionGroupBy) {
	o.GroupBy = v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *LogQueryDefinition) GetIndex() string {
	if o == nil || o.Index == nil {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogQueryDefinition) GetIndexOk() (*string, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *LogQueryDefinition) HasIndex() bool {
	return o != nil && o.Index != nil
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *LogQueryDefinition) SetIndex(v string) {
	o.Index = &v
}

// GetMultiCompute returns the MultiCompute field value if set, zero value otherwise.
func (o *LogQueryDefinition) GetMultiCompute() []LogsQueryCompute {
	if o == nil || o.MultiCompute == nil {
		var ret []LogsQueryCompute
		return ret
	}
	return o.MultiCompute
}

// GetMultiComputeOk returns a tuple with the MultiCompute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogQueryDefinition) GetMultiComputeOk() (*[]LogsQueryCompute, bool) {
	if o == nil || o.MultiCompute == nil {
		return nil, false
	}
	return &o.MultiCompute, true
}

// HasMultiCompute returns a boolean if a field has been set.
func (o *LogQueryDefinition) HasMultiCompute() bool {
	return o != nil && o.MultiCompute != nil
}

// SetMultiCompute gets a reference to the given []LogsQueryCompute and assigns it to the MultiCompute field.
func (o *LogQueryDefinition) SetMultiCompute(v []LogsQueryCompute) {
	o.MultiCompute = v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *LogQueryDefinition) GetSearch() LogQueryDefinitionSearch {
	if o == nil || o.Search == nil {
		var ret LogQueryDefinitionSearch
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogQueryDefinition) GetSearchOk() (*LogQueryDefinitionSearch, bool) {
	if o == nil || o.Search == nil {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *LogQueryDefinition) HasSearch() bool {
	return o != nil && o.Search != nil
}

// SetSearch gets a reference to the given LogQueryDefinitionSearch and assigns it to the Search field.
func (o *LogQueryDefinition) SetSearch(v LogQueryDefinitionSearch) {
	o.Search = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Compute != nil {
		toSerialize["compute"] = o.Compute
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.MultiCompute != nil {
		toSerialize["multi_compute"] = o.MultiCompute
	}
	if o.Search != nil {
		toSerialize["search"] = o.Search
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogQueryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute      *LogsQueryCompute           `json:"compute,omitempty"`
		GroupBy      []LogQueryDefinitionGroupBy `json:"group_by,omitempty"`
		Index        *string                     `json:"index,omitempty"`
		MultiCompute []LogsQueryCompute          `json:"multi_compute,omitempty"`
		Search       *LogQueryDefinitionSearch   `json:"search,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compute", "group_by", "index", "multi_compute", "search"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Compute != nil && all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = all.Compute
	o.GroupBy = all.GroupBy
	o.Index = all.Index
	o.MultiCompute = all.MultiCompute
	if all.Search != nil && all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = all.Search

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
