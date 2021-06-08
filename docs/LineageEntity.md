# LineageEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lineage** | Pointer to [**LineageDTO**](LineageDTO.md) |  | [optional] 

## Methods

### NewLineageEntity

`func NewLineageEntity() *LineageEntity`

NewLineageEntity instantiates a new LineageEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineageEntityWithDefaults

`func NewLineageEntityWithDefaults() *LineageEntity`

NewLineageEntityWithDefaults instantiates a new LineageEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineage

`func (o *LineageEntity) GetLineage() LineageDTO`

GetLineage returns the Lineage field if non-nil, zero value otherwise.

### GetLineageOk

`func (o *LineageEntity) GetLineageOk() (*LineageDTO, bool)`

GetLineageOk returns a tuple with the Lineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineage

`func (o *LineageEntity) SetLineage(v LineageDTO)`

SetLineage sets Lineage field to given value.

### HasLineage

`func (o *LineageEntity) HasLineage() bool`

HasLineage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


