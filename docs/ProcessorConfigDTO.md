# ProcessorConfigDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | Pointer to **map[string]string** | The properties for the processor. Properties whose value is not set will only contain the property name. | [optional] 
**Descriptors** | Pointer to [**map[string]PropertyDescriptorDTO**](PropertyDescriptorDTO.md) | Descriptors for the processor&#39;s properties. | [optional] 
**SchedulingPeriod** | Pointer to **string** | The frequency with which to schedule the processor. The format of the value will depend on th value of schedulingStrategy. | [optional] 
**SchedulingStrategy** | Pointer to **string** | Indcates whether the prcessor should be scheduled to run in event or timer driven mode. | [optional] 
**ExecutionNode** | Pointer to **string** | Indicates the node where the process will execute. | [optional] 
**PenaltyDuration** | Pointer to **string** | The amount of time that is used when the process penalizes a flowfile. | [optional] 
**YieldDuration** | Pointer to **string** | The amount of time that must elapse before this processor is scheduled again after yielding. | [optional] 
**BulletinLevel** | Pointer to **string** | The level at which the processor will report bulletins. | [optional] 
**RunDurationMillis** | Pointer to **int64** | The run duration for the processor in milliseconds. | [optional] 
**ConcurrentlySchedulableTaskCount** | Pointer to **int32** | The number of tasks that should be concurrently schedule for the processor. If the processor doesn&#39;t allow parallol processing then any positive input will be ignored. | [optional] 
**AutoTerminatedRelationships** | Pointer to **[]string** | The names of all relationships that cause a flow file to be terminated if the relationship is not connected elsewhere. This property differs from the &#39;isAutoTerminate&#39; property of the RelationshipDTO in that the RelationshipDTO is meant to depict the current configuration, whereas this property can be set in a DTO when updating a Processor in order to change which Relationships should be auto-terminated. | [optional] 
**Comments** | Pointer to **string** | The comments for the processor. | [optional] 
**CustomUiUrl** | Pointer to **string** | The URL for the processor&#39;s custom configuration UI if applicable. | [optional] 
**LossTolerant** | Pointer to **bool** | Whether the processor is loss tolerant. | [optional] 
**AnnotationData** | Pointer to **string** | The annotation data for the processor used to relay configuration between a custom UI and the procesosr. | [optional] 
**DefaultConcurrentTasks** | Pointer to **map[string]string** | Maps default values for concurrent tasks for each applicable scheduling strategy. | [optional] 
**DefaultSchedulingPeriod** | Pointer to **map[string]string** | Maps default values for scheduling period for each applicable scheduling strategy. | [optional] 

## Methods

### NewProcessorConfigDTO

`func NewProcessorConfigDTO() *ProcessorConfigDTO`

NewProcessorConfigDTO instantiates a new ProcessorConfigDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorConfigDTOWithDefaults

`func NewProcessorConfigDTOWithDefaults() *ProcessorConfigDTO`

NewProcessorConfigDTOWithDefaults instantiates a new ProcessorConfigDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *ProcessorConfigDTO) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ProcessorConfigDTO) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ProcessorConfigDTO) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ProcessorConfigDTO) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetDescriptors

`func (o *ProcessorConfigDTO) GetDescriptors() map[string]PropertyDescriptorDTO`

GetDescriptors returns the Descriptors field if non-nil, zero value otherwise.

### GetDescriptorsOk

`func (o *ProcessorConfigDTO) GetDescriptorsOk() (*map[string]PropertyDescriptorDTO, bool)`

GetDescriptorsOk returns a tuple with the Descriptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptors

`func (o *ProcessorConfigDTO) SetDescriptors(v map[string]PropertyDescriptorDTO)`

SetDescriptors sets Descriptors field to given value.

### HasDescriptors

`func (o *ProcessorConfigDTO) HasDescriptors() bool`

HasDescriptors returns a boolean if a field has been set.

### GetSchedulingPeriod

`func (o *ProcessorConfigDTO) GetSchedulingPeriod() string`

GetSchedulingPeriod returns the SchedulingPeriod field if non-nil, zero value otherwise.

### GetSchedulingPeriodOk

`func (o *ProcessorConfigDTO) GetSchedulingPeriodOk() (*string, bool)`

GetSchedulingPeriodOk returns a tuple with the SchedulingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingPeriod

`func (o *ProcessorConfigDTO) SetSchedulingPeriod(v string)`

SetSchedulingPeriod sets SchedulingPeriod field to given value.

### HasSchedulingPeriod

`func (o *ProcessorConfigDTO) HasSchedulingPeriod() bool`

HasSchedulingPeriod returns a boolean if a field has been set.

### GetSchedulingStrategy

`func (o *ProcessorConfigDTO) GetSchedulingStrategy() string`

GetSchedulingStrategy returns the SchedulingStrategy field if non-nil, zero value otherwise.

### GetSchedulingStrategyOk

`func (o *ProcessorConfigDTO) GetSchedulingStrategyOk() (*string, bool)`

