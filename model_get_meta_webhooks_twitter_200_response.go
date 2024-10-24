/*
dofusdude

# A project for you - the developer. The all-in-one toolbelt for your next Ankama related project.  ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save` - [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save` - [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo` - [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude` - [PHP](https://github.com/dofusdude/dofusdude-php) - [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Discord Integration** Ankama related RSS and Almanax feeds to post to Discord servers with advanced features like filters or mentions. Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 2 Beta** from stable to bleeding edge by replacing /dofus2 with /dofus2beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_ including the dropped languages from the Dofus website _de_ and _it_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Complete** actual data from the game including items invisible to the encyclopedia like quest items.  - 🖼️ **HD Images** rendering game assets to high-res images with up to 800x800 px.  ... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h). 

API version: 0.9.4
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
)

// checks if the GetMetaWebhooksTwitter200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMetaWebhooksTwitter200Response{}

// GetMetaWebhooksTwitter200Response struct for GetMetaWebhooksTwitter200Response
type GetMetaWebhooksTwitter200Response struct {
	Subscriptions []string `json:"subscriptions,omitempty"`
}

// NewGetMetaWebhooksTwitter200Response instantiates a new GetMetaWebhooksTwitter200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMetaWebhooksTwitter200Response() *GetMetaWebhooksTwitter200Response {
	this := GetMetaWebhooksTwitter200Response{}
	return &this
}

// NewGetMetaWebhooksTwitter200ResponseWithDefaults instantiates a new GetMetaWebhooksTwitter200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMetaWebhooksTwitter200ResponseWithDefaults() *GetMetaWebhooksTwitter200Response {
	this := GetMetaWebhooksTwitter200Response{}
	return &this
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise.
func (o *GetMetaWebhooksTwitter200Response) GetSubscriptions() []string {
	if o == nil || IsNil(o.Subscriptions) {
		var ret []string
		return ret
	}
	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMetaWebhooksTwitter200Response) GetSubscriptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Subscriptions) {
		return nil, false
	}
	return o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *GetMetaWebhooksTwitter200Response) HasSubscriptions() bool {
	if o != nil && !IsNil(o.Subscriptions) {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given []string and assigns it to the Subscriptions field.
func (o *GetMetaWebhooksTwitter200Response) SetSubscriptions(v []string) {
	o.Subscriptions = v
}

func (o GetMetaWebhooksTwitter200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMetaWebhooksTwitter200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subscriptions) {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	return toSerialize, nil
}

type NullableGetMetaWebhooksTwitter200Response struct {
	value *GetMetaWebhooksTwitter200Response
	isSet bool
}

func (v NullableGetMetaWebhooksTwitter200Response) Get() *GetMetaWebhooksTwitter200Response {
	return v.value
}

func (v *NullableGetMetaWebhooksTwitter200Response) Set(val *GetMetaWebhooksTwitter200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMetaWebhooksTwitter200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMetaWebhooksTwitter200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMetaWebhooksTwitter200Response(val *GetMetaWebhooksTwitter200Response) *NullableGetMetaWebhooksTwitter200Response {
	return &NullableGetMetaWebhooksTwitter200Response{value: val, isSet: true}
}

func (v NullableGetMetaWebhooksTwitter200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMetaWebhooksTwitter200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


