/*
Delphix API Gateway

Delphix API Gateway API

API version: 1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// StopVDBResponse struct for StopVDBResponse
type StopVDBResponse struct {
	// The initiated job id.
	JobId *string `json:"job_id,omitempty"`
}

// NewStopVDBResponse instantiates a new StopVDBResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStopVDBResponse() *StopVDBResponse {
	this := StopVDBResponse{}
	return &this
}

// NewStopVDBResponseWithDefaults instantiates a new StopVDBResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStopVDBResponseWithDefaults() *StopVDBResponse {
	this := StopVDBResponse{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *StopVDBResponse) GetJobId() string {
	if o == nil || o.JobId == nil {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopVDBResponse) GetJobIdOk() (*string, bool) {
	if o == nil || o.JobId == nil {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *StopVDBResponse) HasJobId() bool {
	if o != nil && o.JobId != nil {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *StopVDBResponse) SetJobId(v string) {
	o.JobId = &v
}

func (o StopVDBResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JobId != nil {
		toSerialize["job_id"] = o.JobId
	}
	return json.Marshal(toSerialize)
}

type NullableStopVDBResponse struct {
	value *StopVDBResponse
	isSet bool
}

func (v NullableStopVDBResponse) Get() *StopVDBResponse {
	return v.value
}

func (v *NullableStopVDBResponse) Set(val *StopVDBResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStopVDBResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStopVDBResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStopVDBResponse(val *StopVDBResponse) *NullableStopVDBResponse {
	return &NullableStopVDBResponse{value: val, isSet: true}
}

func (v NullableStopVDBResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStopVDBResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


