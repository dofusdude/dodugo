# ListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**TranslatedId**](TranslatedId.md) |  | [optional] 
**Level** | Pointer to **int32** |  | [optional] 
**ImageUrls** | Pointer to [**Images**](Images.md) |  | [optional] 
**Recipe** | Pointer to [**[]Recipe**](Recipe.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Conditions** | Pointer to [**NullableConditionNode**](ConditionNode.md) |  | [optional] 
**Effects** | Pointer to [**[]Effect**](Effect.md) |  | [optional] 
**IsWeapon** | Pointer to **NullableBool** |  | [optional] 
**Pods** | Pointer to **NullableInt32** |  | [optional] 
**ParentSet** | Pointer to [**NullableTranslatedId**](TranslatedId.md) |  | [optional] 
**CriticalHitProbability** | Pointer to **NullableInt32** |  | [optional] 
**CriticalHitBonus** | Pointer to **NullableInt32** |  | [optional] 
**MaxCastPerTurn** | Pointer to **NullableInt32** |  | [optional] 
**ApCost** | Pointer to **NullableInt32** |  | [optional] 
**Range** | Pointer to [**Range**](Range.md) |  | [optional] 

## Methods

### NewListItem

`func NewListItem() *ListItem`

NewListItem instantiates a new ListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListItemWithDefaults

`func NewListItemWithDefaults() *ListItem`

NewListItemWithDefaults instantiates a new ListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *ListItem) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *ListItem) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *ListItem) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.

### HasAnkamaId

`func (o *ListItem) HasAnkamaId() bool`

HasAnkamaId returns a boolean if a field has been set.

### GetName

`func (o *ListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ListItem) GetType() TranslatedId`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListItem) GetTypeOk() (*TranslatedId, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListItem) SetType(v TranslatedId)`

SetType sets Type field to given value.

### HasType

`func (o *ListItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLevel

`func (o *ListItem) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ListItem) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ListItem) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ListItem) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetImageUrls

`func (o *ListItem) GetImageUrls() Images`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *ListItem) GetImageUrlsOk() (*Images, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *ListItem) SetImageUrls(v Images)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *ListItem) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.

### GetRecipe

`func (o *ListItem) GetRecipe() []Recipe`

GetRecipe returns the Recipe field if non-nil, zero value otherwise.

### GetRecipeOk

`func (o *ListItem) GetRecipeOk() (*[]Recipe, bool)`

GetRecipeOk returns a tuple with the Recipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipe

`func (o *ListItem) SetRecipe(v []Recipe)`

SetRecipe sets Recipe field to given value.

### HasRecipe

`func (o *ListItem) HasRecipe() bool`

HasRecipe returns a boolean if a field has been set.

### SetRecipeNil

`func (o *ListItem) SetRecipeNil(b bool)`

 SetRecipeNil sets the value for Recipe to be an explicit nil

### UnsetRecipe
`func (o *ListItem) UnsetRecipe()`

UnsetRecipe ensures that no value is present for Recipe, not even an explicit nil
### GetDescription

`func (o *ListItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetConditions

`func (o *ListItem) GetConditions() ConditionNode`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ListItem) GetConditionsOk() (*ConditionNode, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ListItem) SetConditions(v ConditionNode)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ListItem) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *ListItem) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *ListItem) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetEffects

`func (o *ListItem) GetEffects() []Effect`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *ListItem) GetEffectsOk() (*[]Effect, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *ListItem) SetEffects(v []Effect)`

SetEffects sets Effects field to given value.

### HasEffects

`func (o *ListItem) HasEffects() bool`

HasEffects returns a boolean if a field has been set.

### SetEffectsNil

`func (o *ListItem) SetEffectsNil(b bool)`

 SetEffectsNil sets the value for Effects to be an explicit nil

### UnsetEffects
`func (o *ListItem) UnsetEffects()`

UnsetEffects ensures that no value is present for Effects, not even an explicit nil
### GetIsWeapon

`func (o *ListItem) GetIsWeapon() bool`

GetIsWeapon returns the IsWeapon field if non-nil, zero value otherwise.

### GetIsWeaponOk

`func (o *ListItem) GetIsWeaponOk() (*bool, bool)`

GetIsWeaponOk returns a tuple with the IsWeapon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWeapon

`func (o *ListItem) SetIsWeapon(v bool)`

SetIsWeapon sets IsWeapon field to given value.

### HasIsWeapon

`func (o *ListItem) HasIsWeapon() bool`

HasIsWeapon returns a boolean if a field has been set.

### SetIsWeaponNil

`func (o *ListItem) SetIsWeaponNil(b bool)`

 SetIsWeaponNil sets the value for IsWeapon to be an explicit nil

### UnsetIsWeapon
`func (o *ListItem) UnsetIsWeapon()`

UnsetIsWeapon ensures that no value is present for IsWeapon, not even an explicit nil
### GetPods

`func (o *ListItem) GetPods() int32`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *ListItem) GetPodsOk() (*int32, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *ListItem) SetPods(v int32)`

SetPods sets Pods field to given value.

### HasPods

`func (o *ListItem) HasPods() bool`

HasPods returns a boolean if a field has been set.

### SetPodsNil

`func (o *ListItem) SetPodsNil(b bool)`

 SetPodsNil sets the value for Pods to be an explicit nil

### UnsetPods
`func (o *ListItem) UnsetPods()`

UnsetPods ensures that no value is present for Pods, not even an explicit nil
### GetParentSet

