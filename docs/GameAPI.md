# \GameAPI

All URIs are relative to *https://api.dofusdu.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGameSearch**](GameAPI.md#GetGameSearch) | **Get** /{game}/{language}/search | Game Search
[**GetItemsAllSearch**](GameAPI.md#GetItemsAllSearch) | **Get** /{game}/{language}/items/search | Search All Items



## GetGameSearch

> []GetGameSearch200ResponseInner GetGameSearch(ctx, language, game).Query(query).FilterType(filterType).Limit(limit).FieldsItem(fieldsItem).Execute()

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
    game := "dofus2" // string | 
    query := "paztek" // string | search query
    filterType := []string{"FilterType_example"} // []string | only results with all specific type (optional)
    limit := int32(8) // int32 | maximum number of returned results (optional) (default to 8)
    fieldsItem := []string{"FieldsItem_example"} // []string | adds fields from the item search to the list entries if the hit is a item. Multiple comma separated values allowed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameAPI.GetGameSearch(context.Background(), language, game).Query(query).FilterType(filterType).Limit(limit).FieldsItem(fieldsItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameAPI.GetGameSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGameSearch`: []GetGameSearch200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `GameAPI.GetGameSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGameSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **string** | search query | 
 **filterType** | **[]string** | only results with all specific type | 
 **limit** | **int32** | maximum number of returned results | [default to 8]
 **fieldsItem** | **[]string** | adds fields from the item search to the list entries if the hit is a item. Multiple comma separated values allowed. | 

### Return type

[**[]GetGameSearch200ResponseInner**](GetGameSearch200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    openapiclient "github.com/dofusdude/dodugo"
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
    resp, r, err := apiClient.GameAPI.GetItemsAllSearch(context.Background(), language, game).Query(query).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameAPI.GetItemsAllSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemsAllSearch`: []ItemsListEntryTyped
    fmt.Fprintf(os.Stdout, "Response from `GameAPI.GetItemsAllSearch`: %v\n", resp)
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

