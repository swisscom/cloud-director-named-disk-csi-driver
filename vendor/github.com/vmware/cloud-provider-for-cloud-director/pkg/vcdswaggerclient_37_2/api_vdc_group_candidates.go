
/*
 * VMware Cloud Director OpenAPI
 *
 * VMware Cloud Director OpenAPI is a new API that is defined using the OpenAPI standards.<br/> This ReSTful API borrows some elements of the legacy VMware Cloud Director API and establishes new patterns for use as described below. <h4>Authentication</h4> Authentication and Authorization schemes are the same as those for the legacy APIs. You can authenticate using the JWT token via the <code>Authorization</code> header or specifying a session using <code>x-vcloud-authorization</code> (The latter form is deprecated). <h4>Operation Patterns</h4> This API follows the following general guidelines to establish a consistent CRUD pattern: <table> <tr>   <th>Operation</th><th>Description</th><th>Response Code</th><th>Response Content</th> </tr><tr>   <td>GET /items<td>Returns a paginated list of items<td>200<td>Response will include Navigational links to the items in the list. </tr><tr>   <td>POST /items<td>Returns newly created item<td>201<td>Content-Location header links to the newly created item </tr><tr>   <td>GET /items/urn<td>Returns an individual item<td>200<td>A single item using same data type as that included in list above </tr><tr>   <td>PUT /items/urn<td>Updates an individual item<td>200<td>Updated view of the item is returned </tr><tr>   <td>DELETE /items/urn<td>Deletes the item<td>204<td>No content is returned. </tr> </table> <h5>Asynchronous operations</h5> Asynchronous operations are determined by the server. In those cases, instead of responding as described above, the server responds with an HTTP Response code 202 and an empty body. The tracking task (which is the same task as all legacy API operations use) is linked via the URI provided in the <code>Location</code> header.<br/> All API calls can choose to service a request asynchronously or synchronously as determined by the server upon interpreting the request. Operations that choose to exhibit this dual behavior will have both options documented by specifying both response code(s) below. The caller must be prepared to handle responses to such API calls by inspecting the HTTP Response code. <h5>Error Conditions</h5> <b>All</b> operations report errors using the following error reporting rules: <ul>   <li>400: Bad Request - In event of bad request due to incorrect data or other user error</li>   <li>401: Bad Request - If user is unauthenticated or their session has expired</li>   <li>403: Forbidden - If the user is not authorized or the entity does not exist</li> </ul> <h4>OpenAPI Design Concepts and Principles</h4> <ul>   <li>IDs are full Uniform Resource Names (URNs).</li>   <li>OpenAPI's <code>Content-Type</code> is always <code>application/json</code></li>   <li>REST links are in the Link header.</li>   <ul>     <li>Multiple relationships for any link are represented by multiple values in a space-separated list.</li>     <li>Links have a custom VMware Cloud Director-specific &quot;model&quot; attribute that hints at the applicable data         type for the links.</li>     <li>title + rel + model attributes evaluates to a unique link.</li>     <li>Links follow Hypermedia as the Engine of Application State (HATEOAS) principles. Links are present if         certain operations are present and permitted for the user&quot;s current role and the state of the         referred entities.</li>   </ul>   <li>APIs follow a flat structure relying on cross-referencing other entities instead of the navigational style       used by the legacy VMware Cloud Director APIs.</li>   <li>Most endpoints that return a list support filtering and sorting similar to the query service in the legacy       VMware Cloud Director APIs.</li>   <li>Accept header must be included to specify the API version for the request similar to calls to existing legacy       VMware Cloud Director APIs.</li>   <li>Each feature has a version in the path element present in its URL.<br/>       <b>Note</b> API URL's without a version in their paths must be considered experimental.</li> </ul> 
 *
 * API version: 37.2
 * Contact: https://code.vmware.com/support
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"context"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type VdcGroupCandidatesApiService service

/*
VdcGroupCandidatesApiService Get a list of candidate vDCs that can be added to a single vDC Group in the networking context.
Get a list of candidate vDCs. Results can be filtered by the type of vDC group and contexts of vDC ID or network pool ID. Below are the supported contexts: &lt;ul&gt; &lt;li&gt;Local vDC Group&lt;code&gt;(_context&#x3D;&#x3D;LOCAL;_context&#x3D;&#x3D;orgVdcId)&lt;/code&gt; - Returns all vDCs that have the same network pool assigned to the org vDC. &lt;li&gt;Universal vDC Group backed by NSX-V network provider &lt;code&gt;(_context&#x3D;&#x3D;UNIVERSAL;_context&#x3D;&#x3D;orgVdcId;_context&#x3D;&#x3D;networkPoolUniversalId)&lt;/code&gt; - If the orgVdc is specified, returns all vDCs that have the same assigned universal network pool as that of the org vDC. A universal network pool is backed by a broadcast domain that can stretch to multiple VMware Cloud Director sites. If the org vDC is not specified or not found and the universal network pool is specified, return all vDCs whose network pool are backed by that universal network pool. This case supports multisite/associated organization queries where user wants to retrieve remote vDCs on different sites/associated organizations with same backing universal network pool. &lt;li&gt;Existing vDC Group&lt;code&gt;(_context&#x3D;&#x3D;vdcGroupId;_context&#x3D;&#x3D;networkPoolUniversalId)&lt;/code&gt; - If vDC Group is specified, return all vDCs that have the?same assigned network pool as that of the vDC Group and also all vDCs that are currently participating in the vDC Group. The assigned network pool can also be universal if the vDC Group type is UNIVERSAL. If vDC Group is not specified or not found and the universal network pool is specified, return all vDCs whose network pool are backed by that universal network pool. This case supports multisite/associated organization queries where user wants to retrieve remote vDCs on different sites/associated organizations with same backing universal network pool. &lt;/ul&gt; Note that multisite calls to get the candidate vDCs for a local vDC Group will only return the vDCs of the local site&#39;s associated organizations. Remote site&#39;s vDCs are not returned since the org vDC or the vDC group is not found there. 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param page Page to fetch, zero offset.
 * @param pageSize Results per page to fetch.
 * @param optional nil or *VdcGroupCandidatesApiGetNetworkingVdcGroupCandidatesOpts - Optional Parameters:
     * @param "Filter" (optional.String) -  Filter for a query.  FIQL format.
     * @param "SortAsc" (optional.String) -  Field to use for ascending sort
     * @param "SortDesc" (optional.String) -  Field to use for descending sort

@return NetworkingCandidateVdcs
*/

