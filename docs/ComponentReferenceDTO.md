# ComponentReferenceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the component. | [optional] 
**VersionedComponentId** | Pointer to **string** | The ID of the corresponding component that is under version control | [optional] 
**ParentGroupId** | Pointer to **string** | The id of parent process group of this component if applicable. | [optional] 
**Position** | Pointer to [**PositionDTO**](PositionDTO.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the component. | [optional] 

## Methods

### NewComponentReferenceDTO

`func NewComponentReferenceDTO() *ComponentReferenceDTO`

NewComponentReferenceDTO instantiates a new ComponentReferenceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentReferenceDTOWithDefaults

`func NewComponentReferenceDTOWithDefaults() *ComponentReferenceDTO`

NewComponentReferenceDTOWithDefaults instantiates a new ComponentReferenceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComponentReferenceDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComponentReferenceDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComponentReferenceDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComponentReferenceDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersionedComponentId

`func (o *ComponentReferenceDTO) GetVersionedComponentId() string`

GetVersionedComponentId returns the VersionedComponentId field if non-nil, zero value otherwise.

### GetVersionedComponentIdOk

`func (o *ComponentReferenceDTO) GetVersionedComponentIdOk() (*string, bool)`

GetVersionedComponentIdOk returns a tuple with the VersionedComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedComponentId

`func (o *ComponentReferenceDTO) SetVersionedComponentId(v string)`

SetVersionedComponentId sets VersionedComponentId field to given value.

### HasVersionedComponentId

`func (o *ComponentReferenceDTO) HasVersionedComponentId() bool`

HasVersionedComponentId returns a boolean if a field has been set.

### GetParentGroupId

`func (o *ComponentReferenceDTO) GetParentGroupId() string`

GetParentGroupId returns the ParentGroupId field if non-nil, zero value otherwise.

### GetParentGroupIdOk

`func (o *ComponentReferenceDTO) GetParentGroupIdOk() (*string, bool)`

GetParentGroupIdOk returns a tuple with the ParentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroupId

`func (o *ComponentReferenceDTO) SetParentGroupId(v string)`

SetParentGroupId sets ParentGroupId field to given value.

### HasParentGroupId

`func (o *ComponentReferenceDTO) HasParentGroupId() bool`

HasParentGroupId returns a boolean if a field has been set.

### GetPosition

`func (o *ComponentReferenceDTO) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ComponentReferenceDTO) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ComponentReferenceDTO) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ComponentReferenceDTO) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetName

`func (o *ComponentReferenceDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentReferenceDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentReferenceDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComponentReferenceDTO) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


