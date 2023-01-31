package result

func Value[T any](v T) Result[T] {
	return Result[T]{Value: v}
}

func Err[T any](err error) Result[T] {
	return Result[T]{Err: err}
}

type Result[T any] struct {
	Value T
	Err   error
}

func (r Result[T]) HasErr() bool {
	return r.Err != nil
}

func (r Result[T]) HasNotErr() (T, bool) {
	return r.Value, r.Err == nil
}
