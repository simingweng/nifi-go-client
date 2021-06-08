# ProcessorRunStatusDetailsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the processor | [optional] 
**Name** | Pointer to **string** | The name of the processor | [optional] 
**RunStatus** | Pointer to **string** | The run status of the processor | [optional] 
**ValidationErrors** | Pointer to **[]string** | The processor&#39;s validation errors | [optional] 
**ActiveThreadCount** | Pointer to **int32** | The current number of threads that the processor is currently using | [optional] 

## Methods

### NewProcessorRunStatusDetailsDTO

`func NewProcessorRunStatusDetailsDTO() *ProcessorRunStatusDetailsDTO`

NewProcessorRunStatusDetailsDTO instantiates a new ProcessorRunStatusDetailsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorRunStatusDetailsDTOWithDefaults

`func NewProcessorRunStatusDetailsDTOWithDefaults() *ProcessorRunStatusDetailsDTO`

NewProcessorRunStatusDetailsDTOWithDefaults instantiates a new ProcessorRunStatusDetailsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProcessorRunStatusDetailsDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessorRunStatusDetailsDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessorRunStatusDetailsDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProcessorRunStatusDetailsDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProcessorRunStatusDetailsDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProcessorRunStatusDetailsDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProcessorRunStatusDetailsDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProcessorRunStatusDetailsDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRunStatus

`func (o *ProcessorRunStatusDetailsDTO) GetRunStatus() string`

GetRunStatus returns the RunStatus field if non-nil, zero value otherwise.

### GetRunStatusOk

`func (o *ProcessorRunStatusDetailsDTO) GetRunStatusOk() (*string, bool)`

GetRunStatusOk returns a tuple with the RunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStatus

`func (o *ProcessorRunStatusDetailsDTO) SetRunStatus(v string)`

SetRunStatus sets RunStatus field to given value.

### HasRunStatus

`func (o *ProcessorRunStatusDetailsDTO) HasRunStatus() bool`

HasRunStatus returns a boolean if a field has been set.

### GetValidationErrors

`func (o *ProcessorRunStatusDetailsDTO) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *ProcessorRunStatusDetailsDTO) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *ProcessorRunStatusDetailsDTO) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *ProcessorRunStatusDetailsDTO) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### GetActiveThreadCount

`func (o *ProcessorRunStatusDetailsDTO) GetActiveThreadCount() int32`

GetActiveThreadCount returns the ActiveThreadCount field if non-nil, zero value otherwise.

### GetActiveThreadCountOk

`func (o *ProcessorRunStatusDetailsDTO) GetActiveThreadCountOk() (*int32, bool)`

GetActiveThreadCountOk returns a tuple with the ActiveThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveThreadCount

`func (o *ProcessorRunStatusDetailsDTO) SetActiveThreadCount(v int32)`

SetActiveThreadCount sets ActiveThreadCount field to given value.

### HasActiveThreadCount

`func (o *ProcessorRunStatusDetailsDTO) HasActiveThreadCount() bool`

HasActiveThreadCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


