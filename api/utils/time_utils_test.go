package utils

import "testing"

func TestMicrosecondFormatToString(t *testing.T) {
	type args struct {
		micro  int64
		format string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "micro",
			args: args{
				micro:  10000000,
				format: TimeFormatterSECOND,
			},
			want: "1970-01-01 08:00:10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MicrosecondFormatToString(tt.args.micro, tt.args.format); got != tt.want {
				t.Errorf("MicrosecondFormatToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMillisecondFormatToString(t *testing.T) {
	type args struct {
		milli  int64
		format string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "milli",
			args: args{
				milli:  10000,
				format: TimeFormatterSECOND,
			},
			want: "1970-01-01 08:00:10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MillisecondFormatToString(tt.args.milli, tt.args.format); got != tt.want {
				t.Errorf("MillisecondFormatToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNanosecondFormatToString(t *testing.T) {
	type args struct {
		nano   int64
		format string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "nano",
			args: args{
				nano:   10000000000,
				format: TimeFormatterSECOND,
			},
			want: "1970-01-01 08:00:10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NanosecondFormatToString(tt.args.nano, tt.args.format); got != tt.want {
				t.Errorf("NanosecondFormatToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecondFormatToString(t *testing.T) {
	type args struct {
		second int64
		format string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "second",
			args: args{
				second: 10,
				format: TimeFormatterSECOND,
			},
			want: "1970-01-01 08:00:10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SecondFormatToString(tt.args.second, tt.args.format); got != tt.want {
				t.Errorf("SecondFormatToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
