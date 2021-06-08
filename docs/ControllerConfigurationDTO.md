# ControllerConfigurationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxTimerDrivenThreadCount** | Pointer to **int32** | The maximum number of timer driven threads the NiFi has available. | [optional] 
**MaxEventDrivenThreadCount** | Pointer to **int32** | The maximum number of event driven threads the NiFi has available. | [optional] 

## Methods

### NewControllerConfigurationDTO

`func NewControllerConfigurationDTO() *ControllerConfigurationDTO`

NewControllerConfigurationDTO instantiates a new ControllerConfigurationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerConfigurationDTOWithDefaults

`func NewControllerConfigurationDTOWithDefaults() *ControllerConfigurationDTO`

NewControllerConfigurationDTOWithDefaults instantiates a new ControllerConfigurationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxTimerDrivenThreadCount

`func (o *ControllerConfigurationDTO) GetMaxTimerDrivenThreadCount() int32`

GetMaxTimerDrivenThreadCount returns the MaxTimerDrivenThreadCount field if non-nil, zero value otherwise.

### GetMaxTimerDrivenThreadCountOk

`func (o *ControllerConfigurationDTO) GetMaxTimerDrivenThreadCountOk() (*int32, bool)`

GetMaxTimerDrivenThreadCountOk returns a tuple with the MaxTimerDrivenThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTimerDrivenThreadCount

`func (o *ControllerConfigurationDTO) SetMaxTimerDrivenThreadCount(v int32)`

SetMaxTimerDrivenThreadCount sets MaxTimerDrivenThreadCount field to given value.

### HasMaxTimerDrivenThreadCount

`func (o *ControllerConfigurationDTO) HasMaxTimerDrivenThreadCount() bool`

HasMaxTimerDrivenThreadCount returns a boolean if a field has been set.

### GetMaxEventDrivenThreadCount

`func (o *ControllerConfigurationDTO) GetMaxEventDrivenThreadCount() int32`

GetMaxEventDrivenThreadCount returns the MaxEventDrivenThreadCount field if non-nil, zero value otherwise.

### GetMaxEventDrivenThreadCountOk

`func (o *ControllerConfigurationDTO) GetMaxEventDrivenThreadCountOk() (*int32, bool)`

GetMaxEventDrivenThreadCountOk returns a tuple with the MaxEventDrivenThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEventDrivenThreadCount

`func (o *ControllerConfigurationDTO) SetMaxEventDrivenThreadCount(v int32)`

SetMaxEventDrivenThreadCount sets MaxEventDrivenThreadCount field to given value.

### HasMaxEventDrivenThreadCount

`func (o *ControllerConfigurationDTO) HasMaxEventDrivenThreadCount() bool`

HasMaxEventDrivenThreadCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


