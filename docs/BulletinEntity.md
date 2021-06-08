# BulletinEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **string** | When this bulletin was generated. | [optional] 
**NodeAddress** | Pointer to **string** |  | [optional] 
**CanRead** | Pointer to **bool** | Indicates whether the user can read a given resource. | [optional] [readonly] 
**Bulletin** | Pointer to [**BulletinDTO**](BulletinDTO.md) |  | [optional] 

## Methods

### NewBulletinEntity

`func NewBulletinEntity() *BulletinEntity`

NewBulletinEntity instantiates a new BulletinEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulletinEntityWithDefaults

`func NewBulletinEntityWithDefaults() *BulletinEntity`

NewBulletinEntityWithDefaults instantiates a new BulletinEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulletinEntity) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulletinEntity) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulletinEntity) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BulletinEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupId

`func (o *BulletinEntity) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *BulletinEntity) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *BulletinEntity) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *BulletinEntity) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetSourceId

`func (o *BulletinEntity) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BulletinEntity) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BulletinEntity) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *BulletinEntity) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetTimestamp

`func (o *BulletinEntity) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BulletinEntity) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BulletinEntity) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BulletinEntity) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetNodeAddress

`func (o *BulletinEntity) GetNodeAddress() string`

GetNodeAddress returns the NodeAddress field if non-nil, zero value otherwise.

### GetNodeAddressOk

`func (o *BulletinEntity) GetNodeAddressOk() (*string, bool)`

GetNodeAddressOk returns a tuple with the NodeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAddress

`func (o *BulletinEntity) SetNodeAddress(v string)`

SetNodeAddress sets NodeAddress field to given value.

### HasNodeAddress

`func (o *BulletinEntity) HasNodeAddress() bool`

HasNodeAddress returns a boolean if a field has been set.

### GetCanRead

`func (o *BulletinEntity) GetCanRead() bool`

GetCanRead returns the CanRead field if non-nil, zero value otherwise.

### GetCanReadOk

`func (o *BulletinEntity) GetCanReadOk() (*bool, bool)`

GetCanReadOk returns a tuple with the CanRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRead

`func (o *BulletinEntity) SetCanRead(v bool)`

SetCanRead sets CanRead field to given value.

### HasCanRead

`func (o *BulletinEntity) HasCanRead() bool`

HasCanRead returns a boolean if a field has been set.

### GetBulletin

`func (o *BulletinEntity) GetBulletin() BulletinDTO`

GetBulletin returns the Bulletin field if non-nil, zero value otherwise.

### GetBulletinOk

`func (o *BulletinEntity) GetBulletinOk() (*BulletinDTO, bool)`

GetBulletinOk returns a tuple with the Bulletin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletin

`func (o *BulletinEntity) SetBulletin(v BulletinDTO)`

SetBulletin sets Bulletin field to given value.

### HasBulletin

`func (o *BulletinEntity) HasBulletin() bool`

HasBulletin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


