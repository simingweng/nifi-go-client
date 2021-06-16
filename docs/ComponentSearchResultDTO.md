# ComponentSearchResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the component that matched the search. | [optional] 
**GroupId** | Pointer to **string** | The group id of the component that matched the search. | [optional] 
**ParentGroup** | Pointer to [**SearchResultGroupDTO**](SearchResultGroupDTO.md) |  | [optional] 
**VersionedGroup** | Pointer to [**SearchResultGroupDTO**](SearchResultGroupDTO.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the component that matched the search. | [optional] 
**Matches** | Pointer to **[]string** | What matched the search from the component. | [optional] 

## Methods

### NewComponentSearchResultDTO

`func NewComponentSearchResultDTO() *ComponentSearchResultDTO`

NewComponentSearchResultDTO instantiates a new ComponentSearchResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentSearchResultDTOWithDefaults

`func NewComponentSearchResultDTOWithDefaults() *ComponentSearchResultDTO`

NewComponentSearchResultDTOWithDefaults instantiates a new ComponentSearchResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComponentSearchResultDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComponentSearchResultDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComponentSearchResultDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComponentSearchResultDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupId

`func (o *ComponentSearchResultDTO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ComponentSearchResultDTO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ComponentSearchResultDTO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ComponentSearchResultDTO) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetParentGroup

`func (o *ComponentSearchResultDTO) GetParentGroup() SearchResultGroupDTO`

GetParentGroup returns the ParentGroup field if non-nil, zero value otherwise.

### GetParentGroupOk

`func (o *ComponentSearchResultDTO) GetParentGroupOk() (*SearchResultGroupDTO, bool)`

GetParentGroupOk returns a tuple with the ParentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroup

`func (o *ComponentSearchResultDTO) SetParentGroup(v SearchResultGroupDTO)`

SetParentGroup sets ParentGroup field to given value.

### HasParentGroup

`func (o *ComponentSearchResultDTO) HasParentGroup() bool`

HasParentGroup returns a boolean if a field has been set.

### GetVersionedGroup

`func (o *ComponentSearchResultDTO) GetVersionedGroup() SearchResultGroupDTO`

GetVersionedGroup returns the VersionedGroup field if non-nil, zero value otherwise.

### GetVersionedGroupOk

`func (o *ComponentSearchResultDTO) GetVersionedGroupOk() (*SearchResultGroupDTO, bool)`

GetVersionedGroupOk returns a tuple with the VersionedGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedGroup

`func (o *ComponentSearchResultDTO) SetVersionedGroup(v SearchResultGroupDTO)`

SetVersionedGroup sets VersionedGroup field to given value.

### HasVersionedGroup

`func (o *ComponentSearchResultDTO) HasVersionedGroup() bool`

HasVersionedGroup returns a boolean if a field has been set.

### GetName

`func (o *ComponentSearchResultDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentSearchResultDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentSearchResultDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComponentSearchResultDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMatches

`func (o *ComponentSearchResultDTO) GetMatches() []string`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *ComponentSearchResultDTO) GetMatchesOk() (*[]string, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *ComponentSearchResultDTO) SetMatches(v []string)`

SetMatches sets Matches field to given value.

### HasMatches

`func (o *ComponentSearchResultDTO) HasMatches() bool`

HasMatches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


