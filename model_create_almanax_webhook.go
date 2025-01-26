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
	"bytes"
	"fmt"
)

// checks if the CreateAlmanaxWebhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAlmanaxWebhook{}

// CreateAlmanaxWebhook struct for CreateAlmanaxWebhook
type CreateAlmanaxWebhook struct {
	// from all available bonuses (ids) from /dofus3/meta/{language}/almanax/bonuses
	BonusWhitelist []string `json:"bonus_whitelist,omitempty"`
	// from all available bonuses (ids) from /dofus3/meta/{language}/almanax/bonuses
	BonusBlacklist []string `json:"bonus_blacklist,omitempty"`
	// Get the available subscriptions with /meta/webhooks/almanax
	Subscriptions []string `json:"subscriptions"`
	Format string `json:"format"`
	// Discord Webhook URL
	Callback string `json:"callback"`
	DailySettings NullableCreateAlmanaxWebhookDailySettings `json:"daily_settings,omitempty"`
	// If false, it will use common local time formats and weekday translations. If true, the format is YYYY-MM-DD.
	IsoDate NullableBool `json:"iso_date,omitempty"`
	// Almanax bonus ids mapped to array of mentions.
	Mentions map[string][]CreateAlmanaxWebhookMentionsValueInner `json:"mentions,omitempty"`
	// - Daily posts each day, filtering with Black/Whitelist and mentions are applied daily. - Weekly posts the next 7 days (excluding the posting day) once per week at the specified time. With only weekly selected, of all mentions, only prior notices will come through daily. The 7 day preview gets filtered by the Black/Whitelist. - Monthly posts a preview of the next month from first to last date. The post will be on the last day of a month (ignoring day of the week) at the specified time. Mentions and filtering works like weekly. The biggest difference between daily and the other two is that daily always posts the current day while monthly and weekly only show future days. You can always combine the intervals by selecting multiple intervals for one hook or create multiple hooks for the same channel with different settings to get every highly specific combination you want.
	Intervals []string `json:"intervals"`
	// When to post the weekly preview at the specified time.
	WeeklyWeekday NullableString `json:"weekly_weekday,omitempty"`
}

type _CreateAlmanaxWebhook CreateAlmanaxWebhook

// NewCreateAlmanaxWebhook instantiates a new CreateAlmanaxWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAlmanaxWebhook(subscriptions []string, format string, callback string, intervals []string) *CreateAlmanaxWebhook {
	this := CreateAlmanaxWebhook{}
	this.Subscriptions = subscriptions
	this.Format = format
	this.Callback = callback
	var isoDate bool = false
	this.IsoDate = *NewNullableBool(&isoDate)
	this.Intervals = intervals
	return &this
}

// NewCreateAlmanaxWebhookWithDefaults instantiates a new CreateAlmanaxWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAlmanaxWebhookWithDefaults() *CreateAlmanaxWebhook {
	this := CreateAlmanaxWebhook{}
	var isoDate bool = false
	this.IsoDate = *NewNullableBool(&isoDate)
	return &this
}

// GetBonusWhitelist returns the BonusWhitelist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAlmanaxWebhook) GetBonusWhitelist() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.BonusWhitelist
}

// GetBonusWhitelistOk returns a tuple with the BonusWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAlmanaxWebhook) GetBonusWhitelistOk() ([]string, bool) {
	if o == nil || IsNil(o.BonusWhitelist) {
		return nil, false
	}
	return o.BonusWhitelist, true
}

// HasBonusWhitelist returns a boolean if a field has been set.
func (o *CreateAlmanaxWebhook) HasBonusWhitelist() bool {
	if o != nil && !IsNil(o.BonusWhitelist) {
		return true
	}

	return false
}

// SetBonusWhitelist gets a reference to the given []string and assigns it to the BonusWhitelist field.
func (o *CreateAlmanaxWebhook) SetBonusWhitelist(v []string) {
	o.BonusWhitelist = v
}

// GetBonusBlacklist returns the BonusBlacklist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAlmanaxWebhook) GetBonusBlacklist() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.BonusBlacklist
}

// GetBonusBlacklistOk returns a tuple with the BonusBlacklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAlmanaxWebhook) GetBonusBlacklistOk() ([]string, bool) {
	if o == nil || IsNil(o.BonusBlacklist) {
		return nil, false
	}
	return o.BonusBlacklist, true
}

// HasBonusBlacklist returns a boolean if a field has been set.
func (o *CreateAlmanaxWebhook) HasBonusBlacklist() bool {
	if o != nil && !IsNil(o.BonusBlacklist) {
		return true
	}

	return false
}

