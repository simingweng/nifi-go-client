# RemoteProcessGroupStatusSnapshotEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the remote process group. | [optional] 
**RemoteProcessGroupStatusSnapshot** | Pointer to [**RemoteProcessGroupStatusSnapshotDTO**](RemoteProcessGroupStatusSnapshotDTO.md) |  | [optional] 
**CanRead** | Pointer to **bool** | Indicates whether the user can read a given resource. | [optional] [readonly] 

## Methods

### NewRemoteProcessGroupStatusSnapshotEntity

`func NewRemoteProcessGroupStatusSnapshotEntity() *RemoteProcessGroupStatusSnapshotEntity`

NewRemoteProcessGroupStatusSnapshotEntity instantiates a new RemoteProcessGroupStatusSnapshotEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteProcessGroupStatusSnapshotEntityWithDefaults

`func NewRemoteProcessGroupStatusSnapshotEntityWithDefaults() *RemoteProcessGroupStatusSnapshotEntity`

NewRemoteProcessGroupStatusSnapshotEntityWithDefaults instantiates a new RemoteProcessGroupStatusSnapshotEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RemoteProcessGroupStatusSnapshotEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoteProcessGroupStatusSnapshotEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoteProcessGroupStatusSnapshotEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RemoteProcessGroupStatusSnapshotEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteProcessGroupStatusSnapshot

`func (o *RemoteProcessGroupStatusSnapshotEntity) GetRemoteProcessGroupStatusSnapshot() RemoteProcessGroupStatusSnapshotDTO`

GetRemoteProcessGroupStatusSnapshot returns the RemoteProcessGroupStatusSnapshot field if non-nil, zero value otherwise.

### GetRemoteProcessGroupStatusSnapshotOk

`func (o *RemoteProcessGroupStatusSnapshotEntity) GetRemoteProcessGroupStatusSnapshotOk() (*RemoteProcessGroupStatusSnapshotDTO, bool)`

GetRemoteProcessGroupStatusSnapshotOk returns a tuple with the RemoteProcessGroupStatusSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProcessGroupStatusSnapshot

`func (o *RemoteProcessGroupStatusSnapshotEntity) SetRemoteProcessGroupStatusSnapshot(v RemoteProcessGroupStatusSnapshotDTO)`

SetRemoteProcessGroupStatusSnapshot sets RemoteProcessGroupStatusSnapshot field to given value.

### HasRemoteProcessGroupStatusSnapshot

`func (o *RemoteProcessGroupStatusSnapshotEntity) HasRemoteProcessGroupStatusSnapshot() bool`

HasRemoteProcessGroupStatusSnapshot returns a boolean if a field has been set.

### GetCanRead

`func (o *RemoteProcessGroupStatusSnapshotEntity) GetCanRead() bool`

GetCanRead returns the CanRead field if non-nil, zero value otherwise.

### GetCanReadOk

`func (o *RemoteProcessGroupStatusSnapshotEntity) GetCanReadOk() (*bool, bool)`

GetCanReadOk returns a tuple with the CanRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRead

`func (o *RemoteProcessGroupStatusSnapshotEntity) SetCanRead(v bool)`

SetCanRead sets CanRead field to given value.

### HasCanRead

`func (o *RemoteProcessGroupStatusSnapshotEntity) HasCanRead() bool`

HasCanRead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


