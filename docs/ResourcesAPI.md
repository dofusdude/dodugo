# \ResourcesAPI

All URIs are relative to *https://api.dofusdu.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllItemsResourcesList**](ResourcesAPI.md#GetAllItemsResourcesList) | **Get** /{game}/v1/{language}/items/resources/all | List All Resources
[**GetItemsResourceSearch**](ResourcesAPI.md#GetItemsResourceSearch) | **Get** /{game}/v1/{language}/items/resources/search | Search Resources
[**GetItemsResourcesList**](ResourcesAPI.md#GetItemsResourcesList) | **Get** /{game}/v1/{language}/items/resources | List Resources
[**GetItemsResourcesSingle**](ResourcesAPI.md#GetItemsResourcesSingle) | **Get** /{game}/v1/{language}/items/resources/{ankama_id} | Single Resources



## GetAllItemsResourcesList

> ListItems GetAllItemsResourcesList(ctx, language, game).SortLevel(sortLevel).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).AcceptEncoding(acceptEncoding).FilterTypeNameId(filterTypeNameId).Execute()

List All Resources



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
	game := "dofus3" // string | dofus3 | dofus3beta
	sortLevel := "desc" // string | sort the resulting list by level, default unsorted (optional)
	filterMinLevel := int32(160) // int32 | only results which level is equal or above this value (optional)
	filterMaxLevel := int32(190) // int32 | only results which level is equal or below this value (optional)
	acceptEncoding := "acceptEncoding_example" // string | optional compression for saving bandwidth (optional)
	filterTypeNameId := []string{"Inner_example"} // []string | multi-filter results with the english type name. Add with \"wood\" or \"+wood\" and exclude with \"-wood\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.GetAllItemsResourcesList(context.Background(), language, game).SortLevel(sortLevel).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).AcceptEncoding(acceptEncoding).FilterTypeNameId(filterTypeNameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.GetAllItemsResourcesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllItemsResourcesList`: ListItems
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.GetAllItemsResourcesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** | dofus3 | dofus3beta | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllItemsResourcesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortLevel** | **string** | sort the resulting list by level, default unsorted | 
 **filterMinLevel** | **int32** | only results which level is equal or above this value | 
 **filterMaxLevel** | **int32** | only results which level is equal or below this value | 
 **acceptEncoding** | **string** | optional compression for saving bandwidth | 
 **filterTypeNameId** | **[]string** | multi-filter results with the english type name. Add with \&quot;wood\&quot; or \&quot;+wood\&quot; and exclude with \&quot;-wood\&quot;. | 

### Return type

[**ListItems**](ListItems.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsResourceSearch

> []ListItem GetItemsResourceSearch(ctx, language, game).Query(query).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).Limit(limit).FilterTypeNameId(filterTypeNameId).Execute()

Search Resources



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
	game := "dofus3" // string | dofus3 | dofus3beta
	query := "snowdew" // string | case sensitive search query
	filterMinLevel := int32(150) // int32 | only results which level is equal or above this value (optional)
	filterMaxLevel := int32(200) // int32 | only results which level is equal or below this value (optional)
	limit := int32(8) // int32 | maximum number of returned results (optional) (default to 8)
	filterTypeNameId := []string{"Inner_example"} // []string | multi-filter results with the english type name. Add with \"wood\" or \"+wood\" and exclude with \"-wood\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.GetItemsResourceSearch(context.Background(), language, game).Query(query).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).Limit(limit).FilterTypeNameId(filterTypeNameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.GetItemsResourceSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsResourceSearch`: []ListItem
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.GetItemsResourceSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** | dofus3 | dofus3beta | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsResourceSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **string** | case sensitive search query | 
 **filterMinLevel** | **int32** | only results which level is equal or above this value | 
 **filterMaxLevel** | **int32** | only results which level is equal or below this value | 
 **limit** | **int32** | maximum number of returned results | [default to 8]
 **filterTypeNameId** | **[]string** | multi-filter results with the english type name. Add with \&quot;wood\&quot; or \&quot;+wood\&quot; and exclude with \&quot;-wood\&quot;. | 

### Return type

[**[]ListItem**](ListItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsResourcesList

> ListItems GetItemsResourcesList(ctx, language, game).SortLevel(sortLevel).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).PageSize(pageSize).PageNumber(pageNumber).FieldsItem(fieldsItem).FilterTypeNameId(filterTypeNameId).Execute()

List Resources



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
	game := "dofus3" // string | dofus3 | dofus3beta
	sortLevel := "desc" // string | sort the resulting list by level, default unsorted (optional)
	filterMinLevel := int32(160) // int32 | only results which level is equal or above this value (optional)
	filterMaxLevel := int32(190) // int32 | only results which level is equal or below this value (optional)
	pageSize := int32(10) // int32 | size of the results from the list. -1 disables pagination and gets all in one response. (optional)
	pageNumber := int32(1) // int32 | page number based on the current page[size]. So you could get page 1 with 8 entrys and page 2 would have entries 8 to 16. (optional)
	fieldsItem := []string{"FieldsItem_example"} // []string | adds fields from their detail endpoint to the item list entries. Multiple comma separated values allowed. (optional)
	filterTypeNameId := []string{"Inner_example"} // []string | multi-filter results with the english type name. Add with \"wood\" or \"+wood\" and exclude with \"-wood\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.GetItemsResourcesList(context.Background(), language, game).SortLevel(sortLevel).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).PageSize(pageSize).PageNumber(pageNumber).FieldsItem(fieldsItem).FilterTypeNameId(filterTypeNameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.GetItemsResourcesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsResourcesList`: ListItems
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.GetItemsResourcesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** | dofus3 | dofus3beta | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsResourcesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortLevel** | **string** | sort the resulting list by level, default unsorted | 
 **filterMinLevel** | **int32** | only results which level is equal or above this value | 
 **filterMaxLevel** | **int32** | only results which level is equal or below this value | 
 **pageSize** | **int32** | size of the results from the list. -1 disables pagination and gets all in one response. | 
 **pageNumber** | **int32** | page number based on the current page[size]. So you could get page 1 with 8 entrys and page 2 would have entries 8 to 16. | 
 **fieldsItem** | **[]string** | adds fields from their detail endpoint to the item list entries. Multiple comma separated values allowed. | 
 **filterTypeNameId** | **[]string** | multi-filter results with the english type name. Add with \&quot;wood\&quot; or \&quot;+wood\&quot; and exclude with \&quot;-wood\&quot;. | 

### Return type

[**ListItems**](ListItems.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsResourcesSingle

> Resource GetItemsResourcesSingle(ctx, language, ankamaId, game).Execute()

Single Resources



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
	ankamaId := int32(7295) // int32 | identifier
	game := "dofus3" // string | dofus3 | dofus3beta

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.GetItemsResourcesSingle(context.Background(), language, ankamaId, game).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.GetItemsResourcesSingle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsResourcesSingle`: Resource
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.GetItemsResourcesSingle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**ankamaId** | **int32** | identifier | 
**game** | **string** | dofus3 | dofus3beta | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsResourcesSingleRequest struct via the builder pattern


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

