# SetEffectsEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntMinimum** | Pointer to **int32** | minimum int value, can be a single if ignore_int_max and no ignore_int_min | [optional] 
**IntMaximum** | Pointer to **int32** | maximum int value, if not ignore_int_max and not ignore_int_min, the effect has a range value | [optional] 
**Type** | Pointer to [**SetEffectsEntryType**](SetEffectsEntryType.md) |  | [optional] 
**IgnoreIntMin** | Pointer to **bool** | ignore the int min field because the actual value is a string. For readability the templated field is the only important field in this case.  | [optional] 
**IgnoreIntMax** | Pointer to **bool** | ignore the int max field, if ignore_int_min is true, int min is a single value | [optional] 
**Formatted** | Pointer to **string** | all fields from above encoded in a single string | [optional] 
**ItemCombination** | Pointer to **int32** | how many items it needs to trigger this effect with the given set | [optional] 

## Methods

### NewSetEffectsEntry

`func NewSetEffectsEntry() *SetEffectsEntry`

NewSetEffectsEntry instantiates a new SetEffectsEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetEffectsEntryWithDefaults

`func NewSetEffectsEntryWithDefaults() *SetEffectsEntry`

NewSetEffectsEntryWithDefaults instantiates a new SetEffectsEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntMinimum

`func (o *SetEffectsEntry) GetIntMinimum() int32`

GetIntMinimum returns the IntMinimum field if non-nil, zero value otherwise.

### GetIntMinimumOk

`func (o *SetEffectsEntry) GetIntMinimumOk() (*int32, bool)`

GetIntMinimumOk returns a tuple with the IntMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntMinimum

`func (o *SetEffectsEntry) SetIntMinimum(v int32)`

SetIntMinimum sets IntMinimum field to given value.

### HasIntMinimum

`func (o *SetEffectsEntry) HasIntMinimum() bool`

HasIntMinimum returns a boolean if a field has been set.

### GetIntMaximum

`func (o *SetEffectsEntry) GetIntMaximum() int32`

GetIntMaximum returns the IntMaximum field if non-nil, zero value otherwise.

### GetIntMaximumOk

`func (o *SetEffectsEntry) GetIntMaximumOk() (*int32, bool)`

GetIntMaximumOk returns a tuple with the IntMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntMaximum

`func (o *SetEffectsEntry) SetIntMaximum(v int32)`

SetIntMaximum sets IntMaximum field to given value.

### HasIntMaximum

`func (o *SetEffectsEntry) HasIntMaximum() bool`

HasIntMaximum returns a boolean if a field has been set.

### GetType

`func (o *SetEffectsEntry) GetType() SetEffectsEntryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SetEffectsEntry) GetTypeOk() (*SetEffectsEntryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SetEffectsEntry) SetType(v SetEffectsEntryType)`

SetType sets Type field to given value.

### HasType

`func (o *SetEffectsEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIgnoreIntMin

`func (o *SetEffectsEntry) GetIgnoreIntMin() bool`

GetIgnoreIntMin returns the IgnoreIntMin field if non-nil, zero value otherwise.

### GetIgnoreIntMinOk

`func (o *SetEffectsEntry) GetIgnoreIntMinOk() (*bool, bool)`

GetIgnoreIntMinOk returns a tuple with the IgnoreIntMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreIntMin

`func (o *SetEffectsEntry) SetIgnoreIntMin(v bool)`

SetIgnoreIntMin sets IgnoreIntMin field to given value.

### HasIgnoreIntMin

`func (o *SetEffectsEntry) HasIgnoreIntMin() bool`

HasIgnoreIntMin returns a boolean if a field has been set.

### GetIgnoreIntMax

`func (o *SetEffectsEntry) GetIgnoreIntMax() bool`

GetIgnoreIntMax returns the IgnoreIntMax field if non-nil, zero value otherwise.

### GetIgnoreIntMaxOk

`func (o *SetEffectsEntry) GetIgnoreIntMaxOk() (*bool, bool)`

GetIgnoreIntMaxOk returns a tuple with the IgnoreIntMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreIntMax

`func (o *SetEffectsEntry) SetIgnoreIntMax(v bool)`

SetIgnoreIntMax sets IgnoreIntMax field to given value.

### HasIgnoreIntMax

`func (o *SetEffectsEntry) HasIgnoreIntMax() bool`

HasIgnoreIntMax returns a boolean if a field has been set.

### GetFormatted

`func (o *SetEffectsEntry) GetFormatted() string`

GetFormatted returns the Formatted field if non-nil, zero value otherwise.

### GetFormattedOk

`func (o *SetEffectsEntry) GetFormattedOk() (*string, bool)`

GetFormattedOk returns a tuple with the Formatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatted

`func (o *SetEffectsEntry) SetFormatted(v string)`

SetFormatted sets Formatted field to given value.

### HasFormatted

`func (o *SetEffectsEntry) HasFormatted() bool`

HasFormatted returns a boolean if a field has been set.

### GetItemCombination

`func (o *SetEffectsEntry) GetItemCombination() int32`

GetItemCombination returns the ItemCombination field if non-nil, zero value otherwise.

### GetItemCombinationOk

`func (o *SetEffectsEntry) GetItemCombinationOk() (*int32, bool)`

GetItemCombinationOk returns a tuple with the ItemCombination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCombination

`func (o *SetEffectsEntry) SetItemCombination(v int32)`

SetItemCombination sets ItemCombination field to given value.

### HasItemCombination

`func (o *SetEffectsEntry) HasItemCombination() bool`

HasItemCombination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


