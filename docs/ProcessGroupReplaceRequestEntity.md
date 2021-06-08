# ProcessGroupReplaceRequestEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessGroupRevision** | Pointer to [**RevisionDTO**](RevisionDTO.md) |  | [optional] 
**Request** | Pointer to [**ProcessGroupReplaceRequestDTO**](ProcessGroupReplaceRequestDTO.md) |  | [optional] 
**VersionedFlowSnapshot** | Pointer to [**VersionedFlowSnapshot**](VersionedFlowSnapshot.md) |  | [optional] 

## Methods

### NewProcessGroupReplaceRequestEntity

`func NewProcessGroupReplaceRequestEntity() *ProcessGroupReplaceRequestEntity`

NewProcessGroupReplaceRequestEntity instantiates a new ProcessGroupReplaceRequestEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessGroupReplaceRequestEntityWithDefaults

`func NewProcessGroupReplaceRequestEntityWithDefaults() *ProcessGroupReplaceRequestEntity`

NewProcessGroupReplaceRequestEntityWithDefaults instantiates a new ProcessGroupReplaceRequestEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessGroupRevision

`func (o *ProcessGroupReplaceRequestEntity) GetProcessGroupRevision() RevisionDTO`

GetProcessGroupRevision returns the ProcessGroupRevision field if non-nil, zero value otherwise.

### GetProcessGroupRevisionOk

`func (o *ProcessGroupReplaceRequestEntity) GetProcessGroupRevisionOk() (*RevisionDTO, bool)`

GetProcessGroupRevisionOk returns a tuple with the ProcessGroupRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupRevision

`func (o *ProcessGroupReplaceRequestEntity) SetProcessGroupRevision(v RevisionDTO)`

SetProcessGroupRevision sets ProcessGroupRevision field to given value.

### HasProcessGroupRevision

`func (o *ProcessGroupReplaceRequestEntity) HasProcessGroupRevision() bool`

HasProcessGroupRevision returns a boolean if a field has been set.

### GetRequest

`func (o *ProcessGroupReplaceRequestEntity) GetRequest() ProcessGroupReplaceRequestDTO`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *ProcessGroupReplaceRequestEntity) GetRequestOk() (*ProcessGroupReplaceRequestDTO, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *ProcessGroupReplaceRequestEntity) SetRequest(v ProcessGroupReplaceRequestDTO)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *ProcessGroupReplaceRequestEntity) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetVersionedFlowSnapshot

`func (o *ProcessGroupReplaceRequestEntity) GetVersionedFlowSnapshot() VersionedFlowSnapshot`

GetVersionedFlowSnapshot returns the VersionedFlowSnapshot field if non-nil, zero value otherwise.

### GetVersionedFlowSnapshotOk

`func (o *ProcessGroupReplaceRequestEntity) GetVersionedFlowSnapshotOk() (*VersionedFlowSnapshot, bool)`

GetVersionedFlowSnapshotOk returns a tuple with the VersionedFlowSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedFlowSnapshot

`func (o *ProcessGroupReplaceRequestEntity) SetVersionedFlowSnapshot(v VersionedFlowSnapshot)`

SetVersionedFlowSnapshot sets VersionedFlowSnapshot field to given value.

### HasVersionedFlowSnapshot

`func (o *ProcessGroupReplaceRequestEntity) HasVersionedFlowSnapshot() bool`

HasVersionedFlowSnapshot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


