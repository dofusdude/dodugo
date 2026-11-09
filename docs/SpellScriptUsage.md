# SpellScriptUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | Pointer to **int32** |  | [optional] 
**ScriptId** | Pointer to **int32** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Random** | Pointer to **int32** |  | [optional] 
**RandomGroup** | Pointer to **int32** |  | [optional] 
**SequenceGroup** | Pointer to **int32** |  | [optional] 
**ActivationMask** | Pointer to [**TargetMask**](TargetMask.md) |  | [optional] 
**CasterMask** | Pointer to [**TargetMask**](TargetMask.md) |  | [optional] 
**TargetMask** | Pointer to [**TargetMask**](TargetMask.md) |  | [optional] 
**ActivationZone** | Pointer to [**EffectZone**](EffectZone.md) |  | [optional] 
**TargetZone** | Pointer to [**EffectZone**](EffectZone.md) |  | [optional] 
**Condition** | Pointer to **map[string]interface{}** |  | [optional] 
**TargetsAffectedConcurrently** | Pointer to **bool** |  | [optional] 
**TargetTreatedAsCaster** | Pointer to **bool** |  | [optional] 
**SpellLevelIds** | Pointer to **[]int32** | Spell-level definition IDs to which this script binding applies. Match SpellLevel.ankama_id, not Spell.ankama_id or SpellLevel.grade. | [optional] 

## Methods

### NewSpellScriptUsage

`func NewSpellScriptUsage() *SpellScriptUsage`

NewSpellScriptUsage instantiates a new SpellScriptUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpellScriptUsageWithDefaults

`func NewSpellScriptUsageWithDefaults() *SpellScriptUsage`

NewSpellScriptUsageWithDefaults instantiates a new SpellScriptUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *SpellScriptUsage) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *SpellScriptUsage) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *SpellScriptUsage) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.

### HasAnkamaId

`func (o *SpellScriptUsage) HasAnkamaId() bool`

HasAnkamaId returns a boolean if a field has been set.

### GetScriptId

`func (o *SpellScriptUsage) GetScriptId() int32`

GetScriptId returns the ScriptId field if non-nil, zero value otherwise.

### GetScriptIdOk

`func (o *SpellScriptUsage) GetScriptIdOk() (*int32, bool)`

GetScriptIdOk returns a tuple with the ScriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptId

`func (o *SpellScriptUsage) SetScriptId(v int32)`

SetScriptId sets ScriptId field to given value.

### HasScriptId

`func (o *SpellScriptUsage) HasScriptId() bool`

HasScriptId returns a boolean if a field has been set.

### GetOrder

