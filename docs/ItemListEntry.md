# ItemListEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ItemsListEntryTypedType**](ItemsListEntryTypedType.md) |  | [optional] 
**Level** | Pointer to **int32** |  | [optional] 
**ImageUrls** | Pointer to [**ImageUrls**](ImageUrls.md) |  | [optional] 
**Recipe** | Pointer to [**[]RecipeEntry**](RecipeEntry.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Conditions** | Pointer to [**[]ConditionEntry**](ConditionEntry.md) |  | [optional] 
**ConditionTree** | Pointer to [**ConditionTreeNode**](ConditionTreeNode.md) |  | [optional] 
**Effects** | Pointer to [**[]EffectsEntry**](EffectsEntry.md) |  | [optional] 
**IsWeapon** | Pointer to **NullableBool** |  | [optional] 
**Pods** | Pointer to **NullableInt32** |  | [optional] 
**ParentSet** | Pointer to [**NullableItemListEntryParentSet**](ItemListEntryParentSet.md) |  | [optional] 
**CriticalHitProbability** | Pointer to **NullableInt32** |  | [optional] 
**CriticalHitBonus** | Pointer to **NullableInt32** |  | [optional] 
**IsTwoHanded** | Pointer to **NullableBool** |  | [optional] 
**MaxCastPerTurn** | Pointer to **NullableInt32** |  | [optional] 
**ApCost** | Pointer to **NullableInt32** |  | [optional] 
**Range** | Pointer to [**NullableItemListEntryRange**](ItemListEntryRange.md) |  | [optional] 

## Methods

### NewItemListEntry

`func NewItemListEntry() *ItemListEntry`

NewItemListEntry instantiates a new ItemListEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemListEntryWithDefaults

`func NewItemListEntryWithDefaults() *ItemListEntry`

NewItemListEntryWithDefaults instantiates a new ItemListEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *ItemListEntry) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *ItemListEntry) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *ItemListEntry) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.

### HasAnkamaId

`func (o *ItemListEntry) HasAnkamaId() bool`

HasAnkamaId returns a boolean if a field has been set.

### GetName

`func (o *ItemListEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ItemListEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ItemListEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ItemListEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ItemListEntry) GetType() ItemsListEntryTypedType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ItemListEntry) GetTypeOk() (*ItemsListEntryTypedType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ItemListEntry) SetType(v ItemsListEntryTypedType)`

SetType sets Type field to given value.

### HasType

`func (o *ItemListEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLevel

`func (o *ItemListEntry) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ItemListEntry) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ItemListEntry) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ItemListEntry) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetImageUrls

`func (o *ItemListEntry) GetImageUrls() ImageUrls`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *ItemListEntry) GetImageUrlsOk() (*ImageUrls, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *ItemListEntry) SetImageUrls(v ImageUrls)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *ItemListEntry) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.

### GetRecipe

`func (o *ItemListEntry) GetRecipe() []RecipeEntry`

GetRecipe returns the Recipe field if non-nil, zero value otherwise.

### GetRecipeOk

`func (o *ItemListEntry) GetRecipeOk() (*[]RecipeEntry, bool)`

GetRecipeOk returns a tuple with the Recipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipe

`func (o *ItemListEntry) SetRecipe(v []RecipeEntry)`

SetRecipe sets Recipe field to given value.

### HasRecipe

`func (o *ItemListEntry) HasRecipe() bool`

HasRecipe returns a boolean if a field has been set.

### SetRecipeNil

`func (o *ItemListEntry) SetRecipeNil(b bool)`

 SetRecipeNil sets the value for Recipe to be an explicit nil

### UnsetRecipe
`func (o *ItemListEntry) UnsetRecipe()`

UnsetRecipe ensures that no value is present for Recipe, not even an explicit nil
### GetDescription

`func (o *ItemListEntry) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ItemListEntry) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ItemListEntry) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ItemListEntry) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ItemListEntry) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ItemListEntry) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetConditions

`func (o *ItemListEntry) GetConditions() []ConditionEntry`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ItemListEntry) GetConditionsOk() (*[]ConditionEntry, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ItemListEntry) SetConditions(v []ConditionEntry)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ItemListEntry) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *ItemListEntry) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *ItemListEntry) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetConditionTree

`func (o *ItemListEntry) GetConditionTree() ConditionTreeNode`

GetConditionTree returns the ConditionTree field if non-nil, zero value otherwise.

### GetConditionTreeOk

`func (o *ItemListEntry) GetConditionTreeOk() (*ConditionTreeNode, bool)`

GetConditionTreeOk returns a tuple with the ConditionTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionTree

`func (o *ItemListEntry) SetConditionTree(v ConditionTreeNode)`

SetConditionTree sets ConditionTree field to given value.

### HasConditionTree

`func (o *ItemListEntry) HasConditionTree() bool`

HasConditionTree returns a boolean if a field has been set.

### GetEffects

`func (o *ItemListEntry) GetEffects() []EffectsEntry`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *ItemListEntry) GetEffectsOk() (*[]EffectsEntry, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *ItemListEntry) SetEffects(v []EffectsEntry)`

