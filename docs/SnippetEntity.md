# SnippetEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snippet** | Pointer to [**SnippetDTO**](SnippetDTO.md) |  | [optional] 
**DisconnectedNodeAcknowledged** | Pointer to **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] 

## Methods

### NewSnippetEntity

`func NewSnippetEntity() *SnippetEntity`

NewSnippetEntity instantiates a new SnippetEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnippetEntityWithDefaults

`func NewSnippetEntityWithDefaults() *SnippetEntity`

NewSnippetEntityWithDefaults instantiates a new SnippetEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnippet

`func (o *SnippetEntity) GetSnippet() SnippetDTO`

GetSnippet returns the Snippet field if non-nil, zero value otherwise.

### GetSnippetOk

`func (o *SnippetEntity) GetSnippetOk() (*SnippetDTO, bool)`

GetSnippetOk returns a tuple with the Snippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippet

`func (o *SnippetEntity) SetSnippet(v SnippetDTO)`

SetSnippet sets Snippet field to given value.

### HasSnippet

`func (o *SnippetEntity) HasSnippet() bool`

HasSnippet returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *SnippetEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *SnippetEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *SnippetEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *SnippetEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


