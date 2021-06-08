# DifferenceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DifferenceType** | Pointer to **string** | The type of difference | [optional] 
**Difference** | Pointer to **string** | Description of the difference | [optional] 

## Methods

### NewDifferenceDTO

`func NewDifferenceDTO() *DifferenceDTO`

NewDifferenceDTO instantiates a new DifferenceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDifferenceDTOWithDefaults

`func NewDifferenceDTOWithDefaults() *DifferenceDTO`

NewDifferenceDTOWithDefaults instantiates a new DifferenceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDifferenceType

`func (o *DifferenceDTO) GetDifferenceType() string`

GetDifferenceType returns the DifferenceType field if non-nil, zero value otherwise.

### GetDifferenceTypeOk

`func (o *DifferenceDTO) GetDifferenceTypeOk() (*string, bool)`

GetDifferenceTypeOk returns a tuple with the DifferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifferenceType

`func (o *DifferenceDTO) SetDifferenceType(v string)`

SetDifferenceType sets DifferenceType field to given value.

### HasDifferenceType

`func (o *DifferenceDTO) HasDifferenceType() bool`

HasDifferenceType returns a boolean if a field has been set.

### GetDifference

`func (o *DifferenceDTO) GetDifference() string`

GetDifference returns the Difference field if non-nil, zero value otherwise.

### GetDifferenceOk

`func (o *DifferenceDTO) GetDifferenceOk() (*string, bool)`

GetDifferenceOk returns a tuple with the Difference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifference

`func (o *DifferenceDTO) SetDifference(v string)`

SetDifference sets Difference field to given value.

### HasDifference

`func (o *DifferenceDTO) HasDifference() bool`

HasDifference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


