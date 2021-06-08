# ReportingTaskDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the component. | [optional] 
**VersionedComponentId** | Pointer to **string** | The ID of the corresponding component that is under version control | [optional] 
**ParentGroupId** | Pointer to **string** | The id of parent process group of this component if applicable. | [optional] 
**Position** | Pointer to [**PositionDTO**](PositionDTO.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the reporting task. | [optional] 
**Type** | Pointer to **string** | The fully qualified type of the reporting task. | [optional] 
**Bundle** | Pointer to [**BundleDTO**](BundleDTO.md) |  | [optional] 
**State** | Pointer to **string** | The state of the reporting task. | [optional] 
**Comments** | Pointer to **string** | The comments of the reporting task. | [optional] 
**PersistsState** | Pointer to **bool** | Whether the reporting task persists state. | [optional] 
**Restricted** | Pointer to **bool** | Whether the reporting task requires elevated privileges. | [optional] 
**Deprecated** | Pointer to **bool** | Whether the reporting task has been deprecated. | [optional] 
**MultipleVersionsAvailable** | Pointer to **bool** | Whether the reporting task has multiple versions available. | [optional] 
**SchedulingPeriod** | Pointer to **string** | The frequency with which to schedule the reporting task. The format of the value willd epend on the valud of the schedulingStrategy. | [optional] 
**SchedulingStrategy** | Pointer to **string** | The scheduling strategy that determines how the schedulingPeriod value should be interpreted. | [optional] 
**DefaultSchedulingPeriod** | Pointer to **map[string]string** | The default scheduling period for the different scheduling strategies. | [optional] 
**Properties** | Pointer to **map[string]string** | The properties of the reporting task. | [optional] 
**Descriptors** | Pointer to [**map[string]PropertyDescriptorDTO**](PropertyDescriptorDTO.md) | The descriptors for the reporting tasks properties. | [optional] 
**CustomUiUrl** | Pointer to **string** | The URL for the custom configuration UI for the reporting task. | [optional] 
**AnnotationData** | Pointer to **string** | The annotation data for the repoting task. This is how the custom UI relays configuration to the reporting task. | [optional] 
**ValidationErrors** | Pointer to **[]string** | Gets the validation errors from the reporting task. These validation errors represent the problems with the reporting task that must be resolved before it can be scheduled to run. | [optional] 
**ValidationStatus** | Pointer to **string** | Indicates whether the Processor is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the Processor is valid) | [optional] [readonly] 
**ActiveThreadCount** | Pointer to **int32** | The number of active threads for the reporting task. | [optional] 
**ExtensionMissing** | Pointer to **bool** | Whether the underlying extension is missing. | [optional] 

## Methods

### NewReportingTaskDTO

`func NewReportingTaskDTO() *ReportingTaskDTO`

NewReportingTaskDTO instantiates a new ReportingTaskDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportingTaskDTOWithDefaults

`func NewReportingTaskDTOWithDefaults() *ReportingTaskDTO`

NewReportingTaskDTOWithDefaults instantiates a new ReportingTaskDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReportingTaskDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportingTaskDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportingTaskDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReportingTaskDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersionedComponentId

`func (o *ReportingTaskDTO) GetVersionedComponentId() string`

GetVersionedComponentId returns the VersionedComponentId field if non-nil, zero value otherwise.

### GetVersionedComponentIdOk

`func (o *ReportingTaskDTO) GetVersionedComponentIdOk() (*string, bool)`

GetVersionedComponentIdOk returns a tuple with the VersionedComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedComponentId

`func (o *ReportingTaskDTO) SetVersionedComponentId(v string)`

SetVersionedComponentId sets VersionedComponentId field to given value.

### HasVersionedComponentId

`func (o *ReportingTaskDTO) HasVersionedComponentId() bool`

HasVersionedComponentId returns a boolean if a field has been set.

### GetParentGroupId

`func (o *ReportingTaskDTO) GetParentGroupId() string`

GetParentGroupId returns the ParentGroupId field if non-nil, zero value otherwise.

### GetParentGroupIdOk

`func (o *ReportingTaskDTO) GetParentGroupIdOk() (*string, bool)`

GetParentGroupIdOk returns a tuple with the ParentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroupId

`func (o *ReportingTaskDTO) SetParentGroupId(v string)`

SetParentGroupId sets ParentGroupId field to given value.

### HasParentGroupId

`func (o *ReportingTaskDTO) HasParentGroupId() bool`

HasParentGroupId returns a boolean if a field has been set.

### GetPosition

