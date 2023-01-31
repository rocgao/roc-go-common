package result

func Map[T1 any, T2 any](r Result[T1], mapFunc func(T1) T2, errFunc func(error) error) Result[T2] {
	if v1, ok := r.HasNotErr(); ok {
		v2 := mapFunc(v1)
		return Value(v2)
	}

	err := r.Err
	if errFunc != nil {
		err = errFunc(err)
	}
	return Err[T2](err)
}
