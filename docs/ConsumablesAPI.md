# \ConsumablesAPI

All URIs are relative to *https://api.dofusdu.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllItemsConsumablesList**](ConsumablesAPI.md#GetAllItemsConsumablesList) | **Get** /{game}/v1/{language}/items/consumables/all | List All Consumables
[**GetItemsConsumablesList**](ConsumablesAPI.md#GetItemsConsumablesList) | **Get** /{game}/v1/{language}/items/consumables | List Consumables
[**GetItemsConsumablesSearch**](ConsumablesAPI.md#GetItemsConsumablesSearch) | **Get** /{game}/v1/{language}/items/consumables/search | Search Consumables
[**GetItemsConsumablesSingle**](ConsumablesAPI.md#GetItemsConsumablesSingle) | **Get** /{game}/v1/{language}/items/consumables/{ankama_id} | Single Consumables



## GetAllItemsConsumablesList

> ListItems GetAllItemsConsumablesList(ctx, language, game).SortLevel(sortLevel).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).AcceptEncoding(acceptEncoding).FilterTypeNameId(filterTypeNameId).Execute()

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
	game := "dofus3" // string | dofus3 | dofus3beta
	sortLevel := "asc" // string | sort the resulting list by level, default unsorted (optional)
	filterMinLevel := int32(150) // int32 | only results which level is equal or above this value (optional)
	filterMaxLevel := int32(180) // int32 | only results which level is equal or below this value (optional)
	acceptEncoding := "acceptEncoding_example" // string | optional compression for saving bandwidth (optional)
	filterTypeNameId := []string{"Inner_example"} // []string | multi-filter results with the english type name. Add with \"wood\" or \"+wood\" and exclude with \"-wood\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumablesAPI.GetAllItemsConsumablesList(context.Background(), language, game).SortLevel(sortLevel).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).AcceptEncoding(acceptEncoding).FilterTypeNameId(filterTypeNameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumablesAPI.GetAllItemsConsumablesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllItemsConsumablesList`: ListItems
	fmt.Fprintf(os.Stdout, "Response from `ConsumablesAPI.GetAllItemsConsumablesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** | dofus3 | dofus3beta | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllItemsConsumablesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortLevel** | **string** | sort the resulting list by level, default unsorted | 
 **filterMinLevel** | **int32** | only results which level is equal or above this value | 
 **filterMaxLevel** | **int32** | only results which level is equal or below this value | 
 **acceptEncoding** | **string** | optional compression for saving bandwidth | 
 **filterTypeNameId** | **[]string** | multi-filter results with the english type name. Add with \&quot;wood\&quot; or \&quot;+wood\&quot; and exclude with \&quot;-wood\&quot;. | 

### Return type

[**ListItems**](ListItems.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsConsumablesList

> ListItems GetItemsConsumablesList(ctx, language, game).SortLevel(sortLevel).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).PageSize(pageSize).PageNumber(pageNumber).FieldsItem(fieldsItem).FilterTypeNameId(filterTypeNameId).Execute()

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
	game := "dofus3" // string | dofus3 | dofus3beta
	sortLevel := "asc" // string | sort the resulting list by level, default unsorted (optional)
	filterMinLevel := int32(150) // int32 | only results which level is equal or above this value (optional)
	filterMaxLevel := int32(180) // int32 | only results which level is equal or below this value (optional)
	pageSize := int32(2) // int32 | size of the results from the list. -1 disables pagination and gets all in one response. (optional)
	pageNumber := int32(1) // int32 | page number based on the current page[size]. So you could get page 1 with 8 entrys and page 2 would have entries 8 to 16. (optional)
	fieldsItem := []string{"FieldsItem_example"} // []string | adds fields from their detail endpoint to the item list entries. Multiple comma separated values allowed. (optional)
	filterTypeNameId := []string{"Inner_example"} // []string | multi-filter results with the english type name. Add with \"wood\" or \"+wood\" and exclude with \"-wood\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumablesAPI.GetItemsConsumablesList(context.Background(), language, game).SortLevel(sortLevel).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).PageSize(pageSize).PageNumber(pageNumber).FieldsItem(fieldsItem).FilterTypeNameId(filterTypeNameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumablesAPI.GetItemsConsumablesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsConsumablesList`: ListItems
	fmt.Fprintf(os.Stdout, "Response from `ConsumablesAPI.GetItemsConsumablesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** | dofus3 | dofus3beta | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsConsumablesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortLevel** | **string** | sort the resulting list by level, default unsorted | 
 **filterMinLevel** | **int32** | only results which level is equal or above this value | 
 **filterMaxLevel** | **int32** | only results which level is equal or below this value | 
 **pageSize** | **int32** | size of the results from the list. -1 disables pagination and gets all in one response. | 
 **pageNumber** | **int32** | page number based on the current page[size]. So you could get page 1 with 8 entrys and page 2 would have entries 8 to 16. | 
 **fieldsItem** | **[]string** | adds fields from their detail endpoint to the item list entries. Multiple comma separated values allowed. | 
 **filterTypeNameId** | **[]string** | multi-filter results with the english type name. Add with \&quot;wood\&quot; or \&quot;+wood\&quot; and exclude with \&quot;-wood\&quot;. | 

### Return type

[**ListItems**](ListItems.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsConsumablesSearch

> []ListItem GetItemsConsumablesSearch(ctx, language, game).Query(query).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).Limit(limit).FilterTypeNameId(filterTypeNameId).Execute()

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
	game := "dofus3" // string | dofus3 | dofus3beta
	query := "Wholewrite" // string | case sensitive search query
	filterMinLevel := int32(1) // int32 | only results which level is equal or above this value (optional)
	filterMaxLevel := int32(200) // int32 | only results which level is equal or below this value (optional)
	limit := int32(8) // int32 | maximum number of returned results (optional) (default to 8)
	filterTypeNameId := []string{"Inner_example"} // []string | multi-filter results with the english type name. Add with \"wood\" or \"+wood\" and exclude with \"-wood\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumablesAPI.GetItemsConsumablesSearch(context.Background(), language, game).Query(query).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).Limit(limit).FilterTypeNameId(filterTypeNameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumablesAPI.GetItemsConsumablesSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsConsumablesSearch`: []ListItem
	fmt.Fprintf(os.Stdout, "Response from `ConsumablesAPI.GetItemsConsumablesSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** | dofus3 | dofus3beta | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsConsumablesSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **string** | case sensitive search query | 
 **filterMinLevel** | **int32** | only results which level is equal or above this value | 
 **filterMaxLevel** | **int32** | only results which level is equal or below this value | 
 **limit** | **int32** | maximum number of returned results | [default to 8]
 **filterTypeNameId** | **[]string** | multi-filter results with the english type name. Add with \&quot;wood\&quot; or \&quot;+wood\&quot; and exclude with \&quot;-wood\&quot;. | 

### Return type

[**[]ListItem**](ListItem.md)

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
	game := "dofus3" // string | dofus3 | dofus3beta

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
**game** | **string** | dofus3 | dofus3beta | 

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

