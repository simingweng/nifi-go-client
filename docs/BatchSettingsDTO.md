# BatchSettingsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Preferred number of flow files to include in a transaction. | [optional] 
**Size** | Pointer to **string** | Preferred number of bytes to include in a transaction. | [optional] 
**Duration** | Pointer to **string** | Preferred amount of time that a transaction should span. | [optional] 

## Methods

### NewBatchSettingsDTO

`func NewBatchSettingsDTO() *BatchSettingsDTO`

NewBatchSettingsDTO instantiates a new BatchSettingsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchSettingsDTOWithDefaults

`func NewBatchSettingsDTOWithDefaults() *BatchSettingsDTO`

NewBatchSettingsDTOWithDefaults instantiates a new BatchSettingsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *BatchSettingsDTO) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BatchSettingsDTO) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BatchSettingsDTO) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *BatchSettingsDTO) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSize

`func (o *BatchSettingsDTO) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BatchSettingsDTO) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BatchSettingsDTO) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *BatchSettingsDTO) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetDuration

`func (o *BatchSettingsDTO) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *BatchSettingsDTO) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *BatchSettingsDTO) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *BatchSettingsDTO) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


