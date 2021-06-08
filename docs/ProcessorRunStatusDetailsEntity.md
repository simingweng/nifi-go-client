# ProcessorRunStatusDetailsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | Pointer to [**RevisionDTO**](RevisionDTO.md) |  | [optional] 
**Permissions** | Pointer to [**PermissionsDTO**](PermissionsDTO.md) |  | [optional] 
**RunStatusDetails** | Pointer to [**ProcessorRunStatusDetailsDTO**](ProcessorRunStatusDetailsDTO.md) |  | [optional] 

## Methods

### NewProcessorRunStatusDetailsEntity

`func NewProcessorRunStatusDetailsEntity() *ProcessorRunStatusDetailsEntity`

NewProcessorRunStatusDetailsEntity instantiates a new ProcessorRunStatusDetailsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorRunStatusDetailsEntityWithDefaults

`func NewProcessorRunStatusDetailsEntityWithDefaults() *ProcessorRunStatusDetailsEntity`

NewProcessorRunStatusDetailsEntityWithDefaults instantiates a new ProcessorRunStatusDetailsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *ProcessorRunStatusDetailsEntity) GetRevision() RevisionDTO`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ProcessorRunStatusDetailsEntity) GetRevisionOk() (*RevisionDTO, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ProcessorRunStatusDetailsEntity) SetRevision(v RevisionDTO)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ProcessorRunStatusDetailsEntity) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetPermissions

`func (o *ProcessorRunStatusDetailsEntity) GetPermissions() PermissionsDTO`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ProcessorRunStatusDetailsEntity) GetPermissionsOk() (*PermissionsDTO, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ProcessorRunStatusDetailsEntity) SetPermissions(v PermissionsDTO)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ProcessorRunStatusDetailsEntity) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetRunStatusDetails

`func (o *ProcessorRunStatusDetailsEntity) GetRunStatusDetails() ProcessorRunStatusDetailsDTO`

GetRunStatusDetails returns the RunStatusDetails field if non-nil, zero value otherwise.

### GetRunStatusDetailsOk

`func (o *ProcessorRunStatusDetailsEntity) GetRunStatusDetailsOk() (*ProcessorRunStatusDetailsDTO, bool)`

GetRunStatusDetailsOk returns a tuple with the RunStatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStatusDetails

`func (o *ProcessorRunStatusDetailsEntity) SetRunStatusDetails(v ProcessorRunStatusDetailsDTO)`

SetRunStatusDetails sets RunStatusDetails field to given value.

### HasRunStatusDetails

`func (o *ProcessorRunStatusDetailsEntity) HasRunStatusDetails() bool`

HasRunStatusDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


