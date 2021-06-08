# ProcessGroupsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessGroups** | Pointer to [**[]ProcessGroupEntity**](ProcessGroupEntity.md) |  | [optional] 

## Methods

### NewProcessGroupsEntity

`func NewProcessGroupsEntity() *ProcessGroupsEntity`

NewProcessGroupsEntity instantiates a new ProcessGroupsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessGroupsEntityWithDefaults

`func NewProcessGroupsEntityWithDefaults() *ProcessGroupsEntity`

NewProcessGroupsEntityWithDefaults instantiates a new ProcessGroupsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessGroups

`func (o *ProcessGroupsEntity) GetProcessGroups() []ProcessGroupEntity`

GetProcessGroups returns the ProcessGroups field if non-nil, zero value otherwise.

### GetProcessGroupsOk

`func (o *ProcessGroupsEntity) GetProcessGroupsOk() (*[]ProcessGroupEntity, bool)`

GetProcessGroupsOk returns a tuple with the ProcessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroups

`func (o *ProcessGroupsEntity) SetProcessGroups(v []ProcessGroupEntity)`

SetProcessGroups sets ProcessGroups field to given value.

### HasProcessGroups

`func (o *ProcessGroupsEntity) HasProcessGroups() bool`

HasProcessGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


