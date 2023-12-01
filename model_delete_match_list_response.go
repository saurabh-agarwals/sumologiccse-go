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

// checks if the DeleteMatchListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteMatchListResponse{}

// DeleteMatchListResponse struct for DeleteMatchListResponse
type DeleteMatchListResponse struct {
	Data DeleteContextActionResponseData `json:"data"`
}

type _DeleteMatchListResponse DeleteMatchListResponse

// NewDeleteMatchListResponse instantiates a new DeleteMatchListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteMatchListResponse(data DeleteContextActionResponseData) *DeleteMatchListResponse {
	this := DeleteMatchListResponse{}
	this.Data = data
	return &this
}

// NewDeleteMatchListResponseWithDefaults instantiates a new DeleteMatchListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteMatchListResponseWithDefaults() *DeleteMatchListResponse {
	this := DeleteMatchListResponse{}
	return &this
}

// GetData returns the Data field value
func (o *DeleteMatchListResponse) GetData() DeleteContextActionResponseData {
	if o == nil {
		var ret DeleteContextActionResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeleteMatchListResponse) GetDataOk() (*DeleteContextActionResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DeleteMatchListResponse) SetData(v DeleteContextActionResponseData) {
	o.Data = v
}

func (o DeleteMatchListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteMatchListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *DeleteMatchListResponse) UnmarshalJSON(bytes []byte) (err error) {
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

	varDeleteMatchListResponse := _DeleteMatchListResponse{}

	err = json.Unmarshal(bytes, &varDeleteMatchListResponse)

	if err != nil {
		return err
	}

	*o = DeleteMatchListResponse(varDeleteMatchListResponse)

	return err
}

type NullableDeleteMatchListResponse struct {
	value *DeleteMatchListResponse
	isSet bool
}

func (v NullableDeleteMatchListResponse) Get() *DeleteMatchListResponse {
	return v.value
}

func (v *NullableDeleteMatchListResponse) Set(val *DeleteMatchListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteMatchListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteMatchListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteMatchListResponse(val *DeleteMatchListResponse) *NullableDeleteMatchListResponse {
	return &NullableDeleteMatchListResponse{value: val, isSet: true}
}

func (v NullableDeleteMatchListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteMatchListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

