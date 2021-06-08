# ProvenanceLinkDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | Pointer to **string** | The source node id of the link. | [optional] 
**TargetId** | Pointer to **string** | The target node id of the link. | [optional] 
**FlowFileUuid** | Pointer to **string** | The flowfile uuid that traversed the link. | [optional] 
**Timestamp** | Pointer to **string** | The timestamp of the link (based on the destination). | [optional] 
**Millis** | Pointer to **int64** | The timestamp of this link in milliseconds. | [optional] 

## Methods

### NewProvenanceLinkDTO

`func NewProvenanceLinkDTO() *ProvenanceLinkDTO`

NewProvenanceLinkDTO instantiates a new ProvenanceLinkDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvenanceLinkDTOWithDefaults

`func NewProvenanceLinkDTOWithDefaults() *ProvenanceLinkDTO`

NewProvenanceLinkDTOWithDefaults instantiates a new ProvenanceLinkDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *ProvenanceLinkDTO) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ProvenanceLinkDTO) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ProvenanceLinkDTO) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ProvenanceLinkDTO) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetTargetId

`func (o *ProvenanceLinkDTO) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *ProvenanceLinkDTO) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *ProvenanceLinkDTO) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *ProvenanceLinkDTO) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetFlowFileUuid

`func (o *ProvenanceLinkDTO) GetFlowFileUuid() string`

GetFlowFileUuid returns the FlowFileUuid field if non-nil, zero value otherwise.

### GetFlowFileUuidOk

`func (o *ProvenanceLinkDTO) GetFlowFileUuidOk() (*string, bool)`

GetFlowFileUuidOk returns a tuple with the FlowFileUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFileUuid

`func (o *ProvenanceLinkDTO) SetFlowFileUuid(v string)`

SetFlowFileUuid sets FlowFileUuid field to given value.

### HasFlowFileUuid

`func (o *ProvenanceLinkDTO) HasFlowFileUuid() bool`

HasFlowFileUuid returns a boolean if a field has been set.

### GetTimestamp

`func (o *ProvenanceLinkDTO) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProvenanceLinkDTO) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProvenanceLinkDTO) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProvenanceLinkDTO) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetMillis

`func (o *ProvenanceLinkDTO) GetMillis() int64`

GetMillis returns the Millis field if non-nil, zero value otherwise.

### GetMillisOk

`func (o *ProvenanceLinkDTO) GetMillisOk() (*int64, bool)`

GetMillisOk returns a tuple with the Millis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMillis

`func (o *ProvenanceLinkDTO) SetMillis(v int64)`

SetMillis sets Millis field to given value.

### HasMillis

`func (o *ProvenanceLinkDTO) HasMillis() bool`

HasMillis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


