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

// checks if the CreateThresholdRuleRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateThresholdRuleRequestBodyFields{}

// CreateThresholdRuleRequestBodyFields struct for CreateThresholdRuleRequestBodyFields
type CreateThresholdRuleRequestBodyFields struct {
	AssetField *string `json:"assetField,omitempty"`
	Category *string `json:"category,omitempty"`
	Enabled bool `json:"enabled"`
	EntitySelectors []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner `json:"entitySelectors,omitempty"`
	IsPrototype *bool `json:"isPrototype,omitempty"`
	Name string `json:"name"`
	ParentJaskId *string `json:"parentJaskId,omitempty"`
	SummaryExpression *string `json:"summaryExpression,omitempty"`
	Tags []string `json:"tags,omitempty"`
	TuningExpressionIds []string `json:"tuningExpressionIds,omitempty"`
	// The description to be used on the generated Signal
	Description string `json:"description"`
	CountDistinct bool `json:"countDistinct"`
	CountField *string `json:"countField,omitempty"`
	Expression string `json:"expression"`
	Limit int32 `json:"limit"`
	Score int32 `json:"score"`
	Stream string `json:"stream"`
	Version int32 `json:"version"`
	WindowSize string `json:"windowSize"`
	// Can be used instead of windowSize.
	WindowSizeMilliseconds *string `json:"windowSizeMilliseconds,omitempty"`
	GroupByFields []string `json:"groupByFields,omitempty"`
}

type _CreateThresholdRuleRequestBodyFields CreateThresholdRuleRequestBodyFields

// NewCreateThresholdRuleRequestBodyFields instantiates a new CreateThresholdRuleRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateThresholdRuleRequestBodyFields(enabled bool, name string, description string, countDistinct bool, expression string, limit int32, score int32, stream string, version int32, windowSize string) *CreateThresholdRuleRequestBodyFields {
	this := CreateThresholdRuleRequestBodyFields{}
	this.Enabled = enabled
	this.Name = name
	this.Description = description
	this.CountDistinct = countDistinct
	this.Expression = expression
	this.Limit = limit
	this.Score = score
	this.Stream = stream
	this.Version = version
	this.WindowSize = windowSize
	return &this
}

// NewCreateThresholdRuleRequestBodyFieldsWithDefaults instantiates a new CreateThresholdRuleRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateThresholdRuleRequestBodyFieldsWithDefaults() *CreateThresholdRuleRequestBodyFields {
	this := CreateThresholdRuleRequestBodyFields{}
	return &this
}

// GetAssetField returns the AssetField field value if set, zero value otherwise.
func (o *CreateThresholdRuleRequestBodyFields) GetAssetField() string {
	if o == nil || IsNil(o.AssetField) {
		var ret string
		return ret
	}
	return *o.AssetField
}

// GetAssetFieldOk returns a tuple with the AssetField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThresholdRuleRequestBodyFields) GetAssetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.AssetField) {
		return nil, false
	}
	return o.AssetField, true
}

// HasAssetField returns a boolean if a field has been set.
func (o *CreateThresholdRuleRequestBodyFields) HasAssetField() bool {
	if o != nil && !IsNil(o.AssetField) {
		return true
	}

	return false
}

// SetAssetField gets a reference to the given string and assigns it to the AssetField field.
func (o *CreateThresholdRuleRequestBodyFields) SetAssetField(v string) {
	o.AssetField = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *CreateThresholdRuleRequestBodyFields) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThresholdRuleRequestBodyFields) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *CreateThresholdRuleRequestBodyFields) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *CreateThresholdRuleRequestBodyFields) SetCategory(v string) {
	o.Category = &v
}

// GetEnabled returns the Enabled field value
func (o *CreateThresholdRuleRequestBodyFields) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CreateThresholdRuleRequestBodyFields) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CreateThresholdRuleRequestBodyFields) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEntitySelectors returns the EntitySelectors field value if set, zero value otherwise.
func (o *CreateThresholdRuleRequestBodyFields) GetEntitySelectors() []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner {
	if o == nil || IsNil(o.EntitySelectors) {
		var ret []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner
		return ret
	}
	return o.EntitySelectors
}

// GetEntitySelectorsOk returns a tuple with the EntitySelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThresholdRuleRequestBodyFields) GetEntitySelectorsOk() ([]CreateMatchRuleRequestBodyFieldsEntitySelectorsInner, bool) {
	if o == nil || IsNil(o.EntitySelectors) {
		return nil, false
	}
	return o.EntitySelectors, true
}

