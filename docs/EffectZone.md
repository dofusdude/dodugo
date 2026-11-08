# EffectZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Raw** | Pointer to **string** | Original compact zone encoding when the source provides one. | [optional] 
**Shape** | [**ZoneShape**](ZoneShape.md) |  | 
**Size** | Pointer to **NullableInt32** | Effective outer size of the area. A point has size 0; whole-map and custom-cell shapes omit it. | [optional] 
**MinSize** | Pointer to **NullableInt32** | Effective inner size for hollow or minimum-size shapes. Omitted when the shape has no inner boundary. | [optional] 
**CellIds** | **[]int32** |  | 
**DamageDecreaseStepPercent** | Pointer to **int32** |  | [optional] 
**MaxDamageDecreaseApplyCount** | Pointer to **int32** |  | [optional] 
**ForcedDirection** | Pointer to **int32** |  | [optional] 
**IncludeCarried** | Pointer to **bool** |  | [optional] 
**StopAtTarget** | Pointer to **bool** |  | [optional] 
**OnlyAffectInSightLine** | Pointer to **bool** |  | [optional] 

## Methods

### NewEffectZone

`func NewEffectZone(shape ZoneShape, cellIds []int32, ) *EffectZone`

NewEffectZone instantiates a new EffectZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEffectZoneWithDefaults

`func NewEffectZoneWithDefaults() *EffectZone`

NewEffectZoneWithDefaults instantiates a new EffectZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRaw

`func (o *EffectZone) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *EffectZone) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *EffectZone) SetRaw(v string)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *EffectZone) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### GetShape

`func (o *EffectZone) GetShape() ZoneShape`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *EffectZone) GetShapeOk() (*ZoneShape, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *EffectZone) SetShape(v ZoneShape)`

SetShape sets Shape field to given value.


### GetSize

