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


// AlmanaxApiService AlmanaxApi service
type AlmanaxApiService service

type ApiGetAlmanaxDateRequest struct {
	ctx context.Context
	ApiService *AlmanaxApiService
	language string
	date string
}

func (r ApiGetAlmanaxDateRequest) Execute() (*AlmanaxEntry, *http.Response, error) {
	return r.ApiService.GetAlmanaxDateExecute(r)
}

/*
GetAlmanaxDate Single Almanax Date

Get a single date. There are not more details in the returned object than the normal range endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language code
 @param date yyyy-mm-dd
 @return ApiGetAlmanaxDateRequest
*/
func (a *AlmanaxApiService) GetAlmanaxDate(ctx context.Context, language string, date string) ApiGetAlmanaxDateRequest {
	return ApiGetAlmanaxDateRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		date: date,
	}
}

// Execute executes the request
//  @return AlmanaxEntry
func (a *AlmanaxApiService) GetAlmanaxDateExecute(r ApiGetAlmanaxDateRequest) (*AlmanaxEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlmanaxEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlmanaxApiService.GetAlmanaxDate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dofus2/{language}/almanax/{date}"
	localVarPath = strings.Replace(localVarPath, "{"+"language"+"}", url.PathEscape(parameterToString(r.language, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"date"+"}", url.PathEscape(parameterToString(r.date, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetAlmanaxRangeRequest struct {
	ctx context.Context
	ApiService *AlmanaxApiService
	language string
	filterBonusType *string
	rangeFrom *string
	rangeTo *string
	rangeSize *int32
	timezone *string
}

// ids from meta/{language}/almanax/bonuses
func (r ApiGetAlmanaxRangeRequest) FilterBonusType(filterBonusType string) ApiGetAlmanaxRangeRequest {
	r.filterBonusType = &filterBonusType
	return r
}

// yyyy-mm-dd
func (r ApiGetAlmanaxRangeRequest) RangeFrom(rangeFrom string) ApiGetAlmanaxRangeRequest {
	r.rangeFrom = &rangeFrom
	return r
}

// yyyy-mm-dd
func (r ApiGetAlmanaxRangeRequest) RangeTo(rangeTo string) ApiGetAlmanaxRangeRequest {
	r.rangeTo = &rangeTo
	return r
}

// size of the returned range
func (r ApiGetAlmanaxRangeRequest) RangeSize(rangeSize int32) ApiGetAlmanaxRangeRequest {
	r.rangeSize = &rangeSize
	return r
}

// determine what the current time is. If you live in Brazil, \&quot;today\&quot; will be hours apart from Paris. Use your timezone to get results relative to your location.
func (r ApiGetAlmanaxRangeRequest) Timezone(timezone string) ApiGetAlmanaxRangeRequest {
	r.timezone = &timezone
	return r
}

func (r ApiGetAlmanaxRangeRequest) Execute() ([]AlmanaxEntry, *http.Response, error) {
	return r.ApiService.GetAlmanaxRangeExecute(r)
}

/*
GetAlmanaxRange Almanax Range

Get a range of dates, defaults to today + 6 following days but can specified by the query parameters. 

filter[bonus_type] can be used seperately and does not have an effect on the other parameters.

range[from] changes the start date, everything else defaults to 6 following dates from this start date.

range[to] when used without anything else, it will use today as start date and this parameter as end. All ranges are inclusive.

range[from] + range[to] = inclusive range over the specified dates, should never be farther apart than 35 days.

range[from|to] + range[size] no need to specify the date, just following days with [from] (0 is today) or go backwards in time with only [to] and [size].

Not all combinations are listed but this should give you an idea how to they could work.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param language code
 @return ApiGetAlmanaxRangeRequest
*/
func (a *AlmanaxApiService) GetAlmanaxRange(ctx context.Context, language string) ApiGetAlmanaxRangeRequest {
	return ApiGetAlmanaxRangeRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
	}
}

// Execute executes the request
//  @return []AlmanaxEntry
func (a *AlmanaxApiService) GetAlmanaxRangeExecute(r ApiGetAlmanaxRangeRequest) ([]AlmanaxEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AlmanaxEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlmanaxApiService.GetAlmanaxRange")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dofus2/{language}/almanax"
	localVarPath = strings.Replace(localVarPath, "{"+"language"+"}", url.PathEscape(parameterToString(r.language, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterBonusType != nil {
		localVarQueryParams.Add("filter[bonus_type]", parameterToString(*r.filterBonusType, ""))
	}
	if r.rangeFrom != nil {
		localVarQueryParams.Add("range[from]", parameterToString(*r.rangeFrom, ""))
	}
	if r.rangeTo != nil {
		localVarQueryParams.Add("range[to]", parameterToString(*r.rangeTo, ""))
	}
	if r.rangeSize != nil {
		localVarQueryParams.Add("range[size]", parameterToString(*r.rangeSize, ""))
	}
	if r.timezone != nil {
		localVarQueryParams.Add("timezone", parameterToString(*r.timezone, ""))
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
