# PutAlmanaxWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BonusWhitelist** | Pointer to **[]string** | from all available bonuses (ids) from /dofus3/meta/{language}/almanax/bonuses. Delete old entries with empty array []. Just null changes nothing. | [optional] 
**BonusBlacklist** | Pointer to **[]string** | from all available bonuses (ids) from /dofus3/meta/{language}/almanax/bonuses. Delete old entries with empty array []. Just null changes nothing. | [optional] 
**Subscriptions** | Pointer to **[]string** | Get the available subscriptions with /meta/webhooks/almanax | [optional] 
**DailySettings** | Pointer to [**NullableCreateAlmanaxWebhookDailySettings**](CreateAlmanaxWebhookDailySettings.md) |  | [optional] 
**IsoDate** | Pointer to **NullableBool** | If false, it will use common local time formats and weekday translations. If true, the format is YYYY-MM-DD. | [optional] [default to false]
**Mentions** | Pointer to [**map[string][]CreateAlmanaxWebhookMentionsValueInner**](array.md) | Almanax bonus ids mapped to array of mentions. | [optional] 
**Intervals** | Pointer to **[]string** |  | [optional] 
**WeeklyWeekday** | Pointer to **NullableString** | When to post the weekly preview at the specified time. | [optional] 

## Methods

### NewPutAlmanaxWebhook

`func NewPutAlmanaxWebhook() *PutAlmanaxWebhook`

NewPutAlmanaxWebhook instantiates a new PutAlmanaxWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutAlmanaxWebhookWithDefaults

`func NewPutAlmanaxWebhookWithDefaults() *PutAlmanaxWebhook`

NewPutAlmanaxWebhookWithDefaults instantiates a new PutAlmanaxWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBonusWhitelist

`func (o *PutAlmanaxWebhook) GetBonusWhitelist() []string`

GetBonusWhitelist returns the BonusWhitelist field if non-nil, zero value otherwise.

### GetBonusWhitelistOk

`func (o *PutAlmanaxWebhook) GetBonusWhitelistOk() (*[]string, bool)`

GetBonusWhitelistOk returns a tuple with the BonusWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonusWhitelist

`func (o *PutAlmanaxWebhook) SetBonusWhitelist(v []string)`

SetBonusWhitelist sets BonusWhitelist field to given value.

### HasBonusWhitelist

`func (o *PutAlmanaxWebhook) HasBonusWhitelist() bool`

HasBonusWhitelist returns a boolean if a field has been set.

### SetBonusWhitelistNil

`func (o *PutAlmanaxWebhook) SetBonusWhitelistNil(b bool)`

 SetBonusWhitelistNil sets the value for BonusWhitelist to be an explicit nil

### UnsetBonusWhitelist
`func (o *PutAlmanaxWebhook) UnsetBonusWhitelist()`

UnsetBonusWhitelist ensures that no value is present for BonusWhitelist, not even an explicit nil
### GetBonusBlacklist

`func (o *PutAlmanaxWebhook) GetBonusBlacklist() []string`

GetBonusBlacklist returns the BonusBlacklist field if non-nil, zero value otherwise.

### GetBonusBlacklistOk

`func (o *PutAlmanaxWebhook) GetBonusBlacklistOk() (*[]string, bool)`

GetBonusBlacklistOk returns a tuple with the BonusBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonusBlacklist

`func (o *PutAlmanaxWebhook) SetBonusBlacklist(v []string)`

SetBonusBlacklist sets BonusBlacklist field to given value.

### HasBonusBlacklist

`func (o *PutAlmanaxWebhook) HasBonusBlacklist() bool`

HasBonusBlacklist returns a boolean if a field has been set.

### SetBonusBlacklistNil

`func (o *PutAlmanaxWebhook) SetBonusBlacklistNil(b bool)`

 SetBonusBlacklistNil sets the value for BonusBlacklist to be an explicit nil

### UnsetBonusBlacklist
`func (o *PutAlmanaxWebhook) UnsetBonusBlacklist()`

UnsetBonusBlacklist ensures that no value is present for BonusBlacklist, not even an explicit nil
### GetSubscriptions

`func (o *PutAlmanaxWebhook) GetSubscriptions() []string`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *PutAlmanaxWebhook) GetSubscriptionsOk() (*[]string, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *PutAlmanaxWebhook) SetSubscriptions(v []string)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *PutAlmanaxWebhook) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### SetSubscriptionsNil

`func (o *PutAlmanaxWebhook) SetSubscriptionsNil(b bool)`

 SetSubscriptionsNil sets the value for Subscriptions to be an explicit nil

### UnsetSubscriptions
`func (o *PutAlmanaxWebhook) UnsetSubscriptions()`

