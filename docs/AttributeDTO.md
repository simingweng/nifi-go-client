# AttributeDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The attribute name. | [optional] 
**Value** | Pointer to **string** | The attribute value. | [optional] 
**PreviousValue** | Pointer to **string** | The value of the attribute before the event took place. | [optional] 

## Methods

### NewAttributeDTO

`func NewAttributeDTO() *AttributeDTO`

NewAttributeDTO instantiates a new AttributeDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeDTOWithDefaults

`func NewAttributeDTOWithDefaults() *AttributeDTO`

NewAttributeDTOWithDefaults instantiates a new AttributeDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AttributeDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttributeDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttributeDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AttributeDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *AttributeDTO) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AttributeDTO) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AttributeDTO) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AttributeDTO) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPreviousValue

`func (o *AttributeDTO) GetPreviousValue() string`

GetPreviousValue returns the PreviousValue field if non-nil, zero value otherwise.

### GetPreviousValueOk

`func (o *AttributeDTO) GetPreviousValueOk() (*string, bool)`

GetPreviousValueOk returns a tuple with the PreviousValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousValue

`func (o *AttributeDTO) SetPreviousValue(v string)`

SetPreviousValue sets PreviousValue field to given value.

### HasPreviousValue

`func (o *AttributeDTO) HasPreviousValue() bool`

HasPreviousValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


