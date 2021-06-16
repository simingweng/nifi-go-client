# TransactionResultEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowFileSent** | Pointer to **int32** |  | [optional] 
**ResponseCode** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewTransactionResultEntity

`func NewTransactionResultEntity() *TransactionResultEntity`

NewTransactionResultEntity instantiates a new TransactionResultEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionResultEntityWithDefaults

`func NewTransactionResultEntityWithDefaults() *TransactionResultEntity`

NewTransactionResultEntityWithDefaults instantiates a new TransactionResultEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowFileSent

`func (o *TransactionResultEntity) GetFlowFileSent() int32`

GetFlowFileSent returns the FlowFileSent field if non-nil, zero value otherwise.

### GetFlowFileSentOk

`func (o *TransactionResultEntity) GetFlowFileSentOk() (*int32, bool)`

GetFlowFileSentOk returns a tuple with the FlowFileSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFileSent

`func (o *TransactionResultEntity) SetFlowFileSent(v int32)`

SetFlowFileSent sets FlowFileSent field to given value.

### HasFlowFileSent

`func (o *TransactionResultEntity) HasFlowFileSent() bool`

HasFlowFileSent returns a boolean if a field has been set.

### GetResponseCode

`func (o *TransactionResultEntity) GetResponseCode() int32`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *TransactionResultEntity) GetResponseCodeOk() (*int32, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *TransactionResultEntity) SetResponseCode(v int32)`

SetResponseCode sets ResponseCode field to given value.

### HasResponseCode

`func (o *TransactionResultEntity) HasResponseCode() bool`

HasResponseCode returns a boolean if a field has been set.

### GetMessage

`func (o *TransactionResultEntity) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TransactionResultEntity) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TransactionResultEntity) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TransactionResultEntity) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


