# CreateAlmanaxWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BonusWhitelist** | Pointer to **[]string** | from all available bonuses (ids) from /dofus2/meta/{language}/almanax/bonuses | [optional] 
**BonusBlacklist** | Pointer to **[]string** | from all available bonuses (ids) from /dofus2/meta/{language}/almanax/bonuses | [optional] 
**Subscriptions** | **[]string** | Get the available subscriptions with /meta/webhooks/almanax | 
**Format** | **string** |  | 
**Callback** | **string** | Discord Webhook URL | 
**DailySettings** | Pointer to [**NullableCreateAlmanaxWebhookDailySettings**](CreateAlmanaxWebhookDailySettings.md) |  | [optional] 
**IsoDate** | Pointer to **NullableBool** | If false, it will use common local time formats and weekday translations. If true, the format is YYYY-MM-DD. | [optional] [default to false]
**Mentions** | Pointer to [**map[string][]CreateAlmanaxWebhookMentionsValueInner**](array.md) | Almanax bonus ids mapped to array of mentions. | [optional] 
**Intervals** | **[]string** | - Daily posts each day, filtering with Black/Whitelist and mentions are applied daily. - Weekly posts the next 7 days (excluding the posting day) once per week at the specified time. With only weekly selected, of all mentions, only prior notices will come through daily. The 7 day preview gets filtered by the Black/Whitelist. - Monthly posts a preview of the next month from first to last date. The post will be on the last day of a month (ignoring day of the week) at the specified time. Mentions and filtering works like weekly. The biggest difference between daily and the other two is that daily always posts the current day while monthly and weekly only show future days. You can always combine the intervals by selecting multiple intervals for one hook or create multiple hooks for the same channel with different settings to get every highly specific combination you want. | 
**WeeklyWeekday** | Pointer to **NullableString** | When to post the weekly preview at the specified time. | [optional] 

## Methods

### NewCreateAlmanaxWebhook

`func NewCreateAlmanaxWebhook(subscriptions []string, format string, callback string, intervals []string, ) *CreateAlmanaxWebhook`

NewCreateAlmanaxWebhook instantiates a new CreateAlmanaxWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAlmanaxWebhookWithDefaults

`func NewCreateAlmanaxWebhookWithDefaults() *CreateAlmanaxWebhook`

NewCreateAlmanaxWebhookWithDefaults instantiates a new CreateAlmanaxWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBonusWhitelist

`func (o *CreateAlmanaxWebhook) GetBonusWhitelist() []string`

GetBonusWhitelist returns the BonusWhitelist field if non-nil, zero value otherwise.

### GetBonusWhitelistOk

`func (o *CreateAlmanaxWebhook) GetBonusWhitelistOk() (*[]string, bool)`

GetBonusWhitelistOk returns a tuple with the BonusWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonusWhitelist

`func (o *CreateAlmanaxWebhook) SetBonusWhitelist(v []string)`

SetBonusWhitelist sets BonusWhitelist field to given value.

### HasBonusWhitelist

`func (o *CreateAlmanaxWebhook) HasBonusWhitelist() bool`

HasBonusWhitelist returns a boolean if a field has been set.

### SetBonusWhitelistNil

`func (o *CreateAlmanaxWebhook) SetBonusWhitelistNil(b bool)`

 SetBonusWhitelistNil sets the value for BonusWhitelist to be an explicit nil

### UnsetBonusWhitelist
`func (o *CreateAlmanaxWebhook) UnsetBonusWhitelist()`

UnsetBonusWhitelist ensures that no value is present for BonusWhitelist, not even an explicit nil
### GetBonusBlacklist

`func (o *CreateAlmanaxWebhook) GetBonusBlacklist() []string`

GetBonusBlacklist returns the BonusBlacklist field if non-nil, zero value otherwise.

### GetBonusBlacklistOk

`func (o *CreateAlmanaxWebhook) GetBonusBlacklistOk() (*[]string, bool)`

GetBonusBlacklistOk returns a tuple with the BonusBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonusBlacklist

`func (o *CreateAlmanaxWebhook) SetBonusBlacklist(v []string)`

SetBonusBlacklist sets BonusBlacklist field to given value.

### HasBonusBlacklist

`func (o *CreateAlmanaxWebhook) HasBonusBlacklist() bool`

HasBonusBlacklist returns a boolean if a field has been set.

### SetBonusBlacklistNil

`func (o *CreateAlmanaxWebhook) SetBonusBlacklistNil(b bool)`

 SetBonusBlacklistNil sets the value for BonusBlacklist to be an explicit nil

### UnsetBonusBlacklist
`func (o *CreateAlmanaxWebhook) UnsetBonusBlacklist()`

UnsetBonusBlacklist ensures that no value is present for BonusBlacklist, not even an explicit nil
### GetSubscriptions

`func (o *CreateAlmanaxWebhook) GetSubscriptions() []string`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *CreateAlmanaxWebhook) GetSubscriptionsOk() (*[]string, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *CreateAlmanaxWebhook) SetSubscriptions(v []string)`

SetSubscriptions sets Subscriptions field to given value.


### GetFormat

`func (o *CreateAlmanaxWebhook) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateAlmanaxWebhook) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateAlmanaxWebhook) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetCallback

