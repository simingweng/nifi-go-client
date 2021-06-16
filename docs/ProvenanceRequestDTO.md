# ProvenanceRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchTerms** | Pointer to [**map[string]ProvenanceSearchValueDTO**](ProvenanceSearchValueDTO.md) | The search terms used to perform the search. | [optional] 
**ClusterNodeId** | Pointer to **string** | The id of the node in the cluster where this provenance originated. | [optional] 
**StartDate** | Pointer to **string** | The earliest event time to include in the query. | [optional] 
**EndDate** | Pointer to **string** | The latest event time to include in the query. | [optional] 
**MinimumFileSize** | Pointer to **string** | The minimum file size to include in the query. | [optional] 
**MaximumFileSize** | Pointer to **string** | The maximum file size to include in the query. | [optional] 
**MaxResults** | Pointer to **int32** | The maximum number of results to include. | [optional] 
**Summarize** | Pointer to **bool** | Whether or not to summarize provenance events returned. This property is false by default. | [optional] 
**IncrementalResults** | Pointer to **bool** | Whether or not incremental results are returned. If false, provenance events are only returned once the query completes. This property is true by default. | [optional] 

## Methods

### NewProvenanceRequestDTO

`func NewProvenanceRequestDTO() *ProvenanceRequestDTO`

NewProvenanceRequestDTO instantiates a new ProvenanceRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvenanceRequestDTOWithDefaults

`func NewProvenanceRequestDTOWithDefaults() *ProvenanceRequestDTO`

NewProvenanceRequestDTOWithDefaults instantiates a new ProvenanceRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchTerms

`func (o *ProvenanceRequestDTO) GetSearchTerms() map[string]ProvenanceSearchValueDTO`

GetSearchTerms returns the SearchTerms field if non-nil, zero value otherwise.

### GetSearchTermsOk

`func (o *ProvenanceRequestDTO) GetSearchTermsOk() (*map[string]ProvenanceSearchValueDTO, bool)`

GetSearchTermsOk returns a tuple with the SearchTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTerms

`func (o *ProvenanceRequestDTO) SetSearchTerms(v map[string]ProvenanceSearchValueDTO)`

SetSearchTerms sets SearchTerms field to given value.

### HasSearchTerms

`func (o *ProvenanceRequestDTO) HasSearchTerms() bool`

HasSearchTerms returns a boolean if a field has been set.

### GetClusterNodeId

`func (o *ProvenanceRequestDTO) GetClusterNodeId() string`

GetClusterNodeId returns the ClusterNodeId field if non-nil, zero value otherwise.

### GetClusterNodeIdOk

`func (o *ProvenanceRequestDTO) GetClusterNodeIdOk() (*string, bool)`

GetClusterNodeIdOk returns a tuple with the ClusterNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNodeId

`func (o *ProvenanceRequestDTO) SetClusterNodeId(v string)`

SetClusterNodeId sets ClusterNodeId field to given value.

### HasClusterNodeId

`func (o *ProvenanceRequestDTO) HasClusterNodeId() bool`

HasClusterNodeId returns a boolean if a field has been set.

### GetStartDate

`func (o *ProvenanceRequestDTO) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ProvenanceRequestDTO) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ProvenanceRequestDTO) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ProvenanceRequestDTO) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ProvenanceRequestDTO) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ProvenanceRequestDTO) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ProvenanceRequestDTO) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ProvenanceRequestDTO) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetMinimumFileSize

`func (o *ProvenanceRequestDTO) GetMinimumFileSize() string`

GetMinimumFileSize returns the MinimumFileSize field if non-nil, zero value otherwise.

### GetMinimumFileSizeOk

`func (o *ProvenanceRequestDTO) GetMinimumFileSizeOk() (*string, bool)`

GetMinimumFileSizeOk returns a tuple with the MinimumFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumFileSize

`func (o *ProvenanceRequestDTO) SetMinimumFileSize(v string)`

SetMinimumFileSize sets MinimumFileSize field to given value.

### HasMinimumFileSize

`func (o *ProvenanceRequestDTO) HasMinimumFileSize() bool`

HasMinimumFileSize returns a boolean if a field has been set.

### GetMaximumFileSize

`func (o *ProvenanceRequestDTO) GetMaximumFileSize() string`

GetMaximumFileSize returns the MaximumFileSize field if non-nil, zero value otherwise.

### GetMaximumFileSizeOk

`func (o *ProvenanceRequestDTO) GetMaximumFileSizeOk() (*string, bool)`

GetMaximumFileSizeOk returns a tuple with the MaximumFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumFileSize

`func (o *ProvenanceRequestDTO) SetMaximumFileSize(v string)`

SetMaximumFileSize sets MaximumFileSize field to given value.

### HasMaximumFileSize

`func (o *ProvenanceRequestDTO) HasMaximumFileSize() bool`

HasMaximumFileSize returns a boolean if a field has been set.

### GetMaxResults

`func (o *ProvenanceRequestDTO) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *ProvenanceRequestDTO) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *ProvenanceRequestDTO) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *ProvenanceRequestDTO) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetSummarize

`func (o *ProvenanceRequestDTO) GetSummarize() bool`

GetSummarize returns the Summarize field if non-nil, zero value otherwise.

### GetSummarizeOk

`func (o *ProvenanceRequestDTO) GetSummarizeOk() (*bool, bool)`

GetSummarizeOk returns a tuple with the Summarize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummarize

`func (o *ProvenanceRequestDTO) SetSummarize(v bool)`

SetSummarize sets Summarize field to given value.

### HasSummarize

`func (o *ProvenanceRequestDTO) HasSummarize() bool`

HasSummarize returns a boolean if a field has been set.

### GetIncrementalResults

`func (o *ProvenanceRequestDTO) GetIncrementalResults() bool`

GetIncrementalResults returns the IncrementalResults field if non-nil, zero value otherwise.

### GetIncrementalResultsOk

`func (o *ProvenanceRequestDTO) GetIncrementalResultsOk() (*bool, bool)`

GetIncrementalResultsOk returns a tuple with the IncrementalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalResults

`func (o *ProvenanceRequestDTO) SetIncrementalResults(v bool)`

SetIncrementalResults sets IncrementalResults field to given value.

### HasIncrementalResults

`func (o *ProvenanceRequestDTO) HasIncrementalResults() bool`

HasIncrementalResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


