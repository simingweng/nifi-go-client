# FlowFileEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowFile** | Pointer to [**FlowFileDTO**](FlowFileDTO.md) |  | [optional] 

## Methods

### NewFlowFileEntity

`func NewFlowFileEntity() *FlowFileEntity`

NewFlowFileEntity instantiates a new FlowFileEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowFileEntityWithDefaults

`func NewFlowFileEntityWithDefaults() *FlowFileEntity`

NewFlowFileEntityWithDefaults instantiates a new FlowFileEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowFile

`func (o *FlowFileEntity) GetFlowFile() FlowFileDTO`

GetFlowFile returns the FlowFile field if non-nil, zero value otherwise.

### GetFlowFileOk

`func (o *FlowFileEntity) GetFlowFileOk() (*FlowFileDTO, bool)`

GetFlowFileOk returns a tuple with the FlowFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFile

`func (o *FlowFileEntity) SetFlowFile(v FlowFileDTO)`

SetFlowFile sets FlowFile field to given value.

### HasFlowFile

`func (o *FlowFileEntity) HasFlowFile() bool`

HasFlowFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


