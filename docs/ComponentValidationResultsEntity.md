# ComponentValidationResultsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidationResults** | Pointer to [**[]ComponentValidationResultEntity**](ComponentValidationResultEntity.md) | A List of ComponentValidationResultEntity, one for each component that is validated | [optional] 

## Methods

### NewComponentValidationResultsEntity

`func NewComponentValidationResultsEntity() *ComponentValidationResultsEntity`

NewComponentValidationResultsEntity instantiates a new ComponentValidationResultsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentValidationResultsEntityWithDefaults

`func NewComponentValidationResultsEntityWithDefaults() *ComponentValidationResultsEntity`

NewComponentValidationResultsEntityWithDefaults instantiates a new ComponentValidationResultsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidationResults

`func (o *ComponentValidationResultsEntity) GetValidationResults() []ComponentValidationResultEntity`

GetValidationResults returns the ValidationResults field if non-nil, zero value otherwise.

### GetValidationResultsOk

`func (o *ComponentValidationResultsEntity) GetValidationResultsOk() (*[]ComponentValidationResultEntity, bool)`

GetValidationResultsOk returns a tuple with the ValidationResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationResults

`func (o *ComponentValidationResultsEntity) SetValidationResults(v []ComponentValidationResultEntity)`

SetValidationResults sets ValidationResults field to given value.

### HasValidationResults

`func (o *ComponentValidationResultsEntity) HasValidationResults() bool`

HasValidationResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


