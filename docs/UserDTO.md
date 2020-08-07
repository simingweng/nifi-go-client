# UserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the component. | [optional] 
**VersionedComponentId** | **string** | The ID of the corresponding component that is under version control | [optional] 
**ParentGroupId** | **string** | The id of parent process group of this component if applicable. | [optional] 
**Position** | [**PositionDto**](PositionDTO.md) |  | [optional] 
**Identity** | **string** | The identity of the tenant. | [optional] 
**Configurable** | **bool** | Whether this tenant is configurable. | [optional] 
**UserGroups** | [**[]TenantEntity**](TenantEntity.md) | The groups to which the user belongs. This field is read only and it provided for convenience. | [optional] [readonly] 
**AccessPolicies** | [**[]AccessPolicySummaryEntity**](AccessPolicySummaryEntity.md) | The access policies this user belongs to. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


