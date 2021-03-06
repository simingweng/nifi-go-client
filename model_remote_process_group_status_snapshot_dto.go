/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.13.2
 * Contact: dev@nifi.apache.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nifi

import (
	"encoding/json"
)

// RemoteProcessGroupStatusSnapshotDTO struct for RemoteProcessGroupStatusSnapshotDTO
type RemoteProcessGroupStatusSnapshotDTO struct {
	// The id of the remote process group.
	Id *string `json:"id,omitempty"`
	// The id of the parent process group the remote process group resides in.
	GroupId *string `json:"groupId,omitempty"`
	// The name of the remote process group.
	Name *string `json:"name,omitempty"`
	// The URI of the target system.
	TargetUri *string `json:"targetUri,omitempty"`
	// The transmission status of the remote process group.
	TransmissionStatus *string `json:"transmissionStatus,omitempty"`
	// The number of active threads for the remote process group.
	ActiveThreadCount *int32 `json:"activeThreadCount,omitempty"`
	// The number of FlowFiles sent to the remote process group in the last 5 minutes.
	FlowFilesSent *int32 `json:"flowFilesSent,omitempty"`
	// The size of the FlowFiles sent to the remote process group in the last 5 minutes.
	BytesSent *int64 `json:"bytesSent,omitempty"`
	// The count/size of the flowfiles sent to the remote process group in the last 5 minutes.
	Sent *string `json:"sent,omitempty"`
	// The number of FlowFiles received from the remote process group in the last 5 minutes.
	FlowFilesReceived *int32 `json:"flowFilesReceived,omitempty"`
	// The size of the FlowFiles received from the remote process group in the last 5 minutes.
	BytesReceived *int64 `json:"bytesReceived,omitempty"`
	// The count/size of the flowfiles received from the remote process group in the last 5 minutes.
	Received *string `json:"received,omitempty"`
}

// NewRemoteProcessGroupStatusSnapshotDTO instantiates a new RemoteProcessGroupStatusSnapshotDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteProcessGroupStatusSnapshotDTO() *RemoteProcessGroupStatusSnapshotDTO {
	this := RemoteProcessGroupStatusSnapshotDTO{}
	return &this
}

// NewRemoteProcessGroupStatusSnapshotDTOWithDefaults instantiates a new RemoteProcessGroupStatusSnapshotDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteProcessGroupStatusSnapshotDTOWithDefaults() *RemoteProcessGroupStatusSnapshotDTO {
	this := RemoteProcessGroupStatusSnapshotDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RemoteProcessGroupStatusSnapshotDTO) SetId(v string) {
	o.Id = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *RemoteProcessGroupStatusSnapshotDTO) SetGroupId(v string) {
	o.GroupId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RemoteProcessGroupStatusSnapshotDTO) SetName(v string) {
	o.Name = &v
}

// GetTargetUri returns the TargetUri field value if set, zero value otherwise.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetTargetUri() string {
	if o == nil || o.TargetUri == nil {
		var ret string
		return ret
	}
	return *o.TargetUri
}

// GetTargetUriOk returns a tuple with the TargetUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetTargetUriOk() (*string, bool) {
	if o == nil || o.TargetUri == nil {
		return nil, false
	}
	return o.TargetUri, true
}

// HasTargetUri returns a boolean if a field has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) HasTargetUri() bool {
	if o != nil && o.TargetUri != nil {
		return true
	}

	return false
}

// SetTargetUri gets a reference to the given string and assigns it to the TargetUri field.
func (o *RemoteProcessGroupStatusSnapshotDTO) SetTargetUri(v string) {
	o.TargetUri = &v
}

// GetTransmissionStatus returns the TransmissionStatus field value if set, zero value otherwise.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetTransmissionStatus() string {
	if o == nil || o.TransmissionStatus == nil {
		var ret string
		return ret
	}
	return *o.TransmissionStatus
}

// GetTransmissionStatusOk returns a tuple with the TransmissionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetTransmissionStatusOk() (*string, bool) {
	if o == nil || o.TransmissionStatus == nil {
		return nil, false
	}
	return o.TransmissionStatus, true
}

