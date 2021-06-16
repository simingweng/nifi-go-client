# AffectedComponentDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessGroupId** | Pointer to **string** | The UUID of the Process Group that this component is in | [optional] 
**Id** | Pointer to **string** | The UUID of this component | [optional] 
**ReferenceType** | Pointer to **string** | The type of this component | [optional] 
**Name** | Pointer to **string** | The name of this component. | [optional] 
**State** | Pointer to **string** | The scheduled state of a processor or reporting task referencing a controller service. If this component is another controller service, this field represents the controller service state. | [optional] 
**ActiveThreadCount** | Pointer to **int32** | The number of active threads for the referencing component. | [optional] 
**ValidationErrors** | Pointer to **[]string** | The validation errors for the component. | [optional] 

## Methods

### NewAffectedComponentDTO

`func NewAffectedComponentDTO() *AffectedComponentDTO`

NewAffectedComponentDTO instantiates a new AffectedComponentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffectedComponentDTOWithDefaults

`func NewAffectedComponentDTOWithDefaults() *AffectedComponentDTO`

NewAffectedComponentDTOWithDefaults instantiates a new AffectedComponentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessGroupId

`func (o *AffectedComponentDTO) GetProcessGroupId() string`

GetProcessGroupId returns the ProcessGroupId field if non-nil, zero value otherwise.

### GetProcessGroupIdOk

`func (o *AffectedComponentDTO) GetProcessGroupIdOk() (*string, bool)`

GetProcessGroupIdOk returns a tuple with the ProcessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupId

`func (o *AffectedComponentDTO) SetProcessGroupId(v string)`

SetProcessGroupId sets ProcessGroupId field to given value.

### HasProcessGroupId

`func (o *AffectedComponentDTO) HasProcessGroupId() bool`

HasProcessGroupId returns a boolean if a field has been set.

### GetId

`func (o *AffectedComponentDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AffectedComponentDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AffectedComponentDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AffectedComponentDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferenceType

`func (o *AffectedComponentDTO) GetReferenceType() string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *AffectedComponentDTO) GetReferenceTypeOk() (*string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *AffectedComponentDTO) SetReferenceType(v string)`

SetReferenceType sets ReferenceType field to given value.

### HasReferenceType

`func (o *AffectedComponentDTO) HasReferenceType() bool`

HasReferenceType returns a boolean if a field has been set.

### GetName

`func (o *AffectedComponentDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AffectedComponentDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AffectedComponentDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AffectedComponentDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *AffectedComponentDTO) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AffectedComponentDTO) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AffectedComponentDTO) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AffectedComponentDTO) HasState() bool`

HasState returns a boolean if a field has been set.

### GetActiveThreadCount

`func (o *AffectedComponentDTO) GetActiveThreadCount() int32`

GetActiveThreadCount returns the ActiveThreadCount field if non-nil, zero value otherwise.

### GetActiveThreadCountOk

`func (o *AffectedComponentDTO) GetActiveThreadCountOk() (*int32, bool)`

GetActiveThreadCountOk returns a tuple with the ActiveThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveThreadCount

`func (o *AffectedComponentDTO) SetActiveThreadCount(v int32)`

SetActiveThreadCount sets ActiveThreadCount field to given value.

### HasActiveThreadCount

`func (o *AffectedComponentDTO) HasActiveThreadCount() bool`

HasActiveThreadCount returns a boolean if a field has been set.

### GetValidationErrors

`func (o *AffectedComponentDTO) GetValidationErrors() []string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *AffectedComponentDTO) GetValidationErrorsOk() (*[]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *AffectedComponentDTO) SetValidationErrors(v []string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *AffectedComponentDTO) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


