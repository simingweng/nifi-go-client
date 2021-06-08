# RegistryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The registry identifier | [optional] 
**Name** | Pointer to **string** | The registry name | [optional] 
**Description** | Pointer to **string** | The registry description | [optional] 
**Uri** | Pointer to **string** | The registry URI | [optional] 

## Methods

### NewRegistryDTO

`func NewRegistryDTO() *RegistryDTO`

NewRegistryDTO instantiates a new RegistryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryDTOWithDefaults

`func NewRegistryDTOWithDefaults() *RegistryDTO`

NewRegistryDTOWithDefaults instantiates a new RegistryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegistryDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegistryDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegistryDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RegistryDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RegistryDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegistryDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegistryDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegistryDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RegistryDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RegistryDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RegistryDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RegistryDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUri

`func (o *RegistryDTO) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *RegistryDTO) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *RegistryDTO) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *RegistryDTO) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


