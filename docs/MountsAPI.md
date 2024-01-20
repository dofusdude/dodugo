# \MountsAPI

All URIs are relative to *https://api.dofusdu.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllMountsList**](MountsAPI.md#GetAllMountsList) | **Get** /{game}/{language}/mounts/all | List All Mounts
[**GetMountsList**](MountsAPI.md#GetMountsList) | **Get** /{game}/{language}/mounts | List Mounts
[**GetMountsSearch**](MountsAPI.md#GetMountsSearch) | **Get** /{game}/{language}/mounts/search | Search Mounts
[**GetMountsSingle**](MountsAPI.md#GetMountsSingle) | **Get** /{game}/{language}/mounts/{ankama_id} | Single Mounts



## GetAllMountsList

> MountsListPaged GetAllMountsList(ctx, language, game).FilterFamilyName(filterFamilyName).AcceptEncoding(acceptEncoding).Execute()

List All Mounts



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
	filterFamilyName := "Dragoturkey" // string | only results with the translated family name (optional)
	acceptEncoding := "acceptEncoding_example" // string | optional compression for saving bandwidth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MountsAPI.GetAllMountsList(context.Background(), language, game).FilterFamilyName(filterFamilyName).AcceptEncoding(acceptEncoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MountsAPI.GetAllMountsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllMountsList`: MountsListPaged
	fmt.Fprintf(os.Stdout, "Response from `MountsAPI.GetAllMountsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllMountsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filterFamilyName** | **string** | only results with the translated family name | 
 **acceptEncoding** | **string** | optional compression for saving bandwidth | 

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


## GetMountsList

> MountsListPaged GetMountsList(ctx, language, game).FilterFamilyName(filterFamilyName).PageSize(pageSize).PageNumber(pageNumber).FieldsMount(fieldsMount).Execute()

List Mounts



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
	filterFamilyName := "Dragoturkey" // string | only results with the translated family name (optional)
	pageSize := int32(10) // int32 | size of the results from the list. -1 disables pagination and gets all in one response. (optional)
	pageNumber := int32(1) // int32 | page number based on the current page[size]. So you could get page 1 with 8 entrys and page 2 would have entries 8 to 16. (optional)
	fieldsMount := []string{"FieldsMount_example"} // []string | adds fields from their detail endpoint to the item list entries. Multiple comma separated values allowed. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MountsAPI.GetMountsList(context.Background(), language, game).FilterFamilyName(filterFamilyName).PageSize(pageSize).PageNumber(pageNumber).FieldsMount(fieldsMount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MountsAPI.GetMountsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMountsList`: MountsListPaged
	fmt.Fprintf(os.Stdout, "Response from `MountsAPI.GetMountsList`: %v\n", resp)
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
 **fieldsMount** | **[]string** | adds fields from their detail endpoint to the item list entries. Multiple comma separated values allowed. | 

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

> []MountListEntry GetMountsSearch(ctx, language, game).Query(query).FilterFamilyName(filterFamilyName).Limit(limit).Execute()

Search Mounts



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
	query := "Dorée" // string | case sensitive search query
	filterFamilyName := "Dragodinde" // string | only results with the translated family name (optional)
	limit := int32(8) // int32 | maximum number of returned results (optional) (default to 8)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MountsAPI.GetMountsSearch(context.Background(), language, game).Query(query).FilterFamilyName(filterFamilyName).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MountsAPI.GetMountsSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMountsSearch`: []MountListEntry
	fmt.Fprintf(os.Stdout, "Response from `MountsAPI.GetMountsSearch`: %v\n", resp)
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
 **limit** | **int32** | maximum number of returned results | [default to 8]

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
	openapiclient "github.com/dofusdude/dodugo"
)

func main() {
	language := "language_example" // string | a valid language code
	ankamaId := int32(180) // int32 | identifier
	game := "dofus2" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MountsAPI.GetMountsSingle(context.Background(), language, ankamaId, game).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MountsAPI.GetMountsSingle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMountsSingle`: Mount
	fmt.Fprintf(os.Stdout, "Response from `MountsAPI.GetMountsSingle`: %v\n", resp)
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

