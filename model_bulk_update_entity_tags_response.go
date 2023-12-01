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

// checks if the BulkUpdateEntityTagsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkUpdateEntityTagsResponse{}

// BulkUpdateEntityTagsResponse struct for BulkUpdateEntityTagsResponse
type BulkUpdateEntityTagsResponse struct {
	Data []BulkAddEntityTagsResponseDataInner `json:"data"`
}

type _BulkUpdateEntityTagsResponse BulkUpdateEntityTagsResponse

// NewBulkUpdateEntityTagsResponse instantiates a new BulkUpdateEntityTagsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkUpdateEntityTagsResponse(data []BulkAddEntityTagsResponseDataInner) *BulkUpdateEntityTagsResponse {
	this := BulkUpdateEntityTagsResponse{}
	this.Data = data
	return &this
}

// NewBulkUpdateEntityTagsResponseWithDefaults instantiates a new BulkUpdateEntityTagsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkUpdateEntityTagsResponseWithDefaults() *BulkUpdateEntityTagsResponse {
	this := BulkUpdateEntityTagsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BulkUpdateEntityTagsResponse) GetData() []BulkAddEntityTagsResponseDataInner {
	if o == nil {
		var ret []BulkAddEntityTagsResponseDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BulkUpdateEntityTagsResponse) GetDataOk() ([]BulkAddEntityTagsResponseDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BulkUpdateEntityTagsResponse) SetData(v []BulkAddEntityTagsResponseDataInner) {
	o.Data = v
}

func (o BulkUpdateEntityTagsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkUpdateEntityTagsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *BulkUpdateEntityTagsResponse) UnmarshalJSON(bytes []byte) (err error) {
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

	varBulkUpdateEntityTagsResponse := _BulkUpdateEntityTagsResponse{}

	err = json.Unmarshal(bytes, &varBulkUpdateEntityTagsResponse)

	if err != nil {
		return err
	}

	*o = BulkUpdateEntityTagsResponse(varBulkUpdateEntityTagsResponse)

	return err
}

type NullableBulkUpdateEntityTagsResponse struct {
	value *BulkUpdateEntityTagsResponse
	isSet bool
}

func (v NullableBulkUpdateEntityTagsResponse) Get() *BulkUpdateEntityTagsResponse {
	return v.value
}

func (v *NullableBulkUpdateEntityTagsResponse) Set(val *BulkUpdateEntityTagsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkUpdateEntityTagsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkUpdateEntityTagsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkUpdateEntityTagsResponse(val *BulkUpdateEntityTagsResponse) *NullableBulkUpdateEntityTagsResponse {
	return &NullableBulkUpdateEntityTagsResponse{value: val, isSet: true}
}

func (v NullableBulkUpdateEntityTagsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkUpdateEntityTagsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

