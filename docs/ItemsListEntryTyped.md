# ItemsListEntryTyped

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ItemsListEntryTypedType**](ItemsListEntryTypedType.md) |  | [optional] 
**ItemSubtype** | Pointer to **string** | The API item category. Can be used to build the request URL for this particular item. Always english. | [optional] 
**Level** | Pointer to **int32** |  | [optional] 
**ImageUrls** | Pointer to [**ImageUrls**](ImageUrls.md) |  | [optional] 

## Methods

### NewItemsListEntryTyped

`func NewItemsListEntryTyped() *ItemsListEntryTyped`

NewItemsListEntryTyped instantiates a new ItemsListEntryTyped object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemsListEntryTypedWithDefaults

`func NewItemsListEntryTypedWithDefaults() *ItemsListEntryTyped`

NewItemsListEntryTypedWithDefaults instantiates a new ItemsListEntryTyped object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *ItemsListEntryTyped) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *ItemsListEntryTyped) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *ItemsListEntryTyped) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.

### HasAnkamaId

`func (o *ItemsListEntryTyped) HasAnkamaId() bool`

HasAnkamaId returns a boolean if a field has been set.

### GetName

`func (o *ItemsListEntryTyped) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ItemsListEntryTyped) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ItemsListEntryTyped) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ItemsListEntryTyped) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ItemsListEntryTyped) GetType() ItemsListEntryTypedType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ItemsListEntryTyped) GetTypeOk() (*ItemsListEntryTypedType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ItemsListEntryTyped) SetType(v ItemsListEntryTypedType)`

SetType sets Type field to given value.

### HasType

`func (o *ItemsListEntryTyped) HasType() bool`

HasType returns a boolean if a field has been set.

### GetItemSubtype

`func (o *ItemsListEntryTyped) GetItemSubtype() string`

GetItemSubtype returns the ItemSubtype field if non-nil, zero value otherwise.

### GetItemSubtypeOk

`func (o *ItemsListEntryTyped) GetItemSubtypeOk() (*string, bool)`

GetItemSubtypeOk returns a tuple with the ItemSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSubtype

`func (o *ItemsListEntryTyped) SetItemSubtype(v string)`

SetItemSubtype sets ItemSubtype field to given value.

### HasItemSubtype

`func (o *ItemsListEntryTyped) HasItemSubtype() bool`

HasItemSubtype returns a boolean if a field has been set.

### GetLevel

`func (o *ItemsListEntryTyped) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ItemsListEntryTyped) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ItemsListEntryTyped) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ItemsListEntryTyped) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetImageUrls

`func (o *ItemsListEntryTyped) GetImageUrls() ImageUrls`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *ItemsListEntryTyped) GetImageUrlsOk() (*ImageUrls, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *ItemsListEntryTyped) SetImageUrls(v ImageUrls)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *ItemsListEntryTyped) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


