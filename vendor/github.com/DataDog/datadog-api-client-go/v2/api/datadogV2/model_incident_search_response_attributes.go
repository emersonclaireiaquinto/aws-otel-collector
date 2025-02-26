// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentSearchResponseAttributes Attributes returned by an incident search.
type IncidentSearchResponseAttributes struct {
	// Facet data for incidents returned by a search query.
	Facets IncidentSearchResponseFacetsData `json:"facets"`
	// Incidents returned by the search.
	Incidents []IncidentSearchResponseIncidentsData `json:"incidents"`
	// Number of incidents returned by the search.
	Total int32 `json:"total"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewIncidentSearchResponseAttributes instantiates a new IncidentSearchResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentSearchResponseAttributes(facets IncidentSearchResponseFacetsData, incidents []IncidentSearchResponseIncidentsData, total int32) *IncidentSearchResponseAttributes {
	this := IncidentSearchResponseAttributes{}
	this.Facets = facets
	this.Incidents = incidents
	this.Total = total
	return &this
}

// NewIncidentSearchResponseAttributesWithDefaults instantiates a new IncidentSearchResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentSearchResponseAttributesWithDefaults() *IncidentSearchResponseAttributes {
	this := IncidentSearchResponseAttributes{}
	return &this
}

// GetFacets returns the Facets field value.
func (o *IncidentSearchResponseAttributes) GetFacets() IncidentSearchResponseFacetsData {
	if o == nil {
		var ret IncidentSearchResponseFacetsData
		return ret
	}
	return o.Facets
}

// GetFacetsOk returns a tuple with the Facets field value
// and a boolean to check if the value has been set.
func (o *IncidentSearchResponseAttributes) GetFacetsOk() (*IncidentSearchResponseFacetsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Facets, true
}

// SetFacets sets field value.
func (o *IncidentSearchResponseAttributes) SetFacets(v IncidentSearchResponseFacetsData) {
	o.Facets = v
}

// GetIncidents returns the Incidents field value.
func (o *IncidentSearchResponseAttributes) GetIncidents() []IncidentSearchResponseIncidentsData {
	if o == nil {
		var ret []IncidentSearchResponseIncidentsData
		return ret
	}
	return o.Incidents
}

// GetIncidentsOk returns a tuple with the Incidents field value
// and a boolean to check if the value has been set.
func (o *IncidentSearchResponseAttributes) GetIncidentsOk() (*[]IncidentSearchResponseIncidentsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Incidents, true
}

// SetIncidents sets field value.
func (o *IncidentSearchResponseAttributes) SetIncidents(v []IncidentSearchResponseIncidentsData) {
	o.Incidents = v
}

// GetTotal returns the Total field value.
func (o *IncidentSearchResponseAttributes) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *IncidentSearchResponseAttributes) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *IncidentSearchResponseAttributes) SetTotal(v int32) {
	o.Total = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentSearchResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["facets"] = o.Facets
	toSerialize["incidents"] = o.Incidents
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentSearchResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Facets    *IncidentSearchResponseFacetsData      `json:"facets"`
		Incidents *[]IncidentSearchResponseIncidentsData `json:"incidents"`
		Total     *int32                                 `json:"total"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Facets == nil {
		return fmt.Errorf("required field facets missing")
	}
	if all.Incidents == nil {
		return fmt.Errorf("required field incidents missing")
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"facets", "incidents", "total"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Facets.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Facets = *all.Facets
	o.Incidents = *all.Incidents
	o.Total = *all.Total

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
