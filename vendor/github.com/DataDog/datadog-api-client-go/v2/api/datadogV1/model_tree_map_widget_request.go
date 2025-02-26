// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreeMapWidgetRequest An updated treemap widget.
type TreeMapWidgetRequest struct {
	// List of formulas that operate on queries.
	Formulas []WidgetFormula `json:"formulas,omitempty"`
	// The widget metrics query.
	Q *string `json:"q,omitempty"`
	// List of queries that can be returned directly or used in formulas.
	Queries []FormulaAndFunctionQueryDefinition `json:"queries,omitempty"`
	// Timeseries, scalar, or event list response. Event list response formats are supported by Geomap widgets.
	ResponseFormat *FormulaAndFunctionResponseFormat `json:"response_format,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewTreeMapWidgetRequest instantiates a new TreeMapWidgetRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTreeMapWidgetRequest() *TreeMapWidgetRequest {
	this := TreeMapWidgetRequest{}
	return &this
}

// NewTreeMapWidgetRequestWithDefaults instantiates a new TreeMapWidgetRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTreeMapWidgetRequestWithDefaults() *TreeMapWidgetRequest {
	this := TreeMapWidgetRequest{}
	return &this
}

// GetFormulas returns the Formulas field value if set, zero value otherwise.
func (o *TreeMapWidgetRequest) GetFormulas() []WidgetFormula {
	if o == nil || o.Formulas == nil {
		var ret []WidgetFormula
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreeMapWidgetRequest) GetFormulasOk() (*[]WidgetFormula, bool) {
	if o == nil || o.Formulas == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// HasFormulas returns a boolean if a field has been set.
func (o *TreeMapWidgetRequest) HasFormulas() bool {
	return o != nil && o.Formulas != nil
}

// SetFormulas gets a reference to the given []WidgetFormula and assigns it to the Formulas field.
func (o *TreeMapWidgetRequest) SetFormulas(v []WidgetFormula) {
	o.Formulas = v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *TreeMapWidgetRequest) GetQ() string {
	if o == nil || o.Q == nil {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreeMapWidgetRequest) GetQOk() (*string, bool) {
	if o == nil || o.Q == nil {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *TreeMapWidgetRequest) HasQ() bool {
	return o != nil && o.Q != nil
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *TreeMapWidgetRequest) SetQ(v string) {
	o.Q = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *TreeMapWidgetRequest) GetQueries() []FormulaAndFunctionQueryDefinition {
	if o == nil || o.Queries == nil {
		var ret []FormulaAndFunctionQueryDefinition
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreeMapWidgetRequest) GetQueriesOk() (*[]FormulaAndFunctionQueryDefinition, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return &o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *TreeMapWidgetRequest) HasQueries() bool {
	return o != nil && o.Queries != nil
}

// SetQueries gets a reference to the given []FormulaAndFunctionQueryDefinition and assigns it to the Queries field.
func (o *TreeMapWidgetRequest) SetQueries(v []FormulaAndFunctionQueryDefinition) {
	o.Queries = v
}

// GetResponseFormat returns the ResponseFormat field value if set, zero value otherwise.
func (o *TreeMapWidgetRequest) GetResponseFormat() FormulaAndFunctionResponseFormat {
	if o == nil || o.ResponseFormat == nil {
		var ret FormulaAndFunctionResponseFormat
		return ret
	}
	return *o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreeMapWidgetRequest) GetResponseFormatOk() (*FormulaAndFunctionResponseFormat, bool) {
	if o == nil || o.ResponseFormat == nil {
		return nil, false
	}
	return o.ResponseFormat, true
}

// HasResponseFormat returns a boolean if a field has been set.
func (o *TreeMapWidgetRequest) HasResponseFormat() bool {
	return o != nil && o.ResponseFormat != nil
}

// SetResponseFormat gets a reference to the given FormulaAndFunctionResponseFormat and assigns it to the ResponseFormat field.
func (o *TreeMapWidgetRequest) SetResponseFormat(v FormulaAndFunctionResponseFormat) {
	o.ResponseFormat = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TreeMapWidgetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Formulas != nil {
		toSerialize["formulas"] = o.Formulas
	}
	if o.Q != nil {
		toSerialize["q"] = o.Q
	}
	if o.Queries != nil {
		toSerialize["queries"] = o.Queries
	}
	if o.ResponseFormat != nil {
		toSerialize["response_format"] = o.ResponseFormat
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TreeMapWidgetRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Formulas       []WidgetFormula                     `json:"formulas,omitempty"`
		Q              *string                             `json:"q,omitempty"`
		Queries        []FormulaAndFunctionQueryDefinition `json:"queries,omitempty"`
		ResponseFormat *FormulaAndFunctionResponseFormat   `json:"response_format,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"formulas", "q", "queries", "response_format"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Formulas = all.Formulas
	o.Q = all.Q
	o.Queries = all.Queries
	if all.ResponseFormat != nil && !all.ResponseFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.ResponseFormat = all.ResponseFormat
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
