// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"
)

// SpansAggregateBucketValueTimeseries A timeseries array.
type SpansAggregateBucketValueTimeseries struct {
	Items []SpansAggregateBucketValueTimeseriesPoint

	// UnparsedObject contains the raw value of the array if there was an error when deserializing into the struct
	UnparsedObject []interface{} `json:"-"`
}

// NewSpansAggregateBucketValueTimeseries instantiates a new SpansAggregateBucketValueTimeseries object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSpansAggregateBucketValueTimeseries() *SpansAggregateBucketValueTimeseries {
	this := SpansAggregateBucketValueTimeseries{}
	return &this
}

// NewSpansAggregateBucketValueTimeseriesWithDefaults instantiates a new SpansAggregateBucketValueTimeseries object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSpansAggregateBucketValueTimeseriesWithDefaults() *SpansAggregateBucketValueTimeseries {
	this := SpansAggregateBucketValueTimeseries{}
	return &this
}

// MarshalJSON serializes the struct using spec logic.
func (o SpansAggregateBucketValueTimeseries) MarshalJSON() ([]byte, error) {
	toSerialize := make([]interface{}, len(o.Items))
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SpansAggregateBucketValueTimeseries) UnmarshalJSON(bytes []byte) (err error) {
	if err = json.Unmarshal(bytes, &o.Items); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	if o.Items != nil && len(o.Items) > 0 {
		for _, v := range o.Items {
			if v.UnparsedObject != nil {
				return json.Unmarshal(bytes, &o.UnparsedObject)
			}
		}
	}

	return nil
}
