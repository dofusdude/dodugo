# CombatCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **string** |  | [optional] 
**Negated** | Pointer to **bool** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Reference** | Pointer to [**CombatReference**](CombatReference.md) |  | [optional] 

## Methods

### NewCombatCondition

`func NewCombatCondition() *CombatCondition`

NewCombatCondition instantiates a new CombatCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCombatConditionWithDefaults

`func NewCombatConditionWithDefaults() *CombatCondition`

NewCombatConditionWithDefaults instantiates a new CombatCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *CombatCondition) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CombatCondition) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CombatCondition) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *CombatCondition) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetNegated

`func (o *CombatCondition) GetNegated() bool`

GetNegated returns the Negated field if non-nil, zero value otherwise.

### GetNegatedOk

`func (o *CombatCondition) GetNegatedOk() (*bool, bool)`

GetNegatedOk returns a tuple with the Negated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegated

`func (o *CombatCondition) SetNegated(v bool)`

SetNegated sets Negated field to given value.

### HasNegated

`func (o *CombatCondition) HasNegated() bool`

HasNegated returns a boolean if a field has been set.

### GetSubject

`func (o *CombatCondition) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CombatCondition) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CombatCondition) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CombatCondition) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetReference

`func (o *CombatCondition) GetReference() CombatReference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CombatCondition) GetReferenceOk() (*CombatReference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CombatCondition) SetReference(v CombatReference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CombatCondition) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


