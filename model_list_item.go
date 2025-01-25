/*
dofusdude

# Open Ankama Developer Community The all-in-one toolbelt for your next Ankama related project.  ## Versions - [Dofus 2](https://docs.dofusdu.de/dofus2/) - [Dofus 3](https://docs.dofusdu.de/dofus3/)   - v1 [latest] (you are here)   ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save` - [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save` - [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo` - [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude` - [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Almanax Discord Integration** Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 3 Beta** from stable to bleeding edge by replacing /dofus3 with /dofus3beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_, _de_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Official Sources** generated from actual data from the game.  ... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h). 

API version: 1.0.0-rc.9
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
)

// checks if the ListItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListItem{}

// ListItem struct for ListItem
type ListItem struct {
	AnkamaId *int32 `json:"ankama_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *TranslatedId `json:"type,omitempty"`
	Level *int32 `json:"level,omitempty"`
	ImageUrls *Images `json:"image_urls,omitempty"`
	Recipe []Recipe `json:"recipe,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Conditions NullableConditionNode `json:"conditions,omitempty"`
	Effects []Effect `json:"effects,omitempty"`
	IsWeapon NullableBool `json:"is_weapon,omitempty"`
	Pods NullableInt32 `json:"pods,omitempty"`
	ParentSet NullableTranslatedId `json:"parent_set,omitempty"`
	CriticalHitProbability NullableInt32 `json:"critical_hit_probability,omitempty"`
	CriticalHitBonus NullableInt32 `json:"critical_hit_bonus,omitempty"`
	MaxCastPerTurn NullableInt32 `json:"max_cast_per_turn,omitempty"`
	ApCost NullableInt32 `json:"ap_cost,omitempty"`
	Range *Range `json:"range,omitempty"`
}

// NewListItem instantiates a new ListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListItem() *ListItem {
	this := ListItem{}
	return &this
}

// NewListItemWithDefaults instantiates a new ListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListItemWithDefaults() *ListItem {
	this := ListItem{}
	return &this
}

// GetAnkamaId returns the AnkamaId field value if set, zero value otherwise.
func (o *ListItem) GetAnkamaId() int32 {
	if o == nil || IsNil(o.AnkamaId) {
		var ret int32
		return ret
	}
	return *o.AnkamaId
}

// GetAnkamaIdOk returns a tuple with the AnkamaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListItem) GetAnkamaIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AnkamaId) {
		return nil, false
	}
	return o.AnkamaId, true
}

// HasAnkamaId returns a boolean if a field has been set.
func (o *ListItem) HasAnkamaId() bool {
	if o != nil && !IsNil(o.AnkamaId) {
		return true
	}

	return false
}

// SetAnkamaId gets a reference to the given int32 and assigns it to the AnkamaId field.
func (o *ListItem) SetAnkamaId(v int32) {
	o.AnkamaId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListItem) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListItem) GetType() TranslatedId {
	if o == nil || IsNil(o.Type) {
		var ret TranslatedId
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListItem) GetTypeOk() (*TranslatedId, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ListItem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given TranslatedId and assigns it to the Type field.
func (o *ListItem) SetType(v TranslatedId) {
	o.Type = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *ListItem) GetLevel() int32 {
	if o == nil || IsNil(o.Level) {
		var ret int32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListItem) GetLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *ListItem) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given int32 and assigns it to the Level field.
func (o *ListItem) SetLevel(v int32) {
	o.Level = &v
}

// GetImageUrls returns the ImageUrls field value if set, zero value otherwise.
func (o *ListItem) GetImageUrls() Images {
	if o == nil || IsNil(o.ImageUrls) {
		var ret Images
		return ret
	}
	return *o.ImageUrls
}

// GetImageUrlsOk returns a tuple with the ImageUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListItem) GetImageUrlsOk() (*Images, bool) {
	if o == nil || IsNil(o.ImageUrls) {
		return nil, false
	}
	return o.ImageUrls, true
}

// HasImageUrls returns a boolean if a field has been set.
func (o *ListItem) HasImageUrls() bool {
	if o != nil && !IsNil(o.ImageUrls) {
		return true
	}

	return false
}

// SetImageUrls gets a reference to the given Images and assigns it to the ImageUrls field.
func (o *ListItem) SetImageUrls(v Images) {
	o.ImageUrls = &v
}

// GetRecipe returns the Recipe field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListItem) GetRecipe() []Recipe {
	if o == nil {
		var ret []Recipe
		return ret
	}
	return o.Recipe
}

// GetRecipeOk returns a tuple with the Recipe field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListItem) GetRecipeOk() ([]Recipe, bool) {
	if o == nil || IsNil(o.Recipe) {
		return nil, false
	}
	return o.Recipe, true
}

// HasRecipe returns a boolean if a field has been set.
func (o *ListItem) HasRecipe() bool {
	if o != nil && !IsNil(o.Recipe) {
		return true
	}

	return false
}

// SetRecipe gets a reference to the given []Recipe and assigns it to the Recipe field.
func (o *ListItem) SetRecipe(v []Recipe) {
	o.Recipe = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListItem) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ListItem) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ListItem) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ListItem) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ListItem) UnsetDescription() {
	o.Description.Unset()
}

// GetConditions returns the Conditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListItem) GetConditions() ConditionNode {
	if o == nil || IsNil(o.Conditions.Get()) {
		var ret ConditionNode
		return ret
	}
	return *o.Conditions.Get()
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListItem) GetConditionsOk() (*ConditionNode, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions.Get(), o.Conditions.IsSet()
}

// HasConditions returns a boolean if a field has been set.
func (o *ListItem) HasConditions() bool {
	if o != nil && o.Conditions.IsSet() {
		return true
	}

	return false
}

// SetConditions gets a reference to the given NullableConditionNode and assigns it to the Conditions field.
func (o *ListItem) SetConditions(v ConditionNode) {
	o.Conditions.Set(&v)
}
// SetConditionsNil sets the value for Conditions to be an explicit nil
func (o *ListItem) SetConditionsNil() {
	o.Conditions.Set(nil)
}

// UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
func (o *ListItem) UnsetConditions() {
	o.Conditions.Unset()
}

// GetEffects returns the Effects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListItem) GetEffects() []Effect {
	if o == nil {
		var ret []Effect
		return ret
	}
	return o.Effects
}

// GetEffectsOk returns a tuple with the Effects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListItem) GetEffectsOk() ([]Effect, bool) {
	if o == nil || IsNil(o.Effects) {
		return nil, false
	}
	return o.Effects, true
}

// HasEffects returns a boolean if a field has been set.
func (o *ListItem) HasEffects() bool {
	if o != nil && !IsNil(o.Effects) {
		return true
	}

	return false
}

// SetEffects gets a reference to the given []Effect and assigns it to the Effects field.
func (o *ListItem) SetEffects(v []Effect) {
	o.Effects = v
}

// GetIsWeapon returns the IsWeapon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListItem) GetIsWeapon() bool {
	if o == nil || IsNil(o.IsWeapon.Get()) {
		var ret bool
		return ret
	}
	return *o.IsWeapon.Get()
}

// GetIsWeaponOk returns a tuple with the IsWeapon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListItem) GetIsWeaponOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsWeapon.Get(), o.IsWeapon.IsSet()
}

// HasIsWeapon returns a boolean if a field has been set.
func (o *ListItem) HasIsWeapon() bool {
	if o != nil && o.IsWeapon.IsSet() {
		return true
	}

	return false
}

// SetIsWeapon gets a reference to the given NullableBool and assigns it to the IsWeapon field.
func (o *ListItem) SetIsWeapon(v bool) {
	o.IsWeapon.Set(&v)
}
// SetIsWeaponNil sets the value for IsWeapon to be an explicit nil
func (o *ListItem) SetIsWeaponNil() {
	o.IsWeapon.Set(nil)
}

// UnsetIsWeapon ensures that no value is present for IsWeapon, not even an explicit nil
func (o *ListItem) UnsetIsWeapon() {
	o.IsWeapon.Unset()
}

// GetPods returns the Pods field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListItem) GetPods() int32 {
	if o == nil || IsNil(o.Pods.Get()) {
		var ret int32
		return ret
	}
	return *o.Pods.Get()
}

// GetPodsOk returns a tuple with the Pods field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListItem) GetPodsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pods.Get(), o.Pods.IsSet()
}

// HasPods returns a boolean if a field has been set.
func (o *ListItem) HasPods() bool {
	if o != nil && o.Pods.IsSet() {
		return true
	}

	return false
}

// SetPods gets a reference to the given NullableInt32 and assigns it to the Pods field.
func (o *ListItem) SetPods(v int32) {
	o.Pods.Set(&v)
}
// SetPodsNil sets the value for Pods to be an explicit nil
func (o *ListItem) SetPodsNil() {
	o.Pods.Set(nil)
}

// UnsetPods ensures that no value is present for Pods, not even an explicit nil
func (o *ListItem) UnsetPods() {
	o.Pods.Unset()
}

// GetParentSet returns the ParentSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListItem) GetParentSet() TranslatedId {
	if o == nil || IsNil(o.ParentSet.Get()) {
		var ret TranslatedId
		return ret
	}
	return *o.ParentSet.Get()
}

// GetParentSetOk returns a tuple with the ParentSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListItem) GetParentSetOk() (*TranslatedId, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentSet.Get(), o.ParentSet.IsSet()
}

// HasParentSet returns a boolean if a field has been set.
func (o *ListItem) HasParentSet() bool {
	if o != nil && o.ParentSet.IsSet() {
		return true
	}

	return false
}

// SetParentSet gets a reference to the given NullableTranslatedId and assigns it to the ParentSet field.
func (o *ListItem) SetParentSet(v TranslatedId) {
	o.ParentSet.Set(&v)
}
// SetParentSetNil sets the value for ParentSet to be an explicit nil
func (o *ListItem) SetParentSetNil() {
	o.ParentSet.Set(nil)
}

// UnsetParentSet ensures that no value is present for ParentSet, not even an explicit nil
func (o *ListItem) UnsetParentSet() {
	o.ParentSet.Unset()
}

// GetCriticalHitProbability returns the CriticalHitProbability field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListItem) GetCriticalHitProbability() int32 {
	if o == nil || IsNil(o.CriticalHitProbability.Get()) {
		var ret int32
		return ret
	}
	return *o.CriticalHitProbability.Get()
}

// GetCriticalHitProbabilityOk returns a tuple with the CriticalHitProbability field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListItem) GetCriticalHitProbabilityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CriticalHitProbability.Get(), o.CriticalHitProbability.IsSet()
}

// HasCriticalHitProbability returns a boolean if a field has been set.
func (o *ListItem) HasCriticalHitProbability() bool {
	if o != nil && o.CriticalHitProbability.IsSet() {
		return true
	}

	return false
}

// SetCriticalHitProbability gets a reference to the given NullableInt32 and assigns it to the CriticalHitProbability field.
func (o *ListItem) SetCriticalHitProbability(v int32) {
	o.CriticalHitProbability.Set(&v)
}
// SetCriticalHitProbabilityNil sets the value for CriticalHitProbability to be an explicit nil
func (o *ListItem) SetCriticalHitProbabilityNil() {
	o.CriticalHitProbability.Set(nil)
}

// UnsetCriticalHitProbability ensures that no value is present for CriticalHitProbability, not even an explicit nil
func (o *ListItem) UnsetCriticalHitProbability() {
	o.CriticalHitProbability.Unset()
}

// GetCriticalHitBonus returns the CriticalHitBonus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListItem) GetCriticalHitBonus() int32 {
	if o == nil || IsNil(o.CriticalHitBonus.Get()) {
		var ret int32
		return ret
	}
	return *o.CriticalHitBonus.Get()
}

// GetCriticalHitBonusOk returns a tuple with the CriticalHitBonus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListItem) GetCriticalHitBonusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CriticalHitBonus.Get(), o.CriticalHitBonus.IsSet()
}

// HasCriticalHitBonus returns a boolean if a field has been set.
func (o *ListItem) HasCriticalHitBonus() bool {
	if o != nil && o.CriticalHitBonus.IsSet() {
		return true
	}

	return false
}

// SetCriticalHitBonus gets a reference to the given NullableInt32 and assigns it to the CriticalHitBonus field.
func (o *ListItem) SetCriticalHitBonus(v int32) {
	o.CriticalHitBonus.Set(&v)
}
// SetCriticalHitBonusNil sets the value for CriticalHitBonus to be an explicit nil
func (o *ListItem) SetCriticalHitBonusNil() {
	o.CriticalHitBonus.Set(nil)
}

// UnsetCriticalHitBonus ensures that no value is present for CriticalHitBonus, not even an explicit nil
func (o *ListItem) UnsetCriticalHitBonus() {
	o.CriticalHitBonus.Unset()
}

// GetMaxCastPerTurn returns the MaxCastPerTurn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListItem) GetMaxCastPerTurn() int32 {
	if o == nil || IsNil(o.MaxCastPerTurn.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxCastPerTurn.Get()
}

// GetMaxCastPerTurnOk returns a tuple with the MaxCastPerTurn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListItem) GetMaxCastPerTurnOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxCastPerTurn.Get(), o.MaxCastPerTurn.IsSet()
}

// HasMaxCastPerTurn returns a boolean if a field has been set.
func (o *ListItem) HasMaxCastPerTurn() bool {
	if o != nil && o.MaxCastPerTurn.IsSet() {
		return true
	}

	return false
}

// SetMaxCastPerTurn gets a reference to the given NullableInt32 and assigns it to the MaxCastPerTurn field.
func (o *ListItem) SetMaxCastPerTurn(v int32) {
	o.MaxCastPerTurn.Set(&v)
}
// SetMaxCastPerTurnNil sets the value for MaxCastPerTurn to be an explicit nil
func (o *ListItem) SetMaxCastPerTurnNil() {
	o.MaxCastPerTurn.Set(nil)
}

// UnsetMaxCastPerTurn ensures that no value is present for MaxCastPerTurn, not even an explicit nil
func (o *ListItem) UnsetMaxCastPerTurn() {
	o.MaxCastPerTurn.Unset()
}

// GetApCost returns the ApCost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListItem) GetApCost() int32 {
	if o == nil || IsNil(o.ApCost.Get()) {
		var ret int32
		return ret
	}
	return *o.ApCost.Get()
}

// GetApCostOk returns a tuple with the ApCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListItem) GetApCostOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApCost.Get(), o.ApCost.IsSet()
}

// HasApCost returns a boolean if a field has been set.
func (o *ListItem) HasApCost() bool {
	if o != nil && o.ApCost.IsSet() {
		return true
	}

	return false
}

// SetApCost gets a reference to the given NullableInt32 and assigns it to the ApCost field.
func (o *ListItem) SetApCost(v int32) {
	o.ApCost.Set(&v)
}
// SetApCostNil sets the value for ApCost to be an explicit nil
func (o *ListItem) SetApCostNil() {
	o.ApCost.Set(nil)
}

// UnsetApCost ensures that no value is present for ApCost, not even an explicit nil
func (o *ListItem) UnsetApCost() {
	o.ApCost.Unset()
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *ListItem) GetRange() Range {
	if o == nil || IsNil(o.Range) {
		var ret Range
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListItem) GetRangeOk() (*Range, bool) {
	if o == nil || IsNil(o.Range) {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *ListItem) HasRange() bool {
	if o != nil && !IsNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given Range and assigns it to the Range field.
func (o *ListItem) SetRange(v Range) {
	o.Range = &v
}

func (o ListItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnkamaId) {
		toSerialize["ankama_id"] = o.AnkamaId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.ImageUrls) {
		toSerialize["image_urls"] = o.ImageUrls
	}
	if o.Recipe != nil {
		toSerialize["recipe"] = o.Recipe
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Conditions.IsSet() {
		toSerialize["conditions"] = o.Conditions.Get()
	}
	if o.Effects != nil {
		toSerialize["effects"] = o.Effects
	}
	if o.IsWeapon.IsSet() {
		toSerialize["is_weapon"] = o.IsWeapon.Get()
	}
	if o.Pods.IsSet() {
		toSerialize["pods"] = o.Pods.Get()
	}
	if o.ParentSet.IsSet() {
		toSerialize["parent_set"] = o.ParentSet.Get()
	}
	if o.CriticalHitProbability.IsSet() {
		toSerialize["critical_hit_probability"] = o.CriticalHitProbability.Get()
	}
	if o.CriticalHitBonus.IsSet() {
		toSerialize["critical_hit_bonus"] = o.CriticalHitBonus.Get()
	}
	if o.MaxCastPerTurn.IsSet() {
		toSerialize["max_cast_per_turn"] = o.MaxCastPerTurn.Get()
	}
	if o.ApCost.IsSet() {
		toSerialize["ap_cost"] = o.ApCost.Get()
	}
	if !IsNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	return toSerialize, nil
}

type NullableListItem struct {
	value *ListItem
	isSet bool
}

func (v NullableListItem) Get() *ListItem {
	return v.value
}

func (v *NullableListItem) Set(val *ListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListItem(val *ListItem) *NullableListItem {
	return &NullableListItem{value: val, isSet: true}
}

func (v NullableListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


