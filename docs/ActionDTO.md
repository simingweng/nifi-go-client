# ActionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The action id. | [optional] 
**UserIdentity** | Pointer to **string** | The identity of the user that performed the action. | [optional] 
**Timestamp** | Pointer to **string** | The timestamp of the action. | [optional] 
**SourceId** | Pointer to **string** | The id of the source component. | [optional] 
**SourceName** | Pointer to **string** | The name of the source component. | [optional] 
**SourceType** | Pointer to **string** | The type of the source component. | [optional] 
**ComponentDetails** | Pointer to **map[string]interface{}** |  | [optional] 
**Operation** | Pointer to **string** | The operation that was performed. | [optional] 
**ActionDetails** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewActionDTO

`func NewActionDTO() *ActionDTO`

NewActionDTO instantiates a new ActionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionDTOWithDefaults

`func NewActionDTOWithDefaults() *ActionDTO`

NewActionDTOWithDefaults instantiates a new ActionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionDTO) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionDTO) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionDTO) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ActionDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserIdentity

`func (o *ActionDTO) GetUserIdentity() string`

GetUserIdentity returns the UserIdentity field if non-nil, zero value otherwise.

### GetUserIdentityOk

`func (o *ActionDTO) GetUserIdentityOk() (*string, bool)`

GetUserIdentityOk returns a tuple with the UserIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentity

`func (o *ActionDTO) SetUserIdentity(v string)`

SetUserIdentity sets UserIdentity field to given value.

### HasUserIdentity

`func (o *ActionDTO) HasUserIdentity() bool`

HasUserIdentity returns a boolean if a field has been set.

### GetTimestamp

`func (o *ActionDTO) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ActionDTO) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ActionDTO) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ActionDTO) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetSourceId

`func (o *ActionDTO) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ActionDTO) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ActionDTO) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ActionDTO) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceName

`func (o *ActionDTO) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *ActionDTO) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *ActionDTO) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *ActionDTO) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetSourceType

`func (o *ActionDTO) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ActionDTO) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ActionDTO) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *ActionDTO) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetComponentDetails

`func (o *ActionDTO) GetComponentDetails() map[string]interface{}`

GetComponentDetails returns the ComponentDetails field if non-nil, zero value otherwise.

### GetComponentDetailsOk

`func (o *ActionDTO) GetComponentDetailsOk() (*map[string]interface{}, bool)`

GetComponentDetailsOk returns a tuple with the ComponentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentDetails

`func (o *ActionDTO) SetComponentDetails(v map[string]interface{})`

SetComponentDetails sets ComponentDetails field to given value.

### HasComponentDetails

`func (o *ActionDTO) HasComponentDetails() bool`

HasComponentDetails returns a boolean if a field has been set.

### GetOperation

`func (o *ActionDTO) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ActionDTO) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ActionDTO) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *ActionDTO) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetActionDetails

`func (o *ActionDTO) GetActionDetails() map[string]interface{}`

GetActionDetails returns the ActionDetails field if non-nil, zero value otherwise.

### GetActionDetailsOk

`func (o *ActionDTO) GetActionDetailsOk() (*map[string]interface{}, bool)`

GetActionDetailsOk returns a tuple with the ActionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDetails

`func (o *ActionDTO) SetActionDetails(v map[string]interface{})`

SetActionDetails sets ActionDetails field to given value.

### HasActionDetails

`func (o *ActionDTO) HasActionDetails() bool`

HasActionDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


