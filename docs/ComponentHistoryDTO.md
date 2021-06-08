# ComponentHistoryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentId** | Pointer to **string** | The component id. | [optional] 
**PropertyHistory** | Pointer to [**map[string]PropertyHistoryDTO**](PropertyHistoryDTO.md) | The history for the properties of the component. | [optional] 

## Methods

### NewComponentHistoryDTO

`func NewComponentHistoryDTO() *ComponentHistoryDTO`

NewComponentHistoryDTO instantiates a new ComponentHistoryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentHistoryDTOWithDefaults

`func NewComponentHistoryDTOWithDefaults() *ComponentHistoryDTO`

NewComponentHistoryDTOWithDefaults instantiates a new ComponentHistoryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentId

`func (o *ComponentHistoryDTO) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *ComponentHistoryDTO) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *ComponentHistoryDTO) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *ComponentHistoryDTO) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### GetPropertyHistory

`func (o *ComponentHistoryDTO) GetPropertyHistory() map[string]PropertyHistoryDTO`

GetPropertyHistory returns the PropertyHistory field if non-nil, zero value otherwise.

### GetPropertyHistoryOk

`func (o *ComponentHistoryDTO) GetPropertyHistoryOk() (*map[string]PropertyHistoryDTO, bool)`

GetPropertyHistoryOk returns a tuple with the PropertyHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyHistory

`func (o *ComponentHistoryDTO) SetPropertyHistory(v map[string]PropertyHistoryDTO)`

SetPropertyHistory sets PropertyHistory field to given value.

### HasPropertyHistory

`func (o *ComponentHistoryDTO) HasPropertyHistory() bool`

HasPropertyHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


