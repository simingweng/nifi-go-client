# NodeEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Node** | Pointer to [**NodeDTO**](NodeDTO.md) |  | [optional] 

## Methods

### NewNodeEntity

`func NewNodeEntity() *NodeEntity`

NewNodeEntity instantiates a new NodeEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeEntityWithDefaults

`func NewNodeEntityWithDefaults() *NodeEntity`

NewNodeEntityWithDefaults instantiates a new NodeEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNode

`func (o *NodeEntity) GetNode() NodeDTO`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *NodeEntity) GetNodeOk() (*NodeDTO, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *NodeEntity) SetNode(v NodeDTO)`

SetNode sets Node field to given value.

### HasNode

`func (o *NodeEntity) HasNode() bool`

HasNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


