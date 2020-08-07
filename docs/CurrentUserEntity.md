# CurrentUserEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | **string** | The user identity being serialized. | [optional] 
**Anonymous** | **bool** | Whether the current user is anonymous. | [optional] 
**ProvenancePermissions** | [**PermissionsDto**](PermissionsDTO.md) |  | [optional] 
**CountersPermissions** | [**PermissionsDto**](PermissionsDTO.md) |  | [optional] 
**TenantsPermissions** | [**PermissionsDto**](PermissionsDTO.md) |  | [optional] 
**ControllerPermissions** | [**PermissionsDto**](PermissionsDTO.md) |  | [optional] 
**PoliciesPermissions** | [**PermissionsDto**](PermissionsDTO.md) |  | [optional] 
**SystemPermissions** | [**PermissionsDto**](PermissionsDTO.md) |  | [optional] 
**ParameterContextPermissions** | [**PermissionsDto**](PermissionsDTO.md) |  | [optional] 
**RestrictedComponentsPermissions** | [**PermissionsDto**](PermissionsDTO.md) |  | [optional] 
**ComponentRestrictionPermissions** | [**[]ComponentRestrictionPermissionDto**](ComponentRestrictionPermissionDTO.md) | Permissions for specific component restrictions. | [optional] 
**CanVersionFlows** | **bool** | Whether the current user can version flows. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


