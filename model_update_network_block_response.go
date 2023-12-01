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

// checks if the UpdateNetworkBlockResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkBlockResponse{}

// UpdateNetworkBlockResponse struct for UpdateNetworkBlockResponse
type UpdateNetworkBlockResponse struct {
	Data NetworkBlock `json:"data"`
}

type _UpdateNetworkBlockResponse UpdateNetworkBlockResponse

// NewUpdateNetworkBlockResponse instantiates a new UpdateNetworkBlockResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkBlockResponse(data NetworkBlock) *UpdateNetworkBlockResponse {
	this := UpdateNetworkBlockResponse{}
	this.Data = data
	return &this
}

// NewUpdateNetworkBlockResponseWithDefaults instantiates a new UpdateNetworkBlockResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkBlockResponseWithDefaults() *UpdateNetworkBlockResponse {
	this := UpdateNetworkBlockResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateNetworkBlockResponse) GetData() NetworkBlock {
	if o == nil {
		var ret NetworkBlock
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkBlockResponse) GetDataOk() (*NetworkBlock, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdateNetworkBlockResponse) SetData(v NetworkBlock) {
	o.Data = v
}

func (o UpdateNetworkBlockResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkBlockResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *UpdateNetworkBlockResponse) UnmarshalJSON(bytes []byte) (err error) {
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

	varUpdateNetworkBlockResponse := _UpdateNetworkBlockResponse{}

	err = json.Unmarshal(bytes, &varUpdateNetworkBlockResponse)

	if err != nil {
		return err
	}

	*o = UpdateNetworkBlockResponse(varUpdateNetworkBlockResponse)

	return err
}

type NullableUpdateNetworkBlockResponse struct {
	value *UpdateNetworkBlockResponse
	isSet bool
}

func (v NullableUpdateNetworkBlockResponse) Get() *UpdateNetworkBlockResponse {
	return v.value
}

func (v *NullableUpdateNetworkBlockResponse) Set(val *UpdateNetworkBlockResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkBlockResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkBlockResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkBlockResponse(val *UpdateNetworkBlockResponse) *NullableUpdateNetworkBlockResponse {
	return &NullableUpdateNetworkBlockResponse{value: val, isSet: true}
}

func (v NullableUpdateNetworkBlockResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkBlockResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

