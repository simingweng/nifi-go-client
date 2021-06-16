# FlowBreadcrumbEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of this ancestor ProcessGroup. | [optional] 
**Permissions** | Pointer to [**PermissionsDTO**](PermissionsDTO.md) |  | [optional] 
**VersionedFlowState** | Pointer to **string** | The current state of the Process Group, as it relates to the Versioned Flow | [optional] [readonly] 
**Breadcrumb** | Pointer to [**FlowBreadcrumbDTO**](FlowBreadcrumbDTO.md) |  | [optional] 
**ParentBreadcrumb** | Pointer to [**FlowBreadcrumbEntity**](FlowBreadcrumbEntity.md) |  | [optional] 

## Methods

### NewFlowBreadcrumbEntity

`func NewFlowBreadcrumbEntity() *FlowBreadcrumbEntity`

NewFlowBreadcrumbEntity instantiates a new FlowBreadcrumbEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowBreadcrumbEntityWithDefaults

`func NewFlowBreadcrumbEntityWithDefaults() *FlowBreadcrumbEntity`

NewFlowBreadcrumbEntityWithDefaults instantiates a new FlowBreadcrumbEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FlowBreadcrumbEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowBreadcrumbEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowBreadcrumbEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FlowBreadcrumbEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPermissions

`func (o *FlowBreadcrumbEntity) GetPermissions() PermissionsDTO`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *FlowBreadcrumbEntity) GetPermissionsOk() (*PermissionsDTO, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *FlowBreadcrumbEntity) SetPermissions(v PermissionsDTO)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *FlowBreadcrumbEntity) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetVersionedFlowState

`func (o *FlowBreadcrumbEntity) GetVersionedFlowState() string`

GetVersionedFlowState returns the VersionedFlowState field if non-nil, zero value otherwise.

### GetVersionedFlowStateOk

`func (o *FlowBreadcrumbEntity) GetVersionedFlowStateOk() (*string, bool)`

GetVersionedFlowStateOk returns a tuple with the VersionedFlowState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedFlowState

`func (o *FlowBreadcrumbEntity) SetVersionedFlowState(v string)`

SetVersionedFlowState sets VersionedFlowState field to given value.

### HasVersionedFlowState

`func (o *FlowBreadcrumbEntity) HasVersionedFlowState() bool`

HasVersionedFlowState returns a boolean if a field has been set.

### GetBreadcrumb

`func (o *FlowBreadcrumbEntity) GetBreadcrumb() FlowBreadcrumbDTO`

GetBreadcrumb returns the Breadcrumb field if non-nil, zero value otherwise.

### GetBreadcrumbOk

`func (o *FlowBreadcrumbEntity) GetBreadcrumbOk() (*FlowBreadcrumbDTO, bool)`

GetBreadcrumbOk returns a tuple with the Breadcrumb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreadcrumb

`func (o *FlowBreadcrumbEntity) SetBreadcrumb(v FlowBreadcrumbDTO)`

SetBreadcrumb sets Breadcrumb field to given value.

### HasBreadcrumb

`func (o *FlowBreadcrumbEntity) HasBreadcrumb() bool`

HasBreadcrumb returns a boolean if a field has been set.

### GetParentBreadcrumb

`func (o *FlowBreadcrumbEntity) GetParentBreadcrumb() FlowBreadcrumbEntity`

GetParentBreadcrumb returns the ParentBreadcrumb field if non-nil, zero value otherwise.

### GetParentBreadcrumbOk

`func (o *FlowBreadcrumbEntity) GetParentBreadcrumbOk() (*FlowBreadcrumbEntity, bool)`

GetParentBreadcrumbOk returns a tuple with the ParentBreadcrumb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBreadcrumb

`func (o *FlowBreadcrumbEntity) SetParentBreadcrumb(v FlowBreadcrumbEntity)`

SetParentBreadcrumb sets ParentBreadcrumb field to given value.

### HasParentBreadcrumb

`func (o *FlowBreadcrumbEntity) HasParentBreadcrumb() bool`

HasParentBreadcrumb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


