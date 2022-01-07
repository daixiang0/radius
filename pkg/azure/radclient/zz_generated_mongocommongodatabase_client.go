//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package radclient

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// MongoComMongoDatabaseClient contains the methods for the MongoComMongoDatabase group.
// Don't use this type directly, use NewMongoComMongoDatabaseClient() instead.
type MongoComMongoDatabaseClient struct {
	ep string
	pl runtime.Pipeline
	subscriptionID string
}

// NewMongoComMongoDatabaseClient creates a new instance of MongoComMongoDatabaseClient with the specified values.
func NewMongoComMongoDatabaseClient(con *arm.Connection, subscriptionID string) *MongoComMongoDatabaseClient {
	return &MongoComMongoDatabaseClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a mongo.com.MongoDatabase resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MongoComMongoDatabaseClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, applicationName string, mongoDatabaseName string, parameters MongoDatabaseResource, options *MongoComMongoDatabaseBeginCreateOrUpdateOptions) (MongoComMongoDatabaseCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, applicationName, mongoDatabaseName, parameters, options)
	if err != nil {
		return MongoComMongoDatabaseCreateOrUpdatePollerResponse{}, err
	}
	result := MongoComMongoDatabaseCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("MongoComMongoDatabaseClient.CreateOrUpdate", "location", resp, 	client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return MongoComMongoDatabaseCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &MongoComMongoDatabaseCreateOrUpdatePoller {
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a mongo.com.MongoDatabase resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MongoComMongoDatabaseClient) createOrUpdate(ctx context.Context, resourceGroupName string, applicationName string, mongoDatabaseName string, parameters MongoDatabaseResource, options *MongoComMongoDatabaseBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, applicationName, mongoDatabaseName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	 return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *MongoComMongoDatabaseClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, mongoDatabaseName string, parameters MongoDatabaseResource, options *MongoComMongoDatabaseBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/radiusv3/Application/{applicationName}/mongo.com.MongoDatabase/{mongoDatabaseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if mongoDatabaseName == "" {
		return nil, errors.New("parameter mongoDatabaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mongoDatabaseName}", url.PathEscape(mongoDatabaseName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *MongoComMongoDatabaseClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes a mongo.com.MongoDatabase resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MongoComMongoDatabaseClient) BeginDelete(ctx context.Context, resourceGroupName string, applicationName string, mongoDatabaseName string, options *MongoComMongoDatabaseBeginDeleteOptions) (MongoComMongoDatabaseDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, applicationName, mongoDatabaseName, options)
	if err != nil {
		return MongoComMongoDatabaseDeletePollerResponse{}, err
	}
	result := MongoComMongoDatabaseDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("MongoComMongoDatabaseClient.Delete", "location", resp, 	client.pl, client.deleteHandleError)
	if err != nil {
		return MongoComMongoDatabaseDeletePollerResponse{}, err
	}
	result.Poller = &MongoComMongoDatabaseDeletePoller {
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a mongo.com.MongoDatabase resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MongoComMongoDatabaseClient) deleteOperation(ctx context.Context, resourceGroupName string, applicationName string, mongoDatabaseName string, options *MongoComMongoDatabaseBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, applicationName, mongoDatabaseName, options)
	if err != nil {
		return nil, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	 return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MongoComMongoDatabaseClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, mongoDatabaseName string, options *MongoComMongoDatabaseBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/radiusv3/Application/{applicationName}/mongo.com.MongoDatabase/{mongoDatabaseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if mongoDatabaseName == "" {
		return nil, errors.New("parameter mongoDatabaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mongoDatabaseName}", url.PathEscape(mongoDatabaseName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *MongoComMongoDatabaseClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets a mongo.com.MongoDatabase resource by name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MongoComMongoDatabaseClient) Get(ctx context.Context, resourceGroupName string, applicationName string, mongoDatabaseName string, options *MongoComMongoDatabaseGetOptions) (MongoComMongoDatabaseGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationName, mongoDatabaseName, options)
	if err != nil {
		return MongoComMongoDatabaseGetResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return MongoComMongoDatabaseGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MongoComMongoDatabaseGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MongoComMongoDatabaseClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, mongoDatabaseName string, options *MongoComMongoDatabaseGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/radiusv3/Application/{applicationName}/mongo.com.MongoDatabase/{mongoDatabaseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if mongoDatabaseName == "" {
		return nil, errors.New("parameter mongoDatabaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mongoDatabaseName}", url.PathEscape(mongoDatabaseName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MongoComMongoDatabaseClient) getHandleResponse(resp *http.Response) (MongoComMongoDatabaseGetResponse, error) {
	result := MongoComMongoDatabaseGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MongoDatabaseResource); err != nil {
		return MongoComMongoDatabaseGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *MongoComMongoDatabaseClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - List the mongo.com.MongoDatabase resources deployed in the application.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MongoComMongoDatabaseClient) List(ctx context.Context, resourceGroupName string, applicationName string, options *MongoComMongoDatabaseListOptions) (MongoComMongoDatabaseListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, applicationName, options)
	if err != nil {
		return MongoComMongoDatabaseListResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return MongoComMongoDatabaseListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MongoComMongoDatabaseListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *MongoComMongoDatabaseClient) listCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, options *MongoComMongoDatabaseListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/radiusv3/Application/{applicationName}/mongo.com.MongoDatabase"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MongoComMongoDatabaseClient) listHandleResponse(resp *http.Response) (MongoComMongoDatabaseListResponse, error) {
	result := MongoComMongoDatabaseListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MongoDatabaseList); err != nil {
		return MongoComMongoDatabaseListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *MongoComMongoDatabaseClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
