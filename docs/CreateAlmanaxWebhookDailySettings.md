# CreateAlmanaxWebhookDailySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | Pointer to **NullableString** | Timezone of your community to determine midnight. | [optional] [default to "Europe/Paris"]
**MidnightOffset** | Pointer to **NullableInt32** | Hours added to midnight at the selected timezone. 8 &#x3D; 8:00 in the morning. | [optional] [default to 0]

## Methods

### NewCreateAlmanaxWebhookDailySettings

`func NewCreateAlmanaxWebhookDailySettings() *CreateAlmanaxWebhookDailySettings`

NewCreateAlmanaxWebhookDailySettings instantiates a new CreateAlmanaxWebhookDailySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAlmanaxWebhookDailySettingsWithDefaults

`func NewCreateAlmanaxWebhookDailySettingsWithDefaults() *CreateAlmanaxWebhookDailySettings`

NewCreateAlmanaxWebhookDailySettingsWithDefaults instantiates a new CreateAlmanaxWebhookDailySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezone

`func (o *CreateAlmanaxWebhookDailySettings) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *CreateAlmanaxWebhookDailySettings) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *CreateAlmanaxWebhookDailySettings) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *CreateAlmanaxWebhookDailySettings) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *CreateAlmanaxWebhookDailySettings) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *CreateAlmanaxWebhookDailySettings) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetMidnightOffset

`func (o *CreateAlmanaxWebhookDailySettings) GetMidnightOffset() int32`

GetMidnightOffset returns the MidnightOffset field if non-nil, zero value otherwise.

### GetMidnightOffsetOk

`func (o *CreateAlmanaxWebhookDailySettings) GetMidnightOffsetOk() (*int32, bool)`

GetMidnightOffsetOk returns a tuple with the MidnightOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMidnightOffset

`func (o *CreateAlmanaxWebhookDailySettings) SetMidnightOffset(v int32)`

SetMidnightOffset sets MidnightOffset field to given value.

### HasMidnightOffset

`func (o *CreateAlmanaxWebhookDailySettings) HasMidnightOffset() bool`

HasMidnightOffset returns a boolean if a field has been set.

### SetMidnightOffsetNil

`func (o *CreateAlmanaxWebhookDailySettings) SetMidnightOffsetNil(b bool)`

 SetMidnightOffsetNil sets the value for MidnightOffset to be an explicit nil

### UnsetMidnightOffset
`func (o *CreateAlmanaxWebhookDailySettings) UnsetMidnightOffset()`

UnsetMidnightOffset ensures that no value is present for MidnightOffset, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


