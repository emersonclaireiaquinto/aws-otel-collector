// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserInvitationsRequest Object to invite users to join the organization.
type UserInvitationsRequest struct {
	// List of user invitations.
	Data []UserInvitationData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewUserInvitationsRequest instantiates a new UserInvitationsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserInvitationsRequest(data []UserInvitationData) *UserInvitationsRequest {
	this := UserInvitationsRequest{}
	this.Data = data
	return &this
}

// NewUserInvitationsRequestWithDefaults instantiates a new UserInvitationsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserInvitationsRequestWithDefaults() *UserInvitationsRequest {
	this := UserInvitationsRequest{}
	return &this
}

// GetData returns the Data field value.
func (o *UserInvitationsRequest) GetData() []UserInvitationData {
	if o == nil {
		var ret []UserInvitationData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UserInvitationsRequest) GetDataOk() (*[]UserInvitationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *UserInvitationsRequest) SetData(v []UserInvitationData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserInvitationsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserInvitationsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]UserInvitationData `json:"data"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}
	o.Data = *all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
