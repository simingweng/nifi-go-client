# ClusterDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to [**[]NodeDTO**](NodeDTO.md) | The collection of nodes that are part of the cluster. | [optional] 
**Generated** | Pointer to **string** | The timestamp the report was generated. | [optional] 

## Methods

### NewClusterDTO

`func NewClusterDTO() *ClusterDTO`

NewClusterDTO instantiates a new ClusterDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDTOWithDefaults

`func NewClusterDTOWithDefaults() *ClusterDTO`

NewClusterDTOWithDefaults instantiates a new ClusterDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *ClusterDTO) GetNodes() []NodeDTO`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ClusterDTO) GetNodesOk() (*[]NodeDTO, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ClusterDTO) SetNodes(v []NodeDTO)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ClusterDTO) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetGenerated

`func (o *ClusterDTO) GetGenerated() string`

GetGenerated returns the Generated field if non-nil, zero value otherwise.

### GetGeneratedOk

`func (o *ClusterDTO) GetGeneratedOk() (*string, bool)`

GetGeneratedOk returns a tuple with the Generated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerated

`func (o *ClusterDTO) SetGenerated(v string)`

SetGenerated sets Generated field to given value.

### HasGenerated

`func (o *ClusterDTO) HasGenerated() bool`

HasGenerated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


