# FlowBreadcrumbDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the group. | [optional] 
**Name** | Pointer to **string** | The id of the group. | [optional] 
**VersionControlInformation** | Pointer to [**VersionControlInformationDTO**](VersionControlInformationDTO.md) |  | [optional] 

## Methods

### NewFlowBreadcrumbDTO

`func NewFlowBreadcrumbDTO() *FlowBreadcrumbDTO`

NewFlowBreadcrumbDTO instantiates a new FlowBreadcrumbDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowBreadcrumbDTOWithDefaults

`func NewFlowBreadcrumbDTOWithDefaults() *FlowBreadcrumbDTO`

NewFlowBreadcrumbDTOWithDefaults instantiates a new FlowBreadcrumbDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FlowBreadcrumbDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowBreadcrumbDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowBreadcrumbDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FlowBreadcrumbDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FlowBreadcrumbDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlowBreadcrumbDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlowBreadcrumbDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlowBreadcrumbDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersionControlInformation

`func (o *FlowBreadcrumbDTO) GetVersionControlInformation() VersionControlInformationDTO`

GetVersionControlInformation returns the VersionControlInformation field if non-nil, zero value otherwise.

### GetVersionControlInformationOk

`func (o *FlowBreadcrumbDTO) GetVersionControlInformationOk() (*VersionControlInformationDTO, bool)`

GetVersionControlInformationOk returns a tuple with the VersionControlInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionControlInformation

`func (o *FlowBreadcrumbDTO) SetVersionControlInformation(v VersionControlInformationDTO)`

SetVersionControlInformation sets VersionControlInformation field to given value.

### HasVersionControlInformation

`func (o *FlowBreadcrumbDTO) HasVersionControlInformation() bool`

HasVersionControlInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


