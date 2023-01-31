package option

func Some[T any](v T) Option[T] {
	return Option[T]{value: v, none: false}
}

func None[T any]() Option[T] {
	return Option[T]{none: true}
}

type Option[T any] struct {
	value T
	none  bool
}

func (o Option[T]) None() bool {
	return o.none
}

func (o Option[T]) Some() (T, bool) {
	return o.value, o.none
}
