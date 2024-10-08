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

import "@typespec/openapi";

using OpenAPI;

@doc("This specifies a reference to a secret. Secrets are encrypted, often have fine-grained access control, auditing and are recommended to be used to hold sensitive data.")
model SecretReference {
  @doc("The ID of an Applications.Core/SecretStore resource containing sensitive data required for recipe execution.")
  source: string;

  @doc("The key for the secret in the secret store.")
  key: string;
}
