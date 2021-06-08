# ProcessGroupStatusEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessGroupStatus** | Pointer to [**ProcessGroupStatusDTO**](ProcessGroupStatusDTO.md) |  | [optional] 
**CanRead** | Pointer to **bool** | Indicates whether the user can read a given resource. | [optional] [readonly] 

## Methods

### NewProcessGroupStatusEntity

`func NewProcessGroupStatusEntity() *ProcessGroupStatusEntity`

NewProcessGroupStatusEntity instantiates a new ProcessGroupStatusEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessGroupStatusEntityWithDefaults

`func NewProcessGroupStatusEntityWithDefaults() *ProcessGroupStatusEntity`

NewProcessGroupStatusEntityWithDefaults instantiates a new ProcessGroupStatusEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessGroupStatus

`func (o *ProcessGroupStatusEntity) GetProcessGroupStatus() ProcessGroupStatusDTO`

GetProcessGroupStatus returns the ProcessGroupStatus field if non-nil, zero value otherwise.

### GetProcessGroupStatusOk

`func (o *ProcessGroupStatusEntity) GetProcessGroupStatusOk() (*ProcessGroupStatusDTO, bool)`

GetProcessGroupStatusOk returns a tuple with the ProcessGroupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupStatus

`func (o *ProcessGroupStatusEntity) SetProcessGroupStatus(v ProcessGroupStatusDTO)`

SetProcessGroupStatus sets ProcessGroupStatus field to given value.

### HasProcessGroupStatus

`func (o *ProcessGroupStatusEntity) HasProcessGroupStatus() bool`

HasProcessGroupStatus returns a boolean if a field has been set.

### GetCanRead

`func (o *ProcessGroupStatusEntity) GetCanRead() bool`

GetCanRead returns the CanRead field if non-nil, zero value otherwise.

### GetCanReadOk

`func (o *ProcessGroupStatusEntity) GetCanReadOk() (*bool, bool)`

GetCanReadOk returns a tuple with the CanRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRead

`func (o *ProcessGroupStatusEntity) SetCanRead(v bool)`

SetCanRead sets CanRead field to given value.

### HasCanRead

`func (o *ProcessGroupStatusEntity) HasCanRead() bool`

HasCanRead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


