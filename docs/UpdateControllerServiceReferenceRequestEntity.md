# UpdateControllerServiceReferenceRequestEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The identifier of the Controller Service. | [optional] 
**State** | Pointer to **string** | The new state of the references for the controller service. | [optional] 
**ReferencingComponentRevisions** | Pointer to [**map[string]RevisionDTO**](RevisionDTO.md) | The revisions for all referencing components. | [optional] 
**DisconnectedNodeAcknowledged** | Pointer to **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] 

## Methods

### NewUpdateControllerServiceReferenceRequestEntity

`func NewUpdateControllerServiceReferenceRequestEntity() *UpdateControllerServiceReferenceRequestEntity`

NewUpdateControllerServiceReferenceRequestEntity instantiates a new UpdateControllerServiceReferenceRequestEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateControllerServiceReferenceRequestEntityWithDefaults

`func NewUpdateControllerServiceReferenceRequestEntityWithDefaults() *UpdateControllerServiceReferenceRequestEntity`

NewUpdateControllerServiceReferenceRequestEntityWithDefaults instantiates a new UpdateControllerServiceReferenceRequestEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateControllerServiceReferenceRequestEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateControllerServiceReferenceRequestEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateControllerServiceReferenceRequestEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateControllerServiceReferenceRequestEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *UpdateControllerServiceReferenceRequestEntity) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateControllerServiceReferenceRequestEntity) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateControllerServiceReferenceRequestEntity) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UpdateControllerServiceReferenceRequestEntity) HasState() bool`

HasState returns a boolean if a field has been set.

### GetReferencingComponentRevisions

`func (o *UpdateControllerServiceReferenceRequestEntity) GetReferencingComponentRevisions() map[string]RevisionDTO`

GetReferencingComponentRevisions returns the ReferencingComponentRevisions field if non-nil, zero value otherwise.

### GetReferencingComponentRevisionsOk

`func (o *UpdateControllerServiceReferenceRequestEntity) GetReferencingComponentRevisionsOk() (*map[string]RevisionDTO, bool)`

GetReferencingComponentRevisionsOk returns a tuple with the ReferencingComponentRevisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencingComponentRevisions

`func (o *UpdateControllerServiceReferenceRequestEntity) SetReferencingComponentRevisions(v map[string]RevisionDTO)`

SetReferencingComponentRevisions sets ReferencingComponentRevisions field to given value.

### HasReferencingComponentRevisions

`func (o *UpdateControllerServiceReferenceRequestEntity) HasReferencingComponentRevisions() bool`

HasReferencingComponentRevisions returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *UpdateControllerServiceReferenceRequestEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *UpdateControllerServiceReferenceRequestEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *UpdateControllerServiceReferenceRequestEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *UpdateControllerServiceReferenceRequestEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


