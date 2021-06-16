# AffectedComponentEntity

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
**Component** | Pointer to [**AffectedComponentDTO**](AffectedComponentDTO.md) |  | [optional] 
**ProcessGroup** | Pointer to [**ProcessGroupNameDTO**](ProcessGroupNameDTO.md) |  | [optional] 
**ReferenceType** | Pointer to **string** | The type of component referenced | [optional] 

## Methods

### NewAffectedComponentEntity

`func NewAffectedComponentEntity() *AffectedComponentEntity`

NewAffectedComponentEntity instantiates a new AffectedComponentEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffectedComponentEntityWithDefaults

`func NewAffectedComponentEntityWithDefaults() *AffectedComponentEntity`

NewAffectedComponentEntityWithDefaults instantiates a new AffectedComponentEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *AffectedComponentEntity) GetRevision() RevisionDTO`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *AffectedComponentEntity) GetRevisionOk() (*RevisionDTO, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *AffectedComponentEntity) SetRevision(v RevisionDTO)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *AffectedComponentEntity) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetId

`func (o *AffectedComponentEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AffectedComponentEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AffectedComponentEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AffectedComponentEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUri

`func (o *AffectedComponentEntity) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *AffectedComponentEntity) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *AffectedComponentEntity) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *AffectedComponentEntity) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetPosition

`func (o *AffectedComponentEntity) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *AffectedComponentEntity) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *AffectedComponentEntity) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *AffectedComponentEntity) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPermissions

`func (o *AffectedComponentEntity) GetPermissions() PermissionsDTO`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AffectedComponentEntity) GetPermissionsOk() (*PermissionsDTO, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AffectedComponentEntity) SetPermissions(v PermissionsDTO)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AffectedComponentEntity) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetBulletins

`func (o *AffectedComponentEntity) GetBulletins() []BulletinEntity`

GetBulletins returns the Bulletins field if non-nil, zero value otherwise.

### GetBulletinsOk

`func (o *AffectedComponentEntity) GetBulletinsOk() (*[]BulletinEntity, bool)`

GetBulletinsOk returns a tuple with the Bulletins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletins

`func (o *AffectedComponentEntity) SetBulletins(v []BulletinEntity)`

SetBulletins sets Bulletins field to given value.

### HasBulletins

`func (o *AffectedComponentEntity) HasBulletins() bool`

HasBulletins returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *AffectedComponentEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *AffectedComponentEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *AffectedComponentEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *AffectedComponentEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.

### GetComponent

`func (o *AffectedComponentEntity) GetComponent() AffectedComponentDTO`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AffectedComponentEntity) GetComponentOk() (*AffectedComponentDTO, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AffectedComponentEntity) SetComponent(v AffectedComponentDTO)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *AffectedComponentEntity) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetProcessGroup

`func (o *AffectedComponentEntity) GetProcessGroup() ProcessGroupNameDTO`

GetProcessGroup returns the ProcessGroup field if non-nil, zero value otherwise.

### GetProcessGroupOk

`func (o *AffectedComponentEntity) GetProcessGroupOk() (*ProcessGroupNameDTO, bool)`

GetProcessGroupOk returns a tuple with the ProcessGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroup

`func (o *AffectedComponentEntity) SetProcessGroup(v ProcessGroupNameDTO)`

SetProcessGroup sets ProcessGroup field to given value.

### HasProcessGroup

`func (o *AffectedComponentEntity) HasProcessGroup() bool`

HasProcessGroup returns a boolean if a field has been set.

### GetReferenceType

`func (o *AffectedComponentEntity) GetReferenceType() string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *AffectedComponentEntity) GetReferenceTypeOk() (*string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *AffectedComponentEntity) SetReferenceType(v string)`

SetReferenceType sets ReferenceType field to given value.

### HasReferenceType

`func (o *AffectedComponentEntity) HasReferenceType() bool`

HasReferenceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


