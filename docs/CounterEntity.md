# CounterEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counter** | Pointer to [**CounterDTO**](CounterDTO.md) |  | [optional] 

## Methods

### NewCounterEntity

`func NewCounterEntity() *CounterEntity`

NewCounterEntity instantiates a new CounterEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCounterEntityWithDefaults

`func NewCounterEntityWithDefaults() *CounterEntity`

NewCounterEntityWithDefaults instantiates a new CounterEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounter

`func (o *CounterEntity) GetCounter() CounterDTO`

GetCounter returns the Counter field if non-nil, zero value otherwise.

### GetCounterOk

`func (o *CounterEntity) GetCounterOk() (*CounterDTO, bool)`

GetCounterOk returns a tuple with the Counter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounter

`func (o *CounterEntity) SetCounter(v CounterDTO)`

SetCounter sets Counter field to given value.

### HasCounter

`func (o *CounterEntity) HasCounter() bool`

HasCounter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


