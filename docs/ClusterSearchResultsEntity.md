# ClusterSearchResultsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeResults** | Pointer to [**[]NodeSearchResultDTO**](NodeSearchResultDTO.md) |  | [optional] 

## Methods

### NewClusterSearchResultsEntity

`func NewClusterSearchResultsEntity() *ClusterSearchResultsEntity`

NewClusterSearchResultsEntity instantiates a new ClusterSearchResultsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSearchResultsEntityWithDefaults

`func NewClusterSearchResultsEntityWithDefaults() *ClusterSearchResultsEntity`

NewClusterSearchResultsEntityWithDefaults instantiates a new ClusterSearchResultsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeResults

`func (o *ClusterSearchResultsEntity) GetNodeResults() []NodeSearchResultDTO`

GetNodeResults returns the NodeResults field if non-nil, zero value otherwise.

### GetNodeResultsOk

`func (o *ClusterSearchResultsEntity) GetNodeResultsOk() (*[]NodeSearchResultDTO, bool)`

GetNodeResultsOk returns a tuple with the NodeResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeResults

`func (o *ClusterSearchResultsEntity) SetNodeResults(v []NodeSearchResultDTO)`

SetNodeResults sets NodeResults field to given value.

### HasNodeResults

`func (o *ClusterSearchResultsEntity) HasNodeResults() bool`

HasNodeResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


