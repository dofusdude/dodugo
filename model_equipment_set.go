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

// checks if the EquipmentSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EquipmentSet{}

// EquipmentSet struct for EquipmentSet
type EquipmentSet struct {
	AnkamaId *int32 `json:"ankama_id,omitempty"`
	Name *string `json:"name,omitempty"`
	EquipmentIds []int32 `json:"equipment_ids,omitempty"`
	Effects *map[string][]Effect `json:"effects,omitempty"`
	HighestEquipmentLevel *int32 `json:"highest_equipment_level,omitempty"`
	ContainsCosmetics *bool `json:"contains_cosmetics,omitempty"`
	ContainsCosmeticsOnly *bool `json:"contains_cosmetics_only,omitempty"`
}

// NewEquipmentSet instantiates a new EquipmentSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentSet() *EquipmentSet {
	this := EquipmentSet{}
	return &this
}

// NewEquipmentSetWithDefaults instantiates a new EquipmentSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentSetWithDefaults() *EquipmentSet {
	this := EquipmentSet{}
	return &this
}

// GetAnkamaId returns the AnkamaId field value if set, zero value otherwise.
func (o *EquipmentSet) GetAnkamaId() int32 {
	if o == nil || IsNil(o.AnkamaId) {
		var ret int32
		return ret
	}
	return *o.AnkamaId
}

// GetAnkamaIdOk returns a tuple with the AnkamaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSet) GetAnkamaIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AnkamaId) {
		return nil, false
	}
	return o.AnkamaId, true
}

// HasAnkamaId returns a boolean if a field has been set.
func (o *EquipmentSet) HasAnkamaId() bool {
	if o != nil && !IsNil(o.AnkamaId) {
		return true
	}

	return false
}

// SetAnkamaId gets a reference to the given int32 and assigns it to the AnkamaId field.
func (o *EquipmentSet) SetAnkamaId(v int32) {
	o.AnkamaId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EquipmentSet) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSet) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EquipmentSet) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EquipmentSet) SetName(v string) {
	o.Name = &v
}

// GetEquipmentIds returns the EquipmentIds field value if set, zero value otherwise.
func (o *EquipmentSet) GetEquipmentIds() []int32 {
	if o == nil || IsNil(o.EquipmentIds) {
		var ret []int32
		return ret
	}
	return o.EquipmentIds
}

// GetEquipmentIdsOk returns a tuple with the EquipmentIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSet) GetEquipmentIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.EquipmentIds) {
		return nil, false
	}
	return o.EquipmentIds, true
}

// HasEquipmentIds returns a boolean if a field has been set.
func (o *EquipmentSet) HasEquipmentIds() bool {
	if o != nil && !IsNil(o.EquipmentIds) {
		return true
	}

	return false
}

// SetEquipmentIds gets a reference to the given []int32 and assigns it to the EquipmentIds field.
func (o *EquipmentSet) SetEquipmentIds(v []int32) {
	o.EquipmentIds = v
}

// GetEffects returns the Effects field value if set, zero value otherwise.
func (o *EquipmentSet) GetEffects() map[string][]Effect {
	if o == nil || IsNil(o.Effects) {
		var ret map[string][]Effect
		return ret
	}
	return *o.Effects
}

// GetEffectsOk returns a tuple with the Effects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSet) GetEffectsOk() (*map[string][]Effect, bool) {
	if o == nil || IsNil(o.Effects) {
		return nil, false
	}
	return o.Effects, true
}

// HasEffects returns a boolean if a field has been set.
func (o *EquipmentSet) HasEffects() bool {
	if o != nil && !IsNil(o.Effects) {
		return true
	}

	return false
}

// SetEffects gets a reference to the given map[string][]Effect and assigns it to the Effects field.
func (o *EquipmentSet) SetEffects(v map[string][]Effect) {
	o.Effects = &v
}

// GetHighestEquipmentLevel returns the HighestEquipmentLevel field value if set, zero value otherwise.
func (o *EquipmentSet) GetHighestEquipmentLevel() int32 {
	if o == nil || IsNil(o.HighestEquipmentLevel) {
		var ret int32
		return ret
	}
	return *o.HighestEquipmentLevel
}

