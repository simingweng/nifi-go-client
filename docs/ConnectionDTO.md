# ConnectionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the component. | [optional] 
**VersionedComponentId** | Pointer to **string** | The ID of the corresponding component that is under version control | [optional] 
**ParentGroupId** | Pointer to **string** | The id of parent process group of this component if applicable. | [optional] 
**Position** | Pointer to [**PositionDTO**](PositionDTO.md) |  | [optional] 
**Source** | Pointer to [**ConnectableDTO**](ConnectableDTO.md) |  | [optional] 
**Destination** | Pointer to [**ConnectableDTO**](ConnectableDTO.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the connection. | [optional] 
**LabelIndex** | Pointer to **int32** | The index of the bend point where to place the connection label. | [optional] 
**GetzIndex** | Pointer to **int64** | The z index of the connection. | [optional] 
**SelectedRelationships** | Pointer to **[]string** | The selected relationship that comprise the connection. | [optional] 
**AvailableRelationships** | Pointer to **[]string** | The relationships that the source of the connection currently supports. | [optional] [readonly] 
**BackPressureObjectThreshold** | Pointer to **int64** | The object count threshold for determining when back pressure is applied. Updating this value is a passive change in the sense that it won&#39;t impact whether existing files over the limit are affected but it does help feeder processors to stop pushing too much into this work queue. | [optional] 
**BackPressureDataSizeThreshold** | Pointer to **string** | The object data size threshold for determining when back pressure is applied. Updating this value is a passive change in the sense that it won&#39;t impact whether existing files over the limit are affected but it does help feeder processors to stop pushing too much into this work queue. | [optional] 
**FlowFileExpiration** | Pointer to **string** | The amount of time a flow file may be in the flow before it will be automatically aged out of the flow. Once a flow file reaches this age it will be terminated from the flow the next time a processor attempts to start work on it. | [optional] 
**Prioritizers** | Pointer to **[]string** | The comparators used to prioritize the queue. | [optional] 
**Bends** | Pointer to [**[]PositionDTO**](PositionDTO.md) | The bend points on the connection. | [optional] 
**LoadBalanceStrategy** | Pointer to **string** | How to load balance the data in this Connection across the nodes in the cluster. | [optional] 
**LoadBalancePartitionAttribute** | Pointer to **string** | The FlowFile Attribute to use for determining which node a FlowFile will go to if the Load Balancing Strategy is set to PARTITION_BY_ATTRIBUTE | [optional] 
**LoadBalanceCompression** | Pointer to **string** | Whether or not data should be compressed when being transferred between nodes in the cluster. | [optional] 
**LoadBalanceStatus** | Pointer to **string** | The current status of the Connection&#39;s Load Balancing Activities. Status can indicate that Load Balancing is not configured for the connection, that Load Balancing is configured but inactive (not currently transferring data to another node), or that Load Balancing is configured and actively transferring data to another node. | [optional] [readonly] 

## Methods

### NewConnectionDTO

`func NewConnectionDTO() *ConnectionDTO`

NewConnectionDTO instantiates a new ConnectionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionDTOWithDefaults

`func NewConnectionDTOWithDefaults() *ConnectionDTO`

NewConnectionDTOWithDefaults instantiates a new ConnectionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectionDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectionDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectionDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectionDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersionedComponentId

`func (o *ConnectionDTO) GetVersionedComponentId() string`

GetVersionedComponentId returns the VersionedComponentId field if non-nil, zero value otherwise.

### GetVersionedComponentIdOk

`func (o *ConnectionDTO) GetVersionedComponentIdOk() (*string, bool)`

GetVersionedComponentIdOk returns a tuple with the VersionedComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedComponentId

`func (o *ConnectionDTO) SetVersionedComponentId(v string)`

SetVersionedComponentId sets VersionedComponentId field to given value.

### HasVersionedComponentId

`func (o *ConnectionDTO) HasVersionedComponentId() bool`

HasVersionedComponentId returns a boolean if a field has been set.

### GetParentGroupId

`func (o *ConnectionDTO) GetParentGroupId() string`

GetParentGroupId returns the ParentGroupId field if non-nil, zero value otherwise.

### GetParentGroupIdOk

`func (o *ConnectionDTO) GetParentGroupIdOk() (*string, bool)`

GetParentGroupIdOk returns a tuple with the ParentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroupId

`func (o *ConnectionDTO) SetParentGroupId(v string)`

SetParentGroupId sets ParentGroupId field to given value.

### HasParentGroupId

`func (o *ConnectionDTO) HasParentGroupId() bool`

HasParentGroupId returns a boolean if a field has been set.

### GetPosition

`func (o *ConnectionDTO) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ConnectionDTO) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ConnectionDTO) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ConnectionDTO) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetSource

`func (o *ConnectionDTO) GetSource() ConnectableDTO`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ConnectionDTO) GetSourceOk() (*ConnectableDTO, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ConnectionDTO) SetSource(v ConnectableDTO)`

SetSource sets Source field to given value.

### HasSource

`func (o *ConnectionDTO) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDestination

`func (o *ConnectionDTO) GetDestination() ConnectableDTO`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ConnectionDTO) GetDestinationOk() (*ConnectableDTO, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ConnectionDTO) SetDestination(v ConnectableDTO)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *ConnectionDTO) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetName

`func (o *ConnectionDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectionDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabelIndex

`func (o *ConnectionDTO) GetLabelIndex() int32`

GetLabelIndex returns the LabelIndex field if non-nil, zero value otherwise.

### GetLabelIndexOk

`func (o *ConnectionDTO) GetLabelIndexOk() (*int32, bool)`

GetLabelIndexOk returns a tuple with the LabelIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIndex

`func (o *ConnectionDTO) SetLabelIndex(v int32)`

SetLabelIndex sets LabelIndex field to given value.

### HasLabelIndex

`func (o *ConnectionDTO) HasLabelIndex() bool`

HasLabelIndex returns a boolean if a field has been set.

### GetGetzIndex

`func (o *ConnectionDTO) GetGetzIndex() int64`

GetGetzIndex returns the GetzIndex field if non-nil, zero value otherwise.

### GetGetzIndexOk

`func (o *ConnectionDTO) GetGetzIndexOk() (*int64, bool)`

GetGetzIndexOk returns a tuple with the GetzIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetzIndex

`func (o *ConnectionDTO) SetGetzIndex(v int64)`

SetGetzIndex sets GetzIndex field to given value.

### HasGetzIndex

`func (o *ConnectionDTO) HasGetzIndex() bool`

HasGetzIndex returns a boolean if a field has been set.

### GetSelectedRelationships

`func (o *ConnectionDTO) GetSelectedRelationships() []string`

GetSelectedRelationships returns the SelectedRelationships field if non-nil, zero value otherwise.

### GetSelectedRelationshipsOk

`func (o *ConnectionDTO) GetSelectedRelationshipsOk() (*[]string, bool)`

GetSelectedRelationshipsOk returns a tuple with the SelectedRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRelationships

`func (o *ConnectionDTO) SetSelectedRelationships(v []string)`

SetSelectedRelationships sets SelectedRelationships field to given value.

### HasSelectedRelationships

`func (o *ConnectionDTO) HasSelectedRelationships() bool`

HasSelectedRelationships returns a boolean if a field has been set.

### GetAvailableRelationships

`func (o *ConnectionDTO) GetAvailableRelationships() []string`

GetAvailableRelationships returns the AvailableRelationships field if non-nil, zero value otherwise.

### GetAvailableRelationshipsOk

`func (o *ConnectionDTO) GetAvailableRelationshipsOk() (*[]string, bool)`

GetAvailableRelationshipsOk returns a tuple with the AvailableRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableRelationships

`func (o *ConnectionDTO) SetAvailableRelationships(v []string)`

SetAvailableRelationships sets AvailableRelationships field to given value.

### HasAvailableRelationships

`func (o *ConnectionDTO) HasAvailableRelationships() bool`

HasAvailableRelationships returns a boolean if a field has been set.

### GetBackPressureObjectThreshold

`func (o *ConnectionDTO) GetBackPressureObjectThreshold() int64`

GetBackPressureObjectThreshold returns the BackPressureObjectThreshold field if non-nil, zero value otherwise.

### GetBackPressureObjectThresholdOk

`func (o *ConnectionDTO) GetBackPressureObjectThresholdOk() (*int64, bool)`

GetBackPressureObjectThresholdOk returns a tuple with the BackPressureObjectThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackPressureObjectThreshold

`func (o *ConnectionDTO) SetBackPressureObjectThreshold(v int64)`

SetBackPressureObjectThreshold sets BackPressureObjectThreshold field to given value.

### HasBackPressureObjectThreshold

`func (o *ConnectionDTO) HasBackPressureObjectThreshold() bool`

HasBackPressureObjectThreshold returns a boolean if a field has been set.

### GetBackPressureDataSizeThreshold

`func (o *ConnectionDTO) GetBackPressureDataSizeThreshold() string`

GetBackPressureDataSizeThreshold returns the BackPressureDataSizeThreshold field if non-nil, zero value otherwise.

### GetBackPressureDataSizeThresholdOk

`func (o *ConnectionDTO) GetBackPressureDataSizeThresholdOk() (*string, bool)`

GetBackPressureDataSizeThresholdOk returns a tuple with the BackPressureDataSizeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackPressureDataSizeThreshold

`func (o *ConnectionDTO) SetBackPressureDataSizeThreshold(v string)`

SetBackPressureDataSizeThreshold sets BackPressureDataSizeThreshold field to given value.

### HasBackPressureDataSizeThreshold

`func (o *ConnectionDTO) HasBackPressureDataSizeThreshold() bool`

HasBackPressureDataSizeThreshold returns a boolean if a field has been set.

### GetFlowFileExpiration

`func (o *ConnectionDTO) GetFlowFileExpiration() string`

GetFlowFileExpiration returns the FlowFileExpiration field if non-nil, zero value otherwise.

### GetFlowFileExpirationOk

`func (o *ConnectionDTO) GetFlowFileExpirationOk() (*string, bool)`

GetFlowFileExpirationOk returns a tuple with the FlowFileExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFileExpiration

`func (o *ConnectionDTO) SetFlowFileExpiration(v string)`

SetFlowFileExpiration sets FlowFileExpiration field to given value.

### HasFlowFileExpiration

`func (o *ConnectionDTO) HasFlowFileExpiration() bool`

HasFlowFileExpiration returns a boolean if a field has been set.

### GetPrioritizers

`func (o *ConnectionDTO) GetPrioritizers() []string`

GetPrioritizers returns the Prioritizers field if non-nil, zero value otherwise.

### GetPrioritizersOk

`func (o *ConnectionDTO) GetPrioritizersOk() (*[]string, bool)`

GetPrioritizersOk returns a tuple with the Prioritizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrioritizers

`func (o *ConnectionDTO) SetPrioritizers(v []string)`

SetPrioritizers sets Prioritizers field to given value.

### HasPrioritizers

`func (o *ConnectionDTO) HasPrioritizers() bool`

HasPrioritizers returns a boolean if a field has been set.

### GetBends

`func (o *ConnectionDTO) GetBends() []PositionDTO`

GetBends returns the Bends field if non-nil, zero value otherwise.

### GetBendsOk

`func (o *ConnectionDTO) GetBendsOk() (*[]PositionDTO, bool)`

GetBendsOk returns a tuple with the Bends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBends

`func (o *ConnectionDTO) SetBends(v []PositionDTO)`

SetBends sets Bends field to given value.

### HasBends

`func (o *ConnectionDTO) HasBends() bool`

HasBends returns a boolean if a field has been set.

### GetLoadBalanceStrategy

`func (o *ConnectionDTO) GetLoadBalanceStrategy() string`

GetLoadBalanceStrategy returns the LoadBalanceStrategy field if non-nil, zero value otherwise.

### GetLoadBalanceStrategyOk

`func (o *ConnectionDTO) GetLoadBalanceStrategyOk() (*string, bool)`

GetLoadBalanceStrategyOk returns a tuple with the LoadBalanceStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceStrategy

`func (o *ConnectionDTO) SetLoadBalanceStrategy(v string)`

SetLoadBalanceStrategy sets LoadBalanceStrategy field to given value.

### HasLoadBalanceStrategy

`func (o *ConnectionDTO) HasLoadBalanceStrategy() bool`

HasLoadBalanceStrategy returns a boolean if a field has been set.

### GetLoadBalancePartitionAttribute

`func (o *ConnectionDTO) GetLoadBalancePartitionAttribute() string`

GetLoadBalancePartitionAttribute returns the LoadBalancePartitionAttribute field if non-nil, zero value otherwise.

### GetLoadBalancePartitionAttributeOk

`func (o *ConnectionDTO) GetLoadBalancePartitionAttributeOk() (*string, bool)`

GetLoadBalancePartitionAttributeOk returns a tuple with the LoadBalancePartitionAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancePartitionAttribute

`func (o *ConnectionDTO) SetLoadBalancePartitionAttribute(v string)`

SetLoadBalancePartitionAttribute sets LoadBalancePartitionAttribute field to given value.

### HasLoadBalancePartitionAttribute

`func (o *ConnectionDTO) HasLoadBalancePartitionAttribute() bool`

HasLoadBalancePartitionAttribute returns a boolean if a field has been set.

### GetLoadBalanceCompression

`func (o *ConnectionDTO) GetLoadBalanceCompression() string`

GetLoadBalanceCompression returns the LoadBalanceCompression field if non-nil, zero value otherwise.

### GetLoadBalanceCompressionOk

`func (o *ConnectionDTO) GetLoadBalanceCompressionOk() (*string, bool)`

GetLoadBalanceCompressionOk returns a tuple with the LoadBalanceCompression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceCompression

`func (o *ConnectionDTO) SetLoadBalanceCompression(v string)`

SetLoadBalanceCompression sets LoadBalanceCompression field to given value.

### HasLoadBalanceCompression

`func (o *ConnectionDTO) HasLoadBalanceCompression() bool`

HasLoadBalanceCompression returns a boolean if a field has been set.

### GetLoadBalanceStatus

`func (o *ConnectionDTO) GetLoadBalanceStatus() string`

GetLoadBalanceStatus returns the LoadBalanceStatus field if non-nil, zero value otherwise.

### GetLoadBalanceStatusOk

`func (o *ConnectionDTO) GetLoadBalanceStatusOk() (*string, bool)`

GetLoadBalanceStatusOk returns a tuple with the LoadBalanceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceStatus

`func (o *ConnectionDTO) SetLoadBalanceStatus(v string)`

SetLoadBalanceStatus sets LoadBalanceStatus field to given value.

### HasLoadBalanceStatus

`func (o *ConnectionDTO) HasLoadBalanceStatus() bool`

HasLoadBalanceStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


