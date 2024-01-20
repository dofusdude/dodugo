# ConditionTreeLeaf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsOperand** | Pointer to **bool** | always \&quot;true\&quot; for the leaf of a tree | [optional] [default to true]
**Condition** | Pointer to [**ConditionEntry**](ConditionEntry.md) |  | [optional] 

## Methods

### NewConditionTreeLeaf

`func NewConditionTreeLeaf() *ConditionTreeLeaf`

NewConditionTreeLeaf instantiates a new ConditionTreeLeaf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionTreeLeafWithDefaults

`func NewConditionTreeLeafWithDefaults() *ConditionTreeLeaf`

NewConditionTreeLeafWithDefaults instantiates a new ConditionTreeLeaf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsOperand

`func (o *ConditionTreeLeaf) GetIsOperand() bool`

GetIsOperand returns the IsOperand field if non-nil, zero value otherwise.

### GetIsOperandOk

`func (o *ConditionTreeLeaf) GetIsOperandOk() (*bool, bool)`

GetIsOperandOk returns a tuple with the IsOperand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOperand

`func (o *ConditionTreeLeaf) SetIsOperand(v bool)`

SetIsOperand sets IsOperand field to given value.

### HasIsOperand

`func (o *ConditionTreeLeaf) HasIsOperand() bool`

HasIsOperand returns a boolean if a field has been set.

### GetCondition

`func (o *ConditionTreeLeaf) GetCondition() ConditionEntry`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ConditionTreeLeaf) GetConditionOk() (*ConditionEntry, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ConditionTreeLeaf) SetCondition(v ConditionEntry)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ConditionTreeLeaf) HasCondition() bool`

HasCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


