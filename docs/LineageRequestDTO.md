# LineageRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | Pointer to **int64** | The event id that was used to generate this lineage, if applicable. The event id is allowed for any type of lineageRequestType. If the lineageRequestType is FLOWFILE and the flowfile uuid is also included in the request, the event id will be ignored. | [optional] 
**LineageRequestType** | Pointer to **string** | The type of lineage request. PARENTS will return the lineage for the flowfiles that are parents of the specified event. CHILDREN will return the lineage for the flowfiles that are children of the specified event. FLOWFILE will return the lineage for the specified flowfile. | [optional] 
**Uuid** | Pointer to **string** | The flowfile uuid that was used to generate the lineage. The flowfile uuid is only allowed when the lineageRequestType is FLOWFILE and will take precedence over event id. | [optional] 
**ClusterNodeId** | Pointer to **string** | The id of the node where this lineage originated if clustered. | [optional] 

## Methods

### NewLineageRequestDTO

`func NewLineageRequestDTO() *LineageRequestDTO`

NewLineageRequestDTO instantiates a new LineageRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineageRequestDTOWithDefaults

`func NewLineageRequestDTOWithDefaults() *LineageRequestDTO`

NewLineageRequestDTOWithDefaults instantiates a new LineageRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *LineageRequestDTO) GetEventId() int64`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *LineageRequestDTO) GetEventIdOk() (*int64, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *LineageRequestDTO) SetEventId(v int64)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *LineageRequestDTO) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetLineageRequestType

`func (o *LineageRequestDTO) GetLineageRequestType() string`

GetLineageRequestType returns the LineageRequestType field if non-nil, zero value otherwise.

### GetLineageRequestTypeOk

`func (o *LineageRequestDTO) GetLineageRequestTypeOk() (*string, bool)`

GetLineageRequestTypeOk returns a tuple with the LineageRequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineageRequestType

`func (o *LineageRequestDTO) SetLineageRequestType(v string)`

SetLineageRequestType sets LineageRequestType field to given value.

### HasLineageRequestType

`func (o *LineageRequestDTO) HasLineageRequestType() bool`

HasLineageRequestType returns a boolean if a field has been set.

### GetUuid

`func (o *LineageRequestDTO) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *LineageRequestDTO) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *LineageRequestDTO) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *LineageRequestDTO) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetClusterNodeId

`func (o *LineageRequestDTO) GetClusterNodeId() string`

GetClusterNodeId returns the ClusterNodeId field if non-nil, zero value otherwise.

### GetClusterNodeIdOk

`func (o *LineageRequestDTO) GetClusterNodeIdOk() (*string, bool)`

GetClusterNodeIdOk returns a tuple with the ClusterNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNodeId

`func (o *LineageRequestDTO) SetClusterNodeId(v string)`

SetClusterNodeId sets ClusterNodeId field to given value.

### HasClusterNodeId

`func (o *LineageRequestDTO) HasClusterNodeId() bool`

HasClusterNodeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


