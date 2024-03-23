# Weapon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ItemsListEntryTypedType**](ItemsListEntryTypedType.md) |  | [optional] 
**IsWeapon** | Pointer to **bool** | always true when the item is a weapon. Many fields are now available. Always check for this flag first when getting single equipment items. | [optional] 
**Level** | Pointer to **int32** |  | [optional] 
**Pods** | Pointer to **int32** |  | [optional] 
**ImageUrls** | Pointer to [**ImageUrls**](ImageUrls.md) |  | [optional] 
**Effects** | Pointer to [**[]EffectsEntry**](EffectsEntry.md) |  | [optional] 
**Conditions** | Pointer to [**[]ConditionEntry**](ConditionEntry.md) |  | [optional] 
**ConditionTree** | Pointer to [**ConditionTreeNode**](ConditionTreeNode.md) |  | [optional] 
**CriticalHitProbability** | Pointer to **int32** |  | [optional] 
**CriticalHitBonus** | Pointer to **int32** |  | [optional] 
**IsTwoHanded** | Pointer to **bool** |  | [optional] 
**MaxCastPerTurn** | Pointer to **int32** |  | [optional] 
**ApCost** | Pointer to **int32** |  | [optional] 
**Range** | Pointer to [**WeaponRange**](WeaponRange.md) |  | [optional] 
**Recipe** | Pointer to [**[]RecipeEntry**](RecipeEntry.md) |  | [optional] 
**ParentSet** | Pointer to [**NullableItemListEntryParentSet**](ItemListEntryParentSet.md) |  | [optional] 

## Methods

### NewWeapon

`func NewWeapon() *Weapon`

NewWeapon instantiates a new Weapon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeaponWithDefaults

`func NewWeaponWithDefaults() *Weapon`

NewWeaponWithDefaults instantiates a new Weapon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *Weapon) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *Weapon) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *Weapon) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.

### HasAnkamaId

`func (o *Weapon) HasAnkamaId() bool`

HasAnkamaId returns a boolean if a field has been set.

### GetName

`func (o *Weapon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Weapon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Weapon) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Weapon) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Weapon) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Weapon) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Weapon) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Weapon) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *Weapon) GetType() ItemsListEntryTypedType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Weapon) GetTypeOk() (*ItemsListEntryTypedType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Weapon) SetType(v ItemsListEntryTypedType)`

SetType sets Type field to given value.

### HasType

`func (o *Weapon) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsWeapon

`func (o *Weapon) GetIsWeapon() bool`

GetIsWeapon returns the IsWeapon field if non-nil, zero value otherwise.

### GetIsWeaponOk

`func (o *Weapon) GetIsWeaponOk() (*bool, bool)`

GetIsWeaponOk returns a tuple with the IsWeapon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWeapon

`func (o *Weapon) SetIsWeapon(v bool)`

SetIsWeapon sets IsWeapon field to given value.

### HasIsWeapon

`func (o *Weapon) HasIsWeapon() bool`

HasIsWeapon returns a boolean if a field has been set.

### GetLevel

`func (o *Weapon) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *Weapon) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *Weapon) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *Weapon) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetPods

`func (o *Weapon) GetPods() int32`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *Weapon) GetPodsOk() (*int32, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *Weapon) SetPods(v int32)`

SetPods sets Pods field to given value.

### HasPods

`func (o *Weapon) HasPods() bool`

HasPods returns a boolean if a field has been set.

### GetImageUrls

`func (o *Weapon) GetImageUrls() ImageUrls`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *Weapon) GetImageUrlsOk() (*ImageUrls, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *Weapon) SetImageUrls(v ImageUrls)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *Weapon) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.

### GetEffects

`func (o *Weapon) GetEffects() []EffectsEntry`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *Weapon) GetEffectsOk() (*[]EffectsEntry, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *Weapon) SetEffects(v []EffectsEntry)`

SetEffects sets Effects field to given value.

### HasEffects

`func (o *Weapon) HasEffects() bool`

HasEffects returns a boolean if a field has been set.

### SetEffectsNil

`func (o *Weapon) SetEffectsNil(b bool)`

 SetEffectsNil sets the value for Effects to be an explicit nil

### UnsetEffects
`func (o *Weapon) UnsetEffects()`

UnsetEffects ensures that no value is present for Effects, not even an explicit nil
### GetConditions

`func (o *Weapon) GetConditions() []ConditionEntry`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *Weapon) GetConditionsOk() (*[]ConditionEntry, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *Weapon) SetConditions(v []ConditionEntry)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *Weapon) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *Weapon) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *Weapon) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetConditionTree

`func (o *Weapon) GetConditionTree() ConditionTreeNode`

GetConditionTree returns the ConditionTree field if non-nil, zero value otherwise.

### GetConditionTreeOk

`func (o *Weapon) GetConditionTreeOk() (*ConditionTreeNode, bool)`

GetConditionTreeOk returns a tuple with the ConditionTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionTree

`func (o *Weapon) SetConditionTree(v ConditionTreeNode)`

SetConditionTree sets ConditionTree field to given value.

### HasConditionTree

`func (o *Weapon) HasConditionTree() bool`

HasConditionTree returns a boolean if a field has been set.

### GetCriticalHitProbability

`func (o *Weapon) GetCriticalHitProbability() int32`

