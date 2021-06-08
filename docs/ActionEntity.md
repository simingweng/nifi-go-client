# ActionEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **string** | The timestamp of the action. | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**CanRead** | Pointer to **bool** | Indicates whether the user can read a given resource. | [optional] [readonly] 
**Action** | Pointer to [**ActionDTO**](ActionDTO.md) |  | [optional] 

## Methods

### NewActionEntity

`func NewActionEntity() *ActionEntity`

NewActionEntity instantiates a new ActionEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionEntityWithDefaults

`func NewActionEntityWithDefaults() *ActionEntity`

NewActionEntityWithDefaults instantiates a new ActionEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionEntity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionEntity) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionEntity) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ActionEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ActionEntity) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ActionEntity) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ActionEntity) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ActionEntity) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetSourceId

`func (o *ActionEntity) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ActionEntity) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ActionEntity) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ActionEntity) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetCanRead

`func (o *ActionEntity) GetCanRead() bool`

GetCanRead returns the CanRead field if non-nil, zero value otherwise.

### GetCanReadOk

`func (o *ActionEntity) GetCanReadOk() (*bool, bool)`

GetCanReadOk returns a tuple with the CanRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRead

`func (o *ActionEntity) SetCanRead(v bool)`

SetCanRead sets CanRead field to given value.

### HasCanRead

`func (o *ActionEntity) HasCanRead() bool`

HasCanRead returns a boolean if a field has been set.

### GetAction

`func (o *ActionEntity) GetAction() ActionDTO`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ActionEntity) GetActionOk() (*ActionDTO, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ActionEntity) SetAction(v ActionDTO)`

SetAction sets Action field to given value.

### HasAction

`func (o *ActionEntity) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


