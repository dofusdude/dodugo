/*
dofusdude

# Open Ankama Developer Community The all-in-one toolbelt for your next Ankama related project.  ## Versions - [Dofus 2](https://docs.dofusdu.de/dofus2/) - [Dofus 3](https://docs.dofusdu.de/dofus3/)   - v1 [latest] (you are here)   ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save` - [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save` - [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo` - [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude` - [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Almanax Discord Integration** Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 3 Beta** from stable to bleeding edge by replacing /dofus3 with /dofus3beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_, _de_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Official Sources** generated from actual data from the game.  ... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h). 

API version: 1.0.0-rc.3
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
)

// checks if the Resource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Resource{}

// Resource struct for Resource
type Resource struct {
	AnkamaId *int32 `json:"ankama_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Type *TranslatedId `json:"type,omitempty"`
	Level *int32 `json:"level,omitempty"`
	Pods *int32 `json:"pods,omitempty"`
	ImageUrls *Images `json:"image_urls,omitempty"`
	Effects []Effect `json:"effects,omitempty"`
	Conditions NullableConditionNode `json:"conditions,omitempty"`
	Recipe []Recipe `json:"recipe,omitempty"`
}

// NewResource instantiates a new Resource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResource() *Resource {
	this := Resource{}
	return &this
}

// NewResourceWithDefaults instantiates a new Resource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceWithDefaults() *Resource {
	this := Resource{}
	return &this
}

// GetAnkamaId returns the AnkamaId field value if set, zero value otherwise.
func (o *Resource) GetAnkamaId() int32 {
	if o == nil || IsNil(o.AnkamaId) {
		var ret int32
		return ret
	}
	return *o.AnkamaId
}

// GetAnkamaIdOk returns a tuple with the AnkamaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetAnkamaIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AnkamaId) {
		return nil, false
	}
	return o.AnkamaId, true
}

// HasAnkamaId returns a boolean if a field has been set.
func (o *Resource) HasAnkamaId() bool {
	if o != nil && !IsNil(o.AnkamaId) {
		return true
	}

	return false
}

// SetAnkamaId gets a reference to the given int32 and assigns it to the AnkamaId field.
func (o *Resource) SetAnkamaId(v int32) {
	o.AnkamaId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Resource) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Resource) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Resource) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Resource) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Resource) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Resource) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Resource) GetType() TranslatedId {
	if o == nil || IsNil(o.Type) {
		var ret TranslatedId
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetTypeOk() (*TranslatedId, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Resource) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given TranslatedId and assigns it to the Type field.
func (o *Resource) SetType(v TranslatedId) {
	o.Type = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *Resource) GetLevel() int32 {
	if o == nil || IsNil(o.Level) {
		var ret int32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *Resource) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given int32 and assigns it to the Level field.
func (o *Resource) SetLevel(v int32) {
	o.Level = &v
}

// GetPods returns the Pods field value if set, zero value otherwise.
func (o *Resource) GetPods() int32 {
	if o == nil || IsNil(o.Pods) {
		var ret int32
		return ret
	}
	return *o.Pods
}

// GetPodsOk returns a tuple with the Pods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetPodsOk() (*int32, bool) {
	if o == nil || IsNil(o.Pods) {
		return nil, false
	}
	return o.Pods, true
}

// HasPods returns a boolean if a field has been set.
func (o *Resource) HasPods() bool {
	if o != nil && !IsNil(o.Pods) {
		return true
	}

	return false
}

// SetPods gets a reference to the given int32 and assigns it to the Pods field.
func (o *Resource) SetPods(v int32) {
	o.Pods = &v
}

// GetImageUrls returns the ImageUrls field value if set, zero value otherwise.
func (o *Resource) GetImageUrls() Images {
	if o == nil || IsNil(o.ImageUrls) {
		var ret Images
		return ret
	}
	return *o.ImageUrls
}

// GetImageUrlsOk returns a tuple with the ImageUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetImageUrlsOk() (*Images, bool) {
	if o == nil || IsNil(o.ImageUrls) {
		return nil, false
	}
	return o.ImageUrls, true
}

// HasImageUrls returns a boolean if a field has been set.
func (o *Resource) HasImageUrls() bool {
	if o != nil && !IsNil(o.ImageUrls) {
		return true
	}

	return false
}

// SetImageUrls gets a reference to the given Images and assigns it to the ImageUrls field.
func (o *Resource) SetImageUrls(v Images) {
	o.ImageUrls = &v
}

// GetEffects returns the Effects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Resource) GetEffects() []Effect {
	if o == nil {
		var ret []Effect
		return ret
	}
	return o.Effects
}

// GetEffectsOk returns a tuple with the Effects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Resource) GetEffectsOk() ([]Effect, bool) {
	if o == nil || IsNil(o.Effects) {
		return nil, false
	}
	return o.Effects, true
}

// HasEffects returns a boolean if a field has been set.
func (o *Resource) HasEffects() bool {
	if o != nil && !IsNil(o.Effects) {
		return true
	}

	return false
}

// SetEffects gets a reference to the given []Effect and assigns it to the Effects field.
func (o *Resource) SetEffects(v []Effect) {
	o.Effects = v
}

// GetConditions returns the Conditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Resource) GetConditions() ConditionNode {
	if o == nil || IsNil(o.Conditions.Get()) {
		var ret ConditionNode
		return ret
	}
	return *o.Conditions.Get()
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Resource) GetConditionsOk() (*ConditionNode, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions.Get(), o.Conditions.IsSet()
}

// HasConditions returns a boolean if a field has been set.
func (o *Resource) HasConditions() bool {
	if o != nil && o.Conditions.IsSet() {
		return true
	}

	return false
}

// SetConditions gets a reference to the given NullableConditionNode and assigns it to the Conditions field.
func (o *Resource) SetConditions(v ConditionNode) {
	o.Conditions.Set(&v)
}
// SetConditionsNil sets the value for Conditions to be an explicit nil
func (o *Resource) SetConditionsNil() {
	o.Conditions.Set(nil)
}

// UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
func (o *Resource) UnsetConditions() {
	o.Conditions.Unset()
}

// GetRecipe returns the Recipe field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Resource) GetRecipe() []Recipe {
	if o == nil {
		var ret []Recipe
		return ret
	}
	return o.Recipe
}

// GetRecipeOk returns a tuple with the Recipe field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Resource) GetRecipeOk() ([]Recipe, bool) {
	if o == nil || IsNil(o.Recipe) {
		return nil, false
	}
	return o.Recipe, true
}

// HasRecipe returns a boolean if a field has been set.
func (o *Resource) HasRecipe() bool {
	if o != nil && !IsNil(o.Recipe) {
		return true
	}

	return false
}

// SetRecipe gets a reference to the given []Recipe and assigns it to the Recipe field.
func (o *Resource) SetRecipe(v []Recipe) {
	o.Recipe = v
}

func (o Resource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Resource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnkamaId) {
		toSerialize["ankama_id"] = o.AnkamaId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.Pods) {
		toSerialize["pods"] = o.Pods
	}
	if !IsNil(o.ImageUrls) {
		toSerialize["image_urls"] = o.ImageUrls
	}
	if o.Effects != nil {
		toSerialize["effects"] = o.Effects
	}
	if o.Conditions.IsSet() {
		toSerialize["conditions"] = o.Conditions.Get()
	}
	if o.Recipe != nil {
		toSerialize["recipe"] = o.Recipe
	}
	return toSerialize, nil
}

type NullableResource struct {
	value *Resource
	isSet bool
}

func (v NullableResource) Get() *Resource {
	return v.value
}

func (v *NullableResource) Set(val *Resource) {
	v.value = val
	v.isSet = true
}

func (v NullableResource) IsSet() bool {
	return v.isSet
}

func (v *NullableResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResource(val *Resource) *NullableResource {
	return &NullableResource{value: val, isSet: true}
}

func (v NullableResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


