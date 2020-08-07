# FlowConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportsManagedAuthorizer** | **bool** | Whether this NiFi supports a managed authorizer. Managed authorizers can visualize users, groups, and policies in the UI. | [optional] [readonly] 
**SupportsConfigurableAuthorizer** | **bool** | Whether this NiFi supports a configurable authorizer. | [optional] [readonly] 
**SupportsConfigurableUsersAndGroups** | **bool** | Whether this NiFi supports configurable users and groups. | [optional] [readonly] 
**AutoRefreshIntervalSeconds** | **int64** | The interval in seconds between the automatic NiFi refresh requests. | [optional] [readonly] 
**CurrentTime** | **string** | The current time on the system. | [optional] 
**TimeOffset** | **int32** | The time offset of the system. | [optional] 
**DefaultBackPressureObjectThreshold** | **int64** | The default back pressure object threshold. | [optional] 
**DefaultBackPressureDataSizeThreshold** | **string** | The default back pressure data size threshold. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


