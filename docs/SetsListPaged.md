# SetsListPaged

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**LinksPaged**](LinksPaged.md) |  | [optional] 
**Sets** | Pointer to [**[]SetListEntry**](SetListEntry.md) |  | [optional] 

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

### GetSets

`func (o *SetsListPaged) GetSets() []SetListEntry`

GetSets returns the Sets field if non-nil, zero value otherwise.

### GetSetsOk

`func (o *SetsListPaged) GetSetsOk() (*[]SetListEntry, bool)`

GetSetsOk returns a tuple with the Sets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSets

`func (o *SetsListPaged) SetSets(v []SetListEntry)`

SetSets sets Sets field to given value.

### HasSets

`func (o *SetsListPaged) HasSets() bool`

HasSets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