// HasEntitySelectors returns a boolean if a field has been set.
func (o *CreateThresholdRuleRequestBodyFields) HasEntitySelectors() bool {
	if o != nil && !IsNil(o.EntitySelectors) {
		return true
	}

	return false
}

// SetEntitySelectors gets a reference to the given []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner and assigns it to the EntitySelectors field.
func (o *CreateThresholdRuleRequestBodyFields) SetEntitySelectors(v []CreateMatchRuleRequestBodyFieldsEntitySelectorsInner) {
	o.EntitySelectors = v
}

// GetIsPrototype returns the IsPrototype field value if set, zero value otherwise.
func (o *CreateThresholdRuleRequestBodyFields) GetIsPrototype() bool {
	if o == nil || IsNil(o.IsPrototype) {
		var ret bool
		return ret
	}
	return *o.IsPrototype
}

// GetIsPrototypeOk returns a tuple with the IsPrototype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThresholdRuleRequestBodyFields) GetIsPrototypeOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrototype) {
		return nil, false
	}
	return o.IsPrototype, true
}

// HasIsPrototype returns a boolean if a field has been set.
func (o *CreateThresholdRuleRequestBodyFields) HasIsPrototype() bool {
	if o != nil && !IsNil(o.IsPrototype) {
		return true
	}

	return false
}

// SetIsPrototype gets a reference to the given bool and assigns it to the IsPrototype field.
func (o *CreateThresholdRuleRequestBodyFields) SetIsPrototype(v bool) {
	o.IsPrototype = &v
}

// GetName returns the Name field value
func (o *CreateThresholdRuleRequestBodyFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateThresholdRuleRequestBodyFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateThresholdRuleRequestBodyFields) SetName(v string) {
	o.Name = v
}

// GetParentJaskId returns the ParentJaskId field value if set, zero value otherwise.
func (o *CreateThresholdRuleRequestBodyFields) GetParentJaskId() string {
	if o == nil || IsNil(o.ParentJaskId) {
		var ret string
		return ret
	}
	return *o.ParentJaskId
}

// GetParentJaskIdOk returns a tuple with the ParentJaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThresholdRuleRequestBodyFields) GetParentJaskIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentJaskId) {
		return nil, false
	}
	return o.ParentJaskId, true
}

// HasParentJaskId returns a boolean if a field has been set.
func (o *CreateThresholdRuleRequestBodyFields) HasParentJaskId() bool {
	if o != nil && !IsNil(o.ParentJaskId) {
		return true
	}

	return false
}

// SetParentJaskId gets a reference to the given string and assigns it to the ParentJaskId field.
func (o *CreateThresholdRuleRequestBodyFields) SetParentJaskId(v string) {
	o.ParentJaskId = &v
}

// GetSummaryExpression returns the SummaryExpression field value if set, zero value otherwise.
func (o *CreateThresholdRuleRequestBodyFields) GetSummaryExpression() string {
	if o == nil || IsNil(o.SummaryExpression) {
		var ret string
		return ret
	}
	return *o.SummaryExpression
}

// GetSummaryExpressionOk returns a tuple with the SummaryExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThresholdRuleRequestBodyFields) GetSummaryExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.SummaryExpression) {
		return nil, false
	}
	return o.SummaryExpression, true
}

// HasSummaryExpression returns a boolean if a field has been set.
func (o *CreateThresholdRuleRequestBodyFields) HasSummaryExpression() bool {
	if o != nil && !IsNil(o.SummaryExpression) {
		return true
	}

	return false
}

// SetSummaryExpression gets a reference to the given string and assigns it to the SummaryExpression field.
func (o *CreateThresholdRuleRequestBodyFields) SetSummaryExpression(v string) {
	o.SummaryExpression = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateThresholdRuleRequestBodyFields) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThresholdRuleRequestBodyFields) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateThresholdRuleRequestBodyFields) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateThresholdRuleRequestBodyFields) SetTags(v []string) {
	o.Tags = v
}

// GetTuningExpressionIds returns the TuningExpressionIds field value if set, zero value otherwise.
func (o *CreateThresholdRuleRequestBodyFields) GetTuningExpressionIds() []string {
	if o == nil || IsNil(o.TuningExpressionIds) {
		var ret []string
		return ret
	}
	return o.TuningExpressionIds
}

// GetTuningExpressionIdsOk returns a tuple with the TuningExpressionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThresholdRuleRequestBodyFields) GetTuningExpressionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TuningExpressionIds) {
		return nil, false
	}
	return o.TuningExpressionIds, true
}

