// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/goccy/go-json"
)

// HourlyUsageAttributionUsageType Supported products for hourly usage attribution requests.
type HourlyUsageAttributionUsageType string

// List of HourlyUsageAttributionUsageType.
const (
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_API_USAGE                          HourlyUsageAttributionUsageType = "api_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_APM_FARGATE_USAGE                  HourlyUsageAttributionUsageType = "apm_fargate_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_APM_HOST_USAGE                     HourlyUsageAttributionUsageType = "apm_host_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_APM_USM_USAGE                      HourlyUsageAttributionUsageType = "apm_usm_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_APPSEC_FARGATE_USAGE               HourlyUsageAttributionUsageType = "appsec_fargate_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_APPSEC_USAGE                       HourlyUsageAttributionUsageType = "appsec_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_BROWSER_USAGE                      HourlyUsageAttributionUsageType = "browser_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CI_VISIBILITY_ITR_USAGE            HourlyUsageAttributionUsageType = "ci_visibility_itr_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CONTAINER_EXCL_AGENT_USAGE         HourlyUsageAttributionUsageType = "container_excl_agent_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CONTAINER_USAGE                    HourlyUsageAttributionUsageType = "container_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CSPM_CONTAINERS_USAGE              HourlyUsageAttributionUsageType = "cspm_containers_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CSPM_HOSTS_USAGE                   HourlyUsageAttributionUsageType = "cspm_hosts_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CUSTOM_INGESTED_TIMESERIES_USAGE   HourlyUsageAttributionUsageType = "custom_ingested_timeseries_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CUSTOM_TIMESERIES_USAGE            HourlyUsageAttributionUsageType = "custom_timeseries_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CWS_CONTAINERS_USAGE               HourlyUsageAttributionUsageType = "cws_containers_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CWS_HOSTS_USAGE                    HourlyUsageAttributionUsageType = "cws_hosts_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_DBM_HOSTS_USAGE                    HourlyUsageAttributionUsageType = "dbm_hosts_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_DBM_QUERIES_USAGE                  HourlyUsageAttributionUsageType = "dbm_queries_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ESTIMATED_INDEXED_LOGS_USAGE       HourlyUsageAttributionUsageType = "estimated_indexed_logs_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ESTIMATED_INDEXED_SPANS_USAGE      HourlyUsageAttributionUsageType = "estimated_indexed_spans_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ESTIMATED_INGESTED_LOGS_USAGE      HourlyUsageAttributionUsageType = "estimated_ingested_logs_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ESTIMATED_INGESTED_SPANS_USAGE     HourlyUsageAttributionUsageType = "estimated_ingested_spans_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ESTIMATED_RUM_SESSIONS_USAGE       HourlyUsageAttributionUsageType = "estimated_rum_sessions_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_FARGATE_USAGE                      HourlyUsageAttributionUsageType = "fargate_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_FUNCTIONS_USAGE                    HourlyUsageAttributionUsageType = "functions_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_INFRA_HOST_USAGE                   HourlyUsageAttributionUsageType = "infra_host_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_INVOCATIONS_USAGE                  HourlyUsageAttributionUsageType = "invocations_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_MOBILE_APP_TESTING_USAGE           HourlyUsageAttributionUsageType = "mobile_app_testing_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_NDM_NETFLOW_USAGE                  HourlyUsageAttributionUsageType = "ndm_netflow_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_NPM_HOST_USAGE                     HourlyUsageAttributionUsageType = "npm_host_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_OBS_PIPELINE_BYTES_USAGE           HourlyUsageAttributionUsageType = "obs_pipeline_bytes_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_PROFILED_CONTAINER_USAGE           HourlyUsageAttributionUsageType = "profiled_container_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_PROFILED_FARGATE_USAGE             HourlyUsageAttributionUsageType = "profiled_fargate_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_PROFILED_HOST_USAGE                HourlyUsageAttributionUsageType = "profiled_host_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_SDS_SCANNED_BYTES_USAGE            HourlyUsageAttributionUsageType = "sds_scanned_bytes_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_SERVERLESS_APPS_USAGE              HourlyUsageAttributionUsageType = "serverless_apps_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_SNMP_USAGE                         HourlyUsageAttributionUsageType = "snmp_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_UNIVERSAL_SERVICE_MONITORING_USAGE HourlyUsageAttributionUsageType = "universal_service_monitoring_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_VULN_MANAGEMENT_HOSTS_USAGE        HourlyUsageAttributionUsageType = "vuln_management_hosts_usage"
)

var allowedHourlyUsageAttributionUsageTypeEnumValues = []HourlyUsageAttributionUsageType{
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_API_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_APM_FARGATE_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_APM_HOST_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_APM_USM_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_APPSEC_FARGATE_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_APPSEC_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_BROWSER_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CI_VISIBILITY_ITR_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CONTAINER_EXCL_AGENT_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CONTAINER_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CSPM_CONTAINERS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CSPM_HOSTS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CUSTOM_INGESTED_TIMESERIES_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CUSTOM_TIMESERIES_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CWS_CONTAINERS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CWS_HOSTS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_DBM_HOSTS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_DBM_QUERIES_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ESTIMATED_INDEXED_LOGS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ESTIMATED_INDEXED_SPANS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ESTIMATED_INGESTED_LOGS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ESTIMATED_INGESTED_SPANS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ESTIMATED_RUM_SESSIONS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_FARGATE_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_FUNCTIONS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_INFRA_HOST_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_INVOCATIONS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_MOBILE_APP_TESTING_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_NDM_NETFLOW_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_NPM_HOST_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_OBS_PIPELINE_BYTES_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_PROFILED_CONTAINER_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_PROFILED_FARGATE_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_PROFILED_HOST_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_SDS_SCANNED_BYTES_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_SERVERLESS_APPS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_SNMP_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_UNIVERSAL_SERVICE_MONITORING_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_VULN_MANAGEMENT_HOSTS_USAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HourlyUsageAttributionUsageType) GetAllowedValues() []HourlyUsageAttributionUsageType {
	return allowedHourlyUsageAttributionUsageTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HourlyUsageAttributionUsageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HourlyUsageAttributionUsageType(value)
	return nil
}

// NewHourlyUsageAttributionUsageTypeFromValue returns a pointer to a valid HourlyUsageAttributionUsageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHourlyUsageAttributionUsageTypeFromValue(v string) (*HourlyUsageAttributionUsageType, error) {
	ev := HourlyUsageAttributionUsageType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HourlyUsageAttributionUsageType: valid values are %v", v, allowedHourlyUsageAttributionUsageTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HourlyUsageAttributionUsageType) IsValid() bool {
	for _, existing := range allowedHourlyUsageAttributionUsageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HourlyUsageAttributionUsageType value.
func (v HourlyUsageAttributionUsageType) Ptr() *HourlyUsageAttributionUsageType {
	return &v
}
