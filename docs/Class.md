# Class

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | **int32** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**GameplayDescription** | Pointer to **string** |  | [optional] 
**Complexity** | Pointer to **int32** |  | [optional] 
**SortIndex** | Pointer to **int32** |  | [optional] 
**CreatureBonesId** | Pointer to **int32** |  | [optional] 
**Roles** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CharacteristicCosts** | Pointer to [**map[string][]ClassCharacteristicCostsValueInner**](array.md) | Characteristic point-cost tiers keyed by characteristic name. | [optional] 
**ImageUrls** | Pointer to [**Images**](Images.md) |  | [optional] 
**Looks** | [**ClassLooks**](ClassLooks.md) |  | 
**Spells** | [**[]Spell**](Spell.md) | Full spell definitions assigned to this playable class, including their own grades, effects and targeting. A class has no monster-grade dimension. The same spells are independently available from the spells endpoint. | 

## Methods

### NewClass

`func NewClass(ankamaId int32, name string, looks ClassLooks, spells []Spell, ) *Class`

NewClass instantiates a new Class object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassWithDefaults

`func NewClassWithDefaults() *Class`

NewClassWithDefaults instantiates a new Class object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *Class) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *Class) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *Class) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.


### GetName

`func (o *Class) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Class) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Class) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Class) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Class) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Class) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Class) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGameplayDescription

`func (o *Class) GetGameplayDescription() string`

GetGameplayDescription returns the GameplayDescription field if non-nil, zero value otherwise.

### GetGameplayDescriptionOk

`func (o *Class) GetGameplayDescriptionOk() (*string, bool)`

GetGameplayDescriptionOk returns a tuple with the GameplayDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameplayDescription

`func (o *Class) SetGameplayDescription(v string)`

SetGameplayDescription sets GameplayDescription field to given value.

### HasGameplayDescription

`func (o *Class) HasGameplayDescription() bool`

HasGameplayDescription returns a boolean if a field has been set.

### GetComplexity

`func (o *Class) GetComplexity() int32`

GetComplexity returns the Complexity field if non-nil, zero value otherwise.

### GetComplexityOk

`func (o *Class) GetComplexityOk() (*int32, bool)`

GetComplexityOk returns a tuple with the Complexity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexity

`func (o *Class) SetComplexity(v int32)`

SetComplexity sets Complexity field to given value.

### HasComplexity

`func (o *Class) HasComplexity() bool`

HasComplexity returns a boolean if a field has been set.

### GetSortIndex

`func (o *Class) GetSortIndex() int32`

GetSortIndex returns the SortIndex field if non-nil, zero value otherwise.

### GetSortIndexOk

`func (o *Class) GetSortIndexOk() (*int32, bool)`

GetSortIndexOk returns a tuple with the SortIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortIndex

`func (o *Class) SetSortIndex(v int32)`

SetSortIndex sets SortIndex field to given value.

### HasSortIndex

`func (o *Class) HasSortIndex() bool`

HasSortIndex returns a boolean if a field has been set.

### GetCreatureBonesId

`func (o *Class) GetCreatureBonesId() int32`

GetCreatureBonesId returns the CreatureBonesId field if non-nil, zero value otherwise.

### GetCreatureBonesIdOk

`func (o *Class) GetCreatureBonesIdOk() (*int32, bool)`

GetCreatureBonesIdOk returns a tuple with the CreatureBonesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatureBonesId

`func (o *Class) SetCreatureBonesId(v int32)`

SetCreatureBonesId sets CreatureBonesId field to given value.

### HasCreatureBonesId

`func (o *Class) HasCreatureBonesId() bool`

HasCreatureBonesId returns a boolean if a field has been set.

### GetRoles

`func (o *Class) GetRoles() []map[string]interface{}`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Class) GetRolesOk() (*[]map[string]interface{}, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Class) SetRoles(v []map[string]interface{})`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Class) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetCharacteristicCosts

`func (o *Class) GetCharacteristicCosts() map[string][]ClassCharacteristicCostsValueInner`

GetCharacteristicCosts returns the CharacteristicCosts field if non-nil, zero value otherwise.

### GetCharacteristicCostsOk

`func (o *Class) GetCharacteristicCostsOk() (*map[string][]ClassCharacteristicCostsValueInner, bool)`

GetCharacteristicCostsOk returns a tuple with the CharacteristicCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacteristicCosts

`func (o *Class) SetCharacteristicCosts(v map[string][]ClassCharacteristicCostsValueInner)`

SetCharacteristicCosts sets CharacteristicCosts field to given value.

### HasCharacteristicCosts

`func (o *Class) HasCharacteristicCosts() bool`

HasCharacteristicCosts returns a boolean if a field has been set.

### GetImageUrls

`func (o *Class) GetImageUrls() Images`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *Class) GetImageUrlsOk() (*Images, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *Class) SetImageUrls(v Images)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *Class) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.

### GetLooks

`func (o *Class) GetLooks() ClassLooks`

GetLooks returns the Looks field if non-nil, zero value otherwise.

### GetLooksOk

`func (o *Class) GetLooksOk() (*ClassLooks, bool)`

GetLooksOk returns a tuple with the Looks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLooks

`func (o *Class) SetLooks(v ClassLooks)`

SetLooks sets Looks field to given value.


### GetSpells

`func (o *Class) GetSpells() []Spell`

GetSpells returns the Spells field if non-nil, zero value otherwise.

### GetSpellsOk

`func (o *Class) GetSpellsOk() (*[]Spell, bool)`

GetSpellsOk returns a tuple with the Spells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpells

`func (o *Class) SetSpells(v []Spell)`

SetSpells sets Spells field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


