# BulletinBoardDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bulletins** | Pointer to [**[]BulletinEntity**](BulletinEntity.md) | The bulletins in the bulletin board, that matches the supplied request. | [optional] 
**Generated** | Pointer to **string** | The timestamp when this report was generated. | [optional] 

## Methods

### NewBulletinBoardDTO

`func NewBulletinBoardDTO() *BulletinBoardDTO`

NewBulletinBoardDTO instantiates a new BulletinBoardDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulletinBoardDTOWithDefaults

`func NewBulletinBoardDTOWithDefaults() *BulletinBoardDTO`

NewBulletinBoardDTOWithDefaults instantiates a new BulletinBoardDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBulletins

`func (o *BulletinBoardDTO) GetBulletins() []BulletinEntity`

GetBulletins returns the Bulletins field if non-nil, zero value otherwise.

### GetBulletinsOk

`func (o *BulletinBoardDTO) GetBulletinsOk() (*[]BulletinEntity, bool)`

GetBulletinsOk returns a tuple with the Bulletins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletins

`func (o *BulletinBoardDTO) SetBulletins(v []BulletinEntity)`

SetBulletins sets Bulletins field to given value.

### HasBulletins

`func (o *BulletinBoardDTO) HasBulletins() bool`

HasBulletins returns a boolean if a field has been set.

### GetGenerated

`func (o *BulletinBoardDTO) GetGenerated() string`

GetGenerated returns the Generated field if non-nil, zero value otherwise.

### GetGeneratedOk

`func (o *BulletinBoardDTO) GetGeneratedOk() (*string, bool)`

GetGeneratedOk returns a tuple with the Generated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerated

`func (o *BulletinBoardDTO) SetGenerated(v string)`

SetGenerated sets Generated field to given value.

### HasGenerated

`func (o *BulletinBoardDTO) HasGenerated() bool`

HasGenerated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


