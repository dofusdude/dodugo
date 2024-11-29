# Images

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Icon** | Pointer to **string** | 64x64 px, always available | [optional] 
**Sd** | Pointer to **NullableString** | around 2x the resolution of icon | [optional] 
**Hq** | Pointer to **NullableString** | around 2x the resolution of sd | [optional] 
**Hd** | Pointer to **NullableString** | around 2x the resolution of hd | [optional] 

## Methods

### NewImages

`func NewImages() *Images`

NewImages instantiates a new Images object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImagesWithDefaults

`func NewImagesWithDefaults() *Images`

NewImagesWithDefaults instantiates a new Images object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIcon

`func (o *Images) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Images) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Images) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Images) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetSd

`func (o *Images) GetSd() string`

GetSd returns the Sd field if non-nil, zero value otherwise.

### GetSdOk

`func (o *Images) GetSdOk() (*string, bool)`

GetSdOk returns a tuple with the Sd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSd

`func (o *Images) SetSd(v string)`

SetSd sets Sd field to given value.

### HasSd

`func (o *Images) HasSd() bool`

HasSd returns a boolean if a field has been set.

### SetSdNil

`func (o *Images) SetSdNil(b bool)`

 SetSdNil sets the value for Sd to be an explicit nil

### UnsetSd
`func (o *Images) UnsetSd()`

UnsetSd ensures that no value is present for Sd, not even an explicit nil
### GetHq

`func (o *Images) GetHq() string`

GetHq returns the Hq field if non-nil, zero value otherwise.

### GetHqOk

`func (o *Images) GetHqOk() (*string, bool)`

GetHqOk returns a tuple with the Hq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHq

`func (o *Images) SetHq(v string)`

SetHq sets Hq field to given value.

### HasHq

`func (o *Images) HasHq() bool`

HasHq returns a boolean if a field has been set.

### SetHqNil

`func (o *Images) SetHqNil(b bool)`

 SetHqNil sets the value for Hq to be an explicit nil

### UnsetHq
`func (o *Images) UnsetHq()`

UnsetHq ensures that no value is present for Hq, not even an explicit nil
### GetHd

`func (o *Images) GetHd() string`

GetHd returns the Hd field if non-nil, zero value otherwise.

### GetHdOk

`func (o *Images) GetHdOk() (*string, bool)`

GetHdOk returns a tuple with the Hd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHd

`func (o *Images) SetHd(v string)`

SetHd sets Hd field to given value.

### HasHd

`func (o *Images) HasHd() bool`

HasHd returns a boolean if a field has been set.

### SetHdNil

`func (o *Images) SetHdNil(b bool)`

 SetHdNil sets the value for Hd to be an explicit nil

### UnsetHd
`func (o *Images) UnsetHd()`

UnsetHd ensures that no value is present for Hd, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


