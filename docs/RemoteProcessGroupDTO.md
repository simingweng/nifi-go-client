# RemoteProcessGroupDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the component. | [optional] 
**VersionedComponentId** | Pointer to **string** | The ID of the corresponding component that is under version control | [optional] 
**ParentGroupId** | Pointer to **string** | The id of parent process group of this component if applicable. | [optional] 
**Position** | Pointer to [**PositionDTO**](PositionDTO.md) |  | [optional] 
**TargetUri** | Pointer to **string** | The target URI of the remote process group. If target uri is not set, but uris are set, then returns the first url in the urls. If neither target uri nor uris are set, then returns null. | [optional] 
**TargetUris** | Pointer to **string** | The target URI of the remote process group. If target uris is not set but target uri is set, then returns a collection containing the single target uri. If neither target uris nor uris are set, then returns null. | [optional] 
**TargetSecure** | Pointer to **bool** | Whether the target is running securely. | [optional] 
**Name** | Pointer to **string** | The name of the remote process group. | [optional] 
**Comments** | Pointer to **string** | The comments for the remote process group. | [optional] 
**CommunicationsTimeout** | Pointer to **string** | The time period used for the timeout when communicating with the target. | [optional] 
**YieldDuration** | Pointer to **string** | When yielding, this amount of time must elapse before the remote process group is scheduled again. | [optional] 
**TransportProtocol** | Pointer to **string** |  | [optional] 
**LocalNetworkInterface** | Pointer to **string** | The local network interface to send/receive data. If not specified, any local address is used. If clustered, all nodes must have an interface with this identifier. | [optional] 
**ProxyHost** | Pointer to **string** |  | [optional] 
**ProxyPort** | Pointer to **int32** |  | [optional] 
**ProxyUser** | Pointer to **string** |  | [optional] 
**ProxyPassword** | Pointer to **string** |  | [optional] 
**AuthorizationIssues** | Pointer to **[]string** | Any remote authorization issues for the remote process group. | [optional] 
**ValidationErrors** | Pointer to **[]string** | The validation errors for the remote process group. These validation errors represent the problems with the remote process group that must be resolved before it can transmit. | [optional] 
**Transmitting** | Pointer to **bool** | Whether the remote process group is actively transmitting. | [optional] 
**InputPortCount** | Pointer to **int32** | The number of remote input ports currently available on the target. | [optional] 
**OutputPortCount** | Pointer to **int32** | The number of remote output ports currently available on the target. | [optional] 
**ActiveRemoteInputPortCount** | Pointer to **int32** | The number of active remote input ports. | [optional] 
**InactiveRemoteInputPortCount** | Pointer to **int32** | The number of inactive remote input ports. | [optional] 
**ActiveRemoteOutputPortCount** | Pointer to **int32** | The number of active remote output ports. | [optional] 
**InactiveRemoteOutputPortCount** | Pointer to **int32** | The number of inactive remote output ports. | [optional] 
**FlowRefreshed** | Pointer to **string** | The timestamp when this remote process group was last refreshed. | [optional] 
**Contents** | Pointer to [**RemoteProcessGroupContentsDTO**](RemoteProcessGroupContentsDTO.md) |  | [optional] 

## Methods

### NewRemoteProcessGroupDTO

`func NewRemoteProcessGroupDTO() *RemoteProcessGroupDTO`

NewRemoteProcessGroupDTO instantiates a new RemoteProcessGroupDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteProcessGroupDTOWithDefaults

`func NewRemoteProcessGroupDTOWithDefaults() *RemoteProcessGroupDTO`

NewRemoteProcessGroupDTOWithDefaults instantiates a new RemoteProcessGroupDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RemoteProcessGroupDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoteProcessGroupDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoteProcessGroupDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RemoteProcessGroupDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersionedComponentId

`func (o *RemoteProcessGroupDTO) GetVersionedComponentId() string`

GetVersionedComponentId returns the VersionedComponentId field if non-nil, zero value otherwise.

### GetVersionedComponentIdOk

`func (o *RemoteProcessGroupDTO) GetVersionedComponentIdOk() (*string, bool)`

GetVersionedComponentIdOk returns a tuple with the VersionedComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedComponentId

`func (o *RemoteProcessGroupDTO) SetVersionedComponentId(v string)`

SetVersionedComponentId sets VersionedComponentId field to given value.

