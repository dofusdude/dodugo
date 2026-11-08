# MappedEntityCollectionSuperAreasInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | **int32** |  | 
**Name** | **string** |  | 
**ImageUrls** | Pointer to [**Images**](Images.md) |  | [optional] 
**SuperArea** | Pointer to [**CombatReference**](CombatReference.md) |  | [optional] 
**Area** | Pointer to [**CombatReference**](CombatReference.md) |  | [optional] 
**Subareas** | Pointer to [**[]CombatReference**](CombatReference.md) |  | [optional] 
**Monsters** | Pointer to [**[]CombatReference**](CombatReference.md) |  | [optional] 

## Methods

### NewMappedEntityCollectionSuperAreasInner

`func NewMappedEntityCollectionSuperAreasInner(ankamaId int32, name string, ) *MappedEntityCollectionSuperAreasInner`

NewMappedEntityCollectionSuperAreasInner instantiates a new MappedEntityCollectionSuperAreasInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMappedEntityCollectionSuperAreasInnerWithDefaults

`func NewMappedEntityCollectionSuperAreasInnerWithDefaults() *MappedEntityCollectionSuperAreasInner`

NewMappedEntityCollectionSuperAreasInnerWithDefaults instantiates a new MappedEntityCollectionSuperAreasInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *MappedEntityCollectionSuperAreasInner) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *MappedEntityCollectionSuperAreasInner) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *MappedEntityCollectionSuperAreasInner) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.


### GetName

`func (o *MappedEntityCollectionSuperAreasInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MappedEntityCollectionSuperAreasInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MappedEntityCollectionSuperAreasInner) SetName(v string)`

SetName sets Name field to given value.


### GetImageUrls

`func (o *MappedEntityCollectionSuperAreasInner) GetImageUrls() Images`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *MappedEntityCollectionSuperAreasInner) GetImageUrlsOk() (*Images, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *MappedEntityCollectionSuperAreasInner) SetImageUrls(v Images)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *MappedEntityCollectionSuperAreasInner) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.

### GetSuperArea

`func (o *MappedEntityCollectionSuperAreasInner) GetSuperArea() CombatReference`

GetSuperArea returns the SuperArea field if non-nil, zero value otherwise.

### GetSuperAreaOk

`func (o *MappedEntityCollectionSuperAreasInner) GetSuperAreaOk() (*CombatReference, bool)`

GetSuperAreaOk returns a tuple with the SuperArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperArea

`func (o *MappedEntityCollectionSuperAreasInner) SetSuperArea(v CombatReference)`

SetSuperArea sets SuperArea field to given value.

### HasSuperArea

`func (o *MappedEntityCollectionSuperAreasInner) HasSuperArea() bool`

HasSuperArea returns a boolean if a field has been set.

### GetArea

`func (o *MappedEntityCollectionSuperAreasInner) GetArea() CombatReference`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *MappedEntityCollectionSuperAreasInner) GetAreaOk() (*CombatReference, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *MappedEntityCollectionSuperAreasInner) SetArea(v CombatReference)`

SetArea sets Area field to given value.

### HasArea

`func (o *MappedEntityCollectionSuperAreasInner) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetSubareas

`func (o *MappedEntityCollectionSuperAreasInner) GetSubareas() []CombatReference`

GetSubareas returns the Subareas field if non-nil, zero value otherwise.

### GetSubareasOk

`func (o *MappedEntityCollectionSuperAreasInner) GetSubareasOk() (*[]CombatReference, bool)`

GetSubareasOk returns a tuple with the Subareas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubareas

`func (o *MappedEntityCollectionSuperAreasInner) SetSubareas(v []CombatReference)`

SetSubareas sets Subareas field to given value.

### HasSubareas

`func (o *MappedEntityCollectionSuperAreasInner) HasSubareas() bool`

HasSubareas returns a boolean if a field has been set.

### GetMonsters

`func (o *MappedEntityCollectionSuperAreasInner) GetMonsters() []CombatReference`

GetMonsters returns the Monsters field if non-nil, zero value otherwise.

### GetMonstersOk

`func (o *MappedEntityCollectionSuperAreasInner) GetMonstersOk() (*[]CombatReference, bool)`

GetMonstersOk returns a tuple with the Monsters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonsters

`func (o *MappedEntityCollectionSuperAreasInner) SetMonsters(v []CombatReference)`

SetMonsters sets Monsters field to given value.

### HasMonsters

`func (o *MappedEntityCollectionSuperAreasInner) HasMonsters() bool`

HasMonsters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


