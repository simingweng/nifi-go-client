# VariableEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variable** | Pointer to [**VariableDTO**](VariableDTO.md) |  | [optional] 
**CanWrite** | Pointer to **bool** | Indicates whether the user can write a given resource. | [optional] [readonly] 

## Methods

### NewVariableEntity

`func NewVariableEntity() *VariableEntity`

NewVariableEntity instantiates a new VariableEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableEntityWithDefaults

`func NewVariableEntityWithDefaults() *VariableEntity`

NewVariableEntityWithDefaults instantiates a new VariableEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariable

`func (o *VariableEntity) GetVariable() VariableDTO`

GetVariable returns the Variable field if non-nil, zero value otherwise.

### GetVariableOk

`func (o *VariableEntity) GetVariableOk() (*VariableDTO, bool)`

GetVariableOk returns a tuple with the Variable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariable

`func (o *VariableEntity) SetVariable(v VariableDTO)`

SetVariable sets Variable field to given value.

### HasVariable

`func (o *VariableEntity) HasVariable() bool`

HasVariable returns a boolean if a field has been set.

### GetCanWrite

`func (o *VariableEntity) GetCanWrite() bool`

GetCanWrite returns the CanWrite field if non-nil, zero value otherwise.

### GetCanWriteOk

`func (o *VariableEntity) GetCanWriteOk() (*bool, bool)`

GetCanWriteOk returns a tuple with the CanWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWrite

`func (o *VariableEntity) SetCanWrite(v bool)`

SetCanWrite sets CanWrite field to given value.

### HasCanWrite

`func (o *VariableEntity) HasCanWrite() bool`

HasCanWrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


