# RemoteProcessGroupDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the component. | [optional] 
**VersionedComponentId** | **string** | The ID of the corresponding component that is under version control | [optional] 
**ParentGroupId** | **string** | The id of parent process group of this component if applicable. | [optional] 
**Position** | [**PositionDto**](PositionDTO.md) |  | [optional] 
**TargetUri** | **string** | The target URI of the remote process group. If target uri is not set, but uris are set, then returns the first url in the urls. If neither target uri nor uris are set, then returns null. | [optional] 
**TargetUris** | **string** | The target URI of the remote process group. If target uris is not set but target uri is set, then returns a collection containing the single target uri. If neither target uris nor uris are set, then returns null. | [optional] 
**TargetSecure** | **bool** | Whether the target is running securely. | [optional] 
**Name** | **string** | The name of the remote process group. | [optional] 
**Comments** | **string** | The comments for the remote process group. | [optional] 
**CommunicationsTimeout** | **string** | The time period used for the timeout when communicating with the target. | [optional] 
**YieldDuration** | **string** | When yielding, this amount of time must elapse before the remote process group is scheduled again. | [optional] 
**TransportProtocol** | **string** |  | [optional] 
**LocalNetworkInterface** | **string** | The local network interface to send/receive data. If not specified, any local address is used. If clustered, all nodes must have an interface with this identifier. | [optional] 
**ProxyHost** | **string** |  | [optional] 
**ProxyPort** | **int32** |  | [optional] 
**ProxyUser** | **string** |  | [optional] 
**ProxyPassword** | **string** |  | [optional] 
**AuthorizationIssues** | **[]string** | Any remote authorization issues for the remote process group. | [optional] 
**ValidationErrors** | **[]string** | The validation errors for the remote process group. These validation errors represent the problems with the remote process group that must be resolved before it can transmit. | [optional] 
**Transmitting** | **bool** | Whether the remote process group is actively transmitting. | [optional] 
**InputPortCount** | **int32** | The number of remote input ports currently available on the target. | [optional] 
**OutputPortCount** | **int32** | The number of remote output ports currently available on the target. | [optional] 
**ActiveRemoteInputPortCount** | **int32** | The number of active remote input ports. | [optional] 
**InactiveRemoteInputPortCount** | **int32** | The number of inactive remote input ports. | [optional] 
**ActiveRemoteOutputPortCount** | **int32** | The number of active remote output ports. | [optional] 
**InactiveRemoteOutputPortCount** | **int32** | The number of inactive remote output ports. | [optional] 
**FlowRefreshed** | **string** | The timestamp when this remote process group was last refreshed. | [optional] 
**Contents** | [**RemoteProcessGroupContentsDto**](RemoteProcessGroupContentsDTO.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


