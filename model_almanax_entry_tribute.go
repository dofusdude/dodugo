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

// checks if the AlmanaxEntryTribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlmanaxEntryTribute{}

// AlmanaxEntryTribute struct for AlmanaxEntryTribute
type AlmanaxEntryTribute struct {
	Item *AlmanaxEntryTributeItem `json:"item,omitempty"`
	Quantity *int32 `json:"quantity,omitempty"`
}

// NewAlmanaxEntryTribute instantiates a new AlmanaxEntryTribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlmanaxEntryTribute() *AlmanaxEntryTribute {
	this := AlmanaxEntryTribute{}
	return &this
}

// NewAlmanaxEntryTributeWithDefaults instantiates a new AlmanaxEntryTribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlmanaxEntryTributeWithDefaults() *AlmanaxEntryTribute {
	this := AlmanaxEntryTribute{}
	return &this
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *AlmanaxEntryTribute) GetItem() AlmanaxEntryTributeItem {
	if o == nil || IsNil(o.Item) {
		var ret AlmanaxEntryTributeItem
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxEntryTribute) GetItemOk() (*AlmanaxEntryTributeItem, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *AlmanaxEntryTribute) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given AlmanaxEntryTributeItem and assigns it to the Item field.
func (o *AlmanaxEntryTribute) SetItem(v AlmanaxEntryTributeItem) {
	o.Item = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *AlmanaxEntryTribute) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxEntryTribute) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *AlmanaxEntryTribute) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *AlmanaxEntryTribute) SetQuantity(v int32) {
	o.Quantity = &v
}

func (o AlmanaxEntryTribute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlmanaxEntryTribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Item) {
		toSerialize["item"] = o.Item
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullableAlmanaxEntryTribute struct {
	value *AlmanaxEntryTribute
	isSet bool
}

func (v NullableAlmanaxEntryTribute) Get() *AlmanaxEntryTribute {
	return v.value
}

func (v *NullableAlmanaxEntryTribute) Set(val *AlmanaxEntryTribute) {
	v.value = val
	v.isSet = true
}

func (v NullableAlmanaxEntryTribute) IsSet() bool {
	return v.isSet
}

func (v *NullableAlmanaxEntryTribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlmanaxEntryTribute(val *AlmanaxEntryTribute) *NullableAlmanaxEntryTribute {
	return &NullableAlmanaxEntryTribute{value: val, isSet: true}
}

func (v NullableAlmanaxEntryTribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlmanaxEntryTribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


