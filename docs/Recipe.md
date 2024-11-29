# Recipe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemAnkamaId** | Pointer to **int32** |  | [optional] 
**ItemSubtype** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 

## Methods

### NewRecipe

`func NewRecipe() *Recipe`

NewRecipe instantiates a new Recipe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipeWithDefaults

`func NewRecipeWithDefaults() *Recipe`

NewRecipeWithDefaults instantiates a new Recipe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemAnkamaId

`func (o *Recipe) GetItemAnkamaId() int32`

GetItemAnkamaId returns the ItemAnkamaId field if non-nil, zero value otherwise.

### GetItemAnkamaIdOk

`func (o *Recipe) GetItemAnkamaIdOk() (*int32, bool)`

GetItemAnkamaIdOk returns a tuple with the ItemAnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemAnkamaId

`func (o *Recipe) SetItemAnkamaId(v int32)`

SetItemAnkamaId sets ItemAnkamaId field to given value.

### HasItemAnkamaId

`func (o *Recipe) HasItemAnkamaId() bool`

HasItemAnkamaId returns a boolean if a field has been set.

### GetItemSubtype

`func (o *Recipe) GetItemSubtype() string`

GetItemSubtype returns the ItemSubtype field if non-nil, zero value otherwise.

### GetItemSubtypeOk

`func (o *Recipe) GetItemSubtypeOk() (*string, bool)`

GetItemSubtypeOk returns a tuple with the ItemSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSubtype

`func (o *Recipe) SetItemSubtype(v string)`

SetItemSubtype sets ItemSubtype field to given value.

### HasItemSubtype

`func (o *Recipe) HasItemSubtype() bool`

HasItemSubtype returns a boolean if a field has been set.

### GetQuantity

`func (o *Recipe) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Recipe) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Recipe) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Recipe) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


