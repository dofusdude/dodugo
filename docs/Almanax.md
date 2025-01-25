# Almanax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bonus** | Pointer to [**AlmanaxBonus**](AlmanaxBonus.md) |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**Tribute** | Pointer to [**AlmanaxTribute**](AlmanaxTribute.md) |  | [optional] 
**RewardKamas** | Pointer to **NullableInt32** | Amount of Kamas you get as reward for finishing this Almanax quest. | [optional] 
**RewardXp** | Pointer to **NullableInt32** | Optional field that shows when a level is given in the request. Shows the experience points you get this day for finishing this Almanax quest. | [optional] 

## Methods

### NewAlmanax

`func NewAlmanax() *Almanax`

NewAlmanax instantiates a new Almanax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlmanaxWithDefaults

`func NewAlmanaxWithDefaults() *Almanax`

NewAlmanaxWithDefaults instantiates a new Almanax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBonus

`func (o *Almanax) GetBonus() AlmanaxBonus`

GetBonus returns the Bonus field if non-nil, zero value otherwise.

### GetBonusOk

`func (o *Almanax) GetBonusOk() (*AlmanaxBonus, bool)`

GetBonusOk returns a tuple with the Bonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonus

`func (o *Almanax) SetBonus(v AlmanaxBonus)`

SetBonus sets Bonus field to given value.

### HasBonus

`func (o *Almanax) HasBonus() bool`

HasBonus returns a boolean if a field has been set.

### GetDate

`func (o *Almanax) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Almanax) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Almanax) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Almanax) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetTribute

`func (o *Almanax) GetTribute() AlmanaxTribute`

GetTribute returns the Tribute field if non-nil, zero value otherwise.

### GetTributeOk

`func (o *Almanax) GetTributeOk() (*AlmanaxTribute, bool)`

GetTributeOk returns a tuple with the Tribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTribute

`func (o *Almanax) SetTribute(v AlmanaxTribute)`

SetTribute sets Tribute field to given value.

### HasTribute

`func (o *Almanax) HasTribute() bool`

HasTribute returns a boolean if a field has been set.

### GetRewardKamas

`func (o *Almanax) GetRewardKamas() int32`

GetRewardKamas returns the RewardKamas field if non-nil, zero value otherwise.

### GetRewardKamasOk

`func (o *Almanax) GetRewardKamasOk() (*int32, bool)`

GetRewardKamasOk returns a tuple with the RewardKamas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardKamas

`func (o *Almanax) SetRewardKamas(v int32)`

SetRewardKamas sets RewardKamas field to given value.

### HasRewardKamas

`func (o *Almanax) HasRewardKamas() bool`

HasRewardKamas returns a boolean if a field has been set.

### SetRewardKamasNil

`func (o *Almanax) SetRewardKamasNil(b bool)`

 SetRewardKamasNil sets the value for RewardKamas to be an explicit nil

### UnsetRewardKamas
`func (o *Almanax) UnsetRewardKamas()`

UnsetRewardKamas ensures that no value is present for RewardKamas, not even an explicit nil
### GetRewardXp

`func (o *Almanax) GetRewardXp() int32`

GetRewardXp returns the RewardXp field if non-nil, zero value otherwise.

### GetRewardXpOk

`func (o *Almanax) GetRewardXpOk() (*int32, bool)`

GetRewardXpOk returns a tuple with the RewardXp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardXp

`func (o *Almanax) SetRewardXp(v int32)`

SetRewardXp sets RewardXp field to given value.

### HasRewardXp

`func (o *Almanax) HasRewardXp() bool`

HasRewardXp returns a boolean if a field has been set.

### SetRewardXpNil

`func (o *Almanax) SetRewardXpNil(b bool)`

 SetRewardXpNil sets the value for RewardXp to be an explicit nil

### UnsetRewardXp
`func (o *Almanax) UnsetRewardXp()`

UnsetRewardXp ensures that no value is present for RewardXp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


