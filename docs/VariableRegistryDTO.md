# VariableRegistryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variables** | Pointer to [**[]VariableEntity**](VariableEntity.md) | The variables that are available in this Variable Registry | [optional] 
**ProcessGroupId** | Pointer to **string** | The UUID of the Process Group that this Variable Registry belongs to | [optional] 

## Methods

### NewVariableRegistryDTO

`func NewVariableRegistryDTO() *VariableRegistryDTO`

NewVariableRegistryDTO instantiates a new VariableRegistryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableRegistryDTOWithDefaults

`func NewVariableRegistryDTOWithDefaults() *VariableRegistryDTO`

NewVariableRegistryDTOWithDefaults instantiates a new VariableRegistryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariables

`func (o *VariableRegistryDTO) GetVariables() []VariableEntity`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *VariableRegistryDTO) GetVariablesOk() (*[]VariableEntity, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *VariableRegistryDTO) SetVariables(v []VariableEntity)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *VariableRegistryDTO) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetProcessGroupId

`func (o *VariableRegistryDTO) GetProcessGroupId() string`

GetProcessGroupId returns the ProcessGroupId field if non-nil, zero value otherwise.

### GetProcessGroupIdOk

`func (o *VariableRegistryDTO) GetProcessGroupIdOk() (*string, bool)`

GetProcessGroupIdOk returns a tuple with the ProcessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupId

`func (o *VariableRegistryDTO) SetProcessGroupId(v string)`

SetProcessGroupId sets ProcessGroupId field to given value.

### HasProcessGroupId

`func (o *VariableRegistryDTO) HasProcessGroupId() bool`

HasProcessGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