GetCriticalHitProbability returns the CriticalHitProbability field if non-nil, zero value otherwise.

### GetCriticalHitProbabilityOk

`func (o *Weapon) GetCriticalHitProbabilityOk() (*int32, bool)`

GetCriticalHitProbabilityOk returns a tuple with the CriticalHitProbability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalHitProbability

`func (o *Weapon) SetCriticalHitProbability(v int32)`

SetCriticalHitProbability sets CriticalHitProbability field to given value.

### HasCriticalHitProbability

`func (o *Weapon) HasCriticalHitProbability() bool`

HasCriticalHitProbability returns a boolean if a field has been set.

### GetCriticalHitBonus

`func (o *Weapon) GetCriticalHitBonus() int32`

GetCriticalHitBonus returns the CriticalHitBonus field if non-nil, zero value otherwise.

### GetCriticalHitBonusOk

`func (o *Weapon) GetCriticalHitBonusOk() (*int32, bool)`

GetCriticalHitBonusOk returns a tuple with the CriticalHitBonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalHitBonus

`func (o *Weapon) SetCriticalHitBonus(v int32)`

SetCriticalHitBonus sets CriticalHitBonus field to given value.

### HasCriticalHitBonus

`func (o *Weapon) HasCriticalHitBonus() bool`

HasCriticalHitBonus returns a boolean if a field has been set.

### GetIsTwoHanded

`func (o *Weapon) GetIsTwoHanded() bool`

GetIsTwoHanded returns the IsTwoHanded field if non-nil, zero value otherwise.

### GetIsTwoHandedOk

`func (o *Weapon) GetIsTwoHandedOk() (*bool, bool)`

GetIsTwoHandedOk returns a tuple with the IsTwoHanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTwoHanded

`func (o *Weapon) SetIsTwoHanded(v bool)`

SetIsTwoHanded sets IsTwoHanded field to given value.

### HasIsTwoHanded

`func (o *Weapon) HasIsTwoHanded() bool`

HasIsTwoHanded returns a boolean if a field has been set.

### GetMaxCastPerTurn

`func (o *Weapon) GetMaxCastPerTurn() int32`

GetMaxCastPerTurn returns the MaxCastPerTurn field if non-nil, zero value otherwise.

### GetMaxCastPerTurnOk

`func (o *Weapon) GetMaxCastPerTurnOk() (*int32, bool)`

GetMaxCastPerTurnOk returns a tuple with the MaxCastPerTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCastPerTurn

`func (o *Weapon) SetMaxCastPerTurn(v int32)`

SetMaxCastPerTurn sets MaxCastPerTurn field to given value.

### HasMaxCastPerTurn

`func (o *Weapon) HasMaxCastPerTurn() bool`

HasMaxCastPerTurn returns a boolean if a field has been set.

### GetApCost

`func (o *Weapon) GetApCost() int32`

GetApCost returns the ApCost field if non-nil, zero value otherwise.

### GetApCostOk

`func (o *Weapon) GetApCostOk() (*int32, bool)`

GetApCostOk returns a tuple with the ApCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApCost

`func (o *Weapon) SetApCost(v int32)`

SetApCost sets ApCost field to given value.

### HasApCost

`func (o *Weapon) HasApCost() bool`

HasApCost returns a boolean if a field has been set.

### GetRange

`func (o *Weapon) GetRange() WeaponRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *Weapon) GetRangeOk() (*WeaponRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *Weapon) SetRange(v WeaponRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *Weapon) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetRecipe

`func (o *Weapon) GetRecipe() []RecipeEntry`

GetRecipe returns the Recipe field if non-nil, zero value otherwise.

### GetRecipeOk

`func (o *Weapon) GetRecipeOk() (*[]RecipeEntry, bool)`

GetRecipeOk returns a tuple with the Recipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipe

`func (o *Weapon) SetRecipe(v []RecipeEntry)`

SetRecipe sets Recipe field to given value.

### HasRecipe

`func (o *Weapon) HasRecipe() bool`

HasRecipe returns a boolean if a field has been set.

### SetRecipeNil

`func (o *Weapon) SetRecipeNil(b bool)`

 SetRecipeNil sets the value for Recipe to be an explicit nil

### UnsetRecipe
`func (o *Weapon) UnsetRecipe()`

UnsetRecipe ensures that no value is present for Recipe, not even an explicit nil
### GetParentSet

`func (o *Weapon) GetParentSet() ItemListEntryParentSet`

GetParentSet returns the ParentSet field if non-nil, zero value otherwise.

### GetParentSetOk

`func (o *Weapon) GetParentSetOk() (*ItemListEntryParentSet, bool)`

GetParentSetOk returns a tuple with the ParentSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSet

`func (o *Weapon) SetParentSet(v ItemListEntryParentSet)`

SetParentSet sets ParentSet field to given value.

### HasParentSet

`func (o *Weapon) HasParentSet() bool`

HasParentSet returns a boolean if a field has been set.

### SetParentSetNil

`func (o *Weapon) SetParentSetNil(b bool)`

 SetParentSetNil sets the value for ParentSet to be an explicit nil

### UnsetParentSet
`func (o *Weapon) UnsetParentSet()`

UnsetParentSet ensures that no value is present for ParentSet, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


