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

// [START firestore_generated_firestore_apiv1beta1_Client_Listen]

package main

import (
	"context"
	"io"

	firestore "cloud.google.com/go/firestore/apiv1beta1"
	firestorepb "google.golang.org/genproto/googleapis/firestore/v1beta1"
)

func main() {
	// import firestorepb "google.golang.org/genproto/googleapis/firestore/v1beta1"

	ctx := context.Background()
	c, err := firestore.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	stream, err := c.Listen(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	go func() {
		reqs := []*firestorepb.ListenRequest{
			// TODO: Create requests.
		}
		for _, req := range reqs {
			if err := stream.Send(req); err != nil {
				// TODO: Handle error.
			}
		}
		stream.CloseSend()
	}()
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			// TODO: handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

// [END firestore_generated_firestore_apiv1beta1_Client_Listen]