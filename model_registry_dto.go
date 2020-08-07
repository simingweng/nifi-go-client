/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.11.4
 * Contact: dev@nifi.apache.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nifi

// RegistryDto struct for RegistryDto
type RegistryDto struct {
	// The registry identifier
	Id string `json:"id,omitempty"`
	// The registry name
	Name string `json:"name,omitempty"`
	// The registry description
	Description string `json:"description,omitempty"`
	// The registry URI
	Uri string `json:"uri,omitempty"`
}
