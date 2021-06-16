# GarbageCollectionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the garbage collector. | [optional] 
**CollectionCount** | Pointer to **int64** | The number of times garbage collection has run. | [optional] 
**CollectionTime** | Pointer to **string** | The total amount of time spent garbage collecting. | [optional] 
**CollectionMillis** | Pointer to **int64** | The total number of milliseconds spent garbage collecting. | [optional] 

## Methods

### NewGarbageCollectionDTO

`func NewGarbageCollectionDTO() *GarbageCollectionDTO`

NewGarbageCollectionDTO instantiates a new GarbageCollectionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGarbageCollectionDTOWithDefaults

`func NewGarbageCollectionDTOWithDefaults() *GarbageCollectionDTO`

NewGarbageCollectionDTOWithDefaults instantiates a new GarbageCollectionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GarbageCollectionDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GarbageCollectionDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GarbageCollectionDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GarbageCollectionDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCollectionCount

`func (o *GarbageCollectionDTO) GetCollectionCount() int64`

GetCollectionCount returns the CollectionCount field if non-nil, zero value otherwise.

### GetCollectionCountOk

`func (o *GarbageCollectionDTO) GetCollectionCountOk() (*int64, bool)`

GetCollectionCountOk returns a tuple with the CollectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionCount

`func (o *GarbageCollectionDTO) SetCollectionCount(v int64)`

SetCollectionCount sets CollectionCount field to given value.

### HasCollectionCount

`func (o *GarbageCollectionDTO) HasCollectionCount() bool`

HasCollectionCount returns a boolean if a field has been set.

### GetCollectionTime

`func (o *GarbageCollectionDTO) GetCollectionTime() string`

GetCollectionTime returns the CollectionTime field if non-nil, zero value otherwise.

### GetCollectionTimeOk

`func (o *GarbageCollectionDTO) GetCollectionTimeOk() (*string, bool)`

GetCollectionTimeOk returns a tuple with the CollectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTime

`func (o *GarbageCollectionDTO) SetCollectionTime(v string)`

SetCollectionTime sets CollectionTime field to given value.

### HasCollectionTime

`func (o *GarbageCollectionDTO) HasCollectionTime() bool`

HasCollectionTime returns a boolean if a field has been set.

### GetCollectionMillis

`func (o *GarbageCollectionDTO) GetCollectionMillis() int64`

GetCollectionMillis returns the CollectionMillis field if non-nil, zero value otherwise.

### GetCollectionMillisOk

`func (o *GarbageCollectionDTO) GetCollectionMillisOk() (*int64, bool)`

GetCollectionMillisOk returns a tuple with the CollectionMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionMillis

`func (o *GarbageCollectionDTO) SetCollectionMillis(v int64)`

SetCollectionMillis sets CollectionMillis field to given value.

### HasCollectionMillis

`func (o *GarbageCollectionDTO) HasCollectionMillis() bool`

HasCollectionMillis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


