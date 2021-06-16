# VersionedConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | The component&#39;s unique identifier | [optional] 
**Name** | Pointer to **string** | The component&#39;s name | [optional] 
**Comments** | Pointer to **string** | The user-supplied comments for the component | [optional] 
**Position** | Pointer to [**Position**](Position.md) |  | [optional] 
**Source** | Pointer to [**ConnectableComponent**](ConnectableComponent.md) |  | [optional] 
**Destination** | Pointer to [**ConnectableComponent**](ConnectableComponent.md) |  | [optional] 
**LabelIndex** | Pointer to **int32** | The index of the bend point where to place the connection label. | [optional] 
**ZIndex** | Pointer to **int64** | The z index of the connection. | [optional] 
**SelectedRelationships** | Pointer to **[]string** | The selected relationship that comprise the connection. | [optional] 
**BackPressureObjectThreshold** | Pointer to **int64** | The object count threshold for determining when back pressure is applied. Updating this value is a passive change in the sense that it won&#39;t impact whether existing files over the limit are affected but it does help feeder processors to stop pushing too much into this work queue. | [optional] 
**BackPressureDataSizeThreshold** | Pointer to **string** | The object data size threshold for determining when back pressure is applied. Updating this value is a passive change in the sense that it won&#39;t impact whether existing files over the limit are affected but it does help feeder processors to stop pushing too much into this work queue. | [optional] 
**FlowFileExpiration** | Pointer to **string** | The amount of time a flow file may be in the flow before it will be automatically aged out of the flow. Once a flow file reaches this age it will be terminated from the flow the next time a processor attempts to start work on it. | [optional] 
**Prioritizers** | Pointer to **[]string** | The comparators used to prioritize the queue. | [optional] 
**Bends** | Pointer to [**[]Position**](Position.md) | The bend points on the connection. | [optional] 
**LoadBalanceStrategy** | Pointer to **string** | The Strategy to use for load balancing data across the cluster, or null, if no Load Balance Strategy has been specified. | [optional] 
**PartitioningAttribute** | Pointer to **string** | The attribute to use for partitioning data as it is load balanced across the cluster. If the Load Balance Strategy is configured to use PARTITION_BY_ATTRIBUTE, the value returned by this method is the name of the FlowFile Attribute that will be used to determine which node in the cluster should receive a given FlowFile. If the Load Balance Strategy is unset or is set to any other value, the Partitioning Attribute has no effect. | [optional] 
**LoadBalanceCompression** | Pointer to **string** | Whether or not compression should be used when transferring FlowFiles between nodes | [optional] 
**ComponentType** | Pointer to **string** |  | [optional] 
**GroupIdentifier** | Pointer to **string** | The ID of the Process Group that this component belongs to | [optional] 

## Methods

### NewVersionedConnection

`func NewVersionedConnection() *VersionedConnection`

NewVersionedConnection instantiates a new VersionedConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionedConnectionWithDefaults

`func NewVersionedConnectionWithDefaults() *VersionedConnection`

NewVersionedConnectionWithDefaults instantiates a new VersionedConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *VersionedConnection) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *VersionedConnection) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *VersionedConnection) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *VersionedConnection) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *VersionedConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VersionedConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VersionedConnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VersionedConnection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComments

`func (o *VersionedConnection) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VersionedConnection) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VersionedConnection) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VersionedConnection) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetPosition

`func (o *VersionedConnection) GetPosition() Position`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *VersionedConnection) GetPositionOk() (*Position, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *VersionedConnection) SetPosition(v Position)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *VersionedConnection) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetSource

`func (o *VersionedConnection) GetSource() ConnectableComponent`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *VersionedConnection) GetSourceOk() (*ConnectableComponent, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *VersionedConnection) SetSource(v ConnectableComponent)`

SetSource sets Source field to given value.

### HasSource

