# BulletinDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The id of the bulletin. | [optional] 
**NodeAddress** | Pointer to **string** | If clustered, the address of the node from which the bulletin originated. | [optional] 
**Category** | Pointer to **string** | The category of this bulletin. | [optional] 
**GroupId** | Pointer to **string** | The group id of the source component. | [optional] 
**SourceId** | Pointer to **string** | The id of the source component. | [optional] 
**SourceName** | Pointer to **string** | The name of the source component. | [optional] 
**Level** | Pointer to **string** | The level of the bulletin. | [optional] 
**Message** | Pointer to **string** | The bulletin message. | [optional] 
**Timestamp** | Pointer to **string** | When this bulletin was generated. | [optional] 

## Methods

### NewBulletinDTO

`func NewBulletinDTO() *BulletinDTO`

NewBulletinDTO instantiates a new BulletinDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulletinDTOWithDefaults

`func NewBulletinDTOWithDefaults() *BulletinDTO`

NewBulletinDTOWithDefaults instantiates a new BulletinDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulletinDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulletinDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulletinDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BulletinDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNodeAddress

`func (o *BulletinDTO) GetNodeAddress() string`

GetNodeAddress returns the NodeAddress field if non-nil, zero value otherwise.

### GetNodeAddressOk

`func (o *BulletinDTO) GetNodeAddressOk() (*string, bool)`

GetNodeAddressOk returns a tuple with the NodeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAddress

`func (o *BulletinDTO) SetNodeAddress(v string)`

SetNodeAddress sets NodeAddress field to given value.

### HasNodeAddress

`func (o *BulletinDTO) HasNodeAddress() bool`

HasNodeAddress returns a boolean if a field has been set.

### GetCategory

`func (o *BulletinDTO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *BulletinDTO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *BulletinDTO) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *BulletinDTO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetGroupId

`func (o *BulletinDTO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *BulletinDTO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *BulletinDTO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *BulletinDTO) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetSourceId

`func (o *BulletinDTO) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BulletinDTO) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BulletinDTO) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *BulletinDTO) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceName

`func (o *BulletinDTO) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *BulletinDTO) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *BulletinDTO) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *BulletinDTO) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetLevel

`func (o *BulletinDTO) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *BulletinDTO) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *BulletinDTO) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *BulletinDTO) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetMessage

`func (o *BulletinDTO) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BulletinDTO) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BulletinDTO) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BulletinDTO) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTimestamp

`func (o *BulletinDTO) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BulletinDTO) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BulletinDTO) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BulletinDTO) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


