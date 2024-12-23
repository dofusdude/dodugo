# \MetaAPI

All URIs are relative to *https://api.dofusdu.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGameSearchTypes**](MetaAPI.md#GetGameSearchTypes) | **Get** /{game}/v1/meta/search/types | Available Game Search Types
[**GetItemTypes**](MetaAPI.md#GetItemTypes) | **Get** /{game}/v1/meta/items/types | Available Item Types
[**GetMetaAlmanaxBonuses**](MetaAPI.md#GetMetaAlmanaxBonuses) | **Get** /dofus3/v1/meta/{language}/almanax/bonuses | Available Almanax Bonuses
[**GetMetaAlmanaxBonusesSearch**](MetaAPI.md#GetMetaAlmanaxBonusesSearch) | **Get** /dofus3/v1/meta/{language}/almanax/bonuses/search | Search Available Almanax Bonuses
[**GetMetaElements**](MetaAPI.md#GetMetaElements) | **Get** /{game}/v1/meta/elements | Effects and Condition Elements
[**GetMetaVersion**](MetaAPI.md#GetMetaVersion) | **Get** /{game}/v1/meta/version | Game Version



## GetGameSearchTypes

> []string GetGameSearchTypes(ctx, game).Execute()

Available Game Search Types



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
	game := "dofus3" // string | game main 'dofus3' or beta channel 'dofus3beta'

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetaAPI.GetGameSearchTypes(context.Background(), game).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetaAPI.GetGameSearchTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGameSearchTypes`: []string
	fmt.Fprintf(os.Stdout, "Response from `MetaAPI.GetGameSearchTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** | game main &#39;dofus3&#39; or beta channel &#39;dofus3beta&#39; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGameSearchTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetItemTypes

> []string GetItemTypes(ctx, game).Execute()

Available Item Types



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
	game := "dofus3" // string | game main 'dofus3' or beta channel 'dofus3beta'

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetaAPI.GetItemTypes(context.Background(), game).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetaAPI.GetItemTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemTypes`: []string
	fmt.Fprintf(os.Stdout, "Response from `MetaAPI.GetItemTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** | game main &#39;dofus3&#39; or beta channel &#39;dofus3beta&#39; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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
	language := "fr" // string | a valid language code

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
**language** | **string** | a valid language code | 

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


## GetMetaAlmanaxBonusesSearch

> []GetMetaAlmanaxBonuses200ResponseInner GetMetaAlmanaxBonusesSearch(ctx, language).Query(query).Limit(limit).Execute()

Search Available Almanax Bonuses



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
	query := "abond" // string | case sensitive search query
	limit := int32(56) // int32 | maximum number of returned results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetaAPI.GetMetaAlmanaxBonusesSearch(context.Background(), language).Query(query).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetaAPI.GetMetaAlmanaxBonusesSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetaAlmanaxBonusesSearch`: []GetMetaAlmanaxBonuses200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MetaAPI.GetMetaAlmanaxBonusesSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetaAlmanaxBonusesSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **string** | case sensitive search query | 
 **limit** | **int32** | maximum number of returned results | 

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

> []string GetMetaElements(ctx, game).Execute()

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
	game := "dofus3" // string | game main 'dofus3' or beta channel 'dofus3beta'

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetaAPI.GetMetaElements(context.Background(), game).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetaAPI.GetMetaElements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetaElements`: []string
	fmt.Fprintf(os.Stdout, "Response from `MetaAPI.GetMetaElements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** | game main &#39;dofus3&#39; or beta channel &#39;dofus3beta&#39; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetaElementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetMetaVersion

> Version GetMetaVersion(ctx, game).Execute()

Game Version



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
	game := "dofus3" // string | game main 'dofus3' or beta channel 'dofus3beta'

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetaAPI.GetMetaVersion(context.Background(), game).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetaAPI.GetMetaVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetaVersion`: Version
	fmt.Fprintf(os.Stdout, "Response from `MetaAPI.GetMetaVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** | game main &#39;dofus3&#39; or beta channel &#39;dofus3beta&#39; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetaVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Version**](Version.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

