# VersionedRemoteProcessGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | The component&#39;s unique identifier | [optional] 
**Name** | Pointer to **string** | The component&#39;s name | [optional] 
**Comments** | Pointer to **string** | The user-supplied comments for the component | [optional] 
**Position** | Pointer to [**Position**](Position.md) |  | [optional] 
**TargetUri** | Pointer to **string** | [DEPRECATED] The target URI of the remote process group. If target uri is not set, but uris are set, then returns the first uri in the uris. If neither target uri nor uris are set, then returns null. | [optional] 
**TargetUris** | Pointer to **string** | The target URIs of the remote process group. If target uris is not set but target uri is set, then returns the single target uri. If neither target uris nor target uri is set, then returns null. | [optional] 
**CommunicationsTimeout** | Pointer to **string** | The time period used for the timeout when communicating with the target. | [optional] 
**YieldDuration** | Pointer to **string** | When yielding, this amount of time must elapse before the remote process group is scheduled again. | [optional] 
**TransportProtocol** | Pointer to **string** | The Transport Protocol that is used for Site-to-Site communications | [optional] 
**LocalNetworkInterface** | Pointer to **string** | The local network interface to send/receive data. If not specified, any local address is used. If clustered, all nodes must have an interface with this identifier. | [optional] 
**ProxyHost** | Pointer to **string** |  | [optional] 
**ProxyPort** | Pointer to **int32** |  | [optional] 
**ProxyUser** | Pointer to **string** |  | [optional] 
**InputPorts** | Pointer to [**[]VersionedRemoteGroupPort**](VersionedRemoteGroupPort.md) | A Set of Input Ports that can be connected to, in order to send data to the remote NiFi instance | [optional] 
**OutputPorts** | Pointer to [**[]VersionedRemoteGroupPort**](VersionedRemoteGroupPort.md) | A Set of Output Ports that can be connected to, in order to pull data from the remote NiFi instance | [optional] 
**ComponentType** | Pointer to **string** |  | [optional] 
**GroupIdentifier** | Pointer to **string** | The ID of the Process Group that this component belongs to | [optional] 

## Methods

### NewVersionedRemoteProcessGroup

`func NewVersionedRemoteProcessGroup() *VersionedRemoteProcessGroup`

NewVersionedRemoteProcessGroup instantiates a new VersionedRemoteProcessGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionedRemoteProcessGroupWithDefaults

`func NewVersionedRemoteProcessGroupWithDefaults() *VersionedRemoteProcessGroup`

NewVersionedRemoteProcessGroupWithDefaults instantiates a new VersionedRemoteProcessGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *VersionedRemoteProcessGroup) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *VersionedRemoteProcessGroup) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *VersionedRemoteProcessGroup) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *VersionedRemoteProcessGroup) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *VersionedRemoteProcessGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VersionedRemoteProcessGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VersionedRemoteProcessGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VersionedRemoteProcessGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComments

