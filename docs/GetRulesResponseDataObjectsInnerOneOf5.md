# GetRulesResponseDataObjectsInnerOneOf5

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
**DescriptionExpression** | **string** | The description to be used on the generated Signal | 
**NameExpression** | **string** |  | 
**FilterExpression** | **string** |  | 
**ValueExpression** | Pointer to **string** |  | [optional] 
**ValueFields** | Pointer to **[]string** | The value(s) to build an expression from - can be used instead of valueExpression | [optional] 
**GroupByFields** | **[]string** |  | 
**Score** | **int32** |  | 
**Version** | **int32** |  | 
**BaselineLastStarted** | Pointer to **time.Time** |  | [optional] 
**BaselineWindowSize** | **string** | milliseconds | 
**BaselineType** | **string** |  | 
**RetentionWindowSize** | **string** | milliseconds | 
**BaselineWindowSizeOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOf5BaselineWindowSizeOverride**](GetRulesResponseDataObjectsInnerOneOf5BaselineWindowSizeOverride.md) |  | [optional] 
**DescriptionExpressionOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOfNameOverride**](GetRulesResponseDataObjectsInnerOneOfNameOverride.md) |  | [optional] 
**GroupByFieldsOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOfTagsOverride**](GetRulesResponseDataObjectsInnerOneOfTagsOverride.md) |  | [optional] 
**NameExpressionOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOfNameOverride**](GetRulesResponseDataObjectsInnerOneOfNameOverride.md) |  | [optional] 
**RetentionWindowSizeOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOf5BaselineWindowSizeOverride**](GetRulesResponseDataObjectsInnerOneOf5BaselineWindowSizeOverride.md) |  | [optional] 
**ScoreOverride** | Pointer to [**GetRulesResponseDataObjectsInnerOneOf1ScoreOverride**](GetRulesResponseDataObjectsInnerOneOf1ScoreOverride.md) |  | [optional] 
**TuningExpressions** | Pointer to [**[]GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner**](GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner.md) |  | [optional] 

## Methods

### NewGetRulesResponseDataObjectsInnerOneOf5

`func NewGetRulesResponseDataObjectsInnerOneOf5(contentType string, created time.Time, createdBy string, deleted bool, enabled bool, entitySelectors []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner, id string, isPrototype bool, lastUpdated time.Time, lastUpdatedBy string, name string, status GetRulesResponseDataObjectsInnerOneOfStatus, ruleId int32, ruleSource string, ruleType string, signalCount07d int32, signalCount24h int32, summaryExpression string, tags []string, hasOverride bool, descriptionExpression string, nameExpression string, filterExpression string, groupByFields []string, score int32, version int32, baselineWindowSize string, baselineType string, retentionWindowSize string, ) *GetRulesResponseDataObjectsInnerOneOf5`

NewGetRulesResponseDataObjectsInnerOneOf5 instantiates a new GetRulesResponseDataObjectsInnerOneOf5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRulesResponseDataObjectsInnerOneOf5WithDefaults

`func NewGetRulesResponseDataObjectsInnerOneOf5WithDefaults() *GetRulesResponseDataObjectsInnerOneOf5`

NewGetRulesResponseDataObjectsInnerOneOf5WithDefaults instantiates a new GetRulesResponseDataObjectsInnerOneOf5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetRulesResponseDataObjectsInnerOneOf5) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetContentType

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetCreated

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCreatedBy

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDeleted

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetEnabled

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEntitySelectors

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetEntitySelectors() []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner`

GetEntitySelectors returns the EntitySelectors field if non-nil, zero value otherwise.

### GetEntitySelectorsOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetEntitySelectorsOk() (*[]GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner, bool)`

GetEntitySelectorsOk returns a tuple with the EntitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelectors

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetEntitySelectors(v []GetRulesResponseDataObjectsInnerOneOfEntitySelectorsInner)`

SetEntitySelectors sets EntitySelectors field to given value.


### GetId

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetId(v string)`

SetId sets Id field to given value.


### GetIsPrototype

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetIsPrototype() bool`

GetIsPrototype returns the IsPrototype field if non-nil, zero value otherwise.

### GetIsPrototypeOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetIsPrototypeOk() (*bool, bool)`

