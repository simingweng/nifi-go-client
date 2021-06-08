# RemoteProcessGroupPortDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the port. | [optional] 
**TargetId** | Pointer to **string** | The id of the target port. | [optional] 
**VersionedComponentId** | Pointer to **string** | The ID of the corresponding component that is under version control | [optional] 
**GroupId** | Pointer to **string** | The id of the remote process group that the port resides in. | [optional] 
**Name** | Pointer to **string** | The name of the target port. | [optional] 
**Comments** | Pointer to **string** | The comments as configured on the target port. | [optional] 
**ConcurrentlySchedulableTaskCount** | Pointer to **int32** | The number of task that may transmit flowfiles to the target port concurrently. | [optional] 
**Transmitting** | Pointer to **bool** | Whether the remote port is configured for transmission. | [optional] 
**UseCompression** | Pointer to **bool** | Whether the flowfiles are compressed when sent to the target port. | [optional] 
**Exists** | Pointer to **bool** | Whether the target port exists. | [optional] 
**TargetRunning** | Pointer to **bool** | Whether the target port is running. | [optional] 
**Connected** | Pointer to **bool** | Whether the port has either an incoming or outgoing connection. | [optional] 
**BatchSettings** | Pointer to [**BatchSettingsDTO**](BatchSettingsDTO.md) |  | [optional] 

## Methods

### NewRemoteProcessGroupPortDTO

`func NewRemoteProcessGroupPortDTO() *RemoteProcessGroupPortDTO`

NewRemoteProcessGroupPortDTO instantiates a new RemoteProcessGroupPortDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteProcessGroupPortDTOWithDefaults

`func NewRemoteProcessGroupPortDTOWithDefaults() *RemoteProcessGroupPortDTO`

NewRemoteProcessGroupPortDTOWithDefaults instantiates a new RemoteProcessGroupPortDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RemoteProcessGroupPortDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoteProcessGroupPortDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoteProcessGroupPortDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RemoteProcessGroupPortDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTargetId

`func (o *RemoteProcessGroupPortDTO) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *RemoteProcessGroupPortDTO) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *RemoteProcessGroupPortDTO) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *RemoteProcessGroupPortDTO) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetVersionedComponentId

`func (o *RemoteProcessGroupPortDTO) GetVersionedComponentId() string`

GetVersionedComponentId returns the VersionedComponentId field if non-nil, zero value otherwise.

### GetVersionedComponentIdOk

`func (o *RemoteProcessGroupPortDTO) GetVersionedComponentIdOk() (*string, bool)`

GetVersionedComponentIdOk returns a tuple with the VersionedComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedComponentId

`func (o *RemoteProcessGroupPortDTO) SetVersionedComponentId(v string)`

SetVersionedComponentId sets VersionedComponentId field to given value.

### HasVersionedComponentId

`func (o *RemoteProcessGroupPortDTO) HasVersionedComponentId() bool`

HasVersionedComponentId returns a boolean if a field has been set.

### GetGroupId

`func (o *RemoteProcessGroupPortDTO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *RemoteProcessGroupPortDTO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *RemoteProcessGroupPortDTO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *RemoteProcessGroupPortDTO) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *RemoteProcessGroupPortDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemoteProcessGroupPortDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemoteProcessGroupPortDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RemoteProcessGroupPortDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComments

