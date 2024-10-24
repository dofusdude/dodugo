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

// checks if the GetMetaVersion200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMetaVersion200Response{}

// GetMetaVersion200Response struct for GetMetaVersion200Response
type GetMetaVersion200Response struct {
	Version *string `json:"version,omitempty"`
	Release *string `json:"release,omitempty"`
	UpdateStamp *string `json:"update_stamp,omitempty"`
}

// NewGetMetaVersion200Response instantiates a new GetMetaVersion200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMetaVersion200Response() *GetMetaVersion200Response {
	this := GetMetaVersion200Response{}
	return &this
}

// NewGetMetaVersion200ResponseWithDefaults instantiates a new GetMetaVersion200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMetaVersion200ResponseWithDefaults() *GetMetaVersion200Response {
	this := GetMetaVersion200Response{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GetMetaVersion200Response) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMetaVersion200Response) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GetMetaVersion200Response) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *GetMetaVersion200Response) SetVersion(v string) {
	o.Version = &v
}

// GetRelease returns the Release field value if set, zero value otherwise.
func (o *GetMetaVersion200Response) GetRelease() string {
	if o == nil || IsNil(o.Release) {
		var ret string
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMetaVersion200Response) GetReleaseOk() (*string, bool) {
	if o == nil || IsNil(o.Release) {
		return nil, false
	}
	return o.Release, true
}

// HasRelease returns a boolean if a field has been set.
func (o *GetMetaVersion200Response) HasRelease() bool {
	if o != nil && !IsNil(o.Release) {
		return true
	}

	return false
}

// SetRelease gets a reference to the given string and assigns it to the Release field.
func (o *GetMetaVersion200Response) SetRelease(v string) {
	o.Release = &v
}

// GetUpdateStamp returns the UpdateStamp field value if set, zero value otherwise.
func (o *GetMetaVersion200Response) GetUpdateStamp() string {
	if o == nil || IsNil(o.UpdateStamp) {
		var ret string
		return ret
	}
	return *o.UpdateStamp
}

// GetUpdateStampOk returns a tuple with the UpdateStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMetaVersion200Response) GetUpdateStampOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateStamp) {
		return nil, false
	}
	return o.UpdateStamp, true
}

// HasUpdateStamp returns a boolean if a field has been set.
func (o *GetMetaVersion200Response) HasUpdateStamp() bool {
	if o != nil && !IsNil(o.UpdateStamp) {
		return true
	}

	return false
}

// SetUpdateStamp gets a reference to the given string and assigns it to the UpdateStamp field.
func (o *GetMetaVersion200Response) SetUpdateStamp(v string) {
	o.UpdateStamp = &v
}

func (o GetMetaVersion200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMetaVersion200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Release) {
		toSerialize["release"] = o.Release
	}
	if !IsNil(o.UpdateStamp) {
		toSerialize["update_stamp"] = o.UpdateStamp
	}
	return toSerialize, nil
}

type NullableGetMetaVersion200Response struct {
	value *GetMetaVersion200Response
	isSet bool
}

func (v NullableGetMetaVersion200Response) Get() *GetMetaVersion200Response {
	return v.value
}

func (v *NullableGetMetaVersion200Response) Set(val *GetMetaVersion200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMetaVersion200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMetaVersion200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMetaVersion200Response(val *GetMetaVersion200Response) *NullableGetMetaVersion200Response {
	return &NullableGetMetaVersion200Response{value: val, isSet: true}
}

func (v NullableGetMetaVersion200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMetaVersion200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


