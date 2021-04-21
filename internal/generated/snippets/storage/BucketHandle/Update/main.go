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

// [START storage_generated_storage_BucketHandle_Update]

package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
)

func main() {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: handle error.
	}
	// Enable versioning in the bucket, regardless of its previous value.
	attrs, err := client.Bucket("my-bucket").Update(ctx,
		storage.BucketAttrsToUpdate{VersioningEnabled: true})
	if err != nil {
		// TODO: handle error.
	}
	fmt.Println(attrs)
}

// [END storage_generated_storage_BucketHandle_Update]