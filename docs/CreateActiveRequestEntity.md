# CreateActiveRequestEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessGroupId** | Pointer to **string** | The Process Group ID that this active request will update | [optional] 
**DisconnectedNodeAcknowledged** | Pointer to **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] 

## Methods

### NewCreateActiveRequestEntity

`func NewCreateActiveRequestEntity() *CreateActiveRequestEntity`

NewCreateActiveRequestEntity instantiates a new CreateActiveRequestEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateActiveRequestEntityWithDefaults

`func NewCreateActiveRequestEntityWithDefaults() *CreateActiveRequestEntity`

NewCreateActiveRequestEntityWithDefaults instantiates a new CreateActiveRequestEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessGroupId

`func (o *CreateActiveRequestEntity) GetProcessGroupId() string`

GetProcessGroupId returns the ProcessGroupId field if non-nil, zero value otherwise.

### GetProcessGroupIdOk

`func (o *CreateActiveRequestEntity) GetProcessGroupIdOk() (*string, bool)`

GetProcessGroupIdOk returns a tuple with the ProcessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupId

`func (o *CreateActiveRequestEntity) SetProcessGroupId(v string)`

SetProcessGroupId sets ProcessGroupId field to given value.

### HasProcessGroupId

`func (o *CreateActiveRequestEntity) HasProcessGroupId() bool`

HasProcessGroupId returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *CreateActiveRequestEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *CreateActiveRequestEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *CreateActiveRequestEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *CreateActiveRequestEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


