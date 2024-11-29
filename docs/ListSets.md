# ListSets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PagedLinks**](PagedLinks.md) |  | [optional] 
**Sets** | Pointer to [**[]ListSet**](ListSet.md) |  | [optional] 

## Methods

### NewListSets

`func NewListSets() *ListSets`

NewListSets instantiates a new ListSets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSetsWithDefaults

`func NewListSetsWithDefaults() *ListSets`

NewListSetsWithDefaults instantiates a new ListSets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ListSets) GetLinks() PagedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListSets) GetLinksOk() (*PagedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListSets) SetLinks(v PagedLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListSets) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetSets

`func (o *ListSets) GetSets() []ListSet`

GetSets returns the Sets field if non-nil, zero value otherwise.

### GetSetsOk

`func (o *ListSets) GetSetsOk() (*[]ListSet, bool)`

GetSetsOk returns a tuple with the Sets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSets

`func (o *ListSets) SetSets(v []ListSet)`

SetSets sets Sets field to given value.

### HasSets

`func (o *ListSets) HasSets() bool`

HasSets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