SetEffects sets Effects field to given value.

### HasEffects

`func (o *ItemListEntry) HasEffects() bool`

HasEffects returns a boolean if a field has been set.

### SetEffectsNil

`func (o *ItemListEntry) SetEffectsNil(b bool)`

 SetEffectsNil sets the value for Effects to be an explicit nil

### UnsetEffects
`func (o *ItemListEntry) UnsetEffects()`

UnsetEffects ensures that no value is present for Effects, not even an explicit nil
### GetIsWeapon

`func (o *ItemListEntry) GetIsWeapon() bool`

GetIsWeapon returns the IsWeapon field if non-nil, zero value otherwise.

### GetIsWeaponOk

`func (o *ItemListEntry) GetIsWeaponOk() (*bool, bool)`

GetIsWeaponOk returns a tuple with the IsWeapon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWeapon

`func (o *ItemListEntry) SetIsWeapon(v bool)`

SetIsWeapon sets IsWeapon field to given value.

### HasIsWeapon

`func (o *ItemListEntry) HasIsWeapon() bool`

HasIsWeapon returns a boolean if a field has been set.

### SetIsWeaponNil

`func (o *ItemListEntry) SetIsWeaponNil(b bool)`

 SetIsWeaponNil sets the value for IsWeapon to be an explicit nil

### UnsetIsWeapon
`func (o *ItemListEntry) UnsetIsWeapon()`

UnsetIsWeapon ensures that no value is present for IsWeapon, not even an explicit nil
### GetPods

