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

// FlowFileSummaryDTO struct for FlowFileSummaryDTO
type FlowFileSummaryDTO struct {
	// The URI that can be used to access this FlowFile.
	Uri *string `json:"uri,omitempty"`
	// The FlowFile UUID.
	Uuid *string `json:"uuid,omitempty"`
	// The FlowFile filename.
	Filename *string `json:"filename,omitempty"`
	// The FlowFile's position in the queue.
	Position *int32 `json:"position,omitempty"`
	// The FlowFile file size.
	Size *int64 `json:"size,omitempty"`
	// How long this FlowFile has been enqueued.
	QueuedDuration *int64 `json:"queuedDuration,omitempty"`
	// Duration since the FlowFile's greatest ancestor entered the flow.
	LineageDuration *int64 `json:"lineageDuration,omitempty"`
	// How long in milliseconds until the FlowFile penalty expires.
	PenaltyExpiresIn *int64 `json:"penaltyExpiresIn,omitempty"`
	// The id of the node where this FlowFile resides.
	ClusterNodeId *string `json:"clusterNodeId,omitempty"`
	// The label for the node where this FlowFile resides.
	ClusterNodeAddress *string `json:"clusterNodeAddress,omitempty"`
	// If the FlowFile is penalized.
	Penalized *bool `json:"penalized,omitempty"`
}

// NewFlowFileSummaryDTO instantiates a new FlowFileSummaryDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowFileSummaryDTO() *FlowFileSummaryDTO {
	this := FlowFileSummaryDTO{}
	return &this
}

// NewFlowFileSummaryDTOWithDefaults instantiates a new FlowFileSummaryDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowFileSummaryDTOWithDefaults() *FlowFileSummaryDTO {
	this := FlowFileSummaryDTO{}
	return &this
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *FlowFileSummaryDTO) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowFileSummaryDTO) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *FlowFileSummaryDTO) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *FlowFileSummaryDTO) SetUri(v string) {
	o.Uri = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *FlowFileSummaryDTO) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowFileSummaryDTO) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *FlowFileSummaryDTO) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *FlowFileSummaryDTO) SetUuid(v string) {
	o.Uuid = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *FlowFileSummaryDTO) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowFileSummaryDTO) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *FlowFileSummaryDTO) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *FlowFileSummaryDTO) SetFilename(v string) {
	o.Filename = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *FlowFileSummaryDTO) GetPosition() int32 {
	if o == nil || o.Position == nil {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowFileSummaryDTO) GetPositionOk() (*int32, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *FlowFileSummaryDTO) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *FlowFileSummaryDTO) SetPosition(v int32) {
	o.Position = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *FlowFileSummaryDTO) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowFileSummaryDTO) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *FlowFileSummaryDTO) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *FlowFileSummaryDTO) SetSize(v int64) {
	o.Size = &v
}

// GetQueuedDuration returns the QueuedDuration field value if set, zero value otherwise.
func (o *FlowFileSummaryDTO) GetQueuedDuration() int64 {
	if o == nil || o.QueuedDuration == nil {
		var ret int64
		return ret
	}
	return *o.QueuedDuration
}

// GetQueuedDurationOk returns a tuple with the QueuedDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowFileSummaryDTO) GetQueuedDurationOk() (*int64, bool) {
	if o == nil || o.QueuedDuration == nil {
		return nil, false
	}
	return o.QueuedDuration, true
}

// HasQueuedDuration returns a boolean if a field has been set.
func (o *FlowFileSummaryDTO) HasQueuedDuration() bool {
	if o != nil && o.QueuedDuration != nil {
		return true
	}

	return false
}

// SetQueuedDuration gets a reference to the given int64 and assigns it to the QueuedDuration field.
func (o *FlowFileSummaryDTO) SetQueuedDuration(v int64) {
	o.QueuedDuration = &v
}

// GetLineageDuration returns the LineageDuration field value if set, zero value otherwise.
func (o *FlowFileSummaryDTO) GetLineageDuration() int64 {
	if o == nil || o.LineageDuration == nil {
		var ret int64
		return ret
	}
	return *o.LineageDuration
}

// GetLineageDurationOk returns a tuple with the LineageDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowFileSummaryDTO) GetLineageDurationOk() (*int64, bool) {
	if o == nil || o.LineageDuration == nil {
		return nil, false
	}
	return o.LineageDuration, true
}

// HasLineageDuration returns a boolean if a field has been set.
func (o *FlowFileSummaryDTO) HasLineageDuration() bool {
	if o != nil && o.LineageDuration != nil {
		return true
	}

	return false
}

// SetLineageDuration gets a reference to the given int64 and assigns it to the LineageDuration field.
func (o *FlowFileSummaryDTO) SetLineageDuration(v int64) {
	o.LineageDuration = &v
}

// GetPenaltyExpiresIn returns the PenaltyExpiresIn field value if set, zero value otherwise.
func (o *FlowFileSummaryDTO) GetPenaltyExpiresIn() int64 {
	if o == nil || o.PenaltyExpiresIn == nil {
		var ret int64
		return ret
	}
	return *o.PenaltyExpiresIn
}

