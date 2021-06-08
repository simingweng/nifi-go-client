# ConnectionStatisticsSnapshotDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the connection. | [optional] 
**PredictedMillisUntilCountBackpressure** | Pointer to **int64** | The predicted number of milliseconds before the connection will have backpressure applied, based on the queued count. | [optional] 
**PredictedMillisUntilBytesBackpressure** | Pointer to **int64** | The predicted number of milliseconds before the connection will have backpressure applied, based on the total number of bytes in the queue. | [optional] 
**PredictedCountAtNextInterval** | Pointer to **int32** | The predicted number of queued objects at the next configured interval. | [optional] 
**PredictedBytesAtNextInterval** | Pointer to **int64** | The predicted total number of bytes in the queue at the next configured interval. | [optional] 
**PredictedPercentCount** | Pointer to **int32** | The predicted percentage of queued objects at the next configured interval. | [optional] 
**PredictedPercentBytes** | Pointer to **int32** | The predicted percentage of bytes in the queue against current threshold at the next configured interval. | [optional] 
**PredictionIntervalMillis** | Pointer to **int64** | The prediction interval in seconds | [optional] 

## Methods

### NewConnectionStatisticsSnapshotDTO

`func NewConnectionStatisticsSnapshotDTO() *ConnectionStatisticsSnapshotDTO`

NewConnectionStatisticsSnapshotDTO instantiates a new ConnectionStatisticsSnapshotDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionStatisticsSnapshotDTOWithDefaults

`func NewConnectionStatisticsSnapshotDTOWithDefaults() *ConnectionStatisticsSnapshotDTO`

NewConnectionStatisticsSnapshotDTOWithDefaults instantiates a new ConnectionStatisticsSnapshotDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectionStatisticsSnapshotDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectionStatisticsSnapshotDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectionStatisticsSnapshotDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectionStatisticsSnapshotDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPredictedMillisUntilCountBackpressure

`func (o *ConnectionStatisticsSnapshotDTO) GetPredictedMillisUntilCountBackpressure() int64`

GetPredictedMillisUntilCountBackpressure returns the PredictedMillisUntilCountBackpressure field if non-nil, zero value otherwise.

### GetPredictedMillisUntilCountBackpressureOk

`func (o *ConnectionStatisticsSnapshotDTO) GetPredictedMillisUntilCountBackpressureOk() (*int64, bool)`

GetPredictedMillisUntilCountBackpressureOk returns a tuple with the PredictedMillisUntilCountBackpressure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedMillisUntilCountBackpressure

`func (o *ConnectionStatisticsSnapshotDTO) SetPredictedMillisUntilCountBackpressure(v int64)`

SetPredictedMillisUntilCountBackpressure sets PredictedMillisUntilCountBackpressure field to given value.

### HasPredictedMillisUntilCountBackpressure

`func (o *ConnectionStatisticsSnapshotDTO) HasPredictedMillisUntilCountBackpressure() bool`

HasPredictedMillisUntilCountBackpressure returns a boolean if a field has been set.

### GetPredictedMillisUntilBytesBackpressure

`func (o *ConnectionStatisticsSnapshotDTO) GetPredictedMillisUntilBytesBackpressure() int64`

GetPredictedMillisUntilBytesBackpressure returns the PredictedMillisUntilBytesBackpressure field if non-nil, zero value otherwise.

### GetPredictedMillisUntilBytesBackpressureOk

`func (o *ConnectionStatisticsSnapshotDTO) GetPredictedMillisUntilBytesBackpressureOk() (*int64, bool)`

GetPredictedMillisUntilBytesBackpressureOk returns a tuple with the PredictedMillisUntilBytesBackpressure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedMillisUntilBytesBackpressure

`func (o *ConnectionStatisticsSnapshotDTO) SetPredictedMillisUntilBytesBackpressure(v int64)`

SetPredictedMillisUntilBytesBackpressure sets PredictedMillisUntilBytesBackpressure field to given value.

### HasPredictedMillisUntilBytesBackpressure

`func (o *ConnectionStatisticsSnapshotDTO) HasPredictedMillisUntilBytesBackpressure() bool`

HasPredictedMillisUntilBytesBackpressure returns a boolean if a field has been set.

### GetPredictedCountAtNextInterval

`func (o *ConnectionStatisticsSnapshotDTO) GetPredictedCountAtNextInterval() int32`

GetPredictedCountAtNextInterval returns the PredictedCountAtNextInterval field if non-nil, zero value otherwise.

