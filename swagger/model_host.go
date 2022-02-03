/*
 * Delphix API Gateway
 *
 * Delphix API Gateway API
 *
 * API version: 1.0
 * Contact: support@delphix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A physical/virtual server.
type Host struct {
	// The hostname or IP address of this host.
	Hostname string `json:"hostname,omitempty"`
	// The name of the OS on this host.
	OsName string `json:"os_name,omitempty"`
	// The version of the OS on this host.
	OsVersion string `json:"os_version,omitempty"`
	// The total amount of memory on this host in bytes.
	MemorySize int64 `json:"memory_size,omitempty"`
}
