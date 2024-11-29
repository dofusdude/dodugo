# ConditionRelation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsOperand** | Pointer to **bool** | always \&quot;false\&quot; for relations | [optional] [default to false]
**Relation** | Pointer to **string** | \&quot;and\&quot;, \&quot;or\&quot; | [optional] [default to "and"]
**Children** | Pointer to [**[]ConditionNode**](ConditionNode.md) |  | [optional] 

## Methods

### NewConditionRelation

`func NewConditionRelation() *ConditionRelation`

NewConditionRelation instantiates a new ConditionRelation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionRelationWithDefaults

`func NewConditionRelationWithDefaults() *ConditionRelation`

NewConditionRelationWithDefaults instantiates a new ConditionRelation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsOperand

`func (o *ConditionRelation) GetIsOperand() bool`

GetIsOperand returns the IsOperand field if non-nil, zero value otherwise.

### GetIsOperandOk

`func (o *ConditionRelation) GetIsOperandOk() (*bool, bool)`

GetIsOperandOk returns a tuple with the IsOperand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOperand

`func (o *ConditionRelation) SetIsOperand(v bool)`

SetIsOperand sets IsOperand field to given value.

### HasIsOperand

`func (o *ConditionRelation) HasIsOperand() bool`

HasIsOperand returns a boolean if a field has been set.

### GetRelation

`func (o *ConditionRelation) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *ConditionRelation) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *ConditionRelation) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *ConditionRelation) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetChildren

`func (o *ConditionRelation) GetChildren() []*ConditionNode`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ConditionRelation) GetChildrenOk() (*[]*ConditionNode, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ConditionRelation) SetChildren(v []*ConditionNode)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *ConditionRelation) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