`func (o *CreateAlmanaxWebhook) GetCallback() string`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *CreateAlmanaxWebhook) GetCallbackOk() (*string, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *CreateAlmanaxWebhook) SetCallback(v string)`

SetCallback sets Callback field to given value.


### GetDailySettings

`func (o *CreateAlmanaxWebhook) GetDailySettings() CreateAlmanaxWebhookDailySettings`

GetDailySettings returns the DailySettings field if non-nil, zero value otherwise.

### GetDailySettingsOk

`func (o *CreateAlmanaxWebhook) GetDailySettingsOk() (*CreateAlmanaxWebhookDailySettings, bool)`

GetDailySettingsOk returns a tuple with the DailySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySettings

`func (o *CreateAlmanaxWebhook) SetDailySettings(v CreateAlmanaxWebhookDailySettings)`

SetDailySettings sets DailySettings field to given value.

### HasDailySettings

`func (o *CreateAlmanaxWebhook) HasDailySettings() bool`

HasDailySettings returns a boolean if a field has been set.

### SetDailySettingsNil

`func (o *CreateAlmanaxWebhook) SetDailySettingsNil(b bool)`

 SetDailySettingsNil sets the value for DailySettings to be an explicit nil

### UnsetDailySettings
`func (o *CreateAlmanaxWebhook) UnsetDailySettings()`

UnsetDailySettings ensures that no value is present for DailySettings, not even an explicit nil
### GetIsoDate

`func (o *CreateAlmanaxWebhook) GetIsoDate() bool`

GetIsoDate returns the IsoDate field if non-nil, zero value otherwise.

### GetIsoDateOk

`func (o *CreateAlmanaxWebhook) GetIsoDateOk() (*bool, bool)`

GetIsoDateOk returns a tuple with the IsoDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoDate

`func (o *CreateAlmanaxWebhook) SetIsoDate(v bool)`

SetIsoDate sets IsoDate field to given value.

### HasIsoDate

`func (o *CreateAlmanaxWebhook) HasIsoDate() bool`

HasIsoDate returns a boolean if a field has been set.

### SetIsoDateNil

`func (o *CreateAlmanaxWebhook) SetIsoDateNil(b bool)`

 SetIsoDateNil sets the value for IsoDate to be an explicit nil

### UnsetIsoDate
`func (o *CreateAlmanaxWebhook) UnsetIsoDate()`

UnsetIsoDate ensures that no value is present for IsoDate, not even an explicit nil
### GetMentions

`func (o *CreateAlmanaxWebhook) GetMentions() map[string][]CreateAlmanaxWebhookMentionsValueInner`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *CreateAlmanaxWebhook) GetMentionsOk() (*map[string][]CreateAlmanaxWebhookMentionsValueInner, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *CreateAlmanaxWebhook) SetMentions(v map[string][]CreateAlmanaxWebhookMentionsValueInner)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *CreateAlmanaxWebhook) HasMentions() bool`

HasMentions returns a boolean if a field has been set.

### SetMentionsNil

`func (o *CreateAlmanaxWebhook) SetMentionsNil(b bool)`

 SetMentionsNil sets the value for Mentions to be an explicit nil

### UnsetMentions
`func (o *CreateAlmanaxWebhook) UnsetMentions()`

UnsetMentions ensures that no value is present for Mentions, not even an explicit nil
### GetIntervals

`func (o *CreateAlmanaxWebhook) GetIntervals() []string`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *CreateAlmanaxWebhook) GetIntervalsOk() (*[]string, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *CreateAlmanaxWebhook) SetIntervals(v []string)`

SetIntervals sets Intervals field to given value.


### GetWeeklyWeekday

`func (o *CreateAlmanaxWebhook) GetWeeklyWeekday() string`

GetWeeklyWeekday returns the WeeklyWeekday field if non-nil, zero value otherwise.

### GetWeeklyWeekdayOk

`func (o *CreateAlmanaxWebhook) GetWeeklyWeekdayOk() (*string, bool)`

GetWeeklyWeekdayOk returns a tuple with the WeeklyWeekday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklyWeekday

`func (o *CreateAlmanaxWebhook) SetWeeklyWeekday(v string)`

SetWeeklyWeekday sets WeeklyWeekday field to given value.

### HasWeeklyWeekday

`func (o *CreateAlmanaxWebhook) HasWeeklyWeekday() bool`

HasWeeklyWeekday returns a boolean if a field has been set.

### SetWeeklyWeekdayNil

`func (o *CreateAlmanaxWebhook) SetWeeklyWeekdayNil(b bool)`

 SetWeeklyWeekdayNil sets the value for WeeklyWeekday to be an explicit nil

### UnsetWeeklyWeekday
`func (o *CreateAlmanaxWebhook) UnsetWeeklyWeekday()`

UnsetWeeklyWeekday ensures that no value is present for WeeklyWeekday, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


