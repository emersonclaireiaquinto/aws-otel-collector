// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreeMapWidgetDefinition The treemap visualization enables you to display hierarchical and nested data. It is well suited for queries that describe part-whole relationships, such as resource usage by availability zone, data center, or team.
type TreeMapWidgetDefinition struct {
	// (deprecated) The attribute formerly used to determine color in the widget.
	// Deprecated
	ColorBy *TreeMapColorBy `json:"color_by,omitempty"`
	// List of custom links.
	CustomLinks []WidgetCustomLink `json:"custom_links,omitempty"`
	// (deprecated) The attribute formerly used to group elements in the widget.
	// Deprecated
	GroupBy *TreeMapGroupBy `json:"group_by,omitempty"`
	// List of treemap widget requests.
	Requests []TreeMapWidgetRequest `json:"requests"`
	// (deprecated) The attribute formerly used to determine size in the widget.
	// Deprecated
	SizeBy *TreeMapSizeBy `json:"size_by,omitempty"`
	// Time setting for the widget.
	Time *WidgetTime `json:"time,omitempty"`
	// Title of your widget.
	Title *string `json:"title,omitempty"`
	// Type of the treemap widget.
	Type TreeMapWidgetDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewTreeMapWidgetDefinition instantiates a new TreeMapWidgetDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTreeMapWidgetDefinition(requests []TreeMapWidgetRequest, typeVar TreeMapWidgetDefinitionType) *TreeMapWidgetDefinition {
	this := TreeMapWidgetDefinition{}
	var colorBy TreeMapColorBy = TREEMAPCOLORBY_USER
	this.ColorBy = &colorBy
	this.Requests = requests
	this.Type = typeVar
	return &this
}

// NewTreeMapWidgetDefinitionWithDefaults instantiates a new TreeMapWidgetDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTreeMapWidgetDefinitionWithDefaults() *TreeMapWidgetDefinition {
	this := TreeMapWidgetDefinition{}
	var colorBy TreeMapColorBy = TREEMAPCOLORBY_USER
	this.ColorBy = &colorBy
	var typeVar TreeMapWidgetDefinitionType = TREEMAPWIDGETDEFINITIONTYPE_TREEMAP
	this.Type = typeVar
	return &this
}

// GetColorBy returns the ColorBy field value if set, zero value otherwise.
// Deprecated
func (o *TreeMapWidgetDefinition) GetColorBy() TreeMapColorBy {
	if o == nil || o.ColorBy == nil {
		var ret TreeMapColorBy
		return ret
	}
	return *o.ColorBy
}

// GetColorByOk returns a tuple with the ColorBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *TreeMapWidgetDefinition) GetColorByOk() (*TreeMapColorBy, bool) {
	if o == nil || o.ColorBy == nil {
		return nil, false
	}
	return o.ColorBy, true
}

// HasColorBy returns a boolean if a field has been set.
func (o *TreeMapWidgetDefinition) HasColorBy() bool {
	return o != nil && o.ColorBy != nil
}

// SetColorBy gets a reference to the given TreeMapColorBy and assigns it to the ColorBy field.
// Deprecated
func (o *TreeMapWidgetDefinition) SetColorBy(v TreeMapColorBy) {
	o.ColorBy = &v
}

// GetCustomLinks returns the CustomLinks field value if set, zero value otherwise.
func (o *TreeMapWidgetDefinition) GetCustomLinks() []WidgetCustomLink {
	if o == nil || o.CustomLinks == nil {
		var ret []WidgetCustomLink
		return ret
	}
	return o.CustomLinks
}

// GetCustomLinksOk returns a tuple with the CustomLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreeMapWidgetDefinition) GetCustomLinksOk() (*[]WidgetCustomLink, bool) {
	if o == nil || o.CustomLinks == nil {
		return nil, false
	}
	return &o.CustomLinks, true
}

// HasCustomLinks returns a boolean if a field has been set.
func (o *TreeMapWidgetDefinition) HasCustomLinks() bool {
	return o != nil && o.CustomLinks != nil
}

