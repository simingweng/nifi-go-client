# AllowableValueDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | A human readable value that is allowed for the property descriptor. | [optional] 
**Value** | Pointer to **string** | A value that is allowed for the property descriptor. | [optional] 
**Description** | Pointer to **string** | A description for this allowable value. | [optional] 

## Methods

### NewAllowableValueDTO

`func NewAllowableValueDTO() *AllowableValueDTO`

NewAllowableValueDTO instantiates a new AllowableValueDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowableValueDTOWithDefaults

`func NewAllowableValueDTOWithDefaults() *AllowableValueDTO`

NewAllowableValueDTOWithDefaults instantiates a new AllowableValueDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *AllowableValueDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AllowableValueDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AllowableValueDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AllowableValueDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetValue

`func (o *AllowableValueDTO) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AllowableValueDTO) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AllowableValueDTO) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AllowableValueDTO) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDescription

`func (o *AllowableValueDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AllowableValueDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AllowableValueDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AllowableValueDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


