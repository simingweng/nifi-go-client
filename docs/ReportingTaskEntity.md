# ReportingTaskEntity

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
**Component** | Pointer to [**ReportingTaskDTO**](ReportingTaskDTO.md) |  | [optional] 
**OperatePermissions** | Pointer to [**PermissionsDTO**](PermissionsDTO.md) |  | [optional] 
**Status** | Pointer to [**ReportingTaskStatusDTO**](ReportingTaskStatusDTO.md) |  | [optional] 

## Methods

### NewReportingTaskEntity

`func NewReportingTaskEntity() *ReportingTaskEntity`

NewReportingTaskEntity instantiates a new ReportingTaskEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportingTaskEntityWithDefaults

`func NewReportingTaskEntityWithDefaults() *ReportingTaskEntity`

NewReportingTaskEntityWithDefaults instantiates a new ReportingTaskEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *ReportingTaskEntity) GetRevision() RevisionDTO`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ReportingTaskEntity) GetRevisionOk() (*RevisionDTO, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ReportingTaskEntity) SetRevision(v RevisionDTO)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ReportingTaskEntity) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetId

`func (o *ReportingTaskEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportingTaskEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportingTaskEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReportingTaskEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUri

`func (o *ReportingTaskEntity) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ReportingTaskEntity) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ReportingTaskEntity) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ReportingTaskEntity) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetPosition

`func (o *ReportingTaskEntity) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ReportingTaskEntity) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ReportingTaskEntity) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ReportingTaskEntity) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPermissions

`func (o *ReportingTaskEntity) GetPermissions() PermissionsDTO`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ReportingTaskEntity) GetPermissionsOk() (*PermissionsDTO, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ReportingTaskEntity) SetPermissions(v PermissionsDTO)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ReportingTaskEntity) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetBulletins

`func (o *ReportingTaskEntity) GetBulletins() []BulletinEntity`

GetBulletins returns the Bulletins field if non-nil, zero value otherwise.

### GetBulletinsOk

`func (o *ReportingTaskEntity) GetBulletinsOk() (*[]BulletinEntity, bool)`

GetBulletinsOk returns a tuple with the Bulletins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletins

`func (o *ReportingTaskEntity) SetBulletins(v []BulletinEntity)`

SetBulletins sets Bulletins field to given value.

### HasBulletins

`func (o *ReportingTaskEntity) HasBulletins() bool`

HasBulletins returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *ReportingTaskEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *ReportingTaskEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *ReportingTaskEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *ReportingTaskEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.

### GetComponent

`func (o *ReportingTaskEntity) GetComponent() ReportingTaskDTO`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ReportingTaskEntity) GetComponentOk() (*ReportingTaskDTO, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ReportingTaskEntity) SetComponent(v ReportingTaskDTO)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ReportingTaskEntity) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetOperatePermissions

`func (o *ReportingTaskEntity) GetOperatePermissions() PermissionsDTO`

GetOperatePermissions returns the OperatePermissions field if non-nil, zero value otherwise.

### GetOperatePermissionsOk

`func (o *ReportingTaskEntity) GetOperatePermissionsOk() (*PermissionsDTO, bool)`

GetOperatePermissionsOk returns a tuple with the OperatePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatePermissions

`func (o *ReportingTaskEntity) SetOperatePermissions(v PermissionsDTO)`

SetOperatePermissions sets OperatePermissions field to given value.

### HasOperatePermissions

`func (o *ReportingTaskEntity) HasOperatePermissions() bool`

HasOperatePermissions returns a boolean if a field has been set.

### GetStatus

`func (o *ReportingTaskEntity) GetStatus() ReportingTaskStatusDTO`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReportingTaskEntity) GetStatusOk() (*ReportingTaskStatusDTO, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReportingTaskEntity) SetStatus(v ReportingTaskStatusDTO)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReportingTaskEntity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


