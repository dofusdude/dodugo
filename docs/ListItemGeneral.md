# ListItemGeneral

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**TranslatedId**](TranslatedId.md) |  | [optional] 
**ItemSubtype** | Pointer to [**ItemSubtype**](ItemSubtype.md) |  | [optional] 
**Level** | Pointer to **int32** |  | [optional] 
**ImageUrls** | Pointer to [**Images**](Images.md) |  | [optional] 

## Methods

### NewListItemGeneral

`func NewListItemGeneral() *ListItemGeneral`

NewListItemGeneral instantiates a new ListItemGeneral object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListItemGeneralWithDefaults

`func NewListItemGeneralWithDefaults() *ListItemGeneral`

NewListItemGeneralWithDefaults instantiates a new ListItemGeneral object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *ListItemGeneral) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *ListItemGeneral) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *ListItemGeneral) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.

### HasAnkamaId

`func (o *ListItemGeneral) HasAnkamaId() bool`

HasAnkamaId returns a boolean if a field has been set.

### GetName

`func (o *ListItemGeneral) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListItemGeneral) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListItemGeneral) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListItemGeneral) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ListItemGeneral) GetType() TranslatedId`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListItemGeneral) GetTypeOk() (*TranslatedId, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListItemGeneral) SetType(v TranslatedId)`

SetType sets Type field to given value.

### HasType

`func (o *ListItemGeneral) HasType() bool`

HasType returns a boolean if a field has been set.

### GetItemSubtype

`func (o *ListItemGeneral) GetItemSubtype() ItemSubtype`

GetItemSubtype returns the ItemSubtype field if non-nil, zero value otherwise.

### GetItemSubtypeOk

`func (o *ListItemGeneral) GetItemSubtypeOk() (*ItemSubtype, bool)`

GetItemSubtypeOk returns a tuple with the ItemSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSubtype

`func (o *ListItemGeneral) SetItemSubtype(v ItemSubtype)`

SetItemSubtype sets ItemSubtype field to given value.

### HasItemSubtype

`func (o *ListItemGeneral) HasItemSubtype() bool`

HasItemSubtype returns a boolean if a field has been set.

### GetLevel

`func (o *ListItemGeneral) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ListItemGeneral) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ListItemGeneral) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ListItemGeneral) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetImageUrls

`func (o *ListItemGeneral) GetImageUrls() Images`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *ListItemGeneral) GetImageUrlsOk() (*Images, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *ListItemGeneral) SetImageUrls(v Images)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *ListItemGeneral) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


