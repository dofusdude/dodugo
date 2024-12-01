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

// checks if the AlmanaxTributeItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlmanaxTributeItem{}

// AlmanaxTributeItem struct for AlmanaxTributeItem
type AlmanaxTributeItem struct {
	AnkamaId *int32 `json:"ankama_id,omitempty"`
	ImageUrls *Images `json:"image_urls,omitempty"`
	Name *string `json:"name,omitempty"`
	Subtype *string `json:"subtype,omitempty"`
}

// NewAlmanaxTributeItem instantiates a new AlmanaxTributeItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlmanaxTributeItem() *AlmanaxTributeItem {
	this := AlmanaxTributeItem{}
	return &this
}

// NewAlmanaxTributeItemWithDefaults instantiates a new AlmanaxTributeItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlmanaxTributeItemWithDefaults() *AlmanaxTributeItem {
	this := AlmanaxTributeItem{}
	return &this
}

// GetAnkamaId returns the AnkamaId field value if set, zero value otherwise.
func (o *AlmanaxTributeItem) GetAnkamaId() int32 {
	if o == nil || IsNil(o.AnkamaId) {
		var ret int32
		return ret
	}
	return *o.AnkamaId
}

// GetAnkamaIdOk returns a tuple with the AnkamaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxTributeItem) GetAnkamaIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AnkamaId) {
		return nil, false
	}
	return o.AnkamaId, true
}

// HasAnkamaId returns a boolean if a field has been set.
func (o *AlmanaxTributeItem) HasAnkamaId() bool {
	if o != nil && !IsNil(o.AnkamaId) {
		return true
	}

	return false
}

// SetAnkamaId gets a reference to the given int32 and assigns it to the AnkamaId field.
func (o *AlmanaxTributeItem) SetAnkamaId(v int32) {
	o.AnkamaId = &v
}

// GetImageUrls returns the ImageUrls field value if set, zero value otherwise.
func (o *AlmanaxTributeItem) GetImageUrls() Images {
	if o == nil || IsNil(o.ImageUrls) {
		var ret Images
		return ret
	}
	return *o.ImageUrls
}

// GetImageUrlsOk returns a tuple with the ImageUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxTributeItem) GetImageUrlsOk() (*Images, bool) {
	if o == nil || IsNil(o.ImageUrls) {
		return nil, false
	}
	return o.ImageUrls, true
}

// HasImageUrls returns a boolean if a field has been set.
func (o *AlmanaxTributeItem) HasImageUrls() bool {
	if o != nil && !IsNil(o.ImageUrls) {
		return true
	}

	return false
}

// SetImageUrls gets a reference to the given Images and assigns it to the ImageUrls field.
func (o *AlmanaxTributeItem) SetImageUrls(v Images) {
	o.ImageUrls = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlmanaxTributeItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxTributeItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlmanaxTributeItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlmanaxTributeItem) SetName(v string) {
	o.Name = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *AlmanaxTributeItem) GetSubtype() string {
	if o == nil || IsNil(o.Subtype) {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxTributeItem) GetSubtypeOk() (*string, bool) {
	if o == nil || IsNil(o.Subtype) {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *AlmanaxTributeItem) HasSubtype() bool {
	if o != nil && !IsNil(o.Subtype) {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *AlmanaxTributeItem) SetSubtype(v string) {
	o.Subtype = &v
}

func (o AlmanaxTributeItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlmanaxTributeItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnkamaId) {
		toSerialize["ankama_id"] = o.AnkamaId
	}
	if !IsNil(o.ImageUrls) {
		toSerialize["image_urls"] = o.ImageUrls
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Subtype) {
		toSerialize["subtype"] = o.Subtype
	}
	return toSerialize, nil
}

type NullableAlmanaxTributeItem struct {
	value *AlmanaxTributeItem
	isSet bool
}

func (v NullableAlmanaxTributeItem) Get() *AlmanaxTributeItem {
	return v.value
}

func (v *NullableAlmanaxTributeItem) Set(val *AlmanaxTributeItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAlmanaxTributeItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAlmanaxTributeItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlmanaxTributeItem(val *AlmanaxTributeItem) *NullableAlmanaxTributeItem {
	return &NullableAlmanaxTributeItem{value: val, isSet: true}
}

func (v NullableAlmanaxTributeItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlmanaxTributeItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


