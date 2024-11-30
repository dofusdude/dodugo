/*
dofusdude

# Open Ankama Developer Community The all-in-one toolbelt for your next Ankama related project.  ## Versions - [Dofus 2](https://docs.dofusdu.de/dofus2/) - [Dofus 3](https://docs.dofusdu.de/dofus3/)   - v1 [latest] (you are here)   ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save` - [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save` - [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo` - [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude` - [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Almanax Discord Integration** Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 3 Beta** from stable to bleeding edge by replacing /dofus3 with /dofus3beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_, _de_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Official Sources** generated from actual data from the game.  ... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h). 

API version: 1.0.0-rc.5
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
)

// checks if the Recipe type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Recipe{}

// Recipe struct for Recipe
type Recipe struct {
	ItemAnkamaId *int32 `json:"item_ankama_id,omitempty"`
	ItemSubtype *string `json:"item_subtype,omitempty"`
	Quantity *int32 `json:"quantity,omitempty"`
}

// NewRecipe instantiates a new Recipe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipe() *Recipe {
	this := Recipe{}
	return &this
}

// NewRecipeWithDefaults instantiates a new Recipe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipeWithDefaults() *Recipe {
	this := Recipe{}
	return &this
}

// GetItemAnkamaId returns the ItemAnkamaId field value if set, zero value otherwise.
func (o *Recipe) GetItemAnkamaId() int32 {
	if o == nil || IsNil(o.ItemAnkamaId) {
		var ret int32
		return ret
	}
	return *o.ItemAnkamaId
}

// GetItemAnkamaIdOk returns a tuple with the ItemAnkamaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recipe) GetItemAnkamaIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ItemAnkamaId) {
		return nil, false
	}
	return o.ItemAnkamaId, true
}

// HasItemAnkamaId returns a boolean if a field has been set.
func (o *Recipe) HasItemAnkamaId() bool {
	if o != nil && !IsNil(o.ItemAnkamaId) {
		return true
	}

	return false
}

// SetItemAnkamaId gets a reference to the given int32 and assigns it to the ItemAnkamaId field.
func (o *Recipe) SetItemAnkamaId(v int32) {
	o.ItemAnkamaId = &v
}

// GetItemSubtype returns the ItemSubtype field value if set, zero value otherwise.
func (o *Recipe) GetItemSubtype() string {
	if o == nil || IsNil(o.ItemSubtype) {
		var ret string
		return ret
	}
	return *o.ItemSubtype
}

// GetItemSubtypeOk returns a tuple with the ItemSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recipe) GetItemSubtypeOk() (*string, bool) {
	if o == nil || IsNil(o.ItemSubtype) {
		return nil, false
	}
	return o.ItemSubtype, true
}

// HasItemSubtype returns a boolean if a field has been set.
func (o *Recipe) HasItemSubtype() bool {
	if o != nil && !IsNil(o.ItemSubtype) {
		return true
	}

	return false
}

// SetItemSubtype gets a reference to the given string and assigns it to the ItemSubtype field.
func (o *Recipe) SetItemSubtype(v string) {
	o.ItemSubtype = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *Recipe) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recipe) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *Recipe) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *Recipe) SetQuantity(v int32) {
	o.Quantity = &v
}

func (o Recipe) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Recipe) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemAnkamaId) {
		toSerialize["item_ankama_id"] = o.ItemAnkamaId
	}
	if !IsNil(o.ItemSubtype) {
		toSerialize["item_subtype"] = o.ItemSubtype
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullableRecipe struct {
	value *Recipe
	isSet bool
}

func (v NullableRecipe) Get() *Recipe {
	return v.value
}

func (v *NullableRecipe) Set(val *Recipe) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipe) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipe(val *Recipe) *NullableRecipe {
	return &NullableRecipe{value: val, isSet: true}
}

func (v NullableRecipe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


