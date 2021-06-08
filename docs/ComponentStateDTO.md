# ComponentStateDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentId** | Pointer to **string** | The component identifier. | [optional] 
**StateDescription** | Pointer to **string** | Description of the state this component persists. | [optional] 
**ClusterState** | Pointer to [**StateMapDTO**](StateMapDTO.md) |  | [optional] 
**LocalState** | Pointer to [**StateMapDTO**](StateMapDTO.md) |  | [optional] 

## Methods

### NewComponentStateDTO

`func NewComponentStateDTO() *ComponentStateDTO`

NewComponentStateDTO instantiates a new ComponentStateDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentStateDTOWithDefaults

`func NewComponentStateDTOWithDefaults() *ComponentStateDTO`

NewComponentStateDTOWithDefaults instantiates a new ComponentStateDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentId

`func (o *ComponentStateDTO) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *ComponentStateDTO) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *ComponentStateDTO) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *ComponentStateDTO) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### GetStateDescription

`func (o *ComponentStateDTO) GetStateDescription() string`

GetStateDescription returns the StateDescription field if non-nil, zero value otherwise.

### GetStateDescriptionOk

`func (o *ComponentStateDTO) GetStateDescriptionOk() (*string, bool)`

GetStateDescriptionOk returns a tuple with the StateDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateDescription

`func (o *ComponentStateDTO) SetStateDescription(v string)`

SetStateDescription sets StateDescription field to given value.

### HasStateDescription

`func (o *ComponentStateDTO) HasStateDescription() bool`

HasStateDescription returns a boolean if a field has been set.

### GetClusterState

`func (o *ComponentStateDTO) GetClusterState() StateMapDTO`

GetClusterState returns the ClusterState field if non-nil, zero value otherwise.

### GetClusterStateOk

`func (o *ComponentStateDTO) GetClusterStateOk() (*StateMapDTO, bool)`

GetClusterStateOk returns a tuple with the ClusterState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterState

`func (o *ComponentStateDTO) SetClusterState(v StateMapDTO)`

SetClusterState sets ClusterState field to given value.

### HasClusterState

`func (o *ComponentStateDTO) HasClusterState() bool`

HasClusterState returns a boolean if a field has been set.

### GetLocalState

`func (o *ComponentStateDTO) GetLocalState() StateMapDTO`

GetLocalState returns the LocalState field if non-nil, zero value otherwise.

### GetLocalStateOk

`func (o *ComponentStateDTO) GetLocalStateOk() (*StateMapDTO, bool)`

GetLocalStateOk returns a tuple with the LocalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalState

`func (o *ComponentStateDTO) SetLocalState(v StateMapDTO)`

SetLocalState sets LocalState field to given value.

### HasLocalState

`func (o *ComponentStateDTO) HasLocalState() bool`

HasLocalState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


