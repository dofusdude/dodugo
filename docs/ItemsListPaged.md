# ItemsListPaged

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**LinksPaged**](LinksPaged.md) |  | [optional] 
**Items** | Pointer to [**[]ItemListEntry**](ItemListEntry.md) |  | [optional] 

## Methods

### NewItemsListPaged

`func NewItemsListPaged() *ItemsListPaged`

NewItemsListPaged instantiates a new ItemsListPaged object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemsListPagedWithDefaults

`func NewItemsListPagedWithDefaults() *ItemsListPaged`

NewItemsListPagedWithDefaults instantiates a new ItemsListPaged object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ItemsListPaged) GetLinks() LinksPaged`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ItemsListPaged) GetLinksOk() (*LinksPaged, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ItemsListPaged) SetLinks(v LinksPaged)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ItemsListPaged) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetItems

`func (o *ItemsListPaged) GetItems() []ItemListEntry`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ItemsListPaged) GetItemsOk() (*[]ItemListEntry, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ItemsListPaged) SetItems(v []ItemListEntry)`

SetItems sets Items field to given value.

### HasItems

`func (o *ItemsListPaged) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


