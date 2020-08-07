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

// ActivateControllerServicesEntity struct for ActivateControllerServicesEntity
type ActivateControllerServicesEntity struct {
	// The id of the ProcessGroup
	Id string `json:"id,omitempty"`
	// The desired state of the descendant components
	State string `json:"state,omitempty"`
	// Optional services to schedule. If not specified, all authorized descendant controller services will be used.
	Components map[string]RevisionDto `json:"components,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`
}
