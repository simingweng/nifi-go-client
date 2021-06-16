# RevisionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | A client identifier used to make a request. By including a client identifier, the API can allow multiple requests without needing the current revision. Due to the asynchronous nature of requests/responses this was implemented to allow the client to make numerous requests without having to wait for the previous response to come back. | [optional] 
**Version** | Pointer to **int64** | NiFi Registry employs an optimistic locking strategy where the client must include a revision in their request when performing an update. In a response to a mutable flow request, this field represents the updated base version. | [optional] 
**LastModifier** | Pointer to **string** | The user that last modified the entity. | [optional] [readonly] 

## Methods

### NewRevisionInfo

`func NewRevisionInfo() *RevisionInfo`

NewRevisionInfo instantiates a new RevisionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevisionInfoWithDefaults

`func NewRevisionInfoWithDefaults() *RevisionInfo`

NewRevisionInfoWithDefaults instantiates a new RevisionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *RevisionInfo) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *RevisionInfo) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *RevisionInfo) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *RevisionInfo) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetVersion

`func (o *RevisionInfo) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RevisionInfo) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RevisionInfo) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RevisionInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLastModifier

`func (o *RevisionInfo) GetLastModifier() string`

GetLastModifier returns the LastModifier field if non-nil, zero value otherwise.

### GetLastModifierOk

`func (o *RevisionInfo) GetLastModifierOk() (*string, bool)`

GetLastModifierOk returns a tuple with the LastModifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifier

`func (o *RevisionInfo) SetLastModifier(v string)`

SetLastModifier sets LastModifier field to given value.

### HasLastModifier

`func (o *RevisionInfo) HasLastModifier() bool`

HasLastModifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


