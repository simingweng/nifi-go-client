# ControllerServiceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the component. | [optional] 
**VersionedComponentId** | Pointer to **string** | The ID of the corresponding component that is under version control | [optional] 
**ParentGroupId** | Pointer to **string** | The id of parent process group of this component if applicable. | [optional] 
**Position** | Pointer to [**PositionDTO**](PositionDTO.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the controller service. | [optional] 
**Type** | Pointer to **string** | The type of the controller service. | [optional] 
**Bundle** | Pointer to [**BundleDTO**](BundleDTO.md) |  | [optional] 
**ControllerServiceApis** | Pointer to [**[]ControllerServiceApiDTO**](ControllerServiceApiDTO.md) | Lists the APIs this Controller Service implements. | [optional] 
**Comments** | Pointer to **string** | The comments for the controller service. | [optional] 
**State** | Pointer to **string** | The state of the controller service. | [optional] 
**PersistsState** | Pointer to **bool** | Whether the controller service persists state. | [optional] 
**Restricted** | Pointer to **bool** | Whether the controller service requires elevated privileges. | [optional] 
**Deprecated** | Pointer to **bool** | Whether the ontroller service has been deprecated. | [optional] 
**MultipleVersionsAvailable** | Pointer to **bool** | Whether the controller service has multiple versions available. | [optional] 
**Properties** | Pointer to **map[string]string** | The properties of the controller service. | [optional] 
**Descriptors** | Pointer to [**map[string]PropertyDescriptorDTO**](PropertyDescriptorDTO.md) | The descriptors for the controller service properties. | [optional] 
**CustomUiUrl** | Pointer to **string** | The URL for the controller services custom configuration UI if applicable. | [optional] 
**AnnotationData** | Pointer to **string** | The annotation for the controller service. This is how the custom UI relays configuration to the controller service. | [optional] 
**ReferencingComponents** | Pointer to [**[]ControllerServiceReferencingComponentEntity**](ControllerServiceReferencingComponentEntity.md) | All components referencing this controller service. | [optional] 
**ValidationErrors** | Pointer to **[]string** | The validation errors from the controller service. These validation errors represent the problems with the controller service that must be resolved before it can be enabled. | [optional] 
**ValidationStatus** | Pointer to **string** | Indicates whether the ControllerService is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the ControllerService is valid) | [optional] [readonly] 
**ExtensionMissing** | Pointer to **bool** | Whether the underlying extension is missing. | [optional] 

## Methods

### NewControllerServiceDTO

`func NewControllerServiceDTO() *ControllerServiceDTO`

NewControllerServiceDTO instantiates a new ControllerServiceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerServiceDTOWithDefaults

`func NewControllerServiceDTOWithDefaults() *ControllerServiceDTO`

NewControllerServiceDTOWithDefaults instantiates a new ControllerServiceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ControllerServiceDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControllerServiceDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControllerServiceDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ControllerServiceDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersionedComponentId

`func (o *ControllerServiceDTO) GetVersionedComponentId() string`

GetVersionedComponentId returns the VersionedComponentId field if non-nil, zero value otherwise.

### GetVersionedComponentIdOk

`func (o *ControllerServiceDTO) GetVersionedComponentIdOk() (*string, bool)`

GetVersionedComponentIdOk returns a tuple with the VersionedComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedComponentId

`func (o *ControllerServiceDTO) SetVersionedComponentId(v string)`

SetVersionedComponentId sets VersionedComponentId field to given value.

### HasVersionedComponentId

`func (o *ControllerServiceDTO) HasVersionedComponentId() bool`

HasVersionedComponentId returns a boolean if a field has been set.

### GetParentGroupId

`func (o *ControllerServiceDTO) GetParentGroupId() string`

GetParentGroupId returns the ParentGroupId field if non-nil, zero value otherwise.

### GetParentGroupIdOk

`func (o *ControllerServiceDTO) GetParentGroupIdOk() (*string, bool)`

GetParentGroupIdOk returns a tuple with the ParentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroupId

`func (o *ControllerServiceDTO) SetParentGroupId(v string)`

SetParentGroupId sets ParentGroupId field to given value.

### HasParentGroupId

