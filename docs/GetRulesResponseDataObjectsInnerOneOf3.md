# GetRulesResponseDataObjectsInnerOneOf3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**ContentType** | **string** |  | 
**Created** | **time.Time** |  | 
**CreatedBy** | **string** |  | 
**Deleted** | **bool** |  | 
**Enabled** | **bool** |  | 
**EntitySelectors** | [**[]GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner**](GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner.md) |  | 
**Id** | **string** |  | 
**IsPrototype** | **bool** |  | 
**LastUpdated** | **time.Time** |  | 
**LastUpdatedBy** | **string** |  | 
**Name** | **string** |  | 
**ParentJaskId** | Pointer to **string** |  | [optional] 
**Status** | [**GetRulesResponseDataObjectsInnerOneOfStatus**](GetRulesResponseDataObjectsInnerOneOfStatus.md) |  | 
**RuleId** | **int32** |  | 
**RuleSource** | **string** |  | 
**RuleType** | **string** |  | 
**SignalCount07d** | **int32** | The number of Signals generated by this Rule in the past 7 days | 
**SignalCount24h** | **int32** | The number of Signals generated by this Rule in the past 24 hours | 
**SummaryExpression** | **string** |  | 
**Tags** | **[]string** |  | 
**HasOverride** | **bool** |  | 
**TagsOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOfTagsOverride**](GetRulesResponseDataObjectsInnerOneOfTagsOverride.md) |  | [optional] 
**IsPrototypeOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOfIsPrototypeOverride**](GetRulesResponseDataObjectsInnerOneOfIsPrototypeOverride.md) |  | [optional] 
**EntitySelectorsOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride**](GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride.md) |  | [optional] 
**NameOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOfNameOverride**](GetRulesResponseDataObjectsInnerOneOfNameOverride.md) |  | [optional] 
**SummaryExpressionOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOfNameOverride**](GetRulesResponseDataObjectsInnerOneOfNameOverride.md) |  | [optional] 
**AssetField** | **string** |  | 
**Description** | **string** | The description to be used on the generated Signal | 
**ExpressionsAndLimits** | [**[]GetRulesResponseDataObjectsInnerOneOf3ExpressionsAndLimitsInner**](GetRulesResponseDataObjectsInnerOneOf3ExpressionsAndLimitsInner.md) |  | 
**GroupByFields** | **[]string** |  | 
**Ordered** | **bool** |  | 
**Score** | **int32** |  | 
**Stream** | **string** |  | 
**Version** | **int32** |  | 
**WindowSize** | **int32** | milliseconds | 
**WindowSizeName** | **string** |  | 
**DescriptionOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOfNameOverride**](GetRulesResponseDataObjectsInnerOneOfNameOverride.md) |  | [optional] 
**GroupByFieldsOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOfTagsOverride**](GetRulesResponseDataObjectsInnerOneOfTagsOverride.md) |  | [optional] 
**ScoreOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOf1ScoreOverride**](GetRulesResponseDataObjectsInnerOneOf1ScoreOverride.md) |  | [optional] 
**WindowSizeOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOf1WindowSizeOverride**](GetRulesResponseDataObjectsInnerOneOf1WindowSizeOverride.md) |  | [optional] 
**ExpressionsAndLimitsOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOf3ExpressionsAndLimitsOverride**](GetRulesResponseDataObjectsInnerOneOf3ExpressionsAndLimitsOverride.md) |  | [optional] 
**TuningExpressions** | Pointer to [**[]GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner**](GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner.md) |  | [optional] 

## Methods

### NewGetRulesResponseDataObjectsInnerOneOf3