`func (o *ReportingTaskDTO) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ReportingTaskDTO) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ReportingTaskDTO) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ReportingTaskDTO) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetName

`func (o *ReportingTaskDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportingTaskDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportingTaskDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReportingTaskDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ReportingTaskDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReportingTaskDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReportingTaskDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ReportingTaskDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBundle

`func (o *ReportingTaskDTO) GetBundle() BundleDTO`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *ReportingTaskDTO) GetBundleOk() (*BundleDTO, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *ReportingTaskDTO) SetBundle(v BundleDTO)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *ReportingTaskDTO) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### GetState

`func (o *ReportingTaskDTO) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ReportingTaskDTO) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ReportingTaskDTO) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ReportingTaskDTO) HasState() bool`

HasState returns a boolean if a field has been set.

### GetComments

`func (o *ReportingTaskDTO) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ReportingTaskDTO) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ReportingTaskDTO) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ReportingTaskDTO) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetPersistsState

`func (o *ReportingTaskDTO) GetPersistsState() bool`

GetPersistsState returns the PersistsState field if non-nil, zero value otherwise.

### GetPersistsStateOk

`func (o *ReportingTaskDTO) GetPersistsStateOk() (*bool, bool)`

GetPersistsStateOk returns a tuple with the PersistsState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistsState

`func (o *ReportingTaskDTO) SetPersistsState(v bool)`

SetPersistsState sets PersistsState field to given value.

### HasPersistsState

`func (o *ReportingTaskDTO) HasPersistsState() bool`

HasPersistsState returns a boolean if a field has been set.

### GetRestricted

`func (o *ReportingTaskDTO) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *ReportingTaskDTO) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *ReportingTaskDTO) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *ReportingTaskDTO) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### GetDeprecated

`func (o *ReportingTaskDTO) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *ReportingTaskDTO) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *ReportingTaskDTO) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *ReportingTaskDTO) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetMultipleVersionsAvailable

`func (o *ReportingTaskDTO) GetMultipleVersionsAvailable() bool`

GetMultipleVersionsAvailable returns the MultipleVersionsAvailable field if non-nil, zero value otherwise.

### GetMultipleVersionsAvailableOk

`func (o *ReportingTaskDTO) GetMultipleVersionsAvailableOk() (*bool, bool)`

GetMultipleVersionsAvailableOk returns a tuple with the MultipleVersionsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVersionsAvailable

`func (o *ReportingTaskDTO) SetMultipleVersionsAvailable(v bool)`

SetMultipleVersionsAvailable sets MultipleVersionsAvailable field to given value.

### HasMultipleVersionsAvailable

`func (o *ReportingTaskDTO) HasMultipleVersionsAvailable() bool`

HasMultipleVersionsAvailable returns a boolean if a field has been set.

### GetSchedulingPeriod

`func (o *ReportingTaskDTO) GetSchedulingPeriod() string`

GetSchedulingPeriod returns the SchedulingPeriod field if non-nil, zero value otherwise.

### GetSchedulingPeriodOk

`func (o *ReportingTaskDTO) GetSchedulingPeriodOk() (*string, bool)`

GetSchedulingPeriodOk returns a tuple with the SchedulingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingPeriod

`func (o *ReportingTaskDTO) SetSchedulingPeriod(v string)`

SetSchedulingPeriod sets SchedulingPeriod field to given value.

### HasSchedulingPeriod

`func (o *ReportingTaskDTO) HasSchedulingPeriod() bool`

HasSchedulingPeriod returns a boolean if a field has been set.

### GetSchedulingStrategy

`func (o *ReportingTaskDTO) GetSchedulingStrategy() string`

GetSchedulingStrategy returns the SchedulingStrategy field if non-nil, zero value otherwise.

### GetSchedulingStrategyOk

`func (o *ReportingTaskDTO) GetSchedulingStrategyOk() (*string, bool)`

GetSchedulingStrategyOk returns a tuple with the SchedulingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingStrategy

`func (o *ReportingTaskDTO) SetSchedulingStrategy(v string)`

SetSchedulingStrategy sets SchedulingStrategy field to given value.

### HasSchedulingStrategy

`func (o *ReportingTaskDTO) HasSchedulingStrategy() bool`

HasSchedulingStrategy returns a boolean if a field has been set.

### GetDefaultSchedulingPeriod

`func (o *ReportingTaskDTO) GetDefaultSchedulingPeriod() map[string]string`

GetDefaultSchedulingPeriod returns the DefaultSchedulingPeriod field if non-nil, zero value otherwise.

### GetDefaultSchedulingPeriodOk

`func (o *ReportingTaskDTO) GetDefaultSchedulingPeriodOk() (*map[string]string, bool)`

GetDefaultSchedulingPeriodOk returns a tuple with the DefaultSchedulingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSchedulingPeriod

`func (o *ReportingTaskDTO) SetDefaultSchedulingPeriod(v map[string]string)`

SetDefaultSchedulingPeriod sets DefaultSchedulingPeriod field to given value.

### HasDefaultSchedulingPeriod

`func (o *ReportingTaskDTO) HasDefaultSchedulingPeriod() bool`

HasDefaultSchedulingPeriod returns a boolean if a field has been set.

### GetProperties

`func (o *ReportingTaskDTO) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ReportingTaskDTO) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ReportingTaskDTO) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ReportingTaskDTO) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetDescriptors

