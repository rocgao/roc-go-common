package result

import (
	"errors"
	"reflect"
	"testing"
)

func TestResult_HasErr(t *testing.T) {
	type testCase[T any] struct {
		name string
		r    Result[T]
		want bool
	}
	tests := []testCase[string]{
		{
			name: "has value",
			r:    Value("data"),
			want: false,
		},
		{
			name: "has not value",
			r:    Err[string](errors.New("an error")),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.HasErr(); got != tt.want {
				t.Errorf("HasErr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_HasNotErr(t *testing.T) {
	type testCase[T any] struct {
		name  string
		r     Result[T]
		want  T
		want1 bool
	}
	tests := []testCase[string]{
		{
			name:  "has value",
			r:     Value("data"),
			want:  "data",
			want1: true,
		},
		{
			name:  "has value",
			r:     Err[string](errors.New("an error")),
			want:  "",
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.r.HasNotErr()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasNotErr() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("HasNotErr() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
