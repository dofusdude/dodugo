# Spell

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | **int32** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**IconId** | Pointer to **int32** |  | [optional] 
**AdminName** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**ImageUrls** | Pointer to [**Images**](Images.md) |  | [optional] 
**Type** | Pointer to [**CombatReference**](CombatReference.md) |  | [optional] 
**Classes** | Pointer to [**[]CombatReference**](CombatReference.md) |  | [optional] 
**Monsters** | Pointer to [**[]CombatReference**](CombatReference.md) |  | [optional] 
**Levels** | [**[]SpellLevel**](SpellLevel.md) |  | 
**Scripts** | Pointer to [**[]SpellScriptUsage**](SpellScriptUsage.md) |  | [optional] 
**CriticalScripts** | Pointer to [**[]SpellScriptUsage**](SpellScriptUsage.md) |  | [optional] 

## Methods

### NewSpell

`func NewSpell(ankamaId int32, name string, levels []SpellLevel, ) *Spell`

NewSpell instantiates a new Spell object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpellWithDefaults

`func NewSpellWithDefaults() *Spell`

NewSpellWithDefaults instantiates a new Spell object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *Spell) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *Spell) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *Spell) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.


### GetName

`func (o *Spell) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Spell) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Spell) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Spell) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Spell) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Spell) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Spell) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIconId

`func (o *Spell) GetIconId() int32`

GetIconId returns the IconId field if non-nil, zero value otherwise.

### GetIconIdOk

`func (o *Spell) GetIconIdOk() (*int32, bool)`

GetIconIdOk returns a tuple with the IconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconId

`func (o *Spell) SetIconId(v int32)`

SetIconId sets IconId field to given value.

### HasIconId

`func (o *Spell) HasIconId() bool`

HasIconId returns a boolean if a field has been set.

### GetAdminName

`func (o *Spell) GetAdminName() string`

GetAdminName returns the AdminName field if non-nil, zero value otherwise.

### GetAdminNameOk

`func (o *Spell) GetAdminNameOk() (*string, bool)`

GetAdminNameOk returns a tuple with the AdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminName

`func (o *Spell) SetAdminName(v string)`

SetAdminName sets AdminName field to given value.

### HasAdminName

`func (o *Spell) HasAdminName() bool`

HasAdminName returns a boolean if a field has been set.

### GetOrder

`func (o *Spell) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Spell) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Spell) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Spell) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetImageUrls

`func (o *Spell) GetImageUrls() Images`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *Spell) GetImageUrlsOk() (*Images, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *Spell) SetImageUrls(v Images)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *Spell) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.

### GetType

`func (o *Spell) GetType() CombatReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Spell) GetTypeOk() (*CombatReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Spell) SetType(v CombatReference)`

SetType sets Type field to given value.

### HasType

`func (o *Spell) HasType() bool`

HasType returns a boolean if a field has been set.

### GetClasses

`func (o *Spell) GetClasses() []CombatReference`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *Spell) GetClassesOk() (*[]CombatReference, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *Spell) SetClasses(v []CombatReference)`

SetClasses sets Classes field to given value.

### HasClasses

`func (o *Spell) HasClasses() bool`

HasClasses returns a boolean if a field has been set.

### GetMonsters

`func (o *Spell) GetMonsters() []CombatReference`

GetMonsters returns the Monsters field if non-nil, zero value otherwise.

### GetMonstersOk

`func (o *Spell) GetMonstersOk() (*[]CombatReference, bool)`

GetMonstersOk returns a tuple with the Monsters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonsters

`func (o *Spell) SetMonsters(v []CombatReference)`

SetMonsters sets Monsters field to given value.

### HasMonsters

`func (o *Spell) HasMonsters() bool`

HasMonsters returns a boolean if a field has been set.

### GetLevels

`func (o *Spell) GetLevels() []SpellLevel`

GetLevels returns the Levels field if non-nil, zero value otherwise.

### GetLevelsOk

`func (o *Spell) GetLevelsOk() (*[]SpellLevel, bool)`

GetLevelsOk returns a tuple with the Levels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevels

`func (o *Spell) SetLevels(v []SpellLevel)`

SetLevels sets Levels field to given value.


### GetScripts

`func (o *Spell) GetScripts() []SpellScriptUsage`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *Spell) GetScriptsOk() (*[]SpellScriptUsage, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *Spell) SetScripts(v []SpellScriptUsage)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *Spell) HasScripts() bool`

HasScripts returns a boolean if a field has been set.

### GetCriticalScripts

`func (o *Spell) GetCriticalScripts() []SpellScriptUsage`

GetCriticalScripts returns the CriticalScripts field if non-nil, zero value otherwise.

### GetCriticalScriptsOk

`func (o *Spell) GetCriticalScriptsOk() (*[]SpellScriptUsage, bool)`

GetCriticalScriptsOk returns a tuple with the CriticalScripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalScripts

`func (o *Spell) SetCriticalScripts(v []SpellScriptUsage)`

SetCriticalScripts sets CriticalScripts field to given value.

### HasCriticalScripts

`func (o *Spell) HasCriticalScripts() bool`

HasCriticalScripts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


