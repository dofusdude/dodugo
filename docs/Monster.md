# Monster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | **int32** |  | 
**Name** | **string** |  | 
**ImageUrls** | Pointer to [**Images**](Images.md) |  | [optional] 
**Race** | [**MonsterRace**](MonsterRace.md) |  | 
**Grades** | [**[]MonsterGrade**](MonsterGrade.md) | Variants of this monster. Each has its own combat level, statistics and assigned spells. The monster grade number is not a spell grade number or combat level. | 
**SpawnLocations** | [**[]MonsterSpawnLocation**](MonsterSpawnLocation.md) | Known spawn subareas and their geographic parents. This is not live population data or a per-map spawn probability. Maps listed within a subarea are not guaranteed individual spawn points. | 
**Drops** | [**[]MonsterDrop**](MonsterDrop.md) |  | 
**GlobalDrops** | Pointer to [**[]MonsterGlobalDrop**](MonsterGlobalDrop.md) |  | [optional] 
**Appearance** | Pointer to [**ClassAppearance**](ClassAppearance.md) |  | [optional] 
**IsBounty** | Pointer to **bool** |  | [optional] 
**CorrespondingMiniBoss** | Pointer to [**CombatReference**](CombatReference.md) |  | [optional] 
**SummonCost** | Pointer to **int32** |  | [optional] 
**SpeedAdjust** | Pointer to **int32** |  | [optional] 
**ScaleGradeReference** | Pointer to **int32** |  | [optional] 
**Achievements** | Pointer to [**[]CombatReference**](CombatReference.md) |  | [optional] 
**Souls** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CharacteristicRatios** | Pointer to **[][]float64** |  | [optional] 

## Methods

### NewMonster

`func NewMonster(ankamaId int32, name string, race MonsterRace, grades []MonsterGrade, spawnLocations []MonsterSpawnLocation, drops []MonsterDrop, ) *Monster`

NewMonster instantiates a new Monster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonsterWithDefaults

`func NewMonsterWithDefaults() *Monster`

NewMonsterWithDefaults instantiates a new Monster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *Monster) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *Monster) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *Monster) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.


### GetName

`func (o *Monster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Monster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Monster) SetName(v string)`

SetName sets Name field to given value.


### GetImageUrls

`func (o *Monster) GetImageUrls() Images`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *Monster) GetImageUrlsOk() (*Images, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *Monster) SetImageUrls(v Images)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *Monster) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.

### GetRace

`func (o *Monster) GetRace() MonsterRace`

GetRace returns the Race field if non-nil, zero value otherwise.

### GetRaceOk

`func (o *Monster) GetRaceOk() (*MonsterRace, bool)`

GetRaceOk returns a tuple with the Race field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRace

`func (o *Monster) SetRace(v MonsterRace)`

SetRace sets Race field to given value.


### GetGrades

`func (o *Monster) GetGrades() []MonsterGrade`

GetGrades returns the Grades field if non-nil, zero value otherwise.

### GetGradesOk

`func (o *Monster) GetGradesOk() (*[]MonsterGrade, bool)`

GetGradesOk returns a tuple with the Grades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrades

`func (o *Monster) SetGrades(v []MonsterGrade)`

SetGrades sets Grades field to given value.


### GetSpawnLocations

`func (o *Monster) GetSpawnLocations() []MonsterSpawnLocation`

GetSpawnLocations returns the SpawnLocations field if non-nil, zero value otherwise.

### GetSpawnLocationsOk

`func (o *Monster) GetSpawnLocationsOk() (*[]MonsterSpawnLocation, bool)`

GetSpawnLocationsOk returns a tuple with the SpawnLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpawnLocations

`func (o *Monster) SetSpawnLocations(v []MonsterSpawnLocation)`

SetSpawnLocations sets SpawnLocations field to given value.


### GetDrops

`func (o *Monster) GetDrops() []MonsterDrop`

GetDrops returns the Drops field if non-nil, zero value otherwise.

### GetDropsOk

`func (o *Monster) GetDropsOk() (*[]MonsterDrop, bool)`

GetDropsOk returns a tuple with the Drops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrops

`func (o *Monster) SetDrops(v []MonsterDrop)`

SetDrops sets Drops field to given value.


### GetGlobalDrops

`func (o *Monster) GetGlobalDrops() []MonsterGlobalDrop`

GetGlobalDrops returns the GlobalDrops field if non-nil, zero value otherwise.

### GetGlobalDropsOk

`func (o *Monster) GetGlobalDropsOk() (*[]MonsterGlobalDrop, bool)`

GetGlobalDropsOk returns a tuple with the GlobalDrops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalDrops

`func (o *Monster) SetGlobalDrops(v []MonsterGlobalDrop)`

SetGlobalDrops sets GlobalDrops field to given value.

### HasGlobalDrops

`func (o *Monster) HasGlobalDrops() bool`

HasGlobalDrops returns a boolean if a field has been set.

### GetAppearance

`func (o *Monster) GetAppearance() ClassAppearance`

GetAppearance returns the Appearance field if non-nil, zero value otherwise.

### GetAppearanceOk

`func (o *Monster) GetAppearanceOk() (*ClassAppearance, bool)`

GetAppearanceOk returns a tuple with the Appearance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppearance

`func (o *Monster) SetAppearance(v ClassAppearance)`