`func (o *EffectZone) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *EffectZone) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *EffectZone) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *EffectZone) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *EffectZone) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *EffectZone) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetMinSize

`func (o *EffectZone) GetMinSize() int32`

GetMinSize returns the MinSize field if non-nil, zero value otherwise.

### GetMinSizeOk

`func (o *EffectZone) GetMinSizeOk() (*int32, bool)`

GetMinSizeOk returns a tuple with the MinSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSize

`func (o *EffectZone) SetMinSize(v int32)`

SetMinSize sets MinSize field to given value.

### HasMinSize

`func (o *EffectZone) HasMinSize() bool`

HasMinSize returns a boolean if a field has been set.

### SetMinSizeNil

`func (o *EffectZone) SetMinSizeNil(b bool)`

 SetMinSizeNil sets the value for MinSize to be an explicit nil

### UnsetMinSize
`func (o *EffectZone) UnsetMinSize()`

UnsetMinSize ensures that no value is present for MinSize, not even an explicit nil
### GetCellIds

`func (o *EffectZone) GetCellIds() []int32`

GetCellIds returns the CellIds field if non-nil, zero value otherwise.

### GetCellIdsOk

`func (o *EffectZone) GetCellIdsOk() (*[]int32, bool)`

GetCellIdsOk returns a tuple with the CellIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellIds

`func (o *EffectZone) SetCellIds(v []int32)`

SetCellIds sets CellIds field to given value.


### GetDamageDecreaseStepPercent

`func (o *EffectZone) GetDamageDecreaseStepPercent() int32`

GetDamageDecreaseStepPercent returns the DamageDecreaseStepPercent field if non-nil, zero value otherwise.

### GetDamageDecreaseStepPercentOk

`func (o *EffectZone) GetDamageDecreaseStepPercentOk() (*int32, bool)`

GetDamageDecreaseStepPercentOk returns a tuple with the DamageDecreaseStepPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageDecreaseStepPercent

`func (o *EffectZone) SetDamageDecreaseStepPercent(v int32)`

SetDamageDecreaseStepPercent sets DamageDecreaseStepPercent field to given value.

### HasDamageDecreaseStepPercent

`func (o *EffectZone) HasDamageDecreaseStepPercent() bool`

HasDamageDecreaseStepPercent returns a boolean if a field has been set.

### GetMaxDamageDecreaseApplyCount

`func (o *EffectZone) GetMaxDamageDecreaseApplyCount() int32`

GetMaxDamageDecreaseApplyCount returns the MaxDamageDecreaseApplyCount field if non-nil, zero value otherwise.

### GetMaxDamageDecreaseApplyCountOk

`func (o *EffectZone) GetMaxDamageDecreaseApplyCountOk() (*int32, bool)`

GetMaxDamageDecreaseApplyCountOk returns a tuple with the MaxDamageDecreaseApplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDamageDecreaseApplyCount

`func (o *EffectZone) SetMaxDamageDecreaseApplyCount(v int32)`

SetMaxDamageDecreaseApplyCount sets MaxDamageDecreaseApplyCount field to given value.

### HasMaxDamageDecreaseApplyCount

`func (o *EffectZone) HasMaxDamageDecreaseApplyCount() bool`

HasMaxDamageDecreaseApplyCount returns a boolean if a field has been set.

### GetForcedDirection

`func (o *EffectZone) GetForcedDirection() int32`

GetForcedDirection returns the ForcedDirection field if non-nil, zero value otherwise.

### GetForcedDirectionOk

`func (o *EffectZone) GetForcedDirectionOk() (*int32, bool)`

GetForcedDirectionOk returns a tuple with the ForcedDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedDirection

`func (o *EffectZone) SetForcedDirection(v int32)`

SetForcedDirection sets ForcedDirection field to given value.

### HasForcedDirection

`func (o *EffectZone) HasForcedDirection() bool`

HasForcedDirection returns a boolean if a field has been set.

### GetIncludeCarried

`func (o *EffectZone) GetIncludeCarried() bool`

GetIncludeCarried returns the IncludeCarried field if non-nil, zero value otherwise.

### GetIncludeCarriedOk

`func (o *EffectZone) GetIncludeCarriedOk() (*bool, bool)`

GetIncludeCarriedOk returns a tuple with the IncludeCarried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCarried

`func (o *EffectZone) SetIncludeCarried(v bool)`

SetIncludeCarried sets IncludeCarried field to given value.

### HasIncludeCarried

`func (o *EffectZone) HasIncludeCarried() bool`

HasIncludeCarried returns a boolean if a field has been set.

### GetStopAtTarget

`func (o *EffectZone) GetStopAtTarget() bool`

GetStopAtTarget returns the StopAtTarget field if non-nil, zero value otherwise.

### GetStopAtTargetOk

`func (o *EffectZone) GetStopAtTargetOk() (*bool, bool)`

GetStopAtTargetOk returns a tuple with the StopAtTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopAtTarget

`func (o *EffectZone) SetStopAtTarget(v bool)`

SetStopAtTarget sets StopAtTarget field to given value.

### HasStopAtTarget

`func (o *EffectZone) HasStopAtTarget() bool`

HasStopAtTarget returns a boolean if a field has been set.

### GetOnlyAffectInSightLine

`func (o *EffectZone) GetOnlyAffectInSightLine() bool`

GetOnlyAffectInSightLine returns the OnlyAffectInSightLine field if non-nil, zero value otherwise.

### GetOnlyAffectInSightLineOk

`func (o *EffectZone) GetOnlyAffectInSightLineOk() (*bool, bool)`

GetOnlyAffectInSightLineOk returns a tuple with the OnlyAffectInSightLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyAffectInSightLine

`func (o *EffectZone) SetOnlyAffectInSightLine(v bool)`

SetOnlyAffectInSightLine sets OnlyAffectInSightLine field to given value.

### HasOnlyAffectInSightLine

`func (o *EffectZone) HasOnlyAffectInSightLine() bool`

HasOnlyAffectInSightLine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


