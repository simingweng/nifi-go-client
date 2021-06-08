# ExplicitRestrictionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequiredPermission** | Pointer to [**RequiredPermissionDTO**](RequiredPermissionDTO.md) |  | [optional] 
**Explanation** | Pointer to **string** | The description of why the usage of this component is restricted for this required permission. | [optional] 

## Methods

### NewExplicitRestrictionDTO

`func NewExplicitRestrictionDTO() *ExplicitRestrictionDTO`

NewExplicitRestrictionDTO instantiates a new ExplicitRestrictionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExplicitRestrictionDTOWithDefaults

`func NewExplicitRestrictionDTOWithDefaults() *ExplicitRestrictionDTO`

NewExplicitRestrictionDTOWithDefaults instantiates a new ExplicitRestrictionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequiredPermission

`func (o *ExplicitRestrictionDTO) GetRequiredPermission() RequiredPermissionDTO`

GetRequiredPermission returns the RequiredPermission field if non-nil, zero value otherwise.

### GetRequiredPermissionOk

`func (o *ExplicitRestrictionDTO) GetRequiredPermissionOk() (*RequiredPermissionDTO, bool)`

GetRequiredPermissionOk returns a tuple with the RequiredPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredPermission

`func (o *ExplicitRestrictionDTO) SetRequiredPermission(v RequiredPermissionDTO)`

SetRequiredPermission sets RequiredPermission field to given value.

### HasRequiredPermission

`func (o *ExplicitRestrictionDTO) HasRequiredPermission() bool`

HasRequiredPermission returns a boolean if a field has been set.

### GetExplanation

`func (o *ExplicitRestrictionDTO) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *ExplicitRestrictionDTO) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *ExplicitRestrictionDTO) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.

### HasExplanation

`func (o *ExplicitRestrictionDTO) HasExplanation() bool`

HasExplanation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


