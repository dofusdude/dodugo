# RecipeEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemAnkamaId** | Pointer to **int32** |  | [optional] 
**ItemSubtype** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 

## Methods

### NewRecipeEntry

`func NewRecipeEntry() *RecipeEntry`

NewRecipeEntry instantiates a new RecipeEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipeEntryWithDefaults

`func NewRecipeEntryWithDefaults() *RecipeEntry`

NewRecipeEntryWithDefaults instantiates a new RecipeEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemAnkamaId

`func (o *RecipeEntry) GetItemAnkamaId() int32`

GetItemAnkamaId returns the ItemAnkamaId field if non-nil, zero value otherwise.

### GetItemAnkamaIdOk

`func (o *RecipeEntry) GetItemAnkamaIdOk() (*int32, bool)`

GetItemAnkamaIdOk returns a tuple with the ItemAnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemAnkamaId

`func (o *RecipeEntry) SetItemAnkamaId(v int32)`

SetItemAnkamaId sets ItemAnkamaId field to given value.

### HasItemAnkamaId

`func (o *RecipeEntry) HasItemAnkamaId() bool`

HasItemAnkamaId returns a boolean if a field has been set.

### GetItemSubtype

`func (o *RecipeEntry) GetItemSubtype() string`

GetItemSubtype returns the ItemSubtype field if non-nil, zero value otherwise.

### GetItemSubtypeOk

`func (o *RecipeEntry) GetItemSubtypeOk() (*string, bool)`

GetItemSubtypeOk returns a tuple with the ItemSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSubtype

`func (o *RecipeEntry) SetItemSubtype(v string)`

SetItemSubtype sets ItemSubtype field to given value.

### HasItemSubtype

`func (o *RecipeEntry) HasItemSubtype() bool`

HasItemSubtype returns a boolean if a field has been set.

### GetQuantity

`func (o *RecipeEntry) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *RecipeEntry) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *RecipeEntry) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *RecipeEntry) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


