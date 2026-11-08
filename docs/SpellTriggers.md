# SpellTriggers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Raw** | **string** | Original trigger tokens retained for lossless debugging. | 
**Logic** | Pointer to **string** |  | [optional] 
**Events** | [**[]SpellTrigger**](SpellTrigger.md) |  | 

## Methods

### NewSpellTriggers

`func NewSpellTriggers(raw string, events []SpellTrigger, ) *SpellTriggers`

NewSpellTriggers instantiates a new SpellTriggers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpellTriggersWithDefaults

`func NewSpellTriggersWithDefaults() *SpellTriggers`

NewSpellTriggersWithDefaults instantiates a new SpellTriggers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRaw

`func (o *SpellTriggers) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *SpellTriggers) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *SpellTriggers) SetRaw(v string)`

SetRaw sets Raw field to given value.


### GetLogic

`func (o *SpellTriggers) GetLogic() string`

GetLogic returns the Logic field if non-nil, zero value otherwise.

### GetLogicOk

`func (o *SpellTriggers) GetLogicOk() (*string, bool)`

GetLogicOk returns a tuple with the Logic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogic

`func (o *SpellTriggers) SetLogic(v string)`

SetLogic sets Logic field to given value.

### HasLogic

`func (o *SpellTriggers) HasLogic() bool`

HasLogic returns a boolean if a field has been set.

### GetEvents

`func (o *SpellTriggers) GetEvents() []SpellTrigger`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SpellTriggers) GetEventsOk() (*[]SpellTrigger, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SpellTriggers) SetEvents(v []SpellTrigger)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


