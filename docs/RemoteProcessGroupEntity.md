# RemoteProcessGroupEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | Pointer to [**RevisionDTO**](RevisionDTO.md) |  | [optional] 
**Id** | Pointer to **string** | The id of the component. | [optional] 
**Uri** | Pointer to **string** | The URI for futures requests to the component. | [optional] 
**Position** | Pointer to [**PositionDTO**](PositionDTO.md) |  | [optional] 
**Permissions** | Pointer to [**PermissionsDTO**](PermissionsDTO.md) |  | [optional] 
**Bulletins** | Pointer to [**[]BulletinEntity**](BulletinEntity.md) | The bulletins for this component. | [optional] 
**DisconnectedNodeAcknowledged** | Pointer to **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] 
**Component** | Pointer to [**RemoteProcessGroupDTO**](RemoteProcessGroupDTO.md) |  | [optional] 
**Status** | Pointer to [**RemoteProcessGroupStatusDTO**](RemoteProcessGroupStatusDTO.md) |  | [optional] 
**InputPortCount** | Pointer to **int32** | The number of remote input ports currently available on the target. | [optional] 
**OutputPortCount** | Pointer to **int32** | The number of remote output ports currently available on the target. | [optional] 
**OperatePermissions** | Pointer to [**PermissionsDTO**](PermissionsDTO.md) |  | [optional] 

## Methods

### NewRemoteProcessGroupEntity

`func NewRemoteProcessGroupEntity() *RemoteProcessGroupEntity`

NewRemoteProcessGroupEntity instantiates a new RemoteProcessGroupEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteProcessGroupEntityWithDefaults

`func NewRemoteProcessGroupEntityWithDefaults() *RemoteProcessGroupEntity`

NewRemoteProcessGroupEntityWithDefaults instantiates a new RemoteProcessGroupEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *RemoteProcessGroupEntity) GetRevision() RevisionDTO`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *RemoteProcessGroupEntity) GetRevisionOk() (*RevisionDTO, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *RemoteProcessGroupEntity) SetRevision(v RevisionDTO)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *RemoteProcessGroupEntity) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetId

`func (o *RemoteProcessGroupEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoteProcessGroupEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoteProcessGroupEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RemoteProcessGroupEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUri

`func (o *RemoteProcessGroupEntity) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *RemoteProcessGroupEntity) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *RemoteProcessGroupEntity) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *RemoteProcessGroupEntity) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetPosition

`func (o *RemoteProcessGroupEntity) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *RemoteProcessGroupEntity) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *RemoteProcessGroupEntity) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *RemoteProcessGroupEntity) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPermissions

`func (o *RemoteProcessGroupEntity) GetPermissions() PermissionsDTO`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RemoteProcessGroupEntity) GetPermissionsOk() (*PermissionsDTO, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RemoteProcessGroupEntity) SetPermissions(v PermissionsDTO)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *RemoteProcessGroupEntity) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetBulletins

`func (o *RemoteProcessGroupEntity) GetBulletins() []BulletinEntity`

GetBulletins returns the Bulletins field if non-nil, zero value otherwise.

### GetBulletinsOk

`func (o *RemoteProcessGroupEntity) GetBulletinsOk() (*[]BulletinEntity, bool)`

GetBulletinsOk returns a tuple with the Bulletins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletins

`func (o *RemoteProcessGroupEntity) SetBulletins(v []BulletinEntity)`

SetBulletins sets Bulletins field to given value.

### HasBulletins

`func (o *RemoteProcessGroupEntity) HasBulletins() bool`

HasBulletins returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *RemoteProcessGroupEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *RemoteProcessGroupEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *RemoteProcessGroupEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *RemoteProcessGroupEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.

### GetComponent

`func (o *RemoteProcessGroupEntity) GetComponent() RemoteProcessGroupDTO`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *RemoteProcessGroupEntity) GetComponentOk() (*RemoteProcessGroupDTO, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *RemoteProcessGroupEntity) SetComponent(v RemoteProcessGroupDTO)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *RemoteProcessGroupEntity) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetStatus

`func (o *RemoteProcessGroupEntity) GetStatus() RemoteProcessGroupStatusDTO`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RemoteProcessGroupEntity) GetStatusOk() (*RemoteProcessGroupStatusDTO, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RemoteProcessGroupEntity) SetStatus(v RemoteProcessGroupStatusDTO)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RemoteProcessGroupEntity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInputPortCount

`func (o *RemoteProcessGroupEntity) GetInputPortCount() int32`

GetInputPortCount returns the InputPortCount field if non-nil, zero value otherwise.

### GetInputPortCountOk

`func (o *RemoteProcessGroupEntity) GetInputPortCountOk() (*int32, bool)`

GetInputPortCountOk returns a tuple with the InputPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPortCount

`func (o *RemoteProcessGroupEntity) SetInputPortCount(v int32)`

SetInputPortCount sets InputPortCount field to given value.

### HasInputPortCount

`func (o *RemoteProcessGroupEntity) HasInputPortCount() bool`

HasInputPortCount returns a boolean if a field has been set.

### GetOutputPortCount

`func (o *RemoteProcessGroupEntity) GetOutputPortCount() int32`

GetOutputPortCount returns the OutputPortCount field if non-nil, zero value otherwise.

### GetOutputPortCountOk

`func (o *RemoteProcessGroupEntity) GetOutputPortCountOk() (*int32, bool)`

GetOutputPortCountOk returns a tuple with the OutputPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPortCount

`func (o *RemoteProcessGroupEntity) SetOutputPortCount(v int32)`

SetOutputPortCount sets OutputPortCount field to given value.

### HasOutputPortCount

`func (o *RemoteProcessGroupEntity) HasOutputPortCount() bool`

HasOutputPortCount returns a boolean if a field has been set.

### GetOperatePermissions

`func (o *RemoteProcessGroupEntity) GetOperatePermissions() PermissionsDTO`

GetOperatePermissions returns the OperatePermissions field if non-nil, zero value otherwise.

### GetOperatePermissionsOk

`func (o *RemoteProcessGroupEntity) GetOperatePermissionsOk() (*PermissionsDTO, bool)`

GetOperatePermissionsOk returns a tuple with the OperatePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatePermissions

`func (o *RemoteProcessGroupEntity) SetOperatePermissions(v PermissionsDTO)`

SetOperatePermissions sets OperatePermissions field to given value.

### HasOperatePermissions

`func (o *RemoteProcessGroupEntity) HasOperatePermissions() bool`

HasOperatePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