GetIsPrototypeOk returns a tuple with the IsPrototype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrototype

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetIsPrototype(v bool)`

SetIsPrototype sets IsPrototype field to given value.


### GetLastUpdated

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.


### GetName

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetName(v string)`

SetName sets Name field to given value.


### GetParentJaskId

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetParentJaskId() string`

GetParentJaskId returns the ParentJaskId field if non-nil, zero value otherwise.

### GetParentJaskIdOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetParentJaskIdOk() (*string, bool)`

GetParentJaskIdOk returns a tuple with the ParentJaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentJaskId

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetParentJaskId(v string)`

SetParentJaskId sets ParentJaskId field to given value.

### HasParentJaskId

`func (o *GetRulesResponseDataObjectsInnerOneOf5) HasParentJaskId() bool`

HasParentJaskId returns a boolean if a field has been set.

### GetStatus

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetStatus() GetRulesResponseDataObjectsInnerOneOfStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetStatusOk() (*GetRulesResponseDataObjectsInnerOneOfStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetStatus(v GetRulesResponseDataObjectsInnerOneOfStatus)`

SetStatus sets Status field to given value.


### GetRuleId

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetRuleId() int32`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetRuleIdOk() (*int32, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetRuleId(v int32)`

SetRuleId sets RuleId field to given value.


### GetRuleSource

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetRuleSource() string`

GetRuleSource returns the RuleSource field if non-nil, zero value otherwise.

### GetRuleSourceOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetRuleSourceOk() (*string, bool)`

GetRuleSourceOk returns a tuple with the RuleSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSource

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetRuleSource(v string)`

SetRuleSource sets RuleSource field to given value.


### GetRuleType

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.


### GetSignalCount07d

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetSignalCount07d() int32`

GetSignalCount07d returns the SignalCount07d field if non-nil, zero value otherwise.

### GetSignalCount07dOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetSignalCount07dOk() (*int32, bool)`

GetSignalCount07dOk returns a tuple with the SignalCount07d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalCount07d

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetSignalCount07d(v int32)`

SetSignalCount07d sets SignalCount07d field to given value.


### GetSignalCount24h

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetSignalCount24h() int32`

GetSignalCount24h returns the SignalCount24h field if non-nil, zero value otherwise.

### GetSignalCount24hOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetSignalCount24hOk() (*int32, bool)`

GetSignalCount24hOk returns a tuple with the SignalCount24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalCount24h

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetSignalCount24h(v int32)`

SetSignalCount24h sets SignalCount24h field to given value.


### GetSummaryExpression

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetSummaryExpression() string`

GetSummaryExpression returns the SummaryExpression field if non-nil, zero value otherwise.

### GetSummaryExpressionOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetSummaryExpressionOk() (*string, bool)`

GetSummaryExpressionOk returns a tuple with the SummaryExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryExpression

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetSummaryExpression(v string)`

SetSummaryExpression sets SummaryExpression field to given value.


### GetTags

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetHasOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetHasOverride() bool`

GetHasOverride returns the HasOverride field if non-nil, zero value otherwise.

### GetHasOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetHasOverrideOk() (*bool, bool)`

GetHasOverrideOk returns a tuple with the HasOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetHasOverride(v bool)`

SetHasOverride sets HasOverride field to given value.


### GetTagsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetTagsOverride() GetRulesResponseDataObjectsInnerOneOfTagsOverride`

GetTagsOverride returns the TagsOverride field if non-nil, zero value otherwise.

### GetTagsOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetTagsOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfTagsOverride, bool)`

GetTagsOverrideOk returns a tuple with the TagsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetTagsOverride(v GetRulesResponseDataObjectsInnerOneOfTagsOverride)`

SetTagsOverride sets TagsOverride field to given value.

### HasTagsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) HasTagsOverride() bool`

HasTagsOverride returns a boolean if a field has been set.

### GetIsPrototypeOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetIsPrototypeOverride() GetRulesResponseDataObjectsInnerOneOfIsPrototypeOverride`

GetIsPrototypeOverride returns the IsPrototypeOverride field if non-nil, zero value otherwise.

### GetIsPrototypeOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetIsPrototypeOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfIsPrototypeOverride, bool)`

