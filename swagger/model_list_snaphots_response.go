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

type ListSnaphotsResponse struct {
	Items []Snapshot `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}