`func (o *ControllerServiceDTO) HasParentGroupId() bool`

HasParentGroupId returns a boolean if a field has been set.

### GetPosition

`func (o *ControllerServiceDTO) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ControllerServiceDTO) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ControllerServiceDTO) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ControllerServiceDTO) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetName

`func (o *ControllerServiceDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControllerServiceDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControllerServiceDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ControllerServiceDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ControllerServiceDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ControllerServiceDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ControllerServiceDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ControllerServiceDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBundle

`func (o *ControllerServiceDTO) GetBundle() BundleDTO`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *ControllerServiceDTO) GetBundleOk() (*BundleDTO, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *ControllerServiceDTO) SetBundle(v BundleDTO)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *ControllerServiceDTO) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### GetControllerServiceApis

`func (o *ControllerServiceDTO) GetControllerServiceApis() []ControllerServiceApiDTO`

GetControllerServiceApis returns the ControllerServiceApis field if non-nil, zero value otherwise.

### GetControllerServiceApisOk

`func (o *ControllerServiceDTO) GetControllerServiceApisOk() (*[]ControllerServiceApiDTO, bool)`

GetControllerServiceApisOk returns a tuple with the ControllerServiceApis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerServiceApis

`func (o *ControllerServiceDTO) SetControllerServiceApis(v []ControllerServiceApiDTO)`

SetControllerServiceApis sets ControllerServiceApis field to given value.

### HasControllerServiceApis

`func (o *ControllerServiceDTO) HasControllerServiceApis() bool`

HasControllerServiceApis returns a boolean if a field has been set.

### GetComments

`func (o *ControllerServiceDTO) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ControllerServiceDTO) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ControllerServiceDTO) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ControllerServiceDTO) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetState

`func (o *ControllerServiceDTO) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ControllerServiceDTO) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ControllerServiceDTO) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ControllerServiceDTO) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPersistsState

`func (o *ControllerServiceDTO) GetPersistsState() bool`

GetPersistsState returns the PersistsState field if non-nil, zero value otherwise.

### GetPersistsStateOk

`func (o *ControllerServiceDTO) GetPersistsStateOk() (*bool, bool)`

GetPersistsStateOk returns a tuple with the PersistsState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistsState

`func (o *ControllerServiceDTO) SetPersistsState(v bool)`

SetPersistsState sets PersistsState field to given value.

### HasPersistsState

`func (o *ControllerServiceDTO) HasPersistsState() bool`

HasPersistsState returns a boolean if a field has been set.

### GetRestricted

`func (o *ControllerServiceDTO) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *ControllerServiceDTO) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *ControllerServiceDTO) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *ControllerServiceDTO) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### GetDeprecated

`func (o *ControllerServiceDTO) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *ControllerServiceDTO) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *ControllerServiceDTO) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *ControllerServiceDTO) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetMultipleVersionsAvailable

`func (o *ControllerServiceDTO) GetMultipleVersionsAvailable() bool`

GetMultipleVersionsAvailable returns the MultipleVersionsAvailable field if non-nil, zero value otherwise.

### GetMultipleVersionsAvailableOk

`func (o *ControllerServiceDTO) GetMultipleVersionsAvailableOk() (*bool, bool)`

GetMultipleVersionsAvailableOk returns a tuple with the MultipleVersionsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVersionsAvailable

`func (o *ControllerServiceDTO) SetMultipleVersionsAvailable(v bool)`

SetMultipleVersionsAvailable sets MultipleVersionsAvailable field to given value.

### HasMultipleVersionsAvailable

`func (o *ControllerServiceDTO) HasMultipleVersionsAvailable() bool`

HasMultipleVersionsAvailable returns a boolean if a field has been set.

### GetProperties

