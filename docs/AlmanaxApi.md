# \AlmanaxApi

All URIs are relative to *https://api.dofusdu.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAlmanaxDate**](AlmanaxApi.md#GetAlmanaxDate) | **Get** /dofus2/{language}/almanax/{date} | Single Almanax Date
[**GetAlmanaxRange**](AlmanaxApi.md#GetAlmanaxRange) | **Get** /dofus2/{language}/almanax | Almanax Range



## GetAlmanaxDate

> AlmanaxEntry GetAlmanaxDate(ctx, language, date).Execute()

Single Almanax Date



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
    language := "language_example" // string | code
    date := "date_example" // string | yyyy-mm-dd

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlmanaxApi.GetAlmanaxDate(context.Background(), language, date).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlmanaxApi.GetAlmanaxDate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlmanaxDate`: AlmanaxEntry
    fmt.Fprintf(os.Stdout, "Response from `AlmanaxApi.GetAlmanaxDate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | code | 
**date** | **string** | yyyy-mm-dd | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlmanaxDateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AlmanaxEntry**](AlmanaxEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlmanaxRange

> []AlmanaxEntry GetAlmanaxRange(ctx, language).FilterBonusType(filterBonusType).RangeFrom(rangeFrom).RangeTo(rangeTo).RangeSize(rangeSize).Timezone(timezone).Execute()

Almanax Range



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
    language := "language_example" // string | code
    filterBonusType := "experience-points" // string | ids from meta/{language}/almanax/bonuses (optional)
    rangeFrom := "rangeFrom_example" // string | yyyy-mm-dd (optional)
    rangeTo := "rangeTo_example" // string | yyyy-mm-dd (optional)
    rangeSize := int32(56) // int32 | size of the returned range (optional)
    timezone := "timezone_example" // string | determine what the current time is. If you live in Brazil, \"today\" will be hours apart from Paris. Use your timezone to get results relative to your location. (optional) (default to "Europe/Paris")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlmanaxApi.GetAlmanaxRange(context.Background(), language).FilterBonusType(filterBonusType).RangeFrom(rangeFrom).RangeTo(rangeTo).RangeSize(rangeSize).Timezone(timezone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlmanaxApi.GetAlmanaxRange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlmanaxRange`: []AlmanaxEntry
    fmt.Fprintf(os.Stdout, "Response from `AlmanaxApi.GetAlmanaxRange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlmanaxRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterBonusType** | **string** | ids from meta/{language}/almanax/bonuses | 
 **rangeFrom** | **string** | yyyy-mm-dd | 
 **rangeTo** | **string** | yyyy-mm-dd | 
 **rangeSize** | **int32** | size of the returned range | 
 **timezone** | **string** | determine what the current time is. If you live in Brazil, \&quot;today\&quot; will be hours apart from Paris. Use your timezone to get results relative to your location. | [default to &quot;Europe/Paris&quot;]

### Return type

[**[]AlmanaxEntry**](AlmanaxEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

