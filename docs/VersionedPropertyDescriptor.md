# VersionedPropertyDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the property | [optional] 
**DisplayName** | Pointer to **string** | The display name of the property | [optional] 
**IdentifiesControllerService** | Pointer to **bool** | Whether or not the property provides the identifier of a Controller Service | [optional] 
**Sensitive** | Pointer to **bool** | Whether or not the property is considered sensitive | [optional] 

## Methods

### NewVersionedPropertyDescriptor

`func NewVersionedPropertyDescriptor() *VersionedPropertyDescriptor`

NewVersionedPropertyDescriptor instantiates a new VersionedPropertyDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionedPropertyDescriptorWithDefaults

`func NewVersionedPropertyDescriptorWithDefaults() *VersionedPropertyDescriptor`

NewVersionedPropertyDescriptorWithDefaults instantiates a new VersionedPropertyDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VersionedPropertyDescriptor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VersionedPropertyDescriptor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VersionedPropertyDescriptor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VersionedPropertyDescriptor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *VersionedPropertyDescriptor) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VersionedPropertyDescriptor) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VersionedPropertyDescriptor) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *VersionedPropertyDescriptor) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIdentifiesControllerService

`func (o *VersionedPropertyDescriptor) GetIdentifiesControllerService() bool`

GetIdentifiesControllerService returns the IdentifiesControllerService field if non-nil, zero value otherwise.

### GetIdentifiesControllerServiceOk

`func (o *VersionedPropertyDescriptor) GetIdentifiesControllerServiceOk() (*bool, bool)`

GetIdentifiesControllerServiceOk returns a tuple with the IdentifiesControllerService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiesControllerService

`func (o *VersionedPropertyDescriptor) SetIdentifiesControllerService(v bool)`

SetIdentifiesControllerService sets IdentifiesControllerService field to given value.

### HasIdentifiesControllerService

`func (o *VersionedPropertyDescriptor) HasIdentifiesControllerService() bool`

HasIdentifiesControllerService returns a boolean if a field has been set.

### GetSensitive

`func (o *VersionedPropertyDescriptor) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *VersionedPropertyDescriptor) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *VersionedPropertyDescriptor) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.

### HasSensitive

`func (o *VersionedPropertyDescriptor) HasSensitive() bool`

HasSensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


