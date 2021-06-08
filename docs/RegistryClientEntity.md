# RegistryClientEntity

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
**Component** | Pointer to [**RegistryDTO**](RegistryDTO.md) |  | [optional] 

## Methods

### NewRegistryClientEntity

`func NewRegistryClientEntity() *RegistryClientEntity`

NewRegistryClientEntity instantiates a new RegistryClientEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryClientEntityWithDefaults

`func NewRegistryClientEntityWithDefaults() *RegistryClientEntity`

NewRegistryClientEntityWithDefaults instantiates a new RegistryClientEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *RegistryClientEntity) GetRevision() RevisionDTO`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *RegistryClientEntity) GetRevisionOk() (*RevisionDTO, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *RegistryClientEntity) SetRevision(v RevisionDTO)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *RegistryClientEntity) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetId

`func (o *RegistryClientEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegistryClientEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegistryClientEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RegistryClientEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUri

`func (o *RegistryClientEntity) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *RegistryClientEntity) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *RegistryClientEntity) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *RegistryClientEntity) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetPosition

`func (o *RegistryClientEntity) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *RegistryClientEntity) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *RegistryClientEntity) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *RegistryClientEntity) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPermissions

`func (o *RegistryClientEntity) GetPermissions() PermissionsDTO`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RegistryClientEntity) GetPermissionsOk() (*PermissionsDTO, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RegistryClientEntity) SetPermissions(v PermissionsDTO)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *RegistryClientEntity) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetBulletins

`func (o *RegistryClientEntity) GetBulletins() []BulletinEntity`

GetBulletins returns the Bulletins field if non-nil, zero value otherwise.

### GetBulletinsOk

`func (o *RegistryClientEntity) GetBulletinsOk() (*[]BulletinEntity, bool)`

GetBulletinsOk returns a tuple with the Bulletins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletins

`func (o *RegistryClientEntity) SetBulletins(v []BulletinEntity)`

SetBulletins sets Bulletins field to given value.

### HasBulletins

`func (o *RegistryClientEntity) HasBulletins() bool`

HasBulletins returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *RegistryClientEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *RegistryClientEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *RegistryClientEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *RegistryClientEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.

### GetComponent

`func (o *RegistryClientEntity) GetComponent() RegistryDTO`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *RegistryClientEntity) GetComponentOk() (*RegistryDTO, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *RegistryClientEntity) SetComponent(v RegistryDTO)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *RegistryClientEntity) HasComponent() bool`

HasComponent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


