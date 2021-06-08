# ConnectionStatisticsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionStatistics** | Pointer to [**ConnectionStatisticsDTO**](ConnectionStatisticsDTO.md) |  | [optional] 
**CanRead** | Pointer to **bool** | Indicates whether the user can read a given resource. | [optional] [readonly] 

## Methods

### NewConnectionStatisticsEntity

`func NewConnectionStatisticsEntity() *ConnectionStatisticsEntity`

NewConnectionStatisticsEntity instantiates a new ConnectionStatisticsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionStatisticsEntityWithDefaults

`func NewConnectionStatisticsEntityWithDefaults() *ConnectionStatisticsEntity`

NewConnectionStatisticsEntityWithDefaults instantiates a new ConnectionStatisticsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionStatistics

`func (o *ConnectionStatisticsEntity) GetConnectionStatistics() ConnectionStatisticsDTO`

GetConnectionStatistics returns the ConnectionStatistics field if non-nil, zero value otherwise.

### GetConnectionStatisticsOk

`func (o *ConnectionStatisticsEntity) GetConnectionStatisticsOk() (*ConnectionStatisticsDTO, bool)`

GetConnectionStatisticsOk returns a tuple with the ConnectionStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatistics

`func (o *ConnectionStatisticsEntity) SetConnectionStatistics(v ConnectionStatisticsDTO)`

SetConnectionStatistics sets ConnectionStatistics field to given value.

### HasConnectionStatistics

`func (o *ConnectionStatisticsEntity) HasConnectionStatistics() bool`

HasConnectionStatistics returns a boolean if a field has been set.

### GetCanRead

`func (o *ConnectionStatisticsEntity) GetCanRead() bool`

GetCanRead returns the CanRead field if non-nil, zero value otherwise.

### GetCanReadOk

`func (o *ConnectionStatisticsEntity) GetCanReadOk() (*bool, bool)`

GetCanReadOk returns a tuple with the CanRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRead

`func (o *ConnectionStatisticsEntity) SetCanRead(v bool)`

SetCanRead sets CanRead field to given value.

### HasCanRead

`func (o *ConnectionStatisticsEntity) HasCanRead() bool`

HasCanRead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


