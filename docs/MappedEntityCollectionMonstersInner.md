# MappedEntityCollectionMonstersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | **int32** |  | 
**Name** | **string** |  | 
**ImageUrls** | Pointer to [**Images**](Images.md) |  | [optional] 
**Race** | [**MonsterRace**](MonsterRace.md) |  | 
**Grades** | [**[]MonsterGrade**](MonsterGrade.md) | Monster statistics and spells grouped by the monster grade that uses them. | 
**SpawnLocations** | [**[]MonsterSpawnLocation**](MonsterSpawnLocation.md) | Spawn subareas taken from the monster&#39;s authoritative subarea list. | 
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

### NewMappedEntityCollectionMonstersInner

`func NewMappedEntityCollectionMonstersInner(ankamaId int32, name string, race MonsterRace, grades []MonsterGrade, spawnLocations []MonsterSpawnLocation, drops []MonsterDrop, ) *MappedEntityCollectionMonstersInner`

NewMappedEntityCollectionMonstersInner instantiates a new MappedEntityCollectionMonstersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMappedEntityCollectionMonstersInnerWithDefaults

`func NewMappedEntityCollectionMonstersInnerWithDefaults() *MappedEntityCollectionMonstersInner`

NewMappedEntityCollectionMonstersInnerWithDefaults instantiates a new MappedEntityCollectionMonstersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *MappedEntityCollectionMonstersInner) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *MappedEntityCollectionMonstersInner) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *MappedEntityCollectionMonstersInner) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.


### GetName

`func (o *MappedEntityCollectionMonstersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MappedEntityCollectionMonstersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MappedEntityCollectionMonstersInner) SetName(v string)`

SetName sets Name field to given value.


### GetImageUrls

`func (o *MappedEntityCollectionMonstersInner) GetImageUrls() Images`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *MappedEntityCollectionMonstersInner) GetImageUrlsOk() (*Images, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *MappedEntityCollectionMonstersInner) SetImageUrls(v Images)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *MappedEntityCollectionMonstersInner) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.

### GetRace

`func (o *MappedEntityCollectionMonstersInner) GetRace() MonsterRace`

GetRace returns the Race field if non-nil, zero value otherwise.

### GetRaceOk

`func (o *MappedEntityCollectionMonstersInner) GetRaceOk() (*MonsterRace, bool)`

GetRaceOk returns a tuple with the Race field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRace

`func (o *MappedEntityCollectionMonstersInner) SetRace(v MonsterRace)`

SetRace sets Race field to given value.


### GetGrades

`func (o *MappedEntityCollectionMonstersInner) GetGrades() []MonsterGrade`

GetGrades returns the Grades field if non-nil, zero value otherwise.

### GetGradesOk

`func (o *MappedEntityCollectionMonstersInner) GetGradesOk() (*[]MonsterGrade, bool)`

GetGradesOk returns a tuple with the Grades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrades

`func (o *MappedEntityCollectionMonstersInner) SetGrades(v []MonsterGrade)`

SetGrades sets Grades field to given value.


### GetSpawnLocations

`func (o *MappedEntityCollectionMonstersInner) GetSpawnLocations() []MonsterSpawnLocation`

GetSpawnLocations returns the SpawnLocations field if non-nil, zero value otherwise.

### GetSpawnLocationsOk

`func (o *MappedEntityCollectionMonstersInner) GetSpawnLocationsOk() (*[]MonsterSpawnLocation, bool)`

GetSpawnLocationsOk returns a tuple with the SpawnLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpawnLocations

`func (o *MappedEntityCollectionMonstersInner) SetSpawnLocations(v []MonsterSpawnLocation)`

SetSpawnLocations sets SpawnLocations field to given value.


### GetDrops

`func (o *MappedEntityCollectionMonstersInner) GetDrops() []MonsterDrop`

GetDrops returns the Drops field if non-nil, zero value otherwise.

### GetDropsOk

`func (o *MappedEntityCollectionMonstersInner) GetDropsOk() (*[]MonsterDrop, bool)`

