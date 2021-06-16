# ProcessGroupNameDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the Process Group | [optional] 
**Name** | Pointer to **string** | The name of the Process Group, or the ID of the Process Group if the user does not have the READ policy for the Process Group | [optional] 

## Methods

### NewProcessGroupNameDTO

`func NewProcessGroupNameDTO() *ProcessGroupNameDTO`

NewProcessGroupNameDTO instantiates a new ProcessGroupNameDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessGroupNameDTOWithDefaults

`func NewProcessGroupNameDTOWithDefaults() *ProcessGroupNameDTO`

NewProcessGroupNameDTOWithDefaults instantiates a new ProcessGroupNameDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProcessGroupNameDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessGroupNameDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessGroupNameDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProcessGroupNameDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProcessGroupNameDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProcessGroupNameDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProcessGroupNameDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProcessGroupNameDTO) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