`func NewGetRulesResponseDataObjectsInnerOneOf3(contentType string, created time.Time, createdBy string, deleted bool, enabled bool, entitySelectors []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner, id string, isPrototype bool, lastUpdated time.Time, lastUpdatedBy string, name string, status GetRulesResponseDataObjectsInnerOneOfStatus, ruleId int32, ruleSource string, ruleType string, signalCount07d int32, signalCount24h int32, summaryExpression string, tags []string, hasOverride bool, assetField string, description string, expressionsAndLimits []GetRulesResponseDataObjectsInnerOneOf3ExpressionsAndLimitsInner, groupByFields []string, ordered bool, score int32, stream string, version int32, windowSize int32, windowSizeName string, ) *GetRulesResponseDataObjectsInnerOneOf3`

NewGetRulesResponseDataObjectsInnerOneOf3 instantiates a new GetRulesResponseDataObjectsInnerOneOf3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRulesResponseDataObjectsInnerOneOf3WithDefaults

`func NewGetRulesResponseDataObjectsInnerOneOf3WithDefaults() *GetRulesResponseDataObjectsInnerOneOf3`

NewGetRulesResponseDataObjectsInnerOneOf3WithDefaults instantiates a new GetRulesResponseDataObjectsInnerOneOf3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetRulesResponseDataObjectsInnerOneOf3) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetContentType

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetCreated

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCreatedBy

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDeleted

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetEnabled

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEntitySelectors

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetEntitySelectors() []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner`

GetEntitySelectors returns the EntitySelectors field if non-nil, zero value otherwise.

### GetEntitySelectorsOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetEntitySelectorsOk() (*[]GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner, bool)`

GetEntitySelectorsOk returns a tuple with the EntitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelectors

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetEntitySelectors(v []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner)`

SetEntitySelectors sets EntitySelectors field to given value.


### GetId

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetId(v string)`

SetId sets Id field to given value.


### GetIsPrototype

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetIsPrototype() bool`

GetIsPrototype returns the IsPrototype field if non-nil, zero value otherwise.

### GetIsPrototypeOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetIsPrototypeOk() (*bool, bool)`

GetIsPrototypeOk returns a tuple with the IsPrototype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrototype

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetIsPrototype(v bool)`

SetIsPrototype sets IsPrototype field to given value.


### GetLastUpdated

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.


### GetName

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetName(v string)`

SetName sets Name field to given value.


### GetParentJaskId

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetParentJaskId() string`

GetParentJaskId returns the ParentJaskId field if non-nil, zero value otherwise.

### GetParentJaskIdOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetParentJaskIdOk() (*string, bool)`

GetParentJaskIdOk returns a tuple with the ParentJaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentJaskId

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetParentJaskId(v string)`

SetParentJaskId sets ParentJaskId field to given value.

### HasParentJaskId

`func (o *GetRulesResponseDataObjectsInnerOneOf3) HasParentJaskId() bool`

HasParentJaskId returns a boolean if a field has been set.

### GetStatus

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetStatus() GetRulesResponseDataObjectsInnerOneOfStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetStatusOk() (*GetRulesResponseDataObjectsInnerOneOfStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetStatus(v GetRulesResponseDataObjectsInnerOneOfStatus)`

SetStatus sets Status field to given value.


### GetRuleId

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetRuleId() int32`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetRuleIdOk() (*int32, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetRuleId(v int32)`

SetRuleId sets RuleId field to given value.


### GetRuleSource

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetRuleSource() string`

GetRuleSource returns the RuleSource field if non-nil, zero value otherwise.

### GetRuleSourceOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetRuleSourceOk() (*string, bool)`

GetRuleSourceOk returns a tuple with the RuleSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSource

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetRuleSource(v string)`

SetRuleSource sets RuleSource field to given value.


### GetRuleType

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.


### GetSignalCount07d

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetSignalCount07d() int32`

GetSignalCount07d returns the SignalCount07d field if non-nil, zero value otherwise.

### GetSignalCount07dOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetSignalCount07dOk() (*int32, bool)`

