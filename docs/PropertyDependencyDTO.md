# PropertyDependencyDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyName** | Pointer to **string** | The name of the property that is being depended upon | [optional] 
**DependentValues** | Pointer to **[]string** | The values for the property that satisfies the dependency, or null if the dependency is satisfied by the presence of any value for the associated property name | [optional] 

## Methods

### NewPropertyDependencyDTO

`func NewPropertyDependencyDTO() *PropertyDependencyDTO`

NewPropertyDependencyDTO instantiates a new PropertyDependencyDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyDependencyDTOWithDefaults

`func NewPropertyDependencyDTOWithDefaults() *PropertyDependencyDTO`

NewPropertyDependencyDTOWithDefaults instantiates a new PropertyDependencyDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyName

`func (o *PropertyDependencyDTO) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *PropertyDependencyDTO) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *PropertyDependencyDTO) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *PropertyDependencyDTO) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetDependentValues

`func (o *PropertyDependencyDTO) GetDependentValues() []string`

GetDependentValues returns the DependentValues field if non-nil, zero value otherwise.

### GetDependentValuesOk

`func (o *PropertyDependencyDTO) GetDependentValuesOk() (*[]string, bool)`

GetDependentValuesOk returns a tuple with the DependentValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentValues

`func (o *PropertyDependencyDTO) SetDependentValues(v []string)`

SetDependentValues sets DependentValues field to given value.

### HasDependentValues

`func (o *PropertyDependencyDTO) HasDependentValues() bool`

HasDependentValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


