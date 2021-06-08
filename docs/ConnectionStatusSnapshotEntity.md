# ConnectionStatusSnapshotEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the connection. | [optional] 
**ConnectionStatusSnapshot** | Pointer to [**ConnectionStatusSnapshotDTO**](ConnectionStatusSnapshotDTO.md) |  | [optional] 
**CanRead** | Pointer to **bool** | Indicates whether the user can read a given resource. | [optional] [readonly] 

## Methods

### NewConnectionStatusSnapshotEntity

`func NewConnectionStatusSnapshotEntity() *ConnectionStatusSnapshotEntity`

NewConnectionStatusSnapshotEntity instantiates a new ConnectionStatusSnapshotEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionStatusSnapshotEntityWithDefaults

`func NewConnectionStatusSnapshotEntityWithDefaults() *ConnectionStatusSnapshotEntity`

NewConnectionStatusSnapshotEntityWithDefaults instantiates a new ConnectionStatusSnapshotEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectionStatusSnapshotEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectionStatusSnapshotEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectionStatusSnapshotEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectionStatusSnapshotEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConnectionStatusSnapshot

`func (o *ConnectionStatusSnapshotEntity) GetConnectionStatusSnapshot() ConnectionStatusSnapshotDTO`

GetConnectionStatusSnapshot returns the ConnectionStatusSnapshot field if non-nil, zero value otherwise.

### GetConnectionStatusSnapshotOk

`func (o *ConnectionStatusSnapshotEntity) GetConnectionStatusSnapshotOk() (*ConnectionStatusSnapshotDTO, bool)`

GetConnectionStatusSnapshotOk returns a tuple with the ConnectionStatusSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatusSnapshot

`func (o *ConnectionStatusSnapshotEntity) SetConnectionStatusSnapshot(v ConnectionStatusSnapshotDTO)`

SetConnectionStatusSnapshot sets ConnectionStatusSnapshot field to given value.

### HasConnectionStatusSnapshot

`func (o *ConnectionStatusSnapshotEntity) HasConnectionStatusSnapshot() bool`

HasConnectionStatusSnapshot returns a boolean if a field has been set.

### GetCanRead

`func (o *ConnectionStatusSnapshotEntity) GetCanRead() bool`

GetCanRead returns the CanRead field if non-nil, zero value otherwise.

### GetCanReadOk

`func (o *ConnectionStatusSnapshotEntity) GetCanReadOk() (*bool, bool)`

GetCanReadOk returns a tuple with the CanRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRead

`func (o *ConnectionStatusSnapshotEntity) SetCanRead(v bool)`

SetCanRead sets CanRead field to given value.

### HasCanRead

`func (o *ConnectionStatusSnapshotEntity) HasCanRead() bool`

HasCanRead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


