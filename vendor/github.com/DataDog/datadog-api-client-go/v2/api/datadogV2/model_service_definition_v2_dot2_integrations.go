// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceDefinitionV2Dot2Integrations Third party integrations that Datadog supports.
type ServiceDefinitionV2Dot2Integrations struct {
	// Opsgenie integration for the service.
	Opsgenie *ServiceDefinitionV2Dot2Opsgenie `json:"opsgenie,omitempty"`
	// PagerDuty integration for the service.
	Pagerduty *ServiceDefinitionV2Dot2Pagerduty `json:"pagerduty,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewServiceDefinitionV2Dot2Integrations instantiates a new ServiceDefinitionV2Dot2Integrations object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceDefinitionV2Dot2Integrations() *ServiceDefinitionV2Dot2Integrations {
	this := ServiceDefinitionV2Dot2Integrations{}
	return &this
}

// NewServiceDefinitionV2Dot2IntegrationsWithDefaults instantiates a new ServiceDefinitionV2Dot2Integrations object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceDefinitionV2Dot2IntegrationsWithDefaults() *ServiceDefinitionV2Dot2Integrations {
	this := ServiceDefinitionV2Dot2Integrations{}
	return &this
}

// GetOpsgenie returns the Opsgenie field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot2Integrations) GetOpsgenie() ServiceDefinitionV2Dot2Opsgenie {
	if o == nil || o.Opsgenie == nil {
		var ret ServiceDefinitionV2Dot2Opsgenie
		return ret
	}
	return *o.Opsgenie
}

// GetOpsgenieOk returns a tuple with the Opsgenie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot2Integrations) GetOpsgenieOk() (*ServiceDefinitionV2Dot2Opsgenie, bool) {
	if o == nil || o.Opsgenie == nil {
		return nil, false
	}
	return o.Opsgenie, true
}

// HasOpsgenie returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot2Integrations) HasOpsgenie() bool {
	return o != nil && o.Opsgenie != nil
}

// SetOpsgenie gets a reference to the given ServiceDefinitionV2Dot2Opsgenie and assigns it to the Opsgenie field.
func (o *ServiceDefinitionV2Dot2Integrations) SetOpsgenie(v ServiceDefinitionV2Dot2Opsgenie) {
	o.Opsgenie = &v
}

// GetPagerduty returns the Pagerduty field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot2Integrations) GetPagerduty() ServiceDefinitionV2Dot2Pagerduty {
	if o == nil || o.Pagerduty == nil {
		var ret ServiceDefinitionV2Dot2Pagerduty
		return ret
	}
	return *o.Pagerduty
}

// GetPagerdutyOk returns a tuple with the Pagerduty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot2Integrations) GetPagerdutyOk() (*ServiceDefinitionV2Dot2Pagerduty, bool) {
	if o == nil || o.Pagerduty == nil {
		return nil, false
	}
	return o.Pagerduty, true
}

// HasPagerduty returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot2Integrations) HasPagerduty() bool {
	return o != nil && o.Pagerduty != nil
}

// SetPagerduty gets a reference to the given ServiceDefinitionV2Dot2Pagerduty and assigns it to the Pagerduty field.
func (o *ServiceDefinitionV2Dot2Integrations) SetPagerduty(v ServiceDefinitionV2Dot2Pagerduty) {
	o.Pagerduty = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceDefinitionV2Dot2Integrations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Opsgenie != nil {
		toSerialize["opsgenie"] = o.Opsgenie
	}
	if o.Pagerduty != nil {
		toSerialize["pagerduty"] = o.Pagerduty
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceDefinitionV2Dot2Integrations) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Opsgenie  *ServiceDefinitionV2Dot2Opsgenie  `json:"opsgenie,omitempty"`
		Pagerduty *ServiceDefinitionV2Dot2Pagerduty `json:"pagerduty,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"opsgenie", "pagerduty"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Opsgenie != nil && all.Opsgenie.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Opsgenie = all.Opsgenie
	if all.Pagerduty != nil && all.Pagerduty.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Pagerduty = all.Pagerduty

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
