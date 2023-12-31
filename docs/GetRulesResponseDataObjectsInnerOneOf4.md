# GetRulesResponseDataObjectsInnerOneOf4

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
**NameExpression** | **string** |  | 
**DescriptionExpression** | **string** |  | 
**MatchExpression** | **string** |  | 
**GroupByAsset** | **bool** |  | 
**GroupByFields** | **[]string** |  | 
**AggregationFunctions** | [**[]GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner**](GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner.md) |  | 
**ScoreMapping** | [**GetRulesResponseDataObjectsInnerOneOf2ScoreMapping**](GetRulesResponseDataObjectsInnerOneOf2ScoreMapping.md) |  | 
**Stream** | **string** |  | 
**Version** | **int32** |  | 
**WindowSize** | **int32** | milliseconds | 
**WindowSizeName** | **string** |  | 
**DescriptionExpressionOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOfNameOverride**](GetRulesResponseDataObjectsInnerOneOfNameOverride.md) |  | [optional] 
**GroupByFieldsOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOfTagsOverride**](GetRulesResponseDataObjectsInnerOneOfTagsOverride.md) |  | [optional] 
**NameExpressionOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOfNameOverride**](GetRulesResponseDataObjectsInnerOneOfNameOverride.md) |  | [optional] 
**ScoreMappingOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOf2ScoreMappingOverride**](GetRulesResponseDataObjectsInnerOneOf2ScoreMappingOverride.md) |  | [optional] 
**TriggerExpression** | **string** |  | 
**WindowSizeOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOf1WindowSizeOverride**](GetRulesResponseDataObjectsInnerOneOf1WindowSizeOverride.md) |  | [optional] 
**TuningExpressions** | Pointer to [**[]GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner**](GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner.md) |  | [optional] 

## Methods

### NewGetRulesResponseDataObjectsInnerOneOf4

`func NewGetRulesResponseDataObjectsInnerOneOf4(contentType string, created time.Time, createdBy string, deleted bool, enabled bool, entitySelectors []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner, id string, isPrototype bool, lastUpdated time.Time, lastUpdatedBy string, name string, status GetRulesResponseDataObjectsInnerOneOfStatus, ruleId int32, ruleSource string, ruleType string, signalCount07d int32, signalCount24h int32, summaryExpression string, tags []string, hasOverride bool, assetField string, nameExpression string, descriptionExpression string, matchExpression string, groupByAsset bool, groupByFields []string, aggregationFunctions []GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner, scoreMapping GetRulesResponseDataObjectsInnerOneOf2ScoreMapping, stream string, version int32, windowSize int32, windowSizeName string, triggerExpression string, ) *GetRulesResponseDataObjectsInnerOneOf4`

NewGetRulesResponseDataObjectsInnerOneOf4 instantiates a new GetRulesResponseDataObjectsInnerOneOf4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRulesResponseDataObjectsInnerOneOf4WithDefaults

`func NewGetRulesResponseDataObjectsInnerOneOf4WithDefaults() *GetRulesResponseDataObjectsInnerOneOf4`

NewGetRulesResponseDataObjectsInnerOneOf4WithDefaults instantiates a new GetRulesResponseDataObjectsInnerOneOf4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetRulesResponseDataObjectsInnerOneOf4) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetContentType

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetCreated

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCreatedBy

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDeleted

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetEnabled

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEntitySelectors

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetEntitySelectors() []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner`

GetEntitySelectors returns the EntitySelectors field if non-nil, zero value otherwise.

### GetEntitySelectorsOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetEntitySelectorsOk() (*[]GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner, bool)`

GetEntitySelectorsOk returns a tuple with the EntitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelectors

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetEntitySelectors(v []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner)`

SetEntitySelectors sets EntitySelectors field to given value.


### GetId

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetId(v string)`

SetId sets Id field to given value.


