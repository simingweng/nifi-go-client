# VersionedFlowCoordinates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistryUrl** | Pointer to **string** | The URL of the Flow Registry that contains the flow | [optional] 
**BucketId** | Pointer to **string** | The UUID of the bucket that the flow resides in | [optional] 
**FlowId** | Pointer to **string** | The UUID of the flow | [optional] 
**Version** | Pointer to **int32** | The version of the flow | [optional] 
**Latest** | Pointer to **bool** | Whether or not these coordinates point to the latest version of the flow | [optional] 

## Methods

### NewVersionedFlowCoordinates

`func NewVersionedFlowCoordinates() *VersionedFlowCoordinates`

NewVersionedFlowCoordinates instantiates a new VersionedFlowCoordinates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionedFlowCoordinatesWithDefaults

`func NewVersionedFlowCoordinatesWithDefaults() *VersionedFlowCoordinates`

NewVersionedFlowCoordinatesWithDefaults instantiates a new VersionedFlowCoordinates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistryUrl

`func (o *VersionedFlowCoordinates) GetRegistryUrl() string`

GetRegistryUrl returns the RegistryUrl field if non-nil, zero value otherwise.

### GetRegistryUrlOk

`func (o *VersionedFlowCoordinates) GetRegistryUrlOk() (*string, bool)`

GetRegistryUrlOk returns a tuple with the RegistryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryUrl

`func (o *VersionedFlowCoordinates) SetRegistryUrl(v string)`

SetRegistryUrl sets RegistryUrl field to given value.

### HasRegistryUrl

`func (o *VersionedFlowCoordinates) HasRegistryUrl() bool`

HasRegistryUrl returns a boolean if a field has been set.

### GetBucketId

`func (o *VersionedFlowCoordinates) GetBucketId() string`

GetBucketId returns the BucketId field if non-nil, zero value otherwise.

### GetBucketIdOk

`func (o *VersionedFlowCoordinates) GetBucketIdOk() (*string, bool)`

GetBucketIdOk returns a tuple with the BucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketId

`func (o *VersionedFlowCoordinates) SetBucketId(v string)`

SetBucketId sets BucketId field to given value.

### HasBucketId

`func (o *VersionedFlowCoordinates) HasBucketId() bool`

HasBucketId returns a boolean if a field has been set.

### GetFlowId

`func (o *VersionedFlowCoordinates) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *VersionedFlowCoordinates) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *VersionedFlowCoordinates) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *VersionedFlowCoordinates) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetVersion

`func (o *VersionedFlowCoordinates) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VersionedFlowCoordinates) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VersionedFlowCoordinates) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VersionedFlowCoordinates) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLatest

`func (o *VersionedFlowCoordinates) GetLatest() bool`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *VersionedFlowCoordinates) GetLatestOk() (*bool, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *VersionedFlowCoordinates) SetLatest(v bool)`

SetLatest sets Latest field to given value.

### HasLatest

`func (o *VersionedFlowCoordinates) HasLatest() bool`

HasLatest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


