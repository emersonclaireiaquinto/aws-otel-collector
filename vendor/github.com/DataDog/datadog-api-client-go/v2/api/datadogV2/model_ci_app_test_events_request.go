// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppTestEventsRequest The request for a tests search.
type CIAppTestEventsRequest struct {
	// The search and filter query settings.
	Filter *CIAppTestsQueryFilter `json:"filter,omitempty"`
	// Global query options that are used during the query.
	// Only supply timezone or time offset, not both. Otherwise, the query fails.
	Options *CIAppQueryOptions `json:"options,omitempty"`
	// Paging attributes for listing events.
	Page *CIAppQueryPageOptions `json:"page,omitempty"`
	// Sort parameters when querying events.
	Sort *CIAppSort `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewCIAppTestEventsRequest instantiates a new CIAppTestEventsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCIAppTestEventsRequest() *CIAppTestEventsRequest {
	this := CIAppTestEventsRequest{}
	return &this
}

// NewCIAppTestEventsRequestWithDefaults instantiates a new CIAppTestEventsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCIAppTestEventsRequestWithDefaults() *CIAppTestEventsRequest {
	this := CIAppTestEventsRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *CIAppTestEventsRequest) GetFilter() CIAppTestsQueryFilter {
	if o == nil || o.Filter == nil {
		var ret CIAppTestsQueryFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppTestEventsRequest) GetFilterOk() (*CIAppTestsQueryFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *CIAppTestEventsRequest) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given CIAppTestsQueryFilter and assigns it to the Filter field.
func (o *CIAppTestEventsRequest) SetFilter(v CIAppTestsQueryFilter) {
	o.Filter = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *CIAppTestEventsRequest) GetOptions() CIAppQueryOptions {
	if o == nil || o.Options == nil {
		var ret CIAppQueryOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppTestEventsRequest) GetOptionsOk() (*CIAppQueryOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *CIAppTestEventsRequest) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given CIAppQueryOptions and assigns it to the Options field.
func (o *CIAppTestEventsRequest) SetOptions(v CIAppQueryOptions) {
	o.Options = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *CIAppTestEventsRequest) GetPage() CIAppQueryPageOptions {
	if o == nil || o.Page == nil {
		var ret CIAppQueryPageOptions
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppTestEventsRequest) GetPageOk() (*CIAppQueryPageOptions, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *CIAppTestEventsRequest) HasPage() bool {
	return o != nil && o.Page != nil
}

// SetPage gets a reference to the given CIAppQueryPageOptions and assigns it to the Page field.
func (o *CIAppTestEventsRequest) SetPage(v CIAppQueryPageOptions) {
	o.Page = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *CIAppTestEventsRequest) GetSort() CIAppSort {
	if o == nil || o.Sort == nil {
		var ret CIAppSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppTestEventsRequest) GetSortOk() (*CIAppSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *CIAppTestEventsRequest) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given CIAppSort and assigns it to the Sort field.
func (o *CIAppTestEventsRequest) SetSort(v CIAppSort) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CIAppTestEventsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CIAppTestEventsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter  *CIAppTestsQueryFilter `json:"filter,omitempty"`
		Options *CIAppQueryOptions     `json:"options,omitempty"`
		Page    *CIAppQueryPageOptions `json:"page,omitempty"`
		Sort    *CIAppSort             `json:"sort,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filter", "options", "page", "sort"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Filter != nil && all.Filter.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Filter = all.Filter
	if all.Options != nil && all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Options = all.Options
	if all.Page != nil && all.Page.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Page = all.Page
	if all.Sort != nil && !all.Sort.IsValid() {
		hasInvalidField = true
	} else {
		o.Sort = all.Sort
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
