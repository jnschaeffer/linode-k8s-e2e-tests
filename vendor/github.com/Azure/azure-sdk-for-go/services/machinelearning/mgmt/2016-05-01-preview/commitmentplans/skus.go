package commitmentplans

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
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2016-05-01-preview/commitmentplans instead.
// SkusClient is the these APIs allow end users to operate on Azure Machine Learning Commitment Plans resources and
// their child Commitment Association resources. They support CRUD operations for commitment plans, get and list
// operations for commitment associations, moving commitment associations between commitment plans, and retrieving
// commitment plan usage history.
type SkusClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2016-05-01-preview/commitmentplans instead.
// NewSkusClient creates an instance of the SkusClient client.
func NewSkusClient(subscriptionID string) SkusClient {
	return NewSkusClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2016-05-01-preview/commitmentplans instead.
// NewSkusClientWithBaseURI creates an instance of the SkusClient client.
func NewSkusClientWithBaseURI(baseURI string, subscriptionID string) SkusClient {
	return SkusClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2016-05-01-preview/commitmentplans instead.
// List lists the available commitment plan SKUs.
func (client SkusClient) List(ctx context.Context) (result SkuListResult, err error) {
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.SkusClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "commitmentplans.SkusClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.SkusClient", "List", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2016-05-01-preview/commitmentplans instead.
// ListPreparer prepares the List request.
func (client SkusClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearning/skus", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2016-05-01-preview/commitmentplans instead.
// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SkusClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2016-05-01-preview/commitmentplans instead.
// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SkusClient) ListResponder(resp *http.Response) (result SkuListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}