# EffectsEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntMinimum** | Pointer to **int32** | minimum int value, can be a single if ignore_int_max and no ignore_int_min | [optional] 
**IntMaximum** | Pointer to **int32** | maximum int value, if not ignore_int_max and not ignore_int_min, the effect has a range value | [optional] 
**Type** | Pointer to [**EffectsEntryType**](EffectsEntryType.md) |  | [optional] 
**IgnoreIntMin** | Pointer to **bool** | ignore the int min field because the actual value is a string. For readability the templated field is the only important field in this case.  | [optional] 
**IgnoreIntMax** | Pointer to **bool** | ignore the int max field, if ignore_int_min is true, int min is a single value | [optional] 
**Formatted** | Pointer to **string** | all fields from above encoded in a single string | [optional] 

## Methods

### NewEffectsEntry

`func NewEffectsEntry() *EffectsEntry`

NewEffectsEntry instantiates a new EffectsEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEffectsEntryWithDefaults

`func NewEffectsEntryWithDefaults() *EffectsEntry`

NewEffectsEntryWithDefaults instantiates a new EffectsEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntMinimum

`func (o *EffectsEntry) GetIntMinimum() int32`

GetIntMinimum returns the IntMinimum field if non-nil, zero value otherwise.

### GetIntMinimumOk

`func (o *EffectsEntry) GetIntMinimumOk() (*int32, bool)`

GetIntMinimumOk returns a tuple with the IntMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntMinimum

`func (o *EffectsEntry) SetIntMinimum(v int32)`

SetIntMinimum sets IntMinimum field to given value.

### HasIntMinimum

`func (o *EffectsEntry) HasIntMinimum() bool`

HasIntMinimum returns a boolean if a field has been set.

### GetIntMaximum

`func (o *EffectsEntry) GetIntMaximum() int32`

GetIntMaximum returns the IntMaximum field if non-nil, zero value otherwise.

### GetIntMaximumOk

`func (o *EffectsEntry) GetIntMaximumOk() (*int32, bool)`

GetIntMaximumOk returns a tuple with the IntMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntMaximum

`func (o *EffectsEntry) SetIntMaximum(v int32)`

SetIntMaximum sets IntMaximum field to given value.

### HasIntMaximum

`func (o *EffectsEntry) HasIntMaximum() bool`

HasIntMaximum returns a boolean if a field has been set.

### GetType

`func (o *EffectsEntry) GetType() EffectsEntryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EffectsEntry) GetTypeOk() (*EffectsEntryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EffectsEntry) SetType(v EffectsEntryType)`

SetType sets Type field to given value.

### HasType

`func (o *EffectsEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIgnoreIntMin

`func (o *EffectsEntry) GetIgnoreIntMin() bool`

GetIgnoreIntMin returns the IgnoreIntMin field if non-nil, zero value otherwise.

### GetIgnoreIntMinOk

`func (o *EffectsEntry) GetIgnoreIntMinOk() (*bool, bool)`

GetIgnoreIntMinOk returns a tuple with the IgnoreIntMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreIntMin

`func (o *EffectsEntry) SetIgnoreIntMin(v bool)`

SetIgnoreIntMin sets IgnoreIntMin field to given value.

### HasIgnoreIntMin

`func (o *EffectsEntry) HasIgnoreIntMin() bool`

HasIgnoreIntMin returns a boolean if a field has been set.

### GetIgnoreIntMax

`func (o *EffectsEntry) GetIgnoreIntMax() bool`

GetIgnoreIntMax returns the IgnoreIntMax field if non-nil, zero value otherwise.

### GetIgnoreIntMaxOk

`func (o *EffectsEntry) GetIgnoreIntMaxOk() (*bool, bool)`

GetIgnoreIntMaxOk returns a tuple with the IgnoreIntMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreIntMax

`func (o *EffectsEntry) SetIgnoreIntMax(v bool)`

SetIgnoreIntMax sets IgnoreIntMax field to given value.

### HasIgnoreIntMax

`func (o *EffectsEntry) HasIgnoreIntMax() bool`

HasIgnoreIntMax returns a boolean if a field has been set.

### GetFormatted

`func (o *EffectsEntry) GetFormatted() string`

GetFormatted returns the Formatted field if non-nil, zero value otherwise.

### GetFormattedOk

`func (o *EffectsEntry) GetFormattedOk() (*string, bool)`

GetFormattedOk returns a tuple with the Formatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatted

`func (o *EffectsEntry) SetFormatted(v string)`

SetFormatted sets Formatted field to given value.

### HasFormatted

`func (o *EffectsEntry) HasFormatted() bool`

HasFormatted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