GetDropsOk returns a tuple with the Drops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrops

`func (o *MappedEntityCollectionMonstersInner) SetDrops(v []MonsterDrop)`

SetDrops sets Drops field to given value.


### GetGlobalDrops

`func (o *MappedEntityCollectionMonstersInner) GetGlobalDrops() []MonsterGlobalDrop`

GetGlobalDrops returns the GlobalDrops field if non-nil, zero value otherwise.

### GetGlobalDropsOk

`func (o *MappedEntityCollectionMonstersInner) GetGlobalDropsOk() (*[]MonsterGlobalDrop, bool)`

GetGlobalDropsOk returns a tuple with the GlobalDrops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalDrops

`func (o *MappedEntityCollectionMonstersInner) SetGlobalDrops(v []MonsterGlobalDrop)`

SetGlobalDrops sets GlobalDrops field to given value.

### HasGlobalDrops

`func (o *MappedEntityCollectionMonstersInner) HasGlobalDrops() bool`

HasGlobalDrops returns a boolean if a field has been set.

### GetAppearance

`func (o *MappedEntityCollectionMonstersInner) GetAppearance() ClassAppearance`

GetAppearance returns the Appearance field if non-nil, zero value otherwise.

### GetAppearanceOk

`func (o *MappedEntityCollectionMonstersInner) GetAppearanceOk() (*ClassAppearance, bool)`

GetAppearanceOk returns a tuple with the Appearance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppearance

`func (o *MappedEntityCollectionMonstersInner) SetAppearance(v ClassAppearance)`

SetAppearance sets Appearance field to given value.

### HasAppearance

`func (o *MappedEntityCollectionMonstersInner) HasAppearance() bool`

HasAppearance returns a boolean if a field has been set.

### GetIsBounty

`func (o *MappedEntityCollectionMonstersInner) GetIsBounty() bool`

GetIsBounty returns the IsBounty field if non-nil, zero value otherwise.

### GetIsBountyOk

`func (o *MappedEntityCollectionMonstersInner) GetIsBountyOk() (*bool, bool)`

GetIsBountyOk returns a tuple with the IsBounty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBounty

`func (o *MappedEntityCollectionMonstersInner) SetIsBounty(v bool)`

SetIsBounty sets IsBounty field to given value.

### HasIsBounty

`func (o *MappedEntityCollectionMonstersInner) HasIsBounty() bool`

HasIsBounty returns a boolean if a field has been set.

### GetCorrespondingMiniBoss

`func (o *MappedEntityCollectionMonstersInner) GetCorrespondingMiniBoss() CombatReference`

GetCorrespondingMiniBoss returns the CorrespondingMiniBoss field if non-nil, zero value otherwise.

### GetCorrespondingMiniBossOk

`func (o *MappedEntityCollectionMonstersInner) GetCorrespondingMiniBossOk() (*CombatReference, bool)`

GetCorrespondingMiniBossOk returns a tuple with the CorrespondingMiniBoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondingMiniBoss

`func (o *MappedEntityCollectionMonstersInner) SetCorrespondingMiniBoss(v CombatReference)`

SetCorrespondingMiniBoss sets CorrespondingMiniBoss field to given value.

### HasCorrespondingMiniBoss

`func (o *MappedEntityCollectionMonstersInner) HasCorrespondingMiniBoss() bool`

HasCorrespondingMiniBoss returns a boolean if a field has been set.

### GetSummonCost

`func (o *MappedEntityCollectionMonstersInner) GetSummonCost() int32`

GetSummonCost returns the SummonCost field if non-nil, zero value otherwise.

### GetSummonCostOk

`func (o *MappedEntityCollectionMonstersInner) GetSummonCostOk() (*int32, bool)`

GetSummonCostOk returns a tuple with the SummonCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummonCost

`func (o *MappedEntityCollectionMonstersInner) SetSummonCost(v int32)`

SetSummonCost sets SummonCost field to given value.

### HasSummonCost

