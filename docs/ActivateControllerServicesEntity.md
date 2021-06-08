# ActivateControllerServicesEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the ProcessGroup | [optional] 
**State** | Pointer to **string** | The desired state of the descendant components | [optional] 
**Components** | Pointer to [**map[string]RevisionDTO**](RevisionDTO.md) | Optional services to schedule. If not specified, all authorized descendant controller services will be used. | [optional] 
**DisconnectedNodeAcknowledged** | Pointer to **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] 

## Methods

### NewActivateControllerServicesEntity

`func NewActivateControllerServicesEntity() *ActivateControllerServicesEntity`

NewActivateControllerServicesEntity instantiates a new ActivateControllerServicesEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivateControllerServicesEntityWithDefaults

`func NewActivateControllerServicesEntityWithDefaults() *ActivateControllerServicesEntity`

NewActivateControllerServicesEntityWithDefaults instantiates a new ActivateControllerServicesEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivateControllerServicesEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivateControllerServicesEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivateControllerServicesEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActivateControllerServicesEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *ActivateControllerServicesEntity) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ActivateControllerServicesEntity) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ActivateControllerServicesEntity) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ActivateControllerServicesEntity) HasState() bool`

HasState returns a boolean if a field has been set.

### GetComponents

`func (o *ActivateControllerServicesEntity) GetComponents() map[string]RevisionDTO`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ActivateControllerServicesEntity) GetComponentsOk() (*map[string]RevisionDTO, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ActivateControllerServicesEntity) SetComponents(v map[string]RevisionDTO)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *ActivateControllerServicesEntity) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *ActivateControllerServicesEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *ActivateControllerServicesEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *ActivateControllerServicesEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *ActivateControllerServicesEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


