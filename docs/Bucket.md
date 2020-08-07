# Bucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | [**JaxbLink**](JaxbLink.md) |  | [optional] 
**Identifier** | **string** | An ID to uniquely identify this object. | [optional] [readonly] 
**Name** | **string** | The name of the bucket. | 
**CreatedTimestamp** | **int64** | The timestamp of when the bucket was first created. This is set by the server at creation time. | [optional] [readonly] 
**Description** | **string** | A description of the bucket. | [optional] 
**AllowBundleRedeploy** | **bool** | Indicates if this bucket allows the same version of an extension bundle to be redeployed and thus overwrite the existing artifact. By default this is false. | [optional] 
**AllowPublicRead** | **bool** | Indicates if this bucket allows read access to unauthenticated anonymous users | [optional] 
**Permissions** | [**Permissions**](Permissions.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