`func (o *RemoteProcessGroupPortDTO) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *RemoteProcessGroupPortDTO) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *RemoteProcessGroupPortDTO) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *RemoteProcessGroupPortDTO) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetConcurrentlySchedulableTaskCount

`func (o *RemoteProcessGroupPortDTO) GetConcurrentlySchedulableTaskCount() int32`

GetConcurrentlySchedulableTaskCount returns the ConcurrentlySchedulableTaskCount field if non-nil, zero value otherwise.

### GetConcurrentlySchedulableTaskCountOk

`func (o *RemoteProcessGroupPortDTO) GetConcurrentlySchedulableTaskCountOk() (*int32, bool)`

GetConcurrentlySchedulableTaskCountOk returns a tuple with the ConcurrentlySchedulableTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentlySchedulableTaskCount

`func (o *RemoteProcessGroupPortDTO) SetConcurrentlySchedulableTaskCount(v int32)`

SetConcurrentlySchedulableTaskCount sets ConcurrentlySchedulableTaskCount field to given value.

### HasConcurrentlySchedulableTaskCount

`func (o *RemoteProcessGroupPortDTO) HasConcurrentlySchedulableTaskCount() bool`

HasConcurrentlySchedulableTaskCount returns a boolean if a field has been set.

### GetTransmitting

`func (o *RemoteProcessGroupPortDTO) GetTransmitting() bool`

GetTransmitting returns the Transmitting field if non-nil, zero value otherwise.

### GetTransmittingOk

`func (o *RemoteProcessGroupPortDTO) GetTransmittingOk() (*bool, bool)`

GetTransmittingOk returns a tuple with the Transmitting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitting

`func (o *RemoteProcessGroupPortDTO) SetTransmitting(v bool)`

SetTransmitting sets Transmitting field to given value.

### HasTransmitting

`func (o *RemoteProcessGroupPortDTO) HasTransmitting() bool`

HasTransmitting returns a boolean if a field has been set.

### GetUseCompression

`func (o *RemoteProcessGroupPortDTO) GetUseCompression() bool`

GetUseCompression returns the UseCompression field if non-nil, zero value otherwise.

### GetUseCompressionOk

`func (o *RemoteProcessGroupPortDTO) GetUseCompressionOk() (*bool, bool)`

GetUseCompressionOk returns a tuple with the UseCompression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCompression

`func (o *RemoteProcessGroupPortDTO) SetUseCompression(v bool)`

SetUseCompression sets UseCompression field to given value.

### HasUseCompression

`func (o *RemoteProcessGroupPortDTO) HasUseCompression() bool`

HasUseCompression returns a boolean if a field has been set.

### GetExists

`func (o *RemoteProcessGroupPortDTO) GetExists() bool`

GetExists returns the Exists field if non-nil, zero value otherwise.

### GetExistsOk

`func (o *RemoteProcessGroupPortDTO) GetExistsOk() (*bool, bool)`

GetExistsOk returns a tuple with the Exists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExists

`func (o *RemoteProcessGroupPortDTO) SetExists(v bool)`

SetExists sets Exists field to given value.

### HasExists

`func (o *RemoteProcessGroupPortDTO) HasExists() bool`

HasExists returns a boolean if a field has been set.

### GetTargetRunning

`func (o *RemoteProcessGroupPortDTO) GetTargetRunning() bool`

GetTargetRunning returns the TargetRunning field if non-nil, zero value otherwise.

### GetTargetRunningOk

`func (o *RemoteProcessGroupPortDTO) GetTargetRunningOk() (*bool, bool)`

GetTargetRunningOk returns a tuple with the TargetRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRunning

`func (o *RemoteProcessGroupPortDTO) SetTargetRunning(v bool)`

SetTargetRunning sets TargetRunning field to given value.

### HasTargetRunning

`func (o *RemoteProcessGroupPortDTO) HasTargetRunning() bool`

HasTargetRunning returns a boolean if a field has been set.

### GetConnected

`func (o *RemoteProcessGroupPortDTO) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *RemoteProcessGroupPortDTO) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *RemoteProcessGroupPortDTO) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *RemoteProcessGroupPortDTO) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetBatchSettings

`func (o *RemoteProcessGroupPortDTO) GetBatchSettings() BatchSettingsDTO`

GetBatchSettings returns the BatchSettings field if non-nil, zero value otherwise.

### GetBatchSettingsOk

`func (o *RemoteProcessGroupPortDTO) GetBatchSettingsOk() (*BatchSettingsDTO, bool)`

GetBatchSettingsOk returns a tuple with the BatchSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSettings

`func (o *RemoteProcessGroupPortDTO) SetBatchSettings(v BatchSettingsDTO)`

SetBatchSettings sets BatchSettings field to given value.

### HasBatchSettings

`func (o *RemoteProcessGroupPortDTO) HasBatchSettings() bool`

HasBatchSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