GetSignalCount07dOk returns a tuple with the SignalCount07d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalCount07d

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetSignalCount07d(v int32)`

SetSignalCount07d sets SignalCount07d field to given value.


### GetSignalCount24h

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetSignalCount24h() int32`

GetSignalCount24h returns the SignalCount24h field if non-nil, zero value otherwise.

### GetSignalCount24hOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetSignalCount24hOk() (*int32, bool)`

GetSignalCount24hOk returns a tuple with the SignalCount24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalCount24h

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetSignalCount24h(v int32)`

SetSignalCount24h sets SignalCount24h field to given value.


### GetSummaryExpression

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetSummaryExpression() string`

GetSummaryExpression returns the SummaryExpression field if non-nil, zero value otherwise.

### GetSummaryExpressionOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetSummaryExpressionOk() (*string, bool)`

GetSummaryExpressionOk returns a tuple with the SummaryExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryExpression

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetSummaryExpression(v string)`

SetSummaryExpression sets SummaryExpression field to given value.


### GetTags

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetHasOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetHasOverride() bool`

GetHasOverride returns the HasOverride field if non-nil, zero value otherwise.

### GetHasOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetHasOverrideOk() (*bool, bool)`

GetHasOverrideOk returns a tuple with the HasOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetHasOverride(v bool)`

SetHasOverride sets HasOverride field to given value.


### GetTagsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetTagsOverride() GetRulesResponseDataObjectsInnerOneOfTagsOverride`

GetTagsOverride returns the TagsOverride field if non-nil, zero value otherwise.

### GetTagsOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetTagsOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfTagsOverride, bool)`

GetTagsOverrideOk returns a tuple with the TagsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetTagsOverride(v GetRulesResponseDataObjectsInnerOneOfTagsOverride)`

SetTagsOverride sets TagsOverride field to given value.

### HasTagsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) HasTagsOverride() bool`

HasTagsOverride returns a boolean if a field has been set.

### GetIsPrototypeOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetIsPrototypeOverride() GetRulesResponseDataObjectsInnerOneOfIsPrototypeOverride`

GetIsPrototypeOverride returns the IsPrototypeOverride field if non-nil, zero value otherwise.

### GetIsPrototypeOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetIsPrototypeOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfIsPrototypeOverride, bool)`

GetIsPrototypeOverrideOk returns a tuple with the IsPrototypeOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrototypeOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetIsPrototypeOverride(v GetRulesResponseDataObjectsInnerOneOfIsPrototypeOverride)`

SetIsPrototypeOverride sets IsPrototypeOverride field to given value.

### HasIsPrototypeOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) HasIsPrototypeOverride() bool`

HasIsPrototypeOverride returns a boolean if a field has been set.

### GetEntitySelectorsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetEntitySelectorsOverride() GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride`

GetEntitySelectorsOverride returns the EntitySelectorsOverride field if non-nil, zero value otherwise.

### GetEntitySelectorsOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetEntitySelectorsOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride, bool)`

GetEntitySelectorsOverrideOk returns a tuple with the EntitySelectorsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelectorsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetEntitySelectorsOverride(v GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride)`

SetEntitySelectorsOverride sets EntitySelectorsOverride field to given value.

### HasEntitySelectorsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) HasEntitySelectorsOverride() bool`

HasEntitySelectorsOverride returns a boolean if a field has been set.

### GetNameOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetNameOverride() GetRulesResponseDataObjectsInnerOneOfNameOverride`

GetNameOverride returns the NameOverride field if non-nil, zero value otherwise.

### GetNameOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetNameOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfNameOverride, bool)`

GetNameOverrideOk returns a tuple with the NameOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetNameOverride(v GetRulesResponseDataObjectsInnerOneOfNameOverride)`

SetNameOverride sets NameOverride field to given value.

### HasNameOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) HasNameOverride() bool`

HasNameOverride returns a boolean if a field has been set.

### GetSummaryExpressionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetSummaryExpressionOverride() GetRulesResponseDataObjectsInnerOneOfNameOverride`

