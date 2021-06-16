# CountersEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counters** | Pointer to [**CountersDTO**](CountersDTO.md) |  | [optional] 

## Methods

### NewCountersEntity

`func NewCountersEntity() *CountersEntity`

NewCountersEntity instantiates a new CountersEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountersEntityWithDefaults

`func NewCountersEntityWithDefaults() *CountersEntity`

NewCountersEntityWithDefaults instantiates a new CountersEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounters

`func (o *CountersEntity) GetCounters() CountersDTO`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *CountersEntity) GetCountersOk() (*CountersDTO, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *CountersEntity) SetCounters(v CountersDTO)`

SetCounters sets Counters field to given value.

### HasCounters

`func (o *CountersEntity) HasCounters() bool`

HasCounters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


