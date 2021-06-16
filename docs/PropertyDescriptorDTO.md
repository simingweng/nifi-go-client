# PropertyDescriptorDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for the property. | [optional] 
**DisplayName** | Pointer to **string** | The human readable name for the property. | [optional] 
**Description** | Pointer to **string** | The description for the property. Used to relay additional details to a user or provide a mechanism of documenting intent. | [optional] 
**DefaultValue** | Pointer to **string** | The default value for the property. | [optional] 
**AllowableValues** | Pointer to [**[]AllowableValueEntity**](AllowableValueEntity.md) | Allowable values for the property. If empty then the allowed values are not constrained. | [optional] 
**Required** | Pointer to **bool** | Whether the property is required. | [optional] 
**Sensitive** | Pointer to **bool** | Whether the property is sensitive and protected whenever stored or represented. | [optional] 
**Dynamic** | Pointer to **bool** | Whether the property is dynamic (user-defined). | [optional] 
**SupportsEl** | Pointer to **bool** | Whether the property supports expression language. | [optional] 
**ExpressionLanguageScope** | Pointer to **string** | Scope of the Expression Language evaluation for the property. | [optional] 
**IdentifiesControllerService** | Pointer to **string** | If the property identifies a controller service this returns the fully qualified type. | [optional] 
**IdentifiesControllerServiceBundle** | Pointer to [**BundleDTO**](BundleDTO.md) |  | [optional] 
**Dependencies** | Pointer to [**[]PropertyDependencyDTO**](PropertyDependencyDTO.md) | A list of dependencies that must be met in order for this Property to be relevant. If any of these dependencies is not met, the property described by this Property Descriptor is not relevant. | [optional] 

## Methods

### NewPropertyDescriptorDTO

`func NewPropertyDescriptorDTO() *PropertyDescriptorDTO`

NewPropertyDescriptorDTO instantiates a new PropertyDescriptorDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyDescriptorDTOWithDefaults

`func NewPropertyDescriptorDTOWithDefaults() *PropertyDescriptorDTO`

NewPropertyDescriptorDTOWithDefaults instantiates a new PropertyDescriptorDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PropertyDescriptorDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PropertyDescriptorDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PropertyDescriptorDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PropertyDescriptorDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *PropertyDescriptorDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PropertyDescriptorDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PropertyDescriptorDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PropertyDescriptorDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *PropertyDescriptorDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PropertyDescriptorDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PropertyDescriptorDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PropertyDescriptorDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultValue

`func (o *PropertyDescriptorDTO) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *PropertyDescriptorDTO) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *PropertyDescriptorDTO) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *PropertyDescriptorDTO) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetAllowableValues

`func (o *PropertyDescriptorDTO) GetAllowableValues() []AllowableValueEntity`

GetAllowableValues returns the AllowableValues field if non-nil, zero value otherwise.

### GetAllowableValuesOk

`func (o *PropertyDescriptorDTO) GetAllowableValuesOk() (*[]AllowableValueEntity, bool)`

GetAllowableValuesOk returns a tuple with the AllowableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableValues

`func (o *PropertyDescriptorDTO) SetAllowableValues(v []AllowableValueEntity)`

SetAllowableValues sets AllowableValues field to given value.

### HasAllowableValues

`func (o *PropertyDescriptorDTO) HasAllowableValues() bool`

HasAllowableValues returns a boolean if a field has been set.

### GetRequired

`func (o *PropertyDescriptorDTO) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *PropertyDescriptorDTO) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *PropertyDescriptorDTO) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *PropertyDescriptorDTO) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetSensitive

