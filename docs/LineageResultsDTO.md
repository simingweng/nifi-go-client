# LineageResultsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to **[]string** | Any errors that occurred while generating the lineage. | [optional] 
**Nodes** | Pointer to [**[]ProvenanceNodeDTO**](ProvenanceNodeDTO.md) | The nodes in the lineage. | [optional] 
**Links** | Pointer to [**[]ProvenanceLinkDTO**](ProvenanceLinkDTO.md) | The links between the nodes in the lineage. | [optional] 

## Methods

### NewLineageResultsDTO

`func NewLineageResultsDTO() *LineageResultsDTO`

NewLineageResultsDTO instantiates a new LineageResultsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineageResultsDTOWithDefaults

`func NewLineageResultsDTOWithDefaults() *LineageResultsDTO`

NewLineageResultsDTOWithDefaults instantiates a new LineageResultsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *LineageResultsDTO) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *LineageResultsDTO) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *LineageResultsDTO) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *LineageResultsDTO) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetNodes

`func (o *LineageResultsDTO) GetNodes() []ProvenanceNodeDTO`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *LineageResultsDTO) GetNodesOk() (*[]ProvenanceNodeDTO, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *LineageResultsDTO) SetNodes(v []ProvenanceNodeDTO)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *LineageResultsDTO) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetLinks

`func (o *LineageResultsDTO) GetLinks() []ProvenanceLinkDTO`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LineageResultsDTO) GetLinksOk() (*[]ProvenanceLinkDTO, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LineageResultsDTO) SetLinks(v []ProvenanceLinkDTO)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LineageResultsDTO) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


