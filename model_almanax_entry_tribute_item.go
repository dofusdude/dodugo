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

// checks if the AlmanaxEntryTributeItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlmanaxEntryTributeItem{}

// AlmanaxEntryTributeItem struct for AlmanaxEntryTributeItem
type AlmanaxEntryTributeItem struct {
	AnkamaId *int32 `json:"ankama_id,omitempty"`
	ImageUrls *ImageUrls `json:"image_urls,omitempty"`
	Name *string `json:"name,omitempty"`
	Subtype *string `json:"subtype,omitempty"`
}

// NewAlmanaxEntryTributeItem instantiates a new AlmanaxEntryTributeItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlmanaxEntryTributeItem() *AlmanaxEntryTributeItem {
	this := AlmanaxEntryTributeItem{}
	return &this
}

// NewAlmanaxEntryTributeItemWithDefaults instantiates a new AlmanaxEntryTributeItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlmanaxEntryTributeItemWithDefaults() *AlmanaxEntryTributeItem {
	this := AlmanaxEntryTributeItem{}
	return &this
}

// GetAnkamaId returns the AnkamaId field value if set, zero value otherwise.
func (o *AlmanaxEntryTributeItem) GetAnkamaId() int32 {
	if o == nil || IsNil(o.AnkamaId) {
		var ret int32
		return ret
	}
	return *o.AnkamaId
}

// GetAnkamaIdOk returns a tuple with the AnkamaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxEntryTributeItem) GetAnkamaIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AnkamaId) {
		return nil, false
	}
	return o.AnkamaId, true
}

// HasAnkamaId returns a boolean if a field has been set.
func (o *AlmanaxEntryTributeItem) HasAnkamaId() bool {
	if o != nil && !IsNil(o.AnkamaId) {
		return true
	}

	return false
}

// SetAnkamaId gets a reference to the given int32 and assigns it to the AnkamaId field.
func (o *AlmanaxEntryTributeItem) SetAnkamaId(v int32) {
	o.AnkamaId = &v
}

// GetImageUrls returns the ImageUrls field value if set, zero value otherwise.
func (o *AlmanaxEntryTributeItem) GetImageUrls() ImageUrls {
	if o == nil || IsNil(o.ImageUrls) {
		var ret ImageUrls
		return ret
	}
	return *o.ImageUrls
}

// GetImageUrlsOk returns a tuple with the ImageUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxEntryTributeItem) GetImageUrlsOk() (*ImageUrls, bool) {
	if o == nil || IsNil(o.ImageUrls) {
		return nil, false
	}
	return o.ImageUrls, true
}

// HasImageUrls returns a boolean if a field has been set.
func (o *AlmanaxEntryTributeItem) HasImageUrls() bool {
	if o != nil && !IsNil(o.ImageUrls) {
		return true
	}

	return false
}

// SetImageUrls gets a reference to the given ImageUrls and assigns it to the ImageUrls field.
func (o *AlmanaxEntryTributeItem) SetImageUrls(v ImageUrls) {
	o.ImageUrls = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlmanaxEntryTributeItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxEntryTributeItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlmanaxEntryTributeItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlmanaxEntryTributeItem) SetName(v string) {
	o.Name = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *AlmanaxEntryTributeItem) GetSubtype() string {
	if o == nil || IsNil(o.Subtype) {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxEntryTributeItem) GetSubtypeOk() (*string, bool) {
	if o == nil || IsNil(o.Subtype) {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *AlmanaxEntryTributeItem) HasSubtype() bool {
	if o != nil && !IsNil(o.Subtype) {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *AlmanaxEntryTributeItem) SetSubtype(v string) {
	o.Subtype = &v
}

func (o AlmanaxEntryTributeItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlmanaxEntryTributeItem) ToMap() (map[string]interface{}, error) {
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

type NullableAlmanaxEntryTributeItem struct {
	value *AlmanaxEntryTributeItem
	isSet bool
}

func (v NullableAlmanaxEntryTributeItem) Get() *AlmanaxEntryTributeItem {
	return v.value
}

func (v *NullableAlmanaxEntryTributeItem) Set(val *AlmanaxEntryTributeItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAlmanaxEntryTributeItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAlmanaxEntryTributeItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlmanaxEntryTributeItem(val *AlmanaxEntryTributeItem) *NullableAlmanaxEntryTributeItem {
	return &NullableAlmanaxEntryTributeItem{value: val, isSet: true}
}

func (v NullableAlmanaxEntryTributeItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlmanaxEntryTributeItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


