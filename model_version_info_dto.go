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
	"time"
)

// VersionInfoDTO struct for VersionInfoDTO
type VersionInfoDTO struct {
	// The version of this NiFi.
	NiFiVersion *string `json:"niFiVersion,omitempty"`
	// Java JVM vendor
	JavaVendor *string `json:"javaVendor,omitempty"`
	// Java version
	JavaVersion *string `json:"javaVersion,omitempty"`
	// Host operating system name
	OsName *string `json:"osName,omitempty"`
	// Host operating system version
	OsVersion *string `json:"osVersion,omitempty"`
	// Host operating system architecture
	OsArchitecture *string `json:"osArchitecture,omitempty"`
	// Build tag
	BuildTag *string `json:"buildTag,omitempty"`
	// Build revision or commit hash
	BuildRevision *string `json:"buildRevision,omitempty"`
	// Build branch
	BuildBranch *string `json:"buildBranch,omitempty"`
	// Build timestamp
	BuildTimestamp *time.Time `json:"buildTimestamp,omitempty"`
}

// NewVersionInfoDTO instantiates a new VersionInfoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionInfoDTO() *VersionInfoDTO {
	this := VersionInfoDTO{}
	return &this
}

// NewVersionInfoDTOWithDefaults instantiates a new VersionInfoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionInfoDTOWithDefaults() *VersionInfoDTO {
	this := VersionInfoDTO{}
	return &this
}

// GetNiFiVersion returns the NiFiVersion field value if set, zero value otherwise.
func (o *VersionInfoDTO) GetNiFiVersion() string {
	if o == nil || o.NiFiVersion == nil {
		var ret string
		return ret
	}
	return *o.NiFiVersion
}

// GetNiFiVersionOk returns a tuple with the NiFiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfoDTO) GetNiFiVersionOk() (*string, bool) {
	if o == nil || o.NiFiVersion == nil {
		return nil, false
	}
	return o.NiFiVersion, true
}

// HasNiFiVersion returns a boolean if a field has been set.
func (o *VersionInfoDTO) HasNiFiVersion() bool {
	if o != nil && o.NiFiVersion != nil {
		return true
	}

	return false
}

// SetNiFiVersion gets a reference to the given string and assigns it to the NiFiVersion field.
func (o *VersionInfoDTO) SetNiFiVersion(v string) {
	o.NiFiVersion = &v
}

// GetJavaVendor returns the JavaVendor field value if set, zero value otherwise.
func (o *VersionInfoDTO) GetJavaVendor() string {
	if o == nil || o.JavaVendor == nil {
		var ret string
		return ret
	}
	return *o.JavaVendor
}

// GetJavaVendorOk returns a tuple with the JavaVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfoDTO) GetJavaVendorOk() (*string, bool) {
	if o == nil || o.JavaVendor == nil {
		return nil, false
	}
	return o.JavaVendor, true
}

// HasJavaVendor returns a boolean if a field has been set.
func (o *VersionInfoDTO) HasJavaVendor() bool {
	if o != nil && o.JavaVendor != nil {
		return true
	}

	return false
}

// SetJavaVendor gets a reference to the given string and assigns it to the JavaVendor field.
func (o *VersionInfoDTO) SetJavaVendor(v string) {
	o.JavaVendor = &v
}

// GetJavaVersion returns the JavaVersion field value if set, zero value otherwise.
func (o *VersionInfoDTO) GetJavaVersion() string {
	if o == nil || o.JavaVersion == nil {
		var ret string
		return ret
	}
	return *o.JavaVersion
}

// GetJavaVersionOk returns a tuple with the JavaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfoDTO) GetJavaVersionOk() (*string, bool) {
	if o == nil || o.JavaVersion == nil {
		return nil, false
	}
	return o.JavaVersion, true
}

// HasJavaVersion returns a boolean if a field has been set.
func (o *VersionInfoDTO) HasJavaVersion() bool {
	if o != nil && o.JavaVersion != nil {
		return true
	}

	return false
}

// SetJavaVersion gets a reference to the given string and assigns it to the JavaVersion field.
func (o *VersionInfoDTO) SetJavaVersion(v string) {
	o.JavaVersion = &v
}

// GetOsName returns the OsName field value if set, zero value otherwise.
func (o *VersionInfoDTO) GetOsName() string {
	if o == nil || o.OsName == nil {
		var ret string
		return ret
	}
	return *o.OsName
}

// GetOsNameOk returns a tuple with the OsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfoDTO) GetOsNameOk() (*string, bool) {
	if o == nil || o.OsName == nil {
		return nil, false
	}
	return o.OsName, true
}

// HasOsName returns a boolean if a field has been set.
func (o *VersionInfoDTO) HasOsName() bool {
	if o != nil && o.OsName != nil {
		return true
	}

	return false
}

