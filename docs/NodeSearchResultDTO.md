# NodeSearchResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the node that matched the search. | [optional] 
**Address** | Pointer to **string** | The address of the node that matched the search. | [optional] 

## Methods

### NewNodeSearchResultDTO

`func NewNodeSearchResultDTO() *NodeSearchResultDTO`

NewNodeSearchResultDTO instantiates a new NodeSearchResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeSearchResultDTOWithDefaults

`func NewNodeSearchResultDTOWithDefaults() *NodeSearchResultDTO`

NewNodeSearchResultDTOWithDefaults instantiates a new NodeSearchResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodeSearchResultDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeSearchResultDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeSearchResultDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NodeSearchResultDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddress

`func (o *NodeSearchResultDTO) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NodeSearchResultDTO) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NodeSearchResultDTO) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NodeSearchResultDTO) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