`func (o *ControllerServiceDTO) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ControllerServiceDTO) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ControllerServiceDTO) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ControllerServiceDTO) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetDescriptors

`func (o *ControllerServiceDTO) GetDescriptors() map[string]PropertyDescriptorDTO`

GetDescriptors returns the Descriptors field if non-nil, zero value otherwise.

### GetDescriptorsOk

`func (o *ControllerServiceDTO) GetDescriptorsOk() (*map[string]PropertyDescriptorDTO, bool)`

GetDescriptorsOk returns a tuple with the Descriptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptors

`func (o *ControllerServiceDTO) SetDescriptors(v map[string]PropertyDescriptorDTO)`

SetDescriptors sets Descriptors field to given value.

### HasDescriptors

`func (o *ControllerServiceDTO) HasDescriptors() bool`

HasDescriptors returns a boolean if a field has been set.

### GetCustomUiUrl

`func (o *ControllerServiceDTO) GetCustomUiUrl() string`

GetCustomUiUrl returns the CustomUiUrl field if non-nil, zero value otherwise.

### GetCustomUiUrlOk

`func (o *ControllerServiceDTO) GetCustomUiUrlOk() (*string, bool)`

GetCustomUiUrlOk returns a tuple with the CustomUiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUiUrl

`func (o *ControllerServiceDTO) SetCustomUiUrl(v string)`

SetCustomUiUrl sets CustomUiUrl field to given value.

### HasCustomUiUrl

`func (o *ControllerServiceDTO) HasCustomUiUrl() bool`

HasCustomUiUrl returns a boolean if a field has been set.

### GetAnnotationData

`func (o *ControllerServiceDTO) GetAnnotationData() string`

GetAnnotationData returns the AnnotationData field if non-nil, zero value otherwise.

### GetAnnotationDataOk

`func (o *ControllerServiceDTO) GetAnnotationDataOk() (*string, bool)`

GetAnnotationDataOk returns a tuple with the AnnotationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationData

`func (o *ControllerServiceDTO) SetAnnotationData(v string)`

SetAnnotationData sets AnnotationData field to given value.

### HasAnnotationData

`func (o *ControllerServiceDTO) HasAnnotationData() bool`

HasAnnotationData returns a boolean if a field has been set.

### GetReferencingComponents

`func (o *ControllerServiceDTO) GetReferencingComponents() []ControllerServiceReferencingComponentEntity`

GetReferencingComponents returns the ReferencingComponents field if non-nil, zero value otherwise.

### GetReferencingComponentsOk

`func (o *ControllerServiceDTO) GetReferencingComponentsOk() (*[]ControllerServiceReferencingComponentEntity, bool)`

GetReferencingComponentsOk returns a tuple with the ReferencingComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencingComponents

`func (o *ControllerServiceDTO) SetReferencingComponents(v []ControllerServiceReferencingComponentEntity)`

SetReferencingComponents sets ReferencingComponents field to given value.

### HasReferencingComponents

`func (o *ControllerServiceDTO) HasReferencingComponents() bool`

HasReferencingComponents returns a boolean if a field has been set.

### GetValidationErrors

`func (o *ControllerServiceDTO) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *ControllerServiceDTO) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *ControllerServiceDTO) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *ControllerServiceDTO) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### GetValidationStatus

`func (o *ControllerServiceDTO) GetValidationStatus() string`

GetValidationStatus returns the ValidationStatus field if non-nil, zero value otherwise.

### GetValidationStatusOk

`func (o *ControllerServiceDTO) GetValidationStatusOk() (*string, bool)`

GetValidationStatusOk returns a tuple with the ValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStatus

`func (o *ControllerServiceDTO) SetValidationStatus(v string)`

SetValidationStatus sets ValidationStatus field to given value.

### HasValidationStatus

`func (o *ControllerServiceDTO) HasValidationStatus() bool`

HasValidationStatus returns a boolean if a field has been set.

### GetExtensionMissing

`func (o *ControllerServiceDTO) GetExtensionMissing() bool`

GetExtensionMissing returns the ExtensionMissing field if non-nil, zero value otherwise.

### GetExtensionMissingOk

`func (o *ControllerServiceDTO) GetExtensionMissingOk() (*bool, bool)`

GetExtensionMissingOk returns a tuple with the ExtensionMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionMissing

`func (o *ControllerServiceDTO) SetExtensionMissing(v bool)`

SetExtensionMissing sets ExtensionMissing field to given value.

### HasExtensionMissing

`func (o *ControllerServiceDTO) HasExtensionMissing() bool`

HasExtensionMissing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


