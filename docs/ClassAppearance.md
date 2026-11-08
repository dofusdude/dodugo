# ClassAppearance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BonesId** | **int32** |  | 
**SkinIds** | **[]int32** |  | 
**ScalePercent** | [**ClassAppearanceScalePercent**](ClassAppearanceScalePercent.md) |  | 
**DefaultColors** | **[]int32** |  | 

## Methods

### NewClassAppearance

`func NewClassAppearance(bonesId int32, skinIds []int32, scalePercent ClassAppearanceScalePercent, defaultColors []int32, ) *ClassAppearance`

NewClassAppearance instantiates a new ClassAppearance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassAppearanceWithDefaults

`func NewClassAppearanceWithDefaults() *ClassAppearance`

NewClassAppearanceWithDefaults instantiates a new ClassAppearance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBonesId

`func (o *ClassAppearance) GetBonesId() int32`

GetBonesId returns the BonesId field if non-nil, zero value otherwise.

### GetBonesIdOk

`func (o *ClassAppearance) GetBonesIdOk() (*int32, bool)`

GetBonesIdOk returns a tuple with the BonesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonesId

`func (o *ClassAppearance) SetBonesId(v int32)`

SetBonesId sets BonesId field to given value.


### GetSkinIds

`func (o *ClassAppearance) GetSkinIds() []int32`

GetSkinIds returns the SkinIds field if non-nil, zero value otherwise.

### GetSkinIdsOk

`func (o *ClassAppearance) GetSkinIdsOk() (*[]int32, bool)`

GetSkinIdsOk returns a tuple with the SkinIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkinIds

`func (o *ClassAppearance) SetSkinIds(v []int32)`

SetSkinIds sets SkinIds field to given value.


### GetScalePercent

`func (o *ClassAppearance) GetScalePercent() ClassAppearanceScalePercent`

GetScalePercent returns the ScalePercent field if non-nil, zero value otherwise.

### GetScalePercentOk

`func (o *ClassAppearance) GetScalePercentOk() (*ClassAppearanceScalePercent, bool)`

GetScalePercentOk returns a tuple with the ScalePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalePercent

`func (o *ClassAppearance) SetScalePercent(v ClassAppearanceScalePercent)`

SetScalePercent sets ScalePercent field to given value.


### GetDefaultColors

`func (o *ClassAppearance) GetDefaultColors() []int32`

GetDefaultColors returns the DefaultColors field if non-nil, zero value otherwise.

### GetDefaultColorsOk

`func (o *ClassAppearance) GetDefaultColorsOk() (*[]int32, bool)`

GetDefaultColorsOk returns a tuple with the DefaultColors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultColors

`func (o *ClassAppearance) SetDefaultColors(v []int32)`

SetDefaultColors sets DefaultColors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


