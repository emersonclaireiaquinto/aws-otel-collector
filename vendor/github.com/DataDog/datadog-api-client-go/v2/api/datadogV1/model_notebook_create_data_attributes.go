// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookCreateDataAttributes The data attributes of a notebook.
type NotebookCreateDataAttributes struct {
	// List of cells to display in the notebook.
	Cells []NotebookCellCreateRequest `json:"cells"`
	// Metadata associated with the notebook.
	Metadata *NotebookMetadata `json:"metadata,omitempty"`
	// The name of the notebook.
	Name string `json:"name"`
	// Publication status of the notebook. For now, always "published".
	Status *NotebookStatus `json:"status,omitempty"`
	// Notebook global timeframe.
	Time NotebookGlobalTime `json:"time"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewNotebookCreateDataAttributes instantiates a new NotebookCreateDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotebookCreateDataAttributes(cells []NotebookCellCreateRequest, name string, time NotebookGlobalTime) *NotebookCreateDataAttributes {
	this := NotebookCreateDataAttributes{}
	this.Cells = cells
	this.Name = name
	var status NotebookStatus = NOTEBOOKSTATUS_PUBLISHED
	this.Status = &status
	this.Time = time
	return &this
}

// NewNotebookCreateDataAttributesWithDefaults instantiates a new NotebookCreateDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotebookCreateDataAttributesWithDefaults() *NotebookCreateDataAttributes {
	this := NotebookCreateDataAttributes{}
	var status NotebookStatus = NOTEBOOKSTATUS_PUBLISHED
	this.Status = &status
	return &this
}

// GetCells returns the Cells field value.
func (o *NotebookCreateDataAttributes) GetCells() []NotebookCellCreateRequest {
	if o == nil {
		var ret []NotebookCellCreateRequest
		return ret
	}
	return o.Cells
}

// GetCellsOk returns a tuple with the Cells field value
// and a boolean to check if the value has been set.
func (o *NotebookCreateDataAttributes) GetCellsOk() (*[]NotebookCellCreateRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cells, true
}

// SetCells sets field value.
func (o *NotebookCreateDataAttributes) SetCells(v []NotebookCellCreateRequest) {
	o.Cells = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *NotebookCreateDataAttributes) GetMetadata() NotebookMetadata {
	if o == nil || o.Metadata == nil {
		var ret NotebookMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookCreateDataAttributes) GetMetadataOk() (*NotebookMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *NotebookCreateDataAttributes) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given NotebookMetadata and assigns it to the Metadata field.
func (o *NotebookCreateDataAttributes) SetMetadata(v NotebookMetadata) {
	o.Metadata = &v
}

// GetName returns the Name field value.
func (o *NotebookCreateDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NotebookCreateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *NotebookCreateDataAttributes) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NotebookCreateDataAttributes) GetStatus() NotebookStatus {
	if o == nil || o.Status == nil {
		var ret NotebookStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookCreateDataAttributes) GetStatusOk() (*NotebookStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NotebookCreateDataAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given NotebookStatus and assigns it to the Status field.
func (o *NotebookCreateDataAttributes) SetStatus(v NotebookStatus) {
	o.Status = &v
}

// GetTime returns the Time field value.
func (o *NotebookCreateDataAttributes) GetTime() NotebookGlobalTime {
	if o == nil {
		var ret NotebookGlobalTime
		return ret
	}
	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *NotebookCreateDataAttributes) GetTimeOk() (*NotebookGlobalTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value.
func (o *NotebookCreateDataAttributes) SetTime(v NotebookGlobalTime) {
	o.Time = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotebookCreateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["cells"] = o.Cells
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["name"] = o.Name
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	toSerialize["time"] = o.Time

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotebookCreateDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cells    *[]NotebookCellCreateRequest `json:"cells"`
		Metadata *NotebookMetadata            `json:"metadata,omitempty"`
		Name     *string                      `json:"name"`
		Status   *NotebookStatus              `json:"status,omitempty"`
		Time     *NotebookGlobalTime          `json:"time"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Cells == nil {
		return fmt.Errorf("required field cells missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Time == nil {
		return fmt.Errorf("required field time missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cells", "metadata", "name", "status", "time"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Cells = *all.Cells
	if all.Metadata != nil && all.Metadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Metadata = all.Metadata
	o.Name = *all.Name
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Time = *all.Time

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