`func (o *ListItem) GetParentSet() TranslatedId`

GetParentSet returns the ParentSet field if non-nil, zero value otherwise.

### GetParentSetOk

`func (o *ListItem) GetParentSetOk() (*TranslatedId, bool)`

GetParentSetOk returns a tuple with the ParentSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSet

`func (o *ListItem) SetParentSet(v TranslatedId)`

SetParentSet sets ParentSet field to given value.

### HasParentSet

`func (o *ListItem) HasParentSet() bool`

HasParentSet returns a boolean if a field has been set.

### SetParentSetNil

`func (o *ListItem) SetParentSetNil(b bool)`

 SetParentSetNil sets the value for ParentSet to be an explicit nil

### UnsetParentSet
`func (o *ListItem) UnsetParentSet()`

UnsetParentSet ensures that no value is present for ParentSet, not even an explicit nil
### GetCriticalHitProbability

`func (o *ListItem) GetCriticalHitProbability() int32`

GetCriticalHitProbability returns the CriticalHitProbability field if non-nil, zero value otherwise.

### GetCriticalHitProbabilityOk

`func (o *ListItem) GetCriticalHitProbabilityOk() (*int32, bool)`

GetCriticalHitProbabilityOk returns a tuple with the CriticalHitProbability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalHitProbability

`func (o *ListItem) SetCriticalHitProbability(v int32)`

SetCriticalHitProbability sets CriticalHitProbability field to given value.

### HasCriticalHitProbability

`func (o *ListItem) HasCriticalHitProbability() bool`

HasCriticalHitProbability returns a boolean if a field has been set.

### SetCriticalHitProbabilityNil

`func (o *ListItem) SetCriticalHitProbabilityNil(b bool)`

 SetCriticalHitProbabilityNil sets the value for CriticalHitProbability to be an explicit nil

### UnsetCriticalHitProbability
`func (o *ListItem) UnsetCriticalHitProbability()`

UnsetCriticalHitProbability ensures that no value is present for CriticalHitProbability, not even an explicit nil
### GetCriticalHitBonus

`func (o *ListItem) GetCriticalHitBonus() int32`

GetCriticalHitBonus returns the CriticalHitBonus field if non-nil, zero value otherwise.

### GetCriticalHitBonusOk

`func (o *ListItem) GetCriticalHitBonusOk() (*int32, bool)`

GetCriticalHitBonusOk returns a tuple with the CriticalHitBonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalHitBonus

`func (o *ListItem) SetCriticalHitBonus(v int32)`

SetCriticalHitBonus sets CriticalHitBonus field to given value.

### HasCriticalHitBonus

`func (o *ListItem) HasCriticalHitBonus() bool`

HasCriticalHitBonus returns a boolean if a field has been set.

### SetCriticalHitBonusNil

`func (o *ListItem) SetCriticalHitBonusNil(b bool)`

 SetCriticalHitBonusNil sets the value for CriticalHitBonus to be an explicit nil

### UnsetCriticalHitBonus
`func (o *ListItem) UnsetCriticalHitBonus()`

UnsetCriticalHitBonus ensures that no value is present for CriticalHitBonus, not even an explicit nil
### GetMaxCastPerTurn

`func (o *ListItem) GetMaxCastPerTurn() int32`

GetMaxCastPerTurn returns the MaxCastPerTurn field if non-nil, zero value otherwise.

### GetMaxCastPerTurnOk

`func (o *ListItem) GetMaxCastPerTurnOk() (*int32, bool)`

GetMaxCastPerTurnOk returns a tuple with the MaxCastPerTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCastPerTurn

`func (o *ListItem) SetMaxCastPerTurn(v int32)`

SetMaxCastPerTurn sets MaxCastPerTurn field to given value.

### HasMaxCastPerTurn

`func (o *ListItem) HasMaxCastPerTurn() bool`

HasMaxCastPerTurn returns a boolean if a field has been set.

### SetMaxCastPerTurnNil

`func (o *ListItem) SetMaxCastPerTurnNil(b bool)`

 SetMaxCastPerTurnNil sets the value for MaxCastPerTurn to be an explicit nil

### UnsetMaxCastPerTurn
`func (o *ListItem) UnsetMaxCastPerTurn()`

UnsetMaxCastPerTurn ensures that no value is present for MaxCastPerTurn, not even an explicit nil
### GetApCost

`func (o *ListItem) GetApCost() int32`

GetApCost returns the ApCost field if non-nil, zero value otherwise.

### GetApCostOk

`func (o *ListItem) GetApCostOk() (*int32, bool)`

GetApCostOk returns a tuple with the ApCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApCost

`func (o *ListItem) SetApCost(v int32)`

SetApCost sets ApCost field to given value.

### HasApCost

`func (o *ListItem) HasApCost() bool`

HasApCost returns a boolean if a field has been set.

### SetApCostNil

`func (o *ListItem) SetApCostNil(b bool)`

 SetApCostNil sets the value for ApCost to be an explicit nil

### UnsetApCost
`func (o *ListItem) UnsetApCost()`

UnsetApCost ensures that no value is present for ApCost, not even an explicit nil
### GetRange

`func (o *ListItem) GetRange() Range`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *ListItem) GetRangeOk() (*Range, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *ListItem) SetRange(v Range)`

SetRange sets Range field to given value.

### HasRange

`func (o *ListItem) HasRange() bool`

HasRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