// GetHighestEquipmentLevelOk returns a tuple with the HighestEquipmentLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSet) GetHighestEquipmentLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.HighestEquipmentLevel) {
		return nil, false
	}
	return o.HighestEquipmentLevel, true
}

// HasHighestEquipmentLevel returns a boolean if a field has been set.
func (o *EquipmentSet) HasHighestEquipmentLevel() bool {
	if o != nil && !IsNil(o.HighestEquipmentLevel) {
		return true
	}

	return false
}

// SetHighestEquipmentLevel gets a reference to the given int32 and assigns it to the HighestEquipmentLevel field.
func (o *EquipmentSet) SetHighestEquipmentLevel(v int32) {
	o.HighestEquipmentLevel = &v
}

// GetContainsCosmetics returns the ContainsCosmetics field value if set, zero value otherwise.
func (o *EquipmentSet) GetContainsCosmetics() bool {
	if o == nil || IsNil(o.ContainsCosmetics) {
		var ret bool
		return ret
	}
	return *o.ContainsCosmetics
}

// GetContainsCosmeticsOk returns a tuple with the ContainsCosmetics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSet) GetContainsCosmeticsOk() (*bool, bool) {
	if o == nil || IsNil(o.ContainsCosmetics) {
		return nil, false
	}
	return o.ContainsCosmetics, true
}

// HasContainsCosmetics returns a boolean if a field has been set.
func (o *EquipmentSet) HasContainsCosmetics() bool {
	if o != nil && !IsNil(o.ContainsCosmetics) {
		return true
	}

	return false
}

// SetContainsCosmetics gets a reference to the given bool and assigns it to the ContainsCosmetics field.
func (o *EquipmentSet) SetContainsCosmetics(v bool) {
	o.ContainsCosmetics = &v
}

// GetContainsCosmeticsOnly returns the ContainsCosmeticsOnly field value if set, zero value otherwise.
func (o *EquipmentSet) GetContainsCosmeticsOnly() bool {
	if o == nil || IsNil(o.ContainsCosmeticsOnly) {
		var ret bool
		return ret
	}
	return *o.ContainsCosmeticsOnly
}

// GetContainsCosmeticsOnlyOk returns a tuple with the ContainsCosmeticsOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSet) GetContainsCosmeticsOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ContainsCosmeticsOnly) {
		return nil, false
	}
	return o.ContainsCosmeticsOnly, true
}

// HasContainsCosmeticsOnly returns a boolean if a field has been set.
func (o *EquipmentSet) HasContainsCosmeticsOnly() bool {
	if o != nil && !IsNil(o.ContainsCosmeticsOnly) {
		return true
	}

	return false
}

// SetContainsCosmeticsOnly gets a reference to the given bool and assigns it to the ContainsCosmeticsOnly field.
func (o *EquipmentSet) SetContainsCosmeticsOnly(v bool) {
	o.ContainsCosmeticsOnly = &v
}

func (o EquipmentSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquipmentSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnkamaId) {
		toSerialize["ankama_id"] = o.AnkamaId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.EquipmentIds) {
		toSerialize["equipment_ids"] = o.EquipmentIds
	}
	if !IsNil(o.Effects) {
		toSerialize["effects"] = o.Effects
	}
	if !IsNil(o.HighestEquipmentLevel) {
		toSerialize["highest_equipment_level"] = o.HighestEquipmentLevel
	}
	if !IsNil(o.ContainsCosmetics) {
		toSerialize["contains_cosmetics"] = o.ContainsCosmetics
	}
	if !IsNil(o.ContainsCosmeticsOnly) {
		toSerialize["contains_cosmetics_only"] = o.ContainsCosmeticsOnly
	}
	return toSerialize, nil
}

type NullableEquipmentSet struct {
	value *EquipmentSet
	isSet bool
}

func (v NullableEquipmentSet) Get() *EquipmentSet {
	return v.value
}

func (v *NullableEquipmentSet) Set(val *EquipmentSet) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentSet) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentSet(val *EquipmentSet) *NullableEquipmentSet {
	return &NullableEquipmentSet{value: val, isSet: true}
}

func (v NullableEquipmentSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


