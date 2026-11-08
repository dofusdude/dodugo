# WeaponCombat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Zone** | [**EffectZone**](EffectZone.md) |  | 
**CastInLine** | **bool** |  | 
**CastInDiagonal** | **bool** |  | 
**RequiresLineOfSight** | **bool** |  | 

## Methods

### NewWeaponCombat

`func NewWeaponCombat(zone EffectZone, castInLine bool, castInDiagonal bool, requiresLineOfSight bool, ) *WeaponCombat`

NewWeaponCombat instantiates a new WeaponCombat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeaponCombatWithDefaults

`func NewWeaponCombatWithDefaults() *WeaponCombat`

NewWeaponCombatWithDefaults instantiates a new WeaponCombat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZone

`func (o *WeaponCombat) GetZone() EffectZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *WeaponCombat) GetZoneOk() (*EffectZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *WeaponCombat) SetZone(v EffectZone)`

SetZone sets Zone field to given value.


### GetCastInLine

`func (o *WeaponCombat) GetCastInLine() bool`

GetCastInLine returns the CastInLine field if non-nil, zero value otherwise.

### GetCastInLineOk

`func (o *WeaponCombat) GetCastInLineOk() (*bool, bool)`

GetCastInLineOk returns a tuple with the CastInLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastInLine

`func (o *WeaponCombat) SetCastInLine(v bool)`

SetCastInLine sets CastInLine field to given value.


### GetCastInDiagonal

`func (o *WeaponCombat) GetCastInDiagonal() bool`

GetCastInDiagonal returns the CastInDiagonal field if non-nil, zero value otherwise.

### GetCastInDiagonalOk

`func (o *WeaponCombat) GetCastInDiagonalOk() (*bool, bool)`

GetCastInDiagonalOk returns a tuple with the CastInDiagonal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastInDiagonal

`func (o *WeaponCombat) SetCastInDiagonal(v bool)`

SetCastInDiagonal sets CastInDiagonal field to given value.


### GetRequiresLineOfSight

`func (o *WeaponCombat) GetRequiresLineOfSight() bool`

GetRequiresLineOfSight returns the RequiresLineOfSight field if non-nil, zero value otherwise.

### GetRequiresLineOfSightOk

`func (o *WeaponCombat) GetRequiresLineOfSightOk() (*bool, bool)`

GetRequiresLineOfSightOk returns a tuple with the RequiresLineOfSight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresLineOfSight

`func (o *WeaponCombat) SetRequiresLineOfSight(v bool)`

SetRequiresLineOfSight sets RequiresLineOfSight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


