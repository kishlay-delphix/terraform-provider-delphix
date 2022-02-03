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

// Parameters to enable a VDB.
type EnableVdbParameters struct {
	// Whether to attempt a startup of the VDB after the enable.
	AttemptStart bool `json:"attempt_start,omitempty"`
}
