# ConnectionsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | Pointer to [**[]ConnectionEntity**](ConnectionEntity.md) |  | [optional] 

## Methods

### NewConnectionsEntity

`func NewConnectionsEntity() *ConnectionsEntity`

NewConnectionsEntity instantiates a new ConnectionsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionsEntityWithDefaults

`func NewConnectionsEntityWithDefaults() *ConnectionsEntity`

NewConnectionsEntityWithDefaults instantiates a new ConnectionsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *ConnectionsEntity) GetConnections() []ConnectionEntity`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *ConnectionsEntity) GetConnectionsOk() (*[]ConnectionEntity, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *ConnectionsEntity) SetConnections(v []ConnectionEntity)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *ConnectionsEntity) HasConnections() bool`

HasConnections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


