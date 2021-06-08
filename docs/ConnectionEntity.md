# ConnectionEntity

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
**Component** | Pointer to [**ConnectionDTO**](ConnectionDTO.md) |  | [optional] 
**Status** | Pointer to [**ConnectionStatusDTO**](ConnectionStatusDTO.md) |  | [optional] 
**Bends** | Pointer to [**[]PositionDTO**](PositionDTO.md) | The bend points on the connection. | [optional] 
**LabelIndex** | Pointer to **int32** | The index of the bend point where to place the connection label. | [optional] 
**GetzIndex** | Pointer to **int64** | The z index of the connection. | [optional] 
**SourceId** | Pointer to **string** | The identifier of the source of this connection. | [optional] 
**SourceGroupId** | Pointer to **string** | The identifier of the group of the source of this connection. | [optional] 
**SourceType** | **string** | The type of component the source connectable is. | 
**DestinationId** | Pointer to **string** | The identifier of the destination of this connection. | [optional] 
**DestinationGroupId** | Pointer to **string** | The identifier of the group of the destination of this connection. | [optional] 
**DestinationType** | **string** | The type of component the destination connectable is. | 

## Methods

### NewConnectionEntity

`func NewConnectionEntity(sourceType string, destinationType string, ) *ConnectionEntity`

NewConnectionEntity instantiates a new ConnectionEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionEntityWithDefaults

`func NewConnectionEntityWithDefaults() *ConnectionEntity`

NewConnectionEntityWithDefaults instantiates a new ConnectionEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *ConnectionEntity) GetRevision() RevisionDTO`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ConnectionEntity) GetRevisionOk() (*RevisionDTO, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ConnectionEntity) SetRevision(v RevisionDTO)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ConnectionEntity) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetId

`func (o *ConnectionEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectionEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectionEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectionEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUri

`func (o *ConnectionEntity) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ConnectionEntity) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ConnectionEntity) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ConnectionEntity) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetPosition

`func (o *ConnectionEntity) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ConnectionEntity) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ConnectionEntity) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ConnectionEntity) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPermissions

`func (o *ConnectionEntity) GetPermissions() PermissionsDTO`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ConnectionEntity) GetPermissionsOk() (*PermissionsDTO, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ConnectionEntity) SetPermissions(v PermissionsDTO)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ConnectionEntity) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetBulletins

`func (o *ConnectionEntity) GetBulletins() []BulletinEntity`

GetBulletins returns the Bulletins field if non-nil, zero value otherwise.

### GetBulletinsOk

`func (o *ConnectionEntity) GetBulletinsOk() (*[]BulletinEntity, bool)`

GetBulletinsOk returns a tuple with the Bulletins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletins

`func (o *ConnectionEntity) SetBulletins(v []BulletinEntity)`

SetBulletins sets Bulletins field to given value.

### HasBulletins

`func (o *ConnectionEntity) HasBulletins() bool`

HasBulletins returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *ConnectionEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *ConnectionEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *ConnectionEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *ConnectionEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.

### GetComponent

`func (o *ConnectionEntity) GetComponent() ConnectionDTO`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ConnectionEntity) GetComponentOk() (*ConnectionDTO, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ConnectionEntity) SetComponent(v ConnectionDTO)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ConnectionEntity) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetStatus

`func (o *ConnectionEntity) GetStatus() ConnectionStatusDTO`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectionEntity) GetStatusOk() (*ConnectionStatusDTO, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectionEntity) SetStatus(v ConnectionStatusDTO)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConnectionEntity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBends

`func (o *ConnectionEntity) GetBends() []PositionDTO`

GetBends returns the Bends field if non-nil, zero value otherwise.

### GetBendsOk

`func (o *ConnectionEntity) GetBendsOk() (*[]PositionDTO, bool)`

GetBendsOk returns a tuple with the Bends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBends

`func (o *ConnectionEntity) SetBends(v []PositionDTO)`

SetBends sets Bends field to given value.

### HasBends

`func (o *ConnectionEntity) HasBends() bool`

HasBends returns a boolean if a field has been set.

### GetLabelIndex

`func (o *ConnectionEntity) GetLabelIndex() int32`

GetLabelIndex returns the LabelIndex field if non-nil, zero value otherwise.

### GetLabelIndexOk

`func (o *ConnectionEntity) GetLabelIndexOk() (*int32, bool)`

GetLabelIndexOk returns a tuple with the LabelIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIndex

`func (o *ConnectionEntity) SetLabelIndex(v int32)`

SetLabelIndex sets LabelIndex field to given value.

### HasLabelIndex

`func (o *ConnectionEntity) HasLabelIndex() bool`

HasLabelIndex returns a boolean if a field has been set.

### GetGetzIndex

`func (o *ConnectionEntity) GetGetzIndex() int64`

GetGetzIndex returns the GetzIndex field if non-nil, zero value otherwise.

### GetGetzIndexOk

`func (o *ConnectionEntity) GetGetzIndexOk() (*int64, bool)`

GetGetzIndexOk returns a tuple with the GetzIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetzIndex

`func (o *ConnectionEntity) SetGetzIndex(v int64)`

SetGetzIndex sets GetzIndex field to given value.

### HasGetzIndex

`func (o *ConnectionEntity) HasGetzIndex() bool`

HasGetzIndex returns a boolean if a field has been set.

### GetSourceId

`func (o *ConnectionEntity) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ConnectionEntity) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ConnectionEntity) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ConnectionEntity) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceGroupId

`func (o *ConnectionEntity) GetSourceGroupId() string`

GetSourceGroupId returns the SourceGroupId field if non-nil, zero value otherwise.

### GetSourceGroupIdOk

`func (o *ConnectionEntity) GetSourceGroupIdOk() (*string, bool)`

GetSourceGroupIdOk returns a tuple with the SourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroupId

`func (o *ConnectionEntity) SetSourceGroupId(v string)`

SetSourceGroupId sets SourceGroupId field to given value.

### HasSourceGroupId

`func (o *ConnectionEntity) HasSourceGroupId() bool`

HasSourceGroupId returns a boolean if a field has been set.

### GetSourceType

`func (o *ConnectionEntity) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ConnectionEntity) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ConnectionEntity) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetDestinationId

`func (o *ConnectionEntity) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *ConnectionEntity) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *ConnectionEntity) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.

### HasDestinationId

`func (o *ConnectionEntity) HasDestinationId() bool`

HasDestinationId returns a boolean if a field has been set.

### GetDestinationGroupId

`func (o *ConnectionEntity) GetDestinationGroupId() string`

GetDestinationGroupId returns the DestinationGroupId field if non-nil, zero value otherwise.

### GetDestinationGroupIdOk

`func (o *ConnectionEntity) GetDestinationGroupIdOk() (*string, bool)`

GetDestinationGroupIdOk returns a tuple with the DestinationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGroupId

`func (o *ConnectionEntity) SetDestinationGroupId(v string)`

SetDestinationGroupId sets DestinationGroupId field to given value.

### HasDestinationGroupId

`func (o *ConnectionEntity) HasDestinationGroupId() bool`

HasDestinationGroupId returns a boolean if a field has been set.

### GetDestinationType

`func (o *ConnectionEntity) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *ConnectionEntity) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *ConnectionEntity) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


