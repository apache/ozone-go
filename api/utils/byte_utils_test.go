package utils

import (
	"reflect"
	"testing"
)

func TestUintToBitSet(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "uint to bitset",
			args: args{n: 4},
			want: []byte{0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UintToBitSet(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UintToBitSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16toBytes(t *testing.T) {
	type args struct {
		b []byte
		v uint16
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "TestUint16toBytes",
			args: args{
				b: make([]byte, 2),
				v: 1,
			},
			want: []byte{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Uint16toBytes(tt.args.b, tt.args.v)
			if got := tt.args.b; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UintToBitSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32toBytes(t *testing.T) {
	type args struct {
		b []byte
		v uint32
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "TestUint32toBytes",
			args: args{
				b: make([]byte, 4),
				v: 1,
			},
			want: []byte{0, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Uint32toBytes(tt.args.b, tt.args.v)
			if got := tt.args.b; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UintToBitSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64toBytes(t *testing.T) {
	type args struct {
		b []byte
		v uint64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "TestUint64toBytes",
			args: args{
				b: make([]byte, 8),
				v: 1,
			},
			want: []byte{1, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Uint64toBytes(tt.args.b, tt.args.v)
			if got := tt.args.b; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UintToBitSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8toBytes(t *testing.T) {
	type args struct {
		b []byte
		v uint8
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "TestUint8toBytes",
			args: args{
				b: make([]byte, 1),
				v: 1,
			},
			want: []byte{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Uint8toBytes(tt.args.b, tt.args.v)
			if got := tt.args.b; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UintToBitSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
