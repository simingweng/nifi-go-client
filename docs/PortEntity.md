# PortEntity

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
**Component** | Pointer to [**PortDTO**](PortDTO.md) |  | [optional] 
**Status** | Pointer to [**PortStatusDTO**](PortStatusDTO.md) |  | [optional] 
**PortType** | Pointer to **string** |  | [optional] 
**OperatePermissions** | Pointer to [**PermissionsDTO**](PermissionsDTO.md) |  | [optional] 
**AllowRemoteAccess** | Pointer to **bool** | Whether this port can be accessed remotely via Site-to-Site protocol. | [optional] 

## Methods

### NewPortEntity

`func NewPortEntity() *PortEntity`

NewPortEntity instantiates a new PortEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortEntityWithDefaults

`func NewPortEntityWithDefaults() *PortEntity`

NewPortEntityWithDefaults instantiates a new PortEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *PortEntity) GetRevision() RevisionDTO`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *PortEntity) GetRevisionOk() (*RevisionDTO, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *PortEntity) SetRevision(v RevisionDTO)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *PortEntity) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetId

`func (o *PortEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUri

`func (o *PortEntity) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *PortEntity) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *PortEntity) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *PortEntity) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetPosition

`func (o *PortEntity) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PortEntity) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PortEntity) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PortEntity) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPermissions

`func (o *PortEntity) GetPermissions() PermissionsDTO`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PortEntity) GetPermissionsOk() (*PermissionsDTO, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PortEntity) SetPermissions(v PermissionsDTO)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *PortEntity) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetBulletins

`func (o *PortEntity) GetBulletins() []BulletinEntity`

GetBulletins returns the Bulletins field if non-nil, zero value otherwise.

### GetBulletinsOk

`func (o *PortEntity) GetBulletinsOk() (*[]BulletinEntity, bool)`

GetBulletinsOk returns a tuple with the Bulletins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletins

`func (o *PortEntity) SetBulletins(v []BulletinEntity)`

SetBulletins sets Bulletins field to given value.

### HasBulletins

`func (o *PortEntity) HasBulletins() bool`

HasBulletins returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *PortEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *PortEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *PortEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *PortEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.

### GetComponent

`func (o *PortEntity) GetComponent() PortDTO`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *PortEntity) GetComponentOk() (*PortDTO, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *PortEntity) SetComponent(v PortDTO)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *PortEntity) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetStatus

`func (o *PortEntity) GetStatus() PortStatusDTO`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PortEntity) GetStatusOk() (*PortStatusDTO, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PortEntity) SetStatus(v PortStatusDTO)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PortEntity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPortType

`func (o *PortEntity) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *PortEntity) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *PortEntity) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *PortEntity) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### GetOperatePermissions

`func (o *PortEntity) GetOperatePermissions() PermissionsDTO`

GetOperatePermissions returns the OperatePermissions field if non-nil, zero value otherwise.

### GetOperatePermissionsOk

`func (o *PortEntity) GetOperatePermissionsOk() (*PermissionsDTO, bool)`

GetOperatePermissionsOk returns a tuple with the OperatePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatePermissions

`func (o *PortEntity) SetOperatePermissions(v PermissionsDTO)`

SetOperatePermissions sets OperatePermissions field to given value.

### HasOperatePermissions

`func (o *PortEntity) HasOperatePermissions() bool`

HasOperatePermissions returns a boolean if a field has been set.

### GetAllowRemoteAccess

`func (o *PortEntity) GetAllowRemoteAccess() bool`

GetAllowRemoteAccess returns the AllowRemoteAccess field if non-nil, zero value otherwise.

### GetAllowRemoteAccessOk

`func (o *PortEntity) GetAllowRemoteAccessOk() (*bool, bool)`

GetAllowRemoteAccessOk returns a tuple with the AllowRemoteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRemoteAccess

`func (o *PortEntity) SetAllowRemoteAccess(v bool)`

SetAllowRemoteAccess sets AllowRemoteAccess field to given value.

### HasAllowRemoteAccess

`func (o *PortEntity) HasAllowRemoteAccess() bool`

HasAllowRemoteAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


