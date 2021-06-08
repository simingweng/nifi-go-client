# PermissionsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanRead** | Pointer to **bool** | Indicates whether the user can read a given resource. | [optional] [readonly] 
**CanWrite** | Pointer to **bool** | Indicates whether the user can write a given resource. | [optional] [readonly] 

## Methods

### NewPermissionsDTO

`func NewPermissionsDTO() *PermissionsDTO`

NewPermissionsDTO instantiates a new PermissionsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsDTOWithDefaults

`func NewPermissionsDTOWithDefaults() *PermissionsDTO`

NewPermissionsDTOWithDefaults instantiates a new PermissionsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanRead

`func (o *PermissionsDTO) GetCanRead() bool`

GetCanRead returns the CanRead field if non-nil, zero value otherwise.

### GetCanReadOk

`func (o *PermissionsDTO) GetCanReadOk() (*bool, bool)`

GetCanReadOk returns a tuple with the CanRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRead

`func (o *PermissionsDTO) SetCanRead(v bool)`

SetCanRead sets CanRead field to given value.

### HasCanRead

`func (o *PermissionsDTO) HasCanRead() bool`

HasCanRead returns a boolean if a field has been set.

### GetCanWrite

`func (o *PermissionsDTO) GetCanWrite() bool`

GetCanWrite returns the CanWrite field if non-nil, zero value otherwise.

### GetCanWriteOk

`func (o *PermissionsDTO) GetCanWriteOk() (*bool, bool)`

GetCanWriteOk returns a tuple with the CanWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWrite

`func (o *PermissionsDTO) SetCanWrite(v bool)`

SetCanWrite sets CanWrite field to given value.

### HasCanWrite

`func (o *PermissionsDTO) HasCanWrite() bool`

HasCanWrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