// GetPenaltyExpiresInOk returns a tuple with the PenaltyExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowFileSummaryDTO) GetPenaltyExpiresInOk() (*int64, bool) {
	if o == nil || o.PenaltyExpiresIn == nil {
		return nil, false
	}
	return o.PenaltyExpiresIn, true
}

// HasPenaltyExpiresIn returns a boolean if a field has been set.
func (o *FlowFileSummaryDTO) HasPenaltyExpiresIn() bool {
	if o != nil && o.PenaltyExpiresIn != nil {
		return true
	}

	return false
}

// SetPenaltyExpiresIn gets a reference to the given int64 and assigns it to the PenaltyExpiresIn field.
func (o *FlowFileSummaryDTO) SetPenaltyExpiresIn(v int64) {
	o.PenaltyExpiresIn = &v
}

// GetClusterNodeId returns the ClusterNodeId field value if set, zero value otherwise.
func (o *FlowFileSummaryDTO) GetClusterNodeId() string {
	if o == nil || o.ClusterNodeId == nil {
		var ret string
		return ret
	}
	return *o.ClusterNodeId
}

// GetClusterNodeIdOk returns a tuple with the ClusterNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowFileSummaryDTO) GetClusterNodeIdOk() (*string, bool) {
	if o == nil || o.ClusterNodeId == nil {
		return nil, false
	}
	return o.ClusterNodeId, true
}

// HasClusterNodeId returns a boolean if a field has been set.
func (o *FlowFileSummaryDTO) HasClusterNodeId() bool {
	if o != nil && o.ClusterNodeId != nil {
		return true
	}

	return false
}

// SetClusterNodeId gets a reference to the given string and assigns it to the ClusterNodeId field.
func (o *FlowFileSummaryDTO) SetClusterNodeId(v string) {
	o.ClusterNodeId = &v
}

// GetClusterNodeAddress returns the ClusterNodeAddress field value if set, zero value otherwise.
func (o *FlowFileSummaryDTO) GetClusterNodeAddress() string {
	if o == nil || o.ClusterNodeAddress == nil {
		var ret string
		return ret
	}
	return *o.ClusterNodeAddress
}

// GetClusterNodeAddressOk returns a tuple with the ClusterNodeAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowFileSummaryDTO) GetClusterNodeAddressOk() (*string, bool) {
	if o == nil || o.ClusterNodeAddress == nil {
		return nil, false
	}
	return o.ClusterNodeAddress, true
}

// HasClusterNodeAddress returns a boolean if a field has been set.
func (o *FlowFileSummaryDTO) HasClusterNodeAddress() bool {
	if o != nil && o.ClusterNodeAddress != nil {
		return true
	}

	return false
}

// SetClusterNodeAddress gets a reference to the given string and assigns it to the ClusterNodeAddress field.
func (o *FlowFileSummaryDTO) SetClusterNodeAddress(v string) {
	o.ClusterNodeAddress = &v
}

// GetPenalized returns the Penalized field value if set, zero value otherwise.
func (o *FlowFileSummaryDTO) GetPenalized() bool {
	if o == nil || o.Penalized == nil {
		var ret bool
		return ret
	}
	return *o.Penalized
}

// GetPenalizedOk returns a tuple with the Penalized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowFileSummaryDTO) GetPenalizedOk() (*bool, bool) {
	if o == nil || o.Penalized == nil {
		return nil, false
	}
	return o.Penalized, true
}

// HasPenalized returns a boolean if a field has been set.
func (o *FlowFileSummaryDTO) HasPenalized() bool {
	if o != nil && o.Penalized != nil {
		return true
	}

	return false
}

// SetPenalized gets a reference to the given bool and assigns it to the Penalized field.
func (o *FlowFileSummaryDTO) SetPenalized(v bool) {
	o.Penalized = &v
}

func (o FlowFileSummaryDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Filename != nil {
		toSerialize["filename"] = o.Filename
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.QueuedDuration != nil {
		toSerialize["queuedDuration"] = o.QueuedDuration
	}
	if o.LineageDuration != nil {
		toSerialize["lineageDuration"] = o.LineageDuration
	}
	if o.PenaltyExpiresIn != nil {
		toSerialize["penaltyExpiresIn"] = o.PenaltyExpiresIn
	}
	if o.ClusterNodeId != nil {
		toSerialize["clusterNodeId"] = o.ClusterNodeId
	}
	if o.ClusterNodeAddress != nil {
		toSerialize["clusterNodeAddress"] = o.ClusterNodeAddress
	}
	if o.Penalized != nil {
		toSerialize["penalized"] = o.Penalized
	}
	return json.Marshal(toSerialize)
}

type NullableFlowFileSummaryDTO struct {
	value *FlowFileSummaryDTO
	isSet bool
}

func (v NullableFlowFileSummaryDTO) Get() *FlowFileSummaryDTO {
	return v.value
}

func (v *NullableFlowFileSummaryDTO) Set(val *FlowFileSummaryDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowFileSummaryDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowFileSummaryDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowFileSummaryDTO(val *FlowFileSummaryDTO) *NullableFlowFileSummaryDTO {
	return &NullableFlowFileSummaryDTO{value: val, isSet: true}
}

func (v NullableFlowFileSummaryDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowFileSummaryDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
