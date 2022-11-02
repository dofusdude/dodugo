# TwitterWebhook

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

### NewTwitterWebhook

`func NewTwitterWebhook() *TwitterWebhook`

NewTwitterWebhook instantiates a new TwitterWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwitterWebhookWithDefaults

`func NewTwitterWebhookWithDefaults() *TwitterWebhook`

NewTwitterWebhookWithDefaults instantiates a new TwitterWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TwitterWebhook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TwitterWebhook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TwitterWebhook) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TwitterWebhook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWhitelist

`func (o *TwitterWebhook) GetWhitelist() []string`

GetWhitelist returns the Whitelist field if non-nil, zero value otherwise.

### GetWhitelistOk

`func (o *TwitterWebhook) GetWhitelistOk() (*[]string, bool)`

GetWhitelistOk returns a tuple with the Whitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelist

`func (o *TwitterWebhook) SetWhitelist(v []string)`

SetWhitelist sets Whitelist field to given value.

### HasWhitelist

`func (o *TwitterWebhook) HasWhitelist() bool`

HasWhitelist returns a boolean if a field has been set.

### SetWhitelistNil

`func (o *TwitterWebhook) SetWhitelistNil(b bool)`

 SetWhitelistNil sets the value for Whitelist to be an explicit nil

### UnsetWhitelist
`func (o *TwitterWebhook) UnsetWhitelist()`

UnsetWhitelist ensures that no value is present for Whitelist, not even an explicit nil
### GetBlacklist

`func (o *TwitterWebhook) GetBlacklist() []string`

GetBlacklist returns the Blacklist field if non-nil, zero value otherwise.

### GetBlacklistOk

`func (o *TwitterWebhook) GetBlacklistOk() (*[]string, bool)`

GetBlacklistOk returns a tuple with the Blacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklist

`func (o *TwitterWebhook) SetBlacklist(v []string)`

SetBlacklist sets Blacklist field to given value.

### HasBlacklist

`func (o *TwitterWebhook) HasBlacklist() bool`

HasBlacklist returns a boolean if a field has been set.

### SetBlacklistNil

`func (o *TwitterWebhook) SetBlacklistNil(b bool)`

 SetBlacklistNil sets the value for Blacklist to be an explicit nil

### UnsetBlacklist
`func (o *TwitterWebhook) UnsetBlacklist()`

UnsetBlacklist ensures that no value is present for Blacklist, not even an explicit nil
### GetSubscriptions

`func (o *TwitterWebhook) GetSubscriptions() []string`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *TwitterWebhook) GetSubscriptionsOk() (*[]string, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *TwitterWebhook) SetSubscriptions(v []string)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *TwitterWebhook) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetFormat

`func (o *TwitterWebhook) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *TwitterWebhook) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *TwitterWebhook) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *TwitterWebhook) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetPreviewLength

`func (o *TwitterWebhook) GetPreviewLength() int32`

GetPreviewLength returns the PreviewLength field if non-nil, zero value otherwise.

### GetPreviewLengthOk

`func (o *TwitterWebhook) GetPreviewLengthOk() (*int32, bool)`

GetPreviewLengthOk returns a tuple with the PreviewLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewLength

`func (o *TwitterWebhook) SetPreviewLength(v int32)`

SetPreviewLength sets PreviewLength field to given value.

### HasPreviewLength

`func (o *TwitterWebhook) HasPreviewLength() bool`

HasPreviewLength returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TwitterWebhook) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TwitterWebhook) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TwitterWebhook) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TwitterWebhook) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastFiredAt

`func (o *TwitterWebhook) GetLastFiredAt() time.Time`

GetLastFiredAt returns the LastFiredAt field if non-nil, zero value otherwise.

### GetLastFiredAtOk

`func (o *TwitterWebhook) GetLastFiredAtOk() (*time.Time, bool)`

GetLastFiredAtOk returns a tuple with the LastFiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFiredAt

`func (o *TwitterWebhook) SetLastFiredAt(v time.Time)`

SetLastFiredAt sets LastFiredAt field to given value.

### HasLastFiredAt

`func (o *TwitterWebhook) HasLastFiredAt() bool`

HasLastFiredAt returns a boolean if a field has been set.

### SetLastFiredAtNil

`func (o *TwitterWebhook) SetLastFiredAtNil(b bool)`

 SetLastFiredAtNil sets the value for LastFiredAt to be an explicit nil

### UnsetLastFiredAt
`func (o *TwitterWebhook) UnsetLastFiredAt()`

UnsetLastFiredAt ensures that no value is present for LastFiredAt, not even an explicit nil
### GetUpdatedAt

`func (o *TwitterWebhook) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TwitterWebhook) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TwitterWebhook) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TwitterWebhook) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


