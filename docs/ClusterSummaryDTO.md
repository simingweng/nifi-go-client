# ClusterSummaryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectedNodes** | Pointer to **string** | When clustered, reports the number of nodes connected vs the number of nodes in the cluster. | [optional] 
**ConnectedNodeCount** | Pointer to **int32** | The number of nodes that are currently connected to the cluster | [optional] 
**TotalNodeCount** | Pointer to **int32** | The number of nodes in the cluster, regardless of whether or not they are connected | [optional] 
**ConnectedToCluster** | Pointer to **bool** | Whether this NiFi instance is connected to a cluster. | [optional] 
**Clustered** | Pointer to **bool** | Whether this NiFi instance is clustered. | [optional] 

## Methods

### NewClusterSummaryDTO

`func NewClusterSummaryDTO() *ClusterSummaryDTO`

NewClusterSummaryDTO instantiates a new ClusterSummaryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSummaryDTOWithDefaults

`func NewClusterSummaryDTOWithDefaults() *ClusterSummaryDTO`

NewClusterSummaryDTOWithDefaults instantiates a new ClusterSummaryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectedNodes

`func (o *ClusterSummaryDTO) GetConnectedNodes() string`

GetConnectedNodes returns the ConnectedNodes field if non-nil, zero value otherwise.

### GetConnectedNodesOk

`func (o *ClusterSummaryDTO) GetConnectedNodesOk() (*string, bool)`

GetConnectedNodesOk returns a tuple with the ConnectedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedNodes

`func (o *ClusterSummaryDTO) SetConnectedNodes(v string)`

SetConnectedNodes sets ConnectedNodes field to given value.

### HasConnectedNodes

`func (o *ClusterSummaryDTO) HasConnectedNodes() bool`

HasConnectedNodes returns a boolean if a field has been set.

### GetConnectedNodeCount

`func (o *ClusterSummaryDTO) GetConnectedNodeCount() int32`

GetConnectedNodeCount returns the ConnectedNodeCount field if non-nil, zero value otherwise.

### GetConnectedNodeCountOk

`func (o *ClusterSummaryDTO) GetConnectedNodeCountOk() (*int32, bool)`

GetConnectedNodeCountOk returns a tuple with the ConnectedNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedNodeCount

`func (o *ClusterSummaryDTO) SetConnectedNodeCount(v int32)`

SetConnectedNodeCount sets ConnectedNodeCount field to given value.

### HasConnectedNodeCount

`func (o *ClusterSummaryDTO) HasConnectedNodeCount() bool`

HasConnectedNodeCount returns a boolean if a field has been set.

### GetTotalNodeCount

`func (o *ClusterSummaryDTO) GetTotalNodeCount() int32`

GetTotalNodeCount returns the TotalNodeCount field if non-nil, zero value otherwise.

### GetTotalNodeCountOk

`func (o *ClusterSummaryDTO) GetTotalNodeCountOk() (*int32, bool)`

GetTotalNodeCountOk returns a tuple with the TotalNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNodeCount

`func (o *ClusterSummaryDTO) SetTotalNodeCount(v int32)`

SetTotalNodeCount sets TotalNodeCount field to given value.

### HasTotalNodeCount

`func (o *ClusterSummaryDTO) HasTotalNodeCount() bool`

HasTotalNodeCount returns a boolean if a field has been set.

### GetConnectedToCluster

`func (o *ClusterSummaryDTO) GetConnectedToCluster() bool`

GetConnectedToCluster returns the ConnectedToCluster field if non-nil, zero value otherwise.

### GetConnectedToClusterOk

`func (o *ClusterSummaryDTO) GetConnectedToClusterOk() (*bool, bool)`

GetConnectedToClusterOk returns a tuple with the ConnectedToCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedToCluster

`func (o *ClusterSummaryDTO) SetConnectedToCluster(v bool)`

SetConnectedToCluster sets ConnectedToCluster field to given value.

### HasConnectedToCluster

`func (o *ClusterSummaryDTO) HasConnectedToCluster() bool`

HasConnectedToCluster returns a boolean if a field has been set.

### GetClustered

`func (o *ClusterSummaryDTO) GetClustered() bool`

GetClustered returns the Clustered field if non-nil, zero value otherwise.

### GetClusteredOk

`func (o *ClusterSummaryDTO) GetClusteredOk() (*bool, bool)`

GetClusteredOk returns a tuple with the Clustered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClustered

`func (o *ClusterSummaryDTO) SetClustered(v bool)`

SetClustered sets Clustered field to given value.

### HasClustered

`func (o *ClusterSummaryDTO) HasClustered() bool`

HasClustered returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


