# ComponentStateEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentState** | Pointer to [**ComponentStateDTO**](ComponentStateDTO.md) |  | [optional] 

## Methods

### NewComponentStateEntity

`func NewComponentStateEntity() *ComponentStateEntity`

NewComponentStateEntity instantiates a new ComponentStateEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentStateEntityWithDefaults

`func NewComponentStateEntityWithDefaults() *ComponentStateEntity`

NewComponentStateEntityWithDefaults instantiates a new ComponentStateEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentState

`func (o *ComponentStateEntity) GetComponentState() ComponentStateDTO`

GetComponentState returns the ComponentState field if non-nil, zero value otherwise.

### GetComponentStateOk

`func (o *ComponentStateEntity) GetComponentStateOk() (*ComponentStateDTO, bool)`

GetComponentStateOk returns a tuple with the ComponentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentState

`func (o *ComponentStateEntity) SetComponentState(v ComponentStateDTO)`

SetComponentState sets ComponentState field to given value.

### HasComponentState

`func (o *ComponentStateEntity) HasComponentState() bool`

HasComponentState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


