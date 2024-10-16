# \QuestItemsAPI

All URIs are relative to *https://api.dofusdu.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllItemsQuestList**](QuestItemsAPI.md#GetAllItemsQuestList) | **Get** /{game}/{language}/items/quest/all | List All Quest Items
[**GetItemQuestSingle**](QuestItemsAPI.md#GetItemQuestSingle) | **Get** /{game}/{language}/items/quest/{ankama_id} | Single Quest Items
[**GetItemsQuestList**](QuestItemsAPI.md#GetItemsQuestList) | **Get** /{game}/{language}/items/quest | List Quest Items
[**GetItemsQuestSearch**](QuestItemsAPI.md#GetItemsQuestSearch) | **Get** /{game}/{language}/items/quest/search | Search Quest Items



## GetAllItemsQuestList

> ItemsListPaged GetAllItemsQuestList(ctx, language, game).SortLevel(sortLevel).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).AcceptEncoding(acceptEncoding).FilterTypeEnum(filterTypeEnum).Execute()

List All Quest Items



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dofusdude/dodugo"
)

func main() {
	language := "fr" // string | a valid language code
	game := "dofus2" // string | 
	sortLevel := "desc" // string | sort the resulting list by level, default unsorted (optional)
	filterTypeName := "Sufokia" // string | only results with the translated type name (optional)
	filterMinLevel := int32(1) // int32 | only results which level is equal or above this value (optional)
	filterMaxLevel := int32(50) // int32 | only results which level is equal or below this value (optional)
	acceptEncoding := "acceptEncoding_example" // string | optional compression for saving bandwidth (optional)
	filterTypeEnum := []string{"Inner_example"} // []string | multi-filter results with the english type name. Add with \"wood\" or \"+wood\" and exclude with \"-wood\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuestItemsAPI.GetAllItemsQuestList(context.Background(), language, game).SortLevel(sortLevel).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).AcceptEncoding(acceptEncoding).FilterTypeEnum(filterTypeEnum).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuestItemsAPI.GetAllItemsQuestList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllItemsQuestList`: ItemsListPaged
	fmt.Fprintf(os.Stdout, "Response from `QuestItemsAPI.GetAllItemsQuestList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllItemsQuestListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortLevel** | **string** | sort the resulting list by level, default unsorted | 
 **filterTypeName** | **string** | only results with the translated type name | 
 **filterMinLevel** | **int32** | only results which level is equal or above this value | 
 **filterMaxLevel** | **int32** | only results which level is equal or below this value | 
 **acceptEncoding** | **string** | optional compression for saving bandwidth | 
 **filterTypeEnum** | **[]string** | multi-filter results with the english type name. Add with \&quot;wood\&quot; or \&quot;+wood\&quot; and exclude with \&quot;-wood\&quot;. | 

### Return type

[**ItemsListPaged**](ItemsListPaged.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemQuestSingle

> Resource GetItemQuestSingle(ctx, language, ankamaId, game).Execute()

Single Quest Items



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dofusdude/dodugo"
)

func main() {
	language := "language_example" // string | a valid language code
	ankamaId := int32(25256) // int32 | identifier
	game := "dofus2" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuestItemsAPI.GetItemQuestSingle(context.Background(), language, ankamaId, game).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuestItemsAPI.GetItemQuestSingle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemQuestSingle`: Resource
	fmt.Fprintf(os.Stdout, "Response from `QuestItemsAPI.GetItemQuestSingle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**ankamaId** | **int32** | identifier | 
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemQuestSingleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Resource**](Resource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsQuestList

> ItemsListPaged GetItemsQuestList(ctx, language, game).SortLevel(sortLevel).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).PageSize(pageSize).PageNumber(pageNumber).FieldsItem(fieldsItem).FilterTypeEnum(filterTypeEnum).Execute()

List Quest Items



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dofusdude/dodugo"
)

func main() {
	language := "fr" // string | a valid language code
	game := "dofus2" // string | 
	sortLevel := "desc" // string | sort the resulting list by level, default unsorted (optional)
	filterTypeName := "Sufokia" // string | only results with the translated type name (optional)
	filterMinLevel := int32(1) // int32 | only results which level is equal or above this value (optional)
	filterMaxLevel := int32(50) // int32 | only results which level is equal or below this value (optional)
	pageSize := int32(5) // int32 | size of the results from the list. -1 disables pagination and gets all in one response. (optional)
	pageNumber := int32(1) // int32 | page number based on the current page[size]. So you could get page 1 with 8 entrys and page 2 would have entries 8 to 16. (optional)
	fieldsItem := []string{"FieldsItem_example"} // []string | adds fields from their detail endpoint to the item list entries. Multiple comma separated values allowed. (optional)
	filterTypeEnum := []string{"Inner_example"} // []string | multi-filter results with the english type name. Add with \"wood\" or \"+wood\" and exclude with \"-wood\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuestItemsAPI.GetItemsQuestList(context.Background(), language, game).SortLevel(sortLevel).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).PageSize(pageSize).PageNumber(pageNumber).FieldsItem(fieldsItem).FilterTypeEnum(filterTypeEnum).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuestItemsAPI.GetItemsQuestList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsQuestList`: ItemsListPaged
	fmt.Fprintf(os.Stdout, "Response from `QuestItemsAPI.GetItemsQuestList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsQuestListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortLevel** | **string** | sort the resulting list by level, default unsorted | 
 **filterTypeName** | **string** | only results with the translated type name | 
 **filterMinLevel** | **int32** | only results which level is equal or above this value | 
 **filterMaxLevel** | **int32** | only results which level is equal or below this value | 
 **pageSize** | **int32** | size of the results from the list. -1 disables pagination and gets all in one response. | 
 **pageNumber** | **int32** | page number based on the current page[size]. So you could get page 1 with 8 entrys and page 2 would have entries 8 to 16. | 
 **fieldsItem** | **[]string** | adds fields from their detail endpoint to the item list entries. Multiple comma separated values allowed. | 
 **filterTypeEnum** | **[]string** | multi-filter results with the english type name. Add with \&quot;wood\&quot; or \&quot;+wood\&quot; and exclude with \&quot;-wood\&quot;. | 

### Return type

[**ItemsListPaged**](ItemsListPaged.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsQuestSearch

> []ItemListEntry GetItemsQuestSearch(ctx, language, game).Query(query).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).Limit(limit).FilterTypeEnum(filterTypeEnum).Execute()

Search Quest Items



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dofusdude/dodugo"
)

func main() {
	language := "es" // string | a valid language code
	game := "dofus2" // string | 
	query := "Ficha" // string | case sensitive search query
	filterTypeName := "Justicieros" // string | only results with the translated type name (optional)
	filterMinLevel := int32(60) // int32 | only results which level is equal or above this value (optional)
	filterMaxLevel := int32(70) // int32 | only results which level is equal or below this value (optional)
	limit := int32(8) // int32 | maximum number of returned results (optional) (default to 8)
	filterTypeEnum := []string{"Inner_example"} // []string | multi-filter results with the english type name. Add with \"wood\" or \"+wood\" and exclude with \"-wood\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuestItemsAPI.GetItemsQuestSearch(context.Background(), language, game).Query(query).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).Limit(limit).FilterTypeEnum(filterTypeEnum).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuestItemsAPI.GetItemsQuestSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsQuestSearch`: []ItemListEntry
	fmt.Fprintf(os.Stdout, "Response from `QuestItemsAPI.GetItemsQuestSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsQuestSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **string** | case sensitive search query | 
 **filterTypeName** | **string** | only results with the translated type name | 
 **filterMinLevel** | **int32** | only results which level is equal or above this value | 
 **filterMaxLevel** | **int32** | only results which level is equal or below this value | 
 **limit** | **int32** | maximum number of returned results | [default to 8]
 **filterTypeEnum** | **[]string** | multi-filter results with the english type name. Add with \&quot;wood\&quot; or \&quot;+wood\&quot; and exclude with \&quot;-wood\&quot;. | 

### Return type

[**[]ItemListEntry**](ItemListEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

