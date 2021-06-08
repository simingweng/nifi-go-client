# ProvenanceEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provenance** | Pointer to [**ProvenanceDTO**](ProvenanceDTO.md) |  | [optional] 

## Methods

### NewProvenanceEntity

`func NewProvenanceEntity() *ProvenanceEntity`

NewProvenanceEntity instantiates a new ProvenanceEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvenanceEntityWithDefaults

`func NewProvenanceEntityWithDefaults() *ProvenanceEntity`

NewProvenanceEntityWithDefaults instantiates a new ProvenanceEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvenance

`func (o *ProvenanceEntity) GetProvenance() ProvenanceDTO`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *ProvenanceEntity) GetProvenanceOk() (*ProvenanceDTO, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *ProvenanceEntity) SetProvenance(v ProvenanceDTO)`

SetProvenance sets Provenance field to given value.

### HasProvenance

`func (o *ProvenanceEntity) HasProvenance() bool`

HasProvenance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