GetSummaryExpressionOverride returns the SummaryExpressionOverride field if non-nil, zero value otherwise.

### GetSummaryExpressionOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetSummaryExpressionOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfNameOverride, bool)`

GetSummaryExpressionOverrideOk returns a tuple with the SummaryExpressionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryExpressionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetSummaryExpressionOverride(v GetRulesResponseDataObjectsInnerOneOfNameOverride)`

SetSummaryExpressionOverride sets SummaryExpressionOverride field to given value.

### HasSummaryExpressionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) HasSummaryExpressionOverride() bool`

HasSummaryExpressionOverride returns a boolean if a field has been set.

### GetAssetField

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetAssetField() string`

GetAssetField returns the AssetField field if non-nil, zero value otherwise.

### GetAssetFieldOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetAssetFieldOk() (*string, bool)`

GetAssetFieldOk returns a tuple with the AssetField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetField

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetAssetField(v string)`

SetAssetField sets AssetField field to given value.


### GetDescription

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetExpressionsAndLimits

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetExpressionsAndLimits() []GetRulesResponseDataObjectsInnerOneOf3ExpressionsAndLimitsInner`

GetExpressionsAndLimits returns the ExpressionsAndLimits field if non-nil, zero value otherwise.

### GetExpressionsAndLimitsOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetExpressionsAndLimitsOk() (*[]GetRulesResponseDataObjectsInnerOneOf3ExpressionsAndLimitsInner, bool)`

