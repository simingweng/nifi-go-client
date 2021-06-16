# VersionControlInformationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | The ID of the Process Group that is under version control | [optional] 
**RegistryId** | Pointer to **string** | The ID of the registry that the flow is stored in | [optional] 
**RegistryName** | Pointer to **string** | The name of the registry that the flow is stored in | [optional] [readonly] 
**BucketId** | Pointer to **string** | The ID of the bucket that the flow is stored in | [optional] 
**BucketName** | Pointer to **string** | The name of the bucket that the flow is stored in | [optional] [readonly] 
**FlowId** | Pointer to **string** | The ID of the flow | [optional] 
**FlowName** | Pointer to **string** | The name of the flow | [optional] 
**FlowDescription** | Pointer to **string** | The description of the flow | [optional] 
**Version** | Pointer to **int32** | The version of the flow | [optional] 
**State** | Pointer to **string** | The current state of the Process Group, as it relates to the Versioned Flow | [optional] [readonly] 
**StateExplanation** | Pointer to **string** | Explanation of why the group is in the specified state | [optional] [readonly] 

## Methods

### NewVersionControlInformationDTO

`func NewVersionControlInformationDTO() *VersionControlInformationDTO`

NewVersionControlInformationDTO instantiates a new VersionControlInformationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionControlInformationDTOWithDefaults

`func NewVersionControlInformationDTOWithDefaults() *VersionControlInformationDTO`

NewVersionControlInformationDTOWithDefaults instantiates a new VersionControlInformationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *VersionControlInformationDTO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *VersionControlInformationDTO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *VersionControlInformationDTO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *VersionControlInformationDTO) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetRegistryId

`func (o *VersionControlInformationDTO) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *VersionControlInformationDTO) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *VersionControlInformationDTO) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *VersionControlInformationDTO) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.

### GetRegistryName

`func (o *VersionControlInformationDTO) GetRegistryName() string`

GetRegistryName returns the RegistryName field if non-nil, zero value otherwise.

### GetRegistryNameOk

`func (o *VersionControlInformationDTO) GetRegistryNameOk() (*string, bool)`

GetRegistryNameOk returns a tuple with the RegistryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryName

`func (o *VersionControlInformationDTO) SetRegistryName(v string)`

SetRegistryName sets RegistryName field to given value.

### HasRegistryName

`func (o *VersionControlInformationDTO) HasRegistryName() bool`

HasRegistryName returns a boolean if a field has been set.

### GetBucketId

`func (o *VersionControlInformationDTO) GetBucketId() string`

GetBucketId returns the BucketId field if non-nil, zero value otherwise.

### GetBucketIdOk

`func (o *VersionControlInformationDTO) GetBucketIdOk() (*string, bool)`

GetBucketIdOk returns a tuple with the BucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketId

`func (o *VersionControlInformationDTO) SetBucketId(v string)`

SetBucketId sets BucketId field to given value.

### HasBucketId

`func (o *VersionControlInformationDTO) HasBucketId() bool`

HasBucketId returns a boolean if a field has been set.

### GetBucketName

`func (o *VersionControlInformationDTO) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *VersionControlInformationDTO) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *VersionControlInformationDTO) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *VersionControlInformationDTO) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetFlowId

`func (o *VersionControlInformationDTO) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *VersionControlInformationDTO) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *VersionControlInformationDTO) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *VersionControlInformationDTO) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetFlowName

`func (o *VersionControlInformationDTO) GetFlowName() string`

GetFlowName returns the FlowName field if non-nil, zero value otherwise.

### GetFlowNameOk

`func (o *VersionControlInformationDTO) GetFlowNameOk() (*string, bool)`

GetFlowNameOk returns a tuple with the FlowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowName

`func (o *VersionControlInformationDTO) SetFlowName(v string)`

SetFlowName sets FlowName field to given value.

### HasFlowName

`func (o *VersionControlInformationDTO) HasFlowName() bool`

HasFlowName returns a boolean if a field has been set.

### GetFlowDescription

`func (o *VersionControlInformationDTO) GetFlowDescription() string`

GetFlowDescription returns the FlowDescription field if non-nil, zero value otherwise.

### GetFlowDescriptionOk

`func (o *VersionControlInformationDTO) GetFlowDescriptionOk() (*string, bool)`

GetFlowDescriptionOk returns a tuple with the FlowDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDescription

`func (o *VersionControlInformationDTO) SetFlowDescription(v string)`

SetFlowDescription sets FlowDescription field to given value.

### HasFlowDescription

`func (o *VersionControlInformationDTO) HasFlowDescription() bool`

HasFlowDescription returns a boolean if a field has been set.

### GetVersion

`func (o *VersionControlInformationDTO) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VersionControlInformationDTO) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VersionControlInformationDTO) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VersionControlInformationDTO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetState

`func (o *VersionControlInformationDTO) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VersionControlInformationDTO) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VersionControlInformationDTO) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *VersionControlInformationDTO) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateExplanation

`func (o *VersionControlInformationDTO) GetStateExplanation() string`

GetStateExplanation returns the StateExplanation field if non-nil, zero value otherwise.

### GetStateExplanationOk

`func (o *VersionControlInformationDTO) GetStateExplanationOk() (*string, bool)`

GetStateExplanationOk returns a tuple with the StateExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateExplanation

`func (o *VersionControlInformationDTO) SetStateExplanation(v string)`

SetStateExplanation sets StateExplanation field to given value.

### HasStateExplanation

`func (o *VersionControlInformationDTO) HasStateExplanation() bool`

HasStateExplanation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


