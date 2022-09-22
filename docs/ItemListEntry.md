# ItemListEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ItemsListEntryTypedType**](ItemsListEntryTypedType.md) |  | [optional] 
**Level** | Pointer to **int32** |  | [optional] 
**ImageUrls** | Pointer to [**ImageUrls**](ImageUrls.md) |  | [optional] 
**HasRecipe** | Pointer to **NullableBool** |  | [optional] 
**Recipe** | Pointer to [**[]RecipeEntry**](RecipeEntry.md) |  | [optional] 

## Methods

### NewItemListEntry

`func NewItemListEntry() *ItemListEntry`

NewItemListEntry instantiates a new ItemListEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemListEntryWithDefaults

`func NewItemListEntryWithDefaults() *ItemListEntry`

NewItemListEntryWithDefaults instantiates a new ItemListEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *ItemListEntry) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *ItemListEntry) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *ItemListEntry) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.

### HasAnkamaId

`func (o *ItemListEntry) HasAnkamaId() bool`

HasAnkamaId returns a boolean if a field has been set.

### GetName

`func (o *ItemListEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ItemListEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ItemListEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ItemListEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ItemListEntry) GetType() ItemsListEntryTypedType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ItemListEntry) GetTypeOk() (*ItemsListEntryTypedType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ItemListEntry) SetType(v ItemsListEntryTypedType)`

SetType sets Type field to given value.

### HasType

`func (o *ItemListEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLevel

`func (o *ItemListEntry) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ItemListEntry) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ItemListEntry) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ItemListEntry) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetImageUrls

`func (o *ItemListEntry) GetImageUrls() ImageUrls`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *ItemListEntry) GetImageUrlsOk() (*ImageUrls, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *ItemListEntry) SetImageUrls(v ImageUrls)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *ItemListEntry) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.

### GetHasRecipe

`func (o *ItemListEntry) GetHasRecipe() bool`

GetHasRecipe returns the HasRecipe field if non-nil, zero value otherwise.

### GetHasRecipeOk

`func (o *ItemListEntry) GetHasRecipeOk() (*bool, bool)`

GetHasRecipeOk returns a tuple with the HasRecipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRecipe

`func (o *ItemListEntry) SetHasRecipe(v bool)`

SetHasRecipe sets HasRecipe field to given value.

### HasHasRecipe

`func (o *ItemListEntry) HasHasRecipe() bool`

HasHasRecipe returns a boolean if a field has been set.

### SetHasRecipeNil

`func (o *ItemListEntry) SetHasRecipeNil(b bool)`

 SetHasRecipeNil sets the value for HasRecipe to be an explicit nil

### UnsetHasRecipe
`func (o *ItemListEntry) UnsetHasRecipe()`

UnsetHasRecipe ensures that no value is present for HasRecipe, not even an explicit nil
### GetRecipe

`func (o *ItemListEntry) GetRecipe() []RecipeEntry`

GetRecipe returns the Recipe field if non-nil, zero value otherwise.

### GetRecipeOk

`func (o *ItemListEntry) GetRecipeOk() (*[]RecipeEntry, bool)`

GetRecipeOk returns a tuple with the Recipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipe

`func (o *ItemListEntry) SetRecipe(v []RecipeEntry)`

SetRecipe sets Recipe field to given value.

### HasRecipe

`func (o *ItemListEntry) HasRecipe() bool`

HasRecipe returns a boolean if a field has been set.

### SetRecipeNil

`func (o *ItemListEntry) SetRecipeNil(b bool)`

 SetRecipeNil sets the value for Recipe to be an explicit nil

### UnsetRecipe
`func (o *ItemListEntry) UnsetRecipe()`

UnsetRecipe ensures that no value is present for Recipe, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


