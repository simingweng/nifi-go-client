# CounterDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the counter. | [optional] 
**Context** | Pointer to **string** | The context of the counter. | [optional] 
**Name** | Pointer to **string** | The name of the counter. | [optional] 
**ValueCount** | Pointer to **int64** | The value count. | [optional] 
**Value** | Pointer to **string** | The value of the counter. | [optional] 

## Methods

### NewCounterDTO

`func NewCounterDTO() *CounterDTO`

NewCounterDTO instantiates a new CounterDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCounterDTOWithDefaults

`func NewCounterDTOWithDefaults() *CounterDTO`

NewCounterDTOWithDefaults instantiates a new CounterDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CounterDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CounterDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CounterDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CounterDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContext

`func (o *CounterDTO) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CounterDTO) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CounterDTO) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *CounterDTO) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetName

`func (o *CounterDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CounterDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CounterDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CounterDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValueCount

`func (o *CounterDTO) GetValueCount() int64`

GetValueCount returns the ValueCount field if non-nil, zero value otherwise.

### GetValueCountOk

`func (o *CounterDTO) GetValueCountOk() (*int64, bool)`

GetValueCountOk returns a tuple with the ValueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueCount

`func (o *CounterDTO) SetValueCount(v int64)`

SetValueCount sets ValueCount field to given value.

### HasValueCount

`func (o *CounterDTO) HasValueCount() bool`

HasValueCount returns a boolean if a field has been set.

### GetValue

`func (o *CounterDTO) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CounterDTO) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CounterDTO) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CounterDTO) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


