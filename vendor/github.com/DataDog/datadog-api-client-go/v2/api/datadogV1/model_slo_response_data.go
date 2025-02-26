// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOResponseData A service level objective object includes a service level indicator, thresholds
// for one or more timeframes, and metadata (`name`, `description`, `tags`, etc.).
type SLOResponseData struct {
	// A list of SLO monitors IDs that reference this SLO. This field is returned only when `with_configured_alert_ids` parameter is true in query.
	ConfiguredAlertIds []int64 `json:"configured_alert_ids,omitempty"`
	// Creation timestamp (UNIX time in seconds)
	//
	// Always included in service level objective responses.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Object describing the creator of the shared element.
	Creator *Creator `json:"creator,omitempty"`
	// A user-defined description of the service level objective.
	//
	// Always included in service level objective responses (but may be `null`).
	// Optional in create/update requests.
	Description datadog.NullableString `json:"description,omitempty"`
	// A list of (up to 20) monitor groups that narrow the scope of a monitor service level objective.
	//
	// Included in service level objective responses if it is not empty. Optional in
	// create/update requests for monitor service level objectives, but may only be
	// used when then length of the `monitor_ids` field is one.
	Groups []string `json:"groups,omitempty"`
	// A unique identifier for the service level objective object.
	//
	// Always included in service level objective responses.
	Id *string `json:"id,omitempty"`
	// Modification timestamp (UNIX time in seconds)
	//
	// Always included in service level objective responses.
	ModifiedAt *int64 `json:"modified_at,omitempty"`
	// A list of monitor ids that defines the scope of a monitor service level
	// objective. **Required if type is `monitor`**.
	MonitorIds []int64 `json:"monitor_ids,omitempty"`
	// The union of monitor tags for all monitors referenced by the `monitor_ids`
	// field.
	// Always included in service level objective responses for monitor service level
	// objectives (but may be empty). Ignored in create/update requests. Does not
	// affect which monitors are included in the service level objective (that is
	// determined entirely by the `monitor_ids` field).
	MonitorTags []string `json:"monitor_tags,omitempty"`
	// The name of the service level objective object.
	Name *string `json:"name,omitempty"`
	// A metric-based SLO. **Required if type is `metric`**. Note that Datadog only allows the sum by aggregator
	// to be used because this will sum up all request counts instead of averaging them, or taking the max or
	// min of all of those requests.
	Query *ServiceLevelObjectiveQuery `json:"query,omitempty"`
	// A list of tags associated with this service level objective.
	// Always included in service level objective responses (but may be empty).
	// Optional in create/update requests.
	Tags []string `json:"tags,omitempty"`
	// The target threshold such that when the service level indicator is above this
	// threshold over the given timeframe, the objective is being met.
	TargetThreshold *float64 `json:"target_threshold,omitempty"`
	// The thresholds (timeframes and associated targets) for this service level
	// objective object.
	Thresholds []SLOThreshold `json:"thresholds,omitempty"`
	// The SLO time window options.
	Timeframe *SLOTimeframe `json:"timeframe,omitempty"`
	// The type of the service level objective.
	Type *SLOType `json:"type,omitempty"`
	// The optional warning threshold such that when the service level indicator is
	// below this value for the given threshold, but above the target threshold, the
	// objective appears in a "warning" state. This value must be greater than the target
	// threshold.
	WarningThreshold *float64 `json:"warning_threshold,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSLOResponseData instantiates a new SLOResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSLOResponseData() *SLOResponseData {
	this := SLOResponseData{}
	return &this
}

// NewSLOResponseDataWithDefaults instantiates a new SLOResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSLOResponseDataWithDefaults() *SLOResponseData {
	this := SLOResponseData{}
	return &this
}

// GetConfiguredAlertIds returns the ConfiguredAlertIds field value if set, zero value otherwise.
func (o *SLOResponseData) GetConfiguredAlertIds() []int64 {
	if o == nil || o.ConfiguredAlertIds == nil {
		var ret []int64
		return ret
	}
	return o.ConfiguredAlertIds
}

// GetConfiguredAlertIdsOk returns a tuple with the ConfiguredAlertIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOResponseData) GetConfiguredAlertIdsOk() (*[]int64, bool) {
	if o == nil || o.ConfiguredAlertIds == nil {
		return nil, false
	}
	return &o.ConfiguredAlertIds, true
}

// HasConfiguredAlertIds returns a boolean if a field has been set.
func (o *SLOResponseData) HasConfiguredAlertIds() bool {
	return o != nil && o.ConfiguredAlertIds != nil
}

// SetConfiguredAlertIds gets a reference to the given []int64 and assigns it to the ConfiguredAlertIds field.
func (o *SLOResponseData) SetConfiguredAlertIds(v []int64) {
	o.ConfiguredAlertIds = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SLOResponseData) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOResponseData) GetCreatedAtOk() (*int64, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SLOResponseData) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *SLOResponseData) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *SLOResponseData) GetCreator() Creator {
	if o == nil || o.Creator == nil {
		var ret Creator
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOResponseData) GetCreatorOk() (*Creator, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *SLOResponseData) HasCreator() bool {
	return o != nil && o.Creator != nil
}

// SetCreator gets a reference to the given Creator and assigns it to the Creator field.
func (o *SLOResponseData) SetCreator(v Creator) {
	o.Creator = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SLOResponseData) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SLOResponseData) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SLOResponseData) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *SLOResponseData) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *SLOResponseData) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *SLOResponseData) UnsetDescription() {
	o.Description.Unset()
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *SLOResponseData) GetGroups() []string {
	if o == nil || o.Groups == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOResponseData) GetGroupsOk() (*[]string, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return &o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *SLOResponseData) HasGroups() bool {
	return o != nil && o.Groups != nil
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *SLOResponseData) SetGroups(v []string) {
	o.Groups = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SLOResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SLOResponseData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SLOResponseData) SetId(v string) {
	o.Id = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *SLOResponseData) GetModifiedAt() int64 {
	if o == nil || o.ModifiedAt == nil {
		var ret int64
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOResponseData) GetModifiedAtOk() (*int64, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *SLOResponseData) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given int64 and assigns it to the ModifiedAt field.
func (o *SLOResponseData) SetModifiedAt(v int64) {
	o.ModifiedAt = &v
}

// GetMonitorIds returns the MonitorIds field value if set, zero value otherwise.
func (o *SLOResponseData) GetMonitorIds() []int64 {
	if o == nil || o.MonitorIds == nil {
		var ret []int64
		return ret
	}
	return o.MonitorIds
}

// GetMonitorIdsOk returns a tuple with the MonitorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOResponseData) GetMonitorIdsOk() (*[]int64, bool) {
	if o == nil || o.MonitorIds == nil {
		return nil, false
	}
	return &o.MonitorIds, true
}

// HasMonitorIds returns a boolean if a field has been set.
func (o *SLOResponseData) HasMonitorIds() bool {
	return o != nil && o.MonitorIds != nil
}

// SetMonitorIds gets a reference to the given []int64 and assigns it to the MonitorIds field.
func (o *SLOResponseData) SetMonitorIds(v []int64) {
	o.MonitorIds = v
}

// GetMonitorTags returns the MonitorTags field value if set, zero value otherwise.
func (o *SLOResponseData) GetMonitorTags() []string {
	if o == nil || o.MonitorTags == nil {
		var ret []string
		return ret
	}
	return o.MonitorTags
}

// GetMonitorTagsOk returns a tuple with the MonitorTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOResponseData) GetMonitorTagsOk() (*[]string, bool) {
	if o == nil || o.MonitorTags == nil {
		return nil, false
	}
	return &o.MonitorTags, true
}

// HasMonitorTags returns a boolean if a field has been set.
func (o *SLOResponseData) HasMonitorTags() bool {
	return o != nil && o.MonitorTags != nil
}

// SetMonitorTags gets a reference to the given []string and assigns it to the MonitorTags field.
func (o *SLOResponseData) SetMonitorTags(v []string) {
	o.MonitorTags = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SLOResponseData) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOResponseData) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SLOResponseData) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SLOResponseData) SetName(v string) {
	o.Name = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *SLOResponseData) GetQuery() ServiceLevelObjectiveQuery {
	if o == nil || o.Query == nil {
		var ret ServiceLevelObjectiveQuery
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOResponseData) GetQueryOk() (*ServiceLevelObjectiveQuery, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *SLOResponseData) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given ServiceLevelObjectiveQuery and assigns it to the Query field.
func (o *SLOResponseData) SetQuery(v ServiceLevelObjectiveQuery) {
	o.Query = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SLOResponseData) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOResponseData) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SLOResponseData) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SLOResponseData) SetTags(v []string) {
	o.Tags = v
}

// GetTargetThreshold returns the TargetThreshold field value if set, zero value otherwise.
func (o *SLOResponseData) GetTargetThreshold() float64 {
	if o == nil || o.TargetThreshold == nil {
		var ret float64
		return ret
	}
	return *o.TargetThreshold
}

// GetTargetThresholdOk returns a tuple with the TargetThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOResponseData) GetTargetThresholdOk() (*float64, bool) {
	if o == nil || o.TargetThreshold == nil {
		return nil, false
	}
	return o.TargetThreshold, true
}

// HasTargetThreshold returns a boolean if a field has been set.
func (o *SLOResponseData) HasTargetThreshold() bool {
	return o != nil && o.TargetThreshold != nil
}

// SetTargetThreshold gets a reference to the given float64 and assigns it to the TargetThreshold field.
func (o *SLOResponseData) SetTargetThreshold(v float64) {
	o.TargetThreshold = &v
}

// GetThresholds returns the Thresholds field value if set, zero value otherwise.
func (o *SLOResponseData) GetThresholds() []SLOThreshold {
	if o == nil || o.Thresholds == nil {
		var ret []SLOThreshold
		return ret
	}
	return o.Thresholds
}

// GetThresholdsOk returns a tuple with the Thresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOResponseData) GetThresholdsOk() (*[]SLOThreshold, bool) {
	if o == nil || o.Thresholds == nil {
		return nil, false
	}
	return &o.Thresholds, true
}

// HasThresholds returns a boolean if a field has been set.
func (o *SLOResponseData) HasThresholds() bool {
	return o != nil && o.Thresholds != nil
}

// SetThresholds gets a reference to the given []SLOThreshold and assigns it to the Thresholds field.
func (o *SLOResponseData) SetThresholds(v []SLOThreshold) {
	o.Thresholds = v
}

// GetTimeframe returns the Timeframe field value if set, zero value otherwise.
func (o *SLOResponseData) GetTimeframe() SLOTimeframe {
	if o == nil || o.Timeframe == nil {
		var ret SLOTimeframe
		return ret
	}
	return *o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOResponseData) GetTimeframeOk() (*SLOTimeframe, bool) {
	if o == nil || o.Timeframe == nil {
		return nil, false
	}
	return o.Timeframe, true
}

// HasTimeframe returns a boolean if a field has been set.
func (o *SLOResponseData) HasTimeframe() bool {
	return o != nil && o.Timeframe != nil
}

// SetTimeframe gets a reference to the given SLOTimeframe and assigns it to the Timeframe field.
func (o *SLOResponseData) SetTimeframe(v SLOTimeframe) {
	o.Timeframe = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SLOResponseData) GetType() SLOType {
	if o == nil || o.Type == nil {
		var ret SLOType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOResponseData) GetTypeOk() (*SLOType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SLOResponseData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given SLOType and assigns it to the Type field.
func (o *SLOResponseData) SetType(v SLOType) {
	o.Type = &v
}

// GetWarningThreshold returns the WarningThreshold field value if set, zero value otherwise.
func (o *SLOResponseData) GetWarningThreshold() float64 {
	if o == nil || o.WarningThreshold == nil {
		var ret float64
		return ret
	}
	return *o.WarningThreshold
}

// GetWarningThresholdOk returns a tuple with the WarningThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOResponseData) GetWarningThresholdOk() (*float64, bool) {
	if o == nil || o.WarningThreshold == nil {
		return nil, false
	}
	return o.WarningThreshold, true
}

// HasWarningThreshold returns a boolean if a field has been set.
func (o *SLOResponseData) HasWarningThreshold() bool {
	return o != nil && o.WarningThreshold != nil
}

// SetWarningThreshold gets a reference to the given float64 and assigns it to the WarningThreshold field.
func (o *SLOResponseData) SetWarningThreshold(v float64) {
	o.WarningThreshold = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SLOResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ConfiguredAlertIds != nil {
		toSerialize["configured_alert_ids"] = o.ConfiguredAlertIds
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if o.MonitorIds != nil {
		toSerialize["monitor_ids"] = o.MonitorIds
	}
	if o.MonitorTags != nil {
		toSerialize["monitor_tags"] = o.MonitorTags
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TargetThreshold != nil {
		toSerialize["target_threshold"] = o.TargetThreshold
	}
	if o.Thresholds != nil {
		toSerialize["thresholds"] = o.Thresholds
	}
	if o.Timeframe != nil {
		toSerialize["timeframe"] = o.Timeframe
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.WarningThreshold != nil {
		toSerialize["warning_threshold"] = o.WarningThreshold
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SLOResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfiguredAlertIds []int64                     `json:"configured_alert_ids,omitempty"`
		CreatedAt          *int64                      `json:"created_at,omitempty"`
		Creator            *Creator                    `json:"creator,omitempty"`
		Description        datadog.NullableString      `json:"description,omitempty"`
		Groups             []string                    `json:"groups,omitempty"`
		Id                 *string                     `json:"id,omitempty"`
		ModifiedAt         *int64                      `json:"modified_at,omitempty"`
		MonitorIds         []int64                     `json:"monitor_ids,omitempty"`
		MonitorTags        []string                    `json:"monitor_tags,omitempty"`
		Name               *string                     `json:"name,omitempty"`
		Query              *ServiceLevelObjectiveQuery `json:"query,omitempty"`
		Tags               []string                    `json:"tags,omitempty"`
		TargetThreshold    *float64                    `json:"target_threshold,omitempty"`
		Thresholds         []SLOThreshold              `json:"thresholds,omitempty"`
		Timeframe          *SLOTimeframe               `json:"timeframe,omitempty"`
		Type               *SLOType                    `json:"type,omitempty"`
		WarningThreshold   *float64                    `json:"warning_threshold,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"configured_alert_ids", "created_at", "creator", "description", "groups", "id", "modified_at", "monitor_ids", "monitor_tags", "name", "query", "tags", "target_threshold", "thresholds", "timeframe", "type", "warning_threshold"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ConfiguredAlertIds = all.ConfiguredAlertIds
	o.CreatedAt = all.CreatedAt
	if all.Creator != nil && all.Creator.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Creator = all.Creator
	o.Description = all.Description
	o.Groups = all.Groups
	o.Id = all.Id
	o.ModifiedAt = all.ModifiedAt
	o.MonitorIds = all.MonitorIds
	o.MonitorTags = all.MonitorTags
	o.Name = all.Name
	if all.Query != nil && all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Query = all.Query
	o.Tags = all.Tags
	o.TargetThreshold = all.TargetThreshold
	o.Thresholds = all.Thresholds
	if all.Timeframe != nil && !all.Timeframe.IsValid() {
		hasInvalidField = true
	} else {
		o.Timeframe = all.Timeframe
	}
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	o.WarningThreshold = all.WarningThreshold

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
