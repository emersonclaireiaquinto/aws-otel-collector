// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"
)

// CIAppGroupByTotal - A resulting object to put the given computes in over all the matching records.
type CIAppGroupByTotal struct {
	CIAppGroupByTotalBoolean *bool
	CIAppGroupByTotalString  *string
	CIAppGroupByTotalNumber  *float64

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CIAppGroupByTotalBooleanAsCIAppGroupByTotal is a convenience function that returns bool wrapped in CIAppGroupByTotal.
func CIAppGroupByTotalBooleanAsCIAppGroupByTotal(v *bool) CIAppGroupByTotal {
	return CIAppGroupByTotal{CIAppGroupByTotalBoolean: v}
}

// CIAppGroupByTotalStringAsCIAppGroupByTotal is a convenience function that returns string wrapped in CIAppGroupByTotal.
func CIAppGroupByTotalStringAsCIAppGroupByTotal(v *string) CIAppGroupByTotal {
	return CIAppGroupByTotal{CIAppGroupByTotalString: v}
}

// CIAppGroupByTotalNumberAsCIAppGroupByTotal is a convenience function that returns float64 wrapped in CIAppGroupByTotal.
func CIAppGroupByTotalNumberAsCIAppGroupByTotal(v *float64) CIAppGroupByTotal {
	return CIAppGroupByTotal{CIAppGroupByTotalNumber: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CIAppGroupByTotal) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CIAppGroupByTotalBoolean
	err = json.Unmarshal(data, &obj.CIAppGroupByTotalBoolean)
	if err == nil {
		if obj.CIAppGroupByTotalBoolean != nil {
			jsonCIAppGroupByTotalBoolean, _ := json.Marshal(obj.CIAppGroupByTotalBoolean)
			if string(jsonCIAppGroupByTotalBoolean) == "{}" { // empty struct
				obj.CIAppGroupByTotalBoolean = nil
			} else {
				match++
			}
		} else {
			obj.CIAppGroupByTotalBoolean = nil
		}
	} else {
		obj.CIAppGroupByTotalBoolean = nil
	}

	// try to unmarshal data into CIAppGroupByTotalString
	err = json.Unmarshal(data, &obj.CIAppGroupByTotalString)
	if err == nil {
		if obj.CIAppGroupByTotalString != nil {
			jsonCIAppGroupByTotalString, _ := json.Marshal(obj.CIAppGroupByTotalString)
			if string(jsonCIAppGroupByTotalString) == "{}" { // empty struct
				obj.CIAppGroupByTotalString = nil
			} else {
				match++
			}
		} else {
			obj.CIAppGroupByTotalString = nil
		}
	} else {
		obj.CIAppGroupByTotalString = nil
	}

	// try to unmarshal data into CIAppGroupByTotalNumber
	err = json.Unmarshal(data, &obj.CIAppGroupByTotalNumber)
	if err == nil {
		if obj.CIAppGroupByTotalNumber != nil {
			jsonCIAppGroupByTotalNumber, _ := json.Marshal(obj.CIAppGroupByTotalNumber)
			if string(jsonCIAppGroupByTotalNumber) == "{}" { // empty struct
				obj.CIAppGroupByTotalNumber = nil
			} else {
				match++
			}
		} else {
			obj.CIAppGroupByTotalNumber = nil
		}
	} else {
		obj.CIAppGroupByTotalNumber = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CIAppGroupByTotalBoolean = nil
		obj.CIAppGroupByTotalString = nil
		obj.CIAppGroupByTotalNumber = nil
		return json.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CIAppGroupByTotal) MarshalJSON() ([]byte, error) {
	if obj.CIAppGroupByTotalBoolean != nil {
		return json.Marshal(&obj.CIAppGroupByTotalBoolean)
	}

	if obj.CIAppGroupByTotalString != nil {
		return json.Marshal(&obj.CIAppGroupByTotalString)
	}

	if obj.CIAppGroupByTotalNumber != nil {
		return json.Marshal(&obj.CIAppGroupByTotalNumber)
	}

	if obj.UnparsedObject != nil {
		return json.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CIAppGroupByTotal) GetActualInstance() interface{} {
	if obj.CIAppGroupByTotalBoolean != nil {
		return obj.CIAppGroupByTotalBoolean
	}

	if obj.CIAppGroupByTotalString != nil {
		return obj.CIAppGroupByTotalString
	}

	if obj.CIAppGroupByTotalNumber != nil {
		return obj.CIAppGroupByTotalNumber
	}

	// all schemas are nil
	return nil
}
