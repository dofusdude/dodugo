# ConditionNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsOperand** | Pointer to **bool** | always \&quot;true\&quot; for the leaf of a tree | [optional] [default to true]
**Relation** | Pointer to **string** | \&quot;and\&quot;, \&quot;or\&quot; | [optional] [default to "and"]
**Children** | Pointer to [**[]ConditionNode**](ConditionNode.md) |  | [optional] 
**Condition** | Pointer to [**Condition**](Condition.md) |  | [optional] 

## Methods

### NewConditionNode

`func NewConditionNode() *ConditionNode`

NewConditionNode instantiates a new ConditionNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionNodeWithDefaults

`func NewConditionNodeWithDefaults() *ConditionNode`

NewConditionNodeWithDefaults instantiates a new ConditionNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsOperand

`func (o *ConditionNode) GetIsOperand() bool`

GetIsOperand returns the IsOperand field if non-nil, zero value otherwise.

### GetIsOperandOk

`func (o *ConditionNode) GetIsOperandOk() (*bool, bool)`

GetIsOperandOk returns a tuple with the IsOperand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOperand

`func (o *ConditionNode) SetIsOperand(v bool)`

SetIsOperand sets IsOperand field to given value.

### HasIsOperand

`func (o *ConditionNode) HasIsOperand() bool`

HasIsOperand returns a boolean if a field has been set.

### GetRelation

`func (o *ConditionNode) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *ConditionNode) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *ConditionNode) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *ConditionNode) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetChildren

`func (o *ConditionNode) GetChildren() []*ConditionNode`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ConditionNode) GetChildrenOk() (*[]*ConditionNode, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ConditionNode) SetChildren(v []*ConditionNode)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *ConditionNode) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetCondition

`func (o *ConditionNode) GetCondition() Condition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ConditionNode) GetConditionOk() (*Condition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ConditionNode) SetCondition(v Condition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ConditionNode) HasCondition() bool`

HasCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


