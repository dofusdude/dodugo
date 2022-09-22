# \CosmeticsApi

All URIs are relative to *https://api.dofusdu.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCosmeticsList**](CosmeticsApi.md#GetCosmeticsList) | **Get** /{game}/{language}/items/cosmetics | List Cosmetics
[**GetCosmeticsSearch**](CosmeticsApi.md#GetCosmeticsSearch) | **Get** /{game}/{language}/items/cosmetics/search | Search Cosmetics
[**GetCosmeticsSingle**](CosmeticsApi.md#GetCosmeticsSingle) | **Get** /{game}/{language}/items/cosmetics/{ankama_id} | Single Cosmetics



## GetCosmeticsList

> ItemsListPaged GetCosmeticsList(ctx, language, game).SortLevel(sortLevel).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).PageSize(pageSize).PageNumber(pageNumber).FieldsItem(fieldsItem).Execute()

List Cosmetics



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
    language := "language_example" // string | a valid language code
    game := "dofus2" // string | 
    sortLevel := "sortLevel_example" // string | sort the resulting list by level, default unsorted (optional)
    filterTypeName := "filterTypeName_example" // string | only results with the translated type name (optional)
    filterMinLevel := int32(56) // int32 | only results which level is equal or above this value (optional)
    filterMaxLevel := int32(56) // int32 | only results which level is equal or below this value (optional)
    pageSize := int32(56) // int32 | size of the results from the list. -1 disables pagination and gets all in one response. (optional)
    pageNumber := int32(56) // int32 | page number based on the current page[size]. So you could get page 1 with 8 entrys and page 2 would have entries 8 to 16. (optional)
    fieldsItem := "recipe" // string | adds fields from their detail endpoint to the item list entries. Multiple comma separated values allowed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CosmeticsApi.GetCosmeticsList(context.Background(), language, game).SortLevel(sortLevel).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).PageSize(pageSize).PageNumber(pageNumber).FieldsItem(fieldsItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CosmeticsApi.GetCosmeticsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCosmeticsList`: ItemsListPaged
    fmt.Fprintf(os.Stdout, "Response from `CosmeticsApi.GetCosmeticsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCosmeticsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortLevel** | **string** | sort the resulting list by level, default unsorted | 
 **filterTypeName** | **string** | only results with the translated type name | 
 **filterMinLevel** | **int32** | only results which level is equal or above this value | 
 **filterMaxLevel** | **int32** | only results which level is equal or below this value | 
 **pageSize** | **int32** | size of the results from the list. -1 disables pagination and gets all in one response. | 
 **pageNumber** | **int32** | page number based on the current page[size]. So you could get page 1 with 8 entrys and page 2 would have entries 8 to 16. | 
 **fieldsItem** | **string** | adds fields from their detail endpoint to the item list entries. Multiple comma separated values allowed. | 

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


## GetCosmeticsSearch

> []ItemListEntry GetCosmeticsSearch(ctx, language, game).Query(query).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).Execute()

Search Cosmetics



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
    language := "language_example" // string | a valid language code
    game := "dofus2" // string | 
    query := "cape" // string | case sensitive search query
    filterTypeName := "filterTypeName_example" // string | only results with the translated type name (optional)
    filterMinLevel := int32(56) // int32 | only results which level is equal or above this value (optional)
    filterMaxLevel := int32(56) // int32 | only results which level is equal or below this value (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CosmeticsApi.GetCosmeticsSearch(context.Background(), language, game).Query(query).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CosmeticsApi.GetCosmeticsSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCosmeticsSearch`: []ItemListEntry
    fmt.Fprintf(os.Stdout, "Response from `CosmeticsApi.GetCosmeticsSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCosmeticsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **string** | case sensitive search query | 
 **filterTypeName** | **string** | only results with the translated type name | 
 **filterMinLevel** | **int32** | only results which level is equal or above this value | 
 **filterMaxLevel** | **int32** | only results which level is equal or below this value | 

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


## GetCosmeticsSingle

> Cosmetic GetCosmeticsSingle(ctx, language, ankamaId, game).Execute()

Single Cosmetics



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
    language := "language_example" // string | a valid language code
    ankamaId := int32(24132) // int32 | identifier
    game := "dofus2" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CosmeticsApi.GetCosmeticsSingle(context.Background(), language, ankamaId, game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CosmeticsApi.GetCosmeticsSingle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCosmeticsSingle`: Cosmetic
    fmt.Fprintf(os.Stdout, "Response from `CosmeticsApi.GetCosmeticsSingle`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetCosmeticsSingleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Cosmetic**](Cosmetic.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

