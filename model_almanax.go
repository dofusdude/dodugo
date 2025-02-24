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
)

// checks if the Almanax type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Almanax{}

// Almanax struct for Almanax
type Almanax struct {
	Bonus *AlmanaxBonus `json:"bonus,omitempty"`
	Date *string `json:"date,omitempty"`
	Tribute *AlmanaxTribute `json:"tribute,omitempty"`
	// Amount of Kamas you get as reward for finishing this Almanax quest.
	RewardKamas NullableInt32 `json:"reward_kamas,omitempty"`
	// Optional field that shows when a level is given in the request. Shows the experience points you get this day for finishing this Almanax quest.
	RewardXp NullableInt32 `json:"reward_xp,omitempty"`
}

// NewAlmanax instantiates a new Almanax object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlmanax() *Almanax {
	this := Almanax{}
	return &this
}

// NewAlmanaxWithDefaults instantiates a new Almanax object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlmanaxWithDefaults() *Almanax {
	this := Almanax{}
	return &this
}

// GetBonus returns the Bonus field value if set, zero value otherwise.
func (o *Almanax) GetBonus() AlmanaxBonus {
	if o == nil || IsNil(o.Bonus) {
		var ret AlmanaxBonus
		return ret
	}
	return *o.Bonus
}

// GetBonusOk returns a tuple with the Bonus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Almanax) GetBonusOk() (*AlmanaxBonus, bool) {
	if o == nil || IsNil(o.Bonus) {
		return nil, false
	}
	return o.Bonus, true
}

// HasBonus returns a boolean if a field has been set.
func (o *Almanax) HasBonus() bool {
	if o != nil && !IsNil(o.Bonus) {
		return true
	}

	return false
}

// SetBonus gets a reference to the given AlmanaxBonus and assigns it to the Bonus field.
func (o *Almanax) SetBonus(v AlmanaxBonus) {
	o.Bonus = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *Almanax) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Almanax) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *Almanax) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *Almanax) SetDate(v string) {
	o.Date = &v
}

// GetTribute returns the Tribute field value if set, zero value otherwise.
func (o *Almanax) GetTribute() AlmanaxTribute {
	if o == nil || IsNil(o.Tribute) {
		var ret AlmanaxTribute
		return ret
	}
	return *o.Tribute
}

// GetTributeOk returns a tuple with the Tribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Almanax) GetTributeOk() (*AlmanaxTribute, bool) {
	if o == nil || IsNil(o.Tribute) {
		return nil, false
	}
	return o.Tribute, true
}

// HasTribute returns a boolean if a field has been set.
func (o *Almanax) HasTribute() bool {
	if o != nil && !IsNil(o.Tribute) {
		return true
	}

	return false
}

// SetTribute gets a reference to the given AlmanaxTribute and assigns it to the Tribute field.
func (o *Almanax) SetTribute(v AlmanaxTribute) {
	o.Tribute = &v
}

// GetRewardKamas returns the RewardKamas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Almanax) GetRewardKamas() int32 {
	if o == nil || IsNil(o.RewardKamas.Get()) {
		var ret int32
		return ret
	}
	return *o.RewardKamas.Get()
}

// GetRewardKamasOk returns a tuple with the RewardKamas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Almanax) GetRewardKamasOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RewardKamas.Get(), o.RewardKamas.IsSet()
}

// HasRewardKamas returns a boolean if a field has been set.
func (o *Almanax) HasRewardKamas() bool {
	if o != nil && o.RewardKamas.IsSet() {
		return true
	}

	return false
}

// SetRewardKamas gets a reference to the given NullableInt32 and assigns it to the RewardKamas field.
func (o *Almanax) SetRewardKamas(v int32) {
	o.RewardKamas.Set(&v)
}
// SetRewardKamasNil sets the value for RewardKamas to be an explicit nil
func (o *Almanax) SetRewardKamasNil() {
	o.RewardKamas.Set(nil)
}

// UnsetRewardKamas ensures that no value is present for RewardKamas, not even an explicit nil
func (o *Almanax) UnsetRewardKamas() {
	o.RewardKamas.Unset()
}

// GetRewardXp returns the RewardXp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Almanax) GetRewardXp() int32 {
	if o == nil || IsNil(o.RewardXp.Get()) {
		var ret int32
		return ret
	}
	return *o.RewardXp.Get()
}

// GetRewardXpOk returns a tuple with the RewardXp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Almanax) GetRewardXpOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RewardXp.Get(), o.RewardXp.IsSet()
}

// HasRewardXp returns a boolean if a field has been set.
func (o *Almanax) HasRewardXp() bool {
	if o != nil && o.RewardXp.IsSet() {
		return true
	}

	return false
}

// SetRewardXp gets a reference to the given NullableInt32 and assigns it to the RewardXp field.
func (o *Almanax) SetRewardXp(v int32) {
	o.RewardXp.Set(&v)
}
// SetRewardXpNil sets the value for RewardXp to be an explicit nil
func (o *Almanax) SetRewardXpNil() {
	o.RewardXp.Set(nil)
}

// UnsetRewardXp ensures that no value is present for RewardXp, not even an explicit nil
func (o *Almanax) UnsetRewardXp() {
	o.RewardXp.Unset()
}

func (o Almanax) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Almanax) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bonus) {
		toSerialize["bonus"] = o.Bonus
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Tribute) {
		toSerialize["tribute"] = o.Tribute
	}
	if o.RewardKamas.IsSet() {
		toSerialize["reward_kamas"] = o.RewardKamas.Get()
	}
	if o.RewardXp.IsSet() {
		toSerialize["reward_xp"] = o.RewardXp.Get()
	}
	return toSerialize, nil
}

type NullableAlmanax struct {
	value *Almanax
	isSet bool
}

func (v NullableAlmanax) Get() *Almanax {
	return v.value
}

func (v *NullableAlmanax) Set(val *Almanax) {
	v.value = val
	v.isSet = true
}

func (v NullableAlmanax) IsSet() bool {
	return v.isSet
}

func (v *NullableAlmanax) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlmanax(val *Almanax) *NullableAlmanax {
	return &NullableAlmanax{value: val, isSet: true}
}

func (v NullableAlmanax) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlmanax) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


