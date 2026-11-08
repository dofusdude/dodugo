# MonsterGlobalDrop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasConditions** | Pointer to **bool** |  | [optional] 
**Item** | [**CombatReference**](CombatReference.md) |  | 
**MinPercent** | **float64** |  | 
**MaxPercent** | **float64** |  | 
**Condition** | Pointer to **map[string]interface{}** |  | [optional] 
**AlterationId** | Pointer to **int32** |  | [optional] 
**DisableDropModifier** | Pointer to **bool** |  | [optional] 

## Methods

### NewMonsterGlobalDrop

`func NewMonsterGlobalDrop(item CombatReference, minPercent float64, maxPercent float64, ) *MonsterGlobalDrop`

NewMonsterGlobalDrop instantiates a new MonsterGlobalDrop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonsterGlobalDropWithDefaults

`func NewMonsterGlobalDropWithDefaults() *MonsterGlobalDrop`

NewMonsterGlobalDropWithDefaults instantiates a new MonsterGlobalDrop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasConditions

`func (o *MonsterGlobalDrop) GetHasConditions() bool`

GetHasConditions returns the HasConditions field if non-nil, zero value otherwise.

### GetHasConditionsOk

`func (o *MonsterGlobalDrop) GetHasConditionsOk() (*bool, bool)`

GetHasConditionsOk returns a tuple with the HasConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConditions

`func (o *MonsterGlobalDrop) SetHasConditions(v bool)`

SetHasConditions sets HasConditions field to given value.

### HasHasConditions

`func (o *MonsterGlobalDrop) HasHasConditions() bool`

HasHasConditions returns a boolean if a field has been set.

### GetItem

`func (o *MonsterGlobalDrop) GetItem() CombatReference`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *MonsterGlobalDrop) GetItemOk() (*CombatReference, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *MonsterGlobalDrop) SetItem(v CombatReference)`

SetItem sets Item field to given value.


### GetMinPercent

`func (o *MonsterGlobalDrop) GetMinPercent() float64`

GetMinPercent returns the MinPercent field if non-nil, zero value otherwise.

### GetMinPercentOk

`func (o *MonsterGlobalDrop) GetMinPercentOk() (*float64, bool)`

GetMinPercentOk returns a tuple with the MinPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPercent

`func (o *MonsterGlobalDrop) SetMinPercent(v float64)`

SetMinPercent sets MinPercent field to given value.


### GetMaxPercent

`func (o *MonsterGlobalDrop) GetMaxPercent() float64`

GetMaxPercent returns the MaxPercent field if non-nil, zero value otherwise.

### GetMaxPercentOk

`func (o *MonsterGlobalDrop) GetMaxPercentOk() (*float64, bool)`

GetMaxPercentOk returns a tuple with the MaxPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPercent

`func (o *MonsterGlobalDrop) SetMaxPercent(v float64)`

SetMaxPercent sets MaxPercent field to given value.


### GetCondition

`func (o *MonsterGlobalDrop) GetCondition() map[string]interface{}`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *MonsterGlobalDrop) GetConditionOk() (*map[string]interface{}, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *MonsterGlobalDrop) SetCondition(v map[string]interface{})`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *MonsterGlobalDrop) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *MonsterGlobalDrop) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *MonsterGlobalDrop) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetAlterationId

`func (o *MonsterGlobalDrop) GetAlterationId() int32`

GetAlterationId returns the AlterationId field if non-nil, zero value otherwise.

### GetAlterationIdOk

`func (o *MonsterGlobalDrop) GetAlterationIdOk() (*int32, bool)`

GetAlterationIdOk returns a tuple with the AlterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlterationId

`func (o *MonsterGlobalDrop) SetAlterationId(v int32)`

SetAlterationId sets AlterationId field to given value.

### HasAlterationId

`func (o *MonsterGlobalDrop) HasAlterationId() bool`

HasAlterationId returns a boolean if a field has been set.

### GetDisableDropModifier

`func (o *MonsterGlobalDrop) GetDisableDropModifier() bool`

GetDisableDropModifier returns the DisableDropModifier field if non-nil, zero value otherwise.

### GetDisableDropModifierOk

`func (o *MonsterGlobalDrop) GetDisableDropModifierOk() (*bool, bool)`

GetDisableDropModifierOk returns a tuple with the DisableDropModifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDropModifier

`func (o *MonsterGlobalDrop) SetDisableDropModifier(v bool)`

SetDisableDropModifier sets DisableDropModifier field to given value.

### HasDisableDropModifier

`func (o *MonsterGlobalDrop) HasDisableDropModifier() bool`

HasDisableDropModifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


