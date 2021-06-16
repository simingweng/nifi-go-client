# UsersEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Generated** | Pointer to **string** | When this content was generated. | [optional] 
**Users** | Pointer to [**[]UserEntity**](UserEntity.md) |  | [optional] 

## Methods

### NewUsersEntity

`func NewUsersEntity() *UsersEntity`

NewUsersEntity instantiates a new UsersEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersEntityWithDefaults

`func NewUsersEntityWithDefaults() *UsersEntity`

NewUsersEntityWithDefaults instantiates a new UsersEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenerated

`func (o *UsersEntity) GetGenerated() string`

GetGenerated returns the Generated field if non-nil, zero value otherwise.

### GetGeneratedOk

`func (o *UsersEntity) GetGeneratedOk() (*string, bool)`

GetGeneratedOk returns a tuple with the Generated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerated

`func (o *UsersEntity) SetGenerated(v string)`

SetGenerated sets Generated field to given value.

### HasGenerated

`func (o *UsersEntity) HasGenerated() bool`

HasGenerated returns a boolean if a field has been set.

### GetUsers

`func (o *UsersEntity) GetUsers() []UserEntity`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UsersEntity) GetUsersOk() (*[]UserEntity, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UsersEntity) SetUsers(v []UserEntity)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UsersEntity) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


