# AlmanaxWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DailySettings** | Pointer to [**AlmanaxWebhookDailySettings**](AlmanaxWebhookDailySettings.md) |  | [optional] 
**BonusWhitelist** | Pointer to **[]string** | Only post when these bonuses come up. From all available bonuses (ids) from /dofus3/meta/{language}/almanax/bonuses. | [optional] 
**BonusBlacklist** | Pointer to **[]string** | Skip the day when these bonuses come up. From all available bonuses (ids) from /dofus3/meta/{language}/almanax/bonuses | [optional] 
**Subscriptions** | Pointer to **[]string** | Get the available subscriptions with /meta/webhooks/almanax | [optional] 
**IsoDate** | Pointer to **bool** | If false, it will use common local time formats and weekday translations. If true, the format is YYYY-MM-DD. | [optional] [default to false]
**Mentions** | Pointer to [**map[string][]CreateAlmanaxWebhookMentionsValueInner**](array.md) | Almanax bonus ids mapped to array of mentions. | [optional] 
**Intervals** | Pointer to **[]string** | - Daily posts each day, filtering with Black/Whitelist and mentions are applied daily. - Weekly posts the next 7 days (excluding the posting day) once per week at the specified time. With only weekly selected, of all mentions, only prior notices will come through daily. The 7 day preview gets filtered by the Black/Whitelist. - Monthly posts a preview of the next month from first to last date. The post will be on the last day of a month (ignoring day of the week) at the specified time. Mentions and filtering works like weekly. The biggest difference between daily and the other two is that daily always posts the current day while monthly and weekly only show future days. You can always combine the intervals by selecting multiple intervals for one hook or create multiple hooks for the same channel with different settings to get every highly specific combination you want. | [optional] 
**WeeklyWeekday** | Pointer to **NullableString** | When to post the weekly preview at the specified time. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastFiredAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAlmanaxWebhook

`func NewAlmanaxWebhook() *AlmanaxWebhook`

NewAlmanaxWebhook instantiates a new AlmanaxWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlmanaxWebhookWithDefaults

`func NewAlmanaxWebhookWithDefaults() *AlmanaxWebhook`

NewAlmanaxWebhookWithDefaults instantiates a new AlmanaxWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlmanaxWebhook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlmanaxWebhook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlmanaxWebhook) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AlmanaxWebhook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDailySettings

`func (o *AlmanaxWebhook) GetDailySettings() AlmanaxWebhookDailySettings`

GetDailySettings returns the DailySettings field if non-nil, zero value otherwise.

### GetDailySettingsOk

`func (o *AlmanaxWebhook) GetDailySettingsOk() (*AlmanaxWebhookDailySettings, bool)`

GetDailySettingsOk returns a tuple with the DailySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySettings

`func (o *AlmanaxWebhook) SetDailySettings(v AlmanaxWebhookDailySettings)`

SetDailySettings sets DailySettings field to given value.

### HasDailySettings

`func (o *AlmanaxWebhook) HasDailySettings() bool`

HasDailySettings returns a boolean if a field has been set.

### GetBonusWhitelist

`func (o *AlmanaxWebhook) GetBonusWhitelist() []string`

GetBonusWhitelist returns the BonusWhitelist field if non-nil, zero value otherwise.

### GetBonusWhitelistOk

`func (o *AlmanaxWebhook) GetBonusWhitelistOk() (*[]string, bool)`

GetBonusWhitelistOk returns a tuple with the BonusWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonusWhitelist

`func (o *AlmanaxWebhook) SetBonusWhitelist(v []string)`

SetBonusWhitelist sets BonusWhitelist field to given value.

### HasBonusWhitelist

`func (o *AlmanaxWebhook) HasBonusWhitelist() bool`

HasBonusWhitelist returns a boolean if a field has been set.

### SetBonusWhitelistNil

`func (o *AlmanaxWebhook) SetBonusWhitelistNil(b bool)`

 SetBonusWhitelistNil sets the value for BonusWhitelist to be an explicit nil

### UnsetBonusWhitelist
`func (o *AlmanaxWebhook) UnsetBonusWhitelist()`

UnsetBonusWhitelist ensures that no value is present for BonusWhitelist, not even an explicit nil
### GetBonusBlacklist

`func (o *AlmanaxWebhook) GetBonusBlacklist() []string`

GetBonusBlacklist returns the BonusBlacklist field if non-nil, zero value otherwise.

### GetBonusBlacklistOk

`func (o *AlmanaxWebhook) GetBonusBlacklistOk() (*[]string, bool)`

GetBonusBlacklistOk returns a tuple with the BonusBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonusBlacklist

`func (o *AlmanaxWebhook) SetBonusBlacklist(v []string)`

SetBonusBlacklist sets BonusBlacklist field to given value.

### HasBonusBlacklist

`func (o *AlmanaxWebhook) HasBonusBlacklist() bool`

HasBonusBlacklist returns a boolean if a field has been set.

### SetBonusBlacklistNil

