/*
dofusdude

# A project for you - the developer. The all-in-one toolbelt for your next Ankama related project.  ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save` - [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save` - [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo` - [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude` - [PHP](https://github.com/dofusdude/dofusdude-php) - [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Discord Integration** Ankama related RSS and Almanax feeds to post to Discord servers with advanced features like filters or mentions. Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 2 Beta** from stable to bleeding edge by replacing /dofus2 with /dofus2beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_ including the dropped languages from the Dofus website _de_ and _it_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Complete** actual data from the game including items invisible to the encyclopedia like quest items.  - 🖼️ **HD Images** rendering game assets to high-res images with up to 800x800 px.  ... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h). 

API version: 0.9.1
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
)

// checks if the PutTwitterWebhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutTwitterWebhook{}

// PutTwitterWebhook 
type PutTwitterWebhook struct {
	Whitelist []string `json:"whitelist,omitempty"`
	Blacklist []string `json:"blacklist,omitempty"`
	// Get the available subscriptions with /meta/webhooks/twitter
	Subscriptions []string `json:"subscriptions,omitempty"`
	PreviewLength NullableInt32 `json:"preview_length,omitempty"`
}

// NewPutTwitterWebhook instantiates a new PutTwitterWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutTwitterWebhook() *PutTwitterWebhook {
	this := PutTwitterWebhook{}
	return &this
}

// NewPutTwitterWebhookWithDefaults instantiates a new PutTwitterWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutTwitterWebhookWithDefaults() *PutTwitterWebhook {
	this := PutTwitterWebhook{}
	return &this
}

// GetWhitelist returns the Whitelist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PutTwitterWebhook) GetWhitelist() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Whitelist
}

// GetWhitelistOk returns a tuple with the Whitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutTwitterWebhook) GetWhitelistOk() ([]string, bool) {
	if o == nil || IsNil(o.Whitelist) {
		return nil, false
	}
	return o.Whitelist, true
}

// HasWhitelist returns a boolean if a field has been set.
func (o *PutTwitterWebhook) HasWhitelist() bool {
	if o != nil && !IsNil(o.Whitelist) {
		return true
	}

	return false
}

// SetWhitelist gets a reference to the given []string and assigns it to the Whitelist field.
func (o *PutTwitterWebhook) SetWhitelist(v []string) {
	o.Whitelist = v
}

// GetBlacklist returns the Blacklist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PutTwitterWebhook) GetBlacklist() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Blacklist
}

// GetBlacklistOk returns a tuple with the Blacklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutTwitterWebhook) GetBlacklistOk() ([]string, bool) {
	if o == nil || IsNil(o.Blacklist) {
		return nil, false
	}
	return o.Blacklist, true
}

// HasBlacklist returns a boolean if a field has been set.
func (o *PutTwitterWebhook) HasBlacklist() bool {
	if o != nil && !IsNil(o.Blacklist) {
		return true
	}

	return false
}

// SetBlacklist gets a reference to the given []string and assigns it to the Blacklist field.
func (o *PutTwitterWebhook) SetBlacklist(v []string) {
	o.Blacklist = v
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PutTwitterWebhook) GetSubscriptions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutTwitterWebhook) GetSubscriptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Subscriptions) {
		return nil, false
	}
	return o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *PutTwitterWebhook) HasSubscriptions() bool {
	if o != nil && !IsNil(o.Subscriptions) {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given []string and assigns it to the Subscriptions field.
func (o *PutTwitterWebhook) SetSubscriptions(v []string) {
	o.Subscriptions = v
}

// GetPreviewLength returns the PreviewLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PutTwitterWebhook) GetPreviewLength() int32 {
	if o == nil || IsNil(o.PreviewLength.Get()) {
		var ret int32
		return ret
	}
	return *o.PreviewLength.Get()
}

// GetPreviewLengthOk returns a tuple with the PreviewLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutTwitterWebhook) GetPreviewLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviewLength.Get(), o.PreviewLength.IsSet()
}

// HasPreviewLength returns a boolean if a field has been set.
func (o *PutTwitterWebhook) HasPreviewLength() bool {
	if o != nil && o.PreviewLength.IsSet() {
		return true
	}

	return false
}

// SetPreviewLength gets a reference to the given NullableInt32 and assigns it to the PreviewLength field.
func (o *PutTwitterWebhook) SetPreviewLength(v int32) {
	o.PreviewLength.Set(&v)
}
// SetPreviewLengthNil sets the value for PreviewLength to be an explicit nil
func (o *PutTwitterWebhook) SetPreviewLengthNil() {
	o.PreviewLength.Set(nil)
}

// UnsetPreviewLength ensures that no value is present for PreviewLength, not even an explicit nil
func (o *PutTwitterWebhook) UnsetPreviewLength() {
	o.PreviewLength.Unset()
}

func (o PutTwitterWebhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutTwitterWebhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Whitelist != nil {
		toSerialize["whitelist"] = o.Whitelist
	}
	if o.Blacklist != nil {
		toSerialize["blacklist"] = o.Blacklist
	}
	if o.Subscriptions != nil {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	if o.PreviewLength.IsSet() {
		toSerialize["preview_length"] = o.PreviewLength.Get()
	}
	return toSerialize, nil
}

type NullablePutTwitterWebhook struct {
	value *PutTwitterWebhook
	isSet bool
}

func (v NullablePutTwitterWebhook) Get() *PutTwitterWebhook {
	return v.value
}

func (v *NullablePutTwitterWebhook) Set(val *PutTwitterWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullablePutTwitterWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullablePutTwitterWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutTwitterWebhook(val *PutTwitterWebhook) *NullablePutTwitterWebhook {
	return &NullablePutTwitterWebhook{value: val, isSet: true}
}

func (v NullablePutTwitterWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutTwitterWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


