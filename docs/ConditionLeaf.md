# ConditionLeaf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsOperand** | Pointer to **bool** | always \&quot;true\&quot; for the leaf of a tree | [optional] [default to true]
**Condition** | Pointer to [**Condition**](Condition.md) |  | [optional] 

## Methods

### NewConditionLeaf

`func NewConditionLeaf() *ConditionLeaf`

NewConditionLeaf instantiates a new ConditionLeaf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionLeafWithDefaults

`func NewConditionLeafWithDefaults() *ConditionLeaf`

NewConditionLeafWithDefaults instantiates a new ConditionLeaf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsOperand

`func (o *ConditionLeaf) GetIsOperand() bool`

GetIsOperand returns the IsOperand field if non-nil, zero value otherwise.

### GetIsOperandOk

`func (o *ConditionLeaf) GetIsOperandOk() (*bool, bool)`

GetIsOperandOk returns a tuple with the IsOperand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOperand

`func (o *ConditionLeaf) SetIsOperand(v bool)`

SetIsOperand sets IsOperand field to given value.

### HasIsOperand

`func (o *ConditionLeaf) HasIsOperand() bool`

HasIsOperand returns a boolean if a field has been set.

### GetCondition

`func (o *ConditionLeaf) GetCondition() Condition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ConditionLeaf) GetConditionOk() (*Condition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ConditionLeaf) SetCondition(v Condition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ConditionLeaf) HasCondition() bool`

HasCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