`func (o *ReportingTaskDTO) GetDescriptors() map[string]PropertyDescriptorDTO`

GetDescriptors returns the Descriptors field if non-nil, zero value otherwise.

### GetDescriptorsOk

`func (o *ReportingTaskDTO) GetDescriptorsOk() (*map[string]PropertyDescriptorDTO, bool)`

GetDescriptorsOk returns a tuple with the Descriptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptors

`func (o *ReportingTaskDTO) SetDescriptors(v map[string]PropertyDescriptorDTO)`

SetDescriptors sets Descriptors field to given value.

### HasDescriptors

`func (o *ReportingTaskDTO) HasDescriptors() bool`

HasDescriptors returns a boolean if a field has been set.

### GetCustomUiUrl

`func (o *ReportingTaskDTO) GetCustomUiUrl() string`

GetCustomUiUrl returns the CustomUiUrl field if non-nil, zero value otherwise.

### GetCustomUiUrlOk

`func (o *ReportingTaskDTO) GetCustomUiUrlOk() (*string, bool)`

GetCustomUiUrlOk returns a tuple with the CustomUiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUiUrl

`func (o *ReportingTaskDTO) SetCustomUiUrl(v string)`

SetCustomUiUrl sets CustomUiUrl field to given value.

### HasCustomUiUrl

`func (o *ReportingTaskDTO) HasCustomUiUrl() bool`

HasCustomUiUrl returns a boolean if a field has been set.

### GetAnnotationData

`func (o *ReportingTaskDTO) GetAnnotationData() string`

GetAnnotationData returns the AnnotationData field if non-nil, zero value otherwise.

### GetAnnotationDataOk

`func (o *ReportingTaskDTO) GetAnnotationDataOk() (*string, bool)`

GetAnnotationDataOk returns a tuple with the AnnotationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationData

`func (o *ReportingTaskDTO) SetAnnotationData(v string)`

SetAnnotationData sets AnnotationData field to given value.

### HasAnnotationData

`func (o *ReportingTaskDTO) HasAnnotationData() bool`

HasAnnotationData returns a boolean if a field has been set.

### GetValidationErrors

`func (o *ReportingTaskDTO) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *ReportingTaskDTO) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *ReportingTaskDTO) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *ReportingTaskDTO) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### GetValidationStatus

`func (o *ReportingTaskDTO) GetValidationStatus() string`

GetValidationStatus returns the ValidationStatus field if non-nil, zero value otherwise.

### GetValidationStatusOk

`func (o *ReportingTaskDTO) GetValidationStatusOk() (*string, bool)`

GetValidationStatusOk returns a tuple with the ValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStatus

`func (o *ReportingTaskDTO) SetValidationStatus(v string)`

SetValidationStatus sets ValidationStatus field to given value.

### HasValidationStatus

`func (o *ReportingTaskDTO) HasValidationStatus() bool`

HasValidationStatus returns a boolean if a field has been set.

### GetActiveThreadCount

`func (o *ReportingTaskDTO) GetActiveThreadCount() int32`

GetActiveThreadCount returns the ActiveThreadCount field if non-nil, zero value otherwise.

### GetActiveThreadCountOk

`func (o *ReportingTaskDTO) GetActiveThreadCountOk() (*int32, bool)`

GetActiveThreadCountOk returns a tuple with the ActiveThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveThreadCount

`func (o *ReportingTaskDTO) SetActiveThreadCount(v int32)`

SetActiveThreadCount sets ActiveThreadCount field to given value.

### HasActiveThreadCount

`func (o *ReportingTaskDTO) HasActiveThreadCount() bool`

HasActiveThreadCount returns a boolean if a field has been set.

### GetExtensionMissing

`func (o *ReportingTaskDTO) GetExtensionMissing() bool`

GetExtensionMissing returns the ExtensionMissing field if non-nil, zero value otherwise.

### GetExtensionMissingOk

`func (o *ReportingTaskDTO) GetExtensionMissingOk() (*bool, bool)`

GetExtensionMissingOk returns a tuple with the ExtensionMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionMissing

`func (o *ReportingTaskDTO) SetExtensionMissing(v bool)`

SetExtensionMissing sets ExtensionMissing field to given value.

### HasExtensionMissing

`func (o *ReportingTaskDTO) HasExtensionMissing() bool`

HasExtensionMissing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


