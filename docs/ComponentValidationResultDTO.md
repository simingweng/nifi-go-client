# ComponentValidationResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessGroupId** | Pointer to **string** | The UUID of the Process Group that this component is in | [optional] 
**Id** | Pointer to **string** | The UUID of this component | [optional] 
**ReferenceType** | Pointer to **string** | The type of this component | [optional] 
**Name** | Pointer to **string** | The name of this component. | [optional] 
**State** | Pointer to **string** | The scheduled state of a processor or reporting task referencing a controller service. If this component is another controller service, this field represents the controller service state. | [optional] 
**ActiveThreadCount** | Pointer to **int32** | The number of active threads for the referencing component. | [optional] 
**ValidationErrors** | Pointer to **[]string** | The validation errors for the component. | [optional] 
**CurrentlyValid** | Pointer to **bool** | Whether or not the component is currently valid | [optional] 
**ResultsValid** | Pointer to **bool** | Whether or not the component will be valid if the Parameter Context is changed | [optional] 
**ResultantValidationErrors** | Pointer to **[]string** | The validation errors that will apply to the component if the Parameter Context is changed | [optional] 

## Methods

### NewComponentValidationResultDTO

`func NewComponentValidationResultDTO() *ComponentValidationResultDTO`

NewComponentValidationResultDTO instantiates a new ComponentValidationResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentValidationResultDTOWithDefaults

`func NewComponentValidationResultDTOWithDefaults() *ComponentValidationResultDTO`

NewComponentValidationResultDTOWithDefaults instantiates a new ComponentValidationResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessGroupId

`func (o *ComponentValidationResultDTO) GetProcessGroupId() string`

GetProcessGroupId returns the ProcessGroupId field if non-nil, zero value otherwise.

### GetProcessGroupIdOk

`func (o *ComponentValidationResultDTO) GetProcessGroupIdOk() (*string, bool)`

GetProcessGroupIdOk returns a tuple with the ProcessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupId

`func (o *ComponentValidationResultDTO) SetProcessGroupId(v string)`

SetProcessGroupId sets ProcessGroupId field to given value.

### HasProcessGroupId

`func (o *ComponentValidationResultDTO) HasProcessGroupId() bool`

HasProcessGroupId returns a boolean if a field has been set.

### GetId

`func (o *ComponentValidationResultDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComponentValidationResultDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComponentValidationResultDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComponentValidationResultDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferenceType

`func (o *ComponentValidationResultDTO) GetReferenceType() string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *ComponentValidationResultDTO) GetReferenceTypeOk() (*string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *ComponentValidationResultDTO) SetReferenceType(v string)`

SetReferenceType sets ReferenceType field to given value.

### HasReferenceType

`func (o *ComponentValidationResultDTO) HasReferenceType() bool`

HasReferenceType returns a boolean if a field has been set.

### GetName

`func (o *ComponentValidationResultDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentValidationResultDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentValidationResultDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComponentValidationResultDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *ComponentValidationResultDTO) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ComponentValidationResultDTO) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ComponentValidationResultDTO) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ComponentValidationResultDTO) HasState() bool`

HasState returns a boolean if a field has been set.

### GetActiveThreadCount

`func (o *ComponentValidationResultDTO) GetActiveThreadCount() int32`

GetActiveThreadCount returns the ActiveThreadCount field if non-nil, zero value otherwise.

### GetActiveThreadCountOk

`func (o *ComponentValidationResultDTO) GetActiveThreadCountOk() (*int32, bool)`

GetActiveThreadCountOk returns a tuple with the ActiveThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveThreadCount

`func (o *ComponentValidationResultDTO) SetActiveThreadCount(v int32)`

SetActiveThreadCount sets ActiveThreadCount field to given value.

### HasActiveThreadCount

`func (o *ComponentValidationResultDTO) HasActiveThreadCount() bool`

HasActiveThreadCount returns a boolean if a field has been set.

### GetValidationErrors

`func (o *ComponentValidationResultDTO) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *ComponentValidationResultDTO) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *ComponentValidationResultDTO) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *ComponentValidationResultDTO) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### GetCurrentlyValid

`func (o *ComponentValidationResultDTO) GetCurrentlyValid() bool`

GetCurrentlyValid returns the CurrentlyValid field if non-nil, zero value otherwise.

### GetCurrentlyValidOk

`func (o *ComponentValidationResultDTO) GetCurrentlyValidOk() (*bool, bool)`

GetCurrentlyValidOk returns a tuple with the CurrentlyValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentlyValid

`func (o *ComponentValidationResultDTO) SetCurrentlyValid(v bool)`

SetCurrentlyValid sets CurrentlyValid field to given value.

### HasCurrentlyValid

`func (o *ComponentValidationResultDTO) HasCurrentlyValid() bool`

HasCurrentlyValid returns a boolean if a field has been set.

### GetResultsValid

`func (o *ComponentValidationResultDTO) GetResultsValid() bool`

GetResultsValid returns the ResultsValid field if non-nil, zero value otherwise.

### GetResultsValidOk

`func (o *ComponentValidationResultDTO) GetResultsValidOk() (*bool, bool)`

GetResultsValidOk returns a tuple with the ResultsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsValid

`func (o *ComponentValidationResultDTO) SetResultsValid(v bool)`

SetResultsValid sets ResultsValid field to given value.

### HasResultsValid

`func (o *ComponentValidationResultDTO) HasResultsValid() bool`

HasResultsValid returns a boolean if a field has been set.

### GetResultantValidationErrors

`func (o *ComponentValidationResultDTO) GetResultantValidationErrors() []string`

GetResultantValidationErrors returns the ResultantValidationErrors field if non-nil, zero value otherwise.

### GetResultantValidationErrorsOk

`func (o *ComponentValidationResultDTO) GetResultantValidationErrorsOk() (*[]string, bool)`

GetResultantValidationErrorsOk returns a tuple with the ResultantValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultantValidationErrors

`func (o *ComponentValidationResultDTO) SetResultantValidationErrors(v []string)`

SetResultantValidationErrors sets ResultantValidationErrors field to given value.

### HasResultantValidationErrors

`func (o *ComponentValidationResultDTO) HasResultantValidationErrors() bool`

HasResultantValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


