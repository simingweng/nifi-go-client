# AllowableValueEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowableValue** | Pointer to [**AllowableValueDTO**](AllowableValueDTO.md) |  | [optional] 
**CanRead** | Pointer to **bool** | Indicates whether the user can read a given resource. | [optional] [readonly] 

## Methods

### NewAllowableValueEntity

`func NewAllowableValueEntity() *AllowableValueEntity`

NewAllowableValueEntity instantiates a new AllowableValueEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowableValueEntityWithDefaults

`func NewAllowableValueEntityWithDefaults() *AllowableValueEntity`

NewAllowableValueEntityWithDefaults instantiates a new AllowableValueEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowableValue

`func (o *AllowableValueEntity) GetAllowableValue() AllowableValueDTO`

GetAllowableValue returns the AllowableValue field if non-nil, zero value otherwise.

### GetAllowableValueOk

`func (o *AllowableValueEntity) GetAllowableValueOk() (*AllowableValueDTO, bool)`

GetAllowableValueOk returns a tuple with the AllowableValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableValue

`func (o *AllowableValueEntity) SetAllowableValue(v AllowableValueDTO)`

SetAllowableValue sets AllowableValue field to given value.

### HasAllowableValue

`func (o *AllowableValueEntity) HasAllowableValue() bool`

HasAllowableValue returns a boolean if a field has been set.

### GetCanRead

`func (o *AllowableValueEntity) GetCanRead() bool`

GetCanRead returns the CanRead field if non-nil, zero value otherwise.

### GetCanReadOk

`func (o *AllowableValueEntity) GetCanReadOk() (*bool, bool)`

GetCanReadOk returns a tuple with the CanRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRead

`func (o *AllowableValueEntity) SetCanRead(v bool)`

SetCanRead sets CanRead field to given value.

### HasCanRead

`func (o *AllowableValueEntity) HasCanRead() bool`

HasCanRead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


