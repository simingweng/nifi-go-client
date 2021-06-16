# RemoteProcessGroupPortEntity

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
**RemoteProcessGroupPort** | Pointer to [**RemoteProcessGroupPortDTO**](RemoteProcessGroupPortDTO.md) |  | [optional] 
**OperatePermissions** | Pointer to [**PermissionsDTO**](PermissionsDTO.md) |  | [optional] 

## Methods

### NewRemoteProcessGroupPortEntity

`func NewRemoteProcessGroupPortEntity() *RemoteProcessGroupPortEntity`

NewRemoteProcessGroupPortEntity instantiates a new RemoteProcessGroupPortEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteProcessGroupPortEntityWithDefaults

`func NewRemoteProcessGroupPortEntityWithDefaults() *RemoteProcessGroupPortEntity`

NewRemoteProcessGroupPortEntityWithDefaults instantiates a new RemoteProcessGroupPortEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *RemoteProcessGroupPortEntity) GetRevision() RevisionDTO`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *RemoteProcessGroupPortEntity) GetRevisionOk() (*RevisionDTO, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *RemoteProcessGroupPortEntity) SetRevision(v RevisionDTO)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *RemoteProcessGroupPortEntity) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetId

`func (o *RemoteProcessGroupPortEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoteProcessGroupPortEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoteProcessGroupPortEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RemoteProcessGroupPortEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUri

`func (o *RemoteProcessGroupPortEntity) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *RemoteProcessGroupPortEntity) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *RemoteProcessGroupPortEntity) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *RemoteProcessGroupPortEntity) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetPosition

`func (o *RemoteProcessGroupPortEntity) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *RemoteProcessGroupPortEntity) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *RemoteProcessGroupPortEntity) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *RemoteProcessGroupPortEntity) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPermissions

`func (o *RemoteProcessGroupPortEntity) GetPermissions() PermissionsDTO`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RemoteProcessGroupPortEntity) GetPermissionsOk() (*PermissionsDTO, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RemoteProcessGroupPortEntity) SetPermissions(v PermissionsDTO)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *RemoteProcessGroupPortEntity) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetBulletins

`func (o *RemoteProcessGroupPortEntity) GetBulletins() []BulletinEntity`

GetBulletins returns the Bulletins field if non-nil, zero value otherwise.

### GetBulletinsOk

`func (o *RemoteProcessGroupPortEntity) GetBulletinsOk() (*[]BulletinEntity, bool)`

GetBulletinsOk returns a tuple with the Bulletins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletins

`func (o *RemoteProcessGroupPortEntity) SetBulletins(v []BulletinEntity)`

SetBulletins sets Bulletins field to given value.

### HasBulletins

`func (o *RemoteProcessGroupPortEntity) HasBulletins() bool`

HasBulletins returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *RemoteProcessGroupPortEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *RemoteProcessGroupPortEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *RemoteProcessGroupPortEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *RemoteProcessGroupPortEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.

### GetRemoteProcessGroupPort

`func (o *RemoteProcessGroupPortEntity) GetRemoteProcessGroupPort() RemoteProcessGroupPortDTO`

GetRemoteProcessGroupPort returns the RemoteProcessGroupPort field if non-nil, zero value otherwise.

### GetRemoteProcessGroupPortOk

`func (o *RemoteProcessGroupPortEntity) GetRemoteProcessGroupPortOk() (*RemoteProcessGroupPortDTO, bool)`

GetRemoteProcessGroupPortOk returns a tuple with the RemoteProcessGroupPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProcessGroupPort

`func (o *RemoteProcessGroupPortEntity) SetRemoteProcessGroupPort(v RemoteProcessGroupPortDTO)`

SetRemoteProcessGroupPort sets RemoteProcessGroupPort field to given value.

### HasRemoteProcessGroupPort

`func (o *RemoteProcessGroupPortEntity) HasRemoteProcessGroupPort() bool`

HasRemoteProcessGroupPort returns a boolean if a field has been set.

### GetOperatePermissions

`func (o *RemoteProcessGroupPortEntity) GetOperatePermissions() PermissionsDTO`

GetOperatePermissions returns the OperatePermissions field if non-nil, zero value otherwise.

### GetOperatePermissionsOk

`func (o *RemoteProcessGroupPortEntity) GetOperatePermissionsOk() (*PermissionsDTO, bool)`

GetOperatePermissionsOk returns a tuple with the OperatePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatePermissions

`func (o *RemoteProcessGroupPortEntity) SetOperatePermissions(v PermissionsDTO)`

SetOperatePermissions sets OperatePermissions field to given value.

### HasOperatePermissions

`func (o *RemoteProcessGroupPortEntity) HasOperatePermissions() bool`

HasOperatePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


