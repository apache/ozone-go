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
	"reflect"
	"testing"
)

func TestAclInfo_AddRight(t *testing.T) {
	type fields struct {
		acls []Acl
	}
	type args struct {
		acl Acl
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Acl
	}{
		// TODO: Add test cases.
		{
			name:   "add right",
			fields: fields{acls: make([]Acl, 0)},
			args:   args{READ},
			want:   []Acl{READ},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aclInfo := &AclInfo{
				Acls: tt.fields.acls,
			}
			aclInfo.AddRight(tt.args.acl)
			got := aclInfo.GetRights()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListBucket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAclInfo_Clear(t *testing.T) {
	type fields struct {
		acls []Acl
	}
	tests := []struct {
		name   string
		fields fields
		want   []Acl
	}{
		// TODO: Add test cases.
		{
			name:   "clean right",
			fields: fields{acls: make([]Acl, 0)},
			want:   []Acl{},
		},
		{
			name:   "clean right",
			fields: fields{acls: []Acl{READ, WRITE}},
			want:   []Acl{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aclInfo := &AclInfo{
				Acls: tt.fields.acls,
			}
			aclInfo.Clear()
			got := aclInfo.GetRights()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListBucket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAclInfo_FromBytes(t *testing.T) {
	type fields struct {
		acls []Acl
	}
	type args struct {
		rights []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Acl
	}{
		// TODO: Add test cases.
		{
			name:   "from bytes",
			fields: fields{acls: make([]Acl, 0)},
			args:   args{rights: []byte{1}},
			want:   []Acl{READ},
		},
		{
			name:   "from bytes",
			fields: fields{acls: make([]Acl, 0)},
			args:   args{rights: []byte{2}},
			want:   []Acl{WRITE},
		},
		{
			name:   "from bytes",
			fields: fields{acls: make([]Acl, 0)},
			args:   args{rights: []byte{128}},
			want:   []Acl{ALL},
		},
		{
			name:   "from bytes",
			fields: fields{acls: make([]Acl, 0)},
			args:   args{rights: []byte{0, 1}},
			want:   []Acl{NONE},
		},
		{
			name:   "from bytes",
			fields: fields{acls: make([]Acl, 0)},
			args:   args{rights: []byte{3}},
			want:   []Acl{READ, WRITE},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aclInfo := &AclInfo{
				Acls: tt.fields.acls,
			}
			aclInfo.FromBytes(tt.args.rights)
			got := aclInfo.GetRights()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListBucket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAclInfo_GetRights(t *testing.T) {
	type fields struct {
		acls []Acl
	}
	tests := []struct {
		name   string
		fields fields
		want   []Acl
	}{
		// TODO: Add test cases.
		{
			name:   "Get right",
			fields: fields{acls: []Acl{READ, WRITE}},
			want:   []Acl{READ, WRITE},
		},
		{
			name:   "Get right",
			fields: fields{acls: []Acl{}},
			want:   []Acl{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aclInfo := &AclInfo{
				Acls: tt.fields.acls,
			}
			if got := aclInfo.GetRights(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRights() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAclInfo_Len(t *testing.T) {
	type fields struct {
		acls []Acl
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "len",
			fields: fields{acls: []Acl{READ, WRITE}},
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aclInfo := &AclInfo{
				Acls: tt.fields.acls,
			}
			if got := aclInfo.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAclInfo_Less(t *testing.T) {
	type fields struct {
		acls []Acl
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name:   "less",
			fields: fields{acls: []Acl{READ, WRITE}},
			args: args{
				i: 0, j: 1,
			},
			want: true,
		},
		{
			name:   "great",
			fields: fields{acls: []Acl{READ, WRITE}},
			args: args{
				i: 1, j: 0,
			},
			want: false,
		},
		{
			name:   "equal",
			fields: fields{acls: []Acl{READ, WRITE}},
			args: args{
				i: 1, j: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aclInfo := &AclInfo{
				Acls: tt.fields.acls,
			}
			if got := aclInfo.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAclInfo_SetRights(t *testing.T) {
	type fields struct {
		acls []Acl
	}
	type args struct {
		acls []Acl
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Acl
	}{
		// TODO: Add test cases.
		{
			name:   "set rights",
			fields: fields{acls: make([]Acl, 0)},
			args:   args{acls: []Acl{READ}},
			want:   []Acl{READ},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aclInfo := &AclInfo{
				Acls: tt.fields.acls,
			}
			aclInfo.SetRights(tt.args.acls)
			got := aclInfo.GetRights()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListBucket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAclInfo_Swap(t *testing.T) {
	type fields struct {
		acls []Acl
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Acl
	}{
		// TODO: Add test cases.
		{
			name:   "swap",
			fields: fields{acls: []Acl{READ, WRITE}},
			args:   args{i: 0, j: 1},
			want:   []Acl{WRITE, READ},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aclInfo := &AclInfo{
				Acls: tt.fields.acls,
			}
			aclInfo.Swap(tt.args.i, tt.args.j)
			got := aclInfo.GetRights()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListBucket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAclInfo_ToBytes(t *testing.T) {
	type fields struct {
		acls []Acl
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
		{
			name:   "to bytes",
			fields: fields{acls: []Acl{READ, WRITE}},
			want:   []byte{3},
		},
		{
			name:   "to bytes",
			fields: fields{acls: []Acl{READ}},
			want:   []byte{1},
		},
		{
			name:   "to bytes",
			fields: fields{acls: []Acl{WRITE}},
			want:   []byte{2},
		},
		{
			name:   "to bytes",
			fields: fields{acls: []Acl{ALL}},
			want:   []byte{128},
		},
		{
			name:   "to bytes",
			fields: fields{acls: []Acl{NONE}},
			want:   []byte{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aclInfo := &AclInfo{
				Acls: tt.fields.acls,
			}
			if got := aclInfo.ToBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAcl_GetName(t *testing.T) {
	tests := []struct {
		name string
		acl  Acl
		want string
	}{
		// TODO: Add test cases.
		{
			name: "get name",
			acl:  1,
			want: "READ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.acl.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAcl_GetRight(t *testing.T) {
	tests := []struct {
		name string
		acl  Acl
		want uint
	}{
		// TODO: Add test cases.
		{
			name: "get right",
			acl:  READ,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.acl.GetRight(); got != tt.want {
				t.Errorf("GetRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAcl(t *testing.T) {
	type args struct {
		right uint
	}
	tests := []struct {
		name string
		args args
		want Acl
	}{
		// TODO: Add test cases.
		{
			name: "new acl",
			args: args{right: 1},
			want: READ,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAcl(tt.args.right); got != tt.want {
				t.Errorf("NewAcl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAclFromString(t *testing.T) {
	type args struct {
		right string
	}
	tests := []struct {
		name string
		args args
		want Acl
	}{
		// TODO: Add test cases.
		{
			name: "new acl from string",
			args: args{"READ"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAclFromString(tt.args.right); got != tt.want {
				t.Errorf("NewAclFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAclInfo(t *testing.T) {
	tests := []struct {
		name string
		want *AclInfo
	}{
		// TODO: Add test cases.
		{
			name: "new acl info",
			want: &AclInfo{Acls: make([]Acl, 0)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAclInfo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAclInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAclInfo_Strings(t *testing.T) {
	type fields struct {
		Acls []Acl
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
		{
			name:   "empty acl info strings",
			fields: fields{Acls: []Acl{}},
			want:   []string{},
		},
		{
			name:   "acl info strings",
			fields: fields{Acls: []Acl{ALL, READ}},
			want:   []string{"ALL", "READ"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aclInfo := &AclInfo{
				Acls: tt.fields.Acls,
			}
			if got := aclInfo.Strings(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Strings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAclInfo_String(t *testing.T) {
	type fields struct {
		Acls []Acl
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "empty acl info string",
			fields: fields{Acls: []Acl{}},
			want:   "",
		},
		{
			name:   "acl info string",
			fields: fields{Acls: []Acl{ALL}},
			want:   "ALL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aclInfo := &AclInfo{
				Acls: tt.fields.Acls,
			}
			if got := aclInfo.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
