// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START cloudkms_generated_kms_apiv1_KeyManagementClient_GetCryptoKey]

package main

import (
	"context"

	kms "cloud.google.com/go/kms/apiv1"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
)

func main() {
	// import kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"

	ctx := context.Background()
	c, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &kmspb.GetCryptoKeyRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetCryptoKey(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END cloudkms_generated_kms_apiv1_KeyManagementClient_GetCryptoKey]