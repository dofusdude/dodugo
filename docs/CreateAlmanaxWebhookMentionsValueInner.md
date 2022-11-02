# CreateAlmanaxWebhookMentionsValueInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscordId** | Pointer to **int32** | User or role ID directly from Discord. Activate developer mode and right click a user or role to get them. | [optional] 
**IsRole** | Pointer to **bool** | Whether an ID describes a role (true) or user (false). This is needed for formatting the mention command so Discord understands it. | [optional] 
**PingDaysBefore** | Pointer to **NullableInt32** | Get a mention days before the bonus comes up. It will post on the specified time. Also works when the interval is not daily. | [optional] 

## Methods

### NewCreateAlmanaxWebhookMentionsValueInner

`func NewCreateAlmanaxWebhookMentionsValueInner() *CreateAlmanaxWebhookMentionsValueInner`

NewCreateAlmanaxWebhookMentionsValueInner instantiates a new CreateAlmanaxWebhookMentionsValueInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAlmanaxWebhookMentionsValueInnerWithDefaults

`func NewCreateAlmanaxWebhookMentionsValueInnerWithDefaults() *CreateAlmanaxWebhookMentionsValueInner`

NewCreateAlmanaxWebhookMentionsValueInnerWithDefaults instantiates a new CreateAlmanaxWebhookMentionsValueInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscordId

`func (o *CreateAlmanaxWebhookMentionsValueInner) GetDiscordId() int32`

GetDiscordId returns the DiscordId field if non-nil, zero value otherwise.

### GetDiscordIdOk

`func (o *CreateAlmanaxWebhookMentionsValueInner) GetDiscordIdOk() (*int32, bool)`

GetDiscordIdOk returns a tuple with the DiscordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordId

`func (o *CreateAlmanaxWebhookMentionsValueInner) SetDiscordId(v int32)`

SetDiscordId sets DiscordId field to given value.

### HasDiscordId

`func (o *CreateAlmanaxWebhookMentionsValueInner) HasDiscordId() bool`

HasDiscordId returns a boolean if a field has been set.

### GetIsRole

`func (o *CreateAlmanaxWebhookMentionsValueInner) GetIsRole() bool`

GetIsRole returns the IsRole field if non-nil, zero value otherwise.

### GetIsRoleOk

`func (o *CreateAlmanaxWebhookMentionsValueInner) GetIsRoleOk() (*bool, bool)`

GetIsRoleOk returns a tuple with the IsRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRole

`func (o *CreateAlmanaxWebhookMentionsValueInner) SetIsRole(v bool)`

SetIsRole sets IsRole field to given value.

### HasIsRole

`func (o *CreateAlmanaxWebhookMentionsValueInner) HasIsRole() bool`

HasIsRole returns a boolean if a field has been set.

### GetPingDaysBefore

`func (o *CreateAlmanaxWebhookMentionsValueInner) GetPingDaysBefore() int32`

GetPingDaysBefore returns the PingDaysBefore field if non-nil, zero value otherwise.

### GetPingDaysBeforeOk

`func (o *CreateAlmanaxWebhookMentionsValueInner) GetPingDaysBeforeOk() (*int32, bool)`

GetPingDaysBeforeOk returns a tuple with the PingDaysBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingDaysBefore

`func (o *CreateAlmanaxWebhookMentionsValueInner) SetPingDaysBefore(v int32)`

SetPingDaysBefore sets PingDaysBefore field to given value.

### HasPingDaysBefore

`func (o *CreateAlmanaxWebhookMentionsValueInner) HasPingDaysBefore() bool`

HasPingDaysBefore returns a boolean if a field has been set.

### SetPingDaysBeforeNil

`func (o *CreateAlmanaxWebhookMentionsValueInner) SetPingDaysBeforeNil(b bool)`

 SetPingDaysBeforeNil sets the value for PingDaysBefore to be an explicit nil

### UnsetPingDaysBefore
`func (o *CreateAlmanaxWebhookMentionsValueInner) UnsetPingDaysBefore()`

UnsetPingDaysBefore ensures that no value is present for PingDaysBefore, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


