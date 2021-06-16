# ProcessorStatusSnapshotEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the processor. | [optional] 
**ProcessorStatusSnapshot** | Pointer to [**ProcessorStatusSnapshotDTO**](ProcessorStatusSnapshotDTO.md) |  | [optional] 
**CanRead** | Pointer to **bool** | Indicates whether the user can read a given resource. | [optional] [readonly] 

## Methods

### NewProcessorStatusSnapshotEntity

`func NewProcessorStatusSnapshotEntity() *ProcessorStatusSnapshotEntity`

NewProcessorStatusSnapshotEntity instantiates a new ProcessorStatusSnapshotEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorStatusSnapshotEntityWithDefaults

`func NewProcessorStatusSnapshotEntityWithDefaults() *ProcessorStatusSnapshotEntity`

NewProcessorStatusSnapshotEntityWithDefaults instantiates a new ProcessorStatusSnapshotEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProcessorStatusSnapshotEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessorStatusSnapshotEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessorStatusSnapshotEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProcessorStatusSnapshotEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProcessorStatusSnapshot

`func (o *ProcessorStatusSnapshotEntity) GetProcessorStatusSnapshot() ProcessorStatusSnapshotDTO`

GetProcessorStatusSnapshot returns the ProcessorStatusSnapshot field if non-nil, zero value otherwise.

### GetProcessorStatusSnapshotOk

`func (o *ProcessorStatusSnapshotEntity) GetProcessorStatusSnapshotOk() (*ProcessorStatusSnapshotDTO, bool)`

GetProcessorStatusSnapshotOk returns a tuple with the ProcessorStatusSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorStatusSnapshot

`func (o *ProcessorStatusSnapshotEntity) SetProcessorStatusSnapshot(v ProcessorStatusSnapshotDTO)`

SetProcessorStatusSnapshot sets ProcessorStatusSnapshot field to given value.

### HasProcessorStatusSnapshot

`func (o *ProcessorStatusSnapshotEntity) HasProcessorStatusSnapshot() bool`

HasProcessorStatusSnapshot returns a boolean if a field has been set.

### GetCanRead

`func (o *ProcessorStatusSnapshotEntity) GetCanRead() bool`

GetCanRead returns the CanRead field if non-nil, zero value otherwise.

### GetCanReadOk

`func (o *ProcessorStatusSnapshotEntity) GetCanReadOk() (*bool, bool)`

GetCanReadOk returns a tuple with the CanRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRead

`func (o *ProcessorStatusSnapshotEntity) SetCanRead(v bool)`

SetCanRead sets CanRead field to given value.

### HasCanRead

`func (o *ProcessorStatusSnapshotEntity) HasCanRead() bool`

HasCanRead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


