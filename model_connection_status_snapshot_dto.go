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

// ConnectionStatusSnapshotDTO struct for ConnectionStatusSnapshotDTO
type ConnectionStatusSnapshotDTO struct {
	// The id of the connection.
	Id *string `json:"id,omitempty"`
	// The id of the process group the connection belongs to.
	GroupId *string `json:"groupId,omitempty"`
	// The name of the connection.
	Name *string `json:"name,omitempty"`
	// The id of the source of the connection.
	SourceId *string `json:"sourceId,omitempty"`
	// The name of the source of the connection.
	SourceName *string `json:"sourceName,omitempty"`
	// The id of the destination of the connection.
	DestinationId *string `json:"destinationId,omitempty"`
	// The name of the destination of the connection.
	DestinationName *string                                 `json:"destinationName,omitempty"`
	Predictions     *ConnectionStatusPredictionsSnapshotDTO `json:"predictions,omitempty"`
	// The number of FlowFiles that have come into the connection in the last 5 minutes.
	FlowFilesIn *int32 `json:"flowFilesIn,omitempty"`
	// The size of the FlowFiles that have come into the connection in the last 5 minutes.
	BytesIn *int64 `json:"bytesIn,omitempty"`
	// The input count/size for the connection in the last 5 minutes, pretty printed.
	Input *string `json:"input,omitempty"`
	// The number of FlowFiles that have left the connection in the last 5 minutes.
	FlowFilesOut *int32 `json:"flowFilesOut,omitempty"`
	// The number of bytes that have left the connection in the last 5 minutes.
	BytesOut *int64 `json:"bytesOut,omitempty"`
	// The output count/sie for the connection in the last 5 minutes, pretty printed.
	Output *string `json:"output,omitempty"`
	// The number of FlowFiles that are currently queued in the connection.
	FlowFilesQueued *int32 `json:"flowFilesQueued,omitempty"`
	// The size of the FlowFiles that are currently queued in the connection.
	BytesQueued *int64 `json:"bytesQueued,omitempty"`
	// The total count and size of queued flowfiles formatted.
	Queued *string `json:"queued,omitempty"`
	// The total size of flowfiles that are queued formatted.
	QueuedSize *string `json:"queuedSize,omitempty"`
	// The number of flowfiles that are queued, pretty printed.
	QueuedCount *string `json:"queuedCount,omitempty"`
	// Connection percent use regarding queued flow files count and backpressure threshold if configured.
	PercentUseCount *int32 `json:"percentUseCount,omitempty"`
	// Connection percent use regarding queued flow files size and backpressure threshold if configured.
	PercentUseBytes *int32 `json:"percentUseBytes,omitempty"`
}

// NewConnectionStatusSnapshotDTO instantiates a new ConnectionStatusSnapshotDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionStatusSnapshotDTO() *ConnectionStatusSnapshotDTO {
	this := ConnectionStatusSnapshotDTO{}
	return &this
}

