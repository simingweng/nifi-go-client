# ResourceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | The identifier of the resource. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 

## Methods

### NewResourceDTO

`func NewResourceDTO() *ResourceDTO`

NewResourceDTO instantiates a new ResourceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceDTOWithDefaults

`func NewResourceDTOWithDefaults() *ResourceDTO`

NewResourceDTOWithDefaults instantiates a new ResourceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *ResourceDTO) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ResourceDTO) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ResourceDTO) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *ResourceDTO) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *ResourceDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceDTO) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


