# HistoryEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**History** | Pointer to [**HistoryDTO**](HistoryDTO.md) |  | [optional] 

## Methods

### NewHistoryEntity

`func NewHistoryEntity() *HistoryEntity`

NewHistoryEntity instantiates a new HistoryEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryEntityWithDefaults

`func NewHistoryEntityWithDefaults() *HistoryEntity`

NewHistoryEntityWithDefaults instantiates a new HistoryEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistory

`func (o *HistoryEntity) GetHistory() HistoryDTO`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *HistoryEntity) GetHistoryOk() (*HistoryDTO, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *HistoryEntity) SetHistory(v HistoryDTO)`

SetHistory sets History field to given value.

### HasHistory

`func (o *HistoryEntity) HasHistory() bool`

HasHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


