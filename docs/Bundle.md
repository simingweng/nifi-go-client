# Bundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | The group of the bundle | [optional] 
**Artifact** | Pointer to **string** | The artifact of the bundle | [optional] 
**Version** | Pointer to **string** | The version of the bundle | [optional] 

## Methods

### NewBundle

`func NewBundle() *Bundle`

NewBundle instantiates a new Bundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleWithDefaults

`func NewBundleWithDefaults() *Bundle`

NewBundleWithDefaults instantiates a new Bundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *Bundle) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Bundle) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Bundle) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Bundle) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetArtifact

`func (o *Bundle) GetArtifact() string`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *Bundle) GetArtifactOk() (*string, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *Bundle) SetArtifact(v string)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *Bundle) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetVersion

`func (o *Bundle) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Bundle) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Bundle) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Bundle) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


