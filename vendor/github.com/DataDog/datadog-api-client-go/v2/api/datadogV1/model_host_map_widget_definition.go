// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostMapWidgetDefinition The host map widget graphs any metric across your hosts using the same visualization available from the main Host Map page.
type HostMapWidgetDefinition struct {
	// List of custom links.
	CustomLinks []WidgetCustomLink `json:"custom_links,omitempty"`
	// List of tag prefixes to group by.
	Group []string `json:"group,omitempty"`
	// Whether to show the hosts that don’t fit in a group.
	NoGroupHosts *bool `json:"no_group_hosts,omitempty"`
	// Whether to show the hosts with no metrics.
	NoMetricHosts *bool `json:"no_metric_hosts,omitempty"`
	// Which type of node to use in the map.
	NodeType *WidgetNodeType `json:"node_type,omitempty"`
	// Notes on the title.
	Notes *string `json:"notes,omitempty"`
	// List of definitions.
	Requests HostMapWidgetDefinitionRequests `json:"requests"`
	// List of tags used to filter the map.
	Scope []string `json:"scope,omitempty"`
	// The style to apply to the widget.
	Style *HostMapWidgetDefinitionStyle `json:"style,omitempty"`
	// Title of the widget.
	Title *string `json:"title,omitempty"`
	// How to align the text on the widget.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Size of the title.
	TitleSize *string `json:"title_size,omitempty"`
	// Type of the host map widget.
	Type HostMapWidgetDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewHostMapWidgetDefinition instantiates a new HostMapWidgetDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHostMapWidgetDefinition(requests HostMapWidgetDefinitionRequests, typeVar HostMapWidgetDefinitionType) *HostMapWidgetDefinition {
	this := HostMapWidgetDefinition{}
	this.Requests = requests
	this.Type = typeVar
	return &this
}

// NewHostMapWidgetDefinitionWithDefaults instantiates a new HostMapWidgetDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHostMapWidgetDefinitionWithDefaults() *HostMapWidgetDefinition {
	this := HostMapWidgetDefinition{}
	var typeVar HostMapWidgetDefinitionType = HOSTMAPWIDGETDEFINITIONTYPE_HOSTMAP
	this.Type = typeVar
	return &this
}

// GetCustomLinks returns the CustomLinks field value if set, zero value otherwise.
func (o *HostMapWidgetDefinition) GetCustomLinks() []WidgetCustomLink {
	if o == nil || o.CustomLinks == nil {
		var ret []WidgetCustomLink
		return ret
	}
	return o.CustomLinks
}

// GetCustomLinksOk returns a tuple with the CustomLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinition) GetCustomLinksOk() (*[]WidgetCustomLink, bool) {
	if o == nil || o.CustomLinks == nil {
		return nil, false
	}
	return &o.CustomLinks, true
}

// HasCustomLinks returns a boolean if a field has been set.
func (o *HostMapWidgetDefinition) HasCustomLinks() bool {
	return o != nil && o.CustomLinks != nil
}

// SetCustomLinks gets a reference to the given []WidgetCustomLink and assigns it to the CustomLinks field.
func (o *HostMapWidgetDefinition) SetCustomLinks(v []WidgetCustomLink) {
	o.CustomLinks = v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *HostMapWidgetDefinition) GetGroup() []string {
	if o == nil || o.Group == nil {
		var ret []string
		return ret
	}
	return o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinition) GetGroupOk() (*[]string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return &o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *HostMapWidgetDefinition) HasGroup() bool {
	return o != nil && o.Group != nil
}

// SetGroup gets a reference to the given []string and assigns it to the Group field.
func (o *HostMapWidgetDefinition) SetGroup(v []string) {
	o.Group = v
}

// GetNoGroupHosts returns the NoGroupHosts field value if set, zero value otherwise.
func (o *HostMapWidgetDefinition) GetNoGroupHosts() bool {
	if o == nil || o.NoGroupHosts == nil {
		var ret bool
		return ret
	}
	return *o.NoGroupHosts
}

// GetNoGroupHostsOk returns a tuple with the NoGroupHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinition) GetNoGroupHostsOk() (*bool, bool) {
	if o == nil || o.NoGroupHosts == nil {
		return nil, false
	}
	return o.NoGroupHosts, true
}

// HasNoGroupHosts returns a boolean if a field has been set.
func (o *HostMapWidgetDefinition) HasNoGroupHosts() bool {
	return o != nil && o.NoGroupHosts != nil
}

// SetNoGroupHosts gets a reference to the given bool and assigns it to the NoGroupHosts field.
func (o *HostMapWidgetDefinition) SetNoGroupHosts(v bool) {
	o.NoGroupHosts = &v
}

// GetNoMetricHosts returns the NoMetricHosts field value if set, zero value otherwise.
func (o *HostMapWidgetDefinition) GetNoMetricHosts() bool {
	if o == nil || o.NoMetricHosts == nil {
		var ret bool
		return ret
	}
	return *o.NoMetricHosts
}

// GetNoMetricHostsOk returns a tuple with the NoMetricHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinition) GetNoMetricHostsOk() (*bool, bool) {
	if o == nil || o.NoMetricHosts == nil {
		return nil, false
	}
	return o.NoMetricHosts, true
}

// HasNoMetricHosts returns a boolean if a field has been set.
func (o *HostMapWidgetDefinition) HasNoMetricHosts() bool {
	return o != nil && o.NoMetricHosts != nil
}

