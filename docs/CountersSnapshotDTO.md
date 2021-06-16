# CountersSnapshotDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Generated** | Pointer to **string** | The timestamp when the report was generated. | [optional] 
**Counters** | Pointer to [**[]CounterDTO**](CounterDTO.md) | All counters in the NiFi. | [optional] 

## Methods

### NewCountersSnapshotDTO

`func NewCountersSnapshotDTO() *CountersSnapshotDTO`

NewCountersSnapshotDTO instantiates a new CountersSnapshotDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountersSnapshotDTOWithDefaults

`func NewCountersSnapshotDTOWithDefaults() *CountersSnapshotDTO`

NewCountersSnapshotDTOWithDefaults instantiates a new CountersSnapshotDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenerated

`func (o *CountersSnapshotDTO) GetGenerated() string`

GetGenerated returns the Generated field if non-nil, zero value otherwise.

### GetGeneratedOk

`func (o *CountersSnapshotDTO) GetGeneratedOk() (*string, bool)`

GetGeneratedOk returns a tuple with the Generated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerated

`func (o *CountersSnapshotDTO) SetGenerated(v string)`

SetGenerated sets Generated field to given value.

### HasGenerated

`func (o *CountersSnapshotDTO) HasGenerated() bool`

HasGenerated returns a boolean if a field has been set.

### GetCounters

`func (o *CountersSnapshotDTO) GetCounters() []CounterDTO`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *CountersSnapshotDTO) GetCountersOk() (*[]CounterDTO, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *CountersSnapshotDTO) SetCounters(v []CounterDTO)`

SetCounters sets Counters field to given value.

### HasCounters

`func (o *CountersSnapshotDTO) HasCounters() bool`

HasCounters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


