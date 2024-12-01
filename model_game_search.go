/*
dofusdude

# Open Ankama Developer Community The all-in-one toolbelt for your next Ankama related project.  ## Versions - [Dofus 2](https://docs.dofusdu.de/dofus2/) - [Dofus 3](https://docs.dofusdu.de/dofus3/)   - v1 [latest] (you are here)   ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save` - [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save` - [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo` - [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude` - [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Almanax Discord Integration** Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 3 Beta** from stable to bleeding edge by replacing /dofus3 with /dofus3beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_, _de_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Official Sources** generated from actual data from the game.  ... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h). 

API version: 1.0.0-rc.7
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
)

// checks if the GameSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameSearch{}

// GameSearch struct for GameSearch
type GameSearch struct {
	AnkamaId *int32 `json:"ankama_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *GameSearchType `json:"type,omitempty"`
	ItemFields *GameSearchItem `json:"item_fields,omitempty"`
}

// NewGameSearch instantiates a new GameSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameSearch() *GameSearch {
	this := GameSearch{}
	return &this
}

// NewGameSearchWithDefaults instantiates a new GameSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameSearchWithDefaults() *GameSearch {
	this := GameSearch{}
	return &this
}

// GetAnkamaId returns the AnkamaId field value if set, zero value otherwise.
func (o *GameSearch) GetAnkamaId() int32 {
	if o == nil || IsNil(o.AnkamaId) {
		var ret int32
		return ret
	}
	return *o.AnkamaId
}

// GetAnkamaIdOk returns a tuple with the AnkamaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameSearch) GetAnkamaIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AnkamaId) {
		return nil, false
	}
	return o.AnkamaId, true
}

// HasAnkamaId returns a boolean if a field has been set.
func (o *GameSearch) HasAnkamaId() bool {
	if o != nil && !IsNil(o.AnkamaId) {
		return true
	}

	return false
}

// SetAnkamaId gets a reference to the given int32 and assigns it to the AnkamaId field.
func (o *GameSearch) SetAnkamaId(v int32) {
	o.AnkamaId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GameSearch) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameSearch) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GameSearch) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GameSearch) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GameSearch) GetType() GameSearchType {
	if o == nil || IsNil(o.Type) {
		var ret GameSearchType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameSearch) GetTypeOk() (*GameSearchType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GameSearch) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given GameSearchType and assigns it to the Type field.
func (o *GameSearch) SetType(v GameSearchType) {
	o.Type = &v
}

// GetItemFields returns the ItemFields field value if set, zero value otherwise.
func (o *GameSearch) GetItemFields() GameSearchItem {
	if o == nil || IsNil(o.ItemFields) {
		var ret GameSearchItem
		return ret
	}
	return *o.ItemFields
}

// GetItemFieldsOk returns a tuple with the ItemFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameSearch) GetItemFieldsOk() (*GameSearchItem, bool) {
	if o == nil || IsNil(o.ItemFields) {
		return nil, false
	}
	return o.ItemFields, true
}

// HasItemFields returns a boolean if a field has been set.
func (o *GameSearch) HasItemFields() bool {
	if o != nil && !IsNil(o.ItemFields) {
		return true
	}

	return false
}

// SetItemFields gets a reference to the given GameSearchItem and assigns it to the ItemFields field.
func (o *GameSearch) SetItemFields(v GameSearchItem) {
	o.ItemFields = &v
}

func (o GameSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnkamaId) {
		toSerialize["ankama_id"] = o.AnkamaId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ItemFields) {
		toSerialize["item_fields"] = o.ItemFields
	}
	return toSerialize, nil
}

type NullableGameSearch struct {
	value *GameSearch
	isSet bool
}

func (v NullableGameSearch) Get() *GameSearch {
	return v.value
}

func (v *NullableGameSearch) Set(val *GameSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableGameSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableGameSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameSearch(val *GameSearch) *NullableGameSearch {
	return &NullableGameSearch{value: val, isSet: true}
}

func (v NullableGameSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


