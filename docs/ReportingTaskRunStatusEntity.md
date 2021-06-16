# ReportingTaskRunStatusEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | Pointer to [**RevisionDTO**](RevisionDTO.md) |  | [optional] 
**State** | Pointer to **string** | The run status of the ReportingTask. | [optional] 
**DisconnectedNodeAcknowledged** | Pointer to **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] 

## Methods

### NewReportingTaskRunStatusEntity

`func NewReportingTaskRunStatusEntity() *ReportingTaskRunStatusEntity`

NewReportingTaskRunStatusEntity instantiates a new ReportingTaskRunStatusEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportingTaskRunStatusEntityWithDefaults

`func NewReportingTaskRunStatusEntityWithDefaults() *ReportingTaskRunStatusEntity`

NewReportingTaskRunStatusEntityWithDefaults instantiates a new ReportingTaskRunStatusEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *ReportingTaskRunStatusEntity) GetRevision() RevisionDTO`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ReportingTaskRunStatusEntity) GetRevisionOk() (*RevisionDTO, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ReportingTaskRunStatusEntity) SetRevision(v RevisionDTO)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ReportingTaskRunStatusEntity) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetState

`func (o *ReportingTaskRunStatusEntity) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ReportingTaskRunStatusEntity) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ReportingTaskRunStatusEntity) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ReportingTaskRunStatusEntity) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *ReportingTaskRunStatusEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *ReportingTaskRunStatusEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *ReportingTaskRunStatusEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *ReportingTaskRunStatusEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


