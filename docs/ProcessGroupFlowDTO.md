# ProcessGroupFlowDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the component. | [optional] 
**Uri** | Pointer to **string** | The URI for futures requests to the component. | [optional] 
**ParentGroupId** | Pointer to **string** | The id of parent process group of this component if applicable. | [optional] 
**ParameterContext** | Pointer to [**ParameterContextReferenceEntity**](ParameterContextReferenceEntity.md) |  | [optional] 
**Breadcrumb** | Pointer to [**FlowBreadcrumbEntity**](FlowBreadcrumbEntity.md) |  | [optional] 
**Flow** | Pointer to [**FlowDTO**](FlowDTO.md) |  | [optional] 
**LastRefreshed** | Pointer to **string** | The time the flow for the process group was last refreshed. | [optional] 

## Methods

### NewProcessGroupFlowDTO

`func NewProcessGroupFlowDTO() *ProcessGroupFlowDTO`

NewProcessGroupFlowDTO instantiates a new ProcessGroupFlowDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessGroupFlowDTOWithDefaults

`func NewProcessGroupFlowDTOWithDefaults() *ProcessGroupFlowDTO`

NewProcessGroupFlowDTOWithDefaults instantiates a new ProcessGroupFlowDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProcessGroupFlowDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessGroupFlowDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessGroupFlowDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProcessGroupFlowDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUri

`func (o *ProcessGroupFlowDTO) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ProcessGroupFlowDTO) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ProcessGroupFlowDTO) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ProcessGroupFlowDTO) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetParentGroupId

`func (o *ProcessGroupFlowDTO) GetParentGroupId() string`

GetParentGroupId returns the ParentGroupId field if non-nil, zero value otherwise.

### GetParentGroupIdOk

`func (o *ProcessGroupFlowDTO) GetParentGroupIdOk() (*string, bool)`

GetParentGroupIdOk returns a tuple with the ParentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroupId

`func (o *ProcessGroupFlowDTO) SetParentGroupId(v string)`

SetParentGroupId sets ParentGroupId field to given value.

### HasParentGroupId

`func (o *ProcessGroupFlowDTO) HasParentGroupId() bool`

HasParentGroupId returns a boolean if a field has been set.

### GetParameterContext

`func (o *ProcessGroupFlowDTO) GetParameterContext() ParameterContextReferenceEntity`

GetParameterContext returns the ParameterContext field if non-nil, zero value otherwise.

### GetParameterContextOk

`func (o *ProcessGroupFlowDTO) GetParameterContextOk() (*ParameterContextReferenceEntity, bool)`

GetParameterContextOk returns a tuple with the ParameterContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterContext

`func (o *ProcessGroupFlowDTO) SetParameterContext(v ParameterContextReferenceEntity)`

SetParameterContext sets ParameterContext field to given value.

### HasParameterContext

`func (o *ProcessGroupFlowDTO) HasParameterContext() bool`

HasParameterContext returns a boolean if a field has been set.

### GetBreadcrumb

`func (o *ProcessGroupFlowDTO) GetBreadcrumb() FlowBreadcrumbEntity`

GetBreadcrumb returns the Breadcrumb field if non-nil, zero value otherwise.

### GetBreadcrumbOk

`func (o *ProcessGroupFlowDTO) GetBreadcrumbOk() (*FlowBreadcrumbEntity, bool)`

GetBreadcrumbOk returns a tuple with the Breadcrumb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreadcrumb

`func (o *ProcessGroupFlowDTO) SetBreadcrumb(v FlowBreadcrumbEntity)`

SetBreadcrumb sets Breadcrumb field to given value.

### HasBreadcrumb

`func (o *ProcessGroupFlowDTO) HasBreadcrumb() bool`

HasBreadcrumb returns a boolean if a field has been set.

### GetFlow

`func (o *ProcessGroupFlowDTO) GetFlow() FlowDTO`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *ProcessGroupFlowDTO) GetFlowOk() (*FlowDTO, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *ProcessGroupFlowDTO) SetFlow(v FlowDTO)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *ProcessGroupFlowDTO) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetLastRefreshed

`func (o *ProcessGroupFlowDTO) GetLastRefreshed() string`

GetLastRefreshed returns the LastRefreshed field if non-nil, zero value otherwise.

### GetLastRefreshedOk

`func (o *ProcessGroupFlowDTO) GetLastRefreshedOk() (*string, bool)`

GetLastRefreshedOk returns a tuple with the LastRefreshed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefreshed

`func (o *ProcessGroupFlowDTO) SetLastRefreshed(v string)`

SetLastRefreshed sets LastRefreshed field to given value.

### HasLastRefreshed

`func (o *ProcessGroupFlowDTO) HasLastRefreshed() bool`

HasLastRefreshed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


