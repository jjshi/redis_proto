package proto

import (
	"reflect"
	"testing"
)

func TestParseProto(t *testing.T) {
	type args struct {
		protocol []byte
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		//  Add test cases.
		{
			name: "test_status",
			args: args{
				protocol: []byte("+OK\r\n"),
			},
			want:    "OK",
			wantErr: false,
		},

		{
			name: "test_string",
			args: args{
				protocol: []byte("+message\r\n"),
			},
			want:    "message",
			wantErr: false,
		},

		{
			name: "test_int",
			args: args{
				protocol: []byte(":1000\r\n"),
			},
			want:    1000,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseProto(tt.args.protocol)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseProto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseProto() = %v, want %v", got, tt.want)
			}
		})
	}
}
