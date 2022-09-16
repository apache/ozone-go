// Package common TODO
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
package common

import (
	"sync"
	"sync/atomic"

	"github.com/google/uuid"
)

// IdInterface TODO
type IdInterface interface {
	CallIdGetAndIncrement() uint64
	GetClientId() string
	GetTraceId() string
}

// OzoneId TODO
type OzoneId struct {
	callId uint64
}

// OId TODO
var OId *OzoneId

var once sync.Once

func init() {
	GetOId()
}

// GetOId get ozone id with singleton
func GetOId() *OzoneId {
	if OId == nil {
		once.Do(func() {
			OId = &OzoneId{
				callId: 0,
			}
		})
	}
	return OId
}

// CallIdGetAndIncrement get increment call id
func (o *OzoneId) CallIdGetAndIncrement() uint64 {
	return atomic.AddUint64(&o.callId, 1)
}

// GetClientId get uniq client id
func (o *OzoneId) GetClientId() string {
	return uuid.New().String()[0:16]
}

// GetTraceId  get uniq trace id
func (o *OzoneId) GetTraceId() string {
	return uuid.New().String()[0:16]
}
