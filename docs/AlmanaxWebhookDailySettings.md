# AlmanaxWebhookDailySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | Pointer to **string** | Timezone of your community to determine midnight. | [optional] [default to "Europe/Paris"]
**MidnightOffset** | Pointer to **int32** | Hours added to midnight at the selected timezone. 8 &#x3D; 8:00 in the morning. | [optional] [default to 0]

## Methods

### NewAlmanaxWebhookDailySettings

`func NewAlmanaxWebhookDailySettings() *AlmanaxWebhookDailySettings`

NewAlmanaxWebhookDailySettings instantiates a new AlmanaxWebhookDailySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlmanaxWebhookDailySettingsWithDefaults

`func NewAlmanaxWebhookDailySettingsWithDefaults() *AlmanaxWebhookDailySettings`

NewAlmanaxWebhookDailySettingsWithDefaults instantiates a new AlmanaxWebhookDailySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezone

`func (o *AlmanaxWebhookDailySettings) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AlmanaxWebhookDailySettings) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AlmanaxWebhookDailySettings) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *AlmanaxWebhookDailySettings) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetMidnightOffset

`func (o *AlmanaxWebhookDailySettings) GetMidnightOffset() int32`

GetMidnightOffset returns the MidnightOffset field if non-nil, zero value otherwise.

### GetMidnightOffsetOk

`func (o *AlmanaxWebhookDailySettings) GetMidnightOffsetOk() (*int32, bool)`

GetMidnightOffsetOk returns a tuple with the MidnightOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMidnightOffset

`func (o *AlmanaxWebhookDailySettings) SetMidnightOffset(v int32)`

SetMidnightOffset sets MidnightOffset field to given value.

### HasMidnightOffset

`func (o *AlmanaxWebhookDailySettings) HasMidnightOffset() bool`

HasMidnightOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


