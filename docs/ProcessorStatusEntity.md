# ProcessorStatusEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessorStatus** | Pointer to [**ProcessorStatusDTO**](ProcessorStatusDTO.md) |  | [optional] 
**CanRead** | Pointer to **bool** | Indicates whether the user can read a given resource. | [optional] [readonly] 

## Methods

### NewProcessorStatusEntity

`func NewProcessorStatusEntity() *ProcessorStatusEntity`

NewProcessorStatusEntity instantiates a new ProcessorStatusEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorStatusEntityWithDefaults

`func NewProcessorStatusEntityWithDefaults() *ProcessorStatusEntity`

NewProcessorStatusEntityWithDefaults instantiates a new ProcessorStatusEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessorStatus

`func (o *ProcessorStatusEntity) GetProcessorStatus() ProcessorStatusDTO`

GetProcessorStatus returns the ProcessorStatus field if non-nil, zero value otherwise.

### GetProcessorStatusOk

`func (o *ProcessorStatusEntity) GetProcessorStatusOk() (*ProcessorStatusDTO, bool)`

GetProcessorStatusOk returns a tuple with the ProcessorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorStatus

`func (o *ProcessorStatusEntity) SetProcessorStatus(v ProcessorStatusDTO)`

SetProcessorStatus sets ProcessorStatus field to given value.

### HasProcessorStatus

`func (o *ProcessorStatusEntity) HasProcessorStatus() bool`

HasProcessorStatus returns a boolean if a field has been set.

### GetCanRead

`func (o *ProcessorStatusEntity) GetCanRead() bool`

GetCanRead returns the CanRead field if non-nil, zero value otherwise.

### GetCanReadOk

`func (o *ProcessorStatusEntity) GetCanReadOk() (*bool, bool)`

GetCanReadOk returns a tuple with the CanRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRead

`func (o *ProcessorStatusEntity) SetCanRead(v bool)`

SetCanRead sets CanRead field to given value.

### HasCanRead

`func (o *ProcessorStatusEntity) HasCanRead() bool`

HasCanRead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