// HasTransmissionStatus returns a boolean if a field has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) HasTransmissionStatus() bool {
	if o != nil && o.TransmissionStatus != nil {
		return true
	}

	return false
}

// SetTransmissionStatus gets a reference to the given string and assigns it to the TransmissionStatus field.
func (o *RemoteProcessGroupStatusSnapshotDTO) SetTransmissionStatus(v string) {
	o.TransmissionStatus = &v
}

// GetActiveThreadCount returns the ActiveThreadCount field value if set, zero value otherwise.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetActiveThreadCount() int32 {
	if o == nil || o.ActiveThreadCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveThreadCount
}

// GetActiveThreadCountOk returns a tuple with the ActiveThreadCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetActiveThreadCountOk() (*int32, bool) {
	if o == nil || o.ActiveThreadCount == nil {
		return nil, false
	}
	return o.ActiveThreadCount, true
}

// HasActiveThreadCount returns a boolean if a field has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) HasActiveThreadCount() bool {
	if o != nil && o.ActiveThreadCount != nil {
		return true
	}

	return false
}

// SetActiveThreadCount gets a reference to the given int32 and assigns it to the ActiveThreadCount field.
func (o *RemoteProcessGroupStatusSnapshotDTO) SetActiveThreadCount(v int32) {
	o.ActiveThreadCount = &v
}

// GetFlowFilesSent returns the FlowFilesSent field value if set, zero value otherwise.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetFlowFilesSent() int32 {
	if o == nil || o.FlowFilesSent == nil {
		var ret int32
		return ret
	}
	return *o.FlowFilesSent
}

// GetFlowFilesSentOk returns a tuple with the FlowFilesSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetFlowFilesSentOk() (*int32, bool) {
	if o == nil || o.FlowFilesSent == nil {
		return nil, false
	}
	return o.FlowFilesSent, true
}

// HasFlowFilesSent returns a boolean if a field has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) HasFlowFilesSent() bool {
	if o != nil && o.FlowFilesSent != nil {
		return true
	}

	return false
}

// SetFlowFilesSent gets a reference to the given int32 and assigns it to the FlowFilesSent field.
func (o *RemoteProcessGroupStatusSnapshotDTO) SetFlowFilesSent(v int32) {
	o.FlowFilesSent = &v
}

// GetBytesSent returns the BytesSent field value if set, zero value otherwise.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetBytesSent() int64 {
	if o == nil || o.BytesSent == nil {
		var ret int64
		return ret
	}
	return *o.BytesSent
}

// GetBytesSentOk returns a tuple with the BytesSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetBytesSentOk() (*int64, bool) {
	if o == nil || o.BytesSent == nil {
		return nil, false
	}
	return o.BytesSent, true
}

// HasBytesSent returns a boolean if a field has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) HasBytesSent() bool {
	if o != nil && o.BytesSent != nil {
		return true
	}

	return false
}

// SetBytesSent gets a reference to the given int64 and assigns it to the BytesSent field.
func (o *RemoteProcessGroupStatusSnapshotDTO) SetBytesSent(v int64) {
	o.BytesSent = &v
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetSent() string {
	if o == nil || o.Sent == nil {
		var ret string
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetSentOk() (*string, bool) {
	if o == nil || o.Sent == nil {
		return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) HasSent() bool {
	if o != nil && o.Sent != nil {
		return true
	}

	return false
}

// SetSent gets a reference to the given string and assigns it to the Sent field.
func (o *RemoteProcessGroupStatusSnapshotDTO) SetSent(v string) {
	o.Sent = &v
}

// GetFlowFilesReceived returns the FlowFilesReceived field value if set, zero value otherwise.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetFlowFilesReceived() int32 {
	if o == nil || o.FlowFilesReceived == nil {
		var ret int32
		return ret
	}
	return *o.FlowFilesReceived
}

// GetFlowFilesReceivedOk returns a tuple with the FlowFilesReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetFlowFilesReceivedOk() (*int32, bool) {
	if o == nil || o.FlowFilesReceived == nil {
		return nil, false
	}
	return o.FlowFilesReceived, true
}

// HasFlowFilesReceived returns a boolean if a field has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) HasFlowFilesReceived() bool {
	if o != nil && o.FlowFilesReceived != nil {
		return true
	}

	return false
}

// SetFlowFilesReceived gets a reference to the given int32 and assigns it to the FlowFilesReceived field.
func (o *RemoteProcessGroupStatusSnapshotDTO) SetFlowFilesReceived(v int32) {
	o.FlowFilesReceived = &v
}

// GetBytesReceived returns the BytesReceived field value if set, zero value otherwise.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetBytesReceived() int64 {
	if o == nil || o.BytesReceived == nil {
		var ret int64
		return ret
	}
	return *o.BytesReceived
}

