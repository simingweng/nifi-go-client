# ScheduleComponentsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the ProcessGroup | [optional] 
**State** | Pointer to **string** | The desired state of the descendant components | [optional] 
**Components** | Pointer to [**map[string]RevisionDTO**](RevisionDTO.md) | Optional components to schedule. If not specified, all authorized descendant components will be used. | [optional] 
**DisconnectedNodeAcknowledged** | Pointer to **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] 

## Methods

### NewScheduleComponentsEntity

`func NewScheduleComponentsEntity() *ScheduleComponentsEntity`

NewScheduleComponentsEntity instantiates a new ScheduleComponentsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleComponentsEntityWithDefaults

`func NewScheduleComponentsEntityWithDefaults() *ScheduleComponentsEntity`

NewScheduleComponentsEntityWithDefaults instantiates a new ScheduleComponentsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduleComponentsEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleComponentsEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleComponentsEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScheduleComponentsEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *ScheduleComponentsEntity) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ScheduleComponentsEntity) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ScheduleComponentsEntity) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ScheduleComponentsEntity) HasState() bool`

HasState returns a boolean if a field has been set.

### GetComponents

`func (o *ScheduleComponentsEntity) GetComponents() map[string]RevisionDTO`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ScheduleComponentsEntity) GetComponentsOk() (*map[string]RevisionDTO, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ScheduleComponentsEntity) SetComponents(v map[string]RevisionDTO)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *ScheduleComponentsEntity) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *ScheduleComponentsEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *ScheduleComponentsEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *ScheduleComponentsEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *ScheduleComponentsEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


