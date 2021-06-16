# ProvenanceEventDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The event uuid. | [optional] 
**EventId** | Pointer to **int64** | The event id. This is a one up number thats unique per node. | [optional] 
**EventTime** | Pointer to **string** | The timestamp of the event. | [optional] 
**EventDuration** | Pointer to **int64** | The event duration in milliseconds. | [optional] 
**LineageDuration** | Pointer to **int64** | The duration since the lineage began, in milliseconds. | [optional] 
**EventType** | Pointer to **string** | The type of the event. | [optional] 
**FlowFileUuid** | Pointer to **string** | The uuid of the flowfile for the event. | [optional] 
**FileSize** | Pointer to **string** | The size of the flowfile for the event. | [optional] 
**FileSizeBytes** | Pointer to **int64** | The size of the flowfile in bytes for the event. | [optional] 
**ClusterNodeId** | Pointer to **string** | The identifier for the node where the event originated. | [optional] 
**ClusterNodeAddress** | Pointer to **string** | The label for the node where the event originated. | [optional] 
**GroupId** | Pointer to **string** | The id of the group that the component resides in. If the component is no longer in the flow, the group id will not be set. | [optional] 
**ComponentId** | Pointer to **string** | The id of the component that generated the event. | [optional] 
**ComponentType** | Pointer to **string** | The type of the component that generated the event. | [optional] 
**ComponentName** | Pointer to **string** | The name of the component that generated the event. | [optional] 
**SourceSystemFlowFileId** | Pointer to **string** | The source system flowfile id. | [optional] 
**AlternateIdentifierUri** | Pointer to **string** | The alternate identifier uri for the fileflow for the event. | [optional] 
**Attributes** | Pointer to [**[]AttributeDTO**](AttributeDTO.md) | The attributes of the flowfile for the event. | [optional] 
**ParentUuids** | Pointer to **[]string** | The parent uuids for the event. | [optional] 
**ChildUuids** | Pointer to **[]string** | The child uuids for the event. | [optional] 
**TransitUri** | Pointer to **string** | The source/destination system uri if the event was a RECEIVE/SEND. | [optional] 
**Relationship** | Pointer to **string** | The relationship to which the flowfile was routed if the event is of type ROUTE. | [optional] 
**Details** | Pointer to **string** | The event details. | [optional] 
**ContentEqual** | Pointer to **bool** | Whether the input and output content claim is the same. | [optional] 
**InputContentAvailable** | Pointer to **bool** | Whether the input content is still available. | [optional] 
**InputContentClaimSection** | Pointer to **string** | The section in which the input content claim lives. | [optional] 
**InputContentClaimContainer** | Pointer to **string** | The container in which the input content claim lives. | [optional] 
**InputContentClaimIdentifier** | Pointer to **string** | The identifier of the input content claim. | [optional] 
**InputContentClaimOffset** | Pointer to **int64** | The offset into the input content claim where the flowfiles content begins. | [optional] 
**InputContentClaimFileSize** | Pointer to **string** | The file size of the input content claim formatted. | [optional] 
**InputContentClaimFileSizeBytes** | Pointer to **int64** | The file size of the intput content claim in bytes. | [optional] 
**OutputContentAvailable** | Pointer to **bool** | Whether the output content is still available. | [optional] 
**OutputContentClaimSection** | Pointer to **string** | The section in which the output content claim lives. | [optional] 
**OutputContentClaimContainer** | Pointer to **string** | The container in which the output content claim lives. | [optional] 
**OutputContentClaimIdentifier** | Pointer to **string** | The identifier of the output content claim. | [optional] 
**OutputContentClaimOffset** | Pointer to **int64** | The offset into the output content claim where the flowfiles content begins. | [optional] 
**OutputContentClaimFileSize** | Pointer to **string** | The file size of the output content claim formatted. | [optional] 
**OutputContentClaimFileSizeBytes** | Pointer to **int64** | The file size of the output content claim in bytes. | [optional] 
**ReplayAvailable** | Pointer to **bool** | Whether or not replay is available. | [optional] 
**ReplayExplanation** | Pointer to **string** | Explanation as to why replay is unavailable. | [optional] 
**SourceConnectionIdentifier** | Pointer to **string** | The identifier of the queue/connection from which the flowfile was pulled to genereate this event. May be null if the queue/connection is unknown or the flowfile was generated from this event. | [optional] 

