# ProvenanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the provenance query. | [optional] 
**Uri** | **string** | The URI for this query. Used for obtaining/deleting the request at a later time | [optional] 
**SubmissionTime** | **string** | The timestamp when the query was submitted. | [optional] 
**Expiration** | **string** | The timestamp when the query will expire. | [optional] 
**PercentCompleted** | **int32** | The current percent complete. | [optional] 
**Finished** | **bool** | Whether the query has finished. | [optional] 
**Request** | [**ProvenanceRequestDto**](ProvenanceRequestDTO.md) |  | [optional] 
**Results** | [**ProvenanceResultsDto**](ProvenanceResultsDTO.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