### GetIsPrototype

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetIsPrototype() bool`

GetIsPrototype returns the IsPrototype field if non-nil, zero value otherwise.

### GetIsPrototypeOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetIsPrototypeOk() (*bool, bool)`

GetIsPrototypeOk returns a tuple with the IsPrototype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrototype

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetIsPrototype(v bool)`

SetIsPrototype sets IsPrototype field to given value.


### GetLastUpdated

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.


### GetName

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetName(v string)`

SetName sets Name field to given value.


### GetParentJaskId

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetParentJaskId() string`

GetParentJaskId returns the ParentJaskId field if non-nil, zero value otherwise.

### GetParentJaskIdOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetParentJaskIdOk() (*string, bool)`

GetParentJaskIdOk returns a tuple with the ParentJaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentJaskId

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetParentJaskId(v string)`

SetParentJaskId sets ParentJaskId field to given value.

### HasParentJaskId

`func (o *GetRulesResponseDataObjectsInnerOneOf4) HasParentJaskId() bool`

HasParentJaskId returns a boolean if a field has been set.

### GetStatus

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetStatus() GetRulesResponseDataObjectsInnerOneOfStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetStatusOk() (*GetRulesResponseDataObjectsInnerOneOfStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetStatus(v GetRulesResponseDataObjectsInnerOneOfStatus)`

SetStatus sets Status field to given value.


### GetRuleId

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetRuleId() int32`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetRuleIdOk() (*int32, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetRuleId(v int32)`

SetRuleId sets RuleId field to given value.


### GetRuleSource

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetRuleSource() string`

GetRuleSource returns the RuleSource field if non-nil, zero value otherwise.

### GetRuleSourceOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetRuleSourceOk() (*string, bool)`

GetRuleSourceOk returns a tuple with the RuleSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSource

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetRuleSource(v string)`

SetRuleSource sets RuleSource field to given value.


### GetRuleType

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.


### GetSignalCount07d

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetSignalCount07d() int32`

GetSignalCount07d returns the SignalCount07d field if non-nil, zero value otherwise.

### GetSignalCount07dOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetSignalCount07dOk() (*int32, bool)`

GetSignalCount07dOk returns a tuple with the SignalCount07d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalCount07d

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetSignalCount07d(v int32)`

SetSignalCount07d sets SignalCount07d field to given value.


### GetSignalCount24h

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetSignalCount24h() int32`

GetSignalCount24h returns the SignalCount24h field if non-nil, zero value otherwise.

### GetSignalCount24hOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetSignalCount24hOk() (*int32, bool)`

GetSignalCount24hOk returns a tuple with the SignalCount24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalCount24h

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetSignalCount24h(v int32)`

SetSignalCount24h sets SignalCount24h field to given value.


### GetSummaryExpression

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetSummaryExpression() string`

GetSummaryExpression returns the SummaryExpression field if non-nil, zero value otherwise.

### GetSummaryExpressionOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetSummaryExpressionOk() (*string, bool)`

GetSummaryExpressionOk returns a tuple with the SummaryExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryExpression

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetSummaryExpression(v string)`

SetSummaryExpression sets SummaryExpression field to given value.


### GetTags

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetHasOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetHasOverride() bool`

GetHasOverride returns the HasOverride field if non-nil, zero value otherwise.

### GetHasOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetHasOverrideOk() (*bool, bool)`

GetHasOverrideOk returns a tuple with the HasOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetHasOverride(v bool)`

SetHasOverride sets HasOverride field to given value.


### GetTagsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetTagsOverride() GetRulesResponseDataObjectsInnerOneOfTagsOverride`

GetTagsOverride returns the TagsOverride field if non-nil, zero value otherwise.

### GetTagsOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetTagsOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfTagsOverride, bool)`

GetTagsOverrideOk returns a tuple with the TagsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetTagsOverride(v GetRulesResponseDataObjectsInnerOneOfTagsOverride)`

SetTagsOverride sets TagsOverride field to given value.

### HasTagsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) HasTagsOverride() bool`