## Methods

### NewProvenanceEventDTO

`func NewProvenanceEventDTO() *ProvenanceEventDTO`

NewProvenanceEventDTO instantiates a new ProvenanceEventDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvenanceEventDTOWithDefaults

`func NewProvenanceEventDTOWithDefaults() *ProvenanceEventDTO`

NewProvenanceEventDTOWithDefaults instantiates a new ProvenanceEventDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProvenanceEventDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvenanceEventDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvenanceEventDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProvenanceEventDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEventId

`func (o *ProvenanceEventDTO) GetEventId() int64`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *ProvenanceEventDTO) GetEventIdOk() (*int64, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *ProvenanceEventDTO) SetEventId(v int64)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *ProvenanceEventDTO) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventTime

`func (o *ProvenanceEventDTO) GetEventTime() string`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *ProvenanceEventDTO) GetEventTimeOk() (*string, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *ProvenanceEventDTO) SetEventTime(v string)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *ProvenanceEventDTO) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetEventDuration

`func (o *ProvenanceEventDTO) GetEventDuration() int64`

GetEventDuration returns the EventDuration field if non-nil, zero value otherwise.

### GetEventDurationOk

`func (o *ProvenanceEventDTO) GetEventDurationOk() (*int64, bool)`

GetEventDurationOk returns a tuple with the EventDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDuration

`func (o *ProvenanceEventDTO) SetEventDuration(v int64)`

SetEventDuration sets EventDuration field to given value.

### HasEventDuration

`func (o *ProvenanceEventDTO) HasEventDuration() bool`

HasEventDuration returns a boolean if a field has been set.

### GetLineageDuration

`func (o *ProvenanceEventDTO) GetLineageDuration() int64`

GetLineageDuration returns the LineageDuration field if non-nil, zero value otherwise.

### GetLineageDurationOk

`func (o *ProvenanceEventDTO) GetLineageDurationOk() (*int64, bool)`

GetLineageDurationOk returns a tuple with the LineageDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineageDuration

`func (o *ProvenanceEventDTO) SetLineageDuration(v int64)`

SetLineageDuration sets LineageDuration field to given value.

### HasLineageDuration

`func (o *ProvenanceEventDTO) HasLineageDuration() bool`

HasLineageDuration returns a boolean if a field has been set.

### GetEventType

`func (o *ProvenanceEventDTO) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ProvenanceEventDTO) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ProvenanceEventDTO) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ProvenanceEventDTO) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetFlowFileUuid

`func (o *ProvenanceEventDTO) GetFlowFileUuid() string`

GetFlowFileUuid returns the FlowFileUuid field if non-nil, zero value otherwise.

### GetFlowFileUuidOk

`func (o *ProvenanceEventDTO) GetFlowFileUuidOk() (*string, bool)`

GetFlowFileUuidOk returns a tuple with the FlowFileUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFileUuid

`func (o *ProvenanceEventDTO) SetFlowFileUuid(v string)`

SetFlowFileUuid sets FlowFileUuid field to given value.

### HasFlowFileUuid

`func (o *ProvenanceEventDTO) HasFlowFileUuid() bool`

HasFlowFileUuid returns a boolean if a field has been set.

### GetFileSize

`func (o *ProvenanceEventDTO) GetFileSize() string`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *ProvenanceEventDTO) GetFileSizeOk() (*string, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *ProvenanceEventDTO) SetFileSize(v string)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *ProvenanceEventDTO) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFileSizeBytes

`func (o *ProvenanceEventDTO) GetFileSizeBytes() int64`

GetFileSizeBytes returns the FileSizeBytes field if non-nil, zero value otherwise.

### GetFileSizeBytesOk

`func (o *ProvenanceEventDTO) GetFileSizeBytesOk() (*int64, bool)`

GetFileSizeBytesOk returns a tuple with the FileSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeBytes

`func (o *ProvenanceEventDTO) SetFileSizeBytes(v int64)`

SetFileSizeBytes sets FileSizeBytes field to given value.

### HasFileSizeBytes

`func (o *ProvenanceEventDTO) HasFileSizeBytes() bool`

HasFileSizeBytes returns a boolean if a field has been set.

### GetClusterNodeId

`func (o *ProvenanceEventDTO) GetClusterNodeId() string`

GetClusterNodeId returns the ClusterNodeId field if non-nil, zero value otherwise.

### GetClusterNodeIdOk

`func (o *ProvenanceEventDTO) GetClusterNodeIdOk() (*string, bool)`

GetClusterNodeIdOk returns a tuple with the ClusterNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNodeId

`func (o *ProvenanceEventDTO) SetClusterNodeId(v string)`

SetClusterNodeId sets ClusterNodeId field to given value.

### HasClusterNodeId

`func (o *ProvenanceEventDTO) HasClusterNodeId() bool`

HasClusterNodeId returns a boolean if a field has been set.

### GetClusterNodeAddress

`func (o *ProvenanceEventDTO) GetClusterNodeAddress() string`

GetClusterNodeAddress returns the ClusterNodeAddress field if non-nil, zero value otherwise.

### GetClusterNodeAddressOk

`func (o *ProvenanceEventDTO) GetClusterNodeAddressOk() (*string, bool)`

GetClusterNodeAddressOk returns a tuple with the ClusterNodeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNodeAddress

`func (o *ProvenanceEventDTO) SetClusterNodeAddress(v string)`

SetClusterNodeAddress sets ClusterNodeAddress field to given value.

### HasClusterNodeAddress

`func (o *ProvenanceEventDTO) HasClusterNodeAddress() bool`

HasClusterNodeAddress returns a boolean if a field has been set.

### GetGroupId

`func (o *ProvenanceEventDTO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ProvenanceEventDTO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ProvenanceEventDTO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ProvenanceEventDTO) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetComponentId

