/*
dofusdude

# Open Ankama Developer Community The all-in-one toolbelt for your next Ankama related project.  ## Versions - [Dofus 2](https://docs.dofusdu.de/dofus2/) - [Dofus 3](https://docs.dofusdu.de/dofus3/)   - v1 [latest] (you are here)   ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save` - [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save` - [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo` - [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude` - [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Almanax Discord Integration** Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 3 Beta** from stable to bleeding edge by replacing /dofus3 with /dofus3beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_, _de_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Official Sources** generated from actual data from the game.  ... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h). 

API version: 1.0.0-rc.7
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


// AlmanaxAPIService AlmanaxAPI service
type AlmanaxAPIService service

type ApiGetAlmanaxDateRequest struct {
	ctx context.Context
	ApiService *AlmanaxAPIService
	language string
	date string
}

func (r ApiGetAlmanaxDateRequest) Execute() (*Almanax, *http.Response, error) {
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
func (a *AlmanaxAPIService) GetAlmanaxDate(ctx context.Context, language string, date string) ApiGetAlmanaxDateRequest {
	return ApiGetAlmanaxDateRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
		date: date,
	}
}

// Execute executes the request
//  @return Almanax
func (a *AlmanaxAPIService) GetAlmanaxDateExecute(r ApiGetAlmanaxDateRequest) (*Almanax, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Almanax
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlmanaxAPIService.GetAlmanaxDate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dofus2/{language}/almanax/{date}"
	localVarPath = strings.Replace(localVarPath, "{"+"language"+"}", url.PathEscape(parameterValueToString(r.language, "language")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"date"+"}", url.PathEscape(parameterValueToString(r.date, "date")), -1)

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
	ApiService *AlmanaxAPIService
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

func (r ApiGetAlmanaxRangeRequest) Execute() ([]Almanax, *http.Response, error) {
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
func (a *AlmanaxAPIService) GetAlmanaxRange(ctx context.Context, language string) ApiGetAlmanaxRangeRequest {
	return ApiGetAlmanaxRangeRequest{
		ApiService: a,
		ctx: ctx,
		language: language,
	}
}

// Execute executes the request
//  @return []Almanax
func (a *AlmanaxAPIService) GetAlmanaxRangeExecute(r ApiGetAlmanaxRangeRequest) ([]Almanax, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Almanax
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlmanaxAPIService.GetAlmanaxRange")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dofus2/{language}/almanax"
	localVarPath = strings.Replace(localVarPath, "{"+"language"+"}", url.PathEscape(parameterValueToString(r.language, "language")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterBonusType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[bonus_type]", r.filterBonusType, "form", "")
	}
	if r.rangeFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "range[from]", r.rangeFrom, "form", "")
	}
	if r.rangeTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "range[to]", r.rangeTo, "form", "")
	}
	if r.rangeSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "range[size]", r.rangeSize, "form", "")
	}
	if r.timezone != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timezone", r.timezone, "form", "")
	} else {
		var defaultValue string = "Europe/Paris"
		r.timezone = &defaultValue
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
