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
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
)

func TestGetOzoneId(t *testing.T) {
	tests := []struct {
		name string
		want *OzoneId
	}{
		{
			name: "get oid",
			want: GetOId(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOId(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOzoneId_CallIdGetAndIncrement(t *testing.T) {
	type fields struct {
		callId uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		{
			name:   "call id",
			fields: fields{callId: 0},
			want:   1,
		},
	}
	oid := GetOId()
	patch := gomonkey.ApplyMethod(reflect.TypeOf(oid), "CallIdGetAndIncrement", func(_ *OzoneId) uint64 {
		return 1
	})
	defer patch.Reset()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oid.CallIdGetAndIncrement(); got != tt.want {
				t.Errorf("CallIdGetAndIncrement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOzoneId_GetClientId(t *testing.T) {
	type fields struct {
		callId uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name:   "client id",
			fields: fields{},
			want:   "1234567890",
		},
	}
	oid := GetOId()
	patch := gomonkey.ApplyMethod(reflect.TypeOf(oid), "GetClientId", func(_ *OzoneId) string {
		return "1234567890"
	})
	defer patch.Reset()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oid.GetClientId(); got != tt.want {
				t.Errorf("GetClientId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOzoneId_GetTraceId(t *testing.T) {
	type fields struct {
		callId uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "trace id",
			fields: fields{},
			want:   "1234567890",
		},
	}
	oid := GetOId()
	patch := gomonkey.ApplyMethod(reflect.TypeOf(oid), "GetTraceId", func(_ *OzoneId) string {
		return "1234567890"
	})
	defer patch.Reset()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oid.GetTraceId(); got != tt.want {
				t.Errorf("GetTraceId() = %v, want %v", got, tt.want)
			}
		})
	}
}
