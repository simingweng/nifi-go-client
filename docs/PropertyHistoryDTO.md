# PropertyHistoryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreviousValues** | Pointer to [**[]PreviousValueDTO**](PreviousValueDTO.md) | Previous values for a given property. | [optional] 

## Methods

### NewPropertyHistoryDTO

`func NewPropertyHistoryDTO() *PropertyHistoryDTO`

NewPropertyHistoryDTO instantiates a new PropertyHistoryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyHistoryDTOWithDefaults

`func NewPropertyHistoryDTOWithDefaults() *PropertyHistoryDTO`

NewPropertyHistoryDTOWithDefaults instantiates a new PropertyHistoryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreviousValues

`func (o *PropertyHistoryDTO) GetPreviousValues() []PreviousValueDTO`

GetPreviousValues returns the PreviousValues field if non-nil, zero value otherwise.

### GetPreviousValuesOk

`func (o *PropertyHistoryDTO) GetPreviousValuesOk() (*[]PreviousValueDTO, bool)`

GetPreviousValuesOk returns a tuple with the PreviousValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousValues

`func (o *PropertyHistoryDTO) SetPreviousValues(v []PreviousValueDTO)`

SetPreviousValues sets PreviousValues field to given value.

### HasPreviousValues

`func (o *PropertyHistoryDTO) HasPreviousValues() bool`

HasPreviousValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


