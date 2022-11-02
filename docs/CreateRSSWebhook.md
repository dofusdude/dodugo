# CreateRSSWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Whitelist** | Pointer to **[]string** |  | [optional] 
**Blacklist** | Pointer to **[]string** |  | [optional] 
**Subscriptions** | **[]string** | Get the available subscriptions with /meta/webhooks/rss | 
**Format** | **string** |  | 
**PreviewLength** | Pointer to **NullableInt32** |  | [optional] 
**Callback** | **string** | Discord Webhook URL | 

## Methods

### NewCreateRSSWebhook

`func NewCreateRSSWebhook(subscriptions []string, format string, callback string, ) *CreateRSSWebhook`

NewCreateRSSWebhook instantiates a new CreateRSSWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRSSWebhookWithDefaults

`func NewCreateRSSWebhookWithDefaults() *CreateRSSWebhook`

NewCreateRSSWebhookWithDefaults instantiates a new CreateRSSWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWhitelist

`func (o *CreateRSSWebhook) GetWhitelist() []string`

GetWhitelist returns the Whitelist field if non-nil, zero value otherwise.

### GetWhitelistOk

`func (o *CreateRSSWebhook) GetWhitelistOk() (*[]string, bool)`

GetWhitelistOk returns a tuple with the Whitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelist

`func (o *CreateRSSWebhook) SetWhitelist(v []string)`

SetWhitelist sets Whitelist field to given value.

### HasWhitelist

`func (o *CreateRSSWebhook) HasWhitelist() bool`

HasWhitelist returns a boolean if a field has been set.

### SetWhitelistNil

`func (o *CreateRSSWebhook) SetWhitelistNil(b bool)`

 SetWhitelistNil sets the value for Whitelist to be an explicit nil

### UnsetWhitelist
`func (o *CreateRSSWebhook) UnsetWhitelist()`

UnsetWhitelist ensures that no value is present for Whitelist, not even an explicit nil
### GetBlacklist

`func (o *CreateRSSWebhook) GetBlacklist() []string`

GetBlacklist returns the Blacklist field if non-nil, zero value otherwise.

### GetBlacklistOk

`func (o *CreateRSSWebhook) GetBlacklistOk() (*[]string, bool)`

GetBlacklistOk returns a tuple with the Blacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklist

`func (o *CreateRSSWebhook) SetBlacklist(v []string)`

SetBlacklist sets Blacklist field to given value.

### HasBlacklist

`func (o *CreateRSSWebhook) HasBlacklist() bool`

HasBlacklist returns a boolean if a field has been set.

### SetBlacklistNil

`func (o *CreateRSSWebhook) SetBlacklistNil(b bool)`

 SetBlacklistNil sets the value for Blacklist to be an explicit nil

### UnsetBlacklist
`func (o *CreateRSSWebhook) UnsetBlacklist()`

UnsetBlacklist ensures that no value is present for Blacklist, not even an explicit nil
### GetSubscriptions

`func (o *CreateRSSWebhook) GetSubscriptions() []string`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *CreateRSSWebhook) GetSubscriptionsOk() (*[]string, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *CreateRSSWebhook) SetSubscriptions(v []string)`

SetSubscriptions sets Subscriptions field to given value.


### GetFormat

`func (o *CreateRSSWebhook) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateRSSWebhook) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateRSSWebhook) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetPreviewLength

`func (o *CreateRSSWebhook) GetPreviewLength() int32`

GetPreviewLength returns the PreviewLength field if non-nil, zero value otherwise.

### GetPreviewLengthOk

`func (o *CreateRSSWebhook) GetPreviewLengthOk() (*int32, bool)`

GetPreviewLengthOk returns a tuple with the PreviewLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewLength

`func (o *CreateRSSWebhook) SetPreviewLength(v int32)`

SetPreviewLength sets PreviewLength field to given value.

### HasPreviewLength

`func (o *CreateRSSWebhook) HasPreviewLength() bool`

HasPreviewLength returns a boolean if a field has been set.

### SetPreviewLengthNil

`func (o *CreateRSSWebhook) SetPreviewLengthNil(b bool)`

 SetPreviewLengthNil sets the value for PreviewLength to be an explicit nil

### UnsetPreviewLength
`func (o *CreateRSSWebhook) UnsetPreviewLength()`

UnsetPreviewLength ensures that no value is present for PreviewLength, not even an explicit nil
### GetCallback

`func (o *CreateRSSWebhook) GetCallback() string`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *CreateRSSWebhook) GetCallbackOk() (*string, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *CreateRSSWebhook) SetCallback(v string)`

SetCallback sets Callback field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


