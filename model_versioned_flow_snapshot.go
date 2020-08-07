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

// VersionedFlowSnapshot struct for VersionedFlowSnapshot
type VersionedFlowSnapshot struct {
	SnapshotMetadata VersionedFlowSnapshotMetadata `json:"snapshotMetadata"`
	FlowContents     VersionedProcessGroup         `json:"flowContents"`
	// The information about controller services that exist outside this versioned flow, but are referenced by components within the versioned flow.
	ExternalControllerServices map[string]ExternalControllerServiceReference `json:"externalControllerServices,omitempty"`
	// The parameter contexts referenced by process groups in the flow contents. The mapping is from the name of the context to the context instance, and it is expected that any context in this map is referenced by at least one process group in this flow.
	ParameterContexts map[string]VersionedParameterContext `json:"parameterContexts,omitempty"`
	// The optional encoding version of the flow contents.
	FlowEncodingVersion string        `json:"flowEncodingVersion,omitempty"`
	Flow                VersionedFlow `json:"flow,omitempty"`
	Bucket              Bucket        `json:"bucket,omitempty"`
	Latest              bool          `json:"latest,omitempty"`
}
