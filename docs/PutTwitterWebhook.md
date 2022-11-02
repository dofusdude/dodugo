# PutTwitterWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Whitelist** | Pointer to **[]string** |  | [optional] 
**Blacklist** | Pointer to **[]string** |  | [optional] 
**Subscriptions** | Pointer to **[]string** | Get the available subscriptions with /meta/webhooks/twitter | [optional] 
**PreviewLength** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewPutTwitterWebhook

`func NewPutTwitterWebhook() *PutTwitterWebhook`

NewPutTwitterWebhook instantiates a new PutTwitterWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutTwitterWebhookWithDefaults

`func NewPutTwitterWebhookWithDefaults() *PutTwitterWebhook`

NewPutTwitterWebhookWithDefaults instantiates a new PutTwitterWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWhitelist

`func (o *PutTwitterWebhook) GetWhitelist() []string`

GetWhitelist returns the Whitelist field if non-nil, zero value otherwise.

### GetWhitelistOk

`func (o *PutTwitterWebhook) GetWhitelistOk() (*[]string, bool)`

GetWhitelistOk returns a tuple with the Whitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelist

`func (o *PutTwitterWebhook) SetWhitelist(v []string)`

SetWhitelist sets Whitelist field to given value.

### HasWhitelist

`func (o *PutTwitterWebhook) HasWhitelist() bool`

HasWhitelist returns a boolean if a field has been set.

### SetWhitelistNil

`func (o *PutTwitterWebhook) SetWhitelistNil(b bool)`

 SetWhitelistNil sets the value for Whitelist to be an explicit nil

### UnsetWhitelist
`func (o *PutTwitterWebhook) UnsetWhitelist()`

UnsetWhitelist ensures that no value is present for Whitelist, not even an explicit nil
### GetBlacklist

`func (o *PutTwitterWebhook) GetBlacklist() []string`

GetBlacklist returns the Blacklist field if non-nil, zero value otherwise.

### GetBlacklistOk

`func (o *PutTwitterWebhook) GetBlacklistOk() (*[]string, bool)`

GetBlacklistOk returns a tuple with the Blacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklist

`func (o *PutTwitterWebhook) SetBlacklist(v []string)`

SetBlacklist sets Blacklist field to given value.

### HasBlacklist

`func (o *PutTwitterWebhook) HasBlacklist() bool`

HasBlacklist returns a boolean if a field has been set.

### SetBlacklistNil

`func (o *PutTwitterWebhook) SetBlacklistNil(b bool)`

 SetBlacklistNil sets the value for Blacklist to be an explicit nil

### UnsetBlacklist
`func (o *PutTwitterWebhook) UnsetBlacklist()`

UnsetBlacklist ensures that no value is present for Blacklist, not even an explicit nil
### GetSubscriptions

`func (o *PutTwitterWebhook) GetSubscriptions() []string`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *PutTwitterWebhook) GetSubscriptionsOk() (*[]string, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *PutTwitterWebhook) SetSubscriptions(v []string)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *PutTwitterWebhook) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### SetSubscriptionsNil

`func (o *PutTwitterWebhook) SetSubscriptionsNil(b bool)`

 SetSubscriptionsNil sets the value for Subscriptions to be an explicit nil

### UnsetSubscriptions
`func (o *PutTwitterWebhook) UnsetSubscriptions()`

UnsetSubscriptions ensures that no value is present for Subscriptions, not even an explicit nil
### GetPreviewLength

`func (o *PutTwitterWebhook) GetPreviewLength() int32`

GetPreviewLength returns the PreviewLength field if non-nil, zero value otherwise.

### GetPreviewLengthOk

`func (o *PutTwitterWebhook) GetPreviewLengthOk() (*int32, bool)`

GetPreviewLengthOk returns a tuple with the PreviewLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewLength

`func (o *PutTwitterWebhook) SetPreviewLength(v int32)`

SetPreviewLength sets PreviewLength field to given value.

### HasPreviewLength

`func (o *PutTwitterWebhook) HasPreviewLength() bool`

HasPreviewLength returns a boolean if a field has been set.

### SetPreviewLengthNil

`func (o *PutTwitterWebhook) SetPreviewLengthNil(b bool)`

 SetPreviewLengthNil sets the value for PreviewLength to be an explicit nil

### UnsetPreviewLength
`func (o *PutTwitterWebhook) UnsetPreviewLength()`

UnsetPreviewLength ensures that no value is present for PreviewLength, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


