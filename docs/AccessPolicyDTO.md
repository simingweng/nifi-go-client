# AccessPolicyDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the component. | [optional] 
**VersionedComponentId** | **string** | The ID of the corresponding component that is under version control | [optional] 
**ParentGroupId** | **string** | The id of parent process group of this component if applicable. | [optional] 
**Position** | [**PositionDto**](PositionDTO.md) |  | [optional] 
**Resource** | **string** | The resource for this access policy. | [optional] 
**Action** | **string** | The action associated with this access policy. | [optional] 
**ComponentReference** | [**ComponentReferenceEntity**](ComponentReferenceEntity.md) |  | [optional] 
**Configurable** | **bool** | Whether this policy is configurable. | [optional] 
**Users** | [**[]TenantEntity**](TenantEntity.md) | The set of user IDs associated with this access policy. | [optional] 
**UserGroups** | [**[]TenantEntity**](TenantEntity.md) | The set of user group IDs associated with this access policy. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


