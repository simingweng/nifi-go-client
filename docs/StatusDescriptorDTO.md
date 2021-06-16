# StatusDescriptorDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** | The name of the status field. | [optional] 
**Label** | Pointer to **string** | The label for the status field. | [optional] 
**Description** | Pointer to **string** | The description of the status field. | [optional] 
**Formatter** | Pointer to **string** | The formatter for the status descriptor. | [optional] 

## Methods

### NewStatusDescriptorDTO

`func NewStatusDescriptorDTO() *StatusDescriptorDTO`

NewStatusDescriptorDTO instantiates a new StatusDescriptorDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusDescriptorDTOWithDefaults

`func NewStatusDescriptorDTOWithDefaults() *StatusDescriptorDTO`

NewStatusDescriptorDTOWithDefaults instantiates a new StatusDescriptorDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *StatusDescriptorDTO) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *StatusDescriptorDTO) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *StatusDescriptorDTO) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *StatusDescriptorDTO) HasField() bool`

HasField returns a boolean if a field has been set.

### GetLabel

`func (o *StatusDescriptorDTO) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *StatusDescriptorDTO) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *StatusDescriptorDTO) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *StatusDescriptorDTO) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *StatusDescriptorDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StatusDescriptorDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StatusDescriptorDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StatusDescriptorDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFormatter

`func (o *StatusDescriptorDTO) GetFormatter() string`

GetFormatter returns the Formatter field if non-nil, zero value otherwise.

### GetFormatterOk

`func (o *StatusDescriptorDTO) GetFormatterOk() (*string, bool)`

GetFormatterOk returns a tuple with the Formatter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatter

`func (o *StatusDescriptorDTO) SetFormatter(v string)`

SetFormatter sets Formatter field to given value.

### HasFormatter

`func (o *StatusDescriptorDTO) HasFormatter() bool`

HasFormatter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


