# StorageUsageDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | The identifier of this storage location. The identifier will correspond to the identifier keyed in the storage configuration. | [optional] 
**FreeSpace** | Pointer to **string** | Amount of free space. | [optional] 
**TotalSpace** | Pointer to **string** | Amount of total space. | [optional] 
**UsedSpace** | Pointer to **string** | Amount of used space. | [optional] 
**FreeSpaceBytes** | Pointer to **int64** | The number of bytes of free space. | [optional] 
**TotalSpaceBytes** | Pointer to **int64** | The number of bytes of total space. | [optional] 
**UsedSpaceBytes** | Pointer to **int64** | The number of bytes of used space. | [optional] 
**Utilization** | Pointer to **string** | Utilization of this storage location. | [optional] 

## Methods

### NewStorageUsageDTO

`func NewStorageUsageDTO() *StorageUsageDTO`

NewStorageUsageDTO instantiates a new StorageUsageDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageUsageDTOWithDefaults

`func NewStorageUsageDTOWithDefaults() *StorageUsageDTO`

NewStorageUsageDTOWithDefaults instantiates a new StorageUsageDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *StorageUsageDTO) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *StorageUsageDTO) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *StorageUsageDTO) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *StorageUsageDTO) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetFreeSpace

`func (o *StorageUsageDTO) GetFreeSpace() string`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *StorageUsageDTO) GetFreeSpaceOk() (*string, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *StorageUsageDTO) SetFreeSpace(v string)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *StorageUsageDTO) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### GetTotalSpace

`func (o *StorageUsageDTO) GetTotalSpace() string`

GetTotalSpace returns the TotalSpace field if non-nil, zero value otherwise.

### GetTotalSpaceOk

`func (o *StorageUsageDTO) GetTotalSpaceOk() (*string, bool)`

GetTotalSpaceOk returns a tuple with the TotalSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpace

`func (o *StorageUsageDTO) SetTotalSpace(v string)`

SetTotalSpace sets TotalSpace field to given value.

### HasTotalSpace

`func (o *StorageUsageDTO) HasTotalSpace() bool`

HasTotalSpace returns a boolean if a field has been set.

### GetUsedSpace

`func (o *StorageUsageDTO) GetUsedSpace() string`

GetUsedSpace returns the UsedSpace field if non-nil, zero value otherwise.

### GetUsedSpaceOk

`func (o *StorageUsageDTO) GetUsedSpaceOk() (*string, bool)`

GetUsedSpaceOk returns a tuple with the UsedSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSpace

`func (o *StorageUsageDTO) SetUsedSpace(v string)`

SetUsedSpace sets UsedSpace field to given value.

### HasUsedSpace

`func (o *StorageUsageDTO) HasUsedSpace() bool`

HasUsedSpace returns a boolean if a field has been set.

### GetFreeSpaceBytes

`func (o *StorageUsageDTO) GetFreeSpaceBytes() int64`

GetFreeSpaceBytes returns the FreeSpaceBytes field if non-nil, zero value otherwise.

### GetFreeSpaceBytesOk

`func (o *StorageUsageDTO) GetFreeSpaceBytesOk() (*int64, bool)`

GetFreeSpaceBytesOk returns a tuple with the FreeSpaceBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpaceBytes

`func (o *StorageUsageDTO) SetFreeSpaceBytes(v int64)`

SetFreeSpaceBytes sets FreeSpaceBytes field to given value.

### HasFreeSpaceBytes

`func (o *StorageUsageDTO) HasFreeSpaceBytes() bool`

HasFreeSpaceBytes returns a boolean if a field has been set.

### GetTotalSpaceBytes

`func (o *StorageUsageDTO) GetTotalSpaceBytes() int64`

GetTotalSpaceBytes returns the TotalSpaceBytes field if non-nil, zero value otherwise.

### GetTotalSpaceBytesOk

`func (o *StorageUsageDTO) GetTotalSpaceBytesOk() (*int64, bool)`

GetTotalSpaceBytesOk returns a tuple with the TotalSpaceBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpaceBytes

`func (o *StorageUsageDTO) SetTotalSpaceBytes(v int64)`

SetTotalSpaceBytes sets TotalSpaceBytes field to given value.

### HasTotalSpaceBytes

`func (o *StorageUsageDTO) HasTotalSpaceBytes() bool`

HasTotalSpaceBytes returns a boolean if a field has been set.

### GetUsedSpaceBytes

`func (o *StorageUsageDTO) GetUsedSpaceBytes() int64`

GetUsedSpaceBytes returns the UsedSpaceBytes field if non-nil, zero value otherwise.

### GetUsedSpaceBytesOk

`func (o *StorageUsageDTO) GetUsedSpaceBytesOk() (*int64, bool)`

GetUsedSpaceBytesOk returns a tuple with the UsedSpaceBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSpaceBytes

`func (o *StorageUsageDTO) SetUsedSpaceBytes(v int64)`

SetUsedSpaceBytes sets UsedSpaceBytes field to given value.

### HasUsedSpaceBytes

`func (o *StorageUsageDTO) HasUsedSpaceBytes() bool`

HasUsedSpaceBytes returns a boolean if a field has been set.

### GetUtilization

`func (o *StorageUsageDTO) GetUtilization() string`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *StorageUsageDTO) GetUtilizationOk() (*string, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *StorageUsageDTO) SetUtilization(v string)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *StorageUsageDTO) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


