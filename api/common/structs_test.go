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
)

func TestOzoneObjectFromFsPath(t *testing.T) {
	type args struct {
		originPath string
	}
	tests := []struct {
		name string
		args args
		want OzoneObjectAddress
	}{
		{
			name: "full path",
			args: args{originPath: "/user/slave/key"},
			want: OzoneObjectAddress{
				Om:       "",
				Volume:   "user",
				Bucket:   "slave",
				Key:      "key",
				FilePath: "/user/slave/key",
			},
		},
		{
			name: "relative path",
			args: args{originPath: "key"},
			want: OzoneObjectAddress{
				Om:       "",
				Volume:   "user",
				Bucket:   "guest",
				Key:      "key",
				FilePath: "/user/guest/key",
			},
		},
		{
			name: "sub dir path",
			args: args{originPath: "key/subdir/test"},
			want: OzoneObjectAddress{
				Om:       "",
				Volume:   "user",
				Bucket:   "guest",
				Key:      "key/subdir/test",
				FilePath: "/user/guest/key/subdir/test",
			},
		},
		{
			name: "bucket path",
			args: args{originPath: "/user/slave"},
			want: OzoneObjectAddress{
				Om:       "",
				Volume:   "user",
				Bucket:   "slave",
				Key:      "",
				FilePath: "/user/slave",
			},
		},
		{
			name: "volume path",
			args: args{originPath: "/user"},
			want: OzoneObjectAddress{
				Om:       "",
				Volume:   "user",
				Bucket:   "",
				Key:      "",
				FilePath: "/user",
			},
		},
		{
			name: "root path",
			args: args{originPath: "/"},
			want: OzoneObjectAddress{
				Om:       "",
				Volume:   "",
				Bucket:   "",
				Key:      "",
				FilePath: "/",
			},
		},
		{
			name: "relative path",
			args: args{originPath: ""},
			want: OzoneObjectAddress{
				Om:       "",
				Volume:   "user",
				Bucket:   "guest",
				Key:      "",
				FilePath: "/user/guest",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OzoneObjectAddressFromFsPath(tt.args.originPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OzoneObjectAddressFromFsPath() = %v, want %v", got.String(), tt.want.String())
			}
		})
	}
}
