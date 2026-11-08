# MonsterGrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grade** | **int32** |  | 
**Level** | **int32** |  | 
**LifePoints** | Pointer to **int32** |  | [optional] 
**ActionPoints** | Pointer to **int32** |  | [optional] 
**MovementPoints** | Pointer to **int32** |  | [optional] 
**GradeXp** | Pointer to **int32** |  | [optional] 
**NeutralResistance** | Pointer to **int32** |  | [optional] 
**EarthResistance** | Pointer to **int32** |  | [optional] 
**FireResistance** | Pointer to **int32** |  | [optional] 
**WaterResistance** | Pointer to **int32** |  | [optional] 
**AirResistance** | Pointer to **int32** |  | [optional] 
**PaDodge** | Pointer to **int32** |  | [optional] 
**PmDodge** | Pointer to **int32** |  | [optional] 
**DamageBonus** | Pointer to **int32** |  | [optional] 
**PercentDamageBonus** | Pointer to **int32** |  | [optional] 
**CriticalHitBonus** | Pointer to **int32** |  | [optional] 
**InitiativeBonus** | Pointer to **int32** |  | [optional] 
**TackleBonus** | Pointer to **int32** |  | [optional] 
**TackleEvade** | Pointer to **int32** |  | [optional] 
**MaxSummon** | Pointer to **int32** |  | [optional] 
**BonusCharacteristics** | Pointer to **map[string]int32** |  | [optional] 
**Spells** | [**[]MonsterGradeSpellsInner**](MonsterGradeSpellsInner.md) | Spells assigned to this exact monster grade. | 

## Methods

### NewMonsterGrade

`func NewMonsterGrade(grade int32, level int32, spells []MonsterGradeSpellsInner, ) *MonsterGrade`

NewMonsterGrade instantiates a new MonsterGrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonsterGradeWithDefaults

`func NewMonsterGradeWithDefaults() *MonsterGrade`

NewMonsterGradeWithDefaults instantiates a new MonsterGrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrade

