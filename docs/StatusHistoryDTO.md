# StatusHistoryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Generated** | Pointer to **string** | When the status history was generated. | [optional] 
**ComponentDetails** | Pointer to **map[string]string** | A Map of key/value pairs that describe the component that the status history belongs to | [optional] 
**FieldDescriptors** | Pointer to [**[]StatusDescriptorDTO**](StatusDescriptorDTO.md) | The Descriptors that provide information on each of the metrics provided in the status history | [optional] 
**AggregateSnapshots** | Pointer to [**[]StatusSnapshotDTO**](StatusSnapshotDTO.md) | A list of StatusSnapshotDTO objects that provide the actual metric values for the component. If the NiFi instance is clustered, this will represent the aggregate status across all nodes. If the NiFi instance is not clustered, this will represent the status of the entire NiFi instance. | [optional] 
**NodeSnapshots** | Pointer to [**[]NodeStatusSnapshotsDTO**](NodeStatusSnapshotsDTO.md) | The NodeStatusSnapshotsDTO objects that provide the actual metric values for the component, for each node. If the NiFi instance is not clustered, this value will be null. | [optional] 

## Methods

### NewStatusHistoryDTO

`func NewStatusHistoryDTO() *StatusHistoryDTO`

NewStatusHistoryDTO instantiates a new StatusHistoryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusHistoryDTOWithDefaults

`func NewStatusHistoryDTOWithDefaults() *StatusHistoryDTO`

NewStatusHistoryDTOWithDefaults instantiates a new StatusHistoryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenerated

`func (o *StatusHistoryDTO) GetGenerated() string`

GetGenerated returns the Generated field if non-nil, zero value otherwise.

### GetGeneratedOk

`func (o *StatusHistoryDTO) GetGeneratedOk() (*string, bool)`

GetGeneratedOk returns a tuple with the Generated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerated

`func (o *StatusHistoryDTO) SetGenerated(v string)`

SetGenerated sets Generated field to given value.

### HasGenerated

`func (o *StatusHistoryDTO) HasGenerated() bool`

HasGenerated returns a boolean if a field has been set.

### GetComponentDetails

`func (o *StatusHistoryDTO) GetComponentDetails() map[string]string`

GetComponentDetails returns the ComponentDetails field if non-nil, zero value otherwise.

### GetComponentDetailsOk

`func (o *StatusHistoryDTO) GetComponentDetailsOk() (*map[string]string, bool)`

GetComponentDetailsOk returns a tuple with the ComponentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentDetails

`func (o *StatusHistoryDTO) SetComponentDetails(v map[string]string)`

SetComponentDetails sets ComponentDetails field to given value.

### HasComponentDetails

`func (o *StatusHistoryDTO) HasComponentDetails() bool`

HasComponentDetails returns a boolean if a field has been set.

### GetFieldDescriptors

`func (o *StatusHistoryDTO) GetFieldDescriptors() []StatusDescriptorDTO`

GetFieldDescriptors returns the FieldDescriptors field if non-nil, zero value otherwise.

### GetFieldDescriptorsOk

`func (o *StatusHistoryDTO) GetFieldDescriptorsOk() (*[]StatusDescriptorDTO, bool)`

GetFieldDescriptorsOk returns a tuple with the FieldDescriptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldDescriptors

`func (o *StatusHistoryDTO) SetFieldDescriptors(v []StatusDescriptorDTO)`

SetFieldDescriptors sets FieldDescriptors field to given value.

### HasFieldDescriptors

`func (o *StatusHistoryDTO) HasFieldDescriptors() bool`

HasFieldDescriptors returns a boolean if a field has been set.

### GetAggregateSnapshots

`func (o *StatusHistoryDTO) GetAggregateSnapshots() []StatusSnapshotDTO`

GetAggregateSnapshots returns the AggregateSnapshots field if non-nil, zero value otherwise.

### GetAggregateSnapshotsOk

`func (o *StatusHistoryDTO) GetAggregateSnapshotsOk() (*[]StatusSnapshotDTO, bool)`

GetAggregateSnapshotsOk returns a tuple with the AggregateSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateSnapshots

`func (o *StatusHistoryDTO) SetAggregateSnapshots(v []StatusSnapshotDTO)`

SetAggregateSnapshots sets AggregateSnapshots field to given value.

### HasAggregateSnapshots

`func (o *StatusHistoryDTO) HasAggregateSnapshots() bool`

HasAggregateSnapshots returns a boolean if a field has been set.

### GetNodeSnapshots

`func (o *StatusHistoryDTO) GetNodeSnapshots() []NodeStatusSnapshotsDTO`

GetNodeSnapshots returns the NodeSnapshots field if non-nil, zero value otherwise.

### GetNodeSnapshotsOk

`func (o *StatusHistoryDTO) GetNodeSnapshotsOk() (*[]NodeStatusSnapshotsDTO, bool)`

GetNodeSnapshotsOk returns a tuple with the NodeSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSnapshots

`func (o *StatusHistoryDTO) SetNodeSnapshots(v []NodeStatusSnapshotsDTO)`

SetNodeSnapshots sets NodeSnapshots field to given value.

### HasNodeSnapshots

`func (o *StatusHistoryDTO) HasNodeSnapshots() bool`

HasNodeSnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


