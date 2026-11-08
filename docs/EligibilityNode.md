# EligibilityNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsOperand** | Pointer to **bool** |  | [optional] 
**Relation** | Pointer to **string** |  | [optional] 
**Children** | Pointer to [**[]EligibilityNode**](EligibilityNode.md) |  | [optional] 
**Condition** | Pointer to [**EligibilityNodeCondition**](EligibilityNodeCondition.md) |  | [optional] 

## Methods

### NewEligibilityNode

`func NewEligibilityNode() *EligibilityNode`

NewEligibilityNode instantiates a new EligibilityNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEligibilityNodeWithDefaults

`func NewEligibilityNodeWithDefaults() *EligibilityNode`

NewEligibilityNodeWithDefaults instantiates a new EligibilityNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsOperand

`func (o *EligibilityNode) GetIsOperand() bool`

GetIsOperand returns the IsOperand field if non-nil, zero value otherwise.

### GetIsOperandOk

`func (o *EligibilityNode) GetIsOperandOk() (*bool, bool)`

GetIsOperandOk returns a tuple with the IsOperand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOperand

`func (o *EligibilityNode) SetIsOperand(v bool)`

SetIsOperand sets IsOperand field to given value.

### HasIsOperand

`func (o *EligibilityNode) HasIsOperand() bool`

HasIsOperand returns a boolean if a field has been set.

### GetRelation

`func (o *EligibilityNode) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *EligibilityNode) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *EligibilityNode) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *EligibilityNode) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetChildren

`func (o *EligibilityNode) GetChildren() []EligibilityNode`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *EligibilityNode) GetChildrenOk() (*[]EligibilityNode, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *EligibilityNode) SetChildren(v []EligibilityNode)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *EligibilityNode) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetCondition

`func (o *EligibilityNode) GetCondition() EligibilityNodeCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *EligibilityNode) GetConditionOk() (*EligibilityNodeCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *EligibilityNode) SetCondition(v EligibilityNodeCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *EligibilityNode) HasCondition() bool`

HasCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


