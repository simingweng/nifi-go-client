# VersionedFlowUpdateRequestEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessGroupRevision** | Pointer to [**RevisionDTO**](RevisionDTO.md) |  | [optional] 
**Request** | Pointer to [**VersionedFlowUpdateRequestDTO**](VersionedFlowUpdateRequestDTO.md) |  | [optional] 

## Methods

### NewVersionedFlowUpdateRequestEntity

`func NewVersionedFlowUpdateRequestEntity() *VersionedFlowUpdateRequestEntity`

NewVersionedFlowUpdateRequestEntity instantiates a new VersionedFlowUpdateRequestEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionedFlowUpdateRequestEntityWithDefaults

`func NewVersionedFlowUpdateRequestEntityWithDefaults() *VersionedFlowUpdateRequestEntity`

NewVersionedFlowUpdateRequestEntityWithDefaults instantiates a new VersionedFlowUpdateRequestEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessGroupRevision

`func (o *VersionedFlowUpdateRequestEntity) GetProcessGroupRevision() RevisionDTO`

GetProcessGroupRevision returns the ProcessGroupRevision field if non-nil, zero value otherwise.

### GetProcessGroupRevisionOk

`func (o *VersionedFlowUpdateRequestEntity) GetProcessGroupRevisionOk() (*RevisionDTO, bool)`

GetProcessGroupRevisionOk returns a tuple with the ProcessGroupRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupRevision

`func (o *VersionedFlowUpdateRequestEntity) SetProcessGroupRevision(v RevisionDTO)`

SetProcessGroupRevision sets ProcessGroupRevision field to given value.

### HasProcessGroupRevision

`func (o *VersionedFlowUpdateRequestEntity) HasProcessGroupRevision() bool`

HasProcessGroupRevision returns a boolean if a field has been set.

### GetRequest

`func (o *VersionedFlowUpdateRequestEntity) GetRequest() VersionedFlowUpdateRequestDTO`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *VersionedFlowUpdateRequestEntity) GetRequestOk() (*VersionedFlowUpdateRequestDTO, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *VersionedFlowUpdateRequestEntity) SetRequest(v VersionedFlowUpdateRequestDTO)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *VersionedFlowUpdateRequestEntity) HasRequest() bool`

HasRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