// SetCustomLinks gets a reference to the given []WidgetCustomLink and assigns it to the CustomLinks field.
func (o *TreeMapWidgetDefinition) SetCustomLinks(v []WidgetCustomLink) {
	o.CustomLinks = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
// Deprecated
func (o *TreeMapWidgetDefinition) GetGroupBy() TreeMapGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret TreeMapGroupBy
		return ret
	}
	return *o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *TreeMapWidgetDefinition) GetGroupByOk() (*TreeMapGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *TreeMapWidgetDefinition) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given TreeMapGroupBy and assigns it to the GroupBy field.
// Deprecated
func (o *TreeMapWidgetDefinition) SetGroupBy(v TreeMapGroupBy) {
	o.GroupBy = &v
}

// GetRequests returns the Requests field value.
func (o *TreeMapWidgetDefinition) GetRequests() []TreeMapWidgetRequest {
	if o == nil {
		var ret []TreeMapWidgetRequest
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *TreeMapWidgetDefinition) GetRequestsOk() (*[]TreeMapWidgetRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *TreeMapWidgetDefinition) SetRequests(v []TreeMapWidgetRequest) {
	o.Requests = v
}

// GetSizeBy returns the SizeBy field value if set, zero value otherwise.
// Deprecated
func (o *TreeMapWidgetDefinition) GetSizeBy() TreeMapSizeBy {
	if o == nil || o.SizeBy == nil {
		var ret TreeMapSizeBy
		return ret
	}
	return *o.SizeBy
}

// GetSizeByOk returns a tuple with the SizeBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *TreeMapWidgetDefinition) GetSizeByOk() (*TreeMapSizeBy, bool) {
	if o == nil || o.SizeBy == nil {
		return nil, false
	}
	return o.SizeBy, true
}

// HasSizeBy returns a boolean if a field has been set.
func (o *TreeMapWidgetDefinition) HasSizeBy() bool {
	return o != nil && o.SizeBy != nil
}

// SetSizeBy gets a reference to the given TreeMapSizeBy and assigns it to the SizeBy field.
// Deprecated
func (o *TreeMapWidgetDefinition) SetSizeBy(v TreeMapSizeBy) {
	o.SizeBy = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *TreeMapWidgetDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreeMapWidgetDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *TreeMapWidgetDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *TreeMapWidgetDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *TreeMapWidgetDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreeMapWidgetDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *TreeMapWidgetDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *TreeMapWidgetDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value.
func (o *TreeMapWidgetDefinition) GetType() TreeMapWidgetDefinitionType {
	if o == nil {
		var ret TreeMapWidgetDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TreeMapWidgetDefinition) GetTypeOk() (*TreeMapWidgetDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TreeMapWidgetDefinition) SetType(v TreeMapWidgetDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TreeMapWidgetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ColorBy != nil {
		toSerialize["color_by"] = o.ColorBy
	}
	if o.CustomLinks != nil {
		toSerialize["custom_links"] = o.CustomLinks
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["requests"] = o.Requests
	if o.SizeBy != nil {
		toSerialize["size_by"] = o.SizeBy
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TreeMapWidgetDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColorBy     *TreeMapColorBy              `json:"color_by,omitempty"`
		CustomLinks []WidgetCustomLink           `json:"custom_links,omitempty"`
		GroupBy     *TreeMapGroupBy              `json:"group_by,omitempty"`
		Requests    *[]TreeMapWidgetRequest      `json:"requests"`
		SizeBy      *TreeMapSizeBy               `json:"size_by,omitempty"`
		Time        *WidgetTime                  `json:"time,omitempty"`
		Title       *string                      `json:"title,omitempty"`
		Type        *TreeMapWidgetDefinitionType `json:"type"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Requests == nil {
		return fmt.Errorf("required field requests missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"color_by", "custom_links", "group_by", "requests", "size_by", "time", "title", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ColorBy != nil && !all.ColorBy.IsValid() {
		hasInvalidField = true
	} else {
		o.ColorBy = all.ColorBy
	}
	o.CustomLinks = all.CustomLinks
	if all.GroupBy != nil && !all.GroupBy.IsValid() {
		hasInvalidField = true
	} else {
		o.GroupBy = all.GroupBy
	}
	o.Requests = *all.Requests
	if all.SizeBy != nil && !all.SizeBy.IsValid() {
		hasInvalidField = true
	} else {
		o.SizeBy = all.SizeBy
	}
	if all.Time != nil && all.Time.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Time = all.Time
	o.Title = all.Title
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