// SetBonusBlacklist gets a reference to the given []string and assigns it to the BonusBlacklist field.
func (o *CreateAlmanaxWebhook) SetBonusBlacklist(v []string) {
	o.BonusBlacklist = v
}

// GetSubscriptions returns the Subscriptions field value
func (o *CreateAlmanaxWebhook) GetSubscriptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value
// and a boolean to check if the value has been set.
func (o *CreateAlmanaxWebhook) GetSubscriptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// SetSubscriptions sets field value
func (o *CreateAlmanaxWebhook) SetSubscriptions(v []string) {
	o.Subscriptions = v
}

// GetFormat returns the Format field value
func (o *CreateAlmanaxWebhook) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *CreateAlmanaxWebhook) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *CreateAlmanaxWebhook) SetFormat(v string) {
	o.Format = v
}

// GetCallback returns the Callback field value
func (o *CreateAlmanaxWebhook) GetCallback() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value
// and a boolean to check if the value has been set.
func (o *CreateAlmanaxWebhook) GetCallbackOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Callback, true
}

// SetCallback sets field value
func (o *CreateAlmanaxWebhook) SetCallback(v string) {
	o.Callback = v
}

// GetDailySettings returns the DailySettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAlmanaxWebhook) GetDailySettings() CreateAlmanaxWebhookDailySettings {
	if o == nil || IsNil(o.DailySettings.Get()) {
		var ret CreateAlmanaxWebhookDailySettings
		return ret
	}
	return *o.DailySettings.Get()
}

// GetDailySettingsOk returns a tuple with the DailySettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAlmanaxWebhook) GetDailySettingsOk() (*CreateAlmanaxWebhookDailySettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.DailySettings.Get(), o.DailySettings.IsSet()
}

// HasDailySettings returns a boolean if a field has been set.
func (o *CreateAlmanaxWebhook) HasDailySettings() bool {
	if o != nil && o.DailySettings.IsSet() {
		return true
	}

	return false
}

// SetDailySettings gets a reference to the given NullableCreateAlmanaxWebhookDailySettings and assigns it to the DailySettings field.
func (o *CreateAlmanaxWebhook) SetDailySettings(v CreateAlmanaxWebhookDailySettings) {
	o.DailySettings.Set(&v)
}
// SetDailySettingsNil sets the value for DailySettings to be an explicit nil
func (o *CreateAlmanaxWebhook) SetDailySettingsNil() {
	o.DailySettings.Set(nil)
}

// UnsetDailySettings ensures that no value is present for DailySettings, not even an explicit nil
func (o *CreateAlmanaxWebhook) UnsetDailySettings() {
	o.DailySettings.Unset()
}

// GetIsoDate returns the IsoDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAlmanaxWebhook) GetIsoDate() bool {
	if o == nil || IsNil(o.IsoDate.Get()) {
		var ret bool
		return ret
	}
	return *o.IsoDate.Get()
}

// GetIsoDateOk returns a tuple with the IsoDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAlmanaxWebhook) GetIsoDateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsoDate.Get(), o.IsoDate.IsSet()
}

// HasIsoDate returns a boolean if a field has been set.
func (o *CreateAlmanaxWebhook) HasIsoDate() bool {
	if o != nil && o.IsoDate.IsSet() {
		return true
	}

	return false
}

// SetIsoDate gets a reference to the given NullableBool and assigns it to the IsoDate field.
func (o *CreateAlmanaxWebhook) SetIsoDate(v bool) {
	o.IsoDate.Set(&v)
}
// SetIsoDateNil sets the value for IsoDate to be an explicit nil
func (o *CreateAlmanaxWebhook) SetIsoDateNil() {
	o.IsoDate.Set(nil)
}

// UnsetIsoDate ensures that no value is present for IsoDate, not even an explicit nil
func (o *CreateAlmanaxWebhook) UnsetIsoDate() {
	o.IsoDate.Unset()
}

// GetMentions returns the Mentions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAlmanaxWebhook) GetMentions() map[string][]CreateAlmanaxWebhookMentionsValueInner {
	if o == nil {
		var ret map[string][]CreateAlmanaxWebhookMentionsValueInner
		return ret
	}
	return o.Mentions
}

// GetMentionsOk returns a tuple with the Mentions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAlmanaxWebhook) GetMentionsOk() (*map[string][]CreateAlmanaxWebhookMentionsValueInner, bool) {
	if o == nil || IsNil(o.Mentions) {
		return nil, false
	}
	return &o.Mentions, true
}

