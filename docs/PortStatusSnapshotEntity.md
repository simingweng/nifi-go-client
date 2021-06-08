# PortStatusSnapshotEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the port. | [optional] 
**PortStatusSnapshot** | Pointer to [**PortStatusSnapshotDTO**](PortStatusSnapshotDTO.md) |  | [optional] 
**CanRead** | Pointer to **bool** | Indicates whether the user can read a given resource. | [optional] [readonly] 

## Methods

### NewPortStatusSnapshotEntity

`func NewPortStatusSnapshotEntity() *PortStatusSnapshotEntity`

NewPortStatusSnapshotEntity instantiates a new PortStatusSnapshotEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortStatusSnapshotEntityWithDefaults

`func NewPortStatusSnapshotEntityWithDefaults() *PortStatusSnapshotEntity`

NewPortStatusSnapshotEntityWithDefaults instantiates a new PortStatusSnapshotEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortStatusSnapshotEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortStatusSnapshotEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortStatusSnapshotEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortStatusSnapshotEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPortStatusSnapshot

`func (o *PortStatusSnapshotEntity) GetPortStatusSnapshot() PortStatusSnapshotDTO`

GetPortStatusSnapshot returns the PortStatusSnapshot field if non-nil, zero value otherwise.

### GetPortStatusSnapshotOk

`func (o *PortStatusSnapshotEntity) GetPortStatusSnapshotOk() (*PortStatusSnapshotDTO, bool)`

GetPortStatusSnapshotOk returns a tuple with the PortStatusSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStatusSnapshot

`func (o *PortStatusSnapshotEntity) SetPortStatusSnapshot(v PortStatusSnapshotDTO)`

SetPortStatusSnapshot sets PortStatusSnapshot field to given value.

### HasPortStatusSnapshot

`func (o *PortStatusSnapshotEntity) HasPortStatusSnapshot() bool`

HasPortStatusSnapshot returns a boolean if a field has been set.

### GetCanRead

`func (o *PortStatusSnapshotEntity) GetCanRead() bool`

GetCanRead returns the CanRead field if non-nil, zero value otherwise.

### GetCanReadOk

`func (o *PortStatusSnapshotEntity) GetCanReadOk() (*bool, bool)`

GetCanReadOk returns a tuple with the CanRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRead

`func (o *PortStatusSnapshotEntity) SetCanRead(v bool)`

SetCanRead sets CanRead field to given value.

### HasCanRead

`func (o *PortStatusSnapshotEntity) HasCanRead() bool`

HasCanRead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