// NewConnectionStatusSnapshotDTOWithDefaults instantiates a new ConnectionStatusSnapshotDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionStatusSnapshotDTOWithDefaults() *ConnectionStatusSnapshotDTO {
	this := ConnectionStatusSnapshotDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConnectionStatusSnapshotDTO) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusSnapshotDTO) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConnectionStatusSnapshotDTO) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConnectionStatusSnapshotDTO) SetId(v string) {
	o.Id = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ConnectionStatusSnapshotDTO) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusSnapshotDTO) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ConnectionStatusSnapshotDTO) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ConnectionStatusSnapshotDTO) SetGroupId(v string) {
	o.GroupId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConnectionStatusSnapshotDTO) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusSnapshotDTO) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConnectionStatusSnapshotDTO) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConnectionStatusSnapshotDTO) SetName(v string) {
	o.Name = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *ConnectionStatusSnapshotDTO) GetSourceId() string {
	if o == nil || o.SourceId == nil {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusSnapshotDTO) GetSourceIdOk() (*string, bool) {
	if o == nil || o.SourceId == nil {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *ConnectionStatusSnapshotDTO) HasSourceId() bool {
	if o != nil && o.SourceId != nil {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *ConnectionStatusSnapshotDTO) SetSourceId(v string) {
	o.SourceId = &v
}

// GetSourceName returns the SourceName field value if set, zero value otherwise.
func (o *ConnectionStatusSnapshotDTO) GetSourceName() string {
	if o == nil || o.SourceName == nil {
		var ret string
		return ret
	}
	return *o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusSnapshotDTO) GetSourceNameOk() (*string, bool) {
	if o == nil || o.SourceName == nil {
		return nil, false
	}
	return o.SourceName, true
}

// HasSourceName returns a boolean if a field has been set.
func (o *ConnectionStatusSnapshotDTO) HasSourceName() bool {
	if o != nil && o.SourceName != nil {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given string and assigns it to the SourceName field.
func (o *ConnectionStatusSnapshotDTO) SetSourceName(v string) {
	o.SourceName = &v
}

// GetDestinationId returns the DestinationId field value if set, zero value otherwise.
func (o *ConnectionStatusSnapshotDTO) GetDestinationId() string {
	if o == nil || o.DestinationId == nil {
		var ret string
		return ret
	}
	return *o.DestinationId
}

// GetDestinationIdOk returns a tuple with the DestinationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusSnapshotDTO) GetDestinationIdOk() (*string, bool) {
	if o == nil || o.DestinationId == nil {
		return nil, false
	}
	return o.DestinationId, true
}

// HasDestinationId returns a boolean if a field has been set.
func (o *ConnectionStatusSnapshotDTO) HasDestinationId() bool {
	if o != nil && o.DestinationId != nil {
		return true
	}

	return false
}

// SetDestinationId gets a reference to the given string and assigns it to the DestinationId field.
func (o *ConnectionStatusSnapshotDTO) SetDestinationId(v string) {
	o.DestinationId = &v
}

// GetDestinationName returns the DestinationName field value if set, zero value otherwise.
func (o *ConnectionStatusSnapshotDTO) GetDestinationName() string {
	if o == nil || o.DestinationName == nil {
		var ret string
		return ret
	}
	return *o.DestinationName
}

// GetDestinationNameOk returns a tuple with the DestinationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusSnapshotDTO) GetDestinationNameOk() (*string, bool) {
	if o == nil || o.DestinationName == nil {
		return nil, false
	}
	return o.DestinationName, true
}

// HasDestinationName returns a boolean if a field has been set.
func (o *ConnectionStatusSnapshotDTO) HasDestinationName() bool {
	if o != nil && o.DestinationName != nil {
		return true
	}

	return false
}

// SetDestinationName gets a reference to the given string and assigns it to the DestinationName field.
func (o *ConnectionStatusSnapshotDTO) SetDestinationName(v string) {
	o.DestinationName = &v
}

// GetPredictions returns the Predictions field value if set, zero value otherwise.
func (o *ConnectionStatusSnapshotDTO) GetPredictions() ConnectionStatusPredictionsSnapshotDTO {
	if o == nil || o.Predictions == nil {
		var ret ConnectionStatusPredictionsSnapshotDTO
		return ret
	}
	return *o.Predictions
}

// GetPredictionsOk returns a tuple with the Predictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusSnapshotDTO) GetPredictionsOk() (*ConnectionStatusPredictionsSnapshotDTO, bool) {
	if o == nil || o.Predictions == nil {
		return nil, false
	}
	return o.Predictions, true
}

// HasPredictions returns a boolean if a field has been set.
func (o *ConnectionStatusSnapshotDTO) HasPredictions() bool {
	if o != nil && o.Predictions != nil {
		return true
	}

	return false
}

// SetPredictions gets a reference to the given ConnectionStatusPredictionsSnapshotDTO and assigns it to the Predictions field.
func (o *ConnectionStatusSnapshotDTO) SetPredictions(v ConnectionStatusPredictionsSnapshotDTO) {
	o.Predictions = &v
}

// GetFlowFilesIn returns the FlowFilesIn field value if set, zero value otherwise.
func (o *ConnectionStatusSnapshotDTO) GetFlowFilesIn() int32 {
	if o == nil || o.FlowFilesIn == nil {
		var ret int32
		return ret
	}
	return *o.FlowFilesIn
}

// GetFlowFilesInOk returns a tuple with the FlowFilesIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusSnapshotDTO) GetFlowFilesInOk() (*int32, bool) {
	if o == nil || o.FlowFilesIn == nil {
		return nil, false
	}
	return o.FlowFilesIn, true
}

// HasFlowFilesIn returns a boolean if a field has been set.
func (o *ConnectionStatusSnapshotDTO) HasFlowFilesIn() bool {
	if o != nil && o.FlowFilesIn != nil {
		return true
	}

	return false
}

// SetFlowFilesIn gets a reference to the given int32 and assigns it to the FlowFilesIn field.
func (o *ConnectionStatusSnapshotDTO) SetFlowFilesIn(v int32) {
	o.FlowFilesIn = &v
}

// GetBytesIn returns the BytesIn field value if set, zero value otherwise.
func (o *ConnectionStatusSnapshotDTO) GetBytesIn() int64 {
	if o == nil || o.BytesIn == nil {
		var ret int64
		return ret
	}
	return *o.BytesIn
}

// GetBytesInOk returns a tuple with the BytesIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusSnapshotDTO) GetBytesInOk() (*int64, bool) {
	if o == nil || o.BytesIn == nil {
		return nil, false
	}
	return o.BytesIn, true
}