### HasVersionedComponentId

`func (o *RemoteProcessGroupDTO) HasVersionedComponentId() bool`

HasVersionedComponentId returns a boolean if a field has been set.

### GetParentGroupId

`func (o *RemoteProcessGroupDTO) GetParentGroupId() string`

GetParentGroupId returns the ParentGroupId field if non-nil, zero value otherwise.

### GetParentGroupIdOk

`func (o *RemoteProcessGroupDTO) GetParentGroupIdOk() (*string, bool)`

GetParentGroupIdOk returns a tuple with the ParentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroupId

`func (o *RemoteProcessGroupDTO) SetParentGroupId(v string)`

SetParentGroupId sets ParentGroupId field to given value.

### HasParentGroupId

`func (o *RemoteProcessGroupDTO) HasParentGroupId() bool`

HasParentGroupId returns a boolean if a field has been set.

### GetPosition

`func (o *RemoteProcessGroupDTO) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *RemoteProcessGroupDTO) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *RemoteProcessGroupDTO) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *RemoteProcessGroupDTO) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetTargetUri

`func (o *RemoteProcessGroupDTO) GetTargetUri() string`

GetTargetUri returns the TargetUri field if non-nil, zero value otherwise.

### GetTargetUriOk

`func (o *RemoteProcessGroupDTO) GetTargetUriOk() (*string, bool)`

GetTargetUriOk returns a tuple with the TargetUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUri

`func (o *RemoteProcessGroupDTO) SetTargetUri(v string)`

SetTargetUri sets TargetUri field to given value.

### HasTargetUri

`func (o *RemoteProcessGroupDTO) HasTargetUri() bool`

HasTargetUri returns a boolean if a field has been set.

### GetTargetUris

`func (o *RemoteProcessGroupDTO) GetTargetUris() string`

GetTargetUris returns the TargetUris field if non-nil, zero value otherwise.

### GetTargetUrisOk

`func (o *RemoteProcessGroupDTO) GetTargetUrisOk() (*string, bool)`

GetTargetUrisOk returns a tuple with the TargetUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUris

`func (o *RemoteProcessGroupDTO) SetTargetUris(v string)`

SetTargetUris sets TargetUris field to given value.

### HasTargetUris

`func (o *RemoteProcessGroupDTO) HasTargetUris() bool`

HasTargetUris returns a boolean if a field has been set.

### GetTargetSecure

`func (o *RemoteProcessGroupDTO) GetTargetSecure() bool`

GetTargetSecure returns the TargetSecure field if non-nil, zero value otherwise.

### GetTargetSecureOk

`func (o *RemoteProcessGroupDTO) GetTargetSecureOk() (*bool, bool)`

GetTargetSecureOk returns a tuple with the TargetSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSecure

`func (o *RemoteProcessGroupDTO) SetTargetSecure(v bool)`

SetTargetSecure sets TargetSecure field to given value.

### HasTargetSecure

`func (o *RemoteProcessGroupDTO) HasTargetSecure() bool`

HasTargetSecure returns a boolean if a field has been set.

### GetName

`func (o *RemoteProcessGroupDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemoteProcessGroupDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemoteProcessGroupDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RemoteProcessGroupDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComments

`func (o *RemoteProcessGroupDTO) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *RemoteProcessGroupDTO) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *RemoteProcessGroupDTO) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *RemoteProcessGroupDTO) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCommunicationsTimeout

`func (o *RemoteProcessGroupDTO) GetCommunicationsTimeout() string`

GetCommunicationsTimeout returns the CommunicationsTimeout field if non-nil, zero value otherwise.

### GetCommunicationsTimeoutOk

`func (o *RemoteProcessGroupDTO) GetCommunicationsTimeoutOk() (*string, bool)`

GetCommunicationsTimeoutOk returns a tuple with the CommunicationsTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationsTimeout

`func (o *RemoteProcessGroupDTO) SetCommunicationsTimeout(v string)`

SetCommunicationsTimeout sets CommunicationsTimeout field to given value.

### HasCommunicationsTimeout

`func (o *RemoteProcessGroupDTO) HasCommunicationsTimeout() bool`

HasCommunicationsTimeout returns a boolean if a field has been set.

### GetYieldDuration

`func (o *RemoteProcessGroupDTO) GetYieldDuration() string`

GetYieldDuration returns the YieldDuration field if non-nil, zero value otherwise.

### GetYieldDurationOk

`func (o *RemoteProcessGroupDTO) GetYieldDurationOk() (*string, bool)`

GetYieldDurationOk returns a tuple with the YieldDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYieldDuration

`func (o *RemoteProcessGroupDTO) SetYieldDuration(v string)`

SetYieldDuration sets YieldDuration field to given value.

### HasYieldDuration

`func (o *RemoteProcessGroupDTO) HasYieldDuration() bool`

HasYieldDuration returns a boolean if a field has been set.

### GetTransportProtocol

`func (o *RemoteProcessGroupDTO) GetTransportProtocol() string`

GetTransportProtocol returns the TransportProtocol field if non-nil, zero value otherwise.

### GetTransportProtocolOk

`func (o *RemoteProcessGroupDTO) GetTransportProtocolOk() (*string, bool)`

GetTransportProtocolOk returns a tuple with the TransportProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportProtocol

`func (o *RemoteProcessGroupDTO) SetTransportProtocol(v string)`

SetTransportProtocol sets TransportProtocol field to given value.

### HasTransportProtocol

`func (o *RemoteProcessGroupDTO) HasTransportProtocol() bool`

HasTransportProtocol returns a boolean if a field has been set.

### GetLocalNetworkInterface

`func (o *RemoteProcessGroupDTO) GetLocalNetworkInterface() string`

GetLocalNetworkInterface returns the LocalNetworkInterface field if non-nil, zero value otherwise.

### GetLocalNetworkInterfaceOk

`func (o *RemoteProcessGroupDTO) GetLocalNetworkInterfaceOk() (*string, bool)`

GetLocalNetworkInterfaceOk returns a tuple with the LocalNetworkInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNetworkInterface

`func (o *RemoteProcessGroupDTO) SetLocalNetworkInterface(v string)`

SetLocalNetworkInterface sets LocalNetworkInterface field to given value.

### HasLocalNetworkInterface

`func (o *RemoteProcessGroupDTO) HasLocalNetworkInterface() bool`

HasLocalNetworkInterface returns a boolean if a field has been set.

### GetProxyHost

`func (o *RemoteProcessGroupDTO) GetProxyHost() string`

GetProxyHost returns the ProxyHost field if non-nil, zero value otherwise.

### GetProxyHostOk

`func (o *RemoteProcessGroupDTO) GetProxyHostOk() (*string, bool)`

GetProxyHostOk returns a tuple with the ProxyHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyHost

`func (o *RemoteProcessGroupDTO) SetProxyHost(v string)`

SetProxyHost sets ProxyHost field to given value.

### HasProxyHost

`func (o *RemoteProcessGroupDTO) HasProxyHost() bool`

HasProxyHost returns a boolean if a field has been set.

### GetProxyPort

`func (o *RemoteProcessGroupDTO) GetProxyPort() int32`

GetProxyPort returns the ProxyPort field if non-nil, zero value otherwise.

### GetProxyPortOk

`func (o *RemoteProcessGroupDTO) GetProxyPortOk() (*int32, bool)`

GetProxyPortOk returns a tuple with the ProxyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPort

`func (o *RemoteProcessGroupDTO) SetProxyPort(v int32)`

SetProxyPort sets ProxyPort field to given value.

### HasProxyPort

`func (o *RemoteProcessGroupDTO) HasProxyPort() bool`

HasProxyPort returns a boolean if a field has been set.

### GetProxyUser

`func (o *RemoteProcessGroupDTO) GetProxyUser() string`

GetProxyUser returns the ProxyUser field if non-nil, zero value otherwise.

### GetProxyUserOk

`func (o *RemoteProcessGroupDTO) GetProxyUserOk() (*string, bool)`

GetProxyUserOk returns a tuple with the ProxyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUser

`func (o *RemoteProcessGroupDTO) SetProxyUser(v string)`

SetProxyUser sets ProxyUser field to given value.

### HasProxyUser

`func (o *RemoteProcessGroupDTO) HasProxyUser() bool`

HasProxyUser returns a boolean if a field has been set.

### GetProxyPassword

`func (o *RemoteProcessGroupDTO) GetProxyPassword() string`

GetProxyPassword returns the ProxyPassword field if non-nil, zero value otherwise.

### GetProxyPasswordOk

`func (o *RemoteProcessGroupDTO) GetProxyPasswordOk() (*string, bool)`

GetProxyPasswordOk returns a tuple with the ProxyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPassword

`func (o *RemoteProcessGroupDTO) SetProxyPassword(v string)`

SetProxyPassword sets ProxyPassword field to given value.

### HasProxyPassword

`func (o *RemoteProcessGroupDTO) HasProxyPassword() bool`

HasProxyPassword returns a boolean if a field has been set.

### GetAuthorizationIssues

`func (o *RemoteProcessGroupDTO) GetAuthorizationIssues() []string`

GetAuthorizationIssues returns the AuthorizationIssues field if non-nil, zero value otherwise.

### GetAuthorizationIssuesOk

`func (o *RemoteProcessGroupDTO) GetAuthorizationIssuesOk() (*[]string, bool)`

GetAuthorizationIssuesOk returns a tuple with the AuthorizationIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationIssues

`func (o *RemoteProcessGroupDTO) SetAuthorizationIssues(v []string)`

SetAuthorizationIssues sets AuthorizationIssues field to given value.

### HasAuthorizationIssues

`func (o *RemoteProcessGroupDTO) HasAuthorizationIssues() bool`

HasAuthorizationIssues returns a boolean if a field has been set.

### GetValidationErrors

`func (o *RemoteProcessGroupDTO) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *RemoteProcessGroupDTO) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *RemoteProcessGroupDTO) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *RemoteProcessGroupDTO) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### GetTransmitting

