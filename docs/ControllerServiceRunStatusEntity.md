# ControllerServiceRunStatusEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | Pointer to [**RevisionDTO**](RevisionDTO.md) |  | [optional] 
**State** | Pointer to **string** | The run status of the ControllerService. | [optional] 
**DisconnectedNodeAcknowledged** | Pointer to **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] 

## Methods

### NewControllerServiceRunStatusEntity

`func NewControllerServiceRunStatusEntity() *ControllerServiceRunStatusEntity`

NewControllerServiceRunStatusEntity instantiates a new ControllerServiceRunStatusEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerServiceRunStatusEntityWithDefaults

`func NewControllerServiceRunStatusEntityWithDefaults() *ControllerServiceRunStatusEntity`

NewControllerServiceRunStatusEntityWithDefaults instantiates a new ControllerServiceRunStatusEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *ControllerServiceRunStatusEntity) GetRevision() RevisionDTO`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ControllerServiceRunStatusEntity) GetRevisionOk() (*RevisionDTO, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ControllerServiceRunStatusEntity) SetRevision(v RevisionDTO)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ControllerServiceRunStatusEntity) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetState

`func (o *ControllerServiceRunStatusEntity) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ControllerServiceRunStatusEntity) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ControllerServiceRunStatusEntity) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ControllerServiceRunStatusEntity) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *ControllerServiceRunStatusEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *ControllerServiceRunStatusEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *ControllerServiceRunStatusEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *ControllerServiceRunStatusEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


