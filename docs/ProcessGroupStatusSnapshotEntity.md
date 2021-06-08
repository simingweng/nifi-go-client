# ProcessGroupStatusSnapshotEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the process group. | [optional] 
**ProcessGroupStatusSnapshot** | Pointer to [**ProcessGroupStatusSnapshotDTO**](ProcessGroupStatusSnapshotDTO.md) |  | [optional] 
**CanRead** | Pointer to **bool** | Indicates whether the user can read a given resource. | [optional] [readonly] 

## Methods

### NewProcessGroupStatusSnapshotEntity

`func NewProcessGroupStatusSnapshotEntity() *ProcessGroupStatusSnapshotEntity`

NewProcessGroupStatusSnapshotEntity instantiates a new ProcessGroupStatusSnapshotEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessGroupStatusSnapshotEntityWithDefaults

`func NewProcessGroupStatusSnapshotEntityWithDefaults() *ProcessGroupStatusSnapshotEntity`

NewProcessGroupStatusSnapshotEntityWithDefaults instantiates a new ProcessGroupStatusSnapshotEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProcessGroupStatusSnapshotEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessGroupStatusSnapshotEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessGroupStatusSnapshotEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProcessGroupStatusSnapshotEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProcessGroupStatusSnapshot

`func (o *ProcessGroupStatusSnapshotEntity) GetProcessGroupStatusSnapshot() ProcessGroupStatusSnapshotDTO`

GetProcessGroupStatusSnapshot returns the ProcessGroupStatusSnapshot field if non-nil, zero value otherwise.

### GetProcessGroupStatusSnapshotOk

`func (o *ProcessGroupStatusSnapshotEntity) GetProcessGroupStatusSnapshotOk() (*ProcessGroupStatusSnapshotDTO, bool)`

GetProcessGroupStatusSnapshotOk returns a tuple with the ProcessGroupStatusSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupStatusSnapshot

`func (o *ProcessGroupStatusSnapshotEntity) SetProcessGroupStatusSnapshot(v ProcessGroupStatusSnapshotDTO)`

SetProcessGroupStatusSnapshot sets ProcessGroupStatusSnapshot field to given value.

### HasProcessGroupStatusSnapshot

`func (o *ProcessGroupStatusSnapshotEntity) HasProcessGroupStatusSnapshot() bool`

HasProcessGroupStatusSnapshot returns a boolean if a field has been set.

### GetCanRead

`func (o *ProcessGroupStatusSnapshotEntity) GetCanRead() bool`

GetCanRead returns the CanRead field if non-nil, zero value otherwise.

### GetCanReadOk

`func (o *ProcessGroupStatusSnapshotEntity) GetCanReadOk() (*bool, bool)`

GetCanReadOk returns a tuple with the CanRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRead

`func (o *ProcessGroupStatusSnapshotEntity) SetCanRead(v bool)`

SetCanRead sets CanRead field to given value.

### HasCanRead

`func (o *ProcessGroupStatusSnapshotEntity) HasCanRead() bool`

HasCanRead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


