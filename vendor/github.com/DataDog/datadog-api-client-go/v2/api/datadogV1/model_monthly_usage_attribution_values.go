// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonthlyUsageAttributionValues Fields in Usage Summary by tag(s).
type MonthlyUsageAttributionValues struct {
	// The percentage of synthetic API test usage by tag(s).
	ApiPercentage *float64 `json:"api_percentage,omitempty"`
	// The synthetic API test usage by tag(s).
	ApiUsage *float64 `json:"api_usage,omitempty"`
	// The percentage of APM ECS Fargate task usage by tag(s).
	ApmFargatePercentage *float64 `json:"apm_fargate_percentage,omitempty"`
	// The APM ECS Fargate task usage by tag(s).
	ApmFargateUsage *float64 `json:"apm_fargate_usage,omitempty"`
	// The percentage of APM host usage by tag(s).
	ApmHostPercentage *float64 `json:"apm_host_percentage,omitempty"`
	// The APM host usage by tag(s).
	ApmHostUsage *float64 `json:"apm_host_usage,omitempty"`
	// The percentage of APM and Universal Service Monitoring host usage by tag(s).
	ApmUsmPercentage *float64 `json:"apm_usm_percentage,omitempty"`
	// The APM and Universal Service Monitoring host usage by tag(s).
	ApmUsmUsage *float64 `json:"apm_usm_usage,omitempty"`
	// The percentage of Application Security Monitoring ECS Fargate task usage by tag(s).
	AppsecFargatePercentage *float64 `json:"appsec_fargate_percentage,omitempty"`
	// The Application Security Monitoring ECS Fargate task usage by tag(s).
	AppsecFargateUsage *float64 `json:"appsec_fargate_usage,omitempty"`
	// The percentage of Application Security Monitoring host usage by tag(s).
	AppsecPercentage *float64 `json:"appsec_percentage,omitempty"`
	// The Application Security Monitoring host usage by tag(s).
	AppsecUsage *float64 `json:"appsec_usage,omitempty"`
	// The percentage of synthetic browser test usage by tag(s).
	BrowserPercentage *float64 `json:"browser_percentage,omitempty"`
	// The synthetic browser test usage by tag(s).
	BrowserUsage *float64 `json:"browser_usage,omitempty"`
	// The percentage of Git committers for Intelligent Test Runner usage by tag(s).
	CiVisibilityItrPercentage *float64 `json:"ci_visibility_itr_percentage,omitempty"`
	// The Git committers for Intelligent Test Runner usage by tag(s).
	CiVisibilityItrUsage *float64 `json:"ci_visibility_itr_usage,omitempty"`
	// The percentage of container usage without the Datadog Agent by tag(s).
	ContainerExclAgentPercentage *float64 `json:"container_excl_agent_percentage,omitempty"`
	// The container usage without the Datadog Agent by tag(s).
	ContainerExclAgentUsage *float64 `json:"container_excl_agent_usage,omitempty"`
	// The percentage of container usage by tag(s).
	ContainerPercentage *float64 `json:"container_percentage,omitempty"`
	// The container usage by tag(s).
	ContainerUsage *float64 `json:"container_usage,omitempty"`
	// The percentage of Cloud Security Management Pro container usage by tag(s).
	CspmContainersPercentage *float64 `json:"cspm_containers_percentage,omitempty"`
	// The Cloud Security Management Pro container usage by tag(s).
	CspmContainersUsage *float64 `json:"cspm_containers_usage,omitempty"`
	// The percentage of Cloud Security Management Pro host usage by tag(s).
	CspmHostsPercentage *float64 `json:"cspm_hosts_percentage,omitempty"`
	// The Cloud Security Management Pro host usage by tag(s).
	CspmHostsUsage *float64 `json:"cspm_hosts_usage,omitempty"`
	// The percentage of ingested custom metrics usage by tag(s).
	CustomIngestedTimeseriesPercentage *float64 `json:"custom_ingested_timeseries_percentage,omitempty"`
	// The ingested custom metrics usage by tag(s).
	CustomIngestedTimeseriesUsage *float64 `json:"custom_ingested_timeseries_usage,omitempty"`
	// The percentage of indexed custom metrics usage by tag(s).
	CustomTimeseriesPercentage *float64 `json:"custom_timeseries_percentage,omitempty"`
	// The indexed custom metrics usage by tag(s).
	CustomTimeseriesUsage *float64 `json:"custom_timeseries_usage,omitempty"`
	// The percentage of Cloud Workload Security container usage by tag(s).
	CwsContainersPercentage *float64 `json:"cws_containers_percentage,omitempty"`
	// The Cloud Workload Security container usage by tag(s).
	CwsContainersUsage *float64 `json:"cws_containers_usage,omitempty"`
	// The percentage of Cloud Workload Security host usage by tag(s).
	CwsHostsPercentage *float64 `json:"cws_hosts_percentage,omitempty"`
	// The Cloud Workload Security host usage by tag(s).
	CwsHostsUsage *float64 `json:"cws_hosts_usage,omitempty"`
	// The percentage of Database Monitoring host usage by tag(s).
	DbmHostsPercentage *float64 `json:"dbm_hosts_percentage,omitempty"`
	// The Database Monitoring host usage by tag(s).
	DbmHostsUsage *float64 `json:"dbm_hosts_usage,omitempty"`
	// The percentage of Database Monitoring queries usage by tag(s).
	DbmQueriesPercentage *float64 `json:"dbm_queries_percentage,omitempty"`
	// The Database Monitoring queries usage by tag(s).
	DbmQueriesUsage *float64 `json:"dbm_queries_usage,omitempty"`
	// The percentage of estimated live indexed logs usage by tag(s).
	EstimatedIndexedLogsPercentage *float64 `json:"estimated_indexed_logs_percentage,omitempty"`
	// The estimated live indexed logs usage by tag(s).
	EstimatedIndexedLogsUsage *float64 `json:"estimated_indexed_logs_usage,omitempty"`
	// The percentage of estimated indexed spans usage by tag(s).
	EstimatedIndexedSpansPercentage *float64 `json:"estimated_indexed_spans_percentage,omitempty"`
	// The estimated indexed spans usage by tag(s).
	EstimatedIndexedSpansUsage *float64 `json:"estimated_indexed_spans_usage,omitempty"`
	// The percentage of estimated live ingested logs usage by tag(s).
	EstimatedIngestedLogsPercentage *float64 `json:"estimated_ingested_logs_percentage,omitempty"`
	// The estimated live ingested logs usage by tag(s).
	EstimatedIngestedLogsUsage *float64 `json:"estimated_ingested_logs_usage,omitempty"`
	// The percentage of estimated ingested spans usage by tag(s).
	EstimatedIngestedSpansPercentage *float64 `json:"estimated_ingested_spans_percentage,omitempty"`
	// The estimated ingested spans usage by tag(s).
	EstimatedIngestedSpansUsage *float64 `json:"estimated_ingested_spans_usage,omitempty"`
	// The percentage of estimated rum sessions usage by tag(s).
	EstimatedRumSessionsPercentage *float64 `json:"estimated_rum_sessions_percentage,omitempty"`
	// The estimated rum sessions usage by tag(s).
	EstimatedRumSessionsUsage *float64 `json:"estimated_rum_sessions_usage,omitempty"`
	// The percentage of Fargate usage by tags.
	FargatePercentage *float64 `json:"fargate_percentage,omitempty"`
	// The Fargate usage by tags.
	FargateUsage *float64 `json:"fargate_usage,omitempty"`
	// The percentage of Lambda function usage by tag(s).
	FunctionsPercentage *float64 `json:"functions_percentage,omitempty"`
	// The Lambda function usage by tag(s).
	FunctionsUsage *float64 `json:"functions_usage,omitempty"`
	// The percentage of infrastructure host usage by tag(s).
	InfraHostPercentage *float64 `json:"infra_host_percentage,omitempty"`
	// The infrastructure host usage by tag(s).
	InfraHostUsage *float64 `json:"infra_host_usage,omitempty"`
	// The percentage of Lambda invocation usage by tag(s).
	InvocationsPercentage *float64 `json:"invocations_percentage,omitempty"`
	// The Lambda invocation usage by tag(s).
	InvocationsUsage *float64 `json:"invocations_usage,omitempty"`
	// The percentage of Synthetic mobile application test usage by tag(s).
	MobileAppTestingPercentage *float64 `json:"mobile_app_testing_percentage,omitempty"`
	// The Synthetic mobile application test usage by tag(s).
	MobileAppTestingUsage *float64 `json:"mobile_app_testing_usage,omitempty"`
	// The percentage of Network Device Monitoring NetFlow usage by tag(s).
	NdmNetflowPercentage *float64 `json:"ndm_netflow_percentage,omitempty"`
	// The Network Device Monitoring NetFlow usage by tag(s).
	NdmNetflowUsage *float64 `json:"ndm_netflow_usage,omitempty"`
	// The percentage of network host usage by tag(s).
	NpmHostPercentage *float64 `json:"npm_host_percentage,omitempty"`
	// The network host usage by tag(s).
	NpmHostUsage *float64 `json:"npm_host_usage,omitempty"`
	// The percentage of observability pipeline bytes usage by tag(s).
	ObsPipelineBytesPercentage *float64 `json:"obs_pipeline_bytes_percentage,omitempty"`
	// The observability pipeline bytes usage by tag(s).
	ObsPipelineBytesUsage *float64 `json:"obs_pipeline_bytes_usage,omitempty"`
	// The percentage of profiled container usage by tag(s).
	ProfiledContainerPercentage *float64 `json:"profiled_container_percentage,omitempty"`
	// The profiled container usage by tag(s).
	ProfiledContainerUsage *float64 `json:"profiled_container_usage,omitempty"`
	// The percentage of profiled Fargate task usage by tag(s).
	ProfiledFargatePercentage *float64 `json:"profiled_fargate_percentage,omitempty"`
	// The profiled Fargate task usage by tag(s).
	ProfiledFargateUsage *float64 `json:"profiled_fargate_usage,omitempty"`
	// The percentage of profiled hosts usage by tag(s).
	ProfiledHostPercentage *float64 `json:"profiled_host_percentage,omitempty"`
	// The profiled hosts usage by tag(s).
	ProfiledHostUsage *float64 `json:"profiled_host_usage,omitempty"`
	// The percentage of Sensitive Data Scanner usage by tag(s).
	SdsScannedBytesPercentage *float64 `json:"sds_scanned_bytes_percentage,omitempty"`
	// The total Sensitive Data Scanner usage by tag(s).
	SdsScannedBytesUsage *float64 `json:"sds_scanned_bytes_usage,omitempty"`
	// The percentage of Serverless Apps usage by tag(s).
	ServerlessAppsPercentage *float64 `json:"serverless_apps_percentage,omitempty"`
	// The total Serverless Apps usage by tag(s).
	ServerlessAppsUsage *float64 `json:"serverless_apps_usage,omitempty"`
	// The percentage of network device usage by tag(s).
	SnmpPercentage *float64 `json:"snmp_percentage,omitempty"`
	// The network device usage by tag(s).
	SnmpUsage *float64 `json:"snmp_usage,omitempty"`
	// The percentage of universal service monitoring usage by tag(s).
	UniversalServiceMonitoringPercentage *float64 `json:"universal_service_monitoring_percentage,omitempty"`
	// The universal service monitoring usage by tag(s).
	UniversalServiceMonitoringUsage *float64 `json:"universal_service_monitoring_usage,omitempty"`
	// The percentage of Application Vulnerability Management usage by tag(s).
	VulnManagementHostsPercentage *float64 `json:"vuln_management_hosts_percentage,omitempty"`
	// The Application Vulnerability Management usage by tag(s).
	VulnManagementHostsUsage *float64 `json:"vuln_management_hosts_usage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewMonthlyUsageAttributionValues instantiates a new MonthlyUsageAttributionValues object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonthlyUsageAttributionValues() *MonthlyUsageAttributionValues {
	this := MonthlyUsageAttributionValues{}
	return &this
}

// NewMonthlyUsageAttributionValuesWithDefaults instantiates a new MonthlyUsageAttributionValues object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonthlyUsageAttributionValuesWithDefaults() *MonthlyUsageAttributionValues {
	this := MonthlyUsageAttributionValues{}
	return &this
}

// GetApiPercentage returns the ApiPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetApiPercentage() float64 {
	if o == nil || o.ApiPercentage == nil {
		var ret float64
		return ret
	}
	return *o.ApiPercentage
}

// GetApiPercentageOk returns a tuple with the ApiPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetApiPercentageOk() (*float64, bool) {
	if o == nil || o.ApiPercentage == nil {
		return nil, false
	}
	return o.ApiPercentage, true
}

// HasApiPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasApiPercentage() bool {
	return o != nil && o.ApiPercentage != nil
}

// SetApiPercentage gets a reference to the given float64 and assigns it to the ApiPercentage field.
func (o *MonthlyUsageAttributionValues) SetApiPercentage(v float64) {
	o.ApiPercentage = &v
}

// GetApiUsage returns the ApiUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetApiUsage() float64 {
	if o == nil || o.ApiUsage == nil {
		var ret float64
		return ret
	}
	return *o.ApiUsage
}

// GetApiUsageOk returns a tuple with the ApiUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetApiUsageOk() (*float64, bool) {
	if o == nil || o.ApiUsage == nil {
		return nil, false
	}
	return o.ApiUsage, true
}

// HasApiUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasApiUsage() bool {
	return o != nil && o.ApiUsage != nil
}

// SetApiUsage gets a reference to the given float64 and assigns it to the ApiUsage field.
func (o *MonthlyUsageAttributionValues) SetApiUsage(v float64) {
	o.ApiUsage = &v
}

// GetApmFargatePercentage returns the ApmFargatePercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetApmFargatePercentage() float64 {
	if o == nil || o.ApmFargatePercentage == nil {
		var ret float64
		return ret
	}
	return *o.ApmFargatePercentage
}

// GetApmFargatePercentageOk returns a tuple with the ApmFargatePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetApmFargatePercentageOk() (*float64, bool) {
	if o == nil || o.ApmFargatePercentage == nil {
		return nil, false
	}
	return o.ApmFargatePercentage, true
}

// HasApmFargatePercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasApmFargatePercentage() bool {
	return o != nil && o.ApmFargatePercentage != nil
}

// SetApmFargatePercentage gets a reference to the given float64 and assigns it to the ApmFargatePercentage field.
func (o *MonthlyUsageAttributionValues) SetApmFargatePercentage(v float64) {
	o.ApmFargatePercentage = &v
}

// GetApmFargateUsage returns the ApmFargateUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetApmFargateUsage() float64 {
	if o == nil || o.ApmFargateUsage == nil {
		var ret float64
		return ret
	}
	return *o.ApmFargateUsage
}

// GetApmFargateUsageOk returns a tuple with the ApmFargateUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetApmFargateUsageOk() (*float64, bool) {
	if o == nil || o.ApmFargateUsage == nil {
		return nil, false
	}
	return o.ApmFargateUsage, true
}

// HasApmFargateUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasApmFargateUsage() bool {
	return o != nil && o.ApmFargateUsage != nil
}

// SetApmFargateUsage gets a reference to the given float64 and assigns it to the ApmFargateUsage field.
func (o *MonthlyUsageAttributionValues) SetApmFargateUsage(v float64) {
	o.ApmFargateUsage = &v
}

// GetApmHostPercentage returns the ApmHostPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetApmHostPercentage() float64 {
	if o == nil || o.ApmHostPercentage == nil {
		var ret float64
		return ret
	}
	return *o.ApmHostPercentage
}

// GetApmHostPercentageOk returns a tuple with the ApmHostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetApmHostPercentageOk() (*float64, bool) {
	if o == nil || o.ApmHostPercentage == nil {
		return nil, false
	}
	return o.ApmHostPercentage, true
}

// HasApmHostPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasApmHostPercentage() bool {
	return o != nil && o.ApmHostPercentage != nil
}

// SetApmHostPercentage gets a reference to the given float64 and assigns it to the ApmHostPercentage field.
func (o *MonthlyUsageAttributionValues) SetApmHostPercentage(v float64) {
	o.ApmHostPercentage = &v
}

// GetApmHostUsage returns the ApmHostUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetApmHostUsage() float64 {
	if o == nil || o.ApmHostUsage == nil {
		var ret float64
		return ret
	}
	return *o.ApmHostUsage
}

// GetApmHostUsageOk returns a tuple with the ApmHostUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetApmHostUsageOk() (*float64, bool) {
	if o == nil || o.ApmHostUsage == nil {
		return nil, false
	}
	return o.ApmHostUsage, true
}

// HasApmHostUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasApmHostUsage() bool {
	return o != nil && o.ApmHostUsage != nil
}

// SetApmHostUsage gets a reference to the given float64 and assigns it to the ApmHostUsage field.
func (o *MonthlyUsageAttributionValues) SetApmHostUsage(v float64) {
	o.ApmHostUsage = &v
}

// GetApmUsmPercentage returns the ApmUsmPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetApmUsmPercentage() float64 {
	if o == nil || o.ApmUsmPercentage == nil {
		var ret float64
		return ret
	}
	return *o.ApmUsmPercentage
}

// GetApmUsmPercentageOk returns a tuple with the ApmUsmPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetApmUsmPercentageOk() (*float64, bool) {
	if o == nil || o.ApmUsmPercentage == nil {
		return nil, false
	}
	return o.ApmUsmPercentage, true
}

// HasApmUsmPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasApmUsmPercentage() bool {
	return o != nil && o.ApmUsmPercentage != nil
}

// SetApmUsmPercentage gets a reference to the given float64 and assigns it to the ApmUsmPercentage field.
func (o *MonthlyUsageAttributionValues) SetApmUsmPercentage(v float64) {
	o.ApmUsmPercentage = &v
}

// GetApmUsmUsage returns the ApmUsmUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetApmUsmUsage() float64 {
	if o == nil || o.ApmUsmUsage == nil {
		var ret float64
		return ret
	}
	return *o.ApmUsmUsage
}

// GetApmUsmUsageOk returns a tuple with the ApmUsmUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetApmUsmUsageOk() (*float64, bool) {
	if o == nil || o.ApmUsmUsage == nil {
		return nil, false
	}
	return o.ApmUsmUsage, true
}

// HasApmUsmUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasApmUsmUsage() bool {
	return o != nil && o.ApmUsmUsage != nil
}

// SetApmUsmUsage gets a reference to the given float64 and assigns it to the ApmUsmUsage field.
func (o *MonthlyUsageAttributionValues) SetApmUsmUsage(v float64) {
	o.ApmUsmUsage = &v
}

// GetAppsecFargatePercentage returns the AppsecFargatePercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetAppsecFargatePercentage() float64 {
	if o == nil || o.AppsecFargatePercentage == nil {
		var ret float64
		return ret
	}
	return *o.AppsecFargatePercentage
}

// GetAppsecFargatePercentageOk returns a tuple with the AppsecFargatePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetAppsecFargatePercentageOk() (*float64, bool) {
	if o == nil || o.AppsecFargatePercentage == nil {
		return nil, false
	}
	return o.AppsecFargatePercentage, true
}

// HasAppsecFargatePercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasAppsecFargatePercentage() bool {
	return o != nil && o.AppsecFargatePercentage != nil
}

// SetAppsecFargatePercentage gets a reference to the given float64 and assigns it to the AppsecFargatePercentage field.
func (o *MonthlyUsageAttributionValues) SetAppsecFargatePercentage(v float64) {
	o.AppsecFargatePercentage = &v
}

// GetAppsecFargateUsage returns the AppsecFargateUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetAppsecFargateUsage() float64 {
	if o == nil || o.AppsecFargateUsage == nil {
		var ret float64
		return ret
	}
	return *o.AppsecFargateUsage
}

// GetAppsecFargateUsageOk returns a tuple with the AppsecFargateUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetAppsecFargateUsageOk() (*float64, bool) {
	if o == nil || o.AppsecFargateUsage == nil {
		return nil, false
	}
	return o.AppsecFargateUsage, true
}

// HasAppsecFargateUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasAppsecFargateUsage() bool {
	return o != nil && o.AppsecFargateUsage != nil
}

// SetAppsecFargateUsage gets a reference to the given float64 and assigns it to the AppsecFargateUsage field.
func (o *MonthlyUsageAttributionValues) SetAppsecFargateUsage(v float64) {
	o.AppsecFargateUsage = &v
}

// GetAppsecPercentage returns the AppsecPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetAppsecPercentage() float64 {
	if o == nil || o.AppsecPercentage == nil {
		var ret float64
		return ret
	}
	return *o.AppsecPercentage
}

// GetAppsecPercentageOk returns a tuple with the AppsecPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetAppsecPercentageOk() (*float64, bool) {
	if o == nil || o.AppsecPercentage == nil {
		return nil, false
	}
	return o.AppsecPercentage, true
}

// HasAppsecPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasAppsecPercentage() bool {
	return o != nil && o.AppsecPercentage != nil
}

// SetAppsecPercentage gets a reference to the given float64 and assigns it to the AppsecPercentage field.
func (o *MonthlyUsageAttributionValues) SetAppsecPercentage(v float64) {
	o.AppsecPercentage = &v
}

// GetAppsecUsage returns the AppsecUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetAppsecUsage() float64 {
	if o == nil || o.AppsecUsage == nil {
		var ret float64
		return ret
	}
	return *o.AppsecUsage
}

// GetAppsecUsageOk returns a tuple with the AppsecUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetAppsecUsageOk() (*float64, bool) {
	if o == nil || o.AppsecUsage == nil {
		return nil, false
	}
	return o.AppsecUsage, true
}

// HasAppsecUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasAppsecUsage() bool {
	return o != nil && o.AppsecUsage != nil
}

// SetAppsecUsage gets a reference to the given float64 and assigns it to the AppsecUsage field.
func (o *MonthlyUsageAttributionValues) SetAppsecUsage(v float64) {
	o.AppsecUsage = &v
}

// GetBrowserPercentage returns the BrowserPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetBrowserPercentage() float64 {
	if o == nil || o.BrowserPercentage == nil {
		var ret float64
		return ret
	}
	return *o.BrowserPercentage
}

// GetBrowserPercentageOk returns a tuple with the BrowserPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetBrowserPercentageOk() (*float64, bool) {
	if o == nil || o.BrowserPercentage == nil {
		return nil, false
	}
	return o.BrowserPercentage, true
}

// HasBrowserPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasBrowserPercentage() bool {
	return o != nil && o.BrowserPercentage != nil
}

// SetBrowserPercentage gets a reference to the given float64 and assigns it to the BrowserPercentage field.
func (o *MonthlyUsageAttributionValues) SetBrowserPercentage(v float64) {
	o.BrowserPercentage = &v
}

// GetBrowserUsage returns the BrowserUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetBrowserUsage() float64 {
	if o == nil || o.BrowserUsage == nil {
		var ret float64
		return ret
	}
	return *o.BrowserUsage
}

// GetBrowserUsageOk returns a tuple with the BrowserUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetBrowserUsageOk() (*float64, bool) {
	if o == nil || o.BrowserUsage == nil {
		return nil, false
	}
	return o.BrowserUsage, true
}

// HasBrowserUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasBrowserUsage() bool {
	return o != nil && o.BrowserUsage != nil
}

// SetBrowserUsage gets a reference to the given float64 and assigns it to the BrowserUsage field.
func (o *MonthlyUsageAttributionValues) SetBrowserUsage(v float64) {
	o.BrowserUsage = &v
}

// GetCiVisibilityItrPercentage returns the CiVisibilityItrPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCiVisibilityItrPercentage() float64 {
	if o == nil || o.CiVisibilityItrPercentage == nil {
		var ret float64
		return ret
	}
	return *o.CiVisibilityItrPercentage
}

// GetCiVisibilityItrPercentageOk returns a tuple with the CiVisibilityItrPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCiVisibilityItrPercentageOk() (*float64, bool) {
	if o == nil || o.CiVisibilityItrPercentage == nil {
		return nil, false
	}
	return o.CiVisibilityItrPercentage, true
}

// HasCiVisibilityItrPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCiVisibilityItrPercentage() bool {
	return o != nil && o.CiVisibilityItrPercentage != nil
}

// SetCiVisibilityItrPercentage gets a reference to the given float64 and assigns it to the CiVisibilityItrPercentage field.
func (o *MonthlyUsageAttributionValues) SetCiVisibilityItrPercentage(v float64) {
	o.CiVisibilityItrPercentage = &v
}

// GetCiVisibilityItrUsage returns the CiVisibilityItrUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCiVisibilityItrUsage() float64 {
	if o == nil || o.CiVisibilityItrUsage == nil {
		var ret float64
		return ret
	}
	return *o.CiVisibilityItrUsage
}

// GetCiVisibilityItrUsageOk returns a tuple with the CiVisibilityItrUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCiVisibilityItrUsageOk() (*float64, bool) {
	if o == nil || o.CiVisibilityItrUsage == nil {
		return nil, false
	}
	return o.CiVisibilityItrUsage, true
}

// HasCiVisibilityItrUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCiVisibilityItrUsage() bool {
	return o != nil && o.CiVisibilityItrUsage != nil
}

// SetCiVisibilityItrUsage gets a reference to the given float64 and assigns it to the CiVisibilityItrUsage field.
func (o *MonthlyUsageAttributionValues) SetCiVisibilityItrUsage(v float64) {
	o.CiVisibilityItrUsage = &v
}

// GetContainerExclAgentPercentage returns the ContainerExclAgentPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetContainerExclAgentPercentage() float64 {
	if o == nil || o.ContainerExclAgentPercentage == nil {
		var ret float64
		return ret
	}
	return *o.ContainerExclAgentPercentage
}

// GetContainerExclAgentPercentageOk returns a tuple with the ContainerExclAgentPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetContainerExclAgentPercentageOk() (*float64, bool) {
	if o == nil || o.ContainerExclAgentPercentage == nil {
		return nil, false
	}
	return o.ContainerExclAgentPercentage, true
}

// HasContainerExclAgentPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasContainerExclAgentPercentage() bool {
	return o != nil && o.ContainerExclAgentPercentage != nil
}

// SetContainerExclAgentPercentage gets a reference to the given float64 and assigns it to the ContainerExclAgentPercentage field.
func (o *MonthlyUsageAttributionValues) SetContainerExclAgentPercentage(v float64) {
	o.ContainerExclAgentPercentage = &v
}

// GetContainerExclAgentUsage returns the ContainerExclAgentUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetContainerExclAgentUsage() float64 {
	if o == nil || o.ContainerExclAgentUsage == nil {
		var ret float64
		return ret
	}
	return *o.ContainerExclAgentUsage
}

// GetContainerExclAgentUsageOk returns a tuple with the ContainerExclAgentUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetContainerExclAgentUsageOk() (*float64, bool) {
	if o == nil || o.ContainerExclAgentUsage == nil {
		return nil, false
	}
	return o.ContainerExclAgentUsage, true
}

// HasContainerExclAgentUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasContainerExclAgentUsage() bool {
	return o != nil && o.ContainerExclAgentUsage != nil
}

// SetContainerExclAgentUsage gets a reference to the given float64 and assigns it to the ContainerExclAgentUsage field.
func (o *MonthlyUsageAttributionValues) SetContainerExclAgentUsage(v float64) {
	o.ContainerExclAgentUsage = &v
}

// GetContainerPercentage returns the ContainerPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetContainerPercentage() float64 {
	if o == nil || o.ContainerPercentage == nil {
		var ret float64
		return ret
	}
	return *o.ContainerPercentage
}

// GetContainerPercentageOk returns a tuple with the ContainerPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetContainerPercentageOk() (*float64, bool) {
	if o == nil || o.ContainerPercentage == nil {
		return nil, false
	}
	return o.ContainerPercentage, true
}

// HasContainerPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasContainerPercentage() bool {
	return o != nil && o.ContainerPercentage != nil
}

// SetContainerPercentage gets a reference to the given float64 and assigns it to the ContainerPercentage field.
func (o *MonthlyUsageAttributionValues) SetContainerPercentage(v float64) {
	o.ContainerPercentage = &v
}

// GetContainerUsage returns the ContainerUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetContainerUsage() float64 {
	if o == nil || o.ContainerUsage == nil {
		var ret float64
		return ret
	}
	return *o.ContainerUsage
}

// GetContainerUsageOk returns a tuple with the ContainerUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetContainerUsageOk() (*float64, bool) {
	if o == nil || o.ContainerUsage == nil {
		return nil, false
	}
	return o.ContainerUsage, true
}

// HasContainerUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasContainerUsage() bool {
	return o != nil && o.ContainerUsage != nil
}

// SetContainerUsage gets a reference to the given float64 and assigns it to the ContainerUsage field.
func (o *MonthlyUsageAttributionValues) SetContainerUsage(v float64) {
	o.ContainerUsage = &v
}

// GetCspmContainersPercentage returns the CspmContainersPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCspmContainersPercentage() float64 {
	if o == nil || o.CspmContainersPercentage == nil {
		var ret float64
		return ret
	}
	return *o.CspmContainersPercentage
}

// GetCspmContainersPercentageOk returns a tuple with the CspmContainersPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCspmContainersPercentageOk() (*float64, bool) {
	if o == nil || o.CspmContainersPercentage == nil {
		return nil, false
	}
	return o.CspmContainersPercentage, true
}

// HasCspmContainersPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCspmContainersPercentage() bool {
	return o != nil && o.CspmContainersPercentage != nil
}

// SetCspmContainersPercentage gets a reference to the given float64 and assigns it to the CspmContainersPercentage field.
func (o *MonthlyUsageAttributionValues) SetCspmContainersPercentage(v float64) {
	o.CspmContainersPercentage = &v
}

// GetCspmContainersUsage returns the CspmContainersUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCspmContainersUsage() float64 {
	if o == nil || o.CspmContainersUsage == nil {
		var ret float64
		return ret
	}
	return *o.CspmContainersUsage
}

// GetCspmContainersUsageOk returns a tuple with the CspmContainersUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCspmContainersUsageOk() (*float64, bool) {
	if o == nil || o.CspmContainersUsage == nil {
		return nil, false
	}
	return o.CspmContainersUsage, true
}

// HasCspmContainersUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCspmContainersUsage() bool {
	return o != nil && o.CspmContainersUsage != nil
}

// SetCspmContainersUsage gets a reference to the given float64 and assigns it to the CspmContainersUsage field.
func (o *MonthlyUsageAttributionValues) SetCspmContainersUsage(v float64) {
	o.CspmContainersUsage = &v
}

// GetCspmHostsPercentage returns the CspmHostsPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCspmHostsPercentage() float64 {
	if o == nil || o.CspmHostsPercentage == nil {
		var ret float64
		return ret
	}
	return *o.CspmHostsPercentage
}

// GetCspmHostsPercentageOk returns a tuple with the CspmHostsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCspmHostsPercentageOk() (*float64, bool) {
	if o == nil || o.CspmHostsPercentage == nil {
		return nil, false
	}
	return o.CspmHostsPercentage, true
}

// HasCspmHostsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCspmHostsPercentage() bool {
	return o != nil && o.CspmHostsPercentage != nil
}

// SetCspmHostsPercentage gets a reference to the given float64 and assigns it to the CspmHostsPercentage field.
func (o *MonthlyUsageAttributionValues) SetCspmHostsPercentage(v float64) {
	o.CspmHostsPercentage = &v
}

// GetCspmHostsUsage returns the CspmHostsUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCspmHostsUsage() float64 {
	if o == nil || o.CspmHostsUsage == nil {
		var ret float64
		return ret
	}
	return *o.CspmHostsUsage
}

// GetCspmHostsUsageOk returns a tuple with the CspmHostsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCspmHostsUsageOk() (*float64, bool) {
	if o == nil || o.CspmHostsUsage == nil {
		return nil, false
	}
	return o.CspmHostsUsage, true
}

// HasCspmHostsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCspmHostsUsage() bool {
	return o != nil && o.CspmHostsUsage != nil
}

// SetCspmHostsUsage gets a reference to the given float64 and assigns it to the CspmHostsUsage field.
func (o *MonthlyUsageAttributionValues) SetCspmHostsUsage(v float64) {
	o.CspmHostsUsage = &v
}

// GetCustomIngestedTimeseriesPercentage returns the CustomIngestedTimeseriesPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCustomIngestedTimeseriesPercentage() float64 {
	if o == nil || o.CustomIngestedTimeseriesPercentage == nil {
		var ret float64
		return ret
	}
	return *o.CustomIngestedTimeseriesPercentage
}

// GetCustomIngestedTimeseriesPercentageOk returns a tuple with the CustomIngestedTimeseriesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCustomIngestedTimeseriesPercentageOk() (*float64, bool) {
	if o == nil || o.CustomIngestedTimeseriesPercentage == nil {
		return nil, false
	}
	return o.CustomIngestedTimeseriesPercentage, true
}

// HasCustomIngestedTimeseriesPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCustomIngestedTimeseriesPercentage() bool {
	return o != nil && o.CustomIngestedTimeseriesPercentage != nil
}

// SetCustomIngestedTimeseriesPercentage gets a reference to the given float64 and assigns it to the CustomIngestedTimeseriesPercentage field.
func (o *MonthlyUsageAttributionValues) SetCustomIngestedTimeseriesPercentage(v float64) {
	o.CustomIngestedTimeseriesPercentage = &v
}

// GetCustomIngestedTimeseriesUsage returns the CustomIngestedTimeseriesUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCustomIngestedTimeseriesUsage() float64 {
	if o == nil || o.CustomIngestedTimeseriesUsage == nil {
		var ret float64
		return ret
	}
	return *o.CustomIngestedTimeseriesUsage
}

// GetCustomIngestedTimeseriesUsageOk returns a tuple with the CustomIngestedTimeseriesUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCustomIngestedTimeseriesUsageOk() (*float64, bool) {
	if o == nil || o.CustomIngestedTimeseriesUsage == nil {
		return nil, false
	}
	return o.CustomIngestedTimeseriesUsage, true
}

// HasCustomIngestedTimeseriesUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCustomIngestedTimeseriesUsage() bool {
	return o != nil && o.CustomIngestedTimeseriesUsage != nil
}

// SetCustomIngestedTimeseriesUsage gets a reference to the given float64 and assigns it to the CustomIngestedTimeseriesUsage field.
func (o *MonthlyUsageAttributionValues) SetCustomIngestedTimeseriesUsage(v float64) {
	o.CustomIngestedTimeseriesUsage = &v
}

// GetCustomTimeseriesPercentage returns the CustomTimeseriesPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCustomTimeseriesPercentage() float64 {
	if o == nil || o.CustomTimeseriesPercentage == nil {
		var ret float64
		return ret
	}
	return *o.CustomTimeseriesPercentage
}

// GetCustomTimeseriesPercentageOk returns a tuple with the CustomTimeseriesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCustomTimeseriesPercentageOk() (*float64, bool) {
	if o == nil || o.CustomTimeseriesPercentage == nil {
		return nil, false
	}
	return o.CustomTimeseriesPercentage, true
}

// HasCustomTimeseriesPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCustomTimeseriesPercentage() bool {
	return o != nil && o.CustomTimeseriesPercentage != nil
}

// SetCustomTimeseriesPercentage gets a reference to the given float64 and assigns it to the CustomTimeseriesPercentage field.
func (o *MonthlyUsageAttributionValues) SetCustomTimeseriesPercentage(v float64) {
	o.CustomTimeseriesPercentage = &v
}

// GetCustomTimeseriesUsage returns the CustomTimeseriesUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCustomTimeseriesUsage() float64 {
	if o == nil || o.CustomTimeseriesUsage == nil {
		var ret float64
		return ret
	}
	return *o.CustomTimeseriesUsage
}

// GetCustomTimeseriesUsageOk returns a tuple with the CustomTimeseriesUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCustomTimeseriesUsageOk() (*float64, bool) {
	if o == nil || o.CustomTimeseriesUsage == nil {
		return nil, false
	}
	return o.CustomTimeseriesUsage, true
}

// HasCustomTimeseriesUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCustomTimeseriesUsage() bool {
	return o != nil && o.CustomTimeseriesUsage != nil
}

// SetCustomTimeseriesUsage gets a reference to the given float64 and assigns it to the CustomTimeseriesUsage field.
func (o *MonthlyUsageAttributionValues) SetCustomTimeseriesUsage(v float64) {
	o.CustomTimeseriesUsage = &v
}

// GetCwsContainersPercentage returns the CwsContainersPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCwsContainersPercentage() float64 {
	if o == nil || o.CwsContainersPercentage == nil {
		var ret float64
		return ret
	}
	return *o.CwsContainersPercentage
}

// GetCwsContainersPercentageOk returns a tuple with the CwsContainersPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCwsContainersPercentageOk() (*float64, bool) {
	if o == nil || o.CwsContainersPercentage == nil {
		return nil, false
	}
	return o.CwsContainersPercentage, true
}

// HasCwsContainersPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCwsContainersPercentage() bool {
	return o != nil && o.CwsContainersPercentage != nil
}

// SetCwsContainersPercentage gets a reference to the given float64 and assigns it to the CwsContainersPercentage field.
func (o *MonthlyUsageAttributionValues) SetCwsContainersPercentage(v float64) {
	o.CwsContainersPercentage = &v
}

// GetCwsContainersUsage returns the CwsContainersUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCwsContainersUsage() float64 {
	if o == nil || o.CwsContainersUsage == nil {
		var ret float64
		return ret
	}
	return *o.CwsContainersUsage
}

// GetCwsContainersUsageOk returns a tuple with the CwsContainersUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCwsContainersUsageOk() (*float64, bool) {
	if o == nil || o.CwsContainersUsage == nil {
		return nil, false
	}
	return o.CwsContainersUsage, true
}

// HasCwsContainersUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCwsContainersUsage() bool {
	return o != nil && o.CwsContainersUsage != nil
}

// SetCwsContainersUsage gets a reference to the given float64 and assigns it to the CwsContainersUsage field.
func (o *MonthlyUsageAttributionValues) SetCwsContainersUsage(v float64) {
	o.CwsContainersUsage = &v
}

// GetCwsHostsPercentage returns the CwsHostsPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCwsHostsPercentage() float64 {
	if o == nil || o.CwsHostsPercentage == nil {
		var ret float64
		return ret
	}
	return *o.CwsHostsPercentage
}

// GetCwsHostsPercentageOk returns a tuple with the CwsHostsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCwsHostsPercentageOk() (*float64, bool) {
	if o == nil || o.CwsHostsPercentage == nil {
		return nil, false
	}
	return o.CwsHostsPercentage, true
}

// HasCwsHostsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCwsHostsPercentage() bool {
	return o != nil && o.CwsHostsPercentage != nil
}

// SetCwsHostsPercentage gets a reference to the given float64 and assigns it to the CwsHostsPercentage field.
func (o *MonthlyUsageAttributionValues) SetCwsHostsPercentage(v float64) {
	o.CwsHostsPercentage = &v
}

// GetCwsHostsUsage returns the CwsHostsUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCwsHostsUsage() float64 {
	if o == nil || o.CwsHostsUsage == nil {
		var ret float64
		return ret
	}
	return *o.CwsHostsUsage
}

// GetCwsHostsUsageOk returns a tuple with the CwsHostsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCwsHostsUsageOk() (*float64, bool) {
	if o == nil || o.CwsHostsUsage == nil {
		return nil, false
	}
	return o.CwsHostsUsage, true
}

// HasCwsHostsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCwsHostsUsage() bool {
	return o != nil && o.CwsHostsUsage != nil
}

// SetCwsHostsUsage gets a reference to the given float64 and assigns it to the CwsHostsUsage field.
func (o *MonthlyUsageAttributionValues) SetCwsHostsUsage(v float64) {
	o.CwsHostsUsage = &v
}

// GetDbmHostsPercentage returns the DbmHostsPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetDbmHostsPercentage() float64 {
	if o == nil || o.DbmHostsPercentage == nil {
		var ret float64
		return ret
	}
	return *o.DbmHostsPercentage
}

// GetDbmHostsPercentageOk returns a tuple with the DbmHostsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetDbmHostsPercentageOk() (*float64, bool) {
	if o == nil || o.DbmHostsPercentage == nil {
		return nil, false
	}
	return o.DbmHostsPercentage, true
}

// HasDbmHostsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasDbmHostsPercentage() bool {
	return o != nil && o.DbmHostsPercentage != nil
}

// SetDbmHostsPercentage gets a reference to the given float64 and assigns it to the DbmHostsPercentage field.
func (o *MonthlyUsageAttributionValues) SetDbmHostsPercentage(v float64) {
	o.DbmHostsPercentage = &v
}

// GetDbmHostsUsage returns the DbmHostsUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetDbmHostsUsage() float64 {
	if o == nil || o.DbmHostsUsage == nil {
		var ret float64
		return ret
	}
	return *o.DbmHostsUsage
}

// GetDbmHostsUsageOk returns a tuple with the DbmHostsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetDbmHostsUsageOk() (*float64, bool) {
	if o == nil || o.DbmHostsUsage == nil {
		return nil, false
	}
	return o.DbmHostsUsage, true
}

// HasDbmHostsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasDbmHostsUsage() bool {
	return o != nil && o.DbmHostsUsage != nil
}

// SetDbmHostsUsage gets a reference to the given float64 and assigns it to the DbmHostsUsage field.
func (o *MonthlyUsageAttributionValues) SetDbmHostsUsage(v float64) {
	o.DbmHostsUsage = &v
}

// GetDbmQueriesPercentage returns the DbmQueriesPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetDbmQueriesPercentage() float64 {
	if o == nil || o.DbmQueriesPercentage == nil {
		var ret float64
		return ret
	}
	return *o.DbmQueriesPercentage
}

// GetDbmQueriesPercentageOk returns a tuple with the DbmQueriesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetDbmQueriesPercentageOk() (*float64, bool) {
	if o == nil || o.DbmQueriesPercentage == nil {
		return nil, false
	}
	return o.DbmQueriesPercentage, true
}

// HasDbmQueriesPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasDbmQueriesPercentage() bool {
	return o != nil && o.DbmQueriesPercentage != nil
}

// SetDbmQueriesPercentage gets a reference to the given float64 and assigns it to the DbmQueriesPercentage field.
func (o *MonthlyUsageAttributionValues) SetDbmQueriesPercentage(v float64) {
	o.DbmQueriesPercentage = &v
}

// GetDbmQueriesUsage returns the DbmQueriesUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetDbmQueriesUsage() float64 {
	if o == nil || o.DbmQueriesUsage == nil {
		var ret float64
		return ret
	}
	return *o.DbmQueriesUsage
}

// GetDbmQueriesUsageOk returns a tuple with the DbmQueriesUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetDbmQueriesUsageOk() (*float64, bool) {
	if o == nil || o.DbmQueriesUsage == nil {
		return nil, false
	}
	return o.DbmQueriesUsage, true
}

// HasDbmQueriesUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasDbmQueriesUsage() bool {
	return o != nil && o.DbmQueriesUsage != nil
}

// SetDbmQueriesUsage gets a reference to the given float64 and assigns it to the DbmQueriesUsage field.
func (o *MonthlyUsageAttributionValues) SetDbmQueriesUsage(v float64) {
	o.DbmQueriesUsage = &v
}

// GetEstimatedIndexedLogsPercentage returns the EstimatedIndexedLogsPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetEstimatedIndexedLogsPercentage() float64 {
	if o == nil || o.EstimatedIndexedLogsPercentage == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIndexedLogsPercentage
}

// GetEstimatedIndexedLogsPercentageOk returns a tuple with the EstimatedIndexedLogsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetEstimatedIndexedLogsPercentageOk() (*float64, bool) {
	if o == nil || o.EstimatedIndexedLogsPercentage == nil {
		return nil, false
	}
	return o.EstimatedIndexedLogsPercentage, true
}

// HasEstimatedIndexedLogsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedIndexedLogsPercentage() bool {
	return o != nil && o.EstimatedIndexedLogsPercentage != nil
}

// SetEstimatedIndexedLogsPercentage gets a reference to the given float64 and assigns it to the EstimatedIndexedLogsPercentage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedIndexedLogsPercentage(v float64) {
	o.EstimatedIndexedLogsPercentage = &v
}

// GetEstimatedIndexedLogsUsage returns the EstimatedIndexedLogsUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetEstimatedIndexedLogsUsage() float64 {
	if o == nil || o.EstimatedIndexedLogsUsage == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIndexedLogsUsage
}

// GetEstimatedIndexedLogsUsageOk returns a tuple with the EstimatedIndexedLogsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetEstimatedIndexedLogsUsageOk() (*float64, bool) {
	if o == nil || o.EstimatedIndexedLogsUsage == nil {
		return nil, false
	}
	return o.EstimatedIndexedLogsUsage, true
}

// HasEstimatedIndexedLogsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedIndexedLogsUsage() bool {
	return o != nil && o.EstimatedIndexedLogsUsage != nil
}

// SetEstimatedIndexedLogsUsage gets a reference to the given float64 and assigns it to the EstimatedIndexedLogsUsage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedIndexedLogsUsage(v float64) {
	o.EstimatedIndexedLogsUsage = &v
}

// GetEstimatedIndexedSpansPercentage returns the EstimatedIndexedSpansPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetEstimatedIndexedSpansPercentage() float64 {
	if o == nil || o.EstimatedIndexedSpansPercentage == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIndexedSpansPercentage
}

// GetEstimatedIndexedSpansPercentageOk returns a tuple with the EstimatedIndexedSpansPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetEstimatedIndexedSpansPercentageOk() (*float64, bool) {
	if o == nil || o.EstimatedIndexedSpansPercentage == nil {
		return nil, false
	}
	return o.EstimatedIndexedSpansPercentage, true
}

// HasEstimatedIndexedSpansPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedIndexedSpansPercentage() bool {
	return o != nil && o.EstimatedIndexedSpansPercentage != nil
}

// SetEstimatedIndexedSpansPercentage gets a reference to the given float64 and assigns it to the EstimatedIndexedSpansPercentage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedIndexedSpansPercentage(v float64) {
	o.EstimatedIndexedSpansPercentage = &v
}

// GetEstimatedIndexedSpansUsage returns the EstimatedIndexedSpansUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetEstimatedIndexedSpansUsage() float64 {
	if o == nil || o.EstimatedIndexedSpansUsage == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIndexedSpansUsage
}

// GetEstimatedIndexedSpansUsageOk returns a tuple with the EstimatedIndexedSpansUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetEstimatedIndexedSpansUsageOk() (*float64, bool) {
	if o == nil || o.EstimatedIndexedSpansUsage == nil {
		return nil, false
	}
	return o.EstimatedIndexedSpansUsage, true
}

// HasEstimatedIndexedSpansUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedIndexedSpansUsage() bool {
	return o != nil && o.EstimatedIndexedSpansUsage != nil
}

// SetEstimatedIndexedSpansUsage gets a reference to the given float64 and assigns it to the EstimatedIndexedSpansUsage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedIndexedSpansUsage(v float64) {
	o.EstimatedIndexedSpansUsage = &v
}

// GetEstimatedIngestedLogsPercentage returns the EstimatedIngestedLogsPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetEstimatedIngestedLogsPercentage() float64 {
	if o == nil || o.EstimatedIngestedLogsPercentage == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIngestedLogsPercentage
}

// GetEstimatedIngestedLogsPercentageOk returns a tuple with the EstimatedIngestedLogsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetEstimatedIngestedLogsPercentageOk() (*float64, bool) {
	if o == nil || o.EstimatedIngestedLogsPercentage == nil {
		return nil, false
	}
	return o.EstimatedIngestedLogsPercentage, true
}

// HasEstimatedIngestedLogsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedIngestedLogsPercentage() bool {
	return o != nil && o.EstimatedIngestedLogsPercentage != nil
}

// SetEstimatedIngestedLogsPercentage gets a reference to the given float64 and assigns it to the EstimatedIngestedLogsPercentage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedIngestedLogsPercentage(v float64) {
	o.EstimatedIngestedLogsPercentage = &v
}

// GetEstimatedIngestedLogsUsage returns the EstimatedIngestedLogsUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetEstimatedIngestedLogsUsage() float64 {
	if o == nil || o.EstimatedIngestedLogsUsage == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIngestedLogsUsage
}

// GetEstimatedIngestedLogsUsageOk returns a tuple with the EstimatedIngestedLogsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetEstimatedIngestedLogsUsageOk() (*float64, bool) {
	if o == nil || o.EstimatedIngestedLogsUsage == nil {
		return nil, false
	}
	return o.EstimatedIngestedLogsUsage, true
}

// HasEstimatedIngestedLogsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedIngestedLogsUsage() bool {
	return o != nil && o.EstimatedIngestedLogsUsage != nil
}

// SetEstimatedIngestedLogsUsage gets a reference to the given float64 and assigns it to the EstimatedIngestedLogsUsage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedIngestedLogsUsage(v float64) {
	o.EstimatedIngestedLogsUsage = &v
}

// GetEstimatedIngestedSpansPercentage returns the EstimatedIngestedSpansPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetEstimatedIngestedSpansPercentage() float64 {
	if o == nil || o.EstimatedIngestedSpansPercentage == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIngestedSpansPercentage
}

// GetEstimatedIngestedSpansPercentageOk returns a tuple with the EstimatedIngestedSpansPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetEstimatedIngestedSpansPercentageOk() (*float64, bool) {
	if o == nil || o.EstimatedIngestedSpansPercentage == nil {
		return nil, false
	}
	return o.EstimatedIngestedSpansPercentage, true
}

// HasEstimatedIngestedSpansPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedIngestedSpansPercentage() bool {
	return o != nil && o.EstimatedIngestedSpansPercentage != nil
}

// SetEstimatedIngestedSpansPercentage gets a reference to the given float64 and assigns it to the EstimatedIngestedSpansPercentage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedIngestedSpansPercentage(v float64) {
	o.EstimatedIngestedSpansPercentage = &v
}

// GetEstimatedIngestedSpansUsage returns the EstimatedIngestedSpansUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetEstimatedIngestedSpansUsage() float64 {
	if o == nil || o.EstimatedIngestedSpansUsage == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIngestedSpansUsage
}

// GetEstimatedIngestedSpansUsageOk returns a tuple with the EstimatedIngestedSpansUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetEstimatedIngestedSpansUsageOk() (*float64, bool) {
	if o == nil || o.EstimatedIngestedSpansUsage == nil {
		return nil, false
	}
	return o.EstimatedIngestedSpansUsage, true
}

// HasEstimatedIngestedSpansUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedIngestedSpansUsage() bool {
	return o != nil && o.EstimatedIngestedSpansUsage != nil
}

// SetEstimatedIngestedSpansUsage gets a reference to the given float64 and assigns it to the EstimatedIngestedSpansUsage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedIngestedSpansUsage(v float64) {
	o.EstimatedIngestedSpansUsage = &v
}

// GetEstimatedRumSessionsPercentage returns the EstimatedRumSessionsPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetEstimatedRumSessionsPercentage() float64 {
	if o == nil || o.EstimatedRumSessionsPercentage == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedRumSessionsPercentage
}

// GetEstimatedRumSessionsPercentageOk returns a tuple with the EstimatedRumSessionsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetEstimatedRumSessionsPercentageOk() (*float64, bool) {
	if o == nil || o.EstimatedRumSessionsPercentage == nil {
		return nil, false
	}
	return o.EstimatedRumSessionsPercentage, true
}

// HasEstimatedRumSessionsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedRumSessionsPercentage() bool {
	return o != nil && o.EstimatedRumSessionsPercentage != nil
}

// SetEstimatedRumSessionsPercentage gets a reference to the given float64 and assigns it to the EstimatedRumSessionsPercentage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedRumSessionsPercentage(v float64) {
	o.EstimatedRumSessionsPercentage = &v
}

// GetEstimatedRumSessionsUsage returns the EstimatedRumSessionsUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetEstimatedRumSessionsUsage() float64 {
	if o == nil || o.EstimatedRumSessionsUsage == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedRumSessionsUsage
}

// GetEstimatedRumSessionsUsageOk returns a tuple with the EstimatedRumSessionsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetEstimatedRumSessionsUsageOk() (*float64, bool) {
	if o == nil || o.EstimatedRumSessionsUsage == nil {
		return nil, false
	}
	return o.EstimatedRumSessionsUsage, true
}

// HasEstimatedRumSessionsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedRumSessionsUsage() bool {
	return o != nil && o.EstimatedRumSessionsUsage != nil
}

// SetEstimatedRumSessionsUsage gets a reference to the given float64 and assigns it to the EstimatedRumSessionsUsage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedRumSessionsUsage(v float64) {
	o.EstimatedRumSessionsUsage = &v
}

// GetFargatePercentage returns the FargatePercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetFargatePercentage() float64 {
	if o == nil || o.FargatePercentage == nil {
		var ret float64
		return ret
	}
	return *o.FargatePercentage
}

// GetFargatePercentageOk returns a tuple with the FargatePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetFargatePercentageOk() (*float64, bool) {
	if o == nil || o.FargatePercentage == nil {
		return nil, false
	}
	return o.FargatePercentage, true
}

// HasFargatePercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasFargatePercentage() bool {
	return o != nil && o.FargatePercentage != nil
}

// SetFargatePercentage gets a reference to the given float64 and assigns it to the FargatePercentage field.
func (o *MonthlyUsageAttributionValues) SetFargatePercentage(v float64) {
	o.FargatePercentage = &v
}

// GetFargateUsage returns the FargateUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetFargateUsage() float64 {
	if o == nil || o.FargateUsage == nil {
		var ret float64
		return ret
	}
	return *o.FargateUsage
}

// GetFargateUsageOk returns a tuple with the FargateUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetFargateUsageOk() (*float64, bool) {
	if o == nil || o.FargateUsage == nil {
		return nil, false
	}
	return o.FargateUsage, true
}

// HasFargateUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasFargateUsage() bool {
	return o != nil && o.FargateUsage != nil
}

// SetFargateUsage gets a reference to the given float64 and assigns it to the FargateUsage field.
func (o *MonthlyUsageAttributionValues) SetFargateUsage(v float64) {
	o.FargateUsage = &v
}

// GetFunctionsPercentage returns the FunctionsPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetFunctionsPercentage() float64 {
	if o == nil || o.FunctionsPercentage == nil {
		var ret float64
		return ret
	}
	return *o.FunctionsPercentage
}

// GetFunctionsPercentageOk returns a tuple with the FunctionsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetFunctionsPercentageOk() (*float64, bool) {
	if o == nil || o.FunctionsPercentage == nil {
		return nil, false
	}
	return o.FunctionsPercentage, true
}

// HasFunctionsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasFunctionsPercentage() bool {
	return o != nil && o.FunctionsPercentage != nil
}

// SetFunctionsPercentage gets a reference to the given float64 and assigns it to the FunctionsPercentage field.
func (o *MonthlyUsageAttributionValues) SetFunctionsPercentage(v float64) {
	o.FunctionsPercentage = &v
}

// GetFunctionsUsage returns the FunctionsUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetFunctionsUsage() float64 {
	if o == nil || o.FunctionsUsage == nil {
		var ret float64
		return ret
	}
	return *o.FunctionsUsage
}

// GetFunctionsUsageOk returns a tuple with the FunctionsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetFunctionsUsageOk() (*float64, bool) {
	if o == nil || o.FunctionsUsage == nil {
		return nil, false
	}
	return o.FunctionsUsage, true
}

// HasFunctionsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasFunctionsUsage() bool {
	return o != nil && o.FunctionsUsage != nil
}

// SetFunctionsUsage gets a reference to the given float64 and assigns it to the FunctionsUsage field.
func (o *MonthlyUsageAttributionValues) SetFunctionsUsage(v float64) {
	o.FunctionsUsage = &v
}

// GetInfraHostPercentage returns the InfraHostPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetInfraHostPercentage() float64 {
	if o == nil || o.InfraHostPercentage == nil {
		var ret float64
		return ret
	}
	return *o.InfraHostPercentage
}

// GetInfraHostPercentageOk returns a tuple with the InfraHostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetInfraHostPercentageOk() (*float64, bool) {
	if o == nil || o.InfraHostPercentage == nil {
		return nil, false
	}
	return o.InfraHostPercentage, true
}

// HasInfraHostPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasInfraHostPercentage() bool {
	return o != nil && o.InfraHostPercentage != nil
}

// SetInfraHostPercentage gets a reference to the given float64 and assigns it to the InfraHostPercentage field.
func (o *MonthlyUsageAttributionValues) SetInfraHostPercentage(v float64) {
	o.InfraHostPercentage = &v
}

// GetInfraHostUsage returns the InfraHostUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetInfraHostUsage() float64 {
	if o == nil || o.InfraHostUsage == nil {
		var ret float64
		return ret
	}
	return *o.InfraHostUsage
}

// GetInfraHostUsageOk returns a tuple with the InfraHostUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetInfraHostUsageOk() (*float64, bool) {
	if o == nil || o.InfraHostUsage == nil {
		return nil, false
	}
	return o.InfraHostUsage, true
}

// HasInfraHostUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasInfraHostUsage() bool {
	return o != nil && o.InfraHostUsage != nil
}

// SetInfraHostUsage gets a reference to the given float64 and assigns it to the InfraHostUsage field.
func (o *MonthlyUsageAttributionValues) SetInfraHostUsage(v float64) {
	o.InfraHostUsage = &v
}

// GetInvocationsPercentage returns the InvocationsPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetInvocationsPercentage() float64 {
	if o == nil || o.InvocationsPercentage == nil {
		var ret float64
		return ret
	}
	return *o.InvocationsPercentage
}

// GetInvocationsPercentageOk returns a tuple with the InvocationsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetInvocationsPercentageOk() (*float64, bool) {
	if o == nil || o.InvocationsPercentage == nil {
		return nil, false
	}
	return o.InvocationsPercentage, true
}

// HasInvocationsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasInvocationsPercentage() bool {
	return o != nil && o.InvocationsPercentage != nil
}

// SetInvocationsPercentage gets a reference to the given float64 and assigns it to the InvocationsPercentage field.
func (o *MonthlyUsageAttributionValues) SetInvocationsPercentage(v float64) {
	o.InvocationsPercentage = &v
}

// GetInvocationsUsage returns the InvocationsUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetInvocationsUsage() float64 {
	if o == nil || o.InvocationsUsage == nil {
		var ret float64
		return ret
	}
	return *o.InvocationsUsage
}

// GetInvocationsUsageOk returns a tuple with the InvocationsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetInvocationsUsageOk() (*float64, bool) {
	if o == nil || o.InvocationsUsage == nil {
		return nil, false
	}
	return o.InvocationsUsage, true
}

// HasInvocationsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasInvocationsUsage() bool {
	return o != nil && o.InvocationsUsage != nil
}

// SetInvocationsUsage gets a reference to the given float64 and assigns it to the InvocationsUsage field.
func (o *MonthlyUsageAttributionValues) SetInvocationsUsage(v float64) {
	o.InvocationsUsage = &v
}

// GetMobileAppTestingPercentage returns the MobileAppTestingPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetMobileAppTestingPercentage() float64 {
	if o == nil || o.MobileAppTestingPercentage == nil {
		var ret float64
		return ret
	}
	return *o.MobileAppTestingPercentage
}

// GetMobileAppTestingPercentageOk returns a tuple with the MobileAppTestingPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetMobileAppTestingPercentageOk() (*float64, bool) {
	if o == nil || o.MobileAppTestingPercentage == nil {
		return nil, false
	}
	return o.MobileAppTestingPercentage, true
}

// HasMobileAppTestingPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasMobileAppTestingPercentage() bool {
	return o != nil && o.MobileAppTestingPercentage != nil
}

// SetMobileAppTestingPercentage gets a reference to the given float64 and assigns it to the MobileAppTestingPercentage field.
func (o *MonthlyUsageAttributionValues) SetMobileAppTestingPercentage(v float64) {
	o.MobileAppTestingPercentage = &v
}

// GetMobileAppTestingUsage returns the MobileAppTestingUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetMobileAppTestingUsage() float64 {
	if o == nil || o.MobileAppTestingUsage == nil {
		var ret float64
		return ret
	}
	return *o.MobileAppTestingUsage
}

// GetMobileAppTestingUsageOk returns a tuple with the MobileAppTestingUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetMobileAppTestingUsageOk() (*float64, bool) {
	if o == nil || o.MobileAppTestingUsage == nil {
		return nil, false
	}
	return o.MobileAppTestingUsage, true
}

// HasMobileAppTestingUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasMobileAppTestingUsage() bool {
	return o != nil && o.MobileAppTestingUsage != nil
}

// SetMobileAppTestingUsage gets a reference to the given float64 and assigns it to the MobileAppTestingUsage field.
func (o *MonthlyUsageAttributionValues) SetMobileAppTestingUsage(v float64) {
	o.MobileAppTestingUsage = &v
}

// GetNdmNetflowPercentage returns the NdmNetflowPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetNdmNetflowPercentage() float64 {
	if o == nil || o.NdmNetflowPercentage == nil {
		var ret float64
		return ret
	}
	return *o.NdmNetflowPercentage
}

// GetNdmNetflowPercentageOk returns a tuple with the NdmNetflowPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetNdmNetflowPercentageOk() (*float64, bool) {
	if o == nil || o.NdmNetflowPercentage == nil {
		return nil, false
	}
	return o.NdmNetflowPercentage, true
}

// HasNdmNetflowPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasNdmNetflowPercentage() bool {
	return o != nil && o.NdmNetflowPercentage != nil
}

// SetNdmNetflowPercentage gets a reference to the given float64 and assigns it to the NdmNetflowPercentage field.
func (o *MonthlyUsageAttributionValues) SetNdmNetflowPercentage(v float64) {
	o.NdmNetflowPercentage = &v
}

// GetNdmNetflowUsage returns the NdmNetflowUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetNdmNetflowUsage() float64 {
	if o == nil || o.NdmNetflowUsage == nil {
		var ret float64
		return ret
	}
	return *o.NdmNetflowUsage
}

// GetNdmNetflowUsageOk returns a tuple with the NdmNetflowUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetNdmNetflowUsageOk() (*float64, bool) {
	if o == nil || o.NdmNetflowUsage == nil {
		return nil, false
	}
	return o.NdmNetflowUsage, true
}

// HasNdmNetflowUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasNdmNetflowUsage() bool {
	return o != nil && o.NdmNetflowUsage != nil
}

// SetNdmNetflowUsage gets a reference to the given float64 and assigns it to the NdmNetflowUsage field.
func (o *MonthlyUsageAttributionValues) SetNdmNetflowUsage(v float64) {
	o.NdmNetflowUsage = &v
}

// GetNpmHostPercentage returns the NpmHostPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetNpmHostPercentage() float64 {
	if o == nil || o.NpmHostPercentage == nil {
		var ret float64
		return ret
	}
	return *o.NpmHostPercentage
}

// GetNpmHostPercentageOk returns a tuple with the NpmHostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetNpmHostPercentageOk() (*float64, bool) {
	if o == nil || o.NpmHostPercentage == nil {
		return nil, false
	}
	return o.NpmHostPercentage, true
}

// HasNpmHostPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasNpmHostPercentage() bool {
	return o != nil && o.NpmHostPercentage != nil
}

// SetNpmHostPercentage gets a reference to the given float64 and assigns it to the NpmHostPercentage field.
func (o *MonthlyUsageAttributionValues) SetNpmHostPercentage(v float64) {
	o.NpmHostPercentage = &v
}

// GetNpmHostUsage returns the NpmHostUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetNpmHostUsage() float64 {
	if o == nil || o.NpmHostUsage == nil {
		var ret float64
		return ret
	}
	return *o.NpmHostUsage
}

// GetNpmHostUsageOk returns a tuple with the NpmHostUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetNpmHostUsageOk() (*float64, bool) {
	if o == nil || o.NpmHostUsage == nil {
		return nil, false
	}
	return o.NpmHostUsage, true
}

// HasNpmHostUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasNpmHostUsage() bool {
	return o != nil && o.NpmHostUsage != nil
}

// SetNpmHostUsage gets a reference to the given float64 and assigns it to the NpmHostUsage field.
func (o *MonthlyUsageAttributionValues) SetNpmHostUsage(v float64) {
	o.NpmHostUsage = &v
}

// GetObsPipelineBytesPercentage returns the ObsPipelineBytesPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetObsPipelineBytesPercentage() float64 {
	if o == nil || o.ObsPipelineBytesPercentage == nil {
		var ret float64
		return ret
	}
	return *o.ObsPipelineBytesPercentage
}

// GetObsPipelineBytesPercentageOk returns a tuple with the ObsPipelineBytesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetObsPipelineBytesPercentageOk() (*float64, bool) {
	if o == nil || o.ObsPipelineBytesPercentage == nil {
		return nil, false
	}
	return o.ObsPipelineBytesPercentage, true
}

// HasObsPipelineBytesPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasObsPipelineBytesPercentage() bool {
	return o != nil && o.ObsPipelineBytesPercentage != nil
}

// SetObsPipelineBytesPercentage gets a reference to the given float64 and assigns it to the ObsPipelineBytesPercentage field.
func (o *MonthlyUsageAttributionValues) SetObsPipelineBytesPercentage(v float64) {
	o.ObsPipelineBytesPercentage = &v
}

// GetObsPipelineBytesUsage returns the ObsPipelineBytesUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetObsPipelineBytesUsage() float64 {
	if o == nil || o.ObsPipelineBytesUsage == nil {
		var ret float64
		return ret
	}
	return *o.ObsPipelineBytesUsage
}

// GetObsPipelineBytesUsageOk returns a tuple with the ObsPipelineBytesUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetObsPipelineBytesUsageOk() (*float64, bool) {
	if o == nil || o.ObsPipelineBytesUsage == nil {
		return nil, false
	}
	return o.ObsPipelineBytesUsage, true
}

// HasObsPipelineBytesUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasObsPipelineBytesUsage() bool {
	return o != nil && o.ObsPipelineBytesUsage != nil
}

// SetObsPipelineBytesUsage gets a reference to the given float64 and assigns it to the ObsPipelineBytesUsage field.
func (o *MonthlyUsageAttributionValues) SetObsPipelineBytesUsage(v float64) {
	o.ObsPipelineBytesUsage = &v
}

// GetProfiledContainerPercentage returns the ProfiledContainerPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetProfiledContainerPercentage() float64 {
	if o == nil || o.ProfiledContainerPercentage == nil {
		var ret float64
		return ret
	}
	return *o.ProfiledContainerPercentage
}

// GetProfiledContainerPercentageOk returns a tuple with the ProfiledContainerPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetProfiledContainerPercentageOk() (*float64, bool) {
	if o == nil || o.ProfiledContainerPercentage == nil {
		return nil, false
	}
	return o.ProfiledContainerPercentage, true
}

// HasProfiledContainerPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasProfiledContainerPercentage() bool {
	return o != nil && o.ProfiledContainerPercentage != nil
}

// SetProfiledContainerPercentage gets a reference to the given float64 and assigns it to the ProfiledContainerPercentage field.
func (o *MonthlyUsageAttributionValues) SetProfiledContainerPercentage(v float64) {
	o.ProfiledContainerPercentage = &v
}

// GetProfiledContainerUsage returns the ProfiledContainerUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetProfiledContainerUsage() float64 {
	if o == nil || o.ProfiledContainerUsage == nil {
		var ret float64
		return ret
	}
	return *o.ProfiledContainerUsage
}

// GetProfiledContainerUsageOk returns a tuple with the ProfiledContainerUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetProfiledContainerUsageOk() (*float64, bool) {
	if o == nil || o.ProfiledContainerUsage == nil {
		return nil, false
	}
	return o.ProfiledContainerUsage, true
}

// HasProfiledContainerUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasProfiledContainerUsage() bool {
	return o != nil && o.ProfiledContainerUsage != nil
}

// SetProfiledContainerUsage gets a reference to the given float64 and assigns it to the ProfiledContainerUsage field.
func (o *MonthlyUsageAttributionValues) SetProfiledContainerUsage(v float64) {
	o.ProfiledContainerUsage = &v
}

// GetProfiledFargatePercentage returns the ProfiledFargatePercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetProfiledFargatePercentage() float64 {
	if o == nil || o.ProfiledFargatePercentage == nil {
		var ret float64
		return ret
	}
	return *o.ProfiledFargatePercentage
}

// GetProfiledFargatePercentageOk returns a tuple with the ProfiledFargatePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetProfiledFargatePercentageOk() (*float64, bool) {
	if o == nil || o.ProfiledFargatePercentage == nil {
		return nil, false
	}
	return o.ProfiledFargatePercentage, true
}

// HasProfiledFargatePercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasProfiledFargatePercentage() bool {
	return o != nil && o.ProfiledFargatePercentage != nil
}

// SetProfiledFargatePercentage gets a reference to the given float64 and assigns it to the ProfiledFargatePercentage field.
func (o *MonthlyUsageAttributionValues) SetProfiledFargatePercentage(v float64) {
	o.ProfiledFargatePercentage = &v
}

// GetProfiledFargateUsage returns the ProfiledFargateUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetProfiledFargateUsage() float64 {
	if o == nil || o.ProfiledFargateUsage == nil {
		var ret float64
		return ret
	}
	return *o.ProfiledFargateUsage
}

// GetProfiledFargateUsageOk returns a tuple with the ProfiledFargateUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetProfiledFargateUsageOk() (*float64, bool) {
	if o == nil || o.ProfiledFargateUsage == nil {
		return nil, false
	}
	return o.ProfiledFargateUsage, true
}

// HasProfiledFargateUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasProfiledFargateUsage() bool {
	return o != nil && o.ProfiledFargateUsage != nil
}

// SetProfiledFargateUsage gets a reference to the given float64 and assigns it to the ProfiledFargateUsage field.
func (o *MonthlyUsageAttributionValues) SetProfiledFargateUsage(v float64) {
	o.ProfiledFargateUsage = &v
}

// GetProfiledHostPercentage returns the ProfiledHostPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetProfiledHostPercentage() float64 {
	if o == nil || o.ProfiledHostPercentage == nil {
		var ret float64
		return ret
	}
	return *o.ProfiledHostPercentage
}

// GetProfiledHostPercentageOk returns a tuple with the ProfiledHostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetProfiledHostPercentageOk() (*float64, bool) {
	if o == nil || o.ProfiledHostPercentage == nil {
		return nil, false
	}
	return o.ProfiledHostPercentage, true
}

// HasProfiledHostPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasProfiledHostPercentage() bool {
	return o != nil && o.ProfiledHostPercentage != nil
}

// SetProfiledHostPercentage gets a reference to the given float64 and assigns it to the ProfiledHostPercentage field.
func (o *MonthlyUsageAttributionValues) SetProfiledHostPercentage(v float64) {
	o.ProfiledHostPercentage = &v
}

// GetProfiledHostUsage returns the ProfiledHostUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetProfiledHostUsage() float64 {
	if o == nil || o.ProfiledHostUsage == nil {
		var ret float64
		return ret
	}
	return *o.ProfiledHostUsage
}

// GetProfiledHostUsageOk returns a tuple with the ProfiledHostUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetProfiledHostUsageOk() (*float64, bool) {
	if o == nil || o.ProfiledHostUsage == nil {
		return nil, false
	}
	return o.ProfiledHostUsage, true
}

// HasProfiledHostUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasProfiledHostUsage() bool {
	return o != nil && o.ProfiledHostUsage != nil
}

// SetProfiledHostUsage gets a reference to the given float64 and assigns it to the ProfiledHostUsage field.
func (o *MonthlyUsageAttributionValues) SetProfiledHostUsage(v float64) {
	o.ProfiledHostUsage = &v
}

// GetSdsScannedBytesPercentage returns the SdsScannedBytesPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetSdsScannedBytesPercentage() float64 {
	if o == nil || o.SdsScannedBytesPercentage == nil {
		var ret float64
		return ret
	}
	return *o.SdsScannedBytesPercentage
}

// GetSdsScannedBytesPercentageOk returns a tuple with the SdsScannedBytesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetSdsScannedBytesPercentageOk() (*float64, bool) {
	if o == nil || o.SdsScannedBytesPercentage == nil {
		return nil, false
	}
	return o.SdsScannedBytesPercentage, true
}

// HasSdsScannedBytesPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasSdsScannedBytesPercentage() bool {
	return o != nil && o.SdsScannedBytesPercentage != nil
}

// SetSdsScannedBytesPercentage gets a reference to the given float64 and assigns it to the SdsScannedBytesPercentage field.
func (o *MonthlyUsageAttributionValues) SetSdsScannedBytesPercentage(v float64) {
	o.SdsScannedBytesPercentage = &v
}

// GetSdsScannedBytesUsage returns the SdsScannedBytesUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetSdsScannedBytesUsage() float64 {
	if o == nil || o.SdsScannedBytesUsage == nil {
		var ret float64
		return ret
	}
	return *o.SdsScannedBytesUsage
}

// GetSdsScannedBytesUsageOk returns a tuple with the SdsScannedBytesUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetSdsScannedBytesUsageOk() (*float64, bool) {
	if o == nil || o.SdsScannedBytesUsage == nil {
		return nil, false
	}
	return o.SdsScannedBytesUsage, true
}

// HasSdsScannedBytesUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasSdsScannedBytesUsage() bool {
	return o != nil && o.SdsScannedBytesUsage != nil
}

// SetSdsScannedBytesUsage gets a reference to the given float64 and assigns it to the SdsScannedBytesUsage field.
func (o *MonthlyUsageAttributionValues) SetSdsScannedBytesUsage(v float64) {
	o.SdsScannedBytesUsage = &v
}

// GetServerlessAppsPercentage returns the ServerlessAppsPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetServerlessAppsPercentage() float64 {
	if o == nil || o.ServerlessAppsPercentage == nil {
		var ret float64
		return ret
	}
	return *o.ServerlessAppsPercentage
}

// GetServerlessAppsPercentageOk returns a tuple with the ServerlessAppsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetServerlessAppsPercentageOk() (*float64, bool) {
	if o == nil || o.ServerlessAppsPercentage == nil {
		return nil, false
	}
	return o.ServerlessAppsPercentage, true
}

// HasServerlessAppsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasServerlessAppsPercentage() bool {
	return o != nil && o.ServerlessAppsPercentage != nil
}

// SetServerlessAppsPercentage gets a reference to the given float64 and assigns it to the ServerlessAppsPercentage field.
func (o *MonthlyUsageAttributionValues) SetServerlessAppsPercentage(v float64) {
	o.ServerlessAppsPercentage = &v
}

// GetServerlessAppsUsage returns the ServerlessAppsUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetServerlessAppsUsage() float64 {
	if o == nil || o.ServerlessAppsUsage == nil {
		var ret float64
		return ret
	}
	return *o.ServerlessAppsUsage
}

// GetServerlessAppsUsageOk returns a tuple with the ServerlessAppsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetServerlessAppsUsageOk() (*float64, bool) {
	if o == nil || o.ServerlessAppsUsage == nil {
		return nil, false
	}
	return o.ServerlessAppsUsage, true
}

// HasServerlessAppsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasServerlessAppsUsage() bool {
	return o != nil && o.ServerlessAppsUsage != nil
}

// SetServerlessAppsUsage gets a reference to the given float64 and assigns it to the ServerlessAppsUsage field.
func (o *MonthlyUsageAttributionValues) SetServerlessAppsUsage(v float64) {
	o.ServerlessAppsUsage = &v
}

// GetSnmpPercentage returns the SnmpPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetSnmpPercentage() float64 {
	if o == nil || o.SnmpPercentage == nil {
		var ret float64
		return ret
	}
	return *o.SnmpPercentage
}

// GetSnmpPercentageOk returns a tuple with the SnmpPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetSnmpPercentageOk() (*float64, bool) {
	if o == nil || o.SnmpPercentage == nil {
		return nil, false
	}
	return o.SnmpPercentage, true
}

// HasSnmpPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasSnmpPercentage() bool {
	return o != nil && o.SnmpPercentage != nil
}

// SetSnmpPercentage gets a reference to the given float64 and assigns it to the SnmpPercentage field.
func (o *MonthlyUsageAttributionValues) SetSnmpPercentage(v float64) {
	o.SnmpPercentage = &v
}

// GetSnmpUsage returns the SnmpUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetSnmpUsage() float64 {
	if o == nil || o.SnmpUsage == nil {
		var ret float64
		return ret
	}
	return *o.SnmpUsage
}

// GetSnmpUsageOk returns a tuple with the SnmpUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetSnmpUsageOk() (*float64, bool) {
	if o == nil || o.SnmpUsage == nil {
		return nil, false
	}
	return o.SnmpUsage, true
}

// HasSnmpUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasSnmpUsage() bool {
	return o != nil && o.SnmpUsage != nil
}

// SetSnmpUsage gets a reference to the given float64 and assigns it to the SnmpUsage field.
func (o *MonthlyUsageAttributionValues) SetSnmpUsage(v float64) {
	o.SnmpUsage = &v
}

// GetUniversalServiceMonitoringPercentage returns the UniversalServiceMonitoringPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetUniversalServiceMonitoringPercentage() float64 {
	if o == nil || o.UniversalServiceMonitoringPercentage == nil {
		var ret float64
		return ret
	}
	return *o.UniversalServiceMonitoringPercentage
}

// GetUniversalServiceMonitoringPercentageOk returns a tuple with the UniversalServiceMonitoringPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetUniversalServiceMonitoringPercentageOk() (*float64, bool) {
	if o == nil || o.UniversalServiceMonitoringPercentage == nil {
		return nil, false
	}
	return o.UniversalServiceMonitoringPercentage, true
}

// HasUniversalServiceMonitoringPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasUniversalServiceMonitoringPercentage() bool {
	return o != nil && o.UniversalServiceMonitoringPercentage != nil
}

// SetUniversalServiceMonitoringPercentage gets a reference to the given float64 and assigns it to the UniversalServiceMonitoringPercentage field.
func (o *MonthlyUsageAttributionValues) SetUniversalServiceMonitoringPercentage(v float64) {
	o.UniversalServiceMonitoringPercentage = &v
}

// GetUniversalServiceMonitoringUsage returns the UniversalServiceMonitoringUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetUniversalServiceMonitoringUsage() float64 {
	if o == nil || o.UniversalServiceMonitoringUsage == nil {
		var ret float64
		return ret
	}
	return *o.UniversalServiceMonitoringUsage
}

// GetUniversalServiceMonitoringUsageOk returns a tuple with the UniversalServiceMonitoringUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetUniversalServiceMonitoringUsageOk() (*float64, bool) {
	if o == nil || o.UniversalServiceMonitoringUsage == nil {
		return nil, false
	}
	return o.UniversalServiceMonitoringUsage, true
}

// HasUniversalServiceMonitoringUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasUniversalServiceMonitoringUsage() bool {
	return o != nil && o.UniversalServiceMonitoringUsage != nil
}

// SetUniversalServiceMonitoringUsage gets a reference to the given float64 and assigns it to the UniversalServiceMonitoringUsage field.
func (o *MonthlyUsageAttributionValues) SetUniversalServiceMonitoringUsage(v float64) {
	o.UniversalServiceMonitoringUsage = &v
}

// GetVulnManagementHostsPercentage returns the VulnManagementHostsPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetVulnManagementHostsPercentage() float64 {
	if o == nil || o.VulnManagementHostsPercentage == nil {
		var ret float64
		return ret
	}
	return *o.VulnManagementHostsPercentage
}

// GetVulnManagementHostsPercentageOk returns a tuple with the VulnManagementHostsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetVulnManagementHostsPercentageOk() (*float64, bool) {
	if o == nil || o.VulnManagementHostsPercentage == nil {
		return nil, false
	}
	return o.VulnManagementHostsPercentage, true
}

// HasVulnManagementHostsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasVulnManagementHostsPercentage() bool {
	return o != nil && o.VulnManagementHostsPercentage != nil
}

// SetVulnManagementHostsPercentage gets a reference to the given float64 and assigns it to the VulnManagementHostsPercentage field.
func (o *MonthlyUsageAttributionValues) SetVulnManagementHostsPercentage(v float64) {
	o.VulnManagementHostsPercentage = &v
}

// GetVulnManagementHostsUsage returns the VulnManagementHostsUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetVulnManagementHostsUsage() float64 {
	if o == nil || o.VulnManagementHostsUsage == nil {
		var ret float64
		return ret
	}
	return *o.VulnManagementHostsUsage
}

// GetVulnManagementHostsUsageOk returns a tuple with the VulnManagementHostsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetVulnManagementHostsUsageOk() (*float64, bool) {
	if o == nil || o.VulnManagementHostsUsage == nil {
		return nil, false
	}
	return o.VulnManagementHostsUsage, true
}

// HasVulnManagementHostsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasVulnManagementHostsUsage() bool {
	return o != nil && o.VulnManagementHostsUsage != nil
}

// SetVulnManagementHostsUsage gets a reference to the given float64 and assigns it to the VulnManagementHostsUsage field.
func (o *MonthlyUsageAttributionValues) SetVulnManagementHostsUsage(v float64) {
	o.VulnManagementHostsUsage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonthlyUsageAttributionValues) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ApiPercentage != nil {
		toSerialize["api_percentage"] = o.ApiPercentage
	}
	if o.ApiUsage != nil {
		toSerialize["api_usage"] = o.ApiUsage
	}
	if o.ApmFargatePercentage != nil {
		toSerialize["apm_fargate_percentage"] = o.ApmFargatePercentage
	}
	if o.ApmFargateUsage != nil {
		toSerialize["apm_fargate_usage"] = o.ApmFargateUsage
	}
	if o.ApmHostPercentage != nil {
		toSerialize["apm_host_percentage"] = o.ApmHostPercentage
	}
	if o.ApmHostUsage != nil {
		toSerialize["apm_host_usage"] = o.ApmHostUsage
	}
	if o.ApmUsmPercentage != nil {
		toSerialize["apm_usm_percentage"] = o.ApmUsmPercentage
	}
	if o.ApmUsmUsage != nil {
		toSerialize["apm_usm_usage"] = o.ApmUsmUsage
	}
	if o.AppsecFargatePercentage != nil {
		toSerialize["appsec_fargate_percentage"] = o.AppsecFargatePercentage
	}
	if o.AppsecFargateUsage != nil {
		toSerialize["appsec_fargate_usage"] = o.AppsecFargateUsage
	}
	if o.AppsecPercentage != nil {
		toSerialize["appsec_percentage"] = o.AppsecPercentage
	}
	if o.AppsecUsage != nil {
		toSerialize["appsec_usage"] = o.AppsecUsage
	}
	if o.BrowserPercentage != nil {
		toSerialize["browser_percentage"] = o.BrowserPercentage
	}
	if o.BrowserUsage != nil {
		toSerialize["browser_usage"] = o.BrowserUsage
	}
	if o.CiVisibilityItrPercentage != nil {
		toSerialize["ci_visibility_itr_percentage"] = o.CiVisibilityItrPercentage
	}
	if o.CiVisibilityItrUsage != nil {
		toSerialize["ci_visibility_itr_usage"] = o.CiVisibilityItrUsage
	}
	if o.ContainerExclAgentPercentage != nil {
		toSerialize["container_excl_agent_percentage"] = o.ContainerExclAgentPercentage
	}
	if o.ContainerExclAgentUsage != nil {
		toSerialize["container_excl_agent_usage"] = o.ContainerExclAgentUsage
	}
	if o.ContainerPercentage != nil {
		toSerialize["container_percentage"] = o.ContainerPercentage
	}
	if o.ContainerUsage != nil {
		toSerialize["container_usage"] = o.ContainerUsage
	}
	if o.CspmContainersPercentage != nil {
		toSerialize["cspm_containers_percentage"] = o.CspmContainersPercentage
	}
	if o.CspmContainersUsage != nil {
		toSerialize["cspm_containers_usage"] = o.CspmContainersUsage
	}
	if o.CspmHostsPercentage != nil {
		toSerialize["cspm_hosts_percentage"] = o.CspmHostsPercentage
	}
	if o.CspmHostsUsage != nil {
		toSerialize["cspm_hosts_usage"] = o.CspmHostsUsage
	}
	if o.CustomIngestedTimeseriesPercentage != nil {
		toSerialize["custom_ingested_timeseries_percentage"] = o.CustomIngestedTimeseriesPercentage
	}
	if o.CustomIngestedTimeseriesUsage != nil {
		toSerialize["custom_ingested_timeseries_usage"] = o.CustomIngestedTimeseriesUsage
	}
	if o.CustomTimeseriesPercentage != nil {
		toSerialize["custom_timeseries_percentage"] = o.CustomTimeseriesPercentage
	}
	if o.CustomTimeseriesUsage != nil {
		toSerialize["custom_timeseries_usage"] = o.CustomTimeseriesUsage
	}
	if o.CwsContainersPercentage != nil {
		toSerialize["cws_containers_percentage"] = o.CwsContainersPercentage
	}
	if o.CwsContainersUsage != nil {
		toSerialize["cws_containers_usage"] = o.CwsContainersUsage
	}
	if o.CwsHostsPercentage != nil {
		toSerialize["cws_hosts_percentage"] = o.CwsHostsPercentage
	}
	if o.CwsHostsUsage != nil {
		toSerialize["cws_hosts_usage"] = o.CwsHostsUsage
	}
	if o.DbmHostsPercentage != nil {
		toSerialize["dbm_hosts_percentage"] = o.DbmHostsPercentage
	}
	if o.DbmHostsUsage != nil {
		toSerialize["dbm_hosts_usage"] = o.DbmHostsUsage
	}
	if o.DbmQueriesPercentage != nil {
		toSerialize["dbm_queries_percentage"] = o.DbmQueriesPercentage
	}
	if o.DbmQueriesUsage != nil {
		toSerialize["dbm_queries_usage"] = o.DbmQueriesUsage
	}
	if o.EstimatedIndexedLogsPercentage != nil {
		toSerialize["estimated_indexed_logs_percentage"] = o.EstimatedIndexedLogsPercentage
	}
	if o.EstimatedIndexedLogsUsage != nil {
		toSerialize["estimated_indexed_logs_usage"] = o.EstimatedIndexedLogsUsage
	}
	if o.EstimatedIndexedSpansPercentage != nil {
		toSerialize["estimated_indexed_spans_percentage"] = o.EstimatedIndexedSpansPercentage
	}
	if o.EstimatedIndexedSpansUsage != nil {
		toSerialize["estimated_indexed_spans_usage"] = o.EstimatedIndexedSpansUsage
	}
	if o.EstimatedIngestedLogsPercentage != nil {
		toSerialize["estimated_ingested_logs_percentage"] = o.EstimatedIngestedLogsPercentage
	}
	if o.EstimatedIngestedLogsUsage != nil {
		toSerialize["estimated_ingested_logs_usage"] = o.EstimatedIngestedLogsUsage
	}
	if o.EstimatedIngestedSpansPercentage != nil {
		toSerialize["estimated_ingested_spans_percentage"] = o.EstimatedIngestedSpansPercentage
	}
	if o.EstimatedIngestedSpansUsage != nil {
		toSerialize["estimated_ingested_spans_usage"] = o.EstimatedIngestedSpansUsage
	}
	if o.EstimatedRumSessionsPercentage != nil {
		toSerialize["estimated_rum_sessions_percentage"] = o.EstimatedRumSessionsPercentage
	}
	if o.EstimatedRumSessionsUsage != nil {
		toSerialize["estimated_rum_sessions_usage"] = o.EstimatedRumSessionsUsage
	}
	if o.FargatePercentage != nil {
		toSerialize["fargate_percentage"] = o.FargatePercentage
	}
	if o.FargateUsage != nil {
		toSerialize["fargate_usage"] = o.FargateUsage
	}
	if o.FunctionsPercentage != nil {
		toSerialize["functions_percentage"] = o.FunctionsPercentage
	}
	if o.FunctionsUsage != nil {
		toSerialize["functions_usage"] = o.FunctionsUsage
	}
	if o.InfraHostPercentage != nil {
		toSerialize["infra_host_percentage"] = o.InfraHostPercentage
	}
	if o.InfraHostUsage != nil {
		toSerialize["infra_host_usage"] = o.InfraHostUsage
	}
	if o.InvocationsPercentage != nil {
		toSerialize["invocations_percentage"] = o.InvocationsPercentage
	}
	if o.InvocationsUsage != nil {
		toSerialize["invocations_usage"] = o.InvocationsUsage
	}
	if o.MobileAppTestingPercentage != nil {
		toSerialize["mobile_app_testing_percentage"] = o.MobileAppTestingPercentage
	}
	if o.MobileAppTestingUsage != nil {
		toSerialize["mobile_app_testing_usage"] = o.MobileAppTestingUsage
	}
	if o.NdmNetflowPercentage != nil {
		toSerialize["ndm_netflow_percentage"] = o.NdmNetflowPercentage
	}
	if o.NdmNetflowUsage != nil {
		toSerialize["ndm_netflow_usage"] = o.NdmNetflowUsage
	}
	if o.NpmHostPercentage != nil {
		toSerialize["npm_host_percentage"] = o.NpmHostPercentage
	}
	if o.NpmHostUsage != nil {
		toSerialize["npm_host_usage"] = o.NpmHostUsage
	}
	if o.ObsPipelineBytesPercentage != nil {
		toSerialize["obs_pipeline_bytes_percentage"] = o.ObsPipelineBytesPercentage
	}
	if o.ObsPipelineBytesUsage != nil {
		toSerialize["obs_pipeline_bytes_usage"] = o.ObsPipelineBytesUsage
	}
	if o.ProfiledContainerPercentage != nil {
		toSerialize["profiled_container_percentage"] = o.ProfiledContainerPercentage
	}
	if o.ProfiledContainerUsage != nil {
		toSerialize["profiled_container_usage"] = o.ProfiledContainerUsage
	}
	if o.ProfiledFargatePercentage != nil {
		toSerialize["profiled_fargate_percentage"] = o.ProfiledFargatePercentage
	}
	if o.ProfiledFargateUsage != nil {
		toSerialize["profiled_fargate_usage"] = o.ProfiledFargateUsage
	}
	if o.ProfiledHostPercentage != nil {
		toSerialize["profiled_host_percentage"] = o.ProfiledHostPercentage
	}
	if o.ProfiledHostUsage != nil {
		toSerialize["profiled_host_usage"] = o.ProfiledHostUsage
	}
	if o.SdsScannedBytesPercentage != nil {
		toSerialize["sds_scanned_bytes_percentage"] = o.SdsScannedBytesPercentage
	}
	if o.SdsScannedBytesUsage != nil {
		toSerialize["sds_scanned_bytes_usage"] = o.SdsScannedBytesUsage
	}
	if o.ServerlessAppsPercentage != nil {
		toSerialize["serverless_apps_percentage"] = o.ServerlessAppsPercentage
	}
	if o.ServerlessAppsUsage != nil {
		toSerialize["serverless_apps_usage"] = o.ServerlessAppsUsage
	}
	if o.SnmpPercentage != nil {
		toSerialize["snmp_percentage"] = o.SnmpPercentage
	}
	if o.SnmpUsage != nil {
		toSerialize["snmp_usage"] = o.SnmpUsage
	}
	if o.UniversalServiceMonitoringPercentage != nil {
		toSerialize["universal_service_monitoring_percentage"] = o.UniversalServiceMonitoringPercentage
	}
	if o.UniversalServiceMonitoringUsage != nil {
		toSerialize["universal_service_monitoring_usage"] = o.UniversalServiceMonitoringUsage
	}
	if o.VulnManagementHostsPercentage != nil {
		toSerialize["vuln_management_hosts_percentage"] = o.VulnManagementHostsPercentage
	}
	if o.VulnManagementHostsUsage != nil {
		toSerialize["vuln_management_hosts_usage"] = o.VulnManagementHostsUsage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonthlyUsageAttributionValues) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiPercentage                        *float64 `json:"api_percentage,omitempty"`
		ApiUsage                             *float64 `json:"api_usage,omitempty"`
		ApmFargatePercentage                 *float64 `json:"apm_fargate_percentage,omitempty"`
		ApmFargateUsage                      *float64 `json:"apm_fargate_usage,omitempty"`
		ApmHostPercentage                    *float64 `json:"apm_host_percentage,omitempty"`
		ApmHostUsage                         *float64 `json:"apm_host_usage,omitempty"`
		ApmUsmPercentage                     *float64 `json:"apm_usm_percentage,omitempty"`
		ApmUsmUsage                          *float64 `json:"apm_usm_usage,omitempty"`
		AppsecFargatePercentage              *float64 `json:"appsec_fargate_percentage,omitempty"`
		AppsecFargateUsage                   *float64 `json:"appsec_fargate_usage,omitempty"`
		AppsecPercentage                     *float64 `json:"appsec_percentage,omitempty"`
		AppsecUsage                          *float64 `json:"appsec_usage,omitempty"`
		BrowserPercentage                    *float64 `json:"browser_percentage,omitempty"`
		BrowserUsage                         *float64 `json:"browser_usage,omitempty"`
		CiVisibilityItrPercentage            *float64 `json:"ci_visibility_itr_percentage,omitempty"`
		CiVisibilityItrUsage                 *float64 `json:"ci_visibility_itr_usage,omitempty"`
		ContainerExclAgentPercentage         *float64 `json:"container_excl_agent_percentage,omitempty"`
		ContainerExclAgentUsage              *float64 `json:"container_excl_agent_usage,omitempty"`
		ContainerPercentage                  *float64 `json:"container_percentage,omitempty"`
		ContainerUsage                       *float64 `json:"container_usage,omitempty"`
		CspmContainersPercentage             *float64 `json:"cspm_containers_percentage,omitempty"`
		CspmContainersUsage                  *float64 `json:"cspm_containers_usage,omitempty"`
		CspmHostsPercentage                  *float64 `json:"cspm_hosts_percentage,omitempty"`
		CspmHostsUsage                       *float64 `json:"cspm_hosts_usage,omitempty"`
		CustomIngestedTimeseriesPercentage   *float64 `json:"custom_ingested_timeseries_percentage,omitempty"`
		CustomIngestedTimeseriesUsage        *float64 `json:"custom_ingested_timeseries_usage,omitempty"`
		CustomTimeseriesPercentage           *float64 `json:"custom_timeseries_percentage,omitempty"`
		CustomTimeseriesUsage                *float64 `json:"custom_timeseries_usage,omitempty"`
		CwsContainersPercentage              *float64 `json:"cws_containers_percentage,omitempty"`
		CwsContainersUsage                   *float64 `json:"cws_containers_usage,omitempty"`
		CwsHostsPercentage                   *float64 `json:"cws_hosts_percentage,omitempty"`
		CwsHostsUsage                        *float64 `json:"cws_hosts_usage,omitempty"`
		DbmHostsPercentage                   *float64 `json:"dbm_hosts_percentage,omitempty"`
		DbmHostsUsage                        *float64 `json:"dbm_hosts_usage,omitempty"`
		DbmQueriesPercentage                 *float64 `json:"dbm_queries_percentage,omitempty"`
		DbmQueriesUsage                      *float64 `json:"dbm_queries_usage,omitempty"`
		EstimatedIndexedLogsPercentage       *float64 `json:"estimated_indexed_logs_percentage,omitempty"`
		EstimatedIndexedLogsUsage            *float64 `json:"estimated_indexed_logs_usage,omitempty"`
		EstimatedIndexedSpansPercentage      *float64 `json:"estimated_indexed_spans_percentage,omitempty"`
		EstimatedIndexedSpansUsage           *float64 `json:"estimated_indexed_spans_usage,omitempty"`
		EstimatedIngestedLogsPercentage      *float64 `json:"estimated_ingested_logs_percentage,omitempty"`
		EstimatedIngestedLogsUsage           *float64 `json:"estimated_ingested_logs_usage,omitempty"`
		EstimatedIngestedSpansPercentage     *float64 `json:"estimated_ingested_spans_percentage,omitempty"`
		EstimatedIngestedSpansUsage          *float64 `json:"estimated_ingested_spans_usage,omitempty"`
		EstimatedRumSessionsPercentage       *float64 `json:"estimated_rum_sessions_percentage,omitempty"`
		EstimatedRumSessionsUsage            *float64 `json:"estimated_rum_sessions_usage,omitempty"`
		FargatePercentage                    *float64 `json:"fargate_percentage,omitempty"`
		FargateUsage                         *float64 `json:"fargate_usage,omitempty"`
		FunctionsPercentage                  *float64 `json:"functions_percentage,omitempty"`
		FunctionsUsage                       *float64 `json:"functions_usage,omitempty"`
		InfraHostPercentage                  *float64 `json:"infra_host_percentage,omitempty"`
		InfraHostUsage                       *float64 `json:"infra_host_usage,omitempty"`
		InvocationsPercentage                *float64 `json:"invocations_percentage,omitempty"`
		InvocationsUsage                     *float64 `json:"invocations_usage,omitempty"`
		MobileAppTestingPercentage           *float64 `json:"mobile_app_testing_percentage,omitempty"`
		MobileAppTestingUsage                *float64 `json:"mobile_app_testing_usage,omitempty"`
		NdmNetflowPercentage                 *float64 `json:"ndm_netflow_percentage,omitempty"`
		NdmNetflowUsage                      *float64 `json:"ndm_netflow_usage,omitempty"`
		NpmHostPercentage                    *float64 `json:"npm_host_percentage,omitempty"`
		NpmHostUsage                         *float64 `json:"npm_host_usage,omitempty"`
		ObsPipelineBytesPercentage           *float64 `json:"obs_pipeline_bytes_percentage,omitempty"`
		ObsPipelineBytesUsage                *float64 `json:"obs_pipeline_bytes_usage,omitempty"`
		ProfiledContainerPercentage          *float64 `json:"profiled_container_percentage,omitempty"`
		ProfiledContainerUsage               *float64 `json:"profiled_container_usage,omitempty"`
		ProfiledFargatePercentage            *float64 `json:"profiled_fargate_percentage,omitempty"`
		ProfiledFargateUsage                 *float64 `json:"profiled_fargate_usage,omitempty"`
		ProfiledHostPercentage               *float64 `json:"profiled_host_percentage,omitempty"`
		ProfiledHostUsage                    *float64 `json:"profiled_host_usage,omitempty"`
		SdsScannedBytesPercentage            *float64 `json:"sds_scanned_bytes_percentage,omitempty"`
		SdsScannedBytesUsage                 *float64 `json:"sds_scanned_bytes_usage,omitempty"`
		ServerlessAppsPercentage             *float64 `json:"serverless_apps_percentage,omitempty"`
		ServerlessAppsUsage                  *float64 `json:"serverless_apps_usage,omitempty"`
		SnmpPercentage                       *float64 `json:"snmp_percentage,omitempty"`
		SnmpUsage                            *float64 `json:"snmp_usage,omitempty"`
		UniversalServiceMonitoringPercentage *float64 `json:"universal_service_monitoring_percentage,omitempty"`
		UniversalServiceMonitoringUsage      *float64 `json:"universal_service_monitoring_usage,omitempty"`
		VulnManagementHostsPercentage        *float64 `json:"vuln_management_hosts_percentage,omitempty"`
		VulnManagementHostsUsage             *float64 `json:"vuln_management_hosts_usage,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"api_percentage", "api_usage", "apm_fargate_percentage", "apm_fargate_usage", "apm_host_percentage", "apm_host_usage", "apm_usm_percentage", "apm_usm_usage", "appsec_fargate_percentage", "appsec_fargate_usage", "appsec_percentage", "appsec_usage", "browser_percentage", "browser_usage", "ci_visibility_itr_percentage", "ci_visibility_itr_usage", "container_excl_agent_percentage", "container_excl_agent_usage", "container_percentage", "container_usage", "cspm_containers_percentage", "cspm_containers_usage", "cspm_hosts_percentage", "cspm_hosts_usage", "custom_ingested_timeseries_percentage", "custom_ingested_timeseries_usage", "custom_timeseries_percentage", "custom_timeseries_usage", "cws_containers_percentage", "cws_containers_usage", "cws_hosts_percentage", "cws_hosts_usage", "dbm_hosts_percentage", "dbm_hosts_usage", "dbm_queries_percentage", "dbm_queries_usage", "estimated_indexed_logs_percentage", "estimated_indexed_logs_usage", "estimated_indexed_spans_percentage", "estimated_indexed_spans_usage", "estimated_ingested_logs_percentage", "estimated_ingested_logs_usage", "estimated_ingested_spans_percentage", "estimated_ingested_spans_usage", "estimated_rum_sessions_percentage", "estimated_rum_sessions_usage", "fargate_percentage", "fargate_usage", "functions_percentage", "functions_usage", "infra_host_percentage", "infra_host_usage", "invocations_percentage", "invocations_usage", "mobile_app_testing_percentage", "mobile_app_testing_usage", "ndm_netflow_percentage", "ndm_netflow_usage", "npm_host_percentage", "npm_host_usage", "obs_pipeline_bytes_percentage", "obs_pipeline_bytes_usage", "profiled_container_percentage", "profiled_container_usage", "profiled_fargate_percentage", "profiled_fargate_usage", "profiled_host_percentage", "profiled_host_usage", "sds_scanned_bytes_percentage", "sds_scanned_bytes_usage", "serverless_apps_percentage", "serverless_apps_usage", "snmp_percentage", "snmp_usage", "universal_service_monitoring_percentage", "universal_service_monitoring_usage", "vuln_management_hosts_percentage", "vuln_management_hosts_usage"})
	} else {
		return err
	}
	o.ApiPercentage = all.ApiPercentage
	o.ApiUsage = all.ApiUsage
	o.ApmFargatePercentage = all.ApmFargatePercentage
	o.ApmFargateUsage = all.ApmFargateUsage
	o.ApmHostPercentage = all.ApmHostPercentage
	o.ApmHostUsage = all.ApmHostUsage
	o.ApmUsmPercentage = all.ApmUsmPercentage
	o.ApmUsmUsage = all.ApmUsmUsage
	o.AppsecFargatePercentage = all.AppsecFargatePercentage
	o.AppsecFargateUsage = all.AppsecFargateUsage
	o.AppsecPercentage = all.AppsecPercentage
	o.AppsecUsage = all.AppsecUsage
	o.BrowserPercentage = all.BrowserPercentage
	o.BrowserUsage = all.BrowserUsage
	o.CiVisibilityItrPercentage = all.CiVisibilityItrPercentage
	o.CiVisibilityItrUsage = all.CiVisibilityItrUsage
	o.ContainerExclAgentPercentage = all.ContainerExclAgentPercentage
	o.ContainerExclAgentUsage = all.ContainerExclAgentUsage
	o.ContainerPercentage = all.ContainerPercentage
	o.ContainerUsage = all.ContainerUsage
	o.CspmContainersPercentage = all.CspmContainersPercentage
	o.CspmContainersUsage = all.CspmContainersUsage
	o.CspmHostsPercentage = all.CspmHostsPercentage
	o.CspmHostsUsage = all.CspmHostsUsage
	o.CustomIngestedTimeseriesPercentage = all.CustomIngestedTimeseriesPercentage
	o.CustomIngestedTimeseriesUsage = all.CustomIngestedTimeseriesUsage
	o.CustomTimeseriesPercentage = all.CustomTimeseriesPercentage
	o.CustomTimeseriesUsage = all.CustomTimeseriesUsage
	o.CwsContainersPercentage = all.CwsContainersPercentage
	o.CwsContainersUsage = all.CwsContainersUsage
	o.CwsHostsPercentage = all.CwsHostsPercentage
	o.CwsHostsUsage = all.CwsHostsUsage
	o.DbmHostsPercentage = all.DbmHostsPercentage
	o.DbmHostsUsage = all.DbmHostsUsage
	o.DbmQueriesPercentage = all.DbmQueriesPercentage
	o.DbmQueriesUsage = all.DbmQueriesUsage
	o.EstimatedIndexedLogsPercentage = all.EstimatedIndexedLogsPercentage
	o.EstimatedIndexedLogsUsage = all.EstimatedIndexedLogsUsage
	o.EstimatedIndexedSpansPercentage = all.EstimatedIndexedSpansPercentage
	o.EstimatedIndexedSpansUsage = all.EstimatedIndexedSpansUsage
	o.EstimatedIngestedLogsPercentage = all.EstimatedIngestedLogsPercentage
	o.EstimatedIngestedLogsUsage = all.EstimatedIngestedLogsUsage
	o.EstimatedIngestedSpansPercentage = all.EstimatedIngestedSpansPercentage
	o.EstimatedIngestedSpansUsage = all.EstimatedIngestedSpansUsage
	o.EstimatedRumSessionsPercentage = all.EstimatedRumSessionsPercentage
	o.EstimatedRumSessionsUsage = all.EstimatedRumSessionsUsage
	o.FargatePercentage = all.FargatePercentage
	o.FargateUsage = all.FargateUsage
	o.FunctionsPercentage = all.FunctionsPercentage
	o.FunctionsUsage = all.FunctionsUsage
	o.InfraHostPercentage = all.InfraHostPercentage
	o.InfraHostUsage = all.InfraHostUsage
	o.InvocationsPercentage = all.InvocationsPercentage
	o.InvocationsUsage = all.InvocationsUsage
	o.MobileAppTestingPercentage = all.MobileAppTestingPercentage
	o.MobileAppTestingUsage = all.MobileAppTestingUsage
	o.NdmNetflowPercentage = all.NdmNetflowPercentage
	o.NdmNetflowUsage = all.NdmNetflowUsage
	o.NpmHostPercentage = all.NpmHostPercentage
	o.NpmHostUsage = all.NpmHostUsage
	o.ObsPipelineBytesPercentage = all.ObsPipelineBytesPercentage
	o.ObsPipelineBytesUsage = all.ObsPipelineBytesUsage
	o.ProfiledContainerPercentage = all.ProfiledContainerPercentage
	o.ProfiledContainerUsage = all.ProfiledContainerUsage
	o.ProfiledFargatePercentage = all.ProfiledFargatePercentage
	o.ProfiledFargateUsage = all.ProfiledFargateUsage
	o.ProfiledHostPercentage = all.ProfiledHostPercentage
	o.ProfiledHostUsage = all.ProfiledHostUsage
	o.SdsScannedBytesPercentage = all.SdsScannedBytesPercentage
	o.SdsScannedBytesUsage = all.SdsScannedBytesUsage
	o.ServerlessAppsPercentage = all.ServerlessAppsPercentage
	o.ServerlessAppsUsage = all.ServerlessAppsUsage
	o.SnmpPercentage = all.SnmpPercentage
	o.SnmpUsage = all.SnmpUsage
	o.UniversalServiceMonitoringPercentage = all.UniversalServiceMonitoringPercentage
	o.UniversalServiceMonitoringUsage = all.UniversalServiceMonitoringUsage
	o.VulnManagementHostsPercentage = all.VulnManagementHostsPercentage
	o.VulnManagementHostsUsage = all.VulnManagementHostsUsage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
