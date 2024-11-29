# \GameAPI

All URIs are relative to *https://api.dofusdu.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGameSearch**](GameAPI.md#GetGameSearch) | **Get** /{game}/v1/{language}/search | Game Search
[**GetItemsAllSearch**](GameAPI.md#GetItemsAllSearch) | **Get** /{game}/v1/{language}/items/search | Search All Items



## GetGameSearch

> []GameSearch GetGameSearch(ctx, language, game).Query(query).FilterSearchIndex(filterSearchIndex).Limit(limit).FieldsItem(fieldsItem).FilterTypeNameId(filterTypeNameId).Execute()

Game Search



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
	query := "paztek" // string | search query
	filterSearchIndex := []string{"FilterSearchIndex_example"} // []string | only results with all specific type (optional)
	limit := int32(8) // int32 | maximum number of returned results (optional) (default to 8)
	fieldsItem := []string{"FieldsItem_example"} // []string | adds fields from the item search to the list entries if the hit is a item. Multiple comma separated values allowed. (optional)
	filterTypeNameId := []string{"Inner_example"} // []string | multi-filter results with the english item type name, including \"mount\" and \"set\" from filter[type]. Add with \"wood\" or \"+wood\" and exclude with \"-wood\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameAPI.GetGameSearch(context.Background(), language, game).Query(query).FilterSearchIndex(filterSearchIndex).Limit(limit).FieldsItem(fieldsItem).FilterTypeNameId(filterTypeNameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameAPI.GetGameSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGameSearch`: []GameSearch
	fmt.Fprintf(os.Stdout, "Response from `GameAPI.GetGameSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** | dofus3 | dofus3beta | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGameSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **string** | search query | 
 **filterSearchIndex** | **[]string** | only results with all specific type | 
 **limit** | **int32** | maximum number of returned results | [default to 8]
 **fieldsItem** | **[]string** | adds fields from the item search to the list entries if the hit is a item. Multiple comma separated values allowed. | 
 **filterTypeNameId** | **[]string** | multi-filter results with the english item type name, including \&quot;mount\&quot; and \&quot;set\&quot; from filter[type]. Add with \&quot;wood\&quot; or \&quot;+wood\&quot; and exclude with \&quot;-wood\&quot;. | 

### Return type

[**[]GameSearch**](GameSearch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsAllSearch

> []ListItemGeneral GetItemsAllSearch(ctx, language, game).Query(query).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).Limit(limit).FilterTypeNameId(filterTypeNameId).Execute()

Search All Items



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
	game := "dofus3" // string | dofus3 | dofus3beta
	query := "atcham" // string | case sensitive search query
	filterMinLevel := int32(190) // int32 | only results which level is equal or above this value (optional)
	filterMaxLevel := int32(200) // int32 | only results which level is equal or below this value (optional)
	limit := int32(8) // int32 | maximum number of returned results (optional) (default to 8)
	filterTypeNameId := []string{"Inner_example"} // []string | multi-filter results with the english type name. Add with \"wood\" or \"+wood\" and exclude with \"-wood\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameAPI.GetItemsAllSearch(context.Background(), language, game).Query(query).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).Limit(limit).FilterTypeNameId(filterTypeNameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameAPI.GetItemsAllSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsAllSearch`: []ListItemGeneral
	fmt.Fprintf(os.Stdout, "Response from `GameAPI.GetItemsAllSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** | dofus3 | dofus3beta | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsAllSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **string** | case sensitive search query | 
 **filterMinLevel** | **int32** | only results which level is equal or above this value | 
 **filterMaxLevel** | **int32** | only results which level is equal or below this value | 
 **limit** | **int32** | maximum number of returned results | [default to 8]
 **filterTypeNameId** | **[]string** | multi-filter results with the english type name. Add with \&quot;wood\&quot; or \&quot;+wood\&quot; and exclude with \&quot;-wood\&quot;. | 

### Return type

[**[]ListItemGeneral**](ListItemGeneral.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

