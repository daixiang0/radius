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

package converter

import (
	"encoding/json"

	v1 "github.com/radius-project/radius/pkg/armrpc/api/v1"
	"github.com/radius-project/radius/pkg/daprrp/api/v20231001preview"
	"github.com/radius-project/radius/pkg/daprrp/datamodel"
)

// ConfigurationStoreDataModelToVersioned converts a version-agnostic datamodel.DaprConfigurationStore to a versioned model based on the version
// string, returning an error if the version is not supported.
func ConfigurationStoreDataModelToVersioned(model *datamodel.DaprConfigurationStore, version string) (v1.VersionedModelInterface, error) {
	switch version {
	case v20231001preview.Version:
		versioned := &v20231001preview.DaprConfigurationStoreResource{}
		err := versioned.ConvertFrom(model)
		return versioned, err

	default:
		return nil, v1.ErrUnsupportedAPIVersion
	}
}

// ConfigurationStoreDataModelFromVersioned unmarshals a JSON byte slice into a versioned ConfigurationStore resource and converts it
// to a version-agnostic datamodel Configuration Store, returning an error if either operation fails.
func ConfigurationStoreDataModelFromVersioned(content []byte, version string) (*datamodel.DaprConfigurationStore, error) {
	switch version {
	case v20231001preview.Version:
		am := &v20231001preview.DaprConfigurationStoreResource{}
		if err := json.Unmarshal(content, am); err != nil {
			return nil, err
		}
		dm, err := am.ConvertTo()
		if err != nil {
			return nil, err
		}

		return dm.(*datamodel.DaprConfigurationStore), err

	default:
		return nil, v1.ErrUnsupportedAPIVersion
	}
}
