# EquipmentSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**EquipmentIds** | Pointer to **[]int32** |  | [optional] 
**Effects** | Pointer to [**map[string][]Effect**](array.md) |  | [optional] 
**HighestEquipmentLevel** | Pointer to **int32** |  | [optional] 
**ContainsCosmetics** | Pointer to **bool** |  | [optional] 
**ContainsCosmeticsOnly** | Pointer to **bool** |  | [optional] 

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

### GetEffects

`func (o *EquipmentSet) GetEffects() map[string][]Effect`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *EquipmentSet) GetEffectsOk() (*map[string][]Effect, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *EquipmentSet) SetEffects(v map[string][]Effect)`

SetEffects sets Effects field to given value.

### HasEffects

`func (o *EquipmentSet) HasEffects() bool`

HasEffects returns a boolean if a field has been set.

### GetHighestEquipmentLevel

`func (o *EquipmentSet) GetHighestEquipmentLevel() int32`

GetHighestEquipmentLevel returns the HighestEquipmentLevel field if non-nil, zero value otherwise.

### GetHighestEquipmentLevelOk

`func (o *EquipmentSet) GetHighestEquipmentLevelOk() (*int32, bool)`

GetHighestEquipmentLevelOk returns a tuple with the HighestEquipmentLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestEquipmentLevel

`func (o *EquipmentSet) SetHighestEquipmentLevel(v int32)`

SetHighestEquipmentLevel sets HighestEquipmentLevel field to given value.

### HasHighestEquipmentLevel

`func (o *EquipmentSet) HasHighestEquipmentLevel() bool`

HasHighestEquipmentLevel returns a boolean if a field has been set.

### GetContainsCosmetics

`func (o *EquipmentSet) GetContainsCosmetics() bool`

GetContainsCosmetics returns the ContainsCosmetics field if non-nil, zero value otherwise.

### GetContainsCosmeticsOk

`func (o *EquipmentSet) GetContainsCosmeticsOk() (*bool, bool)`

GetContainsCosmeticsOk returns a tuple with the ContainsCosmetics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsCosmetics

`func (o *EquipmentSet) SetContainsCosmetics(v bool)`

SetContainsCosmetics sets ContainsCosmetics field to given value.

### HasContainsCosmetics

`func (o *EquipmentSet) HasContainsCosmetics() bool`

HasContainsCosmetics returns a boolean if a field has been set.

### GetContainsCosmeticsOnly

`func (o *EquipmentSet) GetContainsCosmeticsOnly() bool`

GetContainsCosmeticsOnly returns the ContainsCosmeticsOnly field if non-nil, zero value otherwise.

### GetContainsCosmeticsOnlyOk

`func (o *EquipmentSet) GetContainsCosmeticsOnlyOk() (*bool, bool)`

GetContainsCosmeticsOnlyOk returns a tuple with the ContainsCosmeticsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsCosmeticsOnly

`func (o *EquipmentSet) SetContainsCosmeticsOnly(v bool)`

SetContainsCosmeticsOnly sets ContainsCosmeticsOnly field to given value.

### HasContainsCosmeticsOnly

`func (o *EquipmentSet) HasContainsCosmeticsOnly() bool`

HasContainsCosmeticsOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


