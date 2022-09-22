# Mount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**FamilyName** | Pointer to **string** |  | [optional] 
**ImageUrls** | Pointer to [**ImageUrls**](ImageUrls.md) |  | [optional] 
**HasEffects** | Pointer to **bool** |  | [optional] 
**Effects** | Pointer to [**[]EffectsEntry**](EffectsEntry.md) |  | [optional] 

## Methods

### NewMount

`func NewMount() *Mount`

NewMount instantiates a new Mount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMountWithDefaults

`func NewMountWithDefaults() *Mount`

NewMountWithDefaults instantiates a new Mount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *Mount) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *Mount) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *Mount) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.

### HasAnkamaId

`func (o *Mount) HasAnkamaId() bool`

HasAnkamaId returns a boolean if a field has been set.

### GetName

`func (o *Mount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Mount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Mount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Mount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFamilyName

`func (o *Mount) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *Mount) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *Mount) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *Mount) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### GetImageUrls

`func (o *Mount) GetImageUrls() ImageUrls`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *Mount) GetImageUrlsOk() (*ImageUrls, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *Mount) SetImageUrls(v ImageUrls)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *Mount) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.

### GetHasEffects

`func (o *Mount) GetHasEffects() bool`

GetHasEffects returns the HasEffects field if non-nil, zero value otherwise.

### GetHasEffectsOk

`func (o *Mount) GetHasEffectsOk() (*bool, bool)`

GetHasEffectsOk returns a tuple with the HasEffects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEffects

`func (o *Mount) SetHasEffects(v bool)`

SetHasEffects sets HasEffects field to given value.

### HasHasEffects

`func (o *Mount) HasHasEffects() bool`

HasHasEffects returns a boolean if a field has been set.

### GetEffects

`func (o *Mount) GetEffects() []EffectsEntry`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *Mount) GetEffectsOk() (*[]EffectsEntry, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *Mount) SetEffects(v []EffectsEntry)`

SetEffects sets Effects field to given value.

### HasEffects

`func (o *Mount) HasEffects() bool`

HasEffects returns a boolean if a field has been set.

### SetEffectsNil

`func (o *Mount) SetEffectsNil(b bool)`

 SetEffectsNil sets the value for Effects to be an explicit nil

### UnsetEffects
`func (o *Mount) UnsetEffects()`

UnsetEffects ensures that no value is present for Effects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


