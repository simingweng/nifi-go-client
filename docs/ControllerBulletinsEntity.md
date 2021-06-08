# ControllerBulletinsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bulletins** | Pointer to [**[]BulletinEntity**](BulletinEntity.md) | System level bulletins to be reported to the user. | [optional] 
**ControllerServiceBulletins** | Pointer to [**[]BulletinEntity**](BulletinEntity.md) | Controller service bulletins to be reported to the user. | [optional] 
**ReportingTaskBulletins** | Pointer to [**[]BulletinEntity**](BulletinEntity.md) | Reporting task bulletins to be reported to the user. | [optional] 

## Methods

### NewControllerBulletinsEntity

`func NewControllerBulletinsEntity() *ControllerBulletinsEntity`

NewControllerBulletinsEntity instantiates a new ControllerBulletinsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerBulletinsEntityWithDefaults

`func NewControllerBulletinsEntityWithDefaults() *ControllerBulletinsEntity`

NewControllerBulletinsEntityWithDefaults instantiates a new ControllerBulletinsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBulletins

`func (o *ControllerBulletinsEntity) GetBulletins() []BulletinEntity`

GetBulletins returns the Bulletins field if non-nil, zero value otherwise.

### GetBulletinsOk

`func (o *ControllerBulletinsEntity) GetBulletinsOk() (*[]BulletinEntity, bool)`

GetBulletinsOk returns a tuple with the Bulletins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletins

`func (o *ControllerBulletinsEntity) SetBulletins(v []BulletinEntity)`

SetBulletins sets Bulletins field to given value.

### HasBulletins

`func (o *ControllerBulletinsEntity) HasBulletins() bool`

HasBulletins returns a boolean if a field has been set.

### GetControllerServiceBulletins

`func (o *ControllerBulletinsEntity) GetControllerServiceBulletins() []BulletinEntity`

GetControllerServiceBulletins returns the ControllerServiceBulletins field if non-nil, zero value otherwise.

### GetControllerServiceBulletinsOk

`func (o *ControllerBulletinsEntity) GetControllerServiceBulletinsOk() (*[]BulletinEntity, bool)`

GetControllerServiceBulletinsOk returns a tuple with the ControllerServiceBulletins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerServiceBulletins

`func (o *ControllerBulletinsEntity) SetControllerServiceBulletins(v []BulletinEntity)`

SetControllerServiceBulletins sets ControllerServiceBulletins field to given value.

### HasControllerServiceBulletins

`func (o *ControllerBulletinsEntity) HasControllerServiceBulletins() bool`

HasControllerServiceBulletins returns a boolean if a field has been set.

### GetReportingTaskBulletins

`func (o *ControllerBulletinsEntity) GetReportingTaskBulletins() []BulletinEntity`

GetReportingTaskBulletins returns the ReportingTaskBulletins field if non-nil, zero value otherwise.

### GetReportingTaskBulletinsOk

`func (o *ControllerBulletinsEntity) GetReportingTaskBulletinsOk() (*[]BulletinEntity, bool)`

GetReportingTaskBulletinsOk returns a tuple with the ReportingTaskBulletins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingTaskBulletins

`func (o *ControllerBulletinsEntity) SetReportingTaskBulletins(v []BulletinEntity)`

SetReportingTaskBulletins sets ReportingTaskBulletins field to given value.

### HasReportingTaskBulletins

`func (o *ControllerBulletinsEntity) HasReportingTaskBulletins() bool`

HasReportingTaskBulletins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


