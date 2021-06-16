# CopySnippetRequestEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnippetId** | Pointer to **string** | The identifier of the snippet. | [optional] 
**OriginX** | Pointer to **float64** | The x coordinate of the origin of the bounding box where the new components will be placed. | [optional] 
**OriginY** | Pointer to **float64** | The y coordinate of the origin of the bounding box where the new components will be placed. | [optional] 
**DisconnectedNodeAcknowledged** | Pointer to **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] 

## Methods

### NewCopySnippetRequestEntity

`func NewCopySnippetRequestEntity() *CopySnippetRequestEntity`

NewCopySnippetRequestEntity instantiates a new CopySnippetRequestEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopySnippetRequestEntityWithDefaults

`func NewCopySnippetRequestEntityWithDefaults() *CopySnippetRequestEntity`

NewCopySnippetRequestEntityWithDefaults instantiates a new CopySnippetRequestEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnippetId

`func (o *CopySnippetRequestEntity) GetSnippetId() string`

GetSnippetId returns the SnippetId field if non-nil, zero value otherwise.

### GetSnippetIdOk

`func (o *CopySnippetRequestEntity) GetSnippetIdOk() (*string, bool)`

GetSnippetIdOk returns a tuple with the SnippetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippetId

`func (o *CopySnippetRequestEntity) SetSnippetId(v string)`

SetSnippetId sets SnippetId field to given value.

### HasSnippetId

`func (o *CopySnippetRequestEntity) HasSnippetId() bool`

HasSnippetId returns a boolean if a field has been set.

### GetOriginX

`func (o *CopySnippetRequestEntity) GetOriginX() float64`

GetOriginX returns the OriginX field if non-nil, zero value otherwise.

### GetOriginXOk

`func (o *CopySnippetRequestEntity) GetOriginXOk() (*float64, bool)`

GetOriginXOk returns a tuple with the OriginX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginX

`func (o *CopySnippetRequestEntity) SetOriginX(v float64)`

SetOriginX sets OriginX field to given value.

### HasOriginX

`func (o *CopySnippetRequestEntity) HasOriginX() bool`

HasOriginX returns a boolean if a field has been set.

### GetOriginY

`func (o *CopySnippetRequestEntity) GetOriginY() float64`

GetOriginY returns the OriginY field if non-nil, zero value otherwise.

### GetOriginYOk

`func (o *CopySnippetRequestEntity) GetOriginYOk() (*float64, bool)`

GetOriginYOk returns a tuple with the OriginY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginY

`func (o *CopySnippetRequestEntity) SetOriginY(v float64)`

SetOriginY sets OriginY field to given value.

### HasOriginY

`func (o *CopySnippetRequestEntity) HasOriginY() bool`

HasOriginY returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *CopySnippetRequestEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *CopySnippetRequestEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *CopySnippetRequestEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *CopySnippetRequestEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


