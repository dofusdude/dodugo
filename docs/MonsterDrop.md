# MonsterDrop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | [**CombatReference**](CombatReference.md) |  | 
**RatesPercent** | [**[]MonsterDropRatesPercentInner**](MonsterDropRatesPercentInner.md) | Base drop percentages indexed by MonsterGrade.grade, not spell grade or array position. A value of 5 means 5 percent, not 0.05 percent. These are not spawn percentages. | 
**HasConditions** | **bool** | The source reports restrictions. Can be true without a complete client-side condition tree because server-only predicates are not exposed. | 
**DropId** | Pointer to **int32** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Condition** | Pointer to **map[string]interface{}** |  | [optional] 
**HiddenIfInvalidConditions** | Pointer to **bool** |  | [optional] 
**DisableDropModifier** | Pointer to **bool** |  | [optional] 
**SpecificCoefficients** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewMonsterDrop

`func NewMonsterDrop(item CombatReference, ratesPercent []MonsterDropRatesPercentInner, hasConditions bool, ) *MonsterDrop`

NewMonsterDrop instantiates a new MonsterDrop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonsterDropWithDefaults

`func NewMonsterDropWithDefaults() *MonsterDrop`

NewMonsterDropWithDefaults instantiates a new MonsterDrop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *MonsterDrop) GetItem() CombatReference`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *MonsterDrop) GetItemOk() (*CombatReference, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *MonsterDrop) SetItem(v CombatReference)`

SetItem sets Item field to given value.


### GetRatesPercent

`func (o *MonsterDrop) GetRatesPercent() []MonsterDropRatesPercentInner`

GetRatesPercent returns the RatesPercent field if non-nil, zero value otherwise.

### GetRatesPercentOk

`func (o *MonsterDrop) GetRatesPercentOk() (*[]MonsterDropRatesPercentInner, bool)`

GetRatesPercentOk returns a tuple with the RatesPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatesPercent

`func (o *MonsterDrop) SetRatesPercent(v []MonsterDropRatesPercentInner)`

SetRatesPercent sets RatesPercent field to given value.


### GetHasConditions

`func (o *MonsterDrop) GetHasConditions() bool`

GetHasConditions returns the HasConditions field if non-nil, zero value otherwise.

### GetHasConditionsOk

`func (o *MonsterDrop) GetHasConditionsOk() (*bool, bool)`

GetHasConditionsOk returns a tuple with the HasConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConditions

`func (o *MonsterDrop) SetHasConditions(v bool)`

SetHasConditions sets HasConditions field to given value.


### GetDropId

`func (o *MonsterDrop) GetDropId() int32`

GetDropId returns the DropId field if non-nil, zero value otherwise.

### GetDropIdOk

`func (o *MonsterDrop) GetDropIdOk() (*int32, bool)`

GetDropIdOk returns a tuple with the DropId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropId

`func (o *MonsterDrop) SetDropId(v int32)`

SetDropId sets DropId field to given value.

### HasDropId

`func (o *MonsterDrop) HasDropId() bool`

HasDropId returns a boolean if a field has been set.

### GetCount

`func (o *MonsterDrop) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MonsterDrop) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MonsterDrop) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *MonsterDrop) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCondition

`func (o *MonsterDrop) GetCondition() map[string]interface{}`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *MonsterDrop) GetConditionOk() (*map[string]interface{}, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *MonsterDrop) SetCondition(v map[string]interface{})`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *MonsterDrop) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *MonsterDrop) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *MonsterDrop) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetHiddenIfInvalidConditions

`func (o *MonsterDrop) GetHiddenIfInvalidConditions() bool`

GetHiddenIfInvalidConditions returns the HiddenIfInvalidConditions field if non-nil, zero value otherwise.

### GetHiddenIfInvalidConditionsOk

`func (o *MonsterDrop) GetHiddenIfInvalidConditionsOk() (*bool, bool)`

GetHiddenIfInvalidConditionsOk returns a tuple with the HiddenIfInvalidConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenIfInvalidConditions

`func (o *MonsterDrop) SetHiddenIfInvalidConditions(v bool)`

SetHiddenIfInvalidConditions sets HiddenIfInvalidConditions field to given value.

### HasHiddenIfInvalidConditions

`func (o *MonsterDrop) HasHiddenIfInvalidConditions() bool`

HasHiddenIfInvalidConditions returns a boolean if a field has been set.

### GetDisableDropModifier

`func (o *MonsterDrop) GetDisableDropModifier() bool`

GetDisableDropModifier returns the DisableDropModifier field if non-nil, zero value otherwise.

### GetDisableDropModifierOk

`func (o *MonsterDrop) GetDisableDropModifierOk() (*bool, bool)`

GetDisableDropModifierOk returns a tuple with the DisableDropModifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDropModifier

`func (o *MonsterDrop) SetDisableDropModifier(v bool)`

SetDisableDropModifier sets DisableDropModifier field to given value.

### HasDisableDropModifier

`func (o *MonsterDrop) HasDisableDropModifier() bool`

HasDisableDropModifier returns a boolean if a field has been set.

### GetSpecificCoefficients

`func (o *MonsterDrop) GetSpecificCoefficients() []map[string]interface{}`

GetSpecificCoefficients returns the SpecificCoefficients field if non-nil, zero value otherwise.

### GetSpecificCoefficientsOk

`func (o *MonsterDrop) GetSpecificCoefficientsOk() (*[]map[string]interface{}, bool)`

GetSpecificCoefficientsOk returns a tuple with the SpecificCoefficients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificCoefficients

`func (o *MonsterDrop) SetSpecificCoefficients(v []map[string]interface{})`

SetSpecificCoefficients sets SpecificCoefficients field to given value.

### HasSpecificCoefficients

`func (o *MonsterDrop) HasSpecificCoefficients() bool`

HasSpecificCoefficients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