SetAppearance sets Appearance field to given value.

### HasAppearance

`func (o *Monster) HasAppearance() bool`

HasAppearance returns a boolean if a field has been set.

### GetIsBounty

`func (o *Monster) GetIsBounty() bool`

GetIsBounty returns the IsBounty field if non-nil, zero value otherwise.

### GetIsBountyOk

`func (o *Monster) GetIsBountyOk() (*bool, bool)`

GetIsBountyOk returns a tuple with the IsBounty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBounty

`func (o *Monster) SetIsBounty(v bool)`

SetIsBounty sets IsBounty field to given value.

### HasIsBounty

`func (o *Monster) HasIsBounty() bool`

HasIsBounty returns a boolean if a field has been set.

### GetCorrespondingMiniBoss

`func (o *Monster) GetCorrespondingMiniBoss() CombatReference`

GetCorrespondingMiniBoss returns the CorrespondingMiniBoss field if non-nil, zero value otherwise.

### GetCorrespondingMiniBossOk

`func (o *Monster) GetCorrespondingMiniBossOk() (*CombatReference, bool)`

GetCorrespondingMiniBossOk returns a tuple with the CorrespondingMiniBoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondingMiniBoss

`func (o *Monster) SetCorrespondingMiniBoss(v CombatReference)`

SetCorrespondingMiniBoss sets CorrespondingMiniBoss field to given value.

### HasCorrespondingMiniBoss

`func (o *Monster) HasCorrespondingMiniBoss() bool`

HasCorrespondingMiniBoss returns a boolean if a field has been set.

### GetSummonCost

`func (o *Monster) GetSummonCost() int32`

GetSummonCost returns the SummonCost field if non-nil, zero value otherwise.

### GetSummonCostOk

`func (o *Monster) GetSummonCostOk() (*int32, bool)`

GetSummonCostOk returns a tuple with the SummonCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummonCost

`func (o *Monster) SetSummonCost(v int32)`

SetSummonCost sets SummonCost field to given value.

### HasSummonCost

`func (o *Monster) HasSummonCost() bool`

HasSummonCost returns a boolean if a field has been set.

### GetSpeedAdjust

`func (o *Monster) GetSpeedAdjust() int32`

GetSpeedAdjust returns the SpeedAdjust field if non-nil, zero value otherwise.

### GetSpeedAdjustOk

`func (o *Monster) GetSpeedAdjustOk() (*int32, bool)`

GetSpeedAdjustOk returns a tuple with the SpeedAdjust field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedAdjust

`func (o *Monster) SetSpeedAdjust(v int32)`

SetSpeedAdjust sets SpeedAdjust field to given value.

### HasSpeedAdjust

`func (o *Monster) HasSpeedAdjust() bool`

HasSpeedAdjust returns a boolean if a field has been set.

### GetScaleGradeReference

`func (o *Monster) GetScaleGradeReference() int32`

GetScaleGradeReference returns the ScaleGradeReference field if non-nil, zero value otherwise.

### GetScaleGradeReferenceOk

`func (o *Monster) GetScaleGradeReferenceOk() (*int32, bool)`

GetScaleGradeReferenceOk returns a tuple with the ScaleGradeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleGradeReference

`func (o *Monster) SetScaleGradeReference(v int32)`

SetScaleGradeReference sets ScaleGradeReference field to given value.

### HasScaleGradeReference

`func (o *Monster) HasScaleGradeReference() bool`

HasScaleGradeReference returns a boolean if a field has been set.

### GetAchievements

`func (o *Monster) GetAchievements() []CombatReference`

GetAchievements returns the Achievements field if non-nil, zero value otherwise.

### GetAchievementsOk

`func (o *Monster) GetAchievementsOk() (*[]CombatReference, bool)`

GetAchievementsOk returns a tuple with the Achievements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievements

`func (o *Monster) SetAchievements(v []CombatReference)`

SetAchievements sets Achievements field to given value.

### HasAchievements

`func (o *Monster) HasAchievements() bool`

HasAchievements returns a boolean if a field has been set.

### GetSouls

`func (o *Monster) GetSouls() []map[string]interface{}`

GetSouls returns the Souls field if non-nil, zero value otherwise.

### GetSoulsOk

`func (o *Monster) GetSoulsOk() (*[]map[string]interface{}, bool)`

GetSoulsOk returns a tuple with the Souls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSouls

`func (o *Monster) SetSouls(v []map[string]interface{})`

SetSouls sets Souls field to given value.

### HasSouls

`func (o *Monster) HasSouls() bool`

HasSouls returns a boolean if a field has been set.

### GetCharacteristicRatios

`func (o *Monster) GetCharacteristicRatios() [][]float64`

GetCharacteristicRatios returns the CharacteristicRatios field if non-nil, zero value otherwise.

### GetCharacteristicRatiosOk

`func (o *Monster) GetCharacteristicRatiosOk() (*[][]float64, bool)`

GetCharacteristicRatiosOk returns a tuple with the CharacteristicRatios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacteristicRatios

`func (o *Monster) SetCharacteristicRatios(v [][]float64)`

SetCharacteristicRatios sets CharacteristicRatios field to given value.

### HasCharacteristicRatios

`func (o *Monster) HasCharacteristicRatios() bool`

HasCharacteristicRatios returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


