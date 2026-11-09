# SpellEffect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Localized effect category, such as Air damage. This is not a spell ID or a structured game reference. | [optional] 
**Min** | Pointer to **int32** |  | [optional] 
**Max** | Pointer to **int32** |  | [optional] 
**Duration** | Pointer to **int32** | Effect duration in fight turns as represented by the source. Not a real-time duration in seconds. | [optional] 
**Delay** | Pointer to **int32** | Delay before the effect in fight turns as represented by the source. | [optional] 
**Random** | Pointer to **float32** | Source random-selection value. Do not interpret it as a percentage without the effect&#39;s selection rules. | [optional] 
**Group** | Pointer to **int32** |  | [optional] 
**Dispellable** | Pointer to **int32** |  | [optional] 
**EffectUid** | Pointer to **int32** | Stable identifier of this concrete effect instance. | [optional] 
**TargetId** | Pointer to **int32** |  | [optional] 
**TriggeredDuration** | Pointer to **int32** |  | [optional] 
**DisplayZero** | Pointer to **bool** |  | [optional] 
**Modifier** | Pointer to **int32** |  | [optional] 
**Zone** | Pointer to [**EffectZone**](EffectZone.md) |  | [optional] 
**TargetMask** | Pointer to [**TargetMask**](TargetMask.md) |  | [optional] 
**Triggers** | Pointer to [**SpellTriggers**](SpellTriggers.md) |  | [optional] 
**Templated** | Pointer to **string** | Localized effect text suitable for display. Numeric effect properties remain available separately. | [optional] 

## Methods

### NewSpellEffect

`func NewSpellEffect() *SpellEffect`

NewSpellEffect instantiates a new SpellEffect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpellEffectWithDefaults

`func NewSpellEffectWithDefaults() *SpellEffect`

NewSpellEffectWithDefaults instantiates a new SpellEffect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SpellEffect) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpellEffect) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpellEffect) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SpellEffect) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMin

`func (o *SpellEffect) GetMin() int32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *SpellEffect) GetMinOk() (*int32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *SpellEffect) SetMin(v int32)`

SetMin sets Min field to given value.

### HasMin

`func (o *SpellEffect) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMax

`func (o *SpellEffect) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *SpellEffect) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *SpellEffect) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *SpellEffect) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetDuration

`func (o *SpellEffect) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SpellEffect) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SpellEffect) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SpellEffect) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDelay

`func (o *SpellEffect) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *SpellEffect) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *SpellEffect) SetDelay(v int32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *SpellEffect) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetRandom

`func (o *SpellEffect) GetRandom() float32`

GetRandom returns the Random field if non-nil, zero value otherwise.

### GetRandomOk

`func (o *SpellEffect) GetRandomOk() (*float32, bool)`

GetRandomOk returns a tuple with the Random field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandom

`func (o *SpellEffect) SetRandom(v float32)`

SetRandom sets Random field to given value.

### HasRandom

`func (o *SpellEffect) HasRandom() bool`

HasRandom returns a boolean if a field has been set.

### GetGroup

