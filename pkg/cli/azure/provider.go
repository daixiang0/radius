/*
Copyright 2023 The Radius Authors.

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

package azure

// AzureCredentialKind - Azure credential kinds supported.
type AzureCredentialKind string

const (
	// ProviderDisplayName is the text used in display for Azure.
	ProviderDisplayName                                     = "Azure"
	AzureCredentialKindWorkloadIdentity AzureCredentialKind = "WorkloadIdentity"
	AzureCredentialKindServicePrincipal AzureCredentialKind = "ServicePrincipal"
)

// Provider specifies the properties required to configure Azure provider for cloud resources
type Provider struct {
	SubscriptionID   string
	ResourceGroup    string
	CredentialKind   AzureCredentialKind
	WorkloadIdentity *WorkloadIdentityCredential
	ServicePrincipal *ServicePrincipalCredential
}

// WorkloadIdentityCredential specifies the properties of an Azure service principal
type WorkloadIdentityCredential struct {
	ClientID string
	TenantID string
}

// ServicePrincipal specifies the properties of an Azure service principal
type ServicePrincipalCredential struct {
	ClientID     string
	ClientSecret string
	TenantID     string
}
