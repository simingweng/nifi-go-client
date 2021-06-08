# TemplateDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | Pointer to **string** | The URI for the template. | [optional] 
**Id** | Pointer to **string** | The id of the template. | [optional] 
**GroupId** | Pointer to **string** | The id of the Process Group that the template belongs to. | [optional] 
**Name** | Pointer to **string** | The name of the template. | [optional] 
**Description** | Pointer to **string** | The description of the template. | [optional] 
**Timestamp** | Pointer to **string** | The timestamp when this template was created. | [optional] 
**EncodingVersion** | Pointer to **string** | The encoding version of this template. | [optional] 
**Snippet** | Pointer to [**FlowSnippetDTO**](FlowSnippetDTO.md) |  | [optional] 

## Methods

### NewTemplateDTO

`func NewTemplateDTO() *TemplateDTO`

NewTemplateDTO instantiates a new TemplateDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateDTOWithDefaults

`func NewTemplateDTOWithDefaults() *TemplateDTO`

NewTemplateDTOWithDefaults instantiates a new TemplateDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *TemplateDTO) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *TemplateDTO) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *TemplateDTO) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *TemplateDTO) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetId

`func (o *TemplateDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupId

`func (o *TemplateDTO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *TemplateDTO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *TemplateDTO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *TemplateDTO) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *TemplateDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TemplateDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTimestamp

`func (o *TemplateDTO) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TemplateDTO) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TemplateDTO) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TemplateDTO) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetEncodingVersion

`func (o *TemplateDTO) GetEncodingVersion() string`

GetEncodingVersion returns the EncodingVersion field if non-nil, zero value otherwise.

### GetEncodingVersionOk

`func (o *TemplateDTO) GetEncodingVersionOk() (*string, bool)`

GetEncodingVersionOk returns a tuple with the EncodingVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingVersion

`func (o *TemplateDTO) SetEncodingVersion(v string)`

SetEncodingVersion sets EncodingVersion field to given value.

### HasEncodingVersion

`func (o *TemplateDTO) HasEncodingVersion() bool`

HasEncodingVersion returns a boolean if a field has been set.

### GetSnippet

`func (o *TemplateDTO) GetSnippet() FlowSnippetDTO`

GetSnippet returns the Snippet field if non-nil, zero value otherwise.

### GetSnippetOk

`func (o *TemplateDTO) GetSnippetOk() (*FlowSnippetDTO, bool)`

GetSnippetOk returns a tuple with the Snippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippet

`func (o *TemplateDTO) SetSnippet(v FlowSnippetDTO)`

SetSnippet sets Snippet field to given value.

### HasSnippet

`func (o *TemplateDTO) HasSnippet() bool`

HasSnippet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


