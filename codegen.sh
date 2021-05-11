#!/usr/bin/env bash
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -ex

: ${PROTOC=/usr/bin/protoc}
rm -rf api/proto

mkdir -p api/proto/common
mkdir -p api/proto/hdds
mkdir -p api/proto/ozone
mkdir -p api/proto/datanode
mkdir -p api/proto/ratis

$PROTOC -I $(pwd)/proto $(pwd)/proto/Security.proto --go_out=/tmp/
mv /tmp/github.com/apache/ozone-go/api/proto/common/Security.pb.go api/proto/common/

$PROTOC -I $(pwd)/proto $(pwd)/proto/hdds.proto --go_out=/tmp/
mv /tmp/github.com/apache/ozone-go/api/proto/hdds/hdds.pb.go api/proto/hdds/

$PROTOC -I $(pwd)/proto $(pwd)/proto/DatanodeClientProtocol.proto --go_out=/tmp/
mv /tmp/github.com/apache/ozone-go/api/proto/datanode/DatanodeClientProtocol.pb.go api/proto/datanode/


$PROTOC -I $(pwd)/proto $(pwd)/proto/DatanodeClientProtocol.proto --go_out=plugins=grpc:/tmp/
mv /tmp/github.com/apache/ozone-go/api/proto/datanode/DatanodeClientProtocol.pb.go api/proto/datanode/

$PROTOC -I $(pwd)/proto $(pwd)/proto/OmClientProtocol.proto --go_out=/tmp/
mv /tmp/github.com/apache/ozone-go/api/proto/ozone/OmClientProtocol.pb.go api/proto/ozone/


$PROTOC -I $(pwd)/proto $(pwd)/proto/raft.proto --go_out=/tmp/
mv /tmp/github.com/apache/ozone-go/api/proto/ratis/raft.pb.go api/proto/ratis/


$PROTOC -I $(pwd)/proto $(pwd)/proto/ratis-grpc.proto --go_out=plugins=grpc:/tmp/
mv /tmp/github.com/apache/ozone-go/api/proto/ratis/ratis-grpc.pb.go api/proto/ratis/

