# Effect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntMinimum** | Pointer to **int32** | minimum int value, can be a single if ignore_int_max and no ignore_int_min | [optional] 
**IntMaximum** | Pointer to **int32** | maximum int value, if not ignore_int_max and not ignore_int_min, the effect has a range value | [optional] 
**Type** | Pointer to [**EffectType**](EffectType.md) |  | [optional] 
**IgnoreIntMin** | Pointer to **bool** | ignore the int min field because the actual value is a string. For readability the templated field is the only important field in this case.  | [optional] 
**IgnoreIntMax** | Pointer to **bool** | ignore the int max field, if ignore_int_min is true, int min is a single value | [optional] 
**Formatted** | Pointer to **string** | all fields from above encoded in a single string | [optional] 

## Methods

### NewEffect

`func NewEffect() *Effect`

NewEffect instantiates a new Effect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEffectWithDefaults

`func NewEffectWithDefaults() *Effect`

NewEffectWithDefaults instantiates a new Effect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntMinimum

`func (o *Effect) GetIntMinimum() int32`

GetIntMinimum returns the IntMinimum field if non-nil, zero value otherwise.

### GetIntMinimumOk

`func (o *Effect) GetIntMinimumOk() (*int32, bool)`

GetIntMinimumOk returns a tuple with the IntMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntMinimum

`func (o *Effect) SetIntMinimum(v int32)`

SetIntMinimum sets IntMinimum field to given value.

### HasIntMinimum

`func (o *Effect) HasIntMinimum() bool`

HasIntMinimum returns a boolean if a field has been set.

### GetIntMaximum

`func (o *Effect) GetIntMaximum() int32`

GetIntMaximum returns the IntMaximum field if non-nil, zero value otherwise.

### GetIntMaximumOk

`func (o *Effect) GetIntMaximumOk() (*int32, bool)`

GetIntMaximumOk returns a tuple with the IntMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntMaximum

`func (o *Effect) SetIntMaximum(v int32)`

SetIntMaximum sets IntMaximum field to given value.

### HasIntMaximum

`func (o *Effect) HasIntMaximum() bool`

HasIntMaximum returns a boolean if a field has been set.

### GetType

`func (o *Effect) GetType() EffectType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Effect) GetTypeOk() (*EffectType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Effect) SetType(v EffectType)`

SetType sets Type field to given value.

### HasType

`func (o *Effect) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIgnoreIntMin

`func (o *Effect) GetIgnoreIntMin() bool`

GetIgnoreIntMin returns the IgnoreIntMin field if non-nil, zero value otherwise.

### GetIgnoreIntMinOk

`func (o *Effect) GetIgnoreIntMinOk() (*bool, bool)`

GetIgnoreIntMinOk returns a tuple with the IgnoreIntMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreIntMin

`func (o *Effect) SetIgnoreIntMin(v bool)`

SetIgnoreIntMin sets IgnoreIntMin field to given value.

### HasIgnoreIntMin

`func (o *Effect) HasIgnoreIntMin() bool`

HasIgnoreIntMin returns a boolean if a field has been set.

### GetIgnoreIntMax

`func (o *Effect) GetIgnoreIntMax() bool`

GetIgnoreIntMax returns the IgnoreIntMax field if non-nil, zero value otherwise.

### GetIgnoreIntMaxOk

`func (o *Effect) GetIgnoreIntMaxOk() (*bool, bool)`

GetIgnoreIntMaxOk returns a tuple with the IgnoreIntMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreIntMax

`func (o *Effect) SetIgnoreIntMax(v bool)`

SetIgnoreIntMax sets IgnoreIntMax field to given value.

### HasIgnoreIntMax

`func (o *Effect) HasIgnoreIntMax() bool`

HasIgnoreIntMax returns a boolean if a field has been set.

### GetFormatted

`func (o *Effect) GetFormatted() string`

GetFormatted returns the Formatted field if non-nil, zero value otherwise.

### GetFormattedOk

`func (o *Effect) GetFormattedOk() (*string, bool)`

GetFormattedOk returns a tuple with the Formatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatted

`func (o *Effect) SetFormatted(v string)`

SetFormatted sets Formatted field to given value.

### HasFormatted

`func (o *Effect) HasFormatted() bool`

HasFormatted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


