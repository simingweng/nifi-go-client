# ComponentRestrictionPermissionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequiredPermission** | Pointer to [**RequiredPermissionDTO**](RequiredPermissionDTO.md) |  | [optional] 
**Permissions** | Pointer to [**PermissionsDTO**](PermissionsDTO.md) |  | [optional] 

## Methods

### NewComponentRestrictionPermissionDTO

`func NewComponentRestrictionPermissionDTO() *ComponentRestrictionPermissionDTO`

NewComponentRestrictionPermissionDTO instantiates a new ComponentRestrictionPermissionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentRestrictionPermissionDTOWithDefaults

`func NewComponentRestrictionPermissionDTOWithDefaults() *ComponentRestrictionPermissionDTO`

NewComponentRestrictionPermissionDTOWithDefaults instantiates a new ComponentRestrictionPermissionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequiredPermission

`func (o *ComponentRestrictionPermissionDTO) GetRequiredPermission() RequiredPermissionDTO`

GetRequiredPermission returns the RequiredPermission field if non-nil, zero value otherwise.

### GetRequiredPermissionOk

`func (o *ComponentRestrictionPermissionDTO) GetRequiredPermissionOk() (*RequiredPermissionDTO, bool)`

GetRequiredPermissionOk returns a tuple with the RequiredPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredPermission

`func (o *ComponentRestrictionPermissionDTO) SetRequiredPermission(v RequiredPermissionDTO)`

SetRequiredPermission sets RequiredPermission field to given value.

### HasRequiredPermission

`func (o *ComponentRestrictionPermissionDTO) HasRequiredPermission() bool`

HasRequiredPermission returns a boolean if a field has been set.

### GetPermissions

`func (o *ComponentRestrictionPermissionDTO) GetPermissions() PermissionsDTO`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ComponentRestrictionPermissionDTO) GetPermissionsOk() (*PermissionsDTO, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ComponentRestrictionPermissionDTO) SetPermissions(v PermissionsDTO)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ComponentRestrictionPermissionDTO) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


