// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WebhooksIntegrationUpdateRequest Update request of a Webhooks integration object.
//
// *All properties are optional.*
type WebhooksIntegrationUpdateRequest struct {
	// If `null`, uses no header.
	// If given a JSON payload, these will be headers attached to your webhook.
	CustomHeaders *string `json:"custom_headers,omitempty"`
	// Encoding type. Can be given either `json` or `form`.
	EncodeAs *WebhooksIntegrationEncoding `json:"encode_as,omitempty"`
	// The name of the webhook. It corresponds with `<WEBHOOK_NAME>`.
	// Learn more on how to use it in
	// [monitor notifications](https://docs.datadoghq.com/monitors/notify).
	Name *string `json:"name,omitempty"`
	// If `null`, uses the default payload.
	// If given a JSON payload, the webhook returns the payload
	// specified by the given payload.
	// [Webhooks variable usage](https://docs.datadoghq.com/integrations/webhooks/#usage).
	Payload datadog.NullableString `json:"payload,omitempty"`
	// URL of the webhook.
	Url *string `json:"url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewWebhooksIntegrationUpdateRequest instantiates a new WebhooksIntegrationUpdateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWebhooksIntegrationUpdateRequest() *WebhooksIntegrationUpdateRequest {
	this := WebhooksIntegrationUpdateRequest{}
	var encodeAs WebhooksIntegrationEncoding = WEBHOOKSINTEGRATIONENCODING_JSON
	this.EncodeAs = &encodeAs
	return &this
}

// NewWebhooksIntegrationUpdateRequestWithDefaults instantiates a new WebhooksIntegrationUpdateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWebhooksIntegrationUpdateRequestWithDefaults() *WebhooksIntegrationUpdateRequest {
	this := WebhooksIntegrationUpdateRequest{}
	var encodeAs WebhooksIntegrationEncoding = WEBHOOKSINTEGRATIONENCODING_JSON
	this.EncodeAs = &encodeAs
	return &this
}

// GetCustomHeaders returns the CustomHeaders field value if set, zero value otherwise.
func (o *WebhooksIntegrationUpdateRequest) GetCustomHeaders() string {
	if o == nil || o.CustomHeaders == nil {
		var ret string
		return ret
	}
	return *o.CustomHeaders
}

// GetCustomHeadersOk returns a tuple with the CustomHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksIntegrationUpdateRequest) GetCustomHeadersOk() (*string, bool) {
	if o == nil || o.CustomHeaders == nil {
		return nil, false
	}
	return o.CustomHeaders, true
}

// HasCustomHeaders returns a boolean if a field has been set.
func (o *WebhooksIntegrationUpdateRequest) HasCustomHeaders() bool {
	return o != nil && o.CustomHeaders != nil
}

// SetCustomHeaders gets a reference to the given string and assigns it to the CustomHeaders field.
func (o *WebhooksIntegrationUpdateRequest) SetCustomHeaders(v string) {
	o.CustomHeaders = &v
}

// GetEncodeAs returns the EncodeAs field value if set, zero value otherwise.
func (o *WebhooksIntegrationUpdateRequest) GetEncodeAs() WebhooksIntegrationEncoding {
	if o == nil || o.EncodeAs == nil {
		var ret WebhooksIntegrationEncoding
		return ret
	}
	return *o.EncodeAs
}

// GetEncodeAsOk returns a tuple with the EncodeAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksIntegrationUpdateRequest) GetEncodeAsOk() (*WebhooksIntegrationEncoding, bool) {
	if o == nil || o.EncodeAs == nil {
		return nil, false
	}
	return o.EncodeAs, true
}

// HasEncodeAs returns a boolean if a field has been set.
func (o *WebhooksIntegrationUpdateRequest) HasEncodeAs() bool {
	return o != nil && o.EncodeAs != nil
}

// SetEncodeAs gets a reference to the given WebhooksIntegrationEncoding and assigns it to the EncodeAs field.
func (o *WebhooksIntegrationUpdateRequest) SetEncodeAs(v WebhooksIntegrationEncoding) {
	o.EncodeAs = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WebhooksIntegrationUpdateRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksIntegrationUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WebhooksIntegrationUpdateRequest) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WebhooksIntegrationUpdateRequest) SetName(v string) {
	o.Name = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksIntegrationUpdateRequest) GetPayload() string {
	if o == nil || o.Payload.Get() == nil {
		var ret string
		return ret
	}
	return *o.Payload.Get()
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *WebhooksIntegrationUpdateRequest) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Payload.Get(), o.Payload.IsSet()
}

// HasPayload returns a boolean if a field has been set.
func (o *WebhooksIntegrationUpdateRequest) HasPayload() bool {
	return o != nil && o.Payload.IsSet()
}

// SetPayload gets a reference to the given datadog.NullableString and assigns it to the Payload field.
func (o *WebhooksIntegrationUpdateRequest) SetPayload(v string) {
	o.Payload.Set(&v)
}

// SetPayloadNil sets the value for Payload to be an explicit nil.
func (o *WebhooksIntegrationUpdateRequest) SetPayloadNil() {
	o.Payload.Set(nil)
}

// UnsetPayload ensures that no value is present for Payload, not even an explicit nil.
func (o *WebhooksIntegrationUpdateRequest) UnsetPayload() {
	o.Payload.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WebhooksIntegrationUpdateRequest) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksIntegrationUpdateRequest) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WebhooksIntegrationUpdateRequest) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WebhooksIntegrationUpdateRequest) SetUrl(v string) {
	o.Url = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o WebhooksIntegrationUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.CustomHeaders != nil {
		toSerialize["custom_headers"] = o.CustomHeaders
	}
	if o.EncodeAs != nil {
		toSerialize["encode_as"] = o.EncodeAs
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Payload.IsSet() {
		toSerialize["payload"] = o.Payload.Get()
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WebhooksIntegrationUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomHeaders *string                      `json:"custom_headers,omitempty"`
		EncodeAs      *WebhooksIntegrationEncoding `json:"encode_as,omitempty"`
		Name          *string                      `json:"name,omitempty"`
		Payload       datadog.NullableString       `json:"payload,omitempty"`
		Url           *string                      `json:"url,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"custom_headers", "encode_as", "name", "payload", "url"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CustomHeaders = all.CustomHeaders
	if all.EncodeAs != nil && !all.EncodeAs.IsValid() {
		hasInvalidField = true
	} else {
		o.EncodeAs = all.EncodeAs
	}
	o.Name = all.Name
	o.Payload = all.Payload
	o.Url = all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
