# FlowConfigurationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportsManagedAuthorizer** | Pointer to **bool** | Whether this NiFi supports a managed authorizer. Managed authorizers can visualize users, groups, and policies in the UI. | [optional] [readonly] 
**SupportsConfigurableAuthorizer** | Pointer to **bool** | Whether this NiFi supports a configurable authorizer. | [optional] [readonly] 
**SupportsConfigurableUsersAndGroups** | Pointer to **bool** | Whether this NiFi supports configurable users and groups. | [optional] [readonly] 
**AutoRefreshIntervalSeconds** | Pointer to **int64** | The interval in seconds between the automatic NiFi refresh requests. | [optional] [readonly] 
**CurrentTime** | Pointer to **string** | The current time on the system. | [optional] 
**TimeOffset** | Pointer to **int32** | The time offset of the system. | [optional] 
**DefaultBackPressureObjectThreshold** | Pointer to **int64** | The default back pressure object threshold. | [optional] 
**DefaultBackPressureDataSizeThreshold** | Pointer to **string** | The default back pressure data size threshold. | [optional] 

## Methods

### NewFlowConfigurationDTO

`func NewFlowConfigurationDTO() *FlowConfigurationDTO`

NewFlowConfigurationDTO instantiates a new FlowConfigurationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowConfigurationDTOWithDefaults

`func NewFlowConfigurationDTOWithDefaults() *FlowConfigurationDTO`

NewFlowConfigurationDTOWithDefaults instantiates a new FlowConfigurationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportsManagedAuthorizer

`func (o *FlowConfigurationDTO) GetSupportsManagedAuthorizer() bool`

GetSupportsManagedAuthorizer returns the SupportsManagedAuthorizer field if non-nil, zero value otherwise.

### GetSupportsManagedAuthorizerOk

`func (o *FlowConfigurationDTO) GetSupportsManagedAuthorizerOk() (*bool, bool)`

GetSupportsManagedAuthorizerOk returns a tuple with the SupportsManagedAuthorizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsManagedAuthorizer

`func (o *FlowConfigurationDTO) SetSupportsManagedAuthorizer(v bool)`

SetSupportsManagedAuthorizer sets SupportsManagedAuthorizer field to given value.

### HasSupportsManagedAuthorizer

`func (o *FlowConfigurationDTO) HasSupportsManagedAuthorizer() bool`

HasSupportsManagedAuthorizer returns a boolean if a field has been set.

### GetSupportsConfigurableAuthorizer

`func (o *FlowConfigurationDTO) GetSupportsConfigurableAuthorizer() bool`

GetSupportsConfigurableAuthorizer returns the SupportsConfigurableAuthorizer field if non-nil, zero value otherwise.

### GetSupportsConfigurableAuthorizerOk

`func (o *FlowConfigurationDTO) GetSupportsConfigurableAuthorizerOk() (*bool, bool)`

GetSupportsConfigurableAuthorizerOk returns a tuple with the SupportsConfigurableAuthorizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsConfigurableAuthorizer

`func (o *FlowConfigurationDTO) SetSupportsConfigurableAuthorizer(v bool)`

SetSupportsConfigurableAuthorizer sets SupportsConfigurableAuthorizer field to given value.

### HasSupportsConfigurableAuthorizer

`func (o *FlowConfigurationDTO) HasSupportsConfigurableAuthorizer() bool`

HasSupportsConfigurableAuthorizer returns a boolean if a field has been set.

### GetSupportsConfigurableUsersAndGroups

`func (o *FlowConfigurationDTO) GetSupportsConfigurableUsersAndGroups() bool`

GetSupportsConfigurableUsersAndGroups returns the SupportsConfigurableUsersAndGroups field if non-nil, zero value otherwise.

### GetSupportsConfigurableUsersAndGroupsOk

`func (o *FlowConfigurationDTO) GetSupportsConfigurableUsersAndGroupsOk() (*bool, bool)`

GetSupportsConfigurableUsersAndGroupsOk returns a tuple with the SupportsConfigurableUsersAndGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsConfigurableUsersAndGroups

`func (o *FlowConfigurationDTO) SetSupportsConfigurableUsersAndGroups(v bool)`

SetSupportsConfigurableUsersAndGroups sets SupportsConfigurableUsersAndGroups field to given value.

### HasSupportsConfigurableUsersAndGroups

`func (o *FlowConfigurationDTO) HasSupportsConfigurableUsersAndGroups() bool`

