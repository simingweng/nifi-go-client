# VersionedProcessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | The component&#39;s unique identifier | [optional] 
**Name** | Pointer to **string** | The component&#39;s name | [optional] 
**Comments** | Pointer to **string** | The user-supplied comments for the component | [optional] 
**Position** | Pointer to [**Position**](Position.md) |  | [optional] 
**Bundle** | Pointer to [**Bundle**](Bundle.md) |  | [optional] 
**Style** | Pointer to **map[string]string** | Stylistic data for rendering in a UI | [optional] 
**Type** | Pointer to **string** | The type of Processor | [optional] 
**Properties** | Pointer to **map[string]string** | The properties for the processor. Properties whose value is not set will only contain the property name. | [optional] 
**PropertyDescriptors** | Pointer to [**map[string]VersionedPropertyDescriptor**](VersionedPropertyDescriptor.md) | The property descriptors for the processor. | [optional] 
**AnnotationData** | Pointer to **string** | The annotation data for the processor used to relay configuration between a custom UI and the procesosr. | [optional] 
**SchedulingPeriod** | Pointer to **string** | The frequency with which to schedule the processor. The format of the value will depend on th value of schedulingStrategy. | [optional] 
**SchedulingStrategy** | Pointer to **string** | Indcates whether the prcessor should be scheduled to run in event or timer driven mode. | [optional] 
**ExecutionNode** | Pointer to **string** | Indicates the node where the process will execute. | [optional] 
**PenaltyDuration** | Pointer to **string** | The amout of time that is used when the process penalizes a flowfile. | [optional] 
**YieldDuration** | Pointer to **string** | The amount of time that must elapse before this processor is scheduled again after yielding. | [optional] 
**BulletinLevel** | Pointer to **string** | The level at which the processor will report bulletins. | [optional] 
**RunDurationMillis** | Pointer to **int64** | The run duration for the processor in milliseconds. | [optional] 
**ConcurrentlySchedulableTaskCount** | Pointer to **int32** | The number of tasks that should be concurrently schedule for the processor. If the processor doesn&#39;t allow parallol processing then any positive input will be ignored. | [optional] 
**AutoTerminatedRelationships** | Pointer to **[]string** | The names of all relationships that cause a flow file to be terminated if the relationship is not connected elsewhere. This property differs from the &#39;isAutoTerminate&#39; property of the RelationshipDTO in that the RelationshipDTO is meant to depict the current configuration, whereas this property can be set in a DTO when updating a Processor in order to change which Relationships should be auto-terminated. | [optional] 
**ScheduledState** | Pointer to **string** | The scheduled state of the component | [optional] 
**ComponentType** | Pointer to **string** |  | [optional] 
**GroupIdentifier** | Pointer to **string** | The ID of the Process Group that this component belongs to | [optional] 

## Methods

### NewVersionedProcessor

`func NewVersionedProcessor() *VersionedProcessor`

NewVersionedProcessor instantiates a new VersionedProcessor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionedProcessorWithDefaults

`func NewVersionedProcessorWithDefaults() *VersionedProcessor`

NewVersionedProcessorWithDefaults instantiates a new VersionedProcessor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *VersionedProcessor) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *VersionedProcessor) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *VersionedProcessor) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *VersionedProcessor) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *VersionedProcessor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VersionedProcessor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VersionedProcessor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VersionedProcessor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComments

`func (o *VersionedProcessor) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VersionedProcessor) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VersionedProcessor) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VersionedProcessor) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetPosition

`func (o *VersionedProcessor) GetPosition() Position`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *VersionedProcessor) GetPositionOk() (*Position, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *VersionedProcessor) SetPosition(v Position)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *VersionedProcessor) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetBundle

`func (o *VersionedProcessor) GetBundle() Bundle`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *VersionedProcessor) GetBundleOk() (*Bundle, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *VersionedProcessor) SetBundle(v Bundle)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *VersionedProcessor) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### GetStyle

`func (o *VersionedProcessor) GetStyle() map[string]string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *VersionedProcessor) GetStyleOk() (*map[string]string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *VersionedProcessor) SetStyle(v map[string]string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *VersionedProcessor) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### GetType

`func (o *VersionedProcessor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VersionedProcessor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VersionedProcessor) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VersionedProcessor) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProperties

`func (o *VersionedProcessor) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *VersionedProcessor) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *VersionedProcessor) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *VersionedProcessor) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetPropertyDescriptors

`func (o *VersionedProcessor) GetPropertyDescriptors() map[string]VersionedPropertyDescriptor`

GetPropertyDescriptors returns the PropertyDescriptors field if non-nil, zero value otherwise.

### GetPropertyDescriptorsOk

