/*
Dofusdude

The last API for everything Dofus 🤯  ```js var dofusdude = require(\"dofusdude-js\");  new dofusdude.AllItemsApi().getItemsAllSearch(   \"en\",   \"dofus2\",   \"nidas\",   { filterTypeName: \"hat\" },   (err, data, response) => {     console.log(data[0]);   } ); ```  ### Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) npm i dofusdude-js --save - [Typescript](https://github.com/dofusdude/dofusdude-ts) npm i dofusdude-ts --save  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue at the Docs Repo.  ## Main Features - 🥷 **seamless auto-update** load data in the background when a new Dofus version is released and serving it within 2 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **blazingly fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 🩸 **Dofus 2 Beta** from stable to bleeding edge by replacing /dofus2 with /dofus2beta.  - 🗣️ **multilingual** supporting _en_, _fr_, _es_, _pt_ including the dropped languages from the Dofus website _de_ and _it_.  - 🧠 **search by relevance** allowing typos in name and description, handled by language specific text analysis and indexing by the powerful [Meilisearch](https://www.meilisearch.com) written in Rust.  - 🕵️ **complete** actual data from the game including items invisible to the encyclopedia like quest items.  - 🖼️ **HD images** rendering vector graphics into PNGs up to 800x800 px in the background.   ## Current state - Weapons ✅ - Equipment ✅ - Sets ✅ - Resources ✅ - Consumables ✅ - Pets ✅ - Mounts ✅ - Cosmetics/Ceremonial Items ✅ - Harnesses ✅ - Quest Items ✅ - Almanax ✅  - Monsters ❌ - Classes ❌ - Spells ❌ - Professions ❌   ### Maybes? I don't know what for 🤷 - Sidekicks ❌ - Haven Bags ❌ - Map ❌   ## Future I want this project to be useful and not just add plain categories no one needs. More and more features will be added to enhance the quality based on the needs of you, the developers.  Examples: _I need to know where I can drop the all the items I need to craft set X!_  _Please get a detailed always up-to-date spell description so I can calculate the damage for my set builder site!_  Nearly everything can be done. But I want to make sure somebody also wants it.  If you have anything or you are just interested in the project, join the [Discord](https://discord.gg/3EtHskZD8h).  ### Versioning Updating an API is a hard problem. This is why we'll keep it simple:  Everything you see here on this site, you can use now and forever. Updates could introduce new fields, new paths or parameter but never break backwards compatibility, so no field or parameter will be deleted. Ever.  There is one exception! **The API will _always_ choose being up-to-date over everything else**. So if Ankama decides to drop languages from the game like they did with their website, the API will loose support for them, too.  We can prevent this specific use case with a nice community but even then, it would be hidden behind a feature flag.  ## Get started! 🥳 Scroll down and try it for yourself!  If you are ready to use them in your project, think about [generating a client 🧙](https://github.com/OpenAPITools/openapi-generator) or use one of our pre generated SDKs linked at the top.  Awesome Projects using this API:  - [KaellyBot](https://github.com/Kaysoro/KaellyBot) by Kaysoro - [Dofus Craftlist](https://dofuscraftlist-dev.netlify.app) by Lystina - [AlmanaxApp](https://almanaxapp.netlify.app) by Lystina - [luwnarya.fr](https://luwnarya.fr)  

API version: 0.5.3
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dofusdude-go

import (
	"encoding/json"
)

// Weapon 
type Weapon struct {
	AnkamaId *int32 `json:"ankama_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Type *ItemsListEntryTypedType `json:"type,omitempty"`
	// always true when the item is a weapon. Many fields are now available. Always check for this flag first when getting single equipment items.
	IsWeapon *bool `json:"is_weapon,omitempty"`
	Level *int32 `json:"level,omitempty"`
	Pods *int32 `json:"pods,omitempty"`
	ImageUrls *ImageUrls `json:"image_urls,omitempty"`
	HasEffects *bool `json:"has_effects,omitempty"`
	Effects []EffectsEntry `json:"effects,omitempty"`
	HasConditions *bool `json:"has_conditions,omitempty"`
	Conditions []ConditionEntry `json:"conditions,omitempty"`
	CriticalHitProbability *int32 `json:"critical_hit_probability,omitempty"`
	CriticalHitBonus *int32 `json:"critical_hit_bonus,omitempty"`
	IsTwoHanded *bool `json:"is_two_handed,omitempty"`
	MaxCastPerTurn *int32 `json:"max_cast_per_turn,omitempty"`
	ApCost *int32 `json:"ap_cost,omitempty"`
	Range *WeaponRange `json:"range,omitempty"`
	HasRecipe *bool `json:"has_recipe,omitempty"`
	Recipe []RecipeEntry `json:"recipe,omitempty"`
	HasParentSet *bool `json:"has_parent_set,omitempty"`
	ParentSet NullableEquipmentParentSet `json:"parent_set,omitempty"`
}

// NewWeapon instantiates a new Weapon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeapon() *Weapon {
	this := Weapon{}
	return &this
}

// NewWeaponWithDefaults instantiates a new Weapon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeaponWithDefaults() *Weapon {
	this := Weapon{}
	return &this
}

// GetAnkamaId returns the AnkamaId field value if set, zero value otherwise.
func (o *Weapon) GetAnkamaId() int32 {
	if o == nil || o.AnkamaId == nil {
		var ret int32
		return ret
	}
	return *o.AnkamaId
}

// GetAnkamaIdOk returns a tuple with the AnkamaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetAnkamaIdOk() (*int32, bool) {
	if o == nil || o.AnkamaId == nil {
		return nil, false
	}
	return o.AnkamaId, true
}

// HasAnkamaId returns a boolean if a field has been set.
func (o *Weapon) HasAnkamaId() bool {
	if o != nil && o.AnkamaId != nil {
		return true
	}

	return false
}

// SetAnkamaId gets a reference to the given int32 and assigns it to the AnkamaId field.
func (o *Weapon) SetAnkamaId(v int32) {
	o.AnkamaId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Weapon) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Weapon) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Weapon) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Weapon) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Weapon) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Weapon) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Weapon) GetType() ItemsListEntryTypedType {
	if o == nil || o.Type == nil {
		var ret ItemsListEntryTypedType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetTypeOk() (*ItemsListEntryTypedType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Weapon) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given ItemsListEntryTypedType and assigns it to the Type field.
func (o *Weapon) SetType(v ItemsListEntryTypedType) {
	o.Type = &v
}

// GetIsWeapon returns the IsWeapon field value if set, zero value otherwise.
func (o *Weapon) GetIsWeapon() bool {
	if o == nil || o.IsWeapon == nil {
		var ret bool
		return ret
	}
	return *o.IsWeapon
}

// GetIsWeaponOk returns a tuple with the IsWeapon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetIsWeaponOk() (*bool, bool) {
	if o == nil || o.IsWeapon == nil {
		return nil, false
	}
	return o.IsWeapon, true
}

// HasIsWeapon returns a boolean if a field has been set.
func (o *Weapon) HasIsWeapon() bool {
	if o != nil && o.IsWeapon != nil {
		return true
	}

	return false
}

// SetIsWeapon gets a reference to the given bool and assigns it to the IsWeapon field.
func (o *Weapon) SetIsWeapon(v bool) {
	o.IsWeapon = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *Weapon) GetLevel() int32 {
	if o == nil || o.Level == nil {
		var ret int32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetLevelOk() (*int32, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *Weapon) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given int32 and assigns it to the Level field.
func (o *Weapon) SetLevel(v int32) {
	o.Level = &v
}

// GetPods returns the Pods field value if set, zero value otherwise.
func (o *Weapon) GetPods() int32 {
	if o == nil || o.Pods == nil {
		var ret int32
		return ret
	}
	return *o.Pods
}

// GetPodsOk returns a tuple with the Pods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetPodsOk() (*int32, bool) {
	if o == nil || o.Pods == nil {
		return nil, false
	}
	return o.Pods, true
}

// HasPods returns a boolean if a field has been set.
func (o *Weapon) HasPods() bool {
	if o != nil && o.Pods != nil {
		return true
	}

	return false
}

// SetPods gets a reference to the given int32 and assigns it to the Pods field.
func (o *Weapon) SetPods(v int32) {
	o.Pods = &v
}

// GetImageUrls returns the ImageUrls field value if set, zero value otherwise.
func (o *Weapon) GetImageUrls() ImageUrls {
	if o == nil || o.ImageUrls == nil {
		var ret ImageUrls
		return ret
	}
	return *o.ImageUrls
}

// GetImageUrlsOk returns a tuple with the ImageUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetImageUrlsOk() (*ImageUrls, bool) {
	if o == nil || o.ImageUrls == nil {
		return nil, false
	}
	return o.ImageUrls, true
}

// HasImageUrls returns a boolean if a field has been set.
func (o *Weapon) HasImageUrls() bool {
	if o != nil && o.ImageUrls != nil {
		return true
	}

	return false
}

// SetImageUrls gets a reference to the given ImageUrls and assigns it to the ImageUrls field.
func (o *Weapon) SetImageUrls(v ImageUrls) {
	o.ImageUrls = &v
}

// GetHasEffects returns the HasEffects field value if set, zero value otherwise.
func (o *Weapon) GetHasEffects() bool {
	if o == nil || o.HasEffects == nil {
		var ret bool
		return ret
	}
	return *o.HasEffects
}

// GetHasEffectsOk returns a tuple with the HasEffects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetHasEffectsOk() (*bool, bool) {
	if o == nil || o.HasEffects == nil {
		return nil, false
	}
	return o.HasEffects, true
}

// HasHasEffects returns a boolean if a field has been set.
func (o *Weapon) HasHasEffects() bool {
	if o != nil && o.HasEffects != nil {
		return true
	}

	return false
}

// SetHasEffects gets a reference to the given bool and assigns it to the HasEffects field.
func (o *Weapon) SetHasEffects(v bool) {
	o.HasEffects = &v
}

// GetEffects returns the Effects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Weapon) GetEffects() []EffectsEntry {
	if o == nil {
		var ret []EffectsEntry
		return ret
	}
	return o.Effects
}

// GetEffectsOk returns a tuple with the Effects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Weapon) GetEffectsOk() ([]EffectsEntry, bool) {
	if o == nil || o.Effects == nil {
		return nil, false
	}
	return o.Effects, true
}

// HasEffects returns a boolean if a field has been set.
func (o *Weapon) HasEffects() bool {
	if o != nil && o.Effects != nil {
		return true
	}

	return false
}

// SetEffects gets a reference to the given []EffectsEntry and assigns it to the Effects field.
func (o *Weapon) SetEffects(v []EffectsEntry) {
	o.Effects = v
}

// GetHasConditions returns the HasConditions field value if set, zero value otherwise.
func (o *Weapon) GetHasConditions() bool {
	if o == nil || o.HasConditions == nil {
		var ret bool
		return ret
	}
	return *o.HasConditions
}

// GetHasConditionsOk returns a tuple with the HasConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetHasConditionsOk() (*bool, bool) {
	if o == nil || o.HasConditions == nil {
		return nil, false
	}
	return o.HasConditions, true
}

// HasHasConditions returns a boolean if a field has been set.
func (o *Weapon) HasHasConditions() bool {
	if o != nil && o.HasConditions != nil {
		return true
	}

	return false
}

// SetHasConditions gets a reference to the given bool and assigns it to the HasConditions field.
func (o *Weapon) SetHasConditions(v bool) {
	o.HasConditions = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Weapon) GetConditions() []ConditionEntry {
	if o == nil {
		var ret []ConditionEntry
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Weapon) GetConditionsOk() ([]ConditionEntry, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *Weapon) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []ConditionEntry and assigns it to the Conditions field.
func (o *Weapon) SetConditions(v []ConditionEntry) {
	o.Conditions = v
}

// GetCriticalHitProbability returns the CriticalHitProbability field value if set, zero value otherwise.
func (o *Weapon) GetCriticalHitProbability() int32 {
	if o == nil || o.CriticalHitProbability == nil {
		var ret int32
		return ret
	}
	return *o.CriticalHitProbability
}

// GetCriticalHitProbabilityOk returns a tuple with the CriticalHitProbability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetCriticalHitProbabilityOk() (*int32, bool) {
	if o == nil || o.CriticalHitProbability == nil {
		return nil, false
	}
	return o.CriticalHitProbability, true
}

// HasCriticalHitProbability returns a boolean if a field has been set.
func (o *Weapon) HasCriticalHitProbability() bool {
	if o != nil && o.CriticalHitProbability != nil {
		return true
	}

	return false
}

// SetCriticalHitProbability gets a reference to the given int32 and assigns it to the CriticalHitProbability field.
func (o *Weapon) SetCriticalHitProbability(v int32) {
	o.CriticalHitProbability = &v
}

// GetCriticalHitBonus returns the CriticalHitBonus field value if set, zero value otherwise.
func (o *Weapon) GetCriticalHitBonus() int32 {
	if o == nil || o.CriticalHitBonus == nil {
		var ret int32
		return ret
	}
	return *o.CriticalHitBonus
}

// GetCriticalHitBonusOk returns a tuple with the CriticalHitBonus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetCriticalHitBonusOk() (*int32, bool) {
	if o == nil || o.CriticalHitBonus == nil {
		return nil, false
	}
	return o.CriticalHitBonus, true
}

// HasCriticalHitBonus returns a boolean if a field has been set.
func (o *Weapon) HasCriticalHitBonus() bool {
	if o != nil && o.CriticalHitBonus != nil {
		return true
	}

	return false
}

// SetCriticalHitBonus gets a reference to the given int32 and assigns it to the CriticalHitBonus field.
func (o *Weapon) SetCriticalHitBonus(v int32) {
	o.CriticalHitBonus = &v
}

// GetIsTwoHanded returns the IsTwoHanded field value if set, zero value otherwise.
func (o *Weapon) GetIsTwoHanded() bool {
	if o == nil || o.IsTwoHanded == nil {
		var ret bool
		return ret
	}
	return *o.IsTwoHanded
}

// GetIsTwoHandedOk returns a tuple with the IsTwoHanded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetIsTwoHandedOk() (*bool, bool) {
	if o == nil || o.IsTwoHanded == nil {
		return nil, false
	}
	return o.IsTwoHanded, true
}

// HasIsTwoHanded returns a boolean if a field has been set.
func (o *Weapon) HasIsTwoHanded() bool {
	if o != nil && o.IsTwoHanded != nil {
		return true
	}

	return false
}

// SetIsTwoHanded gets a reference to the given bool and assigns it to the IsTwoHanded field.
func (o *Weapon) SetIsTwoHanded(v bool) {
	o.IsTwoHanded = &v
}

// GetMaxCastPerTurn returns the MaxCastPerTurn field value if set, zero value otherwise.
func (o *Weapon) GetMaxCastPerTurn() int32 {
	if o == nil || o.MaxCastPerTurn == nil {
		var ret int32
		return ret
	}
	return *o.MaxCastPerTurn
}

// GetMaxCastPerTurnOk returns a tuple with the MaxCastPerTurn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetMaxCastPerTurnOk() (*int32, bool) {
	if o == nil || o.MaxCastPerTurn == nil {
		return nil, false
	}
	return o.MaxCastPerTurn, true
}

// HasMaxCastPerTurn returns a boolean if a field has been set.
func (o *Weapon) HasMaxCastPerTurn() bool {
	if o != nil && o.MaxCastPerTurn != nil {
		return true
	}

	return false
}

// SetMaxCastPerTurn gets a reference to the given int32 and assigns it to the MaxCastPerTurn field.
func (o *Weapon) SetMaxCastPerTurn(v int32) {
	o.MaxCastPerTurn = &v
}

// GetApCost returns the ApCost field value if set, zero value otherwise.
func (o *Weapon) GetApCost() int32 {
	if o == nil || o.ApCost == nil {
		var ret int32
		return ret
	}
	return *o.ApCost
}

// GetApCostOk returns a tuple with the ApCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetApCostOk() (*int32, bool) {
	if o == nil || o.ApCost == nil {
		return nil, false
	}
	return o.ApCost, true
}

// HasApCost returns a boolean if a field has been set.
func (o *Weapon) HasApCost() bool {
	if o != nil && o.ApCost != nil {
		return true
	}

	return false
}

// SetApCost gets a reference to the given int32 and assigns it to the ApCost field.
func (o *Weapon) SetApCost(v int32) {
	o.ApCost = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *Weapon) GetRange() WeaponRange {
	if o == nil || o.Range == nil {
		var ret WeaponRange
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetRangeOk() (*WeaponRange, bool) {
	if o == nil || o.Range == nil {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *Weapon) HasRange() bool {
	if o != nil && o.Range != nil {
		return true
	}

	return false
}

// SetRange gets a reference to the given WeaponRange and assigns it to the Range field.
func (o *Weapon) SetRange(v WeaponRange) {
	o.Range = &v
}

// GetHasRecipe returns the HasRecipe field value if set, zero value otherwise.
func (o *Weapon) GetHasRecipe() bool {
	if o == nil || o.HasRecipe == nil {
		var ret bool
		return ret
	}
	return *o.HasRecipe
}

// GetHasRecipeOk returns a tuple with the HasRecipe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetHasRecipeOk() (*bool, bool) {
	if o == nil || o.HasRecipe == nil {
		return nil, false
	}
	return o.HasRecipe, true
}

// HasHasRecipe returns a boolean if a field has been set.
func (o *Weapon) HasHasRecipe() bool {
	if o != nil && o.HasRecipe != nil {
		return true
	}

	return false
}

// SetHasRecipe gets a reference to the given bool and assigns it to the HasRecipe field.
func (o *Weapon) SetHasRecipe(v bool) {
	o.HasRecipe = &v
}

// GetRecipe returns the Recipe field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Weapon) GetRecipe() []RecipeEntry {
	if o == nil {
		var ret []RecipeEntry
		return ret
	}
	return o.Recipe
}

// GetRecipeOk returns a tuple with the Recipe field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Weapon) GetRecipeOk() ([]RecipeEntry, bool) {
	if o == nil || o.Recipe == nil {
		return nil, false
	}
	return o.Recipe, true
}

// HasRecipe returns a boolean if a field has been set.
func (o *Weapon) HasRecipe() bool {
	if o != nil && o.Recipe != nil {
		return true
	}

	return false
}

// SetRecipe gets a reference to the given []RecipeEntry and assigns it to the Recipe field.
func (o *Weapon) SetRecipe(v []RecipeEntry) {
	o.Recipe = v
}

// GetHasParentSet returns the HasParentSet field value if set, zero value otherwise.
func (o *Weapon) GetHasParentSet() bool {
	if o == nil || o.HasParentSet == nil {
		var ret bool
		return ret
	}
	return *o.HasParentSet
}

// GetHasParentSetOk returns a tuple with the HasParentSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Weapon) GetHasParentSetOk() (*bool, bool) {
	if o == nil || o.HasParentSet == nil {
		return nil, false
	}
	return o.HasParentSet, true
}

// HasHasParentSet returns a boolean if a field has been set.
func (o *Weapon) HasHasParentSet() bool {
	if o != nil && o.HasParentSet != nil {
		return true
	}

	return false
}

// SetHasParentSet gets a reference to the given bool and assigns it to the HasParentSet field.
func (o *Weapon) SetHasParentSet(v bool) {
	o.HasParentSet = &v
}

// GetParentSet returns the ParentSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Weapon) GetParentSet() EquipmentParentSet {
	if o == nil || o.ParentSet.Get() == nil {
		var ret EquipmentParentSet
		return ret
	}
	return *o.ParentSet.Get()
}

// GetParentSetOk returns a tuple with the ParentSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Weapon) GetParentSetOk() (*EquipmentParentSet, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentSet.Get(), o.ParentSet.IsSet()
}

// HasParentSet returns a boolean if a field has been set.
func (o *Weapon) HasParentSet() bool {
	if o != nil && o.ParentSet.IsSet() {
		return true
	}

	return false
}

// SetParentSet gets a reference to the given NullableEquipmentParentSet and assigns it to the ParentSet field.
func (o *Weapon) SetParentSet(v EquipmentParentSet) {
	o.ParentSet.Set(&v)
}
// SetParentSetNil sets the value for ParentSet to be an explicit nil
func (o *Weapon) SetParentSetNil() {
	o.ParentSet.Set(nil)
}

// UnsetParentSet ensures that no value is present for ParentSet, not even an explicit nil
func (o *Weapon) UnsetParentSet() {
	o.ParentSet.Unset()
}

func (o Weapon) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AnkamaId != nil {
		toSerialize["ankama_id"] = o.AnkamaId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.IsWeapon != nil {
		toSerialize["is_weapon"] = o.IsWeapon
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.Pods != nil {
		toSerialize["pods"] = o.Pods
	}
	if o.ImageUrls != nil {
		toSerialize["image_urls"] = o.ImageUrls
	}
	if o.HasEffects != nil {
		toSerialize["has_effects"] = o.HasEffects
	}
	if o.Effects != nil {
		toSerialize["effects"] = o.Effects
	}
	if o.HasConditions != nil {
		toSerialize["has_conditions"] = o.HasConditions
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.CriticalHitProbability != nil {
		toSerialize["critical_hit_probability"] = o.CriticalHitProbability
	}
	if o.CriticalHitBonus != nil {
		toSerialize["critical_hit_bonus"] = o.CriticalHitBonus
	}
	if o.IsTwoHanded != nil {
		toSerialize["is_two_handed"] = o.IsTwoHanded
	}
	if o.MaxCastPerTurn != nil {
		toSerialize["max_cast_per_turn"] = o.MaxCastPerTurn
	}
	if o.ApCost != nil {
		toSerialize["ap_cost"] = o.ApCost
	}
	if o.Range != nil {
		toSerialize["range"] = o.Range
	}
	if o.HasRecipe != nil {
		toSerialize["has_recipe"] = o.HasRecipe
	}
	if o.Recipe != nil {
		toSerialize["recipe"] = o.Recipe
	}
	if o.HasParentSet != nil {
		toSerialize["has_parent_set"] = o.HasParentSet
	}
	if o.ParentSet.IsSet() {
		toSerialize["parent_set"] = o.ParentSet.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableWeapon struct {
	value *Weapon
	isSet bool
}

func (v NullableWeapon) Get() *Weapon {
	return v.value
}

func (v *NullableWeapon) Set(val *Weapon) {
	v.value = val
	v.isSet = true
}

func (v NullableWeapon) IsSet() bool {
	return v.isSet
}

func (v *NullableWeapon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeapon(val *Weapon) *NullableWeapon {
	return &NullableWeapon{value: val, isSet: true}
}

func (v NullableWeapon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeapon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

