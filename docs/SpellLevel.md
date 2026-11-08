# SpellLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | Pointer to **int32** |  | [optional] 
**Grade** | Pointer to **int32** |  | [optional] 
**ApCost** | Pointer to **int32** |  | [optional] 
**MinRange** | Pointer to **int32** |  | [optional] 
**MaxRange** | Pointer to **int32** |  | [optional] 
**MaxCastPerTurn** | Pointer to **int32** |  | [optional] 
**MaxCastPerTarget** | Pointer to **int32** |  | [optional] 
**MaxGlobalCastPerTurn** | Pointer to **int32** |  | [optional] 
**MaxGlobalCastPerTarget** | Pointer to **int32** |  | [optional] 
**MinCastInterval** | Pointer to **int32** |  | [optional] 
**InitialCooldown** | Pointer to **int32** |  | [optional] 
**GlobalCooldown** | Pointer to **int32** |  | [optional] 
**IgnorePrisonGlyph** | Pointer to **bool** |  | [optional] 
**SpellBreed** | Pointer to **int32** |  | [optional] 
**Effects** | Pointer to [**[]SpellEffect**](SpellEffect.md) |  | [optional] 
**CriticalEffects** | Pointer to [**[]SpellEffect**](SpellEffect.md) |  | [optional] 

## Methods

### NewSpellLevel

`func NewSpellLevel() *SpellLevel`

NewSpellLevel instantiates a new SpellLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpellLevelWithDefaults

`func NewSpellLevelWithDefaults() *SpellLevel`

NewSpellLevelWithDefaults instantiates a new SpellLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *SpellLevel) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *SpellLevel) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *SpellLevel) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.

### HasAnkamaId

`func (o *SpellLevel) HasAnkamaId() bool`

HasAnkamaId returns a boolean if a field has been set.

### GetGrade