// SetNoMetricHosts gets a reference to the given bool and assigns it to the NoMetricHosts field.
func (o *HostMapWidgetDefinition) SetNoMetricHosts(v bool) {
	o.NoMetricHosts = &v
}

// GetNodeType returns the NodeType field value if set, zero value otherwise.
func (o *HostMapWidgetDefinition) GetNodeType() WidgetNodeType {
	if o == nil || o.NodeType == nil {
		var ret WidgetNodeType
		return ret
	}
	return *o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinition) GetNodeTypeOk() (*WidgetNodeType, bool) {
	if o == nil || o.NodeType == nil {
		return nil, false
	}
	return o.NodeType, true
}

// HasNodeType returns a boolean if a field has been set.
func (o *HostMapWidgetDefinition) HasNodeType() bool {
	return o != nil && o.NodeType != nil
}

// SetNodeType gets a reference to the given WidgetNodeType and assigns it to the NodeType field.
func (o *HostMapWidgetDefinition) SetNodeType(v WidgetNodeType) {
	o.NodeType = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *HostMapWidgetDefinition) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinition) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *HostMapWidgetDefinition) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *HostMapWidgetDefinition) SetNotes(v string) {
	o.Notes = &v
}

// GetRequests returns the Requests field value.
func (o *HostMapWidgetDefinition) GetRequests() HostMapWidgetDefinitionRequests {
	if o == nil {
		var ret HostMapWidgetDefinitionRequests
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinition) GetRequestsOk() (*HostMapWidgetDefinitionRequests, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *HostMapWidgetDefinition) SetRequests(v HostMapWidgetDefinitionRequests) {
	o.Requests = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *HostMapWidgetDefinition) GetScope() []string {
	if o == nil || o.Scope == nil {
		var ret []string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinition) GetScopeOk() (*[]string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return &o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *HostMapWidgetDefinition) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given []string and assigns it to the Scope field.
func (o *HostMapWidgetDefinition) SetScope(v []string) {
	o.Scope = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *HostMapWidgetDefinition) GetStyle() HostMapWidgetDefinitionStyle {
	if o == nil || o.Style == nil {
		var ret HostMapWidgetDefinitionStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinition) GetStyleOk() (*HostMapWidgetDefinitionStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *HostMapWidgetDefinition) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given HostMapWidgetDefinitionStyle and assigns it to the Style field.
func (o *HostMapWidgetDefinition) SetStyle(v HostMapWidgetDefinitionStyle) {
	o.Style = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *HostMapWidgetDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *HostMapWidgetDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *HostMapWidgetDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *HostMapWidgetDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *HostMapWidgetDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *HostMapWidgetDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *HostMapWidgetDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *HostMapWidgetDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *HostMapWidgetDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *HostMapWidgetDefinition) GetType() HostMapWidgetDefinitionType {
	if o == nil {
		var ret HostMapWidgetDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinition) GetTypeOk() (*HostMapWidgetDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *HostMapWidgetDefinition) SetType(v HostMapWidgetDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HostMapWidgetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.CustomLinks != nil {
		toSerialize["custom_links"] = o.CustomLinks
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.NoGroupHosts != nil {
		toSerialize["no_group_hosts"] = o.NoGroupHosts
	}
	if o.NoMetricHosts != nil {
		toSerialize["no_metric_hosts"] = o.NoMetricHosts
	}
	if o.NodeType != nil {
		toSerialize["node_type"] = o.NodeType
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["requests"] = o.Requests
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.TitleAlign != nil {
		toSerialize["title_align"] = o.TitleAlign
	}
	if o.TitleSize != nil {
		toSerialize["title_size"] = o.TitleSize
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HostMapWidgetDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomLinks   []WidgetCustomLink               `json:"custom_links,omitempty"`
		Group         []string                         `json:"group,omitempty"`
		NoGroupHosts  *bool                            `json:"no_group_hosts,omitempty"`
		NoMetricHosts *bool                            `json:"no_metric_hosts,omitempty"`
		NodeType      *WidgetNodeType                  `json:"node_type,omitempty"`
		Notes         *string                          `json:"notes,omitempty"`
		Requests      *HostMapWidgetDefinitionRequests `json:"requests"`
		Scope         []string                         `json:"scope,omitempty"`
		Style         *HostMapWidgetDefinitionStyle    `json:"style,omitempty"`
		Title         *string                          `json:"title,omitempty"`
		TitleAlign    *WidgetTextAlign                 `json:"title_align,omitempty"`
		TitleSize     *string                          `json:"title_size,omitempty"`
		Type          *HostMapWidgetDefinitionType     `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"custom_links", "group", "no_group_hosts", "no_metric_hosts", "node_type", "notes", "requests", "scope", "style", "title", "title_align", "title_size", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CustomLinks = all.CustomLinks
	o.Group = all.Group
	o.NoGroupHosts = all.NoGroupHosts
	o.NoMetricHosts = all.NoMetricHosts
	if all.NodeType != nil && !all.NodeType.IsValid() {
		hasInvalidField = true
	} else {
		o.NodeType = all.NodeType
	}
	o.Notes = all.Notes
	if all.Requests.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Requests = *all.Requests
	o.Scope = all.Scope
	if all.Style != nil && all.Style.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Style = all.Style
	o.Title = all.Title
	if all.TitleAlign != nil && !all.TitleAlign.IsValid() {
		hasInvalidField = true
	} else {
		o.TitleAlign = all.TitleAlign
	}
	o.TitleSize = all.TitleSize
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