// HasBytesIn returns a boolean if a field has been set.
func (o *ConnectionStatusSnapshotDTO) HasBytesIn() bool {
	if o != nil && o.BytesIn != nil {
		return true
	}

	return false
}

// SetBytesIn gets a reference to the given int64 and assigns it to the BytesIn field.
func (o *ConnectionStatusSnapshotDTO) SetBytesIn(v int64) {
	o.BytesIn = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *ConnectionStatusSnapshotDTO) GetInput() string {
	if o == nil || o.Input == nil {
		var ret string
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusSnapshotDTO) GetInputOk() (*string, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *ConnectionStatusSnapshotDTO) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given string and assigns it to the Input field.
func (o *ConnectionStatusSnapshotDTO) SetInput(v string) {
	o.Input = &v
}

// GetFlowFilesOut returns the FlowFilesOut field value if set, zero value otherwise.
func (o *ConnectionStatusSnapshotDTO) GetFlowFilesOut() int32 {
	if o == nil || o.FlowFilesOut == nil {
		var ret int32
		return ret
	}
	return *o.FlowFilesOut
}

// GetFlowFilesOutOk returns a tuple with the FlowFilesOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusSnapshotDTO) GetFlowFilesOutOk() (*int32, bool) {
	if o == nil || o.FlowFilesOut == nil {
		return nil, false
	}
	return o.FlowFilesOut, true
}

// HasFlowFilesOut returns a boolean if a field has been set.
func (o *ConnectionStatusSnapshotDTO) HasFlowFilesOut() bool {
	if o != nil && o.FlowFilesOut != nil {
		return true
	}

	return false
}

// SetFlowFilesOut gets a reference to the given int32 and assigns it to the FlowFilesOut field.
func (o *ConnectionStatusSnapshotDTO) SetFlowFilesOut(v int32) {
	o.FlowFilesOut = &v
}

// GetBytesOut returns the BytesOut field value if set, zero value otherwise.
func (o *ConnectionStatusSnapshotDTO) GetBytesOut() int64 {
	if o == nil || o.BytesOut == nil {
		var ret int64
		return ret
	}
	return *o.BytesOut
}

// GetBytesOutOk returns a tuple with the BytesOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusSnapshotDTO) GetBytesOutOk() (*int64, bool) {
	if o == nil || o.BytesOut == nil {
		return nil, false
	}
	return o.BytesOut, true
}

// HasBytesOut returns a boolean if a field has been set.
func (o *ConnectionStatusSnapshotDTO) HasBytesOut() bool {
	if o != nil && o.BytesOut != nil {
		return true
	}

	return false
}

// SetBytesOut gets a reference to the given int64 and assigns it to the BytesOut field.
func (o *ConnectionStatusSnapshotDTO) SetBytesOut(v int64) {
	o.BytesOut = &v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *ConnectionStatusSnapshotDTO) GetOutput() string {
	if o == nil || o.Output == nil {
		var ret string
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusSnapshotDTO) GetOutputOk() (*string, bool) {
	if o == nil || o.Output == nil {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *ConnectionStatusSnapshotDTO) HasOutput() bool {
	if o != nil && o.Output != nil {
		return true
	}

	return false
}

// SetOutput gets a reference to the given string and assigns it to the Output field.
func (o *ConnectionStatusSnapshotDTO) SetOutput(v string) {
	o.Output = &v
}

// GetFlowFilesQueued returns the FlowFilesQueued field value if set, zero value otherwise.
func (o *ConnectionStatusSnapshotDTO) GetFlowFilesQueued() int32 {
	if o == nil || o.FlowFilesQueued == nil {
		var ret int32
		return ret
	}
	return *o.FlowFilesQueued
}

// GetFlowFilesQueuedOk returns a tuple with the FlowFilesQueued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusSnapshotDTO) GetFlowFilesQueuedOk() (*int32, bool) {
	if o == nil || o.FlowFilesQueued == nil {
		return nil, false
	}
	return o.FlowFilesQueued, true
}

// HasFlowFilesQueued returns a boolean if a field has been set.
func (o *ConnectionStatusSnapshotDTO) HasFlowFilesQueued() bool {
	if o != nil && o.FlowFilesQueued != nil {
		return true
	}

	return false
}

// SetFlowFilesQueued gets a reference to the given int32 and assigns it to the FlowFilesQueued field.
func (o *ConnectionStatusSnapshotDTO) SetFlowFilesQueued(v int32) {
	o.FlowFilesQueued = &v
}

// GetBytesQueued returns the BytesQueued field value if set, zero value otherwise.
func (o *ConnectionStatusSnapshotDTO) GetBytesQueued() int64 {
	if o == nil || o.BytesQueued == nil {
		var ret int64
		return ret
	}
	return *o.BytesQueued
}

// GetBytesQueuedOk returns a tuple with the BytesQueued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusSnapshotDTO) GetBytesQueuedOk() (*int64, bool) {
	if o == nil || o.BytesQueued == nil {
		return nil, false
	}
	return o.BytesQueued, true
}

