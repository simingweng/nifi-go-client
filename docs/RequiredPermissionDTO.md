# RequiredPermissionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The required sub-permission necessary for this restriction. | [optional] 
**Label** | Pointer to **string** | The label for the required sub-permission necessary for this restriction. | [optional] 

## Methods

### NewRequiredPermissionDTO

`func NewRequiredPermissionDTO() *RequiredPermissionDTO`

NewRequiredPermissionDTO instantiates a new RequiredPermissionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequiredPermissionDTOWithDefaults

`func NewRequiredPermissionDTOWithDefaults() *RequiredPermissionDTO`

NewRequiredPermissionDTOWithDefaults instantiates a new RequiredPermissionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequiredPermissionDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequiredPermissionDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequiredPermissionDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RequiredPermissionDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *RequiredPermissionDTO) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RequiredPermissionDTO) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RequiredPermissionDTO) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *RequiredPermissionDTO) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


