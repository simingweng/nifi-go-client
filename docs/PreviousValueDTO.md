# PreviousValueDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreviousValue** | Pointer to **string** | The previous value. | [optional] 
**Timestamp** | Pointer to **string** | The timestamp when the value was modified. | [optional] 
**UserIdentity** | Pointer to **string** | The user who changed the previous value. | [optional] 

## Methods

### NewPreviousValueDTO

`func NewPreviousValueDTO() *PreviousValueDTO`

NewPreviousValueDTO instantiates a new PreviousValueDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviousValueDTOWithDefaults

`func NewPreviousValueDTOWithDefaults() *PreviousValueDTO`

NewPreviousValueDTOWithDefaults instantiates a new PreviousValueDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreviousValue

`func (o *PreviousValueDTO) GetPreviousValue() string`

GetPreviousValue returns the PreviousValue field if non-nil, zero value otherwise.

### GetPreviousValueOk

`func (o *PreviousValueDTO) GetPreviousValueOk() (*string, bool)`

GetPreviousValueOk returns a tuple with the PreviousValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousValue

`func (o *PreviousValueDTO) SetPreviousValue(v string)`

SetPreviousValue sets PreviousValue field to given value.

### HasPreviousValue

`func (o *PreviousValueDTO) HasPreviousValue() bool`

HasPreviousValue returns a boolean if a field has been set.

### GetTimestamp

`func (o *PreviousValueDTO) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PreviousValueDTO) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PreviousValueDTO) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PreviousValueDTO) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUserIdentity

`func (o *PreviousValueDTO) GetUserIdentity() string`

GetUserIdentity returns the UserIdentity field if non-nil, zero value otherwise.

### GetUserIdentityOk

`func (o *PreviousValueDTO) GetUserIdentityOk() (*string, bool)`

GetUserIdentityOk returns a tuple with the UserIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentity

`func (o *PreviousValueDTO) SetUserIdentity(v string)`

SetUserIdentity sets UserIdentity field to given value.

### HasUserIdentity

`func (o *PreviousValueDTO) HasUserIdentity() bool`

HasUserIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