`func (o *MonsterGrade) GetGrade() int32`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *MonsterGrade) GetGradeOk() (*int32, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *MonsterGrade) SetGrade(v int32)`

SetGrade sets Grade field to given value.


### GetLevel

`func (o *MonsterGrade) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *MonsterGrade) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *MonsterGrade) SetLevel(v int32)`

SetLevel sets Level field to given value.


### GetLifePoints

`func (o *MonsterGrade) GetLifePoints() int32`

GetLifePoints returns the LifePoints field if non-nil, zero value otherwise.

### GetLifePointsOk

`func (o *MonsterGrade) GetLifePointsOk() (*int32, bool)`

GetLifePointsOk returns a tuple with the LifePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifePoints

`func (o *MonsterGrade) SetLifePoints(v int32)`

SetLifePoints sets LifePoints field to given value.

### HasLifePoints

`func (o *MonsterGrade) HasLifePoints() bool`

HasLifePoints returns a boolean if a field has been set.

### GetActionPoints

`func (o *MonsterGrade) GetActionPoints() int32`

GetActionPoints returns the ActionPoints field if non-nil, zero value otherwise.

### GetActionPointsOk

`func (o *MonsterGrade) GetActionPointsOk() (*int32, bool)`

GetActionPointsOk returns a tuple with the ActionPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionPoints

`func (o *MonsterGrade) SetActionPoints(v int32)`

SetActionPoints sets ActionPoints field to given value.

### HasActionPoints

`func (o *MonsterGrade) HasActionPoints() bool`

HasActionPoints returns a boolean if a field has been set.

### GetMovementPoints

`func (o *MonsterGrade) GetMovementPoints() int32`

GetMovementPoints returns the MovementPoints field if non-nil, zero value otherwise.

### GetMovementPointsOk

`func (o *MonsterGrade) GetMovementPointsOk() (*int32, bool)`

GetMovementPointsOk returns a tuple with the MovementPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovementPoints

`func (o *MonsterGrade) SetMovementPoints(v int32)`

SetMovementPoints sets MovementPoints field to given value.

### HasMovementPoints

`func (o *MonsterGrade) HasMovementPoints() bool`

HasMovementPoints returns a boolean if a field has been set.

### GetGradeXp

`func (o *MonsterGrade) GetGradeXp() int32`

GetGradeXp returns the GradeXp field if non-nil, zero value otherwise.

### GetGradeXpOk

`func (o *MonsterGrade) GetGradeXpOk() (*int32, bool)`

GetGradeXpOk returns a tuple with the GradeXp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradeXp

`func (o *MonsterGrade) SetGradeXp(v int32)`

SetGradeXp sets GradeXp field to given value.

### HasGradeXp

`func (o *MonsterGrade) HasGradeXp() bool`

HasGradeXp returns a boolean if a field has been set.

### GetNeutralResistance

`func (o *MonsterGrade) GetNeutralResistance() int32`

GetNeutralResistance returns the NeutralResistance field if non-nil, zero value otherwise.

### GetNeutralResistanceOk

`func (o *MonsterGrade) GetNeutralResistanceOk() (*int32, bool)`

GetNeutralResistanceOk returns a tuple with the NeutralResistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeutralResistance

`func (o *MonsterGrade) SetNeutralResistance(v int32)`

SetNeutralResistance sets NeutralResistance field to given value.

### HasNeutralResistance

`func (o *MonsterGrade) HasNeutralResistance() bool`

HasNeutralResistance returns a boolean if a field has been set.

### GetEarthResistance

`func (o *MonsterGrade) GetEarthResistance() int32`

GetEarthResistance returns the EarthResistance field if non-nil, zero value otherwise.

### GetEarthResistanceOk

`func (o *MonsterGrade) GetEarthResistanceOk() (*int32, bool)`

GetEarthResistanceOk returns a tuple with the EarthResistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarthResistance

`func (o *MonsterGrade) SetEarthResistance(v int32)`

SetEarthResistance sets EarthResistance field to given value.

### HasEarthResistance

`func (o *MonsterGrade) HasEarthResistance() bool`

HasEarthResistance returns a boolean if a field has been set.

### GetFireResistance

`func (o *MonsterGrade) GetFireResistance() int32`

GetFireResistance returns the FireResistance field if non-nil, zero value otherwise.

### GetFireResistanceOk

`func (o *MonsterGrade) GetFireResistanceOk() (*int32, bool)`

GetFireResistanceOk returns a tuple with the FireResistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFireResistance

`func (o *MonsterGrade) SetFireResistance(v int32)`

SetFireResistance sets FireResistance field to given value.

### HasFireResistance

`func (o *MonsterGrade) HasFireResistance() bool`

HasFireResistance returns a boolean if a field has been set.

### GetWaterResistance

`func (o *MonsterGrade) GetWaterResistance() int32`

GetWaterResistance returns the WaterResistance field if non-nil, zero value otherwise.

### GetWaterResistanceOk

`func (o *MonsterGrade) GetWaterResistanceOk() (*int32, bool)`

GetWaterResistanceOk returns a tuple with the WaterResistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaterResistance

`func (o *MonsterGrade) SetWaterResistance(v int32)`

SetWaterResistance sets WaterResistance field to given value.

### HasWaterResistance

`func (o *MonsterGrade) HasWaterResistance() bool`

HasWaterResistance returns a boolean if a field has been set.

### GetAirResistance

`func (o *MonsterGrade) GetAirResistance() int32`

GetAirResistance returns the AirResistance field if non-nil, zero value otherwise.

### GetAirResistanceOk

`func (o *MonsterGrade) GetAirResistanceOk() (*int32, bool)`

GetAirResistanceOk returns a tuple with the AirResistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirResistance

`func (o *MonsterGrade) SetAirResistance(v int32)`

SetAirResistance sets AirResistance field to given value.

### HasAirResistance

`func (o *MonsterGrade) HasAirResistance() bool`

HasAirResistance returns a boolean if a field has been set.

### GetPaDodge

`func (o *MonsterGrade) GetPaDodge() int32`

GetPaDodge returns the PaDodge field if non-nil, zero value otherwise.

### GetPaDodgeOk

`func (o *MonsterGrade) GetPaDodgeOk() (*int32, bool)`

GetPaDodgeOk returns a tuple with the PaDodge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaDodge

`func (o *MonsterGrade) SetPaDodge(v int32)`

SetPaDodge sets PaDodge field to given value.

### HasPaDodge

`func (o *MonsterGrade) HasPaDodge() bool`

HasPaDodge returns a boolean if a field has been set.

### GetPmDodge

`func (o *MonsterGrade) GetPmDodge() int32`

GetPmDodge returns the PmDodge field if non-nil, zero value otherwise.

### GetPmDodgeOk

`func (o *MonsterGrade) GetPmDodgeOk() (*int32, bool)`

GetPmDodgeOk returns a tuple with the PmDodge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmDodge

`func (o *MonsterGrade) SetPmDodge(v int32)`

SetPmDodge sets PmDodge field to given value.

### HasPmDodge

`func (o *MonsterGrade) HasPmDodge() bool`

HasPmDodge returns a boolean if a field has been set.

### GetDamageBonus

`func (o *MonsterGrade) GetDamageBonus() int32`

GetDamageBonus returns the DamageBonus field if non-nil, zero value otherwise.

### GetDamageBonusOk

`func (o *MonsterGrade) GetDamageBonusOk() (*int32, bool)`

GetDamageBonusOk returns a tuple with the DamageBonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageBonus

`func (o *MonsterGrade) SetDamageBonus(v int32)`

SetDamageBonus sets DamageBonus field to given value.

### HasDamageBonus

`func (o *MonsterGrade) HasDamageBonus() bool`

HasDamageBonus returns a boolean if a field has been set.

### GetPercentDamageBonus

`func (o *MonsterGrade) GetPercentDamageBonus() int32`

GetPercentDamageBonus returns the PercentDamageBonus field if non-nil, zero value otherwise.

### GetPercentDamageBonusOk

`func (o *MonsterGrade) GetPercentDamageBonusOk() (*int32, bool)`

GetPercentDamageBonusOk returns a tuple with the PercentDamageBonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentDamageBonus

`func (o *MonsterGrade) SetPercentDamageBonus(v int32)`

SetPercentDamageBonus sets PercentDamageBonus field to given value.

### HasPercentDamageBonus

`func (o *MonsterGrade) HasPercentDamageBonus() bool`

HasPercentDamageBonus returns a boolean if a field has been set.

### GetCriticalHitBonus

`func (o *MonsterGrade) GetCriticalHitBonus() int32`

GetCriticalHitBonus returns the CriticalHitBonus field if non-nil, zero value otherwise.

### GetCriticalHitBonusOk

`func (o *MonsterGrade) GetCriticalHitBonusOk() (*int32, bool)`

GetCriticalHitBonusOk returns a tuple with the CriticalHitBonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalHitBonus

`func (o *MonsterGrade) SetCriticalHitBonus(v int32)`

SetCriticalHitBonus sets CriticalHitBonus field to given value.

### HasCriticalHitBonus

`func (o *MonsterGrade) HasCriticalHitBonus() bool`

HasCriticalHitBonus returns a boolean if a field has been set.

### GetInitiativeBonus

`func (o *MonsterGrade) GetInitiativeBonus() int32`

GetInitiativeBonus returns the InitiativeBonus field if non-nil, zero value otherwise.

### GetInitiativeBonusOk

`func (o *MonsterGrade) GetInitiativeBonusOk() (*int32, bool)`

GetInitiativeBonusOk returns a tuple with the InitiativeBonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiativeBonus

`func (o *MonsterGrade) SetInitiativeBonus(v int32)`

SetInitiativeBonus sets InitiativeBonus field to given value.

### HasInitiativeBonus

`func (o *MonsterGrade) HasInitiativeBonus() bool`

HasInitiativeBonus returns a boolean if a field has been set.

### GetTackleBonus

`func (o *MonsterGrade) GetTackleBonus() int32`

GetTackleBonus returns the TackleBonus field if non-nil, zero value otherwise.

### GetTackleBonusOk

`func (o *MonsterGrade) GetTackleBonusOk() (*int32, bool)`

GetTackleBonusOk returns a tuple with the TackleBonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTackleBonus

`func (o *MonsterGrade) SetTackleBonus(v int32)`

SetTackleBonus sets TackleBonus field to given value.

### HasTackleBonus

`func (o *MonsterGrade) HasTackleBonus() bool`

HasTackleBonus returns a boolean if a field has been set.

### GetTackleEvade

`func (o *MonsterGrade) GetTackleEvade() int32`

GetTackleEvade returns the TackleEvade field if non-nil, zero value otherwise.

### GetTackleEvadeOk

`func (o *MonsterGrade) GetTackleEvadeOk() (*int32, bool)`

GetTackleEvadeOk returns a tuple with the TackleEvade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTackleEvade

`func (o *MonsterGrade) SetTackleEvade(v int32)`

SetTackleEvade sets TackleEvade field to given value.

### HasTackleEvade

`func (o *MonsterGrade) HasTackleEvade() bool`

HasTackleEvade returns a boolean if a field has been set.

### GetMaxSummon

`func (o *MonsterGrade) GetMaxSummon() int32`

GetMaxSummon returns the MaxSummon field if non-nil, zero value otherwise.

### GetMaxSummonOk

`func (o *MonsterGrade) GetMaxSummonOk() (*int32, bool)`

GetMaxSummonOk returns a tuple with the MaxSummon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSummon

`func (o *MonsterGrade) SetMaxSummon(v int32)`

SetMaxSummon sets MaxSummon field to given value.

### HasMaxSummon

`func (o *MonsterGrade) HasMaxSummon() bool`

HasMaxSummon returns a boolean if a field has been set.

### GetBonusCharacteristics

`func (o *MonsterGrade) GetBonusCharacteristics() map[string]int32`

GetBonusCharacteristics returns the BonusCharacteristics field if non-nil, zero value otherwise.

### GetBonusCharacteristicsOk

`func (o *MonsterGrade) GetBonusCharacteristicsOk() (*map[string]int32, bool)`

GetBonusCharacteristicsOk returns a tuple with the BonusCharacteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonusCharacteristics

`func (o *MonsterGrade) SetBonusCharacteristics(v map[string]int32)`

SetBonusCharacteristics sets BonusCharacteristics field to given value.

### HasBonusCharacteristics

`func (o *MonsterGrade) HasBonusCharacteristics() bool`

HasBonusCharacteristics returns a boolean if a field has been set.

### GetSpells

`func (o *MonsterGrade) GetSpells() []MonsterGradeSpellsInner`

GetSpells returns the Spells field if non-nil, zero value otherwise.

### GetSpellsOk

`func (o *MonsterGrade) GetSpellsOk() (*[]MonsterGradeSpellsInner, bool)`

GetSpellsOk returns a tuple with the Spells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpells

`func (o *MonsterGrade) SetSpells(v []MonsterGradeSpellsInner)`

SetSpells sets Spells field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