// GetBytesReceivedOk returns a tuple with the BytesReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetBytesReceivedOk() (*int64, bool) {
	if o == nil || o.BytesReceived == nil {
		return nil, false
	}
	return o.BytesReceived, true
}

// HasBytesReceived returns a boolean if a field has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) HasBytesReceived() bool {
	if o != nil && o.BytesReceived != nil {
		return true
	}

	return false
}

// SetBytesReceived gets a reference to the given int64 and assigns it to the BytesReceived field.
func (o *RemoteProcessGroupStatusSnapshotDTO) SetBytesReceived(v int64) {
	o.BytesReceived = &v
}

// GetReceived returns the Received field value if set, zero value otherwise.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetReceived() string {
	if o == nil || o.Received == nil {
		var ret string
		return ret
	}
	return *o.Received
}

// GetReceivedOk returns a tuple with the Received field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) GetReceivedOk() (*string, bool) {
	if o == nil || o.Received == nil {
		return nil, false
	}
	return o.Received, true
}

// HasReceived returns a boolean if a field has been set.
func (o *RemoteProcessGroupStatusSnapshotDTO) HasReceived() bool {
	if o != nil && o.Received != nil {
		return true
	}

	return false
}

// SetReceived gets a reference to the given string and assigns it to the Received field.
func (o *RemoteProcessGroupStatusSnapshotDTO) SetReceived(v string) {
	o.Received = &v
}

func (o RemoteProcessGroupStatusSnapshotDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TargetUri != nil {
		toSerialize["targetUri"] = o.TargetUri
	}
	if o.TransmissionStatus != nil {
		toSerialize["transmissionStatus"] = o.TransmissionStatus
	}
	if o.ActiveThreadCount != nil {
		toSerialize["activeThreadCount"] = o.ActiveThreadCount
	}
	if o.FlowFilesSent != nil {
		toSerialize["flowFilesSent"] = o.FlowFilesSent
	}
	if o.BytesSent != nil {
		toSerialize["bytesSent"] = o.BytesSent
	}
	if o.Sent != nil {
		toSerialize["sent"] = o.Sent
	}
	if o.FlowFilesReceived != nil {
		toSerialize["flowFilesReceived"] = o.FlowFilesReceived
	}
	if o.BytesReceived != nil {
		toSerialize["bytesReceived"] = o.BytesReceived
	}
	if o.Received != nil {
		toSerialize["received"] = o.Received
	}
	return json.Marshal(toSerialize)
}

type NullableRemoteProcessGroupStatusSnapshotDTO struct {
	value *RemoteProcessGroupStatusSnapshotDTO
	isSet bool
}

func (v NullableRemoteProcessGroupStatusSnapshotDTO) Get() *RemoteProcessGroupStatusSnapshotDTO {
	return v.value
}

func (v *NullableRemoteProcessGroupStatusSnapshotDTO) Set(val *RemoteProcessGroupStatusSnapshotDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteProcessGroupStatusSnapshotDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteProcessGroupStatusSnapshotDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteProcessGroupStatusSnapshotDTO(val *RemoteProcessGroupStatusSnapshotDTO) *NullableRemoteProcessGroupStatusSnapshotDTO {
	return &NullableRemoteProcessGroupStatusSnapshotDTO{value: val, isSet: true}
}

func (v NullableRemoteProcessGroupStatusSnapshotDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteProcessGroupStatusSnapshotDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
