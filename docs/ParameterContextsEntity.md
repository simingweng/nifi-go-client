# ParameterContextsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParameterContexts** | Pointer to [**[]ParameterContextEntity**](ParameterContextEntity.md) | The Parameter Contexts | [optional] 
**CurrentTime** | Pointer to **string** | The current time on the system. | [optional] [readonly] 

## Methods

### NewParameterContextsEntity

`func NewParameterContextsEntity() *ParameterContextsEntity`

NewParameterContextsEntity instantiates a new ParameterContextsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterContextsEntityWithDefaults

`func NewParameterContextsEntityWithDefaults() *ParameterContextsEntity`

NewParameterContextsEntityWithDefaults instantiates a new ParameterContextsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameterContexts

`func (o *ParameterContextsEntity) GetParameterContexts() []ParameterContextEntity`

GetParameterContexts returns the ParameterContexts field if non-nil, zero value otherwise.

### GetParameterContextsOk

`func (o *ParameterContextsEntity) GetParameterContextsOk() (*[]ParameterContextEntity, bool)`

GetParameterContextsOk returns a tuple with the ParameterContexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterContexts

`func (o *ParameterContextsEntity) SetParameterContexts(v []ParameterContextEntity)`

SetParameterContexts sets ParameterContexts field to given value.

### HasParameterContexts

`func (o *ParameterContextsEntity) HasParameterContexts() bool`

HasParameterContexts returns a boolean if a field has been set.

### GetCurrentTime

`func (o *ParameterContextsEntity) GetCurrentTime() string`

GetCurrentTime returns the CurrentTime field if non-nil, zero value otherwise.

### GetCurrentTimeOk

`func (o *ParameterContextsEntity) GetCurrentTimeOk() (*string, bool)`

GetCurrentTimeOk returns a tuple with the CurrentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTime

`func (o *ParameterContextsEntity) SetCurrentTime(v string)`

SetCurrentTime sets CurrentTime field to given value.

### HasCurrentTime

`func (o *ParameterContextsEntity) HasCurrentTime() bool`

HasCurrentTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