`func (o *RemoteProcessGroupDTO) GetTransmitting() bool`

GetTransmitting returns the Transmitting field if non-nil, zero value otherwise.

### GetTransmittingOk

`func (o *RemoteProcessGroupDTO) GetTransmittingOk() (*bool, bool)`

GetTransmittingOk returns a tuple with the Transmitting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitting

`func (o *RemoteProcessGroupDTO) SetTransmitting(v bool)`

SetTransmitting sets Transmitting field to given value.

### HasTransmitting

`func (o *RemoteProcessGroupDTO) HasTransmitting() bool`

HasTransmitting returns a boolean if a field has been set.

### GetInputPortCount

`func (o *RemoteProcessGroupDTO) GetInputPortCount() int32`

GetInputPortCount returns the InputPortCount field if non-nil, zero value otherwise.

### GetInputPortCountOk

`func (o *RemoteProcessGroupDTO) GetInputPortCountOk() (*int32, bool)`

GetInputPortCountOk returns a tuple with the InputPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPortCount

`func (o *RemoteProcessGroupDTO) SetInputPortCount(v int32)`

SetInputPortCount sets InputPortCount field to given value.

### HasInputPortCount

`func (o *RemoteProcessGroupDTO) HasInputPortCount() bool`

HasInputPortCount returns a boolean if a field has been set.

