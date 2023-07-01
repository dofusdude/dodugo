# \EquipmentAPI

All URIs are relative to *https://api.dofusdu.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllItemsEquipmentList**](EquipmentAPI.md#GetAllItemsEquipmentList) | **Get** /{game}/{language}/items/equipment/all | List All Equipment
[**GetItemsEquipmentList**](EquipmentAPI.md#GetItemsEquipmentList) | **Get** /{game}/{language}/items/equipment | List Equipment
[**GetItemsEquipmentSearch**](EquipmentAPI.md#GetItemsEquipmentSearch) | **Get** /{game}/{language}/items/equipment/search | Search Equipment
[**GetItemsEquipmentSingle**](EquipmentAPI.md#GetItemsEquipmentSingle) | **Get** /{game}/{language}/items/equipment/{ankama_id} | Single Equipment



## GetAllItemsEquipmentList

> ItemsListPaged GetAllItemsEquipmentList(ctx, language, game).SortLevel(sortLevel).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).AcceptEncoding(acceptEncoding).Execute()

List All Equipment



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
    language := "en" // string | a valid language code
    game := "dofus2" // string | 
    sortLevel := "desc" // string | sort the resulting list by level, default unsorted (optional)
    filterTypeName := "Sword" // string | only results with the translated type name (optional)
    filterMinLevel := int32(10) // int32 | only results which level is equal or above this value (optional)
    filterMaxLevel := int32(60) // int32 | only results which level is equal or below this value (optional)
    acceptEncoding := "acceptEncoding_example" // string | optional compression for saving bandwidth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EquipmentAPI.GetAllItemsEquipmentList(context.Background(), language, game).SortLevel(sortLevel).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).AcceptEncoding(acceptEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EquipmentAPI.GetAllItemsEquipmentList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllItemsEquipmentList`: ItemsListPaged
    fmt.Fprintf(os.Stdout, "Response from `EquipmentAPI.GetAllItemsEquipmentList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllItemsEquipmentListRequest struct via the builder pattern


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


## GetItemsEquipmentList

> ItemsListPaged GetItemsEquipmentList(ctx, language, game).SortLevel(sortLevel).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).PageSize(pageSize).PageNumber(pageNumber).FieldsItem(fieldsItem).Execute()

List Equipment



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
    language := "en" // string | a valid language code
    game := "dofus2" // string | 
    sortLevel := "desc" // string | sort the resulting list by level, default unsorted (optional)
    filterTypeName := "Sword" // string | only results with the translated type name (optional)
    filterMinLevel := int32(10) // int32 | only results which level is equal or above this value (optional)
    filterMaxLevel := int32(60) // int32 | only results which level is equal or below this value (optional)
    pageSize := int32(5) // int32 | size of the results from the list. -1 disables pagination and gets all in one response. (optional)
    pageNumber := int32(1) // int32 | page number based on the current page[size]. So you could get page 1 with 8 entrys and page 2 would have entries 8 to 16. (optional)
    fieldsItem := []string{"FieldsItem_example"} // []string | adds fields from their detail endpoint to the item list entries. Multiple comma separated values allowed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EquipmentAPI.GetItemsEquipmentList(context.Background(), language, game).SortLevel(sortLevel).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).PageSize(pageSize).PageNumber(pageNumber).FieldsItem(fieldsItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EquipmentAPI.GetItemsEquipmentList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemsEquipmentList`: ItemsListPaged
    fmt.Fprintf(os.Stdout, "Response from `EquipmentAPI.GetItemsEquipmentList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsEquipmentListRequest struct via the builder pattern


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


## GetItemsEquipmentSearch

> []ItemListEntry GetItemsEquipmentSearch(ctx, language, game).Query(query).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).Limit(limit).Execute()

Search Equipment



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
    query := "nidas" // string | case sensitive search query
    filterTypeName := "boots" // string | only results with the translated type name (optional)
    filterMinLevel := int32(150) // int32 | only results which level is equal or above this value (optional)
    filterMaxLevel := int32(200) // int32 | only results which level is equal or below this value (optional)
    limit := int32(8) // int32 | maximum number of returned results (optional) (default to 8)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EquipmentAPI.GetItemsEquipmentSearch(context.Background(), language, game).Query(query).FilterTypeName(filterTypeName).FilterMinLevel(filterMinLevel).FilterMaxLevel(filterMaxLevel).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EquipmentAPI.GetItemsEquipmentSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemsEquipmentSearch`: []ItemListEntry
    fmt.Fprintf(os.Stdout, "Response from `EquipmentAPI.GetItemsEquipmentSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsEquipmentSearchRequest struct via the builder pattern


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


## GetItemsEquipmentSingle

> Weapon GetItemsEquipmentSingle(ctx, language, ankamaId, game).Execute()

Single Equipment



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
    ankamaId := int32(26009) // int32 | identifier
    game := "dofus2" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EquipmentAPI.GetItemsEquipmentSingle(context.Background(), language, ankamaId, game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EquipmentAPI.GetItemsEquipmentSingle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemsEquipmentSingle`: Weapon
    fmt.Fprintf(os.Stdout, "Response from `EquipmentAPI.GetItemsEquipmentSingle`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetItemsEquipmentSingleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Weapon**](Weapon.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