`func (o *VersionedRemoteProcessGroup) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VersionedRemoteProcessGroup) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VersionedRemoteProcessGroup) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VersionedRemoteProcessGroup) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetPosition

`func (o *VersionedRemoteProcessGroup) GetPosition() Position`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *VersionedRemoteProcessGroup) GetPositionOk() (*Position, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *VersionedRemoteProcessGroup) SetPosition(v Position)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *VersionedRemoteProcessGroup) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetTargetUri

`func (o *VersionedRemoteProcessGroup) GetTargetUri() string`

GetTargetUri returns the TargetUri field if non-nil, zero value otherwise.

### GetTargetUriOk

`func (o *VersionedRemoteProcessGroup) GetTargetUriOk() (*string, bool)`

GetTargetUriOk returns a tuple with the TargetUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUri

`func (o *VersionedRemoteProcessGroup) SetTargetUri(v string)`

SetTargetUri sets TargetUri field to given value.

### HasTargetUri

`func (o *VersionedRemoteProcessGroup) HasTargetUri() bool`

HasTargetUri returns a boolean if a field has been set.

### GetTargetUris

`func (o *VersionedRemoteProcessGroup) GetTargetUris() string`

GetTargetUris returns the TargetUris field if non-nil, zero value otherwise.

### GetTargetUrisOk

`func (o *VersionedRemoteProcessGroup) GetTargetUrisOk() (*string, bool)`

GetTargetUrisOk returns a tuple with the TargetUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUris

`func (o *VersionedRemoteProcessGroup) SetTargetUris(v string)`

SetTargetUris sets TargetUris field to given value.

### HasTargetUris

`func (o *VersionedRemoteProcessGroup) HasTargetUris() bool`

HasTargetUris returns a boolean if a field has been set.

### GetCommunicationsTimeout

`func (o *VersionedRemoteProcessGroup) GetCommunicationsTimeout() string`

GetCommunicationsTimeout returns the CommunicationsTimeout field if non-nil, zero value otherwise.

### GetCommunicationsTimeoutOk

`func (o *VersionedRemoteProcessGroup) GetCommunicationsTimeoutOk() (*string, bool)`

GetCommunicationsTimeoutOk returns a tuple with the CommunicationsTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationsTimeout

`func (o *VersionedRemoteProcessGroup) SetCommunicationsTimeout(v string)`

SetCommunicationsTimeout sets CommunicationsTimeout field to given value.

### HasCommunicationsTimeout

`func (o *VersionedRemoteProcessGroup) HasCommunicationsTimeout() bool`

HasCommunicationsTimeout returns a boolean if a field has been set.

### GetYieldDuration

`func (o *VersionedRemoteProcessGroup) GetYieldDuration() string`

GetYieldDuration returns the YieldDuration field if non-nil, zero value otherwise.

### GetYieldDurationOk

`func (o *VersionedRemoteProcessGroup) GetYieldDurationOk() (*string, bool)`

GetYieldDurationOk returns a tuple with the YieldDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYieldDuration

`func (o *VersionedRemoteProcessGroup) SetYieldDuration(v string)`

SetYieldDuration sets YieldDuration field to given value.

### HasYieldDuration

`func (o *VersionedRemoteProcessGroup) HasYieldDuration() bool`

HasYieldDuration returns a boolean if a field has been set.

### GetTransportProtocol

`func (o *VersionedRemoteProcessGroup) GetTransportProtocol() string`

GetTransportProtocol returns the TransportProtocol field if non-nil, zero value otherwise.

### GetTransportProtocolOk

`func (o *VersionedRemoteProcessGroup) GetTransportProtocolOk() (*string, bool)`

GetTransportProtocolOk returns a tuple with the TransportProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportProtocol

`func (o *VersionedRemoteProcessGroup) SetTransportProtocol(v string)`

SetTransportProtocol sets TransportProtocol field to given value.

### HasTransportProtocol

`func (o *VersionedRemoteProcessGroup) HasTransportProtocol() bool`

HasTransportProtocol returns a boolean if a field has been set.

### GetLocalNetworkInterface

`func (o *VersionedRemoteProcessGroup) GetLocalNetworkInterface() string`

GetLocalNetworkInterface returns the LocalNetworkInterface field if non-nil, zero value otherwise.

### GetLocalNetworkInterfaceOk

`func (o *VersionedRemoteProcessGroup) GetLocalNetworkInterfaceOk() (*string, bool)`

GetLocalNetworkInterfaceOk returns a tuple with the LocalNetworkInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNetworkInterface

`func (o *VersionedRemoteProcessGroup) SetLocalNetworkInterface(v string)`

SetLocalNetworkInterface sets LocalNetworkInterface field to given value.

### HasLocalNetworkInterface

`func (o *VersionedRemoteProcessGroup) HasLocalNetworkInterface() bool`

HasLocalNetworkInterface returns a boolean if a field has been set.

### GetProxyHost

`func (o *VersionedRemoteProcessGroup) GetProxyHost() string`

GetProxyHost returns the ProxyHost field if non-nil, zero value otherwise.

### GetProxyHostOk

`func (o *VersionedRemoteProcessGroup) GetProxyHostOk() (*string, bool)`

GetProxyHostOk returns a tuple with the ProxyHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyHost

`func (o *VersionedRemoteProcessGroup) SetProxyHost(v string)`

SetProxyHost sets ProxyHost field to given value.

### HasProxyHost

`func (o *VersionedRemoteProcessGroup) HasProxyHost() bool`

HasProxyHost returns a boolean if a field has been set.

### GetProxyPort

`func (o *VersionedRemoteProcessGroup) GetProxyPort() int32`

GetProxyPort returns the ProxyPort field if non-nil, zero value otherwise.

### GetProxyPortOk

`func (o *VersionedRemoteProcessGroup) GetProxyPortOk() (*int32, bool)`

GetProxyPortOk returns a tuple with the ProxyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPort

`func (o *VersionedRemoteProcessGroup) SetProxyPort(v int32)`

SetProxyPort sets ProxyPort field to given value.

### HasProxyPort

`func (o *VersionedRemoteProcessGroup) HasProxyPort() bool`

HasProxyPort returns a boolean if a field has been set.

### GetProxyUser

`func (o *VersionedRemoteProcessGroup) GetProxyUser() string`

GetProxyUser returns the ProxyUser field if non-nil, zero value otherwise.

### GetProxyUserOk

`func (o *VersionedRemoteProcessGroup) GetProxyUserOk() (*string, bool)`

GetProxyUserOk returns a tuple with the ProxyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUser

`func (o *VersionedRemoteProcessGroup) SetProxyUser(v string)`

SetProxyUser sets ProxyUser field to given value.

### HasProxyUser

`func (o *VersionedRemoteProcessGroup) HasProxyUser() bool`

HasProxyUser returns a boolean if a field has been set.

### GetInputPorts

`func (o *VersionedRemoteProcessGroup) GetInputPorts() []VersionedRemoteGroupPort`

GetInputPorts returns the InputPorts field if non-nil, zero value otherwise.

### GetInputPortsOk

`func (o *VersionedRemoteProcessGroup) GetInputPortsOk() (*[]VersionedRemoteGroupPort, bool)`

GetInputPortsOk returns a tuple with the InputPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPorts

`func (o *VersionedRemoteProcessGroup) SetInputPorts(v []VersionedRemoteGroupPort)`

SetInputPorts sets InputPorts field to given value.

### HasInputPorts

`func (o *VersionedRemoteProcessGroup) HasInputPorts() bool`

HasInputPorts returns a boolean if a field has been set.

### GetOutputPorts

`func (o *VersionedRemoteProcessGroup) GetOutputPorts() []VersionedRemoteGroupPort`

GetOutputPorts returns the OutputPorts field if non-nil, zero value otherwise.

### GetOutputPortsOk

`func (o *VersionedRemoteProcessGroup) GetOutputPortsOk() (*[]VersionedRemoteGroupPort, bool)`

GetOutputPortsOk returns a tuple with the OutputPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPorts

`func (o *VersionedRemoteProcessGroup) SetOutputPorts(v []VersionedRemoteGroupPort)`

SetOutputPorts sets OutputPorts field to given value.

### HasOutputPorts

`func (o *VersionedRemoteProcessGroup) HasOutputPorts() bool`

HasOutputPorts returns a boolean if a field has been set.

### GetComponentType

`func (o *VersionedRemoteProcessGroup) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *VersionedRemoteProcessGroup) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *VersionedRemoteProcessGroup) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *VersionedRemoteProcessGroup) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### GetGroupIdentifier

`func (o *VersionedRemoteProcessGroup) GetGroupIdentifier() string`

GetGroupIdentifier returns the GroupIdentifier field if non-nil, zero value otherwise.

### GetGroupIdentifierOk

`func (o *VersionedRemoteProcessGroup) GetGroupIdentifierOk() (*string, bool)`

GetGroupIdentifierOk returns a tuple with the GroupIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIdentifier

`func (o *VersionedRemoteProcessGroup) SetGroupIdentifier(v string)`

SetGroupIdentifier sets GroupIdentifier field to given value.

### HasGroupIdentifier

`func (o *VersionedRemoteProcessGroup) HasGroupIdentifier() bool`

HasGroupIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


