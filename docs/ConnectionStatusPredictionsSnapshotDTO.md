# ConnectionStatusPredictionsSnapshotDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PredictedMillisUntilCountBackpressure** | Pointer to **int64** | The predicted number of milliseconds before the connection will have backpressure applied, based on the queued count. | [optional] 
**PredictedMillisUntilBytesBackpressure** | Pointer to **int64** | The predicted number of milliseconds before the connection will have backpressure applied, based on the total number of bytes in the queue. | [optional] 
**PredictionIntervalSeconds** | Pointer to **int32** | The configured interval (in seconds) for predicting connection queue count and size (and percent usage). | [optional] 
**PredictedCountAtNextInterval** | Pointer to **int32** | The predicted number of queued objects at the next configured interval. | [optional] 
**PredictedBytesAtNextInterval** | Pointer to **int64** | The predicted total number of bytes in the queue at the next configured interval. | [optional] 
**PredictedPercentCount** | Pointer to **int32** | Predicted connection percent use regarding queued flow files count and backpressure threshold if configured. | [optional] 
**PredictedPercentBytes** | Pointer to **int32** | Predicted connection percent use regarding queued flow files size and backpressure threshold if configured. | [optional] 

## Methods

### NewConnectionStatusPredictionsSnapshotDTO

`func NewConnectionStatusPredictionsSnapshotDTO() *ConnectionStatusPredictionsSnapshotDTO`

NewConnectionStatusPredictionsSnapshotDTO instantiates a new ConnectionStatusPredictionsSnapshotDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionStatusPredictionsSnapshotDTOWithDefaults

`func NewConnectionStatusPredictionsSnapshotDTOWithDefaults() *ConnectionStatusPredictionsSnapshotDTO`

NewConnectionStatusPredictionsSnapshotDTOWithDefaults instantiates a new ConnectionStatusPredictionsSnapshotDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPredictedMillisUntilCountBackpressure

`func (o *ConnectionStatusPredictionsSnapshotDTO) GetPredictedMillisUntilCountBackpressure() int64`

GetPredictedMillisUntilCountBackpressure returns the PredictedMillisUntilCountBackpressure field if non-nil, zero value otherwise.

### GetPredictedMillisUntilCountBackpressureOk

`func (o *ConnectionStatusPredictionsSnapshotDTO) GetPredictedMillisUntilCountBackpressureOk() (*int64, bool)`

GetPredictedMillisUntilCountBackpressureOk returns a tuple with the PredictedMillisUntilCountBackpressure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedMillisUntilCountBackpressure

`func (o *ConnectionStatusPredictionsSnapshotDTO) SetPredictedMillisUntilCountBackpressure(v int64)`

SetPredictedMillisUntilCountBackpressure sets PredictedMillisUntilCountBackpressure field to given value.

### HasPredictedMillisUntilCountBackpressure

`func (o *ConnectionStatusPredictionsSnapshotDTO) HasPredictedMillisUntilCountBackpressure() bool`

HasPredictedMillisUntilCountBackpressure returns a boolean if a field has been set.

### GetPredictedMillisUntilBytesBackpressure

`func (o *ConnectionStatusPredictionsSnapshotDTO) GetPredictedMillisUntilBytesBackpressure() int64`

GetPredictedMillisUntilBytesBackpressure returns the PredictedMillisUntilBytesBackpressure field if non-nil, zero value otherwise.

### GetPredictedMillisUntilBytesBackpressureOk

`func (o *ConnectionStatusPredictionsSnapshotDTO) GetPredictedMillisUntilBytesBackpressureOk() (*int64, bool)`

GetPredictedMillisUntilBytesBackpressureOk returns a tuple with the PredictedMillisUntilBytesBackpressure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedMillisUntilBytesBackpressure

`func (o *ConnectionStatusPredictionsSnapshotDTO) SetPredictedMillisUntilBytesBackpressure(v int64)`

SetPredictedMillisUntilBytesBackpressure sets PredictedMillisUntilBytesBackpressure field to given value.

### HasPredictedMillisUntilBytesBackpressure

`func (o *ConnectionStatusPredictionsSnapshotDTO) HasPredictedMillisUntilBytesBackpressure() bool`

HasPredictedMillisUntilBytesBackpressure returns a boolean if a field has been set.

### GetPredictionIntervalSeconds

`func (o *ConnectionStatusPredictionsSnapshotDTO) GetPredictionIntervalSeconds() int32`

GetPredictionIntervalSeconds returns the PredictionIntervalSeconds field if non-nil, zero value otherwise.

### GetPredictionIntervalSecondsOk

