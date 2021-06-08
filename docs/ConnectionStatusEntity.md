# ConnectionStatusEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionStatus** | Pointer to [**ConnectionStatusDTO**](ConnectionStatusDTO.md) |  | [optional] 
**CanRead** | Pointer to **bool** | Indicates whether the user can read a given resource. | [optional] [readonly] 

## Methods

### NewConnectionStatusEntity

`func NewConnectionStatusEntity() *ConnectionStatusEntity`

NewConnectionStatusEntity instantiates a new ConnectionStatusEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionStatusEntityWithDefaults

`func NewConnectionStatusEntityWithDefaults() *ConnectionStatusEntity`

NewConnectionStatusEntityWithDefaults instantiates a new ConnectionStatusEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionStatus

`func (o *ConnectionStatusEntity) GetConnectionStatus() ConnectionStatusDTO`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *ConnectionStatusEntity) GetConnectionStatusOk() (*ConnectionStatusDTO, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *ConnectionStatusEntity) SetConnectionStatus(v ConnectionStatusDTO)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *ConnectionStatusEntity) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetCanRead

`func (o *ConnectionStatusEntity) GetCanRead() bool`

GetCanRead returns the CanRead field if non-nil, zero value otherwise.

### GetCanReadOk

`func (o *ConnectionStatusEntity) GetCanReadOk() (*bool, bool)`

GetCanReadOk returns a tuple with the CanRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRead

`func (o *ConnectionStatusEntity) SetCanRead(v bool)`

SetCanRead sets CanRead field to given value.

### HasCanRead

`func (o *ConnectionStatusEntity) HasCanRead() bool`

HasCanRead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


