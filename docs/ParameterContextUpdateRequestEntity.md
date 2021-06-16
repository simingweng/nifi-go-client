# ParameterContextUpdateRequestEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParameterContextRevision** | Pointer to [**RevisionDTO**](RevisionDTO.md) |  | [optional] 
**Request** | Pointer to [**ParameterContextUpdateRequestDTO**](ParameterContextUpdateRequestDTO.md) |  | [optional] 

## Methods

### NewParameterContextUpdateRequestEntity

`func NewParameterContextUpdateRequestEntity() *ParameterContextUpdateRequestEntity`

NewParameterContextUpdateRequestEntity instantiates a new ParameterContextUpdateRequestEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterContextUpdateRequestEntityWithDefaults

`func NewParameterContextUpdateRequestEntityWithDefaults() *ParameterContextUpdateRequestEntity`

NewParameterContextUpdateRequestEntityWithDefaults instantiates a new ParameterContextUpdateRequestEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameterContextRevision

`func (o *ParameterContextUpdateRequestEntity) GetParameterContextRevision() RevisionDTO`

GetParameterContextRevision returns the ParameterContextRevision field if non-nil, zero value otherwise.

### GetParameterContextRevisionOk

`func (o *ParameterContextUpdateRequestEntity) GetParameterContextRevisionOk() (*RevisionDTO, bool)`

GetParameterContextRevisionOk returns a tuple with the ParameterContextRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterContextRevision

`func (o *ParameterContextUpdateRequestEntity) SetParameterContextRevision(v RevisionDTO)`

SetParameterContextRevision sets ParameterContextRevision field to given value.

### HasParameterContextRevision

`func (o *ParameterContextUpdateRequestEntity) HasParameterContextRevision() bool`

HasParameterContextRevision returns a boolean if a field has been set.

### GetRequest

`func (o *ParameterContextUpdateRequestEntity) GetRequest() ParameterContextUpdateRequestDTO`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *ParameterContextUpdateRequestEntity) GetRequestOk() (*ParameterContextUpdateRequestDTO, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *ParameterContextUpdateRequestEntity) SetRequest(v ParameterContextUpdateRequestDTO)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *ParameterContextUpdateRequestEntity) HasRequest() bool`

HasRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


