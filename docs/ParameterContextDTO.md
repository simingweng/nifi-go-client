# ParameterContextDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The Name of the Parameter Context. | [optional] 
**Description** | Pointer to **string** | The Description of the Parameter Context. | [optional] 
**Parameters** | Pointer to [**[]ParameterEntity**](ParameterEntity.md) | The Parameters for the Parameter Context | [optional] 
**BoundProcessGroups** | Pointer to [**[]ProcessGroupEntity**](ProcessGroupEntity.md) | The Process Groups that are bound to this Parameter Context | [optional] [readonly] 
**Id** | Pointer to **string** | The ID the Parameter Context. | [optional] [readonly] 

## Methods

### NewParameterContextDTO

`func NewParameterContextDTO() *ParameterContextDTO`

NewParameterContextDTO instantiates a new ParameterContextDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterContextDTOWithDefaults

`func NewParameterContextDTOWithDefaults() *ParameterContextDTO`

NewParameterContextDTOWithDefaults instantiates a new ParameterContextDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ParameterContextDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParameterContextDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParameterContextDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ParameterContextDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ParameterContextDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ParameterContextDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ParameterContextDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ParameterContextDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParameters

`func (o *ParameterContextDTO) GetParameters() []ParameterEntity`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ParameterContextDTO) GetParametersOk() (*[]ParameterEntity, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ParameterContextDTO) SetParameters(v []ParameterEntity)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ParameterContextDTO) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetBoundProcessGroups

`func (o *ParameterContextDTO) GetBoundProcessGroups() []ProcessGroupEntity`

GetBoundProcessGroups returns the BoundProcessGroups field if non-nil, zero value otherwise.

### GetBoundProcessGroupsOk

`func (o *ParameterContextDTO) GetBoundProcessGroupsOk() (*[]ProcessGroupEntity, bool)`

GetBoundProcessGroupsOk returns a tuple with the BoundProcessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundProcessGroups

`func (o *ParameterContextDTO) SetBoundProcessGroups(v []ProcessGroupEntity)`

SetBoundProcessGroups sets BoundProcessGroups field to given value.

### HasBoundProcessGroups

`func (o *ParameterContextDTO) HasBoundProcessGroups() bool`

HasBoundProcessGroups returns a boolean if a field has been set.

### GetId

`func (o *ParameterContextDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParameterContextDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParameterContextDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ParameterContextDTO) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


