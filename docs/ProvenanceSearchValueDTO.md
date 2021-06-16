# ProvenanceSearchValueDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | The search value. | [optional] 
**Inverse** | Pointer to **bool** | Query for all except for search value. | [optional] 

## Methods

### NewProvenanceSearchValueDTO

`func NewProvenanceSearchValueDTO() *ProvenanceSearchValueDTO`

NewProvenanceSearchValueDTO instantiates a new ProvenanceSearchValueDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvenanceSearchValueDTOWithDefaults

`func NewProvenanceSearchValueDTOWithDefaults() *ProvenanceSearchValueDTO`

NewProvenanceSearchValueDTOWithDefaults instantiates a new ProvenanceSearchValueDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ProvenanceSearchValueDTO) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProvenanceSearchValueDTO) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProvenanceSearchValueDTO) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProvenanceSearchValueDTO) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetInverse

`func (o *ProvenanceSearchValueDTO) GetInverse() bool`

GetInverse returns the Inverse field if non-nil, zero value otherwise.

### GetInverseOk

`func (o *ProvenanceSearchValueDTO) GetInverseOk() (*bool, bool)`

GetInverseOk returns a tuple with the Inverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverse

`func (o *ProvenanceSearchValueDTO) SetInverse(v bool)`

SetInverse sets Inverse field to given value.

### HasInverse

`func (o *ProvenanceSearchValueDTO) HasInverse() bool`

HasInverse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


