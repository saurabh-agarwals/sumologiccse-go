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

// checks if the BulkDeleteMatchListItemsRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkDeleteMatchListItemsRequestBody{}

// BulkDeleteMatchListItemsRequestBody struct for BulkDeleteMatchListItemsRequestBody
type BulkDeleteMatchListItemsRequestBody struct {
	Ids []string `json:"ids"`
}

type _BulkDeleteMatchListItemsRequestBody BulkDeleteMatchListItemsRequestBody

// NewBulkDeleteMatchListItemsRequestBody instantiates a new BulkDeleteMatchListItemsRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkDeleteMatchListItemsRequestBody(ids []string) *BulkDeleteMatchListItemsRequestBody {
	this := BulkDeleteMatchListItemsRequestBody{}
	this.Ids = ids
	return &this
}

// NewBulkDeleteMatchListItemsRequestBodyWithDefaults instantiates a new BulkDeleteMatchListItemsRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkDeleteMatchListItemsRequestBodyWithDefaults() *BulkDeleteMatchListItemsRequestBody {
	this := BulkDeleteMatchListItemsRequestBody{}
	return &this
}

// GetIds returns the Ids field value
func (o *BulkDeleteMatchListItemsRequestBody) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *BulkDeleteMatchListItemsRequestBody) GetIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *BulkDeleteMatchListItemsRequestBody) SetIds(v []string) {
	o.Ids = v
}

func (o BulkDeleteMatchListItemsRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkDeleteMatchListItemsRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ids"] = o.Ids
	return toSerialize, nil
}

func (o *BulkDeleteMatchListItemsRequestBody) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ids",
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

	varBulkDeleteMatchListItemsRequestBody := _BulkDeleteMatchListItemsRequestBody{}

	err = json.Unmarshal(bytes, &varBulkDeleteMatchListItemsRequestBody)

	if err != nil {
		return err
	}

	*o = BulkDeleteMatchListItemsRequestBody(varBulkDeleteMatchListItemsRequestBody)

	return err
}

type NullableBulkDeleteMatchListItemsRequestBody struct {
	value *BulkDeleteMatchListItemsRequestBody
	isSet bool
}

func (v NullableBulkDeleteMatchListItemsRequestBody) Get() *BulkDeleteMatchListItemsRequestBody {
	return v.value
}

func (v *NullableBulkDeleteMatchListItemsRequestBody) Set(val *BulkDeleteMatchListItemsRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkDeleteMatchListItemsRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkDeleteMatchListItemsRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkDeleteMatchListItemsRequestBody(val *BulkDeleteMatchListItemsRequestBody) *NullableBulkDeleteMatchListItemsRequestBody {
	return &NullableBulkDeleteMatchListItemsRequestBody{value: val, isSet: true}
}

func (v NullableBulkDeleteMatchListItemsRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkDeleteMatchListItemsRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

