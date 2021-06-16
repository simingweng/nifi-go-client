# VersionInfoDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NiFiVersion** | Pointer to **string** | The version of this NiFi. | [optional] 
**JavaVendor** | Pointer to **string** | Java JVM vendor | [optional] 
**JavaVersion** | Pointer to **string** | Java version | [optional] 
**OsName** | Pointer to **string** | Host operating system name | [optional] 
**OsVersion** | Pointer to **string** | Host operating system version | [optional] 
**OsArchitecture** | Pointer to **string** | Host operating system architecture | [optional] 
**BuildTag** | Pointer to **string** | Build tag | [optional] 
**BuildRevision** | Pointer to **string** | Build revision or commit hash | [optional] 
**BuildBranch** | Pointer to **string** | Build branch | [optional] 
**BuildTimestamp** | Pointer to **time.Time** | Build timestamp | [optional] 

## Methods

### NewVersionInfoDTO

`func NewVersionInfoDTO() *VersionInfoDTO`

NewVersionInfoDTO instantiates a new VersionInfoDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionInfoDTOWithDefaults

`func NewVersionInfoDTOWithDefaults() *VersionInfoDTO`

NewVersionInfoDTOWithDefaults instantiates a new VersionInfoDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNiFiVersion

`func (o *VersionInfoDTO) GetNiFiVersion() string`

GetNiFiVersion returns the NiFiVersion field if non-nil, zero value otherwise.

### GetNiFiVersionOk

`func (o *VersionInfoDTO) GetNiFiVersionOk() (*string, bool)`

GetNiFiVersionOk returns a tuple with the NiFiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiFiVersion

`func (o *VersionInfoDTO) SetNiFiVersion(v string)`

SetNiFiVersion sets NiFiVersion field to given value.

### HasNiFiVersion

`func (o *VersionInfoDTO) HasNiFiVersion() bool`

HasNiFiVersion returns a boolean if a field has been set.

### GetJavaVendor

`func (o *VersionInfoDTO) GetJavaVendor() string`

GetJavaVendor returns the JavaVendor field if non-nil, zero value otherwise.

### GetJavaVendorOk

`func (o *VersionInfoDTO) GetJavaVendorOk() (*string, bool)`

GetJavaVendorOk returns a tuple with the JavaVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaVendor

`func (o *VersionInfoDTO) SetJavaVendor(v string)`

SetJavaVendor sets JavaVendor field to given value.

### HasJavaVendor

`func (o *VersionInfoDTO) HasJavaVendor() bool`

HasJavaVendor returns a boolean if a field has been set.

### GetJavaVersion

`func (o *VersionInfoDTO) GetJavaVersion() string`

GetJavaVersion returns the JavaVersion field if non-nil, zero value otherwise.

### GetJavaVersionOk

`func (o *VersionInfoDTO) GetJavaVersionOk() (*string, bool)`

GetJavaVersionOk returns a tuple with the JavaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaVersion

`func (o *VersionInfoDTO) SetJavaVersion(v string)`

SetJavaVersion sets JavaVersion field to given value.

### HasJavaVersion

`func (o *VersionInfoDTO) HasJavaVersion() bool`

HasJavaVersion returns a boolean if a field has been set.

### GetOsName

`func (o *VersionInfoDTO) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *VersionInfoDTO) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *VersionInfoDTO) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *VersionInfoDTO) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### GetOsVersion

`func (o *VersionInfoDTO) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *VersionInfoDTO) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *VersionInfoDTO) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *VersionInfoDTO) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetOsArchitecture

`func (o *VersionInfoDTO) GetOsArchitecture() string`

GetOsArchitecture returns the OsArchitecture field if non-nil, zero value otherwise.

### GetOsArchitectureOk

`func (o *VersionInfoDTO) GetOsArchitectureOk() (*string, bool)`

GetOsArchitectureOk returns a tuple with the OsArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsArchitecture

`func (o *VersionInfoDTO) SetOsArchitecture(v string)`

SetOsArchitecture sets OsArchitecture field to given value.

### HasOsArchitecture

`func (o *VersionInfoDTO) HasOsArchitecture() bool`

HasOsArchitecture returns a boolean if a field has been set.

### GetBuildTag

`func (o *VersionInfoDTO) GetBuildTag() string`

GetBuildTag returns the BuildTag field if non-nil, zero value otherwise.

### GetBuildTagOk

`func (o *VersionInfoDTO) GetBuildTagOk() (*string, bool)`

GetBuildTagOk returns a tuple with the BuildTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildTag

`func (o *VersionInfoDTO) SetBuildTag(v string)`

SetBuildTag sets BuildTag field to given value.

### HasBuildTag

`func (o *VersionInfoDTO) HasBuildTag() bool`

HasBuildTag returns a boolean if a field has been set.

### GetBuildRevision

`func (o *VersionInfoDTO) GetBuildRevision() string`

GetBuildRevision returns the BuildRevision field if non-nil, zero value otherwise.

### GetBuildRevisionOk

`func (o *VersionInfoDTO) GetBuildRevisionOk() (*string, bool)`

GetBuildRevisionOk returns a tuple with the BuildRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildRevision

`func (o *VersionInfoDTO) SetBuildRevision(v string)`

SetBuildRevision sets BuildRevision field to given value.

### HasBuildRevision

`func (o *VersionInfoDTO) HasBuildRevision() bool`

HasBuildRevision returns a boolean if a field has been set.

### GetBuildBranch

`func (o *VersionInfoDTO) GetBuildBranch() string`

GetBuildBranch returns the BuildBranch field if non-nil, zero value otherwise.

### GetBuildBranchOk

`func (o *VersionInfoDTO) GetBuildBranchOk() (*string, bool)`

GetBuildBranchOk returns a tuple with the BuildBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildBranch

`func (o *VersionInfoDTO) SetBuildBranch(v string)`

SetBuildBranch sets BuildBranch field to given value.

### HasBuildBranch

`func (o *VersionInfoDTO) HasBuildBranch() bool`

HasBuildBranch returns a boolean if a field has been set.

### GetBuildTimestamp

`func (o *VersionInfoDTO) GetBuildTimestamp() time.Time`

GetBuildTimestamp returns the BuildTimestamp field if non-nil, zero value otherwise.

### GetBuildTimestampOk

`func (o *VersionInfoDTO) GetBuildTimestampOk() (*time.Time, bool)`

GetBuildTimestampOk returns a tuple with the BuildTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildTimestamp

`func (o *VersionInfoDTO) SetBuildTimestamp(v time.Time)`

SetBuildTimestamp sets BuildTimestamp field to given value.

### HasBuildTimestamp

`func (o *VersionInfoDTO) HasBuildTimestamp() bool`

HasBuildTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


