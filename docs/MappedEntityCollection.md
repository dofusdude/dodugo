# MappedEntityCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classes** | Pointer to [**[]MappedEntityCollectionClassesInner**](MappedEntityCollectionClassesInner.md) |  | [optional] 
**Spells** | Pointer to [**[]MappedEntityCollectionSpellsInner**](MappedEntityCollectionSpellsInner.md) |  | [optional] 
**Monsters** | Pointer to [**[]MappedEntityCollectionMonstersInner**](MappedEntityCollectionMonstersInner.md) |  | [optional] 
**SuperAreas** | Pointer to [**[]MappedEntityCollectionSuperAreasInner**](MappedEntityCollectionSuperAreasInner.md) |  | [optional] 
**Areas** | Pointer to [**[]MappedEntityCollectionSuperAreasInner**](MappedEntityCollectionSuperAreasInner.md) |  | [optional] 
**Subareas** | Pointer to [**[]MappedEntityCollectionSuperAreasInner**](MappedEntityCollectionSuperAreasInner.md) |  | [optional] 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewMappedEntityCollection

`func NewMappedEntityCollection() *MappedEntityCollection`

NewMappedEntityCollection instantiates a new MappedEntityCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMappedEntityCollectionWithDefaults

`func NewMappedEntityCollectionWithDefaults() *MappedEntityCollection`

NewMappedEntityCollectionWithDefaults instantiates a new MappedEntityCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClasses

`func (o *MappedEntityCollection) GetClasses() []MappedEntityCollectionClassesInner`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *MappedEntityCollection) GetClassesOk() (*[]MappedEntityCollectionClassesInner, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *MappedEntityCollection) SetClasses(v []MappedEntityCollectionClassesInner)`

SetClasses sets Classes field to given value.

### HasClasses

`func (o *MappedEntityCollection) HasClasses() bool`

HasClasses returns a boolean if a field has been set.

### GetSpells

`func (o *MappedEntityCollection) GetSpells() []MappedEntityCollectionSpellsInner`

GetSpells returns the Spells field if non-nil, zero value otherwise.

### GetSpellsOk

`func (o *MappedEntityCollection) GetSpellsOk() (*[]MappedEntityCollectionSpellsInner, bool)`

GetSpellsOk returns a tuple with the Spells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpells

`func (o *MappedEntityCollection) SetSpells(v []MappedEntityCollectionSpellsInner)`

SetSpells sets Spells field to given value.

### HasSpells

`func (o *MappedEntityCollection) HasSpells() bool`

HasSpells returns a boolean if a field has been set.

### GetMonsters

`func (o *MappedEntityCollection) GetMonsters() []MappedEntityCollectionMonstersInner`

GetMonsters returns the Monsters field if non-nil, zero value otherwise.

### GetMonstersOk

`func (o *MappedEntityCollection) GetMonstersOk() (*[]MappedEntityCollectionMonstersInner, bool)`

GetMonstersOk returns a tuple with the Monsters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonsters

`func (o *MappedEntityCollection) SetMonsters(v []MappedEntityCollectionMonstersInner)`

SetMonsters sets Monsters field to given value.

### HasMonsters

`func (o *MappedEntityCollection) HasMonsters() bool`

HasMonsters returns a boolean if a field has been set.

### GetSuperAreas

`func (o *MappedEntityCollection) GetSuperAreas() []MappedEntityCollectionSuperAreasInner`

GetSuperAreas returns the SuperAreas field if non-nil, zero value otherwise.

### GetSuperAreasOk

`func (o *MappedEntityCollection) GetSuperAreasOk() (*[]MappedEntityCollectionSuperAreasInner, bool)`

GetSuperAreasOk returns a tuple with the SuperAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperAreas

`func (o *MappedEntityCollection) SetSuperAreas(v []MappedEntityCollectionSuperAreasInner)`

SetSuperAreas sets SuperAreas field to given value.

### HasSuperAreas

`func (o *MappedEntityCollection) HasSuperAreas() bool`

HasSuperAreas returns a boolean if a field has been set.

### GetAreas

`func (o *MappedEntityCollection) GetAreas() []MappedEntityCollectionSuperAreasInner`

GetAreas returns the Areas field if non-nil, zero value otherwise.

### GetAreasOk

`func (o *MappedEntityCollection) GetAreasOk() (*[]MappedEntityCollectionSuperAreasInner, bool)`

GetAreasOk returns a tuple with the Areas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreas

`func (o *MappedEntityCollection) SetAreas(v []MappedEntityCollectionSuperAreasInner)`

SetAreas sets Areas field to given value.

### HasAreas

`func (o *MappedEntityCollection) HasAreas() bool`

HasAreas returns a boolean if a field has been set.

### GetSubareas

`func (o *MappedEntityCollection) GetSubareas() []MappedEntityCollectionSuperAreasInner`

GetSubareas returns the Subareas field if non-nil, zero value otherwise.

### GetSubareasOk

`func (o *MappedEntityCollection) GetSubareasOk() (*[]MappedEntityCollectionSuperAreasInner, bool)`

GetSubareasOk returns a tuple with the Subareas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubareas

`func (o *MappedEntityCollection) SetSubareas(v []MappedEntityCollectionSuperAreasInner)`

SetSubareas sets Subareas field to given value.

### HasSubareas

`func (o *MappedEntityCollection) HasSubareas() bool`

HasSubareas returns a boolean if a field has been set.

### GetLinks

`func (o *MappedEntityCollection) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MappedEntityCollection) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MappedEntityCollection) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MappedEntityCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