`func (o *ProvenanceEventDTO) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *ProvenanceEventDTO) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *ProvenanceEventDTO) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *ProvenanceEventDTO) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### GetComponentType

`func (o *ProvenanceEventDTO) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *ProvenanceEventDTO) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *ProvenanceEventDTO) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *ProvenanceEventDTO) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### GetComponentName

`func (o *ProvenanceEventDTO) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *ProvenanceEventDTO) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *ProvenanceEventDTO) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *ProvenanceEventDTO) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.

### GetSourceSystemFlowFileId

`func (o *ProvenanceEventDTO) GetSourceSystemFlowFileId() string`

GetSourceSystemFlowFileId returns the SourceSystemFlowFileId field if non-nil, zero value otherwise.

### GetSourceSystemFlowFileIdOk

`func (o *ProvenanceEventDTO) GetSourceSystemFlowFileIdOk() (*string, bool)`

GetSourceSystemFlowFileIdOk returns a tuple with the SourceSystemFlowFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSystemFlowFileId

`func (o *ProvenanceEventDTO) SetSourceSystemFlowFileId(v string)`

SetSourceSystemFlowFileId sets SourceSystemFlowFileId field to given value.

### HasSourceSystemFlowFileId

`func (o *ProvenanceEventDTO) HasSourceSystemFlowFileId() bool`

HasSourceSystemFlowFileId returns a boolean if a field has been set.

### GetAlternateIdentifierUri

`func (o *ProvenanceEventDTO) GetAlternateIdentifierUri() string`

GetAlternateIdentifierUri returns the AlternateIdentifierUri field if non-nil, zero value otherwise.

### GetAlternateIdentifierUriOk

`func (o *ProvenanceEventDTO) GetAlternateIdentifierUriOk() (*string, bool)`

GetAlternateIdentifierUriOk returns a tuple with the AlternateIdentifierUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateIdentifierUri

`func (o *ProvenanceEventDTO) SetAlternateIdentifierUri(v string)`

SetAlternateIdentifierUri sets AlternateIdentifierUri field to given value.

### HasAlternateIdentifierUri

`func (o *ProvenanceEventDTO) HasAlternateIdentifierUri() bool`

HasAlternateIdentifierUri returns a boolean if a field has been set.

### GetAttributes

`func (o *ProvenanceEventDTO) GetAttributes() []AttributeDTO`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ProvenanceEventDTO) GetAttributesOk() (*[]AttributeDTO, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ProvenanceEventDTO) SetAttributes(v []AttributeDTO)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ProvenanceEventDTO) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetParentUuids

`func (o *ProvenanceEventDTO) GetParentUuids() []string`

GetParentUuids returns the ParentUuids field if non-nil, zero value otherwise.

### GetParentUuidsOk

`func (o *ProvenanceEventDTO) GetParentUuidsOk() (*[]string, bool)`

GetParentUuidsOk returns a tuple with the ParentUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUuids

`func (o *ProvenanceEventDTO) SetParentUuids(v []string)`

SetParentUuids sets ParentUuids field to given value.

### HasParentUuids

`func (o *ProvenanceEventDTO) HasParentUuids() bool`

HasParentUuids returns a boolean if a field has been set.

### GetChildUuids

`func (o *ProvenanceEventDTO) GetChildUuids() []string`

GetChildUuids returns the ChildUuids field if non-nil, zero value otherwise.

### GetChildUuidsOk

`func (o *ProvenanceEventDTO) GetChildUuidsOk() (*[]string, bool)`

GetChildUuidsOk returns a tuple with the ChildUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildUuids

`func (o *ProvenanceEventDTO) SetChildUuids(v []string)`

SetChildUuids sets ChildUuids field to given value.

### HasChildUuids

`func (o *ProvenanceEventDTO) HasChildUuids() bool`

HasChildUuids returns a boolean if a field has been set.

### GetTransitUri

`func (o *ProvenanceEventDTO) GetTransitUri() string`

GetTransitUri returns the TransitUri field if non-nil, zero value otherwise.

### GetTransitUriOk

`func (o *ProvenanceEventDTO) GetTransitUriOk() (*string, bool)`

GetTransitUriOk returns a tuple with the TransitUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitUri

`func (o *ProvenanceEventDTO) SetTransitUri(v string)`

SetTransitUri sets TransitUri field to given value.

### HasTransitUri

`func (o *ProvenanceEventDTO) HasTransitUri() bool`

HasTransitUri returns a boolean if a field has been set.

### GetRelationship

`func (o *ProvenanceEventDTO) GetRelationship() string`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *ProvenanceEventDTO) GetRelationshipOk() (*string, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *ProvenanceEventDTO) SetRelationship(v string)`

