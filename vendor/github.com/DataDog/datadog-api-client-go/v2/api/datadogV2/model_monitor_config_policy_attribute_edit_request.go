// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorConfigPolicyAttributeEditRequest Policy and policy type for a monitor configuration policy.
type MonitorConfigPolicyAttributeEditRequest struct {
	// Configuration for the policy.
	Policy MonitorConfigPolicyPolicy `json:"policy"`
	// The monitor configuration policy type.
	PolicyType MonitorConfigPolicyType `json:"policy_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewMonitorConfigPolicyAttributeEditRequest instantiates a new MonitorConfigPolicyAttributeEditRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorConfigPolicyAttributeEditRequest(policy MonitorConfigPolicyPolicy, policyType MonitorConfigPolicyType) *MonitorConfigPolicyAttributeEditRequest {
	this := MonitorConfigPolicyAttributeEditRequest{}
	this.Policy = policy
	this.PolicyType = policyType
	return &this
}

// NewMonitorConfigPolicyAttributeEditRequestWithDefaults instantiates a new MonitorConfigPolicyAttributeEditRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorConfigPolicyAttributeEditRequestWithDefaults() *MonitorConfigPolicyAttributeEditRequest {
	this := MonitorConfigPolicyAttributeEditRequest{}
	var policyType MonitorConfigPolicyType = MONITORCONFIGPOLICYTYPE_TAG
	this.PolicyType = policyType
	return &this
}

// GetPolicy returns the Policy field value.
func (o *MonitorConfigPolicyAttributeEditRequest) GetPolicy() MonitorConfigPolicyPolicy {
	if o == nil {
		var ret MonitorConfigPolicyPolicy
		return ret
	}
	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *MonitorConfigPolicyAttributeEditRequest) GetPolicyOk() (*MonitorConfigPolicyPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value.
func (o *MonitorConfigPolicyAttributeEditRequest) SetPolicy(v MonitorConfigPolicyPolicy) {
	o.Policy = v
}

// GetPolicyType returns the PolicyType field value.
func (o *MonitorConfigPolicyAttributeEditRequest) GetPolicyType() MonitorConfigPolicyType {
	if o == nil {
		var ret MonitorConfigPolicyType
		return ret
	}
	return o.PolicyType
}

// GetPolicyTypeOk returns a tuple with the PolicyType field value
// and a boolean to check if the value has been set.
func (o *MonitorConfigPolicyAttributeEditRequest) GetPolicyTypeOk() (*MonitorConfigPolicyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyType, true
}

// SetPolicyType sets field value.
func (o *MonitorConfigPolicyAttributeEditRequest) SetPolicyType(v MonitorConfigPolicyType) {
	o.PolicyType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorConfigPolicyAttributeEditRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["policy"] = o.Policy
	toSerialize["policy_type"] = o.PolicyType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorConfigPolicyAttributeEditRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Policy     *MonitorConfigPolicyPolicy `json:"policy"`
		PolicyType *MonitorConfigPolicyType   `json:"policy_type"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Policy == nil {
		return fmt.Errorf("required field policy missing")
	}
	if all.PolicyType == nil {
		return fmt.Errorf("required field policy_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"policy", "policy_type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Policy = *all.Policy
	if !all.PolicyType.IsValid() {
		hasInvalidField = true
	} else {
		o.PolicyType = *all.PolicyType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
