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
package omclient

import (
    "github.com/apache/ozone-go/api/config"
    "github.com/apache/ozone-go/api/proto/hdds"
    "reflect"
    "testing"
)

func TestGetReplication(t *testing.T) {
    tests := []struct {
        name string
        arg  string
        want *hdds.ReplicationFactor
    }{
        // TODO: Add test cases.
        {
            name: "get replication",
            arg:  "1",
            want: hdds.ReplicationFactor_ONE.Enum(),
        },
        {
            name: "get replication",
            arg:  "3",
            want: hdds.ReplicationFactor_THREE.Enum(),
        },
        {
            name: "get replication",
            arg:  "2",
            want: hdds.ReplicationFactor_THREE.Enum(),
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            config.OzoneConfig.Set(config.OZONE_REPLICATION_KEY, tt.arg)
            if got := GetReplication(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("GetReplication() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestGetReplicationType(t *testing.T) {
    config.OzoneConfig.Set(config.OZONE_REPLICATION_TYPE_KEY, "ratis")
    tests := []struct {
        name string
        arg  string
        want *hdds.ReplicationType
    }{
        // TODO: Add test cases.
        {
            name: "get replication type",
            arg:  "ratis",
            want: hdds.ReplicationType_RATIS.Enum(),
        },
        {
            name: "get replication type",
            arg:  "1",
            want: hdds.ReplicationType_RATIS.Enum(),
        },
        {
            name: "get replication type",
            arg:  "stand_alone",
            want: hdds.ReplicationType_STAND_ALONE.Enum(),
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            config.OzoneConfig.Set(config.OZONE_REPLICATION_TYPE_KEY, tt.arg)
            if got := GetReplicationType(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("GetReplicationType() = %v, want %v", got, tt.want)
			}
		})
	}
}
