package option

func Map[T1 any, T2 any](o Option[T1], mapFunc func(t1 T1) T2) Option[T2] {
	if o.None() {
		return None[T2]()
	}

	v2 := mapFunc(o.value)
	return Some(v2)
}