GetIsPrototypeOverrideOk returns a tuple with the IsPrototypeOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrototypeOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetIsPrototypeOverride(v GetRulesResponseDataObjectsInnerOneOfIsPrototypeOverride)`

SetIsPrototypeOverride sets IsPrototypeOverride field to given value.

### HasIsPrototypeOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) HasIsPrototypeOverride() bool`

HasIsPrototypeOverride returns a boolean if a field has been set.

### GetEntitySelectorsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetEntitySelectorsOverride() GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride`

GetEntitySelectorsOverride returns the EntitySelectorsOverride field if non-nil, zero value otherwise.

### GetEntitySelectorsOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetEntitySelectorsOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride, bool)`

GetEntitySelectorsOverrideOk returns a tuple with the EntitySelectorsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelectorsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetEntitySelectorsOverride(v GetRulesResponseDataObjectsInnerOneOfEntitySelectorsOverride)`

SetEntitySelectorsOverride sets EntitySelectorsOverride field to given value.

### HasEntitySelectorsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) HasEntitySelectorsOverride() bool`

HasEntitySelectorsOverride returns a boolean if a field has been set.

### GetNameOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetNameOverride() GetRulesResponseDataObjectsInnerOneOfNameOverride`

GetNameOverride returns the NameOverride field if non-nil, zero value otherwise.

### GetNameOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetNameOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfNameOverride, bool)`

GetNameOverrideOk returns a tuple with the NameOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetNameOverride(v GetRulesResponseDataObjectsInnerOneOfNameOverride)`

SetNameOverride sets NameOverride field to given value.

### HasNameOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) HasNameOverride() bool`

HasNameOverride returns a boolean if a field has been set.

### GetSummaryExpressionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetSummaryExpressionOverride() GetRulesResponseDataObjectsInnerOneOfNameOverride`

GetSummaryExpressionOverride returns the SummaryExpressionOverride field if non-nil, zero value otherwise.

### GetSummaryExpressionOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetSummaryExpressionOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfNameOverride, bool)`

GetSummaryExpressionOverrideOk returns a tuple with the SummaryExpressionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryExpressionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetSummaryExpressionOverride(v GetRulesResponseDataObjectsInnerOneOfNameOverride)`

SetSummaryExpressionOverride sets SummaryExpressionOverride field to given value.

### HasSummaryExpressionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) HasSummaryExpressionOverride() bool`

HasSummaryExpressionOverride returns a boolean if a field has been set.

### GetDescriptionExpression

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetDescriptionExpression() string`

GetDescriptionExpression returns the DescriptionExpression field if non-nil, zero value otherwise.

### GetDescriptionExpressionOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetDescriptionExpressionOk() (*string, bool)`

GetDescriptionExpressionOk returns a tuple with the DescriptionExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionExpression

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetDescriptionExpression(v string)`

SetDescriptionExpression sets DescriptionExpression field to given value.


### GetNameExpression

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetNameExpression() string`

GetNameExpression returns the NameExpression field if non-nil, zero value otherwise.

### GetNameExpressionOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetNameExpressionOk() (*string, bool)`

GetNameExpressionOk returns a tuple with the NameExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameExpression

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetNameExpression(v string)`

SetNameExpression sets NameExpression field to given value.


### GetFilterExpression

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetFilterExpression() string`

GetFilterExpression returns the FilterExpression field if non-nil, zero value otherwise.

### GetFilterExpressionOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetFilterExpressionOk() (*string, bool)`

GetFilterExpressionOk returns a tuple with the FilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterExpression

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetFilterExpression(v string)`

SetFilterExpression sets FilterExpression field to given value.


### GetValueExpression

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetValueExpression() string`

GetValueExpression returns the ValueExpression field if non-nil, zero value otherwise.

### GetValueExpressionOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetValueExpressionOk() (*string, bool)`

GetValueExpressionOk returns a tuple with the ValueExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueExpression

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetValueExpression(v string)`

SetValueExpression sets ValueExpression field to given value.

### HasValueExpression

`func (o *GetRulesResponseDataObjectsInnerOneOf5) HasValueExpression() bool`

HasValueExpression returns a boolean if a field has been set.

### GetValueFields

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetValueFields() []string`

GetValueFields returns the ValueFields field if non-nil, zero value otherwise.

### GetValueFieldsOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetValueFieldsOk() (*[]string, bool)`

