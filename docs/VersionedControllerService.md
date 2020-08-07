# VersionedControllerService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | The component&#39;s unique identifier | [optional] 
**Name** | **string** | The component&#39;s name | [optional] 
**Comments** | **string** | The user-supplied comments for the component | [optional] 
**Position** | [**Position**](Position.md) |  | [optional] 
**Type** | **string** | The type of the controller service. | [optional] 
**Bundle** | [**Bundle**](Bundle.md) |  | [optional] 
**ControllerServiceApis** | [**[]ControllerServiceApi**](ControllerServiceAPI.md) | Lists the APIs this Controller Service implements. | [optional] 
**Properties** | **map[string]string** | The properties of the controller service. | [optional] 
**PropertyDescriptors** | [**map[string]VersionedPropertyDescriptor**](VersionedPropertyDescriptor.md) | The property descriptors for the processor. | [optional] 
**AnnotationData** | **string** | The annotation for the controller service. This is how the custom UI relays configuration to the controller service. | [optional] 
**ComponentType** | **string** |  | [optional] 
**GroupIdentifier** | **string** | The ID of the Process Group that this component belongs to | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


