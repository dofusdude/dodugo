# EffectsEntryType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IsMeta** | Pointer to **bool** | true if a type is generated from the Api instead of Ankama. In that case, always prefer showing the templated string and omit everything else. The \&quot;name\&quot; field will have an english description of the meta type. An example for such effects are class sets effects. | [optional] 

## Methods

### NewEffectsEntryType

`func NewEffectsEntryType() *EffectsEntryType`

NewEffectsEntryType instantiates a new EffectsEntryType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEffectsEntryTypeWithDefaults

`func NewEffectsEntryTypeWithDefaults() *EffectsEntryType`

NewEffectsEntryTypeWithDefaults instantiates a new EffectsEntryType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EffectsEntryType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EffectsEntryType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EffectsEntryType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EffectsEntryType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *EffectsEntryType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EffectsEntryType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EffectsEntryType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EffectsEntryType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsMeta

`func (o *EffectsEntryType) GetIsMeta() bool`

GetIsMeta returns the IsMeta field if non-nil, zero value otherwise.

### GetIsMetaOk

`func (o *EffectsEntryType) GetIsMetaOk() (*bool, bool)`

GetIsMetaOk returns a tuple with the IsMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMeta

`func (o *EffectsEntryType) SetIsMeta(v bool)`

SetIsMeta sets IsMeta field to given value.

### HasIsMeta

`func (o *EffectsEntryType) HasIsMeta() bool`

HasIsMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


