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
    "github.com/apache/ozone-go/api/proto/datanode"
    "github.com/apache/ozone-go/api/proto/hdds"
    "github.com/apache/ozone-go/api/utils"
    "reflect"
    "testing"
)

func TestChecksumOperator_ComputeChecksum(t *testing.T) {
    type fields struct {
        Checksum *datanode.ChecksumData
    }
    type args struct {
        data []byte
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        wantErr bool
    }{{
        name: "compute none checksum",
        fields: fields{Checksum: &datanode.ChecksumData{
            Type:             datanode.ChecksumType_NONE.Enum(),
            BytesPerChecksum: nil,
            Checksums:        nil,
        }},
        args:    args{data: []byte{0}},
        wantErr: false,
    },
        {
            name: "compute crc32 checksum",
            fields: fields{Checksum: &datanode.ChecksumData{
                Type:             datanode.ChecksumType_CRC32.Enum(),
                BytesPerChecksum: utils.PointUint32(1048576),
                Checksums:        nil,
            }},
            args:    args{data: []byte{0, 1, 2, 3}},
            wantErr: false,
        },
        {
            name: "compute crc32c checksum",
            fields: fields{Checksum: &datanode.ChecksumData{
                Type:             datanode.ChecksumType_CRC32C.Enum(),
                BytesPerChecksum: utils.PointUint32(1048576),
                Checksums:        nil,
            }},
            args:    args{data: []byte{0}},
            wantErr: false,
        },
        {
            name: "compute sha256 checksum",
            fields: fields{Checksum: &datanode.ChecksumData{
                Type:             datanode.ChecksumType_SHA256.Enum(),
                BytesPerChecksum: utils.PointUint32(1048576),
                Checksums:        nil,
            }},
            args:    args{data: []byte{0}},
            wantErr: false,
        },
        {
            name: "compute md5 checksum",
            fields: fields{Checksum: &datanode.ChecksumData{
                Type:             datanode.ChecksumType_MD5.Enum(),
                BytesPerChecksum: utils.PointUint32(1048576),
                Checksums:        nil,
            }},
            args:    args{data: []byte{0}},
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            operator := newChecksumOperator(tt.fields.Checksum)
            if err := operator.ComputeChecksum(tt.args.data); (err != nil) != tt.wantErr {
                t.Errorf("ComputeChecksum() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestChecksumOperator_VerifyChecksum(t *testing.T) {
    type fields struct {
        Checksum *datanode.ChecksumData
    }
    type args struct {
        data []byte
        off  uint64
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        wantErr bool
    }{{
        name: "verify crc32 checksum",
        fields: fields{Checksum: &datanode.ChecksumData{
            Type:             datanode.ChecksumType_CRC32.Enum(),
            BytesPerChecksum: utils.PointUint32(1048576),
            Checksums:        [][]byte{{139, 185, 134, 19}},
        }},
        args: args{
            data: []byte{0, 1, 2, 3},
            off:  0,
        },
        wantErr: false,
    },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            operator := &ChecksumOperator{
                Checksum: tt.fields.Checksum,
            }
            if err := operator.VerifyChecksum(tt.args.data, tt.args.off); (err != nil) != tt.wantErr {
                t.Errorf("VerifyChecksum() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestDatanodeDetailsOperator_Host(t *testing.T) {
    type fields struct {
        DatanodeDetails *hdds.DatanodeDetailsProto
    }
    tests := []struct {
        name   string
        fields fields
        want   string
    }{{
        name: "host",
        fields: fields{DatanodeDetails: &hdds.DatanodeDetailsProto{
            HostName: utils.PointString("localhost"),
        }},
        want: "localhost",
    },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            dn := &DatanodeDetailsOperator{
                DatanodeDetails: tt.fields.DatanodeDetails,
            }
            if got := dn.Host(); got != tt.want {
                t.Errorf("Host() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestDatanodeDetailsOperator_ID(t *testing.T) {
    type fields struct {
        DatanodeDetails *hdds.DatanodeDetailsProto
    }
    tests := []struct {
        name   string
        fields fields
        want   string
    }{{
        name: "id",
        fields: fields{DatanodeDetails: &hdds.DatanodeDetailsProto{
            Uuid: utils.PointString("id"),
        }},
        want: "id",
    },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            dn := &DatanodeDetailsOperator{
                DatanodeDetails: tt.fields.DatanodeDetails,
            }
            if got := dn.ID(); got != tt.want {
                t.Errorf("ID() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestDatanodeDetailsOperator_IP(t *testing.T) {
    type fields struct {
        DatanodeDetails *hdds.DatanodeDetailsProto
    }
    tests := []struct {
        name   string
        fields fields
        want   string
    }{{
        name: "ip",
        fields: fields{DatanodeDetails: &hdds.DatanodeDetailsProto{
            IpAddress: utils.PointString("ip"),
        }},
        want: "ip",
    },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            dn := &DatanodeDetailsOperator{
                DatanodeDetails: tt.fields.DatanodeDetails,
            }
            if got := dn.IP(); got != tt.want {
                t.Errorf("IP() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestDatanodeDetailsOperator_Port(t *testing.T) {
    type fields struct {
        DatanodeDetails *hdds.DatanodeDetailsProto
    }
    type args struct {
        t PipelinePortName
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   uint32
    }{{
        name: "ratis port",
        fields: fields{DatanodeDetails: &hdds.DatanodeDetailsProto{
            Ports: []*hdds.Port{{Name: utils.PointString("RATIS"), Value: utils.PointUint32(9070)}},
        }},
        args: args{t: PipelinePortNameRATIS},
        want: 9070,
    },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            dn := &DatanodeDetailsOperator{
                DatanodeDetails: tt.fields.DatanodeDetails,
            }
            if got := dn.Port(tt.args.t); got != tt.want {
                t.Errorf("Port() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestExcludeListOperator_AddContainerId(t *testing.T) {
    type fields struct {
        datanodes    []string
        containerIds []int64
        pipelineIds  []string
    }
    type args struct {
        containerId int64
    }
    tests := []struct {
        name   string
        fields fields
        args   args
    }{{
        name: "add container id",
        fields: fields{
            datanodes:    []string{},
            containerIds: []int64{},
            pipelineIds:  []string{},
        },
        args: args{containerId: 1},
    },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            operator := &ExcludeListOperator{
                datanodes:    tt.fields.datanodes,
                containerIds: tt.fields.containerIds,
                pipelineIds:  tt.fields.pipelineIds,
            }
            operator.AddContainerId(tt.args.containerId)
        })
    }
}

func TestExcludeListOperator_AddDataNode(t *testing.T) {
    type fields struct {
        datanodes    []string
        containerIds []int64
        pipelineIds  []string
    }
    type args struct {
        datanode string
    }
    tests := []struct {
        name   string
        fields fields
        args   args
    }{{
        name: "add datanode",
        fields: fields{
            datanodes:    []string{},
            containerIds: []int64{},
            pipelineIds:  []string{},
        },
        args: args{datanode: "1"},
    },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            operator := &ExcludeListOperator{
                datanodes:    tt.fields.datanodes,
                containerIds: tt.fields.containerIds,
                pipelineIds:  tt.fields.pipelineIds,
            }
            operator.AddDataNode(tt.args.datanode)
        })
    }
}

func TestExcludeListOperator_AddDataNodes(t *testing.T) {
    type fields struct {
        datanodes    []string
        containerIds []int64
        pipelineIds  []string
    }
    type args struct {
        datanode []string
    }
    tests := []struct {
        name   string
        fields fields
        args   args
    }{{
        name: "add datanodes",
        fields: fields{
            datanodes:    []string{},
            containerIds: []int64{},
            pipelineIds:  []string{},
        },
        args: args{datanode: []string{"1", "2"}},
    },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            operator := &ExcludeListOperator{
                datanodes:    tt.fields.datanodes,
                containerIds: tt.fields.containerIds,
                pipelineIds:  tt.fields.pipelineIds,
            }
            operator.AddDataNodes(tt.args.datanode)
        })
    }
}

func TestExcludeListOperator_AddPipelineId(t *testing.T) {
    type fields struct {
        datanodes    []string
        containerIds []int64
        pipelineIds  []string
    }
    type args struct {
        pipelineId string
    }
    tests := []struct {
        name   string
        fields fields
        args   args
    }{{
        name: "add pipeline id",
        fields: fields{
            datanodes:    []string{},
            containerIds: []int64{},
            pipelineIds:  []string{},
        },
        args: args{pipelineId: "1"},
    },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            operator := &ExcludeListOperator{
                datanodes:    tt.fields.datanodes,
                containerIds: tt.fields.containerIds,
                pipelineIds:  tt.fields.pipelineIds,
            }
            operator.AddPipelineId(tt.args.pipelineId)
        })
    }
}

func TestExcludeListOperator_ToExcludeList(t *testing.T) {
    type fields struct {
        datanodes    []string
        containerIds []int64
        pipelineIds  []string
    }
    tests := []struct {
        name   string
        fields fields
        want   *hdds.ExcludeListProto
    }{{
        name: "to exclude list",
        fields: fields{
            datanodes:    []string{},
            containerIds: []int64{},
            pipelineIds:  []string{},
        },
        want: &hdds.ExcludeListProto{
            Datanodes:    []string{},
            ContainerIds: []int64{},
            PipelineIds:  []*hdds.PipelineID{},
        },
    },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            operator := &ExcludeListOperator{
                datanodes:    tt.fields.datanodes,
                containerIds: tt.fields.containerIds,
                pipelineIds:  tt.fields.pipelineIds,
            }
            if got := operator.ToExcludeList(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ToExcludeList() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestNewChecksumOperatorComputer(t *testing.T) {
    type args struct {
        checksumType     *datanode.ChecksumType
        bytesPerChecksum uint32
    }
    tests := []struct {
        name string
        args args
        want *ChecksumOperator
    }{{
        name: "checksum op",
        args: args{
            checksumType:     datanode.ChecksumType_CRC32.Enum(),
            bytesPerChecksum: 1048576,
        },
        want: &ChecksumOperator{Checksum: &datanode.ChecksumData{
            Type:             datanode.ChecksumType_CRC32.Enum(),
            BytesPerChecksum: utils.PointUint32(1048576),
        }},
    },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewChecksumOperatorComputer(tt.args.checksumType, tt.args.bytesPerChecksum); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewChecksumOperatorComputer() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestNewChecksumOperatorVerifier(t *testing.T) {
    type args struct {
        checksum *datanode.ChecksumData
    }
    tests := []struct {
        name string
        args args
        want *ChecksumOperator
    }{{
        name: "checksum verify",
        args: args{checksum: &datanode.ChecksumData{
            Type:             datanode.ChecksumType_CRC32.Enum(),
            BytesPerChecksum: utils.PointUint32(1048576),
            Checksums:        nil,
        }},
        want: &ChecksumOperator{Checksum: &datanode.ChecksumData{
            Type:             datanode.ChecksumType_CRC32.Enum(),
            BytesPerChecksum: utils.PointUint32(1048576),
        }},
    },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewChecksumOperatorVerifier(tt.args.checksum); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewChecksumOperatorVerifier() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestNewDatanodeDetailsOperator(t *testing.T) {
    type args struct {
        dn *hdds.DatanodeDetailsProto
    }
    tests := []struct {
        name string
        args args
        want *DatanodeDetailsOperator
    }{{
        name: "datanode op",
        args: args{dn: &hdds.DatanodeDetailsProto{}},
        want: &DatanodeDetailsOperator{
            DatanodeDetails: &hdds.DatanodeDetailsProto{},
        },
    },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewDatanodeDetailsOperator(tt.args.dn); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewDatanodeDetailsOperator() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestNewExcludeListOperator(t *testing.T) {
    tests := []struct {
        name string
        want *ExcludeListOperator
    }{{
        name: "exclude op",
        want: &ExcludeListOperator{datanodes: []string{}, containerIds: []int64{}, pipelineIds: []string{}},
    },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewExcludeListOperator(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewExcludeListOperator() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestNewPipelineOperator(t *testing.T) {
    op, _ := NewPipelineOperator(NewPipeline())
    type args struct {
        pipeline *hdds.Pipeline
    }
    tests := []struct {
        name    string
        args    args
        want    *PipelineOperator
        wantErr bool
    }{
        {
            name:    "pipeline op",
            args:    args{pipeline: NewPipeline()},
            want:    op,
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := NewPipelineOperator(tt.args.pipeline)
            if (err != nil) != tt.wantErr {
                t.Errorf("NewPipelineOperator() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewPipelineOperator() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestPipelineOperator_GetClosestNode(t *testing.T) {
    type fields struct {
        CurrentIndex int
        Pipeline     *hdds.Pipeline
    }
    tests := []struct {
        name   string
        fields fields
        want   *hdds.DatanodeDetailsProto
    }{
        {
            name:   "closet node",
            fields: fields{CurrentIndex: 0, Pipeline: NewPipeline()},
            want: &hdds.DatanodeDetailsProto{
                Uuid:      utils.PointString("LeaderID"),
                IpAddress: utils.PointString("0.0.0.0"),
                HostName:  utils.PointString("localhost"),
                Ports:     []*hdds.Port{{Name: utils.PointString("ratis"), Value: utils.PointUint32(9070)}},
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            operator := &PipelineOperator{
                CurrentIndex: tt.fields.CurrentIndex,
                Pipeline:     tt.fields.Pipeline,
            }
            if got := operator.GetClosestNode(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("GetClosestNode() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestPipelineOperator_GetCurrentNode(t *testing.T) {
    type fields struct {
        CurrentIndex int
        Pipeline     *hdds.Pipeline
    }
    tests := []struct {
        name   string
        fields fields
        want   *hdds.DatanodeDetailsProto
    }{
        {
            name:   "current node",
            fields: fields{CurrentIndex: 0, Pipeline: NewPipeline()},
            want: &hdds.DatanodeDetailsProto{
                Uuid:      utils.PointString("LeaderID"),
                IpAddress: utils.PointString("0.0.0.0"),
                HostName:  utils.PointString("localhost"),
                Ports:     []*hdds.Port{{Name: utils.PointString("ratis"), Value: utils.PointUint32(9070)}},
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            operator := &PipelineOperator{
                CurrentIndex: tt.fields.CurrentIndex,
                Pipeline:     tt.fields.Pipeline,
            }
            if got := operator.GetCurrentNode(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("GetCurrentNode() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestPipelineOperator_GetFirstNode(t *testing.T) {
    type fields struct {
        CurrentIndex int
        Pipeline     *hdds.Pipeline
    }
    tests := []struct {
        name   string
        fields fields
        want   *hdds.DatanodeDetailsProto
    }{
        {
            name:   "first node",
            fields: fields{CurrentIndex: 0, Pipeline: NewPipeline()},
            want: &hdds.DatanodeDetailsProto{
                Uuid:      utils.PointString("LeaderID"),
                IpAddress: utils.PointString("0.0.0.0"),
                HostName:  utils.PointString("localhost"),
                Ports:     []*hdds.Port{{Name: utils.PointString("ratis"), Value: utils.PointUint32(9070)}},
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            operator := &PipelineOperator{
                CurrentIndex: tt.fields.CurrentIndex,
                Pipeline:     tt.fields.Pipeline,
            }
            if got := operator.GetFirstNode(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("GetFirstNode() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestPipelineOperator_GetId(t *testing.T) {
    type fields struct {
        CurrentIndex int
        Pipeline     *hdds.Pipeline
    }
    tests := []struct {
        name   string
        fields fields
        want   string
    }{
        {
            name:   "node id",
            fields: fields{CurrentIndex: 0, Pipeline: NewPipeline()},
            want:   "id",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            operator := &PipelineOperator{
                CurrentIndex: tt.fields.CurrentIndex,
                Pipeline:     tt.fields.Pipeline,
            }
            if got := operator.GetId(); got != tt.want {
                t.Errorf("GetId() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestPipelineOperator_GetLeaderAddress(t *testing.T) {
    type fields struct {
        CurrentIndex int
        Pipeline     *hdds.Pipeline
    }
    tests := []struct {
        name   string
        fields fields
        want   string
    }{
        {
            name: "get leader address",
            fields: fields{
                CurrentIndex: 0,
                Pipeline:     NewPipeline(),
            },
            want: "0.0.0.0",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            operator := &PipelineOperator{
                CurrentIndex: tt.fields.CurrentIndex,
                Pipeline:     tt.fields.Pipeline,
            }
            if got := operator.GetLeaderAddress(); got != tt.want {
                t.Errorf("GetLeaderAddress() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestPipelineOperator_GetLeaderNode(t *testing.T) {
    type fields struct {
        CurrentIndex int
        Pipeline     *hdds.Pipeline
    }
    tests := []struct {
        name   string
        fields fields
        want   *hdds.DatanodeDetailsProto
    }{
        {
            name: "get leader node",
            fields: fields{
                CurrentIndex: 0,
                Pipeline:     NewPipeline(),
            },
            want: &hdds.DatanodeDetailsProto{
                Uuid:      utils.PointString("LeaderID"),
                IpAddress: utils.PointString("0.0.0.0"),
                HostName:  utils.PointString("localhost"),
                Ports:     []*hdds.Port{{Name: utils.PointString("ratis"), Value: utils.PointUint32(9070)}},
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            operator := &PipelineOperator{
                CurrentIndex: tt.fields.CurrentIndex,
                Pipeline:     tt.fields.Pipeline,
            }
            if got := operator.GetLeaderNode(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("GetLeaderNode() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestPipelineOperator_GetNodeAddr(t *testing.T) {
    type fields struct {
        CurrentIndex int
        Pipeline     *hdds.Pipeline
    }
    type args struct {
        portName  PipelinePortName
        firstNode bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   string
    }{{
        name: "get node addr",
        fields: fields{
            CurrentIndex: 0,
            Pipeline:     NewPipeline(),
        },
        args: args{
            portName:  "ratis",
            firstNode: false,
        },
        want: "0.0.0.0:9070",
    }}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            operator := &PipelineOperator{
                CurrentIndex: tt.fields.CurrentIndex,
                Pipeline:     tt.fields.Pipeline,
            }
            if got := operator.GetNodeAddr(tt.args.portName, tt.args.firstNode); got != tt.want {
                t.Errorf("GetNodeAddr() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestPipelineOperator_GetPipeline(t *testing.T) {
    type fields struct {
        CurrentIndex int
        Pipeline     *hdds.Pipeline
    }
    tests := []struct {
        name   string
        fields fields
        want   *hdds.Pipeline
    }{
        {
            name:   "get pipeline",
            fields: fields{CurrentIndex: 0, Pipeline: NewPipeline()},
            want:   NewPipeline(),
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            operator := &PipelineOperator{
                CurrentIndex: tt.fields.CurrentIndex,
                Pipeline:     tt.fields.Pipeline,
            }
            if got := operator.GetPipeline(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("GetPipeline() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestPipelineOperator_SetCurrentToNext(t *testing.T) {
    type fields struct {
        CurrentIndex int
        Pipeline     *hdds.Pipeline
    }
    tests := []struct {
        name   string
        fields fields
    }{
        {
            name: "set to next",
            fields: fields{
                CurrentIndex: 0,
                Pipeline:     nil,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            operator := &PipelineOperator{
                CurrentIndex: tt.fields.CurrentIndex,
                Pipeline:     tt.fields.Pipeline,
            }
            operator.SetCurrentToNext()
        })
    }
}

func TestPipelineOperator_getRemotePort(t *testing.T) {
    type fields struct {
        CurrentIndex int
        Pipeline     *hdds.Pipeline
    }
    type args struct {
        mode string
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   uint32
    }{
        {
            name: "get remote port",
            fields: fields{
                CurrentIndex: 0,
                Pipeline: &hdds.Pipeline{
                    Members: []*hdds.DatanodeDetailsProto{{
                        Ports: []*hdds.Port{{
                            Name:  utils.PointString("ratis"),
                            Value: utils.PointUint32(9070),
                        }},
                    }},
                    MemberOrders: []uint32{0},
                    Id:           &hdds.PipelineID{Id: utils.PointString("id")},
                },
            },
            args: args{mode: "ratis"},
            want: 9070,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            operator := &PipelineOperator{
                CurrentIndex: tt.fields.CurrentIndex,
                Pipeline:     tt.fields.Pipeline,
            }
            if got := operator.getRemotePort(tt.args.mode); got != tt.want {
                t.Errorf("getRemotePort() = %v, want %v", got, tt.want)
            }
        })
    }
}

func Test_newChecksumOperator(t *testing.T) {
    type args struct {
        checksum *datanode.ChecksumData
    }
    tests := []struct {
        name string
        args args
        want *ChecksumOperator
    }{{
        name: "new checksum op",
        args: args{checksum: &datanode.ChecksumData{
            Type:             nil,
            BytesPerChecksum: nil,
            Checksums:        nil,
        }},
        want: &ChecksumOperator{&datanode.ChecksumData{
            Type:             nil,
            BytesPerChecksum: nil,
            Checksums:        nil,
        }},
    },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := newChecksumOperator(tt.args.checksum); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("newChecksumOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}
