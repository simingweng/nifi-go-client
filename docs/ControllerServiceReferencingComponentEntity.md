# ControllerServiceReferencingComponentEntity

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
**Component** | Pointer to [**ControllerServiceReferencingComponentDTO**](ControllerServiceReferencingComponentDTO.md) |  | [optional] 
**OperatePermissions** | Pointer to [**PermissionsDTO**](PermissionsDTO.md) |  | [optional] 

## Methods

### NewControllerServiceReferencingComponentEntity

`func NewControllerServiceReferencingComponentEntity() *ControllerServiceReferencingComponentEntity`

NewControllerServiceReferencingComponentEntity instantiates a new ControllerServiceReferencingComponentEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerServiceReferencingComponentEntityWithDefaults

`func NewControllerServiceReferencingComponentEntityWithDefaults() *ControllerServiceReferencingComponentEntity`

NewControllerServiceReferencingComponentEntityWithDefaults instantiates a new ControllerServiceReferencingComponentEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *ControllerServiceReferencingComponentEntity) GetRevision() RevisionDTO`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ControllerServiceReferencingComponentEntity) GetRevisionOk() (*RevisionDTO, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ControllerServiceReferencingComponentEntity) SetRevision(v RevisionDTO)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ControllerServiceReferencingComponentEntity) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetId

`func (o *ControllerServiceReferencingComponentEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControllerServiceReferencingComponentEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControllerServiceReferencingComponentEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ControllerServiceReferencingComponentEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUri

`func (o *ControllerServiceReferencingComponentEntity) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ControllerServiceReferencingComponentEntity) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ControllerServiceReferencingComponentEntity) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ControllerServiceReferencingComponentEntity) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetPosition

`func (o *ControllerServiceReferencingComponentEntity) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ControllerServiceReferencingComponentEntity) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ControllerServiceReferencingComponentEntity) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ControllerServiceReferencingComponentEntity) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPermissions

`func (o *ControllerServiceReferencingComponentEntity) GetPermissions() PermissionsDTO`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ControllerServiceReferencingComponentEntity) GetPermissionsOk() (*PermissionsDTO, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ControllerServiceReferencingComponentEntity) SetPermissions(v PermissionsDTO)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ControllerServiceReferencingComponentEntity) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetBulletins

`func (o *ControllerServiceReferencingComponentEntity) GetBulletins() []BulletinEntity`

GetBulletins returns the Bulletins field if non-nil, zero value otherwise.

### GetBulletinsOk

`func (o *ControllerServiceReferencingComponentEntity) GetBulletinsOk() (*[]BulletinEntity, bool)`

GetBulletinsOk returns a tuple with the Bulletins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletins

`func (o *ControllerServiceReferencingComponentEntity) SetBulletins(v []BulletinEntity)`

SetBulletins sets Bulletins field to given value.

### HasBulletins

`func (o *ControllerServiceReferencingComponentEntity) HasBulletins() bool`

HasBulletins returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *ControllerServiceReferencingComponentEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *ControllerServiceReferencingComponentEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *ControllerServiceReferencingComponentEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *ControllerServiceReferencingComponentEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.

### GetComponent

`func (o *ControllerServiceReferencingComponentEntity) GetComponent() ControllerServiceReferencingComponentDTO`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ControllerServiceReferencingComponentEntity) GetComponentOk() (*ControllerServiceReferencingComponentDTO, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ControllerServiceReferencingComponentEntity) SetComponent(v ControllerServiceReferencingComponentDTO)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ControllerServiceReferencingComponentEntity) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetOperatePermissions

`func (o *ControllerServiceReferencingComponentEntity) GetOperatePermissions() PermissionsDTO`

GetOperatePermissions returns the OperatePermissions field if non-nil, zero value otherwise.

### GetOperatePermissionsOk

`func (o *ControllerServiceReferencingComponentEntity) GetOperatePermissionsOk() (*PermissionsDTO, bool)`

GetOperatePermissionsOk returns a tuple with the OperatePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatePermissions

`func (o *ControllerServiceReferencingComponentEntity) SetOperatePermissions(v PermissionsDTO)`

SetOperatePermissions sets OperatePermissions field to given value.

### HasOperatePermissions

`func (o *ControllerServiceReferencingComponentEntity) HasOperatePermissions() bool`

HasOperatePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


