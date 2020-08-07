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

// StatusHistoryDto struct for StatusHistoryDto
type StatusHistoryDto struct {
	// When the status history was generated.
	Generated string `json:"generated,omitempty"`
	// A Map of key/value pairs that describe the component that the status history belongs to
	ComponentDetails map[string]string `json:"componentDetails,omitempty"`
	// The Descriptors that provide information on each of the metrics provided in the status history
	FieldDescriptors []StatusDescriptorDto `json:"fieldDescriptors,omitempty"`
	// A list of StatusSnapshotDTO objects that provide the actual metric values for the component. If the NiFi instance is clustered, this will represent the aggregate status across all nodes. If the NiFi instance is not clustered, this will represent the status of the entire NiFi instance.
	AggregateSnapshots []StatusSnapshotDto `json:"aggregateSnapshots,omitempty"`
	// The NodeStatusSnapshotsDTO objects that provide the actual metric values for the component, for each node. If the NiFi instance is not clustered, this value will be null.
	NodeSnapshots []NodeStatusSnapshotsDto `json:"nodeSnapshots,omitempty"`
}
