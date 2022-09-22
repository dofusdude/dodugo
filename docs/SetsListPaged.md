# SetsListPaged

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**LinksPaged**](LinksPaged.md) |  | [optional] 
**Items** | Pointer to [**[]SetListEntry**](SetListEntry.md) |  | [optional] 

## Methods

### NewSetsListPaged

`func NewSetsListPaged() *SetsListPaged`

NewSetsListPaged instantiates a new SetsListPaged object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetsListPagedWithDefaults

`func NewSetsListPagedWithDefaults() *SetsListPaged`

NewSetsListPagedWithDefaults instantiates a new SetsListPaged object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SetsListPaged) GetLinks() LinksPaged`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SetsListPaged) GetLinksOk() (*LinksPaged, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SetsListPaged) SetLinks(v LinksPaged)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SetsListPaged) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetItems

`func (o *SetsListPaged) GetItems() []SetListEntry`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SetsListPaged) GetItemsOk() (*[]SetListEntry, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SetsListPaged) SetItems(v []SetListEntry)`

SetItems sets Items field to given value.

### HasItems

`func (o *SetsListPaged) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


