# RelationshipDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The relationship name. | [optional] 
**Description** | Pointer to **string** | The relationship description. | [optional] 
**AutoTerminate** | Pointer to **bool** | Whether or not flowfiles sent to this relationship should auto terminate. | [optional] 

## Methods

### NewRelationshipDTO

`func NewRelationshipDTO() *RelationshipDTO`

NewRelationshipDTO instantiates a new RelationshipDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipDTOWithDefaults

`func NewRelationshipDTOWithDefaults() *RelationshipDTO`

NewRelationshipDTOWithDefaults instantiates a new RelationshipDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RelationshipDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RelationshipDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RelationshipDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RelationshipDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RelationshipDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RelationshipDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RelationshipDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RelationshipDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAutoTerminate

`func (o *RelationshipDTO) GetAutoTerminate() bool`

GetAutoTerminate returns the AutoTerminate field if non-nil, zero value otherwise.

### GetAutoTerminateOk

`func (o *RelationshipDTO) GetAutoTerminateOk() (*bool, bool)`

GetAutoTerminateOk returns a tuple with the AutoTerminate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTerminate

`func (o *RelationshipDTO) SetAutoTerminate(v bool)`

SetAutoTerminate sets AutoTerminate field to given value.

### HasAutoTerminate

`func (o *RelationshipDTO) HasAutoTerminate() bool`

HasAutoTerminate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