// SetOsName gets a reference to the given string and assigns it to the OsName field.
func (o *VersionInfoDTO) SetOsName(v string) {
	o.OsName = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *VersionInfoDTO) GetOsVersion() string {
	if o == nil || o.OsVersion == nil {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfoDTO) GetOsVersionOk() (*string, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *VersionInfoDTO) HasOsVersion() bool {
	if o != nil && o.OsVersion != nil {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *VersionInfoDTO) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetOsArchitecture returns the OsArchitecture field value if set, zero value otherwise.
func (o *VersionInfoDTO) GetOsArchitecture() string {
	if o == nil || o.OsArchitecture == nil {
		var ret string
		return ret
	}
	return *o.OsArchitecture
}

// GetOsArchitectureOk returns a tuple with the OsArchitecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfoDTO) GetOsArchitectureOk() (*string, bool) {
	if o == nil || o.OsArchitecture == nil {
		return nil, false
	}
	return o.OsArchitecture, true
}

// HasOsArchitecture returns a boolean if a field has been set.
func (o *VersionInfoDTO) HasOsArchitecture() bool {
	if o != nil && o.OsArchitecture != nil {
		return true
	}

	return false
}

// SetOsArchitecture gets a reference to the given string and assigns it to the OsArchitecture field.
func (o *VersionInfoDTO) SetOsArchitecture(v string) {
	o.OsArchitecture = &v
}

// GetBuildTag returns the BuildTag field value if set, zero value otherwise.
func (o *VersionInfoDTO) GetBuildTag() string {
	if o == nil || o.BuildTag == nil {
		var ret string
		return ret
	}
	return *o.BuildTag
}

// GetBuildTagOk returns a tuple with the BuildTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfoDTO) GetBuildTagOk() (*string, bool) {
	if o == nil || o.BuildTag == nil {
		return nil, false
	}
	return o.BuildTag, true
}

// HasBuildTag returns a boolean if a field has been set.
func (o *VersionInfoDTO) HasBuildTag() bool {
	if o != nil && o.BuildTag != nil {
		return true
	}

	return false
}

// SetBuildTag gets a reference to the given string and assigns it to the BuildTag field.
func (o *VersionInfoDTO) SetBuildTag(v string) {
	o.BuildTag = &v
}

// GetBuildRevision returns the BuildRevision field value if set, zero value otherwise.
func (o *VersionInfoDTO) GetBuildRevision() string {
	if o == nil || o.BuildRevision == nil {
		var ret string
		return ret
	}
	return *o.BuildRevision
}

// GetBuildRevisionOk returns a tuple with the BuildRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfoDTO) GetBuildRevisionOk() (*string, bool) {
	if o == nil || o.BuildRevision == nil {
		return nil, false
	}
	return o.BuildRevision, true
}

// HasBuildRevision returns a boolean if a field has been set.
func (o *VersionInfoDTO) HasBuildRevision() bool {
	if o != nil && o.BuildRevision != nil {
		return true
	}

	return false
}

// SetBuildRevision gets a reference to the given string and assigns it to the BuildRevision field.
func (o *VersionInfoDTO) SetBuildRevision(v string) {
	o.BuildRevision = &v
}

// GetBuildBranch returns the BuildBranch field value if set, zero value otherwise.
func (o *VersionInfoDTO) GetBuildBranch() string {
	if o == nil || o.BuildBranch == nil {
		var ret string
		return ret
	}
	return *o.BuildBranch
}

// GetBuildBranchOk returns a tuple with the BuildBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfoDTO) GetBuildBranchOk() (*string, bool) {
	if o == nil || o.BuildBranch == nil {
		return nil, false
	}
	return o.BuildBranch, true
}

// HasBuildBranch returns a boolean if a field has been set.
func (o *VersionInfoDTO) HasBuildBranch() bool {
	if o != nil && o.BuildBranch != nil {
		return true
	}

	return false
}

// SetBuildBranch gets a reference to the given string and assigns it to the BuildBranch field.
func (o *VersionInfoDTO) SetBuildBranch(v string) {
	o.BuildBranch = &v
}

// GetBuildTimestamp returns the BuildTimestamp field value if set, zero value otherwise.
func (o *VersionInfoDTO) GetBuildTimestamp() time.Time {
	if o == nil || o.BuildTimestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.BuildTimestamp
}

// GetBuildTimestampOk returns a tuple with the BuildTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfoDTO) GetBuildTimestampOk() (*time.Time, bool) {
	if o == nil || o.BuildTimestamp == nil {
		return nil, false
	}
	return o.BuildTimestamp, true
}

// HasBuildTimestamp returns a boolean if a field has been set.
func (o *VersionInfoDTO) HasBuildTimestamp() bool {
	if o != nil && o.BuildTimestamp != nil {
		return true
	}

	return false
}

// SetBuildTimestamp gets a reference to the given time.Time and assigns it to the BuildTimestamp field.
func (o *VersionInfoDTO) SetBuildTimestamp(v time.Time) {
	o.BuildTimestamp = &v
}

func (o VersionInfoDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NiFiVersion != nil {
		toSerialize["niFiVersion"] = o.NiFiVersion
	}
	if o.JavaVendor != nil {
		toSerialize["javaVendor"] = o.JavaVendor
	}
	if o.JavaVersion != nil {
		toSerialize["javaVersion"] = o.JavaVersion
	}
	if o.OsName != nil {
		toSerialize["osName"] = o.OsName
	}
	if o.OsVersion != nil {
		toSerialize["osVersion"] = o.OsVersion
	}
	if o.OsArchitecture != nil {
		toSerialize["osArchitecture"] = o.OsArchitecture
	}
	if o.BuildTag != nil {
		toSerialize["buildTag"] = o.BuildTag
	}
	if o.BuildRevision != nil {
		toSerialize["buildRevision"] = o.BuildRevision
	}
	if o.BuildBranch != nil {
		toSerialize["buildBranch"] = o.BuildBranch
	}
	if o.BuildTimestamp != nil {
		toSerialize["buildTimestamp"] = o.BuildTimestamp
	}
	return json.Marshal(toSerialize)
}

type NullableVersionInfoDTO struct {
	value *VersionInfoDTO
	isSet bool
}

func (v NullableVersionInfoDTO) Get() *VersionInfoDTO {
	return v.value
}

func (v *NullableVersionInfoDTO) Set(val *VersionInfoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionInfoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionInfoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionInfoDTO(val *VersionInfoDTO) *NullableVersionInfoDTO {
	return &NullableVersionInfoDTO{value: val, isSet: true}
}

func (v NullableVersionInfoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionInfoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
