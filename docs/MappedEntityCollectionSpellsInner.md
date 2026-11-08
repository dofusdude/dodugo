# MappedEntityCollectionSpellsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | **int32** |  | 
**Name** | **string** |  | 
**ImageUrls** | Pointer to [**Images**](Images.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IconId** | Pointer to **int32** |  | [optional] 
**AdminName** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to [**CombatReference**](CombatReference.md) |  | [optional] 
**Classes** | Pointer to [**[]CombatReference**](CombatReference.md) |  | [optional] 
**Monsters** | Pointer to [**[]CombatReference**](CombatReference.md) |  | [optional] 
**Levels** | [**[]SpellLevel**](SpellLevel.md) |  | 
**Scripts** | Pointer to [**[]SpellScriptUsage**](SpellScriptUsage.md) |  | [optional] 
**CriticalScripts** | Pointer to [**[]SpellScriptUsage**](SpellScriptUsage.md) |  | [optional] 

## Methods

### NewMappedEntityCollectionSpellsInner

`func NewMappedEntityCollectionSpellsInner(ankamaId int32, name string, levels []SpellLevel, ) *MappedEntityCollectionSpellsInner`

NewMappedEntityCollectionSpellsInner instantiates a new MappedEntityCollectionSpellsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMappedEntityCollectionSpellsInnerWithDefaults

`func NewMappedEntityCollectionSpellsInnerWithDefaults() *MappedEntityCollectionSpellsInner`

NewMappedEntityCollectionSpellsInnerWithDefaults instantiates a new MappedEntityCollectionSpellsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *MappedEntityCollectionSpellsInner) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *MappedEntityCollectionSpellsInner) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *MappedEntityCollectionSpellsInner) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.


### GetName

`func (o *MappedEntityCollectionSpellsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MappedEntityCollectionSpellsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MappedEntityCollectionSpellsInner) SetName(v string)`

SetName sets Name field to given value.


### GetImageUrls

`func (o *MappedEntityCollectionSpellsInner) GetImageUrls() Images`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *MappedEntityCollectionSpellsInner) GetImageUrlsOk() (*Images, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *MappedEntityCollectionSpellsInner) SetImageUrls(v Images)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *MappedEntityCollectionSpellsInner) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.

### GetDescription

`func (o *MappedEntityCollectionSpellsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MappedEntityCollectionSpellsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MappedEntityCollectionSpellsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MappedEntityCollectionSpellsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIconId

`func (o *MappedEntityCollectionSpellsInner) GetIconId() int32`

GetIconId returns the IconId field if non-nil, zero value otherwise.

### GetIconIdOk

`func (o *MappedEntityCollectionSpellsInner) GetIconIdOk() (*int32, bool)`

GetIconIdOk returns a tuple with the IconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconId

`func (o *MappedEntityCollectionSpellsInner) SetIconId(v int32)`

SetIconId sets IconId field to given value.

### HasIconId

`func (o *MappedEntityCollectionSpellsInner) HasIconId() bool`

HasIconId returns a boolean if a field has been set.

### GetAdminName

`func (o *MappedEntityCollectionSpellsInner) GetAdminName() string`

GetAdminName returns the AdminName field if non-nil, zero value otherwise.

### GetAdminNameOk

`func (o *MappedEntityCollectionSpellsInner) GetAdminNameOk() (*string, bool)`

GetAdminNameOk returns a tuple with the AdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminName

`func (o *MappedEntityCollectionSpellsInner) SetAdminName(v string)`

SetAdminName sets AdminName field to given value.

### HasAdminName

`func (o *MappedEntityCollectionSpellsInner) HasAdminName() bool`

HasAdminName returns a boolean if a field has been set.

### GetOrder

`func (o *MappedEntityCollectionSpellsInner) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MappedEntityCollectionSpellsInner) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *MappedEntityCollectionSpellsInner) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *MappedEntityCollectionSpellsInner) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetType

`func (o *MappedEntityCollectionSpellsInner) GetType() CombatReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MappedEntityCollectionSpellsInner) GetTypeOk() (*CombatReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MappedEntityCollectionSpellsInner) SetType(v CombatReference)`

SetType sets Type field to given value.

### HasType

`func (o *MappedEntityCollectionSpellsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetClasses

`func (o *MappedEntityCollectionSpellsInner) GetClasses() []CombatReference`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *MappedEntityCollectionSpellsInner) GetClassesOk() (*[]CombatReference, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *MappedEntityCollectionSpellsInner) SetClasses(v []CombatReference)`

SetClasses sets Classes field to given value.

### HasClasses

`func (o *MappedEntityCollectionSpellsInner) HasClasses() bool`

HasClasses returns a boolean if a field has been set.

### GetMonsters

`func (o *MappedEntityCollectionSpellsInner) GetMonsters() []CombatReference`

GetMonsters returns the Monsters field if non-nil, zero value otherwise.

### GetMonstersOk

`func (o *MappedEntityCollectionSpellsInner) GetMonstersOk() (*[]CombatReference, bool)`

GetMonstersOk returns a tuple with the Monsters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonsters

`func (o *MappedEntityCollectionSpellsInner) SetMonsters(v []CombatReference)`

SetMonsters sets Monsters field to given value.

### HasMonsters

`func (o *MappedEntityCollectionSpellsInner) HasMonsters() bool`

HasMonsters returns a boolean if a field has been set.

### GetLevels

`func (o *MappedEntityCollectionSpellsInner) GetLevels() []SpellLevel`

GetLevels returns the Levels field if non-nil, zero value otherwise.

### GetLevelsOk

`func (o *MappedEntityCollectionSpellsInner) GetLevelsOk() (*[]SpellLevel, bool)`

GetLevelsOk returns a tuple with the Levels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevels

`func (o *MappedEntityCollectionSpellsInner) SetLevels(v []SpellLevel)`

SetLevels sets Levels field to given value.


### GetScripts

`func (o *MappedEntityCollectionSpellsInner) GetScripts() []SpellScriptUsage`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *MappedEntityCollectionSpellsInner) GetScriptsOk() (*[]SpellScriptUsage, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *MappedEntityCollectionSpellsInner) SetScripts(v []SpellScriptUsage)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *MappedEntityCollectionSpellsInner) HasScripts() bool`

HasScripts returns a boolean if a field has been set.

### GetCriticalScripts

`func (o *MappedEntityCollectionSpellsInner) GetCriticalScripts() []SpellScriptUsage`

GetCriticalScripts returns the CriticalScripts field if non-nil, zero value otherwise.

### GetCriticalScriptsOk

`func (o *MappedEntityCollectionSpellsInner) GetCriticalScriptsOk() (*[]SpellScriptUsage, bool)`

GetCriticalScriptsOk returns a tuple with the CriticalScripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalScripts

`func (o *MappedEntityCollectionSpellsInner) SetCriticalScripts(v []SpellScriptUsage)`

SetCriticalScripts sets CriticalScripts field to given value.

### HasCriticalScripts

`func (o *MappedEntityCollectionSpellsInner) HasCriticalScripts() bool`

HasCriticalScripts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