// HasBytesQueued returns a boolean if a field has been set.
func (o *ConnectionStatusSnapshotDTO) HasBytesQueued() bool {
	if o != nil && o.BytesQueued != nil {
		return true
	}

	return false
}

// SetBytesQueued gets a reference to the given int64 and assigns it to the BytesQueued field.
func (o *ConnectionStatusSnapshotDTO) SetBytesQueued(v int64) {
	o.BytesQueued = &v
}

// GetQueued returns the Queued field value if set, zero value otherwise.
func (o *ConnectionStatusSnapshotDTO) GetQueued() string {
	if o == nil || o.Queued == nil {
		var ret string
		return ret
	}
	return *o.Queued
}

// GetQueuedOk returns a tuple with the Queued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusSnapshotDTO) GetQueuedOk() (*string, bool) {
	if o == nil || o.Queued == nil {
		return nil, false
	}
	return o.Queued, true
}

// HasQueued returns a boolean if a field has been set.
func (o *ConnectionStatusSnapshotDTO) HasQueued() bool {
	if o != nil && o.Queued != nil {
		return true
	}

	return false
}

// SetQueued gets a reference to the given string and assigns it to the Queued field.
func (o *ConnectionStatusSnapshotDTO) SetQueued(v string) {
	o.Queued = &v
}

// GetQueuedSize returns the QueuedSize field value if set, zero value otherwise.
func (o *ConnectionStatusSnapshotDTO) GetQueuedSize() string {
	if o == nil || o.QueuedSize == nil {
		var ret string
		return ret
	}
	return *o.QueuedSize
}

// GetQueuedSizeOk returns a tuple with the QueuedSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusSnapshotDTO) GetQueuedSizeOk() (*string, bool) {
	if o == nil || o.QueuedSize == nil {
		return nil, false
	}
	return o.QueuedSize, true
}

// HasQueuedSize returns a boolean if a field has been set.
func (o *ConnectionStatusSnapshotDTO) HasQueuedSize() bool {
	if o != nil && o.QueuedSize != nil {
		return true
	}

	return false
}

// SetQueuedSize gets a reference to the given string and assigns it to the QueuedSize field.
func (o *ConnectionStatusSnapshotDTO) SetQueuedSize(v string) {
	o.QueuedSize = &v
}

// GetQueuedCount returns the QueuedCount field value if set, zero value otherwise.
func (o *ConnectionStatusSnapshotDTO) GetQueuedCount() string {
	if o == nil || o.QueuedCount == nil {
		var ret string
		return ret
	}
	return *o.QueuedCount
}

// GetQueuedCountOk returns a tuple with the QueuedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusSnapshotDTO) GetQueuedCountOk() (*string, bool) {
	if o == nil || o.QueuedCount == nil {
		return nil, false
	}
	return o.QueuedCount, true
}

// HasQueuedCount returns a boolean if a field has been set.
func (o *ConnectionStatusSnapshotDTO) HasQueuedCount() bool {
	if o != nil && o.QueuedCount != nil {
		return true
	}

	return false
}

// SetQueuedCount gets a reference to the given string and assigns it to the QueuedCount field.
func (o *ConnectionStatusSnapshotDTO) SetQueuedCount(v string) {
	o.QueuedCount = &v
}

// GetPercentUseCount returns the PercentUseCount field value if set, zero value otherwise.
func (o *ConnectionStatusSnapshotDTO) GetPercentUseCount() int32 {
	if o == nil || o.PercentUseCount == nil {
		var ret int32
		return ret
	}
	return *o.PercentUseCount
}

