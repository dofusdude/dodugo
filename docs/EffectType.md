# EffectType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** | Affects target or source actively. | [optional] 
**IsMeta** | Pointer to **bool** | true if a type is generated from the Api instead of Ankama. In that case, always prefer showing the templated string and omit everything else. The \&quot;name\&quot; field will have an english description of the meta type. An example for such effects are class sets effects. | [optional] 

## Methods

### NewEffectType

`func NewEffectType() *EffectType`

NewEffectType instantiates a new EffectType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEffectTypeWithDefaults

`func NewEffectTypeWithDefaults() *EffectType`

NewEffectTypeWithDefaults instantiates a new EffectType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EffectType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EffectType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EffectType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EffectType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EffectType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EffectType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EffectType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EffectType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsActive

`func (o *EffectType) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *EffectType) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *EffectType) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *EffectType) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsMeta

`func (o *EffectType) GetIsMeta() bool`

GetIsMeta returns the IsMeta field if non-nil, zero value otherwise.

### GetIsMetaOk

`func (o *EffectType) GetIsMetaOk() (*bool, bool)`

GetIsMetaOk returns a tuple with the IsMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMeta

`func (o *EffectType) SetIsMeta(v bool)`

SetIsMeta sets IsMeta field to given value.

### HasIsMeta

`func (o *EffectType) HasIsMeta() bool`

HasIsMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


