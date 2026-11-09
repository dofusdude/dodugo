# Condition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **string** | Comparison to apply to the interpreted requirement value. | [optional] 
**IntValue** | Pointer to **int32** |  | [optional] 
**Element** | Pointer to [**TranslatedId**](TranslatedId.md) |  | [optional] 
**Kind** | Pointer to **string** | Interpreted requirement category, such as profession_level or item_received_age. Use this with the operator, values and references rather than decoding original condition tokens. | [optional] 
**Unit** | Pointer to **string** | Unit associated with the requirement value when the interpreted condition supplies one. | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**StringValue** | Pointer to **string** |  | [optional] 
**ValueLabel** | Pointer to **string** |  | [optional] 
**Reference** | Pointer to [**CombatReference**](CombatReference.md) |  | [optional] 
**SecondaryReference** | Pointer to [**CombatReference**](CombatReference.md) |  | [optional] 

## Methods

### NewCondition

`func NewCondition() *Condition`

NewCondition instantiates a new Condition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionWithDefaults

`func NewConditionWithDefaults() *Condition`

NewConditionWithDefaults instantiates a new Condition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *Condition) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Condition) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Condition) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *Condition) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetIntValue

`func (o *Condition) GetIntValue() int32`

GetIntValue returns the IntValue field if non-nil, zero value otherwise.

### GetIntValueOk

`func (o *Condition) GetIntValueOk() (*int32, bool)`

GetIntValueOk returns a tuple with the IntValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntValue

`func (o *Condition) SetIntValue(v int32)`

SetIntValue sets IntValue field to given value.

### HasIntValue

`func (o *Condition) HasIntValue() bool`

HasIntValue returns a boolean if a field has been set.

### GetElement

`func (o *Condition) GetElement() TranslatedId`

GetElement returns the Element field if non-nil, zero value otherwise.

### GetElementOk

`func (o *Condition) GetElementOk() (*TranslatedId, bool)`

GetElementOk returns a tuple with the Element field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElement

`func (o *Condition) SetElement(v TranslatedId)`

SetElement sets Element field to given value.

### HasElement

`func (o *Condition) HasElement() bool`

HasElement returns a boolean if a field has been set.

### GetKind

`func (o *Condition) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Condition) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Condition) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Condition) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetUnit

`func (o *Condition) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Condition) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Condition) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Condition) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetQuantity

`func (o *Condition) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Condition) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Condition) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Condition) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStringValue

`func (o *Condition) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *Condition) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *Condition) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *Condition) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.

### GetValueLabel

`func (o *Condition) GetValueLabel() string`

GetValueLabel returns the ValueLabel field if non-nil, zero value otherwise.

### GetValueLabelOk

`func (o *Condition) GetValueLabelOk() (*string, bool)`

GetValueLabelOk returns a tuple with the ValueLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueLabel

`func (o *Condition) SetValueLabel(v string)`

SetValueLabel sets ValueLabel field to given value.

### HasValueLabel

`func (o *Condition) HasValueLabel() bool`

HasValueLabel returns a boolean if a field has been set.

### GetReference

`func (o *Condition) GetReference() CombatReference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Condition) GetReferenceOk() (*CombatReference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Condition) SetReference(v CombatReference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *Condition) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSecondaryReference

`func (o *Condition) GetSecondaryReference() CombatReference`

GetSecondaryReference returns the SecondaryReference field if non-nil, zero value otherwise.

### GetSecondaryReferenceOk

`func (o *Condition) GetSecondaryReferenceOk() (*CombatReference, bool)`

GetSecondaryReferenceOk returns a tuple with the SecondaryReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryReference

`func (o *Condition) SetSecondaryReference(v CombatReference)`

SetSecondaryReference sets SecondaryReference field to given value.

### HasSecondaryReference

`func (o *Condition) HasSecondaryReference() bool`

HasSecondaryReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


