# \SetsAPI

All URIs are relative to *https://api.dofusdu.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllSetsList**](SetsAPI.md#GetAllSetsList) | **Get** /{game}/{language}/sets/all | List All Sets
[**GetSetsList**](SetsAPI.md#GetSetsList) | **Get** /{game}/{language}/sets | List Sets
[**GetSetsSearch**](SetsAPI.md#GetSetsSearch) | **Get** /{game}/{language}/sets/search | Search Sets
[**GetSetsSingle**](SetsAPI.md#GetSetsSingle) | **Get** /{game}/{language}/sets/{ankama_id} | Single Sets



## GetAllSetsList

> SetsListPaged GetAllSetsList(ctx, language, game).SortLevel(sortLevel).FilterMinHighestEquipmentLevel(filterMinHighestEquipmentLevel).FilterMaxHighestEquipmentLevel(filterMaxHighestEquipmentLevel).AcceptEncoding(acceptEncoding).FilterIsCosmetic(filterIsCosmetic).Execute()

List All Sets



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
	filterMinHighestEquipmentLevel := int32(190) // int32 | only results where the equipment with the highest level is above or equal to this value (optional)
	filterMaxHighestEquipmentLevel := int32(200) // int32 | only results where the equipment with the highest level is below or equal to this value (optional)
	acceptEncoding := "acceptEncoding_example" // string | optional compression for saving bandwidth (optional)
	filterIsCosmetic := true // bool | filter sets based on if they only got cosmetic items in it. If true, the item ids are for the cosmetic endpoints instead of equipment. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SetsAPI.GetAllSetsList(context.Background(), language, game).SortLevel(sortLevel).FilterMinHighestEquipmentLevel(filterMinHighestEquipmentLevel).FilterMaxHighestEquipmentLevel(filterMaxHighestEquipmentLevel).AcceptEncoding(acceptEncoding).FilterIsCosmetic(filterIsCosmetic).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SetsAPI.GetAllSetsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllSetsList`: SetsListPaged
	fmt.Fprintf(os.Stdout, "Response from `SetsAPI.GetAllSetsList`: %v\n", resp)
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
 **filterIsCosmetic** | **bool** | filter sets based on if they only got cosmetic items in it. If true, the item ids are for the cosmetic endpoints instead of equipment. | 

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

> SetsListPaged GetSetsList(ctx, language, game).SortLevel(sortLevel).FilterMinHighestEquipmentLevel(filterMinHighestEquipmentLevel).FilterMaxHighestEquipmentLevel(filterMaxHighestEquipmentLevel).PageSize(pageSize).PageNumber(pageNumber).FieldsSet(fieldsSet).FilterIsCosmetic(filterIsCosmetic).Execute()

List Sets



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
	filterMinHighestEquipmentLevel := int32(190) // int32 | only results where the equipment with the highest level is above or equal to this value (optional)
	filterMaxHighestEquipmentLevel := int32(200) // int32 | only results where the equipment with the highest level is below or equal to this value (optional)
	pageSize := int32(20) // int32 | size of the results from the list. -1 disables pagination and gets all in one response. (optional)
	pageNumber := int32(1) // int32 | page number based on the current page[size]. So you could get page 1 with 8 entrys and page 2 would have entries 8 to 16. (optional)
	fieldsSet := []string{"FieldsSet_example"} // []string | adds fields from their detail endpoint to the item list entries. Multiple comma separated values allowed. (optional)
	filterIsCosmetic := true // bool | filter sets based on if they only got cosmetic items in it. If true, the item ids are for the cosmetic endpoints instead of equipment. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SetsAPI.GetSetsList(context.Background(), language, game).SortLevel(sortLevel).FilterMinHighestEquipmentLevel(filterMinHighestEquipmentLevel).FilterMaxHighestEquipmentLevel(filterMaxHighestEquipmentLevel).PageSize(pageSize).PageNumber(pageNumber).FieldsSet(fieldsSet).FilterIsCosmetic(filterIsCosmetic).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SetsAPI.GetSetsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSetsList`: SetsListPaged
	fmt.Fprintf(os.Stdout, "Response from `SetsAPI.GetSetsList`: %v\n", resp)
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
 **filterIsCosmetic** | **bool** | filter sets based on if they only got cosmetic items in it. If true, the item ids are for the cosmetic endpoints instead of equipment. | 

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

> []SetListEntry GetSetsSearch(ctx, language, game).Query(query).FilterMinHighestEquipmentLevel(filterMinHighestEquipmentLevel).FilterMaxHighestEquipmentLevel(filterMaxHighestEquipmentLevel).Limit(limit).FilterIsCosmetic(filterIsCosmetic).Execute()

Search Sets



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
	query := "Des" // string | case sensitive search query
	filterMinHighestEquipmentLevel := int32(195) // int32 | only results where the equipment with the highest level is above or equal to this value (optional)
	filterMaxHighestEquipmentLevel := int32(200) // int32 | only results where the equipment with the highest level is below or equal to this value (optional)
	limit := int32(8) // int32 | maximum number of returned results (optional) (default to 8)
	filterIsCosmetic := true // bool | filter sets based on if they only got cosmetic items in it. If true, the item ids are for the cosmetic endpoints instead of equipment. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SetsAPI.GetSetsSearch(context.Background(), language, game).Query(query).FilterMinHighestEquipmentLevel(filterMinHighestEquipmentLevel).FilterMaxHighestEquipmentLevel(filterMaxHighestEquipmentLevel).Limit(limit).FilterIsCosmetic(filterIsCosmetic).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SetsAPI.GetSetsSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSetsSearch`: []SetListEntry
	fmt.Fprintf(os.Stdout, "Response from `SetsAPI.GetSetsSearch`: %v\n", resp)
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
 **limit** | **int32** | maximum number of returned results | [default to 8]
 **filterIsCosmetic** | **bool** | filter sets based on if they only got cosmetic items in it. If true, the item ids are for the cosmetic endpoints instead of equipment. | 

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
	openapiclient "github.com/dofusdude/dodugo"
)

func main() {
	language := "language_example" // string | a valid language code
	ankamaId := int32(499) // int32 | identifier
	game := "dofus2" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SetsAPI.GetSetsSingle(context.Background(), language, ankamaId, game).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SetsAPI.GetSetsSingle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSetsSingle`: EquipmentSet
	fmt.Fprintf(os.Stdout, "Response from `SetsAPI.GetSetsSingle`: %v\n", resp)
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

