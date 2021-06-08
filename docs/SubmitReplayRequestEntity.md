# SubmitReplayRequestEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | Pointer to **int64** | The event identifier | [optional] 
**ClusterNodeId** | Pointer to **string** | The identifier of the node where to submit the replay request. | [optional] 

## Methods

### NewSubmitReplayRequestEntity

`func NewSubmitReplayRequestEntity() *SubmitReplayRequestEntity`

NewSubmitReplayRequestEntity instantiates a new SubmitReplayRequestEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitReplayRequestEntityWithDefaults

`func NewSubmitReplayRequestEntityWithDefaults() *SubmitReplayRequestEntity`

NewSubmitReplayRequestEntityWithDefaults instantiates a new SubmitReplayRequestEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *SubmitReplayRequestEntity) GetEventId() int64`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SubmitReplayRequestEntity) GetEventIdOk() (*int64, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SubmitReplayRequestEntity) SetEventId(v int64)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *SubmitReplayRequestEntity) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetClusterNodeId

`func (o *SubmitReplayRequestEntity) GetClusterNodeId() string`

GetClusterNodeId returns the ClusterNodeId field if non-nil, zero value otherwise.

### GetClusterNodeIdOk

`func (o *SubmitReplayRequestEntity) GetClusterNodeIdOk() (*string, bool)`

GetClusterNodeIdOk returns a tuple with the ClusterNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNodeId

`func (o *SubmitReplayRequestEntity) SetClusterNodeId(v string)`

SetClusterNodeId sets ClusterNodeId field to given value.

### HasClusterNodeId

`func (o *SubmitReplayRequestEntity) HasClusterNodeId() bool`

HasClusterNodeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


