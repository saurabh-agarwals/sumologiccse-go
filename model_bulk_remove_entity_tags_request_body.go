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

// checks if the BulkRemoveEntityTagsRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkRemoveEntityTagsRequestBody{}

// BulkRemoveEntityTagsRequestBody struct for BulkRemoveEntityTagsRequestBody
type BulkRemoveEntityTagsRequestBody struct {
	EntityIds []string `json:"entityIds"`
	Tags []string `json:"tags"`
}

type _BulkRemoveEntityTagsRequestBody BulkRemoveEntityTagsRequestBody

// NewBulkRemoveEntityTagsRequestBody instantiates a new BulkRemoveEntityTagsRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkRemoveEntityTagsRequestBody(entityIds []string, tags []string) *BulkRemoveEntityTagsRequestBody {
	this := BulkRemoveEntityTagsRequestBody{}
	this.EntityIds = entityIds
	this.Tags = tags
	return &this
}

// NewBulkRemoveEntityTagsRequestBodyWithDefaults instantiates a new BulkRemoveEntityTagsRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkRemoveEntityTagsRequestBodyWithDefaults() *BulkRemoveEntityTagsRequestBody {
	this := BulkRemoveEntityTagsRequestBody{}
	return &this
}

// GetEntityIds returns the EntityIds field value
func (o *BulkRemoveEntityTagsRequestBody) GetEntityIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EntityIds
}

// GetEntityIdsOk returns a tuple with the EntityIds field value
// and a boolean to check if the value has been set.
func (o *BulkRemoveEntityTagsRequestBody) GetEntityIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntityIds, true
}

// SetEntityIds sets field value
func (o *BulkRemoveEntityTagsRequestBody) SetEntityIds(v []string) {
	o.EntityIds = v
}

// GetTags returns the Tags field value
func (o *BulkRemoveEntityTagsRequestBody) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *BulkRemoveEntityTagsRequestBody) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *BulkRemoveEntityTagsRequestBody) SetTags(v []string) {
	o.Tags = v
}

func (o BulkRemoveEntityTagsRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkRemoveEntityTagsRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entityIds"] = o.EntityIds
	toSerialize["tags"] = o.Tags
	return toSerialize, nil
}

func (o *BulkRemoveEntityTagsRequestBody) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entityIds",
		"tags",
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

	varBulkRemoveEntityTagsRequestBody := _BulkRemoveEntityTagsRequestBody{}

	err = json.Unmarshal(bytes, &varBulkRemoveEntityTagsRequestBody)

	if err != nil {
		return err
	}

	*o = BulkRemoveEntityTagsRequestBody(varBulkRemoveEntityTagsRequestBody)

	return err
}

type NullableBulkRemoveEntityTagsRequestBody struct {
	value *BulkRemoveEntityTagsRequestBody
	isSet bool
}

func (v NullableBulkRemoveEntityTagsRequestBody) Get() *BulkRemoveEntityTagsRequestBody {
	return v.value
}

func (v *NullableBulkRemoveEntityTagsRequestBody) Set(val *BulkRemoveEntityTagsRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkRemoveEntityTagsRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkRemoveEntityTagsRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkRemoveEntityTagsRequestBody(val *BulkRemoveEntityTagsRequestBody) *NullableBulkRemoveEntityTagsRequestBody {
	return &NullableBulkRemoveEntityTagsRequestBody{value: val, isSet: true}
}

func (v NullableBulkRemoveEntityTagsRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkRemoveEntityTagsRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

