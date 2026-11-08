# EligibilityNodeCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **string** |  | [optional] 
**IntValue** | Pointer to **int32** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**Kind** | Pointer to **string** | Verified semantic requirement type, for example item_received_age or profession_level. | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**StringValue** | Pointer to **string** |  | [optional] 
**ValueLabel** | Pointer to **string** |  | [optional] 
**Element** | Pointer to [**EligibilityNodeConditionElement**](EligibilityNodeConditionElement.md) |  | [optional] 
**Reference** | Pointer to [**CombatReference**](CombatReference.md) |  | [optional] 
**SecondaryReference** | Pointer to [**CombatReference**](CombatReference.md) |  | [optional] 

## Methods

### NewEligibilityNodeCondition

`func NewEligibilityNodeCondition() *EligibilityNodeCondition`

NewEligibilityNodeCondition instantiates a new EligibilityNodeCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEligibilityNodeConditionWithDefaults

`func NewEligibilityNodeConditionWithDefaults() *EligibilityNodeCondition`

NewEligibilityNodeConditionWithDefaults instantiates a new EligibilityNodeCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *EligibilityNodeCondition) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *EligibilityNodeCondition) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *EligibilityNodeCondition) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *EligibilityNodeCondition) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetIntValue

`func (o *EligibilityNodeCondition) GetIntValue() int32`

GetIntValue returns the IntValue field if non-nil, zero value otherwise.

### GetIntValueOk

`func (o *EligibilityNodeCondition) GetIntValueOk() (*int32, bool)`

GetIntValueOk returns a tuple with the IntValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntValue

`func (o *EligibilityNodeCondition) SetIntValue(v int32)`

SetIntValue sets IntValue field to given value.

### HasIntValue

`func (o *EligibilityNodeCondition) HasIntValue() bool`

HasIntValue returns a boolean if a field has been set.

### GetQuantity

`func (o *EligibilityNodeCondition) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *EligibilityNodeCondition) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *EligibilityNodeCondition) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *EligibilityNodeCondition) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetKind

`func (o *EligibilityNodeCondition) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EligibilityNodeCondition) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EligibilityNodeCondition) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *EligibilityNodeCondition) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetUnit

`func (o *EligibilityNodeCondition) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *EligibilityNodeCondition) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *EligibilityNodeCondition) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *EligibilityNodeCondition) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetStringValue

`func (o *EligibilityNodeCondition) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *EligibilityNodeCondition) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *EligibilityNodeCondition) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *EligibilityNodeCondition) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.

### GetValueLabel

`func (o *EligibilityNodeCondition) GetValueLabel() string`

GetValueLabel returns the ValueLabel field if non-nil, zero value otherwise.

### GetValueLabelOk

`func (o *EligibilityNodeCondition) GetValueLabelOk() (*string, bool)`

GetValueLabelOk returns a tuple with the ValueLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueLabel

`func (o *EligibilityNodeCondition) SetValueLabel(v string)`

SetValueLabel sets ValueLabel field to given value.

### HasValueLabel

`func (o *EligibilityNodeCondition) HasValueLabel() bool`

HasValueLabel returns a boolean if a field has been set.

### GetElement

`func (o *EligibilityNodeCondition) GetElement() EligibilityNodeConditionElement`

GetElement returns the Element field if non-nil, zero value otherwise.

### GetElementOk

`func (o *EligibilityNodeCondition) GetElementOk() (*EligibilityNodeConditionElement, bool)`

GetElementOk returns a tuple with the Element field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElement

`func (o *EligibilityNodeCondition) SetElement(v EligibilityNodeConditionElement)`

SetElement sets Element field to given value.

### HasElement

`func (o *EligibilityNodeCondition) HasElement() bool`

HasElement returns a boolean if a field has been set.

### GetReference

`func (o *EligibilityNodeCondition) GetReference() CombatReference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *EligibilityNodeCondition) GetReferenceOk() (*CombatReference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *EligibilityNodeCondition) SetReference(v CombatReference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *EligibilityNodeCondition) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSecondaryReference

`func (o *EligibilityNodeCondition) GetSecondaryReference() CombatReference`

GetSecondaryReference returns the SecondaryReference field if non-nil, zero value otherwise.

### GetSecondaryReferenceOk

`func (o *EligibilityNodeCondition) GetSecondaryReferenceOk() (*CombatReference, bool)`

GetSecondaryReferenceOk returns a tuple with the SecondaryReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryReference

`func (o *EligibilityNodeCondition) SetSecondaryReference(v CombatReference)`

SetSecondaryReference sets SecondaryReference field to given value.

### HasSecondaryReference

`func (o *EligibilityNodeCondition) HasSecondaryReference() bool`

HasSecondaryReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