`func (o *ItemListEntry) GetPods() int32`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *ItemListEntry) GetPodsOk() (*int32, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *ItemListEntry) SetPods(v int32)`

SetPods sets Pods field to given value.

### HasPods

`func (o *ItemListEntry) HasPods() bool`

HasPods returns a boolean if a field has been set.

### SetPodsNil

`func (o *ItemListEntry) SetPodsNil(b bool)`

 SetPodsNil sets the value for Pods to be an explicit nil

### UnsetPods
`func (o *ItemListEntry) UnsetPods()`

UnsetPods ensures that no value is present for Pods, not even an explicit nil
### GetParentSet

`func (o *ItemListEntry) GetParentSet() ItemListEntryParentSet`

GetParentSet returns the ParentSet field if non-nil, zero value otherwise.

### GetParentSetOk

`func (o *ItemListEntry) GetParentSetOk() (*ItemListEntryParentSet, bool)`

GetParentSetOk returns a tuple with the ParentSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSet

`func (o *ItemListEntry) SetParentSet(v ItemListEntryParentSet)`

SetParentSet sets ParentSet field to given value.

### HasParentSet

`func (o *ItemListEntry) HasParentSet() bool`

HasParentSet returns a boolean if a field has been set.

### SetParentSetNil

`func (o *ItemListEntry) SetParentSetNil(b bool)`

 SetParentSetNil sets the value for ParentSet to be an explicit nil

### UnsetParentSet
`func (o *ItemListEntry) UnsetParentSet()`

UnsetParentSet ensures that no value is present for ParentSet, not even an explicit nil
### GetCriticalHitProbability

`func (o *ItemListEntry) GetCriticalHitProbability() int32`

GetCriticalHitProbability returns the CriticalHitProbability field if non-nil, zero value otherwise.

### GetCriticalHitProbabilityOk

`func (o *ItemListEntry) GetCriticalHitProbabilityOk() (*int32, bool)`

GetCriticalHitProbabilityOk returns a tuple with the CriticalHitProbability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalHitProbability

`func (o *ItemListEntry) SetCriticalHitProbability(v int32)`

SetCriticalHitProbability sets CriticalHitProbability field to given value.

### HasCriticalHitProbability

`func (o *ItemListEntry) HasCriticalHitProbability() bool`

HasCriticalHitProbability returns a boolean if a field has been set.

### SetCriticalHitProbabilityNil

`func (o *ItemListEntry) SetCriticalHitProbabilityNil(b bool)`

 SetCriticalHitProbabilityNil sets the value for CriticalHitProbability to be an explicit nil

### UnsetCriticalHitProbability
`func (o *ItemListEntry) UnsetCriticalHitProbability()`

UnsetCriticalHitProbability ensures that no value is present for CriticalHitProbability, not even an explicit nil
### GetCriticalHitBonus

`func (o *ItemListEntry) GetCriticalHitBonus() int32`

GetCriticalHitBonus returns the CriticalHitBonus field if non-nil, zero value otherwise.

### GetCriticalHitBonusOk

`func (o *ItemListEntry) GetCriticalHitBonusOk() (*int32, bool)`

GetCriticalHitBonusOk returns a tuple with the CriticalHitBonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalHitBonus

`func (o *ItemListEntry) SetCriticalHitBonus(v int32)`

SetCriticalHitBonus sets CriticalHitBonus field to given value.

### HasCriticalHitBonus

`func (o *ItemListEntry) HasCriticalHitBonus() bool`

HasCriticalHitBonus returns a boolean if a field has been set.

### SetCriticalHitBonusNil

`func (o *ItemListEntry) SetCriticalHitBonusNil(b bool)`

 SetCriticalHitBonusNil sets the value for CriticalHitBonus to be an explicit nil

### UnsetCriticalHitBonus
`func (o *ItemListEntry) UnsetCriticalHitBonus()`

UnsetCriticalHitBonus ensures that no value is present for CriticalHitBonus, not even an explicit nil
### GetIsTwoHanded

`func (o *ItemListEntry) GetIsTwoHanded() bool`

GetIsTwoHanded returns the IsTwoHanded field if non-nil, zero value otherwise.

### GetIsTwoHandedOk

`func (o *ItemListEntry) GetIsTwoHandedOk() (*bool, bool)`

GetIsTwoHandedOk returns a tuple with the IsTwoHanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTwoHanded

`func (o *ItemListEntry) SetIsTwoHanded(v bool)`

SetIsTwoHanded sets IsTwoHanded field to given value.

### HasIsTwoHanded

`func (o *ItemListEntry) HasIsTwoHanded() bool`

HasIsTwoHanded returns a boolean if a field has been set.

### SetIsTwoHandedNil

`func (o *ItemListEntry) SetIsTwoHandedNil(b bool)`

 SetIsTwoHandedNil sets the value for IsTwoHanded to be an explicit nil

### UnsetIsTwoHanded
`func (o *ItemListEntry) UnsetIsTwoHanded()`

UnsetIsTwoHanded ensures that no value is present for IsTwoHanded, not even an explicit nil
### GetMaxCastPerTurn

`func (o *ItemListEntry) GetMaxCastPerTurn() int32`

GetMaxCastPerTurn returns the MaxCastPerTurn field if non-nil, zero value otherwise.

### GetMaxCastPerTurnOk

`func (o *ItemListEntry) GetMaxCastPerTurnOk() (*int32, bool)`

GetMaxCastPerTurnOk returns a tuple with the MaxCastPerTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCastPerTurn

`func (o *ItemListEntry) SetMaxCastPerTurn(v int32)`

SetMaxCastPerTurn sets MaxCastPerTurn field to given value.

### HasMaxCastPerTurn

`func (o *ItemListEntry) HasMaxCastPerTurn() bool`

HasMaxCastPerTurn returns a boolean if a field has been set.

### SetMaxCastPerTurnNil

`func (o *ItemListEntry) SetMaxCastPerTurnNil(b bool)`

 SetMaxCastPerTurnNil sets the value for MaxCastPerTurn to be an explicit nil

### UnsetMaxCastPerTurn
`func (o *ItemListEntry) UnsetMaxCastPerTurn()`

UnsetMaxCastPerTurn ensures that no value is present for MaxCastPerTurn, not even an explicit nil
### GetApCost

`func (o *ItemListEntry) GetApCost() int32`

GetApCost returns the ApCost field if non-nil, zero value otherwise.

### GetApCostOk

`func (o *ItemListEntry) GetApCostOk() (*int32, bool)`

GetApCostOk returns a tuple with the ApCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApCost

`func (o *ItemListEntry) SetApCost(v int32)`

SetApCost sets ApCost field to given value.

### HasApCost

`func (o *ItemListEntry) HasApCost() bool`

HasApCost returns a boolean if a field has been set.

### SetApCostNil

`func (o *ItemListEntry) SetApCostNil(b bool)`

 SetApCostNil sets the value for ApCost to be an explicit nil

### UnsetApCost
`func (o *ItemListEntry) UnsetApCost()`

UnsetApCost ensures that no value is present for ApCost, not even an explicit nil
### GetRange

`func (o *ItemListEntry) GetRange() ItemListEntryRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *ItemListEntry) GetRangeOk() (*ItemListEntryRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *ItemListEntry) SetRange(v ItemListEntryRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *ItemListEntry) HasRange() bool`

HasRange returns a boolean if a field has been set.

### SetRangeNil

`func (o *ItemListEntry) SetRangeNil(b bool)`

 SetRangeNil sets the value for Range to be an explicit nil

### UnsetRange
`func (o *ItemListEntry) UnsetRange()`

UnsetRange ensures that no value is present for Range, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


