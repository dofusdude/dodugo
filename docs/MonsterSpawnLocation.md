# MonsterSpawnLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subarea** | [**CombatReference**](CombatReference.md) |  | 
**Area** | [**CombatReference**](CombatReference.md) |  | 
**SuperArea** | [**CombatReference**](CombatReference.md) |  | 
**SubareaMaps** | [**[]MonsterSpawnLocationSubareaMapsInner**](MonsterSpawnLocationSubareaMapsInner.md) | Maps belonging to the referenced subarea. Map membership alone is not proof that the monster spawns on a particular map. | 
**Preferred** | **bool** | Whether the source marks this as a preferred subarea. A preference flag, not a spawn percentage or guarantee of presence. | 

## Methods

### NewMonsterSpawnLocation

`func NewMonsterSpawnLocation(subarea CombatReference, area CombatReference, superArea CombatReference, subareaMaps []MonsterSpawnLocationSubareaMapsInner, preferred bool, ) *MonsterSpawnLocation`

NewMonsterSpawnLocation instantiates a new MonsterSpawnLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonsterSpawnLocationWithDefaults

`func NewMonsterSpawnLocationWithDefaults() *MonsterSpawnLocation`

NewMonsterSpawnLocationWithDefaults instantiates a new MonsterSpawnLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubarea

`func (o *MonsterSpawnLocation) GetSubarea() CombatReference`

GetSubarea returns the Subarea field if non-nil, zero value otherwise.

### GetSubareaOk

`func (o *MonsterSpawnLocation) GetSubareaOk() (*CombatReference, bool)`

GetSubareaOk returns a tuple with the Subarea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubarea

`func (o *MonsterSpawnLocation) SetSubarea(v CombatReference)`

SetSubarea sets Subarea field to given value.


### GetArea

`func (o *MonsterSpawnLocation) GetArea() CombatReference`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *MonsterSpawnLocation) GetAreaOk() (*CombatReference, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *MonsterSpawnLocation) SetArea(v CombatReference)`

SetArea sets Area field to given value.


### GetSuperArea

`func (o *MonsterSpawnLocation) GetSuperArea() CombatReference`

GetSuperArea returns the SuperArea field if non-nil, zero value otherwise.

### GetSuperAreaOk

`func (o *MonsterSpawnLocation) GetSuperAreaOk() (*CombatReference, bool)`

GetSuperAreaOk returns a tuple with the SuperArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperArea

`func (o *MonsterSpawnLocation) SetSuperArea(v CombatReference)`

SetSuperArea sets SuperArea field to given value.


### GetSubareaMaps

`func (o *MonsterSpawnLocation) GetSubareaMaps() []MonsterSpawnLocationSubareaMapsInner`

GetSubareaMaps returns the SubareaMaps field if non-nil, zero value otherwise.

### GetSubareaMapsOk

`func (o *MonsterSpawnLocation) GetSubareaMapsOk() (*[]MonsterSpawnLocationSubareaMapsInner, bool)`

GetSubareaMapsOk returns a tuple with the SubareaMaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubareaMaps

`func (o *MonsterSpawnLocation) SetSubareaMaps(v []MonsterSpawnLocationSubareaMapsInner)`

SetSubareaMaps sets SubareaMaps field to given value.


### GetPreferred

`func (o *MonsterSpawnLocation) GetPreferred() bool`

GetPreferred returns the Preferred field if non-nil, zero value otherwise.

### GetPreferredOk

`func (o *MonsterSpawnLocation) GetPreferredOk() (*bool, bool)`

GetPreferredOk returns a tuple with the Preferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferred

`func (o *MonsterSpawnLocation) SetPreferred(v bool)`

SetPreferred sets Preferred field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