`func (o *SpellScriptUsage) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *SpellScriptUsage) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *SpellScriptUsage) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *SpellScriptUsage) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetRandom

`func (o *SpellScriptUsage) GetRandom() int32`

GetRandom returns the Random field if non-nil, zero value otherwise.

### GetRandomOk

`func (o *SpellScriptUsage) GetRandomOk() (*int32, bool)`

GetRandomOk returns a tuple with the Random field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandom

`func (o *SpellScriptUsage) SetRandom(v int32)`

SetRandom sets Random field to given value.

### HasRandom

`func (o *SpellScriptUsage) HasRandom() bool`

HasRandom returns a boolean if a field has been set.

### GetRandomGroup

`func (o *SpellScriptUsage) GetRandomGroup() int32`

GetRandomGroup returns the RandomGroup field if non-nil, zero value otherwise.

### GetRandomGroupOk

`func (o *SpellScriptUsage) GetRandomGroupOk() (*int32, bool)`

GetRandomGroupOk returns a tuple with the RandomGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomGroup

`func (o *SpellScriptUsage) SetRandomGroup(v int32)`

SetRandomGroup sets RandomGroup field to given value.

### HasRandomGroup

`func (o *SpellScriptUsage) HasRandomGroup() bool`

HasRandomGroup returns a boolean if a field has been set.

### GetSequenceGroup

`func (o *SpellScriptUsage) GetSequenceGroup() int32`

GetSequenceGroup returns the SequenceGroup field if non-nil, zero value otherwise.

### GetSequenceGroupOk

`func (o *SpellScriptUsage) GetSequenceGroupOk() (*int32, bool)`

GetSequenceGroupOk returns a tuple with the SequenceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceGroup

`func (o *SpellScriptUsage) SetSequenceGroup(v int32)`

SetSequenceGroup sets SequenceGroup field to given value.

### HasSequenceGroup

`func (o *SpellScriptUsage) HasSequenceGroup() bool`

HasSequenceGroup returns a boolean if a field has been set.

### GetActivationMask

`func (o *SpellScriptUsage) GetActivationMask() TargetMask`

GetActivationMask returns the ActivationMask field if non-nil, zero value otherwise.

### GetActivationMaskOk

`func (o *SpellScriptUsage) GetActivationMaskOk() (*TargetMask, bool)`

GetActivationMaskOk returns a tuple with the ActivationMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationMask

`func (o *SpellScriptUsage) SetActivationMask(v TargetMask)`

SetActivationMask sets ActivationMask field to given value.

### HasActivationMask

`func (o *SpellScriptUsage) HasActivationMask() bool`

HasActivationMask returns a boolean if a field has been set.

### GetCasterMask

`func (o *SpellScriptUsage) GetCasterMask() TargetMask`

GetCasterMask returns the CasterMask field if non-nil, zero value otherwise.

### GetCasterMaskOk

`func (o *SpellScriptUsage) GetCasterMaskOk() (*TargetMask, bool)`

GetCasterMaskOk returns a tuple with the CasterMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasterMask

`func (o *SpellScriptUsage) SetCasterMask(v TargetMask)`

SetCasterMask sets CasterMask field to given value.

### HasCasterMask

`func (o *SpellScriptUsage) HasCasterMask() bool`

HasCasterMask returns a boolean if a field has been set.

### GetTargetMask

`func (o *SpellScriptUsage) GetTargetMask() TargetMask`

GetTargetMask returns the TargetMask field if non-nil, zero value otherwise.

### GetTargetMaskOk

`func (o *SpellScriptUsage) GetTargetMaskOk() (*TargetMask, bool)`

GetTargetMaskOk returns a tuple with the TargetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMask

`func (o *SpellScriptUsage) SetTargetMask(v TargetMask)`

SetTargetMask sets TargetMask field to given value.

### HasTargetMask

`func (o *SpellScriptUsage) HasTargetMask() bool`

HasTargetMask returns a boolean if a field has been set.

### GetActivationZone

`func (o *SpellScriptUsage) GetActivationZone() EffectZone`

GetActivationZone returns the ActivationZone field if non-nil, zero value otherwise.

### GetActivationZoneOk

`func (o *SpellScriptUsage) GetActivationZoneOk() (*EffectZone, bool)`

GetActivationZoneOk returns a tuple with the ActivationZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationZone

`func (o *SpellScriptUsage) SetActivationZone(v EffectZone)`

SetActivationZone sets ActivationZone field to given value.

### HasActivationZone

`func (o *SpellScriptUsage) HasActivationZone() bool`

HasActivationZone returns a boolean if a field has been set.

### GetTargetZone

`func (o *SpellScriptUsage) GetTargetZone() EffectZone`

GetTargetZone returns the TargetZone field if non-nil, zero value otherwise.

### GetTargetZoneOk

`func (o *SpellScriptUsage) GetTargetZoneOk() (*EffectZone, bool)`

GetTargetZoneOk returns a tuple with the TargetZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetZone

`func (o *SpellScriptUsage) SetTargetZone(v EffectZone)`

SetTargetZone sets TargetZone field to given value.

### HasTargetZone

`func (o *SpellScriptUsage) HasTargetZone() bool`

HasTargetZone returns a boolean if a field has been set.

### GetCondition

`func (o *SpellScriptUsage) GetCondition() map[string]interface{}`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *SpellScriptUsage) GetConditionOk() (*map[string]interface{}, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *SpellScriptUsage) SetCondition(v map[string]interface{})`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *SpellScriptUsage) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *SpellScriptUsage) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *SpellScriptUsage) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetTargetsAffectedConcurrently

`func (o *SpellScriptUsage) GetTargetsAffectedConcurrently() bool`

GetTargetsAffectedConcurrently returns the TargetsAffectedConcurrently field if non-nil, zero value otherwise.

### GetTargetsAffectedConcurrentlyOk

`func (o *SpellScriptUsage) GetTargetsAffectedConcurrentlyOk() (*bool, bool)`

GetTargetsAffectedConcurrentlyOk returns a tuple with the TargetsAffectedConcurrently field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetsAffectedConcurrently

`func (o *SpellScriptUsage) SetTargetsAffectedConcurrently(v bool)`

SetTargetsAffectedConcurrently sets TargetsAffectedConcurrently field to given value.

### HasTargetsAffectedConcurrently

`func (o *SpellScriptUsage) HasTargetsAffectedConcurrently() bool`

HasTargetsAffectedConcurrently returns a boolean if a field has been set.

### GetTargetTreatedAsCaster

`func (o *SpellScriptUsage) GetTargetTreatedAsCaster() bool`

GetTargetTreatedAsCaster returns the TargetTreatedAsCaster field if non-nil, zero value otherwise.

### GetTargetTreatedAsCasterOk

`func (o *SpellScriptUsage) GetTargetTreatedAsCasterOk() (*bool, bool)`

GetTargetTreatedAsCasterOk returns a tuple with the TargetTreatedAsCaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTreatedAsCaster

`func (o *SpellScriptUsage) SetTargetTreatedAsCaster(v bool)`

SetTargetTreatedAsCaster sets TargetTreatedAsCaster field to given value.

### HasTargetTreatedAsCaster

`func (o *SpellScriptUsage) HasTargetTreatedAsCaster() bool`

HasTargetTreatedAsCaster returns a boolean if a field has been set.

### GetSpellLevelIds

`func (o *SpellScriptUsage) GetSpellLevelIds() []int32`

GetSpellLevelIds returns the SpellLevelIds field if non-nil, zero value otherwise.

### GetSpellLevelIdsOk

`func (o *SpellScriptUsage) GetSpellLevelIdsOk() (*[]int32, bool)`

GetSpellLevelIdsOk returns a tuple with the SpellLevelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpellLevelIds

`func (o *SpellScriptUsage) SetSpellLevelIds(v []int32)`

SetSpellLevelIds sets SpellLevelIds field to given value.

### HasSpellLevelIds

`func (o *SpellScriptUsage) HasSpellLevelIds() bool`

HasSpellLevelIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


