// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/goccy/go-json"
)

// NotebookCellCreateRequest The description of a notebook cell create request.
type NotebookCellCreateRequest struct {
	// The attributes of a notebook cell in create cell request. Valid cell types are `markdown`, `timeseries`, `toplist`, `heatmap`, `distribution`,
	// `log_stream`. [More information on each graph visualization type.](https://docs.datadoghq.com/dashboards/widgets/)
	Attributes NotebookCellCreateRequestAttributes `json:"attributes"`
	// Type of the Notebook Cell resource.
	Type NotebookCellResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewNotebookCellCreateRequest instantiates a new NotebookCellCreateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotebookCellCreateRequest(attributes NotebookCellCreateRequestAttributes, typeVar NotebookCellResourceType) *NotebookCellCreateRequest {
	this := NotebookCellCreateRequest{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewNotebookCellCreateRequestWithDefaults instantiates a new NotebookCellCreateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotebookCellCreateRequestWithDefaults() *NotebookCellCreateRequest {
	this := NotebookCellCreateRequest{}
	var typeVar NotebookCellResourceType = NOTEBOOKCELLRESOURCETYPE_NOTEBOOK_CELLS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *NotebookCellCreateRequest) GetAttributes() NotebookCellCreateRequestAttributes {
	if o == nil {
		var ret NotebookCellCreateRequestAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *NotebookCellCreateRequest) GetAttributesOk() (*NotebookCellCreateRequestAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *NotebookCellCreateRequest) SetAttributes(v NotebookCellCreateRequestAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *NotebookCellCreateRequest) GetType() NotebookCellResourceType {
	if o == nil {
		var ret NotebookCellResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotebookCellCreateRequest) GetTypeOk() (*NotebookCellResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *NotebookCellCreateRequest) SetType(v NotebookCellResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotebookCellCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotebookCellCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *NotebookCellCreateRequestAttributes `json:"attributes"`
		Type       *NotebookCellResourceType            `json:"type"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Attributes = *all.Attributes
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