GetExpressionsAndLimitsOk returns a tuple with the ExpressionsAndLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressionsAndLimits

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetExpressionsAndLimits(v []GetRulesResponseDataObjectsInnerOneOf3ExpressionsAndLimitsInner)`

SetExpressionsAndLimits sets ExpressionsAndLimits field to given value.


### GetGroupByFields

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetGroupByFields() []string`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetGroupByFieldsOk() (*[]string, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetGroupByFields(v []string)`

SetGroupByFields sets GroupByFields field to given value.


### GetOrdered

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetOrdered() bool`

GetOrdered returns the Ordered field if non-nil, zero value otherwise.

### GetOrderedOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetOrderedOk() (*bool, bool)`

GetOrderedOk returns a tuple with the Ordered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdered

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetOrdered(v bool)`

SetOrdered sets Ordered field to given value.


### GetScore

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetScore(v int32)`

SetScore sets Score field to given value.


### GetStream

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetStream(v string)`

SetStream sets Stream field to given value.


### GetVersion

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetWindowSize

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetWindowSize() int32`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetWindowSizeOk() (*int32, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetWindowSize(v int32)`

SetWindowSize sets WindowSize field to given value.


### GetWindowSizeName

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetWindowSizeName() string`

GetWindowSizeName returns the WindowSizeName field if non-nil, zero value otherwise.

### GetWindowSizeNameOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetWindowSizeNameOk() (*string, bool)`

GetWindowSizeNameOk returns a tuple with the WindowSizeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSizeName

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetWindowSizeName(v string)`

SetWindowSizeName sets WindowSizeName field to given value.


### GetDescriptionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetDescriptionOverride() GetRulesResponseDataObjectsInnerOneOfNameOverride`

GetDescriptionOverride returns the DescriptionOverride field if non-nil, zero value otherwise.

### GetDescriptionOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetDescriptionOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfNameOverride, bool)`

GetDescriptionOverrideOk returns a tuple with the DescriptionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetDescriptionOverride(v GetRulesResponseDataObjectsInnerOneOfNameOverride)`

SetDescriptionOverride sets DescriptionOverride field to given value.

### HasDescriptionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) HasDescriptionOverride() bool`

HasDescriptionOverride returns a boolean if a field has been set.

### GetGroupByFieldsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetGroupByFieldsOverride() GetRulesResponseDataObjectsInnerOneOfTagsOverride`

GetGroupByFieldsOverride returns the GroupByFieldsOverride field if non-nil, zero value otherwise.

### GetGroupByFieldsOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetGroupByFieldsOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfTagsOverride, bool)`

GetGroupByFieldsOverrideOk returns a tuple with the GroupByFieldsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFieldsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetGroupByFieldsOverride(v GetRulesResponseDataObjectsInnerOneOfTagsOverride)`

SetGroupByFieldsOverride sets GroupByFieldsOverride field to given value.

### HasGroupByFieldsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) HasGroupByFieldsOverride() bool`

HasGroupByFieldsOverride returns a boolean if a field has been set.

### GetScoreOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetScoreOverride() GetRulesResponseDataObjectsInnerOneOf1ScoreOverride`

GetScoreOverride returns the ScoreOverride field if non-nil, zero value otherwise.

### GetScoreOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetScoreOverrideOk() (*GetRulesResponseDataObjectsInnerOneOf1ScoreOverride, bool)`

GetScoreOverrideOk returns a tuple with the ScoreOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetScoreOverride(v GetRulesResponseDataObjectsInnerOneOf1ScoreOverride)`

SetScoreOverride sets ScoreOverride field to given value.

### HasScoreOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) HasScoreOverride() bool`

HasScoreOverride returns a boolean if a field has been set.

### GetWindowSizeOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetWindowSizeOverride() GetRulesResponseDataObjectsInnerOneOf1WindowSizeOverride`

GetWindowSizeOverride returns the WindowSizeOverride field if non-nil, zero value otherwise.

### GetWindowSizeOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetWindowSizeOverrideOk() (*GetRulesResponseDataObjectsInnerOneOf1WindowSizeOverride, bool)`

GetWindowSizeOverrideOk returns a tuple with the WindowSizeOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSizeOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetWindowSizeOverride(v GetRulesResponseDataObjectsInnerOneOf1WindowSizeOverride)`

SetWindowSizeOverride sets WindowSizeOverride field to given value.

### HasWindowSizeOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) HasWindowSizeOverride() bool`

HasWindowSizeOverride returns a boolean if a field has been set.

### GetExpressionsAndLimitsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetExpressionsAndLimitsOverride() GetRulesResponseDataObjectsInnerOneOf3ExpressionsAndLimitsOverride`

GetExpressionsAndLimitsOverride returns the ExpressionsAndLimitsOverride field if non-nil, zero value otherwise.

### GetExpressionsAndLimitsOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetExpressionsAndLimitsOverrideOk() (*GetRulesResponseDataObjectsInnerOneOf3ExpressionsAndLimitsOverride, bool)`

GetExpressionsAndLimitsOverrideOk returns a tuple with the ExpressionsAndLimitsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressionsAndLimitsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetExpressionsAndLimitsOverride(v GetRulesResponseDataObjectsInnerOneOf3ExpressionsAndLimitsOverride)`

SetExpressionsAndLimitsOverride sets ExpressionsAndLimitsOverride field to given value.

### HasExpressionsAndLimitsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf3) HasExpressionsAndLimitsOverride() bool`

HasExpressionsAndLimitsOverride returns a boolean if a field has been set.

### GetTuningExpressions

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetTuningExpressions() []GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner`

GetTuningExpressions returns the TuningExpressions field if non-nil, zero value otherwise.

### GetTuningExpressionsOk

`func (o *GetRulesResponseDataObjectsInnerOneOf3) GetTuningExpressionsOk() (*[]GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner, bool)`

GetTuningExpressionsOk returns a tuple with the TuningExpressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuningExpressions

`func (o *GetRulesResponseDataObjectsInnerOneOf3) SetTuningExpressions(v []GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner)`

SetTuningExpressions sets TuningExpressions field to given value.

### HasTuningExpressions

`func (o *GetRulesResponseDataObjectsInnerOneOf3) HasTuningExpressions() bool`

HasTuningExpressions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


