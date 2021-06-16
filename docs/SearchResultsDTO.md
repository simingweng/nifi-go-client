# SearchResultsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessorResults** | Pointer to [**[]ComponentSearchResultDTO**](ComponentSearchResultDTO.md) | The processors that matched the search. | [optional] 
**ConnectionResults** | Pointer to [**[]ComponentSearchResultDTO**](ComponentSearchResultDTO.md) | The connections that matched the search. | [optional] 
**ProcessGroupResults** | Pointer to [**[]ComponentSearchResultDTO**](ComponentSearchResultDTO.md) | The process groups that matched the search. | [optional] 
**InputPortResults** | Pointer to [**[]ComponentSearchResultDTO**](ComponentSearchResultDTO.md) | The input ports that matched the search. | [optional] 
**OutputPortResults** | Pointer to [**[]ComponentSearchResultDTO**](ComponentSearchResultDTO.md) | The output ports that matched the search. | [optional] 
**RemoteProcessGroupResults** | Pointer to [**[]ComponentSearchResultDTO**](ComponentSearchResultDTO.md) | The remote process groups that matched the search. | [optional] 
**FunnelResults** | Pointer to [**[]ComponentSearchResultDTO**](ComponentSearchResultDTO.md) | The funnels that matched the search. | [optional] 
**LabelResults** | Pointer to [**[]ComponentSearchResultDTO**](ComponentSearchResultDTO.md) | The labels that matched the search. | [optional] 
**ControllerServiceNodeResults** | Pointer to [**[]ComponentSearchResultDTO**](ComponentSearchResultDTO.md) | The controller service nodes that matched the search | [optional] 
**ParameterContextResults** | Pointer to [**[]ComponentSearchResultDTO**](ComponentSearchResultDTO.md) | The parameter contexts that matched the search. | [optional] 
**ParameterResults** | Pointer to [**[]ComponentSearchResultDTO**](ComponentSearchResultDTO.md) | The parameters that matched the search. | [optional] 

## Methods

### NewSearchResultsDTO

`func NewSearchResultsDTO() *SearchResultsDTO`

NewSearchResultsDTO instantiates a new SearchResultsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultsDTOWithDefaults

`func NewSearchResultsDTOWithDefaults() *SearchResultsDTO`

NewSearchResultsDTOWithDefaults instantiates a new SearchResultsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessorResults

`func (o *SearchResultsDTO) GetProcessorResults() []ComponentSearchResultDTO`

GetProcessorResults returns the ProcessorResults field if non-nil, zero value otherwise.

### GetProcessorResultsOk

`func (o *SearchResultsDTO) GetProcessorResultsOk() (*[]ComponentSearchResultDTO, bool)`

GetProcessorResultsOk returns a tuple with the ProcessorResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorResults

`func (o *SearchResultsDTO) SetProcessorResults(v []ComponentSearchResultDTO)`

SetProcessorResults sets ProcessorResults field to given value.

### HasProcessorResults

`func (o *SearchResultsDTO) HasProcessorResults() bool`

HasProcessorResults returns a boolean if a field has been set.

### GetConnectionResults

`func (o *SearchResultsDTO) GetConnectionResults() []ComponentSearchResultDTO`

GetConnectionResults returns the ConnectionResults field if non-nil, zero value otherwise.

### GetConnectionResultsOk

`func (o *SearchResultsDTO) GetConnectionResultsOk() (*[]ComponentSearchResultDTO, bool)`

GetConnectionResultsOk returns a tuple with the ConnectionResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionResults

`func (o *SearchResultsDTO) SetConnectionResults(v []ComponentSearchResultDTO)`

SetConnectionResults sets ConnectionResults field to given value.

### HasConnectionResults

`func (o *SearchResultsDTO) HasConnectionResults() bool`

HasConnectionResults returns a boolean if a field has been set.

### GetProcessGroupResults

`func (o *SearchResultsDTO) GetProcessGroupResults() []ComponentSearchResultDTO`

GetProcessGroupResults returns the ProcessGroupResults field if non-nil, zero value otherwise.

### GetProcessGroupResultsOk

`func (o *SearchResultsDTO) GetProcessGroupResultsOk() (*[]ComponentSearchResultDTO, bool)`

GetProcessGroupResultsOk returns a tuple with the ProcessGroupResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupResults

`func (o *SearchResultsDTO) SetProcessGroupResults(v []ComponentSearchResultDTO)`

SetProcessGroupResults sets ProcessGroupResults field to given value.

### HasProcessGroupResults

`func (o *SearchResultsDTO) HasProcessGroupResults() bool`

HasProcessGroupResults returns a boolean if a field has been set.

### GetInputPortResults

`func (o *SearchResultsDTO) GetInputPortResults() []ComponentSearchResultDTO`

GetInputPortResults returns the InputPortResults field if non-nil, zero value otherwise.

### GetInputPortResultsOk

`func (o *SearchResultsDTO) GetInputPortResultsOk() (*[]ComponentSearchResultDTO, bool)`

GetInputPortResultsOk returns a tuple with the InputPortResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPortResults

`func (o *SearchResultsDTO) SetInputPortResults(v []ComponentSearchResultDTO)`

SetInputPortResults sets InputPortResults field to given value.

### HasInputPortResults

`func (o *SearchResultsDTO) HasInputPortResults() bool`

HasInputPortResults returns a boolean if a field has been set.

### GetOutputPortResults

`func (o *SearchResultsDTO) GetOutputPortResults() []ComponentSearchResultDTO`

