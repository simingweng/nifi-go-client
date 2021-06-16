# ProvenanceOptionsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchableFields** | Pointer to [**[]ProvenanceSearchableFieldDTO**](ProvenanceSearchableFieldDTO.md) | The available searchable field for the NiFi. | [optional] 

## Methods

### NewProvenanceOptionsDTO

`func NewProvenanceOptionsDTO() *ProvenanceOptionsDTO`

NewProvenanceOptionsDTO instantiates a new ProvenanceOptionsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvenanceOptionsDTOWithDefaults

`func NewProvenanceOptionsDTOWithDefaults() *ProvenanceOptionsDTO`

NewProvenanceOptionsDTOWithDefaults instantiates a new ProvenanceOptionsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchableFields

`func (o *ProvenanceOptionsDTO) GetSearchableFields() []ProvenanceSearchableFieldDTO`

GetSearchableFields returns the SearchableFields field if non-nil, zero value otherwise.

### GetSearchableFieldsOk

`func (o *ProvenanceOptionsDTO) GetSearchableFieldsOk() (*[]ProvenanceSearchableFieldDTO, bool)`

GetSearchableFieldsOk returns a tuple with the SearchableFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchableFields

`func (o *ProvenanceOptionsDTO) SetSearchableFields(v []ProvenanceSearchableFieldDTO)`

SetSearchableFields sets SearchableFields field to given value.

### HasSearchableFields

`func (o *ProvenanceOptionsDTO) HasSearchableFields() bool`

HasSearchableFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


