# VersionedFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | Pointer to [**JaxbLink**](JaxbLink.md) |  | [optional] 
**Identifier** | Pointer to **string** | An ID to uniquely identify this object. | [optional] [readonly] 
**Name** | **string** | The name of the item. | 
**Description** | Pointer to **string** | A description of the item. | [optional] 
**BucketIdentifier** | **string** | The identifier of the bucket this items belongs to. This cannot be changed after the item is created. | 
**BucketName** | Pointer to **string** | The name of the bucket this items belongs to. | [optional] [readonly] 
**CreatedTimestamp** | Pointer to **int64** | The timestamp of when the item was created, as milliseconds since epoch. | [optional] [readonly] 
**ModifiedTimestamp** | Pointer to **int64** | The timestamp of when the item was last modified, as milliseconds since epoch. | [optional] [readonly] 
**Type** | **string** | The type of item. | 
**Permissions** | Pointer to [**Permissions**](Permissions.md) |  | [optional] 
**VersionCount** | Pointer to **int64** | The number of versions of this flow. | [optional] [readonly] 
**Revision** | Pointer to [**RevisionInfo**](RevisionInfo.md) |  | [optional] 

## Methods

### NewVersionedFlow

`func NewVersionedFlow(name string, bucketIdentifier string, type_ string, ) *VersionedFlow`

NewVersionedFlow instantiates a new VersionedFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionedFlowWithDefaults

`func NewVersionedFlowWithDefaults() *VersionedFlow`

NewVersionedFlowWithDefaults instantiates a new VersionedFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLink

`func (o *VersionedFlow) GetLink() JaxbLink`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *VersionedFlow) GetLinkOk() (*JaxbLink, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *VersionedFlow) SetLink(v JaxbLink)`

SetLink sets Link field to given value.

### HasLink

`func (o *VersionedFlow) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetIdentifier

`func (o *VersionedFlow) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *VersionedFlow) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *VersionedFlow) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *VersionedFlow) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *VersionedFlow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VersionedFlow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VersionedFlow) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VersionedFlow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VersionedFlow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VersionedFlow) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VersionedFlow) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBucketIdentifier

`func (o *VersionedFlow) GetBucketIdentifier() string`

GetBucketIdentifier returns the BucketIdentifier field if non-nil, zero value otherwise.

### GetBucketIdentifierOk

`func (o *VersionedFlow) GetBucketIdentifierOk() (*string, bool)`

GetBucketIdentifierOk returns a tuple with the BucketIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketIdentifier

`func (o *VersionedFlow) SetBucketIdentifier(v string)`

SetBucketIdentifier sets BucketIdentifier field to given value.


### GetBucketName

`func (o *VersionedFlow) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *VersionedFlow) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *VersionedFlow) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *VersionedFlow) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *VersionedFlow) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *VersionedFlow) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *VersionedFlow) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *VersionedFlow) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetModifiedTimestamp

`func (o *VersionedFlow) GetModifiedTimestamp() int64`

GetModifiedTimestamp returns the ModifiedTimestamp field if non-nil, zero value otherwise.

### GetModifiedTimestampOk

`func (o *VersionedFlow) GetModifiedTimestampOk() (*int64, bool)`

GetModifiedTimestampOk returns a tuple with the ModifiedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTimestamp

`func (o *VersionedFlow) SetModifiedTimestamp(v int64)`

SetModifiedTimestamp sets ModifiedTimestamp field to given value.

### HasModifiedTimestamp

`func (o *VersionedFlow) HasModifiedTimestamp() bool`

HasModifiedTimestamp returns a boolean if a field has been set.

### GetType

`func (o *VersionedFlow) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VersionedFlow) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VersionedFlow) SetType(v string)`

SetType sets Type field to given value.


### GetPermissions

`func (o *VersionedFlow) GetPermissions() Permissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *VersionedFlow) GetPermissionsOk() (*Permissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *VersionedFlow) SetPermissions(v Permissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *VersionedFlow) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetVersionCount

`func (o *VersionedFlow) GetVersionCount() int64`

GetVersionCount returns the VersionCount field if non-nil, zero value otherwise.

### GetVersionCountOk

`func (o *VersionedFlow) GetVersionCountOk() (*int64, bool)`

GetVersionCountOk returns a tuple with the VersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCount

`func (o *VersionedFlow) SetVersionCount(v int64)`

SetVersionCount sets VersionCount field to given value.

### HasVersionCount

`func (o *VersionedFlow) HasVersionCount() bool`

HasVersionCount returns a boolean if a field has been set.

### GetRevision

`func (o *VersionedFlow) GetRevision() RevisionInfo`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *VersionedFlow) GetRevisionOk() (*RevisionInfo, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *VersionedFlow) SetRevision(v RevisionInfo)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *VersionedFlow) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


