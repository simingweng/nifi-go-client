# PortStatusEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortStatus** | Pointer to [**PortStatusDTO**](PortStatusDTO.md) |  | [optional] 
**CanRead** | Pointer to **bool** | Indicates whether the user can read a given resource. | [optional] [readonly] 

## Methods

### NewPortStatusEntity

`func NewPortStatusEntity() *PortStatusEntity`

NewPortStatusEntity instantiates a new PortStatusEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortStatusEntityWithDefaults

`func NewPortStatusEntityWithDefaults() *PortStatusEntity`

NewPortStatusEntityWithDefaults instantiates a new PortStatusEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortStatus

`func (o *PortStatusEntity) GetPortStatus() PortStatusDTO`

GetPortStatus returns the PortStatus field if non-nil, zero value otherwise.

### GetPortStatusOk

`func (o *PortStatusEntity) GetPortStatusOk() (*PortStatusDTO, bool)`

GetPortStatusOk returns a tuple with the PortStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStatus

`func (o *PortStatusEntity) SetPortStatus(v PortStatusDTO)`

SetPortStatus sets PortStatus field to given value.

### HasPortStatus

`func (o *PortStatusEntity) HasPortStatus() bool`

HasPortStatus returns a boolean if a field has been set.

### GetCanRead

`func (o *PortStatusEntity) GetCanRead() bool`

GetCanRead returns the CanRead field if non-nil, zero value otherwise.

### GetCanReadOk

`func (o *PortStatusEntity) GetCanReadOk() (*bool, bool)`

GetCanReadOk returns a tuple with the CanRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRead

`func (o *PortStatusEntity) SetCanRead(v bool)`

SetCanRead sets CanRead field to given value.

### HasCanRead

`func (o *PortStatusEntity) HasCanRead() bool`

HasCanRead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


