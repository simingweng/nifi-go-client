# VariableRegistryEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessGroupRevision** | Pointer to [**RevisionDTO**](RevisionDTO.md) |  | [optional] 
**VariableRegistry** | Pointer to [**VariableRegistryDTO**](VariableRegistryDTO.md) |  | [optional] 
**DisconnectedNodeAcknowledged** | Pointer to **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] 

## Methods

### NewVariableRegistryEntity

`func NewVariableRegistryEntity() *VariableRegistryEntity`

NewVariableRegistryEntity instantiates a new VariableRegistryEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableRegistryEntityWithDefaults

`func NewVariableRegistryEntityWithDefaults() *VariableRegistryEntity`

NewVariableRegistryEntityWithDefaults instantiates a new VariableRegistryEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessGroupRevision

`func (o *VariableRegistryEntity) GetProcessGroupRevision() RevisionDTO`

GetProcessGroupRevision returns the ProcessGroupRevision field if non-nil, zero value otherwise.

### GetProcessGroupRevisionOk

`func (o *VariableRegistryEntity) GetProcessGroupRevisionOk() (*RevisionDTO, bool)`

GetProcessGroupRevisionOk returns a tuple with the ProcessGroupRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupRevision

`func (o *VariableRegistryEntity) SetProcessGroupRevision(v RevisionDTO)`

SetProcessGroupRevision sets ProcessGroupRevision field to given value.

### HasProcessGroupRevision

`func (o *VariableRegistryEntity) HasProcessGroupRevision() bool`

HasProcessGroupRevision returns a boolean if a field has been set.

### GetVariableRegistry

`func (o *VariableRegistryEntity) GetVariableRegistry() VariableRegistryDTO`

GetVariableRegistry returns the VariableRegistry field if non-nil, zero value otherwise.

### GetVariableRegistryOk

`func (o *VariableRegistryEntity) GetVariableRegistryOk() (*VariableRegistryDTO, bool)`

GetVariableRegistryOk returns a tuple with the VariableRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableRegistry

`func (o *VariableRegistryEntity) SetVariableRegistry(v VariableRegistryDTO)`

SetVariableRegistry sets VariableRegistry field to given value.

### HasVariableRegistry

`func (o *VariableRegistryEntity) HasVariableRegistry() bool`

HasVariableRegistry returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *VariableRegistryEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *VariableRegistryEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *VariableRegistryEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *VariableRegistryEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


