/*
dofusdude

# Open Ankama Developer Community The all-in-one toolbelt for your next Ankama related project.  ## Versions - [Dofus 2](https://docs.dofusdu.de/dofus2/) - [Dofus 3](https://docs.dofusdu.de/dofus3/)   - v1 [latest] (you are here)   ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save` - [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save` - [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo` - [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude` - [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Almanax Discord Integration** Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 3 Beta** from stable to bleeding edge by replacing /dofus3 with /dofus3beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_, _de_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Official Sources** generated from actual data from the game.  ... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h). 

API version: 1.0.0-rc.5
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
)

// checks if the PutAlmanaxWebhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutAlmanaxWebhook{}

// PutAlmanaxWebhook struct for PutAlmanaxWebhook
type PutAlmanaxWebhook struct {
	// from all available bonuses (ids) from /dofus3/meta/{language}/almanax/bonuses. Delete old entries with empty array []. Just null changes nothing.
	BonusWhitelist []string `json:"bonus_whitelist,omitempty"`
	// from all available bonuses (ids) from /dofus3/meta/{language}/almanax/bonuses. Delete old entries with empty array []. Just null changes nothing.
	BonusBlacklist []string `json:"bonus_blacklist,omitempty"`
	// Get the available subscriptions with /meta/webhooks/almanax
	Subscriptions []string `json:"subscriptions,omitempty"`
	DailySettings NullableCreateAlmanaxWebhookDailySettings `json:"daily_settings,omitempty"`
	// If false, it will use common local time formats and weekday translations. If true, the format is YYYY-MM-DD.
	IsoDate NullableBool `json:"iso_date,omitempty"`
	// Almanax bonus ids mapped to array of mentions.
	Mentions *map[string][]CreateAlmanaxWebhookMentionsValueInner `json:"mentions,omitempty"`
	Intervals []string `json:"intervals,omitempty"`
	// When to post the weekly preview at the specified time.
	WeeklyWeekday NullableString `json:"weekly_weekday,omitempty"`
}

// NewPutAlmanaxWebhook instantiates a new PutAlmanaxWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutAlmanaxWebhook() *PutAlmanaxWebhook {
	this := PutAlmanaxWebhook{}
	var isoDate bool = false
	this.IsoDate = *NewNullableBool(&isoDate)
	return &this
}

// NewPutAlmanaxWebhookWithDefaults instantiates a new PutAlmanaxWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutAlmanaxWebhookWithDefaults() *PutAlmanaxWebhook {
	this := PutAlmanaxWebhook{}
	var isoDate bool = false
	this.IsoDate = *NewNullableBool(&isoDate)
	return &this
}

// GetBonusWhitelist returns the BonusWhitelist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PutAlmanaxWebhook) GetBonusWhitelist() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.BonusWhitelist
}

// GetBonusWhitelistOk returns a tuple with the BonusWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutAlmanaxWebhook) GetBonusWhitelistOk() ([]string, bool) {
	if o == nil || IsNil(o.BonusWhitelist) {
		return nil, false
	}
	return o.BonusWhitelist, true
}

// HasBonusWhitelist returns a boolean if a field has been set.
func (o *PutAlmanaxWebhook) HasBonusWhitelist() bool {
	if o != nil && !IsNil(o.BonusWhitelist) {
		return true
	}

	return false
}

// SetBonusWhitelist gets a reference to the given []string and assigns it to the BonusWhitelist field.
func (o *PutAlmanaxWebhook) SetBonusWhitelist(v []string) {
	o.BonusWhitelist = v
}

// GetBonusBlacklist returns the BonusBlacklist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PutAlmanaxWebhook) GetBonusBlacklist() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.BonusBlacklist
}

// GetBonusBlacklistOk returns a tuple with the BonusBlacklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutAlmanaxWebhook) GetBonusBlacklistOk() ([]string, bool) {
	if o == nil || IsNil(o.BonusBlacklist) {
		return nil, false
	}
	return o.BonusBlacklist, true
}

// HasBonusBlacklist returns a boolean if a field has been set.
func (o *PutAlmanaxWebhook) HasBonusBlacklist() bool {
	if o != nil && !IsNil(o.BonusBlacklist) {
		return true
	}

	return false
}

// SetBonusBlacklist gets a reference to the given []string and assigns it to the BonusBlacklist field.
func (o *PutAlmanaxWebhook) SetBonusBlacklist(v []string) {
	o.BonusBlacklist = v
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PutAlmanaxWebhook) GetSubscriptions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutAlmanaxWebhook) GetSubscriptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Subscriptions) {
		return nil, false
	}
	return o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *PutAlmanaxWebhook) HasSubscriptions() bool {
	if o != nil && !IsNil(o.Subscriptions) {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given []string and assigns it to the Subscriptions field.
func (o *PutAlmanaxWebhook) SetSubscriptions(v []string) {
	o.Subscriptions = v
}

// GetDailySettings returns the DailySettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PutAlmanaxWebhook) GetDailySettings() CreateAlmanaxWebhookDailySettings {
	if o == nil || IsNil(o.DailySettings.Get()) {
		var ret CreateAlmanaxWebhookDailySettings
		return ret
	}
	return *o.DailySettings.Get()
}

// GetDailySettingsOk returns a tuple with the DailySettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutAlmanaxWebhook) GetDailySettingsOk() (*CreateAlmanaxWebhookDailySettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.DailySettings.Get(), o.DailySettings.IsSet()
}

// HasDailySettings returns a boolean if a field has been set.
func (o *PutAlmanaxWebhook) HasDailySettings() bool {
	if o != nil && o.DailySettings.IsSet() {
		return true
	}

	return false
}

// SetDailySettings gets a reference to the given NullableCreateAlmanaxWebhookDailySettings and assigns it to the DailySettings field.
func (o *PutAlmanaxWebhook) SetDailySettings(v CreateAlmanaxWebhookDailySettings) {
	o.DailySettings.Set(&v)
}
// SetDailySettingsNil sets the value for DailySettings to be an explicit nil
func (o *PutAlmanaxWebhook) SetDailySettingsNil() {
	o.DailySettings.Set(nil)
}

// UnsetDailySettings ensures that no value is present for DailySettings, not even an explicit nil
func (o *PutAlmanaxWebhook) UnsetDailySettings() {
	o.DailySettings.Unset()
}

// GetIsoDate returns the IsoDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PutAlmanaxWebhook) GetIsoDate() bool {
	if o == nil || IsNil(o.IsoDate.Get()) {
		var ret bool
		return ret
	}
	return *o.IsoDate.Get()
}

// GetIsoDateOk returns a tuple with the IsoDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutAlmanaxWebhook) GetIsoDateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsoDate.Get(), o.IsoDate.IsSet()
}

// HasIsoDate returns a boolean if a field has been set.
func (o *PutAlmanaxWebhook) HasIsoDate() bool {
	if o != nil && o.IsoDate.IsSet() {
		return true
	}

	return false
}

// SetIsoDate gets a reference to the given NullableBool and assigns it to the IsoDate field.
func (o *PutAlmanaxWebhook) SetIsoDate(v bool) {
	o.IsoDate.Set(&v)
}
// SetIsoDateNil sets the value for IsoDate to be an explicit nil
func (o *PutAlmanaxWebhook) SetIsoDateNil() {
	o.IsoDate.Set(nil)
}

// UnsetIsoDate ensures that no value is present for IsoDate, not even an explicit nil
func (o *PutAlmanaxWebhook) UnsetIsoDate() {
	o.IsoDate.Unset()
}

// GetMentions returns the Mentions field value if set, zero value otherwise.
func (o *PutAlmanaxWebhook) GetMentions() map[string][]CreateAlmanaxWebhookMentionsValueInner {
	if o == nil || IsNil(o.Mentions) {
		var ret map[string][]CreateAlmanaxWebhookMentionsValueInner
		return ret
	}
	return *o.Mentions
}

// GetMentionsOk returns a tuple with the Mentions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutAlmanaxWebhook) GetMentionsOk() (*map[string][]CreateAlmanaxWebhookMentionsValueInner, bool) {
	if o == nil || IsNil(o.Mentions) {
		return nil, false
	}
	return o.Mentions, true
}

// HasMentions returns a boolean if a field has been set.
func (o *PutAlmanaxWebhook) HasMentions() bool {
	if o != nil && !IsNil(o.Mentions) {
		return true
	}

	return false
}

// SetMentions gets a reference to the given map[string][]CreateAlmanaxWebhookMentionsValueInner and assigns it to the Mentions field.
func (o *PutAlmanaxWebhook) SetMentions(v map[string][]CreateAlmanaxWebhookMentionsValueInner) {
	o.Mentions = &v
}

// GetIntervals returns the Intervals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PutAlmanaxWebhook) GetIntervals() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Intervals
}

// GetIntervalsOk returns a tuple with the Intervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutAlmanaxWebhook) GetIntervalsOk() ([]string, bool) {
	if o == nil || IsNil(o.Intervals) {
		return nil, false
	}
	return o.Intervals, true
}

// HasIntervals returns a boolean if a field has been set.
func (o *PutAlmanaxWebhook) HasIntervals() bool {
	if o != nil && !IsNil(o.Intervals) {
		return true
	}

	return false
}

// SetIntervals gets a reference to the given []string and assigns it to the Intervals field.
func (o *PutAlmanaxWebhook) SetIntervals(v []string) {
	o.Intervals = v
}

// GetWeeklyWeekday returns the WeeklyWeekday field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PutAlmanaxWebhook) GetWeeklyWeekday() string {
	if o == nil || IsNil(o.WeeklyWeekday.Get()) {
		var ret string
		return ret
	}
	return *o.WeeklyWeekday.Get()
}

// GetWeeklyWeekdayOk returns a tuple with the WeeklyWeekday field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutAlmanaxWebhook) GetWeeklyWeekdayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WeeklyWeekday.Get(), o.WeeklyWeekday.IsSet()
}

// HasWeeklyWeekday returns a boolean if a field has been set.
func (o *PutAlmanaxWebhook) HasWeeklyWeekday() bool {
	if o != nil && o.WeeklyWeekday.IsSet() {
		return true
	}

	return false
}

// SetWeeklyWeekday gets a reference to the given NullableString and assigns it to the WeeklyWeekday field.
func (o *PutAlmanaxWebhook) SetWeeklyWeekday(v string) {
	o.WeeklyWeekday.Set(&v)
}
// SetWeeklyWeekdayNil sets the value for WeeklyWeekday to be an explicit nil
func (o *PutAlmanaxWebhook) SetWeeklyWeekdayNil() {
	o.WeeklyWeekday.Set(nil)
}

// UnsetWeeklyWeekday ensures that no value is present for WeeklyWeekday, not even an explicit nil
func (o *PutAlmanaxWebhook) UnsetWeeklyWeekday() {
	o.WeeklyWeekday.Unset()
}

func (o PutAlmanaxWebhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutAlmanaxWebhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BonusWhitelist != nil {
		toSerialize["bonus_whitelist"] = o.BonusWhitelist
	}
	if o.BonusBlacklist != nil {
		toSerialize["bonus_blacklist"] = o.BonusBlacklist
	}
	if o.Subscriptions != nil {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	if o.DailySettings.IsSet() {
		toSerialize["daily_settings"] = o.DailySettings.Get()
	}
	if o.IsoDate.IsSet() {
		toSerialize["iso_date"] = o.IsoDate.Get()
	}
	if !IsNil(o.Mentions) {
		toSerialize["mentions"] = o.Mentions
	}
	if o.Intervals != nil {
		toSerialize["intervals"] = o.Intervals
	}
	if o.WeeklyWeekday.IsSet() {
		toSerialize["weekly_weekday"] = o.WeeklyWeekday.Get()
	}
	return toSerialize, nil
}

type NullablePutAlmanaxWebhook struct {
	value *PutAlmanaxWebhook
	isSet bool
}

func (v NullablePutAlmanaxWebhook) Get() *PutAlmanaxWebhook {
	return v.value
}

func (v *NullablePutAlmanaxWebhook) Set(val *PutAlmanaxWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullablePutAlmanaxWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullablePutAlmanaxWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutAlmanaxWebhook(val *PutAlmanaxWebhook) *NullablePutAlmanaxWebhook {
	return &NullablePutAlmanaxWebhook{value: val, isSet: true}
}

func (v NullablePutAlmanaxWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutAlmanaxWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


