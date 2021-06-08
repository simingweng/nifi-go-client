# VersionedControllerService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | The component&#39;s unique identifier | [optional] 
**Name** | Pointer to **string** | The component&#39;s name | [optional] 
**Comments** | Pointer to **string** | The user-supplied comments for the component | [optional] 
**Position** | Pointer to [**Position**](Position.md) |  | [optional] 
**Type** | Pointer to **string** | The type of the controller service. | [optional] 
**Bundle** | Pointer to [**Bundle**](Bundle.md) |  | [optional] 
**ControllerServiceApis** | Pointer to [**[]ControllerServiceAPI**](ControllerServiceAPI.md) | Lists the APIs this Controller Service implements. | [optional] 
**Properties** | Pointer to **map[string]string** | The properties of the controller service. | [optional] 
**PropertyDescriptors** | Pointer to [**map[string]VersionedPropertyDescriptor**](VersionedPropertyDescriptor.md) | The property descriptors for the processor. | [optional] 
**AnnotationData** | Pointer to **string** | The annotation for the controller service. This is how the custom UI relays configuration to the controller service. | [optional] 
**ComponentType** | Pointer to **string** |  | [optional] 
**GroupIdentifier** | Pointer to **string** | The ID of the Process Group that this component belongs to | [optional] 

## Methods

### NewVersionedControllerService

`func NewVersionedControllerService() *VersionedControllerService`

NewVersionedControllerService instantiates a new VersionedControllerService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionedControllerServiceWithDefaults

`func NewVersionedControllerServiceWithDefaults() *VersionedControllerService`

NewVersionedControllerServiceWithDefaults instantiates a new VersionedControllerService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *VersionedControllerService) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *VersionedControllerService) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *VersionedControllerService) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *VersionedControllerService) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *VersionedControllerService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VersionedControllerService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VersionedControllerService) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VersionedControllerService) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComments

`func (o *VersionedControllerService) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VersionedControllerService) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VersionedControllerService) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VersionedControllerService) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetPosition

`func (o *VersionedControllerService) GetPosition() Position`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *VersionedControllerService) GetPositionOk() (*Position, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *VersionedControllerService) SetPosition(v Position)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *VersionedControllerService) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetType

`func (o *VersionedControllerService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VersionedControllerService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VersionedControllerService) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VersionedControllerService) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBundle

`func (o *VersionedControllerService) GetBundle() Bundle`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *VersionedControllerService) GetBundleOk() (*Bundle, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *VersionedControllerService) SetBundle(v Bundle)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *VersionedControllerService) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### GetControllerServiceApis

`func (o *VersionedControllerService) GetControllerServiceApis() []ControllerServiceAPI`

GetControllerServiceApis returns the ControllerServiceApis field if non-nil, zero value otherwise.

### GetControllerServiceApisOk

`func (o *VersionedControllerService) GetControllerServiceApisOk() (*[]ControllerServiceAPI, bool)`

GetControllerServiceApisOk returns a tuple with the ControllerServiceApis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerServiceApis

`func (o *VersionedControllerService) SetControllerServiceApis(v []ControllerServiceAPI)`

SetControllerServiceApis sets ControllerServiceApis field to given value.

### HasControllerServiceApis

`func (o *VersionedControllerService) HasControllerServiceApis() bool`

HasControllerServiceApis returns a boolean if a field has been set.

### GetProperties

`func (o *VersionedControllerService) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *VersionedControllerService) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *VersionedControllerService) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *VersionedControllerService) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetPropertyDescriptors

`func (o *VersionedControllerService) GetPropertyDescriptors() map[string]VersionedPropertyDescriptor`

GetPropertyDescriptors returns the PropertyDescriptors field if non-nil, zero value otherwise.

### GetPropertyDescriptorsOk

`func (o *VersionedControllerService) GetPropertyDescriptorsOk() (*map[string]VersionedPropertyDescriptor, bool)`

GetPropertyDescriptorsOk returns a tuple with the PropertyDescriptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyDescriptors

`func (o *VersionedControllerService) SetPropertyDescriptors(v map[string]VersionedPropertyDescriptor)`

SetPropertyDescriptors sets PropertyDescriptors field to given value.

### HasPropertyDescriptors

`func (o *VersionedControllerService) HasPropertyDescriptors() bool`

HasPropertyDescriptors returns a boolean if a field has been set.

### GetAnnotationData

`func (o *VersionedControllerService) GetAnnotationData() string`

GetAnnotationData returns the AnnotationData field if non-nil, zero value otherwise.

### GetAnnotationDataOk

`func (o *VersionedControllerService) GetAnnotationDataOk() (*string, bool)`

GetAnnotationDataOk returns a tuple with the AnnotationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationData

`func (o *VersionedControllerService) SetAnnotationData(v string)`

SetAnnotationData sets AnnotationData field to given value.

### HasAnnotationData

`func (o *VersionedControllerService) HasAnnotationData() bool`

HasAnnotationData returns a boolean if a field has been set.

### GetComponentType

`func (o *VersionedControllerService) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *VersionedControllerService) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *VersionedControllerService) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *VersionedControllerService) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### GetGroupIdentifier

`func (o *VersionedControllerService) GetGroupIdentifier() string`

GetGroupIdentifier returns the GroupIdentifier field if non-nil, zero value otherwise.

### GetGroupIdentifierOk

`func (o *VersionedControllerService) GetGroupIdentifierOk() (*string, bool)`

GetGroupIdentifierOk returns a tuple with the GroupIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIdentifier

`func (o *VersionedControllerService) SetGroupIdentifier(v string)`

SetGroupIdentifier sets GroupIdentifier field to given value.

### HasGroupIdentifier

`func (o *VersionedControllerService) HasGroupIdentifier() bool`

HasGroupIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