`func (o *ConnectionStatusPredictionsSnapshotDTO) GetPredictionIntervalSecondsOk() (*int32, bool)`

GetPredictionIntervalSecondsOk returns a tuple with the PredictionIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictionIntervalSeconds

`func (o *ConnectionStatusPredictionsSnapshotDTO) SetPredictionIntervalSeconds(v int32)`

SetPredictionIntervalSeconds sets PredictionIntervalSeconds field to given value.

### HasPredictionIntervalSeconds

`func (o *ConnectionStatusPredictionsSnapshotDTO) HasPredictionIntervalSeconds() bool`

HasPredictionIntervalSeconds returns a boolean if a field has been set.

### GetPredictedCountAtNextInterval

`func (o *ConnectionStatusPredictionsSnapshotDTO) GetPredictedCountAtNextInterval() int32`

GetPredictedCountAtNextInterval returns the PredictedCountAtNextInterval field if non-nil, zero value otherwise.

### GetPredictedCountAtNextIntervalOk

`func (o *ConnectionStatusPredictionsSnapshotDTO) GetPredictedCountAtNextIntervalOk() (*int32, bool)`

GetPredictedCountAtNextIntervalOk returns a tuple with the PredictedCountAtNextInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedCountAtNextInterval

`func (o *ConnectionStatusPredictionsSnapshotDTO) SetPredictedCountAtNextInterval(v int32)`

SetPredictedCountAtNextInterval sets PredictedCountAtNextInterval field to given value.

### HasPredictedCountAtNextInterval

`func (o *ConnectionStatusPredictionsSnapshotDTO) HasPredictedCountAtNextInterval() bool`

HasPredictedCountAtNextInterval returns a boolean if a field has been set.

### GetPredictedBytesAtNextInterval

`func (o *ConnectionStatusPredictionsSnapshotDTO) GetPredictedBytesAtNextInterval() int64`

GetPredictedBytesAtNextInterval returns the PredictedBytesAtNextInterval field if non-nil, zero value otherwise.

### GetPredictedBytesAtNextIntervalOk

`func (o *ConnectionStatusPredictionsSnapshotDTO) GetPredictedBytesAtNextIntervalOk() (*int64, bool)`

GetPredictedBytesAtNextIntervalOk returns a tuple with the PredictedBytesAtNextInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedBytesAtNextInterval

`func (o *ConnectionStatusPredictionsSnapshotDTO) SetPredictedBytesAtNextInterval(v int64)`

SetPredictedBytesAtNextInterval sets PredictedBytesAtNextInterval field to given value.

### HasPredictedBytesAtNextInterval

`func (o *ConnectionStatusPredictionsSnapshotDTO) HasPredictedBytesAtNextInterval() bool`

HasPredictedBytesAtNextInterval returns a boolean if a field has been set.

### GetPredictedPercentCount

`func (o *ConnectionStatusPredictionsSnapshotDTO) GetPredictedPercentCount() int32`

GetPredictedPercentCount returns the PredictedPercentCount field if non-nil, zero value otherwise.

### GetPredictedPercentCountOk

`func (o *ConnectionStatusPredictionsSnapshotDTO) GetPredictedPercentCountOk() (*int32, bool)`

GetPredictedPercentCountOk returns a tuple with the PredictedPercentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedPercentCount

`func (o *ConnectionStatusPredictionsSnapshotDTO) SetPredictedPercentCount(v int32)`

SetPredictedPercentCount sets PredictedPercentCount field to given value.

### HasPredictedPercentCount

`func (o *ConnectionStatusPredictionsSnapshotDTO) HasPredictedPercentCount() bool`

HasPredictedPercentCount returns a boolean if a field has been set.

### GetPredictedPercentBytes

`func (o *ConnectionStatusPredictionsSnapshotDTO) GetPredictedPercentBytes() int32`

GetPredictedPercentBytes returns the PredictedPercentBytes field if non-nil, zero value otherwise.

### GetPredictedPercentBytesOk

`func (o *ConnectionStatusPredictionsSnapshotDTO) GetPredictedPercentBytesOk() (*int32, bool)`

GetPredictedPercentBytesOk returns a tuple with the PredictedPercentBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedPercentBytes

`func (o *ConnectionStatusPredictionsSnapshotDTO) SetPredictedPercentBytes(v int32)`

SetPredictedPercentBytes sets PredictedPercentBytes field to given value.

### HasPredictedPercentBytes

`func (o *ConnectionStatusPredictionsSnapshotDTO) HasPredictedPercentBytes() bool`

HasPredictedPercentBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


