# \MountsApi

All URIs are relative to *https://api.dofusdu.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMountsList**](MountsApi.md#GetMountsList) | **Get** /{game}/{language}/mounts | List Mounts
[**GetMountsSearch**](MountsApi.md#GetMountsSearch) | **Get** /{game}/{language}/mounts/search | Search Mounts
[**GetMountsSingle**](MountsApi.md#GetMountsSingle) | **Get** /{game}/{language}/mounts/{ankama_id} | Single Mounts



## GetMountsList

> MountsListPaged GetMountsList(ctx, language, game).FilterFamilyName(filterFamilyName).PageSize(pageSize).PageNumber(pageNumber).Execute()

List Mounts



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
    filterFamilyName := "Dragoturkey" // string | only results with the translated family name (optional)
    pageSize := int32(10) // int32 | size of the results from the list. -1 disables pagination and gets all in one response. (optional)
    pageNumber := int32(1) // int32 | page number based on the current page[size]. So you could get page 1 with 8 entrys and page 2 would have entries 8 to 16. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MountsApi.GetMountsList(context.Background(), language, game).FilterFamilyName(filterFamilyName).PageSize(pageSize).PageNumber(pageNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MountsApi.GetMountsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMountsList`: MountsListPaged
    fmt.Fprintf(os.Stdout, "Response from `MountsApi.GetMountsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMountsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filterFamilyName** | **string** | only results with the translated family name | 
 **pageSize** | **int32** | size of the results from the list. -1 disables pagination and gets all in one response. | 
 **pageNumber** | **int32** | page number based on the current page[size]. So you could get page 1 with 8 entrys and page 2 would have entries 8 to 16. | 

### Return type

[**MountsListPaged**](MountsListPaged.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMountsSearch

> []MountListEntry GetMountsSearch(ctx, language, game).Query(query).FilterFamilyName(filterFamilyName).Execute()

Search Mounts



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
    query := "Dorée" // string | case sensitive search query
    filterFamilyName := "Dragodinde" // string | only results with the translated family name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MountsApi.GetMountsSearch(context.Background(), language, game).Query(query).FilterFamilyName(filterFamilyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MountsApi.GetMountsSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMountsSearch`: []MountListEntry
    fmt.Fprintf(os.Stdout, "Response from `MountsApi.GetMountsSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMountsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **string** | case sensitive search query | 
 **filterFamilyName** | **string** | only results with the translated family name | 

### Return type

[**[]MountListEntry**](MountListEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMountsSingle

> Mount GetMountsSingle(ctx, language, ankamaId, game).Execute()

Single Mounts



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
    ankamaId := int32(180) // int32 | identifier
    game := "dofus2" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MountsApi.GetMountsSingle(context.Background(), language, ankamaId, game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MountsApi.GetMountsSingle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMountsSingle`: Mount
    fmt.Fprintf(os.Stdout, "Response from `MountsApi.GetMountsSingle`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetMountsSingleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Mount**](Mount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

