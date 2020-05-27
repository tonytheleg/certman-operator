package consumption

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// UsageDetailsClient is the consumption management client provides access to consumption resources for Azure
// Enterprise Subscriptions.
type UsageDetailsClient struct {
	BaseClient
}

// NewUsageDetailsClient creates an instance of the UsageDetailsClient client.
func NewUsageDetailsClient(subscriptionID string) UsageDetailsClient {
	return NewUsageDetailsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUsageDetailsClientWithBaseURI creates an instance of the UsageDetailsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewUsageDetailsClientWithBaseURI(baseURI string, subscriptionID string) UsageDetailsClient {
	return UsageDetailsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists the usage details for a scope by current billing period. Usage details are available via this API only
// for May 1, 2014 or later.
// Parameters:
// expand - may be used to expand the properties/additionalProperties or properties/meterDetails within a list
// of usage details. By default, these fields are not included when listing usage details.
// filter - may be used to filter usageDetails by properties/usageEnd (Utc time), properties/usageStart (Utc
// time), properties/resourceGroup, properties/instanceName, properties/instanceId or tags. The filter supports
// 'eq', 'lt', 'gt', 'le', 'ge', and 'and'. It does not currently support 'ne', 'or', or 'not'. Tag filter is a
// key value pair string where key and value is separated by a colon (:).
// skiptoken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls.
// top - may be used to limit the number of results to the most recent N usageDetails.
// apply - oData apply expression to aggregate usageDetails by tags or (tags and properties/usageStart)
func (client UsageDetailsClient) List(ctx context.Context, expand string, filter string, skiptoken string, top *int32, apply string) (result UsageDetailsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsageDetailsClient.List")
		defer func() {
			sc := -1
			if result.udlr.Response.Response != nil {
				sc = result.udlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("consumption.UsageDetailsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, expand, filter, skiptoken, top, apply)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.udlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "List", resp, "Failure sending request")
		return
	}

	result.udlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client UsageDetailsClient) ListPreparer(ctx context.Context, expand string, filter string, skiptoken string, top *int32, apply string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(apply) > 0 {
		queryParameters["$apply"] = autorest.Encode("query", apply)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Consumption/usageDetails", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client UsageDetailsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client UsageDetailsClient) ListResponder(resp *http.Response) (result UsageDetailsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client UsageDetailsClient) listNextResults(ctx context.Context, lastResults UsageDetailsListResult) (result UsageDetailsListResult, err error) {
	req, err := lastResults.usageDetailsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client UsageDetailsClient) ListComplete(ctx context.Context, expand string, filter string, skiptoken string, top *int32, apply string) (result UsageDetailsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsageDetailsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, expand, filter, skiptoken, top, apply)
	return
}

// ListByBillingPeriod lists the usage details for a scope by billing period. Usage details are available via this API
// only for May 1, 2014 or later.
// Parameters:
// billingPeriodName - billing Period Name.
// expand - may be used to expand the properties/additionalProperties or properties/meterDetails within a list
// of usage details. By default, these fields are not included when listing usage details.
// filter - may be used to filter usageDetails by properties/usageEnd (Utc time), properties/usageStart (Utc
// time), properties/resourceGroup, properties/instanceName or properties/instanceId. The filter supports 'eq',
// 'lt', 'gt', 'le', 'ge', and 'and'. It does not currently support 'ne', 'or', or 'not'. Tag filter is a key
// value pair string where key and value is separated by a colon (:).
// apply - oData apply expression to aggregate usageDetails by tags or (tags and properties/usageStart) for
// specified billing period
// skiptoken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls.
// top - may be used to limit the number of results to the most recent N usageDetails.
func (client UsageDetailsClient) ListByBillingPeriod(ctx context.Context, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result UsageDetailsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsageDetailsClient.ListByBillingPeriod")
		defer func() {
			sc := -1
			if result.udlr.Response.Response != nil {
				sc = result.udlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("consumption.UsageDetailsClient", "ListByBillingPeriod", err.Error())
	}

	result.fn = client.listByBillingPeriodNextResults
	req, err := client.ListByBillingPeriodPreparer(ctx, billingPeriodName, expand, filter, apply, skiptoken, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "ListByBillingPeriod", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingPeriodSender(req)
	if err != nil {
		result.udlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "ListByBillingPeriod", resp, "Failure sending request")
		return
	}

	result.udlr, err = client.ListByBillingPeriodResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "ListByBillingPeriod", resp, "Failure responding to request")
	}

	return
}

// ListByBillingPeriodPreparer prepares the ListByBillingPeriod request.
func (client UsageDetailsClient) ListByBillingPeriodPreparer(ctx context.Context, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingPeriodName": autorest.Encode("path", billingPeriodName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(apply) > 0 {
		queryParameters["$apply"] = autorest.Encode("query", apply)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}/providers/Microsoft.Consumption/usageDetails", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingPeriodSender sends the ListByBillingPeriod request. The method will close the
// http.Response Body if it receives an error.
func (client UsageDetailsClient) ListByBillingPeriodSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByBillingPeriodResponder handles the response to the ListByBillingPeriod request. The method always
// closes the http.Response Body.
func (client UsageDetailsClient) ListByBillingPeriodResponder(resp *http.Response) (result UsageDetailsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingPeriodNextResults retrieves the next set of results, if any.
func (client UsageDetailsClient) listByBillingPeriodNextResults(ctx context.Context, lastResults UsageDetailsListResult) (result UsageDetailsListResult, err error) {
	req, err := lastResults.usageDetailsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "listByBillingPeriodNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingPeriodSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "listByBillingPeriodNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingPeriodResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "listByBillingPeriodNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingPeriodComplete enumerates all values, automatically crossing page boundaries as required.
func (client UsageDetailsClient) ListByBillingPeriodComplete(ctx context.Context, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result UsageDetailsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsageDetailsClient.ListByBillingPeriod")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingPeriod(ctx, billingPeriodName, expand, filter, apply, skiptoken, top)
	return
}