GetValueFieldsOk returns a tuple with the ValueFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFields

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetValueFields(v []string)`

SetValueFields sets ValueFields field to given value.

### HasValueFields

`func (o *GetRulesResponseDataObjectsInnerOneOf5) HasValueFields() bool`

HasValueFields returns a boolean if a field has been set.

### GetGroupByFields

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetGroupByFields() []string`

GetGroupByFields returns the GroupByFields field if non-nil, zero value otherwise.

### GetGroupByFieldsOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetGroupByFieldsOk() (*[]string, bool)`

GetGroupByFieldsOk returns a tuple with the GroupByFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFields

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetGroupByFields(v []string)`

SetGroupByFields sets GroupByFields field to given value.


### GetScore

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetScore(v int32)`

SetScore sets Score field to given value.


### GetVersion

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetBaselineLastStarted

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetBaselineLastStarted() time.Time`

GetBaselineLastStarted returns the BaselineLastStarted field if non-nil, zero value otherwise.

### GetBaselineLastStartedOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetBaselineLastStartedOk() (*time.Time, bool)`

GetBaselineLastStartedOk returns a tuple with the BaselineLastStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineLastStarted

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetBaselineLastStarted(v time.Time)`

SetBaselineLastStarted sets BaselineLastStarted field to given value.

### HasBaselineLastStarted

`func (o *GetRulesResponseDataObjectsInnerOneOf5) HasBaselineLastStarted() bool`

HasBaselineLastStarted returns a boolean if a field has been set.

### GetBaselineWindowSize

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetBaselineWindowSize() string`

GetBaselineWindowSize returns the BaselineWindowSize field if non-nil, zero value otherwise.

### GetBaselineWindowSizeOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetBaselineWindowSizeOk() (*string, bool)`

GetBaselineWindowSizeOk returns a tuple with the BaselineWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineWindowSize

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetBaselineWindowSize(v string)`

SetBaselineWindowSize sets BaselineWindowSize field to given value.


### GetBaselineType

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetBaselineType() string`

GetBaselineType returns the BaselineType field if non-nil, zero value otherwise.

### GetBaselineTypeOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetBaselineTypeOk() (*string, bool)`

GetBaselineTypeOk returns a tuple with the BaselineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineType

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetBaselineType(v string)`

SetBaselineType sets BaselineType field to given value.


### GetRetentionWindowSize

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetRetentionWindowSize() string`

GetRetentionWindowSize returns the RetentionWindowSize field if non-nil, zero value otherwise.

### GetRetentionWindowSizeOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetRetentionWindowSizeOk() (*string, bool)`

GetRetentionWindowSizeOk returns a tuple with the RetentionWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionWindowSize

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetRetentionWindowSize(v string)`

SetRetentionWindowSize sets RetentionWindowSize field to given value.


### GetBaselineWindowSizeOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetBaselineWindowSizeOverride() GetRulesResponseDataObjectsInnerOneOf5BaselineWindowSizeOverride`

GetBaselineWindowSizeOverride returns the BaselineWindowSizeOverride field if non-nil, zero value otherwise.

### GetBaselineWindowSizeOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetBaselineWindowSizeOverrideOk() (*GetRulesResponseDataObjectsInnerOneOf5BaselineWindowSizeOverride, bool)`

GetBaselineWindowSizeOverrideOk returns a tuple with the BaselineWindowSizeOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineWindowSizeOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetBaselineWindowSizeOverride(v GetRulesResponseDataObjectsInnerOneOf5BaselineWindowSizeOverride)`

SetBaselineWindowSizeOverride sets BaselineWindowSizeOverride field to given value.

### HasBaselineWindowSizeOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) HasBaselineWindowSizeOverride() bool`

HasBaselineWindowSizeOverride returns a boolean if a field has been set.

### GetDescriptionExpressionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetDescriptionExpressionOverride() GetRulesResponseDataObjectsInnerOneOfNameOverride`

GetDescriptionExpressionOverride returns the DescriptionExpressionOverride field if non-nil, zero value otherwise.

### GetDescriptionExpressionOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetDescriptionExpressionOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfNameOverride, bool)`

GetDescriptionExpressionOverrideOk returns a tuple with the DescriptionExpressionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionExpressionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetDescriptionExpressionOverride(v GetRulesResponseDataObjectsInnerOneOfNameOverride)`

SetDescriptionExpressionOverride sets DescriptionExpressionOverride field to given value.

### HasDescriptionExpressionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) HasDescriptionExpressionOverride() bool`

