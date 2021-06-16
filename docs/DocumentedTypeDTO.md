# DocumentedTypeDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The fully qualified name of the type. | [optional] 
**Bundle** | Pointer to [**BundleDTO**](BundleDTO.md) |  | [optional] 
**ControllerServiceApis** | Pointer to [**[]ControllerServiceApiDTO**](ControllerServiceApiDTO.md) | If this type represents a ControllerService, this lists the APIs it implements. | [optional] 
**Description** | Pointer to **string** | The description of the type. | [optional] 
**Restricted** | Pointer to **bool** | Whether this type is restricted. | [optional] 
**UsageRestriction** | Pointer to **string** | The optional description of why the usage of this component is restricted. | [optional] 
**ExplicitRestrictions** | Pointer to [**[]ExplicitRestrictionDTO**](ExplicitRestrictionDTO.md) | An optional collection of explicit restrictions. If specified, these explicit restrictions will be enfored. | [optional] 
**DeprecationReason** | Pointer to **string** | The description of why the usage of this component is restricted. | [optional] 
**Tags** | Pointer to **[]string** | The tags associated with this type. | [optional] 

## Methods

### NewDocumentedTypeDTO

`func NewDocumentedTypeDTO() *DocumentedTypeDTO`

NewDocumentedTypeDTO instantiates a new DocumentedTypeDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentedTypeDTOWithDefaults

`func NewDocumentedTypeDTOWithDefaults() *DocumentedTypeDTO`

NewDocumentedTypeDTOWithDefaults instantiates a new DocumentedTypeDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DocumentedTypeDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DocumentedTypeDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DocumentedTypeDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DocumentedTypeDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBundle

`func (o *DocumentedTypeDTO) GetBundle() BundleDTO`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *DocumentedTypeDTO) GetBundleOk() (*BundleDTO, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *DocumentedTypeDTO) SetBundle(v BundleDTO)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *DocumentedTypeDTO) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### GetControllerServiceApis

`func (o *DocumentedTypeDTO) GetControllerServiceApis() []ControllerServiceApiDTO`

GetControllerServiceApis returns the ControllerServiceApis field if non-nil, zero value otherwise.

### GetControllerServiceApisOk

`func (o *DocumentedTypeDTO) GetControllerServiceApisOk() (*[]ControllerServiceApiDTO, bool)`

GetControllerServiceApisOk returns a tuple with the ControllerServiceApis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerServiceApis

`func (o *DocumentedTypeDTO) SetControllerServiceApis(v []ControllerServiceApiDTO)`

SetControllerServiceApis sets ControllerServiceApis field to given value.

### HasControllerServiceApis

`func (o *DocumentedTypeDTO) HasControllerServiceApis() bool`

HasControllerServiceApis returns a boolean if a field has been set.

### GetDescription

`func (o *DocumentedTypeDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DocumentedTypeDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DocumentedTypeDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DocumentedTypeDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRestricted

`func (o *DocumentedTypeDTO) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *DocumentedTypeDTO) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *DocumentedTypeDTO) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *DocumentedTypeDTO) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### GetUsageRestriction

`func (o *DocumentedTypeDTO) GetUsageRestriction() string`

GetUsageRestriction returns the UsageRestriction field if non-nil, zero value otherwise.

### GetUsageRestrictionOk

`func (o *DocumentedTypeDTO) GetUsageRestrictionOk() (*string, bool)`

GetUsageRestrictionOk returns a tuple with the UsageRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageRestriction

`func (o *DocumentedTypeDTO) SetUsageRestriction(v string)`

SetUsageRestriction sets UsageRestriction field to given value.

### HasUsageRestriction

`func (o *DocumentedTypeDTO) HasUsageRestriction() bool`

HasUsageRestriction returns a boolean if a field has been set.

### GetExplicitRestrictions

`func (o *DocumentedTypeDTO) GetExplicitRestrictions() []ExplicitRestrictionDTO`

GetExplicitRestrictions returns the ExplicitRestrictions field if non-nil, zero value otherwise.

### GetExplicitRestrictionsOk

`func (o *DocumentedTypeDTO) GetExplicitRestrictionsOk() (*[]ExplicitRestrictionDTO, bool)`

GetExplicitRestrictionsOk returns a tuple with the ExplicitRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitRestrictions

`func (o *DocumentedTypeDTO) SetExplicitRestrictions(v []ExplicitRestrictionDTO)`

SetExplicitRestrictions sets ExplicitRestrictions field to given value.

### HasExplicitRestrictions

`func (o *DocumentedTypeDTO) HasExplicitRestrictions() bool`

HasExplicitRestrictions returns a boolean if a field has been set.

### GetDeprecationReason

`func (o *DocumentedTypeDTO) GetDeprecationReason() string`

GetDeprecationReason returns the DeprecationReason field if non-nil, zero value otherwise.

### GetDeprecationReasonOk

`func (o *DocumentedTypeDTO) GetDeprecationReasonOk() (*string, bool)`

GetDeprecationReasonOk returns a tuple with the DeprecationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecationReason

`func (o *DocumentedTypeDTO) SetDeprecationReason(v string)`

SetDeprecationReason sets DeprecationReason field to given value.

### HasDeprecationReason

`func (o *DocumentedTypeDTO) HasDeprecationReason() bool`

HasDeprecationReason returns a boolean if a field has been set.

### GetTags

`func (o *DocumentedTypeDTO) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DocumentedTypeDTO) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DocumentedTypeDTO) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DocumentedTypeDTO) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


