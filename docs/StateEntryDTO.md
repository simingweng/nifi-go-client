# StateEntryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The key for this state. | [optional] 
**Value** | Pointer to **string** | The value for this state. | [optional] 
**ClusterNodeId** | Pointer to **string** | The identifier for the node where the state originated. | [optional] 
**ClusterNodeAddress** | Pointer to **string** | The label for the node where the state originated. | [optional] 

## Methods

### NewStateEntryDTO

`func NewStateEntryDTO() *StateEntryDTO`

NewStateEntryDTO instantiates a new StateEntryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateEntryDTOWithDefaults

`func NewStateEntryDTOWithDefaults() *StateEntryDTO`

NewStateEntryDTOWithDefaults instantiates a new StateEntryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *StateEntryDTO) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *StateEntryDTO) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *StateEntryDTO) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *StateEntryDTO) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *StateEntryDTO) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StateEntryDTO) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StateEntryDTO) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *StateEntryDTO) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetClusterNodeId

`func (o *StateEntryDTO) GetClusterNodeId() string`

GetClusterNodeId returns the ClusterNodeId field if non-nil, zero value otherwise.

### GetClusterNodeIdOk

`func (o *StateEntryDTO) GetClusterNodeIdOk() (*string, bool)`

GetClusterNodeIdOk returns a tuple with the ClusterNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNodeId

`func (o *StateEntryDTO) SetClusterNodeId(v string)`

SetClusterNodeId sets ClusterNodeId field to given value.

### HasClusterNodeId

`func (o *StateEntryDTO) HasClusterNodeId() bool`

HasClusterNodeId returns a boolean if a field has been set.

### GetClusterNodeAddress

`func (o *StateEntryDTO) GetClusterNodeAddress() string`

GetClusterNodeAddress returns the ClusterNodeAddress field if non-nil, zero value otherwise.

### GetClusterNodeAddressOk

`func (o *StateEntryDTO) GetClusterNodeAddressOk() (*string, bool)`

GetClusterNodeAddressOk returns a tuple with the ClusterNodeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNodeAddress

`func (o *StateEntryDTO) SetClusterNodeAddress(v string)`

SetClusterNodeAddress sets ClusterNodeAddress field to given value.

### HasClusterNodeAddress

`func (o *StateEntryDTO) HasClusterNodeAddress() bool`

HasClusterNodeAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


