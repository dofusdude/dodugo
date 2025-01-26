/*
dofusdude

# Open Ankama Developer Community The all-in-one toolbelt for your next Ankama related project.  ## Versions - [Dofus 2](https://docs.dofusdu.de/dofus2/) - [Dofus 3](https://docs.dofusdu.de/dofus3/)   - v1 [latest] (you are here)   ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save` - [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save` - [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo` - [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude` - [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Almanax Discord Integration** Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 3 Beta** from stable to bleeding edge by replacing /dofus3 with /dofus3beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_, _de_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Official Sources** generated from actual data from the game.  ... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h). 

API version: 1.0.0
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// CosmeticsAPIService CosmeticsAPI service
type CosmeticsAPIService service

type ApiGetAllCosmeticsListRequest struct {
	ctx context.Context
	ApiService *CosmeticsAPIService
	language string
	game string
	sortLevel *string
	filterMinLevel *int32
	filterMaxLevel *int32
	acceptEncoding *string
	filterTypeNameId *[]string
}

// sort the resulting list by level, default unsorted
func (r ApiGetAllCosmeticsListRequest) SortLevel(sortLevel string) ApiGetAllCosmeticsListRequest {
	r.sortLevel = &sortLevel
	return r
}

// only results which level is equal or above this value
func (r ApiGetAllCosmeticsListRequest) FilterMinLevel(filterMinLevel int32) ApiGetAllCosmeticsListRequest {
	r.filterMinLevel = &filterMinLevel
	return r
}

// only results which level is equal or below this value
func (r ApiGetAllCosmeticsListRequest) FilterMaxLevel(filterMaxLevel int32) ApiGetAllCosmeticsListRequest {
	r.filterMaxLevel = &filterMaxLevel
	return r
}

// optional compression for saving bandwidth
func (r ApiGetAllCosmeticsListRequest) AcceptEncoding(acceptEncoding string) ApiGetAllCosmeticsListRequest {
	r.acceptEncoding = &acceptEncoding
	return r
}

// multi-filter results with the english type name. Add with \&quot;wood\&quot; or \&quot;+wood\&quot; and exclude with \&quot;-wood\&quot;.
func (r ApiGetAllCosmeticsListRequest) FilterTypeNameId(filterTypeNameId []string) ApiGetAllCosmeticsListRequest {
	r.filterTypeNameId = &filterTypeNameId
	return r
}

func (r ApiGetAllCosmeticsListRequest) Execute() (*ListItems, *http.Response, error) {
	return r.ApiService.GetAllCosmeticsListExecute(r)
}

/*
GetAllCosmeticsList List All Cosmetics

Retrieve all cosmetic items with one request. This endpoint is just an alias for the a list with disabled pagination (page[size]=-1) and all fields[type] set.

If you want everything unfiltered, delete the other query parameters.

Be careful with testing or (god forbid) using /all in your browser, the returned json is huge and will slow down the browser!

Tip: set the HTTP Header 'Accept-Encoding: gzip' for saving bandwidth. You will need to uncompress it on your end.
Example with cURL:
```
curl -sH 'Accept-Encoding: gzip' <api-endpoint> | gunzip -
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param game game main 'dofus3' or beta channel 'dofus3beta'
 @return ApiGetAllCosmeticsListRequest
*/
func (a *CosmeticsAPIService) GetAllCosmeticsList(ctx context.Context, language string, game string) ApiGetAllCosmeticsListRequest {
	return ApiGetAllCosmeticsListRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		game: game,
	}
}

// Execute executes the request
//  @return ListItems
func (a *CosmeticsAPIService) GetAllCosmeticsListExecute(r ApiGetAllCosmeticsListRequest) (*ListItems, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListItems
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CosmeticsAPIService.GetAllCosmeticsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/v1/{language}/items/cosmetics/all"
	localVarPath = strings.Replace(localVarPath, "{"+"language"+"}", url.PathEscape(parameterValueToString(r.language, "language")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"game"+"}", url.PathEscape(parameterValueToString(r.game, "game")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.language) < 2 {
		return localVarReturnValue, nil, reportError("language must have at least 2 elements")
	}
	if strlen(r.language) > 2 {
		return localVarReturnValue, nil, reportError("language must have less than 2 elements")
	}

	if r.sortLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort[level]", r.sortLevel, "form", "")
	}
	if r.filterMinLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[min_level]", r.filterMinLevel, "form", "")
	}
	if r.filterMaxLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[max_level]", r.filterMaxLevel, "form", "")
	}
	if r.filterTypeNameId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[type.name_id]", r.filterTypeNameId, "form", "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptEncoding != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Encoding", r.acceptEncoding, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCosmeticsListRequest struct {
	ctx context.Context
	ApiService *CosmeticsAPIService
	language string
	game string
	sortLevel *string
	filterMinLevel *int32
	filterMaxLevel *int32
	pageSize *int32
	pageNumber *int32
	fieldsItem *[]string
	filterTypeNameId *[]string
}

