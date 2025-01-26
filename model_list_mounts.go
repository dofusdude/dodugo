/*
dofusdude

# Open Ankama Developer Community The all-in-one toolbelt for your next Ankama related project.  ## Versions - [Dofus 2](https://docs.dofusdu.de/dofus2/) - [Dofus 3](https://docs.dofusdu.de/dofus3/)   - v1 [latest] (you are here)   ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save` - [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save` - [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo` - [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude` - [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Almanax Discord Integration** Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 3 Beta** from stable to bleeding edge by replacing /dofus3 with /dofus3beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_, _de_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Official Sources** generated from actual data from the game.  ... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h). 

API version: 1.0.0-rc.10
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
)

// checks if the ListMounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMounts{}

// ListMounts struct for ListMounts
type ListMounts struct {
	Links *PagedLinks `json:"_links,omitempty"`
	Items []Mount `json:"items,omitempty"`
}

// NewListMounts instantiates a new ListMounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMounts() *ListMounts {
	this := ListMounts{}
	return &this
}

// NewListMountsWithDefaults instantiates a new ListMounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMountsWithDefaults() *ListMounts {
	this := ListMounts{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListMounts) GetLinks() PagedLinks {
	if o == nil || IsNil(o.Links) {
		var ret PagedLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMounts) GetLinksOk() (*PagedLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListMounts) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PagedLinks and assigns it to the Links field.
func (o *ListMounts) SetLinks(v PagedLinks) {
	o.Links = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListMounts) GetItems() []Mount {
	if o == nil || IsNil(o.Items) {
		var ret []Mount
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMounts) GetItemsOk() ([]Mount, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListMounts) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Mount and assigns it to the Items field.
func (o *ListMounts) SetItems(v []Mount) {
	o.Items = v
}

func (o ListMounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableListMounts struct {
	value *ListMounts
	isSet bool
}

func (v NullableListMounts) Get() *ListMounts {
	return v.value
}

func (v *NullableListMounts) Set(val *ListMounts) {
	v.value = val
	v.isSet = true
}

func (v NullableListMounts) IsSet() bool {
	return v.isSet
}

func (v *NullableListMounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMounts(val *ListMounts) *NullableListMounts {
	return &NullableListMounts{value: val, isSet: true}
}

func (v NullableListMounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


