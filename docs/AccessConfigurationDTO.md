# AccessConfigurationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportsLogin** | Pointer to **bool** | Indicates whether or not this NiFi supports user login. | [optional] [readonly] 

## Methods

### NewAccessConfigurationDTO

`func NewAccessConfigurationDTO() *AccessConfigurationDTO`

NewAccessConfigurationDTO instantiates a new AccessConfigurationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessConfigurationDTOWithDefaults

`func NewAccessConfigurationDTOWithDefaults() *AccessConfigurationDTO`

NewAccessConfigurationDTOWithDefaults instantiates a new AccessConfigurationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportsLogin

`func (o *AccessConfigurationDTO) GetSupportsLogin() bool`

GetSupportsLogin returns the SupportsLogin field if non-nil, zero value otherwise.

### GetSupportsLoginOk

`func (o *AccessConfigurationDTO) GetSupportsLoginOk() (*bool, bool)`

GetSupportsLoginOk returns a tuple with the SupportsLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsLogin

`func (o *AccessConfigurationDTO) SetSupportsLogin(v bool)`

SetSupportsLogin sets SupportsLogin field to given value.

### HasSupportsLogin

`func (o *AccessConfigurationDTO) HasSupportsLogin() bool`

HasSupportsLogin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


