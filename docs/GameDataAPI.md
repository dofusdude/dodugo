# \GameDataAPI

All URIs are relative to *https://api.dofusdu.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllAreas**](GameDataAPI.md#GetAllAreas) | **Get** /{game}/v1/{language}/areas/all | Get all areas
[**GetAllClasses**](GameDataAPI.md#GetAllClasses) | **Get** /{game}/v1/{language}/classes/all | Get all classes
[**GetAllMonsters**](GameDataAPI.md#GetAllMonsters) | **Get** /{game}/v1/{language}/monsters/all | Get all monsters
[**GetAllSpells**](GameDataAPI.md#GetAllSpells) | **Get** /{game}/v1/{language}/spells/all | Get all spells
[**GetAllSubareas**](GameDataAPI.md#GetAllSubareas) | **Get** /{game}/v1/{language}/subareas/all | Get all subareas
[**GetAllSuperAreas**](GameDataAPI.md#GetAllSuperAreas) | **Get** /{game}/v1/{language}/super-areas/all | Get all super-areas
[**GetAreaById**](GameDataAPI.md#GetAreaById) | **Get** /{game}/v1/{language}/areas/{ankama_id} | Get one area
[**GetAreas**](GameDataAPI.md#GetAreas) | **Get** /{game}/v1/{language}/areas | List areas
[**GetClassById**](GameDataAPI.md#GetClassById) | **Get** /{game}/v1/{language}/classes/{ankama_id} | Get one class
[**GetClasses**](GameDataAPI.md#GetClasses) | **Get** /{game}/v1/{language}/classes | List classes
[**GetMonsterById**](GameDataAPI.md#GetMonsterById) | **Get** /{game}/v1/{language}/monsters/{ankama_id} | Get one monster
[**GetMonsters**](GameDataAPI.md#GetMonsters) | **Get** /{game}/v1/{language}/monsters | List monsters
[**GetSpellById**](GameDataAPI.md#GetSpellById) | **Get** /{game}/v1/{language}/spells/{ankama_id} | Get one spell
[**GetSpells**](GameDataAPI.md#GetSpells) | **Get** /{game}/v1/{language}/spells | List spells
[**GetSubareaById**](GameDataAPI.md#GetSubareaById) | **Get** /{game}/v1/{language}/subareas/{ankama_id} | Get one subarea
[**GetSubareas**](GameDataAPI.md#GetSubareas) | **Get** /{game}/v1/{language}/subareas | List subareas
[**GetSuperAreaById**](GameDataAPI.md#GetSuperAreaById) | **Get** /{game}/v1/{language}/super-areas/{ankama_id} | Get one super-area
[**GetSuperAreas**](GameDataAPI.md#GetSuperAreas) | **Get** /{game}/v1/{language}/super-areas | List super-areas
[**SearchAreas**](GameDataAPI.md#SearchAreas) | **Get** /{game}/v1/{language}/areas/search | Search areas
[**SearchClasses**](GameDataAPI.md#SearchClasses) | **Get** /{game}/v1/{language}/classes/search | Search classes
[**SearchMonsters**](GameDataAPI.md#SearchMonsters) | **Get** /{game}/v1/{language}/monsters/search | Search monsters
[**SearchSpells**](GameDataAPI.md#SearchSpells) | **Get** /{game}/v1/{language}/spells/search | Search spells
[**SearchSubareas**](GameDataAPI.md#SearchSubareas) | **Get** /{game}/v1/{language}/subareas/search | Search subareas
[**SearchSuperAreas**](GameDataAPI.md#SearchSuperAreas) | **Get** /{game}/v1/{language}/super-areas/search | Search super-areas



## GetAllAreas

> MappedEntityCollection GetAllAreas(ctx, game, language).Execute()

Get all areas

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
	game := "game_example" // string | 
	language := "language_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.GetAllAreas(context.Background(), game, language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.GetAllAreas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllAreas`: MappedEntityCollection
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.GetAllAreas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAreasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MappedEntityCollection**](MappedEntityCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllClasses

> MappedEntityCollection GetAllClasses(ctx, game, language).Execute()

Get all classes

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
	game := "game_example" // string | 
	language := "language_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.GetAllClasses(context.Background(), game, language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.GetAllClasses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllClasses`: MappedEntityCollection
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.GetAllClasses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllClassesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MappedEntityCollection**](MappedEntityCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllMonsters

> MappedEntityCollection GetAllMonsters(ctx, game, language).Include(include).Execute()

Get all monsters

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
	game := "game_example" // string | 
	language := "language_example" // string | 
	include := "include_example" // string | Optional spell expansion for monster responses. Leave include unset to return spell URLs. Use `spells.grade` to embed the assigned spell grade. Use `spells` to embed every available grade of each assigned spell. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.GetAllMonsters(context.Background(), game, language).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.GetAllMonsters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllMonsters`: MappedEntityCollection
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.GetAllMonsters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllMonstersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **include** | **string** | Optional spell expansion for monster responses. Leave include unset to return spell URLs. Use &#x60;spells.grade&#x60; to embed the assigned spell grade. Use &#x60;spells&#x60; to embed every available grade of each assigned spell. | 

### Return type

[**MappedEntityCollection**](MappedEntityCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllSpells

> MappedEntityCollection GetAllSpells(ctx, game, language).Execute()

Get all spells

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
	game := "game_example" // string | 
	language := "language_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.GetAllSpells(context.Background(), game, language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.GetAllSpells``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllSpells`: MappedEntityCollection
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.GetAllSpells`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSpellsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MappedEntityCollection**](MappedEntityCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllSubareas

> MappedEntityCollection GetAllSubareas(ctx, game, language).Execute()

Get all subareas

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
	game := "game_example" // string | 
	language := "language_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.GetAllSubareas(context.Background(), game, language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.GetAllSubareas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllSubareas`: MappedEntityCollection
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.GetAllSubareas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSubareasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MappedEntityCollection**](MappedEntityCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllSuperAreas

> MappedEntityCollection GetAllSuperAreas(ctx, game, language).Execute()

Get all super-areas

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
	game := "game_example" // string | 
	language := "language_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.GetAllSuperAreas(context.Background(), game, language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.GetAllSuperAreas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllSuperAreas`: MappedEntityCollection
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.GetAllSuperAreas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSuperAreasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MappedEntityCollection**](MappedEntityCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAreaById

> WorldArea GetAreaById(ctx, game, language, ankamaId).Execute()

Get one area

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
	game := "game_example" // string | 
	language := "language_example" // string | 
	ankamaId := int32(0) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.GetAreaById(context.Background(), game, language, ankamaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.GetAreaById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAreaById`: WorldArea
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.GetAreaById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 
**ankamaId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAreaByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorldArea**](WorldArea.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAreas

> MappedEntityCollection GetAreas(ctx, game, language).PageNumber(pageNumber).PageSize(pageSize).FieldsAreas(fieldsAreas).Execute()

List areas

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
	game := "game_example" // string | 
	language := "language_example" // string | 
	pageNumber := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	fieldsAreas := "fieldsAreas_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.GetAreas(context.Background(), game, language).PageNumber(pageNumber).PageSize(pageSize).FieldsAreas(fieldsAreas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.GetAreas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAreas`: MappedEntityCollection
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.GetAreas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAreasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageNumber** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]
 **fieldsAreas** | **string** |  | 

### Return type

[**MappedEntityCollection**](MappedEntityCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClassById

> Class GetClassById(ctx, game, language, ankamaId).Execute()

Get one class

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
	game := "game_example" // string | 
	language := "language_example" // string | 
	ankamaId := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.GetClassById(context.Background(), game, language, ankamaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.GetClassById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClassById`: Class
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.GetClassById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 
**ankamaId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClassByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Class**](Class.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClasses

> MappedEntityCollection GetClasses(ctx, game, language).PageNumber(pageNumber).PageSize(pageSize).FieldsClasses(fieldsClasses).Execute()

List classes

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
	game := "game_example" // string | 
	language := "language_example" // string | 
	pageNumber := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	fieldsClasses := "fieldsClasses_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.GetClasses(context.Background(), game, language).PageNumber(pageNumber).PageSize(pageSize).FieldsClasses(fieldsClasses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.GetClasses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClasses`: MappedEntityCollection
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.GetClasses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClassesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageNumber** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]
 **fieldsClasses** | **string** |  | 

### Return type

[**MappedEntityCollection**](MappedEntityCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonsterById

> Monster GetMonsterById(ctx, game, language, ankamaId).Include(include).Execute()

Get one monster

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
	game := "game_example" // string | 
	language := "language_example" // string | 
	ankamaId := int32(31) // int32 | 
	include := "include_example" // string | Optional spell expansion for monster responses. Leave include unset to return spell URLs. Use `spells.grade` to embed the assigned spell grade. Use `spells` to embed every available grade of each assigned spell. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.GetMonsterById(context.Background(), game, language, ankamaId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.GetMonsterById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMonsterById`: Monster
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.GetMonsterById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 
**ankamaId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonsterByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **include** | **string** | Optional spell expansion for monster responses. Leave include unset to return spell URLs. Use &#x60;spells.grade&#x60; to embed the assigned spell grade. Use &#x60;spells&#x60; to embed every available grade of each assigned spell. | 

### Return type

[**Monster**](Monster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonsters

> MappedEntityCollection GetMonsters(ctx, game, language).PageNumber(pageNumber).PageSize(pageSize).FieldsMonsters(fieldsMonsters).Include(include).Execute()

List monsters

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
	game := "game_example" // string | 
	language := "language_example" // string | 
	pageNumber := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	fieldsMonsters := "fieldsMonsters_example" // string |  (optional)
	include := "include_example" // string | Optional spell expansion for monster responses. Leave include unset to return spell URLs. Use `spells.grade` to embed the assigned spell grade. Use `spells` to embed every available grade of each assigned spell. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.GetMonsters(context.Background(), game, language).PageNumber(pageNumber).PageSize(pageSize).FieldsMonsters(fieldsMonsters).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.GetMonsters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMonsters`: MappedEntityCollection
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.GetMonsters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonstersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageNumber** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]
 **fieldsMonsters** | **string** |  | 
 **include** | **string** | Optional spell expansion for monster responses. Leave include unset to return spell URLs. Use &#x60;spells.grade&#x60; to embed the assigned spell grade. Use &#x60;spells&#x60; to embed every available grade of each assigned spell. | 

### Return type

[**MappedEntityCollection**](MappedEntityCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpellById

> Spell GetSpellById(ctx, game, language, ankamaId).Grade(grade).Execute()

Get one spell

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
	game := "game_example" // string | 
	language := "language_example" // string | 
	ankamaId := int32(201) // int32 | 
	grade := int32(56) // int32 | Return only the requested mapped spell grade in `levels`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.GetSpellById(context.Background(), game, language, ankamaId).Grade(grade).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.GetSpellById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpellById`: Spell
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.GetSpellById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 
**ankamaId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpellByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **grade** | **int32** | Return only the requested mapped spell grade in &#x60;levels&#x60;. | 

### Return type

[**Spell**](Spell.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpells

> MappedEntityCollection GetSpells(ctx, game, language).PageNumber(pageNumber).PageSize(pageSize).FieldsSpells(fieldsSpells).Execute()

List spells

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
	game := "game_example" // string | 
	language := "language_example" // string | 
	pageNumber := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	fieldsSpells := "fieldsSpells_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.GetSpells(context.Background(), game, language).PageNumber(pageNumber).PageSize(pageSize).FieldsSpells(fieldsSpells).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.GetSpells``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpells`: MappedEntityCollection
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.GetSpells`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpellsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageNumber** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]
 **fieldsSpells** | **string** |  | 

### Return type

[**MappedEntityCollection**](MappedEntityCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubareaById

> WorldArea GetSubareaById(ctx, game, language, ankamaId).Execute()

Get one subarea

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
	game := "game_example" // string | 
	language := "language_example" // string | 
	ankamaId := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.GetSubareaById(context.Background(), game, language, ankamaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.GetSubareaById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubareaById`: WorldArea
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.GetSubareaById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 
**ankamaId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubareaByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorldArea**](WorldArea.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubareas

> MappedEntityCollection GetSubareas(ctx, game, language).PageNumber(pageNumber).PageSize(pageSize).FieldsSubareas(fieldsSubareas).Execute()

List subareas

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
	game := "game_example" // string | 
	language := "language_example" // string | 
	pageNumber := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	fieldsSubareas := "fieldsSubareas_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.GetSubareas(context.Background(), game, language).PageNumber(pageNumber).PageSize(pageSize).FieldsSubareas(fieldsSubareas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.GetSubareas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubareas`: MappedEntityCollection
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.GetSubareas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubareasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageNumber** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]
 **fieldsSubareas** | **string** |  | 

### Return type

[**MappedEntityCollection**](MappedEntityCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuperAreaById

> WorldArea GetSuperAreaById(ctx, game, language, ankamaId).Execute()

Get one super-area

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
	game := "game_example" // string | 
	language := "language_example" // string | 
	ankamaId := int32(0) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.GetSuperAreaById(context.Background(), game, language, ankamaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.GetSuperAreaById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSuperAreaById`: WorldArea
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.GetSuperAreaById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 
**ankamaId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSuperAreaByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorldArea**](WorldArea.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuperAreas

> MappedEntityCollection GetSuperAreas(ctx, game, language).PageNumber(pageNumber).PageSize(pageSize).FieldsSuperAreas(fieldsSuperAreas).Execute()

List super-areas

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
	game := "game_example" // string | 
	language := "language_example" // string | 
	pageNumber := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	fieldsSuperAreas := "fieldsSuperAreas_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.GetSuperAreas(context.Background(), game, language).PageNumber(pageNumber).PageSize(pageSize).FieldsSuperAreas(fieldsSuperAreas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.GetSuperAreas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSuperAreas`: MappedEntityCollection
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.GetSuperAreas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSuperAreasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageNumber** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]
 **fieldsSuperAreas** | **string** |  | 

### Return type

[**MappedEntityCollection**](MappedEntityCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAreas

> []MappedEntityListEntry SearchAreas(ctx, game, language).Query(query).Limit(limit).Execute()

Search areas

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
	game := "game_example" // string | 
	language := "language_example" // string | 
	query := "query_example" // string | 
	limit := int32(56) // int32 |  (optional) (default to 8)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.SearchAreas(context.Background(), game, language).Query(query).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.SearchAreas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAreas`: []MappedEntityListEntry
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.SearchAreas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchAreasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **string** |  | 
 **limit** | **int32** |  | [default to 8]

### Return type

[**[]MappedEntityListEntry**](MappedEntityListEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchClasses

> []MappedEntityListEntry SearchClasses(ctx, game, language).Query(query).Limit(limit).Execute()

Search classes

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
	game := "game_example" // string | 
	language := "language_example" // string | 
	query := "query_example" // string | 
	limit := int32(56) // int32 |  (optional) (default to 8)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.SearchClasses(context.Background(), game, language).Query(query).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.SearchClasses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchClasses`: []MappedEntityListEntry
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.SearchClasses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchClassesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **string** |  | 
 **limit** | **int32** |  | [default to 8]

### Return type

[**[]MappedEntityListEntry**](MappedEntityListEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchMonsters

> []MappedEntityListEntry SearchMonsters(ctx, game, language).Query(query).Limit(limit).Execute()

Search monsters

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
	game := "game_example" // string | 
	language := "language_example" // string | 
	query := "query_example" // string | 
	limit := int32(56) // int32 |  (optional) (default to 8)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.SearchMonsters(context.Background(), game, language).Query(query).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.SearchMonsters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchMonsters`: []MappedEntityListEntry
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.SearchMonsters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchMonstersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **string** |  | 
 **limit** | **int32** |  | [default to 8]

### Return type

[**[]MappedEntityListEntry**](MappedEntityListEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSpells

> []MappedEntityListEntry SearchSpells(ctx, game, language).Query(query).Limit(limit).Execute()

Search spells

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
	game := "game_example" // string | 
	language := "language_example" // string | 
	query := "query_example" // string | 
	limit := int32(56) // int32 |  (optional) (default to 8)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.SearchSpells(context.Background(), game, language).Query(query).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.SearchSpells``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSpells`: []MappedEntityListEntry
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.SearchSpells`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSpellsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **string** |  | 
 **limit** | **int32** |  | [default to 8]

### Return type

[**[]MappedEntityListEntry**](MappedEntityListEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSubareas

> []MappedEntityListEntry SearchSubareas(ctx, game, language).Query(query).Limit(limit).Execute()

Search subareas

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
	game := "game_example" // string | 
	language := "language_example" // string | 
	query := "query_example" // string | 
	limit := int32(56) // int32 |  (optional) (default to 8)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.SearchSubareas(context.Background(), game, language).Query(query).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.SearchSubareas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSubareas`: []MappedEntityListEntry
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.SearchSubareas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSubareasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **string** |  | 
 **limit** | **int32** |  | [default to 8]

### Return type

[**[]MappedEntityListEntry**](MappedEntityListEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSuperAreas

> []MappedEntityListEntry SearchSuperAreas(ctx, game, language).Query(query).Limit(limit).Execute()

Search super-areas

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
	game := "game_example" // string | 
	language := "language_example" // string | 
	query := "query_example" // string | 
	limit := int32(56) // int32 |  (optional) (default to 8)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameDataAPI.SearchSuperAreas(context.Background(), game, language).Query(query).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameDataAPI.SearchSuperAreas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSuperAreas`: []MappedEntityListEntry
	fmt.Fprintf(os.Stdout, "Response from `GameDataAPI.SearchSuperAreas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 
**language** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSuperAreasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **string** |  | 
 **limit** | **int32** |  | [default to 8]

### Return type

[**[]MappedEntityListEntry**](MappedEntityListEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

