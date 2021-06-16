# StartVersionControlRequestEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionedFlow** | Pointer to [**VersionedFlowDTO**](VersionedFlowDTO.md) |  | [optional] 
**ProcessGroupRevision** | Pointer to [**RevisionDTO**](RevisionDTO.md) |  | [optional] 
**DisconnectedNodeAcknowledged** | Pointer to **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] 

## Methods

### NewStartVersionControlRequestEntity

`func NewStartVersionControlRequestEntity() *StartVersionControlRequestEntity`

NewStartVersionControlRequestEntity instantiates a new StartVersionControlRequestEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartVersionControlRequestEntityWithDefaults

`func NewStartVersionControlRequestEntityWithDefaults() *StartVersionControlRequestEntity`

NewStartVersionControlRequestEntityWithDefaults instantiates a new StartVersionControlRequestEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionedFlow

`func (o *StartVersionControlRequestEntity) GetVersionedFlow() VersionedFlowDTO`

GetVersionedFlow returns the VersionedFlow field if non-nil, zero value otherwise.

### GetVersionedFlowOk

`func (o *StartVersionControlRequestEntity) GetVersionedFlowOk() (*VersionedFlowDTO, bool)`

GetVersionedFlowOk returns a tuple with the VersionedFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedFlow

`func (o *StartVersionControlRequestEntity) SetVersionedFlow(v VersionedFlowDTO)`

SetVersionedFlow sets VersionedFlow field to given value.

### HasVersionedFlow

`func (o *StartVersionControlRequestEntity) HasVersionedFlow() bool`

HasVersionedFlow returns a boolean if a field has been set.

### GetProcessGroupRevision

`func (o *StartVersionControlRequestEntity) GetProcessGroupRevision() RevisionDTO`

GetProcessGroupRevision returns the ProcessGroupRevision field if non-nil, zero value otherwise.

### GetProcessGroupRevisionOk

`func (o *StartVersionControlRequestEntity) GetProcessGroupRevisionOk() (*RevisionDTO, bool)`

GetProcessGroupRevisionOk returns a tuple with the ProcessGroupRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupRevision

`func (o *StartVersionControlRequestEntity) SetProcessGroupRevision(v RevisionDTO)`

SetProcessGroupRevision sets ProcessGroupRevision field to given value.

### HasProcessGroupRevision

`func (o *StartVersionControlRequestEntity) HasProcessGroupRevision() bool`

HasProcessGroupRevision returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *StartVersionControlRequestEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *StartVersionControlRequestEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *StartVersionControlRequestEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *StartVersionControlRequestEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


