# HistoryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | The number of number of actions that matched the search criteria.. | [optional] 
**LastRefreshed** | Pointer to **string** | The timestamp when the report was generated. | [optional] 
**Actions** | Pointer to [**[]ActionEntity**](ActionEntity.md) | The actions. | [optional] 

## Methods

### NewHistoryDTO

`func NewHistoryDTO() *HistoryDTO`

NewHistoryDTO instantiates a new HistoryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryDTOWithDefaults

`func NewHistoryDTOWithDefaults() *HistoryDTO`

NewHistoryDTOWithDefaults instantiates a new HistoryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *HistoryDTO) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *HistoryDTO) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *HistoryDTO) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *HistoryDTO) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetLastRefreshed

`func (o *HistoryDTO) GetLastRefreshed() string`

GetLastRefreshed returns the LastRefreshed field if non-nil, zero value otherwise.

### GetLastRefreshedOk

`func (o *HistoryDTO) GetLastRefreshedOk() (*string, bool)`

GetLastRefreshedOk returns a tuple with the LastRefreshed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefreshed

`func (o *HistoryDTO) SetLastRefreshed(v string)`

SetLastRefreshed sets LastRefreshed field to given value.

### HasLastRefreshed

`func (o *HistoryDTO) HasLastRefreshed() bool`

HasLastRefreshed returns a boolean if a field has been set.

### GetActions

`func (o *HistoryDTO) GetActions() []ActionEntity`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *HistoryDTO) GetActionsOk() (*[]ActionEntity, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *HistoryDTO) SetActions(v []ActionEntity)`

SetActions sets Actions field to given value.

### HasActions

`func (o *HistoryDTO) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


