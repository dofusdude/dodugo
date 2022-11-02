/*
Dofusdude

# A project for you - the developer. The free, always-up-to-date, low-latency, insert-buzzword-here Ankama API for your next cool project!  ## Client SDKs Don't write types or functions yourself - I already (kinda) did! 😉 - [Javascript](https://github.com/dofusdude/dofusdude-js) npm i dofusdude-js --save - [Typescript](https://github.com/dofusdude/dofusdude-ts) npm i dofusdude-ts --save - [Go](https://github.com/dofusdude/dodugo) go get -u github.com/dofusdude/dodugo - [Python](https://github.com/dofusdude/dofusdude-py) pip install dofusdude  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 2 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Discord Integration** Ankama related Twitter, RSS and Almanax feeds to post to Discord servers with advanced features like filters or mentions. Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 2 Beta** from stable to bleeding edge by replacing /dofus2 with /dofus2beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_ including the dropped languages from the Dofus website _de_ and _it_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing by the powerful [Meilisearch](https://www.meilisearch.com) written in Rust.  - 🕵️ **Complete** actual data from the game including items invisible to the encyclopedia like quest items.  - 🖼️ **HD Images** rendering vector graphics into PNGs up to 800x800 px in the background.   ## Current state - Weapons ✅ - Equipment ✅ - Sets ✅ - Resources ✅ - Consumables ✅ - Pets ✅ - Mounts ✅ - Cosmetics/Ceremonial Items ✅ - Harnesses ✅ - Quest Items ✅ - Almanax ✅ - Monsters ❌ - Spells ❌  ... and much more on the Roadmap on my Discord.   ## Deploy now. Use forever. Everything you see here on this site, you can use now and forever. Updates could introduce new fields, new paths or parameter but never break backwards compatibility, so no field or parameter will be deleted.  There is one exception! **The API will _always_ choose being up-to-date over everything else**. So if Ankama decides to drop languages from the game like they did with their website, the API will loose support for them, too.  ## Only the beginning... 🤯 I want this project to be useful and not just add plain GET-categories no one needs.  There is a long list of features I want to add (see the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h)). But they are all focussed on you, the developers. So please let me know what you need. I will change the list based on demand.  # Get started! 🥳 Scroll down and try it for yourself!  Or see how these other awesome projects use it: - [KaellyBot](https://github.com/Kaysoro/KaellyBot) by Kaysoro - [Dofus Craftlist](https://dofuscraftlist-dev.netlify.app) by Lystina - [AlmanaxApp](https://almanaxapp.netlify.app) by Lystina - [luwnarya.fr](https://luwnarya.fr)  I highly recommend using the SDKs for quick results. I use them myself for microservices for the API.  ## Thank you! I highly welcome everyone on my [Discord](https://discord.gg/3EtHskZD8h) to just talk about projects and use cases or give feedback of any kind.  The servers have a fixed monthly cost to provide very fast responses. If you want to help me keeping them running or simply  donate, consider becoming a [GitHub Sponsor](https://github.com/sponsors/dofusdude).  

API version: 0.7.0
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// QuestItemsApiService QuestItemsApi service
type QuestItemsApiService service

type ApiGetAllItemsQuestListRequest struct {
	ctx context.Context
	ApiService *QuestItemsApiService
	language string
	game string
	sortLevel *string
	filterTypeName *string
	filterMinLevel *int32
	filterMaxLevel *int32
	acceptEncoding *string
}

// sort the resulting list by level, default unsorted
func (r ApiGetAllItemsQuestListRequest) SortLevel(sortLevel string) ApiGetAllItemsQuestListRequest {
	r.sortLevel = &sortLevel
	return r
}

// only results with the translated type name
func (r ApiGetAllItemsQuestListRequest) FilterTypeName(filterTypeName string) ApiGetAllItemsQuestListRequest {
	r.filterTypeName = &filterTypeName
	return r
}

// only results which level is equal or above this value
func (r ApiGetAllItemsQuestListRequest) FilterMinLevel(filterMinLevel int32) ApiGetAllItemsQuestListRequest {
	r.filterMinLevel = &filterMinLevel
	return r
}

// only results which level is equal or below this value
func (r ApiGetAllItemsQuestListRequest) FilterMaxLevel(filterMaxLevel int32) ApiGetAllItemsQuestListRequest {
	r.filterMaxLevel = &filterMaxLevel
	return r
}

// optional compression for saving bandwidth
func (r ApiGetAllItemsQuestListRequest) AcceptEncoding(acceptEncoding string) ApiGetAllItemsQuestListRequest {
	r.acceptEncoding = &acceptEncoding
	return r
}

func (r ApiGetAllItemsQuestListRequest) Execute() (*ItemsListPaged, *http.Response, error) {
	return r.ApiService.GetAllItemsQuestListExecute(r)
}

/*
GetAllItemsQuestList List All Quest Items

Retrieve all quest items with one request. This endpoint is just an alias for the a list with disabled pagination (page[size]=-1) and all fields[type] set.

If you want everything unfiltered, delete the other query parameters.

Be careful with testing or (god forbid) using /all in your browser, the returned json is huge and will slow down the browser!

Tip: set the HTTP Header 'Accept-Encoding: gzip' for saving bandwidth. You will need to uncompress it on your end.
Example with cURL:
```
curl -sH 'Accept-Encoding: gzip' <api-endpoint> | gunzip -
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param game
 @return ApiGetAllItemsQuestListRequest
*/
func (a *QuestItemsApiService) GetAllItemsQuestList(ctx context.Context, language string, game string) ApiGetAllItemsQuestListRequest {
	return ApiGetAllItemsQuestListRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		game: game,
	}
}

