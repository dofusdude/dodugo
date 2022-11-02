# \WebhooksApi

All URIs are relative to *https://api.dofusdu.de*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWebhooksAlmanaxId**](WebhooksApi.md#DeleteWebhooksAlmanaxId) | **Delete** /webhooks/almanax/{id} | Unregister Almanax Hook
[**DeleteWebhooksRssId**](WebhooksApi.md#DeleteWebhooksRssId) | **Delete** /webhooks/rss/{id} | Unregister RSS Hook
[**DeleteWebhooksTwitterId**](WebhooksApi.md#DeleteWebhooksTwitterId) | **Delete** /webhooks/twitter/{id} | Unregister Twitter Hook
[**GetMetaWebhooksAlmanax**](WebhooksApi.md#GetMetaWebhooksAlmanax) | **Get** /meta/webhooks/almanax | Get Almanax Hook Metainfo
[**GetMetaWebhooksRss**](WebhooksApi.md#GetMetaWebhooksRss) | **Get** /meta/webhooks/rss | Get RSS Hook Metainfo
[**GetMetaWebhooksTwitter**](WebhooksApi.md#GetMetaWebhooksTwitter) | **Get** /meta/webhooks/twitter | Get Twitter Hook Metainfo
[**GetWebhooksAlmanaxId**](WebhooksApi.md#GetWebhooksAlmanaxId) | **Get** /webhooks/almanax/{id} | Get Almanax Hook
[**GetWebhooksRssId**](WebhooksApi.md#GetWebhooksRssId) | **Get** /webhooks/rss/{id} | Get RSS Hook
[**GetWebhooksTwitterId**](WebhooksApi.md#GetWebhooksTwitterId) | **Get** /webhooks/twitter/{id} | Get Twitter Hook
[**PostWebhooksAlmanax**](WebhooksApi.md#PostWebhooksAlmanax) | **Post** /webhooks/almanax | Register Almanax Hook
[**PostWebhooksRss**](WebhooksApi.md#PostWebhooksRss) | **Post** /webhooks/rss | Register RSS Hook
[**PostWebhooksTwitter**](WebhooksApi.md#PostWebhooksTwitter) | **Post** /webhooks/twitter | Register Twitter Hook
[**PutWebhooksAlmanaxId**](WebhooksApi.md#PutWebhooksAlmanaxId) | **Put** /webhooks/almanax/{id} | Update Almanax Hook
[**PutWebhooksRssId**](WebhooksApi.md#PutWebhooksRssId) | **Put** /webhooks/rss/{id} | Update RSS Hook
[**PutWebhooksTwitterId**](WebhooksApi.md#PutWebhooksTwitterId) | **Put** /webhooks/twitter/{id} | Update Twitter Hook



## DeleteWebhooksAlmanaxId

> DeleteWebhooksAlmanaxId(ctx, id).Execute()

Unregister Almanax Hook



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.DeleteWebhooksAlmanaxId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.DeleteWebhooksAlmanaxId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhooksAlmanaxIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhooksRssId

> DeleteWebhooksRssId(ctx, id).Execute()

Unregister RSS Hook



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.DeleteWebhooksRssId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.DeleteWebhooksRssId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhooksRssIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhooksTwitterId

> DeleteWebhooksTwitterId(ctx, id).Execute()

Unregister Twitter Hook



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.DeleteWebhooksTwitterId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.DeleteWebhooksTwitterId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhooksTwitterIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetaWebhooksAlmanax

> GetMetaWebhooksTwitter200Response GetMetaWebhooksAlmanax(ctx).Execute()

Get Almanax Hook Metainfo



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.GetMetaWebhooksAlmanax(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GetMetaWebhooksAlmanax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetaWebhooksAlmanax`: GetMetaWebhooksTwitter200Response
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GetMetaWebhooksAlmanax`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetaWebhooksAlmanaxRequest struct via the builder pattern


### Return type

[**GetMetaWebhooksTwitter200Response**](GetMetaWebhooksTwitter200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetaWebhooksRss

> GetMetaWebhooksTwitter200Response GetMetaWebhooksRss(ctx).Execute()

Get RSS Hook Metainfo



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.GetMetaWebhooksRss(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GetMetaWebhooksRss``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetaWebhooksRss`: GetMetaWebhooksTwitter200Response
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GetMetaWebhooksRss`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetaWebhooksRssRequest struct via the builder pattern


### Return type

[**GetMetaWebhooksTwitter200Response**](GetMetaWebhooksTwitter200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetaWebhooksTwitter

> GetMetaWebhooksTwitter200Response GetMetaWebhooksTwitter(ctx).Execute()

Get Twitter Hook Metainfo



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.GetMetaWebhooksTwitter(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GetMetaWebhooksTwitter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetaWebhooksTwitter`: GetMetaWebhooksTwitter200Response
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GetMetaWebhooksTwitter`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetaWebhooksTwitterRequest struct via the builder pattern


### Return type

[**GetMetaWebhooksTwitter200Response**](GetMetaWebhooksTwitter200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooksAlmanaxId

> AlmanaxWebhook GetWebhooksAlmanaxId(ctx, id).Execute()

Get Almanax Hook



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.GetWebhooksAlmanaxId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GetWebhooksAlmanaxId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhooksAlmanaxId`: AlmanaxWebhook
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GetWebhooksAlmanaxId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksAlmanaxIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlmanaxWebhook**](AlmanaxWebhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooksRssId

> RssWebhook GetWebhooksRssId(ctx, id).Execute()

Get RSS Hook



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.GetWebhooksRssId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GetWebhooksRssId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhooksRssId`: RssWebhook
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GetWebhooksRssId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksRssIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RssWebhook**](RssWebhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooksTwitterId

> TwitterWebhook GetWebhooksTwitterId(ctx, id).Execute()

Get Twitter Hook



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.GetWebhooksTwitterId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GetWebhooksTwitterId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhooksTwitterId`: TwitterWebhook
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GetWebhooksTwitterId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksTwitterIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TwitterWebhook**](TwitterWebhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWebhooksAlmanax

> PostWebhooksAlmanax(ctx).CreateAlmanaxWebhook(createAlmanaxWebhook).Execute()

Register Almanax Hook



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
    createAlmanaxWebhook := *openapiclient.NewCreateAlmanaxWebhook([]string{"Subscriptions_example"}, "Format_example", "https://discord.com/api/webhooks/XYZ", []string{"Intervals_example"}) // CreateAlmanaxWebhook |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.PostWebhooksAlmanax(context.Background()).CreateAlmanaxWebhook(createAlmanaxWebhook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.PostWebhooksAlmanax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWebhooksAlmanaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAlmanaxWebhook** | [**CreateAlmanaxWebhook**](CreateAlmanaxWebhook.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWebhooksRss

> PostWebhooksRss(ctx).CreateRSSWebhook(createRSSWebhook).Execute()

Register RSS Hook



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
    createRSSWebhook := *openapiclient.NewCreateRSSWebhook([]string{"Subscriptions_example"}, "Format_example", "https://discord.com/api/webhooks/XYZ") // CreateRSSWebhook |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.PostWebhooksRss(context.Background()).CreateRSSWebhook(createRSSWebhook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.PostWebhooksRss``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWebhooksRssRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRSSWebhook** | [**CreateRSSWebhook**](CreateRSSWebhook.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWebhooksTwitter

> PostWebhooksTwitter(ctx).CreateTwitterWebhook(createTwitterWebhook).Execute()

Register Twitter Hook



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
    createTwitterWebhook := *openapiclient.NewCreateTwitterWebhook([]string{"Subscriptions_example"}, "Format_example", "https://discord.com/api/webhooks/XYZ") // CreateTwitterWebhook |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.PostWebhooksTwitter(context.Background()).CreateTwitterWebhook(createTwitterWebhook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.PostWebhooksTwitter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWebhooksTwitterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTwitterWebhook** | [**CreateTwitterWebhook**](CreateTwitterWebhook.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutWebhooksAlmanaxId

> AlmanaxWebhook PutWebhooksAlmanaxId(ctx, id).PutAlmanaxWebhook(putAlmanaxWebhook).Execute()

Update Almanax Hook



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
    id := "id_example" // string | 
    putAlmanaxWebhook := *openapiclient.NewPutAlmanaxWebhook() // PutAlmanaxWebhook |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.PutWebhooksAlmanaxId(context.Background(), id).PutAlmanaxWebhook(putAlmanaxWebhook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.PutWebhooksAlmanaxId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutWebhooksAlmanaxId`: AlmanaxWebhook
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.PutWebhooksAlmanaxId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutWebhooksAlmanaxIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putAlmanaxWebhook** | [**PutAlmanaxWebhook**](PutAlmanaxWebhook.md) |  | 

### Return type

[**AlmanaxWebhook**](AlmanaxWebhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutWebhooksRssId

> RssWebhook PutWebhooksRssId(ctx, id).PutRSSWebhook(putRSSWebhook).Execute()

Update RSS Hook



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
    id := "id_example" // string | 
    putRSSWebhook := *openapiclient.NewPutRSSWebhook() // PutRSSWebhook |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.PutWebhooksRssId(context.Background(), id).PutRSSWebhook(putRSSWebhook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.PutWebhooksRssId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutWebhooksRssId`: RssWebhook
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.PutWebhooksRssId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutWebhooksRssIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putRSSWebhook** | [**PutRSSWebhook**](PutRSSWebhook.md) |  | 

### Return type

[**RssWebhook**](RssWebhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutWebhooksTwitterId

> TwitterWebhook PutWebhooksTwitterId(ctx, id).PutTwitterWebhook(putTwitterWebhook).Execute()

Update Twitter Hook



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
    id := "id_example" // string | 
    putTwitterWebhook := *openapiclient.NewPutTwitterWebhook() // PutTwitterWebhook |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.PutWebhooksTwitterId(context.Background(), id).PutTwitterWebhook(putTwitterWebhook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.PutWebhooksTwitterId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutWebhooksTwitterId`: TwitterWebhook
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.PutWebhooksTwitterId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutWebhooksTwitterIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putTwitterWebhook** | [**PutTwitterWebhook**](PutTwitterWebhook.md) |  | 

### Return type

[**TwitterWebhook**](TwitterWebhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

