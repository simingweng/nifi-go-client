# ConnectableDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the connectable component. | 
**VersionedComponentId** | Pointer to **string** | The ID of the corresponding component that is under version control | [optional] 
**Type** | **string** | The type of component the connectable is. | 
**GroupId** | **string** | The id of the group that the connectable component resides in | 
**Name** | Pointer to **string** | The name of the connectable component | [optional] 
**Running** | Pointer to **bool** | Reflects the current state of the connectable component. | [optional] 
**Transmitting** | Pointer to **bool** | If the connectable component represents a remote port, indicates if the target is configured to transmit. | [optional] 
**Exists** | Pointer to **bool** | If the connectable component represents a remote port, indicates if the target exists. | [optional] 
**Comments** | Pointer to **string** | The comments for the connectable component. | [optional] 

## Methods

### NewConnectableDTO

`func NewConnectableDTO(id string, type_ string, groupId string, ) *ConnectableDTO`

NewConnectableDTO instantiates a new ConnectableDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectableDTOWithDefaults

`func NewConnectableDTOWithDefaults() *ConnectableDTO`

NewConnectableDTOWithDefaults instantiates a new ConnectableDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectableDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectableDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectableDTO) SetId(v string)`

SetId sets Id field to given value.


### GetVersionedComponentId

`func (o *ConnectableDTO) GetVersionedComponentId() string`

GetVersionedComponentId returns the VersionedComponentId field if non-nil, zero value otherwise.

### GetVersionedComponentIdOk

`func (o *ConnectableDTO) GetVersionedComponentIdOk() (*string, bool)`

GetVersionedComponentIdOk returns a tuple with the VersionedComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedComponentId

`func (o *ConnectableDTO) SetVersionedComponentId(v string)`

SetVersionedComponentId sets VersionedComponentId field to given value.

### HasVersionedComponentId

`func (o *ConnectableDTO) HasVersionedComponentId() bool`

HasVersionedComponentId returns a boolean if a field has been set.

### GetType

`func (o *ConnectableDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectableDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectableDTO) SetType(v string)`

SetType sets Type field to given value.


### GetGroupId

`func (o *ConnectableDTO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ConnectableDTO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ConnectableDTO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetName

`func (o *ConnectableDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectableDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectableDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectableDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRunning

`func (o *ConnectableDTO) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *ConnectableDTO) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *ConnectableDTO) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *ConnectableDTO) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetTransmitting

`func (o *ConnectableDTO) GetTransmitting() bool`

GetTransmitting returns the Transmitting field if non-nil, zero value otherwise.

### GetTransmittingOk

`func (o *ConnectableDTO) GetTransmittingOk() (*bool, bool)`

GetTransmittingOk returns a tuple with the Transmitting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitting

`func (o *ConnectableDTO) SetTransmitting(v bool)`

SetTransmitting sets Transmitting field to given value.

### HasTransmitting

`func (o *ConnectableDTO) HasTransmitting() bool`

HasTransmitting returns a boolean if a field has been set.

### GetExists

`func (o *ConnectableDTO) GetExists() bool`

GetExists returns the Exists field if non-nil, zero value otherwise.

### GetExistsOk

`func (o *ConnectableDTO) GetExistsOk() (*bool, bool)`

GetExistsOk returns a tuple with the Exists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExists

`func (o *ConnectableDTO) SetExists(v bool)`

SetExists sets Exists field to given value.

### HasExists

`func (o *ConnectableDTO) HasExists() bool`

HasExists returns a boolean if a field has been set.

### GetComments

`func (o *ConnectableDTO) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ConnectableDTO) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ConnectableDTO) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ConnectableDTO) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


