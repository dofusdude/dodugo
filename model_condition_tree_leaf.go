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

// checks if the ConditionTreeLeaf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConditionTreeLeaf{}

// ConditionTreeLeaf struct for ConditionTreeLeaf
type ConditionTreeLeaf struct {
	// always \"true\" for the leaf of a tree
	IsOperand *bool `json:"is_operand,omitempty"`
	Condition *ConditionEntry `json:"condition,omitempty"`
}

// NewConditionTreeLeaf instantiates a new ConditionTreeLeaf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionTreeLeaf() *ConditionTreeLeaf {
	this := ConditionTreeLeaf{}
	var isOperand bool = true
	this.IsOperand = &isOperand
	return &this
}

// NewConditionTreeLeafWithDefaults instantiates a new ConditionTreeLeaf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionTreeLeafWithDefaults() *ConditionTreeLeaf {
	this := ConditionTreeLeaf{}
	var isOperand bool = true
	this.IsOperand = &isOperand
	return &this
}

// GetIsOperand returns the IsOperand field value if set, zero value otherwise.
func (o *ConditionTreeLeaf) GetIsOperand() bool {
	if o == nil || IsNil(o.IsOperand) {
		var ret bool
		return ret
	}
	return *o.IsOperand
}

// GetIsOperandOk returns a tuple with the IsOperand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionTreeLeaf) GetIsOperandOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOperand) {
		return nil, false
	}
	return o.IsOperand, true
}

// HasIsOperand returns a boolean if a field has been set.
func (o *ConditionTreeLeaf) HasIsOperand() bool {
	if o != nil && !IsNil(o.IsOperand) {
		return true
	}

	return false
}

// SetIsOperand gets a reference to the given bool and assigns it to the IsOperand field.
func (o *ConditionTreeLeaf) SetIsOperand(v bool) {
	o.IsOperand = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *ConditionTreeLeaf) GetCondition() ConditionEntry {
	if o == nil || IsNil(o.Condition) {
		var ret ConditionEntry
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionTreeLeaf) GetConditionOk() (*ConditionEntry, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *ConditionTreeLeaf) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given ConditionEntry and assigns it to the Condition field.
func (o *ConditionTreeLeaf) SetCondition(v ConditionEntry) {
	o.Condition = &v
}

func (o ConditionTreeLeaf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConditionTreeLeaf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsOperand) {
		toSerialize["is_operand"] = o.IsOperand
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	return toSerialize, nil
}

type NullableConditionTreeLeaf struct {
	value *ConditionTreeLeaf
	isSet bool
}

func (v NullableConditionTreeLeaf) Get() *ConditionTreeLeaf {
	return v.value
}

func (v *NullableConditionTreeLeaf) Set(val *ConditionTreeLeaf) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionTreeLeaf) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionTreeLeaf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionTreeLeaf(val *ConditionTreeLeaf) *NullableConditionTreeLeaf {
	return &NullableConditionTreeLeaf{value: val, isSet: true}
}

func (v NullableConditionTreeLeaf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionTreeLeaf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


