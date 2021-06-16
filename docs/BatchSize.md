# BatchSize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Preferred number of flow files to include in a transaction. | [optional] 
**Size** | Pointer to **string** | Preferred number of bytes to include in a transaction. | [optional] 
**Duration** | Pointer to **string** | Preferred amount of time that a transaction should span. | [optional] 

## Methods

### NewBatchSize

`func NewBatchSize() *BatchSize`

NewBatchSize instantiates a new BatchSize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchSizeWithDefaults

`func NewBatchSizeWithDefaults() *BatchSize`

NewBatchSizeWithDefaults instantiates a new BatchSize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *BatchSize) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BatchSize) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BatchSize) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *BatchSize) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSize

`func (o *BatchSize) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BatchSize) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BatchSize) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *BatchSize) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetDuration

`func (o *BatchSize) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *BatchSize) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *BatchSize) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *BatchSize) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


