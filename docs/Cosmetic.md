# Cosmetic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**CosmeticType**](CosmeticType.md) |  | [optional] 
**Level** | Pointer to **int32** |  | [optional] 
**Pods** | Pointer to **int32** |  | [optional] 
**ImageUrls** | Pointer to [**ImageUrls**](ImageUrls.md) |  | [optional] 
**Effects** | Pointer to [**[]EffectsEntry**](EffectsEntry.md) |  | [optional] 
**Conditions** | Pointer to [**[]ConditionEntry**](ConditionEntry.md) |  | [optional] 
**ConditionTree** | Pointer to [**ConditionTreeNode**](ConditionTreeNode.md) |  | [optional] 
**Recipe** | Pointer to [**[]RecipeEntry**](RecipeEntry.md) |  | [optional] 

## Methods

### NewCosmetic

`func NewCosmetic() *Cosmetic`

NewCosmetic instantiates a new Cosmetic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCosmeticWithDefaults

`func NewCosmeticWithDefaults() *Cosmetic`

NewCosmeticWithDefaults instantiates a new Cosmetic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *Cosmetic) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *Cosmetic) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *Cosmetic) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.

### HasAnkamaId

`func (o *Cosmetic) HasAnkamaId() bool`

HasAnkamaId returns a boolean if a field has been set.

### GetName

`func (o *Cosmetic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cosmetic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cosmetic) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Cosmetic) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Cosmetic) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Cosmetic) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Cosmetic) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Cosmetic) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *Cosmetic) GetType() CosmeticType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Cosmetic) GetTypeOk() (*CosmeticType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Cosmetic) SetType(v CosmeticType)`

SetType sets Type field to given value.

### HasType

`func (o *Cosmetic) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLevel

`func (o *Cosmetic) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *Cosmetic) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *Cosmetic) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *Cosmetic) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetPods

`func (o *Cosmetic) GetPods() int32`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *Cosmetic) GetPodsOk() (*int32, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *Cosmetic) SetPods(v int32)`

SetPods sets Pods field to given value.

### HasPods

`func (o *Cosmetic) HasPods() bool`

HasPods returns a boolean if a field has been set.

### GetImageUrls

`func (o *Cosmetic) GetImageUrls() ImageUrls`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *Cosmetic) GetImageUrlsOk() (*ImageUrls, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *Cosmetic) SetImageUrls(v ImageUrls)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *Cosmetic) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.

### GetEffects

`func (o *Cosmetic) GetEffects() []EffectsEntry`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *Cosmetic) GetEffectsOk() (*[]EffectsEntry, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *Cosmetic) SetEffects(v []EffectsEntry)`

SetEffects sets Effects field to given value.

### HasEffects

`func (o *Cosmetic) HasEffects() bool`

HasEffects returns a boolean if a field has been set.

### SetEffectsNil

`func (o *Cosmetic) SetEffectsNil(b bool)`

 SetEffectsNil sets the value for Effects to be an explicit nil

### UnsetEffects
`func (o *Cosmetic) UnsetEffects()`

UnsetEffects ensures that no value is present for Effects, not even an explicit nil
### GetConditions

`func (o *Cosmetic) GetConditions() []ConditionEntry`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *Cosmetic) GetConditionsOk() (*[]ConditionEntry, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *Cosmetic) SetConditions(v []ConditionEntry)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *Cosmetic) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *Cosmetic) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *Cosmetic) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetConditionTree

`func (o *Cosmetic) GetConditionTree() ConditionTreeNode`

GetConditionTree returns the ConditionTree field if non-nil, zero value otherwise.

### GetConditionTreeOk

`func (o *Cosmetic) GetConditionTreeOk() (*ConditionTreeNode, bool)`

GetConditionTreeOk returns a tuple with the ConditionTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionTree

`func (o *Cosmetic) SetConditionTree(v ConditionTreeNode)`

SetConditionTree sets ConditionTree field to given value.

### HasConditionTree

`func (o *Cosmetic) HasConditionTree() bool`

HasConditionTree returns a boolean if a field has been set.

### GetRecipe

`func (o *Cosmetic) GetRecipe() []RecipeEntry`

GetRecipe returns the Recipe field if non-nil, zero value otherwise.

### GetRecipeOk

`func (o *Cosmetic) GetRecipeOk() (*[]RecipeEntry, bool)`

GetRecipeOk returns a tuple with the Recipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipe

`func (o *Cosmetic) SetRecipe(v []RecipeEntry)`

SetRecipe sets Recipe field to given value.

### HasRecipe

`func (o *Cosmetic) HasRecipe() bool`

HasRecipe returns a boolean if a field has been set.

### SetRecipeNil

`func (o *Cosmetic) SetRecipeNil(b bool)`

 SetRecipeNil sets the value for Recipe to be an explicit nil

### UnsetRecipe
`func (o *Cosmetic) UnsetRecipe()`

UnsetRecipe ensures that no value is present for Recipe, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


