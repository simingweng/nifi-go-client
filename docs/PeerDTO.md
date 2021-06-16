# PeerDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | The hostname of this peer. | [optional] 
**Port** | Pointer to **int32** | The port number of this peer. | [optional] 
**Secure** | Pointer to **bool** | Returns if this peer connection is secure. | [optional] 
**FlowFileCount** | Pointer to **int32** | The number of flowFiles this peer holds. | [optional] 

## Methods

### NewPeerDTO

`func NewPeerDTO() *PeerDTO`

NewPeerDTO instantiates a new PeerDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeerDTOWithDefaults

`func NewPeerDTOWithDefaults() *PeerDTO`

NewPeerDTOWithDefaults instantiates a new PeerDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *PeerDTO) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *PeerDTO) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *PeerDTO) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *PeerDTO) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *PeerDTO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PeerDTO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PeerDTO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PeerDTO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecure

`func (o *PeerDTO) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *PeerDTO) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *PeerDTO) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *PeerDTO) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetFlowFileCount

`func (o *PeerDTO) GetFlowFileCount() int32`

GetFlowFileCount returns the FlowFileCount field if non-nil, zero value otherwise.

### GetFlowFileCountOk

`func (o *PeerDTO) GetFlowFileCountOk() (*int32, bool)`

GetFlowFileCountOk returns a tuple with the FlowFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowFileCount

`func (o *PeerDTO) SetFlowFileCount(v int32)`

SetFlowFileCount sets FlowFileCount field to given value.

### HasFlowFileCount

`func (o *PeerDTO) HasFlowFileCount() bool`

HasFlowFileCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


