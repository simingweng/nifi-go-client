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

// UpdateControllerServiceReferenceRequestEntity struct for UpdateControllerServiceReferenceRequestEntity
type UpdateControllerServiceReferenceRequestEntity struct {
	// The identifier of the Controller Service.
	Id string `json:"id,omitempty"`
	// The new state of the references for the controller service.
	State string `json:"state,omitempty"`
	// The revisions for all referencing components.
	ReferencingComponentRevisions map[string]RevisionDto `json:"referencingComponentRevisions,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`
}