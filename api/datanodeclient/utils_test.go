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
package datanodeclient

import (
	"testing"
)

func TestIsReadOnly(t *testing.T) {

	type args struct {
		cmdType string
	}
	for _, readTye := range readTypeList {
		tests := []struct {
			name string
			args args
			want bool
		}{
			// TODO: Add test cases.
			{
				name: "read " + readTye,
				args: args{cmdType: readTye},
				want: true,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := IsReadOnly(tt.args.cmdType); got != tt.want {
					t.Errorf("IsReadOnly() = %v, want %v", got, tt.want)
				}
			})
		}
	}
	for _, writeTye := range writeTypeList {
		tests := []struct {
			name string
			args args
			want bool
		}{
			// TODO: Add test cases.
			{
				name: "read " + writeTye,
				args: args{cmdType: writeTye},
				want: false,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := IsReadOnly(tt.args.cmdType); got != tt.want {
					t.Errorf("IsReadOnly() = %v, want %v", got, tt.want)
				}
			})
		}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "read default",
			args: args{cmdType: "default"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsReadOnly(tt.args.cmdType); got != tt.want {
				t.Errorf("IsReadOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}
