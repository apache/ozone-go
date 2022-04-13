// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package api

import (
    "bytes"
    dnClient "github.com/apache/ozone-go/api/datanodeclient"
    "github.com/apache/ozone-go/api/omclient"
    "testing"
)

func TestOzoneClient_GetKey(t *testing.T) {
    type fields struct {
        OmClient       *omclient.OmClient
        xceiverManager *dnClient.XceiverClientManager
    }
    type args struct {
        volume string
        bucket string
        key    string
    }
    tests := []struct {
        name            string
        fields          fields
        args            args
        wantDestination string
        wantErr         bool
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            ozoneClient := &OzoneClient{
                OmClient:       tt.fields.OmClient,
                xceiverManager: tt.fields.xceiverManager,
            }
            destination := &bytes.Buffer{}
            err := ozoneClient.GetKey(tt.args.volume, tt.args.bucket, tt.args.key, destination)
            if (err != nil) != tt.wantErr {
                t.Errorf("GetKey() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if gotDestination := destination.String(); gotDestination != tt.wantDestination {
                t.Errorf("GetKey() gotDestination = %v, want %v", gotDestination, tt.wantDestination)
			}
		})
	}
}
