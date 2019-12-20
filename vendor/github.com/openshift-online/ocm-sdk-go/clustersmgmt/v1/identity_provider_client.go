/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// IdentityProviderClient is the client of the 'identity_provider' resource.
//
// Manages a specific identity provider.
type IdentityProviderClient struct {
	transport http.RoundTripper
	path      string
	metric    string
}

// NewIdentityProviderClient creates a new client for the 'identity_provider'
// resource using the given transport to sned the requests and receive the
// responses.
func NewIdentityProviderClient(transport http.RoundTripper, path string, metric string) *IdentityProviderClient {
	client := new(IdentityProviderClient)
	client.transport = transport
	client.path = path
	client.metric = metric
	return client
}

// Delete creates a request for the 'delete' method.
//
// Deletes the identity provider.
func (c *IdentityProviderClient) Delete() *IdentityProviderDeleteRequest {
	request := new(IdentityProviderDeleteRequest)
	request.transport = c.transport
	request.path = c.path
	request.metric = c.metric
	return request
}

// Get creates a request for the 'get' method.
//
// Retrieves the details of the identity provider.
func (c *IdentityProviderClient) Get() *IdentityProviderGetRequest {
	request := new(IdentityProviderGetRequest)
	request.transport = c.transport
	request.path = c.path
	request.metric = c.metric
	return request
}

// IdentityProviderPollRequest is the request for the Poll method.
type IdentityProviderPollRequest struct {
	request    *IdentityProviderGetRequest
	interval   time.Duration
	statuses   []int
	predicates []func(interface{}) bool
}

// Parameter adds a query parameter to all the requests that will be used to retrieve the object.
func (r *IdentityProviderPollRequest) Parameter(name string, value interface{}) *IdentityProviderPollRequest {
	r.request.Parameter(name, value)
	return r
}

// Header adds a request header to all the requests that will be used to retrieve the object.
func (r *IdentityProviderPollRequest) Header(name string, value interface{}) *IdentityProviderPollRequest {
	r.request.Header(name, value)
	return r
}

// Interval sets the polling interval. This parameter is mandatory and must be greater than zero.
func (r *IdentityProviderPollRequest) Interval(value time.Duration) *IdentityProviderPollRequest {
	r.interval = value
	return r
}

// Status set the expected status of the response. Multiple values can be set calling this method
// multiple times. The response will be considered successful if the status is any of those values.
func (r *IdentityProviderPollRequest) Status(value int) *IdentityProviderPollRequest {
	r.statuses = append(r.statuses, value)
	return r
}

// Predicate adds a predicate that the response should satisfy be considered successful. Multiple
// predicates can be set calling this method multiple times. The response will be considered successful
// if all the predicates are satisfied.
func (r *IdentityProviderPollRequest) Predicate(value func(*IdentityProviderGetResponse) bool) *IdentityProviderPollRequest {
	r.predicates = append(r.predicates, func(response interface{}) bool {
		return value(response.(*IdentityProviderGetResponse))
	})
	return r
}

// StartContext starts the polling loop. Responses will be considered successful if the status is one of
// the values specified with the Status method and if all the predicates specified with the Predicate
// method return nil.
//
// The context must have a timeout or deadline, otherwise this method will immediately return an error.
func (r *IdentityProviderPollRequest) StartContext(ctx context.Context) (response *IdentityProviderPollResponse, err error) {
	result, err := helpers.PollContext(ctx, r.interval, r.statuses, r.predicates, r.task)
	if result != nil {
		response = &IdentityProviderPollResponse{
			response: result.(*IdentityProviderGetResponse),
		}
	}
	return
}

// task adapts the types of the request/response types so that they can be used with the generic
// polling function from the helpers package.
func (r *IdentityProviderPollRequest) task(ctx context.Context) (status int, result interface{}, err error) {
	response, err := r.request.SendContext(ctx)
	if response != nil {
		status = response.Status()
		result = response
	}
	return
}

