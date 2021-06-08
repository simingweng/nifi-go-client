# LineageDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of this lineage query. | [optional] 
**Uri** | Pointer to **string** | The URI for this lineage query for later retrieval and deletion. | [optional] 
**SubmissionTime** | Pointer to **string** | When the lineage query was submitted. | [optional] 
**Expiration** | Pointer to **string** | When the lineage query will expire. | [optional] 
**PercentCompleted** | Pointer to **int32** | The percent complete for the lineage query. | [optional] 
**Finished** | Pointer to **bool** | Whether the lineage query has finished. | [optional] 
**Request** | Pointer to [**LineageRequestDTO**](LineageRequestDTO.md) |  | [optional] 
**Results** | Pointer to [**LineageResultsDTO**](LineageResultsDTO.md) |  | [optional] 

## Methods

### NewLineageDTO

`func NewLineageDTO() *LineageDTO`

NewLineageDTO instantiates a new LineageDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineageDTOWithDefaults

`func NewLineageDTOWithDefaults() *LineageDTO`

NewLineageDTOWithDefaults instantiates a new LineageDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LineageDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LineageDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LineageDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LineageDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUri

`func (o *LineageDTO) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *LineageDTO) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *LineageDTO) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *LineageDTO) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetSubmissionTime

`func (o *LineageDTO) GetSubmissionTime() string`

GetSubmissionTime returns the SubmissionTime field if non-nil, zero value otherwise.

### GetSubmissionTimeOk

`func (o *LineageDTO) GetSubmissionTimeOk() (*string, bool)`

GetSubmissionTimeOk returns a tuple with the SubmissionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionTime

`func (o *LineageDTO) SetSubmissionTime(v string)`

SetSubmissionTime sets SubmissionTime field to given value.

### HasSubmissionTime

`func (o *LineageDTO) HasSubmissionTime() bool`

HasSubmissionTime returns a boolean if a field has been set.

### GetExpiration

`func (o *LineageDTO) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *LineageDTO) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *LineageDTO) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *LineageDTO) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetPercentCompleted

`func (o *LineageDTO) GetPercentCompleted() int32`

GetPercentCompleted returns the PercentCompleted field if non-nil, zero value otherwise.

### GetPercentCompletedOk

`func (o *LineageDTO) GetPercentCompletedOk() (*int32, bool)`

GetPercentCompletedOk returns a tuple with the PercentCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentCompleted

`func (o *LineageDTO) SetPercentCompleted(v int32)`

SetPercentCompleted sets PercentCompleted field to given value.

### HasPercentCompleted

`func (o *LineageDTO) HasPercentCompleted() bool`

HasPercentCompleted returns a boolean if a field has been set.

### GetFinished

`func (o *LineageDTO) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *LineageDTO) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *LineageDTO) SetFinished(v bool)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *LineageDTO) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetRequest

`func (o *LineageDTO) GetRequest() LineageRequestDTO`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *LineageDTO) GetRequestOk() (*LineageRequestDTO, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *LineageDTO) SetRequest(v LineageRequestDTO)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *LineageDTO) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResults

`func (o *LineageDTO) GetResults() LineageResultsDTO`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *LineageDTO) GetResultsOk() (*LineageResultsDTO, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *LineageDTO) SetResults(v LineageResultsDTO)`

SetResults sets Results field to given value.

### HasResults

`func (o *LineageDTO) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


