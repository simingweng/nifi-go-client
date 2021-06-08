# CreateTemplateRequestEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the template. | [optional] 
**Description** | Pointer to **string** | The description of the template. | [optional] 
**SnippetId** | Pointer to **string** | The identifier of the snippet. | [optional] 
**DisconnectedNodeAcknowledged** | Pointer to **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] 

## Methods

### NewCreateTemplateRequestEntity

`func NewCreateTemplateRequestEntity() *CreateTemplateRequestEntity`

NewCreateTemplateRequestEntity instantiates a new CreateTemplateRequestEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTemplateRequestEntityWithDefaults

`func NewCreateTemplateRequestEntityWithDefaults() *CreateTemplateRequestEntity`

NewCreateTemplateRequestEntityWithDefaults instantiates a new CreateTemplateRequestEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTemplateRequestEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTemplateRequestEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTemplateRequestEntity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateTemplateRequestEntity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateTemplateRequestEntity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTemplateRequestEntity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTemplateRequestEntity) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTemplateRequestEntity) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSnippetId

`func (o *CreateTemplateRequestEntity) GetSnippetId() string`

GetSnippetId returns the SnippetId field if non-nil, zero value otherwise.

### GetSnippetIdOk

`func (o *CreateTemplateRequestEntity) GetSnippetIdOk() (*string, bool)`

GetSnippetIdOk returns a tuple with the SnippetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippetId

`func (o *CreateTemplateRequestEntity) SetSnippetId(v string)`

SetSnippetId sets SnippetId field to given value.

### HasSnippetId

`func (o *CreateTemplateRequestEntity) HasSnippetId() bool`

HasSnippetId returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *CreateTemplateRequestEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *CreateTemplateRequestEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *CreateTemplateRequestEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *CreateTemplateRequestEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


