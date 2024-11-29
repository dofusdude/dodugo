# GameSearchItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**GameSearchType**](GameSearchType.md) |  | [optional] 
**Level** | Pointer to **NullableInt32** |  | [optional] 
**ImageUrls** | Pointer to [**Images**](Images.md) |  | [optional] 

## Methods

### NewGameSearchItem

`func NewGameSearchItem() *GameSearchItem`

NewGameSearchItem instantiates a new GameSearchItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameSearchItemWithDefaults

`func NewGameSearchItemWithDefaults() *GameSearchItem`

NewGameSearchItemWithDefaults instantiates a new GameSearchItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameSearchItem) GetType() GameSearchType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameSearchItem) GetTypeOk() (*GameSearchType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameSearchItem) SetType(v GameSearchType)`

SetType sets Type field to given value.

### HasType

`func (o *GameSearchItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLevel

`func (o *GameSearchItem) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *GameSearchItem) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *GameSearchItem) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *GameSearchItem) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### SetLevelNil

`func (o *GameSearchItem) SetLevelNil(b bool)`

 SetLevelNil sets the value for Level to be an explicit nil

### UnsetLevel
`func (o *GameSearchItem) UnsetLevel()`

UnsetLevel ensures that no value is present for Level, not even an explicit nil
### GetImageUrls

`func (o *GameSearchItem) GetImageUrls() Images`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *GameSearchItem) GetImageUrlsOk() (*Images, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *GameSearchItem) SetImageUrls(v Images)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *GameSearchItem) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


