// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsGroupByHistogram Used to perform a histogram computation (only for measure facets).
// Note: at most 100 buckets are allowed, the number of buckets is (max - min)/interval.
type LogsGroupByHistogram struct {
	// The bin size of the histogram buckets
	Interval float64 `json:"interval"`
	// The maximum value for the measure used in the histogram
	// (values greater than this one are filtered out)
	Max float64 `json:"max"`
	// The minimum value for the measure used in the histogram
	// (values smaller than this one are filtered out)
	Min float64 `json:"min"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewLogsGroupByHistogram instantiates a new LogsGroupByHistogram object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsGroupByHistogram(interval float64, max float64, min float64) *LogsGroupByHistogram {
	this := LogsGroupByHistogram{}
	this.Interval = interval
	this.Max = max
	this.Min = min
	return &this
}

// NewLogsGroupByHistogramWithDefaults instantiates a new LogsGroupByHistogram object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsGroupByHistogramWithDefaults() *LogsGroupByHistogram {
	this := LogsGroupByHistogram{}
	return &this
}

// GetInterval returns the Interval field value.
func (o *LogsGroupByHistogram) GetInterval() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *LogsGroupByHistogram) GetIntervalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value.
func (o *LogsGroupByHistogram) SetInterval(v float64) {
	o.Interval = v
}

// GetMax returns the Max field value.
func (o *LogsGroupByHistogram) GetMax() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Max
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
func (o *LogsGroupByHistogram) GetMaxOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Max, true
}

// SetMax sets field value.
func (o *LogsGroupByHistogram) SetMax(v float64) {
	o.Max = v
}

// GetMin returns the Min field value.
func (o *LogsGroupByHistogram) GetMin() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Min
}

// GetMinOk returns a tuple with the Min field value
// and a boolean to check if the value has been set.
func (o *LogsGroupByHistogram) GetMinOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Min, true
}

// SetMin sets field value.
func (o *LogsGroupByHistogram) SetMin(v float64) {
	o.Min = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsGroupByHistogram) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["interval"] = o.Interval
	toSerialize["max"] = o.Max
	toSerialize["min"] = o.Min

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsGroupByHistogram) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Interval *float64 `json:"interval"`
		Max      *float64 `json:"max"`
		Min      *float64 `json:"min"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Interval == nil {
		return fmt.Errorf("required field interval missing")
	}
	if all.Max == nil {
		return fmt.Errorf("required field max missing")
	}
	if all.Min == nil {
		return fmt.Errorf("required field min missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"interval", "max", "min"})
	} else {
		return err
	}
	o.Interval = *all.Interval
	o.Max = *all.Max
	o.Min = *all.Min

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
