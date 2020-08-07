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

// VersionControlInformationDto struct for VersionControlInformationDto
type VersionControlInformationDto struct {
	// The ID of the Process Group that is under version control
	GroupId string `json:"groupId,omitempty"`
	// The ID of the registry that the flow is stored in
	RegistryId string `json:"registryId,omitempty"`
	// The name of the registry that the flow is stored in
	RegistryName string `json:"registryName,omitempty"`
	// The ID of the bucket that the flow is stored in
	BucketId string `json:"bucketId,omitempty"`
	// The name of the bucket that the flow is stored in
	BucketName string `json:"bucketName,omitempty"`
	// The ID of the flow
	FlowId string `json:"flowId,omitempty"`
	// The name of the flow
	FlowName string `json:"flowName,omitempty"`
	// The description of the flow
	FlowDescription string `json:"flowDescription,omitempty"`
	// The version of the flow
	Version int32 `json:"version,omitempty"`
	// The current state of the Process Group, as it relates to the Versioned Flow
	State string `json:"state,omitempty"`
	// Explanation of why the group is in the specified state
	StateExplanation string `json:"stateExplanation,omitempty"`
}
