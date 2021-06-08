# DropRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id for this drop request. | [optional] 
**Uri** | Pointer to **string** | The URI for future requests to this drop request. | [optional] 
**SubmissionTime** | Pointer to **string** | The timestamp when the query was submitted. | [optional] 
**LastUpdated** | Pointer to **string** | The last time this drop request was updated. | [optional] 
**PercentCompleted** | Pointer to **int32** | The current percent complete. | [optional] 
**Finished** | Pointer to **bool** | Whether the query has finished. | [optional] 
**FailureReason** | Pointer to **string** | The reason, if any, that this drop request failed. | [optional] 
**CurrentCount** | Pointer to **int32** | The number of flow files currently queued. | [optional] 
**CurrentSize** | Pointer to **int64** | The size of flow files currently queued in bytes. | [optional] 
**Current** | Pointer to **string** | The count and size of flow files currently queued. | [optional] 
**OriginalCount** | Pointer to **int32** | The number of flow files to be dropped as a result of this request. | [optional] 
**OriginalSize** | Pointer to **int64** | The size of flow files to be dropped as a result of this request in bytes. | [optional] 
**Original** | Pointer to **string** | The count and size of flow files to be dropped as a result of this request. | [optional] 
**DroppedCount** | Pointer to **int32** | The number of flow files that have been dropped thus far. | [optional] 
**DroppedSize** | Pointer to **int64** | The size of flow files that have been dropped thus far in bytes. | [optional] 
**Dropped** | Pointer to **string** | The count and size of flow files that have been dropped thus far. | [optional] 
**State** | Pointer to **string** | The current state of the drop request. | [optional] 

## Methods

### NewDropRequestDTO

`func NewDropRequestDTO() *DropRequestDTO`

NewDropRequestDTO instantiates a new DropRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDropRequestDTOWithDefaults

`func NewDropRequestDTOWithDefaults() *DropRequestDTO`

NewDropRequestDTOWithDefaults instantiates a new DropRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DropRequestDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DropRequestDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DropRequestDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DropRequestDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUri

`func (o *DropRequestDTO) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *DropRequestDTO) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *DropRequestDTO) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *DropRequestDTO) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetSubmissionTime

`func (o *DropRequestDTO) GetSubmissionTime() string`

GetSubmissionTime returns the SubmissionTime field if non-nil, zero value otherwise.

### GetSubmissionTimeOk

`func (o *DropRequestDTO) GetSubmissionTimeOk() (*string, bool)`

GetSubmissionTimeOk returns a tuple with the SubmissionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionTime

`func (o *DropRequestDTO) SetSubmissionTime(v string)`

SetSubmissionTime sets SubmissionTime field to given value.

### HasSubmissionTime

`func (o *DropRequestDTO) HasSubmissionTime() bool`

HasSubmissionTime returns a boolean if a field has been set.

### GetLastUpdated

`func (o *DropRequestDTO) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *DropRequestDTO) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *DropRequestDTO) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *DropRequestDTO) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetPercentCompleted

`func (o *DropRequestDTO) GetPercentCompleted() int32`

GetPercentCompleted returns the PercentCompleted field if non-nil, zero value otherwise.

### GetPercentCompletedOk

`func (o *DropRequestDTO) GetPercentCompletedOk() (*int32, bool)`

GetPercentCompletedOk returns a tuple with the PercentCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentCompleted

`func (o *DropRequestDTO) SetPercentCompleted(v int32)`

SetPercentCompleted sets PercentCompleted field to given value.

### HasPercentCompleted

`func (o *DropRequestDTO) HasPercentCompleted() bool`

HasPercentCompleted returns a boolean if a field has been set.

### GetFinished

`func (o *DropRequestDTO) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *DropRequestDTO) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *DropRequestDTO) SetFinished(v bool)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *DropRequestDTO) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetFailureReason

`func (o *DropRequestDTO) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *DropRequestDTO) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *DropRequestDTO) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *DropRequestDTO) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetCurrentCount

`func (o *DropRequestDTO) GetCurrentCount() int32`

GetCurrentCount returns the CurrentCount field if non-nil, zero value otherwise.

### GetCurrentCountOk

`func (o *DropRequestDTO) GetCurrentCountOk() (*int32, bool)`

GetCurrentCountOk returns a tuple with the CurrentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCount

`func (o *DropRequestDTO) SetCurrentCount(v int32)`

SetCurrentCount sets CurrentCount field to given value.

