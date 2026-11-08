# MonsterRace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | **int32** |  | 
**Name** | **string** |  | 
**SuperRace** | [**CombatReference**](CombatReference.md) |  | 

## Methods

### NewMonsterRace

`func NewMonsterRace(ankamaId int32, name string, superRace CombatReference, ) *MonsterRace`

NewMonsterRace instantiates a new MonsterRace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonsterRaceWithDefaults

`func NewMonsterRaceWithDefaults() *MonsterRace`

NewMonsterRaceWithDefaults instantiates a new MonsterRace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *MonsterRace) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *MonsterRace) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *MonsterRace) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.


### GetName

`func (o *MonsterRace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonsterRace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonsterRace) SetName(v string)`

SetName sets Name field to given value.


### GetSuperRace

`func (o *MonsterRace) GetSuperRace() CombatReference`

GetSuperRace returns the SuperRace field if non-nil, zero value otherwise.

### GetSuperRaceOk

`func (o *MonsterRace) GetSuperRaceOk() (*CombatReference, bool)`

GetSuperRaceOk returns a tuple with the SuperRace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperRace

`func (o *MonsterRace) SetSuperRace(v CombatReference)`

SetSuperRace sets SuperRace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


