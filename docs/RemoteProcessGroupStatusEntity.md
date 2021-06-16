# RemoteProcessGroupStatusEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteProcessGroupStatus** | Pointer to [**RemoteProcessGroupStatusDTO**](RemoteProcessGroupStatusDTO.md) |  | [optional] 
**CanRead** | Pointer to **bool** | Indicates whether the user can read a given resource. | [optional] [readonly] 

## Methods

### NewRemoteProcessGroupStatusEntity

`func NewRemoteProcessGroupStatusEntity() *RemoteProcessGroupStatusEntity`

NewRemoteProcessGroupStatusEntity instantiates a new RemoteProcessGroupStatusEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteProcessGroupStatusEntityWithDefaults

`func NewRemoteProcessGroupStatusEntityWithDefaults() *RemoteProcessGroupStatusEntity`

NewRemoteProcessGroupStatusEntityWithDefaults instantiates a new RemoteProcessGroupStatusEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteProcessGroupStatus

`func (o *RemoteProcessGroupStatusEntity) GetRemoteProcessGroupStatus() RemoteProcessGroupStatusDTO`

GetRemoteProcessGroupStatus returns the RemoteProcessGroupStatus field if non-nil, zero value otherwise.

### GetRemoteProcessGroupStatusOk

`func (o *RemoteProcessGroupStatusEntity) GetRemoteProcessGroupStatusOk() (*RemoteProcessGroupStatusDTO, bool)`

GetRemoteProcessGroupStatusOk returns a tuple with the RemoteProcessGroupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProcessGroupStatus

`func (o *RemoteProcessGroupStatusEntity) SetRemoteProcessGroupStatus(v RemoteProcessGroupStatusDTO)`

SetRemoteProcessGroupStatus sets RemoteProcessGroupStatus field to given value.

### HasRemoteProcessGroupStatus

`func (o *RemoteProcessGroupStatusEntity) HasRemoteProcessGroupStatus() bool`

HasRemoteProcessGroupStatus returns a boolean if a field has been set.

### GetCanRead

`func (o *RemoteProcessGroupStatusEntity) GetCanRead() bool`

GetCanRead returns the CanRead field if non-nil, zero value otherwise.

### GetCanReadOk

`func (o *RemoteProcessGroupStatusEntity) GetCanReadOk() (*bool, bool)`

GetCanReadOk returns a tuple with the CanRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRead

`func (o *RemoteProcessGroupStatusEntity) SetCanRead(v bool)`

SetCanRead sets CanRead field to given value.

### HasCanRead

`func (o *RemoteProcessGroupStatusEntity) HasCanRead() bool`

HasCanRead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