type VdcGroupCandidatesApiGetNetworkingVdcGroupCandidatesOpts struct { 
	Filter optional.String
	SortAsc optional.String
	SortDesc optional.String
}

func (a *VdcGroupCandidatesApiService) GetNetworkingVdcGroupCandidates(ctx context.Context, page int32, pageSize int32, localVarOptionals *VdcGroupCandidatesApiGetNetworkingVdcGroupCandidatesOpts) (NetworkingCandidateVdcs, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue NetworkingCandidateVdcs
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/1.0.0/vdcGroups/networkingCandidateVdcs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if page < 1 {
		return localVarReturnValue, nil, reportError("page must be greater than 1")
	}
	if pageSize < 0 {
		return localVarReturnValue, nil, reportError("pageSize must be greater than 0")
	}
	if pageSize > 128 {
		return localVarReturnValue, nil, reportError("pageSize must be less than 128")
	}

	if localVarOptionals != nil && localVarOptionals.Filter.IsSet() {
		localVarQueryParams.Add("filter", parameterToString(localVarOptionals.Filter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SortAsc.IsSet() {
		localVarQueryParams.Add("sortAsc", parameterToString(localVarOptionals.SortAsc.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SortDesc.IsSet() {
		localVarQueryParams.Add("sortDesc", parameterToString(localVarOptionals.SortDesc.Value(), ""))
	}
	localVarQueryParams.Add("page", parameterToString(page, ""))
	localVarQueryParams.Add("pageSize", parameterToString(pageSize, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json;version=37.2"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v NetworkingCandidateVdcs
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