HasSupportsConfigurableUsersAndGroups returns a boolean if a field has been set.

### GetAutoRefreshIntervalSeconds

`func (o *FlowConfigurationDTO) GetAutoRefreshIntervalSeconds() int64`

GetAutoRefreshIntervalSeconds returns the AutoRefreshIntervalSeconds field if non-nil, zero value otherwise.

### GetAutoRefreshIntervalSecondsOk

`func (o *FlowConfigurationDTO) GetAutoRefreshIntervalSecondsOk() (*int64, bool)`

GetAutoRefreshIntervalSecondsOk returns a tuple with the AutoRefreshIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRefreshIntervalSeconds

`func (o *FlowConfigurationDTO) SetAutoRefreshIntervalSeconds(v int64)`

SetAutoRefreshIntervalSeconds sets AutoRefreshIntervalSeconds field to given value.

### HasAutoRefreshIntervalSeconds

`func (o *FlowConfigurationDTO) HasAutoRefreshIntervalSeconds() bool`

HasAutoRefreshIntervalSeconds returns a boolean if a field has been set.

### GetCurrentTime

`func (o *FlowConfigurationDTO) GetCurrentTime() string`

GetCurrentTime returns the CurrentTime field if non-nil, zero value otherwise.

### GetCurrentTimeOk

`func (o *FlowConfigurationDTO) GetCurrentTimeOk() (*string, bool)`

GetCurrentTimeOk returns a tuple with the CurrentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTime

`func (o *FlowConfigurationDTO) SetCurrentTime(v string)`

SetCurrentTime sets CurrentTime field to given value.

### HasCurrentTime

`func (o *FlowConfigurationDTO) HasCurrentTime() bool`

HasCurrentTime returns a boolean if a field has been set.

### GetTimeOffset

`func (o *FlowConfigurationDTO) GetTimeOffset() int32`

GetTimeOffset returns the TimeOffset field if non-nil, zero value otherwise.

### GetTimeOffsetOk

`func (o *FlowConfigurationDTO) GetTimeOffsetOk() (*int32, bool)`

GetTimeOffsetOk returns a tuple with the TimeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOffset

`func (o *FlowConfigurationDTO) SetTimeOffset(v int32)`

SetTimeOffset sets TimeOffset field to given value.

### HasTimeOffset

`func (o *FlowConfigurationDTO) HasTimeOffset() bool`

HasTimeOffset returns a boolean if a field has been set.

### GetDefaultBackPressureObjectThreshold

`func (o *FlowConfigurationDTO) GetDefaultBackPressureObjectThreshold() int64`

GetDefaultBackPressureObjectThreshold returns the DefaultBackPressureObjectThreshold field if non-nil, zero value otherwise.

### GetDefaultBackPressureObjectThresholdOk

`func (o *FlowConfigurationDTO) GetDefaultBackPressureObjectThresholdOk() (*int64, bool)`

GetDefaultBackPressureObjectThresholdOk returns a tuple with the DefaultBackPressureObjectThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBackPressureObjectThreshold

`func (o *FlowConfigurationDTO) SetDefaultBackPressureObjectThreshold(v int64)`

SetDefaultBackPressureObjectThreshold sets DefaultBackPressureObjectThreshold field to given value.

### HasDefaultBackPressureObjectThreshold

`func (o *FlowConfigurationDTO) HasDefaultBackPressureObjectThreshold() bool`

HasDefaultBackPressureObjectThreshold returns a boolean if a field has been set.

### GetDefaultBackPressureDataSizeThreshold

`func (o *FlowConfigurationDTO) GetDefaultBackPressureDataSizeThreshold() string`

GetDefaultBackPressureDataSizeThreshold returns the DefaultBackPressureDataSizeThreshold field if non-nil, zero value otherwise.

### GetDefaultBackPressureDataSizeThresholdOk

`func (o *FlowConfigurationDTO) GetDefaultBackPressureDataSizeThresholdOk() (*string, bool)`

GetDefaultBackPressureDataSizeThresholdOk returns a tuple with the DefaultBackPressureDataSizeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBackPressureDataSizeThreshold

`func (o *FlowConfigurationDTO) SetDefaultBackPressureDataSizeThreshold(v string)`

SetDefaultBackPressureDataSizeThreshold sets DefaultBackPressureDataSizeThreshold field to given value.

### HasDefaultBackPressureDataSizeThreshold

`func (o *FlowConfigurationDTO) HasDefaultBackPressureDataSizeThreshold() bool`

HasDefaultBackPressureDataSizeThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


