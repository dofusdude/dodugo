/*
dofusdude

# Open Ankama Developer Community The all-in-one toolbelt for your next Ankama related project.  ## Versions - [Dofus 2](https://docs.dofusdu.de/dofus2/) - [Dofus 3](https://docs.dofusdu.de/dofus3/)   - v1 [latest] (you are here)   ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save` - [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save` - [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo` - [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude` - [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Almanax Discord Integration** Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 3 Beta** from stable to bleeding edge by replacing /dofus3 with /dofus3beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_, _de_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Official Sources** generated from actual data from the game.  ... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h). 

API version: 1.0.0-rc.6
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
)

// checks if the GetMetaAlmanaxBonuses200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMetaAlmanaxBonuses200ResponseInner{}

// GetMetaAlmanaxBonuses200ResponseInner struct for GetMetaAlmanaxBonuses200ResponseInner
type GetMetaAlmanaxBonuses200ResponseInner struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewGetMetaAlmanaxBonuses200ResponseInner instantiates a new GetMetaAlmanaxBonuses200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMetaAlmanaxBonuses200ResponseInner() *GetMetaAlmanaxBonuses200ResponseInner {
	this := GetMetaAlmanaxBonuses200ResponseInner{}
	return &this
}

// NewGetMetaAlmanaxBonuses200ResponseInnerWithDefaults instantiates a new GetMetaAlmanaxBonuses200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMetaAlmanaxBonuses200ResponseInnerWithDefaults() *GetMetaAlmanaxBonuses200ResponseInner {
	this := GetMetaAlmanaxBonuses200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetMetaAlmanaxBonuses200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMetaAlmanaxBonuses200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetMetaAlmanaxBonuses200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetMetaAlmanaxBonuses200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetMetaAlmanaxBonuses200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMetaAlmanaxBonuses200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetMetaAlmanaxBonuses200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetMetaAlmanaxBonuses200ResponseInner) SetName(v string) {
	o.Name = &v
}

func (o GetMetaAlmanaxBonuses200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMetaAlmanaxBonuses200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGetMetaAlmanaxBonuses200ResponseInner struct {
	value *GetMetaAlmanaxBonuses200ResponseInner
	isSet bool
}

func (v NullableGetMetaAlmanaxBonuses200ResponseInner) Get() *GetMetaAlmanaxBonuses200ResponseInner {
	return v.value
}

func (v *NullableGetMetaAlmanaxBonuses200ResponseInner) Set(val *GetMetaAlmanaxBonuses200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMetaAlmanaxBonuses200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMetaAlmanaxBonuses200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMetaAlmanaxBonuses200ResponseInner(val *GetMetaAlmanaxBonuses200ResponseInner) *NullableGetMetaAlmanaxBonuses200ResponseInner {
	return &NullableGetMetaAlmanaxBonuses200ResponseInner{value: val, isSet: true}
}

func (v NullableGetMetaAlmanaxBonuses200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMetaAlmanaxBonuses200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


