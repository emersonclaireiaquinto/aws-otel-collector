// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSAccountListResponse List of enabled AWS accounts.
type AWSAccountListResponse struct {
	// List of enabled AWS accounts.
	Accounts []AWSAccount `json:"accounts,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewAWSAccountListResponse instantiates a new AWSAccountListResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSAccountListResponse() *AWSAccountListResponse {
	this := AWSAccountListResponse{}
	return &this
}

// NewAWSAccountListResponseWithDefaults instantiates a new AWSAccountListResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSAccountListResponseWithDefaults() *AWSAccountListResponse {
	this := AWSAccountListResponse{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *AWSAccountListResponse) GetAccounts() []AWSAccount {
	if o == nil || o.Accounts == nil {
		var ret []AWSAccount
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountListResponse) GetAccountsOk() (*[]AWSAccount, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return &o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *AWSAccountListResponse) HasAccounts() bool {
	return o != nil && o.Accounts != nil
}

// SetAccounts gets a reference to the given []AWSAccount and assigns it to the Accounts field.
func (o *AWSAccountListResponse) SetAccounts(v []AWSAccount) {
	o.Accounts = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSAccountListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSAccountListResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Accounts []AWSAccount `json:"accounts,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"accounts"})
	} else {
		return err
	}
	o.Accounts = all.Accounts

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
