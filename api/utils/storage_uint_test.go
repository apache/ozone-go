package utils

import "testing"

func TestBytesToHuman(t *testing.T) {
	type args struct {
		length uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "byte to human",
			args: args{length: 0},
			want: "0 B",
		},
		{
			name: "byte to human",
			args: args{length: 1},
			want: "1.00 B",
		},
		{
			name: "byte to human",
			args: args{length: 1024},
			want: "1.00 KB",
		},
		{
			name: "byte to human",
			args: args{length: 1024 * 1024},
			want: "1.00 MB",
		},
		{
			name: "byte to human",
			args: args{length: 1024 * 1024 * 1024},
			want: "1.00 GB",
		},
		{
			name: "byte to human",
			args: args{length: 1024 * 1024 * 1024 * 1024},
			want: "1.00 TB",
		},
		{
			name: "byte to human",
			args: args{length: 1024 * 1024 * 1024 * 1024 * 1024},
			want: "1.00 PB",
		},
		{
			name: "byte to human",
			args: args{length: 1024 * 1024 * 1024 * 1024 * 1024 * 1024},
			want: "1.00 EB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToHuman(tt.args.length); got != tt.want {
				t.Errorf("BytesToHuman() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorageSize_Parse(t *testing.T) {
	type fields struct {
		Uint  StorageUint
		Value uint64
	}
	type args struct {
		storage string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "parse size",
			fields: fields{
				Uint:  B,
				Value: 0,
			},
			args:    args{"0B"},
			wantErr: false,
		},
		{
			name: "parse size",
			fields: fields{
				Uint:  KB,
				Value: 1,
			},
			args:    args{"1KB"},
			wantErr: false,
		},
		{
			name: "parse size",
			fields: fields{
				Uint:  MB,
				Value: 1,
			},
			args:    args{"1MB"},
			wantErr: false,
		},
		{
			name: "parse size",
			fields: fields{
				Uint:  GB,
				Value: 1,
			},
			args:    args{"1GB"},
			wantErr: false,
		},
		{
			name: "parse size",
			fields: fields{
				Uint:  TB,
				Value: 1,
			},
			args:    args{"1TB"},
			wantErr: false,
		},
		{
			name: "parse size",
			fields: fields{
				Uint:  EB,
				Value: 1,
			},
			args:    args{"1EB"},
			wantErr: false,
		},

		{
			name: "parse size err",
			fields: fields{
				Uint:  B,
				Value: 0,
			},
			args:    args{"0B1"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StorageSize{
				Uint:  tt.fields.Uint,
				Value: tt.fields.Value,
			}
			if err := s.Parse(tt.args.storage); (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorageSize_Size(t *testing.T) {
	type fields struct {
		Uint  StorageUint
		Value uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
		{
			name: "size",
			fields: fields{
				Uint:  KB,
				Value: 1,
			},
			want: 1024,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StorageSize{
				Uint:  tt.fields.Uint,
				Value: tt.fields.Value,
			}
			if got := s.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorageSize_String(t *testing.T) {
	type fields struct {
		Uint  StorageUint
		Value uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "string",
			fields: fields{
				Uint:  KB,
				Value: 1,
			},
			want: "1KB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StorageSize{
				Uint:  tt.fields.Uint,
				Value: tt.fields.Value,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorageSize_conv(t *testing.T) {
	type fields struct {
		Uint  StorageUint
		Value uint64
	}
	type args struct {
		storage string
		uint    StorageUint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "conv",
			fields: fields{
				Uint:  KB,
				Value: 1,
			},
			args: args{
				storage: "1KB",
				uint:    KB,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StorageSize{
				Uint:  tt.fields.Uint,
				Value: tt.fields.Value,
			}
			if err := s.conv(tt.args.storage, tt.args.uint); (err != nil) != tt.wantErr {
				t.Errorf("conv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorageUint_String(t *testing.T) {
	tests := []struct {
		name string
		su   StorageUint
		want string
	}{
		// TODO: Add test cases.
		{
			name: "string",
			su:   KB,
			want: "KB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.su.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorageUint_Value(t *testing.T) {
	tests := []struct {
		name string
		su   StorageUint
		want uint64
	}{
		// TODO: Add test cases.
		{
			name: "value",
			su:   KB,
			want: 1024,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.su.Value(); got != tt.want {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorageUint_fromString(t *testing.T) {
	type args struct {
		uintString string
	}
	tests := []struct {
		name string
		st   StorageUint
		args args
		want StorageUint
	}{
		// TODO: Add test cases.
		{
			name: "from string",
			st:   KB,
			args: args{"B"},
			want: B,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.st.fromString(tt.args.uintString); got != tt.want {
				t.Errorf("fromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
