package utils

import (
	"reflect"
	"testing"
)

func TestPointBool(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "false",
			args: args{b: false},
			want: false,
		},
		{
			name: "true",
			args: args{b: true},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PointBool(tt.args.b); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("PointBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPointInt32(t *testing.T) {
	type args struct {
		i int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
		{
			name: "pointvalue",
			args: args{0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PointInt32(tt.args.i); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("PointInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPointInt64(t *testing.T) {
	type args struct {
		i int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "pointvalue",
			args: args{0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PointInt64(tt.args.i); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("PointInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPointString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "pointvalue",
			args: args{"0"},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PointString(tt.args.s); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("PointString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPointUint32(t *testing.T) {
	type args struct {
		i uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
		{
			name: "pointvalue",
			args: args{0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PointUint32(tt.args.i); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("PointUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPointUint64(t *testing.T) {
	type args struct {
		i uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
		{
			name: "pointvalue",
			args: args{0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PointUint64(tt.args.i); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("PointUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}
