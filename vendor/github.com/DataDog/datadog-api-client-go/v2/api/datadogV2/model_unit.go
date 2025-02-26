// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Unit Object containing the metric unit family, scale factor, name, and short name.
type Unit struct {
	// Unit family, allows for conversion between units of the same family, for scaling.
	Family *string `json:"family,omitempty"`
	// Unit name
	Name *string `json:"name,omitempty"`
	// Plural form of the unit name.
	Plural *string `json:"plural,omitempty"`
	// Factor for scaling between units of the same family.
	ScaleFactor *float64 `json:"scale_factor,omitempty"`
	// Abbreviation of the unit.
	ShortName *string `json:"short_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewUnit instantiates a new Unit object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUnit() *Unit {
	this := Unit{}
	return &this
}

// NewUnitWithDefaults instantiates a new Unit object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUnitWithDefaults() *Unit {
	this := Unit{}
	return &this
}

// GetFamily returns the Family field value if set, zero value otherwise.
func (o *Unit) GetFamily() string {
	if o == nil || o.Family == nil {
		var ret string
		return ret
	}
	return *o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Unit) GetFamilyOk() (*string, bool) {
	if o == nil || o.Family == nil {
		return nil, false
	}
	return o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *Unit) HasFamily() bool {
	return o != nil && o.Family != nil
}

// SetFamily gets a reference to the given string and assigns it to the Family field.
func (o *Unit) SetFamily(v string) {
	o.Family = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Unit) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Unit) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Unit) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Unit) SetName(v string) {
	o.Name = &v
}

// GetPlural returns the Plural field value if set, zero value otherwise.
func (o *Unit) GetPlural() string {
	if o == nil || o.Plural == nil {
		var ret string
		return ret
	}
	return *o.Plural
}

// GetPluralOk returns a tuple with the Plural field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Unit) GetPluralOk() (*string, bool) {
	if o == nil || o.Plural == nil {
		return nil, false
	}
	return o.Plural, true
}

// HasPlural returns a boolean if a field has been set.
func (o *Unit) HasPlural() bool {
	return o != nil && o.Plural != nil
}

// SetPlural gets a reference to the given string and assigns it to the Plural field.
func (o *Unit) SetPlural(v string) {
	o.Plural = &v
}

// GetScaleFactor returns the ScaleFactor field value if set, zero value otherwise.
func (o *Unit) GetScaleFactor() float64 {
	if o == nil || o.ScaleFactor == nil {
		var ret float64
		return ret
	}
	return *o.ScaleFactor
}

// GetScaleFactorOk returns a tuple with the ScaleFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Unit) GetScaleFactorOk() (*float64, bool) {
	if o == nil || o.ScaleFactor == nil {
		return nil, false
	}
	return o.ScaleFactor, true
}

// HasScaleFactor returns a boolean if a field has been set.
func (o *Unit) HasScaleFactor() bool {
	return o != nil && o.ScaleFactor != nil
}

// SetScaleFactor gets a reference to the given float64 and assigns it to the ScaleFactor field.
func (o *Unit) SetScaleFactor(v float64) {
	o.ScaleFactor = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *Unit) GetShortName() string {
	if o == nil || o.ShortName == nil {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Unit) GetShortNameOk() (*string, bool) {
	if o == nil || o.ShortName == nil {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *Unit) HasShortName() bool {
	return o != nil && o.ShortName != nil
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *Unit) SetShortName(v string) {
	o.ShortName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Unit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Family != nil {
		toSerialize["family"] = o.Family
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Plural != nil {
		toSerialize["plural"] = o.Plural
	}
	if o.ScaleFactor != nil {
		toSerialize["scale_factor"] = o.ScaleFactor
	}
	if o.ShortName != nil {
		toSerialize["short_name"] = o.ShortName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Unit) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Family      *string  `json:"family,omitempty"`
		Name        *string  `json:"name,omitempty"`
		Plural      *string  `json:"plural,omitempty"`
		ScaleFactor *float64 `json:"scale_factor,omitempty"`
		ShortName   *string  `json:"short_name,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"family", "name", "plural", "scale_factor", "short_name"})
	} else {
		return err
	}
	o.Family = all.Family
	o.Name = all.Name
	o.Plural = all.Plural
	o.ScaleFactor = all.ScaleFactor
	o.ShortName = all.ShortName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableUnit handles when a null is used for Unit.
type NullableUnit struct {
	value *Unit
	isSet bool
}

// Get returns the associated value.
func (v NullableUnit) Get() *Unit {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableUnit) Set(val *Unit) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableUnit) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableUnit) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableUnit initializes the struct as if Set has been called.
func NewNullableUnit(val *Unit) *NullableUnit {
	return &NullableUnit{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
