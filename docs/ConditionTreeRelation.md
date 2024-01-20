# ConditionTreeRelation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsOperand** | Pointer to **bool** | always \&quot;false\&quot; for relations | [optional] [default to false]
**Relation** | Pointer to **string** | \&quot;and\&quot;, \&quot;or\&quot; | [optional] [default to "and"]
**Children** | Pointer to [**[]ConditionTreeNode**](ConditionTreeNode.md) |  | [optional] 

## Methods

### NewConditionTreeRelation

`func NewConditionTreeRelation() *ConditionTreeRelation`

NewConditionTreeRelation instantiates a new ConditionTreeRelation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionTreeRelationWithDefaults

`func NewConditionTreeRelationWithDefaults() *ConditionTreeRelation`

NewConditionTreeRelationWithDefaults instantiates a new ConditionTreeRelation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsOperand

`func (o *ConditionTreeRelation) GetIsOperand() bool`

GetIsOperand returns the IsOperand field if non-nil, zero value otherwise.

### GetIsOperandOk

`func (o *ConditionTreeRelation) GetIsOperandOk() (*bool, bool)`

GetIsOperandOk returns a tuple with the IsOperand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOperand

`func (o *ConditionTreeRelation) SetIsOperand(v bool)`

SetIsOperand sets IsOperand field to given value.

### HasIsOperand

`func (o *ConditionTreeRelation) HasIsOperand() bool`

HasIsOperand returns a boolean if a field has been set.

### GetRelation

`func (o *ConditionTreeRelation) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *ConditionTreeRelation) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *ConditionTreeRelation) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *ConditionTreeRelation) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetChildren

`func (o *ConditionTreeRelation) GetChildren() []ConditionTreeNode`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ConditionTreeRelation) GetChildrenOk() (*[]ConditionTreeNode, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ConditionTreeRelation) SetChildren(v []ConditionTreeNode)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *ConditionTreeRelation) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


