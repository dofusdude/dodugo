/*
dofusdude

# Open Ankama Developer Community The all-in-one toolbelt for your next Ankama related project.  ## Versions - [Dofus 2](https://docs.dofusdu.de/dofus2/) - [Dofus 3](https://docs.dofusdu.de/dofus3/)   - v1 [latest] (you are here)   ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save` - [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save` - [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo` - [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude` - [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Almanax Discord Integration** Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 3 Beta** from stable to bleeding edge by replacing /dofus3 with /dofus3beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_, _de_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Official Sources** generated from actual data from the game.  ... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h). 

API version: 1.0.0-rc.3
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
	"time"
)

// checks if the TwitterWebhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TwitterWebhook{}

// TwitterWebhook struct for TwitterWebhook
type TwitterWebhook struct {
	Id *string `json:"id,omitempty"`
	Whitelist []string `json:"whitelist,omitempty"`
	Blacklist []string `json:"blacklist,omitempty"`
	Subscriptions []string `json:"subscriptions,omitempty"`
	Format *string `json:"format,omitempty"`
	PreviewLength *int32 `json:"preview_length,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	LastFiredAt NullableTime `json:"last_fired_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewTwitterWebhook instantiates a new TwitterWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTwitterWebhook() *TwitterWebhook {
	this := TwitterWebhook{}
	return &this
}

// NewTwitterWebhookWithDefaults instantiates a new TwitterWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTwitterWebhookWithDefaults() *TwitterWebhook {
	this := TwitterWebhook{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TwitterWebhook) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwitterWebhook) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TwitterWebhook) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TwitterWebhook) SetId(v string) {
	o.Id = &v
}

// GetWhitelist returns the Whitelist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TwitterWebhook) GetWhitelist() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Whitelist
}

// GetWhitelistOk returns a tuple with the Whitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TwitterWebhook) GetWhitelistOk() ([]string, bool) {
	if o == nil || IsNil(o.Whitelist) {
		return nil, false
	}
	return o.Whitelist, true
}

// HasWhitelist returns a boolean if a field has been set.
func (o *TwitterWebhook) HasWhitelist() bool {
	if o != nil && !IsNil(o.Whitelist) {
		return true
	}

	return false
}

// SetWhitelist gets a reference to the given []string and assigns it to the Whitelist field.
func (o *TwitterWebhook) SetWhitelist(v []string) {
	o.Whitelist = v
}

// GetBlacklist returns the Blacklist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TwitterWebhook) GetBlacklist() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Blacklist
}

// GetBlacklistOk returns a tuple with the Blacklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TwitterWebhook) GetBlacklistOk() ([]string, bool) {
	if o == nil || IsNil(o.Blacklist) {
		return nil, false
	}
	return o.Blacklist, true
}

// HasBlacklist returns a boolean if a field has been set.
func (o *TwitterWebhook) HasBlacklist() bool {
	if o != nil && !IsNil(o.Blacklist) {
		return true
	}

	return false
}

// SetBlacklist gets a reference to the given []string and assigns it to the Blacklist field.
func (o *TwitterWebhook) SetBlacklist(v []string) {
	o.Blacklist = v
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise.
func (o *TwitterWebhook) GetSubscriptions() []string {
	if o == nil || IsNil(o.Subscriptions) {
		var ret []string
		return ret
	}
	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwitterWebhook) GetSubscriptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Subscriptions) {
		return nil, false
	}
	return o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *TwitterWebhook) HasSubscriptions() bool {
	if o != nil && !IsNil(o.Subscriptions) {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given []string and assigns it to the Subscriptions field.
func (o *TwitterWebhook) SetSubscriptions(v []string) {
	o.Subscriptions = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *TwitterWebhook) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwitterWebhook) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *TwitterWebhook) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *TwitterWebhook) SetFormat(v string) {
	o.Format = &v
}

// GetPreviewLength returns the PreviewLength field value if set, zero value otherwise.
func (o *TwitterWebhook) GetPreviewLength() int32 {
	if o == nil || IsNil(o.PreviewLength) {
		var ret int32
		return ret
	}
	return *o.PreviewLength
}

// GetPreviewLengthOk returns a tuple with the PreviewLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwitterWebhook) GetPreviewLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.PreviewLength) {
		return nil, false
	}
	return o.PreviewLength, true
}

// HasPreviewLength returns a boolean if a field has been set.
func (o *TwitterWebhook) HasPreviewLength() bool {
	if o != nil && !IsNil(o.PreviewLength) {
		return true
	}

	return false
}

// SetPreviewLength gets a reference to the given int32 and assigns it to the PreviewLength field.
func (o *TwitterWebhook) SetPreviewLength(v int32) {
	o.PreviewLength = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TwitterWebhook) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwitterWebhook) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TwitterWebhook) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TwitterWebhook) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastFiredAt returns the LastFiredAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TwitterWebhook) GetLastFiredAt() time.Time {
	if o == nil || IsNil(o.LastFiredAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastFiredAt.Get()
}

// GetLastFiredAtOk returns a tuple with the LastFiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TwitterWebhook) GetLastFiredAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastFiredAt.Get(), o.LastFiredAt.IsSet()
}

// HasLastFiredAt returns a boolean if a field has been set.
func (o *TwitterWebhook) HasLastFiredAt() bool {
	if o != nil && o.LastFiredAt.IsSet() {
		return true
	}

	return false
}

// SetLastFiredAt gets a reference to the given NullableTime and assigns it to the LastFiredAt field.
func (o *TwitterWebhook) SetLastFiredAt(v time.Time) {
	o.LastFiredAt.Set(&v)
}
// SetLastFiredAtNil sets the value for LastFiredAt to be an explicit nil
func (o *TwitterWebhook) SetLastFiredAtNil() {
	o.LastFiredAt.Set(nil)
}

// UnsetLastFiredAt ensures that no value is present for LastFiredAt, not even an explicit nil
func (o *TwitterWebhook) UnsetLastFiredAt() {
	o.LastFiredAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TwitterWebhook) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwitterWebhook) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TwitterWebhook) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *TwitterWebhook) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o TwitterWebhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TwitterWebhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Whitelist != nil {
		toSerialize["whitelist"] = o.Whitelist
	}
	if o.Blacklist != nil {
		toSerialize["blacklist"] = o.Blacklist
	}
	if !IsNil(o.Subscriptions) {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.PreviewLength) {
		toSerialize["preview_length"] = o.PreviewLength
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

type NullableTwitterWebhook struct {
	value *TwitterWebhook
	isSet bool
}

func (v NullableTwitterWebhook) Get() *TwitterWebhook {
	return v.value
}

func (v *NullableTwitterWebhook) Set(val *TwitterWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableTwitterWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableTwitterWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTwitterWebhook(val *TwitterWebhook) *NullableTwitterWebhook {
	return &NullableTwitterWebhook{value: val, isSet: true}
}

func (v NullableTwitterWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTwitterWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


