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

// GetEntitiesResponseDataObjectsInnerInventoryInner - struct for GetEntitiesResponseDataObjectsInnerInventoryInner
type GetEntitiesResponseDataObjectsInnerInventoryInner struct {
	GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf
	GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf1 *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf1
}

// GetEntitiesResponseDataObjectsInnerInventoryInnerOneOfAsGetEntitiesResponseDataObjectsInnerInventoryInner is a convenience function that returns GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf wrapped in GetEntitiesResponseDataObjectsInnerInventoryInner
func GetEntitiesResponseDataObjectsInnerInventoryInnerOneOfAsGetEntitiesResponseDataObjectsInnerInventoryInner(v *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) GetEntitiesResponseDataObjectsInnerInventoryInner {
	return GetEntitiesResponseDataObjectsInnerInventoryInner{
		GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf: v,
	}
}

// GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf1AsGetEntitiesResponseDataObjectsInnerInventoryInner is a convenience function that returns GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf1 wrapped in GetEntitiesResponseDataObjectsInnerInventoryInner
func GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf1AsGetEntitiesResponseDataObjectsInnerInventoryInner(v *GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf1) GetEntitiesResponseDataObjectsInnerInventoryInner {
	return GetEntitiesResponseDataObjectsInnerInventoryInner{
		GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf1: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetEntitiesResponseDataObjectsInnerInventoryInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf
	err = newStrictDecoder(data).Decode(&dst.GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf)
	if err == nil {
		jsonGetEntitiesResponseDataObjectsInnerInventoryInnerOneOf, _ := json.Marshal(dst.GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf)
		if string(jsonGetEntitiesResponseDataObjectsInnerInventoryInnerOneOf) == "{}" { // empty struct
			dst.GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf = nil
		} else {
			match++
		}
	} else {
		dst.GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf = nil
	}

	// try to unmarshal data into GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf1
	err = newStrictDecoder(data).Decode(&dst.GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf1)
	if err == nil {
		jsonGetEntitiesResponseDataObjectsInnerInventoryInnerOneOf1, _ := json.Marshal(dst.GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf1)
		if string(jsonGetEntitiesResponseDataObjectsInnerInventoryInnerOneOf1) == "{}" { // empty struct
			dst.GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf = nil
		dst.GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetEntitiesResponseDataObjectsInnerInventoryInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetEntitiesResponseDataObjectsInnerInventoryInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetEntitiesResponseDataObjectsInnerInventoryInner) MarshalJSON() ([]byte, error) {
	if src.GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf != nil {
		return json.Marshal(&src.GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf)
	}

	if src.GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf1 != nil {
		return json.Marshal(&src.GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetEntitiesResponseDataObjectsInnerInventoryInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf != nil {
		return obj.GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf
	}

	if obj.GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf1 != nil {
		return obj.GetEntitiesResponseDataObjectsInnerInventoryInnerOneOf1
	}

	// all schemas are nil
	return nil
}

type NullableGetEntitiesResponseDataObjectsInnerInventoryInner struct {
	value *GetEntitiesResponseDataObjectsInnerInventoryInner
	isSet bool
}

func (v NullableGetEntitiesResponseDataObjectsInnerInventoryInner) Get() *GetEntitiesResponseDataObjectsInnerInventoryInner {
	return v.value
}

func (v *NullableGetEntitiesResponseDataObjectsInnerInventoryInner) Set(val *GetEntitiesResponseDataObjectsInnerInventoryInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEntitiesResponseDataObjectsInnerInventoryInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEntitiesResponseDataObjectsInnerInventoryInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEntitiesResponseDataObjectsInnerInventoryInner(val *GetEntitiesResponseDataObjectsInnerInventoryInner) *NullableGetEntitiesResponseDataObjectsInnerInventoryInner {
	return &NullableGetEntitiesResponseDataObjectsInnerInventoryInner{value: val, isSet: true}
}

func (v NullableGetEntitiesResponseDataObjectsInnerInventoryInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEntitiesResponseDataObjectsInnerInventoryInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

