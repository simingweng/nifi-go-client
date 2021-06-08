# ControllerServiceEntity

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
**ParentGroupId** | Pointer to **string** | The id of parent process group of this ControllerService. | [optional] 
**Component** | Pointer to [**ControllerServiceDTO**](ControllerServiceDTO.md) |  | [optional] 
**OperatePermissions** | Pointer to [**PermissionsDTO**](PermissionsDTO.md) |  | [optional] 
**Status** | Pointer to [**ControllerServiceStatusDTO**](ControllerServiceStatusDTO.md) |  | [optional] 

## Methods

### NewControllerServiceEntity

`func NewControllerServiceEntity() *ControllerServiceEntity`

NewControllerServiceEntity instantiates a new ControllerServiceEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerServiceEntityWithDefaults

`func NewControllerServiceEntityWithDefaults() *ControllerServiceEntity`

NewControllerServiceEntityWithDefaults instantiates a new ControllerServiceEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *ControllerServiceEntity) GetRevision() RevisionDTO`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ControllerServiceEntity) GetRevisionOk() (*RevisionDTO, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ControllerServiceEntity) SetRevision(v RevisionDTO)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ControllerServiceEntity) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetId

`func (o *ControllerServiceEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControllerServiceEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControllerServiceEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ControllerServiceEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUri

`func (o *ControllerServiceEntity) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ControllerServiceEntity) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ControllerServiceEntity) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ControllerServiceEntity) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetPosition

`func (o *ControllerServiceEntity) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ControllerServiceEntity) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ControllerServiceEntity) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ControllerServiceEntity) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPermissions

`func (o *ControllerServiceEntity) GetPermissions() PermissionsDTO`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ControllerServiceEntity) GetPermissionsOk() (*PermissionsDTO, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ControllerServiceEntity) SetPermissions(v PermissionsDTO)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ControllerServiceEntity) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetBulletins

`func (o *ControllerServiceEntity) GetBulletins() []BulletinEntity`

GetBulletins returns the Bulletins field if non-nil, zero value otherwise.

### GetBulletinsOk

`func (o *ControllerServiceEntity) GetBulletinsOk() (*[]BulletinEntity, bool)`

GetBulletinsOk returns a tuple with the Bulletins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletins

`func (o *ControllerServiceEntity) SetBulletins(v []BulletinEntity)`

SetBulletins sets Bulletins field to given value.

### HasBulletins

`func (o *ControllerServiceEntity) HasBulletins() bool`

HasBulletins returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *ControllerServiceEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *ControllerServiceEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *ControllerServiceEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *ControllerServiceEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.

### GetParentGroupId

`func (o *ControllerServiceEntity) GetParentGroupId() string`

GetParentGroupId returns the ParentGroupId field if non-nil, zero value otherwise.

### GetParentGroupIdOk

`func (o *ControllerServiceEntity) GetParentGroupIdOk() (*string, bool)`

GetParentGroupIdOk returns a tuple with the ParentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroupId

`func (o *ControllerServiceEntity) SetParentGroupId(v string)`

SetParentGroupId sets ParentGroupId field to given value.

### HasParentGroupId

`func (o *ControllerServiceEntity) HasParentGroupId() bool`

HasParentGroupId returns a boolean if a field has been set.

### GetComponent

`func (o *ControllerServiceEntity) GetComponent() ControllerServiceDTO`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ControllerServiceEntity) GetComponentOk() (*ControllerServiceDTO, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ControllerServiceEntity) SetComponent(v ControllerServiceDTO)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ControllerServiceEntity) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetOperatePermissions

`func (o *ControllerServiceEntity) GetOperatePermissions() PermissionsDTO`

GetOperatePermissions returns the OperatePermissions field if non-nil, zero value otherwise.

### GetOperatePermissionsOk

`func (o *ControllerServiceEntity) GetOperatePermissionsOk() (*PermissionsDTO, bool)`

GetOperatePermissionsOk returns a tuple with the OperatePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatePermissions

`func (o *ControllerServiceEntity) SetOperatePermissions(v PermissionsDTO)`

SetOperatePermissions sets OperatePermissions field to given value.

### HasOperatePermissions

`func (o *ControllerServiceEntity) HasOperatePermissions() bool`

HasOperatePermissions returns a boolean if a field has been set.

### GetStatus

`func (o *ControllerServiceEntity) GetStatus() ControllerServiceStatusDTO`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ControllerServiceEntity) GetStatusOk() (*ControllerServiceStatusDTO, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ControllerServiceEntity) SetStatus(v ControllerServiceStatusDTO)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ControllerServiceEntity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