// sort the resulting list by level, default unsorted
func (r ApiGetCosmeticsListRequest) SortLevel(sortLevel string) ApiGetCosmeticsListRequest {
	r.sortLevel = &sortLevel
	return r
}

// only results which level is equal or above this value
func (r ApiGetCosmeticsListRequest) FilterMinLevel(filterMinLevel int32) ApiGetCosmeticsListRequest {
	r.filterMinLevel = &filterMinLevel
	return r
}

// only results which level is equal or below this value
func (r ApiGetCosmeticsListRequest) FilterMaxLevel(filterMaxLevel int32) ApiGetCosmeticsListRequest {
	r.filterMaxLevel = &filterMaxLevel
	return r
}

// size of the results from the list. -1 disables pagination and gets all in one response.
func (r ApiGetCosmeticsListRequest) PageSize(pageSize int32) ApiGetCosmeticsListRequest {
	r.pageSize = &pageSize
	return r
}

// page number based on the current page[size]. So you could get page 1 with 8 entrys and page 2 would have entries 8 to 16.
func (r ApiGetCosmeticsListRequest) PageNumber(pageNumber int32) ApiGetCosmeticsListRequest {
	r.pageNumber = &pageNumber
	return r
}

// adds fields from their detail endpoint to the item list entries. Multiple comma separated values allowed.
func (r ApiGetCosmeticsListRequest) FieldsItem(fieldsItem []string) ApiGetCosmeticsListRequest {
	r.fieldsItem = &fieldsItem
	return r
}

// multi-filter results with the english type name. Add with \&quot;wood\&quot; or \&quot;+wood\&quot; and exclude with \&quot;-wood\&quot;.
func (r ApiGetCosmeticsListRequest) FilterTypeNameId(filterTypeNameId []string) ApiGetCosmeticsListRequest {
	r.filterTypeNameId = &filterTypeNameId
	return r
}

func (r ApiGetCosmeticsListRequest) Execute() (*ListItems, *http.Response, error) {
	return r.ApiService.GetCosmeticsListExecute(r)
}

/*
GetCosmeticsList List Cosmetics

Retrieve a list of cosmetic items.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param game game main 'dofus3' or beta channel 'dofus3beta'
 @return ApiGetCosmeticsListRequest
*/
func (a *CosmeticsAPIService) GetCosmeticsList(ctx context.Context, language string, game string) ApiGetCosmeticsListRequest {
	return ApiGetCosmeticsListRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		game: game,
	}
}

// Execute executes the request
//  @return ListItems
func (a *CosmeticsAPIService) GetCosmeticsListExecute(r ApiGetCosmeticsListRequest) (*ListItems, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListItems
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CosmeticsAPIService.GetCosmeticsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/v1/{language}/items/cosmetics"
	localVarPath = strings.Replace(localVarPath, "{"+"language"+"}", url.PathEscape(parameterValueToString(r.language, "language")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"game"+"}", url.PathEscape(parameterValueToString(r.game, "game")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.language) < 2 {
		return localVarReturnValue, nil, reportError("language must have at least 2 elements")
	}
	if strlen(r.language) > 2 {
		return localVarReturnValue, nil, reportError("language must have less than 2 elements")
	}

	if r.sortLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort[level]", r.sortLevel, "form", "")
	}
	if r.filterMinLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[min_level]", r.filterMinLevel, "form", "")
	}
	if r.filterMaxLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[max_level]", r.filterMaxLevel, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page[size]", r.pageSize, "form", "")
	}
	if r.pageNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page[number]", r.pageNumber, "form", "")
	}
	if r.fieldsItem != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[item]", r.fieldsItem, "form", "csv")
	}
	if r.filterTypeNameId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[type.name_id]", r.filterTypeNameId, "form", "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCosmeticsSearchRequest struct {
	ctx context.Context
	ApiService *CosmeticsAPIService
	language string
	game string
	query *string
	filterMinLevel *int32
	filterMaxLevel *int32
	limit *int32
	filterTypeNameId *[]string
}

// case sensitive search query
func (r ApiGetCosmeticsSearchRequest) Query(query string) ApiGetCosmeticsSearchRequest {
	r.query = &query
	return r
}