HasTagsOverride returns a boolean if a field has been set.

### GetIsPrototypeOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetIsPrototypeOverride() GetRulesResponseDataObjectsInnerOneOfIsPrototypeOverride`

GetIsPrototypeOverride returns the IsPrototypeOverride field if non-nil, zero value otherwise.

### GetIsPrototypeOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetIsPrototypeOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfIsPrototypeOverride, bool)`

GetIsPrototypeOverrideOk returns a tuple with the IsPrototypeOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrototypeOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetIsPrototypeOverride(v GetRulesResponseDataObjectsInnerOneOfIsPrototypeOverride)`

SetIsPrototypeOverride sets IsPrototypeOverride field to given value.

### HasIsPrototypeOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) HasIsPrototypeOverride() bool`

HasIsPrototypeOverride returns a boolean if a field has been set.

### GetEntitySelectorsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetEntitySelectorsOverride() GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride`

GetEntitySelectorsOverride returns the EntitySelectorsOverride field if non-nil, zero value otherwise.

### GetEntitySelectorsOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetEntitySelectorsOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride, bool)`

GetEntitySelectorsOverrideOk returns a tuple with the EntitySelectorsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelectorsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetEntitySelectorsOverride(v GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride)`

SetEntitySelectorsOverride sets EntitySelectorsOverride field to given value.

### HasEntitySelectorsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) HasEntitySelectorsOverride() bool`

HasEntitySelectorsOverride returns a boolean if a field has been set.

### GetNameOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetNameOverride() GetRulesResponseDataObjectsInnerOneOfNameOverride`

GetNameOverride returns the NameOverride field if non-nil, zero value otherwise.

### GetNameOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetNameOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfNameOverride, bool)`

GetNameOverrideOk returns a tuple with the NameOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetNameOverride(v GetRulesResponseDataObjectsInnerOneOfNameOverride)`

SetNameOverride sets NameOverride field to given value.

### HasNameOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) HasNameOverride() bool`

HasNameOverride returns a boolean if a field has been set.

### GetSummaryExpressionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetSummaryExpressionOverride() GetRulesResponseDataObjectsInnerOneOfNameOverride`

GetSummaryExpressionOverride returns the SummaryExpressionOverride field if non-nil, zero value otherwise.

### GetSummaryExpressionOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetSummaryExpressionOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfNameOverride, bool)`

GetSummaryExpressionOverrideOk returns a tuple with the SummaryExpressionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryExpressionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetSummaryExpressionOverride(v GetRulesResponseDataObjectsInnerOneOfNameOverride)`

SetSummaryExpressionOverride sets SummaryExpressionOverride field to given value.

### HasSummaryExpressionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) HasSummaryExpressionOverride() bool`

HasSummaryExpressionOverride returns a boolean if a field has been set.

### GetAssetField

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetAssetField() string`

GetAssetField returns the AssetField field if non-nil, zero value otherwise.

### GetAssetFieldOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetAssetFieldOk() (*string, bool)`

GetAssetFieldOk returns a tuple with the AssetField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetField

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetAssetField(v string)`

SetAssetField sets AssetField field to given value.


### GetNameExpression

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetNameExpression() string`

GetNameExpression returns the NameExpression field if non-nil, zero value otherwise.

### GetNameExpressionOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetNameExpressionOk() (*string, bool)`

GetNameExpressionOk returns a tuple with the NameExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameExpression

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetNameExpression(v string)`

SetNameExpression sets NameExpression field to given value.


### GetDescriptionExpression

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetDescriptionExpression() string`

GetDescriptionExpression returns the DescriptionExpression field if non-nil, zero value otherwise.

### GetDescriptionExpressionOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetDescriptionExpressionOk() (*string, bool)`

GetDescriptionExpressionOk returns a tuple with the DescriptionExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionExpression

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetDescriptionExpression(v string)`

SetDescriptionExpression sets DescriptionExpression field to given value.


### GetMatchExpression

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetMatchExpression() string`

