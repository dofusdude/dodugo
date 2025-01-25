# \AlmanaxAPI

All URIs are relative to *https://api.dofusdu.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAlmanaxDate**](AlmanaxAPI.md#GetAlmanaxDate) | **Get** /dofus3/v1/{language}/almanax/{date} | Single Almanax Date
[**GetAlmanaxRange**](AlmanaxAPI.md#GetAlmanaxRange) | **Get** /dofus3/v1/{language}/almanax | Almanax Range



## GetAlmanaxDate

> Almanax GetAlmanaxDate(ctx, language, date).Level(level).Execute()

Single Almanax Date



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/dofusdude/dodugo"
)

func main() {
	language := "fr" // string | code
	date := time.Now() // string | yyyy-mm-dd
	level := int32(56) // int32 | character level for the reward_xp field (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlmanaxAPI.GetAlmanaxDate(context.Background(), language, date).Level(level).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlmanaxAPI.GetAlmanaxDate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlmanaxDate`: Almanax
	fmt.Fprintf(os.Stdout, "Response from `AlmanaxAPI.GetAlmanaxDate`: %v\n", resp)
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


 **level** | **int32** | character level for the reward_xp field | 

### Return type

[**Almanax**](Almanax.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlmanaxRange

> []Almanax GetAlmanaxRange(ctx, language).FilterBonusType(filterBonusType).RangeFrom(rangeFrom).RangeTo(rangeTo).RangeSize(rangeSize).Timezone(timezone).Level(level).Execute()

Almanax Range



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/dofusdude/dodugo"
)

func main() {
	language := "fr" // string | code
	filterBonusType := "experience-points" // string | ids from meta/{language}/almanax/bonuses (optional)
	rangeFrom := time.Now() // string | yyyy-mm-dd (optional)
	rangeTo := time.Now() // string | yyyy-mm-dd (optional)
	rangeSize := int32(-1) // int32 | Size of the returned range. Disable to fully use the range by setting size to -1. (optional)
	timezone := "Europe/Paris" // string | determine what the current time is. If you live in Brazil, \"today\" will be hours apart from Paris. Use your timezone to get results relative to your location. (optional) (default to "Europe/Paris")
	level := int32(56) // int32 | character level for the reward_xp field (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlmanaxAPI.GetAlmanaxRange(context.Background(), language).FilterBonusType(filterBonusType).RangeFrom(rangeFrom).RangeTo(rangeTo).RangeSize(rangeSize).Timezone(timezone).Level(level).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlmanaxAPI.GetAlmanaxRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlmanaxRange`: []Almanax
	fmt.Fprintf(os.Stdout, "Response from `AlmanaxAPI.GetAlmanaxRange`: %v\n", resp)
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
 **rangeSize** | **int32** | Size of the returned range. Disable to fully use the range by setting size to -1. | 
 **timezone** | **string** | determine what the current time is. If you live in Brazil, \&quot;today\&quot; will be hours apart from Paris. Use your timezone to get results relative to your location. | [default to &quot;Europe/Paris&quot;]
 **level** | **int32** | character level for the reward_xp field | 

### Return type

[**[]Almanax**](Almanax.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