// Execute executes the request
//  @return ItemsListPaged
func (a *QuestItemsApiService) GetAllItemsQuestListExecute(r ApiGetAllItemsQuestListRequest) (*ItemsListPaged, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ItemsListPaged
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuestItemsApiService.GetAllItemsQuestList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/{language}/items/quest/all"
	localVarPath = strings.Replace(localVarPath, "{"+"language"+"}", url.PathEscape(parameterToString(r.language, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"game"+"}", url.PathEscape(parameterToString(r.game, "")), -1)

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
		localVarQueryParams.Add("sort[level]", parameterToString(*r.sortLevel, ""))
	}
	if r.filterTypeName != nil {
		localVarQueryParams.Add("filter[type_name]", parameterToString(*r.filterTypeName, ""))
	}
	if r.filterMinLevel != nil {
		localVarQueryParams.Add("filter[min_level]", parameterToString(*r.filterMinLevel, ""))
	}
	if r.filterMaxLevel != nil {
		localVarQueryParams.Add("filter[max_level]", parameterToString(*r.filterMaxLevel, ""))
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
		localVarHeaderParams["Accept-Encoding"] = parameterToString(*r.acceptEncoding, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

type ApiGetItemQuestSingleRequest struct {
	ctx context.Context
	ApiService *QuestItemsApiService
	language string
	ankamaId int32
	game string
}

func (r ApiGetItemQuestSingleRequest) Execute() (*Resource, *http.Response, error) {
	return r.ApiService.GetItemQuestSingleExecute(r)
}

/*
GetItemQuestSingle Single Quest Items

Retrieve a specific quest item with id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param ankamaId identifier
 @param game
 @return ApiGetItemQuestSingleRequest
*/
func (a *QuestItemsApiService) GetItemQuestSingle(ctx context.Context, language string, ankamaId int32, game string) ApiGetItemQuestSingleRequest {
	return ApiGetItemQuestSingleRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		ankamaId: ankamaId,
		game: game,
	}
}

// Execute executes the request
//  @return Resource
func (a *QuestItemsApiService) GetItemQuestSingleExecute(r ApiGetItemQuestSingleRequest) (*Resource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Resource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuestItemsApiService.GetItemQuestSingle")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/{language}/items/quest/{ankama_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"language"+"}", url.PathEscape(parameterToString(r.language, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ankama_id"+"}", url.PathEscape(parameterToString(r.ankamaId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"game"+"}", url.PathEscape(parameterToString(r.game, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

type ApiGetItemsQuestListRequest struct {
	ctx context.Context
	ApiService *QuestItemsApiService
	language string
	game string
	sortLevel *string
	filterTypeName *string
	filterMinLevel *int32
	filterMaxLevel *int32
	pageSize *int32
	pageNumber *int32
	fieldsItem *[]string
}

// sort the resulting list by level, default unsorted
func (r ApiGetItemsQuestListRequest) SortLevel(sortLevel string) ApiGetItemsQuestListRequest {
	r.sortLevel = &sortLevel
	return r
}

// only results with the translated type name
func (r ApiGetItemsQuestListRequest) FilterTypeName(filterTypeName string) ApiGetItemsQuestListRequest {
	r.filterTypeName = &filterTypeName
	return r
}

// only results which level is equal or above this value
func (r ApiGetItemsQuestListRequest) FilterMinLevel(filterMinLevel int32) ApiGetItemsQuestListRequest {
	r.filterMinLevel = &filterMinLevel
	return r
}

// only results which level is equal or below this value
func (r ApiGetItemsQuestListRequest) FilterMaxLevel(filterMaxLevel int32) ApiGetItemsQuestListRequest {
	r.filterMaxLevel = &filterMaxLevel
	return r
}

// size of the results from the list. -1 disables pagination and gets all in one response.
func (r ApiGetItemsQuestListRequest) PageSize(pageSize int32) ApiGetItemsQuestListRequest {
	r.pageSize = &pageSize
	return r
}

// page number based on the current page[size]. So you could get page 1 with 8 entrys and page 2 would have entries 8 to 16.
func (r ApiGetItemsQuestListRequest) PageNumber(pageNumber int32) ApiGetItemsQuestListRequest {
	r.pageNumber = &pageNumber
	return r
}

// adds fields from their detail endpoint to the item list entries. Multiple comma separated values allowed.
func (r ApiGetItemsQuestListRequest) FieldsItem(fieldsItem []string) ApiGetItemsQuestListRequest {
	r.fieldsItem = &fieldsItem
	return r
}

func (r ApiGetItemsQuestListRequest) Execute() (*ItemsListPaged, *http.Response, error) {
	return r.ApiService.GetItemsQuestListExecute(r)
}

/*
GetItemsQuestList List Quest Items

Retrieve a list of quest items.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param game
 @return ApiGetItemsQuestListRequest
*/
func (a *QuestItemsApiService) GetItemsQuestList(ctx context.Context, language string, game string) ApiGetItemsQuestListRequest {
	return ApiGetItemsQuestListRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		game: game,
	}
}

// Execute executes the request
//  @return ItemsListPaged
func (a *QuestItemsApiService) GetItemsQuestListExecute(r ApiGetItemsQuestListRequest) (*ItemsListPaged, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ItemsListPaged
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuestItemsApiService.GetItemsQuestList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/{language}/items/quest"
	localVarPath = strings.Replace(localVarPath, "{"+"language"+"}", url.PathEscape(parameterToString(r.language, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"game"+"}", url.PathEscape(parameterToString(r.game, "")), -1)

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
		localVarQueryParams.Add("sort[level]", parameterToString(*r.sortLevel, ""))
	}
	if r.filterTypeName != nil {
		localVarQueryParams.Add("filter[type_name]", parameterToString(*r.filterTypeName, ""))
	}
	if r.filterMinLevel != nil {
		localVarQueryParams.Add("filter[min_level]", parameterToString(*r.filterMinLevel, ""))
	}
	if r.filterMaxLevel != nil {
		localVarQueryParams.Add("filter[max_level]", parameterToString(*r.filterMaxLevel, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page[size]", parameterToString(*r.pageSize, ""))
	}
	if r.pageNumber != nil {
		localVarQueryParams.Add("page[number]", parameterToString(*r.pageNumber, ""))
	}
	if r.fieldsItem != nil {
		localVarQueryParams.Add("fields[item]", parameterToString(*r.fieldsItem, "csv"))
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

type ApiGetItemsQuestSearchRequest struct {
	ctx context.Context
	ApiService *QuestItemsApiService
	language string
	game string
	query *string
	filterTypeName *string
	filterMinLevel *int32
	filterMaxLevel *int32
}

// case sensitive search query
func (r ApiGetItemsQuestSearchRequest) Query(query string) ApiGetItemsQuestSearchRequest {
	r.query = &query
	return r
}

// only results with the translated type name
func (r ApiGetItemsQuestSearchRequest) FilterTypeName(filterTypeName string) ApiGetItemsQuestSearchRequest {
	r.filterTypeName = &filterTypeName
	return r
}

// only results which level is equal or above this value
func (r ApiGetItemsQuestSearchRequest) FilterMinLevel(filterMinLevel int32) ApiGetItemsQuestSearchRequest {
	r.filterMinLevel = &filterMinLevel
	return r
}

// only results which level is equal or below this value
func (r ApiGetItemsQuestSearchRequest) FilterMaxLevel(filterMaxLevel int32) ApiGetItemsQuestSearchRequest {
	r.filterMaxLevel = &filterMaxLevel
	return r
}

func (r ApiGetItemsQuestSearchRequest) Execute() ([]ItemListEntry, *http.Response, error) {
	return r.ApiService.GetItemsQuestSearchExecute(r)
}

/*
GetItemsQuestSearch Search Quest Items

Search in all names and descriptions of quest items with a query.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language a valid language code
 @param game
 @return ApiGetItemsQuestSearchRequest
*/
func (a *QuestItemsApiService) GetItemsQuestSearch(ctx context.Context, language string, game string) ApiGetItemsQuestSearchRequest {
	return ApiGetItemsQuestSearchRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		game: game,
	}
}

// Execute executes the request
//  @return []ItemListEntry
func (a *QuestItemsApiService) GetItemsQuestSearchExecute(r ApiGetItemsQuestSearchRequest) ([]ItemListEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ItemListEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuestItemsApiService.GetItemsQuestSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{game}/{language}/items/quest/search"
	localVarPath = strings.Replace(localVarPath, "{"+"language"+"}", url.PathEscape(parameterToString(r.language, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"game"+"}", url.PathEscape(parameterToString(r.game, "")), -1)

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

	localVarQueryParams.Add("query", parameterToString(*r.query, ""))
	if r.filterTypeName != nil {
		localVarQueryParams.Add("filter[type_name]", parameterToString(*r.filterTypeName, ""))
	}
	if r.filterMinLevel != nil {
		localVarQueryParams.Add("filter[min_level]", parameterToString(*r.filterMinLevel, ""))
	}
	if r.filterMaxLevel != nil {
		localVarQueryParams.Add("filter[max_level]", parameterToString(*r.filterMaxLevel, ""))
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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