GetMatchExpression returns the MatchExpression field if non-nil, zero value otherwise.

### GetMatchExpressionOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetMatchExpressionOk() (*string, bool)`

GetMatchExpressionOk returns a tuple with the MatchExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchExpression

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetMatchExpression(v string)`

SetMatchExpression sets MatchExpression field to given value.


### GetGroupByAsset

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetGroupByAsset() bool`

GetGroupByAsset returns the GroupByAsset field if non-nil, zero value otherwise.

### GetGroupByAssetOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetGroupByAssetOk() (*bool, bool)`

GetGroupByAssetOk returns a tuple with the GroupByAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByAsset

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetGroupByAsset(v bool)`

SetGroupByAsset sets GroupByAsset field to given value.


### GetGroupByFields

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetGroupByFields() []string`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetGroupByFieldsOk() (*[]string, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetGroupByFields(v []string)`

SetGroupByFields sets GroupByFields field to given value.


### GetAggregationFunctions

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetAggregationFunctions() []GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner`

GetAggregationFunctions returns the AggregationFunctions field if non-nil, zero value otherwise.

### GetAggregationFunctionsOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetAggregationFunctionsOk() (*[]GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner, bool)`

GetAggregationFunctionsOk returns a tuple with the AggregationFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationFunctions

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetAggregationFunctions(v []GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner)`

SetAggregationFunctions sets AggregationFunctions field to given value.


### GetScoreMapping

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetScoreMapping() GetRulesResponseDataObjectsInnerOneOf2ScoreMapping`

GetScoreMapping returns the ScoreMapping field if non-nil, zero value otherwise.

### GetScoreMappingOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetScoreMappingOk() (*GetRulesResponseDataObjectsInnerOneOf2ScoreMapping, bool)`

GetScoreMappingOk returns a tuple with the ScoreMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreMapping

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetScoreMapping(v GetRulesResponseDataObjectsInnerOneOf2ScoreMapping)`

SetScoreMapping sets ScoreMapping field to given value.


### GetStream

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetStream(v string)`

SetStream sets Stream field to given value.


### GetVersion

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetWindowSize

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetWindowSize() int32`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetWindowSizeOk() (*int32, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetWindowSize(v int32)`

SetWindowSize sets WindowSize field to given value.


### GetWindowSizeName

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetWindowSizeName() string`

GetWindowSizeName returns the WindowSizeName field if non-nil, zero value otherwise.

### GetWindowSizeNameOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetWindowSizeNameOk() (*string, bool)`

GetWindowSizeNameOk returns a tuple with the WindowSizeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSizeName

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetWindowSizeName(v string)`

SetWindowSizeName sets WindowSizeName field to given value.


### GetDescriptionExpressionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetDescriptionExpressionOverride() GetRulesResponseDataObjectsInnerOneOfNameOverride`

GetDescriptionExpressionOverride returns the DescriptionExpressionOverride field if non-nil, zero value otherwise.

### GetDescriptionExpressionOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetDescriptionExpressionOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfNameOverride, bool)`

GetDescriptionExpressionOverrideOk returns a tuple with the DescriptionExpressionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionExpressionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetDescriptionExpressionOverride(v GetRulesResponseDataObjectsInnerOneOfNameOverride)`

SetDescriptionExpressionOverride sets DescriptionExpressionOverride field to given value.

### HasDescriptionExpressionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) HasDescriptionExpressionOverride() bool`

HasDescriptionExpressionOverride returns a boolean if a field has been set.

### GetGroupByFieldsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetGroupByFieldsOverride() GetRulesResponseDataObjectsInnerOneOfTagsOverride`

GetGroupByFieldsOverride returns the GroupByFieldsOverride field if non-nil, zero value otherwise.

### GetGroupByFieldsOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetGroupByFieldsOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfTagsOverride, bool)`

GetGroupByFieldsOverrideOk returns a tuple with the GroupByFieldsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFieldsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetGroupByFieldsOverride(v GetRulesResponseDataObjectsInnerOneOfTagsOverride)`

SetGroupByFieldsOverride sets GroupByFieldsOverride field to given value.

### HasGroupByFieldsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) HasGroupByFieldsOverride() bool`

HasGroupByFieldsOverride returns a boolean if a field has been set.

### GetNameExpressionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetNameExpressionOverride() GetRulesResponseDataObjectsInnerOneOfNameOverride`

GetNameExpressionOverride returns the NameExpressionOverride field if non-nil, zero value otherwise.

### GetNameExpressionOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetNameExpressionOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfNameOverride, bool)`

GetNameExpressionOverrideOk returns a tuple with the NameExpressionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameExpressionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetNameExpressionOverride(v GetRulesResponseDataObjectsInnerOneOfNameOverride)`

SetNameExpressionOverride sets NameExpressionOverride field to given value.

### HasNameExpressionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) HasNameExpressionOverride() bool`

HasNameExpressionOverride returns a boolean if a field has been set.

### GetScoreMappingOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetScoreMappingOverride() GetRulesResponseDataObjectsInnerOneOf2ScoreMappingOverride`

GetScoreMappingOverride returns the ScoreMappingOverride field if non-nil, zero value otherwise.

### GetScoreMappingOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetScoreMappingOverrideOk() (*GetRulesResponseDataObjectsInnerOneOf2ScoreMappingOverride, bool)`

GetScoreMappingOverrideOk returns a tuple with the ScoreMappingOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreMappingOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetScoreMappingOverride(v GetRulesResponseDataObjectsInnerOneOf2ScoreMappingOverride)`

SetScoreMappingOverride sets ScoreMappingOverride field to given value.

### HasScoreMappingOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) HasScoreMappingOverride() bool`

HasScoreMappingOverride returns a boolean if a field has been set.

### GetTriggerExpression

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetTriggerExpression() string`

GetTriggerExpression returns the TriggerExpression field if non-nil, zero value otherwise.

### GetTriggerExpressionOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetTriggerExpressionOk() (*string, bool)`

GetTriggerExpressionOk returns a tuple with the TriggerExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerExpression

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetTriggerExpression(v string)`

SetTriggerExpression sets TriggerExpression field to given value.


### GetWindowSizeOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetWindowSizeOverride() GetRulesResponseDataObjectsInnerOneOf1WindowSizeOverride`

GetWindowSizeOverride returns the WindowSizeOverride field if non-nil, zero value otherwise.

### GetWindowSizeOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetWindowSizeOverrideOk() (*GetRulesResponseDataObjectsInnerOneOf1WindowSizeOverride, bool)`

GetWindowSizeOverrideOk returns a tuple with the WindowSizeOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSizeOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetWindowSizeOverride(v GetRulesResponseDataObjectsInnerOneOf1WindowSizeOverride)`

SetWindowSizeOverride sets WindowSizeOverride field to given value.

### HasWindowSizeOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf4) HasWindowSizeOverride() bool`

HasWindowSizeOverride returns a boolean if a field has been set.

### GetTuningExpressions

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetTuningExpressions() []GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner`

GetTuningExpressions returns the TuningExpressions field if non-nil, zero value otherwise.

### GetTuningExpressionsOk

`func (o *GetRulesResponseDataObjectsInnerOneOf4) GetTuningExpressionsOk() (*[]GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner, bool)`

GetTuningExpressionsOk returns a tuple with the TuningExpressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuningExpressions

`func (o *GetRulesResponseDataObjectsInnerOneOf4) SetTuningExpressions(v []GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner)`

SetTuningExpressions sets TuningExpressions field to given value.

### HasTuningExpressions

`func (o *GetRulesResponseDataObjectsInnerOneOf4) HasTuningExpressions() bool`

HasTuningExpressions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


