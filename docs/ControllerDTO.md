# ControllerDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the NiFi. | [optional] 
**Name** | Pointer to **string** | The name of the NiFi. | [optional] 
**Comments** | Pointer to **string** | The comments for the NiFi. | [optional] 
**RunningCount** | Pointer to **int32** | The number of running components in the NiFi. | [optional] 
**StoppedCount** | Pointer to **int32** | The number of stopped components in the NiFi. | [optional] 
**InvalidCount** | Pointer to **int32** | The number of invalid components in the NiFi. | [optional] 
**DisabledCount** | Pointer to **int32** | The number of disabled components in the NiFi. | [optional] 
**ActiveRemotePortCount** | Pointer to **int32** | The number of active remote ports contained in the NiFi. | [optional] 
**InactiveRemotePortCount** | Pointer to **int32** | The number of inactive remote ports contained in the NiFi. | [optional] 
**InputPortCount** | Pointer to **int32** | The number of input ports contained in the NiFi. | [optional] 
**OutputPortCount** | Pointer to **int32** | The number of output ports in the NiFi. | [optional] 
**RemoteSiteListeningPort** | Pointer to **int32** | The Socket Port on which this instance is listening for Remote Transfers of Flow Files. If this instance is not configured to receive Flow Files from remote instances, this will be null. | [optional] 
**RemoteSiteHttpListeningPort** | Pointer to **int32** | The HTTP(S) Port on which this instance is listening for Remote Transfers of Flow Files. If this instance is not configured to receive Flow Files from remote instances, this will be null. | [optional] 
**SiteToSiteSecure** | Pointer to **bool** | Indicates whether or not Site-to-Site communications with this instance is secure (2-way authentication). | [optional] 
**InstanceId** | Pointer to **string** | If clustered, the id of the Cluster Manager, otherwise the id of the NiFi. | [optional] 
**InputPorts** | Pointer to [**[]PortDTO**](PortDTO.md) | The input ports available to send data to for the NiFi. | [optional] 
**OutputPorts** | Pointer to [**[]PortDTO**](PortDTO.md) | The output ports available to received data from the NiFi. | [optional] 

## Methods

### NewControllerDTO

`func NewControllerDTO() *ControllerDTO`

NewControllerDTO instantiates a new ControllerDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerDTOWithDefaults

`func NewControllerDTOWithDefaults() *ControllerDTO`

NewControllerDTOWithDefaults instantiates a new ControllerDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ControllerDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControllerDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControllerDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ControllerDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ControllerDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControllerDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControllerDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ControllerDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComments

