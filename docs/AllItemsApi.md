# \AllItemsApi

All URIs are relative to *https://api.dofusdu.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetItemsAllSearch**](AllItemsApi.md#GetItemsAllSearch) | **Get** /{game}/{language}/items/search | Search All Items



## GetItemsAllSearch

> []ItemsListEntryTyped GetItemsAllSearch(ctx, language, game).Query(query).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).Limit(limit).Execute()

Search All Items



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    language := "fr" // string | a valid language code
    game := "dofus2" // string | 
    query := "atcham" // string | case sensitive search query
    filterTypeName := "Bottes" // string | only results with the translated type name across all item_subtypes (optional)
    filterMinLevel := int32(190) // int32 | only results which level is equal or above this value (optional)
    filterMaxLevel := int32(200) // int32 | only results which level is equal or below this value (optional)
    limit := int32(8) // int32 | maximum number of returned results (optional) (default to 8)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllItemsApi.GetItemsAllSearch(context.Background(), language, game).Query(query).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllItemsApi.GetItemsAllSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemsAllSearch`: []ItemsListEntryTyped
    fmt.Fprintf(os.Stdout, "Response from `AllItemsApi.GetItemsAllSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsAllSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **string** | case sensitive search query | 
 **filterTypeName** | **string** | only results with the translated type name across all item_subtypes | 
 **filterMinLevel** | **int32** | only results which level is equal or above this value | 
 **filterMaxLevel** | **int32** | only results which level is equal or below this value | 
 **limit** | **int32** | maximum number of returned results | [default to 8]

### Return type

[**[]ItemsListEntryTyped**](ItemsListEntryTyped.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