`func (o *MappedEntityCollectionMonstersInner) HasSummonCost() bool`

HasSummonCost returns a boolean if a field has been set.

### GetSpeedAdjust

`func (o *MappedEntityCollectionMonstersInner) GetSpeedAdjust() int32`

GetSpeedAdjust returns the SpeedAdjust field if non-nil, zero value otherwise.

### GetSpeedAdjustOk

`func (o *MappedEntityCollectionMonstersInner) GetSpeedAdjustOk() (*int32, bool)`

GetSpeedAdjustOk returns a tuple with the SpeedAdjust field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedAdjust

`func (o *MappedEntityCollectionMonstersInner) SetSpeedAdjust(v int32)`

SetSpeedAdjust sets SpeedAdjust field to given value.

### HasSpeedAdjust

`func (o *MappedEntityCollectionMonstersInner) HasSpeedAdjust() bool`

HasSpeedAdjust returns a boolean if a field has been set.

### GetScaleGradeReference

`func (o *MappedEntityCollectionMonstersInner) GetScaleGradeReference() int32`

GetScaleGradeReference returns the ScaleGradeReference field if non-nil, zero value otherwise.

### GetScaleGradeReferenceOk

`func (o *MappedEntityCollectionMonstersInner) GetScaleGradeReferenceOk() (*int32, bool)`

GetScaleGradeReferenceOk returns a tuple with the ScaleGradeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleGradeReference

`func (o *MappedEntityCollectionMonstersInner) SetScaleGradeReference(v int32)`

SetScaleGradeReference sets ScaleGradeReference field to given value.

### HasScaleGradeReference

`func (o *MappedEntityCollectionMonstersInner) HasScaleGradeReference() bool`

HasScaleGradeReference returns a boolean if a field has been set.

### GetAchievements

`func (o *MappedEntityCollectionMonstersInner) GetAchievements() []CombatReference`

GetAchievements returns the Achievements field if non-nil, zero value otherwise.

### GetAchievementsOk

`func (o *MappedEntityCollectionMonstersInner) GetAchievementsOk() (*[]CombatReference, bool)`

GetAchievementsOk returns a tuple with the Achievements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievements

`func (o *MappedEntityCollectionMonstersInner) SetAchievements(v []CombatReference)`

SetAchievements sets Achievements field to given value.

### HasAchievements

`func (o *MappedEntityCollectionMonstersInner) HasAchievements() bool`

HasAchievements returns a boolean if a field has been set.

### GetSouls

`func (o *MappedEntityCollectionMonstersInner) GetSouls() []map[string]interface{}`

GetSouls returns the Souls field if non-nil, zero value otherwise.

### GetSoulsOk

`func (o *MappedEntityCollectionMonstersInner) GetSoulsOk() (*[]map[string]interface{}, bool)`

GetSoulsOk returns a tuple with the Souls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSouls

`func (o *MappedEntityCollectionMonstersInner) SetSouls(v []map[string]interface{})`

SetSouls sets Souls field to given value.

### HasSouls

`func (o *MappedEntityCollectionMonstersInner) HasSouls() bool`

HasSouls returns a boolean if a field has been set.

### GetCharacteristicRatios

`func (o *MappedEntityCollectionMonstersInner) GetCharacteristicRatios() [][]float64`

GetCharacteristicRatios returns the CharacteristicRatios field if non-nil, zero value otherwise.

### GetCharacteristicRatiosOk

`func (o *MappedEntityCollectionMonstersInner) GetCharacteristicRatiosOk() (*[][]float64, bool)`

GetCharacteristicRatiosOk returns a tuple with the CharacteristicRatios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacteristicRatios

`func (o *MappedEntityCollectionMonstersInner) SetCharacteristicRatios(v [][]float64)`

SetCharacteristicRatios sets CharacteristicRatios field to given value.

### HasCharacteristicRatios

`func (o *MappedEntityCollectionMonstersInner) HasCharacteristicRatios() bool`

HasCharacteristicRatios returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


