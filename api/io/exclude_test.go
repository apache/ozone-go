package io

import (
    "github.com/apache/ozone-go/api/proto/hdds"
    "github.com/apache/ozone-go/api/utils"
    "reflect"
    "testing"
)

func TestKeyOutputStream_AddContainerIdToExcludeList(t *testing.T) {
    type fields struct {
        excludeList *hdds.ExcludeListProto
    }
    type args struct {
        id int64
    }
    tests := []struct {
        name   string
        fields fields
        args   args
    }{
        // TODO: Add test cases.
        {
            name: "AddContainerIdToExcludeList",
            fields: fields{excludeList: &hdds.ExcludeListProto{
                Datanodes:    make([]string, 0),
                ContainerIds: make([]int64, 0),
                PipelineIds:  make([]*hdds.PipelineID, 0),
            }},
            args: args{1},
        },
        {
            name: "AddContainerIdToExcludeList",
            fields: fields{excludeList: &hdds.ExcludeListProto{
                Datanodes:    make([]string, 0),
                ContainerIds: []int64{1},
                PipelineIds:  make([]*hdds.PipelineID, 0),
            }},
            args: args{1},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            k := &KeyOutputStream{
                excludeList: tt.fields.excludeList,
            }
            k.AddContainerIdToExcludeList(tt.args.id)
        })
    }
}

func TestKeyOutputStream_AddDNToExcludeList(t *testing.T) {
    type fields struct {
        excludeList *hdds.ExcludeListProto
    }
    type args struct {
        dn string
    }
    tests := []struct {
        name   string
        fields fields
        args   args
    }{
        // TODO: Add test cases.
        {
            name: "AddDNToExcludeList",
            fields: fields{excludeList: &hdds.ExcludeListProto{
                Datanodes:    make([]string, 0),
                ContainerIds: make([]int64, 0),
                PipelineIds:  make([]*hdds.PipelineID, 0),
            }},
            args: args{dn: "1"},
        },
        {
            name: "AddDNToExcludeList",
            fields: fields{excludeList: &hdds.ExcludeListProto{
                Datanodes:    []string{"1"},
                ContainerIds: make([]int64, 0),
                PipelineIds:  make([]*hdds.PipelineID, 0),
            }},
            args: args{dn: "1"},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            k := &KeyOutputStream{
                excludeList: tt.fields.excludeList,
            }
            k.AddDNToExcludeList(tt.args.dn)
        })
    }
}

func TestKeyOutputStream_AddPipeLineDNToExcludeList(t *testing.T) {
    type fields struct {
        excludeList *hdds.ExcludeListProto
    }
    type args struct {
        pipeline *hdds.Pipeline
    }
    tests := []struct {
        name   string
        fields fields
        args   args
    }{
        // TODO: Add test cases.
        {
            name: "AddPipeLineDNToExcludeList",
            fields: fields{excludeList: &hdds.ExcludeListProto{
                Datanodes:    make([]string, 0),
                ContainerIds: make([]int64, 0),
                PipelineIds:  make([]*hdds.PipelineID, 0),
            }},
            args: args{&hdds.Pipeline{Id: &hdds.PipelineID{
                Id:      utils.PointString("1"),
                Uuid128: &hdds.UUID{},
            },
                Members: make([]*hdds.DatanodeDetailsProto, 0)}},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            k := &KeyOutputStream{
                excludeList: tt.fields.excludeList,
            }
            k.AddPipeLineDNToExcludeList(tt.args.pipeline)
        })
    }
}

func TestKeyOutputStream_AddPipelineToExcludeList(t *testing.T) {
    type fields struct {
        excludeList *hdds.ExcludeListProto
    }
    type args struct {
        id *hdds.PipelineID
    }
    tests := []struct {
        name   string
        fields fields
        args   args
    }{
        // TODO: Add test cases.
        {
            name: "AddPipelineToExcludeList",
            fields: fields{excludeList: &hdds.ExcludeListProto{
                Datanodes:    make([]string, 0),
                ContainerIds: make([]int64, 0),
                PipelineIds:  make([]*hdds.PipelineID, 0),
            }},
            args: args{&hdds.PipelineID{
                Id:      utils.PointString("1"),
                Uuid128: &hdds.UUID{},
            }},
        },
        {
            name: "AddPipelineToExcludeList",
            fields: fields{excludeList: &hdds.ExcludeListProto{
                Datanodes:    []string{"1"},
                ContainerIds: make([]int64, 0),
                PipelineIds: []*hdds.PipelineID{&hdds.PipelineID{
                    Id:      utils.PointString("1"),
                    Uuid128: &hdds.UUID{},
                }},
            }},
            args: args{&hdds.PipelineID{
                Id:      utils.PointString("1"),
                Uuid128: &hdds.UUID{},
            }},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            k := &KeyOutputStream{
                excludeList: tt.fields.excludeList,
            }
            k.AddPipelineToExcludeList(tt.args.id)
        })
    }
}

func TestKeyOutputStream_GetExcludeList(t *testing.T) {
    type fields struct {
        excludeList *hdds.ExcludeListProto
    }
    tests := []struct {
        name   string
        fields fields
        want   *hdds.ExcludeListProto
    }{
        // TODO: Add test cases.
        {
            name: "",
            fields: fields{excludeList: &hdds.ExcludeListProto{
                Datanodes:    make([]string, 0),
                ContainerIds: make([]int64, 0),
                PipelineIds:  make([]*hdds.PipelineID, 0),
            }},
            want: &hdds.ExcludeListProto{
                Datanodes:    make([]string, 0),
                ContainerIds: make([]int64, 0),
                PipelineIds:  make([]*hdds.PipelineID, 0),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            k := &KeyOutputStream{
                excludeList: tt.fields.excludeList,
            }
            if got := k.GetExcludeList(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("GetExcludeList() = %v, want %v", got, tt.want)
			}
		})
	}
}
