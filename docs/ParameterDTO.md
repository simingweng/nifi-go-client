# ParameterDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the Parameter | [optional] 
**Description** | Pointer to **string** | The description of the Parameter | [optional] 
**Sensitive** | Pointer to **bool** | Whether or not the Parameter is sensitive | [optional] 
**Value** | Pointer to **string** | The value of the Parameter | [optional] 
**ValueRemoved** | Pointer to **bool** | Whether or not the value of the Parameter was removed. When a request is made to change a parameter, the value may be null. The absence of the value may be used either to indicate that the value is not to be changed, or that the value is to be set to null (i.e., removed). This denotes which of the two scenarios is being encountered. | [optional] 
**ReferencingComponents** | Pointer to [**[]AffectedComponentEntity**](AffectedComponentEntity.md) | The set of all components in the flow that are referencing this Parameter | [optional] 

## Methods

### NewParameterDTO

`func NewParameterDTO() *ParameterDTO`

NewParameterDTO instantiates a new ParameterDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterDTOWithDefaults

`func NewParameterDTOWithDefaults() *ParameterDTO`

NewParameterDTOWithDefaults instantiates a new ParameterDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ParameterDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParameterDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParameterDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ParameterDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ParameterDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ParameterDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ParameterDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ParameterDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSensitive

`func (o *ParameterDTO) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *ParameterDTO) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *ParameterDTO) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.

### HasSensitive

`func (o *ParameterDTO) HasSensitive() bool`

HasSensitive returns a boolean if a field has been set.

### GetValue

`func (o *ParameterDTO) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ParameterDTO) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ParameterDTO) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ParameterDTO) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueRemoved

`func (o *ParameterDTO) GetValueRemoved() bool`

GetValueRemoved returns the ValueRemoved field if non-nil, zero value otherwise.

### GetValueRemovedOk

`func (o *ParameterDTO) GetValueRemovedOk() (*bool, bool)`

GetValueRemovedOk returns a tuple with the ValueRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueRemoved

`func (o *ParameterDTO) SetValueRemoved(v bool)`

SetValueRemoved sets ValueRemoved field to given value.

### HasValueRemoved

`func (o *ParameterDTO) HasValueRemoved() bool`

HasValueRemoved returns a boolean if a field has been set.

### GetReferencingComponents

`func (o *ParameterDTO) GetReferencingComponents() []AffectedComponentEntity`

GetReferencingComponents returns the ReferencingComponents field if non-nil, zero value otherwise.

### GetReferencingComponentsOk

`func (o *ParameterDTO) GetReferencingComponentsOk() (*[]AffectedComponentEntity, bool)`

GetReferencingComponentsOk returns a tuple with the ReferencingComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencingComponents

`func (o *ParameterDTO) SetReferencingComponents(v []AffectedComponentEntity)`

SetReferencingComponents sets ReferencingComponents field to given value.

### HasReferencingComponents

`func (o *ParameterDTO) HasReferencingComponents() bool`

HasReferencingComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