`func (o *ControllerDTO) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ControllerDTO) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ControllerDTO) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ControllerDTO) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetRunningCount

`func (o *ControllerDTO) GetRunningCount() int32`

GetRunningCount returns the RunningCount field if non-nil, zero value otherwise.

### GetRunningCountOk

`func (o *ControllerDTO) GetRunningCountOk() (*int32, bool)`

GetRunningCountOk returns a tuple with the RunningCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningCount

`func (o *ControllerDTO) SetRunningCount(v int32)`

SetRunningCount sets RunningCount field to given value.

### HasRunningCount

`func (o *ControllerDTO) HasRunningCount() bool`

HasRunningCount returns a boolean if a field has been set.

### GetStoppedCount

`func (o *ControllerDTO) GetStoppedCount() int32`

GetStoppedCount returns the StoppedCount field if non-nil, zero value otherwise.

### GetStoppedCountOk

`func (o *ControllerDTO) GetStoppedCountOk() (*int32, bool)`

GetStoppedCountOk returns a tuple with the StoppedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedCount

`func (o *ControllerDTO) SetStoppedCount(v int32)`

SetStoppedCount sets StoppedCount field to given value.

### HasStoppedCount

`func (o *ControllerDTO) HasStoppedCount() bool`

HasStoppedCount returns a boolean if a field has been set.

### GetInvalidCount

`func (o *ControllerDTO) GetInvalidCount() int32`

GetInvalidCount returns the InvalidCount field if non-nil, zero value otherwise.

### GetInvalidCountOk

`func (o *ControllerDTO) GetInvalidCountOk() (*int32, bool)`

GetInvalidCountOk returns a tuple with the InvalidCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidCount

`func (o *ControllerDTO) SetInvalidCount(v int32)`

SetInvalidCount sets InvalidCount field to given value.

### HasInvalidCount

`func (o *ControllerDTO) HasInvalidCount() bool`

HasInvalidCount returns a boolean if a field has been set.

### GetDisabledCount

`func (o *ControllerDTO) GetDisabledCount() int32`

GetDisabledCount returns the DisabledCount field if non-nil, zero value otherwise.

### GetDisabledCountOk

`func (o *ControllerDTO) GetDisabledCountOk() (*int32, bool)`

GetDisabledCountOk returns a tuple with the DisabledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledCount

`func (o *ControllerDTO) SetDisabledCount(v int32)`

SetDisabledCount sets DisabledCount field to given value.

### HasDisabledCount

`func (o *ControllerDTO) HasDisabledCount() bool`

HasDisabledCount returns a boolean if a field has been set.

### GetActiveRemotePortCount

`func (o *ControllerDTO) GetActiveRemotePortCount() int32`

GetActiveRemotePortCount returns the ActiveRemotePortCount field if non-nil, zero value otherwise.

### GetActiveRemotePortCountOk

`func (o *ControllerDTO) GetActiveRemotePortCountOk() (*int32, bool)`

GetActiveRemotePortCountOk returns a tuple with the ActiveRemotePortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRemotePortCount

`func (o *ControllerDTO) SetActiveRemotePortCount(v int32)`

SetActiveRemotePortCount sets ActiveRemotePortCount field to given value.

### HasActiveRemotePortCount

`func (o *ControllerDTO) HasActiveRemotePortCount() bool`

HasActiveRemotePortCount returns a boolean if a field has been set.

### GetInactiveRemotePortCount

`func (o *ControllerDTO) GetInactiveRemotePortCount() int32`

GetInactiveRemotePortCount returns the InactiveRemotePortCount field if non-nil, zero value otherwise.

### GetInactiveRemotePortCountOk

`func (o *ControllerDTO) GetInactiveRemotePortCountOk() (*int32, bool)`

GetInactiveRemotePortCountOk returns a tuple with the InactiveRemotePortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveRemotePortCount

`func (o *ControllerDTO) SetInactiveRemotePortCount(v int32)`

SetInactiveRemotePortCount sets InactiveRemotePortCount field to given value.

### HasInactiveRemotePortCount

`func (o *ControllerDTO) HasInactiveRemotePortCount() bool`

HasInactiveRemotePortCount returns a boolean if a field has been set.

### GetInputPortCount

`func (o *ControllerDTO) GetInputPortCount() int32`

GetInputPortCount returns the InputPortCount field if non-nil, zero value otherwise.

### GetInputPortCountOk

`func (o *ControllerDTO) GetInputPortCountOk() (*int32, bool)`

GetInputPortCountOk returns a tuple with the InputPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPortCount

`func (o *ControllerDTO) SetInputPortCount(v int32)`

SetInputPortCount sets InputPortCount field to given value.

### HasInputPortCount

`func (o *ControllerDTO) HasInputPortCount() bool`

HasInputPortCount returns a boolean if a field has been set.

### GetOutputPortCount

`func (o *ControllerDTO) GetOutputPortCount() int32`

GetOutputPortCount returns the OutputPortCount field if non-nil, zero value otherwise.

### GetOutputPortCountOk

`func (o *ControllerDTO) GetOutputPortCountOk() (*int32, bool)`

GetOutputPortCountOk returns a tuple with the OutputPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPortCount

`func (o *ControllerDTO) SetOutputPortCount(v int32)`

SetOutputPortCount sets OutputPortCount field to given value.

### HasOutputPortCount

`func (o *ControllerDTO) HasOutputPortCount() bool`

HasOutputPortCount returns a boolean if a field has been set.

### GetRemoteSiteListeningPort

`func (o *ControllerDTO) GetRemoteSiteListeningPort() int32`

GetRemoteSiteListeningPort returns the RemoteSiteListeningPort field if non-nil, zero value otherwise.

### GetRemoteSiteListeningPortOk

`func (o *ControllerDTO) GetRemoteSiteListeningPortOk() (*int32, bool)`

GetRemoteSiteListeningPortOk returns a tuple with the RemoteSiteListeningPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSiteListeningPort

`func (o *ControllerDTO) SetRemoteSiteListeningPort(v int32)`

SetRemoteSiteListeningPort sets RemoteSiteListeningPort field to given value.

### HasRemoteSiteListeningPort

`func (o *ControllerDTO) HasRemoteSiteListeningPort() bool`

HasRemoteSiteListeningPort returns a boolean if a field has been set.

### GetRemoteSiteHttpListeningPort

`func (o *ControllerDTO) GetRemoteSiteHttpListeningPort() int32`

GetRemoteSiteHttpListeningPort returns the RemoteSiteHttpListeningPort field if non-nil, zero value otherwise.

### GetRemoteSiteHttpListeningPortOk

`func (o *ControllerDTO) GetRemoteSiteHttpListeningPortOk() (*int32, bool)`

GetRemoteSiteHttpListeningPortOk returns a tuple with the RemoteSiteHttpListeningPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSiteHttpListeningPort

`func (o *ControllerDTO) SetRemoteSiteHttpListeningPort(v int32)`

SetRemoteSiteHttpListeningPort sets RemoteSiteHttpListeningPort field to given value.

### HasRemoteSiteHttpListeningPort

`func (o *ControllerDTO) HasRemoteSiteHttpListeningPort() bool`

HasRemoteSiteHttpListeningPort returns a boolean if a field has been set.

### GetSiteToSiteSecure

`func (o *ControllerDTO) GetSiteToSiteSecure() bool`

GetSiteToSiteSecure returns the SiteToSiteSecure field if non-nil, zero value otherwise.

### GetSiteToSiteSecureOk

`func (o *ControllerDTO) GetSiteToSiteSecureOk() (*bool, bool)`

GetSiteToSiteSecureOk returns a tuple with the SiteToSiteSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteToSiteSecure

`func (o *ControllerDTO) SetSiteToSiteSecure(v bool)`

SetSiteToSiteSecure sets SiteToSiteSecure field to given value.

### HasSiteToSiteSecure

`func (o *ControllerDTO) HasSiteToSiteSecure() bool`

HasSiteToSiteSecure returns a boolean if a field has been set.

### GetInstanceId

`func (o *ControllerDTO) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ControllerDTO) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ControllerDTO) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ControllerDTO) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInputPorts

`func (o *ControllerDTO) GetInputPorts() []PortDTO`

GetInputPorts returns the InputPorts field if non-nil, zero value otherwise.

### GetInputPortsOk

`func (o *ControllerDTO) GetInputPortsOk() (*[]PortDTO, bool)`

GetInputPortsOk returns a tuple with the InputPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPorts

`func (o *ControllerDTO) SetInputPorts(v []PortDTO)`

SetInputPorts sets InputPorts field to given value.

### HasInputPorts

`func (o *ControllerDTO) HasInputPorts() bool`

HasInputPorts returns a boolean if a field has been set.

### GetOutputPorts

`func (o *ControllerDTO) GetOutputPorts() []PortDTO`

GetOutputPorts returns the OutputPorts field if non-nil, zero value otherwise.

### GetOutputPortsOk

`func (o *ControllerDTO) GetOutputPortsOk() (*[]PortDTO, bool)`

GetOutputPortsOk returns a tuple with the OutputPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPorts

`func (o *ControllerDTO) SetOutputPorts(v []PortDTO)`

SetOutputPorts sets OutputPorts field to given value.

### HasOutputPorts

`func (o *ControllerDTO) HasOutputPorts() bool`

HasOutputPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


