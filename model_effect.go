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

// checks if the Effect type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Effect{}

// Effect struct for Effect
type Effect struct {
	// minimum int value, can be a single if ignore_int_max and no ignore_int_min
	IntMinimum *int32 `json:"int_minimum,omitempty"`
	// maximum int value, if not ignore_int_max and not ignore_int_min, the effect has a range value
	IntMaximum *int32 `json:"int_maximum,omitempty"`
	Type *EffectType `json:"type,omitempty"`
	// ignore the int min field because the actual value is a string. For readability the templated field is the only important field in this case. 
	IgnoreIntMin *bool `json:"ignore_int_min,omitempty"`
	// ignore the int max field, if ignore_int_min is true, int min is a single value
	IgnoreIntMax *bool `json:"ignore_int_max,omitempty"`
	// all fields from above encoded in a single string
	Formatted *string `json:"formatted,omitempty"`
}

// NewEffect instantiates a new Effect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEffect() *Effect {
	this := Effect{}
	return &this
}

// NewEffectWithDefaults instantiates a new Effect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEffectWithDefaults() *Effect {
	this := Effect{}
	return &this
}

// GetIntMinimum returns the IntMinimum field value if set, zero value otherwise.
func (o *Effect) GetIntMinimum() int32 {
	if o == nil || IsNil(o.IntMinimum) {
		var ret int32
		return ret
	}
	return *o.IntMinimum
}

// GetIntMinimumOk returns a tuple with the IntMinimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Effect) GetIntMinimumOk() (*int32, bool) {
	if o == nil || IsNil(o.IntMinimum) {
		return nil, false
	}
	return o.IntMinimum, true
}

// HasIntMinimum returns a boolean if a field has been set.
func (o *Effect) HasIntMinimum() bool {
	if o != nil && !IsNil(o.IntMinimum) {
		return true
	}

	return false
}

// SetIntMinimum gets a reference to the given int32 and assigns it to the IntMinimum field.
func (o *Effect) SetIntMinimum(v int32) {
	o.IntMinimum = &v
}

// GetIntMaximum returns the IntMaximum field value if set, zero value otherwise.
func (o *Effect) GetIntMaximum() int32 {
	if o == nil || IsNil(o.IntMaximum) {
		var ret int32
		return ret
	}
	return *o.IntMaximum
}

// GetIntMaximumOk returns a tuple with the IntMaximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Effect) GetIntMaximumOk() (*int32, bool) {
	if o == nil || IsNil(o.IntMaximum) {
		return nil, false
	}
	return o.IntMaximum, true
}

// HasIntMaximum returns a boolean if a field has been set.
func (o *Effect) HasIntMaximum() bool {
	if o != nil && !IsNil(o.IntMaximum) {
		return true
	}

	return false
}

// SetIntMaximum gets a reference to the given int32 and assigns it to the IntMaximum field.
func (o *Effect) SetIntMaximum(v int32) {
	o.IntMaximum = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Effect) GetType() EffectType {
	if o == nil || IsNil(o.Type) {
		var ret EffectType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Effect) GetTypeOk() (*EffectType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Effect) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EffectType and assigns it to the Type field.
func (o *Effect) SetType(v EffectType) {
	o.Type = &v
}

// GetIgnoreIntMin returns the IgnoreIntMin field value if set, zero value otherwise.
func (o *Effect) GetIgnoreIntMin() bool {
	if o == nil || IsNil(o.IgnoreIntMin) {
		var ret bool
		return ret
	}
	return *o.IgnoreIntMin
}

// GetIgnoreIntMinOk returns a tuple with the IgnoreIntMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Effect) GetIgnoreIntMinOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreIntMin) {
		return nil, false
	}
	return o.IgnoreIntMin, true
}

// HasIgnoreIntMin returns a boolean if a field has been set.
func (o *Effect) HasIgnoreIntMin() bool {
	if o != nil && !IsNil(o.IgnoreIntMin) {
		return true
	}

	return false
}

// SetIgnoreIntMin gets a reference to the given bool and assigns it to the IgnoreIntMin field.
func (o *Effect) SetIgnoreIntMin(v bool) {
	o.IgnoreIntMin = &v
}

// GetIgnoreIntMax returns the IgnoreIntMax field value if set, zero value otherwise.
func (o *Effect) GetIgnoreIntMax() bool {
	if o == nil || IsNil(o.IgnoreIntMax) {
		var ret bool
		return ret
	}
	return *o.IgnoreIntMax
}

// GetIgnoreIntMaxOk returns a tuple with the IgnoreIntMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Effect) GetIgnoreIntMaxOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreIntMax) {
		return nil, false
	}
	return o.IgnoreIntMax, true
}

// HasIgnoreIntMax returns a boolean if a field has been set.
func (o *Effect) HasIgnoreIntMax() bool {
	if o != nil && !IsNil(o.IgnoreIntMax) {
		return true
	}

	return false
}

// SetIgnoreIntMax gets a reference to the given bool and assigns it to the IgnoreIntMax field.
func (o *Effect) SetIgnoreIntMax(v bool) {
	o.IgnoreIntMax = &v
}

// GetFormatted returns the Formatted field value if set, zero value otherwise.
func (o *Effect) GetFormatted() string {
	if o == nil || IsNil(o.Formatted) {
		var ret string
		return ret
	}
	return *o.Formatted
}

// GetFormattedOk returns a tuple with the Formatted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Effect) GetFormattedOk() (*string, bool) {
	if o == nil || IsNil(o.Formatted) {
		return nil, false
	}
	return o.Formatted, true
}

// HasFormatted returns a boolean if a field has been set.
func (o *Effect) HasFormatted() bool {
	if o != nil && !IsNil(o.Formatted) {
		return true
	}

	return false
}

// SetFormatted gets a reference to the given string and assigns it to the Formatted field.
func (o *Effect) SetFormatted(v string) {
	o.Formatted = &v
}

func (o Effect) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Effect) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IntMinimum) {
		toSerialize["int_minimum"] = o.IntMinimum
	}
	if !IsNil(o.IntMaximum) {
		toSerialize["int_maximum"] = o.IntMaximum
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.IgnoreIntMin) {
		toSerialize["ignore_int_min"] = o.IgnoreIntMin
	}
	if !IsNil(o.IgnoreIntMax) {
		toSerialize["ignore_int_max"] = o.IgnoreIntMax
	}
	if !IsNil(o.Formatted) {
		toSerialize["formatted"] = o.Formatted
	}
	return toSerialize, nil
}

type NullableEffect struct {
	value *Effect
	isSet bool
}

func (v NullableEffect) Get() *Effect {
	return v.value
}

func (v *NullableEffect) Set(val *Effect) {
	v.value = val
	v.isSet = true
}

func (v NullableEffect) IsSet() bool {
	return v.isSet
}

func (v *NullableEffect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEffect(val *Effect) *NullableEffect {
	return &NullableEffect{value: val, isSet: true}
}

func (v NullableEffect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEffect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


