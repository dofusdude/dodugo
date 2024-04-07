# AlmanaxEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bonus** | Pointer to [**AlmanaxEntryBonus**](AlmanaxEntryBonus.md) |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**Tribute** | Pointer to [**AlmanaxEntryTribute**](AlmanaxEntryTribute.md) |  | [optional] 
**RewardKamas** | Pointer to **NullableInt32** | Amount of Kamas you get as reward for finishing this Almanax quest. | [optional] 

## Methods

### NewAlmanaxEntry

`func NewAlmanaxEntry() *AlmanaxEntry`

NewAlmanaxEntry instantiates a new AlmanaxEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlmanaxEntryWithDefaults

`func NewAlmanaxEntryWithDefaults() *AlmanaxEntry`

NewAlmanaxEntryWithDefaults instantiates a new AlmanaxEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBonus

`func (o *AlmanaxEntry) GetBonus() AlmanaxEntryBonus`

GetBonus returns the Bonus field if non-nil, zero value otherwise.

### GetBonusOk

`func (o *AlmanaxEntry) GetBonusOk() (*AlmanaxEntryBonus, bool)`

GetBonusOk returns a tuple with the Bonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonus

`func (o *AlmanaxEntry) SetBonus(v AlmanaxEntryBonus)`

SetBonus sets Bonus field to given value.

### HasBonus

`func (o *AlmanaxEntry) HasBonus() bool`

HasBonus returns a boolean if a field has been set.

### GetDate

`func (o *AlmanaxEntry) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AlmanaxEntry) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AlmanaxEntry) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *AlmanaxEntry) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetTribute

`func (o *AlmanaxEntry) GetTribute() AlmanaxEntryTribute`

GetTribute returns the Tribute field if non-nil, zero value otherwise.

### GetTributeOk

`func (o *AlmanaxEntry) GetTributeOk() (*AlmanaxEntryTribute, bool)`

GetTributeOk returns a tuple with the Tribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTribute

`func (o *AlmanaxEntry) SetTribute(v AlmanaxEntryTribute)`

SetTribute sets Tribute field to given value.

### HasTribute

`func (o *AlmanaxEntry) HasTribute() bool`

HasTribute returns a boolean if a field has been set.

### GetRewardKamas

`func (o *AlmanaxEntry) GetRewardKamas() int32`

GetRewardKamas returns the RewardKamas field if non-nil, zero value otherwise.

### GetRewardKamasOk

`func (o *AlmanaxEntry) GetRewardKamasOk() (*int32, bool)`

GetRewardKamasOk returns a tuple with the RewardKamas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardKamas

`func (o *AlmanaxEntry) SetRewardKamas(v int32)`

SetRewardKamas sets RewardKamas field to given value.

### HasRewardKamas

`func (o *AlmanaxEntry) HasRewardKamas() bool`

HasRewardKamas returns a boolean if a field has been set.

### SetRewardKamasNil

`func (o *AlmanaxEntry) SetRewardKamasNil(b bool)`

 SetRewardKamasNil sets the value for RewardKamas to be an explicit nil

### UnsetRewardKamas
`func (o *AlmanaxEntry) UnsetRewardKamas()`

UnsetRewardKamas ensures that no value is present for RewardKamas, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


