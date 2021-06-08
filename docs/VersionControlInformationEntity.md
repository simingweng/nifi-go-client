# VersionControlInformationEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessGroupRevision** | Pointer to [**RevisionDTO**](RevisionDTO.md) |  | [optional] 
**DisconnectedNodeAcknowledged** | Pointer to **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] 
**VersionControlInformation** | Pointer to [**VersionControlInformationDTO**](VersionControlInformationDTO.md) |  | [optional] 

## Methods

### NewVersionControlInformationEntity

`func NewVersionControlInformationEntity() *VersionControlInformationEntity`

NewVersionControlInformationEntity instantiates a new VersionControlInformationEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionControlInformationEntityWithDefaults

`func NewVersionControlInformationEntityWithDefaults() *VersionControlInformationEntity`

NewVersionControlInformationEntityWithDefaults instantiates a new VersionControlInformationEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessGroupRevision

`func (o *VersionControlInformationEntity) GetProcessGroupRevision() RevisionDTO`

GetProcessGroupRevision returns the ProcessGroupRevision field if non-nil, zero value otherwise.

### GetProcessGroupRevisionOk

`func (o *VersionControlInformationEntity) GetProcessGroupRevisionOk() (*RevisionDTO, bool)`

GetProcessGroupRevisionOk returns a tuple with the ProcessGroupRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupRevision

`func (o *VersionControlInformationEntity) SetProcessGroupRevision(v RevisionDTO)`

SetProcessGroupRevision sets ProcessGroupRevision field to given value.

### HasProcessGroupRevision

`func (o *VersionControlInformationEntity) HasProcessGroupRevision() bool`

HasProcessGroupRevision returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *VersionControlInformationEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *VersionControlInformationEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *VersionControlInformationEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *VersionControlInformationEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.

### GetVersionControlInformation

`func (o *VersionControlInformationEntity) GetVersionControlInformation() VersionControlInformationDTO`

GetVersionControlInformation returns the VersionControlInformation field if non-nil, zero value otherwise.

### GetVersionControlInformationOk

`func (o *VersionControlInformationEntity) GetVersionControlInformationOk() (*VersionControlInformationDTO, bool)`

GetVersionControlInformationOk returns a tuple with the VersionControlInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionControlInformation

`func (o *VersionControlInformationEntity) SetVersionControlInformation(v VersionControlInformationDTO)`

SetVersionControlInformation sets VersionControlInformation field to given value.

### HasVersionControlInformation

`func (o *VersionControlInformationEntity) HasVersionControlInformation() bool`

HasVersionControlInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


