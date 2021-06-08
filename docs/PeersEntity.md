# PeersEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Peers** | Pointer to [**[]PeerDTO**](PeerDTO.md) |  | [optional] 

## Methods

### NewPeersEntity

`func NewPeersEntity() *PeersEntity`

NewPeersEntity instantiates a new PeersEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeersEntityWithDefaults

`func NewPeersEntityWithDefaults() *PeersEntity`

NewPeersEntityWithDefaults instantiates a new PeersEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeers

`func (o *PeersEntity) GetPeers() []PeerDTO`

GetPeers returns the Peers field if non-nil, zero value otherwise.

### GetPeersOk

`func (o *PeersEntity) GetPeersOk() (*[]PeerDTO, bool)`

GetPeersOk returns a tuple with the Peers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeers

`func (o *PeersEntity) SetPeers(v []PeerDTO)`

SetPeers sets Peers field to given value.

### HasPeers

`func (o *PeersEntity) HasPeers() bool`

HasPeers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