// GetPercentUseCountOk returns a tuple with the PercentUseCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusSnapshotDTO) GetPercentUseCountOk() (*int32, bool) {
	if o == nil || o.PercentUseCount == nil {
		return nil, false
	}
	return o.PercentUseCount, true
}

// HasPercentUseCount returns a boolean if a field has been set.
func (o *ConnectionStatusSnapshotDTO) HasPercentUseCount() bool {
	if o != nil && o.PercentUseCount != nil {
		return true
	}

	return false
}

// SetPercentUseCount gets a reference to the given int32 and assigns it to the PercentUseCount field.
func (o *ConnectionStatusSnapshotDTO) SetPercentUseCount(v int32) {
	o.PercentUseCount = &v
}

// GetPercentUseBytes returns the PercentUseBytes field value if set, zero value otherwise.
func (o *ConnectionStatusSnapshotDTO) GetPercentUseBytes() int32 {
	if o == nil || o.PercentUseBytes == nil {
		var ret int32
		return ret
	}
	return *o.PercentUseBytes
}

// GetPercentUseBytesOk returns a tuple with the PercentUseBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusSnapshotDTO) GetPercentUseBytesOk() (*int32, bool) {
	if o == nil || o.PercentUseBytes == nil {
		return nil, false
	}
	return o.PercentUseBytes, true
}

// HasPercentUseBytes returns a boolean if a field has been set.
func (o *ConnectionStatusSnapshotDTO) HasPercentUseBytes() bool {
	if o != nil && o.PercentUseBytes != nil {
		return true
	}

	return false
}

// SetPercentUseBytes gets a reference to the given int32 and assigns it to the PercentUseBytes field.
func (o *ConnectionStatusSnapshotDTO) SetPercentUseBytes(v int32) {
	o.PercentUseBytes = &v
}

func (o ConnectionStatusSnapshotDTO) MarshalJSON() ([]byte, error) {
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
	if o.SourceId != nil {
		toSerialize["sourceId"] = o.SourceId
	}
	if o.SourceName != nil {
		toSerialize["sourceName"] = o.SourceName
	}
	if o.DestinationId != nil {
		toSerialize["destinationId"] = o.DestinationId
	}
	if o.DestinationName != nil {
		toSerialize["destinationName"] = o.DestinationName
	}
	if o.Predictions != nil {
		toSerialize["predictions"] = o.Predictions
	}
	if o.FlowFilesIn != nil {
		toSerialize["flowFilesIn"] = o.FlowFilesIn
	}
	if o.BytesIn != nil {
		toSerialize["bytesIn"] = o.BytesIn
	}
	if o.Input != nil {
		toSerialize["input"] = o.Input
	}
	if o.FlowFilesOut != nil {
		toSerialize["flowFilesOut"] = o.FlowFilesOut
	}
	if o.BytesOut != nil {
		toSerialize["bytesOut"] = o.BytesOut
	}
	if o.Output != nil {
		toSerialize["output"] = o.Output
	}
	if o.FlowFilesQueued != nil {
		toSerialize["flowFilesQueued"] = o.FlowFilesQueued
	}
	if o.BytesQueued != nil {
		toSerialize["bytesQueued"] = o.BytesQueued
	}
	if o.Queued != nil {
		toSerialize["queued"] = o.Queued
	}
	if o.QueuedSize != nil {
		toSerialize["queuedSize"] = o.QueuedSize
	}
	if o.QueuedCount != nil {
		toSerialize["queuedCount"] = o.QueuedCount
	}
	if o.PercentUseCount != nil {
		toSerialize["percentUseCount"] = o.PercentUseCount
	}
	if o.PercentUseBytes != nil {
		toSerialize["percentUseBytes"] = o.PercentUseBytes
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionStatusSnapshotDTO struct {
	value *ConnectionStatusSnapshotDTO
	isSet bool
}

func (v NullableConnectionStatusSnapshotDTO) Get() *ConnectionStatusSnapshotDTO {
	return v.value
}

func (v *NullableConnectionStatusSnapshotDTO) Set(val *ConnectionStatusSnapshotDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionStatusSnapshotDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionStatusSnapshotDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionStatusSnapshotDTO(val *ConnectionStatusSnapshotDTO) *NullableConnectionStatusSnapshotDTO {
	return &NullableConnectionStatusSnapshotDTO{value: val, isSet: true}
}

func (v NullableConnectionStatusSnapshotDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionStatusSnapshotDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