### HasCurrentCount

`func (o *DropRequestDTO) HasCurrentCount() bool`

HasCurrentCount returns a boolean if a field has been set.

### GetCurrentSize

`func (o *DropRequestDTO) GetCurrentSize() int64`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *DropRequestDTO) GetCurrentSizeOk() (*int64, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *DropRequestDTO) SetCurrentSize(v int64)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *DropRequestDTO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetCurrent

`func (o *DropRequestDTO) GetCurrent() string`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *DropRequestDTO) GetCurrentOk() (*string, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *DropRequestDTO) SetCurrent(v string)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *DropRequestDTO) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetOriginalCount

`func (o *DropRequestDTO) GetOriginalCount() int32`

GetOriginalCount returns the OriginalCount field if non-nil, zero value otherwise.

### GetOriginalCountOk

`func (o *DropRequestDTO) GetOriginalCountOk() (*int32, bool)`

GetOriginalCountOk returns a tuple with the OriginalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCount

`func (o *DropRequestDTO) SetOriginalCount(v int32)`

SetOriginalCount sets OriginalCount field to given value.

### HasOriginalCount

`func (o *DropRequestDTO) HasOriginalCount() bool`

HasOriginalCount returns a boolean if a field has been set.

### GetOriginalSize

`func (o *DropRequestDTO) GetOriginalSize() int64`

GetOriginalSize returns the OriginalSize field if non-nil, zero value otherwise.

### GetOriginalSizeOk

`func (o *DropRequestDTO) GetOriginalSizeOk() (*int64, bool)`

GetOriginalSizeOk returns a tuple with the OriginalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSize

`func (o *DropRequestDTO) SetOriginalSize(v int64)`

SetOriginalSize sets OriginalSize field to given value.

### HasOriginalSize

`func (o *DropRequestDTO) HasOriginalSize() bool`

HasOriginalSize returns a boolean if a field has been set.

### GetOriginal

`func (o *DropRequestDTO) GetOriginal() string`

GetOriginal returns the Original field if non-nil, zero value otherwise.

### GetOriginalOk

`func (o *DropRequestDTO) GetOriginalOk() (*string, bool)`

GetOriginalOk returns a tuple with the Original field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginal

`func (o *DropRequestDTO) SetOriginal(v string)`

SetOriginal sets Original field to given value.

### HasOriginal

`func (o *DropRequestDTO) HasOriginal() bool`

HasOriginal returns a boolean if a field has been set.

### GetDroppedCount

`func (o *DropRequestDTO) GetDroppedCount() int32`

GetDroppedCount returns the DroppedCount field if non-nil, zero value otherwise.

### GetDroppedCountOk

`func (o *DropRequestDTO) GetDroppedCountOk() (*int32, bool)`

GetDroppedCountOk returns a tuple with the DroppedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroppedCount

`func (o *DropRequestDTO) SetDroppedCount(v int32)`

SetDroppedCount sets DroppedCount field to given value.

### HasDroppedCount

`func (o *DropRequestDTO) HasDroppedCount() bool`

HasDroppedCount returns a boolean if a field has been set.

### GetDroppedSize

`func (o *DropRequestDTO) GetDroppedSize() int64`

GetDroppedSize returns the DroppedSize field if non-nil, zero value otherwise.

### GetDroppedSizeOk

`func (o *DropRequestDTO) GetDroppedSizeOk() (*int64, bool)`

GetDroppedSizeOk returns a tuple with the DroppedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroppedSize

`func (o *DropRequestDTO) SetDroppedSize(v int64)`

SetDroppedSize sets DroppedSize field to given value.

### HasDroppedSize

`func (o *DropRequestDTO) HasDroppedSize() bool`

HasDroppedSize returns a boolean if a field has been set.

### GetDropped

`func (o *DropRequestDTO) GetDropped() string`

GetDropped returns the Dropped field if non-nil, zero value otherwise.

### GetDroppedOk

`func (o *DropRequestDTO) GetDroppedOk() (*string, bool)`

GetDroppedOk returns a tuple with the Dropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropped

`func (o *DropRequestDTO) SetDropped(v string)`

SetDropped sets Dropped field to given value.

### HasDropped

`func (o *DropRequestDTO) HasDropped() bool`

HasDropped returns a boolean if a field has been set.

### GetState

`func (o *DropRequestDTO) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DropRequestDTO) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DropRequestDTO) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DropRequestDTO) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


