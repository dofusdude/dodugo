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

// checks if the CreateAlmanaxWebhookDailySettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAlmanaxWebhookDailySettings{}

// CreateAlmanaxWebhookDailySettings struct for CreateAlmanaxWebhookDailySettings
type CreateAlmanaxWebhookDailySettings struct {
	// Timezone of your community to determine midnight.
	Timezone NullableString `json:"timezone,omitempty"`
	// Hours added to midnight at the selected timezone. 8 = 8:00 in the morning.
	MidnightOffset NullableInt32 `json:"midnight_offset,omitempty"`
}

// NewCreateAlmanaxWebhookDailySettings instantiates a new CreateAlmanaxWebhookDailySettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAlmanaxWebhookDailySettings() *CreateAlmanaxWebhookDailySettings {
	this := CreateAlmanaxWebhookDailySettings{}
	var timezone string = "Europe/Paris"
	this.Timezone = *NewNullableString(&timezone)
	var midnightOffset int32 = 0
	this.MidnightOffset = *NewNullableInt32(&midnightOffset)
	return &this
}

// NewCreateAlmanaxWebhookDailySettingsWithDefaults instantiates a new CreateAlmanaxWebhookDailySettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAlmanaxWebhookDailySettingsWithDefaults() *CreateAlmanaxWebhookDailySettings {
	this := CreateAlmanaxWebhookDailySettings{}
	var timezone string = "Europe/Paris"
	this.Timezone = *NewNullableString(&timezone)
	var midnightOffset int32 = 0
	this.MidnightOffset = *NewNullableInt32(&midnightOffset)
	return &this
}

// GetTimezone returns the Timezone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAlmanaxWebhookDailySettings) GetTimezone() string {
	if o == nil || IsNil(o.Timezone.Get()) {
		var ret string
		return ret
	}
	return *o.Timezone.Get()
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAlmanaxWebhookDailySettings) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timezone.Get(), o.Timezone.IsSet()
}

// HasTimezone returns a boolean if a field has been set.
func (o *CreateAlmanaxWebhookDailySettings) HasTimezone() bool {
	if o != nil && o.Timezone.IsSet() {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given NullableString and assigns it to the Timezone field.
func (o *CreateAlmanaxWebhookDailySettings) SetTimezone(v string) {
	o.Timezone.Set(&v)
}
// SetTimezoneNil sets the value for Timezone to be an explicit nil
func (o *CreateAlmanaxWebhookDailySettings) SetTimezoneNil() {
	o.Timezone.Set(nil)
}

// UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
func (o *CreateAlmanaxWebhookDailySettings) UnsetTimezone() {
	o.Timezone.Unset()
}

// GetMidnightOffset returns the MidnightOffset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAlmanaxWebhookDailySettings) GetMidnightOffset() int32 {
	if o == nil || IsNil(o.MidnightOffset.Get()) {
		var ret int32
		return ret
	}
	return *o.MidnightOffset.Get()
}

// GetMidnightOffsetOk returns a tuple with the MidnightOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAlmanaxWebhookDailySettings) GetMidnightOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MidnightOffset.Get(), o.MidnightOffset.IsSet()
}

// HasMidnightOffset returns a boolean if a field has been set.
func (o *CreateAlmanaxWebhookDailySettings) HasMidnightOffset() bool {
	if o != nil && o.MidnightOffset.IsSet() {
		return true
	}

	return false
}

// SetMidnightOffset gets a reference to the given NullableInt32 and assigns it to the MidnightOffset field.
func (o *CreateAlmanaxWebhookDailySettings) SetMidnightOffset(v int32) {
	o.MidnightOffset.Set(&v)
}
// SetMidnightOffsetNil sets the value for MidnightOffset to be an explicit nil
func (o *CreateAlmanaxWebhookDailySettings) SetMidnightOffsetNil() {
	o.MidnightOffset.Set(nil)
}

// UnsetMidnightOffset ensures that no value is present for MidnightOffset, not even an explicit nil
func (o *CreateAlmanaxWebhookDailySettings) UnsetMidnightOffset() {
	o.MidnightOffset.Unset()
}

func (o CreateAlmanaxWebhookDailySettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAlmanaxWebhookDailySettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Timezone.IsSet() {
		toSerialize["timezone"] = o.Timezone.Get()
	}
	if o.MidnightOffset.IsSet() {
		toSerialize["midnight_offset"] = o.MidnightOffset.Get()
	}
	return toSerialize, nil
}

type NullableCreateAlmanaxWebhookDailySettings struct {
	value *CreateAlmanaxWebhookDailySettings
	isSet bool
}

func (v NullableCreateAlmanaxWebhookDailySettings) Get() *CreateAlmanaxWebhookDailySettings {
	return v.value
}

func (v *NullableCreateAlmanaxWebhookDailySettings) Set(val *CreateAlmanaxWebhookDailySettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAlmanaxWebhookDailySettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAlmanaxWebhookDailySettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAlmanaxWebhookDailySettings(val *CreateAlmanaxWebhookDailySettings) *NullableCreateAlmanaxWebhookDailySettings {
	return &NullableCreateAlmanaxWebhookDailySettings{value: val, isSet: true}
}

func (v NullableCreateAlmanaxWebhookDailySettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAlmanaxWebhookDailySettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


