# \SetsApi

All URIs are relative to *https://api.dofusdu.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllSetsList**](SetsApi.md#GetAllSetsList) | **Get** /{game}/{language}/sets/all | List All Sets
[**GetSetsList**](SetsApi.md#GetSetsList) | **Get** /{game}/{language}/sets | List Sets
[**GetSetsSearch**](SetsApi.md#GetSetsSearch) | **Get** /{game}/{language}/sets/search | Search Sets
[**GetSetsSingle**](SetsApi.md#GetSetsSingle) | **Get** /{game}/{language}/sets/{ankama_id} | Single Sets



## GetAllSetsList

> SetsListPaged GetAllSetsList(ctx, language, game).SortLevel(sortLevel).FilterMinHighestEquipmentLevel(filterMinHighestEquipmentLevel).FilterMaxHighestEquipmentLevel(filterMaxHighestEquipmentLevel).AcceptEncoding(acceptEncoding).Execute()

List All Sets



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
    sortLevel := "asc" // string | sort the resulting list by level, default unsorted (optional)
    filterMinHighestEquipmentLevel := int32(190) // int32 | only results where the equipment with the highest level is above or equal to this value (optional)
    filterMaxHighestEquipmentLevel := int32(200) // int32 | only results where the equipment with the highest level is below or equal to this value (optional)
    acceptEncoding := "acceptEncoding_example" // string | optional compression for saving bandwidth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SetsApi.GetAllSetsList(context.Background(), language, game).SortLevel(sortLevel).FilterMinHighestEquipmentLevel(filterMinHighestEquipmentLevel).FilterMaxHighestEquipmentLevel(filterMaxHighestEquipmentLevel).AcceptEncoding(acceptEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SetsApi.GetAllSetsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllSetsList`: SetsListPaged
    fmt.Fprintf(os.Stdout, "Response from `SetsApi.GetAllSetsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSetsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortLevel** | **string** | sort the resulting list by level, default unsorted | 
 **filterMinHighestEquipmentLevel** | **int32** | only results where the equipment with the highest level is above or equal to this value | 
 **filterMaxHighestEquipmentLevel** | **int32** | only results where the equipment with the highest level is below or equal to this value | 
 **acceptEncoding** | **string** | optional compression for saving bandwidth | 

### Return type

[**SetsListPaged**](SetsListPaged.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSetsList

> SetsListPaged GetSetsList(ctx, language, game).SortLevel(sortLevel).FilterMinHighestEquipmentLevel(filterMinHighestEquipmentLevel).FilterMaxHighestEquipmentLevel(filterMaxHighestEquipmentLevel).PageSize(pageSize).PageNumber(pageNumber).FieldsSet(fieldsSet).Execute()

List Sets



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
    sortLevel := "asc" // string | sort the resulting list by level, default unsorted (optional)
    filterMinHighestEquipmentLevel := int32(190) // int32 | only results where the equipment with the highest level is above or equal to this value (optional)
    filterMaxHighestEquipmentLevel := int32(200) // int32 | only results where the equipment with the highest level is below or equal to this value (optional)
    pageSize := int32(20) // int32 | size of the results from the list. -1 disables pagination and gets all in one response. (optional)
    pageNumber := int32(1) // int32 | page number based on the current page[size]. So you could get page 1 with 8 entrys and page 2 would have entries 8 to 16. (optional)
    fieldsSet := []string{"FieldsSet_example"} // []string | adds fields from their detail endpoint to the item list entries. Multiple comma separated values allowed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SetsApi.GetSetsList(context.Background(), language, game).SortLevel(sortLevel).FilterMinHighestEquipmentLevel(filterMinHighestEquipmentLevel).FilterMaxHighestEquipmentLevel(filterMaxHighestEquipmentLevel).PageSize(pageSize).PageNumber(pageNumber).FieldsSet(fieldsSet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SetsApi.GetSetsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSetsList`: SetsListPaged
    fmt.Fprintf(os.Stdout, "Response from `SetsApi.GetSetsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSetsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortLevel** | **string** | sort the resulting list by level, default unsorted | 
 **filterMinHighestEquipmentLevel** | **int32** | only results where the equipment with the highest level is above or equal to this value | 
 **filterMaxHighestEquipmentLevel** | **int32** | only results where the equipment with the highest level is below or equal to this value | 
 **pageSize** | **int32** | size of the results from the list. -1 disables pagination and gets all in one response. | 
 **pageNumber** | **int32** | page number based on the current page[size]. So you could get page 1 with 8 entrys and page 2 would have entries 8 to 16. | 
 **fieldsSet** | **[]string** | adds fields from their detail endpoint to the item list entries. Multiple comma separated values allowed. | 

### Return type

[**SetsListPaged**](SetsListPaged.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSetsSearch

> []SetListEntry GetSetsSearch(ctx, language, game).Query(query).FilterMinHighestEquipmentLevel(filterMinHighestEquipmentLevel).FilterMaxHighestEquipmentLevel(filterMaxHighestEquipmentLevel).Execute()

Search Sets



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
    query := "Des" // string | case sensitive search query
    filterMinHighestEquipmentLevel := int32(195) // int32 | only results where the equipment with the highest level is above or equal to this value (optional)
    filterMaxHighestEquipmentLevel := int32(200) // int32 | only results where the equipment with the highest level is below or equal to this value (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SetsApi.GetSetsSearch(context.Background(), language, game).Query(query).FilterMinHighestEquipmentLevel(filterMinHighestEquipmentLevel).FilterMaxHighestEquipmentLevel(filterMaxHighestEquipmentLevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SetsApi.GetSetsSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSetsSearch`: []SetListEntry
    fmt.Fprintf(os.Stdout, "Response from `SetsApi.GetSetsSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | a valid language code | 
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSetsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **string** | case sensitive search query | 
 **filterMinHighestEquipmentLevel** | **int32** | only results where the equipment with the highest level is above or equal to this value | 
 **filterMaxHighestEquipmentLevel** | **int32** | only results where the equipment with the highest level is below or equal to this value | 

### Return type

[**[]SetListEntry**](SetListEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSetsSingle

> EquipmentSet GetSetsSingle(ctx, language, ankamaId, game).Execute()

Single Sets



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
    ankamaId := int32(499) // int32 | identifier
    game := "dofus2" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SetsApi.GetSetsSingle(context.Background(), language, ankamaId, game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SetsApi.GetSetsSingle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSetsSingle`: EquipmentSet
    fmt.Fprintf(os.Stdout, "Response from `SetsApi.GetSetsSingle`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetSetsSingleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**EquipmentSet**](EquipmentSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

