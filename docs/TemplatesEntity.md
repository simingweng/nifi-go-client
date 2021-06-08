# TemplatesEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Templates** | Pointer to [**[]TemplateEntity**](TemplateEntity.md) |  | [optional] 
**Generated** | Pointer to **string** | When this content was generated. | [optional] 

## Methods

### NewTemplatesEntity

`func NewTemplatesEntity() *TemplatesEntity`

NewTemplatesEntity instantiates a new TemplatesEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesEntityWithDefaults

`func NewTemplatesEntityWithDefaults() *TemplatesEntity`

NewTemplatesEntityWithDefaults instantiates a new TemplatesEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplates

`func (o *TemplatesEntity) GetTemplates() []TemplateEntity`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *TemplatesEntity) GetTemplatesOk() (*[]TemplateEntity, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *TemplatesEntity) SetTemplates(v []TemplateEntity)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *TemplatesEntity) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### GetGenerated

`func (o *TemplatesEntity) GetGenerated() string`

GetGenerated returns the Generated field if non-nil, zero value otherwise.

### GetGeneratedOk

`func (o *TemplatesEntity) GetGeneratedOk() (*string, bool)`

GetGeneratedOk returns a tuple with the Generated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerated

`func (o *TemplatesEntity) SetGenerated(v string)`

SetGenerated sets Generated field to given value.

### HasGenerated

`func (o *TemplatesEntity) HasGenerated() bool`

HasGenerated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


