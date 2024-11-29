# ListMounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PagedLinks**](PagedLinks.md) |  | [optional] 
**Items** | Pointer to [**[]Mount**](Mount.md) |  | [optional] 

## Methods

### NewListMounts

`func NewListMounts() *ListMounts`

NewListMounts instantiates a new ListMounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMountsWithDefaults

`func NewListMountsWithDefaults() *ListMounts`

NewListMountsWithDefaults instantiates a new ListMounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ListMounts) GetLinks() PagedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListMounts) GetLinksOk() (*PagedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListMounts) SetLinks(v PagedLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListMounts) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetItems

`func (o *ListMounts) GetItems() []Mount`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListMounts) GetItemsOk() (*[]Mount, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListMounts) SetItems(v []Mount)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListMounts) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


