# VersionedParameterContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the context | [optional] 
**Description** | Pointer to **string** | The description of the parameter context | [optional] 
**Parameters** | Pointer to [**[]VersionedParameter**](VersionedParameter.md) | The parameters in the context | [optional] 

## Methods

### NewVersionedParameterContext

`func NewVersionedParameterContext() *VersionedParameterContext`

NewVersionedParameterContext instantiates a new VersionedParameterContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionedParameterContextWithDefaults

`func NewVersionedParameterContextWithDefaults() *VersionedParameterContext`

NewVersionedParameterContextWithDefaults instantiates a new VersionedParameterContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VersionedParameterContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VersionedParameterContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VersionedParameterContext) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VersionedParameterContext) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *VersionedParameterContext) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VersionedParameterContext) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VersionedParameterContext) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VersionedParameterContext) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParameters

`func (o *VersionedParameterContext) GetParameters() []VersionedParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *VersionedParameterContext) GetParametersOk() (*[]VersionedParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *VersionedParameterContext) SetParameters(v []VersionedParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *VersionedParameterContext) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