### GetOutputPortCount

`func (o *RemoteProcessGroupDTO) GetOutputPortCount() int32`

GetOutputPortCount returns the OutputPortCount field if non-nil, zero value otherwise.

### GetOutputPortCountOk

`func (o *RemoteProcessGroupDTO) GetOutputPortCountOk() (*int32, bool)`

GetOutputPortCountOk returns a tuple with the OutputPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPortCount

`func (o *RemoteProcessGroupDTO) SetOutputPortCount(v int32)`

SetOutputPortCount sets OutputPortCount field to given value.

### HasOutputPortCount

`func (o *RemoteProcessGroupDTO) HasOutputPortCount() bool`

HasOutputPortCount returns a boolean if a field has been set.

### GetActiveRemoteInputPortCount

`func (o *RemoteProcessGroupDTO) GetActiveRemoteInputPortCount() int32`

GetActiveRemoteInputPortCount returns the ActiveRemoteInputPortCount field if non-nil, zero value otherwise.

### GetActiveRemoteInputPortCountOk

`func (o *RemoteProcessGroupDTO) GetActiveRemoteInputPortCountOk() (*int32, bool)`

GetActiveRemoteInputPortCountOk returns a tuple with the ActiveRemoteInputPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRemoteInputPortCount

`func (o *RemoteProcessGroupDTO) SetActiveRemoteInputPortCount(v int32)`

SetActiveRemoteInputPortCount sets ActiveRemoteInputPortCount field to given value.

### HasActiveRemoteInputPortCount

`func (o *RemoteProcessGroupDTO) HasActiveRemoteInputPortCount() bool`

