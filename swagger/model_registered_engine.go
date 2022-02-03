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

// A registered engine's connection and authentication settings.
type RegisteredEngine struct {
	Id int64 `json:"id,omitempty"`
	Name string `json:"name"`
	Hostname string `json:"hostname"`
	// Id of the primary user for this engine. The primary user must be an engine admin.
	PrimaryUser int64 `json:"primary_user,omitempty"`
	// Allow connections to the engine over HTTPs without validating the TLS certificate. Even though the connection to the engine might be performed over HTTPs, setting this property eliminates the protection against a man-in-the-middle attach for connections to this engine. Instead, consider creating a truststore with a Certificate Authority to validate the engine's certificate, and set the truststore_path propery. 
	InsecureSsl bool `json:"insecure_ssl,omitempty"`
	// Ignore validation of the name associated to the TLS certificate when connecting to the engine over HTTPs. Setting this value must only be done if the TLS certificate of the engine does not match the hostname, and the TLS configuration of the engine cannot be fixed. Setting this property reduces the protection against a man-in-the-middle attack for connections to this engine. This is ignored if insecure_ssl is set. 
	UnsafeSslHostnameCheck bool `json:"unsafe_ssl_hostname_check,omitempty"`
	// File name of a truststore which can be used to validate the TLS certificate of the engine. The truststore must be available at /etc/config/certs/<truststore_filename> 
	TruststoreFilename string `json:"truststore_filename,omitempty"`
	// Password to read the truststore. 
	TruststorePassword string `json:"truststore_password,omitempty"`
}
