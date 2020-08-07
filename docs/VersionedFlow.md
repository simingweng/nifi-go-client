# VersionedFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | [**JaxbLink**](JaxbLink.md) |  | [optional] 
**Identifier** | **string** | An ID to uniquely identify this object. | [optional] [readonly] 
**Name** | **string** | The name of the item. | 
**Description** | **string** | A description of the item. | [optional] 
**BucketIdentifier** | **string** | The identifier of the bucket this items belongs to. This cannot be changed after the item is created. | 
**BucketName** | **string** | The name of the bucket this items belongs to. | [optional] [readonly] 
**CreatedTimestamp** | **int64** | The timestamp of when the item was created, as milliseconds since epoch. | [optional] [readonly] 
**ModifiedTimestamp** | **int64** | The timestamp of when the item was last modified, as milliseconds since epoch. | [optional] [readonly] 
**Type** | **string** | The type of item. | 
**Permissions** | [**Permissions**](Permissions.md) |  | [optional] 
**VersionCount** | **int64** | The number of versions of this flow. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