### GetPredictedCountAtNextIntervalOk

`func (o *ConnectionStatisticsSnapshotDTO) GetPredictedCountAtNextIntervalOk() (*int32, bool)`

GetPredictedCountAtNextIntervalOk returns a tuple with the PredictedCountAtNextInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedCountAtNextInterval

`func (o *ConnectionStatisticsSnapshotDTO) SetPredictedCountAtNextInterval(v int32)`

SetPredictedCountAtNextInterval sets PredictedCountAtNextInterval field to given value.

### HasPredictedCountAtNextInterval

`func (o *ConnectionStatisticsSnapshotDTO) HasPredictedCountAtNextInterval() bool`

HasPredictedCountAtNextInterval returns a boolean if a field has been set.

### GetPredictedBytesAtNextInterval

`func (o *ConnectionStatisticsSnapshotDTO) GetPredictedBytesAtNextInterval() int64`

GetPredictedBytesAtNextInterval returns the PredictedBytesAtNextInterval field if non-nil, zero value otherwise.

### GetPredictedBytesAtNextIntervalOk

`func (o *ConnectionStatisticsSnapshotDTO) GetPredictedBytesAtNextIntervalOk() (*int64, bool)`

GetPredictedBytesAtNextIntervalOk returns a tuple with the PredictedBytesAtNextInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedBytesAtNextInterval

`func (o *ConnectionStatisticsSnapshotDTO) SetPredictedBytesAtNextInterval(v int64)`

SetPredictedBytesAtNextInterval sets PredictedBytesAtNextInterval field to given value.

### HasPredictedBytesAtNextInterval

`func (o *ConnectionStatisticsSnapshotDTO) HasPredictedBytesAtNextInterval() bool`

HasPredictedBytesAtNextInterval returns a boolean if a field has been set.

### GetPredictedPercentCount

`func (o *ConnectionStatisticsSnapshotDTO) GetPredictedPercentCount() int32`

GetPredictedPercentCount returns the PredictedPercentCount field if non-nil, zero value otherwise.

### GetPredictedPercentCountOk

`func (o *ConnectionStatisticsSnapshotDTO) GetPredictedPercentCountOk() (*int32, bool)`

GetPredictedPercentCountOk returns a tuple with the PredictedPercentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedPercentCount

`func (o *ConnectionStatisticsSnapshotDTO) SetPredictedPercentCount(v int32)`

SetPredictedPercentCount sets PredictedPercentCount field to given value.

### HasPredictedPercentCount

`func (o *ConnectionStatisticsSnapshotDTO) HasPredictedPercentCount() bool`

HasPredictedPercentCount returns a boolean if a field has been set.

### GetPredictedPercentBytes

`func (o *ConnectionStatisticsSnapshotDTO) GetPredictedPercentBytes() int32`

GetPredictedPercentBytes returns the PredictedPercentBytes field if non-nil, zero value otherwise.

### GetPredictedPercentBytesOk

`func (o *ConnectionStatisticsSnapshotDTO) GetPredictedPercentBytesOk() (*int32, bool)`

GetPredictedPercentBytesOk returns a tuple with the PredictedPercentBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedPercentBytes

`func (o *ConnectionStatisticsSnapshotDTO) SetPredictedPercentBytes(v int32)`

SetPredictedPercentBytes sets PredictedPercentBytes field to given value.

### HasPredictedPercentBytes

`func (o *ConnectionStatisticsSnapshotDTO) HasPredictedPercentBytes() bool`

HasPredictedPercentBytes returns a boolean if a field has been set.

### GetPredictionIntervalMillis

`func (o *ConnectionStatisticsSnapshotDTO) GetPredictionIntervalMillis() int64`

GetPredictionIntervalMillis returns the PredictionIntervalMillis field if non-nil, zero value otherwise.

### GetPredictionIntervalMillisOk

`func (o *ConnectionStatisticsSnapshotDTO) GetPredictionIntervalMillisOk() (*int64, bool)`

GetPredictionIntervalMillisOk returns a tuple with the PredictionIntervalMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictionIntervalMillis

`func (o *ConnectionStatisticsSnapshotDTO) SetPredictionIntervalMillis(v int64)`

SetPredictionIntervalMillis sets PredictionIntervalMillis field to given value.

### HasPredictionIntervalMillis

`func (o *ConnectionStatisticsSnapshotDTO) HasPredictionIntervalMillis() bool`

HasPredictionIntervalMillis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


