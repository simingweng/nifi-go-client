# VersionedFlowSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotMetadata** | [**VersionedFlowSnapshotMetadata**](VersionedFlowSnapshotMetadata.md) |  | 
**FlowContents** | [**VersionedProcessGroup**](VersionedProcessGroup.md) |  | 
**ExternalControllerServices** | Pointer to [**map[string]ExternalControllerServiceReference**](ExternalControllerServiceReference.md) | The information about controller services that exist outside this versioned flow, but are referenced by components within the versioned flow. | [optional] 
**ParameterContexts** | Pointer to [**map[string]VersionedParameterContext**](VersionedParameterContext.md) | The parameter contexts referenced by process groups in the flow contents. The mapping is from the name of the context to the context instance, and it is expected that any context in this map is referenced by at least one process group in this flow. | [optional] 
**FlowEncodingVersion** | Pointer to **string** | The optional encoding version of the flow contents. | [optional] 
**Flow** | Pointer to [**VersionedFlow**](VersionedFlow.md) |  | [optional] 
**Bucket** | Pointer to [**Bucket**](Bucket.md) |  | [optional] 
**Latest** | Pointer to **bool** |  | [optional] 

## Methods

### NewVersionedFlowSnapshot

`func NewVersionedFlowSnapshot(snapshotMetadata VersionedFlowSnapshotMetadata, flowContents VersionedProcessGroup, ) *VersionedFlowSnapshot`

NewVersionedFlowSnapshot instantiates a new VersionedFlowSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionedFlowSnapshotWithDefaults

`func NewVersionedFlowSnapshotWithDefaults() *VersionedFlowSnapshot`

NewVersionedFlowSnapshotWithDefaults instantiates a new VersionedFlowSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotMetadata

`func (o *VersionedFlowSnapshot) GetSnapshotMetadata() VersionedFlowSnapshotMetadata`

GetSnapshotMetadata returns the SnapshotMetadata field if non-nil, zero value otherwise.

### GetSnapshotMetadataOk

`func (o *VersionedFlowSnapshot) GetSnapshotMetadataOk() (*VersionedFlowSnapshotMetadata, bool)`

GetSnapshotMetadataOk returns a tuple with the SnapshotMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotMetadata

`func (o *VersionedFlowSnapshot) SetSnapshotMetadata(v VersionedFlowSnapshotMetadata)`

SetSnapshotMetadata sets SnapshotMetadata field to given value.


### GetFlowContents

`func (o *VersionedFlowSnapshot) GetFlowContents() VersionedProcessGroup`

GetFlowContents returns the FlowContents field if non-nil, zero value otherwise.

### GetFlowContentsOk

`func (o *VersionedFlowSnapshot) GetFlowContentsOk() (*VersionedProcessGroup, bool)`

GetFlowContentsOk returns a tuple with the FlowContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowContents

`func (o *VersionedFlowSnapshot) SetFlowContents(v VersionedProcessGroup)`

SetFlowContents sets FlowContents field to given value.


### GetExternalControllerServices

`func (o *VersionedFlowSnapshot) GetExternalControllerServices() map[string]ExternalControllerServiceReference`

GetExternalControllerServices returns the ExternalControllerServices field if non-nil, zero value otherwise.

### GetExternalControllerServicesOk

`func (o *VersionedFlowSnapshot) GetExternalControllerServicesOk() (*map[string]ExternalControllerServiceReference, bool)`

GetExternalControllerServicesOk returns a tuple with the ExternalControllerServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalControllerServices

`func (o *VersionedFlowSnapshot) SetExternalControllerServices(v map[string]ExternalControllerServiceReference)`

SetExternalControllerServices sets ExternalControllerServices field to given value.

### HasExternalControllerServices

`func (o *VersionedFlowSnapshot) HasExternalControllerServices() bool`

HasExternalControllerServices returns a boolean if a field has been set.

### GetParameterContexts

`func (o *VersionedFlowSnapshot) GetParameterContexts() map[string]VersionedParameterContext`

GetParameterContexts returns the ParameterContexts field if non-nil, zero value otherwise.

### GetParameterContextsOk

`func (o *VersionedFlowSnapshot) GetParameterContextsOk() (*map[string]VersionedParameterContext, bool)`

GetParameterContextsOk returns a tuple with the ParameterContexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterContexts

`func (o *VersionedFlowSnapshot) SetParameterContexts(v map[string]VersionedParameterContext)`

SetParameterContexts sets ParameterContexts field to given value.

### HasParameterContexts

`func (o *VersionedFlowSnapshot) HasParameterContexts() bool`

HasParameterContexts returns a boolean if a field has been set.

### GetFlowEncodingVersion

`func (o *VersionedFlowSnapshot) GetFlowEncodingVersion() string`

GetFlowEncodingVersion returns the FlowEncodingVersion field if non-nil, zero value otherwise.

### GetFlowEncodingVersionOk

`func (o *VersionedFlowSnapshot) GetFlowEncodingVersionOk() (*string, bool)`

GetFlowEncodingVersionOk returns a tuple with the FlowEncodingVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowEncodingVersion

`func (o *VersionedFlowSnapshot) SetFlowEncodingVersion(v string)`

SetFlowEncodingVersion sets FlowEncodingVersion field to given value.

### HasFlowEncodingVersion

`func (o *VersionedFlowSnapshot) HasFlowEncodingVersion() bool`

HasFlowEncodingVersion returns a boolean if a field has been set.

### GetFlow

`func (o *VersionedFlowSnapshot) GetFlow() VersionedFlow`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *VersionedFlowSnapshot) GetFlowOk() (*VersionedFlow, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *VersionedFlowSnapshot) SetFlow(v VersionedFlow)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *VersionedFlowSnapshot) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetBucket

`func (o *VersionedFlowSnapshot) GetBucket() Bucket`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *VersionedFlowSnapshot) GetBucketOk() (*Bucket, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *VersionedFlowSnapshot) SetBucket(v Bucket)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *VersionedFlowSnapshot) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetLatest

`func (o *VersionedFlowSnapshot) GetLatest() bool`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *VersionedFlowSnapshot) GetLatestOk() (*bool, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *VersionedFlowSnapshot) SetLatest(v bool)`

SetLatest sets Latest field to given value.

### HasLatest

`func (o *VersionedFlowSnapshot) HasLatest() bool`

HasLatest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


