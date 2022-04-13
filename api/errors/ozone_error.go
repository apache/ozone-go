// Package errors is ozone error info package
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
package errors

import (
	"errors"
	"fmt"
)

var (
	// PermissionErr TODO
	PermissionErr = errors.New("permission denied")

	// FileExistErr TODO
	FileExistErr = errors.New("file/directory already exists")
	// FileNotExistErr TODO
	FileNotExistErr = errors.New("file/directory does not exist")
	// PathIsNotFileErr TODO
	PathIsNotFileErr = errors.New("path is not a file")
	// PathIsNotDirErr TODO
	PathIsNotDirErr = errors.New("is not a directory")
	// FileCreateErr TODO
	FileCreateErr = errors.New("create file error")
	// CopyToLocalErr TODO
	CopyToLocalErr = errors.New("copy to local file error")

	ChecksumWriteErr = errors.New("checksum write data error")

	// ClosedPipeErr TODO
	ClosedPipeErr = errors.New("io: read/write on closed pipe")

	// BufferFullErr TODO
	BufferFullErr = errors.New("limit buffer: buffer full")
	// StandaloneReceiveErr TODO
	StandaloneReceiveErr = errors.New("standalone receive error")
	// RatisReceiveErr TODO
	RatisReceiveErr = errors.New("standalone receive error")
)

// OzoneOMClientError TODO
type OzoneOMClientError struct {
	CmdType string
	Message string
	Status  string
	Success bool
}

// Error TODO
func (e *OzoneOMClientError) Error() string {
	return fmt.Sprintf("Cmd: %v, Success: %v, Message: %v, Status: %v", e.CmdType, e.Success, e.Message, e.Status)
}

// OzoneDataNodeClientError TODO
type OzoneDataNodeClientError struct {
	CmdType string
	Message string
	Result  string
}

// Error TODO
func (e *OzoneDataNodeClientError) Error() string {
	return fmt.Sprintf("Cmd: %v, Message: %v, Result: %v", e.CmdType, e.Message, e.Result)
}

// ChecksumMismatchError TODO
type ChecksumMismatchError struct {
	ChecksumType string
	From         []byte
	CompareTo    []byte
}

// Error TODO
func (e *ChecksumMismatchError) Error() string {
	return fmt.Sprintf("Checksum mismatch, type: %v, from: %v, to %v", e.ChecksumType, e.From, e.CompareTo)
}
