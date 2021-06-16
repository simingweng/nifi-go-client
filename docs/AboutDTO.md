# AboutDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The title to be used on the page and in the about dialog. | [optional] 
**Version** | Pointer to **string** | The version of this NiFi. | [optional] 
**Uri** | Pointer to **string** | The URI for the NiFi. | [optional] 
**ContentViewerUrl** | Pointer to **string** | The URL for the content viewer if configured. | [optional] 
**Timezone** | Pointer to **string** | The timezone of the NiFi instance. | [optional] [readonly] 
**BuildTag** | Pointer to **string** | Build tag | [optional] 
**BuildRevision** | Pointer to **string** | Build revision or commit hash | [optional] 
**BuildBranch** | Pointer to **string** | Build branch | [optional] 
**BuildTimestamp** | Pointer to **string** | Build timestamp | [optional] 

## Methods

### NewAboutDTO

`func NewAboutDTO() *AboutDTO`

NewAboutDTO instantiates a new AboutDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAboutDTOWithDefaults

`func NewAboutDTOWithDefaults() *AboutDTO`

NewAboutDTOWithDefaults instantiates a new AboutDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AboutDTO) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AboutDTO) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AboutDTO) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AboutDTO) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVersion

`func (o *AboutDTO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AboutDTO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AboutDTO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AboutDTO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUri

`func (o *AboutDTO) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *AboutDTO) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *AboutDTO) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *AboutDTO) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetContentViewerUrl

`func (o *AboutDTO) GetContentViewerUrl() string`

GetContentViewerUrl returns the ContentViewerUrl field if non-nil, zero value otherwise.

### GetContentViewerUrlOk

`func (o *AboutDTO) GetContentViewerUrlOk() (*string, bool)`

GetContentViewerUrlOk returns a tuple with the ContentViewerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentViewerUrl

`func (o *AboutDTO) SetContentViewerUrl(v string)`

SetContentViewerUrl sets ContentViewerUrl field to given value.

### HasContentViewerUrl

`func (o *AboutDTO) HasContentViewerUrl() bool`

HasContentViewerUrl returns a boolean if a field has been set.

### GetTimezone

`func (o *AboutDTO) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AboutDTO) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AboutDTO) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *AboutDTO) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetBuildTag

`func (o *AboutDTO) GetBuildTag() string`

GetBuildTag returns the BuildTag field if non-nil, zero value otherwise.

### GetBuildTagOk

`func (o *AboutDTO) GetBuildTagOk() (*string, bool)`

GetBuildTagOk returns a tuple with the BuildTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildTag

`func (o *AboutDTO) SetBuildTag(v string)`

SetBuildTag sets BuildTag field to given value.

### HasBuildTag

`func (o *AboutDTO) HasBuildTag() bool`

HasBuildTag returns a boolean if a field has been set.

### GetBuildRevision

`func (o *AboutDTO) GetBuildRevision() string`

GetBuildRevision returns the BuildRevision field if non-nil, zero value otherwise.

### GetBuildRevisionOk

`func (o *AboutDTO) GetBuildRevisionOk() (*string, bool)`

GetBuildRevisionOk returns a tuple with the BuildRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildRevision

`func (o *AboutDTO) SetBuildRevision(v string)`

SetBuildRevision sets BuildRevision field to given value.

### HasBuildRevision

`func (o *AboutDTO) HasBuildRevision() bool`

HasBuildRevision returns a boolean if a field has been set.

### GetBuildBranch

`func (o *AboutDTO) GetBuildBranch() string`

GetBuildBranch returns the BuildBranch field if non-nil, zero value otherwise.

### GetBuildBranchOk

`func (o *AboutDTO) GetBuildBranchOk() (*string, bool)`

GetBuildBranchOk returns a tuple with the BuildBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildBranch

`func (o *AboutDTO) SetBuildBranch(v string)`

SetBuildBranch sets BuildBranch field to given value.

### HasBuildBranch

`func (o *AboutDTO) HasBuildBranch() bool`

HasBuildBranch returns a boolean if a field has been set.

### GetBuildTimestamp

`func (o *AboutDTO) GetBuildTimestamp() string`

GetBuildTimestamp returns the BuildTimestamp field if non-nil, zero value otherwise.

### GetBuildTimestampOk

`func (o *AboutDTO) GetBuildTimestampOk() (*string, bool)`

GetBuildTimestampOk returns a tuple with the BuildTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildTimestamp

`func (o *AboutDTO) SetBuildTimestamp(v string)`

SetBuildTimestamp sets BuildTimestamp field to given value.

### HasBuildTimestamp

`func (o *AboutDTO) HasBuildTimestamp() bool`

HasBuildTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


