// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAPITestResultFull Object returned describing a API test result.
type SyntheticsAPITestResultFull struct {
	// Object describing the API test configuration.
	Check *SyntheticsAPITestResultFullCheck `json:"check,omitempty"`
	// When the API test was conducted.
	CheckTime *float64 `json:"check_time,omitempty"`
	// Version of the API test used.
	CheckVersion *int64 `json:"check_version,omitempty"`
	// Locations for which to query the API test results.
	ProbeDc *string `json:"probe_dc,omitempty"`
	// Object containing results for your Synthetic API test.
	Result *SyntheticsAPITestResultData `json:"result,omitempty"`
	// ID of the API test result.
	ResultId *string `json:"result_id,omitempty"`
	// The status of your Synthetic monitor.
	// * `O` for not triggered
	// * `1` for triggered
	// * `2` for no data
	Status *SyntheticsTestMonitorStatus `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSyntheticsAPITestResultFull instantiates a new SyntheticsAPITestResultFull object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsAPITestResultFull() *SyntheticsAPITestResultFull {
	this := SyntheticsAPITestResultFull{}
	return &this
}

// NewSyntheticsAPITestResultFullWithDefaults instantiates a new SyntheticsAPITestResultFull object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsAPITestResultFullWithDefaults() *SyntheticsAPITestResultFull {
	this := SyntheticsAPITestResultFull{}
	return &this
}

// GetCheck returns the Check field value if set, zero value otherwise.
func (o *SyntheticsAPITestResultFull) GetCheck() SyntheticsAPITestResultFullCheck {
	if o == nil || o.Check == nil {
		var ret SyntheticsAPITestResultFullCheck
		return ret
	}
	return *o.Check
}

// GetCheckOk returns a tuple with the Check field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestResultFull) GetCheckOk() (*SyntheticsAPITestResultFullCheck, bool) {
	if o == nil || o.Check == nil {
		return nil, false
	}
	return o.Check, true
}

// HasCheck returns a boolean if a field has been set.
func (o *SyntheticsAPITestResultFull) HasCheck() bool {
	return o != nil && o.Check != nil
}

// SetCheck gets a reference to the given SyntheticsAPITestResultFullCheck and assigns it to the Check field.
func (o *SyntheticsAPITestResultFull) SetCheck(v SyntheticsAPITestResultFullCheck) {
	o.Check = &v
}

// GetCheckTime returns the CheckTime field value if set, zero value otherwise.
func (o *SyntheticsAPITestResultFull) GetCheckTime() float64 {
	if o == nil || o.CheckTime == nil {
		var ret float64
		return ret
	}
	return *o.CheckTime
}

// GetCheckTimeOk returns a tuple with the CheckTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestResultFull) GetCheckTimeOk() (*float64, bool) {
	if o == nil || o.CheckTime == nil {
		return nil, false
	}
	return o.CheckTime, true
}

// HasCheckTime returns a boolean if a field has been set.
func (o *SyntheticsAPITestResultFull) HasCheckTime() bool {
	return o != nil && o.CheckTime != nil
}

// SetCheckTime gets a reference to the given float64 and assigns it to the CheckTime field.
func (o *SyntheticsAPITestResultFull) SetCheckTime(v float64) {
	o.CheckTime = &v
}

// GetCheckVersion returns the CheckVersion field value if set, zero value otherwise.
func (o *SyntheticsAPITestResultFull) GetCheckVersion() int64 {
	if o == nil || o.CheckVersion == nil {
		var ret int64
		return ret
	}
	return *o.CheckVersion
}

// GetCheckVersionOk returns a tuple with the CheckVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestResultFull) GetCheckVersionOk() (*int64, bool) {
	if o == nil || o.CheckVersion == nil {
		return nil, false
	}
	return o.CheckVersion, true
}

// HasCheckVersion returns a boolean if a field has been set.
func (o *SyntheticsAPITestResultFull) HasCheckVersion() bool {
	return o != nil && o.CheckVersion != nil
}

// SetCheckVersion gets a reference to the given int64 and assigns it to the CheckVersion field.
func (o *SyntheticsAPITestResultFull) SetCheckVersion(v int64) {
	o.CheckVersion = &v
}

// GetProbeDc returns the ProbeDc field value if set, zero value otherwise.
func (o *SyntheticsAPITestResultFull) GetProbeDc() string {
	if o == nil || o.ProbeDc == nil {
		var ret string
		return ret
	}
	return *o.ProbeDc
}

// GetProbeDcOk returns a tuple with the ProbeDc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestResultFull) GetProbeDcOk() (*string, bool) {
	if o == nil || o.ProbeDc == nil {
		return nil, false
	}
	return o.ProbeDc, true
}

// HasProbeDc returns a boolean if a field has been set.
func (o *SyntheticsAPITestResultFull) HasProbeDc() bool {
	return o != nil && o.ProbeDc != nil
}

// SetProbeDc gets a reference to the given string and assigns it to the ProbeDc field.
func (o *SyntheticsAPITestResultFull) SetProbeDc(v string) {
	o.ProbeDc = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *SyntheticsAPITestResultFull) GetResult() SyntheticsAPITestResultData {
	if o == nil || o.Result == nil {
		var ret SyntheticsAPITestResultData
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestResultFull) GetResultOk() (*SyntheticsAPITestResultData, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *SyntheticsAPITestResultFull) HasResult() bool {
	return o != nil && o.Result != nil
}

// SetResult gets a reference to the given SyntheticsAPITestResultData and assigns it to the Result field.
func (o *SyntheticsAPITestResultFull) SetResult(v SyntheticsAPITestResultData) {
	o.Result = &v
}

// GetResultId returns the ResultId field value if set, zero value otherwise.
func (o *SyntheticsAPITestResultFull) GetResultId() string {
	if o == nil || o.ResultId == nil {
		var ret string
		return ret
	}
	return *o.ResultId
}

// GetResultIdOk returns a tuple with the ResultId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestResultFull) GetResultIdOk() (*string, bool) {
	if o == nil || o.ResultId == nil {
		return nil, false
	}
	return o.ResultId, true
}

// HasResultId returns a boolean if a field has been set.
func (o *SyntheticsAPITestResultFull) HasResultId() bool {
	return o != nil && o.ResultId != nil
}

// SetResultId gets a reference to the given string and assigns it to the ResultId field.
func (o *SyntheticsAPITestResultFull) SetResultId(v string) {
	o.ResultId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyntheticsAPITestResultFull) GetStatus() SyntheticsTestMonitorStatus {
	if o == nil || o.Status == nil {
		var ret SyntheticsTestMonitorStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestResultFull) GetStatusOk() (*SyntheticsTestMonitorStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyntheticsAPITestResultFull) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given SyntheticsTestMonitorStatus and assigns it to the Status field.
func (o *SyntheticsAPITestResultFull) SetStatus(v SyntheticsTestMonitorStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsAPITestResultFull) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Check != nil {
		toSerialize["check"] = o.Check
	}
	if o.CheckTime != nil {
		toSerialize["check_time"] = o.CheckTime
	}
	if o.CheckVersion != nil {
		toSerialize["check_version"] = o.CheckVersion
	}
	if o.ProbeDc != nil {
		toSerialize["probe_dc"] = o.ProbeDc
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.ResultId != nil {
		toSerialize["result_id"] = o.ResultId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsAPITestResultFull) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Check        *SyntheticsAPITestResultFullCheck `json:"check,omitempty"`
		CheckTime    *float64                          `json:"check_time,omitempty"`
		CheckVersion *int64                            `json:"check_version,omitempty"`
		ProbeDc      *string                           `json:"probe_dc,omitempty"`
		Result       *SyntheticsAPITestResultData      `json:"result,omitempty"`
		ResultId     *string                           `json:"result_id,omitempty"`
		Status       *SyntheticsTestMonitorStatus      `json:"status,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"check", "check_time", "check_version", "probe_dc", "result", "result_id", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Check != nil && all.Check.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Check = all.Check
	o.CheckTime = all.CheckTime
	o.CheckVersion = all.CheckVersion
	o.ProbeDc = all.ProbeDc
	if all.Result != nil && all.Result.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Result = all.Result
	o.ResultId = all.ResultId
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
