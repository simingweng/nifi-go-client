# ProvenanceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the provenance query. | [optional] 
**Uri** | Pointer to **string** | The URI for this query. Used for obtaining/deleting the request at a later time | [optional] 
**SubmissionTime** | Pointer to **string** | The timestamp when the query was submitted. | [optional] 
**Expiration** | Pointer to **string** | The timestamp when the query will expire. | [optional] 
**PercentCompleted** | Pointer to **int32** | The current percent complete. | [optional] 
**Finished** | Pointer to **bool** | Whether the query has finished. | [optional] 
**Request** | Pointer to [**ProvenanceRequestDTO**](ProvenanceRequestDTO.md) |  | [optional] 
**Results** | Pointer to [**ProvenanceResultsDTO**](ProvenanceResultsDTO.md) |  | [optional] 

## Methods

### NewProvenanceDTO

`func NewProvenanceDTO() *ProvenanceDTO`

NewProvenanceDTO instantiates a new ProvenanceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvenanceDTOWithDefaults

`func NewProvenanceDTOWithDefaults() *ProvenanceDTO`

NewProvenanceDTOWithDefaults instantiates a new ProvenanceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProvenanceDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvenanceDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvenanceDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProvenanceDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUri

`func (o *ProvenanceDTO) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ProvenanceDTO) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ProvenanceDTO) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ProvenanceDTO) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetSubmissionTime

`func (o *ProvenanceDTO) GetSubmissionTime() string`

GetSubmissionTime returns the SubmissionTime field if non-nil, zero value otherwise.

### GetSubmissionTimeOk

`func (o *ProvenanceDTO) GetSubmissionTimeOk() (*string, bool)`

GetSubmissionTimeOk returns a tuple with the SubmissionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionTime

`func (o *ProvenanceDTO) SetSubmissionTime(v string)`

SetSubmissionTime sets SubmissionTime field to given value.

### HasSubmissionTime

`func (o *ProvenanceDTO) HasSubmissionTime() bool`

HasSubmissionTime returns a boolean if a field has been set.

### GetExpiration

`func (o *ProvenanceDTO) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *ProvenanceDTO) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *ProvenanceDTO) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *ProvenanceDTO) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetPercentCompleted

`func (o *ProvenanceDTO) GetPercentCompleted() int32`

GetPercentCompleted returns the PercentCompleted field if non-nil, zero value otherwise.

### GetPercentCompletedOk

`func (o *ProvenanceDTO) GetPercentCompletedOk() (*int32, bool)`

GetPercentCompletedOk returns a tuple with the PercentCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentCompleted

`func (o *ProvenanceDTO) SetPercentCompleted(v int32)`

SetPercentCompleted sets PercentCompleted field to given value.

### HasPercentCompleted

`func (o *ProvenanceDTO) HasPercentCompleted() bool`

HasPercentCompleted returns a boolean if a field has been set.

### GetFinished

`func (o *ProvenanceDTO) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *ProvenanceDTO) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *ProvenanceDTO) SetFinished(v bool)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *ProvenanceDTO) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetRequest

`func (o *ProvenanceDTO) GetRequest() ProvenanceRequestDTO`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *ProvenanceDTO) GetRequestOk() (*ProvenanceRequestDTO, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *ProvenanceDTO) SetRequest(v ProvenanceRequestDTO)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *ProvenanceDTO) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResults

`func (o *ProvenanceDTO) GetResults() ProvenanceResultsDTO`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ProvenanceDTO) GetResultsOk() (*ProvenanceResultsDTO, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ProvenanceDTO) SetResults(v ProvenanceResultsDTO)`

SetResults sets Results field to given value.

### HasResults

`func (o *ProvenanceDTO) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


