# ProcessGroupFlowEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to [**PermissionsDTO**](PermissionsDTO.md) |  | [optional] 
**ProcessGroupFlow** | Pointer to [**ProcessGroupFlowDTO**](ProcessGroupFlowDTO.md) |  | [optional] 

## Methods

### NewProcessGroupFlowEntity

`func NewProcessGroupFlowEntity() *ProcessGroupFlowEntity`

NewProcessGroupFlowEntity instantiates a new ProcessGroupFlowEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessGroupFlowEntityWithDefaults

`func NewProcessGroupFlowEntityWithDefaults() *ProcessGroupFlowEntity`

NewProcessGroupFlowEntityWithDefaults instantiates a new ProcessGroupFlowEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *ProcessGroupFlowEntity) GetPermissions() PermissionsDTO`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ProcessGroupFlowEntity) GetPermissionsOk() (*PermissionsDTO, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ProcessGroupFlowEntity) SetPermissions(v PermissionsDTO)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ProcessGroupFlowEntity) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetProcessGroupFlow

`func (o *ProcessGroupFlowEntity) GetProcessGroupFlow() ProcessGroupFlowDTO`

GetProcessGroupFlow returns the ProcessGroupFlow field if non-nil, zero value otherwise.

### GetProcessGroupFlowOk

`func (o *ProcessGroupFlowEntity) GetProcessGroupFlowOk() (*ProcessGroupFlowDTO, bool)`

GetProcessGroupFlowOk returns a tuple with the ProcessGroupFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupFlow

`func (o *ProcessGroupFlowEntity) SetProcessGroupFlow(v ProcessGroupFlowDTO)`

SetProcessGroupFlow sets ProcessGroupFlow field to given value.

### HasProcessGroupFlow

`func (o *ProcessGroupFlowEntity) HasProcessGroupFlow() bool`

HasProcessGroupFlow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


