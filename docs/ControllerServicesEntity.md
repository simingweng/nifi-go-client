# ControllerServicesEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentTime** | Pointer to **string** | The current time on the system. | [optional] 
**ControllerServices** | Pointer to [**[]ControllerServiceEntity**](ControllerServiceEntity.md) |  | [optional] 

## Methods

### NewControllerServicesEntity

`func NewControllerServicesEntity() *ControllerServicesEntity`

NewControllerServicesEntity instantiates a new ControllerServicesEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerServicesEntityWithDefaults

`func NewControllerServicesEntityWithDefaults() *ControllerServicesEntity`

NewControllerServicesEntityWithDefaults instantiates a new ControllerServicesEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentTime

`func (o *ControllerServicesEntity) GetCurrentTime() string`

GetCurrentTime returns the CurrentTime field if non-nil, zero value otherwise.

### GetCurrentTimeOk

`func (o *ControllerServicesEntity) GetCurrentTimeOk() (*string, bool)`

GetCurrentTimeOk returns a tuple with the CurrentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTime

`func (o *ControllerServicesEntity) SetCurrentTime(v string)`

SetCurrentTime sets CurrentTime field to given value.

### HasCurrentTime

`func (o *ControllerServicesEntity) HasCurrentTime() bool`

HasCurrentTime returns a boolean if a field has been set.

### GetControllerServices

`func (o *ControllerServicesEntity) GetControllerServices() []ControllerServiceEntity`

GetControllerServices returns the ControllerServices field if non-nil, zero value otherwise.

### GetControllerServicesOk

`func (o *ControllerServicesEntity) GetControllerServicesOk() (*[]ControllerServiceEntity, bool)`

GetControllerServicesOk returns a tuple with the ControllerServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerServices

`func (o *ControllerServicesEntity) SetControllerServices(v []ControllerServiceEntity)`

SetControllerServices sets ControllerServices field to given value.

### HasControllerServices

`func (o *ControllerServicesEntity) HasControllerServices() bool`

HasControllerServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


