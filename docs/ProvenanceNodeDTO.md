# ProvenanceNodeDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the node. | [optional] 
**FlowFileUuid** | Pointer to **string** | The uuid of the flowfile associated with the provenance event. | [optional] 
**ParentUuids** | Pointer to **[]string** | The uuid of the parent flowfiles of the provenance event. | [optional] 
**ChildUuids** | Pointer to **[]string** | The uuid of the childrent flowfiles of the provenance event. | [optional] 
**ClusterNodeIdentifier** | Pointer to **string** | The identifier of the node that this event/flowfile originated from. | [optional] 
**Type** | Pointer to **string** | The type of the node. | [optional] 
**EventType** | Pointer to **string** | If the type is EVENT, this is the type of event. | [optional] 
**Millis** | Pointer to **int64** | The timestamp of the node in milliseconds. | [optional] 
**Timestamp** | Pointer to **string** | The timestamp of the node formatted. | [optional] 

## Methods

### NewProvenanceNodeDTO

`func NewProvenanceNodeDTO() *ProvenanceNodeDTO`

NewProvenanceNodeDTO instantiates a new ProvenanceNodeDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvenanceNodeDTOWithDefaults

`func NewProvenanceNodeDTOWithDefaults() *ProvenanceNodeDTO`

NewProvenanceNodeDTOWithDefaults instantiates a new ProvenanceNodeDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProvenanceNodeDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvenanceNodeDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvenanceNodeDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProvenanceNodeDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFlowFileUuid

`func (o *ProvenanceNodeDTO) GetFlowFileUuid() string`

GetFlowFileUuid returns the FlowFileUuid field if non-nil, zero value otherwise.

### GetFlowFileUuidOk

`func (o *ProvenanceNodeDTO) GetFlowFileUuidOk() (*string, bool)`

GetFlowFileUuidOk returns a tuple with the FlowFileUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFileUuid

`func (o *ProvenanceNodeDTO) SetFlowFileUuid(v string)`

SetFlowFileUuid sets FlowFileUuid field to given value.

### HasFlowFileUuid

`func (o *ProvenanceNodeDTO) HasFlowFileUuid() bool`

HasFlowFileUuid returns a boolean if a field has been set.

### GetParentUuids

`func (o *ProvenanceNodeDTO) GetParentUuids() []string`

GetParentUuids returns the ParentUuids field if non-nil, zero value otherwise.

### GetParentUuidsOk

`func (o *ProvenanceNodeDTO) GetParentUuidsOk() (*[]string, bool)`

GetParentUuidsOk returns a tuple with the ParentUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUuids

`func (o *ProvenanceNodeDTO) SetParentUuids(v []string)`

SetParentUuids sets ParentUuids field to given value.

### HasParentUuids

`func (o *ProvenanceNodeDTO) HasParentUuids() bool`

HasParentUuids returns a boolean if a field has been set.

### GetChildUuids

`func (o *ProvenanceNodeDTO) GetChildUuids() []string`

GetChildUuids returns the ChildUuids field if non-nil, zero value otherwise.

### GetChildUuidsOk

`func (o *ProvenanceNodeDTO) GetChildUuidsOk() (*[]string, bool)`

GetChildUuidsOk returns a tuple with the ChildUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildUuids

`func (o *ProvenanceNodeDTO) SetChildUuids(v []string)`

SetChildUuids sets ChildUuids field to given value.

### HasChildUuids

`func (o *ProvenanceNodeDTO) HasChildUuids() bool`

HasChildUuids returns a boolean if a field has been set.

### GetClusterNodeIdentifier

`func (o *ProvenanceNodeDTO) GetClusterNodeIdentifier() string`

GetClusterNodeIdentifier returns the ClusterNodeIdentifier field if non-nil, zero value otherwise.

### GetClusterNodeIdentifierOk

`func (o *ProvenanceNodeDTO) GetClusterNodeIdentifierOk() (*string, bool)`

GetClusterNodeIdentifierOk returns a tuple with the ClusterNodeIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNodeIdentifier

`func (o *ProvenanceNodeDTO) SetClusterNodeIdentifier(v string)`

SetClusterNodeIdentifier sets ClusterNodeIdentifier field to given value.

### HasClusterNodeIdentifier

`func (o *ProvenanceNodeDTO) HasClusterNodeIdentifier() bool`

HasClusterNodeIdentifier returns a boolean if a field has been set.

### GetType

`func (o *ProvenanceNodeDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProvenanceNodeDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProvenanceNodeDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProvenanceNodeDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEventType

`func (o *ProvenanceNodeDTO) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ProvenanceNodeDTO) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ProvenanceNodeDTO) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ProvenanceNodeDTO) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetMillis

`func (o *ProvenanceNodeDTO) GetMillis() int64`

GetMillis returns the Millis field if non-nil, zero value otherwise.

### GetMillisOk

`func (o *ProvenanceNodeDTO) GetMillisOk() (*int64, bool)`

GetMillisOk returns a tuple with the Millis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMillis

`func (o *ProvenanceNodeDTO) SetMillis(v int64)`

SetMillis sets Millis field to given value.

### HasMillis

`func (o *ProvenanceNodeDTO) HasMillis() bool`

HasMillis returns a boolean if a field has been set.

### GetTimestamp

`func (o *ProvenanceNodeDTO) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProvenanceNodeDTO) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProvenanceNodeDTO) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProvenanceNodeDTO) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


