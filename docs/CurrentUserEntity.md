# CurrentUserEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | Pointer to **string** | The user identity being serialized. | [optional] 
**Anonymous** | Pointer to **bool** | Whether the current user is anonymous. | [optional] 
**ProvenancePermissions** | Pointer to [**PermissionsDTO**](PermissionsDTO.md) |  | [optional] 
**CountersPermissions** | Pointer to [**PermissionsDTO**](PermissionsDTO.md) |  | [optional] 
**TenantsPermissions** | Pointer to [**PermissionsDTO**](PermissionsDTO.md) |  | [optional] 
**ControllerPermissions** | Pointer to [**PermissionsDTO**](PermissionsDTO.md) |  | [optional] 
**PoliciesPermissions** | Pointer to [**PermissionsDTO**](PermissionsDTO.md) |  | [optional] 
**SystemPermissions** | Pointer to [**PermissionsDTO**](PermissionsDTO.md) |  | [optional] 
**ParameterContextPermissions** | Pointer to [**PermissionsDTO**](PermissionsDTO.md) |  | [optional] 
**RestrictedComponentsPermissions** | Pointer to [**PermissionsDTO**](PermissionsDTO.md) |  | [optional] 
**ComponentRestrictionPermissions** | Pointer to [**[]ComponentRestrictionPermissionDTO**](ComponentRestrictionPermissionDTO.md) | Permissions for specific component restrictions. | [optional] 
**CanVersionFlows** | Pointer to **bool** | Whether the current user can version flows. | [optional] 

## Methods

### NewCurrentUserEntity

`func NewCurrentUserEntity() *CurrentUserEntity`

NewCurrentUserEntity instantiates a new CurrentUserEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentUserEntityWithDefaults

`func NewCurrentUserEntityWithDefaults() *CurrentUserEntity`

NewCurrentUserEntityWithDefaults instantiates a new CurrentUserEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *CurrentUserEntity) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CurrentUserEntity) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CurrentUserEntity) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *CurrentUserEntity) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetAnonymous

`func (o *CurrentUserEntity) GetAnonymous() bool`

GetAnonymous returns the Anonymous field if non-nil, zero value otherwise.

### GetAnonymousOk

`func (o *CurrentUserEntity) GetAnonymousOk() (*bool, bool)`

GetAnonymousOk returns a tuple with the Anonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymous

`func (o *CurrentUserEntity) SetAnonymous(v bool)`

SetAnonymous sets Anonymous field to given value.

### HasAnonymous

`func (o *CurrentUserEntity) HasAnonymous() bool`

HasAnonymous returns a boolean if a field has been set.

### GetProvenancePermissions

`func (o *CurrentUserEntity) GetProvenancePermissions() PermissionsDTO`

GetProvenancePermissions returns the ProvenancePermissions field if non-nil, zero value otherwise.

### GetProvenancePermissionsOk

`func (o *CurrentUserEntity) GetProvenancePermissionsOk() (*PermissionsDTO, bool)`

GetProvenancePermissionsOk returns a tuple with the ProvenancePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenancePermissions

`func (o *CurrentUserEntity) SetProvenancePermissions(v PermissionsDTO)`

SetProvenancePermissions sets ProvenancePermissions field to given value.

### HasProvenancePermissions

`func (o *CurrentUserEntity) HasProvenancePermissions() bool`

HasProvenancePermissions returns a boolean if a field has been set.

### GetCountersPermissions

`func (o *CurrentUserEntity) GetCountersPermissions() PermissionsDTO`

GetCountersPermissions returns the CountersPermissions field if non-nil, zero value otherwise.

### GetCountersPermissionsOk

`func (o *CurrentUserEntity) GetCountersPermissionsOk() (*PermissionsDTO, bool)`

GetCountersPermissionsOk returns a tuple with the CountersPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountersPermissions

`func (o *CurrentUserEntity) SetCountersPermissions(v PermissionsDTO)`

SetCountersPermissions sets CountersPermissions field to given value.

### HasCountersPermissions

`func (o *CurrentUserEntity) HasCountersPermissions() bool`

HasCountersPermissions returns a boolean if a field has been set.

### GetTenantsPermissions

`func (o *CurrentUserEntity) GetTenantsPermissions() PermissionsDTO`

GetTenantsPermissions returns the TenantsPermissions field if non-nil, zero value otherwise.

### GetTenantsPermissionsOk

`func (o *CurrentUserEntity) GetTenantsPermissionsOk() (*PermissionsDTO, bool)`

GetTenantsPermissionsOk returns a tuple with the TenantsPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantsPermissions

`func (o *CurrentUserEntity) SetTenantsPermissions(v PermissionsDTO)`

SetTenantsPermissions sets TenantsPermissions field to given value.

### HasTenantsPermissions

`func (o *CurrentUserEntity) HasTenantsPermissions() bool`

HasTenantsPermissions returns a boolean if a field has been set.

### GetControllerPermissions

`func (o *CurrentUserEntity) GetControllerPermissions() PermissionsDTO`

GetControllerPermissions returns the ControllerPermissions field if non-nil, zero value otherwise.

### GetControllerPermissionsOk

`func (o *CurrentUserEntity) GetControllerPermissionsOk() (*PermissionsDTO, bool)`

GetControllerPermissionsOk returns a tuple with the ControllerPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerPermissions

