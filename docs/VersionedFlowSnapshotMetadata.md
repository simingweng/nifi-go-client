# VersionedFlowSnapshotMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | Pointer to [**JaxbLink**](JaxbLink.md) |  | [optional] 
**BucketIdentifier** | **string** | The identifier of the bucket this snapshot belongs to. | 
**FlowIdentifier** | **string** | The identifier of the flow this snapshot belongs to. | 
**Version** | **int32** | The version of this snapshot of the flow. | 
**Timestamp** | Pointer to **int64** | The timestamp when the flow was saved, as milliseconds since epoch. | [optional] [readonly] 
**Author** | Pointer to **string** | The user that created this snapshot of the flow. | [optional] [readonly] 
**Comments** | Pointer to **string** | The comments provided by the user when creating the snapshot. | [optional] 

## Methods

### NewVersionedFlowSnapshotMetadata

`func NewVersionedFlowSnapshotMetadata(bucketIdentifier string, flowIdentifier string, version int32, ) *VersionedFlowSnapshotMetadata`

NewVersionedFlowSnapshotMetadata instantiates a new VersionedFlowSnapshotMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionedFlowSnapshotMetadataWithDefaults

`func NewVersionedFlowSnapshotMetadataWithDefaults() *VersionedFlowSnapshotMetadata`

NewVersionedFlowSnapshotMetadataWithDefaults instantiates a new VersionedFlowSnapshotMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLink

`func (o *VersionedFlowSnapshotMetadata) GetLink() JaxbLink`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *VersionedFlowSnapshotMetadata) GetLinkOk() (*JaxbLink, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *VersionedFlowSnapshotMetadata) SetLink(v JaxbLink)`

SetLink sets Link field to given value.

### HasLink

`func (o *VersionedFlowSnapshotMetadata) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetBucketIdentifier

`func (o *VersionedFlowSnapshotMetadata) GetBucketIdentifier() string`

GetBucketIdentifier returns the BucketIdentifier field if non-nil, zero value otherwise.

### GetBucketIdentifierOk

`func (o *VersionedFlowSnapshotMetadata) GetBucketIdentifierOk() (*string, bool)`

GetBucketIdentifierOk returns a tuple with the BucketIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketIdentifier

`func (o *VersionedFlowSnapshotMetadata) SetBucketIdentifier(v string)`

SetBucketIdentifier sets BucketIdentifier field to given value.


### GetFlowIdentifier

`func (o *VersionedFlowSnapshotMetadata) GetFlowIdentifier() string`

GetFlowIdentifier returns the FlowIdentifier field if non-nil, zero value otherwise.

### GetFlowIdentifierOk

`func (o *VersionedFlowSnapshotMetadata) GetFlowIdentifierOk() (*string, bool)`

GetFlowIdentifierOk returns a tuple with the FlowIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowIdentifier

`func (o *VersionedFlowSnapshotMetadata) SetFlowIdentifier(v string)`

SetFlowIdentifier sets FlowIdentifier field to given value.


### GetVersion

`func (o *VersionedFlowSnapshotMetadata) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VersionedFlowSnapshotMetadata) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VersionedFlowSnapshotMetadata) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetTimestamp

`func (o *VersionedFlowSnapshotMetadata) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *VersionedFlowSnapshotMetadata) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *VersionedFlowSnapshotMetadata) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *VersionedFlowSnapshotMetadata) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetAuthor

`func (o *VersionedFlowSnapshotMetadata) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *VersionedFlowSnapshotMetadata) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *VersionedFlowSnapshotMetadata) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *VersionedFlowSnapshotMetadata) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetComments

`func (o *VersionedFlowSnapshotMetadata) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VersionedFlowSnapshotMetadata) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VersionedFlowSnapshotMetadata) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VersionedFlowSnapshotMetadata) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


