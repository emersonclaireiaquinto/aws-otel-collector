// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IPRanges IP ranges.
type IPRanges struct {
	// Available prefix information for the Agent endpoints.
	Agents *IPPrefixesAgents `json:"agents,omitempty"`
	// Available prefix information for the API endpoints.
	Api *IPPrefixesAPI `json:"api,omitempty"`
	// Available prefix information for the APM endpoints.
	Apm *IPPrefixesAPM `json:"apm,omitempty"`
	// Available prefix information for all Datadog endpoints.
	Global *IPPrefixesGlobal `json:"global,omitempty"`
	// Available prefix information for the Logs endpoints.
	Logs *IPPrefixesLogs `json:"logs,omitempty"`
	// Date when last updated, in the form `YYYY-MM-DD-hh-mm-ss`.
	Modified *string `json:"modified,omitempty"`
	// Available prefix information for the Orchestrator endpoints.
	Orchestrator *IPPrefixesOrchestrator `json:"orchestrator,omitempty"`
	// Available prefix information for the Process endpoints.
	Process *IPPrefixesProcess `json:"process,omitempty"`
	// Available prefix information for the Remote Configuration endpoints.
	RemoteConfiguration *IPPrefixesRemoteConfiguration `json:"remote-configuration,omitempty"`
	// Available prefix information for the Synthetics endpoints.
	Synthetics *IPPrefixesSynthetics `json:"synthetics,omitempty"`
	// Available prefix information for the Synthetics Private Locations endpoints.
	SyntheticsPrivateLocations *IPPrefixesSyntheticsPrivateLocations `json:"synthetics-private-locations,omitempty"`
	// Version of the IP list.
	Version *int64 `json:"version,omitempty"`
	// Available prefix information for the Webhook endpoints.
	Webhooks *IPPrefixesWebhooks `json:"webhooks,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewIPRanges instantiates a new IPRanges object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIPRanges() *IPRanges {
	this := IPRanges{}
	return &this
}

// NewIPRangesWithDefaults instantiates a new IPRanges object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIPRangesWithDefaults() *IPRanges {
	this := IPRanges{}
	return &this
}

// GetAgents returns the Agents field value if set, zero value otherwise.
func (o *IPRanges) GetAgents() IPPrefixesAgents {
	if o == nil || o.Agents == nil {
		var ret IPPrefixesAgents
		return ret
	}
	return *o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRanges) GetAgentsOk() (*IPPrefixesAgents, bool) {
	if o == nil || o.Agents == nil {
		return nil, false
	}
	return o.Agents, true
}

// HasAgents returns a boolean if a field has been set.
func (o *IPRanges) HasAgents() bool {
	return o != nil && o.Agents != nil
}

// SetAgents gets a reference to the given IPPrefixesAgents and assigns it to the Agents field.
func (o *IPRanges) SetAgents(v IPPrefixesAgents) {
	o.Agents = &v
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *IPRanges) GetApi() IPPrefixesAPI {
	if o == nil || o.Api == nil {
		var ret IPPrefixesAPI
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRanges) GetApiOk() (*IPPrefixesAPI, bool) {
	if o == nil || o.Api == nil {
		return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *IPRanges) HasApi() bool {
	return o != nil && o.Api != nil
}

// SetApi gets a reference to the given IPPrefixesAPI and assigns it to the Api field.
func (o *IPRanges) SetApi(v IPPrefixesAPI) {
	o.Api = &v
}

// GetApm returns the Apm field value if set, zero value otherwise.
func (o *IPRanges) GetApm() IPPrefixesAPM {
	if o == nil || o.Apm == nil {
		var ret IPPrefixesAPM
		return ret
	}
	return *o.Apm
}

// GetApmOk returns a tuple with the Apm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRanges) GetApmOk() (*IPPrefixesAPM, bool) {
	if o == nil || o.Apm == nil {
		return nil, false
	}
	return o.Apm, true
}

// HasApm returns a boolean if a field has been set.
func (o *IPRanges) HasApm() bool {
	return o != nil && o.Apm != nil
}

// SetApm gets a reference to the given IPPrefixesAPM and assigns it to the Apm field.
func (o *IPRanges) SetApm(v IPPrefixesAPM) {
	o.Apm = &v
}

// GetGlobal returns the Global field value if set, zero value otherwise.
func (o *IPRanges) GetGlobal() IPPrefixesGlobal {
	if o == nil || o.Global == nil {
		var ret IPPrefixesGlobal
		return ret
	}
	return *o.Global
}

// GetGlobalOk returns a tuple with the Global field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRanges) GetGlobalOk() (*IPPrefixesGlobal, bool) {
	if o == nil || o.Global == nil {
		return nil, false
	}
	return o.Global, true
}

// HasGlobal returns a boolean if a field has been set.
func (o *IPRanges) HasGlobal() bool {
	return o != nil && o.Global != nil
}

// SetGlobal gets a reference to the given IPPrefixesGlobal and assigns it to the Global field.
func (o *IPRanges) SetGlobal(v IPPrefixesGlobal) {
	o.Global = &v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *IPRanges) GetLogs() IPPrefixesLogs {
	if o == nil || o.Logs == nil {
		var ret IPPrefixesLogs
		return ret
	}
	return *o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRanges) GetLogsOk() (*IPPrefixesLogs, bool) {
	if o == nil || o.Logs == nil {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *IPRanges) HasLogs() bool {
	return o != nil && o.Logs != nil
}

// SetLogs gets a reference to the given IPPrefixesLogs and assigns it to the Logs field.
func (o *IPRanges) SetLogs(v IPPrefixesLogs) {
	o.Logs = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *IPRanges) GetModified() string {
	if o == nil || o.Modified == nil {
		var ret string
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRanges) GetModifiedOk() (*string, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *IPRanges) HasModified() bool {
	return o != nil && o.Modified != nil
}

// SetModified gets a reference to the given string and assigns it to the Modified field.
func (o *IPRanges) SetModified(v string) {
	o.Modified = &v
}

// GetOrchestrator returns the Orchestrator field value if set, zero value otherwise.
func (o *IPRanges) GetOrchestrator() IPPrefixesOrchestrator {
	if o == nil || o.Orchestrator == nil {
		var ret IPPrefixesOrchestrator
		return ret
	}
	return *o.Orchestrator
}

// GetOrchestratorOk returns a tuple with the Orchestrator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRanges) GetOrchestratorOk() (*IPPrefixesOrchestrator, bool) {
	if o == nil || o.Orchestrator == nil {
		return nil, false
	}
	return o.Orchestrator, true
}

// HasOrchestrator returns a boolean if a field has been set.
func (o *IPRanges) HasOrchestrator() bool {
	return o != nil && o.Orchestrator != nil
}

// SetOrchestrator gets a reference to the given IPPrefixesOrchestrator and assigns it to the Orchestrator field.
func (o *IPRanges) SetOrchestrator(v IPPrefixesOrchestrator) {
	o.Orchestrator = &v
}

// GetProcess returns the Process field value if set, zero value otherwise.
func (o *IPRanges) GetProcess() IPPrefixesProcess {
	if o == nil || o.Process == nil {
		var ret IPPrefixesProcess
		return ret
	}
	return *o.Process
}

// GetProcessOk returns a tuple with the Process field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRanges) GetProcessOk() (*IPPrefixesProcess, bool) {
	if o == nil || o.Process == nil {
		return nil, false
	}
	return o.Process, true
}

// HasProcess returns a boolean if a field has been set.
func (o *IPRanges) HasProcess() bool {
	return o != nil && o.Process != nil
}

// SetProcess gets a reference to the given IPPrefixesProcess and assigns it to the Process field.
func (o *IPRanges) SetProcess(v IPPrefixesProcess) {
	o.Process = &v
}

// GetRemoteConfiguration returns the RemoteConfiguration field value if set, zero value otherwise.
func (o *IPRanges) GetRemoteConfiguration() IPPrefixesRemoteConfiguration {
	if o == nil || o.RemoteConfiguration == nil {
		var ret IPPrefixesRemoteConfiguration
		return ret
	}
	return *o.RemoteConfiguration
}

// GetRemoteConfigurationOk returns a tuple with the RemoteConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRanges) GetRemoteConfigurationOk() (*IPPrefixesRemoteConfiguration, bool) {
	if o == nil || o.RemoteConfiguration == nil {
		return nil, false
	}
	return o.RemoteConfiguration, true
}

// HasRemoteConfiguration returns a boolean if a field has been set.
func (o *IPRanges) HasRemoteConfiguration() bool {
	return o != nil && o.RemoteConfiguration != nil
}

// SetRemoteConfiguration gets a reference to the given IPPrefixesRemoteConfiguration and assigns it to the RemoteConfiguration field.
func (o *IPRanges) SetRemoteConfiguration(v IPPrefixesRemoteConfiguration) {
	o.RemoteConfiguration = &v
}

// GetSynthetics returns the Synthetics field value if set, zero value otherwise.
func (o *IPRanges) GetSynthetics() IPPrefixesSynthetics {
	if o == nil || o.Synthetics == nil {
		var ret IPPrefixesSynthetics
		return ret
	}
	return *o.Synthetics
}

// GetSyntheticsOk returns a tuple with the Synthetics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRanges) GetSyntheticsOk() (*IPPrefixesSynthetics, bool) {
	if o == nil || o.Synthetics == nil {
		return nil, false
	}
	return o.Synthetics, true
}

// HasSynthetics returns a boolean if a field has been set.
func (o *IPRanges) HasSynthetics() bool {
	return o != nil && o.Synthetics != nil
}

// SetSynthetics gets a reference to the given IPPrefixesSynthetics and assigns it to the Synthetics field.
func (o *IPRanges) SetSynthetics(v IPPrefixesSynthetics) {
	o.Synthetics = &v
}

// GetSyntheticsPrivateLocations returns the SyntheticsPrivateLocations field value if set, zero value otherwise.
func (o *IPRanges) GetSyntheticsPrivateLocations() IPPrefixesSyntheticsPrivateLocations {
	if o == nil || o.SyntheticsPrivateLocations == nil {
		var ret IPPrefixesSyntheticsPrivateLocations
		return ret
	}
	return *o.SyntheticsPrivateLocations
}

// GetSyntheticsPrivateLocationsOk returns a tuple with the SyntheticsPrivateLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRanges) GetSyntheticsPrivateLocationsOk() (*IPPrefixesSyntheticsPrivateLocations, bool) {
	if o == nil || o.SyntheticsPrivateLocations == nil {
		return nil, false
	}
	return o.SyntheticsPrivateLocations, true
}

// HasSyntheticsPrivateLocations returns a boolean if a field has been set.
func (o *IPRanges) HasSyntheticsPrivateLocations() bool {
	return o != nil && o.SyntheticsPrivateLocations != nil
}

// SetSyntheticsPrivateLocations gets a reference to the given IPPrefixesSyntheticsPrivateLocations and assigns it to the SyntheticsPrivateLocations field.
func (o *IPRanges) SetSyntheticsPrivateLocations(v IPPrefixesSyntheticsPrivateLocations) {
	o.SyntheticsPrivateLocations = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *IPRanges) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRanges) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *IPRanges) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *IPRanges) SetVersion(v int64) {
	o.Version = &v
}

// GetWebhooks returns the Webhooks field value if set, zero value otherwise.
func (o *IPRanges) GetWebhooks() IPPrefixesWebhooks {
	if o == nil || o.Webhooks == nil {
		var ret IPPrefixesWebhooks
		return ret
	}
	return *o.Webhooks
}

// GetWebhooksOk returns a tuple with the Webhooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRanges) GetWebhooksOk() (*IPPrefixesWebhooks, bool) {
	if o == nil || o.Webhooks == nil {
		return nil, false
	}
	return o.Webhooks, true
}

// HasWebhooks returns a boolean if a field has been set.
func (o *IPRanges) HasWebhooks() bool {
	return o != nil && o.Webhooks != nil
}

// SetWebhooks gets a reference to the given IPPrefixesWebhooks and assigns it to the Webhooks field.
func (o *IPRanges) SetWebhooks(v IPPrefixesWebhooks) {
	o.Webhooks = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IPRanges) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Agents != nil {
		toSerialize["agents"] = o.Agents
	}
	if o.Api != nil {
		toSerialize["api"] = o.Api
	}
	if o.Apm != nil {
		toSerialize["apm"] = o.Apm
	}
	if o.Global != nil {
		toSerialize["global"] = o.Global
	}
	if o.Logs != nil {
		toSerialize["logs"] = o.Logs
	}
	if o.Modified != nil {
		toSerialize["modified"] = o.Modified
	}
	if o.Orchestrator != nil {
		toSerialize["orchestrator"] = o.Orchestrator
	}
	if o.Process != nil {
		toSerialize["process"] = o.Process
	}
	if o.RemoteConfiguration != nil {
		toSerialize["remote-configuration"] = o.RemoteConfiguration
	}
	if o.Synthetics != nil {
		toSerialize["synthetics"] = o.Synthetics
	}
	if o.SyntheticsPrivateLocations != nil {
		toSerialize["synthetics-private-locations"] = o.SyntheticsPrivateLocations
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Webhooks != nil {
		toSerialize["webhooks"] = o.Webhooks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IPRanges) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Agents                     *IPPrefixesAgents                     `json:"agents,omitempty"`
		Api                        *IPPrefixesAPI                        `json:"api,omitempty"`
		Apm                        *IPPrefixesAPM                        `json:"apm,omitempty"`
		Global                     *IPPrefixesGlobal                     `json:"global,omitempty"`
		Logs                       *IPPrefixesLogs                       `json:"logs,omitempty"`
		Modified                   *string                               `json:"modified,omitempty"`
		Orchestrator               *IPPrefixesOrchestrator               `json:"orchestrator,omitempty"`
		Process                    *IPPrefixesProcess                    `json:"process,omitempty"`
		RemoteConfiguration        *IPPrefixesRemoteConfiguration        `json:"remote-configuration,omitempty"`
		Synthetics                 *IPPrefixesSynthetics                 `json:"synthetics,omitempty"`
		SyntheticsPrivateLocations *IPPrefixesSyntheticsPrivateLocations `json:"synthetics-private-locations,omitempty"`
		Version                    *int64                                `json:"version,omitempty"`
		Webhooks                   *IPPrefixesWebhooks                   `json:"webhooks,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"agents", "api", "apm", "global", "logs", "modified", "orchestrator", "process", "remote-configuration", "synthetics", "synthetics-private-locations", "version", "webhooks"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Agents != nil && all.Agents.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Agents = all.Agents
	if all.Api != nil && all.Api.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Api = all.Api
	if all.Apm != nil && all.Apm.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Apm = all.Apm
	if all.Global != nil && all.Global.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Global = all.Global
	if all.Logs != nil && all.Logs.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Logs = all.Logs
	o.Modified = all.Modified
	if all.Orchestrator != nil && all.Orchestrator.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Orchestrator = all.Orchestrator
	if all.Process != nil && all.Process.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Process = all.Process
	if all.RemoteConfiguration != nil && all.RemoteConfiguration.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RemoteConfiguration = all.RemoteConfiguration
	if all.Synthetics != nil && all.Synthetics.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Synthetics = all.Synthetics
	if all.SyntheticsPrivateLocations != nil && all.SyntheticsPrivateLocations.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SyntheticsPrivateLocations = all.SyntheticsPrivateLocations
	o.Version = all.Version
	if all.Webhooks != nil && all.Webhooks.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Webhooks = all.Webhooks

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
