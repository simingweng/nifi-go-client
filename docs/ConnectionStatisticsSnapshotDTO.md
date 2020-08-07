# ConnectionStatisticsSnapshotDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the connection. | [optional] 
**PredictedMillisUntilCountBackpressure** | **int64** | The predicted number of milliseconds before the connection will have backpressure applied, based on the queued count. | [optional] 
**PredictedMillisUntilBytesBackpressure** | **int64** | The predicted number of milliseconds before the connection will have backpressure applied, based on the total number of bytes in the queue. | [optional] 
**PredictedCountAtNextInterval** | **int32** | The predicted number of queued objects at the next configured interval. | [optional] 
**PredictedBytesAtNextInterval** | **int64** | The predicted total number of bytes in the queue at the next configured interval. | [optional] 
**PredictedPercentCount** | **int32** | The predicted percentage of queued objects at the next configured interval. | [optional] 
**PredictedPercentBytes** | **int32** | The predicted percentage of bytes in the queue against current threshold at the next configured interval. | [optional] 
**PredictionIntervalMillis** | **int64** | The prediction interval in seconds | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


