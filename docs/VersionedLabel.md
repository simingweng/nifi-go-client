# VersionedLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | The component&#39;s unique identifier | [optional] 
**Name** | Pointer to **string** | The component&#39;s name | [optional] 
**Comments** | Pointer to **string** | The user-supplied comments for the component | [optional] 
**Position** | Pointer to [**Position**](Position.md) |  | [optional] 
**Label** | Pointer to **string** | The text that appears in the label. | [optional] 
**Width** | Pointer to **float64** | The width of the label in pixels when at a 1:1 scale. | [optional] 
**Height** | Pointer to **float64** | The height of the label in pixels when at a 1:1 scale. | [optional] 
**Style** | Pointer to **map[string]string** | The styles for this label (font-size : 12px, background-color : #eee, etc). | [optional] 
**ComponentType** | Pointer to **string** |  | [optional] 
**GroupIdentifier** | Pointer to **string** | The ID of the Process Group that this component belongs to | [optional] 

## Methods

### NewVersionedLabel

`func NewVersionedLabel() *VersionedLabel`

NewVersionedLabel instantiates a new VersionedLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionedLabelWithDefaults

`func NewVersionedLabelWithDefaults() *VersionedLabel`

NewVersionedLabelWithDefaults instantiates a new VersionedLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *VersionedLabel) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *VersionedLabel) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *VersionedLabel) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *VersionedLabel) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *VersionedLabel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VersionedLabel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VersionedLabel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VersionedLabel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComments

`func (o *VersionedLabel) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VersionedLabel) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VersionedLabel) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VersionedLabel) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetPosition

`func (o *VersionedLabel) GetPosition() Position`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *VersionedLabel) GetPositionOk() (*Position, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *VersionedLabel) SetPosition(v Position)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *VersionedLabel) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetLabel

`func (o *VersionedLabel) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VersionedLabel) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VersionedLabel) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *VersionedLabel) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetWidth

`func (o *VersionedLabel) GetWidth() float64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *VersionedLabel) GetWidthOk() (*float64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *VersionedLabel) SetWidth(v float64)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *VersionedLabel) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *VersionedLabel) GetHeight() float64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *VersionedLabel) GetHeightOk() (*float64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *VersionedLabel) SetHeight(v float64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *VersionedLabel) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetStyle

`func (o *VersionedLabel) GetStyle() map[string]string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *VersionedLabel) GetStyleOk() (*map[string]string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *VersionedLabel) SetStyle(v map[string]string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *VersionedLabel) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### GetComponentType

`func (o *VersionedLabel) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *VersionedLabel) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *VersionedLabel) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *VersionedLabel) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### GetGroupIdentifier

`func (o *VersionedLabel) GetGroupIdentifier() string`

GetGroupIdentifier returns the GroupIdentifier field if non-nil, zero value otherwise.

### GetGroupIdentifierOk

`func (o *VersionedLabel) GetGroupIdentifierOk() (*string, bool)`

GetGroupIdentifierOk returns a tuple with the GroupIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIdentifier

`func (o *VersionedLabel) SetGroupIdentifier(v string)`

SetGroupIdentifier sets GroupIdentifier field to given value.

### HasGroupIdentifier

`func (o *VersionedLabel) HasGroupIdentifier() bool`

HasGroupIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


