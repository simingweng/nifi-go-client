# ComponentDifferenceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentType** | Pointer to **string** | The type of component | [optional] 
**ComponentId** | Pointer to **string** | The ID of the component | [optional] 
**ComponentName** | Pointer to **string** | The name of the component | [optional] 
**ProcessGroupId** | Pointer to **string** | The ID of the Process Group that the component belongs to | [optional] 
**Differences** | Pointer to [**[]DifferenceDTO**](DifferenceDTO.md) | The differences in the component between the two flows | [optional] 

## Methods

### NewComponentDifferenceDTO

`func NewComponentDifferenceDTO() *ComponentDifferenceDTO`

NewComponentDifferenceDTO instantiates a new ComponentDifferenceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentDifferenceDTOWithDefaults

`func NewComponentDifferenceDTOWithDefaults() *ComponentDifferenceDTO`

NewComponentDifferenceDTOWithDefaults instantiates a new ComponentDifferenceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentType

`func (o *ComponentDifferenceDTO) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *ComponentDifferenceDTO) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *ComponentDifferenceDTO) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *ComponentDifferenceDTO) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### GetComponentId

`func (o *ComponentDifferenceDTO) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *ComponentDifferenceDTO) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *ComponentDifferenceDTO) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *ComponentDifferenceDTO) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### GetComponentName

`func (o *ComponentDifferenceDTO) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *ComponentDifferenceDTO) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *ComponentDifferenceDTO) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *ComponentDifferenceDTO) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.

### GetProcessGroupId

`func (o *ComponentDifferenceDTO) GetProcessGroupId() string`

GetProcessGroupId returns the ProcessGroupId field if non-nil, zero value otherwise.

### GetProcessGroupIdOk

`func (o *ComponentDifferenceDTO) GetProcessGroupIdOk() (*string, bool)`

GetProcessGroupIdOk returns a tuple with the ProcessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupId

`func (o *ComponentDifferenceDTO) SetProcessGroupId(v string)`

SetProcessGroupId sets ProcessGroupId field to given value.

### HasProcessGroupId

`func (o *ComponentDifferenceDTO) HasProcessGroupId() bool`

HasProcessGroupId returns a boolean if a field has been set.

### GetDifferences

`func (o *ComponentDifferenceDTO) GetDifferences() []DifferenceDTO`

GetDifferences returns the Differences field if non-nil, zero value otherwise.

### GetDifferencesOk

`func (o *ComponentDifferenceDTO) GetDifferencesOk() (*[]DifferenceDTO, bool)`

GetDifferencesOk returns a tuple with the Differences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifferences

`func (o *ComponentDifferenceDTO) SetDifferences(v []DifferenceDTO)`

SetDifferences sets Differences field to given value.

### HasDifferences

`func (o *ComponentDifferenceDTO) HasDifferences() bool`

HasDifferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


