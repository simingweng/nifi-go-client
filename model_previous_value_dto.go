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

// PreviousValueDto struct for PreviousValueDto
type PreviousValueDto struct {
	// The previous value.
	PreviousValue string `json:"previousValue,omitempty"`
	// The timestamp when the value was modified.
	Timestamp string `json:"timestamp,omitempty"`
	// The user who changed the previous value.
	UserIdentity string `json:"userIdentity,omitempty"`
}