// HasMentions returns a boolean if a field has been set.
func (o *CreateAlmanaxWebhook) HasMentions() bool {
	if o != nil && !IsNil(o.Mentions) {
		return true
	}

	return false
}

// SetMentions gets a reference to the given map[string][]CreateAlmanaxWebhookMentionsValueInner and assigns it to the Mentions field.
func (o *CreateAlmanaxWebhook) SetMentions(v map[string][]CreateAlmanaxWebhookMentionsValueInner) {
	o.Mentions = v
}

// GetIntervals returns the Intervals field value
func (o *CreateAlmanaxWebhook) GetIntervals() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Intervals
}

// GetIntervalsOk returns a tuple with the Intervals field value
// and a boolean to check if the value has been set.
func (o *CreateAlmanaxWebhook) GetIntervalsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Intervals, true
}

// SetIntervals sets field value
func (o *CreateAlmanaxWebhook) SetIntervals(v []string) {
	o.Intervals = v
}

// GetWeeklyWeekday returns the WeeklyWeekday field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAlmanaxWebhook) GetWeeklyWeekday() string {
	if o == nil || IsNil(o.WeeklyWeekday.Get()) {
		var ret string
		return ret
	}
	return *o.WeeklyWeekday.Get()
}

// GetWeeklyWeekdayOk returns a tuple with the WeeklyWeekday field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAlmanaxWebhook) GetWeeklyWeekdayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WeeklyWeekday.Get(), o.WeeklyWeekday.IsSet()
}

// HasWeeklyWeekday returns a boolean if a field has been set.
func (o *CreateAlmanaxWebhook) HasWeeklyWeekday() bool {
	if o != nil && o.WeeklyWeekday.IsSet() {
		return true
	}

	return false
}

// SetWeeklyWeekday gets a reference to the given NullableString and assigns it to the WeeklyWeekday field.
func (o *CreateAlmanaxWebhook) SetWeeklyWeekday(v string) {
	o.WeeklyWeekday.Set(&v)
}
// SetWeeklyWeekdayNil sets the value for WeeklyWeekday to be an explicit nil
func (o *CreateAlmanaxWebhook) SetWeeklyWeekdayNil() {
	o.WeeklyWeekday.Set(nil)
}

// UnsetWeeklyWeekday ensures that no value is present for WeeklyWeekday, not even an explicit nil
func (o *CreateAlmanaxWebhook) UnsetWeeklyWeekday() {
	o.WeeklyWeekday.Unset()
}

func (o CreateAlmanaxWebhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAlmanaxWebhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BonusWhitelist != nil {
		toSerialize["bonus_whitelist"] = o.BonusWhitelist
	}
	if o.BonusBlacklist != nil {
		toSerialize["bonus_blacklist"] = o.BonusBlacklist
	}
	toSerialize["subscriptions"] = o.Subscriptions
	toSerialize["format"] = o.Format
	toSerialize["callback"] = o.Callback
	if o.DailySettings.IsSet() {
		toSerialize["daily_settings"] = o.DailySettings.Get()
	}
	if o.IsoDate.IsSet() {
		toSerialize["iso_date"] = o.IsoDate.Get()
	}
	if o.Mentions != nil {
		toSerialize["mentions"] = o.Mentions
	}
	toSerialize["intervals"] = o.Intervals
	if o.WeeklyWeekday.IsSet() {
		toSerialize["weekly_weekday"] = o.WeeklyWeekday.Get()
	}
	return toSerialize, nil
}

func (o *CreateAlmanaxWebhook) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subscriptions",
		"format",
		"callback",
		"intervals",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateAlmanaxWebhook := _CreateAlmanaxWebhook{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateAlmanaxWebhook)

	if err != nil {
		return err
	}

	*o = CreateAlmanaxWebhook(varCreateAlmanaxWebhook)

	return err
}

type NullableCreateAlmanaxWebhook struct {
	value *CreateAlmanaxWebhook
	isSet bool
}

func (v NullableCreateAlmanaxWebhook) Get() *CreateAlmanaxWebhook {
	return v.value
}

func (v *NullableCreateAlmanaxWebhook) Set(val *CreateAlmanaxWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAlmanaxWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAlmanaxWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAlmanaxWebhook(val *CreateAlmanaxWebhook) *NullableCreateAlmanaxWebhook {
	return &NullableCreateAlmanaxWebhook{value: val, isSet: true}
}

func (v NullableCreateAlmanaxWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAlmanaxWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


