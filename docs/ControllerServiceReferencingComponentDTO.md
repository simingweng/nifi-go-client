# ControllerServiceReferencingComponentDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | The group id for the component referencing a controller service. If this component is another controller service or a reporting task, this field is blank. | [optional] 
**Id** | Pointer to **string** | The id of the component referencing a controller service. | [optional] 
**Name** | Pointer to **string** | The name of the component referencing a controller service. | [optional] 
**Type** | Pointer to **string** | The type of the component referencing a controller service in simple Java class name format without package name. | [optional] 
**State** | Pointer to **string** | The scheduled state of a processor or reporting task referencing a controller service. If this component is another controller service, this field represents the controller service state. | [optional] 
**Properties** | Pointer to **map[string]string** | The properties for the component. | [optional] 
**Descriptors** | Pointer to [**map[string]PropertyDescriptorDTO**](PropertyDescriptorDTO.md) | The descriptors for the component properties. | [optional] 
**ValidationErrors** | Pointer to **[]string** | The validation errors for the component. | [optional] 
**ReferenceType** | Pointer to **string** | The type of reference this is. | [optional] 
**ActiveThreadCount** | Pointer to **int32** | The number of active threads for the referencing component. | [optional] 
**ReferenceCycle** | Pointer to **bool** | If the referencing component represents a controller service, this indicates whether it has already been represented in this hierarchy. | [optional] 
**ReferencingComponents** | Pointer to [**[]ControllerServiceReferencingComponentEntity**](ControllerServiceReferencingComponentEntity.md) | If the referencing component represents a controller service, these are the components that reference it. | [optional] 

## Methods

### NewControllerServiceReferencingComponentDTO

`func NewControllerServiceReferencingComponentDTO() *ControllerServiceReferencingComponentDTO`

NewControllerServiceReferencingComponentDTO instantiates a new ControllerServiceReferencingComponentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerServiceReferencingComponentDTOWithDefaults

`func NewControllerServiceReferencingComponentDTOWithDefaults() *ControllerServiceReferencingComponentDTO`

NewControllerServiceReferencingComponentDTOWithDefaults instantiates a new ControllerServiceReferencingComponentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *ControllerServiceReferencingComponentDTO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ControllerServiceReferencingComponentDTO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ControllerServiceReferencingComponentDTO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ControllerServiceReferencingComponentDTO) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *ControllerServiceReferencingComponentDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControllerServiceReferencingComponentDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControllerServiceReferencingComponentDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ControllerServiceReferencingComponentDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ControllerServiceReferencingComponentDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControllerServiceReferencingComponentDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControllerServiceReferencingComponentDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ControllerServiceReferencingComponentDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ControllerServiceReferencingComponentDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ControllerServiceReferencingComponentDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ControllerServiceReferencingComponentDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ControllerServiceReferencingComponentDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *ControllerServiceReferencingComponentDTO) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ControllerServiceReferencingComponentDTO) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ControllerServiceReferencingComponentDTO) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ControllerServiceReferencingComponentDTO) HasState() bool`

HasState returns a boolean if a field has been set.

### GetProperties

`func (o *ControllerServiceReferencingComponentDTO) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ControllerServiceReferencingComponentDTO) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ControllerServiceReferencingComponentDTO) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ControllerServiceReferencingComponentDTO) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetDescriptors

`func (o *ControllerServiceReferencingComponentDTO) GetDescriptors() map[string]PropertyDescriptorDTO`

GetDescriptors returns the Descriptors field if non-nil, zero value otherwise.

### GetDescriptorsOk

`func (o *ControllerServiceReferencingComponentDTO) GetDescriptorsOk() (*map[string]PropertyDescriptorDTO, bool)`

GetDescriptorsOk returns a tuple with the Descriptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptors

`func (o *ControllerServiceReferencingComponentDTO) SetDescriptors(v map[string]PropertyDescriptorDTO)`

SetDescriptors sets Descriptors field to given value.

### HasDescriptors

`func (o *ControllerServiceReferencingComponentDTO) HasDescriptors() bool`

HasDescriptors returns a boolean if a field has been set.

### GetValidationErrors

`func (o *ControllerServiceReferencingComponentDTO) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *ControllerServiceReferencingComponentDTO) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *ControllerServiceReferencingComponentDTO) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *ControllerServiceReferencingComponentDTO) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### GetReferenceType

`func (o *ControllerServiceReferencingComponentDTO) GetReferenceType() string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *ControllerServiceReferencingComponentDTO) GetReferenceTypeOk() (*string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *ControllerServiceReferencingComponentDTO) SetReferenceType(v string)`

SetReferenceType sets ReferenceType field to given value.

### HasReferenceType

`func (o *ControllerServiceReferencingComponentDTO) HasReferenceType() bool`

HasReferenceType returns a boolean if a field has been set.

### GetActiveThreadCount

`func (o *ControllerServiceReferencingComponentDTO) GetActiveThreadCount() int32`

GetActiveThreadCount returns the ActiveThreadCount field if non-nil, zero value otherwise.

### GetActiveThreadCountOk

`func (o *ControllerServiceReferencingComponentDTO) GetActiveThreadCountOk() (*int32, bool)`

GetActiveThreadCountOk returns a tuple with the ActiveThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveThreadCount

`func (o *ControllerServiceReferencingComponentDTO) SetActiveThreadCount(v int32)`

SetActiveThreadCount sets ActiveThreadCount field to given value.

### HasActiveThreadCount

`func (o *ControllerServiceReferencingComponentDTO) HasActiveThreadCount() bool`

HasActiveThreadCount returns a boolean if a field has been set.

### GetReferenceCycle

`func (o *ControllerServiceReferencingComponentDTO) GetReferenceCycle() bool`

GetReferenceCycle returns the ReferenceCycle field if non-nil, zero value otherwise.

### GetReferenceCycleOk

`func (o *ControllerServiceReferencingComponentDTO) GetReferenceCycleOk() (*bool, bool)`

GetReferenceCycleOk returns a tuple with the ReferenceCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCycle

`func (o *ControllerServiceReferencingComponentDTO) SetReferenceCycle(v bool)`

SetReferenceCycle sets ReferenceCycle field to given value.

### HasReferenceCycle

`func (o *ControllerServiceReferencingComponentDTO) HasReferenceCycle() bool`

HasReferenceCycle returns a boolean if a field has been set.

### GetReferencingComponents

`func (o *ControllerServiceReferencingComponentDTO) GetReferencingComponents() []ControllerServiceReferencingComponentEntity`

GetReferencingComponents returns the ReferencingComponents field if non-nil, zero value otherwise.

### GetReferencingComponentsOk

`func (o *ControllerServiceReferencingComponentDTO) GetReferencingComponentsOk() (*[]ControllerServiceReferencingComponentEntity, bool)`

GetReferencingComponentsOk returns a tuple with the ReferencingComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencingComponents

`func (o *ControllerServiceReferencingComponentDTO) SetReferencingComponents(v []ControllerServiceReferencingComponentEntity)`

SetReferencingComponents sets ReferencingComponents field to given value.

### HasReferencingComponents

`func (o *ControllerServiceReferencingComponentDTO) HasReferencingComponents() bool`

HasReferencingComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


