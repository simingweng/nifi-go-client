# ParameterContextValidationRequestEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | Pointer to [**ParameterContextValidationRequestDTO**](ParameterContextValidationRequestDTO.md) |  | [optional] 
**DisconnectedNodeAcknowledged** | Pointer to **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] 

## Methods

### NewParameterContextValidationRequestEntity

`func NewParameterContextValidationRequestEntity() *ParameterContextValidationRequestEntity`

NewParameterContextValidationRequestEntity instantiates a new ParameterContextValidationRequestEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterContextValidationRequestEntityWithDefaults

`func NewParameterContextValidationRequestEntityWithDefaults() *ParameterContextValidationRequestEntity`

NewParameterContextValidationRequestEntityWithDefaults instantiates a new ParameterContextValidationRequestEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *ParameterContextValidationRequestEntity) GetRequest() ParameterContextValidationRequestDTO`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *ParameterContextValidationRequestEntity) GetRequestOk() (*ParameterContextValidationRequestDTO, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *ParameterContextValidationRequestEntity) SetRequest(v ParameterContextValidationRequestDTO)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *ParameterContextValidationRequestEntity) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetDisconnectedNodeAcknowledged

`func (o *ParameterContextValidationRequestEntity) GetDisconnectedNodeAcknowledged() bool`

GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field if non-nil, zero value otherwise.

### GetDisconnectedNodeAcknowledgedOk

`func (o *ParameterContextValidationRequestEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool)`

GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedNodeAcknowledged

`func (o *ParameterContextValidationRequestEntity) SetDisconnectedNodeAcknowledged(v bool)`

SetDisconnectedNodeAcknowledged sets DisconnectedNodeAcknowledged field to given value.

### HasDisconnectedNodeAcknowledged

`func (o *ParameterContextValidationRequestEntity) HasDisconnectedNodeAcknowledged() bool`

HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


