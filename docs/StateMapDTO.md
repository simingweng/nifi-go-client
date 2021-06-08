# StateMapDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | Pointer to **string** | The scope of this StateMap. | [optional] 
**TotalEntryCount** | Pointer to **int32** | The total number of state entries. When the state map is lengthy, only of portion of the entries are returned. | [optional] 
**State** | Pointer to [**[]StateEntryDTO**](StateEntryDTO.md) | The state. | [optional] 

## Methods

### NewStateMapDTO

`func NewStateMapDTO() *StateMapDTO`

NewStateMapDTO instantiates a new StateMapDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateMapDTOWithDefaults

`func NewStateMapDTOWithDefaults() *StateMapDTO`

NewStateMapDTOWithDefaults instantiates a new StateMapDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *StateMapDTO) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *StateMapDTO) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *StateMapDTO) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *StateMapDTO) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTotalEntryCount

`func (o *StateMapDTO) GetTotalEntryCount() int32`

GetTotalEntryCount returns the TotalEntryCount field if non-nil, zero value otherwise.

### GetTotalEntryCountOk

`func (o *StateMapDTO) GetTotalEntryCountOk() (*int32, bool)`

GetTotalEntryCountOk returns a tuple with the TotalEntryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntryCount

`func (o *StateMapDTO) SetTotalEntryCount(v int32)`

SetTotalEntryCount sets TotalEntryCount field to given value.

### HasTotalEntryCount

`func (o *StateMapDTO) HasTotalEntryCount() bool`

HasTotalEntryCount returns a boolean if a field has been set.

### GetState

`func (o *StateMapDTO) GetState() []StateEntryDTO`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StateMapDTO) GetStateOk() (*[]StateEntryDTO, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StateMapDTO) SetState(v []StateEntryDTO)`

SetState sets State field to given value.

### HasState

`func (o *StateMapDTO) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


