# ResourcesEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to [**[]ResourceDTO**](ResourceDTO.md) |  | [optional] 

## Methods

### NewResourcesEntity

`func NewResourcesEntity() *ResourcesEntity`

NewResourcesEntity instantiates a new ResourcesEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesEntityWithDefaults

`func NewResourcesEntityWithDefaults() *ResourcesEntity`

NewResourcesEntityWithDefaults instantiates a new ResourcesEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *ResourcesEntity) GetResources() []ResourceDTO`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ResourcesEntity) GetResourcesOk() (*[]ResourceDTO, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ResourcesEntity) SetResources(v []ResourceDTO)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ResourcesEntity) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


