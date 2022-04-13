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
	"reflect"
	"testing"
)

func Test_limitBuffer_Bytes(t *testing.T) {
	type fields struct {
		buffer *bytes.Buffer
		limit  uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
		{
			name:   "limit buffer bytes",
			fields: fields{buffer: bytes.NewBuffer([]byte{0}), limit: 10},
			want:   []byte{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &limitBuffer{
				buffer: tt.fields.buffer,
				limit:  tt.fields.limit,
			}
			if got := lb.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_limitBuffer_HasRemaining(t *testing.T) {
	type fields struct {
		buffer *bytes.Buffer
		limit  uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "remaining",
			fields: fields{
				buffer: bytes.NewBuffer([]byte{0}),
				limit:  10,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &limitBuffer{
				buffer: tt.fields.buffer,
				limit:  tt.fields.limit,
			}
			if got := lb.HasRemaining(); got != tt.want {
				t.Errorf("HasRemaining() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_limitBuffer_Len(t *testing.T) {
	type fields struct {
		buffer *bytes.Buffer
		limit  uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
		{
			name: "len",
			fields: fields{
				buffer: bytes.NewBuffer([]byte{0}),
				limit:  10,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &limitBuffer{
				buffer: tt.fields.buffer,
				limit:  tt.fields.limit,
			}
			if got := lb.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_limitBuffer_Limit(t *testing.T) {
	type fields struct {
		buffer *bytes.Buffer
		limit  uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
		{
			name: "limit",
			fields: fields{
				buffer: bytes.NewBuffer([]byte{0}),
				limit:  10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &limitBuffer{
				buffer: tt.fields.buffer,
				limit:  tt.fields.limit,
			}
			if got := lb.Limit(); got != tt.want {
				t.Errorf("Limit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_limitBuffer_Next(t *testing.T) {
	type fields struct {
		buffer *bytes.Buffer
		limit  uint64
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
		{
			name: "next",
			fields: fields{
				buffer: bytes.NewBuffer([]byte{0, 1}),
				limit:  10,
			},
			args: args{n: 2},
			want: []byte{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &limitBuffer{
				buffer: tt.fields.buffer,
				limit:  tt.fields.limit,
			}
			if got := lb.Next(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_limitBuffer_Read(t *testing.T) {
	type fields struct {
		buffer *bytes.Buffer
		limit  uint64
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "read",
			fields: fields{
				buffer: bytes.NewBuffer([]byte{0, 1}),
				limit:  10,
			},
			args:    args{make([]byte, 2)},
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &limitBuffer{
				buffer: tt.fields.buffer,
				limit:  tt.fields.limit,
			}
			got, err := lb.Read(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_limitBuffer_Remaining(t *testing.T) {
	type fields struct {
		buffer *bytes.Buffer
		limit  uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
		{
			name: "reamining",
			fields: fields{
				buffer: bytes.NewBuffer([]byte{0, 1}),
				limit:  10,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &limitBuffer{
				buffer: tt.fields.buffer,
				limit:  tt.fields.limit,
			}
			if got := lb.Remaining(); got != tt.want {
				t.Errorf("Remaining() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_limitBuffer_Reset(t *testing.T) {
	type fields struct {
		buffer *bytes.Buffer
		limit  uint64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{
			name: "reset",
			fields: fields{
				buffer: bytes.NewBuffer([]byte{0, 1}),
				limit:  10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &limitBuffer{
				buffer: tt.fields.buffer,
				limit:  tt.fields.limit,
			}
			lb.Reset()
		})
	}
}

func Test_limitBuffer_Truncate(t *testing.T) {
	type fields struct {
		buffer *bytes.Buffer
		limit  uint64
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "truncate",
			fields: fields{
				buffer: bytes.NewBuffer([]byte{0, 1}),
				limit:  10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &limitBuffer{
				buffer: tt.fields.buffer,
				limit:  tt.fields.limit,
			}
			lb.Truncate(tt.args.n)
		})
	}
}

func Test_limitBuffer_Write(t *testing.T) {
	type fields struct {
		buffer *bytes.Buffer
		limit  uint64
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "write",
			fields: fields{
				buffer: bytes.NewBuffer(make([]byte, 2)),
				limit:  10,
			},
			args:    args{p: []byte{0, 1}},
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &limitBuffer{
				buffer: tt.fields.buffer,
				limit:  tt.fields.limit,
			}
			got, err := lb.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Write() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newLimitBuffer(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 4096, 4096))
	buffer.Reset()
	type args struct {
		limit uint64
	}
	tests := []struct {
		name string
		args args
		want *limitBuffer
	}{
		// TODO: Add test cases.
		{
			name: "new limit buffer",
			args: args{limit: 4096},
			want: &limitBuffer{
				buffer: buffer,
				limit:  4096,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newLimitBuffer(tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newLimitBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}
