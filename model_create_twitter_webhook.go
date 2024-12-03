/*
dofusdude

# Open Ankama Developer Community The all-in-one toolbelt for your next Ankama related project.  ## Versions - [Dofus 2](https://docs.dofusdu.de/dofus2/) - [Dofus 3](https://docs.dofusdu.de/dofus3/)   - v1 [latest] (you are here)   ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save` - [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save` - [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo` - [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude` - [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Almanax Discord Integration** Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 3 Beta** from stable to bleeding edge by replacing /dofus3 with /dofus3beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_, _de_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Official Sources** generated from actual data from the game.  ... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h). 

API version: 1.0.0-rc.8
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateTwitterWebhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTwitterWebhook{}

// CreateTwitterWebhook 
type CreateTwitterWebhook struct {
	Whitelist []string `json:"whitelist,omitempty"`
	Blacklist []string `json:"blacklist,omitempty"`
	// Get the available subscriptions with /meta/webhooks/twitter
	Subscriptions []string `json:"subscriptions"`
	Format string `json:"format"`
	PreviewLength NullableInt32 `json:"preview_length,omitempty"`
	// Discord Webhook URL
	Callback string `json:"callback"`
}

type _CreateTwitterWebhook CreateTwitterWebhook

// NewCreateTwitterWebhook instantiates a new CreateTwitterWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTwitterWebhook(subscriptions []string, format string, callback string) *CreateTwitterWebhook {
	this := CreateTwitterWebhook{}
	this.Subscriptions = subscriptions
	this.Format = format
	this.Callback = callback
	return &this
}

// NewCreateTwitterWebhookWithDefaults instantiates a new CreateTwitterWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTwitterWebhookWithDefaults() *CreateTwitterWebhook {
	this := CreateTwitterWebhook{}
	return &this
}

// GetWhitelist returns the Whitelist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTwitterWebhook) GetWhitelist() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Whitelist
}

// GetWhitelistOk returns a tuple with the Whitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTwitterWebhook) GetWhitelistOk() ([]string, bool) {
	if o == nil || IsNil(o.Whitelist) {
		return nil, false
	}
	return o.Whitelist, true
}

// HasWhitelist returns a boolean if a field has been set.
func (o *CreateTwitterWebhook) HasWhitelist() bool {
	if o != nil && !IsNil(o.Whitelist) {
		return true
	}

	return false
}

// SetWhitelist gets a reference to the given []string and assigns it to the Whitelist field.
func (o *CreateTwitterWebhook) SetWhitelist(v []string) {
	o.Whitelist = v
}

// GetBlacklist returns the Blacklist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTwitterWebhook) GetBlacklist() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Blacklist
}

// GetBlacklistOk returns a tuple with the Blacklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTwitterWebhook) GetBlacklistOk() ([]string, bool) {
	if o == nil || IsNil(o.Blacklist) {
		return nil, false
	}
	return o.Blacklist, true
}

// HasBlacklist returns a boolean if a field has been set.
func (o *CreateTwitterWebhook) HasBlacklist() bool {
	if o != nil && !IsNil(o.Blacklist) {
		return true
	}

	return false
}

// SetBlacklist gets a reference to the given []string and assigns it to the Blacklist field.
func (o *CreateTwitterWebhook) SetBlacklist(v []string) {
	o.Blacklist = v
}

// GetSubscriptions returns the Subscriptions field value
func (o *CreateTwitterWebhook) GetSubscriptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value
// and a boolean to check if the value has been set.
func (o *CreateTwitterWebhook) GetSubscriptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// SetSubscriptions sets field value
func (o *CreateTwitterWebhook) SetSubscriptions(v []string) {
	o.Subscriptions = v
}

// GetFormat returns the Format field value
func (o *CreateTwitterWebhook) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *CreateTwitterWebhook) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *CreateTwitterWebhook) SetFormat(v string) {
	o.Format = v
}

// GetPreviewLength returns the PreviewLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTwitterWebhook) GetPreviewLength() int32 {
	if o == nil || IsNil(o.PreviewLength.Get()) {
		var ret int32
		return ret
	}
	return *o.PreviewLength.Get()
}

// GetPreviewLengthOk returns a tuple with the PreviewLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTwitterWebhook) GetPreviewLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviewLength.Get(), o.PreviewLength.IsSet()
}

// HasPreviewLength returns a boolean if a field has been set.
func (o *CreateTwitterWebhook) HasPreviewLength() bool {
	if o != nil && o.PreviewLength.IsSet() {
		return true
	}

	return false
}

// SetPreviewLength gets a reference to the given NullableInt32 and assigns it to the PreviewLength field.
func (o *CreateTwitterWebhook) SetPreviewLength(v int32) {
	o.PreviewLength.Set(&v)
}
// SetPreviewLengthNil sets the value for PreviewLength to be an explicit nil
func (o *CreateTwitterWebhook) SetPreviewLengthNil() {
	o.PreviewLength.Set(nil)
}

// UnsetPreviewLength ensures that no value is present for PreviewLength, not even an explicit nil
func (o *CreateTwitterWebhook) UnsetPreviewLength() {
	o.PreviewLength.Unset()
}

// GetCallback returns the Callback field value
func (o *CreateTwitterWebhook) GetCallback() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value
// and a boolean to check if the value has been set.
func (o *CreateTwitterWebhook) GetCallbackOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Callback, true
}

// SetCallback sets field value
func (o *CreateTwitterWebhook) SetCallback(v string) {
	o.Callback = v
}

func (o CreateTwitterWebhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTwitterWebhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Whitelist != nil {
		toSerialize["whitelist"] = o.Whitelist
	}
	if o.Blacklist != nil {
		toSerialize["blacklist"] = o.Blacklist
	}
	toSerialize["subscriptions"] = o.Subscriptions
	toSerialize["format"] = o.Format
	if o.PreviewLength.IsSet() {
		toSerialize["preview_length"] = o.PreviewLength.Get()
	}
	toSerialize["callback"] = o.Callback
	return toSerialize, nil
}

func (o *CreateTwitterWebhook) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subscriptions",
		"format",
		"callback",
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

	varCreateTwitterWebhook := _CreateTwitterWebhook{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTwitterWebhook)

	if err != nil {
		return err
	}

	*o = CreateTwitterWebhook(varCreateTwitterWebhook)

	return err
}

type NullableCreateTwitterWebhook struct {
	value *CreateTwitterWebhook
	isSet bool
}

func (v NullableCreateTwitterWebhook) Get() *CreateTwitterWebhook {
	return v.value
}

func (v *NullableCreateTwitterWebhook) Set(val *CreateTwitterWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTwitterWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTwitterWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTwitterWebhook(val *CreateTwitterWebhook) *NullableCreateTwitterWebhook {
	return &NullableCreateTwitterWebhook{value: val, isSet: true}
}

func (v NullableCreateTwitterWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTwitterWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


