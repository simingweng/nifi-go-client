# ParameterContextEntity

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
**Component** | Pointer to [**ParameterContextDTO**](ParameterContextDTO.md) |  | [optional] 

## Methods

### NewParameterContextEntity

`func NewParameterContextEntity() *ParameterContextEntity`

NewParameterContextEntity instantiates a new ParameterContextEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterContextEntityWithDefaults

`func NewParameterContextEntityWithDefaults() *ParameterContextEntity`

NewParameterContextEntityWithDefaults instantiates a new ParameterContextEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *ParameterContextEntity) GetRevision() RevisionDTO`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ParameterContextEntity) GetRevisionOk() (*RevisionDTO, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ParameterContextEntity) SetRevision(v RevisionDTO)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ParameterContextEntity) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetId

`func (o *ParameterContextEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParameterContextEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParameterContextEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ParameterContextEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUri

`func (o *ParameterContextEntity) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ParameterContextEntity) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ParameterContextEntity) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ParameterContextEntity) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetPosition

`func (o *ParameterContextEntity) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ParameterContextEntity) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ParameterContextEntity) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ParameterContextEntity) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPermissions

`func (o *ParameterContextEntity) GetPermissions() PermissionsDTO`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ParameterContextEntity) GetPermissionsOk() (*PermissionsDTO, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ParameterContextEntity) SetPermissions(v PermissionsDTO)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ParameterContextEntity) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetBulletins

`func (o *ParameterContextEntity) GetBulletins() []BulletinEntity`

GetBulletins returns the Bulletins field if non-nil, zero value otherwise.

### GetBulletinsOk

`func (o *ParameterContextEntity) GetBulletinsOk() (*[]BulletinEntity, bool)`

GetBulletinsOk returns a tuple with the Bulletins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletins

`func (o *ParameterContextEntity) SetBulletins(v []BulletinEntity)`

SetBulletins sets Bulletins field to given value.

### HasBulletins

`func (o *ParameterContextEntity) HasBulletins() bool`

HasBulletins returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *ParameterContextEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *ParameterContextEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *ParameterContextEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *ParameterContextEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.

### GetComponent

`func (o *ParameterContextEntity) GetComponent() ParameterContextDTO`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ParameterContextEntity) GetComponentOk() (*ParameterContextDTO, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ParameterContextEntity) SetComponent(v ParameterContextDTO)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ParameterContextEntity) HasComponent() bool`

HasComponent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