`func (o *VersionedProcessor) GetPropertyDescriptorsOk() (*map[string]VersionedPropertyDescriptor, bool)`

GetPropertyDescriptorsOk returns a tuple with the PropertyDescriptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyDescriptors

`func (o *VersionedProcessor) SetPropertyDescriptors(v map[string]VersionedPropertyDescriptor)`

SetPropertyDescriptors sets PropertyDescriptors field to given value.

### HasPropertyDescriptors

`func (o *VersionedProcessor) HasPropertyDescriptors() bool`

HasPropertyDescriptors returns a boolean if a field has been set.

### GetAnnotationData

`func (o *VersionedProcessor) GetAnnotationData() string`

GetAnnotationData returns the AnnotationData field if non-nil, zero value otherwise.

### GetAnnotationDataOk

`func (o *VersionedProcessor) GetAnnotationDataOk() (*string, bool)`

GetAnnotationDataOk returns a tuple with the AnnotationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationData

`func (o *VersionedProcessor) SetAnnotationData(v string)`

SetAnnotationData sets AnnotationData field to given value.

### HasAnnotationData

`func (o *VersionedProcessor) HasAnnotationData() bool`

HasAnnotationData returns a boolean if a field has been set.

### GetSchedulingPeriod

`func (o *VersionedProcessor) GetSchedulingPeriod() string`

GetSchedulingPeriod returns the SchedulingPeriod field if non-nil, zero value otherwise.

### GetSchedulingPeriodOk

`func (o *VersionedProcessor) GetSchedulingPeriodOk() (*string, bool)`

GetSchedulingPeriodOk returns a tuple with the SchedulingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingPeriod

`func (o *VersionedProcessor) SetSchedulingPeriod(v string)`

SetSchedulingPeriod sets SchedulingPeriod field to given value.

### HasSchedulingPeriod

`func (o *VersionedProcessor) HasSchedulingPeriod() bool`

HasSchedulingPeriod returns a boolean if a field has been set.

### GetSchedulingStrategy

`func (o *VersionedProcessor) GetSchedulingStrategy() string`

GetSchedulingStrategy returns the SchedulingStrategy field if non-nil, zero value otherwise.

### GetSchedulingStrategyOk

`func (o *VersionedProcessor) GetSchedulingStrategyOk() (*string, bool)`

GetSchedulingStrategyOk returns a tuple with the SchedulingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingStrategy

`func (o *VersionedProcessor) SetSchedulingStrategy(v string)`

SetSchedulingStrategy sets SchedulingStrategy field to given value.

### HasSchedulingStrategy

`func (o *VersionedProcessor) HasSchedulingStrategy() bool`

HasSchedulingStrategy returns a boolean if a field has been set.

### GetExecutionNode

`func (o *VersionedProcessor) GetExecutionNode() string`

GetExecutionNode returns the ExecutionNode field if non-nil, zero value otherwise.

### GetExecutionNodeOk

`func (o *VersionedProcessor) GetExecutionNodeOk() (*string, bool)`

GetExecutionNodeOk returns a tuple with the ExecutionNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionNode

`func (o *VersionedProcessor) SetExecutionNode(v string)`

SetExecutionNode sets ExecutionNode field to given value.

### HasExecutionNode

`func (o *VersionedProcessor) HasExecutionNode() bool`

HasExecutionNode returns a boolean if a field has been set.

### GetPenaltyDuration

`func (o *VersionedProcessor) GetPenaltyDuration() string`

GetPenaltyDuration returns the PenaltyDuration field if non-nil, zero value otherwise.

### GetPenaltyDurationOk

`func (o *VersionedProcessor) GetPenaltyDurationOk() (*string, bool)`

GetPenaltyDurationOk returns a tuple with the PenaltyDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPenaltyDuration

`func (o *VersionedProcessor) SetPenaltyDuration(v string)`

SetPenaltyDuration sets PenaltyDuration field to given value.

### HasPenaltyDuration

`func (o *VersionedProcessor) HasPenaltyDuration() bool`

HasPenaltyDuration returns a boolean if a field has been set.

### GetYieldDuration

`func (o *VersionedProcessor) GetYieldDuration() string`

GetYieldDuration returns the YieldDuration field if non-nil, zero value otherwise.

### GetYieldDurationOk

`func (o *VersionedProcessor) GetYieldDurationOk() (*string, bool)`

GetYieldDurationOk returns a tuple with the YieldDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYieldDuration

`func (o *VersionedProcessor) SetYieldDuration(v string)`

SetYieldDuration sets YieldDuration field to given value.

### HasYieldDuration

`func (o *VersionedProcessor) HasYieldDuration() bool`

HasYieldDuration returns a boolean if a field has been set.

