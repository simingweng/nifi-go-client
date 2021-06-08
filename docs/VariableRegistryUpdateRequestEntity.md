# VariableRegistryUpdateRequestEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | Pointer to [**VariableRegistryUpdateRequestDTO**](VariableRegistryUpdateRequestDTO.md) |  | [optional] 
**ProcessGroupRevision** | Pointer to [**RevisionDTO**](RevisionDTO.md) |  | [optional] 

## Methods

### NewVariableRegistryUpdateRequestEntity

`func NewVariableRegistryUpdateRequestEntity() *VariableRegistryUpdateRequestEntity`

NewVariableRegistryUpdateRequestEntity instantiates a new VariableRegistryUpdateRequestEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableRegistryUpdateRequestEntityWithDefaults

`func NewVariableRegistryUpdateRequestEntityWithDefaults() *VariableRegistryUpdateRequestEntity`

NewVariableRegistryUpdateRequestEntityWithDefaults instantiates a new VariableRegistryUpdateRequestEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *VariableRegistryUpdateRequestEntity) GetRequest() VariableRegistryUpdateRequestDTO`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *VariableRegistryUpdateRequestEntity) GetRequestOk() (*VariableRegistryUpdateRequestDTO, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *VariableRegistryUpdateRequestEntity) SetRequest(v VariableRegistryUpdateRequestDTO)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *VariableRegistryUpdateRequestEntity) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetProcessGroupRevision

`func (o *VariableRegistryUpdateRequestEntity) GetProcessGroupRevision() RevisionDTO`

GetProcessGroupRevision returns the ProcessGroupRevision field if non-nil, zero value otherwise.

### GetProcessGroupRevisionOk

`func (o *VariableRegistryUpdateRequestEntity) GetProcessGroupRevisionOk() (*RevisionDTO, bool)`

GetProcessGroupRevisionOk returns a tuple with the ProcessGroupRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupRevision

`func (o *VariableRegistryUpdateRequestEntity) SetProcessGroupRevision(v RevisionDTO)`

SetProcessGroupRevision sets ProcessGroupRevision field to given value.

### HasProcessGroupRevision

`func (o *VariableRegistryUpdateRequestEntity) HasProcessGroupRevision() bool`

HasProcessGroupRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


