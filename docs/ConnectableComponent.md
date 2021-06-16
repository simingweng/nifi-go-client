# ConnectableComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the connectable component. | 
**Type** | **string** | The type of component the connectable is. | 
**GroupId** | **string** | The id of the group that the connectable component resides in | 
**Name** | Pointer to **string** | The name of the connectable component | [optional] 
**Comments** | Pointer to **string** | The comments for the connectable component. | [optional] 

## Methods

### NewConnectableComponent

`func NewConnectableComponent(id string, type_ string, groupId string, ) *ConnectableComponent`

NewConnectableComponent instantiates a new ConnectableComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectableComponentWithDefaults

`func NewConnectableComponentWithDefaults() *ConnectableComponent`

NewConnectableComponentWithDefaults instantiates a new ConnectableComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectableComponent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectableComponent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectableComponent) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ConnectableComponent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectableComponent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectableComponent) SetType(v string)`

SetType sets Type field to given value.


### GetGroupId

`func (o *ConnectableComponent) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ConnectableComponent) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ConnectableComponent) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetName

`func (o *ConnectableComponent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectableComponent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectableComponent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectableComponent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComments

`func (o *ConnectableComponent) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ConnectableComponent) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ConnectableComponent) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ConnectableComponent) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


