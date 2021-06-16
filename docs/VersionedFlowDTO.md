# VersionedFlowDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistryId** | Pointer to **string** | The ID of the registry that the flow is tracked to | [optional] 
**BucketId** | Pointer to **string** | The ID of the bucket where the flow is stored | [optional] 
**FlowId** | Pointer to **string** | The ID of the flow | [optional] 
**FlowName** | Pointer to **string** | The name of the flow | [optional] 
**Description** | Pointer to **string** | A description of the flow | [optional] 
**Comments** | Pointer to **string** | Comments for the changeset | [optional] 
**Action** | Pointer to **string** | The action being performed | [optional] 

## Methods

### NewVersionedFlowDTO

`func NewVersionedFlowDTO() *VersionedFlowDTO`

NewVersionedFlowDTO instantiates a new VersionedFlowDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionedFlowDTOWithDefaults

`func NewVersionedFlowDTOWithDefaults() *VersionedFlowDTO`

NewVersionedFlowDTOWithDefaults instantiates a new VersionedFlowDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistryId

`func (o *VersionedFlowDTO) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *VersionedFlowDTO) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *VersionedFlowDTO) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *VersionedFlowDTO) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.

### GetBucketId

`func (o *VersionedFlowDTO) GetBucketId() string`

GetBucketId returns the BucketId field if non-nil, zero value otherwise.

### GetBucketIdOk

`func (o *VersionedFlowDTO) GetBucketIdOk() (*string, bool)`

GetBucketIdOk returns a tuple with the BucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketId

`func (o *VersionedFlowDTO) SetBucketId(v string)`

SetBucketId sets BucketId field to given value.

### HasBucketId

`func (o *VersionedFlowDTO) HasBucketId() bool`

HasBucketId returns a boolean if a field has been set.

### GetFlowId

`func (o *VersionedFlowDTO) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *VersionedFlowDTO) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *VersionedFlowDTO) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *VersionedFlowDTO) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetFlowName

`func (o *VersionedFlowDTO) GetFlowName() string`

GetFlowName returns the FlowName field if non-nil, zero value otherwise.

### GetFlowNameOk

`func (o *VersionedFlowDTO) GetFlowNameOk() (*string, bool)`

GetFlowNameOk returns a tuple with the FlowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowName

`func (o *VersionedFlowDTO) SetFlowName(v string)`

SetFlowName sets FlowName field to given value.

### HasFlowName

`func (o *VersionedFlowDTO) HasFlowName() bool`

HasFlowName returns a boolean if a field has been set.

### GetDescription

`func (o *VersionedFlowDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VersionedFlowDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VersionedFlowDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VersionedFlowDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *VersionedFlowDTO) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VersionedFlowDTO) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VersionedFlowDTO) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VersionedFlowDTO) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetAction

`func (o *VersionedFlowDTO) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VersionedFlowDTO) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VersionedFlowDTO) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *VersionedFlowDTO) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


