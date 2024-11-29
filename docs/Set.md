# Set

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**EquipmentIds** | Pointer to **[]int32** |  | [optional] 
**Effects** | Pointer to [**map[string][]Effect**](array.md) |  | [optional] 
**HighestEquipmentLevel** | Pointer to **int32** |  | [optional] 
**IsCosmetic** | Pointer to **bool** |  | [optional] 

## Methods

### NewSet

`func NewSet() *Set`

NewSet instantiates a new Set object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetWithDefaults

`func NewSetWithDefaults() *Set`

NewSetWithDefaults instantiates a new Set object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *Set) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *Set) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *Set) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.

### HasAnkamaId

`func (o *Set) HasAnkamaId() bool`

HasAnkamaId returns a boolean if a field has been set.

### GetName

`func (o *Set) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Set) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Set) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Set) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEquipmentIds

`func (o *Set) GetEquipmentIds() []int32`

GetEquipmentIds returns the EquipmentIds field if non-nil, zero value otherwise.

### GetEquipmentIdsOk

`func (o *Set) GetEquipmentIdsOk() (*[]int32, bool)`

GetEquipmentIdsOk returns a tuple with the EquipmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentIds

`func (o *Set) SetEquipmentIds(v []int32)`

SetEquipmentIds sets EquipmentIds field to given value.

### HasEquipmentIds

`func (o *Set) HasEquipmentIds() bool`

HasEquipmentIds returns a boolean if a field has been set.

### GetEffects

`func (o *Set) GetEffects() map[string][]Effect`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *Set) GetEffectsOk() (*map[string][]Effect, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *Set) SetEffects(v map[string][]Effect)`

SetEffects sets Effects field to given value.

### HasEffects

`func (o *Set) HasEffects() bool`

HasEffects returns a boolean if a field has been set.

### GetHighestEquipmentLevel

`func (o *Set) GetHighestEquipmentLevel() int32`

GetHighestEquipmentLevel returns the HighestEquipmentLevel field if non-nil, zero value otherwise.

### GetHighestEquipmentLevelOk

`func (o *Set) GetHighestEquipmentLevelOk() (*int32, bool)`

GetHighestEquipmentLevelOk returns a tuple with the HighestEquipmentLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestEquipmentLevel

`func (o *Set) SetHighestEquipmentLevel(v int32)`

SetHighestEquipmentLevel sets HighestEquipmentLevel field to given value.

### HasHighestEquipmentLevel

`func (o *Set) HasHighestEquipmentLevel() bool`

HasHighestEquipmentLevel returns a boolean if a field has been set.

### GetIsCosmetic

`func (o *Set) GetIsCosmetic() bool`

GetIsCosmetic returns the IsCosmetic field if non-nil, zero value otherwise.

### GetIsCosmeticOk

`func (o *Set) GetIsCosmeticOk() (*bool, bool)`

GetIsCosmeticOk returns a tuple with the IsCosmetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCosmetic

`func (o *Set) SetIsCosmetic(v bool)`

SetIsCosmetic sets IsCosmetic field to given value.

### HasIsCosmetic

`func (o *Set) HasIsCosmetic() bool`

HasIsCosmetic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