HasActiveRemoteInputPortCount returns a boolean if a field has been set.

### GetInactiveRemoteInputPortCount

`func (o *RemoteProcessGroupDTO) GetInactiveRemoteInputPortCount() int32`

GetInactiveRemoteInputPortCount returns the InactiveRemoteInputPortCount field if non-nil, zero value otherwise.

### GetInactiveRemoteInputPortCountOk

`func (o *RemoteProcessGroupDTO) GetInactiveRemoteInputPortCountOk() (*int32, bool)`

GetInactiveRemoteInputPortCountOk returns a tuple with the InactiveRemoteInputPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveRemoteInputPortCount

`func (o *RemoteProcessGroupDTO) SetInactiveRemoteInputPortCount(v int32)`

SetInactiveRemoteInputPortCount sets InactiveRemoteInputPortCount field to given value.

### HasInactiveRemoteInputPortCount

`func (o *RemoteProcessGroupDTO) HasInactiveRemoteInputPortCount() bool`

HasInactiveRemoteInputPortCount returns a boolean if a field has been set.

### GetActiveRemoteOutputPortCount

`func (o *RemoteProcessGroupDTO) GetActiveRemoteOutputPortCount() int32`

GetActiveRemoteOutputPortCount returns the ActiveRemoteOutputPortCount field if non-nil, zero value otherwise.

### GetActiveRemoteOutputPortCountOk

`func (o *RemoteProcessGroupDTO) GetActiveRemoteOutputPortCountOk() (*int32, bool)`

GetActiveRemoteOutputPortCountOk returns a tuple with the ActiveRemoteOutputPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRemoteOutputPortCount

`func (o *RemoteProcessGroupDTO) SetActiveRemoteOutputPortCount(v int32)`

SetActiveRemoteOutputPortCount sets ActiveRemoteOutputPortCount field to given value.

### HasActiveRemoteOutputPortCount

`func (o *RemoteProcessGroupDTO) HasActiveRemoteOutputPortCount() bool`

HasActiveRemoteOutputPortCount returns a boolean if a field has been set.

### GetInactiveRemoteOutputPortCount

`func (o *RemoteProcessGroupDTO) GetInactiveRemoteOutputPortCount() int32`

GetInactiveRemoteOutputPortCount returns the InactiveRemoteOutputPortCount field if non-nil, zero value otherwise.

### GetInactiveRemoteOutputPortCountOk

`func (o *RemoteProcessGroupDTO) GetInactiveRemoteOutputPortCountOk() (*int32, bool)`

GetInactiveRemoteOutputPortCountOk returns a tuple with the InactiveRemoteOutputPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveRemoteOutputPortCount

`func (o *RemoteProcessGroupDTO) SetInactiveRemoteOutputPortCount(v int32)`

SetInactiveRemoteOutputPortCount sets InactiveRemoteOutputPortCount field to given value.

### HasInactiveRemoteOutputPortCount

`func (o *RemoteProcessGroupDTO) HasInactiveRemoteOutputPortCount() bool`

HasInactiveRemoteOutputPortCount returns a boolean if a field has been set.

### GetFlowRefreshed

`func (o *RemoteProcessGroupDTO) GetFlowRefreshed() string`

GetFlowRefreshed returns the FlowRefreshed field if non-nil, zero value otherwise.

### GetFlowRefreshedOk

`func (o *RemoteProcessGroupDTO) GetFlowRefreshedOk() (*string, bool)`

GetFlowRefreshedOk returns a tuple with the FlowRefreshed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowRefreshed

`func (o *RemoteProcessGroupDTO) SetFlowRefreshed(v string)`

SetFlowRefreshed sets FlowRefreshed field to given value.

### HasFlowRefreshed

`func (o *RemoteProcessGroupDTO) HasFlowRefreshed() bool`

HasFlowRefreshed returns a boolean if a field has been set.

### GetContents

`func (o *RemoteProcessGroupDTO) GetContents() RemoteProcessGroupContentsDTO`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *RemoteProcessGroupDTO) GetContentsOk() (*RemoteProcessGroupContentsDTO, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *RemoteProcessGroupDTO) SetContents(v RemoteProcessGroupContentsDTO)`

SetContents sets Contents field to given value.

### HasContents

`func (o *RemoteProcessGroupDTO) HasContents() bool`

HasContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