`func (o *CurrentUserEntity) SetControllerPermissions(v PermissionsDTO)`

SetControllerPermissions sets ControllerPermissions field to given value.

### HasControllerPermissions

`func (o *CurrentUserEntity) HasControllerPermissions() bool`

HasControllerPermissions returns a boolean if a field has been set.

### GetPoliciesPermissions

`func (o *CurrentUserEntity) GetPoliciesPermissions() PermissionsDTO`

GetPoliciesPermissions returns the PoliciesPermissions field if non-nil, zero value otherwise.

### GetPoliciesPermissionsOk

`func (o *CurrentUserEntity) GetPoliciesPermissionsOk() (*PermissionsDTO, bool)`

GetPoliciesPermissionsOk returns a tuple with the PoliciesPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoliciesPermissions

`func (o *CurrentUserEntity) SetPoliciesPermissions(v PermissionsDTO)`

SetPoliciesPermissions sets PoliciesPermissions field to given value.

### HasPoliciesPermissions

`func (o *CurrentUserEntity) HasPoliciesPermissions() bool`

HasPoliciesPermissions returns a boolean if a field has been set.

### GetSystemPermissions

`func (o *CurrentUserEntity) GetSystemPermissions() PermissionsDTO`

GetSystemPermissions returns the SystemPermissions field if non-nil, zero value otherwise.

### GetSystemPermissionsOk

`func (o *CurrentUserEntity) GetSystemPermissionsOk() (*PermissionsDTO, bool)`

GetSystemPermissionsOk returns a tuple with the SystemPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPermissions

`func (o *CurrentUserEntity) SetSystemPermissions(v PermissionsDTO)`

SetSystemPermissions sets SystemPermissions field to given value.

### HasSystemPermissions

`func (o *CurrentUserEntity) HasSystemPermissions() bool`

HasSystemPermissions returns a boolean if a field has been set.

### GetParameterContextPermissions

`func (o *CurrentUserEntity) GetParameterContextPermissions() PermissionsDTO`

GetParameterContextPermissions returns the ParameterContextPermissions field if non-nil, zero value otherwise.

### GetParameterContextPermissionsOk

`func (o *CurrentUserEntity) GetParameterContextPermissionsOk() (*PermissionsDTO, bool)`

GetParameterContextPermissionsOk returns a tuple with the ParameterContextPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterContextPermissions

`func (o *CurrentUserEntity) SetParameterContextPermissions(v PermissionsDTO)`

SetParameterContextPermissions sets ParameterContextPermissions field to given value.

### HasParameterContextPermissions

`func (o *CurrentUserEntity) HasParameterContextPermissions() bool`

HasParameterContextPermissions returns a boolean if a field has been set.

### GetRestrictedComponentsPermissions

`func (o *CurrentUserEntity) GetRestrictedComponentsPermissions() PermissionsDTO`

GetRestrictedComponentsPermissions returns the RestrictedComponentsPermissions field if non-nil, zero value otherwise.

### GetRestrictedComponentsPermissionsOk

`func (o *CurrentUserEntity) GetRestrictedComponentsPermissionsOk() (*PermissionsDTO, bool)`

GetRestrictedComponentsPermissionsOk returns a tuple with the RestrictedComponentsPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedComponentsPermissions

`func (o *CurrentUserEntity) SetRestrictedComponentsPermissions(v PermissionsDTO)`

SetRestrictedComponentsPermissions sets RestrictedComponentsPermissions field to given value.

### HasRestrictedComponentsPermissions

`func (o *CurrentUserEntity) HasRestrictedComponentsPermissions() bool`

HasRestrictedComponentsPermissions returns a boolean if a field has been set.

### GetComponentRestrictionPermissions

`func (o *CurrentUserEntity) GetComponentRestrictionPermissions() []ComponentRestrictionPermissionDTO`

GetComponentRestrictionPermissions returns the ComponentRestrictionPermissions field if non-nil, zero value otherwise.

### GetComponentRestrictionPermissionsOk

`func (o *CurrentUserEntity) GetComponentRestrictionPermissionsOk() (*[]ComponentRestrictionPermissionDTO, bool)`

GetComponentRestrictionPermissionsOk returns a tuple with the ComponentRestrictionPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentRestrictionPermissions

`func (o *CurrentUserEntity) SetComponentRestrictionPermissions(v []ComponentRestrictionPermissionDTO)`

SetComponentRestrictionPermissions sets ComponentRestrictionPermissions field to given value.

### HasComponentRestrictionPermissions

`func (o *CurrentUserEntity) HasComponentRestrictionPermissions() bool`

HasComponentRestrictionPermissions returns a boolean if a field has been set.

### GetCanVersionFlows

`func (o *CurrentUserEntity) GetCanVersionFlows() bool`

GetCanVersionFlows returns the CanVersionFlows field if non-nil, zero value otherwise.

### GetCanVersionFlowsOk

`func (o *CurrentUserEntity) GetCanVersionFlowsOk() (*bool, bool)`

GetCanVersionFlowsOk returns a tuple with the CanVersionFlows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanVersionFlows

`func (o *CurrentUserEntity) SetCanVersionFlows(v bool)`

SetCanVersionFlows sets CanVersionFlows field to given value.

### HasCanVersionFlows

`func (o *CurrentUserEntity) HasCanVersionFlows() bool`

HasCanVersionFlows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


