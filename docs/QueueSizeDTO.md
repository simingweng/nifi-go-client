# QueueSizeDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByteCount** | Pointer to **int64** | The size of objects in a queue. | [optional] 
**ObjectCount** | Pointer to **int32** | The count of objects in a queue. | [optional] 

## Methods

### NewQueueSizeDTO

`func NewQueueSizeDTO() *QueueSizeDTO`

NewQueueSizeDTO instantiates a new QueueSizeDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueSizeDTOWithDefaults

`func NewQueueSizeDTOWithDefaults() *QueueSizeDTO`

NewQueueSizeDTOWithDefaults instantiates a new QueueSizeDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByteCount

`func (o *QueueSizeDTO) GetByteCount() int64`

GetByteCount returns the ByteCount field if non-nil, zero value otherwise.

### GetByteCountOk

`func (o *QueueSizeDTO) GetByteCountOk() (*int64, bool)`

GetByteCountOk returns a tuple with the ByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteCount

`func (o *QueueSizeDTO) SetByteCount(v int64)`

SetByteCount sets ByteCount field to given value.

### HasByteCount

`func (o *QueueSizeDTO) HasByteCount() bool`

HasByteCount returns a boolean if a field has been set.

### GetObjectCount

`func (o *QueueSizeDTO) GetObjectCount() int32`

GetObjectCount returns the ObjectCount field if non-nil, zero value otherwise.

### GetObjectCountOk

`func (o *QueueSizeDTO) GetObjectCountOk() (*int32, bool)`

GetObjectCountOk returns a tuple with the ObjectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectCount

`func (o *QueueSizeDTO) SetObjectCount(v int32)`

SetObjectCount sets ObjectCount field to given value.

### HasObjectCount

`func (o *QueueSizeDTO) HasObjectCount() bool`

HasObjectCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


