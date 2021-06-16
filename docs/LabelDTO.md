# LabelDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the component. | [optional] 
**VersionedComponentId** | Pointer to **string** | The ID of the corresponding component that is under version control | [optional] 
**ParentGroupId** | Pointer to **string** | The id of parent process group of this component if applicable. | [optional] 
**Position** | Pointer to [**PositionDTO**](PositionDTO.md) |  | [optional] 
**Label** | Pointer to **string** | The text that appears in the label. | [optional] 
**Width** | Pointer to **float64** | The width of the label in pixels when at a 1:1 scale. | [optional] 
**Height** | Pointer to **float64** | The height of the label in pixels when at a 1:1 scale. | [optional] 
**Style** | Pointer to **map[string]string** | The styles for this label (font-size : 12px, background-color : #eee, etc). | [optional] 

## Methods

### NewLabelDTO

`func NewLabelDTO() *LabelDTO`

NewLabelDTO instantiates a new LabelDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelDTOWithDefaults

`func NewLabelDTOWithDefaults() *LabelDTO`

NewLabelDTOWithDefaults instantiates a new LabelDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LabelDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LabelDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LabelDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LabelDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersionedComponentId

`func (o *LabelDTO) GetVersionedComponentId() string`

GetVersionedComponentId returns the VersionedComponentId field if non-nil, zero value otherwise.

### GetVersionedComponentIdOk

`func (o *LabelDTO) GetVersionedComponentIdOk() (*string, bool)`

GetVersionedComponentIdOk returns a tuple with the VersionedComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedComponentId

`func (o *LabelDTO) SetVersionedComponentId(v string)`

SetVersionedComponentId sets VersionedComponentId field to given value.

### HasVersionedComponentId

`func (o *LabelDTO) HasVersionedComponentId() bool`

HasVersionedComponentId returns a boolean if a field has been set.

### GetParentGroupId

`func (o *LabelDTO) GetParentGroupId() string`

GetParentGroupId returns the ParentGroupId field if non-nil, zero value otherwise.

### GetParentGroupIdOk

`func (o *LabelDTO) GetParentGroupIdOk() (*string, bool)`

GetParentGroupIdOk returns a tuple with the ParentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroupId

`func (o *LabelDTO) SetParentGroupId(v string)`

SetParentGroupId sets ParentGroupId field to given value.

### HasParentGroupId

`func (o *LabelDTO) HasParentGroupId() bool`

HasParentGroupId returns a boolean if a field has been set.

### GetPosition

`func (o *LabelDTO) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *LabelDTO) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *LabelDTO) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *LabelDTO) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetLabel

`func (o *LabelDTO) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LabelDTO) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LabelDTO) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *LabelDTO) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetWidth

`func (o *LabelDTO) GetWidth() float64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *LabelDTO) GetWidthOk() (*float64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *LabelDTO) SetWidth(v float64)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *LabelDTO) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *LabelDTO) GetHeight() float64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *LabelDTO) GetHeightOk() (*float64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *LabelDTO) SetHeight(v float64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *LabelDTO) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetStyle

`func (o *LabelDTO) GetStyle() map[string]string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *LabelDTO) GetStyleOk() (*map[string]string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *LabelDTO) SetStyle(v map[string]string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *LabelDTO) HasStyle() bool`

HasStyle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


