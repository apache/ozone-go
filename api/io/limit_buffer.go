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
    "bytes"
    ozoneError "github.com/apache/ozone-go/api/errors"
)

type limitBuffer struct {
    buffer *bytes.Buffer
    limit  uint64
}

func newLimitBuffer(limit uint64) *limitBuffer {
    buffer := bytes.NewBuffer(make([]byte, limit, limit))
    buffer.Reset()
    return &limitBuffer{buffer: buffer, limit: limit}
}

// Write TODO
func (lb *limitBuffer) Write(p []byte) (int, error) {
    if !lb.HasRemaining() {
        return 0, ozoneError.BufferFullErr
    }

    remaining := lb.Remaining()
    if uint64(len(p)) > remaining {
        return lb.buffer.Write(p[:remaining])
    }

    return lb.buffer.Write(p)
}

// Read TODO
func (lb *limitBuffer) Read(p []byte) (int, error) {
    return lb.buffer.Read(p)
}

// Bytes TODO
func (lb *limitBuffer) Bytes() []byte {
    return lb.buffer.Bytes()
}

// Remaining TODO
func (lb *limitBuffer) Remaining() uint64 {
    return lb.Limit() - lb.Len()
}

// HasRemaining TODO
func (lb *limitBuffer) HasRemaining() bool {
    return lb.Limit() > lb.Len()
}

// Reset TODO
func (lb *limitBuffer) Reset() {
    lb.buffer.Reset()
}

// Next TODO
func (lb *limitBuffer) Next(n int) []byte {
    return lb.buffer.Next(n)
}

// Len TODO
func (lb *limitBuffer) Len() uint64 {
    return uint64(lb.buffer.Len())
}

// Limit TODO
func (lb *limitBuffer) Limit() uint64 {
    return lb.limit
}

// Truncate TODO
func (lb *limitBuffer) Truncate(n int) {
    lb.buffer.Truncate(n)
}