GetSchedulingStrategyOk returns a tuple with the SchedulingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingStrategy

`func (o *ProcessorConfigDTO) SetSchedulingStrategy(v string)`

SetSchedulingStrategy sets SchedulingStrategy field to given value.

### HasSchedulingStrategy

`func (o *ProcessorConfigDTO) HasSchedulingStrategy() bool`

HasSchedulingStrategy returns a boolean if a field has been set.

### GetExecutionNode

`func (o *ProcessorConfigDTO) GetExecutionNode() string`

GetExecutionNode returns the ExecutionNode field if non-nil, zero value otherwise.

### GetExecutionNodeOk

`func (o *ProcessorConfigDTO) GetExecutionNodeOk() (*string, bool)`

GetExecutionNodeOk returns a tuple with the ExecutionNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionNode

`func (o *ProcessorConfigDTO) SetExecutionNode(v string)`

SetExecutionNode sets ExecutionNode field to given value.

### HasExecutionNode

`func (o *ProcessorConfigDTO) HasExecutionNode() bool`

HasExecutionNode returns a boolean if a field has been set.

### GetPenaltyDuration

`func (o *ProcessorConfigDTO) GetPenaltyDuration() string`

GetPenaltyDuration returns the PenaltyDuration field if non-nil, zero value otherwise.

### GetPenaltyDurationOk

`func (o *ProcessorConfigDTO) GetPenaltyDurationOk() (*string, bool)`

GetPenaltyDurationOk returns a tuple with the PenaltyDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPenaltyDuration

`func (o *ProcessorConfigDTO) SetPenaltyDuration(v string)`

SetPenaltyDuration sets PenaltyDuration field to given value.

### HasPenaltyDuration

`func (o *ProcessorConfigDTO) HasPenaltyDuration() bool`

HasPenaltyDuration returns a boolean if a field has been set.

### GetYieldDuration

`func (o *ProcessorConfigDTO) GetYieldDuration() string`

GetYieldDuration returns the YieldDuration field if non-nil, zero value otherwise.

### GetYieldDurationOk

`func (o *ProcessorConfigDTO) GetYieldDurationOk() (*string, bool)`

GetYieldDurationOk returns a tuple with the YieldDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYieldDuration

`func (o *ProcessorConfigDTO) SetYieldDuration(v string)`

SetYieldDuration sets YieldDuration field to given value.

### HasYieldDuration

`func (o *ProcessorConfigDTO) HasYieldDuration() bool`

HasYieldDuration returns a boolean if a field has been set.

### GetBulletinLevel

`func (o *ProcessorConfigDTO) GetBulletinLevel() string`

GetBulletinLevel returns the BulletinLevel field if non-nil, zero value otherwise.

### GetBulletinLevelOk

`func (o *ProcessorConfigDTO) GetBulletinLevelOk() (*string, bool)`

GetBulletinLevelOk returns a tuple with the BulletinLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletinLevel

`func (o *ProcessorConfigDTO) SetBulletinLevel(v string)`

SetBulletinLevel sets BulletinLevel field to given value.

### HasBulletinLevel

`func (o *ProcessorConfigDTO) HasBulletinLevel() bool`

HasBulletinLevel returns a boolean if a field has been set.

### GetRunDurationMillis

`func (o *ProcessorConfigDTO) GetRunDurationMillis() int64`

GetRunDurationMillis returns the RunDurationMillis field if non-nil, zero value otherwise.

### GetRunDurationMillisOk

`func (o *ProcessorConfigDTO) GetRunDurationMillisOk() (*int64, bool)`

GetRunDurationMillisOk returns a tuple with the RunDurationMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunDurationMillis

`func (o *ProcessorConfigDTO) SetRunDurationMillis(v int64)`

SetRunDurationMillis sets RunDurationMillis field to given value.

### HasRunDurationMillis

`func (o *ProcessorConfigDTO) HasRunDurationMillis() bool`

HasRunDurationMillis returns a boolean if a field has been set.

### GetConcurrentlySchedulableTaskCount

`func (o *ProcessorConfigDTO) GetConcurrentlySchedulableTaskCount() int32`

GetConcurrentlySchedulableTaskCount returns the ConcurrentlySchedulableTaskCount field if non-nil, zero value otherwise.

### GetConcurrentlySchedulableTaskCountOk

`func (o *ProcessorConfigDTO) GetConcurrentlySchedulableTaskCountOk() (*int32, bool)`

GetConcurrentlySchedulableTaskCountOk returns a tuple with the ConcurrentlySchedulableTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentlySchedulableTaskCount

`func (o *ProcessorConfigDTO) SetConcurrentlySchedulableTaskCount(v int32)`

SetConcurrentlySchedulableTaskCount sets ConcurrentlySchedulableTaskCount field to given value.

### HasConcurrentlySchedulableTaskCount

`func (o *ProcessorConfigDTO) HasConcurrentlySchedulableTaskCount() bool`

HasConcurrentlySchedulableTaskCount returns a boolean if a field has been set.