HasDescriptionExpressionOverride returns a boolean if a field has been set.

### GetGroupByFieldsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetGroupByFieldsOverride() GetRulesResponseDataObjectsInnerOneOfTagsOverride`

GetGroupByFieldsOverride returns the GroupByFieldsOverride field if non-nil, zero value otherwise.

### GetGroupByFieldsOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetGroupByFieldsOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfTagsOverride, bool)`

GetGroupByFieldsOverrideOk returns a tuple with the GroupByFieldsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByFieldsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetGroupByFieldsOverride(v GetRulesResponseDataObjectsInnerOneOfTagsOverride)`

SetGroupByFieldsOverride sets GroupByFieldsOverride field to given value.

### HasGroupByFieldsOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) HasGroupByFieldsOverride() bool`

HasGroupByFieldsOverride returns a boolean if a field has been set.

### GetNameExpressionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetNameExpressionOverride() GetRulesResponseDataObjectsInnerOneOfNameOverride`

GetNameExpressionOverride returns the NameExpressionOverride field if non-nil, zero value otherwise.

### GetNameExpressionOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetNameExpressionOverrideOk() (*GetRulesResponseDataObjectsInnerOneOfNameOverride, bool)`

GetNameExpressionOverrideOk returns a tuple with the NameExpressionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameExpressionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetNameExpressionOverride(v GetRulesResponseDataObjectsInnerOneOfNameOverride)`

SetNameExpressionOverride sets NameExpressionOverride field to given value.

### HasNameExpressionOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) HasNameExpressionOverride() bool`

HasNameExpressionOverride returns a boolean if a field has been set.

### GetRetentionWindowSizeOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetRetentionWindowSizeOverride() GetRulesResponseDataObjectsInnerOneOf5BaselineWindowSizeOverride`

GetRetentionWindowSizeOverride returns the RetentionWindowSizeOverride field if non-nil, zero value otherwise.

### GetRetentionWindowSizeOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetRetentionWindowSizeOverrideOk() (*GetRulesResponseDataObjectsInnerOneOf5BaselineWindowSizeOverride, bool)`

GetRetentionWindowSizeOverrideOk returns a tuple with the RetentionWindowSizeOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionWindowSizeOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetRetentionWindowSizeOverride(v GetRulesResponseDataObjectsInnerOneOf5BaselineWindowSizeOverride)`

SetRetentionWindowSizeOverride sets RetentionWindowSizeOverride field to given value.

### HasRetentionWindowSizeOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) HasRetentionWindowSizeOverride() bool`

HasRetentionWindowSizeOverride returns a boolean if a field has been set.

### GetScoreOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetScoreOverride() GetRulesResponseDataObjectsInnerOneOf1ScoreOverride`

GetScoreOverride returns the ScoreOverride field if non-nil, zero value otherwise.

### GetScoreOverrideOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetScoreOverrideOk() (*GetRulesResponseDataObjectsInnerOneOf1ScoreOverride, bool)`

GetScoreOverrideOk returns a tuple with the ScoreOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetScoreOverride(v GetRulesResponseDataObjectsInnerOneOf1ScoreOverride)`

SetScoreOverride sets ScoreOverride field to given value.

### HasScoreOverride

`func (o *GetRulesResponseDataObjectsInnerOneOf5) HasScoreOverride() bool`

HasScoreOverride returns a boolean if a field has been set.

### GetTuningExpressions

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetTuningExpressions() []GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner`

GetTuningExpressions returns the TuningExpressions field if non-nil, zero value otherwise.

### GetTuningExpressionsOk

`func (o *GetRulesResponseDataObjectsInnerOneOf5) GetTuningExpressionsOk() (*[]GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner, bool)`

GetTuningExpressionsOk returns a tuple with the TuningExpressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuningExpressions

`func (o *GetRulesResponseDataObjectsInnerOneOf5) SetTuningExpressions(v []GetRulesResponseDataObjectsInnerOneOfTuningExpressionsInner)`

SetTuningExpressions sets TuningExpressions field to given value.

### HasTuningExpressions

`func (o *GetRulesResponseDataObjectsInnerOneOf5) HasTuningExpressions() bool`

HasTuningExpressions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