`func (o *SpellEffect) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SpellEffect) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SpellEffect) SetGroup(v int32)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *SpellEffect) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetDispellable

`func (o *SpellEffect) GetDispellable() int32`

GetDispellable returns the Dispellable field if non-nil, zero value otherwise.

### GetDispellableOk

`func (o *SpellEffect) GetDispellableOk() (*int32, bool)`

GetDispellableOk returns a tuple with the Dispellable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispellable

`func (o *SpellEffect) SetDispellable(v int32)`

SetDispellable sets Dispellable field to given value.

### HasDispellable

`func (o *SpellEffect) HasDispellable() bool`

HasDispellable returns a boolean if a field has been set.

### GetEffectUid

`func (o *SpellEffect) GetEffectUid() int32`

GetEffectUid returns the EffectUid field if non-nil, zero value otherwise.

### GetEffectUidOk

`func (o *SpellEffect) GetEffectUidOk() (*int32, bool)`

GetEffectUidOk returns a tuple with the EffectUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectUid

`func (o *SpellEffect) SetEffectUid(v int32)`

SetEffectUid sets EffectUid field to given value.

### HasEffectUid

`func (o *SpellEffect) HasEffectUid() bool`

HasEffectUid returns a boolean if a field has been set.

### GetTargetId

`func (o *SpellEffect) GetTargetId() int32`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *SpellEffect) GetTargetIdOk() (*int32, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *SpellEffect) SetTargetId(v int32)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *SpellEffect) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetTriggeredDuration

`func (o *SpellEffect) GetTriggeredDuration() int32`

GetTriggeredDuration returns the TriggeredDuration field if non-nil, zero value otherwise.

### GetTriggeredDurationOk

`func (o *SpellEffect) GetTriggeredDurationOk() (*int32, bool)`

GetTriggeredDurationOk returns a tuple with the TriggeredDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredDuration

`func (o *SpellEffect) SetTriggeredDuration(v int32)`

SetTriggeredDuration sets TriggeredDuration field to given value.

### HasTriggeredDuration

`func (o *SpellEffect) HasTriggeredDuration() bool`

HasTriggeredDuration returns a boolean if a field has been set.

### GetDisplayZero

`func (o *SpellEffect) GetDisplayZero() bool`

GetDisplayZero returns the DisplayZero field if non-nil, zero value otherwise.

### GetDisplayZeroOk

`func (o *SpellEffect) GetDisplayZeroOk() (*bool, bool)`

GetDisplayZeroOk returns a tuple with the DisplayZero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayZero

`func (o *SpellEffect) SetDisplayZero(v bool)`

SetDisplayZero sets DisplayZero field to given value.

### HasDisplayZero

`func (o *SpellEffect) HasDisplayZero() bool`

HasDisplayZero returns a boolean if a field has been set.

### GetModifier

`func (o *SpellEffect) GetModifier() int32`

GetModifier returns the Modifier field if non-nil, zero value otherwise.

### GetModifierOk

`func (o *SpellEffect) GetModifierOk() (*int32, bool)`

GetModifierOk returns a tuple with the Modifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifier

`func (o *SpellEffect) SetModifier(v int32)`

SetModifier sets Modifier field to given value.

### HasModifier

`func (o *SpellEffect) HasModifier() bool`

HasModifier returns a boolean if a field has been set.

### GetZone

`func (o *SpellEffect) GetZone() EffectZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *SpellEffect) GetZoneOk() (*EffectZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *SpellEffect) SetZone(v EffectZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *SpellEffect) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetTargetMask

`func (o *SpellEffect) GetTargetMask() TargetMask`

GetTargetMask returns the TargetMask field if non-nil, zero value otherwise.

### GetTargetMaskOk

`func (o *SpellEffect) GetTargetMaskOk() (*TargetMask, bool)`

GetTargetMaskOk returns a tuple with the TargetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMask

`func (o *SpellEffect) SetTargetMask(v TargetMask)`

SetTargetMask sets TargetMask field to given value.

### HasTargetMask

`func (o *SpellEffect) HasTargetMask() bool`

HasTargetMask returns a boolean if a field has been set.

### GetTriggers

`func (o *SpellEffect) GetTriggers() SpellTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *SpellEffect) GetTriggersOk() (*SpellTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *SpellEffect) SetTriggers(v SpellTriggers)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *SpellEffect) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetTemplated

`func (o *SpellEffect) GetTemplated() string`

GetTemplated returns the Templated field if non-nil, zero value otherwise.

### GetTemplatedOk

`func (o *SpellEffect) GetTemplatedOk() (*string, bool)`

GetTemplatedOk returns a tuple with the Templated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplated

`func (o *SpellEffect) SetTemplated(v string)`

SetTemplated sets Templated field to given value.

### HasTemplated

`func (o *SpellEffect) HasTemplated() bool`

HasTemplated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


