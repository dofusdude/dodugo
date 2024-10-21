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

// checks if the ConditionEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConditionEntry{}

// ConditionEntry 
type ConditionEntry struct {
	Operator *string `json:"operator,omitempty"`
	IntValue *int32 `json:"int_value,omitempty"`
	Element *ItemsListEntryTypedType `json:"element,omitempty"`
}

// NewConditionEntry instantiates a new ConditionEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionEntry() *ConditionEntry {
	this := ConditionEntry{}
	return &this
}

// NewConditionEntryWithDefaults instantiates a new ConditionEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionEntryWithDefaults() *ConditionEntry {
	this := ConditionEntry{}
	return &this
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *ConditionEntry) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionEntry) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *ConditionEntry) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *ConditionEntry) SetOperator(v string) {
	o.Operator = &v
}

// GetIntValue returns the IntValue field value if set, zero value otherwise.
func (o *ConditionEntry) GetIntValue() int32 {
	if o == nil || IsNil(o.IntValue) {
		var ret int32
		return ret
	}
	return *o.IntValue
}

// GetIntValueOk returns a tuple with the IntValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionEntry) GetIntValueOk() (*int32, bool) {
	if o == nil || IsNil(o.IntValue) {
		return nil, false
	}
	return o.IntValue, true
}

// HasIntValue returns a boolean if a field has been set.
func (o *ConditionEntry) HasIntValue() bool {
	if o != nil && !IsNil(o.IntValue) {
		return true
	}

	return false
}

// SetIntValue gets a reference to the given int32 and assigns it to the IntValue field.
func (o *ConditionEntry) SetIntValue(v int32) {
	o.IntValue = &v
}

// GetElement returns the Element field value if set, zero value otherwise.
func (o *ConditionEntry) GetElement() ItemsListEntryTypedType {
	if o == nil || IsNil(o.Element) {
		var ret ItemsListEntryTypedType
		return ret
	}
	return *o.Element
}

// GetElementOk returns a tuple with the Element field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionEntry) GetElementOk() (*ItemsListEntryTypedType, bool) {
	if o == nil || IsNil(o.Element) {
		return nil, false
	}
	return o.Element, true
}

// HasElement returns a boolean if a field has been set.
func (o *ConditionEntry) HasElement() bool {
	if o != nil && !IsNil(o.Element) {
		return true
	}

	return false
}

// SetElement gets a reference to the given ItemsListEntryTypedType and assigns it to the Element field.
func (o *ConditionEntry) SetElement(v ItemsListEntryTypedType) {
	o.Element = &v
}

func (o ConditionEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConditionEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.IntValue) {
		toSerialize["int_value"] = o.IntValue
	}
	if !IsNil(o.Element) {
		toSerialize["element"] = o.Element
	}
	return toSerialize, nil
}

type NullableConditionEntry struct {
	value *ConditionEntry
	isSet bool
}

func (v NullableConditionEntry) Get() *ConditionEntry {
	return v.value
}

func (v *NullableConditionEntry) Set(val *ConditionEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionEntry(val *ConditionEntry) *NullableConditionEntry {
	return &NullableConditionEntry{value: val, isSet: true}
}

func (v NullableConditionEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