SetRelationship sets Relationship field to given value.

### HasRelationship

`func (o *ProvenanceEventDTO) HasRelationship() bool`

HasRelationship returns a boolean if a field has been set.

### GetDetails

`func (o *ProvenanceEventDTO) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ProvenanceEventDTO) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ProvenanceEventDTO) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ProvenanceEventDTO) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetContentEqual

`func (o *ProvenanceEventDTO) GetContentEqual() bool`

GetContentEqual returns the ContentEqual field if non-nil, zero value otherwise.

### GetContentEqualOk

`func (o *ProvenanceEventDTO) GetContentEqualOk() (*bool, bool)`

GetContentEqualOk returns a tuple with the ContentEqual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentEqual

`func (o *ProvenanceEventDTO) SetContentEqual(v bool)`

SetContentEqual sets ContentEqual field to given value.

### HasContentEqual

`func (o *ProvenanceEventDTO) HasContentEqual() bool`

HasContentEqual returns a boolean if a field has been set.

### GetInputContentAvailable

`func (o *ProvenanceEventDTO) GetInputContentAvailable() bool`

GetInputContentAvailable returns the InputContentAvailable field if non-nil, zero value otherwise.

### GetInputContentAvailableOk

`func (o *ProvenanceEventDTO) GetInputContentAvailableOk() (*bool, bool)`

GetInputContentAvailableOk returns a tuple with the InputContentAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputContentAvailable

`func (o *ProvenanceEventDTO) SetInputContentAvailable(v bool)`

SetInputContentAvailable sets InputContentAvailable field to given value.

### HasInputContentAvailable

`func (o *ProvenanceEventDTO) HasInputContentAvailable() bool`

HasInputContentAvailable returns a boolean if a field has been set.

### GetInputContentClaimSection

`func (o *ProvenanceEventDTO) GetInputContentClaimSection() string`

GetInputContentClaimSection returns the InputContentClaimSection field if non-nil, zero value otherwise.

### GetInputContentClaimSectionOk

`func (o *ProvenanceEventDTO) GetInputContentClaimSectionOk() (*string, bool)`

GetInputContentClaimSectionOk returns a tuple with the InputContentClaimSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputContentClaimSection

`func (o *ProvenanceEventDTO) SetInputContentClaimSection(v string)`

SetInputContentClaimSection sets InputContentClaimSection field to given value.

### HasInputContentClaimSection

`func (o *ProvenanceEventDTO) HasInputContentClaimSection() bool`

HasInputContentClaimSection returns a boolean if a field has been set.

### GetInputContentClaimContainer

`func (o *ProvenanceEventDTO) GetInputContentClaimContainer() string`

GetInputContentClaimContainer returns the InputContentClaimContainer field if non-nil, zero value otherwise.

### GetInputContentClaimContainerOk

`func (o *ProvenanceEventDTO) GetInputContentClaimContainerOk() (*string, bool)`

GetInputContentClaimContainerOk returns a tuple with the InputContentClaimContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputContentClaimContainer

`func (o *ProvenanceEventDTO) SetInputContentClaimContainer(v string)`

SetInputContentClaimContainer sets InputContentClaimContainer field to given value.

### HasInputContentClaimContainer

`func (o *ProvenanceEventDTO) HasInputContentClaimContainer() bool`

HasInputContentClaimContainer returns a boolean if a field has been set.

### GetInputContentClaimIdentifier

`func (o *ProvenanceEventDTO) GetInputContentClaimIdentifier() string`

GetInputContentClaimIdentifier returns the InputContentClaimIdentifier field if non-nil, zero value otherwise.

### GetInputContentClaimIdentifierOk

`func (o *ProvenanceEventDTO) GetInputContentClaimIdentifierOk() (*string, bool)`

GetInputContentClaimIdentifierOk returns a tuple with the InputContentClaimIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputContentClaimIdentifier

`func (o *ProvenanceEventDTO) SetInputContentClaimIdentifier(v string)`

SetInputContentClaimIdentifier sets InputContentClaimIdentifier field to given value.

### HasInputContentClaimIdentifier

`func (o *ProvenanceEventDTO) HasInputContentClaimIdentifier() bool`

HasInputContentClaimIdentifier returns a boolean if a field has been set.

### GetInputContentClaimOffset

`func (o *ProvenanceEventDTO) GetInputContentClaimOffset() int64`

GetInputContentClaimOffset returns the InputContentClaimOffset field if non-nil, zero value otherwise.

### GetInputContentClaimOffsetOk

`func (o *ProvenanceEventDTO) GetInputContentClaimOffsetOk() (*int64, bool)`

GetInputContentClaimOffsetOk returns a tuple with the InputContentClaimOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputContentClaimOffset

`func (o *ProvenanceEventDTO) SetInputContentClaimOffset(v int64)`

SetInputContentClaimOffset sets InputContentClaimOffset field to given value.

### HasInputContentClaimOffset

`func (o *ProvenanceEventDTO) HasInputContentClaimOffset() bool`

HasInputContentClaimOffset returns a boolean if a field has been set.

### GetInputContentClaimFileSize

`func (o *ProvenanceEventDTO) GetInputContentClaimFileSize() string`

GetInputContentClaimFileSize returns the InputContentClaimFileSize field if non-nil, zero value otherwise.

### GetInputContentClaimFileSizeOk

`func (o *ProvenanceEventDTO) GetInputContentClaimFileSizeOk() (*string, bool)`

GetInputContentClaimFileSizeOk returns a tuple with the InputContentClaimFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputContentClaimFileSize

`func (o *ProvenanceEventDTO) SetInputContentClaimFileSize(v string)`

SetInputContentClaimFileSize sets InputContentClaimFileSize field to given value.

### HasInputContentClaimFileSize

`func (o *ProvenanceEventDTO) HasInputContentClaimFileSize() bool`

HasInputContentClaimFileSize returns a boolean if a field has been set.

### GetInputContentClaimFileSizeBytes

`func (o *ProvenanceEventDTO) GetInputContentClaimFileSizeBytes() int64`

GetInputContentClaimFileSizeBytes returns the InputContentClaimFileSizeBytes field if non-nil, zero value otherwise.

### GetInputContentClaimFileSizeBytesOk

`func (o *ProvenanceEventDTO) GetInputContentClaimFileSizeBytesOk() (*int64, bool)`

GetInputContentClaimFileSizeBytesOk returns a tuple with the InputContentClaimFileSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputContentClaimFileSizeBytes

`func (o *ProvenanceEventDTO) SetInputContentClaimFileSizeBytes(v int64)`

SetInputContentClaimFileSizeBytes sets InputContentClaimFileSizeBytes field to given value.

### HasInputContentClaimFileSizeBytes

`func (o *ProvenanceEventDTO) HasInputContentClaimFileSizeBytes() bool`

HasInputContentClaimFileSizeBytes returns a boolean if a field has been set.

### GetOutputContentAvailable

`func (o *ProvenanceEventDTO) GetOutputContentAvailable() bool`

GetOutputContentAvailable returns the OutputContentAvailable field if non-nil, zero value otherwise.

### GetOutputContentAvailableOk

`func (o *ProvenanceEventDTO) GetOutputContentAvailableOk() (*bool, bool)`

GetOutputContentAvailableOk returns a tuple with the OutputContentAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputContentAvailable

`func (o *ProvenanceEventDTO) SetOutputContentAvailable(v bool)`

SetOutputContentAvailable sets OutputContentAvailable field to given value.

### HasOutputContentAvailable

`func (o *ProvenanceEventDTO) HasOutputContentAvailable() bool`

HasOutputContentAvailable returns a boolean if a field has been set.

### GetOutputContentClaimSection

`func (o *ProvenanceEventDTO) GetOutputContentClaimSection() string`

GetOutputContentClaimSection returns the OutputContentClaimSection field if non-nil, zero value otherwise.

### GetOutputContentClaimSectionOk

`func (o *ProvenanceEventDTO) GetOutputContentClaimSectionOk() (*string, bool)`

GetOutputContentClaimSectionOk returns a tuple with the OutputContentClaimSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputContentClaimSection

`func (o *ProvenanceEventDTO) SetOutputContentClaimSection(v string)`

SetOutputContentClaimSection sets OutputContentClaimSection field to given value.

### HasOutputContentClaimSection

`func (o *ProvenanceEventDTO) HasOutputContentClaimSection() bool`

HasOutputContentClaimSection returns a boolean if a field has been set.

### GetOutputContentClaimContainer

`func (o *ProvenanceEventDTO) GetOutputContentClaimContainer() string`

GetOutputContentClaimContainer returns the OutputContentClaimContainer field if non-nil, zero value otherwise.

### GetOutputContentClaimContainerOk

`func (o *ProvenanceEventDTO) GetOutputContentClaimContainerOk() (*string, bool)`

GetOutputContentClaimContainerOk returns a tuple with the OutputContentClaimContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputContentClaimContainer

`func (o *ProvenanceEventDTO) SetOutputContentClaimContainer(v string)`

SetOutputContentClaimContainer sets OutputContentClaimContainer field to given value.

### HasOutputContentClaimContainer

`func (o *ProvenanceEventDTO) HasOutputContentClaimContainer() bool`

HasOutputContentClaimContainer returns a boolean if a field has been set.

### GetOutputContentClaimIdentifier

`func (o *ProvenanceEventDTO) GetOutputContentClaimIdentifier() string`

GetOutputContentClaimIdentifier returns the OutputContentClaimIdentifier field if non-nil, zero value otherwise.

### GetOutputContentClaimIdentifierOk

`func (o *ProvenanceEventDTO) GetOutputContentClaimIdentifierOk() (*string, bool)`

GetOutputContentClaimIdentifierOk returns a tuple with the OutputContentClaimIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputContentClaimIdentifier

`func (o *ProvenanceEventDTO) SetOutputContentClaimIdentifier(v string)`

SetOutputContentClaimIdentifier sets OutputContentClaimIdentifier field to given value.

### HasOutputContentClaimIdentifier

`func (o *ProvenanceEventDTO) HasOutputContentClaimIdentifier() bool`

HasOutputContentClaimIdentifier returns a boolean if a field has been set.

### GetOutputContentClaimOffset

`func (o *ProvenanceEventDTO) GetOutputContentClaimOffset() int64`

GetOutputContentClaimOffset returns the OutputContentClaimOffset field if non-nil, zero value otherwise.

### GetOutputContentClaimOffsetOk

`func (o *ProvenanceEventDTO) GetOutputContentClaimOffsetOk() (*int64, bool)`

GetOutputContentClaimOffsetOk returns a tuple with the OutputContentClaimOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputContentClaimOffset

`func (o *ProvenanceEventDTO) SetOutputContentClaimOffset(v int64)`

SetOutputContentClaimOffset sets OutputContentClaimOffset field to given value.

### HasOutputContentClaimOffset

`func (o *ProvenanceEventDTO) HasOutputContentClaimOffset() bool`

HasOutputContentClaimOffset returns a boolean if a field has been set.

### GetOutputContentClaimFileSize

`func (o *ProvenanceEventDTO) GetOutputContentClaimFileSize() string`

GetOutputContentClaimFileSize returns the OutputContentClaimFileSize field if non-nil, zero value otherwise.

### GetOutputContentClaimFileSizeOk

`func (o *ProvenanceEventDTO) GetOutputContentClaimFileSizeOk() (*string, bool)`

GetOutputContentClaimFileSizeOk returns a tuple with the OutputContentClaimFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputContentClaimFileSize

`func (o *ProvenanceEventDTO) SetOutputContentClaimFileSize(v string)`

SetOutputContentClaimFileSize sets OutputContentClaimFileSize field to given value.

### HasOutputContentClaimFileSize

`func (o *ProvenanceEventDTO) HasOutputContentClaimFileSize() bool`

HasOutputContentClaimFileSize returns a boolean if a field has been set.

### GetOutputContentClaimFileSizeBytes

`func (o *ProvenanceEventDTO) GetOutputContentClaimFileSizeBytes() int64`

GetOutputContentClaimFileSizeBytes returns the OutputContentClaimFileSizeBytes field if non-nil, zero value otherwise.

### GetOutputContentClaimFileSizeBytesOk

`func (o *ProvenanceEventDTO) GetOutputContentClaimFileSizeBytesOk() (*int64, bool)`

GetOutputContentClaimFileSizeBytesOk returns a tuple with the OutputContentClaimFileSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputContentClaimFileSizeBytes

`func (o *ProvenanceEventDTO) SetOutputContentClaimFileSizeBytes(v int64)`

SetOutputContentClaimFileSizeBytes sets OutputContentClaimFileSizeBytes field to given value.

### HasOutputContentClaimFileSizeBytes

`func (o *ProvenanceEventDTO) HasOutputContentClaimFileSizeBytes() bool`

HasOutputContentClaimFileSizeBytes returns a boolean if a field has been set.

### GetReplayAvailable

`func (o *ProvenanceEventDTO) GetReplayAvailable() bool`

GetReplayAvailable returns the ReplayAvailable field if non-nil, zero value otherwise.

### GetReplayAvailableOk

`func (o *ProvenanceEventDTO) GetReplayAvailableOk() (*bool, bool)`

GetReplayAvailableOk returns a tuple with the ReplayAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayAvailable

`func (o *ProvenanceEventDTO) SetReplayAvailable(v bool)`

SetReplayAvailable sets ReplayAvailable field to given value.

### HasReplayAvailable

`func (o *ProvenanceEventDTO) HasReplayAvailable() bool`

HasReplayAvailable returns a boolean if a field has been set.

### GetReplayExplanation

`func (o *ProvenanceEventDTO) GetReplayExplanation() string`

GetReplayExplanation returns the ReplayExplanation field if non-nil, zero value otherwise.

### GetReplayExplanationOk

`func (o *ProvenanceEventDTO) GetReplayExplanationOk() (*string, bool)`

GetReplayExplanationOk returns a tuple with the ReplayExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayExplanation

`func (o *ProvenanceEventDTO) SetReplayExplanation(v string)`

SetReplayExplanation sets ReplayExplanation field to given value.

### HasReplayExplanation

`func (o *ProvenanceEventDTO) HasReplayExplanation() bool`

HasReplayExplanation returns a boolean if a field has been set.

### GetSourceConnectionIdentifier

`func (o *ProvenanceEventDTO) GetSourceConnectionIdentifier() string`

GetSourceConnectionIdentifier returns the SourceConnectionIdentifier field if non-nil, zero value otherwise.

### GetSourceConnectionIdentifierOk

`func (o *ProvenanceEventDTO) GetSourceConnectionIdentifierOk() (*string, bool)`

GetSourceConnectionIdentifierOk returns a tuple with the SourceConnectionIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConnectionIdentifier

`func (o *ProvenanceEventDTO) SetSourceConnectionIdentifier(v string)`

SetSourceConnectionIdentifier sets SourceConnectionIdentifier field to given value.

### HasSourceConnectionIdentifier

`func (o *ProvenanceEventDTO) HasSourceConnectionIdentifier() bool`

HasSourceConnectionIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


