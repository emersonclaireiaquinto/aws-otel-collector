// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainerImagesResponse List of Container Images.
type ContainerImagesResponse struct {
	// Array of Container Image objects.
	Data []ContainerImageItem `json:"data,omitempty"`
	// Pagination links.
	Links *ContainerImagesResponseLinks `json:"links,omitempty"`
	// Response metadata object.
	Meta *ContainerImageMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewContainerImagesResponse instantiates a new ContainerImagesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewContainerImagesResponse() *ContainerImagesResponse {
	this := ContainerImagesResponse{}
	return &this
}

// NewContainerImagesResponseWithDefaults instantiates a new ContainerImagesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewContainerImagesResponseWithDefaults() *ContainerImagesResponse {
	this := ContainerImagesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ContainerImagesResponse) GetData() []ContainerImageItem {
	if o == nil || o.Data == nil {
		var ret []ContainerImageItem
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImagesResponse) GetDataOk() (*[]ContainerImageItem, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ContainerImagesResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []ContainerImageItem and assigns it to the Data field.
func (o *ContainerImagesResponse) SetData(v []ContainerImageItem) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ContainerImagesResponse) GetLinks() ContainerImagesResponseLinks {
	if o == nil || o.Links == nil {
		var ret ContainerImagesResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImagesResponse) GetLinksOk() (*ContainerImagesResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ContainerImagesResponse) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given ContainerImagesResponseLinks and assigns it to the Links field.
func (o *ContainerImagesResponse) SetLinks(v ContainerImagesResponseLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ContainerImagesResponse) GetMeta() ContainerImageMeta {
	if o == nil || o.Meta == nil {
		var ret ContainerImageMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImagesResponse) GetMetaOk() (*ContainerImageMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ContainerImagesResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given ContainerImageMeta and assigns it to the Meta field.
func (o *ContainerImagesResponse) SetMeta(v ContainerImageMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ContainerImagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ContainerImagesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data  []ContainerImageItem          `json:"data,omitempty"`
		Links *ContainerImagesResponseLinks `json:"links,omitempty"`
		Meta  *ContainerImageMeta           `json:"meta,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "links", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = all.Data
	if all.Links != nil && all.Links.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Links = all.Links
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
