# FlowComparisonEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentDifferences** | Pointer to [**[]ComponentDifferenceDTO**](ComponentDifferenceDTO.md) | The list of differences for each component in the flow that is not the same between the two flows | [optional] 

## Methods

### NewFlowComparisonEntity

`func NewFlowComparisonEntity() *FlowComparisonEntity`

NewFlowComparisonEntity instantiates a new FlowComparisonEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowComparisonEntityWithDefaults

`func NewFlowComparisonEntityWithDefaults() *FlowComparisonEntity`

NewFlowComparisonEntityWithDefaults instantiates a new FlowComparisonEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentDifferences

`func (o *FlowComparisonEntity) GetComponentDifferences() []ComponentDifferenceDTO`

GetComponentDifferences returns the ComponentDifferences field if non-nil, zero value otherwise.

### GetComponentDifferencesOk

`func (o *FlowComparisonEntity) GetComponentDifferencesOk() (*[]ComponentDifferenceDTO, bool)`

GetComponentDifferencesOk returns a tuple with the ComponentDifferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentDifferences

`func (o *FlowComparisonEntity) SetComponentDifferences(v []ComponentDifferenceDTO)`

SetComponentDifferences sets ComponentDifferences field to given value.

### HasComponentDifferences

`func (o *FlowComparisonEntity) HasComponentDifferences() bool`

HasComponentDifferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