// only results which level is equal or above this value
func (r ApiGetCosmeticsSearchRequest) FilterMinLevel(filterMinLevel int32) ApiGetCosmeticsSearchRequest {
	r.filterMinLevel = &filterMinLevel
	return r
}

// only results which level is equal or below this value
func (r ApiGetCosmeticsSearchRequest) FilterMaxLevel(filterMaxLevel int32) ApiGetCosmeticsSearchRequest {
	r.filterMaxLevel = &filterMaxLevel
	return r
}

// maximum number of returned results
func (r ApiGetCosmeticsSearchRequest) Limit(limit int32) ApiGetCosmeticsSearchRequest {
	r.limit = &limit
	return r
}

// multi-filter results with the english type name. Add with \&quot;wood\&quot; or \&quot;+wood\&quot; and exclude with \&quot;-wood\&quot;.
func (r ApiGetCosmeticsSearchRequest) FilterTypeNameId(filterTypeNameId []string) ApiGetCosmeticsSearchRequest {
	r.filterTypeNameId = &filterTypeNameId
	return r
}

func (r ApiGetCosmeticsSearchRequest) Execute() ([]ListItem, *http.Response, error) {
	return r.ApiService.GetCosmeticsSearchExecute(r)
}

/*
GetCosmeticsSearch Search Cosmetics

Search in all names and descriptions of cosmetic items with a query.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param game game main 'dofus3' or beta channel 'dofus3beta'
 @return ApiGetCosmeticsSearchRequest
*/
func (a *CosmeticsAPIService) GetCosmeticsSearch(ctx context.Context, language string, game string) ApiGetCosmeticsSearchRequest {
	return ApiGetCosmeticsSearchRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		game: game,
	}
}

// Execute executes the request
//  @return []ListItem
func (a *CosmeticsAPIService) GetCosmeticsSearchExecute(r ApiGetCosmeticsSearchRequest) ([]ListItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ListItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CosmeticsAPIService.GetCosmeticsSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/v1/{language}/items/cosmetics/search"
	localVarPath = strings.Replace(localVarPath, "{"+"language"+"}", url.PathEscape(parameterValueToString(r.language, "language")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"game"+"}", url.PathEscape(parameterValueToString(r.game, "game")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.language) < 2 {
		return localVarReturnValue, nil, reportError("language must have at least 2 elements")
	}
	if strlen(r.language) > 2 {
		return localVarReturnValue, nil, reportError("language must have less than 2 elements")
	}
	if r.query == nil {
		return localVarReturnValue, nil, reportError("query is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "query", r.query, "form", "")
	if r.filterMinLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[min_level]", r.filterMinLevel, "form", "")
	}
	if r.filterMaxLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[max_level]", r.filterMaxLevel, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 8
		r.limit = &defaultValue
	}
	if r.filterTypeNameId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[type.name_id]", r.filterTypeNameId, "form", "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCosmeticsSingleRequest struct {
	ctx context.Context
	ApiService *CosmeticsAPIService
	language string
	ankamaId int32
	game string
}

func (r ApiGetCosmeticsSingleRequest) Execute() (*Equipment, *http.Response, error) {
	return r.ApiService.GetCosmeticsSingleExecute(r)
}

/*
GetCosmeticsSingle Single Cosmetics

Retrieve a specific cosmetic item with id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param ankamaId identifier
 @param game game main 'dofus3' or beta channel 'dofus3beta'
 @return ApiGetCosmeticsSingleRequest
*/
func (a *CosmeticsAPIService) GetCosmeticsSingle(ctx context.Context, language string, ankamaId int32, game string) ApiGetCosmeticsSingleRequest {
	return ApiGetCosmeticsSingleRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		ankamaId: ankamaId,
		game: game,
	}
}

// Execute executes the request
//  @return Equipment
func (a *CosmeticsAPIService) GetCosmeticsSingleExecute(r ApiGetCosmeticsSingleRequest) (*Equipment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Equipment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CosmeticsAPIService.GetCosmeticsSingle")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/v1/{language}/items/cosmetics/{ankama_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"language"+"}", url.PathEscape(parameterValueToString(r.language, "language")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ankama_id"+"}", url.PathEscape(parameterValueToString(r.ankamaId, "ankamaId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"game"+"}", url.PathEscape(parameterValueToString(r.game, "game")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.language) < 2 {
		return localVarReturnValue, nil, reportError("language must have at least 2 elements")
	}
	if strlen(r.language) > 2 {
		return localVarReturnValue, nil, reportError("language must have less than 2 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