// HasTuningExpressionIds returns a boolean if a field has been set.
func (o *CreateThresholdRuleRequestBodyFields) HasTuningExpressionIds() bool {
	if o != nil && !IsNil(o.TuningExpressionIds) {
		return true
	}

	return false
}

// SetTuningExpressionIds gets a reference to the given []string and assigns it to the TuningExpressionIds field.
func (o *CreateThresholdRuleRequestBodyFields) SetTuningExpressionIds(v []string) {
	o.TuningExpressionIds = v
}

// GetDescription returns the Description field value
func (o *CreateThresholdRuleRequestBodyFields) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateThresholdRuleRequestBodyFields) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateThresholdRuleRequestBodyFields) SetDescription(v string) {
	o.Description = v
}

// GetCountDistinct returns the CountDistinct field value
func (o *CreateThresholdRuleRequestBodyFields) GetCountDistinct() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CountDistinct
}

// GetCountDistinctOk returns a tuple with the CountDistinct field value
// and a boolean to check if the value has been set.
func (o *CreateThresholdRuleRequestBodyFields) GetCountDistinctOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountDistinct, true
}

// SetCountDistinct sets field value
func (o *CreateThresholdRuleRequestBodyFields) SetCountDistinct(v bool) {
	o.CountDistinct = v
}

// GetCountField returns the CountField field value if set, zero value otherwise.
func (o *CreateThresholdRuleRequestBodyFields) GetCountField() string {
	if o == nil || IsNil(o.CountField) {
		var ret string
		return ret
	}
	return *o.CountField
}

// GetCountFieldOk returns a tuple with the CountField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThresholdRuleRequestBodyFields) GetCountFieldOk() (*string, bool) {
	if o == nil || IsNil(o.CountField) {
		return nil, false
	}
	return o.CountField, true
}

// HasCountField returns a boolean if a field has been set.
func (o *CreateThresholdRuleRequestBodyFields) HasCountField() bool {
	if o != nil && !IsNil(o.CountField) {
		return true
	}

	return false
}

// SetCountField gets a reference to the given string and assigns it to the CountField field.
func (o *CreateThresholdRuleRequestBodyFields) SetCountField(v string) {
	o.CountField = &v
}

// GetExpression returns the Expression field value
func (o *CreateThresholdRuleRequestBodyFields) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *CreateThresholdRuleRequestBodyFields) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *CreateThresholdRuleRequestBodyFields) SetExpression(v string) {
	o.Expression = v
}

// GetLimit returns the Limit field value
func (o *CreateThresholdRuleRequestBodyFields) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *CreateThresholdRuleRequestBodyFields) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *CreateThresholdRuleRequestBodyFields) SetLimit(v int32) {
	o.Limit = v
}

// GetScore returns the Score field value
func (o *CreateThresholdRuleRequestBodyFields) GetScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *CreateThresholdRuleRequestBodyFields) GetScoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *CreateThresholdRuleRequestBodyFields) SetScore(v int32) {
	o.Score = v
}

// GetStream returns the Stream field value
func (o *CreateThresholdRuleRequestBodyFields) GetStream() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stream
}

// GetStreamOk returns a tuple with the Stream field value
// and a boolean to check if the value has been set.
func (o *CreateThresholdRuleRequestBodyFields) GetStreamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stream, true
}

// SetStream sets field value
func (o *CreateThresholdRuleRequestBodyFields) SetStream(v string) {
	o.Stream = v
}

// GetVersion returns the Version field value
func (o *CreateThresholdRuleRequestBodyFields) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *CreateThresholdRuleRequestBodyFields) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *CreateThresholdRuleRequestBodyFields) SetVersion(v int32) {
	o.Version = v
}

// GetWindowSize returns the WindowSize field value
func (o *CreateThresholdRuleRequestBodyFields) GetWindowSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WindowSize
}

// GetWindowSizeOk returns a tuple with the WindowSize field value
// and a boolean to check if the value has been set.
func (o *CreateThresholdRuleRequestBodyFields) GetWindowSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WindowSize, true
}

// SetWindowSize sets field value
func (o *CreateThresholdRuleRequestBodyFields) SetWindowSize(v string) {
	o.WindowSize = v
}

// GetWindowSizeMilliseconds returns the WindowSizeMilliseconds field value if set, zero value otherwise.
func (o *CreateThresholdRuleRequestBodyFields) GetWindowSizeMilliseconds() string {
	if o == nil || IsNil(o.WindowSizeMilliseconds) {
		var ret string
		return ret
	}
	return *o.WindowSizeMilliseconds
}

