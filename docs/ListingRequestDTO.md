# ListingRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id for this listing request. | [optional] 
**Uri** | Pointer to **string** | The URI for future requests to this listing request. | [optional] 
**SubmissionTime** | Pointer to **string** | The timestamp when the query was submitted. | [optional] 
**LastUpdated** | Pointer to **string** | The last time this listing request was updated. | [optional] 
**PercentCompleted** | Pointer to **int32** | The current percent complete. | [optional] 
**Finished** | Pointer to **bool** | Whether the query has finished. | [optional] 
**FailureReason** | Pointer to **string** | The reason, if any, that this listing request failed. | [optional] 
**MaxResults** | Pointer to **int32** | The maximum number of FlowFileSummary objects to return | [optional] 
**State** | Pointer to **string** | The current state of the listing request. | [optional] 
**QueueSize** | Pointer to [**QueueSizeDTO**](QueueSizeDTO.md) |  | [optional] 
**FlowFileSummaries** | Pointer to [**[]FlowFileSummaryDTO**](FlowFileSummaryDTO.md) | The FlowFile summaries. The summaries will be populated once the request has completed. | [optional] 
**SourceRunning** | Pointer to **bool** | Whether the source of the connection is running | [optional] 
**DestinationRunning** | Pointer to **bool** | Whether the destination of the connection is running | [optional] 

## Methods

### NewListingRequestDTO

`func NewListingRequestDTO() *ListingRequestDTO`

NewListingRequestDTO instantiates a new ListingRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingRequestDTOWithDefaults

`func NewListingRequestDTOWithDefaults() *ListingRequestDTO`

NewListingRequestDTOWithDefaults instantiates a new ListingRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListingRequestDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListingRequestDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListingRequestDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListingRequestDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUri

`func (o *ListingRequestDTO) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ListingRequestDTO) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ListingRequestDTO) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ListingRequestDTO) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetSubmissionTime

`func (o *ListingRequestDTO) GetSubmissionTime() string`

GetSubmissionTime returns the SubmissionTime field if non-nil, zero value otherwise.

### GetSubmissionTimeOk

`func (o *ListingRequestDTO) GetSubmissionTimeOk() (*string, bool)`

GetSubmissionTimeOk returns a tuple with the SubmissionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionTime

`func (o *ListingRequestDTO) SetSubmissionTime(v string)`

SetSubmissionTime sets SubmissionTime field to given value.

### HasSubmissionTime

`func (o *ListingRequestDTO) HasSubmissionTime() bool`

HasSubmissionTime returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListingRequestDTO) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListingRequestDTO) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListingRequestDTO) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListingRequestDTO) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetPercentCompleted

`func (o *ListingRequestDTO) GetPercentCompleted() int32`

GetPercentCompleted returns the PercentCompleted field if non-nil, zero value otherwise.

### GetPercentCompletedOk

`func (o *ListingRequestDTO) GetPercentCompletedOk() (*int32, bool)`

GetPercentCompletedOk returns a tuple with the PercentCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentCompleted

`func (o *ListingRequestDTO) SetPercentCompleted(v int32)`

SetPercentCompleted sets PercentCompleted field to given value.

### HasPercentCompleted

`func (o *ListingRequestDTO) HasPercentCompleted() bool`

HasPercentCompleted returns a boolean if a field has been set.

### GetFinished

`func (o *ListingRequestDTO) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *ListingRequestDTO) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *ListingRequestDTO) SetFinished(v bool)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *ListingRequestDTO) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetFailureReason

`func (o *ListingRequestDTO) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *ListingRequestDTO) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *ListingRequestDTO) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *ListingRequestDTO) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetMaxResults

`func (o *ListingRequestDTO) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *ListingRequestDTO) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *ListingRequestDTO) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *ListingRequestDTO) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetState

`func (o *ListingRequestDTO) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ListingRequestDTO) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ListingRequestDTO) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ListingRequestDTO) HasState() bool`

HasState returns a boolean if a field has been set.

### GetQueueSize

`func (o *ListingRequestDTO) GetQueueSize() QueueSizeDTO`

GetQueueSize returns the QueueSize field if non-nil, zero value otherwise.

### GetQueueSizeOk

`func (o *ListingRequestDTO) GetQueueSizeOk() (*QueueSizeDTO, bool)`

GetQueueSizeOk returns a tuple with the QueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueSize

`func (o *ListingRequestDTO) SetQueueSize(v QueueSizeDTO)`

SetQueueSize sets QueueSize field to given value.

### HasQueueSize

`func (o *ListingRequestDTO) HasQueueSize() bool`

HasQueueSize returns a boolean if a field has been set.

### GetFlowFileSummaries

`func (o *ListingRequestDTO) GetFlowFileSummaries() []FlowFileSummaryDTO`

GetFlowFileSummaries returns the FlowFileSummaries field if non-nil, zero value otherwise.

### GetFlowFileSummariesOk

`func (o *ListingRequestDTO) GetFlowFileSummariesOk() (*[]FlowFileSummaryDTO, bool)`

GetFlowFileSummariesOk returns a tuple with the FlowFileSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFileSummaries

`func (o *ListingRequestDTO) SetFlowFileSummaries(v []FlowFileSummaryDTO)`

SetFlowFileSummaries sets FlowFileSummaries field to given value.

### HasFlowFileSummaries

`func (o *ListingRequestDTO) HasFlowFileSummaries() bool`

HasFlowFileSummaries returns a boolean if a field has been set.

### GetSourceRunning

`func (o *ListingRequestDTO) GetSourceRunning() bool`

GetSourceRunning returns the SourceRunning field if non-nil, zero value otherwise.

### GetSourceRunningOk

`func (o *ListingRequestDTO) GetSourceRunningOk() (*bool, bool)`

GetSourceRunningOk returns a tuple with the SourceRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRunning

`func (o *ListingRequestDTO) SetSourceRunning(v bool)`

SetSourceRunning sets SourceRunning field to given value.

### HasSourceRunning

`func (o *ListingRequestDTO) HasSourceRunning() bool`

HasSourceRunning returns a boolean if a field has been set.

### GetDestinationRunning

`func (o *ListingRequestDTO) GetDestinationRunning() bool`

GetDestinationRunning returns the DestinationRunning field if non-nil, zero value otherwise.

### GetDestinationRunningOk

`func (o *ListingRequestDTO) GetDestinationRunningOk() (*bool, bool)`

GetDestinationRunningOk returns a tuple with the DestinationRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationRunning

`func (o *ListingRequestDTO) SetDestinationRunning(v bool)`

SetDestinationRunning sets DestinationRunning field to given value.

### HasDestinationRunning

`func (o *ListingRequestDTO) HasDestinationRunning() bool`

HasDestinationRunning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