`func (o *VersionedConnection) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDestination

`func (o *VersionedConnection) GetDestination() ConnectableComponent`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *VersionedConnection) GetDestinationOk() (*ConnectableComponent, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *VersionedConnection) SetDestination(v ConnectableComponent)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *VersionedConnection) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetLabelIndex

`func (o *VersionedConnection) GetLabelIndex() int32`

GetLabelIndex returns the LabelIndex field if non-nil, zero value otherwise.

### GetLabelIndexOk

`func (o *VersionedConnection) GetLabelIndexOk() (*int32, bool)`

GetLabelIndexOk returns a tuple with the LabelIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIndex

`func (o *VersionedConnection) SetLabelIndex(v int32)`

SetLabelIndex sets LabelIndex field to given value.

### HasLabelIndex

`func (o *VersionedConnection) HasLabelIndex() bool`

HasLabelIndex returns a boolean if a field has been set.

### GetZIndex

`func (o *VersionedConnection) GetZIndex() int64`

GetZIndex returns the ZIndex field if non-nil, zero value otherwise.

### GetZIndexOk

`func (o *VersionedConnection) GetZIndexOk() (*int64, bool)`

GetZIndexOk returns a tuple with the ZIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZIndex

`func (o *VersionedConnection) SetZIndex(v int64)`

SetZIndex sets ZIndex field to given value.

### HasZIndex

`func (o *VersionedConnection) HasZIndex() bool`

HasZIndex returns a boolean if a field has been set.

### GetSelectedRelationships

`func (o *VersionedConnection) GetSelectedRelationships() []string`

GetSelectedRelationships returns the SelectedRelationships field if non-nil, zero value otherwise.

### GetSelectedRelationshipsOk

`func (o *VersionedConnection) GetSelectedRelationshipsOk() (*[]string, bool)`

GetSelectedRelationshipsOk returns a tuple with the SelectedRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRelationships

`func (o *VersionedConnection) SetSelectedRelationships(v []string)`

SetSelectedRelationships sets SelectedRelationships field to given value.

### HasSelectedRelationships

`func (o *VersionedConnection) HasSelectedRelationships() bool`

HasSelectedRelationships returns a boolean if a field has been set.

### GetBackPressureObjectThreshold

`func (o *VersionedConnection) GetBackPressureObjectThreshold() int64`

GetBackPressureObjectThreshold returns the BackPressureObjectThreshold field if non-nil, zero value otherwise.

### GetBackPressureObjectThresholdOk

`func (o *VersionedConnection) GetBackPressureObjectThresholdOk() (*int64, bool)`

GetBackPressureObjectThresholdOk returns a tuple with the BackPressureObjectThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackPressureObjectThreshold

`func (o *VersionedConnection) SetBackPressureObjectThreshold(v int64)`

SetBackPressureObjectThreshold sets BackPressureObjectThreshold field to given value.

### HasBackPressureObjectThreshold

`func (o *VersionedConnection) HasBackPressureObjectThreshold() bool`

HasBackPressureObjectThreshold returns a boolean if a field has been set.

### GetBackPressureDataSizeThreshold

`func (o *VersionedConnection) GetBackPressureDataSizeThreshold() string`

GetBackPressureDataSizeThreshold returns the BackPressureDataSizeThreshold field if non-nil, zero value otherwise.

### GetBackPressureDataSizeThresholdOk

`func (o *VersionedConnection) GetBackPressureDataSizeThresholdOk() (*string, bool)`

GetBackPressureDataSizeThresholdOk returns a tuple with the BackPressureDataSizeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackPressureDataSizeThreshold

`func (o *VersionedConnection) SetBackPressureDataSizeThreshold(v string)`

SetBackPressureDataSizeThreshold sets BackPressureDataSizeThreshold field to given value.

### HasBackPressureDataSizeThreshold

`func (o *VersionedConnection) HasBackPressureDataSizeThreshold() bool`

HasBackPressureDataSizeThreshold returns a boolean if a field has been set.

### GetFlowFileExpiration

`func (o *VersionedConnection) GetFlowFileExpiration() string`

GetFlowFileExpiration returns the FlowFileExpiration field if non-nil, zero value otherwise.

### GetFlowFileExpirationOk

`func (o *VersionedConnection) GetFlowFileExpirationOk() (*string, bool)`

GetFlowFileExpirationOk returns a tuple with the FlowFileExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFileExpiration

`func (o *VersionedConnection) SetFlowFileExpiration(v string)`

SetFlowFileExpiration sets FlowFileExpiration field to given value.

### HasFlowFileExpiration

`func (o *VersionedConnection) HasFlowFileExpiration() bool`

HasFlowFileExpiration returns a boolean if a field has been set.

### GetPrioritizers

`func (o *VersionedConnection) GetPrioritizers() []string`

GetPrioritizers returns the Prioritizers field if non-nil, zero value otherwise.

### GetPrioritizersOk

`func (o *VersionedConnection) GetPrioritizersOk() (*[]string, bool)`

GetPrioritizersOk returns a tuple with the Prioritizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrioritizers

`func (o *VersionedConnection) SetPrioritizers(v []string)`

SetPrioritizers sets Prioritizers field to given value.

### HasPrioritizers

`func (o *VersionedConnection) HasPrioritizers() bool`

HasPrioritizers returns a boolean if a field has been set.

### GetBends

`func (o *VersionedConnection) GetBends() []Position`

GetBends returns the Bends field if non-nil, zero value otherwise.

### GetBendsOk

`func (o *VersionedConnection) GetBendsOk() (*[]Position, bool)`

GetBendsOk returns a tuple with the Bends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBends

`func (o *VersionedConnection) SetBends(v []Position)`

SetBends sets Bends field to given value.

### HasBends

`func (o *VersionedConnection) HasBends() bool`

HasBends returns a boolean if a field has been set.

### GetLoadBalanceStrategy

`func (o *VersionedConnection) GetLoadBalanceStrategy() string`

GetLoadBalanceStrategy returns the LoadBalanceStrategy field if non-nil, zero value otherwise.

### GetLoadBalanceStrategyOk

`func (o *VersionedConnection) GetLoadBalanceStrategyOk() (*string, bool)`

GetLoadBalanceStrategyOk returns a tuple with the LoadBalanceStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceStrategy

`func (o *VersionedConnection) SetLoadBalanceStrategy(v string)`

SetLoadBalanceStrategy sets LoadBalanceStrategy field to given value.

### HasLoadBalanceStrategy

`func (o *VersionedConnection) HasLoadBalanceStrategy() bool`

HasLoadBalanceStrategy returns a boolean if a field has been set.

### GetPartitioningAttribute

`func (o *VersionedConnection) GetPartitioningAttribute() string`

GetPartitioningAttribute returns the PartitioningAttribute field if non-nil, zero value otherwise.

### GetPartitioningAttributeOk

`func (o *VersionedConnection) GetPartitioningAttributeOk() (*string, bool)`

GetPartitioningAttributeOk returns a tuple with the PartitioningAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitioningAttribute

`func (o *VersionedConnection) SetPartitioningAttribute(v string)`

SetPartitioningAttribute sets PartitioningAttribute field to given value.

### HasPartitioningAttribute

`func (o *VersionedConnection) HasPartitioningAttribute() bool`

HasPartitioningAttribute returns a boolean if a field has been set.

### GetLoadBalanceCompression

`func (o *VersionedConnection) GetLoadBalanceCompression() string`

GetLoadBalanceCompression returns the LoadBalanceCompression field if non-nil, zero value otherwise.

### GetLoadBalanceCompressionOk

`func (o *VersionedConnection) GetLoadBalanceCompressionOk() (*string, bool)`

GetLoadBalanceCompressionOk returns a tuple with the LoadBalanceCompression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceCompression

`func (o *VersionedConnection) SetLoadBalanceCompression(v string)`

SetLoadBalanceCompression sets LoadBalanceCompression field to given value.

### HasLoadBalanceCompression

`func (o *VersionedConnection) HasLoadBalanceCompression() bool`

HasLoadBalanceCompression returns a boolean if a field has been set.

### GetComponentType

`func (o *VersionedConnection) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *VersionedConnection) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *VersionedConnection) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *VersionedConnection) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### GetGroupIdentifier

`func (o *VersionedConnection) GetGroupIdentifier() string`

GetGroupIdentifier returns the GroupIdentifier field if non-nil, zero value otherwise.

### GetGroupIdentifierOk

`func (o *VersionedConnection) GetGroupIdentifierOk() (*string, bool)`

GetGroupIdentifierOk returns a tuple with the GroupIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIdentifier

`func (o *VersionedConnection) SetGroupIdentifier(v string)`

SetGroupIdentifier sets GroupIdentifier field to given value.

### HasGroupIdentifier

`func (o *VersionedConnection) HasGroupIdentifier() bool`

HasGroupIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


