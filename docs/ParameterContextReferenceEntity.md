# ParameterContextReferenceEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the component. | [optional] 
**Permissions** | Pointer to [**PermissionsDTO**](PermissionsDTO.md) |  | [optional] 
**Component** | Pointer to [**ParameterContextReferenceDTO**](ParameterContextReferenceDTO.md) |  | [optional] 

## Methods

### NewParameterContextReferenceEntity

`func NewParameterContextReferenceEntity() *ParameterContextReferenceEntity`

NewParameterContextReferenceEntity instantiates a new ParameterContextReferenceEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterContextReferenceEntityWithDefaults

`func NewParameterContextReferenceEntityWithDefaults() *ParameterContextReferenceEntity`

NewParameterContextReferenceEntityWithDefaults instantiates a new ParameterContextReferenceEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ParameterContextReferenceEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParameterContextReferenceEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParameterContextReferenceEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ParameterContextReferenceEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPermissions

`func (o *ParameterContextReferenceEntity) GetPermissions() PermissionsDTO`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ParameterContextReferenceEntity) GetPermissionsOk() (*PermissionsDTO, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ParameterContextReferenceEntity) SetPermissions(v PermissionsDTO)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ParameterContextReferenceEntity) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetComponent

`func (o *ParameterContextReferenceEntity) GetComponent() ParameterContextReferenceDTO`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ParameterContextReferenceEntity) GetComponentOk() (*ParameterContextReferenceDTO, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ParameterContextReferenceEntity) SetComponent(v ParameterContextReferenceDTO)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ParameterContextReferenceEntity) HasComponent() bool`

HasComponent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


