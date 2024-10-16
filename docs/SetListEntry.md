# SetListEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Items** | Pointer to **int32** | amount | [optional] 
**Level** | Pointer to **int32** |  | [optional] 
**Effects** | Pointer to [**[][]SetEffectsEntry**]([]SetEffectsEntry.md) |  | [optional] 
**EquipmentIds** | Pointer to **[]int32** |  | [optional] 
**IsCosmetic** | Pointer to **bool** |  | [optional] 

## Methods

### NewSetListEntry

`func NewSetListEntry() *SetListEntry`

NewSetListEntry instantiates a new SetListEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetListEntryWithDefaults

`func NewSetListEntryWithDefaults() *SetListEntry`

NewSetListEntryWithDefaults instantiates a new SetListEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *SetListEntry) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *SetListEntry) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *SetListEntry) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.

### HasAnkamaId

`func (o *SetListEntry) HasAnkamaId() bool`

HasAnkamaId returns a boolean if a field has been set.

### GetName

`func (o *SetListEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SetListEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SetListEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SetListEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetItems

`func (o *SetListEntry) GetItems() int32`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SetListEntry) GetItemsOk() (*int32, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SetListEntry) SetItems(v int32)`

SetItems sets Items field to given value.

### HasItems

`func (o *SetListEntry) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLevel

`func (o *SetListEntry) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *SetListEntry) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *SetListEntry) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *SetListEntry) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetEffects

`func (o *SetListEntry) GetEffects() [][]SetEffectsEntry`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *SetListEntry) GetEffectsOk() (*[][]SetEffectsEntry, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *SetListEntry) SetEffects(v [][]SetEffectsEntry)`

SetEffects sets Effects field to given value.

### HasEffects

`func (o *SetListEntry) HasEffects() bool`

HasEffects returns a boolean if a field has been set.

### SetEffectsNil

`func (o *SetListEntry) SetEffectsNil(b bool)`

 SetEffectsNil sets the value for Effects to be an explicit nil

### UnsetEffects
`func (o *SetListEntry) UnsetEffects()`

UnsetEffects ensures that no value is present for Effects, not even an explicit nil
### GetEquipmentIds

`func (o *SetListEntry) GetEquipmentIds() []int32`

GetEquipmentIds returns the EquipmentIds field if non-nil, zero value otherwise.

### GetEquipmentIdsOk

`func (o *SetListEntry) GetEquipmentIdsOk() (*[]int32, bool)`

GetEquipmentIdsOk returns a tuple with the EquipmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentIds

`func (o *SetListEntry) SetEquipmentIds(v []int32)`

SetEquipmentIds sets EquipmentIds field to given value.

### HasEquipmentIds

`func (o *SetListEntry) HasEquipmentIds() bool`

HasEquipmentIds returns a boolean if a field has been set.

### SetEquipmentIdsNil

`func (o *SetListEntry) SetEquipmentIdsNil(b bool)`

 SetEquipmentIdsNil sets the value for EquipmentIds to be an explicit nil

### UnsetEquipmentIds
`func (o *SetListEntry) UnsetEquipmentIds()`

UnsetEquipmentIds ensures that no value is present for EquipmentIds, not even an explicit nil
### GetIsCosmetic

`func (o *SetListEntry) GetIsCosmetic() bool`

GetIsCosmetic returns the IsCosmetic field if non-nil, zero value otherwise.

### GetIsCosmeticOk

`func (o *SetListEntry) GetIsCosmeticOk() (*bool, bool)`

GetIsCosmeticOk returns a tuple with the IsCosmetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCosmetic

`func (o *SetListEntry) SetIsCosmetic(v bool)`

SetIsCosmetic sets IsCosmetic field to given value.

### HasIsCosmetic

`func (o *SetListEntry) HasIsCosmetic() bool`

HasIsCosmetic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


