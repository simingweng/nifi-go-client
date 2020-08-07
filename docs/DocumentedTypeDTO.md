# DocumentedTypeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The fully qualified name of the type. | [optional] 
**Bundle** | [**BundleDto**](BundleDTO.md) |  | [optional] 
**ControllerServiceApis** | [**[]ControllerServiceApiDto**](ControllerServiceApiDTO.md) | If this type represents a ControllerService, this lists the APIs it implements. | [optional] 
**Description** | **string** | The description of the type. | [optional] 
**Restricted** | **bool** | Whether this type is restricted. | [optional] 
**UsageRestriction** | **string** | The optional description of why the usage of this component is restricted. | [optional] 
**ExplicitRestrictions** | [**[]ExplicitRestrictionDto**](ExplicitRestrictionDTO.md) | An optional collection of explicit restrictions. If specified, these explicit restrictions will be enfored. | [optional] 
**DeprecationReason** | **string** | The description of why the usage of this component is restricted. | [optional] 
**Tags** | **[]string** | The tags associated with this type. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


