# RegistryClientsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Registries** | Pointer to [**[]RegistryClientEntity**](RegistryClientEntity.md) |  | [optional] 

## Methods

### NewRegistryClientsEntity

`func NewRegistryClientsEntity() *RegistryClientsEntity`

NewRegistryClientsEntity instantiates a new RegistryClientsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryClientsEntityWithDefaults

`func NewRegistryClientsEntityWithDefaults() *RegistryClientsEntity`

NewRegistryClientsEntityWithDefaults instantiates a new RegistryClientsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistries

`func (o *RegistryClientsEntity) GetRegistries() []RegistryClientEntity`

GetRegistries returns the Registries field if non-nil, zero value otherwise.

### GetRegistriesOk

`func (o *RegistryClientsEntity) GetRegistriesOk() (*[]RegistryClientEntity, bool)`

GetRegistriesOk returns a tuple with the Registries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistries

`func (o *RegistryClientsEntity) SetRegistries(v []RegistryClientEntity)`

SetRegistries sets Registries field to given value.

### HasRegistries

`func (o *RegistryClientsEntity) HasRegistries() bool`

HasRegistries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


