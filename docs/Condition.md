# Condition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **string** |  | [optional] 
**IntValue** | Pointer to **int32** |  | [optional] 
**Element** | Pointer to [**TranslatedId**](TranslatedId.md) |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


