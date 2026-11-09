# WorldArea

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | **int32** |  | 
**Name** | **string** |  | 
**WorldMapId** | Pointer to **int32** | World-map context supplied on super-area records. | [optional] 
**Areas** | Pointer to [**[]CombatReference**](CombatReference.md) | Child area references supplied on super-area records. | [optional] 
**Level** | Pointer to **int32** | Suggested geographic difficulty level supplied on subarea records. Not a monster grade or the level of every monster there. | [optional] 
**DungeonId** | Pointer to **int32** | Dungeon association when supplied by the subarea data. | [optional] 
**Maps** | Pointer to [**[]WorldMapReference**](WorldMapReference.md) | Map membership supplied on subarea records. Does not imply a spawn on every map. | [optional] 
**SuperArea** | Pointer to [**CombatReference**](CombatReference.md) |  | [optional] 
**Area** | Pointer to [**CombatReference**](CombatReference.md) |  | [optional] 
**Subareas** | Pointer to [**[]CombatReference**](CombatReference.md) |  | [optional] 
**Monsters** | Pointer to [**[]CombatReference**](CombatReference.md) |  | [optional] 

## Methods

### NewWorldArea

`func NewWorldArea(ankamaId int32, name string, ) *WorldArea`

NewWorldArea instantiates a new WorldArea object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorldAreaWithDefaults

`func NewWorldAreaWithDefaults() *WorldArea`

NewWorldAreaWithDefaults instantiates a new WorldArea object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *WorldArea) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *WorldArea) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *WorldArea) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.


### GetName

`func (o *WorldArea) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorldArea) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorldArea) SetName(v string)`

SetName sets Name field to given value.


### GetWorldMapId

`func (o *WorldArea) GetWorldMapId() int32`

GetWorldMapId returns the WorldMapId field if non-nil, zero value otherwise.

### GetWorldMapIdOk

`func (o *WorldArea) GetWorldMapIdOk() (*int32, bool)`

GetWorldMapIdOk returns a tuple with the WorldMapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorldMapId

`func (o *WorldArea) SetWorldMapId(v int32)`

SetWorldMapId sets WorldMapId field to given value.

### HasWorldMapId

`func (o *WorldArea) HasWorldMapId() bool`

HasWorldMapId returns a boolean if a field has been set.

### GetAreas

`func (o *WorldArea) GetAreas() []CombatReference`

GetAreas returns the Areas field if non-nil, zero value otherwise.

### GetAreasOk

`func (o *WorldArea) GetAreasOk() (*[]CombatReference, bool)`

GetAreasOk returns a tuple with the Areas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreas

`func (o *WorldArea) SetAreas(v []CombatReference)`

SetAreas sets Areas field to given value.

### HasAreas

`func (o *WorldArea) HasAreas() bool`

HasAreas returns a boolean if a field has been set.

### GetLevel

`func (o *WorldArea) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *WorldArea) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *WorldArea) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *WorldArea) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetDungeonId

`func (o *WorldArea) GetDungeonId() int32`

GetDungeonId returns the DungeonId field if non-nil, zero value otherwise.

### GetDungeonIdOk

`func (o *WorldArea) GetDungeonIdOk() (*int32, bool)`

GetDungeonIdOk returns a tuple with the DungeonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDungeonId

`func (o *WorldArea) SetDungeonId(v int32)`

SetDungeonId sets DungeonId field to given value.

### HasDungeonId

`func (o *WorldArea) HasDungeonId() bool`

HasDungeonId returns a boolean if a field has been set.

### GetMaps

`func (o *WorldArea) GetMaps() []WorldMapReference`

GetMaps returns the Maps field if non-nil, zero value otherwise.

### GetMapsOk

`func (o *WorldArea) GetMapsOk() (*[]WorldMapReference, bool)`

GetMapsOk returns a tuple with the Maps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaps

`func (o *WorldArea) SetMaps(v []WorldMapReference)`

SetMaps sets Maps field to given value.

### HasMaps

`func (o *WorldArea) HasMaps() bool`

HasMaps returns a boolean if a field has been set.

### GetSuperArea

`func (o *WorldArea) GetSuperArea() CombatReference`

GetSuperArea returns the SuperArea field if non-nil, zero value otherwise.

### GetSuperAreaOk

`func (o *WorldArea) GetSuperAreaOk() (*CombatReference, bool)`

GetSuperAreaOk returns a tuple with the SuperArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperArea

`func (o *WorldArea) SetSuperArea(v CombatReference)`

SetSuperArea sets SuperArea field to given value.

### HasSuperArea

`func (o *WorldArea) HasSuperArea() bool`

HasSuperArea returns a boolean if a field has been set.

### GetArea

`func (o *WorldArea) GetArea() CombatReference`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *WorldArea) GetAreaOk() (*CombatReference, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *WorldArea) SetArea(v CombatReference)`

SetArea sets Area field to given value.

### HasArea

`func (o *WorldArea) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetSubareas

`func (o *WorldArea) GetSubareas() []CombatReference`

GetSubareas returns the Subareas field if non-nil, zero value otherwise.

### GetSubareasOk

`func (o *WorldArea) GetSubareasOk() (*[]CombatReference, bool)`

GetSubareasOk returns a tuple with the Subareas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubareas

`func (o *WorldArea) SetSubareas(v []CombatReference)`

SetSubareas sets Subareas field to given value.

### HasSubareas

`func (o *WorldArea) HasSubareas() bool`

HasSubareas returns a boolean if a field has been set.

### GetMonsters

`func (o *WorldArea) GetMonsters() []CombatReference`

GetMonsters returns the Monsters field if non-nil, zero value otherwise.

### GetMonstersOk

`func (o *WorldArea) GetMonstersOk() (*[]CombatReference, bool)`

GetMonstersOk returns a tuple with the Monsters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonsters

`func (o *WorldArea) SetMonsters(v []CombatReference)`

SetMonsters sets Monsters field to given value.

### HasMonsters

`func (o *WorldArea) HasMonsters() bool`

HasMonsters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