### GetAutoTerminatedRelationships

`func (o *ProcessorConfigDTO) GetAutoTerminatedRelationships() []string`

GetAutoTerminatedRelationships returns the AutoTerminatedRelationships field if non-nil, zero value otherwise.

### GetAutoTerminatedRelationshipsOk

`func (o *ProcessorConfigDTO) GetAutoTerminatedRelationshipsOk() (*[]string, bool)`

GetAutoTerminatedRelationshipsOk returns a tuple with the AutoTerminatedRelationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTerminatedRelationships

`func (o *ProcessorConfigDTO) SetAutoTerminatedRelationships(v []string)`

SetAutoTerminatedRelationships sets AutoTerminatedRelationships field to given value.

### HasAutoTerminatedRelationships

`func (o *ProcessorConfigDTO) HasAutoTerminatedRelationships() bool`

HasAutoTerminatedRelationships returns a boolean if a field has been set.

### GetComments

`func (o *ProcessorConfigDTO) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ProcessorConfigDTO) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ProcessorConfigDTO) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ProcessorConfigDTO) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCustomUiUrl

`func (o *ProcessorConfigDTO) GetCustomUiUrl() string`

GetCustomUiUrl returns the CustomUiUrl field if non-nil, zero value otherwise.

### GetCustomUiUrlOk

`func (o *ProcessorConfigDTO) GetCustomUiUrlOk() (*string, bool)`

GetCustomUiUrlOk returns a tuple with the CustomUiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUiUrl

`func (o *ProcessorConfigDTO) SetCustomUiUrl(v string)`

SetCustomUiUrl sets CustomUiUrl field to given value.

### HasCustomUiUrl

`func (o *ProcessorConfigDTO) HasCustomUiUrl() bool`

HasCustomUiUrl returns a boolean if a field has been set.

### GetLossTolerant

`func (o *ProcessorConfigDTO) GetLossTolerant() bool`

GetLossTolerant returns the LossTolerant field if non-nil, zero value otherwise.

### GetLossTolerantOk

`func (o *ProcessorConfigDTO) GetLossTolerantOk() (*bool, bool)`

GetLossTolerantOk returns a tuple with the LossTolerant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossTolerant

`func (o *ProcessorConfigDTO) SetLossTolerant(v bool)`

SetLossTolerant sets LossTolerant field to given value.

### HasLossTolerant

`func (o *ProcessorConfigDTO) HasLossTolerant() bool`

HasLossTolerant returns a boolean if a field has been set.

### GetAnnotationData

`func (o *ProcessorConfigDTO) GetAnnotationData() string`

GetAnnotationData returns the AnnotationData field if non-nil, zero value otherwise.

### GetAnnotationDataOk

`func (o *ProcessorConfigDTO) GetAnnotationDataOk() (*string, bool)`

GetAnnotationDataOk returns a tuple with the AnnotationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationData

`func (o *ProcessorConfigDTO) SetAnnotationData(v string)`

SetAnnotationData sets AnnotationData field to given value.

### HasAnnotationData

`func (o *ProcessorConfigDTO) HasAnnotationData() bool`

HasAnnotationData returns a boolean if a field has been set.

### GetDefaultConcurrentTasks

`func (o *ProcessorConfigDTO) GetDefaultConcurrentTasks() map[string]string`

GetDefaultConcurrentTasks returns the DefaultConcurrentTasks field if non-nil, zero value otherwise.

### GetDefaultConcurrentTasksOk

`func (o *ProcessorConfigDTO) GetDefaultConcurrentTasksOk() (*map[string]string, bool)`

GetDefaultConcurrentTasksOk returns a tuple with the DefaultConcurrentTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConcurrentTasks

`func (o *ProcessorConfigDTO) SetDefaultConcurrentTasks(v map[string]string)`

SetDefaultConcurrentTasks sets DefaultConcurrentTasks field to given value.

### HasDefaultConcurrentTasks

`func (o *ProcessorConfigDTO) HasDefaultConcurrentTasks() bool`

HasDefaultConcurrentTasks returns a boolean if a field has been set.

### GetDefaultSchedulingPeriod

`func (o *ProcessorConfigDTO) GetDefaultSchedulingPeriod() map[string]string`

GetDefaultSchedulingPeriod returns the DefaultSchedulingPeriod field if non-nil, zero value otherwise.

### GetDefaultSchedulingPeriodOk

`func (o *ProcessorConfigDTO) GetDefaultSchedulingPeriodOk() (*map[string]string, bool)`

GetDefaultSchedulingPeriodOk returns a tuple with the DefaultSchedulingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSchedulingPeriod

`func (o *ProcessorConfigDTO) SetDefaultSchedulingPeriod(v map[string]string)`

SetDefaultSchedulingPeriod sets DefaultSchedulingPeriod field to given value.

### HasDefaultSchedulingPeriod

`func (o *ProcessorConfigDTO) HasDefaultSchedulingPeriod() bool`

HasDefaultSchedulingPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


