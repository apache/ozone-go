// Package omclient TODO
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
    "strings"

    log "github.com/sirupsen/logrus"
)

// GetReplication TODO
func GetReplication() *hdds.ReplicationFactor {
    factor, err := config.OzoneConfig.GetReplicationFactor()
    if err != nil {
        log.Errorf("Get replication factor error: %s, set to default factor 3", err.Error())
        factor = 3
    }
    switch factor {
    case 1:
        return hdds.ReplicationFactor_ONE.Enum()
    case 3:
        return hdds.ReplicationFactor_THREE.Enum()
    default:
        return hdds.ReplicationFactor_THREE.Enum()
    }
}

// GetReplicationType TODO
func GetReplicationType() *hdds.ReplicationType {
    typ := strings.ToUpper(config.OzoneConfig.GetReplicationType())
    switch typ {
    case hdds.ReplicationType_RATIS.String():
        return hdds.ReplicationType_RATIS.Enum()
    case hdds.ReplicationType_STAND_ALONE.String():
        return hdds.ReplicationType_STAND_ALONE.Enum()
    default:
        return hdds.ReplicationType_RATIS.Enum()
	}
}