`func (o *PropertyDescriptorDTO) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *PropertyDescriptorDTO) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *PropertyDescriptorDTO) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.

### HasSensitive

`func (o *PropertyDescriptorDTO) HasSensitive() bool`

HasSensitive returns a boolean if a field has been set.

### GetDynamic

`func (o *PropertyDescriptorDTO) GetDynamic() bool`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *PropertyDescriptorDTO) GetDynamicOk() (*bool, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *PropertyDescriptorDTO) SetDynamic(v bool)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *PropertyDescriptorDTO) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.

### GetSupportsEl

`func (o *PropertyDescriptorDTO) GetSupportsEl() bool`

GetSupportsEl returns the SupportsEl field if non-nil, zero value otherwise.

### GetSupportsElOk

`func (o *PropertyDescriptorDTO) GetSupportsElOk() (*bool, bool)`

GetSupportsElOk returns a tuple with the SupportsEl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsEl

`func (o *PropertyDescriptorDTO) SetSupportsEl(v bool)`

SetSupportsEl sets SupportsEl field to given value.

### HasSupportsEl

`func (o *PropertyDescriptorDTO) HasSupportsEl() bool`

HasSupportsEl returns a boolean if a field has been set.

### GetExpressionLanguageScope

`func (o *PropertyDescriptorDTO) GetExpressionLanguageScope() string`

GetExpressionLanguageScope returns the ExpressionLanguageScope field if non-nil, zero value otherwise.

### GetExpressionLanguageScopeOk

`func (o *PropertyDescriptorDTO) GetExpressionLanguageScopeOk() (*string, bool)`

GetExpressionLanguageScopeOk returns a tuple with the ExpressionLanguageScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressionLanguageScope

`func (o *PropertyDescriptorDTO) SetExpressionLanguageScope(v string)`

SetExpressionLanguageScope sets ExpressionLanguageScope field to given value.

### HasExpressionLanguageScope

`func (o *PropertyDescriptorDTO) HasExpressionLanguageScope() bool`

HasExpressionLanguageScope returns a boolean if a field has been set.

### GetIdentifiesControllerService

`func (o *PropertyDescriptorDTO) GetIdentifiesControllerService() string`

GetIdentifiesControllerService returns the IdentifiesControllerService field if non-nil, zero value otherwise.

### GetIdentifiesControllerServiceOk

`func (o *PropertyDescriptorDTO) GetIdentifiesControllerServiceOk() (*string, bool)`

GetIdentifiesControllerServiceOk returns a tuple with the IdentifiesControllerService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiesControllerService

`func (o *PropertyDescriptorDTO) SetIdentifiesControllerService(v string)`

SetIdentifiesControllerService sets IdentifiesControllerService field to given value.

### HasIdentifiesControllerService

`func (o *PropertyDescriptorDTO) HasIdentifiesControllerService() bool`

HasIdentifiesControllerService returns a boolean if a field has been set.

### GetIdentifiesControllerServiceBundle

`func (o *PropertyDescriptorDTO) GetIdentifiesControllerServiceBundle() BundleDTO`

GetIdentifiesControllerServiceBundle returns the IdentifiesControllerServiceBundle field if non-nil, zero value otherwise.

### GetIdentifiesControllerServiceBundleOk

`func (o *PropertyDescriptorDTO) GetIdentifiesControllerServiceBundleOk() (*BundleDTO, bool)`

GetIdentifiesControllerServiceBundleOk returns a tuple with the IdentifiesControllerServiceBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiesControllerServiceBundle

`func (o *PropertyDescriptorDTO) SetIdentifiesControllerServiceBundle(v BundleDTO)`

SetIdentifiesControllerServiceBundle sets IdentifiesControllerServiceBundle field to given value.

### HasIdentifiesControllerServiceBundle

`func (o *PropertyDescriptorDTO) HasIdentifiesControllerServiceBundle() bool`

HasIdentifiesControllerServiceBundle returns a boolean if a field has been set.

### GetDependencies

`func (o *PropertyDescriptorDTO) GetDependencies() []PropertyDependencyDTO`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *PropertyDescriptorDTO) GetDependenciesOk() (*[]PropertyDependencyDTO, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *PropertyDescriptorDTO) SetDependencies(v []PropertyDependencyDTO)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *PropertyDescriptorDTO) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


