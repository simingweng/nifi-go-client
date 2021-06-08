# VersionControlComponentMappingEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionControlComponentMapping** | Pointer to **map[string]string** | The mapping of Versioned Component Identifiers to instance ID&#39;s | [optional] 
**ProcessGroupRevision** | Pointer to [**RevisionDTO**](RevisionDTO.md) |  | [optional] 
**DisconnectedNodeAcknowledged** | Pointer to **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] 
**VersionControlInformation** | Pointer to [**VersionControlInformationDTO**](VersionControlInformationDTO.md) |  | [optional] 

## Methods

### NewVersionControlComponentMappingEntity

`func NewVersionControlComponentMappingEntity() *VersionControlComponentMappingEntity`

NewVersionControlComponentMappingEntity instantiates a new VersionControlComponentMappingEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionControlComponentMappingEntityWithDefaults

`func NewVersionControlComponentMappingEntityWithDefaults() *VersionControlComponentMappingEntity`

NewVersionControlComponentMappingEntityWithDefaults instantiates a new VersionControlComponentMappingEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionControlComponentMapping

`func (o *VersionControlComponentMappingEntity) GetVersionControlComponentMapping() map[string]string`

GetVersionControlComponentMapping returns the VersionControlComponentMapping field if non-nil, zero value otherwise.

### GetVersionControlComponentMappingOk

`func (o *VersionControlComponentMappingEntity) GetVersionControlComponentMappingOk() (*map[string]string, bool)`

GetVersionControlComponentMappingOk returns a tuple with the VersionControlComponentMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionControlComponentMapping

`func (o *VersionControlComponentMappingEntity) SetVersionControlComponentMapping(v map[string]string)`

SetVersionControlComponentMapping sets VersionControlComponentMapping field to given value.

### HasVersionControlComponentMapping

`func (o *VersionControlComponentMappingEntity) HasVersionControlComponentMapping() bool`

HasVersionControlComponentMapping returns a boolean if a field has been set.

### GetProcessGroupRevision

`func (o *VersionControlComponentMappingEntity) GetProcessGroupRevision() RevisionDTO`

GetProcessGroupRevision returns the ProcessGroupRevision field if non-nil, zero value otherwise.

### GetProcessGroupRevisionOk

`func (o *VersionControlComponentMappingEntity) GetProcessGroupRevisionOk() (*RevisionDTO, bool)`

GetProcessGroupRevisionOk returns a tuple with the ProcessGroupRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupRevision

`func (o *VersionControlComponentMappingEntity) SetProcessGroupRevision(v RevisionDTO)`

SetProcessGroupRevision sets ProcessGroupRevision field to given value.

### HasProcessGroupRevision

`func (o *VersionControlComponentMappingEntity) HasProcessGroupRevision() bool`

HasProcessGroupRevision returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *VersionControlComponentMappingEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *VersionControlComponentMappingEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *VersionControlComponentMappingEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *VersionControlComponentMappingEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.

### GetVersionControlInformation

`func (o *VersionControlComponentMappingEntity) GetVersionControlInformation() VersionControlInformationDTO`

GetVersionControlInformation returns the VersionControlInformation field if non-nil, zero value otherwise.

### GetVersionControlInformationOk

`func (o *VersionControlComponentMappingEntity) GetVersionControlInformationOk() (*VersionControlInformationDTO, bool)`

GetVersionControlInformationOk returns a tuple with the VersionControlInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionControlInformation

`func (o *VersionControlComponentMappingEntity) SetVersionControlInformation(v VersionControlInformationDTO)`

SetVersionControlInformation sets VersionControlInformation field to given value.

### HasVersionControlInformation

`func (o *VersionControlComponentMappingEntity) HasVersionControlInformation() bool`

HasVersionControlInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


