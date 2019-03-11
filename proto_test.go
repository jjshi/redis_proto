package proto

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestReader_Parse(t *testing.T) {
	type fields struct {
		rd *bufio.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "test_status",
			fields: fields{
				rd: bufio.NewReader(strings.NewReader("+OK\r\n")),
			},
			want:    "OK",
			wantErr: false,
		},

		{
			name: "test_string",
			fields: fields{
				rd: bufio.NewReader(strings.NewReader("$message\r\n")),
			},
			want:    "message",
			wantErr: false,
		},

		{
			name: "test_int",
			fields: fields{
				rd: bufio.NewReader(strings.NewReader(":1000\r\n")),
			},
			want:    1000,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				rd: tt.fields.rd,
			}
			got, err := r.Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reader.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
