# StatusHistoryEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusHistory** | Pointer to [**StatusHistoryDTO**](StatusHistoryDTO.md) |  | [optional] 
**CanRead** | Pointer to **bool** | Indicates whether the user can read a given resource. | [optional] [readonly] 

## Methods

### NewStatusHistoryEntity

`func NewStatusHistoryEntity() *StatusHistoryEntity`

NewStatusHistoryEntity instantiates a new StatusHistoryEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusHistoryEntityWithDefaults

`func NewStatusHistoryEntityWithDefaults() *StatusHistoryEntity`

NewStatusHistoryEntityWithDefaults instantiates a new StatusHistoryEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusHistory

`func (o *StatusHistoryEntity) GetStatusHistory() StatusHistoryDTO`

GetStatusHistory returns the StatusHistory field if non-nil, zero value otherwise.

### GetStatusHistoryOk

`func (o *StatusHistoryEntity) GetStatusHistoryOk() (*StatusHistoryDTO, bool)`

GetStatusHistoryOk returns a tuple with the StatusHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusHistory

`func (o *StatusHistoryEntity) SetStatusHistory(v StatusHistoryDTO)`

SetStatusHistory sets StatusHistory field to given value.

### HasStatusHistory

`func (o *StatusHistoryEntity) HasStatusHistory() bool`

HasStatusHistory returns a boolean if a field has been set.

### GetCanRead

`func (o *StatusHistoryEntity) GetCanRead() bool`

GetCanRead returns the CanRead field if non-nil, zero value otherwise.

### GetCanReadOk

`func (o *StatusHistoryEntity) GetCanReadOk() (*bool, bool)`

GetCanReadOk returns a tuple with the CanRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRead

`func (o *StatusHistoryEntity) SetCanRead(v bool)`

SetCanRead sets CanRead field to given value.

### HasCanRead

`func (o *StatusHistoryEntity) HasCanRead() bool`

HasCanRead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


