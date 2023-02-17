package nullresult

import "github.com/rocgao/roc-go-common/result"

func Err[T any](err error) NullResult[T] {
	return NullResult[T]{Result: result.Err[T](err)}
}

func Value[T any](v T) NullResult[T] {
	return NullResult[T]{Result: result.Value(v)}
}

func None[T any]() NullResult[T] {
	return NullResult[T]{none: true}
}

type NullResult[T any] struct {
	result.Result[T]
	none bool
}

func (n NullResult[T]) None() bool {
	return n.none
}
