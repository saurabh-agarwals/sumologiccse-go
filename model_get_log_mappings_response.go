/*
Sumo Logic CSE API

 https://help.sumologic.com/APIs 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologiccse

import (
	"encoding/json"
	"fmt"
)

// checks if the GetLogMappingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLogMappingsResponse{}

// GetLogMappingsResponse struct for GetLogMappingsResponse
type GetLogMappingsResponse struct {
	Data GetLogMappingsResponseData `json:"data"`
}

type _GetLogMappingsResponse GetLogMappingsResponse

// NewGetLogMappingsResponse instantiates a new GetLogMappingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLogMappingsResponse(data GetLogMappingsResponseData) *GetLogMappingsResponse {
	this := GetLogMappingsResponse{}
	this.Data = data
	return &this
}

// NewGetLogMappingsResponseWithDefaults instantiates a new GetLogMappingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLogMappingsResponseWithDefaults() *GetLogMappingsResponse {
	this := GetLogMappingsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetLogMappingsResponse) GetData() GetLogMappingsResponseData {
	if o == nil {
		var ret GetLogMappingsResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetLogMappingsResponse) GetDataOk() (*GetLogMappingsResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetLogMappingsResponse) SetData(v GetLogMappingsResponseData) {
	o.Data = v
}

func (o GetLogMappingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLogMappingsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetLogMappingsResponse) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetLogMappingsResponse := _GetLogMappingsResponse{}

	err = json.Unmarshal(bytes, &varGetLogMappingsResponse)

	if err != nil {
		return err
	}

	*o = GetLogMappingsResponse(varGetLogMappingsResponse)

	return err
}

type NullableGetLogMappingsResponse struct {
	value *GetLogMappingsResponse
	isSet bool
}

func (v NullableGetLogMappingsResponse) Get() *GetLogMappingsResponse {
	return v.value
}

func (v *NullableGetLogMappingsResponse) Set(val *GetLogMappingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLogMappingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLogMappingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLogMappingsResponse(val *GetLogMappingsResponse) *NullableGetLogMappingsResponse {
	return &NullableGetLogMappingsResponse{value: val, isSet: true}
}

func (v NullableGetLogMappingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLogMappingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

