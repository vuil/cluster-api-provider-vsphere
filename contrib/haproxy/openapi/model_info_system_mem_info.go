/*
 * HAProxy Data Plane API
 *
 * API for editing and managing haproxy instances. Provides process information, configuration management, haproxy stats and logs.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.2
 * Contact: support@haproxy.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// InfoSystemMemInfo struct for InfoSystemMemInfo
type InfoSystemMemInfo struct {
	DataplaneapiMemory int32 `json:"dataplaneapi_memory,omitempty"`
	FreeMemory         int32 `json:"free_memory,omitempty"`
	TotalMemory        int32 `json:"total_memory,omitempty"`
}
