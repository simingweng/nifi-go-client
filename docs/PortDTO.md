# PortDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the component. | [optional] 
**VersionedComponentId** | **string** | The ID of the corresponding component that is under version control | [optional] 
**ParentGroupId** | **string** | The id of parent process group of this component if applicable. | [optional] 
**Position** | [**PositionDto**](PositionDTO.md) |  | [optional] 
**Name** | **string** | The name of the port. | [optional] 
**Comments** | **string** | The comments for the port. | [optional] 
**State** | **string** | The state of the port. | [optional] 
**Type** | **string** | The type of port. | [optional] 
**Transmitting** | **bool** | Whether the port has incoming or output connections to a remote NiFi. This is only applicable when the port is allowed to be accessed remotely. | [optional] 
**ConcurrentlySchedulableTaskCount** | **int32** | The number of tasks that should be concurrently scheduled for the port. | [optional] 
**UserAccessControl** | **[]string** | The users that are allowed to access the port. | [optional] 
**GroupAccessControl** | **[]string** | The user groups that are allowed to access the port. | [optional] 
**AllowRemoteAccess** | **bool** | Whether this port can be accessed remotely via Site-to-Site protocol. | [optional] 
**ValidationErrors** | **[]string** | Gets the validation errors from this port. These validation errors represent the problems with the port that must be resolved before it can be started. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