GetOutputPortResults returns the OutputPortResults field if non-nil, zero value otherwise.

### GetOutputPortResultsOk

`func (o *SearchResultsDTO) GetOutputPortResultsOk() (*[]ComponentSearchResultDTO, bool)`

GetOutputPortResultsOk returns a tuple with the OutputPortResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPortResults

`func (o *SearchResultsDTO) SetOutputPortResults(v []ComponentSearchResultDTO)`

SetOutputPortResults sets OutputPortResults field to given value.

### HasOutputPortResults

`func (o *SearchResultsDTO) HasOutputPortResults() bool`

HasOutputPortResults returns a boolean if a field has been set.

### GetRemoteProcessGroupResults

`func (o *SearchResultsDTO) GetRemoteProcessGroupResults() []ComponentSearchResultDTO`

GetRemoteProcessGroupResults returns the RemoteProcessGroupResults field if non-nil, zero value otherwise.

### GetRemoteProcessGroupResultsOk

`func (o *SearchResultsDTO) GetRemoteProcessGroupResultsOk() (*[]ComponentSearchResultDTO, bool)`

GetRemoteProcessGroupResultsOk returns a tuple with the RemoteProcessGroupResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProcessGroupResults

`func (o *SearchResultsDTO) SetRemoteProcessGroupResults(v []ComponentSearchResultDTO)`

SetRemoteProcessGroupResults sets RemoteProcessGroupResults field to given value.

### HasRemoteProcessGroupResults

`func (o *SearchResultsDTO) HasRemoteProcessGroupResults() bool`

HasRemoteProcessGroupResults returns a boolean if a field has been set.

### GetFunnelResults

`func (o *SearchResultsDTO) GetFunnelResults() []ComponentSearchResultDTO`

GetFunnelResults returns the FunnelResults field if non-nil, zero value otherwise.

### GetFunnelResultsOk

`func (o *SearchResultsDTO) GetFunnelResultsOk() (*[]ComponentSearchResultDTO, bool)`

GetFunnelResultsOk returns a tuple with the FunnelResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelResults

`func (o *SearchResultsDTO) SetFunnelResults(v []ComponentSearchResultDTO)`

SetFunnelResults sets FunnelResults field to given value.

### HasFunnelResults

`func (o *SearchResultsDTO) HasFunnelResults() bool`

HasFunnelResults returns a boolean if a field has been set.

### GetLabelResults

`func (o *SearchResultsDTO) GetLabelResults() []ComponentSearchResultDTO`

GetLabelResults returns the LabelResults field if non-nil, zero value otherwise.

### GetLabelResultsOk

`func (o *SearchResultsDTO) GetLabelResultsOk() (*[]ComponentSearchResultDTO, bool)`

GetLabelResultsOk returns a tuple with the LabelResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelResults

`func (o *SearchResultsDTO) SetLabelResults(v []ComponentSearchResultDTO)`

SetLabelResults sets LabelResults field to given value.

### HasLabelResults

`func (o *SearchResultsDTO) HasLabelResults() bool`

HasLabelResults returns a boolean if a field has been set.

### GetControllerServiceNodeResults

`func (o *SearchResultsDTO) GetControllerServiceNodeResults() []ComponentSearchResultDTO`

GetControllerServiceNodeResults returns the ControllerServiceNodeResults field if non-nil, zero value otherwise.

### GetControllerServiceNodeResultsOk

`func (o *SearchResultsDTO) GetControllerServiceNodeResultsOk() (*[]ComponentSearchResultDTO, bool)`

GetControllerServiceNodeResultsOk returns a tuple with the ControllerServiceNodeResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerServiceNodeResults

`func (o *SearchResultsDTO) SetControllerServiceNodeResults(v []ComponentSearchResultDTO)`

SetControllerServiceNodeResults sets ControllerServiceNodeResults field to given value.

### HasControllerServiceNodeResults

`func (o *SearchResultsDTO) HasControllerServiceNodeResults() bool`

HasControllerServiceNodeResults returns a boolean if a field has been set.

### GetParameterContextResults

`func (o *SearchResultsDTO) GetParameterContextResults() []ComponentSearchResultDTO`

GetParameterContextResults returns the ParameterContextResults field if non-nil, zero value otherwise.

### GetParameterContextResultsOk

`func (o *SearchResultsDTO) GetParameterContextResultsOk() (*[]ComponentSearchResultDTO, bool)`

GetParameterContextResultsOk returns a tuple with the ParameterContextResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterContextResults

`func (o *SearchResultsDTO) SetParameterContextResults(v []ComponentSearchResultDTO)`

SetParameterContextResults sets ParameterContextResults field to given value.

### HasParameterContextResults

`func (o *SearchResultsDTO) HasParameterContextResults() bool`

HasParameterContextResults returns a boolean if a field has been set.

### GetParameterResults

`func (o *SearchResultsDTO) GetParameterResults() []ComponentSearchResultDTO`

GetParameterResults returns the ParameterResults field if non-nil, zero value otherwise.

### GetParameterResultsOk

`func (o *SearchResultsDTO) GetParameterResultsOk() (*[]ComponentSearchResultDTO, bool)`

GetParameterResultsOk returns a tuple with the ParameterResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterResults

`func (o *SearchResultsDTO) SetParameterResults(v []ComponentSearchResultDTO)`

SetParameterResults sets ParameterResults field to given value.

### HasParameterResults

`func (o *SearchResultsDTO) HasParameterResults() bool`

HasParameterResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


