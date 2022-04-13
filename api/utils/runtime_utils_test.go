package utils

import (
	"fmt"
	"testing"
)

func TestRunTimesUntilSuccessfully(t *testing.T) {
	type args struct {
		f     func() error
		times int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "func run succ",
			args: args{f: func() error {
				return nil
			}, times: 3},
			wantErr: false,
		},
		{
			name: "func run error",
			args: args{f: func() error {
				return fmt.Errorf("err")
			}, times: 3},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RunTimesUntilSuccess(tt.args.f, tt.args.times); (err != nil) != tt.wantErr {
				t.Errorf("RunTimesUntilSuccess() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