`func (o *SpellLevel) GetGrade() int32`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *SpellLevel) GetGradeOk() (*int32, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *SpellLevel) SetGrade(v int32)`

SetGrade sets Grade field to given value.

### HasGrade

`func (o *SpellLevel) HasGrade() bool`

HasGrade returns a boolean if a field has been set.

### GetApCost

`func (o *SpellLevel) GetApCost() int32`

GetApCost returns the ApCost field if non-nil, zero value otherwise.

### GetApCostOk

`func (o *SpellLevel) GetApCostOk() (*int32, bool)`

GetApCostOk returns a tuple with the ApCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApCost

`func (o *SpellLevel) SetApCost(v int32)`

SetApCost sets ApCost field to given value.

### HasApCost

`func (o *SpellLevel) HasApCost() bool`

HasApCost returns a boolean if a field has been set.

### GetMinRange

`func (o *SpellLevel) GetMinRange() int32`

GetMinRange returns the MinRange field if non-nil, zero value otherwise.

### GetMinRangeOk

`func (o *SpellLevel) GetMinRangeOk() (*int32, bool)`

GetMinRangeOk returns a tuple with the MinRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRange

`func (o *SpellLevel) SetMinRange(v int32)`

SetMinRange sets MinRange field to given value.

### HasMinRange

`func (o *SpellLevel) HasMinRange() bool`

HasMinRange returns a boolean if a field has been set.

### GetMaxRange

`func (o *SpellLevel) GetMaxRange() int32`

GetMaxRange returns the MaxRange field if non-nil, zero value otherwise.

### GetMaxRangeOk

`func (o *SpellLevel) GetMaxRangeOk() (*int32, bool)`

GetMaxRangeOk returns a tuple with the MaxRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRange

`func (o *SpellLevel) SetMaxRange(v int32)`

SetMaxRange sets MaxRange field to given value.

### HasMaxRange

`func (o *SpellLevel) HasMaxRange() bool`

HasMaxRange returns a boolean if a field has been set.

### GetMaxCastPerTurn

`func (o *SpellLevel) GetMaxCastPerTurn() int32`

GetMaxCastPerTurn returns the MaxCastPerTurn field if non-nil, zero value otherwise.

### GetMaxCastPerTurnOk

`func (o *SpellLevel) GetMaxCastPerTurnOk() (*int32, bool)`

GetMaxCastPerTurnOk returns a tuple with the MaxCastPerTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCastPerTurn

`func (o *SpellLevel) SetMaxCastPerTurn(v int32)`

SetMaxCastPerTurn sets MaxCastPerTurn field to given value.

### HasMaxCastPerTurn

`func (o *SpellLevel) HasMaxCastPerTurn() bool`

HasMaxCastPerTurn returns a boolean if a field has been set.

### GetMaxCastPerTarget

`func (o *SpellLevel) GetMaxCastPerTarget() int32`

GetMaxCastPerTarget returns the MaxCastPerTarget field if non-nil, zero value otherwise.

### GetMaxCastPerTargetOk

`func (o *SpellLevel) GetMaxCastPerTargetOk() (*int32, bool)`

GetMaxCastPerTargetOk returns a tuple with the MaxCastPerTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCastPerTarget

`func (o *SpellLevel) SetMaxCastPerTarget(v int32)`

SetMaxCastPerTarget sets MaxCastPerTarget field to given value.

### HasMaxCastPerTarget

`func (o *SpellLevel) HasMaxCastPerTarget() bool`

HasMaxCastPerTarget returns a boolean if a field has been set.

### GetMaxGlobalCastPerTurn

`func (o *SpellLevel) GetMaxGlobalCastPerTurn() int32`

GetMaxGlobalCastPerTurn returns the MaxGlobalCastPerTurn field if non-nil, zero value otherwise.

### GetMaxGlobalCastPerTurnOk

`func (o *SpellLevel) GetMaxGlobalCastPerTurnOk() (*int32, bool)`

GetMaxGlobalCastPerTurnOk returns a tuple with the MaxGlobalCastPerTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGlobalCastPerTurn

`func (o *SpellLevel) SetMaxGlobalCastPerTurn(v int32)`

SetMaxGlobalCastPerTurn sets MaxGlobalCastPerTurn field to given value.

### HasMaxGlobalCastPerTurn

`func (o *SpellLevel) HasMaxGlobalCastPerTurn() bool`

HasMaxGlobalCastPerTurn returns a boolean if a field has been set.

### GetMaxGlobalCastPerTarget

`func (o *SpellLevel) GetMaxGlobalCastPerTarget() int32`

GetMaxGlobalCastPerTarget returns the MaxGlobalCastPerTarget field if non-nil, zero value otherwise.

### GetMaxGlobalCastPerTargetOk

`func (o *SpellLevel) GetMaxGlobalCastPerTargetOk() (*int32, bool)`

GetMaxGlobalCastPerTargetOk returns a tuple with the MaxGlobalCastPerTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGlobalCastPerTarget

`func (o *SpellLevel) SetMaxGlobalCastPerTarget(v int32)`

SetMaxGlobalCastPerTarget sets MaxGlobalCastPerTarget field to given value.

### HasMaxGlobalCastPerTarget

`func (o *SpellLevel) HasMaxGlobalCastPerTarget() bool`

HasMaxGlobalCastPerTarget returns a boolean if a field has been set.

### GetMinCastInterval

`func (o *SpellLevel) GetMinCastInterval() int32`

GetMinCastInterval returns the MinCastInterval field if non-nil, zero value otherwise.

### GetMinCastIntervalOk

`func (o *SpellLevel) GetMinCastIntervalOk() (*int32, bool)`

GetMinCastIntervalOk returns a tuple with the MinCastInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCastInterval

`func (o *SpellLevel) SetMinCastInterval(v int32)`

SetMinCastInterval sets MinCastInterval field to given value.

### HasMinCastInterval

`func (o *SpellLevel) HasMinCastInterval() bool`

HasMinCastInterval returns a boolean if a field has been set.

### GetInitialCooldown

`func (o *SpellLevel) GetInitialCooldown() int32`

GetInitialCooldown returns the InitialCooldown field if non-nil, zero value otherwise.

### GetInitialCooldownOk

`func (o *SpellLevel) GetInitialCooldownOk() (*int32, bool)`

GetInitialCooldownOk returns a tuple with the InitialCooldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialCooldown

`func (o *SpellLevel) SetInitialCooldown(v int32)`

SetInitialCooldown sets InitialCooldown field to given value.

### HasInitialCooldown

`func (o *SpellLevel) HasInitialCooldown() bool`

HasInitialCooldown returns a boolean if a field has been set.

### GetGlobalCooldown

`func (o *SpellLevel) GetGlobalCooldown() int32`

GetGlobalCooldown returns the GlobalCooldown field if non-nil, zero value otherwise.

### GetGlobalCooldownOk

`func (o *SpellLevel) GetGlobalCooldownOk() (*int32, bool)`

GetGlobalCooldownOk returns a tuple with the GlobalCooldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalCooldown

`func (o *SpellLevel) SetGlobalCooldown(v int32)`

SetGlobalCooldown sets GlobalCooldown field to given value.

### HasGlobalCooldown

`func (o *SpellLevel) HasGlobalCooldown() bool`

HasGlobalCooldown returns a boolean if a field has been set.

### GetIgnorePrisonGlyph

`func (o *SpellLevel) GetIgnorePrisonGlyph() bool`

GetIgnorePrisonGlyph returns the IgnorePrisonGlyph field if non-nil, zero value otherwise.

### GetIgnorePrisonGlyphOk

`func (o *SpellLevel) GetIgnorePrisonGlyphOk() (*bool, bool)`

GetIgnorePrisonGlyphOk returns a tuple with the IgnorePrisonGlyph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnorePrisonGlyph

`func (o *SpellLevel) SetIgnorePrisonGlyph(v bool)`

SetIgnorePrisonGlyph sets IgnorePrisonGlyph field to given value.

### HasIgnorePrisonGlyph

`func (o *SpellLevel) HasIgnorePrisonGlyph() bool`

HasIgnorePrisonGlyph returns a boolean if a field has been set.

### GetSpellBreed

`func (o *SpellLevel) GetSpellBreed() int32`

GetSpellBreed returns the SpellBreed field if non-nil, zero value otherwise.

### GetSpellBreedOk

`func (o *SpellLevel) GetSpellBreedOk() (*int32, bool)`

GetSpellBreedOk returns a tuple with the SpellBreed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpellBreed

`func (o *SpellLevel) SetSpellBreed(v int32)`

SetSpellBreed sets SpellBreed field to given value.

### HasSpellBreed

`func (o *SpellLevel) HasSpellBreed() bool`

HasSpellBreed returns a boolean if a field has been set.

### GetEffects

`func (o *SpellLevel) GetEffects() []SpellEffect`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *SpellLevel) GetEffectsOk() (*[]SpellEffect, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *SpellLevel) SetEffects(v []SpellEffect)`

SetEffects sets Effects field to given value.

### HasEffects

`func (o *SpellLevel) HasEffects() bool`

HasEffects returns a boolean if a field has been set.

### GetCriticalEffects

`func (o *SpellLevel) GetCriticalEffects() []SpellEffect`

GetCriticalEffects returns the CriticalEffects field if non-nil, zero value otherwise.

### GetCriticalEffectsOk

`func (o *SpellLevel) GetCriticalEffectsOk() (*[]SpellEffect, bool)`

GetCriticalEffectsOk returns a tuple with the CriticalEffects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalEffects

`func (o *SpellLevel) SetCriticalEffects(v []SpellEffect)`

SetCriticalEffects sets CriticalEffects field to given value.

### HasCriticalEffects

`func (o *SpellLevel) HasCriticalEffects() bool`

HasCriticalEffects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


