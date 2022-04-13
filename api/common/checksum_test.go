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
    "github.com/apache/ozone-go/api/proto/datanode"
    "reflect"
    "testing"
)

func TestCRC32(t *testing.T) {
    type args struct {
        data []byte
    }
    tests := []struct {
        name string
        args args
        want []byte
    }{
        {
            name: "crc32",
            args: args{[]byte{0, 1, 2, 3}},
            want: []byte{19, 134, 185, 139, 0, 0, 0, 0},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := CRC32(tt.args.data); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("CRC32() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestCRC32c(t *testing.T) {
    type args struct {
        data []byte
    }
    tests := []struct {
        name string
        args args
        want []byte
    }{
        {
            name: "crc32c",
            args: args{[]byte{0, 1, 2, 3}},
            want: []byte{163, 26, 51, 217, 0, 0, 0, 0},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := CRC32c(tt.args.data); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("CRC32c() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestChecksumTypeFromCode(t *testing.T) {
    type args struct {
        code int32
    }
    tests := []struct {
        name string
        args args
        want datanode.ChecksumType
    }{

        {
            name: "none",
            args: args{code: 1},
            want: datanode.ChecksumType_NONE,
        },
        {
            name: "crc32",
            args: args{code: 2},
            want: datanode.ChecksumType_CRC32,
        },
        {
            name: "crc32c",
            args: args{code: 3},
            want: datanode.ChecksumType_CRC32C,
        },
        {
            name: "sha256",
            args: args{code: 4},
            want: datanode.ChecksumType_SHA256,
        },
        {
            name: "md5",
            args: args{code: 5},
            want: datanode.ChecksumType_MD5,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ChecksumTypeFromCode(tt.args.code); got != tt.want {
                t.Errorf("ChecksumTypeFromCode() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestChecksumTypeFromName(t *testing.T) {
    type args struct {
        name string
    }
    tests := []struct {
        name string
        args args
        want datanode.ChecksumType
    }{
        {
            name: "crc32",
            args: args{name: "crc32"},
            want: datanode.ChecksumType_CRC32,
        },
        {
            name: "sha256",
            args: args{name: "sha256"},
            want: datanode.ChecksumType_SHA256,
        },
        {
            name: "crc32c",
            args: args{name: "crc32c"},
            want: datanode.ChecksumType_CRC32C,
        },
        {
            name: "md5",
            args: args{name: "md5"},
            want: datanode.ChecksumType_MD5,
        },
        {
            name: "none",
            args: args{name: "none"},
            want: datanode.ChecksumType_NONE,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ChecksumTypeFromName(tt.args.name); got != tt.want {
                t.Errorf("ChecksumTypeFromName() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestMD5(t *testing.T) {
    type args struct {
        data []byte
    }
    tests := []struct {
        name string
        args args
        want []byte
    }{
        {
            name: "md5",
            args: args{[]byte{0, 1, 2, 3}},
            want: []byte{55, 181, 154, 253, 89, 39, 37, 249, 48, 94, 72, 74, 93, 127, 81, 104},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := MD5(tt.args.data); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("MD5() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSHA256(t *testing.T) {
    type args struct {
        data []byte
    }
    tests := []struct {
        name string
        args args
        want []byte
    }{
        {
            name: "sha256",
            args: args{[]byte{0, 1, 2, 3}},
            want: []byte{5, 78, 222, 193, 208, 33, 31, 98, 79, 237, 12, 188, 169, 212, 249, 64, 11, 14, 73, 28, 67, 116, 42, 242, 197, 176, 171, 235, 240, 201, 144, 216},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := SHA256(tt.args.data); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SHA256() = %v, want %v", got, tt.want)
			}
		})
	}
}
