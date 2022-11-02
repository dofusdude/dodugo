# PutRSSWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Whitelist** | Pointer to **[]string** |  | [optional] 
**Blacklist** | Pointer to **[]string** |  | [optional] 
**Subscriptions** | Pointer to **[]string** | Get the available subscriptions with /meta/webhooks/rss | [optional] 
**PreviewLength** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewPutRSSWebhook

`func NewPutRSSWebhook() *PutRSSWebhook`

NewPutRSSWebhook instantiates a new PutRSSWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutRSSWebhookWithDefaults

`func NewPutRSSWebhookWithDefaults() *PutRSSWebhook`

NewPutRSSWebhookWithDefaults instantiates a new PutRSSWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWhitelist

`func (o *PutRSSWebhook) GetWhitelist() []string`

GetWhitelist returns the Whitelist field if non-nil, zero value otherwise.

### GetWhitelistOk

`func (o *PutRSSWebhook) GetWhitelistOk() (*[]string, bool)`

GetWhitelistOk returns a tuple with the Whitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelist

`func (o *PutRSSWebhook) SetWhitelist(v []string)`

SetWhitelist sets Whitelist field to given value.

### HasWhitelist

`func (o *PutRSSWebhook) HasWhitelist() bool`

HasWhitelist returns a boolean if a field has been set.

### SetWhitelistNil

`func (o *PutRSSWebhook) SetWhitelistNil(b bool)`

 SetWhitelistNil sets the value for Whitelist to be an explicit nil

### UnsetWhitelist
`func (o *PutRSSWebhook) UnsetWhitelist()`

UnsetWhitelist ensures that no value is present for Whitelist, not even an explicit nil
### GetBlacklist

`func (o *PutRSSWebhook) GetBlacklist() []string`

GetBlacklist returns the Blacklist field if non-nil, zero value otherwise.

### GetBlacklistOk

`func (o *PutRSSWebhook) GetBlacklistOk() (*[]string, bool)`

GetBlacklistOk returns a tuple with the Blacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklist

`func (o *PutRSSWebhook) SetBlacklist(v []string)`

SetBlacklist sets Blacklist field to given value.

### HasBlacklist

`func (o *PutRSSWebhook) HasBlacklist() bool`

HasBlacklist returns a boolean if a field has been set.

### SetBlacklistNil

`func (o *PutRSSWebhook) SetBlacklistNil(b bool)`

 SetBlacklistNil sets the value for Blacklist to be an explicit nil

### UnsetBlacklist
`func (o *PutRSSWebhook) UnsetBlacklist()`

UnsetBlacklist ensures that no value is present for Blacklist, not even an explicit nil
### GetSubscriptions

`func (o *PutRSSWebhook) GetSubscriptions() []string`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *PutRSSWebhook) GetSubscriptionsOk() (*[]string, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *PutRSSWebhook) SetSubscriptions(v []string)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *PutRSSWebhook) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### SetSubscriptionsNil

`func (o *PutRSSWebhook) SetSubscriptionsNil(b bool)`

 SetSubscriptionsNil sets the value for Subscriptions to be an explicit nil

### UnsetSubscriptions
`func (o *PutRSSWebhook) UnsetSubscriptions()`

UnsetSubscriptions ensures that no value is present for Subscriptions, not even an explicit nil
### GetPreviewLength

`func (o *PutRSSWebhook) GetPreviewLength() int32`

GetPreviewLength returns the PreviewLength field if non-nil, zero value otherwise.

### GetPreviewLengthOk

`func (o *PutRSSWebhook) GetPreviewLengthOk() (*int32, bool)`

GetPreviewLengthOk returns a tuple with the PreviewLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewLength

`func (o *PutRSSWebhook) SetPreviewLength(v int32)`

SetPreviewLength sets PreviewLength field to given value.

### HasPreviewLength

`func (o *PutRSSWebhook) HasPreviewLength() bool`

HasPreviewLength returns a boolean if a field has been set.

### SetPreviewLengthNil

`func (o *PutRSSWebhook) SetPreviewLengthNil(b bool)`

 SetPreviewLengthNil sets the value for PreviewLength to be an explicit nil

### UnsetPreviewLength
`func (o *PutRSSWebhook) UnsetPreviewLength()`

UnsetPreviewLength ensures that no value is present for PreviewLength, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


