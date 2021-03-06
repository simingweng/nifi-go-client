/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.13.2
 * Contact: dev@nifi.apache.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nifi

import (
	"encoding/json"
)

// VersionedLabel struct for VersionedLabel
type VersionedLabel struct {
	// The component's unique identifier
	Identifier *string `json:"identifier,omitempty"`
	// The component's name
	Name *string `json:"name,omitempty"`
	// The user-supplied comments for the component
	Comments *string   `json:"comments,omitempty"`
	Position *Position `json:"position,omitempty"`
	// The text that appears in the label.
	Label *string `json:"label,omitempty"`
	// The width of the label in pixels when at a 1:1 scale.
	Width *float64 `json:"width,omitempty"`
	// The height of the label in pixels when at a 1:1 scale.
	Height *float64 `json:"height,omitempty"`
	// The styles for this label (font-size : 12px, background-color : #eee, etc).
	Style         *map[string]string `json:"style,omitempty"`
	ComponentType *string            `json:"componentType,omitempty"`
	// The ID of the Process Group that this component belongs to
	GroupIdentifier *string `json:"groupIdentifier,omitempty"`
}

// NewVersionedLabel instantiates a new VersionedLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionedLabel() *VersionedLabel {
	this := VersionedLabel{}
	return &this
}

// NewVersionedLabelWithDefaults instantiates a new VersionedLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionedLabelWithDefaults() *VersionedLabel {
	this := VersionedLabel{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *VersionedLabel) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedLabel) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *VersionedLabel) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *VersionedLabel) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VersionedLabel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedLabel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VersionedLabel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VersionedLabel) SetName(v string) {
	o.Name = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *VersionedLabel) GetComments() string {
	if o == nil || o.Comments == nil {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedLabel) GetCommentsOk() (*string, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *VersionedLabel) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *VersionedLabel) SetComments(v string) {
	o.Comments = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *VersionedLabel) GetPosition() Position {
	if o == nil || o.Position == nil {
		var ret Position
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedLabel) GetPositionOk() (*Position, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *VersionedLabel) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given Position and assigns it to the Position field.
func (o *VersionedLabel) SetPosition(v Position) {
	o.Position = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *VersionedLabel) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedLabel) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *VersionedLabel) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *VersionedLabel) SetLabel(v string) {
	o.Label = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *VersionedLabel) GetWidth() float64 {
	if o == nil || o.Width == nil {
		var ret float64
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedLabel) GetWidthOk() (*float64, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *VersionedLabel) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float64 and assigns it to the Width field.
func (o *VersionedLabel) SetWidth(v float64) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *VersionedLabel) GetHeight() float64 {
	if o == nil || o.Height == nil {
		var ret float64
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedLabel) GetHeightOk() (*float64, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *VersionedLabel) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float64 and assigns it to the Height field.
func (o *VersionedLabel) SetHeight(v float64) {
	o.Height = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *VersionedLabel) GetStyle() map[string]string {
	if o == nil || o.Style == nil {
		var ret map[string]string
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedLabel) GetStyleOk() (*map[string]string, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *VersionedLabel) HasStyle() bool {
	if o != nil && o.Style != nil {
		return true
	}

	return false
}

// SetStyle gets a reference to the given map[string]string and assigns it to the Style field.
func (o *VersionedLabel) SetStyle(v map[string]string) {
	o.Style = &v
}

// GetComponentType returns the ComponentType field value if set, zero value otherwise.
func (o *VersionedLabel) GetComponentType() string {
	if o == nil || o.ComponentType == nil {
		var ret string
		return ret
	}
	return *o.ComponentType
}

// GetComponentTypeOk returns a tuple with the ComponentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedLabel) GetComponentTypeOk() (*string, bool) {
	if o == nil || o.ComponentType == nil {
		return nil, false
	}
	return o.ComponentType, true
}

// HasComponentType returns a boolean if a field has been set.
func (o *VersionedLabel) HasComponentType() bool {
	if o != nil && o.ComponentType != nil {
		return true
	}

	return false
}

// SetComponentType gets a reference to the given string and assigns it to the ComponentType field.
func (o *VersionedLabel) SetComponentType(v string) {
	o.ComponentType = &v
}

// GetGroupIdentifier returns the GroupIdentifier field value if set, zero value otherwise.
func (o *VersionedLabel) GetGroupIdentifier() string {
	if o == nil || o.GroupIdentifier == nil {
		var ret string
		return ret
	}
	return *o.GroupIdentifier
}

// GetGroupIdentifierOk returns a tuple with the GroupIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedLabel) GetGroupIdentifierOk() (*string, bool) {
	if o == nil || o.GroupIdentifier == nil {
		return nil, false
	}
	return o.GroupIdentifier, true
}

// HasGroupIdentifier returns a boolean if a field has been set.
func (o *VersionedLabel) HasGroupIdentifier() bool {
	if o != nil && o.GroupIdentifier != nil {
		return true
	}

	return false
}

// SetGroupIdentifier gets a reference to the given string and assigns it to the GroupIdentifier field.
func (o *VersionedLabel) SetGroupIdentifier(v string) {
	o.GroupIdentifier = &v
}

func (o VersionedLabel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Width != nil {
		toSerialize["width"] = o.Width
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	if o.ComponentType != nil {
		toSerialize["componentType"] = o.ComponentType
	}
	if o.GroupIdentifier != nil {
		toSerialize["groupIdentifier"] = o.GroupIdentifier
	}
	return json.Marshal(toSerialize)
}

type NullableVersionedLabel struct {
	value *VersionedLabel
	isSet bool
}

func (v NullableVersionedLabel) Get() *VersionedLabel {
	return v.value
}

func (v *NullableVersionedLabel) Set(val *VersionedLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionedLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionedLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionedLabel(val *VersionedLabel) *NullableVersionedLabel {
	return &NullableVersionedLabel{value: val, isSet: true}
}

func (v NullableVersionedLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionedLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
