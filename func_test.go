package structmap

import (
	"reflect"
	"testing"
)

func TestStructToMap(t *testing.T) {
	type args struct {
		str any
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "simple convert",
			args: args{
				str: struct {
					St     string
					In16   int16
					In32   int32
					In64   int64
					Uint16 uint16
					Uint32 uint32
					Uint64 uint64
					Fl32   float32
					Fl64   float64
					Bo     bool
				}{
					St:     "string_content",
					In16:   16,
					In32:   32,
					In64:   64,
					Uint16: 16,
					Uint32: 32,
					Uint64: 64,
					Fl32:   32.32,
					Fl64:   64.64,
					Bo:     true,
				},
			},
			want: map[string]string{
				"bo":     "true",
				"fl32":   "32.32",
				"fl64":   "64.64",
				"in16":   "16",
				"in32":   "32",
				"in64":   "64",
				"st":     "string_content",
				"uint16": "16",
				"uint32": "32",
				"uint64": "64",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructToMap(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapToStruct(t *testing.T) {
	type args struct {
		mmap map[string]string
		s    any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "simple convert",
			args: args{
				mmap: map[string]string{"bo": "true", "fl": "12.12", "in": "12", "st": "string_content"},
				s: struct {
					St string
					In int16
					Fl float32
					Bo bool
				}{},
			},
			want: struct {
				St string
				In int16
				Fl float32
				Bo bool
			}{
				St: "string_content",
				In: 12,
				Fl: 12.12,
				Bo: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapToStruct(tt.args.mmap, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapToStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
