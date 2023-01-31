package result

import (
	"errors"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	type args[T1 any, T2 any] struct {
		r       Result[T1]
		mapFunc func(T1) T2
		errFunc func(error) error
	}
	type testCase[T1 any, T2 any] struct {
		name string
		args args[T1, T2]
		want Result[T2]
	}
	tests := []testCase[string, int]{
		{
			name: "has error",
			args: args[string, int]{
				r:       Err[string](errors.New("an error")),
				mapFunc: nil,
				errFunc: nil,
			},
			want: Err[int](errors.New("an error")),
		},
		{
			name: "has error",
			args: args[string, int]{
				r:       Err[string](errors.New("an error")),
				mapFunc: nil,
				errFunc: func(e error) error {
					return errors.New("an new error")
				},
			},
			want: Err[int](errors.New("an new error")),
		}, {
			name: "has not error",
			args: args[string, int]{
				r: Value("data"),
				mapFunc: func(t1 string) int {
					return len(t1)
				},
				errFunc: nil,
			},
			want: Value(4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.r, tt.args.mapFunc, tt.args.errFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
