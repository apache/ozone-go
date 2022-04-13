// Package io TODO
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
package io

import (
    "math"
    dnClient "github.com/apache/ozone-go/api/datanodeclient"
)

// CommitWatcher TODO
type CommitWatcher struct {
    xceiverClient              dnClient.XceiverClientSpi
    CommitIndex2flushedDataMap map[uint64]int // k: put block logIndex    v: flushedChunkIndex
}

func (commitWatcher *CommitWatcher) watchForCommit(commitIndex uint64) (*dnClient.XceiverClientReply, error) {
    return commitWatcher.xceiverClient.WatchForCommit(commitIndex)
}

// UpdateCommitInfoMap TODO
func (commitWatcher *CommitWatcher) UpdateCommitInfoMap(commitIndex uint64, flushedChunkIndex int) {
    commitWatcher.CommitIndex2flushedDataMap[commitIndex] = flushedChunkIndex
}

// WatchOnFirstIndex TODO
func (commitWatcher *CommitWatcher) WatchOnFirstIndex() (*dnClient.XceiverClientReply, error) {
    if len(commitWatcher.CommitIndex2flushedDataMap) > 0 {
        min := uint64(math.MaxUint64)
        for index := range commitWatcher.CommitIndex2flushedDataMap {
            if index < min {
                min = index
            }
        }
        return commitWatcher.watchForCommit(min)
    } else {
        return &dnClient.XceiverClientReply{LogIndex: 0}, nil
    }
}

// WatchOnLastIndex TODO
func (commitWatcher *CommitWatcher) WatchOnLastIndex() (*dnClient.XceiverClientReply, error) {
    if len(commitWatcher.CommitIndex2flushedDataMap) > 0 {
        max := uint64(0)
        for index := range commitWatcher.CommitIndex2flushedDataMap {
            if index > max {
                max = index
            }
        }
        return commitWatcher.watchForCommit(max)
    } else {
        return &dnClient.XceiverClientReply{LogIndex: 0}, nil
	}
}