// IdentityProviderPollResponse is the response for the Poll method.
type IdentityProviderPollResponse struct {
	response *IdentityProviderGetResponse
}

// Status returns the response status code.
func (r *IdentityProviderPollResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.response.Status()
}

// Header returns header of the response.
func (r *IdentityProviderPollResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.response.Header()
}

// Error returns the response error.
func (r *IdentityProviderPollResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.response.Error()
}

// Body returns the value of the 'body' parameter.
//
//
func (r *IdentityProviderPollResponse) Body() *IdentityProvider {
	return r.response.Body()
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *IdentityProviderPollResponse) GetBody() (value *IdentityProvider, ok bool) {
	return r.response.GetBody()
}

// Poll creates a request to repeatedly retrieve the object till the response has one of a given set
// of states and satisfies a set of predicates.
func (c *IdentityProviderClient) Poll() *IdentityProviderPollRequest {
	return &IdentityProviderPollRequest{
		request: c.Get(),
	}
}

// IdentityProviderDeleteRequest is the request for the 'delete' method.
type IdentityProviderDeleteRequest struct {
	transport http.RoundTripper
	path      string
	metric    string
	query     url.Values
	header    http.Header
}

// Parameter adds a query parameter.
func (r *IdentityProviderDeleteRequest) Parameter(name string, value interface{}) *IdentityProviderDeleteRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *IdentityProviderDeleteRequest) Header(name string, value interface{}) *IdentityProviderDeleteRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *IdentityProviderDeleteRequest) Send() (result *IdentityProviderDeleteResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *IdentityProviderDeleteRequest) SendContext(ctx context.Context) (result *IdentityProviderDeleteResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.SetHeader(r.header, r.metric)
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "DELETE",
		URL:    uri,
		Header: header,
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = new(IdentityProviderDeleteResponse)
	result.status = response.StatusCode
	result.header = response.Header
	if result.status >= 400 {
		result.err, err = errors.UnmarshalError(response.Body)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	return
}

// IdentityProviderDeleteResponse is the response for the 'delete' method.
type IdentityProviderDeleteResponse struct {
	status int
	header http.Header
	err    *errors.Error
}

// Status returns the response status code.
func (r *IdentityProviderDeleteResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *IdentityProviderDeleteResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *IdentityProviderDeleteResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// IdentityProviderGetRequest is the request for the 'get' method.
type IdentityProviderGetRequest struct {
	transport http.RoundTripper
	path      string
	metric    string
	query     url.Values
	header    http.Header
}

// Parameter adds a query parameter.
func (r *IdentityProviderGetRequest) Parameter(name string, value interface{}) *IdentityProviderGetRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *IdentityProviderGetRequest) Header(name string, value interface{}) *IdentityProviderGetRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *IdentityProviderGetRequest) Send() (result *IdentityProviderGetResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *IdentityProviderGetRequest) SendContext(ctx context.Context) (result *IdentityProviderGetResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.SetHeader(r.header, r.metric)
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "GET",
		URL:    uri,
		Header: header,
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = new(IdentityProviderGetResponse)
	result.status = response.StatusCode
	result.header = response.Header
	if result.status >= 400 {
		result.err, err = errors.UnmarshalError(response.Body)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = result.unmarshal(response.Body)
	if err != nil {
		return
	}
	return
}

// IdentityProviderGetResponse is the response for the 'get' method.
type IdentityProviderGetResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *IdentityProvider
}

// Status returns the response status code.
func (r *IdentityProviderGetResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *IdentityProviderGetResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *IdentityProviderGetResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *IdentityProviderGetResponse) Body() *IdentityProvider {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *IdentityProviderGetResponse) GetBody() (value *IdentityProvider, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// unmarshal is the method used internally to unmarshal responses to the
// 'get' method.
func (r *IdentityProviderGetResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(identityProviderData)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	r.body, err = data.unwrap()
	if err != nil {
		return err
	}
	return err
}