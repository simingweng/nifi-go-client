# AccessPolicySummaryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the component. | [optional] 
**VersionedComponentId** | Pointer to **string** | The ID of the corresponding component that is under version control | [optional] 
**ParentGroupId** | Pointer to **string** | The id of parent process group of this component if applicable. | [optional] 
**Position** | Pointer to [**PositionDTO**](PositionDTO.md) |  | [optional] 
**Resource** | Pointer to **string** | The resource for this access policy. | [optional] 
**Action** | Pointer to **string** | The action associated with this access policy. | [optional] 
**ComponentReference** | Pointer to [**ComponentReferenceEntity**](ComponentReferenceEntity.md) |  | [optional] 
**Configurable** | Pointer to **bool** | Whether this policy is configurable. | [optional] 

## Methods

### NewAccessPolicySummaryDTO

`func NewAccessPolicySummaryDTO() *AccessPolicySummaryDTO`

NewAccessPolicySummaryDTO instantiates a new AccessPolicySummaryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPolicySummaryDTOWithDefaults

`func NewAccessPolicySummaryDTOWithDefaults() *AccessPolicySummaryDTO`

NewAccessPolicySummaryDTOWithDefaults instantiates a new AccessPolicySummaryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessPolicySummaryDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessPolicySummaryDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessPolicySummaryDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccessPolicySummaryDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersionedComponentId

`func (o *AccessPolicySummaryDTO) GetVersionedComponentId() string`

GetVersionedComponentId returns the VersionedComponentId field if non-nil, zero value otherwise.

### GetVersionedComponentIdOk

`func (o *AccessPolicySummaryDTO) GetVersionedComponentIdOk() (*string, bool)`

GetVersionedComponentIdOk returns a tuple with the VersionedComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedComponentId

`func (o *AccessPolicySummaryDTO) SetVersionedComponentId(v string)`

SetVersionedComponentId sets VersionedComponentId field to given value.

### HasVersionedComponentId

`func (o *AccessPolicySummaryDTO) HasVersionedComponentId() bool`

HasVersionedComponentId returns a boolean if a field has been set.

### GetParentGroupId

`func (o *AccessPolicySummaryDTO) GetParentGroupId() string`

GetParentGroupId returns the ParentGroupId field if non-nil, zero value otherwise.

### GetParentGroupIdOk

`func (o *AccessPolicySummaryDTO) GetParentGroupIdOk() (*string, bool)`

GetParentGroupIdOk returns a tuple with the ParentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroupId

`func (o *AccessPolicySummaryDTO) SetParentGroupId(v string)`

SetParentGroupId sets ParentGroupId field to given value.

### HasParentGroupId

`func (o *AccessPolicySummaryDTO) HasParentGroupId() bool`

HasParentGroupId returns a boolean if a field has been set.

### GetPosition

`func (o *AccessPolicySummaryDTO) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *AccessPolicySummaryDTO) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *AccessPolicySummaryDTO) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *AccessPolicySummaryDTO) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetResource

`func (o *AccessPolicySummaryDTO) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AccessPolicySummaryDTO) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AccessPolicySummaryDTO) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *AccessPolicySummaryDTO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetAction

`func (o *AccessPolicySummaryDTO) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AccessPolicySummaryDTO) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AccessPolicySummaryDTO) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AccessPolicySummaryDTO) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetComponentReference

`func (o *AccessPolicySummaryDTO) GetComponentReference() ComponentReferenceEntity`

GetComponentReference returns the ComponentReference field if non-nil, zero value otherwise.

### GetComponentReferenceOk

`func (o *AccessPolicySummaryDTO) GetComponentReferenceOk() (*ComponentReferenceEntity, bool)`

GetComponentReferenceOk returns a tuple with the ComponentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentReference

`func (o *AccessPolicySummaryDTO) SetComponentReference(v ComponentReferenceEntity)`

SetComponentReference sets ComponentReference field to given value.

### HasComponentReference

`func (o *AccessPolicySummaryDTO) HasComponentReference() bool`

HasComponentReference returns a boolean if a field has been set.

### GetConfigurable

`func (o *AccessPolicySummaryDTO) GetConfigurable() bool`

GetConfigurable returns the Configurable field if non-nil, zero value otherwise.

### GetConfigurableOk

`func (o *AccessPolicySummaryDTO) GetConfigurableOk() (*bool, bool)`

GetConfigurableOk returns a tuple with the Configurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurable

`func (o *AccessPolicySummaryDTO) SetConfigurable(v bool)`

SetConfigurable sets Configurable field to given value.

### HasConfigurable

`func (o *AccessPolicySummaryDTO) HasConfigurable() bool`

HasConfigurable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


