# NodeEventDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **string** | The timestamp of the node event. | [optional] 
**Category** | Pointer to **string** | The category of the node event. | [optional] 
**Message** | Pointer to **string** | The message in the node event. | [optional] 

## Methods

### NewNodeEventDTO

`func NewNodeEventDTO() *NodeEventDTO`

NewNodeEventDTO instantiates a new NodeEventDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeEventDTOWithDefaults

`func NewNodeEventDTOWithDefaults() *NodeEventDTO`

NewNodeEventDTOWithDefaults instantiates a new NodeEventDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *NodeEventDTO) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *NodeEventDTO) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *NodeEventDTO) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *NodeEventDTO) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetCategory

`func (o *NodeEventDTO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *NodeEventDTO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *NodeEventDTO) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *NodeEventDTO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetMessage

`func (o *NodeEventDTO) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NodeEventDTO) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NodeEventDTO) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NodeEventDTO) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


