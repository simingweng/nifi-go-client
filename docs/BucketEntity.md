# BucketEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Bucket** | Pointer to [**BucketDTO**](BucketDTO.md) |  | [optional] 
**Permissions** | Pointer to [**PermissionsDTO**](PermissionsDTO.md) |  | [optional] 

## Methods

### NewBucketEntity

`func NewBucketEntity() *BucketEntity`

NewBucketEntity instantiates a new BucketEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketEntityWithDefaults

`func NewBucketEntityWithDefaults() *BucketEntity`

NewBucketEntityWithDefaults instantiates a new BucketEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BucketEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BucketEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BucketEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BucketEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBucket

`func (o *BucketEntity) GetBucket() BucketDTO`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *BucketEntity) GetBucketOk() (*BucketDTO, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *BucketEntity) SetBucket(v BucketDTO)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *BucketEntity) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetPermissions

`func (o *BucketEntity) GetPermissions() PermissionsDTO`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *BucketEntity) GetPermissionsOk() (*PermissionsDTO, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *BucketEntity) SetPermissions(v PermissionsDTO)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *BucketEntity) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


