# BundleDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | The group of the bundle. | [optional] 
**Artifact** | Pointer to **string** | The artifact of the bundle. | [optional] 
**Version** | Pointer to **string** | The version of the bundle. | [optional] 

## Methods

### NewBundleDTO

`func NewBundleDTO() *BundleDTO`

NewBundleDTO instantiates a new BundleDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleDTOWithDefaults

`func NewBundleDTOWithDefaults() *BundleDTO`

NewBundleDTOWithDefaults instantiates a new BundleDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *BundleDTO) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BundleDTO) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BundleDTO) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BundleDTO) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetArtifact

`func (o *BundleDTO) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *BundleDTO) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *BundleDTO) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *BundleDTO) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetVersion

`func (o *BundleDTO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BundleDTO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BundleDTO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BundleDTO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


