# AccessStatusDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | Pointer to **string** | The user identity. | [optional] [readonly] 
**Status** | Pointer to **string** | The user access status. | [optional] [readonly] 
**Message** | Pointer to **string** | Additional details about the user access status. | [optional] [readonly] 

## Methods

### NewAccessStatusDTO

`func NewAccessStatusDTO() *AccessStatusDTO`

NewAccessStatusDTO instantiates a new AccessStatusDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessStatusDTOWithDefaults

`func NewAccessStatusDTOWithDefaults() *AccessStatusDTO`

NewAccessStatusDTOWithDefaults instantiates a new AccessStatusDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *AccessStatusDTO) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *AccessStatusDTO) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *AccessStatusDTO) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *AccessStatusDTO) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetStatus

`func (o *AccessStatusDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccessStatusDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccessStatusDTO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccessStatusDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *AccessStatusDTO) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AccessStatusDTO) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AccessStatusDTO) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AccessStatusDTO) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


