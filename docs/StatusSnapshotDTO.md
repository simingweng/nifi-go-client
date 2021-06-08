# StatusSnapshotDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | The timestamp of the snapshot. | [optional] 
**StatusMetrics** | Pointer to **map[string]int64** | The status metrics. | [optional] 

## Methods

### NewStatusSnapshotDTO

`func NewStatusSnapshotDTO() *StatusSnapshotDTO`

NewStatusSnapshotDTO instantiates a new StatusSnapshotDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusSnapshotDTOWithDefaults

`func NewStatusSnapshotDTOWithDefaults() *StatusSnapshotDTO`

NewStatusSnapshotDTOWithDefaults instantiates a new StatusSnapshotDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *StatusSnapshotDTO) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *StatusSnapshotDTO) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *StatusSnapshotDTO) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *StatusSnapshotDTO) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetStatusMetrics

`func (o *StatusSnapshotDTO) GetStatusMetrics() map[string]int64`

GetStatusMetrics returns the StatusMetrics field if non-nil, zero value otherwise.

### GetStatusMetricsOk

`func (o *StatusSnapshotDTO) GetStatusMetricsOk() (*map[string]int64, bool)`

GetStatusMetricsOk returns a tuple with the StatusMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMetrics

`func (o *StatusSnapshotDTO) SetStatusMetrics(v map[string]int64)`

SetStatusMetrics sets StatusMetrics field to given value.

### HasStatusMetrics

`func (o *StatusSnapshotDTO) HasStatusMetrics() bool`

HasStatusMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


