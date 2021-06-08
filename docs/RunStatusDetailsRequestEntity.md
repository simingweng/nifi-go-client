# RunStatusDetailsRequestEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessorIds** | Pointer to **[]string** | The IDs of all processors whose run status details should be provided | [optional] 

## Methods

### NewRunStatusDetailsRequestEntity

`func NewRunStatusDetailsRequestEntity() *RunStatusDetailsRequestEntity`

NewRunStatusDetailsRequestEntity instantiates a new RunStatusDetailsRequestEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunStatusDetailsRequestEntityWithDefaults

`func NewRunStatusDetailsRequestEntityWithDefaults() *RunStatusDetailsRequestEntity`

NewRunStatusDetailsRequestEntityWithDefaults instantiates a new RunStatusDetailsRequestEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessorIds

`func (o *RunStatusDetailsRequestEntity) GetProcessorIds() []string`

GetProcessorIds returns the ProcessorIds field if non-nil, zero value otherwise.

### GetProcessorIdsOk

`func (o *RunStatusDetailsRequestEntity) GetProcessorIdsOk() (*[]string, bool)`

GetProcessorIdsOk returns a tuple with the ProcessorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorIds

`func (o *RunStatusDetailsRequestEntity) SetProcessorIds(v []string)`

SetProcessorIds sets ProcessorIds field to given value.

### HasProcessorIds

`func (o *RunStatusDetailsRequestEntity) HasProcessorIds() bool`

HasProcessorIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


