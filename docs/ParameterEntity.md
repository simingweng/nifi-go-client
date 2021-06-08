# ParameterEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanWrite** | Pointer to **bool** | Indicates whether the user can write a given resource. | [optional] [readonly] 
**Parameter** | Pointer to [**ParameterDTO**](ParameterDTO.md) |  | [optional] 

## Methods

### NewParameterEntity

`func NewParameterEntity() *ParameterEntity`

NewParameterEntity instantiates a new ParameterEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterEntityWithDefaults

`func NewParameterEntityWithDefaults() *ParameterEntity`

NewParameterEntityWithDefaults instantiates a new ParameterEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanWrite

`func (o *ParameterEntity) GetCanWrite() bool`

GetCanWrite returns the CanWrite field if non-nil, zero value otherwise.

### GetCanWriteOk

`func (o *ParameterEntity) GetCanWriteOk() (*bool, bool)`

GetCanWriteOk returns a tuple with the CanWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWrite

`func (o *ParameterEntity) SetCanWrite(v bool)`

SetCanWrite sets CanWrite field to given value.

### HasCanWrite

`func (o *ParameterEntity) HasCanWrite() bool`

HasCanWrite returns a boolean if a field has been set.

### GetParameter

`func (o *ParameterEntity) GetParameter() ParameterDTO`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ParameterEntity) GetParameterOk() (*ParameterDTO, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *ParameterEntity) SetParameter(v ParameterDTO)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *ParameterEntity) HasParameter() bool`

HasParameter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


