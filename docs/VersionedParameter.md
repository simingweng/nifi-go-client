# VersionedParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the parameter | [optional] 
**Description** | Pointer to **string** | The description of the param | [optional] 
**Sensitive** | Pointer to **bool** | Whether or not the parameter value is sensitive | [optional] 
**Value** | Pointer to **string** | The value of the parameter | [optional] 

## Methods

### NewVersionedParameter

`func NewVersionedParameter() *VersionedParameter`

NewVersionedParameter instantiates a new VersionedParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionedParameterWithDefaults

`func NewVersionedParameterWithDefaults() *VersionedParameter`

NewVersionedParameterWithDefaults instantiates a new VersionedParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VersionedParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VersionedParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VersionedParameter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VersionedParameter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *VersionedParameter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VersionedParameter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VersionedParameter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VersionedParameter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSensitive

`func (o *VersionedParameter) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *VersionedParameter) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *VersionedParameter) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.

### HasSensitive

`func (o *VersionedParameter) HasSensitive() bool`

HasSensitive returns a boolean if a field has been set.

### GetValue

`func (o *VersionedParameter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VersionedParameter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VersionedParameter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *VersionedParameter) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


