# ProcessorEntity

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
**Component** | Pointer to [**ProcessorDTO**](ProcessorDTO.md) |  | [optional] 
**InputRequirement** | Pointer to **string** | The input requirement for this processor. | [optional] 
**Status** | Pointer to [**ProcessorStatusDTO**](ProcessorStatusDTO.md) |  | [optional] 
**OperatePermissions** | Pointer to [**PermissionsDTO**](PermissionsDTO.md) |  | [optional] 

## Methods

### NewProcessorEntity

`func NewProcessorEntity() *ProcessorEntity`

NewProcessorEntity instantiates a new ProcessorEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorEntityWithDefaults

`func NewProcessorEntityWithDefaults() *ProcessorEntity`

NewProcessorEntityWithDefaults instantiates a new ProcessorEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *ProcessorEntity) GetRevision() RevisionDTO`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ProcessorEntity) GetRevisionOk() (*RevisionDTO, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ProcessorEntity) SetRevision(v RevisionDTO)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ProcessorEntity) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetId

`func (o *ProcessorEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessorEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessorEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProcessorEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUri

`func (o *ProcessorEntity) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ProcessorEntity) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ProcessorEntity) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ProcessorEntity) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetPosition

`func (o *ProcessorEntity) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ProcessorEntity) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ProcessorEntity) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ProcessorEntity) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPermissions

`func (o *ProcessorEntity) GetPermissions() PermissionsDTO`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ProcessorEntity) GetPermissionsOk() (*PermissionsDTO, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ProcessorEntity) SetPermissions(v PermissionsDTO)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ProcessorEntity) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetBulletins

`func (o *ProcessorEntity) GetBulletins() []BulletinEntity`

GetBulletins returns the Bulletins field if non-nil, zero value otherwise.

### GetBulletinsOk

`func (o *ProcessorEntity) GetBulletinsOk() (*[]BulletinEntity, bool)`

GetBulletinsOk returns a tuple with the Bulletins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletins

`func (o *ProcessorEntity) SetBulletins(v []BulletinEntity)`

SetBulletins sets Bulletins field to given value.

### HasBulletins

`func (o *ProcessorEntity) HasBulletins() bool`

HasBulletins returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *ProcessorEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *ProcessorEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *ProcessorEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *ProcessorEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.

### GetComponent

`func (o *ProcessorEntity) GetComponent() ProcessorDTO`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ProcessorEntity) GetComponentOk() (*ProcessorDTO, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ProcessorEntity) SetComponent(v ProcessorDTO)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ProcessorEntity) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetInputRequirement

`func (o *ProcessorEntity) GetInputRequirement() string`

GetInputRequirement returns the InputRequirement field if non-nil, zero value otherwise.

### GetInputRequirementOk

`func (o *ProcessorEntity) GetInputRequirementOk() (*string, bool)`

GetInputRequirementOk returns a tuple with the InputRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputRequirement

`func (o *ProcessorEntity) SetInputRequirement(v string)`

SetInputRequirement sets InputRequirement field to given value.

### HasInputRequirement

`func (o *ProcessorEntity) HasInputRequirement() bool`

HasInputRequirement returns a boolean if a field has been set.

### GetStatus

`func (o *ProcessorEntity) GetStatus() ProcessorStatusDTO`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProcessorEntity) GetStatusOk() (*ProcessorStatusDTO, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProcessorEntity) SetStatus(v ProcessorStatusDTO)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProcessorEntity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOperatePermissions

`func (o *ProcessorEntity) GetOperatePermissions() PermissionsDTO`

GetOperatePermissions returns the OperatePermissions field if non-nil, zero value otherwise.

### GetOperatePermissionsOk

`func (o *ProcessorEntity) GetOperatePermissionsOk() (*PermissionsDTO, bool)`

GetOperatePermissionsOk returns a tuple with the OperatePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatePermissions

`func (o *ProcessorEntity) SetOperatePermissions(v PermissionsDTO)`

SetOperatePermissions sets OperatePermissions field to given value.

### HasOperatePermissions

`func (o *ProcessorEntity) HasOperatePermissions() bool`

HasOperatePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


