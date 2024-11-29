# GameSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**GameSearchType**](GameSearchType.md) |  | [optional] 
**ItemFields** | Pointer to [**GameSearchItem**](GameSearchItem.md) |  | [optional] 

## Methods

### NewGameSearch

`func NewGameSearch() *GameSearch`

NewGameSearch instantiates a new GameSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameSearchWithDefaults

`func NewGameSearchWithDefaults() *GameSearch`

NewGameSearchWithDefaults instantiates a new GameSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *GameSearch) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *GameSearch) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *GameSearch) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.

### HasAnkamaId

`func (o *GameSearch) HasAnkamaId() bool`

HasAnkamaId returns a boolean if a field has been set.

### GetName

`func (o *GameSearch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GameSearch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GameSearch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GameSearch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *GameSearch) GetType() GameSearchType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameSearch) GetTypeOk() (*GameSearchType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameSearch) SetType(v GameSearchType)`

SetType sets Type field to given value.

### HasType

`func (o *GameSearch) HasType() bool`

HasType returns a boolean if a field has been set.

### GetItemFields

`func (o *GameSearch) GetItemFields() GameSearchItem`

GetItemFields returns the ItemFields field if non-nil, zero value otherwise.

### GetItemFieldsOk

`func (o *GameSearch) GetItemFieldsOk() (*GameSearchItem, bool)`

GetItemFieldsOk returns a tuple with the ItemFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemFields

`func (o *GameSearch) SetItemFields(v GameSearchItem)`

SetItemFields sets ItemFields field to given value.

### HasItemFields

`func (o *GameSearch) HasItemFields() bool`

HasItemFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


