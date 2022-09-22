# ConditionEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **string** |  | [optional] 
**IntValue** | Pointer to **int32** |  | [optional] 
**Element** | Pointer to [**EffectsEntryType**](EffectsEntryType.md) |  | [optional] 

## Methods

### NewConditionEntry

`func NewConditionEntry() *ConditionEntry`

NewConditionEntry instantiates a new ConditionEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionEntryWithDefaults

`func NewConditionEntryWithDefaults() *ConditionEntry`

NewConditionEntryWithDefaults instantiates a new ConditionEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *ConditionEntry) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ConditionEntry) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ConditionEntry) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *ConditionEntry) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetIntValue

`func (o *ConditionEntry) GetIntValue() int32`

GetIntValue returns the IntValue field if non-nil, zero value otherwise.

### GetIntValueOk

`func (o *ConditionEntry) GetIntValueOk() (*int32, bool)`

GetIntValueOk returns a tuple with the IntValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntValue

`func (o *ConditionEntry) SetIntValue(v int32)`

SetIntValue sets IntValue field to given value.

### HasIntValue

`func (o *ConditionEntry) HasIntValue() bool`

HasIntValue returns a boolean if a field has been set.

### GetElement

`func (o *ConditionEntry) GetElement() EffectsEntryType`

GetElement returns the Element field if non-nil, zero value otherwise.

### GetElementOk

`func (o *ConditionEntry) GetElementOk() (*EffectsEntryType, bool)`

GetElementOk returns a tuple with the Element field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElement

`func (o *ConditionEntry) SetElement(v EffectsEntryType)`

SetElement sets Element field to given value.

### HasElement

`func (o *ConditionEntry) HasElement() bool`

HasElement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


