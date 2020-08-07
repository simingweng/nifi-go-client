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

// BundleDto struct for BundleDto
type BundleDto struct {
	// The group of the bundle.
	Group string `json:"group,omitempty"`
	// The artifact of the bundle.
	Artifact string `json:"artifact,omitempty"`
	// The version of the bundle.
	Version string `json:"version,omitempty"`
}
