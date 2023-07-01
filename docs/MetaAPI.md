# \MetaAPI

All URIs are relative to *https://api.dofusdu.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetaAlmanaxBonuses**](MetaAPI.md#GetMetaAlmanaxBonuses) | **Get** /dofus2/meta/{language}/almanax/bonuses | Available Almanax Bonuses
[**GetMetaElements**](MetaAPI.md#GetMetaElements) | **Get** /dofus2/meta/elements | Effects and Condition Elements



## GetMetaAlmanaxBonuses

> []GetMetaAlmanaxBonuses200ResponseInner GetMetaAlmanaxBonuses(ctx, language).Execute()

Available Almanax Bonuses



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
    language := "fr" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaAPI.GetMetaAlmanaxBonuses(context.Background(), language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaAPI.GetMetaAlmanaxBonuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetaAlmanaxBonuses`: []GetMetaAlmanaxBonuses200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `MetaAPI.GetMetaAlmanaxBonuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetaAlmanaxBonusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetMetaAlmanaxBonuses200ResponseInner**](GetMetaAlmanaxBonuses200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetaElements

> []string GetMetaElements(ctx).Execute()

Effects and Condition Elements



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaAPI.GetMetaElements(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaAPI.GetMetaElements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetaElements`: []string
    fmt.Fprintf(os.Stdout, "Response from `MetaAPI.GetMetaElements`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetaElementsRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

