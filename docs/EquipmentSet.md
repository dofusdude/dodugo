# EquipmentSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**EquipmentIds** | Pointer to **[]int32** |  | [optional] 
**HasEffects** | Pointer to **bool** |  | [optional] 
**Effects** | Pointer to [**[][]EffectsEntry**]([]EffectsEntry.md) |  | [optional] 
**Level** | Pointer to **int32** |  | [optional] 

## Methods

### NewEquipmentSet

`func NewEquipmentSet() *EquipmentSet`

NewEquipmentSet instantiates a new EquipmentSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentSetWithDefaults

`func NewEquipmentSetWithDefaults() *EquipmentSet`

NewEquipmentSetWithDefaults instantiates a new EquipmentSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *EquipmentSet) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *EquipmentSet) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *EquipmentSet) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.

### HasAnkamaId

`func (o *EquipmentSet) HasAnkamaId() bool`

HasAnkamaId returns a boolean if a field has been set.

### GetName

`func (o *EquipmentSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EquipmentSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EquipmentSet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EquipmentSet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEquipmentIds

`func (o *EquipmentSet) GetEquipmentIds() []int32`

GetEquipmentIds returns the EquipmentIds field if non-nil, zero value otherwise.

### GetEquipmentIdsOk

`func (o *EquipmentSet) GetEquipmentIdsOk() (*[]int32, bool)`

GetEquipmentIdsOk returns a tuple with the EquipmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentIds

`func (o *EquipmentSet) SetEquipmentIds(v []int32)`

SetEquipmentIds sets EquipmentIds field to given value.

### HasEquipmentIds

`func (o *EquipmentSet) HasEquipmentIds() bool`

HasEquipmentIds returns a boolean if a field has been set.

### GetHasEffects

`func (o *EquipmentSet) GetHasEffects() bool`

GetHasEffects returns the HasEffects field if non-nil, zero value otherwise.

### GetHasEffectsOk

`func (o *EquipmentSet) GetHasEffectsOk() (*bool, bool)`

GetHasEffectsOk returns a tuple with the HasEffects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEffects

`func (o *EquipmentSet) SetHasEffects(v bool)`

SetHasEffects sets HasEffects field to given value.

### HasHasEffects

`func (o *EquipmentSet) HasHasEffects() bool`

HasHasEffects returns a boolean if a field has been set.

### GetEffects

`func (o *EquipmentSet) GetEffects() [][]EffectsEntry`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *EquipmentSet) GetEffectsOk() (*[][]EffectsEntry, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *EquipmentSet) SetEffects(v [][]EffectsEntry)`

SetEffects sets Effects field to given value.

### HasEffects

`func (o *EquipmentSet) HasEffects() bool`

HasEffects returns a boolean if a field has been set.

### SetEffectsNil

`func (o *EquipmentSet) SetEffectsNil(b bool)`

 SetEffectsNil sets the value for Effects to be an explicit nil

### UnsetEffects
`func (o *EquipmentSet) UnsetEffects()`

UnsetEffects ensures that no value is present for Effects, not even an explicit nil
### GetLevel

`func (o *EquipmentSet) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *EquipmentSet) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *EquipmentSet) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *EquipmentSet) HasLevel() bool`

HasLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