UnsetSubscriptions ensures that no value is present for Subscriptions, not even an explicit nil
### GetDailySettings

`func (o *PutAlmanaxWebhook) GetDailySettings() CreateAlmanaxWebhookDailySettings`

GetDailySettings returns the DailySettings field if non-nil, zero value otherwise.

### GetDailySettingsOk

`func (o *PutAlmanaxWebhook) GetDailySettingsOk() (*CreateAlmanaxWebhookDailySettings, bool)`

GetDailySettingsOk returns a tuple with the DailySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySettings

`func (o *PutAlmanaxWebhook) SetDailySettings(v CreateAlmanaxWebhookDailySettings)`

SetDailySettings sets DailySettings field to given value.

### HasDailySettings

`func (o *PutAlmanaxWebhook) HasDailySettings() bool`

HasDailySettings returns a boolean if a field has been set.

### SetDailySettingsNil

`func (o *PutAlmanaxWebhook) SetDailySettingsNil(b bool)`

 SetDailySettingsNil sets the value for DailySettings to be an explicit nil

### UnsetDailySettings
`func (o *PutAlmanaxWebhook) UnsetDailySettings()`

UnsetDailySettings ensures that no value is present for DailySettings, not even an explicit nil
### GetIsoDate

`func (o *PutAlmanaxWebhook) GetIsoDate() bool`

GetIsoDate returns the IsoDate field if non-nil, zero value otherwise.

### GetIsoDateOk

`func (o *PutAlmanaxWebhook) GetIsoDateOk() (*bool, bool)`

GetIsoDateOk returns a tuple with the IsoDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoDate

`func (o *PutAlmanaxWebhook) SetIsoDate(v bool)`

SetIsoDate sets IsoDate field to given value.

### HasIsoDate

`func (o *PutAlmanaxWebhook) HasIsoDate() bool`

HasIsoDate returns a boolean if a field has been set.

### SetIsoDateNil

`func (o *PutAlmanaxWebhook) SetIsoDateNil(b bool)`

 SetIsoDateNil sets the value for IsoDate to be an explicit nil

### UnsetIsoDate
`func (o *PutAlmanaxWebhook) UnsetIsoDate()`

UnsetIsoDate ensures that no value is present for IsoDate, not even an explicit nil
### GetMentions

`func (o *PutAlmanaxWebhook) GetMentions() map[string][]CreateAlmanaxWebhookMentionsValueInner`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *PutAlmanaxWebhook) GetMentionsOk() (*map[string][]CreateAlmanaxWebhookMentionsValueInner, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *PutAlmanaxWebhook) SetMentions(v map[string][]CreateAlmanaxWebhookMentionsValueInner)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *PutAlmanaxWebhook) HasMentions() bool`

HasMentions returns a boolean if a field has been set.

### GetIntervals

`func (o *PutAlmanaxWebhook) GetIntervals() []string`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *PutAlmanaxWebhook) GetIntervalsOk() (*[]string, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *PutAlmanaxWebhook) SetIntervals(v []string)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *PutAlmanaxWebhook) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.

### SetIntervalsNil

`func (o *PutAlmanaxWebhook) SetIntervalsNil(b bool)`

 SetIntervalsNil sets the value for Intervals to be an explicit nil

### UnsetIntervals
`func (o *PutAlmanaxWebhook) UnsetIntervals()`

UnsetIntervals ensures that no value is present for Intervals, not even an explicit nil
### GetWeeklyWeekday

`func (o *PutAlmanaxWebhook) GetWeeklyWeekday() string`

GetWeeklyWeekday returns the WeeklyWeekday field if non-nil, zero value otherwise.

### GetWeeklyWeekdayOk

`func (o *PutAlmanaxWebhook) GetWeeklyWeekdayOk() (*string, bool)`

GetWeeklyWeekdayOk returns a tuple with the WeeklyWeekday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklyWeekday

`func (o *PutAlmanaxWebhook) SetWeeklyWeekday(v string)`

SetWeeklyWeekday sets WeeklyWeekday field to given value.

### HasWeeklyWeekday

`func (o *PutAlmanaxWebhook) HasWeeklyWeekday() bool`

HasWeeklyWeekday returns a boolean if a field has been set.

### SetWeeklyWeekdayNil

`func (o *PutAlmanaxWebhook) SetWeeklyWeekdayNil(b bool)`

 SetWeeklyWeekdayNil sets the value for WeeklyWeekday to be an explicit nil

### UnsetWeeklyWeekday
`func (o *PutAlmanaxWebhook) UnsetWeeklyWeekday()`

UnsetWeeklyWeekday ensures that no value is present for WeeklyWeekday, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


