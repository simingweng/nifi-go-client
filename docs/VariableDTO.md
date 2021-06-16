# VariableDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the variable | [optional] 
**Value** | Pointer to **string** | The value of the variable | [optional] 
**ProcessGroupId** | Pointer to **string** | The ID of the Process Group where this Variable is defined | [optional] [readonly] 
**AffectedComponents** | Pointer to [**[]AffectedComponentEntity**](AffectedComponentEntity.md) | A set of all components that will be affected if the value of this variable is changed | [optional] [readonly] 

## Methods

### NewVariableDTO

`func NewVariableDTO() *VariableDTO`

NewVariableDTO instantiates a new VariableDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableDTOWithDefaults

`func NewVariableDTOWithDefaults() *VariableDTO`

NewVariableDTOWithDefaults instantiates a new VariableDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VariableDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VariableDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VariableDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VariableDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *VariableDTO) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VariableDTO) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VariableDTO) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *VariableDTO) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetProcessGroupId

`func (o *VariableDTO) GetProcessGroupId() string`

GetProcessGroupId returns the ProcessGroupId field if non-nil, zero value otherwise.

### GetProcessGroupIdOk

`func (o *VariableDTO) GetProcessGroupIdOk() (*string, bool)`

GetProcessGroupIdOk returns a tuple with the ProcessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupId

`func (o *VariableDTO) SetProcessGroupId(v string)`

SetProcessGroupId sets ProcessGroupId field to given value.

### HasProcessGroupId

`func (o *VariableDTO) HasProcessGroupId() bool`

HasProcessGroupId returns a boolean if a field has been set.

### GetAffectedComponents

`func (o *VariableDTO) GetAffectedComponents() []AffectedComponentEntity`

GetAffectedComponents returns the AffectedComponents field if non-nil, zero value otherwise.

### GetAffectedComponentsOk

`func (o *VariableDTO) GetAffectedComponentsOk() (*[]AffectedComponentEntity, bool)`

GetAffectedComponentsOk returns a tuple with the AffectedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedComponents

`func (o *VariableDTO) SetAffectedComponents(v []AffectedComponentEntity)`

SetAffectedComponents sets AffectedComponents field to given value.

### HasAffectedComponents

`func (o *VariableDTO) HasAffectedComponents() bool`

HasAffectedComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


