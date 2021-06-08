# ControllerServiceApiDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The fully qualified name of the service interface. | [optional] 
**Bundle** | Pointer to [**BundleDTO**](BundleDTO.md) |  | [optional] 

## Methods

### NewControllerServiceApiDTO

`func NewControllerServiceApiDTO() *ControllerServiceApiDTO`

NewControllerServiceApiDTO instantiates a new ControllerServiceApiDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerServiceApiDTOWithDefaults

`func NewControllerServiceApiDTOWithDefaults() *ControllerServiceApiDTO`

NewControllerServiceApiDTOWithDefaults instantiates a new ControllerServiceApiDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ControllerServiceApiDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ControllerServiceApiDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ControllerServiceApiDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ControllerServiceApiDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBundle

`func (o *ControllerServiceApiDTO) GetBundle() BundleDTO`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *ControllerServiceApiDTO) GetBundleOk() (*BundleDTO, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *ControllerServiceApiDTO) SetBundle(v BundleDTO)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *ControllerServiceApiDTO) HasBundle() bool`

HasBundle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


