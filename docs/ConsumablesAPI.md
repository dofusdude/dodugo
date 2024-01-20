# \ConsumablesAPI

All URIs are relative to *https://api.dofusdu.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllItemsConsumablesList**](ConsumablesAPI.md#GetAllItemsConsumablesList) | **Get** /{game}/{language}/items/consumables/all | List All Consumables
[**GetItemsConsumablesList**](ConsumablesAPI.md#GetItemsConsumablesList) | **Get** /{game}/{language}/items/consumables | List Consumables
[**GetItemsConsumablesSearch**](ConsumablesAPI.md#GetItemsConsumablesSearch) | **Get** /{game}/{language}/items/consumables/search | Search Consumables
[**GetItemsConsumablesSingle**](ConsumablesAPI.md#GetItemsConsumablesSingle) | **Get** /{game}/{language}/items/consumables/{ankama_id} | Single Consumables



## GetAllItemsConsumablesList

> ItemsListPaged GetAllItemsConsumablesList(ctx, language, game).SortLevel(sortLevel).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).AcceptEncoding(acceptEncoding).Execute()

List All Consumables



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
	sortLevel := "asc" // string | sort the resulting list by level, default unsorted (optional)
	filterTypeName := "Chest" // string | only results with the translated type name (optional)
	filterMinLevel := int32(150) // int32 | only results which level is equal or above this value (optional)
	filterMaxLevel := int32(180) // int32 | only results which level is equal or below this value (optional)
	acceptEncoding := "acceptEncoding_example" // string | optional compression for saving bandwidth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumablesAPI.GetAllItemsConsumablesList(context.Background(), language, game).SortLevel(sortLevel).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).AcceptEncoding(acceptEncoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumablesAPI.GetAllItemsConsumablesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllItemsConsumablesList`: ItemsListPaged
	fmt.Fprintf(os.Stdout, "Response from `ConsumablesAPI.GetAllItemsConsumablesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllItemsConsumablesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortLevel** | **string** | sort the resulting list by level, default unsorted | 
 **filterTypeName** | **string** | only results with the translated type name | 
 **filterMinLevel** | **int32** | only results which level is equal or above this value | 
 **filterMaxLevel** | **int32** | only results which level is equal or below this value | 
 **acceptEncoding** | **string** | optional compression for saving bandwidth | 

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


## GetItemsConsumablesList

> ItemsListPaged GetItemsConsumablesList(ctx, language, game).SortLevel(sortLevel).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).PageSize(pageSize).PageNumber(pageNumber).FieldsItem(fieldsItem).Execute()

List Consumables



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
	sortLevel := "asc" // string | sort the resulting list by level, default unsorted (optional)
	filterTypeName := "Chest" // string | only results with the translated type name (optional)
	filterMinLevel := int32(150) // int32 | only results which level is equal or above this value (optional)
	filterMaxLevel := int32(180) // int32 | only results which level is equal or below this value (optional)
	pageSize := int32(2) // int32 | size of the results from the list. -1 disables pagination and gets all in one response. (optional)
	pageNumber := int32(1) // int32 | page number based on the current page[size]. So you could get page 1 with 8 entrys and page 2 would have entries 8 to 16. (optional)
	fieldsItem := []string{"FieldsItem_example"} // []string | adds fields from their detail endpoint to the item list entries. Multiple comma separated values allowed. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumablesAPI.GetItemsConsumablesList(context.Background(), language, game).SortLevel(sortLevel).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).PageSize(pageSize).PageNumber(pageNumber).FieldsItem(fieldsItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumablesAPI.GetItemsConsumablesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsConsumablesList`: ItemsListPaged
	fmt.Fprintf(os.Stdout, "Response from `ConsumablesAPI.GetItemsConsumablesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsConsumablesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortLevel** | **string** | sort the resulting list by level, default unsorted | 
 **filterTypeName** | **string** | only results with the translated type name | 
 **filterMinLevel** | **int32** | only results which level is equal or above this value | 
 **filterMaxLevel** | **int32** | only results which level is equal or below this value | 
 **pageSize** | **int32** | size of the results from the list. -1 disables pagination and gets all in one response. | 
 **pageNumber** | **int32** | page number based on the current page[size]. So you could get page 1 with 8 entrys and page 2 would have entries 8 to 16. | 
 **fieldsItem** | **[]string** | adds fields from their detail endpoint to the item list entries. Multiple comma separated values allowed. | 

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


## GetItemsConsumablesSearch

> []ItemListEntry GetItemsConsumablesSearch(ctx, language, game).Query(query).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).Limit(limit).Execute()

Search Consumables



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
	query := "Wholewrite" // string | case sensitive search query
	filterTypeName := "Bread" // string | only results with the translated type name (optional)
	filterMinLevel := int32(1) // int32 | only results which level is equal or above this value (optional)
	filterMaxLevel := int32(200) // int32 | only results which level is equal or below this value (optional)
	limit := int32(8) // int32 | maximum number of returned results (optional) (default to 8)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumablesAPI.GetItemsConsumablesSearch(context.Background(), language, game).Query(query).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumablesAPI.GetItemsConsumablesSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsConsumablesSearch`: []ItemListEntry
	fmt.Fprintf(os.Stdout, "Response from `ConsumablesAPI.GetItemsConsumablesSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsConsumablesSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **string** | case sensitive search query | 
 **filterTypeName** | **string** | only results with the translated type name | 
 **filterMinLevel** | **int32** | only results which level is equal or above this value | 
 **filterMaxLevel** | **int32** | only results which level is equal or below this value | 
 **limit** | **int32** | maximum number of returned results | [default to 8]

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


## GetItemsConsumablesSingle

> Resource GetItemsConsumablesSingle(ctx, language, ankamaId, game).Execute()

Single Consumables



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
	ankamaId := int32(17206) // int32 | identifier
	game := "dofus2" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumablesAPI.GetItemsConsumablesSingle(context.Background(), language, ankamaId, game).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumablesAPI.GetItemsConsumablesSingle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsConsumablesSingle`: Resource
	fmt.Fprintf(os.Stdout, "Response from `ConsumablesAPI.GetItemsConsumablesSingle`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetItemsConsumablesSingleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Resource**](Resource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

