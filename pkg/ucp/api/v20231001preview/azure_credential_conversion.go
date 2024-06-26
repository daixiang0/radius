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

package v20231001preview

import (
	"fmt"

	v1 "github.com/radius-project/radius/pkg/armrpc/api/v1"
	"github.com/radius-project/radius/pkg/to"
	"github.com/radius-project/radius/pkg/ucp/datamodel"
)

const (
	// AzureCredentialType represents the ucp azure crendetial type value.
	AzureCredentialType = "System.Azure/credentials"
)

// ConvertTo converts from the versioned Credential resource to version-agnostic datamodel.
func (cr *AzureCredentialResource) ConvertTo() (v1.DataModelInterface, error) {
	prop, err := cr.getDataModelCredentialProperties()
	if err != nil {
		return nil, err
	}

	converted := &datamodel.AzureCredential{
		BaseResource: v1.BaseResource{
			TrackedResource: v1.TrackedResource{
				ID:       to.String(cr.ID),
				Name:     to.String(cr.Name),
				Type:     to.String(cr.Type),
				Location: to.String(cr.Location),
				Tags:     to.StringMap(cr.Tags),
			},
			InternalMetadata: v1.InternalMetadata{
				UpdatedAPIVersion: Version,
			},
		},
		Properties: prop,
	}

	return converted, nil
}

func (cr *AzureCredentialResource) getDataModelCredentialProperties() (*datamodel.AzureCredentialResourceProperties, error) {
	if cr.Properties == nil {
		return nil, &v1.ErrModelConversion{PropertyName: "$.properties", ValidValue: "not nil"}
	}

	switch p := cr.Properties.(type) {
	case *AzureServicePrincipalProperties:
		var storage *datamodel.CredentialStorageProperties

		switch c := p.Storage.(type) {
		case *InternalCredentialStorageProperties:
			if c.Kind == nil {
				return nil, &v1.ErrModelConversion{PropertyName: "$.properties", ValidValue: "not nil"}
			}
			storage = &datamodel.CredentialStorageProperties{
				Kind: datamodel.InternalStorageKind,
				InternalCredential: &datamodel.InternalCredentialStorageProperties{
					SecretName: to.String(c.SecretName),
				},
			}
		case nil:
			return nil, &v1.ErrModelConversion{PropertyName: "$.properties.storage", ValidValue: "not nil"}
		default:
			return nil, &v1.ErrModelConversion{PropertyName: "$.properties.storage.kind", ValidValue: fmt.Sprintf("one of %q", PossibleCredentialStorageKindValues())}
		}

		return &datamodel.AzureCredentialResourceProperties{
			Kind: datamodel.AzureServicePrincipalCredentialKind,
			AzureCredential: &datamodel.AzureCredentialProperties{
				Kind: datamodel.AzureServicePrincipalCredentialKind,
				ServicePrincipal: &datamodel.AzureServicePrincipalCredentialProperties{
					TenantID:     to.String(p.TenantID),
					ClientID:     to.String(p.ClientID),
					ClientSecret: to.String(p.ClientSecret),
				},
			},
			Storage: storage,
		}, nil
	case *AzureWorkloadIdentityProperties:
		var storage *datamodel.CredentialStorageProperties

		switch c := p.Storage.(type) {
		case *InternalCredentialStorageProperties:
			if c.Kind == nil {
				return nil, &v1.ErrModelConversion{PropertyName: "$.properties", ValidValue: "not nil"}
			}
			storage = &datamodel.CredentialStorageProperties{
				Kind: datamodel.InternalStorageKind,
				InternalCredential: &datamodel.InternalCredentialStorageProperties{
					SecretName: to.String(c.SecretName),
				},
			}
		case nil:
			return nil, &v1.ErrModelConversion{PropertyName: "$.properties.storage", ValidValue: "not nil"}
		default:
			return nil, &v1.ErrModelConversion{PropertyName: "$.properties.storage.kind", ValidValue: fmt.Sprintf("one of %q", PossibleCredentialStorageKindValues())}
		}

		return &datamodel.AzureCredentialResourceProperties{
			Kind: datamodel.AzureWorkloadIdentityCredentialKind,
			AzureCredential: &datamodel.AzureCredentialProperties{
				Kind: datamodel.AzureWorkloadIdentityCredentialKind,
				WorkloadIdentity: &datamodel.AzureWorkloadIdentityCredentialProperties{
					TenantID: to.String(p.TenantID),
					ClientID: to.String(p.ClientID),
				},
			},
			Storage: storage,
		}, nil
	default:
		return nil, v1.ErrInvalidModelConversion
	}
}

// ConvertFrom converts from version-agnostic datamodel to the versioned Credential resource.
func (dst *AzureCredentialResource) ConvertFrom(src v1.DataModelInterface) error {
	dm, ok := src.(*datamodel.AzureCredential)
	if !ok {
		return v1.ErrInvalidModelConversion
	}

	dst.ID = &dm.ID
	dst.Name = &dm.Name
	dst.Type = &dm.Type
	dst.Location = &dm.Location
	dst.Tags = *to.StringMapPtr(dm.Tags)

	var storage CredentialStoragePropertiesClassification
	switch dm.Properties.Storage.Kind {
	case datamodel.InternalStorageKind:
		storage = &InternalCredentialStorageProperties{
			Kind:       to.Ptr(CredentialStorageKindInternal),
			SecretName: to.Ptr(dm.Properties.Storage.InternalCredential.SecretName),
		}
	default:
		return v1.ErrInvalidModelConversion
	}

	// DO NOT convert any secret values to versioned model.
	switch dm.Properties.Kind {
	case datamodel.AzureServicePrincipalCredentialKind:
		if dm.Properties.AzureCredential.ServicePrincipal == nil {
			return v1.ErrInvalidModelConversion
		}
		dst.Properties = &AzureServicePrincipalProperties{
			Kind:     to.Ptr(AzureCredentialKind(dm.Properties.Kind)),
			ClientID: to.Ptr(dm.Properties.AzureCredential.ServicePrincipal.ClientID),
			TenantID: to.Ptr(dm.Properties.AzureCredential.ServicePrincipal.TenantID),
			Storage:  storage,
		}
	case datamodel.AzureWorkloadIdentityCredentialKind:
		if dm.Properties.AzureCredential.WorkloadIdentity == nil {
			return v1.ErrInvalidModelConversion
		}
		dst.Properties = &AzureWorkloadIdentityProperties{
			Kind:     to.Ptr(AzureCredentialKind(dm.Properties.Kind)),
			ClientID: to.Ptr(dm.Properties.AzureCredential.WorkloadIdentity.ClientID),
			TenantID: to.Ptr(dm.Properties.AzureCredential.WorkloadIdentity.TenantID),
			Storage:  storage,
		}
	default:
		return v1.ErrInvalidModelConversion
	}

	return nil
}
