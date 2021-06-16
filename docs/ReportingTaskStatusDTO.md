# ReportingTaskStatusDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunStatus** | Pointer to **string** | The run status of this ReportingTask | [optional] [readonly] 
**ValidationStatus** | Pointer to **string** | Indicates whether the component is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the component is valid) | [optional] [readonly] 
**ActiveThreadCount** | Pointer to **int32** | The number of active threads for the component. | [optional] 

## Methods

### NewReportingTaskStatusDTO

`func NewReportingTaskStatusDTO() *ReportingTaskStatusDTO`

NewReportingTaskStatusDTO instantiates a new ReportingTaskStatusDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportingTaskStatusDTOWithDefaults

`func NewReportingTaskStatusDTOWithDefaults() *ReportingTaskStatusDTO`

NewReportingTaskStatusDTOWithDefaults instantiates a new ReportingTaskStatusDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunStatus

`func (o *ReportingTaskStatusDTO) GetRunStatus() string`

GetRunStatus returns the RunStatus field if non-nil, zero value otherwise.

### GetRunStatusOk

`func (o *ReportingTaskStatusDTO) GetRunStatusOk() (*string, bool)`

GetRunStatusOk returns a tuple with the RunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStatus

`func (o *ReportingTaskStatusDTO) SetRunStatus(v string)`

SetRunStatus sets RunStatus field to given value.

### HasRunStatus

`func (o *ReportingTaskStatusDTO) HasRunStatus() bool`

HasRunStatus returns a boolean if a field has been set.

### GetValidationStatus

`func (o *ReportingTaskStatusDTO) GetValidationStatus() string`

GetValidationStatus returns the ValidationStatus field if non-nil, zero value otherwise.

### GetValidationStatusOk

`func (o *ReportingTaskStatusDTO) GetValidationStatusOk() (*string, bool)`

GetValidationStatusOk returns a tuple with the ValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStatus

`func (o *ReportingTaskStatusDTO) SetValidationStatus(v string)`

SetValidationStatus sets ValidationStatus field to given value.

### HasValidationStatus

`func (o *ReportingTaskStatusDTO) HasValidationStatus() bool`

HasValidationStatus returns a boolean if a field has been set.

### GetActiveThreadCount

`func (o *ReportingTaskStatusDTO) GetActiveThreadCount() int32`

GetActiveThreadCount returns the ActiveThreadCount field if non-nil, zero value otherwise.

### GetActiveThreadCountOk

`func (o *ReportingTaskStatusDTO) GetActiveThreadCountOk() (*int32, bool)`

GetActiveThreadCountOk returns a tuple with the ActiveThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveThreadCount

`func (o *ReportingTaskStatusDTO) SetActiveThreadCount(v int32)`

SetActiveThreadCount sets ActiveThreadCount field to given value.

### HasActiveThreadCount

`func (o *ReportingTaskStatusDTO) HasActiveThreadCount() bool`

HasActiveThreadCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