### GetBulletinLevel

`func (o *VersionedProcessor) GetBulletinLevel() string`

GetBulletinLevel returns the BulletinLevel field if non-nil, zero value otherwise.

### GetBulletinLevelOk

`func (o *VersionedProcessor) GetBulletinLevelOk() (*string, bool)`

GetBulletinLevelOk returns a tuple with the BulletinLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletinLevel

`func (o *VersionedProcessor) SetBulletinLevel(v string)`

SetBulletinLevel sets BulletinLevel field to given value.

### HasBulletinLevel

`func (o *VersionedProcessor) HasBulletinLevel() bool`

HasBulletinLevel returns a boolean if a field has been set.

### GetRunDurationMillis

`func (o *VersionedProcessor) GetRunDurationMillis() int64`

GetRunDurationMillis returns the RunDurationMillis field if non-nil, zero value otherwise.

### GetRunDurationMillisOk

`func (o *VersionedProcessor) GetRunDurationMillisOk() (*int64, bool)`

GetRunDurationMillisOk returns a tuple with the RunDurationMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunDurationMillis

`func (o *VersionedProcessor) SetRunDurationMillis(v int64)`

SetRunDurationMillis sets RunDurationMillis field to given value.

### HasRunDurationMillis

`func (o *VersionedProcessor) HasRunDurationMillis() bool`

HasRunDurationMillis returns a boolean if a field has been set.

### GetConcurrentlySchedulableTaskCount

`func (o *VersionedProcessor) GetConcurrentlySchedulableTaskCount() int32`

GetConcurrentlySchedulableTaskCount returns the ConcurrentlySchedulableTaskCount field if non-nil, zero value otherwise.

### GetConcurrentlySchedulableTaskCountOk

`func (o *VersionedProcessor) GetConcurrentlySchedulableTaskCountOk() (*int32, bool)`

GetConcurrentlySchedulableTaskCountOk returns a tuple with the ConcurrentlySchedulableTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentlySchedulableTaskCount

`func (o *VersionedProcessor) SetConcurrentlySchedulableTaskCount(v int32)`

SetConcurrentlySchedulableTaskCount sets ConcurrentlySchedulableTaskCount field to given value.

### HasConcurrentlySchedulableTaskCount

`func (o *VersionedProcessor) HasConcurrentlySchedulableTaskCount() bool`

HasConcurrentlySchedulableTaskCount returns a boolean if a field has been set.

### GetAutoTerminatedRelationships

`func (o *VersionedProcessor) GetAutoTerminatedRelationships() []string`

GetAutoTerminatedRelationships returns the AutoTerminatedRelationships field if non-nil, zero value otherwise.

### GetAutoTerminatedRelationshipsOk

`func (o *VersionedProcessor) GetAutoTerminatedRelationshipsOk() (*[]string, bool)`

GetAutoTerminatedRelationshipsOk returns a tuple with the AutoTerminatedRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTerminatedRelationships

`func (o *VersionedProcessor) SetAutoTerminatedRelationships(v []string)`

SetAutoTerminatedRelationships sets AutoTerminatedRelationships field to given value.

### HasAutoTerminatedRelationships

`func (o *VersionedProcessor) HasAutoTerminatedRelationships() bool`

HasAutoTerminatedRelationships returns a boolean if a field has been set.

### GetScheduledState

`func (o *VersionedProcessor) GetScheduledState() string`

GetScheduledState returns the ScheduledState field if non-nil, zero value otherwise.

### GetScheduledStateOk

`func (o *VersionedProcessor) GetScheduledStateOk() (*string, bool)`

GetScheduledStateOk returns a tuple with the ScheduledState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledState

`func (o *VersionedProcessor) SetScheduledState(v string)`

SetScheduledState sets ScheduledState field to given value.

### HasScheduledState

`func (o *VersionedProcessor) HasScheduledState() bool`

HasScheduledState returns a boolean if a field has been set.

### GetComponentType

`func (o *VersionedProcessor) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *VersionedProcessor) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *VersionedProcessor) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *VersionedProcessor) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### GetGroupIdentifier

`func (o *VersionedProcessor) GetGroupIdentifier() string`

GetGroupIdentifier returns the GroupIdentifier field if non-nil, zero value otherwise.

### GetGroupIdentifierOk

`func (o *VersionedProcessor) GetGroupIdentifierOk() (*string, bool)`

GetGroupIdentifierOk returns a tuple with the GroupIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIdentifier

`func (o *VersionedProcessor) SetGroupIdentifier(v string)`

SetGroupIdentifier sets GroupIdentifier field to given value.

### HasGroupIdentifier

`func (o *VersionedProcessor) HasGroupIdentifier() bool`

HasGroupIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


