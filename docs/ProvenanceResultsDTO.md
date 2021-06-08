# ProvenanceResultsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProvenanceEvents** | Pointer to [**[]ProvenanceEventDTO**](ProvenanceEventDTO.md) | The provenance events that matched the search criteria. | [optional] 
**Total** | Pointer to **string** | The total number of results formatted. | [optional] 
**TotalCount** | Pointer to **int64** | The total number of results. | [optional] 
**Generated** | Pointer to **string** | Then the search was performed. | [optional] 
**OldestEvent** | Pointer to **string** | The oldest event available in the provenance repository. | [optional] 
**TimeOffset** | Pointer to **int32** | The time offset of the server that&#39;s used for event time. | [optional] 
**Errors** | Pointer to **[]string** | Any errors that occurred while performing the provenance request. | [optional] 

## Methods

### NewProvenanceResultsDTO

`func NewProvenanceResultsDTO() *ProvenanceResultsDTO`

NewProvenanceResultsDTO instantiates a new ProvenanceResultsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvenanceResultsDTOWithDefaults

`func NewProvenanceResultsDTOWithDefaults() *ProvenanceResultsDTO`

NewProvenanceResultsDTOWithDefaults instantiates a new ProvenanceResultsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvenanceEvents

`func (o *ProvenanceResultsDTO) GetProvenanceEvents() []ProvenanceEventDTO`

GetProvenanceEvents returns the ProvenanceEvents field if non-nil, zero value otherwise.

### GetProvenanceEventsOk

`func (o *ProvenanceResultsDTO) GetProvenanceEventsOk() (*[]ProvenanceEventDTO, bool)`

GetProvenanceEventsOk returns a tuple with the ProvenanceEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenanceEvents

`func (o *ProvenanceResultsDTO) SetProvenanceEvents(v []ProvenanceEventDTO)`

SetProvenanceEvents sets ProvenanceEvents field to given value.

### HasProvenanceEvents

`func (o *ProvenanceResultsDTO) HasProvenanceEvents() bool`

HasProvenanceEvents returns a boolean if a field has been set.

### GetTotal

`func (o *ProvenanceResultsDTO) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ProvenanceResultsDTO) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ProvenanceResultsDTO) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ProvenanceResultsDTO) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalCount

`func (o *ProvenanceResultsDTO) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ProvenanceResultsDTO) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ProvenanceResultsDTO) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ProvenanceResultsDTO) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetGenerated

`func (o *ProvenanceResultsDTO) GetGenerated() string`

GetGenerated returns the Generated field if non-nil, zero value otherwise.

### GetGeneratedOk

`func (o *ProvenanceResultsDTO) GetGeneratedOk() (*string, bool)`

GetGeneratedOk returns a tuple with the Generated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerated

`func (o *ProvenanceResultsDTO) SetGenerated(v string)`

SetGenerated sets Generated field to given value.

### HasGenerated

`func (o *ProvenanceResultsDTO) HasGenerated() bool`

HasGenerated returns a boolean if a field has been set.

### GetOldestEvent

`func (o *ProvenanceResultsDTO) GetOldestEvent() string`

GetOldestEvent returns the OldestEvent field if non-nil, zero value otherwise.

### GetOldestEventOk

`func (o *ProvenanceResultsDTO) GetOldestEventOk() (*string, bool)`

GetOldestEventOk returns a tuple with the OldestEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldestEvent

`func (o *ProvenanceResultsDTO) SetOldestEvent(v string)`

SetOldestEvent sets OldestEvent field to given value.

### HasOldestEvent

`func (o *ProvenanceResultsDTO) HasOldestEvent() bool`

HasOldestEvent returns a boolean if a field has been set.

### GetTimeOffset

`func (o *ProvenanceResultsDTO) GetTimeOffset() int32`

GetTimeOffset returns the TimeOffset field if non-nil, zero value otherwise.

### GetTimeOffsetOk

`func (o *ProvenanceResultsDTO) GetTimeOffsetOk() (*int32, bool)`

GetTimeOffsetOk returns a tuple with the TimeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOffset

`func (o *ProvenanceResultsDTO) SetTimeOffset(v int32)`

SetTimeOffset sets TimeOffset field to given value.

### HasTimeOffset

`func (o *ProvenanceResultsDTO) HasTimeOffset() bool`

HasTimeOffset returns a boolean if a field has been set.

### GetErrors

`func (o *ProvenanceResultsDTO) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ProvenanceResultsDTO) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ProvenanceResultsDTO) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ProvenanceResultsDTO) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


