// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostMapWidgetDefinitionRequests List of definitions.
type HostMapWidgetDefinitionRequests struct {
	// Updated host map.
	Fill *HostMapRequest `json:"fill,omitempty"`
	// Updated host map.
	Size *HostMapRequest `json:"size,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewHostMapWidgetDefinitionRequests instantiates a new HostMapWidgetDefinitionRequests object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHostMapWidgetDefinitionRequests() *HostMapWidgetDefinitionRequests {
	this := HostMapWidgetDefinitionRequests{}
	return &this
}

// NewHostMapWidgetDefinitionRequestsWithDefaults instantiates a new HostMapWidgetDefinitionRequests object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHostMapWidgetDefinitionRequestsWithDefaults() *HostMapWidgetDefinitionRequests {
	this := HostMapWidgetDefinitionRequests{}
	return &this
}

// GetFill returns the Fill field value if set, zero value otherwise.
func (o *HostMapWidgetDefinitionRequests) GetFill() HostMapRequest {
	if o == nil || o.Fill == nil {
		var ret HostMapRequest
		return ret
	}
	return *o.Fill
}

// GetFillOk returns a tuple with the Fill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinitionRequests) GetFillOk() (*HostMapRequest, bool) {
	if o == nil || o.Fill == nil {
		return nil, false
	}
	return o.Fill, true
}

// HasFill returns a boolean if a field has been set.
func (o *HostMapWidgetDefinitionRequests) HasFill() bool {
	return o != nil && o.Fill != nil
}

// SetFill gets a reference to the given HostMapRequest and assigns it to the Fill field.
func (o *HostMapWidgetDefinitionRequests) SetFill(v HostMapRequest) {
	o.Fill = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *HostMapWidgetDefinitionRequests) GetSize() HostMapRequest {
	if o == nil || o.Size == nil {
		var ret HostMapRequest
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinitionRequests) GetSizeOk() (*HostMapRequest, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *HostMapWidgetDefinitionRequests) HasSize() bool {
	return o != nil && o.Size != nil
}

// SetSize gets a reference to the given HostMapRequest and assigns it to the Size field.
func (o *HostMapWidgetDefinitionRequests) SetSize(v HostMapRequest) {
	o.Size = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HostMapWidgetDefinitionRequests) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Fill != nil {
		toSerialize["fill"] = o.Fill
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HostMapWidgetDefinitionRequests) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Fill *HostMapRequest `json:"fill,omitempty"`
		Size *HostMapRequest `json:"size,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fill", "size"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Fill != nil && all.Fill.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Fill = all.Fill
	if all.Size != nil && all.Size.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Size = all.Size

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
