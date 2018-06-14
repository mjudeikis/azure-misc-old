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
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/consumption/mgmt/2017-12-30-preview/consumption instead.
// BudgetsClient is the consumption management client provides access to consumption resources for Azure Enterprise
// Subscriptions.
type BudgetsClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/consumption/mgmt/2017-12-30-preview/consumption instead.
// NewBudgetsClient creates an instance of the BudgetsClient client.
func NewBudgetsClient(subscriptionID string, name string) BudgetsClient {
	return NewBudgetsClientWithBaseURI(DefaultBaseURI, subscriptionID, name)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/consumption/mgmt/2017-12-30-preview/consumption instead.
// NewBudgetsClientWithBaseURI creates an instance of the BudgetsClient client.
func NewBudgetsClientWithBaseURI(baseURI string, subscriptionID string, name string) BudgetsClient {
	return BudgetsClient{NewWithBaseURI(baseURI, subscriptionID, name)}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/consumption/mgmt/2017-12-30-preview/consumption instead.
// CreateOrUpdate the operation to create or update a budget. Update operation requires latest eTag to be set in the
// request mandatorily. You may obtain the latest eTag by performing a get operation. Create operation does not require
// eTag.
//
// parameters is parameters supplied to the Create Budget operation.
func (client BudgetsClient) CreateOrUpdate(ctx context.Context, parameters Budget) (result Budget, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.BudgetProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.BudgetProperties.Category", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.BudgetProperties.Amount", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.BudgetProperties.TimePeriod", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.BudgetProperties.TimePeriod.StartDate", Name: validation.Null, Rule: true, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("consumption.BudgetsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/consumption/mgmt/2017-12-30-preview/consumption instead.
// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client BudgetsClient) CreateOrUpdatePreparer(ctx context.Context, parameters Budget) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":           autorest.Encode("path", client.Name),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Consumption/budgets/{name}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/consumption/mgmt/2017-12-30-preview/consumption instead.
// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client BudgetsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/consumption/mgmt/2017-12-30-preview/consumption instead.
// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client BudgetsClient) CreateOrUpdateResponder(resp *http.Response) (result Budget, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/consumption/mgmt/2017-12-30-preview/consumption instead.
// Delete the operation to delete a budget.
func (client BudgetsClient) Delete(ctx context.Context) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/consumption/mgmt/2017-12-30-preview/consumption instead.
// DeletePreparer prepares the Delete request.
func (client BudgetsClient) DeletePreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":           autorest.Encode("path", client.Name),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Consumption/budgets/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/consumption/mgmt/2017-12-30-preview/consumption instead.
// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client BudgetsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/consumption/mgmt/2017-12-30-preview/consumption instead.
// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client BudgetsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/consumption/mgmt/2017-12-30-preview/consumption instead.
// Get gets the budget for a subscription by budget name.
func (client BudgetsClient) Get(ctx context.Context) (result Budget, err error) {
	req, err := client.GetPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/consumption/mgmt/2017-12-30-preview/consumption instead.
// GetPreparer prepares the Get request.
func (client BudgetsClient) GetPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":           autorest.Encode("path", client.Name),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Consumption/budgets/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/consumption/mgmt/2017-12-30-preview/consumption instead.
// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client BudgetsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/consumption/mgmt/2017-12-30-preview/consumption instead.
// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client BudgetsClient) GetResponder(resp *http.Response) (result Budget, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/consumption/mgmt/2017-12-30-preview/consumption instead.
// List lists all budgets for a subscription.
func (client BudgetsClient) List(ctx context.Context) (result BudgetsListResult, err error) {
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "List", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/consumption/mgmt/2017-12-30-preview/consumption instead.
// ListPreparer prepares the List request.
func (client BudgetsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Consumption/budgets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/consumption/mgmt/2017-12-30-preview/consumption instead.
// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client BudgetsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/consumption/mgmt/2017-12-30-preview/consumption instead.
// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client BudgetsClient) ListResponder(resp *http.Response) (result BudgetsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
