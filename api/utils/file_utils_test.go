package utils

import "testing"

func TestExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test file exist",
			args: args{path: "file_utils_test.go"},
			want: true,
		},
		{
			name: "test file not exist",
			args: args{path: "file_utils.go1"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exists(tt.args.path); got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "path is file",
			args: args{path: "file_utils_test.go"},
			want: false,
		},
		{
			name: "path is dir",
			args: args{path: "../utils"},
			want: true,
		},
		{
			name: "path not eixst",
			args: args{path: "test_not_exist"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDir(tt.args.path); got != tt.want {
				t.Errorf("IsDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "path is file",
			args: args{path: "file_utils_test.go"},
			want: true,
		},
		{
			name: "path is dir",
			args: args{path: "../utils"},
			want: false,
		},
		{
			name: "path is not exist",
			args: args{path: "test_not_exist"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFile(tt.args.path); got != tt.want {
				t.Errorf("IsFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
