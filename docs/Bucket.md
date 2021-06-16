# Bucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | Pointer to [**JaxbLink**](JaxbLink.md) |  | [optional] 
**Identifier** | Pointer to **string** | An ID to uniquely identify this object. | [optional] [readonly] 
**Name** | **string** | The name of the bucket. | 
**CreatedTimestamp** | Pointer to **int64** | The timestamp of when the bucket was first created. This is set by the server at creation time. | [optional] [readonly] 
**Description** | Pointer to **string** | A description of the bucket. | [optional] 
**AllowBundleRedeploy** | Pointer to **bool** | Indicates if this bucket allows the same version of an extension bundle to be redeployed and thus overwrite the existing artifact. By default this is false. | [optional] 
**AllowPublicRead** | Pointer to **bool** | Indicates if this bucket allows read access to unauthenticated anonymous users | [optional] 
**Permissions** | Pointer to [**Permissions**](Permissions.md) |  | [optional] 
**Revision** | Pointer to [**RevisionInfo**](RevisionInfo.md) |  | [optional] 

## Methods

### NewBucket

`func NewBucket(name string, ) *Bucket`

NewBucket instantiates a new Bucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketWithDefaults

`func NewBucketWithDefaults() *Bucket`

NewBucketWithDefaults instantiates a new Bucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLink

`func (o *Bucket) GetLink() JaxbLink`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *Bucket) GetLinkOk() (*JaxbLink, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *Bucket) SetLink(v JaxbLink)`

SetLink sets Link field to given value.

### HasLink

`func (o *Bucket) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetIdentifier

`func (o *Bucket) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Bucket) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Bucket) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *Bucket) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *Bucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Bucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Bucket) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedTimestamp

`func (o *Bucket) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *Bucket) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *Bucket) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *Bucket) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetDescription

`func (o *Bucket) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Bucket) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Bucket) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Bucket) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAllowBundleRedeploy

`func (o *Bucket) GetAllowBundleRedeploy() bool`

GetAllowBundleRedeploy returns the AllowBundleRedeploy field if non-nil, zero value otherwise.

### GetAllowBundleRedeployOk

`func (o *Bucket) GetAllowBundleRedeployOk() (*bool, bool)`

GetAllowBundleRedeployOk returns a tuple with the AllowBundleRedeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBundleRedeploy

`func (o *Bucket) SetAllowBundleRedeploy(v bool)`

SetAllowBundleRedeploy sets AllowBundleRedeploy field to given value.

### HasAllowBundleRedeploy

`func (o *Bucket) HasAllowBundleRedeploy() bool`

HasAllowBundleRedeploy returns a boolean if a field has been set.

### GetAllowPublicRead

`func (o *Bucket) GetAllowPublicRead() bool`

GetAllowPublicRead returns the AllowPublicRead field if non-nil, zero value otherwise.

### GetAllowPublicReadOk

`func (o *Bucket) GetAllowPublicReadOk() (*bool, bool)`

GetAllowPublicReadOk returns a tuple with the AllowPublicRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPublicRead

`func (o *Bucket) SetAllowPublicRead(v bool)`

SetAllowPublicRead sets AllowPublicRead field to given value.

### HasAllowPublicRead

`func (o *Bucket) HasAllowPublicRead() bool`

HasAllowPublicRead returns a boolean if a field has been set.

### GetPermissions

`func (o *Bucket) GetPermissions() Permissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Bucket) GetPermissionsOk() (*Permissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Bucket) SetPermissions(v Permissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Bucket) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetRevision

`func (o *Bucket) GetRevision() RevisionInfo`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Bucket) GetRevisionOk() (*RevisionInfo, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Bucket) SetRevision(v RevisionInfo)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *Bucket) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


