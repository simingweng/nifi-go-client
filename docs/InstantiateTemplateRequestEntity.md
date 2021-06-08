# InstantiateTemplateRequestEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginX** | Pointer to **float64** | The x coordinate of the origin of the bounding box where the new components will be placed. | [optional] 
**OriginY** | Pointer to **float64** | The y coordinate of the origin of the bounding box where the new components will be placed. | [optional] 
**TemplateId** | Pointer to **string** | The identifier of the template. | [optional] 
**EncodingVersion** | Pointer to **string** | The encoding version of the flow snippet. If not specified, this is automatically populated by the node receiving the user request. If the snippet is specified, the version will be the latest. If the snippet is not specified, the version will come from the underlying template. These details need to be replicated throughout the cluster to ensure consistency. | [optional] 
**Snippet** | Pointer to [**FlowSnippetDTO**](FlowSnippetDTO.md) |  | [optional] 
**DisconnectedNodeAcknowledged** | Pointer to **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] 

## Methods

### NewInstantiateTemplateRequestEntity

`func NewInstantiateTemplateRequestEntity() *InstantiateTemplateRequestEntity`

NewInstantiateTemplateRequestEntity instantiates a new InstantiateTemplateRequestEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstantiateTemplateRequestEntityWithDefaults

`func NewInstantiateTemplateRequestEntityWithDefaults() *InstantiateTemplateRequestEntity`

NewInstantiateTemplateRequestEntityWithDefaults instantiates a new InstantiateTemplateRequestEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginX

`func (o *InstantiateTemplateRequestEntity) GetOriginX() float64`

GetOriginX returns the OriginX field if non-nil, zero value otherwise.

### GetOriginXOk

`func (o *InstantiateTemplateRequestEntity) GetOriginXOk() (*float64, bool)`

GetOriginXOk returns a tuple with the OriginX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginX

`func (o *InstantiateTemplateRequestEntity) SetOriginX(v float64)`

SetOriginX sets OriginX field to given value.

### HasOriginX

`func (o *InstantiateTemplateRequestEntity) HasOriginX() bool`

HasOriginX returns a boolean if a field has been set.

### GetOriginY

`func (o *InstantiateTemplateRequestEntity) GetOriginY() float64`

GetOriginY returns the OriginY field if non-nil, zero value otherwise.

### GetOriginYOk

`func (o *InstantiateTemplateRequestEntity) GetOriginYOk() (*float64, bool)`

GetOriginYOk returns a tuple with the OriginY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginY

`func (o *InstantiateTemplateRequestEntity) SetOriginY(v float64)`

SetOriginY sets OriginY field to given value.

### HasOriginY

`func (o *InstantiateTemplateRequestEntity) HasOriginY() bool`

HasOriginY returns a boolean if a field has been set.

### GetTemplateId

`func (o *InstantiateTemplateRequestEntity) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *InstantiateTemplateRequestEntity) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *InstantiateTemplateRequestEntity) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *InstantiateTemplateRequestEntity) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetEncodingVersion

`func (o *InstantiateTemplateRequestEntity) GetEncodingVersion() string`

GetEncodingVersion returns the EncodingVersion field if non-nil, zero value otherwise.

### GetEncodingVersionOk

`func (o *InstantiateTemplateRequestEntity) GetEncodingVersionOk() (*string, bool)`

GetEncodingVersionOk returns a tuple with the EncodingVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingVersion

`func (o *InstantiateTemplateRequestEntity) SetEncodingVersion(v string)`

SetEncodingVersion sets EncodingVersion field to given value.

### HasEncodingVersion

`func (o *InstantiateTemplateRequestEntity) HasEncodingVersion() bool`

HasEncodingVersion returns a boolean if a field has been set.

### GetSnippet

`func (o *InstantiateTemplateRequestEntity) GetSnippet() FlowSnippetDTO`

GetSnippet returns the Snippet field if non-nil, zero value otherwise.

### GetSnippetOk

`func (o *InstantiateTemplateRequestEntity) GetSnippetOk() (*FlowSnippetDTO, bool)`

GetSnippetOk returns a tuple with the Snippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippet

`func (o *InstantiateTemplateRequestEntity) SetSnippet(v FlowSnippetDTO)`

SetSnippet sets Snippet field to given value.

### HasSnippet

`func (o *InstantiateTemplateRequestEntity) HasSnippet() bool`

HasSnippet returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *InstantiateTemplateRequestEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *InstantiateTemplateRequestEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *InstantiateTemplateRequestEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *InstantiateTemplateRequestEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


