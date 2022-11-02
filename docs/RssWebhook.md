# RssWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Whitelist** | Pointer to **[]string** |  | [optional] 
**Blacklist** | Pointer to **[]string** |  | [optional] 
**Subscriptions** | Pointer to **[]string** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**PreviewLength** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastFiredAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRssWebhook

`func NewRssWebhook() *RssWebhook`

NewRssWebhook instantiates a new RssWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRssWebhookWithDefaults

`func NewRssWebhookWithDefaults() *RssWebhook`

NewRssWebhookWithDefaults instantiates a new RssWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RssWebhook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RssWebhook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RssWebhook) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RssWebhook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWhitelist

`func (o *RssWebhook) GetWhitelist() []string`

GetWhitelist returns the Whitelist field if non-nil, zero value otherwise.

### GetWhitelistOk

`func (o *RssWebhook) GetWhitelistOk() (*[]string, bool)`

GetWhitelistOk returns a tuple with the Whitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelist

`func (o *RssWebhook) SetWhitelist(v []string)`

SetWhitelist sets Whitelist field to given value.

### HasWhitelist

`func (o *RssWebhook) HasWhitelist() bool`

HasWhitelist returns a boolean if a field has been set.

### SetWhitelistNil

`func (o *RssWebhook) SetWhitelistNil(b bool)`

 SetWhitelistNil sets the value for Whitelist to be an explicit nil

### UnsetWhitelist
`func (o *RssWebhook) UnsetWhitelist()`

UnsetWhitelist ensures that no value is present for Whitelist, not even an explicit nil
### GetBlacklist

`func (o *RssWebhook) GetBlacklist() []string`

GetBlacklist returns the Blacklist field if non-nil, zero value otherwise.

### GetBlacklistOk

`func (o *RssWebhook) GetBlacklistOk() (*[]string, bool)`

GetBlacklistOk returns a tuple with the Blacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklist

`func (o *RssWebhook) SetBlacklist(v []string)`

SetBlacklist sets Blacklist field to given value.

### HasBlacklist

`func (o *RssWebhook) HasBlacklist() bool`

HasBlacklist returns a boolean if a field has been set.

### SetBlacklistNil

`func (o *RssWebhook) SetBlacklistNil(b bool)`

 SetBlacklistNil sets the value for Blacklist to be an explicit nil

### UnsetBlacklist
`func (o *RssWebhook) UnsetBlacklist()`

UnsetBlacklist ensures that no value is present for Blacklist, not even an explicit nil
### GetSubscriptions

`func (o *RssWebhook) GetSubscriptions() []string`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *RssWebhook) GetSubscriptionsOk() (*[]string, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *RssWebhook) SetSubscriptions(v []string)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *RssWebhook) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetFormat

`func (o *RssWebhook) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *RssWebhook) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *RssWebhook) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *RssWebhook) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetPreviewLength

`func (o *RssWebhook) GetPreviewLength() int32`

GetPreviewLength returns the PreviewLength field if non-nil, zero value otherwise.

### GetPreviewLengthOk

`func (o *RssWebhook) GetPreviewLengthOk() (*int32, bool)`

GetPreviewLengthOk returns a tuple with the PreviewLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewLength

`func (o *RssWebhook) SetPreviewLength(v int32)`

SetPreviewLength sets PreviewLength field to given value.

### HasPreviewLength

`func (o *RssWebhook) HasPreviewLength() bool`

HasPreviewLength returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RssWebhook) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RssWebhook) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RssWebhook) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RssWebhook) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastFiredAt

`func (o *RssWebhook) GetLastFiredAt() time.Time`

GetLastFiredAt returns the LastFiredAt field if non-nil, zero value otherwise.

### GetLastFiredAtOk

`func (o *RssWebhook) GetLastFiredAtOk() (*time.Time, bool)`

GetLastFiredAtOk returns a tuple with the LastFiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFiredAt

`func (o *RssWebhook) SetLastFiredAt(v time.Time)`

SetLastFiredAt sets LastFiredAt field to given value.

### HasLastFiredAt

`func (o *RssWebhook) HasLastFiredAt() bool`

HasLastFiredAt returns a boolean if a field has been set.

### SetLastFiredAtNil

`func (o *RssWebhook) SetLastFiredAtNil(b bool)`

 SetLastFiredAtNil sets the value for LastFiredAt to be an explicit nil

### UnsetLastFiredAt
`func (o *RssWebhook) UnsetLastFiredAt()`

UnsetLastFiredAt ensures that no value is present for LastFiredAt, not even an explicit nil
### GetUpdatedAt

`func (o *RssWebhook) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RssWebhook) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RssWebhook) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RssWebhook) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


