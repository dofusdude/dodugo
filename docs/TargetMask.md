# TargetMask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Raw** | **string** | Original game mask retained for lossless debugging. | 
**Logic** | Pointer to **string** |  | [optional] 
**Conditions** | Pointer to [**[]CombatCondition**](CombatCondition.md) |  | [optional] 
**Selectors** | [**[]TargetSelector**](TargetSelector.md) |  | 

## Methods

### NewTargetMask

`func NewTargetMask(raw string, selectors []TargetSelector, ) *TargetMask`

NewTargetMask instantiates a new TargetMask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetMaskWithDefaults

`func NewTargetMaskWithDefaults() *TargetMask`

NewTargetMaskWithDefaults instantiates a new TargetMask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRaw

`func (o *TargetMask) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *TargetMask) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *TargetMask) SetRaw(v string)`

SetRaw sets Raw field to given value.


### GetLogic

`func (o *TargetMask) GetLogic() string`

GetLogic returns the Logic field if non-nil, zero value otherwise.

### GetLogicOk

`func (o *TargetMask) GetLogicOk() (*string, bool)`

GetLogicOk returns a tuple with the Logic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogic

`func (o *TargetMask) SetLogic(v string)`

SetLogic sets Logic field to given value.

### HasLogic

`func (o *TargetMask) HasLogic() bool`

HasLogic returns a boolean if a field has been set.

### GetConditions

`func (o *TargetMask) GetConditions() []CombatCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *TargetMask) GetConditionsOk() (*[]CombatCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *TargetMask) SetConditions(v []CombatCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *TargetMask) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetSelectors

`func (o *TargetMask) GetSelectors() []TargetSelector`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *TargetMask) GetSelectorsOk() (*[]TargetSelector, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *TargetMask) SetSelectors(v []TargetSelector)`

SetSelectors sets Selectors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


