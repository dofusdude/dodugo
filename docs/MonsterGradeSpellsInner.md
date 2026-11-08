# MonsterGradeSpellsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | **int32** |  | 
**Name** | **string** |  | 
**Grade** | **int32** | The spell grade assigned to this monster grade. | 
**Url** | **string** | Grade-specific spell URL; already includes the &#x60;grade&#x60; query parameter. | 
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

### NewMonsterGradeSpellsInner

`func NewMonsterGradeSpellsInner(ankamaId int32, name string, grade int32, url string, levels []SpellLevel, ) *MonsterGradeSpellsInner`

NewMonsterGradeSpellsInner instantiates a new MonsterGradeSpellsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonsterGradeSpellsInnerWithDefaults

`func NewMonsterGradeSpellsInnerWithDefaults() *MonsterGradeSpellsInner`

NewMonsterGradeSpellsInnerWithDefaults instantiates a new MonsterGradeSpellsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *MonsterGradeSpellsInner) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *MonsterGradeSpellsInner) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *MonsterGradeSpellsInner) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.


### GetName

`func (o *MonsterGradeSpellsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonsterGradeSpellsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonsterGradeSpellsInner) SetName(v string)`

SetName sets Name field to given value.


### GetGrade

`func (o *MonsterGradeSpellsInner) GetGrade() int32`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *MonsterGradeSpellsInner) GetGradeOk() (*int32, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *MonsterGradeSpellsInner) SetGrade(v int32)`

SetGrade sets Grade field to given value.


### GetUrl

`func (o *MonsterGradeSpellsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MonsterGradeSpellsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MonsterGradeSpellsInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetImageUrls

`func (o *MonsterGradeSpellsInner) GetImageUrls() Images`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *MonsterGradeSpellsInner) GetImageUrlsOk() (*Images, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *MonsterGradeSpellsInner) SetImageUrls(v Images)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *MonsterGradeSpellsInner) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.

### GetDescription

`func (o *MonsterGradeSpellsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MonsterGradeSpellsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MonsterGradeSpellsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MonsterGradeSpellsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIconId

`func (o *MonsterGradeSpellsInner) GetIconId() int32`

GetIconId returns the IconId field if non-nil, zero value otherwise.

### GetIconIdOk

`func (o *MonsterGradeSpellsInner) GetIconIdOk() (*int32, bool)`

GetIconIdOk returns a tuple with the IconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconId

`func (o *MonsterGradeSpellsInner) SetIconId(v int32)`

SetIconId sets IconId field to given value.

### HasIconId

`func (o *MonsterGradeSpellsInner) HasIconId() bool`

HasIconId returns a boolean if a field has been set.

### GetAdminName

`func (o *MonsterGradeSpellsInner) GetAdminName() string`

GetAdminName returns the AdminName field if non-nil, zero value otherwise.

### GetAdminNameOk

`func (o *MonsterGradeSpellsInner) GetAdminNameOk() (*string, bool)`

GetAdminNameOk returns a tuple with the AdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminName

`func (o *MonsterGradeSpellsInner) SetAdminName(v string)`

SetAdminName sets AdminName field to given value.

### HasAdminName

`func (o *MonsterGradeSpellsInner) HasAdminName() bool`

HasAdminName returns a boolean if a field has been set.

### GetOrder

`func (o *MonsterGradeSpellsInner) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MonsterGradeSpellsInner) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *MonsterGradeSpellsInner) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *MonsterGradeSpellsInner) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetType

`func (o *MonsterGradeSpellsInner) GetType() CombatReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonsterGradeSpellsInner) GetTypeOk() (*CombatReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonsterGradeSpellsInner) SetType(v CombatReference)`

SetType sets Type field to given value.

### HasType

`func (o *MonsterGradeSpellsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetClasses

`func (o *MonsterGradeSpellsInner) GetClasses() []CombatReference`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *MonsterGradeSpellsInner) GetClassesOk() (*[]CombatReference, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *MonsterGradeSpellsInner) SetClasses(v []CombatReference)`

SetClasses sets Classes field to given value.

### HasClasses

`func (o *MonsterGradeSpellsInner) HasClasses() bool`

HasClasses returns a boolean if a field has been set.

### GetMonsters

`func (o *MonsterGradeSpellsInner) GetMonsters() []CombatReference`

GetMonsters returns the Monsters field if non-nil, zero value otherwise.

### GetMonstersOk

`func (o *MonsterGradeSpellsInner) GetMonstersOk() (*[]CombatReference, bool)`

GetMonstersOk returns a tuple with the Monsters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonsters

`func (o *MonsterGradeSpellsInner) SetMonsters(v []CombatReference)`

SetMonsters sets Monsters field to given value.

### HasMonsters

`func (o *MonsterGradeSpellsInner) HasMonsters() bool`

HasMonsters returns a boolean if a field has been set.

### GetLevels

`func (o *MonsterGradeSpellsInner) GetLevels() []SpellLevel`

GetLevels returns the Levels field if non-nil, zero value otherwise.

### GetLevelsOk

`func (o *MonsterGradeSpellsInner) GetLevelsOk() (*[]SpellLevel, bool)`

GetLevelsOk returns a tuple with the Levels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevels

`func (o *MonsterGradeSpellsInner) SetLevels(v []SpellLevel)`

SetLevels sets Levels field to given value.


### GetScripts

`func (o *MonsterGradeSpellsInner) GetScripts() []SpellScriptUsage`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *MonsterGradeSpellsInner) GetScriptsOk() (*[]SpellScriptUsage, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *MonsterGradeSpellsInner) SetScripts(v []SpellScriptUsage)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *MonsterGradeSpellsInner) HasScripts() bool`

HasScripts returns a boolean if a field has been set.

### GetCriticalScripts

`func (o *MonsterGradeSpellsInner) GetCriticalScripts() []SpellScriptUsage`

GetCriticalScripts returns the CriticalScripts field if non-nil, zero value otherwise.

### GetCriticalScriptsOk

`func (o *MonsterGradeSpellsInner) GetCriticalScriptsOk() (*[]SpellScriptUsage, bool)`

GetCriticalScriptsOk returns a tuple with the CriticalScripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalScripts

`func (o *MonsterGradeSpellsInner) SetCriticalScripts(v []SpellScriptUsage)`

SetCriticalScripts sets CriticalScripts field to given value.

### HasCriticalScripts

`func (o *MonsterGradeSpellsInner) HasCriticalScripts() bool`

HasCriticalScripts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


