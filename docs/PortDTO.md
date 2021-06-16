# PortDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the component. | [optional] 
**VersionedComponentId** | Pointer to **string** | The ID of the corresponding component that is under version control | [optional] 
**ParentGroupId** | Pointer to **string** | The id of parent process group of this component if applicable. | [optional] 
**Position** | Pointer to [**PositionDTO**](PositionDTO.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the port. | [optional] 
**Comments** | Pointer to **string** | The comments for the port. | [optional] 
**State** | Pointer to **string** | The state of the port. | [optional] 
**Type** | Pointer to **string** | The type of port. | [optional] 
**Transmitting** | Pointer to **bool** | Whether the port has incoming or output connections to a remote NiFi. This is only applicable when the port is allowed to be accessed remotely. | [optional] 
**ConcurrentlySchedulableTaskCount** | Pointer to **int32** | The number of tasks that should be concurrently scheduled for the port. | [optional] 
**UserAccessControl** | Pointer to **[]string** | The users that are allowed to access the port. | [optional] 
**GroupAccessControl** | Pointer to **[]string** | The user groups that are allowed to access the port. | [optional] 
**AllowRemoteAccess** | Pointer to **bool** | Whether this port can be accessed remotely via Site-to-Site protocol. | [optional] 
**ValidationErrors** | Pointer to **[]string** | Gets the validation errors from this port. These validation errors represent the problems with the port that must be resolved before it can be started. | [optional] 

## Methods

### NewPortDTO

`func NewPortDTO() *PortDTO`

NewPortDTO instantiates a new PortDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortDTOWithDefaults

`func NewPortDTOWithDefaults() *PortDTO`

NewPortDTOWithDefaults instantiates a new PortDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersionedComponentId

`func (o *PortDTO) GetVersionedComponentId() string`

GetVersionedComponentId returns the VersionedComponentId field if non-nil, zero value otherwise.

### GetVersionedComponentIdOk

`func (o *PortDTO) GetVersionedComponentIdOk() (*string, bool)`

GetVersionedComponentIdOk returns a tuple with the VersionedComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedComponentId

`func (o *PortDTO) SetVersionedComponentId(v string)`

SetVersionedComponentId sets VersionedComponentId field to given value.

### HasVersionedComponentId

`func (o *PortDTO) HasVersionedComponentId() bool`

HasVersionedComponentId returns a boolean if a field has been set.

### GetParentGroupId

`func (o *PortDTO) GetParentGroupId() string`

GetParentGroupId returns the ParentGroupId field if non-nil, zero value otherwise.

### GetParentGroupIdOk

`func (o *PortDTO) GetParentGroupIdOk() (*string, bool)`

GetParentGroupIdOk returns a tuple with the ParentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroupId

`func (o *PortDTO) SetParentGroupId(v string)`

SetParentGroupId sets ParentGroupId field to given value.

### HasParentGroupId

`func (o *PortDTO) HasParentGroupId() bool`

HasParentGroupId returns a boolean if a field has been set.

### GetPosition

`func (o *PortDTO) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PortDTO) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PortDTO) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PortDTO) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetName

`func (o *PortDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PortDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComments

`func (o *PortDTO) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PortDTO) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PortDTO) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PortDTO) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetState

`func (o *PortDTO) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PortDTO) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PortDTO) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PortDTO) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *PortDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PortDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PortDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PortDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTransmitting

`func (o *PortDTO) GetTransmitting() bool`

GetTransmitting returns the Transmitting field if non-nil, zero value otherwise.

### GetTransmittingOk

`func (o *PortDTO) GetTransmittingOk() (*bool, bool)`

GetTransmittingOk returns a tuple with the Transmitting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitting

`func (o *PortDTO) SetTransmitting(v bool)`

SetTransmitting sets Transmitting field to given value.

### HasTransmitting

`func (o *PortDTO) HasTransmitting() bool`

HasTransmitting returns a boolean if a field has been set.

### GetConcurrentlySchedulableTaskCount

`func (o *PortDTO) GetConcurrentlySchedulableTaskCount() int32`

GetConcurrentlySchedulableTaskCount returns the ConcurrentlySchedulableTaskCount field if non-nil, zero value otherwise.

### GetConcurrentlySchedulableTaskCountOk

`func (o *PortDTO) GetConcurrentlySchedulableTaskCountOk() (*int32, bool)`

GetConcurrentlySchedulableTaskCountOk returns a tuple with the ConcurrentlySchedulableTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentlySchedulableTaskCount

`func (o *PortDTO) SetConcurrentlySchedulableTaskCount(v int32)`

SetConcurrentlySchedulableTaskCount sets ConcurrentlySchedulableTaskCount field to given value.

### HasConcurrentlySchedulableTaskCount

`func (o *PortDTO) HasConcurrentlySchedulableTaskCount() bool`

HasConcurrentlySchedulableTaskCount returns a boolean if a field has been set.

### GetUserAccessControl

`func (o *PortDTO) GetUserAccessControl() []string`

GetUserAccessControl returns the UserAccessControl field if non-nil, zero value otherwise.

### GetUserAccessControlOk

`func (o *PortDTO) GetUserAccessControlOk() (*[]string, bool)`

GetUserAccessControlOk returns a tuple with the UserAccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccessControl

`func (o *PortDTO) SetUserAccessControl(v []string)`

SetUserAccessControl sets UserAccessControl field to given value.

### HasUserAccessControl

`func (o *PortDTO) HasUserAccessControl() bool`

HasUserAccessControl returns a boolean if a field has been set.

### GetGroupAccessControl

`func (o *PortDTO) GetGroupAccessControl() []string`

GetGroupAccessControl returns the GroupAccessControl field if non-nil, zero value otherwise.

### GetGroupAccessControlOk

`func (o *PortDTO) GetGroupAccessControlOk() (*[]string, bool)`

GetGroupAccessControlOk returns a tuple with the GroupAccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAccessControl

`func (o *PortDTO) SetGroupAccessControl(v []string)`

SetGroupAccessControl sets GroupAccessControl field to given value.

### HasGroupAccessControl

`func (o *PortDTO) HasGroupAccessControl() bool`

HasGroupAccessControl returns a boolean if a field has been set.

### GetAllowRemoteAccess

`func (o *PortDTO) GetAllowRemoteAccess() bool`

GetAllowRemoteAccess returns the AllowRemoteAccess field if non-nil, zero value otherwise.

### GetAllowRemoteAccessOk

`func (o *PortDTO) GetAllowRemoteAccessOk() (*bool, bool)`

GetAllowRemoteAccessOk returns a tuple with the AllowRemoteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRemoteAccess

`func (o *PortDTO) SetAllowRemoteAccess(v bool)`

SetAllowRemoteAccess sets AllowRemoteAccess field to given value.

### HasAllowRemoteAccess

`func (o *PortDTO) HasAllowRemoteAccess() bool`

HasAllowRemoteAccess returns a boolean if a field has been set.

### GetValidationErrors

`func (o *PortDTO) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *PortDTO) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *PortDTO) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *PortDTO) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


