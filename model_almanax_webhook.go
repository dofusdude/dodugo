/*
dofusdude

# Open Ankama Developer Community The all-in-one toolbelt for your next Ankama related project.  ## Versions - [Dofus 2](https://docs.dofusdu.de/dofus2/) - [Dofus 3](https://docs.dofusdu.de/dofus3/)   - v1 [latest] (you are here)   ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save` - [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save` - [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo` - [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude` - [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Almanax Discord Integration** Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 3 Beta** from stable to bleeding edge by replacing /dofus3 with /dofus3beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_, _de_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Official Sources** generated from actual data from the game.  ... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h). 

API version: 1.0.0
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
	"time"
)

// checks if the AlmanaxWebhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlmanaxWebhook{}

// AlmanaxWebhook 
type AlmanaxWebhook struct {
	Id *string `json:"id,omitempty"`
	DailySettings *AlmanaxWebhookDailySettings `json:"daily_settings,omitempty"`
	// Only post when these bonuses come up. From all available bonuses (ids) from /dofus3/meta/{language}/almanax/bonuses.
	BonusWhitelist []string `json:"bonus_whitelist,omitempty"`
	// Skip the day when these bonuses come up. From all available bonuses (ids) from /dofus3/meta/{language}/almanax/bonuses
	BonusBlacklist []string `json:"bonus_blacklist,omitempty"`
	// Get the available subscriptions with /meta/webhooks/almanax
	Subscriptions []string `json:"subscriptions,omitempty"`
	// If false, it will use common local time formats and weekday translations. If true, the format is YYYY-MM-DD.
	IsoDate *bool `json:"iso_date,omitempty"`
	// Almanax bonus ids mapped to array of mentions.
	Mentions map[string][]CreateAlmanaxWebhookMentionsValueInner `json:"mentions,omitempty"`
	// - Daily posts each day, filtering with Black/Whitelist and mentions are applied daily. - Weekly posts the next 7 days (excluding the posting day) once per week at the specified time. With only weekly selected, of all mentions, only prior notices will come through daily. The 7 day preview gets filtered by the Black/Whitelist. - Monthly posts a preview of the next month from first to last date. The post will be on the last day of a month (ignoring day of the week) at the specified time. Mentions and filtering works like weekly. The biggest difference between daily and the other two is that daily always posts the current day while monthly and weekly only show future days. You can always combine the intervals by selecting multiple intervals for one hook or create multiple hooks for the same channel with different settings to get every highly specific combination you want.
	Intervals []string `json:"intervals,omitempty"`
	// When to post the weekly preview at the specified time.
	WeeklyWeekday NullableString `json:"weekly_weekday,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	LastFiredAt NullableTime `json:"last_fired_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewAlmanaxWebhook instantiates a new AlmanaxWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlmanaxWebhook() *AlmanaxWebhook {
	this := AlmanaxWebhook{}
	var isoDate bool = false
	this.IsoDate = &isoDate
	return &this
}

// NewAlmanaxWebhookWithDefaults instantiates a new AlmanaxWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlmanaxWebhookWithDefaults() *AlmanaxWebhook {
	this := AlmanaxWebhook{}
	var isoDate bool = false
	this.IsoDate = &isoDate
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlmanaxWebhook) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxWebhook) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlmanaxWebhook) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AlmanaxWebhook) SetId(v string) {
	o.Id = &v
}

// GetDailySettings returns the DailySettings field value if set, zero value otherwise.
func (o *AlmanaxWebhook) GetDailySettings() AlmanaxWebhookDailySettings {
	if o == nil || IsNil(o.DailySettings) {
		var ret AlmanaxWebhookDailySettings
		return ret
	}
	return *o.DailySettings
}

// GetDailySettingsOk returns a tuple with the DailySettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxWebhook) GetDailySettingsOk() (*AlmanaxWebhookDailySettings, bool) {
	if o == nil || IsNil(o.DailySettings) {
		return nil, false
	}
	return o.DailySettings, true
}

// HasDailySettings returns a boolean if a field has been set.
func (o *AlmanaxWebhook) HasDailySettings() bool {
	if o != nil && !IsNil(o.DailySettings) {
		return true
	}

	return false
}

// SetDailySettings gets a reference to the given AlmanaxWebhookDailySettings and assigns it to the DailySettings field.
func (o *AlmanaxWebhook) SetDailySettings(v AlmanaxWebhookDailySettings) {
	o.DailySettings = &v
}

// GetBonusWhitelist returns the BonusWhitelist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlmanaxWebhook) GetBonusWhitelist() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.BonusWhitelist
}

// GetBonusWhitelistOk returns a tuple with the BonusWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlmanaxWebhook) GetBonusWhitelistOk() ([]string, bool) {
	if o == nil || IsNil(o.BonusWhitelist) {
		return nil, false
	}
	return o.BonusWhitelist, true
}

// HasBonusWhitelist returns a boolean if a field has been set.
func (o *AlmanaxWebhook) HasBonusWhitelist() bool {
	if o != nil && !IsNil(o.BonusWhitelist) {
		return true
	}

	return false
}

// SetBonusWhitelist gets a reference to the given []string and assigns it to the BonusWhitelist field.
func (o *AlmanaxWebhook) SetBonusWhitelist(v []string) {
	o.BonusWhitelist = v
}

// GetBonusBlacklist returns the BonusBlacklist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlmanaxWebhook) GetBonusBlacklist() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.BonusBlacklist
}

// GetBonusBlacklistOk returns a tuple with the BonusBlacklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlmanaxWebhook) GetBonusBlacklistOk() ([]string, bool) {
	if o == nil || IsNil(o.BonusBlacklist) {
		return nil, false
	}
	return o.BonusBlacklist, true
}

// HasBonusBlacklist returns a boolean if a field has been set.
func (o *AlmanaxWebhook) HasBonusBlacklist() bool {
	if o != nil && !IsNil(o.BonusBlacklist) {
		return true
	}

	return false
}

// SetBonusBlacklist gets a reference to the given []string and assigns it to the BonusBlacklist field.
func (o *AlmanaxWebhook) SetBonusBlacklist(v []string) {
	o.BonusBlacklist = v
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise.
func (o *AlmanaxWebhook) GetSubscriptions() []string {
	if o == nil || IsNil(o.Subscriptions) {
		var ret []string
		return ret
	}
	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxWebhook) GetSubscriptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Subscriptions) {
		return nil, false
	}
	return o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *AlmanaxWebhook) HasSubscriptions() bool {
	if o != nil && !IsNil(o.Subscriptions) {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given []string and assigns it to the Subscriptions field.
func (o *AlmanaxWebhook) SetSubscriptions(v []string) {
	o.Subscriptions = v
}

// GetIsoDate returns the IsoDate field value if set, zero value otherwise.
func (o *AlmanaxWebhook) GetIsoDate() bool {
	if o == nil || IsNil(o.IsoDate) {
		var ret bool
		return ret
	}
	return *o.IsoDate
}

// GetIsoDateOk returns a tuple with the IsoDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxWebhook) GetIsoDateOk() (*bool, bool) {
	if o == nil || IsNil(o.IsoDate) {
		return nil, false
	}
	return o.IsoDate, true
}

// HasIsoDate returns a boolean if a field has been set.
func (o *AlmanaxWebhook) HasIsoDate() bool {
	if o != nil && !IsNil(o.IsoDate) {
		return true
	}

	return false
}

// SetIsoDate gets a reference to the given bool and assigns it to the IsoDate field.
func (o *AlmanaxWebhook) SetIsoDate(v bool) {
	o.IsoDate = &v
}

// GetMentions returns the Mentions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlmanaxWebhook) GetMentions() map[string][]CreateAlmanaxWebhookMentionsValueInner {
	if o == nil {
		var ret map[string][]CreateAlmanaxWebhookMentionsValueInner
		return ret
	}
	return o.Mentions
}

// GetMentionsOk returns a tuple with the Mentions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlmanaxWebhook) GetMentionsOk() (*map[string][]CreateAlmanaxWebhookMentionsValueInner, bool) {
	if o == nil || IsNil(o.Mentions) {
		return nil, false
	}
	return &o.Mentions, true
}

// HasMentions returns a boolean if a field has been set.
func (o *AlmanaxWebhook) HasMentions() bool {
	if o != nil && !IsNil(o.Mentions) {
		return true
	}

	return false
}

// SetMentions gets a reference to the given map[string][]CreateAlmanaxWebhookMentionsValueInner and assigns it to the Mentions field.
func (o *AlmanaxWebhook) SetMentions(v map[string][]CreateAlmanaxWebhookMentionsValueInner) {
	o.Mentions = v
}

// GetIntervals returns the Intervals field value if set, zero value otherwise.
func (o *AlmanaxWebhook) GetIntervals() []string {
	if o == nil || IsNil(o.Intervals) {
		var ret []string
		return ret
	}
	return o.Intervals
}

// GetIntervalsOk returns a tuple with the Intervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxWebhook) GetIntervalsOk() ([]string, bool) {
	if o == nil || IsNil(o.Intervals) {
		return nil, false
	}
	return o.Intervals, true
}

// HasIntervals returns a boolean if a field has been set.
func (o *AlmanaxWebhook) HasIntervals() bool {
	if o != nil && !IsNil(o.Intervals) {
		return true
	}

	return false
}

// SetIntervals gets a reference to the given []string and assigns it to the Intervals field.
func (o *AlmanaxWebhook) SetIntervals(v []string) {
	o.Intervals = v
}

// GetWeeklyWeekday returns the WeeklyWeekday field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlmanaxWebhook) GetWeeklyWeekday() string {
	if o == nil || IsNil(o.WeeklyWeekday.Get()) {
		var ret string
		return ret
	}
	return *o.WeeklyWeekday.Get()
}

// GetWeeklyWeekdayOk returns a tuple with the WeeklyWeekday field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlmanaxWebhook) GetWeeklyWeekdayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WeeklyWeekday.Get(), o.WeeklyWeekday.IsSet()
}

// HasWeeklyWeekday returns a boolean if a field has been set.
func (o *AlmanaxWebhook) HasWeeklyWeekday() bool {
	if o != nil && o.WeeklyWeekday.IsSet() {
		return true
	}

	return false
}

// SetWeeklyWeekday gets a reference to the given NullableString and assigns it to the WeeklyWeekday field.
func (o *AlmanaxWebhook) SetWeeklyWeekday(v string) {
	o.WeeklyWeekday.Set(&v)
}
// SetWeeklyWeekdayNil sets the value for WeeklyWeekday to be an explicit nil
func (o *AlmanaxWebhook) SetWeeklyWeekdayNil() {
	o.WeeklyWeekday.Set(nil)
}

// UnsetWeeklyWeekday ensures that no value is present for WeeklyWeekday, not even an explicit nil
func (o *AlmanaxWebhook) UnsetWeeklyWeekday() {
	o.WeeklyWeekday.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AlmanaxWebhook) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxWebhook) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AlmanaxWebhook) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AlmanaxWebhook) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastFiredAt returns the LastFiredAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlmanaxWebhook) GetLastFiredAt() time.Time {
	if o == nil || IsNil(o.LastFiredAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastFiredAt.Get()
}

// GetLastFiredAtOk returns a tuple with the LastFiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlmanaxWebhook) GetLastFiredAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastFiredAt.Get(), o.LastFiredAt.IsSet()
}

// HasLastFiredAt returns a boolean if a field has been set.
func (o *AlmanaxWebhook) HasLastFiredAt() bool {
	if o != nil && o.LastFiredAt.IsSet() {
		return true
	}

	return false
}

// SetLastFiredAt gets a reference to the given NullableTime and assigns it to the LastFiredAt field.
func (o *AlmanaxWebhook) SetLastFiredAt(v time.Time) {
	o.LastFiredAt.Set(&v)
}
// SetLastFiredAtNil sets the value for LastFiredAt to be an explicit nil
func (o *AlmanaxWebhook) SetLastFiredAtNil() {
	o.LastFiredAt.Set(nil)
}

// UnsetLastFiredAt ensures that no value is present for LastFiredAt, not even an explicit nil
func (o *AlmanaxWebhook) UnsetLastFiredAt() {
	o.LastFiredAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AlmanaxWebhook) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxWebhook) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AlmanaxWebhook) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AlmanaxWebhook) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o AlmanaxWebhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlmanaxWebhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.DailySettings) {
		toSerialize["daily_settings"] = o.DailySettings
	}
	if o.BonusWhitelist != nil {
		toSerialize["bonus_whitelist"] = o.BonusWhitelist
	}
	if o.BonusBlacklist != nil {
		toSerialize["bonus_blacklist"] = o.BonusBlacklist
	}
	if !IsNil(o.Subscriptions) {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	if !IsNil(o.IsoDate) {
		toSerialize["iso_date"] = o.IsoDate
	}
	if o.Mentions != nil {
		toSerialize["mentions"] = o.Mentions
	}
	if !IsNil(o.Intervals) {
		toSerialize["intervals"] = o.Intervals
	}
	if o.WeeklyWeekday.IsSet() {
		toSerialize["weekly_weekday"] = o.WeeklyWeekday.Get()
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastFiredAt.IsSet() {
		toSerialize["last_fired_at"] = o.LastFiredAt.Get()
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableAlmanaxWebhook struct {
	value *AlmanaxWebhook
	isSet bool
}

func (v NullableAlmanaxWebhook) Get() *AlmanaxWebhook {
	return v.value
}

func (v *NullableAlmanaxWebhook) Set(val *AlmanaxWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableAlmanaxWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableAlmanaxWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlmanaxWebhook(val *AlmanaxWebhook) *NullableAlmanaxWebhook {
	return &NullableAlmanaxWebhook{value: val, isSet: true}
}

func (v NullableAlmanaxWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlmanaxWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


