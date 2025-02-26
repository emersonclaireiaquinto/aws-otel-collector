// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentAttachmentUpdateResponse The response object containing the created or updated incident attachments.
type IncidentAttachmentUpdateResponse struct {
	// An array of incident attachments. Only the attachments that were created or updated by the request are
	// returned.
	Data []IncidentAttachmentData `json:"data"`
	// Included related resources that the user requested.
	Included []IncidentAttachmentsResponseIncludedItem `json:"included,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewIncidentAttachmentUpdateResponse instantiates a new IncidentAttachmentUpdateResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentAttachmentUpdateResponse(data []IncidentAttachmentData) *IncidentAttachmentUpdateResponse {
	this := IncidentAttachmentUpdateResponse{}
	this.Data = data
	return &this
}

// NewIncidentAttachmentUpdateResponseWithDefaults instantiates a new IncidentAttachmentUpdateResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentAttachmentUpdateResponseWithDefaults() *IncidentAttachmentUpdateResponse {
	this := IncidentAttachmentUpdateResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *IncidentAttachmentUpdateResponse) GetData() []IncidentAttachmentData {
	if o == nil {
		var ret []IncidentAttachmentData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentAttachmentUpdateResponse) GetDataOk() (*[]IncidentAttachmentData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *IncidentAttachmentUpdateResponse) SetData(v []IncidentAttachmentData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *IncidentAttachmentUpdateResponse) GetIncluded() []IncidentAttachmentsResponseIncludedItem {
	if o == nil || o.Included == nil {
		var ret []IncidentAttachmentsResponseIncludedItem
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentAttachmentUpdateResponse) GetIncludedOk() (*[]IncidentAttachmentsResponseIncludedItem, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *IncidentAttachmentUpdateResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []IncidentAttachmentsResponseIncludedItem and assigns it to the Included field.
func (o *IncidentAttachmentUpdateResponse) SetIncluded(v []IncidentAttachmentsResponseIncludedItem) {
	o.Included = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentAttachmentUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentAttachmentUpdateResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *[]IncidentAttachmentData                 `json:"data"`
		Included []IncidentAttachmentsResponseIncludedItem `json:"included,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included"})
	} else {
		return err
	}
	o.Data = *all.Data
	o.Included = all.Included

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
