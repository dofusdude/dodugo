# Equipment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**TranslatedId**](TranslatedId.md) |  | [optional] 
**IsWeapon** | Pointer to **bool** |  | [optional] 
**Level** | Pointer to **int32** |  | [optional] 
**Pods** | Pointer to **int32** |  | [optional] 
**ImageUrls** | Pointer to [**Images**](Images.md) |  | [optional] 
**Effects** | Pointer to [**[]Effect**](Effect.md) |  | [optional] 
**Conditions** | Pointer to [**NullableConditionNode**](ConditionNode.md) |  | [optional] 
**Recipe** | Pointer to [**[]Recipe**](Recipe.md) |  | [optional] 
**ParentSet** | Pointer to [**NullableTranslatedId**](TranslatedId.md) |  | [optional] 

## Methods

### NewEquipment

`func NewEquipment() *Equipment`

NewEquipment instantiates a new Equipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentWithDefaults

`func NewEquipmentWithDefaults() *Equipment`

NewEquipmentWithDefaults instantiates a new Equipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *Equipment) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *Equipment) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *Equipment) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.

### HasAnkamaId

`func (o *Equipment) HasAnkamaId() bool`

HasAnkamaId returns a boolean if a field has been set.

### GetName

`func (o *Equipment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Equipment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Equipment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Equipment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Equipment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Equipment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Equipment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Equipment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *Equipment) GetType() TranslatedId`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Equipment) GetTypeOk() (*TranslatedId, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Equipment) SetType(v TranslatedId)`

SetType sets Type field to given value.

### HasType

`func (o *Equipment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsWeapon

`func (o *Equipment) GetIsWeapon() bool`

GetIsWeapon returns the IsWeapon field if non-nil, zero value otherwise.

### GetIsWeaponOk

`func (o *Equipment) GetIsWeaponOk() (*bool, bool)`

GetIsWeaponOk returns a tuple with the IsWeapon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWeapon

`func (o *Equipment) SetIsWeapon(v bool)`

SetIsWeapon sets IsWeapon field to given value.

### HasIsWeapon

`func (o *Equipment) HasIsWeapon() bool`

HasIsWeapon returns a boolean if a field has been set.

### GetLevel

`func (o *Equipment) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *Equipment) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *Equipment) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *Equipment) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetPods

`func (o *Equipment) GetPods() int32`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *Equipment) GetPodsOk() (*int32, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *Equipment) SetPods(v int32)`

SetPods sets Pods field to given value.

### HasPods

`func (o *Equipment) HasPods() bool`

HasPods returns a boolean if a field has been set.

### GetImageUrls

`func (o *Equipment) GetImageUrls() Images`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *Equipment) GetImageUrlsOk() (*Images, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *Equipment) SetImageUrls(v Images)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *Equipment) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.

### GetEffects

`func (o *Equipment) GetEffects() []Effect`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *Equipment) GetEffectsOk() (*[]Effect, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *Equipment) SetEffects(v []Effect)`

SetEffects sets Effects field to given value.

### HasEffects

`func (o *Equipment) HasEffects() bool`

HasEffects returns a boolean if a field has been set.

### SetEffectsNil

`func (o *Equipment) SetEffectsNil(b bool)`

 SetEffectsNil sets the value for Effects to be an explicit nil

### UnsetEffects
`func (o *Equipment) UnsetEffects()`

UnsetEffects ensures that no value is present for Effects, not even an explicit nil
### GetConditions

`func (o *Equipment) GetConditions() ConditionNode`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *Equipment) GetConditionsOk() (*ConditionNode, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *Equipment) SetConditions(v ConditionNode)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *Equipment) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *Equipment) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *Equipment) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetRecipe

`func (o *Equipment) GetRecipe() []Recipe`

GetRecipe returns the Recipe field if non-nil, zero value otherwise.

### GetRecipeOk

`func (o *Equipment) GetRecipeOk() (*[]Recipe, bool)`

GetRecipeOk returns a tuple with the Recipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipe

`func (o *Equipment) SetRecipe(v []Recipe)`

SetRecipe sets Recipe field to given value.

### HasRecipe

`func (o *Equipment) HasRecipe() bool`

HasRecipe returns a boolean if a field has been set.

### SetRecipeNil

`func (o *Equipment) SetRecipeNil(b bool)`

 SetRecipeNil sets the value for Recipe to be an explicit nil

### UnsetRecipe
`func (o *Equipment) UnsetRecipe()`

UnsetRecipe ensures that no value is present for Recipe, not even an explicit nil
### GetParentSet

`func (o *Equipment) GetParentSet() TranslatedId`

GetParentSet returns the ParentSet field if non-nil, zero value otherwise.

### GetParentSetOk

`func (o *Equipment) GetParentSetOk() (*TranslatedId, bool)`

GetParentSetOk returns a tuple with the ParentSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSet

`func (o *Equipment) SetParentSet(v TranslatedId)`

SetParentSet sets ParentSet field to given value.

### HasParentSet

`func (o *Equipment) HasParentSet() bool`

HasParentSet returns a boolean if a field has been set.

### SetParentSetNil

`func (o *Equipment) SetParentSetNil(b bool)`

 SetParentSetNil sets the value for ParentSet to be an explicit nil

### UnsetParentSet
`func (o *Equipment) UnsetParentSet()`

UnsetParentSet ensures that no value is present for ParentSet, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