// GetWindowSizeMillisecondsOk returns a tuple with the WindowSizeMilliseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThresholdRuleRequestBodyFields) GetWindowSizeMillisecondsOk() (*string, bool) {
	if o == nil || IsNil(o.WindowSizeMilliseconds) {
		return nil, false
	}
	return o.WindowSizeMilliseconds, true
}

// HasWindowSizeMilliseconds returns a boolean if a field has been set.
func (o *CreateThresholdRuleRequestBodyFields) HasWindowSizeMilliseconds() bool {
	if o != nil && !IsNil(o.WindowSizeMilliseconds) {
		return true
	}

	return false
}

// SetWindowSizeMilliseconds gets a reference to the given string and assigns it to the WindowSizeMilliseconds field.
func (o *CreateThresholdRuleRequestBodyFields) SetWindowSizeMilliseconds(v string) {
	o.WindowSizeMilliseconds = &v
}

// GetGroupByFields returns the GroupByFields field value if set, zero value otherwise.
func (o *CreateThresholdRuleRequestBodyFields) GetGroupByFields() []string {
	if o == nil || IsNil(o.GroupByFields) {
		var ret []string
		return ret
	}
	return o.GroupByFields
}

// GetGroupByFieldsOk returns a tuple with the GroupByFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThresholdRuleRequestBodyFields) GetGroupByFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupByFields) {
		return nil, false
	}
	return o.GroupByFields, true
}

// HasGroupByFields returns a boolean if a field has been set.
func (o *CreateThresholdRuleRequestBodyFields) HasGroupByFields() bool {
	if o != nil && !IsNil(o.GroupByFields) {
		return true
	}

	return false
}

// SetGroupByFields gets a reference to the given []string and assigns it to the GroupByFields field.
func (o *CreateThresholdRuleRequestBodyFields) SetGroupByFields(v []string) {
	o.GroupByFields = v
}

func (o CreateThresholdRuleRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateThresholdRuleRequestBodyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetField) {
		toSerialize["assetField"] = o.AssetField
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.EntitySelectors) {
		toSerialize["entitySelectors"] = o.EntitySelectors
	}
	if !IsNil(o.IsPrototype) {
		toSerialize["isPrototype"] = o.IsPrototype
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.ParentJaskId) {
		toSerialize["parentJaskId"] = o.ParentJaskId
	}
	if !IsNil(o.SummaryExpression) {
		toSerialize["summaryExpression"] = o.SummaryExpression
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.TuningExpressionIds) {
		toSerialize["tuningExpressionIds"] = o.TuningExpressionIds
	}
	toSerialize["description"] = o.Description
	toSerialize["countDistinct"] = o.CountDistinct
	if !IsNil(o.CountField) {
		toSerialize["countField"] = o.CountField
	}
	toSerialize["expression"] = o.Expression
	toSerialize["limit"] = o.Limit
	toSerialize["score"] = o.Score
	toSerialize["stream"] = o.Stream
	toSerialize["version"] = o.Version
	toSerialize["windowSize"] = o.WindowSize
	if !IsNil(o.WindowSizeMilliseconds) {
		toSerialize["windowSizeMilliseconds"] = o.WindowSizeMilliseconds
	}
	if !IsNil(o.GroupByFields) {
		toSerialize["groupByFields"] = o.GroupByFields
	}
	return toSerialize, nil
}

func (o *CreateThresholdRuleRequestBodyFields) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"name",
		"description",
		"countDistinct",
		"expression",
		"limit",
		"score",
		"stream",
		"version",
		"windowSize",
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

	varCreateThresholdRuleRequestBodyFields := _CreateThresholdRuleRequestBodyFields{}

	err = json.Unmarshal(bytes, &varCreateThresholdRuleRequestBodyFields)

	if err != nil {
		return err
	}

	*o = CreateThresholdRuleRequestBodyFields(varCreateThresholdRuleRequestBodyFields)

	return err
}

type NullableCreateThresholdRuleRequestBodyFields struct {
	value *CreateThresholdRuleRequestBodyFields
	isSet bool
}

func (v NullableCreateThresholdRuleRequestBodyFields) Get() *CreateThresholdRuleRequestBodyFields {
	return v.value
}

func (v *NullableCreateThresholdRuleRequestBodyFields) Set(val *CreateThresholdRuleRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateThresholdRuleRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateThresholdRuleRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateThresholdRuleRequestBodyFields(val *CreateThresholdRuleRequestBodyFields) *NullableCreateThresholdRuleRequestBodyFields {
	return &NullableCreateThresholdRuleRequestBodyFields{value: val, isSet: true}
}

func (v NullableCreateThresholdRuleRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateThresholdRuleRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

