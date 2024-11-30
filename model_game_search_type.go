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
)

// checks if the GameSearchType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameSearchType{}

// GameSearchType struct for GameSearchType
type GameSearchType struct {
	NameId *string `json:"name_id,omitempty"`
}

// NewGameSearchType instantiates a new GameSearchType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameSearchType() *GameSearchType {
	this := GameSearchType{}
	return &this
}

// NewGameSearchTypeWithDefaults instantiates a new GameSearchType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameSearchTypeWithDefaults() *GameSearchType {
	this := GameSearchType{}
	return &this
}

// GetNameId returns the NameId field value if set, zero value otherwise.
func (o *GameSearchType) GetNameId() string {
	if o == nil || IsNil(o.NameId) {
		var ret string
		return ret
	}
	return *o.NameId
}

// GetNameIdOk returns a tuple with the NameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameSearchType) GetNameIdOk() (*string, bool) {
	if o == nil || IsNil(o.NameId) {
		return nil, false
	}
	return o.NameId, true
}

// HasNameId returns a boolean if a field has been set.
func (o *GameSearchType) HasNameId() bool {
	if o != nil && !IsNil(o.NameId) {
		return true
	}

	return false
}

// SetNameId gets a reference to the given string and assigns it to the NameId field.
func (o *GameSearchType) SetNameId(v string) {
	o.NameId = &v
}

func (o GameSearchType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameSearchType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NameId) {
		toSerialize["name_id"] = o.NameId
	}
	return toSerialize, nil
}

type NullableGameSearchType struct {
	value *GameSearchType
	isSet bool
}

func (v NullableGameSearchType) Get() *GameSearchType {
	return v.value
}

func (v *NullableGameSearchType) Set(val *GameSearchType) {
	v.value = val
	v.isSet = true
}

func (v NullableGameSearchType) IsSet() bool {
	return v.isSet
}

func (v *NullableGameSearchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameSearchType(val *GameSearchType) *NullableGameSearchType {
	return &NullableGameSearchType{value: val, isSet: true}
}

func (v NullableGameSearchType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameSearchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


