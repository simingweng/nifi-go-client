# BucketDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The bucket identifier | [optional] 
**Name** | Pointer to **string** | The bucket name | [optional] 
**Description** | Pointer to **string** | The bucket description | [optional] 
**Created** | Pointer to **int64** | The created timestamp of this bucket | [optional] 

## Methods

### NewBucketDTO

`func NewBucketDTO() *BucketDTO`

NewBucketDTO instantiates a new BucketDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketDTOWithDefaults

`func NewBucketDTOWithDefaults() *BucketDTO`

NewBucketDTOWithDefaults instantiates a new BucketDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BucketDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BucketDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BucketDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BucketDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BucketDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BucketDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BucketDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BucketDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *BucketDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BucketDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BucketDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BucketDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreated

`func (o *BucketDTO) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BucketDTO) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BucketDTO) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *BucketDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