`func (o *AlmanaxWebhook) SetBonusBlacklistNil(b bool)`

 SetBonusBlacklistNil sets the value for BonusBlacklist to be an explicit nil

### UnsetBonusBlacklist
`func (o *AlmanaxWebhook) UnsetBonusBlacklist()`

UnsetBonusBlacklist ensures that no value is present for BonusBlacklist, not even an explicit nil
### GetSubscriptions

`func (o *AlmanaxWebhook) GetSubscriptions() []string`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *AlmanaxWebhook) GetSubscriptionsOk() (*[]string, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *AlmanaxWebhook) SetSubscriptions(v []string)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *AlmanaxWebhook) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetIsoDate

`func (o *AlmanaxWebhook) GetIsoDate() bool`

GetIsoDate returns the IsoDate field if non-nil, zero value otherwise.

### GetIsoDateOk

`func (o *AlmanaxWebhook) GetIsoDateOk() (*bool, bool)`

GetIsoDateOk returns a tuple with the IsoDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoDate

`func (o *AlmanaxWebhook) SetIsoDate(v bool)`

SetIsoDate sets IsoDate field to given value.

### HasIsoDate

`func (o *AlmanaxWebhook) HasIsoDate() bool`

HasIsoDate returns a boolean if a field has been set.

### GetMentions

`func (o *AlmanaxWebhook) GetMentions() map[string][]CreateAlmanaxWebhookMentionsValueInner`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *AlmanaxWebhook) GetMentionsOk() (*map[string][]CreateAlmanaxWebhookMentionsValueInner, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *AlmanaxWebhook) SetMentions(v map[string][]CreateAlmanaxWebhookMentionsValueInner)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *AlmanaxWebhook) HasMentions() bool`

HasMentions returns a boolean if a field has been set.

### SetMentionsNil

`func (o *AlmanaxWebhook) SetMentionsNil(b bool)`

 SetMentionsNil sets the value for Mentions to be an explicit nil

### UnsetMentions
`func (o *AlmanaxWebhook) UnsetMentions()`

UnsetMentions ensures that no value is present for Mentions, not even an explicit nil
### GetIntervals

`func (o *AlmanaxWebhook) GetIntervals() []string`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *AlmanaxWebhook) GetIntervalsOk() (*[]string, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *AlmanaxWebhook) SetIntervals(v []string)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *AlmanaxWebhook) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.

### GetWeeklyWeekday

`func (o *AlmanaxWebhook) GetWeeklyWeekday() string`

GetWeeklyWeekday returns the WeeklyWeekday field if non-nil, zero value otherwise.

### GetWeeklyWeekdayOk

`func (o *AlmanaxWebhook) GetWeeklyWeekdayOk() (*string, bool)`

GetWeeklyWeekdayOk returns a tuple with the WeeklyWeekday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklyWeekday

`func (o *AlmanaxWebhook) SetWeeklyWeekday(v string)`

SetWeeklyWeekday sets WeeklyWeekday field to given value.

### HasWeeklyWeekday

`func (o *AlmanaxWebhook) HasWeeklyWeekday() bool`

HasWeeklyWeekday returns a boolean if a field has been set.

### SetWeeklyWeekdayNil

`func (o *AlmanaxWebhook) SetWeeklyWeekdayNil(b bool)`

 SetWeeklyWeekdayNil sets the value for WeeklyWeekday to be an explicit nil

### UnsetWeeklyWeekday
`func (o *AlmanaxWebhook) UnsetWeeklyWeekday()`

UnsetWeeklyWeekday ensures that no value is present for WeeklyWeekday, not even an explicit nil
### GetCreatedAt

`func (o *AlmanaxWebhook) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AlmanaxWebhook) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AlmanaxWebhook) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AlmanaxWebhook) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastFiredAt

`func (o *AlmanaxWebhook) GetLastFiredAt() time.Time`

GetLastFiredAt returns the LastFiredAt field if non-nil, zero value otherwise.

### GetLastFiredAtOk

`func (o *AlmanaxWebhook) GetLastFiredAtOk() (*time.Time, bool)`

GetLastFiredAtOk returns a tuple with the LastFiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFiredAt

`func (o *AlmanaxWebhook) SetLastFiredAt(v time.Time)`

SetLastFiredAt sets LastFiredAt field to given value.

### HasLastFiredAt

`func (o *AlmanaxWebhook) HasLastFiredAt() bool`

HasLastFiredAt returns a boolean if a field has been set.

### SetLastFiredAtNil

`func (o *AlmanaxWebhook) SetLastFiredAtNil(b bool)`

 SetLastFiredAtNil sets the value for LastFiredAt to be an explicit nil

### UnsetLastFiredAt
`func (o *AlmanaxWebhook) UnsetLastFiredAt()`

UnsetLastFiredAt ensures that no value is present for LastFiredAt, not even an explicit nil
### GetUpdatedAt

`func (o *AlmanaxWebhook) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AlmanaxWebhook) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AlmanaxWebhook) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AlmanaxWebhook) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


